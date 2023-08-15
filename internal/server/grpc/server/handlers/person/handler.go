package person

import (
	pb "party-calc/internal/server/grpc/proto"
	"party-calc/internal/service"
)

type PersonHandler struct {
	pb.PersonServiceServer
	service *service.PersonService
}

func NewPersonHandler(s *service.PersonService) *PersonHandler {
	return &PersonHandler{service: s}
}
