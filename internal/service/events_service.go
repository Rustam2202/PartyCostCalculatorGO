package service

import (
	"party-calc/internal/domain"
	"time"
)

type EventRepository interface {
	Add(ev *domain.Event) error
	GetById(id int64) (*domain.Event, error)
	GetByName(name string) (*domain.Event, error)
	Update(ev *domain.Event) error
	DeleteById(id int64) error
	DeleteByName(name string) error
}

type EventService struct {
	repo EventRepository
}

func NewEventService(r EventRepository) *EventService {
	return &EventService{repo: r}
}

func (p *EventService) NewEvent(name string, date time.Time) error {
	err := p.repo.Add(&domain.Event{Name: name, Date: date})
	if err != nil {
		return err
	}
	return nil
}

func (p *EventService) GetEventById(id int64) (*domain.Event, error) {
	ev, err := p.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return ev, nil
}

func (p *EventService) GetEventByName(name string) (*domain.Event, error) {
	ev, err := p.repo.GetByName(name)
	if err != nil {
		return nil, err
	}
	return ev, nil
}

func (p *EventService) UpdateEvent(id int64, name string, date time.Time) error {
	err := p.repo.Update(&domain.Event{Id: id, Name: name, Date: date})
	if err != nil {
		return err
	}
	return nil
}

func (p *EventService) DeleteEventById(id int64) error {
	err := p.repo.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}

func (p *EventService) DeleteEventByName(name string) error {
	err := p.repo.DeleteByName(name)
	if err != nil {
		return err
	}
	return nil
}
