package event

import (
	pb "party-calc/internal/server/grpc/proto"
	"party-calc/internal/service"
)

type EventHandler struct {
	pb.EventServiceServer
	service *service.EventService
}

func NewEventHandler(s *service.EventService) *EventHandler {
	return &EventHandler{service: s}
}
