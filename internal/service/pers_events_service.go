package service

import (
	"errors"
	"party-calc/internal/database/models"
	"party-calc/internal/database/repository"
)

type PersEventsService struct {
	repo *repository.PersEventsRepository
}

func NewPersEventsService(r *repository.PersEventsRepository) *PersEventsService {
	return &PersEventsService{repo: r}
}

func (p *PersEventsService) AddPersonToPersEvent(perName, evName string, spent float64, factor int) (int64, error) {
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

func (p *PersEventsService) GetPerson(perName string) (PersonData, error) {
	per, err := p.repo.PersRepo.Get(&models.Person{Name: perName})
	if err != nil {
		return PersonData{}, err
	}
	//perFromPersEvents, err := p.repo.GetPerson(&models.PersonsAndEvents{Id: per.Id})
	result := PersonData{
		Id:   per.Id,
		Name: per.Name,
		//Owe:  map[string]float32{},
	}
	if err != nil {
		return PersonData{}, err
	}
	return result, nil
}

func (p *PersEventsService) GetPersonsByEvent(eventId int64) (EventData, error) {
	// persEvents by Event-Id
	persEvSlice, err := p.repo.GetEvent(&models.PersonsAndEvents{Id: eventId})
	if err != nil {
		return EventData{}, err
	}
	var result EventData
	// get event Name and Data 
	if len(persEvSlice) != 0 {
		ev, err := p.repo.EventsRepo.Get(&models.Event{Id: persEvSlice[0].EventId})
		if err != nil {
			return EventData{}, err
		}
		result.Name = ev.Name
		result.Date = ev.Date
	} else {
		return EventData{}, errors.New("not Persons in Event")
	}
	// add Persons to EventData
	for _, pe := range persEvSlice {
		per, err := p.repo.PersRepo.Get(&models.Person{Id: pe.PersonId})
		if err != nil {
			return EventData{}, err
		}
		pd := PersonData{
			Id:   per.Id,
			Name: per.Name,
			Spent: pe.Spent,
			//Owe:  map[string]float32{},
		}
		// add other fields of EventData 
		result.Owes = append(result.Owes, pd)
		result.AllPersonsCount += pe.Factor
		result.TotalAmount += pe.Spent
	}
	result.AverageAmount = result.TotalAmount / float64(result.AllPersonsCount)
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
