package main

import "testing"

type testType struct {
	name   string
	input  []Person
	output PartyData
}

var test1 = testType{
	name: "Test with left preponderance iterator",
	input: []Person{
		{name: "Alex", spent: 90, indeptedTo: make(map[string]float32)},
		{name: "Marry", spent: 55, indeptedTo: make(map[string]float32)},
		{name: "Jhon", spent: 0, indeptedTo: make(map[string]float32)},
		{name: "Mike", spent: 25, indeptedTo: make(map[string]float32)},
		{name: "Suzan", spent: 30, indeptedTo: make(map[string]float32)},
		{name: "Bob", spent: 0, indeptedTo: make(map[string]float32)},
		{name: "Jack", spent: 5, indeptedTo: make(map[string]float32)},
	},
	output: PartyData{
		persons: []Person{
			{name: "Alex", spent: 90, balance: 0, indeptedTo: make(map[string]float32)},
			{name: "Marry", spent: 55, indeptedTo: make(map[string]float32)},
			{name: "Jhon", spent: 0, indeptedTo: make(map[string]float32)},
			{name: "Mike", spent: 25, indeptedTo: make(map[string]float32)},
			{name: "Suzan", spent: 30, indeptedTo: make(map[string]float32)},
			{name: "Bob", spent: 0, indeptedTo: make(map[string]float32)},
			{name: "Jack", spent: 5, indeptedTo: make(map[string]float32)},
		},
		average_amount: 29.285715,
		total_amount:   205,
	},
}

func TestCalculateDebts(t *testing.T) {

}
