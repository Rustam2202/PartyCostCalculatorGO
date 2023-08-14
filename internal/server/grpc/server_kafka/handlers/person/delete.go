package person

import (
	"context"
	"party-calc/internal/server/grpc/proto"

	pm "github.com/golang/protobuf/proto"
)

func (h *PersonHandler) DeletePerson(ctx context.Context, pb *proto.Id) (*proto.Response, error) {
	msg, err := pm.Marshal(pb)
	if err != nil {
		return &proto.Response{Response: "Failed to add Person-delete-request to kafka"}, err
	}
	err = h.producer.WriteMessage("person-delete", msg)
	if err != nil {
		return &proto.Response{Response: "Failed to add Person-delete-request to kafka"}, err
	}
	return &proto.Response{Response: "Person-delete-request added to kafka"}, nil
}
