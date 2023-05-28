package database

import (
	"fmt"
	"party-calc/internal/config"
	"party-calc/internal/database/models"
	"party-calc/internal/logger"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	logger.IntializeLogger()
	config.LoadConfig()
	var db DataBase

	err := db.Open()
	if err != nil {
		t.Fatalf("Failed to open database: %s", err)
	}
	err = db.db.Ping()
	if err != nil {
		t.Fatalf("Failed to ping database: %s", err)
	}
	err = db.db.Close()
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
