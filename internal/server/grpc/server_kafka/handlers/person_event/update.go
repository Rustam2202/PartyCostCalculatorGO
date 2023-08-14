package personevent

import (
	"context"
	"party-calc/internal/server/grpc/proto"

	pm "github.com/golang/protobuf/proto"
)

func (h *PersonEventHandler) UpdatePersonsEvents(ctx context.Context, pb *proto.PersonEventUpdateRequest) (*proto.Response, error) {
	msg, err := pm.Marshal(pb)
	if err != nil {
		return &proto.Response{Response: "Failed to add PersonEvent-update-request to kafka"}, err
	}
	err = h.producer.WriteMessage("person-event-update", msg)
	if err != nil {
		return &proto.Response{Response: "Failed to add PersonEvent-update-request to kafka"}, err
	}
	return &proto.Response{Response: "PersonEvent-update-request added to kafka"}, nil
}
