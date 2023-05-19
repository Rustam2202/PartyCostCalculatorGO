package database

import (
	"database/sql"
	"fmt"
	"time"

	"party-calc/internal/config"
	"party-calc/internal/person"

	// "party-calc/internal/database/models"
	"party-calc/internal/logger"
	"party-calc/internal/service"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var cfg = &config.Cfg.DataBase

type DataBase struct {
	db *sql.DB
}

func (db *DataBase) Open() error {
	var err error
	psqlconn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname)
	db.db, err = sql.Open("postgres", psqlconn)
	if err != nil {
		logger.Logger.Error("Can't open database: ", zap.Error(err))
		return err
	}
	// defer db.db.Close()
	return nil
}

func (db *DataBase) CreateTables() error {
	db.Open()
	createTables := `
	CREATE TABLE IF NOT EXISTS persons (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		spent INTEGER,
		factor INTEGER
	);
	CREATE TABLE IF NOT EXISTS events (
		id SERIAL PRIMARY KEY,
		name TEXT,
		date DATE 
	);
	CREATE TABLE IF NOT EXISTS pers_events (
		id SERIAL PRIMARY KEY,
		person INTEGER,
		event INTEGER
	)`
	_, err := db.db.Exec(createTables)
	if err != nil {
		logger.Logger.Error("Failed to create teables:", zap.Error(err))
		return err
	}
	return nil
}

func (db *DataBase) AddPerson(name string, spent, factor, eventId int) (int64, error) {
	err := db.Open()
	if err != nil {
		return 0, err
	}
	stmt, err := db.db.Prepare(`INSERT INTO persons (name, spent, factor)
		VALUES($1,$2,$3)`)
	if err != nil {
		logger.Logger.Error("Statement prepare is incorrect: ", zap.Error(err))
		return 0, err
	}
	result, err := stmt.Exec(name, spent, factor)
	if err != nil {
		logger.Logger.Error("Failed to Execute Insert operation: ", zap.Error(err))
		return 0, err
	}
	personId, _ := result.LastInsertId()
	return personId, nil
}

func (db *DataBase) GetPerson(id int64) (person.Person, error) {
	var per person.Person
	var ids int
	var name string
	var spent int
	var factor int
	row := db.db.QueryRow(`SELECT * FROM persons WHERE id = $1`, id)
		//Scan(&per.Id, &per.Name, &per.Spent, &per.Factor)
	err:=	row.Scan(&ids, &name, &spent, &factor)
	if err != nil {
		logger.Logger.Error("Failed to Scan data: ", zap.Error(err))
		return person.Person{}, err
	}
	return per, nil
}

func (db *DataBase) UpdatePerson(id int64, name string, spent int, factor int) error {
	stmt, err := db.db.Prepare(`UPDATE persons SET name=$1 spent=$2, factor=$3 WHERE id=$4`)
	if err != nil {
		logger.Logger.Error("Statement prepare is incorrect: ", zap.Error(err))
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, spent, factor)
	if err != nil {
		logger.Logger.Error("Failed to Execute Update operation: ", zap.Error(err))
		return err
	}
	return nil
}

func (db *DataBase) DeletePerson(id int64) error {
	stmt, err := db.db.Prepare(`DELETE FROM persons WHERE id=$1`)
	if err != nil {
		logger.Logger.Error("Statement prepare is incorrect: ", zap.Error(err))
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		logger.Logger.Error("Failed to Execute Delete operation:: ", zap.Error(err))
		return err
	}
	return nil
}

func (db *DataBase) AddEvent(name string, date time.Time) int64 {
	stmt, err := db.db.Prepare(`INSERT INTO events (name, date)
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
func (db *DataBase) GetEvent(id int) int64 {
	var ev service.PartyData
	err := db.db.QueryRow(`SELECT id, name, spent, factor FROM events
		WHERE id=$1`, id).Scan(&ev.Id)
	if err != nil {

	}
	return 0
}

func UpdateEvent(id int) {

}

func DeleteEvent(id int) {

}
