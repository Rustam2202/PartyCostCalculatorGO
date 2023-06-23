package repository

import (
	"database/sql"
	"errors"
	"party-calc/internal/database"
	"party-calc/internal/database/models"
	"party-calc/internal/logger"

	"go.uber.org/zap"
)

type PersEventsRepository struct {
	db         *database.DataBase
	PersRepo   *PersonRepository
	EventsRepo *EventRepository
}

func NewPersEventsRepository(db *database.DataBase, pr *PersonRepository, evr *EventRepository) *PersEventsRepository {
	return &PersEventsRepository{
		db:         db,
		PersRepo:   pr,
		EventsRepo: evr,
	}
}

func (r *PersEventsRepository) Create(pe *models.PersonsAndEvents) (int64, error) {
	var lastInsertedId int64
	err := r.db.DB.QueryRow(`
		INSERT INTO pers_events (Person, Event, Spent, Factor) 
		VALUES ($1, $2, $3, $4) RETURNING Id;
		UPDATE events SET Total = Total + $3 WHERE Id = $1
		`, pe.EventId, pe.PersonId, pe.Spent, pe.Factor).Scan(&lastInsertedId)
	if err != nil {
		logger.Logger.Error("Failed to INSERT to 'pers_events' or UPDATE 'events': ",
			zap.Error(err))
		return 0, err
	}
	return lastInsertedId, nil
}

func (r *PersEventsRepository) GetPerson(pe *models.PersonsAndEvents) (models.PersonsAndEvents, error) {
	var row *sql.Row
	if pe.PersonId != 0 {
		row = r.db.DB.QueryRow(`SELECT * FROM pers_events WHERE Person = $1`, pe.EventId)
	} else {
		return models.PersonsAndEvents{}, errors.New("incorrect input, 'Person Id' mustn't be zero")
	}
	var result models.PersonsAndEvents
	err := row.Scan(&result.Id, &result.PersonId, &result.EventId, &result.Spent, &result.Factor)
	if err != nil {
		logger.Logger.Error("Failed to Scan data from pers_events:", zap.Error(err))
		return models.PersonsAndEvents{}, err
	}
	return result, nil
}

func (r *PersEventsRepository) GetEvent(pe *models.PersonsAndEvents) ([]models.PersonsAndEvents, error) {
	var rows *sql.Rows
	var err error
	if pe.EventId != 0 {
		rows, err = r.db.DB.Query(`SELECT * FROM pers_events WHERE Event = $1`, pe.PersonId)
		if err != nil {
			logger.Logger.Error("Failed to Scan data from pers_events:", zap.Error(err))
			return nil, err
		}
	} else {
		return nil, errors.New("incorrect input, 'Event Id' mustn't be zero")
	}
	var result []models.PersonsAndEvents
	for rows.Next() {
		var perEv models.PersonsAndEvents
		err = rows.Scan(perEv.Id, perEv.PersonId, perEv.EventId, perEv.Spent, perEv.Factor)
		if err != nil {
			logger.Logger.Error("Failed to Scan data from pers_events:", zap.Error(err))
			return nil, err
		}
	}
	return result, nil
}

func (r *PersEventsRepository) Update(oldData, NewData *models.PersonsAndEvents) error {
	old, err := r.GetPerson(oldData)
	if err != nil {
		//logger.Logger.Error("Failed to Execute Update operation: ", zap.Error(err))
		return err
	}
	_, err = r.db.DB.Exec(`
		UPDATE pers_events SET spent=$3, factor=$4 WHERE Event=$1, Person=$2;
		UPDATE events SET Total=Total+$3-$5 WHERE id=$1`,
		oldData.Id, oldData.PersonId, NewData.Spent, NewData.Factor, old.Spent,
	)
	if err != nil {
		logger.Logger.Error("Failed to Execute Update operation: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *PersEventsRepository) Delete(perEv *models.PersonsAndEvents) error {
	_, err := r.db.DB.Exec(`
		DELETE FROM pers_evenets WHERE Person=$1;
		UPDATE events SET Total=Total-$2 WHERE id=$1`,
		perEv.Id, perEv.Spent)
	if err != nil {
		logger.Logger.Error("Failed to Execute Delete operation: ", zap.Error(err))
		return err
	}
	return nil
}
