package event

import (
	"context"
	k "party-calc/internal/kafka"
	"party-calc/internal/server/grpc/proto"

	pm "google.golang.org/protobuf/proto"
)

func (h *EventHandler) Update(ctx context.Context, pb *proto.EventUpdateRequest) (*proto.Response, error) {
	msg, err := pm.Marshal(pb)
	if err != nil {
		return &proto.Response{Response: "Failed to add Event-update-request to kafka"}, err
	}
	err = h.producer.WriteMessage(ctx, k.EventUpdate, msg)
	if err != nil {
		return &proto.Response{Response: "Failed to add Event-update-request to kafka"}, err
	}
	return &proto.Response{Response: "Event-update-request added to kafka"}, nil
}
