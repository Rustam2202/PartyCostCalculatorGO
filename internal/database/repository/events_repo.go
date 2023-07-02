package repository

import (
	"context"
	"party-calc/internal/database"
	"party-calc/internal/domain"
	"party-calc/internal/logger"
	"time"

	"go.uber.org/zap"
)

type EventRepository struct {
	Db *database.DataBase
}

func NewEventRepository(db *database.DataBase) *EventRepository {
	return &EventRepository{Db: db}
}

func (r *EventRepository) Add(ev *domain.Event) error {
	var lastInsertedId int64
	err := r.Db.DBPGX.QueryRow(context.Background(),
		`INSERT INTO events (name, date) VALUES($1,$2) RETURNING Id`,
		ev.Name, ev.Date.Format("2006-01-02")).Scan(&lastInsertedId)
	if err != nil {
		logger.Logger.Error("Failed Insert to 'events' table: ", zap.Error(err))
		return err
	}
	ev.Id = lastInsertedId
	return nil
}

func (r *EventRepository) GetById(id int64) (*domain.Event, error) {
	var result domain.Event
	var date string
	err := r.Db.DBPGX.QueryRow(context.Background(),
		`SELECT * FROM events WHERE id=$1`, id).Scan(&result.Id, &result.Name, &date)
	if err != nil {
		logger.Logger.Error("Failed Scan data from 'events' by id: ", zap.Error(err))
		return nil, err
	}
	rows, err := r.Db.DBPGX.Query(context.Background(),
		`SELECT * FROM persons_events WHERE event_id=$1`, result.Id)
	if err != nil {
		logger.Logger.Error("Failed take 'persons_ids' from 'events' table by id: ", zap.Error(err))
		return nil, err
	}
	for rows.Next() {
		var personId int64
		err = rows.Scan(&personId)
		if err != nil {
			logger.Logger.Error("Failed scan 'person_id' from 'events' rows: ", zap.Error(err))
			return nil, err
		}
		result.PersonIds = append(result.PersonIds, personId)
	}
	result.Date, _ = time.Parse("2006-01-02", date)
	return &result, nil
}

func (r *EventRepository) GetByName(name string) (*domain.Event, error) {
	var result domain.Event
	var date string
	err := r.Db.DBPGX.QueryRow(context.Background(),
		`SELECT * FROM events WHERE name=$1`, name).Scan(&result.Id, &result.Name, &date)
	if err != nil {
		logger.Logger.Error("Failed Scan data from 'events' by name: ", zap.Error(err))
		return nil, err
	}
	rows, err := r.Db.DBPGX.Query(context.Background(),
		`SELECT * FROM persons_events WHERE event_id=$1`, result.Id)
	if err != nil {
		logger.Logger.Error("Failed take 'persons_ids' from 'events' table by name: ", zap.Error(err))
		return nil, err
	}
	for rows.Next() {
		var personId int64
		err = rows.Scan(&personId)
		if err != nil {
			logger.Logger.Error("Failed scan 'person_id' from 'events' rows: ", zap.Error(err))
			return nil, err
		}
		result.PersonIds = append(result.PersonIds, personId)
	}
	result.Date, _ = time.Parse("2006-01-02", date)
	return &result, nil
}

func (r *EventRepository) Update(ev *domain.Event) error {
	_, err := r.Db.DBPGX.Exec(context.Background(),
		`UPDATE events SET name=$2, date=$3 WHERE id=$1`, ev.Id, ev.Name, ev.Date.Format("2006-01-02"))
	if err != nil {
		logger.Logger.Error("Failed Update in 'events' table: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *EventRepository) DeleteById(id int64) error {
	_, err := r.Db.DBPGX.Exec(context.Background(),
		`DELETE FROM events WHERE id=$1`, id)
	if err != nil {
		logger.Logger.Error("Failed Delete in 'events' table: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *EventRepository) DeleteByName(name string) error {
	_, err := r.Db.DBPGX.Exec(context.Background(),
		`DELETE FROM events WHERE name=$1`, name)
	if err != nil {
		logger.Logger.Error("Failed Delete in 'events' table: ", zap.Error(err))
		return err
	}
	return nil
}
