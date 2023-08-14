package personevent

import (
	"context"
	"party-calc/internal/server/grpc/proto"

	pm "github.com/golang/protobuf/proto"
)

func (h *PersonEventHandler) AddPersonToPersonsEvent(ctx context.Context, pb *proto.PersonEventCreateRequest) (*proto.Response, error) {
	msg, err := pm.Marshal(pb)
	if err != nil {
		return &proto.Response{Response: "Failed to add PersonEvent-create-request to kafka"}, err
	}
	err = h.producer.WriteMessage("person-event-create", msg)
	if err != nil {
		return &proto.Response{Response: "Failed to add PersonEvent-create-request to kafka"}, err
	}
	return &proto.Response{Response: "PersonEvent-create-request added to kafka"}, nil
}
