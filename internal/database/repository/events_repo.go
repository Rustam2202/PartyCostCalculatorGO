package repository

import (
	"party-calc/internal/database"
	"party-calc/internal/database/models"
	"party-calc/internal/logger"
	"time"

	"go.uber.org/zap"
)

type EventRepository struct {
	db *database.DataBase
}

func NewEventRepository(db *database.DataBase) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) Add(ev *models.Event) (int64, error) {
	var lastInsertedId int64
	err := r.db.DB.QueryRow(`INSERT INTO events (name, date) VALUES($1,$2) RETURNING Id`,
		ev.Name, ev.Date.Format("2006-01-02")).Scan(&lastInsertedId)
	if err != nil {
		logger.Logger.Error("Couldn`t execute Insert operation", zap.Error(err))
		return 0, err
	}
	return lastInsertedId, nil
}

func (r *EventRepository) Get(ev *models.Event) (models.Event, error) {
	var result models.Event
	var date string
	err := r.db.DB.QueryRow(`SELECT * FROM events WHERE name=$1`, ev.Name).
		Scan(&result.Id, &result.Name, &date, &result.TotalAmount)
	if err != nil {
		logger.Logger.Error("Failed to Scan data from events:", zap.Error(err))
		return models.Event{}, err
	}
	result.Date, _ = time.Parse("2006-01-02", date)
	return result, nil
}

func (r *EventRepository) Update(evOld, evNew *models.Event) error {
	_, err := r.db.DB.Exec(`UPDATE events SET name=$2, date=$3 WHERE id=$1`,
		evOld.Id, evNew.Name, evNew.Date.Format("2006-01-02"))
	if err != nil {
		logger.Logger.Error("Failed to Execute Update operation: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *EventRepository) Delete(ev *models.Event) error {
	_, err := r.db.DB.Exec(`DELETE FROM events WHERE name=$1`, ev.Name)
	if err != nil {
		logger.Logger.Error("Failed to Execute Delete operation: ", zap.Error(err))
		return err
	}
	return nil
}
