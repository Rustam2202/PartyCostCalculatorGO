package service

import (
	"math"
	"sort"

	//"party-calc/internal/config"
	"party-calc/internal/person"
)

type PartyData struct {
	Id              int             `json:"id"`
	Persons         []person.Person `json:"persons"`
	AllPersonsCount uint            `json:"persons_count"`
	AverageAmount   float64         `json:"average"`
	TotalAmount     uint            `json:"total"`
}

func CalculateDebts(input person.Persons) PartyData {
	var result = PartyData{
		Persons: input.Persons,
	}
	for i := 0; i < len(result.Persons); i++ {
		if result.Persons[i].Factor == 0 { // if "participants" not declareted in json, then one participant
			result.Persons[i].Factor = 1
		}
	}
	result.CalculateTotalAndAverageAmount()
	result.CalculateBalances()
	sort.SliceStable(result.Persons, func(i, j int) bool {
		return result.Persons[i].Balance < result.Persons[j].Balance
	})

	i := 0                       // left iterator
	j := len(result.Persons) - 1 // right iterator
	for i < j {
		left := &result.Persons[i]
		right := &result.Persons[j]
		absLeftBalance := math.Abs(left.Balance)

		if absLeftBalance == 0 {
			i++
			continue
		}
		if right.Balance == 0 {
			j--
			continue
		}

		if absLeftBalance >= right.Balance {
			// if absLeftBalance < config.Config.RoundRate {
			// 	left.Balance = 0
			// 	i++
			// 	continue
			// }
			if right.Owe == nil {
				right.Owe = make(map[string]float64)
			}
			right.Owe[left.Name] = right.Balance
			left.Balance += right.Balance
			right.Balance = 0
			j--
		} else if absLeftBalance < right.Balance {
			if right.Owe == nil {
				right.Owe = make(map[string]float64)
			}
			right.Owe[left.Name] = absLeftBalance
			right.Balance -= absLeftBalance
			// if right.Balance < config.Cfg.RoundRate {
			// 	right.Balance = 0
			// }
			left.Balance = 0
			i++
		}
	}
	return result
}

func (data *PartyData) CalculateTotalAndAverageAmount() {
	for _, p := range data.Persons {
		data.TotalAmount += p.Spent
	}
	for _, p := range data.Persons {
		data.AllPersonsCount += p.Factor
	}
	data.AverageAmount = float64(data.TotalAmount) / float64(data.AllPersonsCount)
}

func (data *PartyData) CalculateBalances() {
	for i := 0; i < len(data.Persons); i++ {
		data.Persons[i].Balance = data.AverageAmount*float64(data.Persons[i].Factor) - float64(data.Persons[i].Spent)
	}
}
