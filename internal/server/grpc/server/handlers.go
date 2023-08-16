package server

import (
	"party-calc/internal/server/grpc/server/handlers/person"
	"party-calc/internal/server/grpc/server/handlers/event"
	personevent "party-calc/internal/server/grpc/server/handlers/person_event"
	"party-calc/internal/server/grpc/server/handlers/calculation"
	"party-calc/internal/service"
)

type GRPCHandlers struct {
	PersonHandler     *person.PersonHandler
	EventHandler      *event.EventHandler
	PersEventsHandler *personevent.PersonEventHandler
	CalcHandler       *calculation.CalcHandler
}

func NewGRPCHandlers(services *service.Services) *GRPCHandlers {
	return &GRPCHandlers{
		PersonHandler:     person.NewPersonHandler(services.PersonService),
		EventHandler:      event.NewEventHandler(services.EventService),
		PersEventsHandler: personevent.NewPersonEventHandler(services.PersonEventService),
		CalcHandler:       calculation.NewCalcHandler(services.CalcService),
	}
}
