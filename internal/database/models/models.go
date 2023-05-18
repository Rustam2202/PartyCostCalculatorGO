package models

import "time"

type Persons struct {
	Id           int
	Name         string
	Spent        int
	Participants int
	EventsIds    []int
}

type IndeptedToData struct {
	Id         int
	DebtorId   int
	CreditorId int
	Sum        float32
	PartyId    int
}

type Event struct {
	Id         int
	PersonsIds []int
	Date       time.Time
	Name       string
}
