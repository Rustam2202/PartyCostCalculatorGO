package service

import (
	"party-calc/internal/database/models"
	"party-calc/internal/database/repository"
)

type PersEventsService struct {
	repo *repository.PersEventsRepository
}

func NewPersEventsService(r *repository.PersEventsRepository) *PersEventsService {
	return &PersEventsService{repo: r}
}

func (p *PersEventsService) AddPersonToPersEvents(perName, evName string, spent float64, factor int) (int64, error) {
	per, err := p.repo.PersRepo.Get(&models.Person{Name: perName})
	if err != nil {
		return 0, err
	}
	ev, err := p.repo.EventsRepo.Get(&models.Event{Name: evName})
	if err != nil {
		return 0, err
	}
	id, err := p.repo.Create(&models.PersonsAndEvents{
		PersonId: per.Id,
		EventId:  ev.Id,
		Spent:    spent,
		Factor:   factor,
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (p *PersEventsService) GetPersonFromPersEvents(perName string) (models.PersonsAndEvents, error) {
	per, err := p.repo.PersRepo.Get(&models.Person{Name: perName})
	if err != nil {
		return models.PersonsAndEvents{}, err
	}
	result, err := p.repo.Get(&models.PersonsAndEvents{Id: per.Id})
	if err != nil {
		return models.PersonsAndEvents{}, err
	}
	return result, nil
}

func (p *PersEventsService) UpdatePerson(perName, newEventName string, newSpent float64, newFactor int) error {
	perOld, err := p.repo.PersRepo.Get(&models.Person{Name: perName})
	if err != nil {
		return err
	}
	evId, err := p.repo.EventsRepo.Get(&models.Event{Name: newEventName})
	if err != nil {
		return err
	}
	err = p.repo.Update(
		&models.PersonsAndEvents{PersonId: perOld.Id},
		&models.PersonsAndEvents{PersonId: perOld.Id, EventId: evId.Id, Spent: newSpent, Factor: newFactor})
	if err != nil {
		return err
	}
	return nil
}

func (p *PersEventsService) DeletePerson(perName string) error {
	per, err := p.repo.PersRepo.Get(&models.Person{Name: perName})
	if err != nil {
		return err
	}
	err = p.repo.Delete(&models.PersonsAndEvents{PersonId: per.Id})
	if err != nil {
		return err
	}
	return nil
}
