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
			{Name: "Person 1", Spent: 1000, Factor: 1, IndeptedTo: map[string]float64{}},
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
			{Name: "Person 1", Spent: 1000, Factor: 1, IndeptedTo: map[string]float64{}},
			{Name: "Person 2", Spent: 200, Factor: 1, IndeptedTo: map[string]float64{"Person 1": 400}},
		},
		AllPersonsCount: 2,
		AverageAmount:   600,
		TotalAmount:     1200},
}

var threePersons = testStruct{
	testName: "Three Persons",
	input: person.Persons{Persons: []person.Person{
		{Name: "Person 1", Spent: 1000},
		{Name: "Person 2", Spent: 500},
		{Name: "Person 3", Spent: 0},
	}},

	want: service.PartyData{
		Persons: []person.Person{
			{Name: "Person 1", Spent: 1000, Factor: 1, IndeptedTo: map[string]float64{}},
			{Name: "Person 2", Spent: 500, Factor: 1, IndeptedTo: map[string]float64{}},
			{Name: "Person 3", Spent: 0, Factor: 1, IndeptedTo: map[string]float64{"Person 1": 500}},
		},
		AllPersonsCount: 3,
		AverageAmount:   500,
		TotalAmount:     1500},
}
