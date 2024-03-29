package repository

import (
	"context"
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

func (r *PersEventsRepository) Get(ctx context.Context, id int64) (*domain.PersonsAndEvents, error) {
	var result domain.PersonsAndEvents
	var per domain.Person
	var ev domain.Event
	err := r.Db.DBPGX.QueryRow(ctx,
		`SELECT id, person_id, event_id, spent, factor FROM persons_events WHERE id=$1`, id).
		Scan(&result.Id, &result.PersonId, &result.EventId, &result.Spent, &result.Factor)
	if err != nil {
		return nil, err
	}
	err = r.Db.DBPGX.QueryRow(ctx,
		`SELECT * FROM persons WHERE id=$1`, result.PersonId).Scan(&per.Id, &per.Name)
	if err != nil {
		return nil, err
	}
	result.Person = per
	err = r.Db.DBPGX.QueryRow(ctx,
		`SELECT * FROM events WHERE id=$1`, result.EventId).Scan(&ev.Id, &ev.Name, &ev.Date)
	if err != nil {
		return nil, err
	}
	result.Event = ev
	return &result, nil
}

func (r *PersEventsRepository) fillPersonAndEventsArrayForGet(ctx context.Context, rows pgx.Rows) ([]domain.PersonsAndEvents, error) {
	var result []domain.PersonsAndEvents
	for rows.Next() {
		var perEv domain.PersonsAndEvents
		err := rows.Scan(&perEv.Id, &perEv.PersonId, &perEv.EventId, &perEv.Spent, &perEv.Factor)
		if err != nil {
			return nil, err
		}
		result = append(result, perEv)
	}
	for i := 0; i < len(result); i++ {
		var per domain.Person
		err := r.Db.DBPGX.QueryRow(ctx,
			`SELECT * FROM persons WHERE id=$1`, result[i].PersonId).
			Scan(&per.Id, &per.Name)
		if err != nil {
			logger.Logger.Error("Failed Scan data from 'persons' by id: ", zap.Error(err))
			return nil, err
		}
		result[i].Person = per
	}
	for i := 0; i < len(result); i++ {
		var ev domain.Event
		err := r.Db.DBPGX.QueryRow(ctx,
			`SELECT id, name, date FROM events WHERE id=$1`, result[i].EventId).
			Scan(&ev.Id, &ev.Name, &ev.Date)
		if err != nil {
			logger.Logger.Error("Failed Scan data from 'events' by id: ", zap.Error(err))
			return nil, err
		}
		result[i].Event = ev
	}
	return result, nil
}

func (r *PersEventsRepository) GetByPersonId(ctx context.Context, personId int64) ([]domain.PersonsAndEvents, error) {
	var rows pgx.Rows
	var err error
	rows, err = r.Db.DBPGX.Query(ctx,
		`SELECT * FROM persons_events WHERE person_id=$1`, personId)
	if err != nil {
		return nil, err
	}
	result, err := r.fillPersonAndEventsArrayForGet(ctx, rows)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *PersEventsRepository) GetByEventId(ctx context.Context, eventId int64) ([]domain.PersonsAndEvents, error) {
	var rows pgx.Rows
	var err error
	rows, err = r.Db.DBPGX.Query(ctx,
		`SELECT id, person_id, event_id, spent, factor FROM persons_events WHERE event_id=$1`, eventId)
	if err != nil {
		return nil, err
	}
	result, err := r.fillPersonAndEventsArrayForGet(ctx, rows)
	if err != nil {
		return nil, err
	}
	return result, nil
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
