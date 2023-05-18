package service

import (
	"math"
	"sort"

	"party-calc/internal/config"
	"party-calc/internal/person"
)

type PartyData struct {
	Persons         []person.Person `json:"persons"`
	AllPersonsCount uint
	AverageAmount   float64
	TotalAmount     uint
}

func CalculateDebts(input person.Persons) PartyData {
	var result = PartyData{
		Persons: input.Persons,
	}
	for i := 0; i < len(result.Persons); i++ {
		result.Persons[i].IndeptedTo = make(map[string]float64)
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
		if absLeftBalance >= right.Balance {
			if absLeftBalance < config.Cfg.RoundRate {
				left.Balance = 0
				i++
				continue
			}
			right.IndeptedTo[left.Name] = right.Balance
			left.Balance += right.Balance
			right.Balance = 0
			j--
		} else {
			right.IndeptedTo[left.Name] = absLeftBalance
			right.Balance -= absLeftBalance
			if right.Balance < config.Cfg.RoundRate {
				right.Balance = 0
			}
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

/*
func (data *PartyData) ShowPayments() {
	fmt.Println(data.PrintSpents())
	fmt.Println(data.PrintPayments())
}

func (data *PartyData) PrintToFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		logger.Logger.Error("Problem with creating file")
		panic(nil)
	}
	fmt.Fprintln(file, data.PrintSpents())
	fmt.Fprintln(file, data.PrintPayments())
}

func (data *PartyData) PrintSpents() string {
	var result string
	switch config.Cfg.Language {
	case language.ENG:
		result += "   Participants:\n"
	case language.RUS:
		result += "   Участники:\n"
	}

	for _, p := range data.Persons {
		switch config.Cfg.Language {
		case language.ENG:
			result += fmt.Sprintf("%s (x%d) spent: %d\n", p.Name, p.Participants, p.Spent)
		case language.RUS:
			result += fmt.Sprintf("%s (x%d) потрачено: %d\n", p.Name, p.Participants, p.Spent)
		}
	}

	return result
}

func (data *PartyData) PrintPayments() string {
	var result string
	switch config.Cfg.Language {
	case language.ENG:
		result += "   Payments:\n"
	case language.RUS:
		result += "   Выплаты:\n"
	}

	for _, p := range data.Persons {
		if len(p.IndeptedTo) > 0 {
			switch config.Cfg.Language {
			case language.ENG:
				result += fmt.Sprintf("%s owes to:\n", p.Name)
			case language.RUS:
				result += fmt.Sprintf("%s выплачивает:\n", p.Name)
			}

			for name, debt := range p.IndeptedTo {
				result += fmt.Sprintf("  %s %.f\n", name, debt) // .f - format output of integer and decimal
			}
		}
	}

	switch config.Cfg.Language {
	case language.ENG:
		result += fmt.Sprintf("\nAverage to person: %0.1f\n", data.AverageAmount)
	case language.RUS:
		result += fmt.Sprintf("\nСреднее на человека: %0.1f\n", data.AverageAmount)
	}
	return result
}
*/
