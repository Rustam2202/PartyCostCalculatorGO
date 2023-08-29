package person

import (
	"context"
	k "party-calc/internal/kafka"
	"party-calc/internal/server/grpc/proto"

	pm "google.golang.org/protobuf/proto"
)

func (h *PersonHandler) Create(ctx context.Context, pb *proto.PersonCreateRequest) (*proto.Response, error) {
	msg, err := pm.Marshal(pb)
	if err != nil {
		return &proto.Response{Response: "Failed to add Person-create-request to kafka"}, err
	}
	err = h.producer.WriteMessage(ctx, k.PersonCreate, msg)
	if err != nil {
		return &proto.Response{Response: "Failed to add Person-create-request to kafka"}, err
	}
	return &proto.Response{Response: "Person-create-request added to kafka"}, nil
}
