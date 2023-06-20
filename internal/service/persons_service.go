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

func (p *PersonService) GetPerson(name string) (models.Person, error) {
	per, err := p.repo.Get(&models.Person{Name: name})
	if err != nil {
		return models.Person{}, err
	}
	return per, nil
}

func (p *PersonService) UpdatePerson(name, newName string) error {
	perOld, err := p.GetPerson(name)
	if err != nil {
		return err
	}
	err = p.repo.Update(&perOld, &models.Person{Name: newName})
	if err != nil {
		return err
	}
	return nil
}

func (p *PersonService) DeletePerson(name string) error {
	err := p.repo.Delete(&models.Person{Name: name})
	if err != nil {
		return err
	}
	return nil
}
