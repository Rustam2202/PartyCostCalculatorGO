package event

import (
	"context"
	"party-calc/internal/server/grpc/proto"
)

func (h *EventHandler) DeleteEvent(ctx context.Context, pb *proto.Id) (*proto.Response, error) {
	err := h.service.DeleteEventById(ctx, pb.Id)
	if err != nil {
		return &proto.Response{Response: "Failed to delete Event from db"}, err
	}
	return &proto.Response{Response: "Event deleted from db"}, nil
}
