package personevent

import (
	"context"
	"fmt"
	"party-calc/internal/server/grpc/proto"
)

func (h *PersonEventHandler) AddPersonToPersonsEvent(ctx context.Context, pb *proto.PersonEventCreateRequest) (*proto.Response, error) {
	id, err := h.service.AddPersonToPersEvent(ctx, pb.PersonId, pb.EventId, pb.Spent, int(pb.Factor))
	if err != nil {
		return &proto.Response{Response: "Failed to add PersonEvent to db"}, err
	}
	return &proto.Response{Response: fmt.Sprintf("Person added to db with id: %d", id)}, nil
}
