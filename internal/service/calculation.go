package service

import (
	"context"
	"math"
	"sort"
	"time"
)

// type (
// 	debetor   string
// 	recepient string
// 	owes      map[recepient]float64
// )

type EventData struct {
	Name         string
	Date         time.Time
	AverageSpent float64
	TotalSpent   float64
	PersonsCount int
	//PersonsOwes    map[debetor]owes
	PersonsOwes map[string]map[string]float64 // map[debetor]map[recepient]float64
}

type balance struct {
	perId   int64
	perName string
	balance float64
}

type CalcService struct {
	PersonService        *PersonService
	EventsService        *EventService
	PersonsEventsService *PersonsEventsService
}

func NewCalcService(ps *PersonService, es *EventService, pes *PersonsEventsService) *CalcService {
	return &CalcService{
		PersonService:        ps,
		EventsService:        es,
		PersonsEventsService: pes,
	}
}

func (s *CalcService) createResponse(ctx context.Context, eventId int64) (EventData, []balance, error) {
	var result EventData
	perEvsArr, _ := s.PersonsEventsService.GetByEventId(ctx, eventId)
	result.Name = perEvsArr[0].Event.Name
	result.Date = perEvsArr[0].Event.Date
	for _, pe := range perEvsArr {
		result.TotalSpent += pe.Spent
		result.PersonsCount += pe.Factor
	}
	result.AverageSpent = result.TotalSpent / float64(result.PersonsCount)
	var balances []balance
	for _, pe := range perEvsArr {
		balances = append(balances,
			balance{perId: pe.PersonId, perName: pe.Person.Name, balance: pe.Spent - result.AverageSpent*float64(pe.Factor)})
	}
	sort.SliceStable(balances, func(i, j int) bool {
		return balances[i].balance < balances[j].balance
	})
	return result, balances, nil
}

func (r *EventData) calculateBalances(balances []balance) {
	// i = most indepted Person, j = most portable Person
	for i, j := 0, len(balances)-1; i < j; {
		switch {
		// if Balance of 'i' great them 'j' and the it's left to next 'j+1' Person
		case balances[i].balance+balances[j].balance < 0:
			if r.PersonsOwes[balances[i].perName] == nil {
				r.PersonsOwes = make(map[string]map[string]float64)
			}
			if r.PersonsOwes[balances[i].perName] == nil {
				r.PersonsOwes[balances[i].perName] = make(map[string]float64)
			}
			r.PersonsOwes[balances[i].perName][balances[j].perName] = math.Abs(balances[j].balance)
			balances[i].balance += balances[j].balance
			balances[j].balance = 0
			j--
		// if Balance of 'i' less them 'j' and 'j' should take from 'i+1' Person
		case balances[i].balance+balances[j].balance >= 0 &&
			(balances[i].balance != 0 && balances[j].balance != 0):
			if r.PersonsOwes[balances[i].perName] == nil {
				r.PersonsOwes = make(map[string]map[string]float64)
			}
			if r.PersonsOwes[balances[i].perName] == nil {
				r.PersonsOwes[balances[i].perName] = make(map[string]float64)
			}
			r.PersonsOwes[balances[i].perName][balances[j].perName] = math.Abs(balances[i].balance)
			balances[j].balance += balances[i].balance
			balances[i].balance = 0
			i++
		case balances[i].balance == 0:
			i++
		case balances[j].balance == 0:
			j--
		}
	}
}

func (s *CalcService) CalculateEvent(ctx context.Context, eventId int64) (EventData, error) {
	result, balances, _ := s.createResponse(ctx, eventId)
	result.calculateBalances(balances)
	return result, nil
}
