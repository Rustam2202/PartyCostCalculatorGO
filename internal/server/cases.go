package server

import (
	"party-calc/internal/person"
	"party-calc/internal/service"
)

type testStruct struct {
	testName string
	input    person.Persons //`json:"persons"`
	want     service.PartyData
}

var onePerson = testStruct{
	testName: "One Person",
	input: person.Persons{Persons: []person.Person{
		{Name: "Person 1", Spent: 1000},
	}},

	want: service.PartyData{
		Persons: []person.Person{
			{Name: "Person 1", Spent: 1000, Factor: 1, Owe: map[string]float64{}},
		},
		AllPersonsCount: 1,
		AverageAmount:   1000,
		TotalAmount:     1000},
}

var twoPersons = testStruct{
	testName: "Two Persons",
	input: person.Persons{Persons: []person.Person{
		{Name: "Person 1", Spent: 1000},
		{Name: "Person 2", Spent: 200},
	}},

	want: service.PartyData{
		Persons: []person.Person{
			{Name: "Person 1", Spent: 1000, Factor: 1, Owe: map[string]float64{}},
			{Name: "Person 2", Spent: 200, Factor: 1, Owe: map[string]float64{"Person 1": 400}},
		},
		AllPersonsCount: 2,
		AverageAmount:   600,
		TotalAmount:     1200},
}

var threePersons = testStruct{
	testName: "Three Persons",
	input: person.Persons{Persons: []person.Person{
		{Name: "Alex and Kate", Spent: 800, Factor: 2},
		{Name: "Peter", Spent: 600},
		{Name: "Ivan"},
	}},

	want: service.PartyData{
		Persons: []person.Person{
			{Name: "Peter", Spent: 600, Factor: 1, Owe: nil},
			{Name: "Alex and Kate", Spent: 800, Factor: 2, Owe: nil},
			{Name: "Ivan", Spent: 0, Factor: 1, Owe: map[string]float64{"Peter": 250, "Alex and Kate": 100}},
		},
		AllPersonsCount: 4,
		AverageAmount:   350,
		TotalAmount:     1400},
}

var fivePersons = testStruct{
	testName: "Five Persons",
	input: person.Persons{Persons: []person.Person{
		{Name: "Person 1", Spent: 1000},
		{Name: "Person 2", Spent: 800},
		{Name: "Person 3", Spent: 0},
		{Name: "Person 4", Spent: 0},
		{Name: "Person 5", Spent: 0},
	}},

	want: service.PartyData{
		Persons: []person.Person{
			{Name: "Person 1", Spent: 1000, Factor: 1, Owe: nil},
			{Name: "Person 2", Spent: 800, Factor: 1, Owe: nil},
			{Name: "Person 3", Spent: 0, Factor: 1, Owe: map[string]float64{"Person 2": 360}},
			{Name: "Person 4", Spent: 0, Factor: 1, Owe: map[string]float64{"Person 1": 280, "Person 2": 80}},
			{Name: "Person 5", Spent: 0, Factor: 1, Owe: map[string]float64{"Person 1": 360}},
		},
		AllPersonsCount: 5,
		AverageAmount:   360,
		TotalAmount:     1800},
}

// JSONs not equal because different positions of blocks
var sixPersons = testStruct{
	testName: "Six Persons",
	input: person.Persons{Persons: []person.Person{
		{Name: "Person 1", Spent: 1000, Factor: 2},
		{Name: "Person 2", Spent: 800, Factor: 1},
		{Name: "Person 3", Spent: 300, Factor: 3},
		{Name: "Person 4", Spent: 0, Factor: 1},
		{Name: "Person 5", Spent: 0, Factor: 2},
		{Name: "Person 6", Spent: 0, Factor: 3},
	}},

	want: service.PartyData{
		Persons: []person.Person{
			{Name: "Person 1", Spent: 1000, Factor: 2},
			{Name: "Person 2", Spent: 800, Factor: 1},
			{Name: "Person 3", Spent: 300, Factor: 3, Owe: map[string]float64{"Person 2": 225}},
			{Name: "Person 4", Spent: 0, Factor: 1, Owe: map[string]float64{"Person 2": 175}},
			{Name: "Person 5", Spent: 0, Factor: 2, Owe: map[string]float64{"Person 1": 125, "Person 2": 225}},
			{Name: "Person 6", Spent: 0, Factor: 3, Owe: map[string]float64{"Person 1": 525}},
		},
		AllPersonsCount: 12,
		AverageAmount:   175,
		TotalAmount:     2100},
}
