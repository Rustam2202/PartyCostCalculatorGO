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
		//	err := r.Db.DB.QueryRow(`INSERT INTO events (name, date) VALUES($1,$2) RETURNING Id`,
		ev.Name, ev.Date.Format("2006-01-02")).Scan(&lastInsertedId)
	if err != nil {
		logger.Logger.Error("Couldn`t execute Insert operation", zap.Error(err))
		return err
	}
	ev.Id = lastInsertedId
	return nil
}

func (r *EventRepository) GetById(id int64) (*domain.Event, error) {
	var result domain.Event
	var date string
	err := r.Db.DBPGX.QueryRow(context.Background(),
		`SELECT * FROM events WHERE id=$1`, id).
		// err := r.Db.DB.QueryRow(`SELECT * FROM events WHERE id=$1`, id).
		Scan(&result.Id, &result.Name, &date)
	if err != nil {
		logger.Logger.Error("Failed to Scan data from events:", zap.Error(err))
		return nil, err
	}
	result.Date, _ = time.Parse("2006-01-02", date)
	return &result, nil
}

func (r *EventRepository) GetByName(name string) (*domain.Event, error) {
	var result domain.Event
	var date string
	err := r.Db.DBPGX.QueryRow(context.Background(),
		`SELECT * FROM events WHERE name=$1`, name).
		//	err := r.Db.DB.QueryRow(`SELECT * FROM events WHERE name=$1`, name).
		Scan(&result.Id, &result.Name, &date)
	if err != nil {
		logger.Logger.Error("Failed to Scan data from events:", zap.Error(err))
		return nil, err
	}
	result.Date, _ = time.Parse("2006-01-02", date)
	return &result, nil
}

func (r *EventRepository) Update(ev *domain.Event) error {
	_, err := r.Db.DBPGX.Exec(context.Background(),
		`UPDATE events SET name=$2, date=$3 WHERE id=$1`,
		ev.Id, ev.Name, ev.Date.Format("2006-01-02"))
	// _, err := r.Db.DB.Exec(`UPDATE events SET name=$2, date=$3 WHERE id=$1`,
	// 	ev.Id, ev.Name, ev.Date.Format("2006-01-02"))
	if err != nil {
		logger.Logger.Error("Failed to Execute Update operation: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *EventRepository) DeleteById(id int64) error {
	_, err := r.Db.DBPGX.Exec(context.Background(),
		`DELETE FROM events WHERE id=$1`, id)
	//	_, err := r.Db.DB.Exec(`DELETE FROM events WHERE id=$1`, id)
	if err != nil {
		logger.Logger.Error("Failed to Execute Delete operation: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *EventRepository) DeleteByName(name string) error {
	_, err := r.Db.DBPGX.Exec(context.Background(),
		`DELETE FROM events WHERE name=$1`, name)
	//	_, err := r.Db.DB.Exec(`DELETE FROM events WHERE name=$1`, name)
	if err != nil {
		logger.Logger.Error("Failed to Execute Delete operation: ", zap.Error(err))
		return err
	}
	return nil
}
