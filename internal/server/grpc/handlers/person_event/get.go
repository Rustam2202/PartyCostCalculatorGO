package personevent

import (
	"context"
	"party-calc/internal/server/grpc/proto"
)

func (h *PersonEventHandler) Get(ctx context.Context, pb *proto.Id) (*proto.PersonEvent, error) {
	perEv, err := h.service.Get(ctx, pb.Id)
	if err != nil {
		return nil, err
	}
	result := proto.PersonEvent{
		Id:       perEv.Id,
		PersonId: perEv.PersonId,
		EventId:  perEv.EventId,
		Spent:    float32(perEv.Spent),
		Factor:   int32(perEv.Factor),
		Person: &proto.Person{
			Id:     perEv.Person.Id,
			Name:   perEv.Person.Name,
			Events: []*proto.Event{},
		},
		Event: &proto.Event{
			Id:      perEv.Event.Id,
			Name:    perEv.Event.Name,
			Date:    perEv.Event.Date.Format("2006-01-02"),
			Persons: []*proto.Person{},
		},
	}
	return &result, nil
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
