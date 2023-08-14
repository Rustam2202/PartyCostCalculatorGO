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

func TestNewPersonsEvents(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(ctx)

	repo := repository.NewPersonsEventsRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectQuery("INSERT INTO persons_events").
		WithArgs(int64(2), int64(3), 9.8, 1).
		WillReturnRows(pgxmock.NewRows([]string{"Id"}).
			AddRow(int64(1)))

	serv := NewPersonsEventsService(repo)
	id, err := serv.AddPersonToPersEvent(ctx, 2, 3, 9.8, 1)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.EqualValues(t, 1, id)
}

func TestGetPersonsEventsByPersonId(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(ctx)

	repo := repository.NewPersonsEventsRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectQuery("SELECT (.+) FROM persons_events").
		WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "PersonId", "EventId", "Spent", "Factor"}).
			AddRow(int64(3), int64(1), int64(4), 9.8, 1).
			AddRow(int64(5), int64(1), int64(6), 0.5, 2))

	mock.ExpectQuery("SELECT (.+) FROM persons").
		WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name"}).
			AddRow(int64(1), "Person 1"))

	mock.ExpectQuery("SELECT id, name, date FROM events").
		WithArgs(int64(4)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name", "Date"}).
			AddRows([]any{int64(4), "New Year", time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local)}))

	mock.ExpectQuery("SELECT (.+) FROM persons").
		WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name"}).
			AddRow(int64(1), "Person 1"))

	mock.ExpectQuery("SELECT id, name, date FROM events").
		WithArgs(int64(6)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name", "Date"}).
			AddRows([]any{int64(6), "Old New Year", time.Date(2022, 01, 13, 23, 59, 59, 0, time.Local)}))

	serv := NewPersonsEventsService(repo)
	perEv, err := serv.GetByPersonId(ctx, 1)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.EqualValues(t, 3, perEv[0].Id)
	assert.EqualValues(t, 1, perEv[0].PersonId)
	assert.EqualValues(t, 4, perEv[0].EventId)
	assert.EqualValues(t, 9.8, perEv[0].Spent)
	assert.EqualValues(t, 1, perEv[0].Factor)
	assert.EqualValues(t, 1, perEv[0].Person.Id)
	assert.EqualValues(t, "Person 1", perEv[0].Person.Name)
	assert.EqualValues(t, 4, perEv[0].Event.Id)
	assert.EqualValues(t, "New Year", perEv[0].Event.Name)
	assert.EqualValues(t, time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local), perEv[0].Event.Date)

	assert.EqualValues(t, 5, perEv[1].Id)
	assert.EqualValues(t, 1, perEv[1].PersonId)
	assert.EqualValues(t, 6, perEv[1].EventId)
	assert.EqualValues(t, 0.5, perEv[1].Spent)
	assert.EqualValues(t, 2, perEv[1].Factor)
	assert.EqualValues(t, 1, perEv[1].Person.Id)
	assert.EqualValues(t, "Person 1", perEv[1].Person.Name)
	assert.EqualValues(t, 6, perEv[1].Event.Id)
	assert.EqualValues(t, "Old New Year", perEv[1].Event.Name)
	assert.EqualValues(t, time.Date(2022, 01, 13, 23, 59, 59, 0, time.Local), perEv[1].Event.Date)
}

func TestGetPersonsEventsByEventId(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(ctx)

	repo := repository.NewPersonsEventsRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectQuery("SELECT (.+) FROM persons_events").
		WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "PersonId", "EventId", "Spent", "Factor"}).
			AddRow(int64(3), int64(4), int64(1), 9.8, 1).
			AddRow(int64(5), int64(6), int64(1), 0.5, 2).
			AddRow(int64(7), int64(8), int64(1), 0.0, 3))

	mock.ExpectQuery("SELECT (.+) FROM persons").
		WithArgs(int64(4)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name"}).
			AddRow(int64(4), "Person 4"))
	mock.ExpectQuery("SELECT id, name, date FROM events").
		WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name", "Date"}).
			AddRows([]any{int64(1), "New Year", time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local)}))

	mock.ExpectQuery("SELECT (.+) FROM persons").
		WithArgs(int64(6)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name"}).
			AddRow(int64(6), "Person 6"))
	mock.ExpectQuery("SELECT id, name, date FROM events").
		WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name", "Date"}).
			AddRows([]any{int64(1), "New Year", time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local)}))

	mock.ExpectQuery("SELECT (.+) FROM persons").
		WithArgs(int64(8)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name"}).
			AddRow(int64(8), "Person 8"))
	mock.ExpectQuery("SELECT id, name, date FROM events").
		WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name", "Date"}).
			AddRows([]any{int64(1), "New Year", time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local)}))

	serv := NewPersonsEventsService(repo)
	perEv, err := serv.GetByEventId(ctx, 1)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.EqualValues(t, 3, perEv[0].Id)
	assert.EqualValues(t, 4, perEv[0].PersonId)
	assert.EqualValues(t, 1, perEv[0].EventId)
	assert.EqualValues(t, 9.8, perEv[0].Spent)
	assert.EqualValues(t, 1, perEv[0].Factor)
	assert.EqualValues(t, 4, perEv[0].Person.Id)
	assert.EqualValues(t, "Person 4", perEv[0].Person.Name)
	assert.EqualValues(t, 1, perEv[0].Event.Id)
	assert.EqualValues(t, "New Year", perEv[0].Event.Name)
	assert.EqualValues(t, time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local), perEv[0].Event.Date)

	assert.EqualValues(t, 5, perEv[1].Id)
	assert.EqualValues(t, 6, perEv[1].PersonId)
	assert.EqualValues(t, 1, perEv[1].EventId)
	assert.EqualValues(t, 0.5, perEv[1].Spent)
	assert.EqualValues(t, 2, perEv[1].Factor)
	assert.EqualValues(t, 6, perEv[1].Person.Id)
	assert.EqualValues(t, "Person 6", perEv[1].Person.Name)
	assert.EqualValues(t, 1, perEv[1].Event.Id)
	assert.EqualValues(t, "New Year", perEv[2].Event.Name)
	assert.EqualValues(t, time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local), perEv[2].Event.Date)

	assert.EqualValues(t, 7, perEv[2].Id)
	assert.EqualValues(t, 8, perEv[2].PersonId)
	assert.EqualValues(t, 1, perEv[2].EventId)
	assert.EqualValues(t, 0, perEv[2].Spent)
	assert.EqualValues(t, 3, perEv[2].Factor)
	assert.EqualValues(t, 8, perEv[2].Person.Id)
	assert.EqualValues(t, "Person 8", perEv[2].Person.Name)
	assert.EqualValues(t, 1, perEv[2].Event.Id)
	assert.EqualValues(t, "New Year", perEv[2].Event.Name)
	assert.EqualValues(t, time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local), perEv[2].Event.Date)
}

func TestUpdatePersonsEvents(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating mock database connection: %s", err)
	}

	repo := repository.NewPersonsEventsRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectExec("UPDATE persons_events").
		WithArgs(int64(1), int64(2), int64(3), 9.8, 1).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	serv := NewPersonsEventsService(repo)
	err = serv.Update(ctx, 1, 2, 3, 9.8, 1)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeletePersonsEvents(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating mock database connection: %s", err)
	}

	repo := repository.NewPersonsEventsRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectExec("DELETE FROM persons_events").
		WithArgs(int64(1)).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	serv := NewPersonsEventsService(repo)
	err = serv.Delete(ctx, int64(1))

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
