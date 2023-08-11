package personevent

import (
	"context"
	"party-calc/internal/server/grpc/proto"
	pb "party-calc/internal/server/grpc/proto"
	"party-calc/internal/service"
)

type PersonEventHandler struct {
	pb.PersonsEventsServiceServer
	service *service.PersonsEventsService
}

func NewPersonEventHandler(s *service.PersonsEventsService) *PersonEventHandler {
	return &PersonEventHandler{service: s}
}

func (h *PersonEventHandler) AddPersonToPersonsEvent(ctx context.Context, pb *proto.PersonEventCreateRequest) (*proto.Id, error) {
	id, err := h.service.AddPersonToPersEvent(ctx, pb.PersonId, pb.EventId, pb.Spent, int(pb.Factor))
	if err != nil {
		return nil, err
	}
	return &proto.Id{Id: id}, nil
}

func (h *PersonEventHandler) GetByPersonId(ctx context.Context, pb *proto.Id) (*proto.PersonEventsGetResponse, error) {
	perEv, err := h.service.GetByPersonId(ctx, pb.Id)
	if err != nil {
		return nil, err
	}
	result := proto.PersonEventsGetResponse{}
	for _, pe := range perEv {
		persEvsToPB := proto.PersonEvent{
			Id:       pe.Id,
			PersonId: pe.PersonId,
			EventId:  pe.EventId,
			Spent:    float32(pe.Spent),
			Factor:   int32(pe.Factor),
			Person: &proto.Person{
				Id:   pe.Person.Id,
				Name: pe.Person.Name,
			},
			Event: &proto.Event{
				Id:   pe.Event.Id,
				Name: pe.Event.Name,
				Date: pe.Event.Date.Format("2006-01-02"),
			},
		}
		result.PersonsEvents = append(result.PersonsEvents, &persEvsToPB)
	}
	return &result, nil
}

func (h *PersonEventHandler) UpdatePersonsEvents(ctx context.Context, pb *proto.PersonEventUpdateRequest) (*proto.Empty, error) {
	err := h.service.Update(ctx, pb.Id, pb.PersonId, pb.EventId, pb.Spent, int(pb.Factor))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *PersonEventHandler) DeletePersonsEvents(ctx context.Context, pb *proto.Id) (*proto.Empty, error) {
	err := h.service.Delete(ctx, pb.Id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
