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

func TestNewEvent(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(ctx)

	repo := repository.NewEventRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectQuery("INSERT INTO events").
		WithArgs("New Year", "2021-12-31").
		WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRow(int64(1)))

	serv := NewEventService(repo)
	id, err := serv.NewEvent(ctx, "New Year", time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local))

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.EqualValues(t, 1, id)
}

func TestGetEventById(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(ctx)

	repo := repository.NewEventRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectQuery("SELECT id, name, date FROM events").
		WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name", "Date"}).
			AddRows([]any{int64(1), "New Year", time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local)}))

	mock.ExpectQuery("SELECT person_id FROM persons_events").
		WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id"}).
			AddRow(int64(1)).AddRow(int64(2)).AddRow(int64(3)))

	mock.ExpectQuery("SELECT id, name FROM persons").
		WithArgs([]int64{1, 2, 3}).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name"}).
			AddRow(int64(1), "Person 1").
			AddRow(int64(2), "Person 2").
			AddRow(int64(3), "Person 3"))

	serv := NewEventService(repo)
	ev, err := serv.GetEventById(ctx, 1)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.EqualValues(t, 1, ev.Id)
	assert.Equal(t, "New Year", ev.Name)
	assert.Equal(t, time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local), ev.Date)
	assert.Equal(t, 3, len(ev.Persons))
	assert.EqualValues(t, 3, ev.Persons[2].Id)
	assert.Equal(t, "Person 2", ev.Persons[1].Name)
}

func TestGetEventByName(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(ctx)

	repo := repository.NewEventRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectQuery("SELECT id, name, date FROM events").
		WithArgs("New Year").
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name", "Date"}).
			AddRows([]any{int64(1), "New Year", time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local)}))

	mock.ExpectQuery("SELECT person_id FROM persons_events").
		WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id"}).
			AddRow(int64(1)).AddRow(int64(2)).AddRow(int64(3)))

	mock.ExpectQuery("SELECT id, name FROM persons").
		WithArgs([]int64{1, 2, 3}).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name"}).
			AddRow(int64(1), "Person 1").
			AddRow(int64(2), "Person 2").
			AddRow(int64(3), "Person 3"))

	serv := NewEventService(repo)
	ev, err := serv.GetEventByName(ctx, "New Year")

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.EqualValues(t, 1, ev.Id)
	assert.Equal(t, "New Year", ev.Name)
	assert.Equal(t, time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local), ev.Date)
	assert.Equal(t, 3, len(ev.Persons))
	assert.EqualValues(t, 3, ev.Persons[2].Id)
	assert.Equal(t, "Person 2", ev.Persons[1].Name)
}

func TestUpdateEvent(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating mock database connection: %s", err)
	}

	repo := repository.NewEventRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectExec("UPDATE events").
		WithArgs(int64(1), "New Year", "2021-12-31").
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	serv := NewEventService(repo)
	err = serv.UpdateEvent(ctx, 1, "New Year", time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local))

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteEventById(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating mock database connection: %s", err)
	}

	repo := repository.NewEventRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectExec("DELETE FROM events").
		WithArgs(int64(1)).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	serv := NewEventService(repo)
	err = serv.DeleteEventById(ctx, 1)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteEventByName(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating mock database connection: %s", err)
	}

	repo := repository.NewEventRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectExec("DELETE FROM events").
		WithArgs("New Year").
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	serv := NewEventService(repo)
	err = serv.DeleteEventByName(ctx, "New Year")

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
