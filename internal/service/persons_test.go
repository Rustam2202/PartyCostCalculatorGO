package service

import (
	"context"
	"party-calc/internal/database"
	"party-calc/internal/domain"
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
	per := domain.Person{Name: "John Doe"}
	err = serv.repo.Create(ctx, &per)

	assert.NoError(t, err)
	assert.EqualValues(t, 1, per.Id)
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
	assert.EqualValues(t, 1, per.Id)
	assert.Equal(t, "John Doe", per.Name)
}
