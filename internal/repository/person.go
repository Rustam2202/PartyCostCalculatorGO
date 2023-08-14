package repository

import (
	"context"
	"party-calc/internal/database"
	"party-calc/internal/domain"
	"party-calc/internal/logger"

	"go.uber.org/zap"
)

type PersonRepository struct {
	Db *database.DataBase
}

func NewPersonRepository(db *database.DataBase) *PersonRepository {
	return &PersonRepository{Db: db}
}

func (r *PersonRepository) Create(ctx context.Context, per *domain.Person) error {
	var lastInsertedId int64
	err := r.Db.DBPGX.QueryRow(ctx,
		`INSERT INTO persons (name) VALUES($1) RETURNING Id`, per.Name).Scan(&lastInsertedId)
	if err != nil {
		logger.Logger.Error("Failed Insert to 'persons' table: ", zap.Error(err))
		return err
	}
	per.Id = lastInsertedId
	return nil
}

func (r *PersonRepository) GetById(ctx context.Context, id int64) (*domain.Person, error) {
	var result domain.Person
	err := r.Db.DBPGX.QueryRow(ctx,
		`SELECT * FROM persons WHERE id=$1`, id).Scan(&result.Id, &result.Name)
	if err != nil {
		logger.Logger.Error("Failed Scan data from 'persons' by id: ", zap.Error(err))
		return nil, err
	}

	// search events ids of person
	rows, err := r.Db.DBPGX.Query(ctx,
		`SELECT event_id FROM persons_events WHERE person_id=$1`, result.Id)
	if err != nil {
		logger.Logger.Error("Failed take 'event_ids' from 'persons' table by id: ", zap.Error(err))
		return nil, err
	}
	var eventIds []int64
	for rows.Next() {
		var evId int64
		err = rows.Scan(&evId)
		if err != nil {
			return nil, err
		}
		eventIds = append(eventIds, evId)
	}

	// append events to Person model
	rows, err = r.Db.DBPGX.Query(ctx,
		`SELECT id, name, date FROM events WHERE id=ANY($1)`, eventIds)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var event domain.Event
		err = rows.Scan(&event.Id, &event.Name, &event.Date)
		if err != nil {
			logger.Logger.Error("Failed scan 'event_id' from 'event_ids' rows: ", zap.Error(err))
			return nil, err
		}
		result.Events = append(result.Events, event)
	}
	return &result, nil
}

func (r *PersonRepository) GetByName(ctx context.Context, name string) (*domain.Person, error) {
	var result domain.Person
	err := r.Db.DBPGX.QueryRow(ctx,
		`SELECT * FROM persons WHERE name=$1`, name).Scan(&result.Id, &result.Name)
	if err != nil {
		logger.Logger.Error("Failed Scan data from 'persons' by name: ", zap.Error(err))
		return nil, err
	}

	// search events ids of person
	rows, err := r.Db.DBPGX.Query(ctx,
		`SELECT event_id FROM persons_events WHERE person_id=$1`, result.Id)
	if err != nil {
		logger.Logger.Error("Failed take 'event_ids' from 'persons' table by id: ", zap.Error(err))
		return nil, err
	}
	var eventIds []int64
	for rows.Next() {
		var evId int64
		err = rows.Scan(&evId)
		if err != nil {
			return nil, err
		}
		eventIds = append(eventIds, evId)
	}

	// append events to Person model
	rows, err = r.Db.DBPGX.Query(ctx,
		`SELECT id, name, date FROM events WHERE id=ANY($1)`, eventIds)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var event domain.Event
		err = rows.Scan(&event.Id, &event.Name, &event.Date)
		if err != nil {
			logger.Logger.Error("Failed scan 'event_id' from 'event_ids' rows: ", zap.Error(err))
			return nil, err
		}
		result.Events = append(result.Events, event)
	}
	return &result, nil
}

func (r *PersonRepository) Update(ctx context.Context, per *domain.Person) error {
	_, err := r.Db.DBPGX.Exec(ctx,
		`UPDATE persons SET name=$2 WHERE id=$1`, per.Id, per.Name)
	if err != nil {
		logger.Logger.Error("Failed Update in 'persons' table: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *PersonRepository) DeleteById(ctx context.Context, id int64) error {
	_, err := r.Db.DBPGX.Exec(ctx,
		`DELETE FROM persons WHERE id=$1`, id)
	if err != nil {
		logger.Logger.Error("Failed Delete in 'persons' table: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *PersonRepository) DeleteByName(ctx context.Context, name string) error {
	_, err := r.Db.DBPGX.Exec(ctx,
		`DELETE FROM persons WHERE name=$1`, name)
	if err != nil {
		logger.Logger.Error("Failed Delete in 'persons' table: ", zap.Error(err))
		return err
	}
	return nil
}
