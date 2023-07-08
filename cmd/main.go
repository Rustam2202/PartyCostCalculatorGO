package main

import (
	"party-calc/internal/config"
	"party-calc/internal/database"
	"party-calc/internal/logger"
	"party-calc/internal/repository"
	"party-calc/internal/server"
	"party-calc/internal/server/handlers"
	"party-calc/internal/server/handlers/events"
	"party-calc/internal/server/handlers/persons"
	"party-calc/internal/server/handlers/persons_events"
	"party-calc/internal/service"
)

func main() {
	cfg := config.LoadConfig()
	logger.IntializeLogger(cfg.LoggerConfig)
	db := database.NewPGX(cfg.DatabaseConfig)

	personsRepo := repository.NewPersonRepository(db)
	eventsRepo := repository.NewEventRepository(db)
	persEventsRepo := repository.NewPersEventsRepository(db, personsRepo, eventsRepo)

	personService := service.NewPersonService(personsRepo)
	eventService := service.NewEventService(eventsRepo)
	persEventService := service.NewPersonsEventsService(persEventsRepo)
	calcService := service.NewCalcService(eventService, persEventService)

	personHandler := persons.NewPersonHandler(personService)
	eventHandler := events.NewEventHandler(eventService)
	persEventHandler := personsevents.NewPersEventsHandler(persEventService)
	calcHandler := handlers.NewCalcHandler(calcService)

	srv := server.NewServer(cfg.ServerConfig, personHandler, eventHandler, persEventHandler, calcHandler)
	srv.Start()
}
