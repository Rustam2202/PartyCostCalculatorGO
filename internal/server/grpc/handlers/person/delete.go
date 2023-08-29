package person

import (
	"context"
	k "party-calc/internal/kafka"
	"party-calc/internal/server/grpc/proto"

	pm "google.golang.org/protobuf/proto"
)

func (h *PersonHandler) Delete(ctx context.Context, pb *proto.Id) (*proto.Response, error) {
	msg, err := pm.Marshal(pb)
	if err != nil {
		return &proto.Response{Response: "Failed to add Person-delete-request to kafka"}, err
	}
	err = h.producer.WriteMessage(ctx, k.PersonDelete, msg)
	if err != nil {
		return &proto.Response{Response: "Failed to add Person-delete-request to kafka"}, err
	}
	return &proto.Response{Response: "Person-delete-request added to kafka"}, nil
}
