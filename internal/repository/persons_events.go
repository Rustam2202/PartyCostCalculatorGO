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

func NewPersonsEventsRepository(db *database.DataBase) *PersEventsRepository {
	return &PersEventsRepository{Db: db}
}

func (r *PersEventsRepository) Create(ctx context.Context, pe *domain.PersonsAndEvents) error {
	var lastInsertedId int64
	err := r.Db.DBPGX.QueryRow(ctx,
		`INSERT INTO persons_events (person_id, event_id, spent, factor) 
		VALUES ($1, $2, $3, $4) RETURNING Id`,
		pe.PersonId, pe.EventId, pe.Spent, pe.Factor).Scan(&lastInsertedId)
	if err != nil {
		logger.Logger.Error("Failed Insert to 'persons_events' table: ", zap.Error(err))
		return err
	}
	pe.Id = lastInsertedId
	return nil
}

func (r *PersEventsRepository) GetByPersonId(ctx context.Context, id int64) (*domain.PersonsAndEvents, error) {
	var row pgx.Row
	if id != 0 {
		row = r.Db.DBPGX.QueryRow(ctx,
			`SELECT * FROM persons_events WHERE person_id=$1`, id)
	} else {
		return nil, errors.New("incorrect input, 'Person Id' mustn't be zero")
	}
	var result domain.PersonsAndEvents
	err := row.Scan(&result.Id, &result.PersonId, &result.EventId, &result.Spent, &result.Factor)
	if err != nil {
		logger.Logger.Error("Failed to Scan data from 'persons_events' table: ", zap.Error(err))
		return nil, err
	}
	return &result, nil
}

func (r *PersEventsRepository) GetByEventId(ctx context.Context, id int64) (*domain.PersonsAndEvents, error) {
	var row pgx.Row
	if id != 0 {
		row = r.Db.DBPGX.QueryRow(ctx,
			`SELECT * FROM persons_events WHERE event_id=$1`, id)
	} else {
		return nil, errors.New("incorrect input, 'Event Id' mustn't be zero")
	}
	var result domain.PersonsAndEvents
	err := row.Scan(&result.Id, &result.PersonId, &result.EventId, &result.Spent, &result.Factor)
	if err != nil {
		logger.Logger.Error("Failed to Scan data from 'persons_events' table: ", zap.Error(err))
		return nil, err
	}
	return &result, nil
}

func (r *PersEventsRepository) Update(ctx context.Context, pe *domain.PersonsAndEvents) error {
	_, err := r.Db.DBPGX.Exec(ctx,
		`UPDATE persons_events SET person_id=$2, event_id=$3, spent=$4, factor=$5 WHERE id=$1;`,
		pe.Id, pe.PersonId, pe.EventId, pe.Spent, pe.Factor,
	)
	if err != nil {
		logger.Logger.Error("Failed Update in 'persons_events' table: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *PersEventsRepository) Delete(ctx context.Context, id int64) error {
	_, err := r.Db.DBPGX.Exec(ctx,
		`DELETE FROM persons_events WHERE id=$1;`, id)
	if err != nil {
		logger.Logger.Error("Failed Delete in 'persons_events' table: ", zap.Error(err))
		return err
	}
	return nil
}
