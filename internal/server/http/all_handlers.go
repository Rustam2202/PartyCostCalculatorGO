package http

import (
	"party-calc/internal/server/http/handlers/calculation"
	"party-calc/internal/server/http/handlers/events"
	"party-calc/internal/server/http/handlers/persons"
	personsevents "party-calc/internal/server/http/handlers/persons_events"
	"party-calc/internal/service"
)

type HTTPHandlers struct {
	PersonHandler     *persons.PersonHandler
	EventHandler      *events.EventHandler
	PersEventsHandler *personsevents.PersEventsHandler
	CalcHandler       *calculation.CalcHandler
}

func NewHTTPHandlers(services *service.Services) *HTTPHandlers {
	return &HTTPHandlers{
		PersonHandler:     persons.NewPersonHandler(services.PersonService),
		EventHandler:      events.NewEventHandler(services.EventService),
		PersEventsHandler: personsevents.NewPersEventsHandler(services.PersonEventService),
		CalcHandler:       calculation.NewCalcHandler(services.CalcService),
	}
}
