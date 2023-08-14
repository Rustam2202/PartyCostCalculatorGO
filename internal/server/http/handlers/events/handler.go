package events

import "party-calc/internal/service"

type EventHandler struct {
	service *service.EventService
}

func NewEventHandler(s *service.EventService) *EventHandler {
	return &EventHandler{service: s}
}
