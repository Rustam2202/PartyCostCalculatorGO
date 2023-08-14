package domain

import "time"

type Person struct {
	Id     int64  `json:"id" default:"123456789"`
	Name   string `json:"name" default:"Some Person name"`
	Events []Event
}

type Event struct {
	Id      int64     `json:"id" default:"987654321"`
	Name    string    `json:"name" default:"Some Event name"`
	Date    time.Time `json:"date" default:"2020-12-31"`
	Persons []Person
}

type PersonsAndEvents struct {
	Id       int64   `json:"id" default:"9223372036854775807"`
	PersonId int64   `json:"person_id" default:"123456789"`
	EventId  int64   `json:"event_id" default:"987654321"`
	Spent    float64 `json:"spent" default:"123.45"`
	Factor   int     `json:"factor" default:"1"`
	Person   Person
	Event    Event
}
