package repository

import (
	"context"
	"party-calc/internal/database"
	"party-calc/internal/domain"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v2"
)

func TestCreate(t *testing.T) {
	mockDB,mock, err :=sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	//defer mock.Close()

	conn,err:=pgx.Connect(context.Background(), "mock connection string")

	repo := NewPersonRepository(&database.DataBase{DBPGX: mock})

	mock.ExpectQuery("INSERT INTO persons").
		WithArgs("John Doe").
		WillReturnRows(pgxmock.NewRows([]string{"Id", "Name"}).AddRow(1, "John Doe"))

	person := &domain.Person{Name: "John Doe"}

	err = repo.Create(person)
	if err != nil {
		t.Errorf("Error occurred while creating person: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("There were unfulfilled expectations: %s", err)
	}

	if person.Id != 1 {
		t.Errorf("Expected person ID to be 1, got %d", person.Id)
	}
}
