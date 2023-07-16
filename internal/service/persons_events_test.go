package service

import (
	"context"
	"party-calc/internal/database"
	"party-calc/internal/domain"
	"party-calc/internal/repository"
	"testing"

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
			AddRow(int64(1), int64(2), int64(3), 9.8, 1))

	serv := NewPersonsEventsService(repo)
	perEv1, err := serv.GetByPersonId(ctx, 1)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.EqualValues(t, 1, perEv1.Id)
	assert.EqualValues(t, 2, perEv1.PersonId)
	assert.EqualValues(t, 3, perEv1.EventId)
	assert.EqualValues(t, 9.8, perEv1.Spent)
	assert.EqualValues(t, 1, perEv1.Factor)

	mock.ExpectQuery("SELECT (.+) FROM persons_events").
		WithArgs(int64(2)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "PersonId", "EventId", "Spent", "Factor"}).
			AddRow(int64(2), int64(3), int64(3), 5.0, 2))

	perEv2, err := repo.GetByPersonId(ctx, int64(2))
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.EqualValues(t, 2, perEv2.Id)
	assert.EqualValues(t, 3, perEv2.PersonId)
	assert.EqualValues(t, 3, perEv2.EventId)
	assert.EqualValues(t, 5, perEv2.Spent)
	assert.EqualValues(t, 2, perEv2.Factor)
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
			AddRow(int64(1), int64(2), int64(3), 9.8, 1))

	serv := NewPersonsEventsService(repo)
	perEv := &domain.PersonsAndEvents{}
	perEv, err = serv.GetByEventId(ctx, 1)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.EqualValues(t, 1, perEv.Id)
	assert.EqualValues(t, 2, perEv.PersonId)
	assert.EqualValues(t, 3, perEv.EventId)
	assert.EqualValues(t, 9.8, perEv.Spent)
	assert.EqualValues(t, 1, perEv.Factor)
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
