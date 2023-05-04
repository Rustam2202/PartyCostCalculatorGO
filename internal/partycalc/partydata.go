package partycalc

import (
	"fmt"
	"math"
	"os"
	"sort"
)

type PartyData struct {
	persons         []Person
	AllPersonsCount uint
	average_amount  float32
	total_amount    uint
}

func (data *PartyData) CalculateTotalAndAverageAmount() {
	for _, p := range data.persons {
		data.total_amount += p.Spent
	}
	for _, p := range data.persons {
		data.AllPersonsCount += p.Participants
	}
	data.average_amount = float32(data.total_amount) / float32(data.AllPersonsCount)
}

func (data *PartyData) CalculateBalances() {
	for i := 0; i < len(data.persons); i++ {
		data.persons[i].Balance = data.average_amount*float32(data.persons[i].Participants) - float32(data.persons[i].Spent)
	}
}

func CalculateDebts(input Persons, errorRate float32) PartyData {
	var result = PartyData{
		persons: input.Persons,
	}
	for i := 0; i < len(result.persons); i++ {
		result.persons[i].IndeptedTo = make(map[string]float32) // need `not nil` map to write Balance
		if result.persons[i].Participants == 0 {                // if "participants" not declareted in json, then one person
			result.persons[i].Participants = 1
		}
	}
	result.CalculateTotalAndAverageAmount()
	result.CalculateBalances()
	sort.SliceStable(result.persons, func(i, j int) bool {
		return result.persons[i].Balance < result.persons[j].Balance
	})

	i := 0                       // left iterator
	j := len(result.persons) - 1 // right iterator
	for i < j {
		left := &result.persons[i]
		right := &result.persons[j]
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

func (data *PartyData) CheckCalculation(input Persons) {

	var totalSpent uint
	var averagePerPerson float32
	var allPersonCount uint
	var allPayments float32
	var balances map[string]float32

	for _, p := range input.Persons {
		totalSpent += p.Spent
		if p.Participants == 0 {
			allPersonCount++
		} else {
			allPersonCount += p.Participants
		}
	}
	averagePerPerson = float32(totalSpent) / float32(allPersonCount)

	for _, p := range data.persons {
		for _, p2 := range p.IndeptedTo {
			allPayments += p2
		}
	}

	fmt.Printf("Average per person: %f\n", averagePerPerson)
	for name, balance := range balances {
		fmt.Printf("%s has %f balance", name, balance)
	}
}

func (data *PartyData) ShowPayments(lang Language) {
	fmt.Println(data.PrintSpents(lang))
	fmt.Println(data.PrintPayments(lang))
}

func (data *PartyData) PrintSpents(lang Language) string {
	var result string
	if lang == ENG {
		result += "   Participants:\n"
	} else if lang == RUS {
		result += "   Участники:\n"
	}
	//result += "   Participants:\n"
	for _, p := range data.persons {
		if lang == ENG {
			result += fmt.Sprintf("%s (x%d) spent: %d\n", p.Name, p.Participants, p.Spent)
		} else if lang == RUS {
			result += fmt.Sprintf("%s (x%d) потрачено: %d\n", p.Name, p.Participants, p.Spent)
		}
		// result += fmt.Sprintf("%s (x%d) spent: %d\n", p.Name, p.Participants, p.Spent)
	}
	return result
}

func (data *PartyData) PrintPayments(lang Language) string {
	var result string
	if lang == ENG {
		result += "   Payments:\n"
	} else if lang == RUS {
		result += "   Выплаты:\n"
	}
	//result += "   Payments:\n"
	for _, p := range data.persons {
		if len(p.IndeptedTo) > 0 {
			if lang == ENG {
				result += fmt.Sprintf("%s owes to:\n", p.Name)
			} else if lang == RUS {
				result += fmt.Sprintf("%s выплачивает:\n", p.Name)
			}
			//result += fmt.Sprintf("%s owes to:\n", p.Name)
			for name, debt := range p.IndeptedTo {
				result += fmt.Sprintf("  %s %.f\n", name, debt) // .f - format output of integer and decimal
			}
		}
	}
	if lang == ENG {
		result += fmt.Sprintf("\nAverage to person: %0.1f\n", data.average_amount)
	} else if lang == RUS {
		result += fmt.Sprintf("\nСреднее на человека: %0.1f\n", data.average_amount)
	}
	return result
}

func (data *PartyData) PrintToFile(fileName string, lang Language) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(file, data.PrintSpents(lang))
	fmt.Fprintln(file, data.PrintPayments(lang))
}