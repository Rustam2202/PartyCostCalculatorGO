package tests

import (
	"testing"

	c "party-calc/internal"
	p "party-calc/internal/person"
)

type testType struct {
	name   string
	input  []p.Person
	output c.PartyData
}

var test1 = testType{
	name: "Test with left preponderance iterator",
	input: []p.Person{
		{Name: "Alex", Spent: 90, IndeptedTo: make(map[string]float32)},
		{Name: "Marry", Spent: 55, IndeptedTo: make(map[string]float32)},
		{Name: "Jhon", Spent: 0, IndeptedTo: make(map[string]float32)},
		{Name: "Mike", Spent: 25, IndeptedTo: make(map[string]float32)},
		{Name: "Suzan", Spent: 30, IndeptedTo: make(map[string]float32)},
		{Name: "Bob", Spent: 0, IndeptedTo: make(map[string]float32)},
		{Name: "Jack", Spent: 5, IndeptedTo: make(map[string]float32)},
	},
	/*
	output: PartyData{
		persons: []p.Person{
			{Name: "Alex", Spent: 90, Balance: 0, IndeptedTo: make(map[string]float32)},
			{Name: "Marry", Spent: 55, IndeptedTo: make(map[string]float32)},
			{Name: "Jhon", Spent: 0, IndeptedTo: make(map[string]float32)},
			{Name: "Mike", Spent: 25, IndeptedTo: make(map[string]float32)},
			{Name: "Suzan", Spent: 30, IndeptedTo: make(map[string]float32)},
			{Name: "Bob", Spent: 0, IndeptedTo: make(map[string]float32)},
			{Name: "Jack", Spent: 5, IndeptedTo: make(map[string]float32)},
		},
		average_amount: 29.285715,
		total_amount:   205,
	},
	*/
}

func TestCalculateDebts(t *testing.T) {

}
