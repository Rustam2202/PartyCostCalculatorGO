package internal

import (
	"fmt"
	"math"
	"os"
	"sort"

	"party-calc/internal/language"
	"party-calc/internal/person"
	"party-calc/utils"
)

type PartyData struct {
	Persons         []person.Person `json:"persons"`
	AllPersonsCount uint
	AverageAmount   float32
	TotalAmount     uint
}

func CalculateDebts(input person.Persons, errorRate float32) PartyData {
	var result = PartyData{
		Persons: input.Persons,
	}
	for i := 0; i < len(result.Persons); i++ {
		result.Persons[i].IndeptedTo = make(map[string]float32)
		if result.Persons[i].Participants == 0 { // if "participants" not declareted in json, then one participant
			result.Persons[i].Participants = 1
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
		absLeftBalance := math.Abs(float64(left.Balance))
		if absLeftBalance >= float64(right.Balance) {
			if absLeftBalance < float64(errorRate) {
				left.Balance = 0
				i++
				continue
			}
			right.IndeptedTo[left.Name] = right.Balance
			left.Balance += right.Balance
			right.Balance = 0
			j--
		} else {
			right.IndeptedTo[left.Name] = float32(absLeftBalance)
			right.Balance -= float32(absLeftBalance)
			if float64(right.Balance) < float64(errorRate) {
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
		data.AllPersonsCount += p.Participants
	}
	data.AverageAmount = float32(data.TotalAmount) / float32(data.AllPersonsCount)
}

func (data *PartyData) CalculateBalances() {
	for i := 0; i < len(data.Persons); i++ {
		data.Persons[i].Balance = data.AverageAmount*float32(data.Persons[i].Participants) - float32(data.Persons[i].Spent)
	}
}

func (data *PartyData) ShowPayments() {
	fmt.Println(data.PrintSpents())
	fmt.Println(data.PrintPayments())
}

func (data *PartyData) PrintToFile(fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		utils.Logger.Error("Problem with creating file")
		panic(nil)
	}
	fmt.Fprintln(file, data.PrintSpents())
	fmt.Fprintln(file, data.PrintPayments())
}

func (data *PartyData) PrintSpents() string {
	var result string
	switch utils.Cfg.Language {
	case language.ENG:
		result += "   Participants:\n"
	case language.RUS:
		result += "   Участники:\n"
	}

	for _, p := range data.Persons {
		switch utils.Cfg.Language {
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
	switch utils.Cfg.Language {
	case language.ENG:
		result += "   Payments:\n"
	case language.RUS:
		result += "   Выплаты:\n"
	}

	for _, p := range data.Persons {
		if len(p.IndeptedTo) > 0 {
			switch utils.Cfg.Language {
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

	switch utils.Cfg.Language {
	case language.ENG:
		result += fmt.Sprintf("\nAverage to person: %0.1f\n", data.AverageAmount)
	case language.RUS:
		result += fmt.Sprintf("\nСреднее на человека: %0.1f\n", data.AverageAmount)
	}

	return result
}
