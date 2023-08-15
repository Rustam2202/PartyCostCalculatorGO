package person

import (
	"context"
	"party-calc/internal/server/grpc/proto"
)

func (h *PersonHandler) DeletePerson(ctx context.Context, pb *proto.Id) (*proto.Response, error) {
	err := h.service.DeletePersonById(ctx, pb.Id)
	if err != nil {
		return &proto.Response{Response: "Failed to delete Person from db"}, err
	}
	return &proto.Response{Response: "Person deleted from db"}, nil
}
