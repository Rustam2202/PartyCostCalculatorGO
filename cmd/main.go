package main

import (
	"party-calc/internal/config"
	"party-calc/internal/database"
	"party-calc/internal/logger"
	"party-calc/internal/repository"
	"party-calc/internal/server/grpc"
	"party-calc/internal/server/http"
	http_calc "party-calc/internal/server/http/handlers/calculation"
	http_ev "party-calc/internal/server/http/handlers/events"
	http_per "party-calc/internal/server/http/handlers/persons"
	http_per_ev "party-calc/internal/server/http/handlers/persons_events"

	grpc_calc "party-calc/internal/server/grpc/server/handlers/calculation"
	grpc_ev "party-calc/internal/server/grpc/server/handlers/event"
	grpc_per "party-calc/internal/server/grpc/server/handlers/person"
	grpc_per_ev "party-calc/internal/server/grpc/server/handlers/person_event"

	"party-calc/internal/kafka/consumer"
	"party-calc/internal/service"
)

func main() {
	cfg := config.LoadConfig()
	logger.IntializeLogger(cfg.LoggerConfig)
	db, err := database.NewPGX(cfg.DatabaseConfig)
	if err != nil {
		return
	}

	personsRepo := repository.NewPersonRepository(db)
	eventsRepo := repository.NewEventRepository(db)
	persEventsRepo := repository.NewPersonsEventsRepository(db)

	personService := service.NewPersonService(personsRepo)
	eventService := service.NewEventService(eventsRepo)
	persEventService := service.NewPersonsEventsService(persEventsRepo)
	calcService := service.NewCalcService(personService, eventService, persEventService)

	personHTTPHandler := http_per.NewPersonHandler(personService)
	eventHTTPHandler := http_ev.NewEventHandler(eventService)
	personEventHTTPHandler := http_per_ev.NewPersEventsHandler(persEventService)
	calcHTTPHandler := http_calc.NewCalcHandler(calcService)

	personGRPCHandler := grpc_per.NewPersonHandler(personService)
	eventGRPCHandler := grpc_ev.NewEventHandler(eventService)
	personEventGRPCHandler := grpc_per_ev.NewPersonEventHandler(persEventService)
	calcGRPCHandler := grpc_calc.NewCalcHandler(calcService)

	httpServer := http.NewServer(cfg.ServerConfig, personHTTPHandler, eventHTTPHandler, personEventHTTPHandler, calcHTTPHandler)
	go httpServer.Start()

	services := service.NewServices(personsRepo, eventsRepo, persEventsRepo)
	kafkaConsumer := consumer.NewKafkaConsumer(cfg.KafkaConfig, services)
	kafkaConsumer.RunKafkaConsumer()

	grpcServer := grpc.NewServer(personGRPCHandler, eventGRPCHandler, personEventGRPCHandler, calcGRPCHandler)
	grpcServer.Start()
}
