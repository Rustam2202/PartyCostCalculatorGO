package service

import (
	"context"
	"fmt"
	"party-calc/internal/database"
	"party-calc/internal/domain"
	"party-calc/internal/repository"
	"testing"
	"time"

	"github.com/pashagolub/pgxmock/v2"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

type expect struct {
	total   decimal.Decimal
	count   int
	average decimal.Decimal
	owes    debtors
}

type testCase struct {
	testName      string
	persons       []domain.Person
	events        []domain.Event
	personsEvents []domain.PersonsAndEvents
	RoundRate     int32
	exp           expect
}

type calculationTestCases struct {
	ctx      context.Context
	mock     pgxmock.PgxConnIface
	serv     *CalcService
	testCase *testCase
}

func (c *calculationTestCases) createMockAndService(t *testing.T) {
	c.ctx = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	c.mock = mock

	repoPer := repository.NewPersonRepository(&database.DataBase{DBPGX: mock})
	repoEv := repository.NewEventRepository(&database.DataBase{DBPGX: mock})
	repoPerEv := repository.NewPersonsEventsRepository(&database.DataBase{DBPGX: mock})
	c.serv = NewCalcService(NewPersonService(repoPer), NewEventService(repoEv), NewPersonsEventsService(repoPerEv))
}

func (c *calculationTestCases) fillMock() {
	// personsEvents mock
	var personsRows [][]any
	for _, pe := range c.testCase.personsEvents {
		personsRows = append(personsRows, []any{pe.Id, pe.PersonId, pe.EventId, pe.Spent, pe.Factor})
	}
	c.mock.ExpectQuery("SELECT (.+) FROM persons_events").
		WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "PersonId", "EventId", "Spent", "Factor"}).
			AddRows(personsRows...))

	// persons + events mock
	for _, p := range c.testCase.persons {
		c.mock.ExpectQuery("SELECT (.+) FROM persons").WithArgs(p.Id).
			WillReturnRows(pgxmock.NewRows([]string{"Id", "Name"}).
				AddRow(p.Id, p.Name))
		for _, e := range p.Events {
			c.mock.ExpectQuery("SELECT id, name, date FROM events").
				WithArgs(e.Id).
				WillReturnRows(pgxmock.NewRows([]string{"Id", "Name", "Date"}).
					AddRow(e.Id, e.Name, e.Date))
		}
	}
}

func (c *calculationTestCases) assertCheck(t *testing.T, result *EventData, err error) {
	assert.NoError(t, err,
		fmt.Sprintf("Some error in %s", c.testCase.testName))
	assert.NoError(t, c.mock.ExpectationsWereMet(),
		fmt.Sprintf("Some error with mock in %s", c.testCase.testName))
	assert.EqualValues(t, c.testCase.exp.count, result.AllPeronsCount,
		fmt.Sprintf("Counts not equal in %s", c.testCase.testName))
	assert.True(t, c.testCase.exp.total.Equal(result.TotalSpent),
		fmt.Sprintf("Totals not equal in \"%s\"", c.testCase.testName))
	assert.True(t, c.testCase.exp.average.Equal(result.AverageSpent.Round(c.testCase.RoundRate)),
		fmt.Sprintf("Averages not equal in \"%s\": %s != %s",
			c.testCase.testName, c.testCase.exp.average.String(), result.AverageSpent.String()))

	for k, v := range c.testCase.exp.owes {
		for k1, v1 := range v {
			//	fmt.Println(result.Owes[k][k1].Round(2))
			assert.True(t, result.Owes[k][k1].Round(c.testCase.RoundRate).Equal(v1),
				fmt.Sprintf("Owe not equal in \"%s\": %s != %s",
					c.testCase.testName, v1.String(), result.Owes[k][k1].String()))
		}
	}
}

func createTestCases() []func() *testCase {
	return []func() *testCase{
		testCase1,
		testCase2,
		testCase3,
		testCase4,
		testCase5,
	}
}

func TestRun(t *testing.T) {
	var c calculationTestCases
	c.createMockAndService(t)
	cases := createTestCases()
	for _, tc := range cases {
		c.testCase = tc()
		c.fillMock()
		result, err := c.serv.CalculateEvent(c.ctx, 1, c.testCase.RoundRate)
		//	fmt.Println(result.Owes, c.testCase.exp.owes)
		c.assertCheck(t, result, err)
	}
}

func testCase1() *testCase {
	var tc testCase
	tc.testName = "#1. One debetor to two recepients"
	tc.events = []domain.Event{
		{
			Id:      1,
			Name:    "New Year",
			Date:    time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local),
			Persons: tc.persons},
	}
	tc.persons = []domain.Person{
		{Id: 4, Name: "Person 4", Events: tc.events},
		{Id: 6, Name: "Person 6", Events: tc.events},
		{Id: 8, Name: "Person 8", Events: tc.events},
	}
	tc.personsEvents = []domain.PersonsAndEvents{
		{
			Id:       5,
			PersonId: 4,
			EventId:  1,
			Spent:    6,
			Factor:   1,
			Person:   tc.persons[0],
			Event:    tc.events[0],
		},
		{
			Id:       7,
			PersonId: 6,
			EventId:  1,
			Spent:    18,
			Factor:   2,
			Person:   tc.persons[1],
			Event:    tc.events[0],
		},
		{
			Id:       9,
			PersonId: 8,
			EventId:  1,
			Spent:    0,
			Factor:   3,
			Person:   tc.persons[2],
			Event:    tc.events[0],
		},
	}
	tc.RoundRate = 1
	tc.exp = expect{
		total:   decimal.NewFromFloat(24),
		count:   6,
		average: decimal.NewFromFloat(4),
		owes:    debtors{"Person 8": {"Person 4": decimal.NewFromFloat(2.0), "Person 6": decimal.NewFromFloat(10.0)}},
	}
	return &tc
}

func testCase2() *testCase {
	var tc testCase
	tc.testName = "#2. Two debetors to one recepient"
	tc.events = []domain.Event{
		{
			Id:      1,
			Name:    "New Year",
			Date:    time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local),
			Persons: tc.persons},
	}
	tc.persons = []domain.Person{
		{Id: 4, Name: "Person 4", Events: tc.events},
		{Id: 6, Name: "Person 6", Events: tc.events},
		{Id: 8, Name: "Person 8", Events: tc.events},
	}
	tc.personsEvents = []domain.PersonsAndEvents{
		{
			Id:       5,
			PersonId: 4,
			EventId:  1,
			Spent:    0,
			Factor:   1,
			Person:   tc.persons[0],
			Event:    tc.events[0],
		},
		{
			Id:       7,
			PersonId: 6,
			EventId:  1,
			Spent:    6,
			Factor:   2,
			Person:   tc.persons[1],
			Event:    tc.events[0],
		},
		{
			Id:       9,
			PersonId: 8,
			EventId:  1,
			Spent:    18,
			Factor:   3,
			Person:   tc.persons[2],
			Event:    tc.events[0],
		},
	}
	tc.RoundRate = 1
	tc.exp = expect{
		total:   decimal.NewFromFloat(24),
		count:   6,
		average: decimal.NewFromFloat(4),
		owes: debtors{
			"Person 4": {"Person 8": decimal.NewFromFloat(4.0)},
			"Person 6": {"Person 8": decimal.NewFromFloat(2.0)},
		},
	}
	return &tc
}

func testCase3() *testCase {
	var tc testCase
	tc.testName = "#3. All spent without debts"
	tc.events = []domain.Event{
		{
			Id:      1,
			Name:    "New Year",
			Date:    time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local),
			Persons: tc.persons},
	}
	tc.persons = []domain.Person{
		{Id: 4, Name: "Person 4", Events: tc.events},
		{Id: 6, Name: "Person 6", Events: tc.events},
		{Id: 8, Name: "Person 8", Events: tc.events},
	}
	tc.personsEvents = []domain.PersonsAndEvents{
		{
			Id:       5,
			PersonId: 4,
			EventId:  1,
			Spent:    6,
			Factor:   1,
			Person:   tc.persons[0],
			Event:    tc.events[0],
		},
		{
			Id:       7,
			PersonId: 6,
			EventId:  1,
			Spent:    12,
			Factor:   2,
			Person:   tc.persons[1],
			Event:    tc.events[0],
		},
		{
			Id:       9,
			PersonId: 8,
			EventId:  1,
			Spent:    18,
			Factor:   3,
			Person:   tc.persons[2],
			Event:    tc.events[0],
		},
	}
	tc.RoundRate = 1
	tc.exp = expect{
		total:   decimal.NewFromFloat(36),
		count:   6,
		average: decimal.NewFromFloat(6),
		owes:    nil,
	}
	return &tc
}

func testCase4() *testCase {
	var tc testCase
	tc.testName = "#4. All spents equal zero"
	tc.events = []domain.Event{
		{
			Id:      1,
			Name:    "New Year",
			Date:    time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local),
			Persons: tc.persons},
	}
	tc.persons = []domain.Person{
		{Id: 4, Name: "Person 4", Events: tc.events},
		{Id: 6, Name: "Person 6", Events: tc.events},
		{Id: 8, Name: "Person 8", Events: tc.events},
	}
	tc.personsEvents = []domain.PersonsAndEvents{
		{
			Id:       5,
			PersonId: 4,
			EventId:  1,
			Spent:    0,
			Factor:   1,
			Person:   tc.persons[0],
			Event:    tc.events[0],
		},
		{
			Id:       7,
			PersonId: 6,
			EventId:  1,
			Spent:    0,
			Factor:   2,
			Person:   tc.persons[1],
			Event:    tc.events[0],
		},
		{
			Id:       9,
			PersonId: 8,
			EventId:  1,
			Spent:    0,
			Factor:   3,
			Person:   tc.persons[2],
			Event:    tc.events[0],
		},
	}
	tc.RoundRate = 1
	tc.exp = expect{
		total:   decimal.Zero,
		count:   6,
		average: decimal.Zero,
		owes:    nil,
	}
	return &tc
}

func testCase5() *testCase {
	var tc testCase
	tc.testName = "#5. Owes with round rate 0.01"
	tc.events = []domain.Event{
		{
			Id:   1,
			Name: "New Year",
			Date: time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local),
		},
	}
	tc.persons = []domain.Person{
		{Id: 4, Name: "Person 4", Events: tc.events},
		{Id: 6, Name: "Person 6", Events: tc.events},
		{Id: 8, Name: "Person 8", Events: tc.events},
		{Id: 9, Name: "Person 9", Events: tc.events},
	}
	tc.personsEvents = []domain.PersonsAndEvents{
		{
			Id:       5,
			PersonId: 4,
			EventId:  1,
			Spent:    0,
			Factor:   2,
			Person:   tc.persons[0],
			Event:    tc.events[0],
		},
		{
			Id:       7,
			PersonId: 6,
			EventId:  1,
			Spent:    45,
			Factor:   2,
			Person:   tc.persons[1],
			Event:    tc.events[0],
		},
		{
			Id:       9,
			PersonId: 8,
			EventId:  1,
			Spent:    0,
			Factor:   2,
			Person:   tc.persons[2],
			Event:    tc.events[0],
		},
		{
			Id:       11,
			PersonId: 9,
			EventId:  1,
			Spent:    90,
			Factor:   1,
			Person:   tc.persons[3],
			Event:    tc.events[0],
		},
	}
	tc.RoundRate = 2
	tc.exp = expect{
		total:   decimal.NewFromFloat(135),
		count:   7,
		average: decimal.NewFromFloat(19.29),
		owes: debtors{
			"Person 4": {"Person 9": decimal.NewFromFloat(38.57)},
			"Person 8": {
				"Person 9": decimal.NewFromFloat(32.14),
				"Person 6": decimal.NewFromFloat(6.43),
			},
		},
	}
	return &tc
}
