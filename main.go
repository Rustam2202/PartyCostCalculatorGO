package main

import (
	"fmt"
	"math"
	"sort"
)

type Person struct {
	name       string
	spent      uint
	balance    float32
	indeptedTo map[string]float32
}

type PartyData struct {
	persons        []Person
	average_amount float32
	total_amount   uint
}

func (data *PartyData) CalculateTotalAndAverageAmount() {
	for _, p := range data.persons {
		data.total_amount += p.spent
	}
	data.average_amount = float32(data.total_amount) / float32(len(data.persons))
}

func (data *PartyData) CalculateBalances() {
	for i := 0; i < len(data.persons); i++ {
		data.persons[i].balance = data.average_amount - float32(data.persons[i].spent)
	}
}

func (data *PartyData) CalculateDebts(errorRate float32) {
	SortPersons(data.persons)
	i := 0
	j := len(data.persons) - 1
	for i < j {
		left := &data.persons[i]
		right := &data.persons[j]
		absLeftBalance := math.Abs(float64(left.balance))
		if absLeftBalance >= float64(right.balance) {
			if absLeftBalance >= float64(errorRate) {
				left.balance += right.balance
				right.indeptedTo[left.name] = right.balance
			}
			right.balance = 0
			j--
		} else {
			if absLeftBalance >= float64(errorRate) {
				right.balance -= float32(absLeftBalance)
				right.indeptedTo[left.name] = right.balance
			}
			left.balance = 0
			i++
		}
	}
}

func SortPersons(person []Person) {
	sort.SliceStable(person, func(i, j int) bool {
		return person[i].balance < person[j].balance
	})
}

func (data *PartyData) CheckCalculation() {
	var sum_debts float64
	var sum_spents float64
	for _, p := range data.persons {
		if len(p.indeptedTo) > 0 {
			for _, v := range p.indeptedTo {
				sum_debts += float64(v)
			}
		}
		if p.spent >= uint(data.average_amount) {
			sum_spents += math.Abs((float64(p.spent) - float64(data.average_amount)))
		}
	}
	difference := math.Abs(sum_debts - sum_spents)
	fmt.Println(fmt.Sprintf("Difference after calculation is %0.2f", difference))
}

func (data *PartyData) ShowPayments() {
	for _, p := range data.persons {
		if len(p.indeptedTo) > 0 {
			fmt.Println(fmt.Sprintf("%s owes to:", p.name))
			for k, v := range p.indeptedTo {
				fmt.Println(fmt.Sprintf("  %s %0.2f ", k, v))
			}
		}
	}
}

func main() {
	persons := []Person{
		{name: "Alex", spent: 90, indeptedTo: make(map[string]float32)},
		{name: "Marry", spent: 55, indeptedTo: make(map[string]float32)},
		{name: "Jhon", spent: 0, indeptedTo: make(map[string]float32)},
		{name: "Mike", spent: 25, indeptedTo: make(map[string]float32)},
		{name: "Suzan", spent: 30, indeptedTo: make(map[string]float32)},
		{name: "Bob", spent: 0, indeptedTo: make(map[string]float32)},
		{name: "Jack", spent: 5, indeptedTo: make(map[string]float32)},
	}
	data := PartyData{persons: persons}

	data.CalculateTotalAndAverageAmount()
	data.CalculateBalances()
	data.CalculateDebts(1)
	data.CheckCalculation()
	data.ShowPayments()
}
