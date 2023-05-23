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

func (db *DataBase) DropTables() error {
	err := db.Open()
	if err != nil {
		return err
	}
	defer db.db.Close()

	dropTables := `
		DROP TABLE IF EXISTS persons;
		DROP TABLE IF EXISTS events;
		DROP TABLE IF EXISTS pers_events
		`
	_, err = db.db.Exec(dropTables)
	if err != nil {
		logger.Logger.Error("Failed to drop teables:", zap.Error(err))
		return err
	}
	return nil
}

func (db *DataBase) AddPerson(per person.Person, eventId int) (int64, error) {
	err := db.Open()
	if err != nil {
		return 0, err
	}
	defer db.db.Close()

	result, err := db.db.Exec(`INSERT INTO persons (name, spent, factor)
		VALUES($1,$2,$3)`, per.Name, per.Spent, per.Factor)

	if err != nil {
		logger.Logger.Error("Failed to Execute Insert to 'persons' table: ", zap.Error(err))
		return 0, err
	}
	personId, _ := result.LastInsertId() // ?? always returns 0

	// _, err = db.db.Exec(`INSERT INTO pers_events (person)
	// 	VALUES($1) WHERE id=$2`, personId, eventId)
	// if err != nil {
	// 	logger.Logger.Error("Failed to Execute Insert to 'pers_events' table: ", zap.Error(err))
	// 	return 0, err
	// }

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
		logger.Logger.Error("Failed to Scan data from persons: ", zap.Error(err))
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

	_, err = db.db.Exec(
		`UPDATE persons SET name=$1, spent=$2, factor=$3 WHERE id=$4`,
		name, spent, factor, id)
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

	_, err = db.db.Exec(`DELETE FROM persons WHERE id=$1`, id)
	if err != nil {
		logger.Logger.Error("Failed to Execute Delete operation: ", zap.Error(err))
		return err
	}
	return nil
}

func (db *DataBase) AddEvent(name string, date time.Time) (int64, error) {
	err := db.Open()
	if err != nil {
		return 0, err
	}
	defer db.db.Close()

	result, err := db.db.Exec(`INSERT INTO events (name, date) VALUES($1,$2);`, name, date.Format("2006-01-02"))
	if err != nil {
		logger.Logger.Error("Couldn`t execute Insert operation", zap.Error(err))
		return 0, err
	}
	eventId, _ := result.LastInsertId() // ?? always returns 0
	return eventId, nil
}

func (db *DataBase) GetEvent(name string) (service.PartyData, error) {
	err := db.Open()
	if err != nil {
		return service.PartyData{}, err
	}
	defer db.db.Close()

	var ev service.PartyData
	err = db.db.QueryRow(`SELECT * FROM events WHERE name=$1`, name).Scan(&ev.Id)
	if err != nil {
		logger.Logger.Error("Failed to Scan data from events:", zap.Error(err))
		return service.PartyData{}, err
	}
	return ev, nil
}

func (db *DataBase) UpdateEvent(id int, name string, date time.Time) error {
	err := db.Open()
	if err != nil {
		return err
	}
	defer db.db.Close()

	// stmt, err := db.db.Prepare(`UPDATE events SET name=$1, fate=$2 WHERE id=$3`)
	// if err != nil {
	// 	logger.Logger.Error("Statement prepare is incorrect: ", zap.Error(err))
	// 	return err
	// }
	_, err = db.db.Exec(`UPDATE events SET name=$1, date=$2 WHERE id=$3`, name, date, id)
	//_, err = stmt.Exec(name, date)
	if err != nil {
		logger.Logger.Error("Failed to Execute Update operation: ", zap.Error(err))
		return err
	}
	return nil
}

func (db *DataBase) DeleteEvent(id int) error {
	err := db.Open()
	if err != nil {
		return err
	}
	defer db.db.Close()

	// stmt, err := db.db.Prepare(`DELETE FROM evenets WHERE id=$1`)
	// if err != nil {
	// 	logger.Logger.Error("Statement prepare is incorrect: ", zap.Error(err))
	// 	return err
	// }

	_, err = db.db.Exec(`DELETE FROM evenets WHERE id=$1`, id)
	if err != nil {
		logger.Logger.Error("Failed to Execute Delete operation: ", zap.Error(err))
		return err
	}
	return nil
}
