package service

import (
	"context"
	"math"
	"sort"
	"time"
)

type EventData struct {
<<<<<<< HEAD
	Name           string    `json:"event_name" default:"Some Event name"`
	Date           time.Time `json:"event_date" default:"2020-12-31"`
	AverageSpent   float64   `json:"average_spent" default:"33.33"`
	TotalSpent     float64   `json:"total_spent" default:"100"`
	AllPeronsCount int      `json:"all_persons_count" default:"3"`
	RoundRate      float64   `json:"round_rate" default:"0.01"`
	Owes           map[string]map[string]float64
=======
	Name      string
	Date      time.Time
	Average   float64
	Total     float64
	Count     int
	RoundRate float64
	Owes      map[string]map[string]float64
>>>>>>> 9a4ae566fa2b8141fcb08531d73bacef0f5edb34
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

func (s *CalcService) createEventData(ctx context.Context, eventId int64, roundFactor float64) (EventData, []balance, error) {
	var result EventData
	perEvsArr, _ := s.PersonsEventsService.GetByEventId(ctx, eventId)
	result.Name = perEvsArr[0].Event.Name
	result.Date = perEvsArr[0].Event.Date
	result.RoundRate = roundFactor
	for _, pe := range perEvsArr {
		result.TotalSpent += pe.Spent
		result.AllPeronsCount += pe.Factor
	}
	result.AverageSpent = result.TotalSpent / float64(result.AllPeronsCount)
	var balances []balance
	for _, pe := range perEvsArr {
		balances = append(balances,
			balance{
				perId:   pe.PersonId,
				perName: pe.Person.Name,
<<<<<<< HEAD
				balance: pe.Spent - result.AverageSpent*float64(pe.Factor)})
=======
				balance: pe.Spent - result.Average*float64(pe.Factor)})
>>>>>>> 9a4ae566fa2b8141fcb08531d73bacef0f5edb34
	}
	sort.SliceStable(balances, func(i, j int) bool {
		return balances[i].balance < balances[j].balance
	})
	return result, balances, nil
}

func roundAndAbs(numb, roundFactor float64) float64 {
	return math.Abs(math.Round(numb/roundFactor) * roundFactor)
}

func (r *EventData) calculateBalances(balances []balance) {
	// i = most indepted Person, j = most portable Person
	for i, j := 0, len(balances)-1; i < j; {
		switch {
		// if Balance of 'i' great them 'j' and the it's left to next 'j+1' Person
		case balances[i].balance+balances[j].balance < 0:
			if r.Owes == nil {
				r.Owes = make(map[string]map[string]float64)
			}
			if r.Owes[balances[i].perName] == nil {
				r.Owes[balances[i].perName] = make(map[string]float64)
			}
			r.Owes[balances[i].perName][balances[j].perName] = roundAndAbs(balances[j].balance, r.RoundRate) // math.Abs(balances[j].balance)
			balances[i].balance += balances[j].balance
			balances[j].balance = 0
			j--
		// if Balance of 'i' less them 'j' and 'j' should take from 'i+1' Person
		case balances[i].balance+balances[j].balance >= 0 &&
			(balances[i].balance != 0 && balances[j].balance != 0):
			if r.Owes == nil {
				r.Owes = make(map[string]map[string]float64)
			}
			if r.Owes[balances[i].perName] == nil {
				r.Owes[balances[i].perName] = make(map[string]float64)
			}
			r.Owes[balances[i].perName][balances[j].perName] = roundAndAbs(balances[i].balance, r.RoundRate) //math.Abs(balances[i].balance)
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

func (s *CalcService) CalculateEvent(ctx context.Context, eventId int64, roundFactor float64) (EventData, error) {
	result, balances, _ := s.createEventData(ctx, eventId, roundFactor)
	result.calculateBalances(balances)
	return result, nil
}
