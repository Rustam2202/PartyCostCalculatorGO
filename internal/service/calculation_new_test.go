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

func TestCreateResponse(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()

	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(ctx)

	//repoPer := repository.NewPersonRepository(&database.DataBase{DBPGX: mock})
	repoEv := repository.NewEventRepository(&database.DataBase{DBPGX: mock})
	repoPerEv := repository.NewPersonsEventsRepository(&database.DataBase{DBPGX: mock})

	// Get from events
	// mock.ExpectQuery("SELECT id, name, date FROM events").
	// 	WithArgs(int64(1)).
	// 	WillReturnRows(pgxmock.NewRows([]string{"Id", "Name", "Date"}).
	// 		AddRows([]any{int64(1), "New Year", time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local)}))
	// mock.ExpectQuery("SELECT person_id FROM persons_events").
	// 	WithArgs(int64(1)).
	// 	WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRow(int64(1)).AddRow(int64(2)).AddRow(int64(3)))
	// mock.ExpectQuery("SELECT id, name FROM persons").
	// 	WithArgs([]int64{1, 2, 3}).
	// 	WillReturnRows(pgxmock.NewRows([]string{"Id", "Name"}).
	// 		AddRow(int64(1), "Person 1").
	// 		AddRow(int64(2), "Person 2").
	// 		AddRow(int64(3), "Person 3"))

	// Get from PersonsEvents
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

	servCalc := NewCalcService(NewEventService(repoEv), NewPersonsEventsService(repoPerEv))
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
