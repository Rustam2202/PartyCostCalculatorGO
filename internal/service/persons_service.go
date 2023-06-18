package service

import (
	"party-calc/internal/database"
	"party-calc/internal/database/models"
)

type PersonService struct {
	repo *database.PersonRepository
}

func (p *PersonService) NewPerson(name string) (int64, error) {
	id, err := p.repo.Create(&models.Person{Name: name})
	if err != nil {
		return 0, err
	}
	return id, nil
}
