package main

import (
	//"fmt"
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

func (data *PartyData) CalculateAll() {
	SortPersons(data.persons)
	i := 0
	j := len(data.persons) - 1
	for i < j {
		left := &data.persons[i]
		right := &data.persons[j]
		absLeftBalance := math.Abs(float64(left.balance))
		if absLeftBalance >= float64(right.balance) {
			left.balance += right.balance
			right.indeptedTo[left.name] = right.balance
			right.balance = 0
			j--
		} else {
			right.balance -= float32(absLeftBalance)
			right.indeptedTo[left.name] = right.balance
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

func main() {
	persons := []Person{
		{name: "Alex", spent: 90},
		{name: "Marry", spent: 55},
		{name: "Jhon", spent: 0},
		{name: "Mike", spent: 25},
		{name: "Suzan", spent: 30},
		{name: "Bob", spent: 0},
		{name: "Jack", spent: 5},
	}
	m := map[string]float32{}
	data := PartyData{persons: persons, indeptedTo: m}
	data.CalculateTotalAndAverageAmount()
	data.CalculateBalances()
	data.CalculateAll()
}
