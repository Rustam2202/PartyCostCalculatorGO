package repository

import (
	"context"
	"party-calc/internal/database"
	"party-calc/internal/domain"
	"party-calc/internal/logger"

	"go.uber.org/zap"
)

type EventRepository struct {
	Db *database.DataBase
}

func NewEventRepository(db *database.DataBase) *EventRepository {
	return &EventRepository{Db: db}
}

func (r *EventRepository) Create(ctx context.Context, ev *domain.Event) error {
	var lastInsertedId int64
	err := r.Db.DBPGX.QueryRow(ctx,
		`INSERT INTO events (name, date) VALUES($1,$2) RETURNING Id`,
		ev.Name, ev.Date.Format("2006-01-02")).Scan(&lastInsertedId)
	if err != nil {
		logger.Logger.Error("Failed Insert to 'events' table: ", zap.Error(err))
		return err
	}
	ev.Id = lastInsertedId
	return nil
}

func (r *EventRepository) GetById(ctx context.Context, id int64) (*domain.Event, error) {
	var result domain.Event
	err := r.Db.DBPGX.QueryRow(ctx,
		`SELECT id, name, date FROM events WHERE id=$1`, id).Scan(&result.Id, &result.Name, &result.Date)
	if err != nil {
		logger.Logger.Error("Failed Scan data from 'events' by id: ", zap.Error(err))
		return nil, err
	}

	// search persons ids existed in event
	rows, err := r.Db.DBPGX.Query(ctx,
		`SELECT person_id FROM persons_events WHERE event_id=$1`, result.Id)
	if err != nil {
		logger.Logger.Error("Failed take 'persons_ids' from 'events' table by id: ", zap.Error(err))
		return nil, err
	}
	var perIds []int64
	for rows.Next() {
		var perId int64
		err = rows.Scan(&perId)
		if err != nil {
			return nil, err
		}
		perIds = append(perIds, perId)
	}

	// append persons to event
	rows, err = r.Db.DBPGX.Query(ctx,
		`SELECT id, name FROM persons WHERE id=ANY($1)`, perIds)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var per domain.Person
		err = rows.Scan(&per.Id, &per.Name)
		if err != nil {
			logger.Logger.Error("Failed scan 'person_id' from 'events' rows: ", zap.Error(err))
			return nil, err
		}
		result.Persons = append(result.Persons, per)
	}
	return &result, nil
}

func (r *EventRepository) GetByName(ctx context.Context, name string) (*domain.Event, error) {
	var result domain.Event
	err := r.Db.DBPGX.QueryRow(ctx,
		`SELECT id, name, date FROM events WHERE name=$1`, name).Scan(&result.Id, &result.Name, &result.Date)
	if err != nil {
		logger.Logger.Error("Failed Scan data from 'events' by id: ", zap.Error(err))
		return nil, err
	}

	// search persons ids existed in event
	rows, err := r.Db.DBPGX.Query(ctx,
		`SELECT person_id FROM persons_events WHERE event_id=$1`, result.Id)
	if err != nil {
		logger.Logger.Error("Failed take 'persons_ids' from 'events' table by id: ", zap.Error(err))
		return nil, err
	}
	var perIds []int64
	for rows.Next() {
		var perId int64
		err = rows.Scan(&perId)
		if err != nil {
			return nil, err
		}
		perIds = append(perIds, perId)
	}

	// append persons to event
	rows, err = r.Db.DBPGX.Query(ctx,
		`SELECT id, name FROM persons WHERE id=ANY($1)`, perIds)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var per domain.Person
		err = rows.Scan(&per.Id, &per.Name)
		if err != nil {
			logger.Logger.Error("Failed scan 'person_id' from 'events' rows: ", zap.Error(err))
			return nil, err
		}
		result.Persons = append(result.Persons, per)
	}
	return &result, nil
}

func (r *EventRepository) Update(ctx context.Context, ev *domain.Event) error {
	_, err := r.Db.DBPGX.Exec(ctx,
		`UPDATE events SET name=$2, date=$3 WHERE id=$1`, ev.Id, ev.Name, ev.Date.Format("2006-01-02"))
	if err != nil {
		logger.Logger.Error("Failed Update in 'events' table: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *EventRepository) DeleteById(ctx context.Context, id int64) error {
	_, err := r.Db.DBPGX.Exec(ctx,
		`DELETE FROM events WHERE id=$1`, id)
	if err != nil {
		logger.Logger.Error("Failed Delete in 'events' table: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *EventRepository) DeleteByName(ctx context.Context, name string) error {
	_, err := r.Db.DBPGX.Exec(ctx,
		`DELETE FROM events WHERE name=$1`, name)
	if err != nil {
		logger.Logger.Error("Failed Delete in 'events' table: ", zap.Error(err))
		return err
	}
	return nil
}
