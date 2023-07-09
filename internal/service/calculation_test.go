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

	//repoPer := repository.NewPersonRepository(&database.DataBase{DBPGX: mock})
	repoEv := repository.NewEventRepository(&database.DataBase{DBPGX: mock})
	repoPerEv := repository.NewPersonsEventsRepository(&database.DataBase{DBPGX: mock})

	// Add Persons
	{
		mock.ExpectQuery("INSERT INTO persons").WithArgs("Person 1").
			WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRow(int64(1)))
		mock.ExpectQuery("INSERT INTO persons").WithArgs("Person 2").
			WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRow(int64(2)))
		mock.ExpectQuery("INSERT INTO persons").WithArgs("Person 3").
			WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRow(int64(3)))
		mock.ExpectQuery("INSERT INTO persons").WithArgs("Person 4").
			WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRow(int64(4)))

		// Add Events
		mock.ExpectQuery("INSERT INTO events").WithArgs("New Year", "2021-12-31").
			WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRow(int64(1)))
		mock.ExpectQuery("INSERT INTO events").WithArgs("Old New Year", "2022-01-14").
			WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRow(int64(2)))

		// Add PersonsEvents
		mock.ExpectQuery("INSERT INTO persons_events").WithArgs(int64(1), int64(1), 10, 2).
			WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRow(int64(1)))
		mock.ExpectQuery("INSERT INTO persons_events").WithArgs(int64(2), int64(1), 5, 1).
			WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRow(int64(1)))
		mock.ExpectQuery("INSERT INTO persons_events").WithArgs(int64(3), int64(1), 0, 1).
			WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRow(int64(1)))

		mock.ExpectQuery("INSERT INTO persons_events").WithArgs(int64(1), int64(2), 0, 3).
			WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRow(int64(1)))
		mock.ExpectQuery("INSERT INTO persons_events").WithArgs(int64(3), int64(2), 2, 2).
			WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRow(int64(1)))
		mock.ExpectQuery("INSERT INTO persons_events").WithArgs(int64(4), int64(2), 6, 1).
			WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRow(int64(1)))
	}

	// Get Events
	mock.ExpectQuery("SELECT (.+) FROM events").WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name", "Date"}).
			AddRows([]any{int64(1), "New Year", time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local)}))
	mock.ExpectQuery("SELECT (.+) FROM events").WithArgs(int64(2)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name", "Date"}).
			AddRows([]any{int64(1), "Old New Year", time.Date(2022, 01, 14, 23, 59, 59, 0, time.Local)}))

	mock.ExpectQuery("SELECT person_id FROM persons_events").WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRows([]any{int64(1)}))
	mock.ExpectQuery("SELECT person_id FROM persons_events").WithArgs(int64(2)).
		WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRows([]any{int64(1)}))

	mock.ExpectQuery("SELECT id, name FROM persons").WithArgs([]int64{1}).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name"}).AddRows([]any{int64(1), "Person 1"}))
	mock.ExpectQuery("SELECT id, name FROM persons").WithArgs([]int64{2}).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name"}).AddRows([]any{int64(2), "Person 2"}))
	mock.ExpectQuery("SELECT id, name FROM persons").WithArgs([]int64{3}).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name"}).AddRows([]any{int64(3), "Person 3"}))
	mock.ExpectQuery("SELECT id, name FROM persons").WithArgs([]int64{4}).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name"}).AddRows([]any{int64(4), "Person 4"}))

	servCalc := NewCalcService(NewEventService(repoEv), NewPersonsEventsService(repoPerEv))

	result, err := servCalc.CalcEvent(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, nil, result.Balances)
}
