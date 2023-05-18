package database

import (
	"database/sql"
	"fmt"
	"time"

	//	"party-calc/internal/person"
	"party-calc/internal/config"
	"party-calc/internal/database/models"
	"party-calc/internal/logger"

	_ "github.com/lib/pq"
)

var cfg = config.Cfg.DataBase

var DB *sql.DB

func Open() {
	var err error
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Dbname)
	//	connStr := "postgres://postgres:password@localhost/persons?sslmode=disable"
	DB, err = sql.Open("postgres", psqlconn)
	if err != nil {
		logger.Logger.Error("Can't open database")
	}
	defer DB.Close()

	createTable := `CREATE TABLE IF NOT EXISTS persons (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		spent INTEGER,
		participants INTEGER,
		events []INTEGER
	)`
	_, err = DB.Exec(createTable)
	if err != nil {
		logger.Logger.Error("Can't create teable")
	}
}

func AddPerson(name string, spent, factor, eventId int) int64 {
	stmt, err := DB.Prepare(`INSERT INTO persons (name, spent, participants, event)
		VALUES(?,?,?,?)`)
	if err != nil {
		logger.Logger.Error("Statement prepare is incorrect")
	}
	result, err := stmt.Exec(name, spent, factor, eventId)
	if err != nil {
		logger.Logger.Error("Couldn`t execute Insert operation")
	}
	personId, _ := result.LastInsertId()
	return personId
}

func GetPerson(id int) {
	var person models.Person
	err := DB.QueryRow(`SELECT id, name, spent, factor FROM persons
		WHERE id=$1`, id).Scan(&person.Id, &person.Name, &person.Spent, &person.Factor)
	if err != nil {

	}
}

func UpdatePerson(id int, name string, spent int, factor int) {
	stmt, err := DB.Prepare(`UPDATE persons SET name=$1 spent=$2, factor=$3 WHERE id=$4`)
	if err != nil {

	}
	defer stmt.Close()
	_, err = stmt.Exec(name, spent, factor)
}

func DeletePerson(id int) {
	stmt, err := DB.Prepare(`DELETE FROM persons WHERE id=$1`)
	if err != nil {

	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
}

func AddEvent(name string, date time.Time) int64 {
	stmt, err := DB.Prepare(`INSERT INTO events (name, date)
		VALUES(?,?)`)
	if err != nil {
		logger.Logger.Error("Statement prepare is incorrect")
	}
	result, err := stmt.Exec(name, date)
	if err != nil {
		logger.Logger.Error("Couldn`t execute Insert operation")
	}
	eventId, _ := result.LastInsertId()
	return eventId
}

// use interface for call with id, name or time arguments
func GetEvent(id int) int64 {

	return 0
}

func UpdateEvent(id int) {

}

func DeleteEvent(id int) {

}
