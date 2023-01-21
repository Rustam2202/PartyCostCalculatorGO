package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"sort"
)

type Person struct {
	Name       string `json:"name"`
	Spent      uint   `json:"spent"`
	Balance    float32
	IndeptedTo map[string]float32
}

type Persons struct {
	Persons []Person `json:"persons"`
}

type PartyData struct {
	persons        []Person
	average_amount float32
	total_amount   uint
}

func (data *PartyData) CalculateTotalAndAverageAmount() {
	for _, p := range data.persons {
		data.total_amount += p.Spent
	}
	data.average_amount = float32(data.total_amount) / float32(len(data.persons))
}

func (data *PartyData) CalculateBalances() {
	for i := 0; i < len(data.persons); i++ {
		data.persons[i].Balance = data.average_amount - float32(data.persons[i].Spent)
	}
}

func (data *PartyData) SortPersons(person []Person) {
	sort.SliceStable(person, func(i, j int) bool {
		return person[i].Balance < person[j].Balance
	})
}

func CalculateDebts(input Persons, errorRate float32) PartyData {
	var result = PartyData{
		persons: input.Persons,
	}
	for i := 0; i < len(result.persons); i++ {
		result.persons[i].IndeptedTo = make(map[string]float32) // need not nil map to write Balance
	}
	result.CalculateTotalAndAverageAmount()
	result.CalculateBalances()
	sort.SliceStable(result.persons, func(i, j int) bool {
		return result.persons[i].Balance < result.persons[j].Balance
	})
	// SortPersons(result.persons)
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

func (data *PartyData) CheckCalculation() {
	var sum_debts float64
	var sum_spents float64
	for _, p := range data.persons {
		if len(p.IndeptedTo) > 0 {
			for _, v := range p.IndeptedTo {
				sum_debts += float64(v)
			}
		}
		if p.Spent >= uint(data.average_amount) {
			sum_spents += math.Abs((float64(p.Spent) - float64(data.average_amount)))
		}
	}
	difference := math.Abs(sum_debts - sum_spents)
	fmt.Println("Sum of debts: ", sum_debts)
	fmt.Println("Total spent: ", sum_spents)
	fmt.Println("Difference after calculation is ", difference)
}

func (data *PartyData) ShowPayments() {
	for _, p := range data.persons {
		if len(p.IndeptedTo) > 0 {
			fmt.Println(p.Name, "owes to:")
			for k, v := range p.IndeptedTo {
				fmt.Println("  ", k, v)
			}
		}
	}
}

func main() {
	jsonInput, err := os.Open("persons1.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonInput.Close()

	byteValue, _ := ioutil.ReadAll(jsonInput)
	var personsFromJSON Persons
	json.Unmarshal(byteValue, &personsFromJSON)

	fmt.Println("Input persons:")
	for _, p := range personsFromJSON.Persons {
		fmt.Printf("%s spent: %d\n", p.Name, p.Spent)
	}

	result := CalculateDebts(personsFromJSON, 1)
	result.ShowPayments()

	//Test1()
	//Test2()
}

func Test1() {
	persons := []Person{
		{Name: "Alex", Spent: 90, IndeptedTo: make(map[string]float32)},
		{Name: "Marry", Spent: 55, IndeptedTo: make(map[string]float32)},
		{Name: "Jhon", Spent: 0, IndeptedTo: make(map[string]float32)},
		{Name: "Mike", Spent: 25, IndeptedTo: make(map[string]float32)},
		{Name: "Suzan", Spent: 30, IndeptedTo: make(map[string]float32)},
		{Name: "Bob", Spent: 0, IndeptedTo: make(map[string]float32)},
		{Name: "Jack", Spent: 5, IndeptedTo: make(map[string]float32)},
	}
	data := PartyData{persons: persons}

	data.CalculateTotalAndAverageAmount()
	data.CalculateBalances()
	//	data.CalculateDebts(1)
	data.CheckCalculation()
	data.ShowPayments()
}

func Test2() {
	persons := []Person{
		{Name: "Alex", Spent: 0, IndeptedTo: make(map[string]float32)},
		{Name: "Marry", Spent: 2000, IndeptedTo: make(map[string]float32)},
		{Name: "Jhon", Spent: 4900, IndeptedTo: make(map[string]float32)},
		{Name: "Mike", Spent: 0, IndeptedTo: make(map[string]float32)},
		{Name: "Suzan", Spent: 750, IndeptedTo: make(map[string]float32)},
		{Name: "Bob", Spent: 0, IndeptedTo: make(map[string]float32)},
		{Name: "Jack", Spent: 12000, IndeptedTo: make(map[string]float32)},
		{Name: "Pite", Spent: 49500, IndeptedTo: make(map[string]float32)},
	}
	data := PartyData{persons: persons}

	data.CalculateTotalAndAverageAmount()
	data.CalculateBalances()
	//	data.CalculateDebts(1)
	data.CheckCalculation()
	data.ShowPayments()
}
