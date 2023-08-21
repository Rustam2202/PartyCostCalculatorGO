package event

import (
	"context"
	k "party-calc/internal/kafka"
	"party-calc/internal/server/grpc/proto"

	pm "google.golang.org/protobuf/proto"
)

func (h *EventHandler) DeleteEvent(ctx context.Context, pb *proto.Id) (*proto.Response, error) {
	msg, err := pm.Marshal(pb)
	if err != nil {
		return &proto.Response{Response: "Failed to add Event-delete-request to kafka"}, err
	}
	err = h.producer.WriteMessage(ctx, k.EventDelete, msg)
	if err != nil {
		return &proto.Response{Response: "Failed to add Event-delete-request to kafka"}, err
	}
	return &proto.Response{Response: "Event-delete-request added to kafka"}, nil
}
