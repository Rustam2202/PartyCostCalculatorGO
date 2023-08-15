package event

import (
	"context"
	"fmt"
	"party-calc/internal/server/grpc/proto"
	"time"
)

func (h *EventHandler) AddEvent(ctx context.Context, pb *proto.EventCreateRequest) (*proto.Response, error) {
	date, err := time.Parse("2006-01-02", pb.Date)
	if err != nil {
		return &proto.Response{Response: "Failed to parse date"}, err
	}
	id, err := h.service.NewEvent(ctx, pb.Name, date)
	if err != nil {
		return &proto.Response{Response: "Failed to add Event to db"}, err
	}
	return &proto.Response{Response: fmt.Sprintf("Event added to db with id: %d", id)}, nil
}
