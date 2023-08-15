package person

import (
	"context"
	"fmt"
	"party-calc/internal/server/grpc/proto"
)

func (h *PersonHandler) AddPerson(ctx context.Context, pb *proto.PersonCreateRequest) (*proto.Response, error) {
	id, err := h.service.NewPerson(ctx, pb.Name)
	if err != nil {
		return &proto.Response{Response: "Failed to add Person to db"}, err
	}
	return &proto.Response{Response: fmt.Sprintf("Person added to db with id: %d", id)}, nil
}
