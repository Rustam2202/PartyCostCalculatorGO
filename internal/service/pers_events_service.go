package service

import (
	"party-calc/internal/domain"
)

type PersonsEventsRepository interface {
	Create(pe *domain.PersonsAndEvents) error
	GetByPersonId(id int64) (*domain.PersonsAndEvents, error)
	GetByEventId(id int64) (*domain.PersonsAndEvents, error)
	Update(pe *domain.PersonsAndEvents) error
	Delete(id int64) error
}

type PersonsEventsService struct {
	repo PersonsEventsRepository
}

func NewPersonsEventsService(r PersonsEventsRepository) *PersonsEventsService {
	return &PersonsEventsService{repo: r}
}

func (p *PersonsEventsService) AddPersonToPersEvent(personId, eventId int64, spent float64, factor int) error {
	err := p.repo.Create(&domain.PersonsAndEvents{
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

func (p *PersonsEventsService) GetByPersonId(id int64) (*domain.PersonsAndEvents, error) {
	result, err := p.repo.GetByPersonId(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *PersonsEventsService) GetByEventId(id int64) (*domain.PersonsAndEvents, error) {
	result, err := p.repo.GetByEventId(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *PersonsEventsService) UpdatePerson(id, personId, eventId int64, spent float64, factor int) error {
	err := p.repo.Update(&domain.PersonsAndEvents{
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

func (p *PersonsEventsService) DeletePerson(id int64) error {
	err := p.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
