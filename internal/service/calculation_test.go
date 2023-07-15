package service

import (
	"context"
	"party-calc/internal/database"
	"party-calc/internal/repository"
	"testing"
	"time"

	"github.com/pashagolub/pgxmock/v2"
	"github.com/stretchr/testify/assert"
)

func TestCalcEvent(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()

	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(ctx)

	repoPer := repository.NewPersonRepository(&database.DataBase{DBPGX: mock})
	repoEv := repository.NewEventRepository(&database.DataBase{DBPGX: mock})
	repoPerEv := repository.NewPersonsEventsRepository(&database.DataBase{DBPGX: mock})

	// Get from events
	mock.ExpectQuery("SELECT id, name, date FROM events").
		WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name", "Date"}).
			AddRows([]any{int64(1), "New Year", time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local)}))
	mock.ExpectQuery("SELECT person_id FROM persons_events").
		WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRow(int64(1)).AddRow(int64(2)).AddRow(int64(3)))
	mock.ExpectQuery("SELECT id, name FROM persons").
		WithArgs([]int64{1, 2, 3}).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name"}).
			AddRow(int64(1), "Person 1").
			AddRow(int64(2), "Person 2").
			AddRow(int64(3), "Person 3"))

	// Get PersonsEvents
	mock.ExpectQuery("SELECT (.+) FROM persons_events").
		WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "PersonId", "EventId", "Spent", "Factor"}).
			AddRow(int64(1), int64(1), int64(1), 18.0, 3))
	mock.ExpectQuery("SELECT (.+) FROM persons_events").
		WithArgs(int64(2)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "PersonId", "EventId", "Spent", "Factor"}).
			AddRow(int64(2), int64(2), int64(1), 3.0, 1))
	mock.ExpectQuery("SELECT (.+) FROM persons_events").
		WithArgs(int64(3)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "PersonId", "EventId", "Spent", "Factor"}).
			AddRow(int64(3), int64(3), int64(1), 0.0, 2))

	servCalc := NewCalcService(NewPersonService(repoPer), NewEventService(repoEv), NewPersonsEventsService(repoPerEv))
	result, err := servCalc.CalcEvent(ctx, 1)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.Equal(t, 3, len(result.Persons))
	assert.EqualValues(t, 6, result.AllPersonsCount)
	assert.EqualValues(t, 3.5, result.AverageAmount)
	assert.EqualValues(t, 6, result.AllPersonsCount)
	assert.EqualValues(t, 21, result.TotalAmount)

	assert.Equal(t, map[string]float64(map[string]float64(nil)), result.Persons[0].Owe)
	assert.Equal(t, map[string]float64{"Person 1": 0.5}, result.Persons[1].Owe)
	assert.Equal(t, map[string]float64{"Person 1": 7.0}, result.Persons[2].Owe)
}

func TestEventData_fillAndSortBalances(t *testing.T) {
	tests := []struct {
		name   string
		ed     *EventData
		expect []PersonBalance
	}{
		{
			name: "Test 1",
			ed: &EventData{
				Name: "",
				Date: time.Time{},
				Persons: []PersonData{
					{
						Id:     1,
						Spent:  100,
						Factor: 1,
					},
					{
						Id:     2,
						Spent:  50,
						Factor: 1,
					},
					{
						Id:     3,
						Spent:  0,
						Factor: 1,
					},
				},
				AllPersonsCount: 3,
				AverageAmount:   50,
				TotalAmount:     150,
			},
			expect: []PersonBalance{
				{
					Person:  &PersonData{Id: 3},
					Balance: -50,
				},
				{
					Person:  &PersonData{Id: 2},
					Balance: 0,
				},
				{
					Person:  &PersonData{Id: 1},
					Balance: 50,
				},
			},
		},
		{
			name: "Test 2",
			ed: &EventData{
				Name: "",
				Date: time.Time{},
				Persons: []PersonData{
					{
						Id:     1,
						Spent:  0,
						Factor: 3,
					},
					{
						Id:     2,
						Spent:  100,
						Factor: 1,
					},
					{
						Id:     3,
						Spent:  50,
						Factor: 1,
					},
				},
				AllPersonsCount: 5,
				AverageAmount:   30,
				TotalAmount:     150,
			},
			expect: []PersonBalance{
				{
					Person:  &PersonData{Id: 1},
					Balance: -90,
				},
				{
					Person:  &PersonData{Id: 3},
					Balance: 20,
				},
				{
					Person:  &PersonData{Id: 2},
					Balance: 70,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ed.fillAndSortBalances()
			for i := 0; i < len(tt.expect); i++ {
				assert.EqualValues(t, tt.expect[i].Person.Id, tt.ed.Balances[i].Person.Id)
				assert.EqualValues(t, tt.expect[i].Balance, tt.ed.Balances[i].Balance)
			}
		},
		)
	}
}

func TestEventData_calculateOwes(t *testing.T) {
	// Test 1
	{
		persons := []PersonData{
			{Id: 1, Name: "Person 1"},
			{Id: 2, Name: "Person 2"},
			{Id: 3, Name: "Person 3"},
		}
		evData := EventData{
			Persons: persons,
			Balances: []PersonBalance{
				{
					Person:  &persons[2],
					Balance: -50,
				},
				{
					Person:  &persons[1],
					Balance: 10,
				},
				{
					Person:  &persons[0],
					Balance: 40,
				},
			},
		}
		expect := []PersonData{
			{
				Id:  3,
				Owe: map[string]float64{"Person 1": 50, "Person 2": 10},
			},
		}

		evData.calculateOwes()
		assert.EqualValues(t, evData.Persons[2].Owe, expect[0].Owe, "Test 1 failed")
	}

	// Test 2
	{
		persons := []PersonData{
			{Id: 1, Name: "Person 1"},
			{Id: 2, Name: "Person 2"},
			{Id: 3, Name: "Person 3"},
			{Id: 4, Name: "Person 4"},
		}
		evData := EventData{
			Persons: persons,
			Balances: []PersonBalance{
				{
					Person:  &persons[0],
					Balance: -100,
				},
				{
					Person:  &persons[1],
					Balance: -100,
				},
				{
					Person:  &persons[2],
					Balance: 50,
				},
				{
					Person:  &persons[3],
					Balance: 150,
				},
			},
		}
		expect := []PersonData{
			{
				Id:  1,
				Owe: map[string]float64{"Person 4": 100},
			},
			{
				Id:  2,
				Owe: map[string]float64{"Person 3": 50, "Person 4": 50},
			},
		}
		evData.calculateOwes()
		// for i, v := range evData.Persons {

		// }
		assert.EqualValues(t, evData.Persons[2].Owe, expect[0].Owe)
	}
}
