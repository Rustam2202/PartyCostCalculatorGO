package repository

import (
	"context"
	"errors"
	"party-calc/internal/database"
	"party-calc/internal/domain"
	"party-calc/internal/logger"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type PersEventsRepository struct {
	Db *database.DataBase
}

func NewPersEventsRepository(db *database.DataBase, pr *PersonRepository, evr *EventRepository) *PersEventsRepository {
	return &PersEventsRepository{Db: db}
}

func (r *PersEventsRepository) Create(pe *domain.PersonsAndEvents) error {
	var lastInsertedId int64
	err := r.Db.DBPGX.QueryRow(context.Background(),
		`INSERT INTO persons_events (person_id, event_id, spent, factor) 
		VALUES ($1, $2, $3, $4) RETURNING Id`,
		pe.PersonId, pe.EventId, pe.Spent, pe.Factor).Scan(&lastInsertedId)
	//r.Db.DB.QueryRow(`UPDATE events SET Total = Total + $2 WHERE Id = $1`, pe.EventId, pe.Spent)
	if err != nil {
		logger.Logger.Error("Failed to INSERT to 'pers_events' or UPDATE 'events': ",
			zap.Error(err))
		return err
	}
	pe.Id = lastInsertedId
	return nil
}

func (r *PersEventsRepository) GetByPersonId(id int64) (*domain.PersonsAndEvents, error) {
	var row pgx.Row
	
	if id != 0 {
		row = r.Db.DBPGX.QueryRow(context.Background(),
			`SELECT * FROM persons_events WHERE person_id=$1`, id)
	} else {
		return nil, errors.New("incorrect input, 'Person Id' mustn't be zero")
	}
	var result domain.PersonsAndEvents
	err := row.Scan(&result.Id, &result.PersonId, &result.EventId, &result.Spent, &result.Factor)
	if err != nil {
		logger.Logger.Error("Failed to Scan data from pers_events:", zap.Error(err))
		return nil, err
	}
	return &result, nil
}

func (r *PersEventsRepository) GetByEventId(id int64) (*domain.PersonsAndEvents, error) {
	var row pgx.Row
	if id != 0 {
		row = r.Db.DBPGX.QueryRow(context.Background(),
			`SELECT * FROM persons_events WHERE event_id=$1`, id)
	} else {
		return nil, errors.New("incorrect input, 'Event Id' mustn't be zero")
	}
	var result domain.PersonsAndEvents
	err := row.Scan(&result.Id, &result.PersonId, &result.EventId, &result.Spent, &result.Factor)
	if err != nil {
		logger.Logger.Error("Failed to Scan data from pers_events:", zap.Error(err))
		return nil, err
	}
	return &result, nil
}

func (r *PersEventsRepository) Update(pe *domain.PersonsAndEvents) error {
	_, err := r.Db.DBPGX.Exec(context.Background(),
		`UPDATE persons_events SET person_id=$2, event_id=$3, spent=$4, factor=$5 WHERE id=$1;`,
		pe.Id, pe.PersonId, pe.EventId, pe.Spent, pe.Factor,
	)
	if err != nil {
		logger.Logger.Error("Failed to Execute Update operation: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *PersEventsRepository) Delete(id int64) error {
	_, err := r.Db.DBPGX.Exec(context.Background(),
		`DELETE FROM persons_events WHERE id=$1;`, id)
	if err != nil {
		logger.Logger.Error("Failed to Execute Delete operation: ", zap.Error(err))
		return err
	}
	return nil
}
