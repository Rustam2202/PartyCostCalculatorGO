package service

import (
	"sort"
	"time"
)

type PersonData struct {
	Id     int64
	Name   string
	Spent  float64
	Factor int
	Owe    map[string]float64
}

type PersonBalance struct {
	Person  *PersonData
	Balance float64
}

type EventData struct {
	Name            string
	Date            time.Time
	Persons         []PersonData
	AllPersonsCount int
	AverageAmount   float64
	TotalAmount     float64
	Balances        []PersonBalance
}

type CalcService struct {
	//repo                 *repository.PersEventsRepository
	PersonsService       *PersonService
	EventsService        *EventService
	PersonsEventsService *PersonsEventsService
}

func NewCalcService(ps *PersonService, es *EventService, pes *PersonsEventsService) *CalcService {
	return &CalcService{
		PersonsService:       ps,
		EventsService:        es,
		PersonsEventsService: pes,
	}
}

func (s *CalcService) createEventData(eventName string) (*EventData, error) {
	var result EventData
	event, _ := s.EventsService.GetEventByName(eventName)
	result.Name = event.Name
	result.Date = event.Date
	for _, personId := range event.PersonIds {
		var perData PersonData
		per, _ := s.PersonsService.GetPersonById(personId)
		perEv, _ := s.PersonsEventsService.GetByPersonId(personId)
		perData.Name = per.Name
		perData.Spent = perEv.Spent
		perData.Factor = perEv.Factor
		result.Persons = append(result.Persons, perData)
	}
	return &result, nil
}

func (ed *EventData) fillAndSortBalances() {
	for i := 0; i < ed.AllPersonsCount-1; i++ {
		ed.Balances = append(ed.Balances, PersonBalance{
			Person:  &ed.Persons[i],
			Balance: ed.Persons[i].Spent - ed.AverageAmount*float64(ed.Persons[i].Factor),
		})
	}
	sort.SliceStable(ed.Balances, func(i, j int) bool {
		return ed.Balances[i].Balance < ed.Balances[j].Balance
	})
}

func (ev *EventData) calculateOwes() {
	for i, j := 0, len(ev.Balances)-1; i < j; {
		switch {
		case ev.Balances[i].Balance+ev.Balances[j].Balance > 0:
			if ev.Balances[i].Person.Owe == nil {
				ev.Balances[i].Person.Owe = map[string]float64{}
			}
			ev.Balances[i].Person.Owe[ev.Balances[j].Person.Name] = -ev.Balances[i].Balance
			ev.Balances[j].Balance += ev.Balances[i].Balance
			ev.Balances[i].Balance = 0
			i++
		case ev.Balances[i].Balance+ev.Balances[j].Balance <= 0:
			if ev.Balances[i].Person.Owe == nil {
				ev.Balances[i].Person.Owe = map[string]float64{}
			}
			ev.Balances[i].Person.Owe[ev.Balances[j].Person.Name] = ev.Balances[j].Balance
			ev.Balances[i].Balance += ev.Balances[j].Balance
			ev.Balances[j].Balance = 0
			j--
		case ev.Balances[i].Balance == 0:
			i++
		case ev.Balances[j].Balance == 0:
			j--
		}
	}
	ev.Balances = nil
}

func (s *CalcService) CalcPerson(perName, evName string) (PersonData, error) {
	
	return PersonData{}, nil
}

func (s *CalcService) CalcEvent(name string) (EventData, error) {
	ed, err := s.createEventData(name)
	if err != nil {
		return EventData{}, err
	}
	ed.fillAndSortBalances()
	ed.calculateOwes()
	return *ed, nil
}
