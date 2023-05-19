package models

import "time"

type Person struct {
	Id     int
	Name   string
	Spent  int
	Factor int
	//EventsIds    []int
}

type Event struct {
	Id   int
	Date time.Time // postgres format: `YYYY-MM-DD`
	Name string
	//PersonsIds []int

}

type PersonsAndEvents struct {
	Id       int
	PersonId int
	EventId  int
}

type IndeptedToData struct {
	Id         int
	DebtorId   int
	CreditorId int
	Sum        float32
	PartyId    int
}
