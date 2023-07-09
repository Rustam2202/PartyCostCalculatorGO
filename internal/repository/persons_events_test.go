package repository

import (
	"context"
	"party-calc/internal/database"
	"party-calc/internal/domain"
	"testing"

	"github.com/pashagolub/pgxmock/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreatePersonsEvents(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(ctx)

	repo := NewPersonsEventsRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectQuery("INSERT INTO persons_events").
		WithArgs(int64(2), int64(3), 9.8, 1).
		WillReturnRows(pgxmock.NewRows([]string{"Id"}).
			AddRow(int64(1)))

	perEv := &domain.PersonsAndEvents{PersonId: 2, EventId: 3, Spent: 9.8, Factor: 1}
	assert.EqualValues(t, 0, perEv.Id)
	err = repo.Create(ctx, perEv)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.EqualValues(t, 1, perEv.Id)
}

func TestGetPersonsEventsByPersonId(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(ctx)

	repo := NewPersonsEventsRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectQuery("SELECT (.+) FROM persons_events").
		WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "PersonId", "EventId", "Spent", "Factor"}).
			AddRow(int64(1), int64(2), int64(3), 9.8, 1))

	perEv := &domain.PersonsAndEvents{}

	perEv, err = repo.GetByPersonId(ctx, int64(1))

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.EqualValues(t, 1, perEv.Id)
	assert.EqualValues(t, 2, perEv.PersonId)
	assert.EqualValues(t, 3, perEv.EventId)
	assert.EqualValues(t, 9.8, perEv.Spent)
	assert.EqualValues(t, 1, perEv.Factor)
}

func TestGetPersonsEventsByEventId(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(ctx)

	repo := NewPersonsEventsRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectQuery("SELECT (.+) FROM persons_events").
		WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "PersonId", "EventId", "Spent", "Factor"}).
			AddRow(int64(1), int64(2), int64(3), 9.8, 1))

	perEv := &domain.PersonsAndEvents{}

	perEv, err = repo.GetByPersonId(ctx, int64(1))
	if err != nil {
		t.Errorf("Error occurred while creating person: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}
	if perEv.Id != 1 {
		t.Errorf("Expected person ID to be 1, got %d", perEv.Id)
	}
}

func TestUpdatePersonsEvents(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating mock database connection: %s", err)
	}

	repo := NewPersonsEventsRepository(&database.DataBase{DBPGX: mock})

	perEv := &domain.PersonsAndEvents{
		Id:       1,
		PersonId: 2,
		EventId:  3,
		Spent:    9.8,
		Factor:   1,
	}

	mock.ExpectExec("UPDATE persons_events").
		WithArgs(int64(1), int64(2), int64(3), 9.8, 1).
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err = repo.Update(ctx, perEv)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeletePersonsEvents(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating mock database connection: %s", err)
	}

	repo := NewPersonsEventsRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectExec("DELETE FROM persons_events").
		WithArgs(int64(1)).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = repo.Delete(ctx, int64(1))

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
