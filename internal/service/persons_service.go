package service

import (
	"party-calc/internal/domain"
)

type PersonRepository interface {
	Create(per *domain.Person) error
	GetById(id int64) (*domain.Person, error)
	GetByName(name string) (*domain.Person, error)
	Update(per *domain.Person) error
	DeleteById(id int64) error
	DeleteByName(name string) error
}

type PersonService struct {
	repo PersonRepository
}

func NewPersonService(r PersonRepository) *PersonService {
	return &PersonService{repo: r}
}

func (p *PersonService) NewPerson(name string) (int64, error) {
	per := domain.Person{Name: name}
	err := p.repo.Create(&per)
	if err != nil {
		return 0, err
	}
	return per.Id, nil
}

func (p *PersonService) GetPersonById(id int64) (*domain.Person, error) {
	result, err := p.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *PersonService) GetPersonByName(name string) (*domain.Person, error) {
	result, err := p.repo.GetByName(name)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *PersonService) UpdatePerson(id int64, name string) error {
	err := p.repo.Update(&domain.Person{Id: id, Name: name})
	if err != nil {
		return err
	}
	return nil
}

func (p *PersonService) DeletePersonById(id int64) error {
	err := p.repo.DeleteById(id)
	if err != nil {
		return err
	}
	return nil
}

func (p *PersonService) DeletePersonByName(name string) error {
	err := p.repo.DeleteByName(name)
	if err != nil {
		return err
	}
	return nil
}
