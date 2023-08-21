package personevent

import (
	"context"
	k "party-calc/internal/kafka"
	"party-calc/internal/server/grpc/proto"

	pm "google.golang.org/protobuf/proto"
)

func (h *PersonEventHandler) AddPersonToPersonsEvent(ctx context.Context, pb *proto.PersonEventCreateRequest) (*proto.Response, error) {
	msg, err := pm.Marshal(pb)
	if err != nil {
		return &proto.Response{Response: "Failed to add PersonEvent-create-request to kafka"}, err
	}
	err = h.producer.WriteMessage(ctx, k.PersonEventCreate, msg)
	if err != nil {
		return &proto.Response{Response: "Failed to add PersonEvent-create-request to kafka"}, err
	}
	return &proto.Response{Response: "PersonEvent-create-request added to kafka"}, nil
}
