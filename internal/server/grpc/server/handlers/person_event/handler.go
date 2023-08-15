package personevent

import (
	pb "party-calc/internal/server/grpc/proto"
	"party-calc/internal/service"
)

type PersonEventHandler struct {
	pb.PersonsEventsServiceServer
	service *service.PersonsEventsService
}

func NewPersonEventHandler(s *service.PersonsEventsService) *PersonEventHandler {
	return &PersonEventHandler{service: s}
}
