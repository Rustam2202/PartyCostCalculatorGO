package service

import (
	"context"
	"party-calc/internal/domain"
)

type PersonsEventsRepository interface {
	Create(ctx context.Context, pe *domain.PersonsAndEvents) error
	GetByPersonId(ctx context.Context, id int64) (*domain.PersonsAndEvents, error)
	GetByEventId(ctx context.Context, id int64) (*domain.PersonsAndEvents, error)
	Update(ctx context.Context, pe *domain.PersonsAndEvents) error
	Delete(ctx context.Context, id int64) error
}

type PersonsEventsService struct {
	repo PersonsEventsRepository
}

func NewPersonsEventsService(r PersonsEventsRepository) *PersonsEventsService {
	return &PersonsEventsService{repo: r}
}

func (p *PersonsEventsService) AddPersonToPersEvent(ctx context.Context, personId, eventId int64, spent float64, factor int) (int64, error) {
	perEv := domain.PersonsAndEvents{
		PersonId: personId,
		EventId:  eventId,
		Spent:    spent,
		Factor:   factor,
	}
	err := p.repo.Create(ctx, &perEv)
	if err != nil {
		return 0, err
	}
	return perEv.Id, nil
}

func (p *PersonsEventsService) GetByPersonId(ctx context.Context, id int64) (*domain.PersonsAndEvents, error) {
	result, err := p.repo.GetByPersonId(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *PersonsEventsService) GetByEventId(ctx context.Context, id int64) (*domain.PersonsAndEvents, error) {
	result, err := p.repo.GetByEventId(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *PersonsEventsService) Update(ctx context.Context, id, personId, eventId int64, spent float64, factor int) error {
	err := p.repo.Update(ctx, &domain.PersonsAndEvents{
		Id:       id,
		PersonId: personId,
		EventId:  eventId,
		Spent:    spent,
		Factor:   factor,
	})
	if err != nil {
		return err
	}
	return nil
}

func (p *PersonsEventsService) Delete(ctx context.Context, id int64) error {
	err := p.repo.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
