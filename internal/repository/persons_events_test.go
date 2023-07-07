package repository

import (
	// "party-calc/internal/database"
	// "party-calc/internal/database/models"
	 "testing"
	// "time"

	// "github.com/DATA-DOG/go-sqlmock"
)

func TestCreatePersEvents(t *testing.T) {
	/*
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	repoPer := PersonRepository{Db: &database.DataBase{DB: db}}
	repoEv := EventRepository{Db: &database.DataBase{DB: db}}
	repoPerEv := PersEventsRepository{Db: &database.DataBase{DB: db}}

	perName := "Person 1"
	mock.ExpectQuery(`INSERT INTO persons (.+)`).WithArgs(perName).WillReturnRows(sqlmock.NewRows([]string{"Id"}).AddRow(1))
	person := models.Person{Name: perName}
	 err = repoPer.Create(&person)

	mock.ExpectQuery("INSERT INTO events (.+) VALUES(.+) RETURNING Id").
		WithArgs("Test Event", "2022-01-01").
		WillReturnRows(sqlmock.NewRows([]string{"Id"}).AddRow(1))
	ev := models.Event{Name: "Test Event", Date: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)}
	_, err = repoEv.Add(&ev)

	mock.ExpectQuery(`INSERT INTO pers_events (Person, Event, Spent, Factor) 
	VALUES (.+) RETURNING Id`).WithArgs(1, 1, 120.95, 2)
	mock.ExpectQuery(`UPDATE events SET Total = Total + (.+) WHERE Id = (.+)`).WithArgs(120.95, 2)
	
	persEv := models.PersonsAndEvents{PersonId: 1, EventId: 1, Spent: 120.95, Factor: 2}
	result, err := repoPerEv.Create(&persEv)

	if err != nil {
		t.Errorf("Add returned an error: %v", err)
	}
	if result != 1 {
		t.Errorf("Add returned the wrong result: got %d, want %d", result, 1)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
	*/
}
