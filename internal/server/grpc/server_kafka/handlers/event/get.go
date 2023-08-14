package event

import (
	"context"
	"party-calc/internal/server/grpc/proto"
)

func (h *EventHandler) GetEvent(ctx context.Context, pb *proto.Id) (*proto.Event, error) {
	event, err := h.service.GetEventById(ctx, pb.Id)
	if err != nil {
		return nil, err
	}
	result := &proto.Event{
		Id:      event.Id,
		Name:    event.Name,
		Date:    event.Date.Format("2006-01-02"),
		Persons: []*proto.Person{},
	}
	for _, ev := range event.Persons {
		result.Persons = append(result.Persons, &proto.Person{
			Id:   ev.Id,
			Name: ev.Name,
		})
	}
	return result, nil
}
