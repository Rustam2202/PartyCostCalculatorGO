package event

import (
	"context"
	"party-calc/internal/server/grpc/proto"

	pm "github.com/golang/protobuf/proto"
)

func (h *EventHandler) AddEvent(ctx context.Context, pb *proto.EventCreateRequest) (*proto.Response, error) {
	msg, err := pm.Marshal(pb)
	if err != nil {
		return &proto.Response{Response: "Failed to add Event-create-request to kafka"}, err
	}
	err = h.producer.WriteMessage("event-create", msg)
	if err != nil {
		return &proto.Response{Response: "Failed to add Event-create-request to kafka"}, err
	}
	return &proto.Response{Response: "Event-create-request added to kafka"}, nil
}
