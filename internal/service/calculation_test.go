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

func TestCalculateEvent1(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()

	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(ctx)

	repoPer := repository.NewPersonRepository(&database.DataBase{DBPGX: mock})
	repoEv := repository.NewEventRepository(&database.DataBase{DBPGX: mock})
	repoPerEv := repository.NewPersonsEventsRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectQuery("SELECT (.+) FROM persons_events").WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "PersonId", "EventId", "Spent", "Factor"}).
			AddRow(int64(3), int64(4), int64(1), 6.0, 1).
			AddRow(int64(5), int64(6), int64(1), 18.0, 2).
			AddRow(int64(7), int64(8), int64(1), 0.0, 3))

	mock.ExpectQuery("SELECT (.+) FROM persons").WithArgs(int64(4)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name"}).AddRow(int64(4), "Person 4"))
	mock.ExpectQuery("SELECT id, name, date FROM events").WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name", "Date"}).
			AddRows([]any{int64(1), "New Year", time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local)}))

	mock.ExpectQuery("SELECT (.+) FROM persons").WithArgs(int64(6)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name"}).AddRow(int64(6), "Person 6"))
	mock.ExpectQuery("SELECT id, name, date FROM events").WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name", "Date"}).
			AddRows([]any{int64(1), "New Year", time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local)}))

	mock.ExpectQuery("SELECT (.+) FROM persons").WithArgs(int64(8)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name"}).AddRow(int64(8), "Person 8"))
	mock.ExpectQuery("SELECT id, name, date FROM events").WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name", "Date"}).
			AddRows([]any{int64(1), "New Year", time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local)}))

	servCalc := NewCalcService(NewPersonService(repoPer), NewEventService(repoEv), NewPersonsEventsService(repoPerEv))
	result, err := servCalc.CalculateEvent(ctx, 1)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	//	assert.EqualValues(t, 1, len(result.Owes))
	assert.EqualValues(t, 24, result.TotalSpent)
	assert.EqualValues(t, 4, result.AverageSpent)
	assert.EqualValues(t, 6, result.PersonsCount)
	assert.EqualValues(t, map[string]map[string]float64{
		"Person 8": {"Person 6": 10.0, "Person 4": 2.0},
	}, result.PersonsOwes)
}
