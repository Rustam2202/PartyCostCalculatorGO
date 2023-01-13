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
	temp_persons := MakeNoZeroBalance(data.persons)
	if len(temp_persons) > 0 {
		SortPersons(temp_persons)
		PayDebt(temp_persons)
		data.CalculateAll()
	}
}

func MakeNoZeroBalance(person []Person) []Person {
	var noZeroBal []Person
	for _, p := range person {
		if p.balance != 0 {
			noZeroBal = append(noZeroBal, p)
		}
	}
	return noZeroBal
}

func SortPersons(person []Person) {
	sort.SliceStable(person, func(i, j int) bool {
		return person[i].balance < person[j].balance
	})
}

func PayDebt(person []Person) {
	debetor := &person[len(person)-1]
	recipient := &person[0]
	if math.Abs( float64(recipient.balance)) >= float64( debetor.balance) {
		recipient.balance += debetor.balance
		debetor.balance = 0
	} else {
		recipient.balance = 0
		debetor.balance -= recipient.balance
	}
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

	data := PartyData{persons: persons}
	data.CalculateTotalAndAverageAmount()
	data.CalculateBalances()
	data.CalculateAll()
}
