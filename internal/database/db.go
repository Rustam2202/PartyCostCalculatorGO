package database

import (
	"database/sql"
	"fmt"

	"party-calc/internal/config"
	"party-calc/internal/database/models"

	"party-calc/internal/logger"

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
	return nil
}

func (db *DataBase) AddPerson(per models.Person) (int64, error) {
	err := db.Open()
	if err != nil {
		return 0, err
	}
	defer db.db.Close()

	var lastInsertedId int64
	err = db.db.QueryRow(`INSERT INTO persons (name) VALUES($1) RETURNING Id`, per.Name).Scan(&lastInsertedId)
	//result, err := db.db.Exec(`INSERT INTO persons (name) VALUES($1) RETURNING Id`, per.Name)
	if err != nil {
		logger.Logger.Error("Failed to Execute Insert to 'persons' table: ", zap.Error(err))
		return 0, err
	}
	return lastInsertedId, nil
}

func (db *DataBase) GetPerson(name string) (models.Person, error) {
	err := db.Open()
	if err != nil {
		return models.Person{}, err
	}
	defer db.db.Close()

	var per models.Person
	err = db.db.QueryRow(`SELECT * FROM persons WHERE name = $1`, name).
		Scan(&per.Id, &per.Name)
	if err != nil {
		logger.Logger.Error("Failed to Scan data from persons: ", zap.Error(err))
		return models.Person{}, err
	}
	return per, nil
}

func (db *DataBase) UpdatePerson(id int64, per models.Person) error {
	err := db.Open()
	if err != nil {
		return err
	}
	defer db.db.Close()

	_, err = db.db.Exec(
		`UPDATE persons SET name=$1 WHERE id=$2`,
		per.Name, id)
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

func (db *DataBase) AddEvent(event models.Event) (int64, error) {
	err := db.Open()
	if err != nil {
		return 0, err
	}
	defer db.db.Close()

	var lastInsertedId int64
	err = db.db.QueryRow(`INSERT INTO events (name, date) VALUES($1,$2) RETURNING Id;`,
		event.Name, event.Date.Format("2006-01-02")).Scan(&lastInsertedId)
	if err != nil {
		logger.Logger.Error("Couldn`t execute Insert operation", zap.Error(err))
		return 0, err
	}
	return lastInsertedId, nil
}

func (db *DataBase) GetEvent(name string) (models.Event, error) {
	err := db.Open()
	if err != nil {
		return models.Event{}, err
	}
	defer db.db.Close()

	var ev models.Event
	err = db.db.QueryRow(`SELECT * FROM events WHERE name=$1`, name).Scan(&ev.Id, &ev.Name, &ev.Date)
	if err != nil {
		logger.Logger.Error("Failed to Scan data from events:", zap.Error(err))
		return models.Event{}, err
	}
	return ev, nil
}

func (db *DataBase) UpdateEvent(id int, ev models.Event) error {
	err := db.Open()
	if err != nil {
		return err
	}
	defer db.db.Close()

	_, err = db.db.Exec(`UPDATE events SET name=$1, date=$2 WHERE id=$3`, ev.Name, ev.Date, id)
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

	_, err = db.db.Exec(`DELETE FROM evenets WHERE id=$1`, id)
	if err != nil {
		logger.Logger.Error("Failed to Execute Delete operation: ", zap.Error(err))
		return err
	}
	return nil
}

func (db *DataBase) AddPersonToEvent(evId, perId int64) error {
	err := db.Open()
	if err != nil {
		return err
	}
	defer db.db.Close()

	_, err = db.db.Exec(`INSERT INTO pers_events (Person, Event) VALUES ($1,$2)`, evId, perId)
	if err != nil {
		logger.Logger.Error("Failed to Execute Insert to 'pers_events' table: ", zap.Error(err))
		return err
	}
	return nil
}

func (db *DataBase) AddPersonToEventWithSpent(evId, perId int64, spent float32, factor int) error {
	err := db.Open()
	if err != nil {
		return err
	}
	defer db.db.Close()

	_, err = db.db.Exec(`INSERT INTO pers_events (Person, Event, Spent, Factor) 
		VALUES ($1,$2,$3,$4); 
		UPDATE events SET Total=Total+$3 WHERE Id=$1
		`, evId, perId, spent, factor)
	if err != nil {
		logger.Logger.Error("Failed to Execute Insert to 'pers_events' table: ", zap.Error(err))
		return err
	}
	return nil
}

func (db *DataBase) GetPersEvents(name string) (models.PersonsAndEvents, error) {
	err := db.Open()
	if err != nil {
		return models.PersonsAndEvents{}, err
	}
	defer db.db.Close()

	var pe models.PersonsAndEvents
	err = db.db.QueryRow(`SELECT * FROM pers_events WHERE name=$1`, name).
		Scan(&pe.Id, &pe.PersonId, &pe.EventId, &pe.Spent, &pe.Factor)
	if err != nil {
		logger.Logger.Error("Failed to Scan data from pers_events:", zap.Error(err))
		return models.PersonsAndEvents{}, err
	}
	return pe, nil
}

func (db *DataBase) UpdatePersEvents(evId, perId int64, spent float32, factor int) error {
	err := db.Open()
	if err != nil {
		return err
	}
	defer db.db.Close()

	_, err = db.db.Exec(`UPDATE pers_events SET spent=$3, factor=$4 WHERE Event=$1, Person=$2`, evId, perId, spent, factor)
	if err != nil {
		logger.Logger.Error("Failed to Execute Update operation: ", zap.Error(err))
		return err
	}
	return nil
}

func (db *DataBase) DeletePersEvents(id int64) error {
	err := db.Open()
	if err != nil {
		return err
	}
	defer db.db.Close()

	_, err = db.db.Exec(`DELETE FROM pers_evenets WHERE id=$1`, id)
	if err != nil {
		logger.Logger.Error("Failed to Execute Delete operation: ", zap.Error(err))
		return err
	}
	return nil
}