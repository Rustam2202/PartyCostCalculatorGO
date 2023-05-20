package database

import (
	"party-calc/internal/config"
	"party-calc/internal/logger"
	"testing"

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
	var db DataBase
	logger.IntializeLogger()
	config.LoadConfig()

	id, err := db.AddPerson("Person 1", 1000, 3, 1)
	if err != nil {
		t.Errorf("Failed to ADD Person: %s", err)
	}

	per, err := db.GetPerson("Person 1")
	if err != nil {
		t.Errorf("Failed to GET Person: %s", err)
	}
	assert.Equal(t, "Person 1", per.Name)
	assert.Equal(t, 1000, per.Spent)
	assert.Equal(t, 3, per.Factor)

	err = db.DeletePerson(id)
	if err != nil {
		t.Errorf("Failed to DELETE Person: %s", err)
	}

	_, err = db.GetPerson("Person 1")
	if err == nil {
		t.Errorf("Expected error: %s", err)
	}
}
