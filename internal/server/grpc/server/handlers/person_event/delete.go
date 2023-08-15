package personevent

import (
	"context"
	"party-calc/internal/server/grpc/proto"
)

func (h *PersonEventHandler) DeletePersonsEvents(ctx context.Context, pb *proto.Id) (*proto.Response, error) {
	err := h.service.Delete(ctx, pb.Id)
	if err != nil {
		return &proto.Response{Response: "Failed to delete PersonEvent from db"}, err
	}
	return &proto.Response{Response: "PersonEvent deleted from db"}, nil
}
