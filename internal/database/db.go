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
	err := db.Open()
	if err != nil {
		return err
	}
	defer db.db.Close()

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
	_, err = db.db.Exec(createTables)
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
	defer db.db.Close()

	result, err := db.db.Exec(`INSERT INTO persons (name, spent, factor)
		VALUES($1,$2,$3)`, name, spent, factor)

	if err != nil {
		logger.Logger.Error("Failed to Execute Insert to 'persons' table: ", zap.Error(err))
		return 0, err
	}
	personId, _ := result.LastInsertId() // ?? always returns 0

	_, err = db.db.Exec(`INSERT INTO pers_events (person)
		VALUES($1) WHERE id=$2`, personId, eventId)
	if err != nil {
		logger.Logger.Error("Failed to Execute Insert to 'pers_events' table: ", zap.Error(err))
		return 0, err
	}

	return personId, nil
}

func (db *DataBase) GetPerson(name string) (person.Person, error) {
	err := db.Open()
	if err != nil {
		return person.Person{}, err
	}
	defer db.db.Close()

	var per person.Person
	err = db.db.QueryRow(`SELECT * FROM persons WHERE name = $1`, name).
		Scan(&per.Id, &per.Name, &per.Spent, &per.Factor)
	if err != nil {
		logger.Logger.Error("Failed to Scan data: ", zap.Error(err))
		return person.Person{}, err
	}
	return per, nil
}

func (db *DataBase) UpdatePerson(id int64, name string, spent int, factor int) error {
	err := db.Open()
	if err != nil {
		return err
	}
	defer db.db.Close()

	stmt, err := db.db.Prepare(`UPDATE persons SET name=$1 spent=$2, factor=$3 WHERE id=$4`)
	if err != nil {
		logger.Logger.Error("Statement prepare is incorrect: ", zap.Error(err))
		return err
	}
	_, err = stmt.Exec(name, spent, factor)
	if err != nil {
		logger.Logger.Error("Failed to Execute Update operation: ", zap.Error(err))
		return err
	}
	return nil
}

func (db *DataBase) DeletePerson(id int64) error {
	err := db.Open()
	if err != nil {
		return err
	}
	defer db.db.Close()

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

func (db *DataBase) AddEvent(name string, date time.Time) (int64, error) {
	result, err := db.db.Exec(`
		INSERT INTO events (name, date) VALUES($1,$2);`,name, date)
	if err != nil {
		logger.Logger.Error("Couldn`t execute Insert operation")
		return 0, err
	}
	eventId, _ := result.LastInsertId()
	return eventId, nil
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
