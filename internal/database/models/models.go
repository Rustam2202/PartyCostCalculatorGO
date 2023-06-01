package models

import "time"

type Person struct {
	Id   int
	Name string
	//LastName  string
}

type Event struct {
	Id            int
	Name          string
	Date          time.Time // postgres format: `YYYY-MM-DD`
	// AverageAmount float32
	TotalAmount   float64
}

type PersonsAndEvents struct {
	Id       int
	PersonId int
	EventId  int
	Spent    float64
	Factor   int
}
