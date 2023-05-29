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
	TotalAmount   float32
}

type PersonsAndEvents struct {
	Id       int
	PersonId int
	EventId  int
	Spent    float32
	Factor   int
}
