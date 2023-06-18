package repository

import (
	"party-calc/internal/database"
	"party-calc/internal/database/models"
	"party-calc/internal/logger"

	"go.uber.org/zap"
)

type PersEventsRepository struct {
	db *database.DataBase
}

func NewPersEventsRepository(db *database.DataBase) *PersEventsRepository {
	return &PersEventsRepository{db: db}
}

func (r *PersEventsRepository) Add(pe *models.PersonsAndEvents) (int64, error) {
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

func (r *PersEventsRepository) Get(perName string) (models.PersonsAndEvents, error) {
	var pe models.PersonsAndEvents
	err := r.db.DB.QueryRow(`SELECT * FROM pers_events WHERE Person = $1`, perName).
		Scan(&pe.Id, &pe.PersonId, &pe.EventId, &pe.Spent, &pe.Factor)
	if err != nil {
		logger.Logger.Error("Failed to Scan data from pers_events:", zap.Error(err))
		return models.PersonsAndEvents{}, err
	}
	return pe, nil
}

func (r *PersEventsRepository) GetPersFromEvents(id int64) (models.PersonsAndEvents, error) {
	var pe models.PersonsAndEvents
	err := r.db.DB.QueryRow(`SELECT * FROM pers_events WHERE id=$1`, id).
		Scan(&pe.Id, &pe.PersonId, &pe.EventId, &pe.Spent, &pe.Factor)
	if err != nil {
		logger.Logger.Error("Failed to Scan data from pers_events:", zap.Error(err))
		return models.PersonsAndEvents{}, err
	}
	return pe, nil
}

func (r*PersEventsRepository) UpdatePersEvents(evId, perId int64, spent float64, factor int) error {
	per, _ := r.GetPersFromEvents(perId)
	_, err := r.db.DB.Exec(`
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

func (r*PersEventsRepository) Delete(perId int64) error {
	per, _ := r.db.GetPersFromEvents(perId)
	_, err := r.db.DB.Exec(`
		DELETE FROM pers_evenets WHERE Person=$1;
		UPDATE events SET Total=Total-$2 WHERE id=$1
	`, perId, per.Spent)
	if err != nil {
		logger.Logger.Error("Failed to Execute Delete operation: ", zap.Error(err))
		return err
	}
	return nil
}
