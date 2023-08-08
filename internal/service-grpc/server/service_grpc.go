package grpc

import (
	"context"
	"party-calc/internal/domain"
	pb "party-calc/internal/service-grpc/proto"
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

func addPerson() {
}

func (p *PersonService) getPerson(in pb.Id) pb.PersonData {
	per, _ := p.repo.GetById(context.Background(), in.Id)
	var eventsPb []*pb.EventData
	for _, ev := range per.Events {
		var evData pb.EventData
		evData.Id.Id = ev.Id
		evData.Name.Name = ev.Name
		evData.Date = ev.Date.Format("2006-01-02") // ?? check date format
		eventsPb = append(eventsPb, &evData)
	}
	return pb.PersonData{
		Id:     &pb.Id{Id: per.Id},
		Name:   &pb.Name{Name: per.Name},
		Events: eventsPb,
	}
}
