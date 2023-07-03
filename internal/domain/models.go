package domain

import "time"

type Person struct {
	Id       int64
	Name     string
	//EventIds []int64
	Events   []Event
}

type Event struct {
	Id        int64
	Name      string
	Date      time.Time // postgres format: `YYYY-MM-DD`
	//PersonIds []int64
	Persons   []Person
}

type PersonsAndEvents struct {
	Id       int64
	PersonId int64
	EventId  int64
	Spent    float64
	Factor   int
	Person   Person
	Event    Event
}
