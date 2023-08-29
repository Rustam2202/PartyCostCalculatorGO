package event

import (
	"context"
	k "party-calc/internal/kafka"
	"party-calc/internal/server/grpc/proto"

	pm "google.golang.org/protobuf/proto"
)

func (h *EventHandler) Create(ctx context.Context, pb *proto.EventCreateRequest) (*proto.Response, error) {
	msg, err := pm.Marshal(pb)
	if err != nil {
		return &proto.Response{Response: "Failed to add Event-create-request to kafka"}, err
	}
	err = h.producer.WriteMessage(ctx, k.EventCreate, msg)
	if err != nil {
		return &proto.Response{Response: "Failed to add Event-create-request to kafka"}, err
	}
	return &proto.Response{Response: "Event-create-request added to kafka"}, nil
}
