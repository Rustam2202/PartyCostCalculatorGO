package service

import (
	"context"
	"errors"
	"sort"
	"time"

	"github.com/shopspring/decimal"
)

type (
	recepients map[string]decimal.Decimal
	debtors    map[string]recepients
)

type EventData struct {
	Name         string          `json:"event_name" default:"Some Event name"`
	Date         time.Time       `json:"event_date" default:"2020-12-31"`
	AverageSpent decimal.Decimal `json:"average_spent" default:"33.33"`
	TotalSpent   decimal.Decimal `json:"total_spent" default:"100"`
	// AllPeronsCount exist persons mutiply them factors
	AllPeronsCount int32 `json:"all_persons_count" default:"3"`
	RoundRate      int32 `json:"round_rate" default:"2"`
	Owes           debtors
}

type balance struct {
	perId   int64
	perName string
	balance decimal.Decimal
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

func (s *CalcService) createEventData(ctx context.Context, eventId int64, RoundRate int32) (*EventData, []balance, error) {
	var result EventData
	perEvsArr, err := s.PersonsEventsService.GetByEventId(ctx, eventId)
	if err != nil {
		return nil, nil, err
	}
	if len(perEvsArr) < 3 {
		return nil, nil, errors.New("persons count less then 3")
	}
	result.Name = perEvsArr[0].Event.Name
	result.Date = perEvsArr[0].Event.Date
	result.RoundRate = RoundRate
	for _, pe := range perEvsArr {
		result.TotalSpent = decimal.Sum(result.TotalSpent, decimal.NewFromFloat(pe.Spent))
		result.AllPeronsCount += int32(pe.Factor)
	}
	result.AverageSpent = result.TotalSpent.Div(decimal.NewFromInt32(result.AllPeronsCount))
	var balances []balance
	for _, pe := range perEvsArr {
		balances = append(balances,
			balance{
				perId:   pe.PersonId,
				perName: pe.Person.Name,
				balance: decimal.NewFromFloat(pe.Spent).
					Sub(result.AverageSpent.
						Mul(decimal.NewFromInt32(int32(pe.Factor)))),
			},
		)
	}
	sort.SliceStable(balances, func(i, j int) bool {
		less := balances[i].balance.Cmp(balances[j].balance)
		if less == -1 {
			return true
		} else {
			return false
		}
	})
	return &result, balances, nil
}

func (r *EventData) calculateBalances(bls []balance) {
	// i-person with lowest balance and alway less or equal zero,
	// j-person with greatest balance and alway great or equal zero
	for i, j := 0, len(bls)-1; i < j; {
		switch {
		// if absolute i-balance great them j-balance (i-person can owe more then one j-persons)
		case decimal.Zero.Cmp(bls[i].balance.Add(bls[j].balance)) >= 0:
			if r.Owes == nil {
				r.Owes = make(debtors)
			}
			if r.Owes[bls[i].perName] == nil {
				r.Owes[bls[i].perName] = make(recepients)
			}
			r.Owes[bls[i].perName][bls[j].perName] = bls[j].balance.Abs()
			bls[i].balance = bls[i].balance.Add(bls[j].balance)
			bls[j].balance = decimal.Zero
			j--
		// if absolute i-balance less them j-balance (i-person couldn't owe j-persons more)
		case decimal.Zero.Cmp(bls[i].balance.Add(bls[j].balance)) < 0 &&
			!bls[i].balance.Equal(decimal.Zero) && !bls[j].balance.Equal(decimal.Zero):
			if r.Owes == nil {
				r.Owes = make(debtors)
			}
			if r.Owes[bls[i].perName] == nil {
				r.Owes[bls[i].perName] = make(recepients)
			}
			r.Owes[bls[i].perName][bls[j].perName] = bls[i].balance.Abs()
			bls[j].balance = bls[j].balance.Add(bls[i].balance)
			bls[i].balance = decimal.Zero
			i++
		case bls[i].balance.Equal(decimal.Zero):
			i++
		case bls[j].balance.Equal(decimal.Zero):
			j--
		}
	}
}

func (s *CalcService) CalculateEvent(ctx context.Context, eventId int64, roundFactor int32) (*EventData, error) {
	result, balances, err := s.createEventData(ctx, eventId, roundFactor)
	if err != nil {
		return nil, err
	}
	result.calculateBalances(balances)
	return result, nil
}
