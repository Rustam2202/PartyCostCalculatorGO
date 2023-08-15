package personevent

import (
	"context"
	"party-calc/internal/server/grpc/proto"
)

func (h *PersonEventHandler) UpdatePersonsEvents(ctx context.Context, pb *proto.PersonEventUpdateRequest) (*proto.Response, error) {
	err := h.service.Update(ctx, pb.Id, pb.PersonId, pb.EventId, pb.Spent, int(pb.Factor))
	if err != nil {
		return &proto.Response{Response: "Failed to update PersonEvent in db"}, err
	}
	return &proto.Response{Response: "PersonEvent updated in db"}, nil
}
