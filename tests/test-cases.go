package tests

import (
	"party-calc/internal"
	"party-calc/internal/person"
)

type data struct {
	Input      person.Persons `json:"persons"`
	Want       string
	WantStruct internal.PartyData
}

var case1 = data{
	Input: person.Persons{[]person.Person{
		{Name: "Person 1", Spent: 1000,},
		{Name: "Person 2", Spent: 500},
		{Name: "Person 3", Spent: 0},
	}},
	Want: `{
		"persons": [
			{
				"name": "Person 1",
				"spent": 1000,
				"participants": 1,
				"Balance": 0,
				"IndeptedTo": {}
			},
			{
				"name": "Person 2",
				"spent": 500,
				"participants": 1,
				"Balance": 0,
				"IndeptedTo": {}
			},
			{
				"name": "Person 34",
				"spent": 0,
				"participants": 1,
				"Balance": 0,
				"IndeptedTo": {
					"Person 1": 500
					}
			}
		],
		"AllPersonsCount": 3,
		"AverageAmount": 7500,
		"TotalAmount": 15700
		}
	}`,
	WantStruct: internal.PartyData{
		Persons: []person.Person{
			{Name: "Person 1", Spent: 1000, Participants: 1, Balance: 0, IndeptedTo: nil},
			{Name: "Person 2", Spent: 500, Participants: 1, Balance: 0, IndeptedTo: nil},
			{Name: "Person 3", Spent: 0, Participants: 1, Balance: 0, IndeptedTo: map[string]float32{"Person 1": 500}}, // make(map[string]float32{"Person 1", 500})
		},
		AllPersonsCount: 0,
		AverageAmount:   0,
		TotalAmount:     0},
}
