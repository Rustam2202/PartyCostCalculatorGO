package service

import (
	"party-calc/internal/database/models"
	"party-calc/internal/database/repository"
	"time"
)

type EventService struct {
	repo *repository.EventRepository
}

func NewEventService(r *repository.EventRepository) *EventService {
	return &EventService{repo: r}
}

func (p *EventService) NewEvent(name string, date time.Time) (int64, error) {
	id, err := p.repo.Add(&models.Event{Name: name, Date: date})
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *EventService) GetEvent(name string) (models.Event, error) {
	ev, err := p.repo.Get(&models.Event{Name: name})
	if err != nil {
		return models.Event{}, err
	}
	return ev, nil
}

func (p *EventService) UpdateEvent(nameOld, nameNew string, dateNew time.Time) error {
	evOld, err := p.GetEvent(nameOld)
	if err != nil {
		return err
	}
	err = p.repo.Update(&evOld, &models.Event{Name: nameNew, Date: dateNew})
	if err != nil {
		return err
	}
	return nil
}

func (p *EventService) DeleteEvent(name string) error {
	err := p.repo.Delete(&models.Event{Name: name})
	if err != nil {
		return err
	}
	return nil
}
