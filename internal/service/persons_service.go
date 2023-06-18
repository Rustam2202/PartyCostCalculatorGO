package service

import (
	"party-calc/internal/database/models"
	"party-calc/internal/database/repository"
)

type PersonService struct {
	repo *repository.PersonRepository
}

func NewPersonService(r *repository.PersonRepository) *PersonService {
	return &PersonService{repo: r}
}

func (p *PersonService) NewPerson(name string) (int64, error) {
	id, err := p.repo.Create(&models.Person{Name: name})
	if err != nil {
		return 0, err
	}
	return id, nil
}
