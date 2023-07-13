package service

import (
	"context"
	"party-calc/internal/domain"
	"time"
)

type EventRepository interface {
	Create(ctx context.Context, ev *domain.Event) error
	GetById(ctx context.Context, id int64) (*domain.Event, error)
	GetByName(ctx context.Context, name string) (*domain.Event, error)
	Update(ctx context.Context, ev *domain.Event) error
	DeleteById(ctx context.Context, id int64) error
	DeleteByName(ctx context.Context, name string) error
}

type EventService struct {
	repo EventRepository
}

func NewEventService(r EventRepository) *EventService {
	return &EventService{repo: r}
}

func (p *EventService) NewEvent(ctx context.Context, name string, date time.Time) (int64, error) {
	ev := domain.Event{Name: name, Date: date}
	err := p.repo.Create(ctx, &ev)
	if err != nil {
		return 0, err
	}
	return ev.Id, nil
}

func (p *EventService) GetEventById(ctx context.Context, id int64) (*domain.Event, error) {
	ev, err := p.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return ev, nil
}

func (p *EventService) GetEventByName(ctx context.Context, name string) (*domain.Event, error) {
	ev, err := p.repo.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return ev, nil
}

func (p *EventService) UpdateEvent(ctx context.Context, id int64, name string, date time.Time) error {
	err := p.repo.Update(ctx, &domain.Event{Id: id, Name: name, Date: date})
	if err != nil {
		return err
	}
	return nil
}

func (p *EventService) DeleteEventById(ctx context.Context, id int64) error {
	err := p.repo.DeleteById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *EventService) DeleteEventByName(ctx context.Context, name string) error {
	err := p.repo.DeleteByName(ctx, name)
	if err != nil {
		return err
	}
	return nil
}
