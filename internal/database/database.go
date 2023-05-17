package database

import (
	"database/sql"
	"fmt"
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

	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT NOT NULL
	)
`
	_, err = DB.Exec(createTable)
	if err != nil {
		logger.Logger.Error("Can't create teable")
	}

}

func Insert() {

}
