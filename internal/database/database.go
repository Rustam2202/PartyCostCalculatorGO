package database

import (
	"database/sql"
	"fmt"
	"time"

	//	"party-calc/internal/person"
	"party-calc/internal/config"
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

func AddPersonToDB(name string, spent, participants, eventId int) {
	stmt, err := DB.Prepare(`INSERT INTO persons (name, spent, participants, event)
		VALUES(?,?,?,?)`)
	if err != nil {
		logger.Logger.Error("Statement prepare is incorrect")
	}
	_, err = stmt.Exec(name, spent, participants, eventId)
	if err != nil {
		logger.Logger.Error("Couldn`t execute Insert operation")
	}
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
