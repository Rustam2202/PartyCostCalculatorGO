package event

import (
	"context"
	"party-calc/internal/server/grpc/proto"
	pb "party-calc/internal/server/grpc/proto"
	"party-calc/internal/service"
	"time"
)

type EventHandler struct {
	pb.EventServiceServer
	service *service.EventService
}

func NewEventHandler(s *service.EventService) *EventHandler {
	return &EventHandler{service: s}
}

func (h *EventHandler) AddEvent(ctx context.Context, pb *proto.EventCreateRequest) (*proto.Id, error) {
	date, err := time.Parse("2006-01-02", pb.Date)
	if err != nil {
		return nil, err
	}
	id, err := h.service.NewEvent(ctx, pb.Name, date)
	if err != nil {
		return nil, err
	}
	return &proto.Id{Id: id}, nil
}

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

func (h *EventHandler) UpdateEvent(ctx context.Context, pb *proto.EventUpdate) (*proto.Empty, error) {
	date, err := time.Parse("2006-01-02", pb.Date)
	if err != nil {
		return nil, err
	}
	err = h.service.UpdateEvent(ctx, pb.Id, pb.Name, date)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *EventHandler) DeleteEvent(ctx context.Context, pb *proto.Id) (*proto.Empty, error) {
	err := h.service.DeleteEventById(ctx, pb.Id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
