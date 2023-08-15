package person

import (
	"context"
	"party-calc/internal/server/grpc/proto"
)

func (h *PersonHandler) UpdatePerson(ctx context.Context, pb *proto.PersonUpdateRequest) (*proto.Response, error) {
	err := h.service.UpdatePerson(ctx, pb.Id, pb.Name)
	if err != nil {
		return &proto.Response{Response: "Failed to update Person in db"}, err
	}
	return &proto.Response{Response: "Person updated in db"}, nil
}
