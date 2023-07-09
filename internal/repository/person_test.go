package repository

import (
	"context"
	"party-calc/internal/database"
	"party-calc/internal/domain"
	"testing"
	"time"

	"github.com/pashagolub/pgxmock/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreatePerson(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(ctx)

	repo := NewPersonRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectQuery("INSERT INTO persons").
		WithArgs("John Doe").
		WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRow(int64(1)))

	person := &domain.Person{Name: "John Doe"}
	err = repo.Create(ctx, person)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.EqualValues(t, 1, person.Id)
}

func TestGetPersonById(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(ctx)

	repo := NewPersonRepository(&database.DataBase{DBPGX: mock})

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

	person := &domain.Person{}
	person, err = repo.GetById(ctx, int64(1))

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.EqualValues(t, 1, person.Id)
	assert.Equal(t, "John Doe", person.Name)
}

func TestUpdatePerson(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating mock database connection: %s", err)
	}

	repo := NewPersonRepository(&database.DataBase{DBPGX: mock})

	person := &domain.Person{
		Id:   1,
		Name: "Doe John",
	}

	mock.ExpectExec("UPDATE persons").
		WithArgs(int64(1), "Doe John").
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err = repo.Update(ctx, person)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeletePerson(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating mock database connection: %s", err)
	}

	repo := NewPersonRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectExec("DELETE FROM persons").
		WithArgs(int64(1)).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = repo.DeleteById(ctx, int64(1))

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
