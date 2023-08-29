package personevent

import (
	"context"
	k "party-calc/internal/kafka"
	"party-calc/internal/server/grpc/proto"

	pm "google.golang.org/protobuf/proto"
)

func (h *PersonEventHandler) Delete(ctx context.Context, pb *proto.Id) (*proto.Response, error) {
	msg, err := pm.Marshal(pb)
	if err != nil {
		return &proto.Response{Response: "Failed to add PersonEvent-delete-request to kafka"}, err
	}
	err = h.producer.WriteMessage(ctx, k.PersonEventDelete, msg)
	if err != nil {
		return &proto.Response{Response: "Failed to add PersonEvent-delete-request to kafka"}, err
	}
	return &proto.Response{Response: "PersonEvent-delete-request added to kafka"}, nil
}
