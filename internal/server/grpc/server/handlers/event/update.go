package event

import (
	"context"
	"party-calc/internal/server/grpc/proto"
	"time"
)

func (h *EventHandler) UpdateEvent(ctx context.Context, pb *proto.EventUpdate) (*proto.Response, error) {
	date, err := time.Parse("2006-01-02", pb.Date)
	if err != nil {
		return &proto.Response{Response: "Failed to parse date"}, err
	}
	err = h.service.UpdateEvent(ctx, pb.Id, pb.Name, date)
	if err != nil {
		return &proto.Response{Response: "Failed to add Event to db"}, err
	}
	return &proto.Response{Response: "Event updated in db"}, nil
}
