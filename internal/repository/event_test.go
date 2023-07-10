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

func TestAddEvent(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(ctx)

	repo := NewEventRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectQuery("INSERT INTO events").
		WithArgs("New Year", "2021-12-31").
		WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRow(int64(1)))

	ev := &domain.Event{Name: "New Year", Date: time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local)}

	err = repo.Create(ctx, ev)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.EqualValues(t, 1, ev.Id)
	assert.Equal(t, "New Year", ev.Name)
	assert.Equal(t, time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local), ev.Date)
}

func TestGetEventById(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close(ctx)

	repo := NewEventRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectQuery("SELECT id, name, date FROM events").
		WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name", "Date"}).
			AddRows([]any{int64(1), "New Year", time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local)}))
	// mock.ExpectQuery("SELECT id, name, date FROM events").
	// 	WithArgs(int64(2)).
	// 	WillReturnRows(pgxmock.NewRows([]string{"Id", "Name", "Date"}).
	// 		AddRow(int64(1), "Old New Year", time.Date(2022, 01, 14, 23, 59, 59, 0, time.Local))).RowsWillBeClosed()

	mock.ExpectQuery("SELECT person_id FROM persons_events").
		WithArgs(int64(1)).
		WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRow(int64(1)).AddRow(int64(2)).AddRow(int64(3)))
	// mock.ExpectQuery("SELECT person_id FROM persons_events").
	// 	WithArgs(int64(2)).
	// 	WillReturnRows(pgxmock.NewRows([]string{"Id"}).AddRow(int64(1)).AddRow(int64(3)).AddRow(int64(4)))

	mock.ExpectQuery("SELECT id, name FROM persons").
		WithArgs([]int64{1,2,3}).
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name"}).
			AddRow(int64(1), "Person 1").
			AddRow(int64(2), "Person 2").
			AddRow(int64(3), "Person 3"))

	ev := &domain.Event{}

	ev, err = repo.GetById(ctx, int64(1))

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
	assert.EqualValues(t, 1, ev.Id)
	assert.Equal(t, "New Year", ev.Name)
	assert.Equal(t, time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local), ev.Date)
}

func TestUpdateEvent(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating mock database connection: %s", err)
	}

	repo := NewEventRepository(&database.DataBase{DBPGX: mock})

	ev := &domain.Event{
		Id:   1,
		Name: "New Year",
		Date: time.Date(2021, 12, 31, 23, 59, 59, 0, time.Local),
	}

	mock.ExpectExec("UPDATE events").
		WithArgs(int64(1), "New Year", "2021-12-31").
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))

	err = repo.Update(ctx, ev)

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteEvent(t *testing.T) {
	var ctx context.Context = context.TODO()
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating mock database connection: %s", err)
	}

	repo := NewEventRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectExec("DELETE FROM events").
		WithArgs(int64(1)).
		WillReturnResult(pgxmock.NewResult("DELETE", 1))

	err = repo.DeleteById(ctx, int64(1))

	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
