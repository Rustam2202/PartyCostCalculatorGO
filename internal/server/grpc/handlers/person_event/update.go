package personevent

import (
	"context"
	k "party-calc/internal/kafka"
	"party-calc/internal/server/grpc/proto"

	pm "google.golang.org/protobuf/proto"
)

func (h *PersonEventHandler) Update(ctx context.Context, pb *proto.PersonEventUpdateRequest) (*proto.Response, error) {
	msg, err := pm.Marshal(pb)
	if err != nil {
		return &proto.Response{Response: "Failed to add PersonEvent-update-request to kafka"}, err
	}
	err = h.producer.WriteMessage(ctx, k.PersonEventUpdate, msg)
	if err != nil {
		return &proto.Response{Response: "Failed to add PersonEvent-update-request to kafka"}, err
	}
	return &proto.Response{Response: "PersonEvent-update-request added to kafka"}, nil
}
