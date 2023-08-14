package person

import (
	"context"
	"party-calc/internal/server/grpc/proto"

	pm "github.com/golang/protobuf/proto"
)

func (h *PersonHandler) UpdatePerson(ctx context.Context, pb *proto.PersonUpdateRequest) (*proto.Response, error) {
	msg, err := pm.Marshal(pb)
	if err != nil {
		return &proto.Response{Response: "Failed to add Person-update-request to kafka"}, err
	}
	err = h.producer.WriteMessage("person-update", msg)
	if err != nil {
		return &proto.Response{Response: "Failed to add Person-update-request to kafka"}, err
	}
	return &proto.Response{Response: "Person-update-request added to kafka"}, nil
}
