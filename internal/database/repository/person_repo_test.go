package repository

import (
	"party-calc/internal/database"
	"party-calc/internal/database/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/pashagolub/pgxmock/v2"
)

func TestCreate(t *testing.T) {
	mock, err := pgxmock.NewPool()
	if err != nil {
		t.Fatal(err)
	}
	defer mock.Close()

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE products").
		WillReturnResult(pgxmock.NewResult("UPDATE", 1))
	mock.ExpectExec("INSERT INTO product_viewers").
		WithArgs(2, 3).
		WillReturnResult(pgxmock.NewResult("INSERT", 1))
	mock.ExpectCommit()

	// now we execute our method
	if err = recordStats(mock, 2, 3); err != nil {
		t.Errorf("error was not expected while updating: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	/*
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()
	repo := PersonRepository{Db: &database.DataBase{DBPGX: db}}

	name := "Person 1"
	mock.ExpectQuery(`INSERT INTO persons (.+)`).WithArgs(name).WillReturnRows(sqlmock.NewRows([]string{"Id"}).AddRow(1))

	person := models.Person{Name: name}
	err = repo.Create(&person)

	if err != nil {
		t.Errorf("Create returned an error: %v", err)
	}
	if person.Id != 1 {
		t.Errorf("Create returned the wrong ID: %d", person.Id)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
	*/
}

func TestGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := PersonRepository{Db: &database.DataBase{DB: db}}

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Person 1").AddRow(2, "Person 2")
	mock.ExpectQuery(`SELECT (.+) FROM persons WHERE name = (.+)`).WithArgs("Person 1").WillReturnRows(rows)
	mock.ExpectQuery(`SELECT (.+) FROM persons WHERE id = (.+)`).WithArgs(2).WillReturnRows(rows)

	result1, err1 := repo.GetByName("Person 2")
	result2, err2 := repo.GetById(1)

	exp1 := models.Person{Id: 1, Name: "Person 1"}
	exp2 := models.Person{Id: 2, Name: "Person 2"}

	if err1 != nil {
		t.Errorf("Get returned an error: %v", err1)
	}
	if err2 != nil {
		t.Errorf("Get returned an error: %v", err2)
	}
	if *result1 != exp1 {
		t.Errorf("Get returned the wrong result: %v", result1)
	}
	if *result2 != exp2 {
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

	repo := PersonRepository{Db: &database.DataBase{DB: db}}

	mock.ExpectExec("UPDATE persons SET name=(.+) WHERE id=(.+)").
		WithArgs("Person 2", 1).
		WillReturnResult(sqlmock.NewResult(0, 1))

	perNew := models.Person{Id: 1, Name: "Person 2"}
	err = repo.Update(&perNew)

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

	repo := PersonRepository{Db: &database.DataBase{DB: db}}

	mock.ExpectExec("DELETE FROM persons WHERE name=(.+)").
		WithArgs("Person 1").
		WillReturnResult(sqlmock.NewResult(0, 1))

	per := models.Person{Name: "Person 1"}
	err = repo.DeleteByName(per.Name)

	if err != nil {
		t.Errorf("Delete returned an error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %v", err)
	}
}
