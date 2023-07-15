package service

import (
	"context"
	"sort"
	"time"
)

type Response struct {
	Name    string
	Date    time.Time
	Average float64
	Total   float64
	Count   int
	Owes    map[string]map[string]float64
}

type balance struct {
	perId   int64
	perName string
	balance float64
}

func (s *CalcService) createResponse(ctx context.Context, id int64) (Response, []balance, error) {
	var result Response
	perEvs, _ := s.PersonsEventsService.GetByEventId(ctx, id)
	result.Name = perEvs[0].Event.Name
	result.Date = perEvs[0].Event.Date
	for _, pe := range perEvs {
		result.Total += pe.Spent
		result.Count += pe.Factor
	}
	result.Average = result.Total / float64(result.Count)

	var balances []balance
	for _, pe := range perEvs {
		balances = append(balances,
			balance{perId: pe.PersonId, perName: pe.Person.Name, balance: pe.Spent - result.Average})
	}

	sort.SliceStable(balances, func(i, j int) bool {
		return balances[i].balance < balances[j].balance
	})

	return result, balances, nil
}

func (r *Response) calculateBalances(balances []balance) {
	// i = most indepted Person, j = most portable Person
	for i, j := 0, len(balances)-1; i < j; {
		switch {
		// if Balance of 'i' great them 'j' and the it's left to next 'j+1' Person
		case balances[i].balance+balances[j].balance < 0:
			if r.Owes[balances[i].perName] == nil {
				r.Owes[balances[i].perName] = map[string]float64{}
			}
			r.Owes[balances[i].perName][balances[j].perName] = balances[j].balance
			balances[i].balance += balances[j].balance
			balances[j].balance = 0
			i++
		// if Balance of 'i' less them 'j' and 'j' should take from 'i+1' Person
		case balances[i].balance+balances[j].balance >= 0 &&
			(balances[i].balance != 0 && balances[j].balance != 0):
			if r.Owes[balances[i].perName] == nil {
				r.Owes[balances[i].perName] = map[string]float64{}
			}
			r.Owes[balances[i].perName][balances[j].perName] = balances[i].balance
			balances[j].balance += balances[i].balance
			balances[i].balance = 0
			j--
		case balances[i].balance == 0:
			i++
		case balances[j].balance == 0:
			j--
		}
	}
}

func (s *CalcService) CalculateEvent(ctx context.Context, eventId int64) Response {
	result, balances, _ := s.createResponse(ctx, eventId)
	result.calculateBalances(balances)
	return result
}
