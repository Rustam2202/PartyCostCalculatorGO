package person

import (
	"context"
	"party-calc/internal/server/grpc/proto"
)

func (h *PersonHandler) Get(ctx context.Context, pb *proto.Id) (*proto.Person, error) {
	person, err := h.service.GetPersonById(ctx, pb.Id)
	if err != nil {
		return &proto.Person{}, err
	}
	result := &proto.Person{
		Id:     person.Id,
		Name:   person.Name,
		Events: []*proto.Event{},
	}
	for _, ev := range person.Events {
		result.Events = append(result.Events, &proto.Event{
			Id:   ev.Id,
			Name: ev.Name,
			Date: ev.Date.Format("2006-01-02"),
		})
	}
	return result, nil
}
