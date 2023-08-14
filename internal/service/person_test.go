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

func TestNewPerson(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(ctx)

	repo := repository.NewPersonRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectQuery("INSERT INTO persons").
		WithArgs("John Doe").
		WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRow(int64(1)))

	serv := NewPersonService(repo)
	id, err := serv.NewPerson(ctx, "John Doe")

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.EqualValues(t, 1, id)
}

func TestGetPersonById(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(ctx)

	repo := repository.NewPersonRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectQuery("SELECT (.+) FROM persons").
		WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name"}).AddRows([]any{int64(1), "John Doe"}))

	mock.ExpectQuery("SELECT event_id FROM persons_events").
		WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRows([]any{int64(1)}))

	mock.ExpectQuery("SELECT id, name, date FROM events").
		WithArgs([]int64{1}).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name", "Date"}).
			AddRows([]any{int64(1), "New Year", time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local)}))

	serv := NewPersonService(repo)
	per, err := serv.GetPersonById(ctx, 1)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.EqualValues(t, 1, per.Id)
	assert.Equal(t, "John Doe", per.Name)
	assert.EqualValues(t, 1, per.Events[0].Id)
	assert.Equal(t, "New Year", per.Events[0].Name)
}

func TestGetPersonByName(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(ctx)

	repo := repository.NewPersonRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectQuery("SELECT (.+) FROM persons").
		WithArgs("John Doe").
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name"}).AddRows([]any{int64(1), "John Doe"}))

	mock.ExpectQuery("SELECT event_id FROM persons_events").
		WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRow(int64(2)).AddRow(int64(3)))

	mock.ExpectQuery("SELECT id, name, date FROM events").
		WithArgs([]int64{2, 3}).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name", "Date"}).
			AddRows([]any{int64(2), "New Year", time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local)}).
			AddRows([]any{int64(3), "Old New Year", time.Date(2022, 01, 14, 23, 59, 59, 0, time.Local)}))

	serv := NewPersonService(repo)
	per, err := serv.GetPersonByName(ctx, "John Doe")

	assert.NoError(t, err)
	assert.EqualValues(t, 1, per.Id)
	assert.Equal(t, "John Doe", per.Name)
	assert.EqualValues(t, 2, per.Events[0].Id)
	assert.Equal(t, "New Year", per.Events[0].Name)
	assert.EqualValues(t, 3, per.Events[1].Id)
	assert.Equal(t, "Old New Year", per.Events[1].Name)
}

func TestUpdatePerson(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating mock database connection: %s", err)
	}

	repo := repository.NewPersonRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectExec("UPDATE persons").
		WithArgs(int64(1), "Doe John").
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	serv := NewPersonService(repo)
	err = serv.UpdatePerson(ctx, 1, "Doe John")

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeletePersonById(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating mock database connection: %s", err)
	}

	repo := repository.NewPersonRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectExec("DELETE FROM persons").
		WithArgs(int64(1)).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	serv := NewPersonService(repo)
	err = serv.DeletePersonById(ctx, 1)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeletePersonByName(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating mock database connection: %s", err)
	}

	repo := repository.NewPersonRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectExec("DELETE FROM persons").
		WithArgs("John Doe").
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	serv := NewPersonService(repo)
	err = serv.DeletePersonByName(ctx, "John Doe")

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
