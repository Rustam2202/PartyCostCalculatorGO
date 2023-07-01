package domain

import "time"

type Person struct {
	Id     int64
	Name   string
	Events []Event
}

type Event struct {
	Id          int64
	Name        string
	Date        time.Time // postgres format: `YYYY-MM-DD`
	Persons     []Person
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
