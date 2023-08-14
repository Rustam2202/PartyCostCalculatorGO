package event

import (
	"context"
	"party-calc/internal/server/grpc/proto"

	pm "github.com/golang/protobuf/proto"
)

func (h *EventHandler) UpdateEvent(ctx context.Context, pb *proto.EventUpdate) (*proto.Response, error) {
	msg, err := pm.Marshal(pb)
	if err != nil {
		return &proto.Response{Response: "Failed to add Event-update-request to kafka"}, err
	}
	err = h.producer.WriteMessage("event-update", msg)
	if err != nil {
		return &proto.Response{Response: "Failed to add Event-update-request to kafka"}, err
	}
	return &proto.Response{Response: "Event-update-request added to kafka"}, nil
}
