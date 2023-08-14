package personsevents

import "party-calc/internal/service"

type PersEventsHandler struct {
	service *service.PersonsEventsService
}

func NewPersEventsHandler(s *service.PersonsEventsService) *PersEventsHandler {
	return &PersEventsHandler{service: s}
}
