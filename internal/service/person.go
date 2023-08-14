package service

import (
	"context"
	"party-calc/internal/domain"
)

type PersonRepository interface {
	Create(ctx context.Context, per *domain.Person) error
	GetById(ctx context.Context, id int64) (*domain.Person, error)
	GetByName(ctx context.Context, name string) (*domain.Person, error)
	Update(ctx context.Context, per *domain.Person) error
	DeleteById(ctx context.Context, id int64) error
	DeleteByName(ctx context.Context, name string) error
}

type PersonService struct {
	repo PersonRepository
}

func NewPersonService(r PersonRepository) *PersonService {
	return &PersonService{repo: r}
}

func (p *PersonService) NewPerson(ctx context.Context, name string) (int64, error) {
	per := domain.Person{Name: name}
	err := p.repo.Create(ctx, &per)
	if err != nil {
		return 0, err
	}
	return per.Id, nil
}

func (p *PersonService) GetPersonById(ctx context.Context, id int64) (*domain.Person, error) {
	result, err := p.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *PersonService) GetPersonByName(ctx context.Context, name string) (*domain.Person, error) {
	result, err := p.repo.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *PersonService) UpdatePerson(ctx context.Context, id int64, name string) error {
	err := p.repo.Update(ctx, &domain.Person{Id: id, Name: name})
	if err != nil {
		return err
	}
	return nil
}

func (p *PersonService) DeletePersonById(ctx context.Context, id int64) error {
	err := p.repo.DeleteById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *PersonService) DeletePersonByName(ctx context.Context, name string) error {
	err := p.repo.DeleteByName(ctx, name)
	if err != nil {
		return err
	}
	return nil
}
