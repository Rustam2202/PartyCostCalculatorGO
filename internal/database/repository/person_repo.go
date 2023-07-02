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

func (r *PersonRepository) Create(per *domain.Person) error {
	var lastInsertedId int64
	err := r.Db.DBPGX.QueryRow(context.Background(),
		`INSERT INTO persons (name) VALUES($1) RETURNING Id`, per.Name).Scan(&lastInsertedId)
	if err != nil {
		logger.Logger.Error("Failed Insert to 'persons' table: ", zap.Error(err))
		return err
	}
	per.Id = lastInsertedId
	return nil
}

func (r *PersonRepository) GetById(id int64) (*domain.Person, error) {
	var result domain.Person
	err := r.Db.DBPGX.QueryRow(context.Background(),
		`SELECT * FROM persons WHERE id=$1`, id).Scan(&result.Id, &result.Name)
	if err != nil {
		logger.Logger.Error("Failed Scan data from 'persons' by id: ", zap.Error(err))
		return nil, err
	}
	rows, err := r.Db.DBPGX.Query(context.Background(),
		`SELECT event_id FROM persons_events WHERE person_id=$1`, id)
	if err != nil {
		logger.Logger.Error("Failed take 'event_ids' from 'persons' table by id: ", zap.Error(err))
		return nil, err
	}
	for rows.Next() {
		var eventId int64
		err = rows.Scan(&eventId)
		if err != nil {
			logger.Logger.Error("Failed scan 'event_id' from 'event_ids' rows: ", zap.Error(err))
			return nil, err
		}
		result.EventIds = append(result.EventIds, eventId)
	}
	return &result, nil
}

func (r *PersonRepository) GetByName(name string) (*domain.Person, error) {
	var result domain.Person
	err := r.Db.DBPGX.QueryRow(context.Background(),
		`SELECT * FROM persons WHERE name=$1`, name).Scan(&result.Id, &result.Name)
	if err != nil {
		logger.Logger.Error("Failed take 'event_ids' from 'persons' table by name: ", zap.Error(err))
		return nil, err
	}
	rows, err := r.Db.DBPGX.Query(context.Background(),
		`SELECT event_id FROM persons_events WHERE person_id=$1`, result.Id)
	if err != nil {
		logger.Logger.Error("Failed take 'event_ids' from 'persons' by name: ", zap.Error(err))
		return nil, err
	}
	for rows.Next() {
		var eventId int64
		err = rows.Scan(&eventId)
		if err != nil {
			logger.Logger.Error("Failed scan 'event_id' from 'event_ids' rows: ", zap.Error(err))
			return nil, err
		}
		result.EventIds = append(result.EventIds, eventId)
	}
	return &result, nil
}

func (r *PersonRepository) Update(per *domain.Person) error {
	_, err := r.Db.DBPGX.Exec(context.Background(),
		`UPDATE persons SET name=$2 WHERE id=$1`, per.Id, per.Name)
	if err != nil {
		logger.Logger.Error("Failed Update in 'persons' table: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *PersonRepository) DeleteById(id int64) error {
	_, err := r.Db.DBPGX.Exec(context.Background(),
		`DELETE FROM persons WHERE id=$1`, id)
	if err != nil {
		logger.Logger.Error("Failed Delete in 'persons' table: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *PersonRepository) DeleteByName(name string) error {
	_, err := r.Db.DBPGX.Exec(context.Background(),
		`DELETE FROM persons WHERE name=$1`, name)
	if err != nil {
		logger.Logger.Error("Failed Delete in 'persons' table: ", zap.Error(err))
		return err
	}
	return nil
}
