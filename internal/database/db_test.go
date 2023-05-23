package database

import (
	"party-calc/internal/config"
	"party-calc/internal/logger"
	"party-calc/internal/person"
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

func TestCreateTebles(t *testing.T) {
	logger.IntializeLogger()
	config.LoadConfig()
	var db DataBase
	err := db.CreateTables()
	if err != nil {
		t.Fatalf("Failed to create tables: %s", err)
	}
}

func TestCRUDPersons(t *testing.T) {
	logger.IntializeLogger()
	config.LoadConfig()
	var db DataBase
	db.CreateTables()
	defer db.DropTables()

	var per = person.Person{Name: "Person 1", Spent: 1000}

	_, err := db.AddPerson(per, 1)
	if err != nil {
		t.Errorf("Failed to ADD Person: %s", err)
	}

	per_get, err := db.GetPerson(per.Name)
	if err != nil {
		t.Errorf("Failed to GET Person: %s", err)
	}
	assert.Equal(t, per.Name, per_get.Name)
	assert.Equal(t, per.Spent, per_get.Spent)
	assert.Equal(t, per.Factor, per_get.Factor)

	err = db.UpdatePerson(1, "Person 2", 1200, 2)
	if err != nil {
		t.Errorf("Failed to UPDATE Person: %s", err)
	}
	per_get, err = db.GetPerson("Person 2")
	if err != nil {
		t.Errorf("Failed to GET Person: %s", err)
	}
	assert.Equal(t, "Person 2", per_get.Name)
	assert.Equal(t, uint(1200), per_get.Spent)
	assert.Equal(t, uint(2), per_get.Factor)

	err = db.DeletePerson(1)
	if err != nil {
		t.Errorf("Failed to DELETE Person: %s", err)
	}
}

func TestCRUDEvents(t *testing.T) {
	logger.IntializeLogger()
	config.LoadConfig()
	var db DataBase
	db.CreateTables()
	defer db.DropTables()

	_, err := db.AddEvent("New year", time.Date(2022, 12, 31, 00, 00, 0, 00, time.Local))
	if err != nil {
		t.Errorf("Failed to INSERT Event: %s", err)
	}
	
}
