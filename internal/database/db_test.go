package database

import (
	"fmt"
	"party-calc/internal/config"
	"party-calc/internal/database/models"
	"party-calc/internal/logger"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestAddPerson(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database connection: %v", err)
	}
	defer db.Close()

	d := &DataBase{DB: db}

	name := "John Doe"
	expectedId := int64(1)
	mock.ExpectQuery("INSERT INTO persons (.+) RETURNING Id").
		WithArgs(name).
		WillReturnRows(sqlmock.NewRows([]string{"Id"}).AddRow(expectedId))

	id, err := d.AddPerson(models.Person{Name: name})
	if err != nil {
		t.Fatalf("AddPerson returned an error: %v", err)
	}

	if id != expectedId {
		t.Errorf("AddPerson returned an incorrect ID. Expected %d, got %d", expectedId, id)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestGetPerson(t *testing.T) {
	name := "Person 1"
	mockBD := sqlmock.NewRows([]string{"id", name}).AddRow(1, name)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	database := &DataBase{DB: db}

	mock.ExpectQuery("SELECT * FROM persons WHERE name = $1").WithArgs(name).WillReturnRows(mockBD)

	per, err := database.GetPerson("Person 1")
	if err != nil {
		t.Fatalf("Failed to get person: %v", err)
	}

	if per.Id != 1 || per.Name != "Person 1" {
		t.Fatalf("Returned person does not match mock data")
	}
}

func TestUpdatePerson(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database connection: %v", err)
	}
	defer db.Close()

	database := &DataBase{DB: db}

	id := int64(1)
	per := models.Person{Name: "John Doe"}

	query := "UPDATE persons SET name=\\$1 WHERE id=\\$2"
	result := sqlmock.NewResult(0, 1)

	mock.ExpectExec(query).
		WithArgs(per.Name, id).
		WillReturnResult(result)

	err = database.UpdatePerson(id, per)
	if err != nil {
		t.Fatalf("Failed to update person: %v", err)
	}

	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatalf("Unfulfilled expectations: %v", err)
	}
}

func TestDeletePerson(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database connection: %v", err)
	}
	defer db.Close()

	database := &DataBase{DB: db}

	query := "DELETE FROM persons WHERE id=\\?"
	result := sqlmock.NewResult(0, 1)

	mock.ExpectExec(query).WithArgs(1).WillReturnResult(result)

	err = database.DeletePerson(1)

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func TestMain(t *testing.T) {
	logger.IntializeLogger()
	config.LoadConfig()
	var db DataBase

	err := db.Open()
	if err != nil {
		t.Fatalf("Failed to open database: %s", err)
	}
	err = db.DB.Ping()
	if err != nil {
		t.Fatalf("Failed to ping database: %s", err)
	}
	err = db.DB.Close()
	if err != nil {
		t.Fatalf("Failed to close database: %s", err)
	}
}

func TestCRUDPersons(t *testing.T) {
	logger.IntializeLogger()
	config.LoadConfig()
	var db DataBase

	lastId, err := db.AddPerson(models.Person{Name: "Person 5"})
	if err != nil {
		t.Errorf("Failed to ADD Person: %s", err)
	}
	fmt.Println(lastId)

	per_get, err := db.GetPerson("Person 1")
	if err != nil {
		t.Errorf("Failed to GET Person: %s", err)
	}
	assert.Equal(t, "Person 1", per_get.Name)

	err = db.UpdatePerson(1, models.Person{Name: "Person 2"})
	if err != nil {
		t.Errorf("Failed to UPDATE Person: %s", err)
	}
	per_get, err = db.GetPerson("Person 2")
	if err != nil {
		t.Errorf("Failed to GET Person: %s", err)
	}
	assert.Equal(t, "Person 2", per_get.Name)

	err = db.DeletePerson(1)
	if err != nil {
		t.Errorf("Failed to DELETE Person: %s", err)
	}
}

func TestCRUDEvents(t *testing.T) {
	logger.IntializeLogger()
	config.LoadConfig()
	var db DataBase

	_, err := db.AddEvent(models.Event{Name: "New year", Date: time.Date(2022, 12, 31, 00, 00, 0, 00, time.Local)})
	if err != nil {
		t.Errorf("Failed to INSERT Event: %s", err)
	}

}
