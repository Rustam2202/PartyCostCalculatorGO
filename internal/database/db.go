package database

import (
	"database/sql"
	"fmt"
	"time"

	"party-calc/internal/database/models"

	"party-calc/internal/logger"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type DataBase struct {
	DB  *sql.DB
	CFG DatabaseConfig
}

func (db *DataBase) Open(dbCfg DatabaseConfig) error {
	var err error

	psqlconn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable") //dbCfg.Database.User, dbCfg.Database.Password, dbCfg.Database.Host,
	//	dbCfg.Database.Port, dbCfg.Database.Dbname

	db.DB, err = sql.Open("postgres", psqlconn)
	if err != nil {
		logger.Logger.Error("Can't open database: ", zap.Error(err))
		return err
	}
	return nil
}

func (db *DataBase) AddPerson(per models.Person) (int64, error) {
	var lastInsertedId int64
	err := db.DB.QueryRow(`INSERT INTO persons (name) VALUES($1) RETURNING Id`, per.Name).
		Scan(&lastInsertedId)
	if err != nil {
		logger.Logger.Error("Failed to Execute Insert to 'persons' table: ", zap.Error(err))
		return 0, err
	}
	return lastInsertedId, nil
}

func (db *DataBase) GetPerson(name string) (models.Person, error) {
	var per models.Person
	err := db.DB.QueryRow(`SELECT * FROM persons WHERE name = $1`, name).
		Scan(&per.Id, &per.Name)
	if err != nil {
		logger.Logger.Error("Failed to Scan data from persons: ", zap.Error(err))
		return models.Person{}, err
	}
	return per, nil
}

func (db *DataBase) UpdatePerson(id int64, per models.Person) error {
	_, err := db.DB.Exec(
		`UPDATE persons SET name=$1 WHERE id=$2`,
		per.Name, id)
	if err != nil {
		logger.Logger.Error("Failed to Execute Update operation: ", zap.Error(err))
		return err
	}
	return nil
}

func (db *DataBase) DeletePerson(id int64) error {
	_, err := db.DB.Exec(`DELETE FROM persons WHERE id=$1`, id)
	if err != nil {
		logger.Logger.Error("Failed to Execute Delete operation: ", zap.Error(err))
		return err
	}
	return nil
}

func (db *DataBase) AddEvent(event models.Event) (int64, error) {
	var lastInsertedId int64
	err := db.DB.QueryRow(`INSERT INTO events (name, date) VALUES($1,$2) RETURNING Id;`,
		event.Name, event.Date.Format("2006-01-02")).Scan(&lastInsertedId)
	if err != nil {
		logger.Logger.Error("Couldn`t execute Insert operation", zap.Error(err))
		return 0, err
	}
	return lastInsertedId, nil
}

func (db *DataBase) GetEvent(name string) (models.Event, error) {
	var ev models.Event
	var date string
	err := db.DB.QueryRow(`SELECT * FROM events WHERE name=$1`, name).
		Scan(&ev.Id, &ev.Name, &date, &ev.TotalAmount)
	if err != nil {
		logger.Logger.Error("Failed to Scan data from events:", zap.Error(err))
		return models.Event{}, err
	}
	ev.Date, _ = time.Parse("2006-01-02", date)
	return ev, nil
}

func (db *DataBase) UpdateEvent(id int64, ev models.Event) error {
	_, err := db.DB.Exec(`UPDATE events SET name=$2, date=$3 WHERE id=$1`,
		id, ev.Name, ev.Date.Format("2006-01-02"))
	if err != nil {
		logger.Logger.Error("Failed to Execute Update operation: ", zap.Error(err))
		return err
	}
	return nil
}

func (db *DataBase) DeleteEvent(id int64) error {
	_, err := db.DB.Exec(`DELETE FROM events WHERE id=$1`, id)
	if err != nil {
		logger.Logger.Error("Failed to Execute Delete operation: ", zap.Error(err))
		return err
	}
	return nil
}

func (db *DataBase) AddPersonToEvent(evId, perId int64) (int64, error) {
	var lastInsertedId int64
	err := db.DB.QueryRow(`INSERT INTO pers_events (Person, Event) VALUES ($1,$2) RETURNING Id`,
		evId, perId).Scan(&lastInsertedId)
	if err != nil {
		logger.Logger.Error("Failed to Execute Insert to 'pers_events' table: ", zap.Error(err))
		return 0, err
	}
	return lastInsertedId, nil
}

func (db *DataBase) AddPersonToEventWithSpent(
	evId, perId int64, spent float64, factor int) (int64, error) {
	var lastInsertedId int64
	err := db.DB.QueryRow(`
		INSERT INTO pers_events (Person, Event, Spent, Factor) 
		VALUES ($1, $2, $3, $4) RETURNING Id;
		UPDATE events SET Total = Total + $3 WHERE Id = $1
		`, evId, perId, spent, factor).Scan(&lastInsertedId)
	if err != nil {
		logger.Logger.Error("Failed to INSERT to 'pers_events' or UPDATE 'events': ",
			zap.Error(err))
		return 0, err
	}
	return lastInsertedId, nil
}

func (db *DataBase) GetPersEvents(name string) (models.PersonsAndEvents, error) {
	var pe models.PersonsAndEvents
	err := db.DB.QueryRow(`SELECT * FROM pers_events WHERE name=$1`, name).
		Scan(&pe.Id, &pe.PersonId, &pe.EventId, &pe.Spent, &pe.Factor)
	if err != nil {
		logger.Logger.Error("Failed to Scan data from pers_events:", zap.Error(err))
		return models.PersonsAndEvents{}, err
	}
	return pe, nil
}

func (db *DataBase) GetPersFromEvents(id int64) (models.PersonsAndEvents, error) {
	var pe models.PersonsAndEvents
	err := db.DB.QueryRow(`SELECT * FROM pers_events WHERE id=$1`, id).
		Scan(&pe.Id, &pe.PersonId, &pe.EventId, &pe.Spent, &pe.Factor)
	if err != nil {
		logger.Logger.Error("Failed to Scan data from pers_events:", zap.Error(err))
		return models.PersonsAndEvents{}, err
	}
	return pe, nil
}

func (db *DataBase) UpdatePersEvents(evId, perId int64, spent float64, factor int) error {
	per, _ := db.GetPersFromEvents(perId)
	_, err := db.DB.Exec(`
		UPDATE pers_events SET spent=$3, factor=$4 WHERE Event=$1, Person=$2;
		UPDATE events SET Total=Total+$3-$5 WHERE id=$1
		`,
		evId, perId, spent, factor, per.Spent)
	if err != nil {
		logger.Logger.Error("Failed to Execute Update operation: ", zap.Error(err))
		return err
	}
	return nil
}

func (db *DataBase) DeletePersonFromEvents(perId int64) error {
	per, _ := db.GetPersFromEvents(perId)
	_, err := db.DB.Exec(`
		DELETE FROM pers_evenets WHERE Person=$1;
		UPDATE events SET Total=Total-$2 WHERE id=$1
	`, perId, per.Spent)
	if err != nil {
		logger.Logger.Error("Failed to Execute Delete operation: ", zap.Error(err))
		return err
	}
	return nil
}
