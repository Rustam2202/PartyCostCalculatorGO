package models

import "time"

type Person struct {
	Id   int64
	Name string
	//LastName  string
}

type Event struct {
	Id   int64
	Name string
	Date time.Time // postgres format: `YYYY-MM-DD`
	// AverageAmount float32
	TotalAmount float64
}

type PersonsAndEvents struct {
	Id       int64
	PersonId int64
	EventId  int64
	Spent    float64
	Factor   int
}
