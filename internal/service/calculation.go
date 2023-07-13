package service

import (
	"context"
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
	EventsService        *EventService
	PersonsEventsService *PersonsEventsService
}

func NewCalcService(es *EventService, pes *PersonsEventsService) *CalcService {
	return &CalcService{
		EventsService:        es,
		PersonsEventsService: pes,
	}
}

// Create EventData and fill all fields but Balances
func (s *CalcService) createEventData(ctx context.Context, id int64) (*EventData, error) {
	var result EventData
	event, err := s.EventsService.GetEventById(ctx, id)
	if err != nil {
		return nil, err
	}
	result.Name = event.Name
	result.Date = event.Date
	for _, p := range event.Persons {
		perEv, err := s.PersonsEventsService.GetByPersonId(ctx, p.Id)
		if err != nil {
			return nil, err
		}
		var perData PersonData
		perData.Name = p.Name
		perData.Spent = perEv.Spent
		perData.Factor = perEv.Factor
		result.Persons = append(result.Persons, perData)
		result.AllPersonsCount += perEv.Factor
		result.TotalAmount += perEv.Spent
	}
	result.AverageAmount = result.TotalAmount / float64(result.AllPersonsCount)
	return &result, nil
}

// Fill Balances of Persons by them spents (taking into average) and
// sorting from most indebted to most portable.
func (ed *EventData) fillAndSortBalances() {
	for i := 0; i < len(ed.Persons); i++ {
		ed.Balances = append(ed.Balances, PersonBalance{
			Person:  &ed.Persons[i],
			Balance: ed.Persons[i].Spent - ed.AverageAmount*float64(ed.Persons[i].Factor),
		})
	}
	sort.SliceStable(ed.Balances, func(i, j int) bool {
		return ed.Balances[i].Balance < ed.Balances[j].Balance
	})
}

// Calculate owes of indepted to portable Persons by them Balances.
// Calculation continues until all Balances are equal to zero.
func (ev *EventData) calculateOwes() {
	// i = most indepted Person, j = most portable Person
	for i, j := 0, len(ev.Balances)-1; i < j; {
		switch {
		// if Balance of 'i' great them 'j' and the it's left to next 'j+1' Person
		case ev.Balances[i].Balance+ev.Balances[j].Balance > 0:
			if ev.Balances[i].Person.Owe == nil {
				ev.Balances[i].Person.Owe = map[string]float64{}
			}
			ev.Balances[i].Person.Owe[ev.Balances[j].Person.Name] = -ev.Balances[i].Balance
			ev.Balances[j].Balance += ev.Balances[i].Balance
			ev.Balances[i].Balance = 0
			i++
			// if Balance of 'i' less them 'j' and 'j' should take from 'i+1' Person
		case ev.Balances[i].Balance+ev.Balances[j].Balance <= 0 &&
			(ev.Balances[i].Balance != 0 && ev.Balances[j].Balance != 0):
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

func (s *CalcService) CalcEvent(ctx context.Context, id int64) (*EventData, error) {
	ed, err := s.createEventData(ctx, id)
	if err != nil {
		return nil, err
	}
	ed.fillAndSortBalances()
	ed.calculateOwes()
	return ed, nil
}
