package tests

import (
	"party-calc/internal/database"
	"party-calc/internal/database/models"
	"party-calc/internal/database/repository"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestAddEvent(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := repository.EventRepository{Db: &database.DataBase{DB: db}}

	mock.ExpectQuery("INSERT INTO events (.+) VALUES(.+) RETURNING Id").
		WithArgs("Test Event", "2022-01-01").
		WillReturnRows(sqlmock.NewRows([]string{"Id"}).AddRow(1))

	ev := models.Event{Name: "Test Event", Date: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)}
	result, err := repo.Add(&ev)

	if err != nil {
		t.Errorf("Add returned an error: %v", err)
	}
	if result != 1 {
		t.Errorf("Add returned the wrong result: got %d, want %d", result, 1)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestGetEvent(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := repository.EventRepository{Db: &database.DataBase{DB: db}}

	rows := sqlmock.NewRows([]string{"Id", "Name", "Date", "TotalAmount"}).
		AddRow(1, "Test Event", "2022-01-01", 0)
	mock.ExpectQuery("SELECT (.+) FROM events WHERE name=(.+)").
		WithArgs("Test Event").
		WillReturnRows(rows)

	ev := models.Event{Name: "Test Event"}
	result, err := repo.Get(&ev)

	expected := models.Event{
		Id:          1,
		Name:        "Test Event",
		Date:        time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		TotalAmount: 0,
	}
	if err != nil {
		t.Errorf("Get returned an error: %v", err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Get returned the wrong result: got %v, want %v", result, expected)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestUpdateEvent(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := repository.EventRepository{Db: &database.DataBase{DB: db}}

	mock.ExpectExec("UPDATE events SET name=(.+), date=(.+) WHERE id=(.+)").
		WithArgs(1, "New Test Event", "2022-01-01").
		WillReturnResult(sqlmock.NewResult(0, 1))

	evOld := models.Event{Id: 1, Name: "Test Event", Date: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)}
	evNew := models.Event{Id: 1, Name: "New Test Event", Date: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)}
	err = repo.Update(&evOld, &evNew)

	if err != nil {
		t.Errorf("Update returned an error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}

}

func TestDeleteEvent(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := repository.EventRepository{Db: &database.DataBase{DB: db}}

	mock.ExpectExec("DELETE FROM events WHERE name=(.+)").
		WithArgs("Test Event").
		WillReturnResult(sqlmock.NewResult(0, 1))

	ev := models.Event{Name: "Test Event"}
	err = repo.Delete(&ev)

	if err != nil {
		t.Errorf("Delete returned an error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}
