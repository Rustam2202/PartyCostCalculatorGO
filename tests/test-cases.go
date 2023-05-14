package tests

import (
	"party-calc/internal"
	"party-calc/internal/person"
)

type testStruct struct {
	testName string
	input    person.Persons //`json:"persons"`
	want     internal.PartyData
}

var onePerson = testStruct{
	testName: "One Person",
	input: person.Persons{Persons: []person.Person{
		{Name: "Person 1", Spent: 1000},
	}},

	want: internal.PartyData{
		Persons: []person.Person{
			{Name: "Person 1", Spent: 1000, Participants: 1, IndeptedTo: map[string]float32{}},
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

	want: internal.PartyData{
		Persons: []person.Person{
			{Name: "Person 1", Spent: 1000, Participants: 1, IndeptedTo: map[string]float32{}},
			{Name: "Person 2", Spent: 200, Participants: 1, IndeptedTo: map[string]float32{"Person 1": 400}},
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

	want: internal.PartyData{
		Persons: []person.Person{
			{Name: "Person 1", Spent: 1000, Participants: 1, IndeptedTo: map[string]float32{}},
			{Name: "Person 2", Spent: 500, Participants: 1, IndeptedTo: map[string]float32{}},
			{Name: "Person 3", Spent: 0, Participants: 1, IndeptedTo: map[string]float32{"Person 1": 500}},
		},
		AllPersonsCount: 3,
		AverageAmount:   500,
		TotalAmount:     1500},
}
