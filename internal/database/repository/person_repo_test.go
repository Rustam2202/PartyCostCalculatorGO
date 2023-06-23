package repository

import (
	"party-calc/internal/database"
	"party-calc/internal/database/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := PersonRepository{db: &database.DataBase{DB: db}}

	name := "Person 1"
	mock.ExpectQuery(`INSERT INTO persons (.+)`).WithArgs(name).WillReturnRows(sqlmock.NewRows([]string{"Id"}).AddRow(1))

	person := models.Person{Name: name}
	id, err := repo.Create(&person)

	if err != nil {
		t.Errorf("Create returned an error: %v", err)
	}
	if id != 1 {
		t.Errorf("Create returned the wrong ID: %d", id)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := PersonRepository{db: &database.DataBase{DB: db}}

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Person 1").AddRow(2, "Person 2")
	mock.ExpectQuery(`SELECT (.+) FROM persons WHERE name = (.+)`).WithArgs("Person 1").WillReturnRows(rows)
	mock.ExpectQuery(`SELECT (.+) FROM persons WHERE id = (.+)`).WithArgs(2).WillReturnRows(rows)

	per1 := models.Person{Name: "Person 1"}
	per2 := models.Person{Id: 2}
	result1, err1 := repo.Get(&per1)
	result2, err2 := repo.Get(&per2)

	exp1 := models.Person{Id: 1, Name: "Person 1"}
	exp2 := models.Person{Id: 2, Name: "Person 2"}

	if err1 != nil {
		t.Errorf("Get returned an error: %v", err1)
	}
	if err2 != nil {
		t.Errorf("Get returned an error: %v", err2)
	}
	if result1 != exp1 {
		t.Errorf("Get returned the wrong result: %v", result1)
	}
	if result2 != exp2 {
		t.Errorf("Get returned the wrong result: %v", result2)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}

func TestUpdate(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("failed to create mock database: %v", err)
    }
    defer db.Close()

    repo := PersonRepository{db: &database.DataBase{DB: db}}

    mock.ExpectExec("UPDATE persons SET name=(.+) WHERE id=(.+)").
        WithArgs("Person 2", 1).
        WillReturnResult(sqlmock.NewResult(0, 1))

    perOld := models.Person{Id: 1, Name: "Person 1"}
    perNew := models.Person{Id: 1, Name: "Person 2"}
    err = repo.Update(&perOld, &perNew)

    if err != nil {
        t.Errorf("Update returned an error: %v", err)
    }

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %v", err)
    }
}

func TestDelete(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("failed to create mock database: %v", err)
    }
    defer db.Close()

    repo := PersonRepository{db: &database.DataBase{DB: db}}

    mock.ExpectExec("DELETE FROM persons WHERE name=(.+)").
        WithArgs("Person 1").
        WillReturnResult(sqlmock.NewResult(0, 1))

    per := models.Person{Name: "Person 1"}
    err = repo.Delete(&per)

    if err != nil {
        t.Errorf("Delete returned an error: %v", err)
    }

    if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("there were unfulfilled expectations: %v", err)
    }
}

