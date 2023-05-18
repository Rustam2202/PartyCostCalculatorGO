package service

import (
	"party-calc/internal/person"
	"testing"
)

func Test1(t *testing.T) {
	var per1, per2, per3 person.Person
	var pers person.Persons
	per1.InitPerson()
	per2.InitPerson()
	per3.InitPerson()

	per1.Name = "Person 1"
	per1.Factor = 2
	per1.Spent = 1000

	per2.Name = "Person 2"
	per2.Spent = 500

	per3.Name = "Person 3"

	pers.AddPerson(per1)
	pers.AddPerson(per2)
	pers.AddPerson(per3)

	//	result := CalculateDebts(pers)
}
