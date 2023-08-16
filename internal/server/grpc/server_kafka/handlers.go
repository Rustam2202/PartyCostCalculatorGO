package serverkafka

import (
	"party-calc/internal/kafka/producer"
	"party-calc/internal/server/grpc/server_kafka/handlers/calculation"
	"party-calc/internal/server/grpc/server_kafka/handlers/event"
	"party-calc/internal/server/grpc/server_kafka/handlers/person"
	personevent "party-calc/internal/server/grpc/server_kafka/handlers/person_event"
	"party-calc/internal/service"
)

type GRPCKafkaHandlers struct {
	PersonHandler     *person.PersonHandler
	EventHandler      *event.EventHandler
	PersEventsHandler *personevent.PersonEventHandler
	CalcHandler       *calculation.CalcHandler
}

func NewGRPCKafkaHandlers(services *service.Services, p *producer.KafkaProducer) *GRPCKafkaHandlers {
	return &GRPCKafkaHandlers{
		PersonHandler:     person.NewPersonHandler(services.PersonService, p),
		EventHandler:      event.NewEventHandler(services.EventService, p),
		PersEventsHandler: personevent.NewPersonEventHandler(services.PersonEventService, p),
		CalcHandler:       calculation.NewCalcHandler(services.CalcService),
	}
}
