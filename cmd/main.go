package main

import (
	"party-calc/internal/config"
	"party-calc/internal/database"
	"party-calc/internal/database/repository"
	"party-calc/internal/logger"
	"party-calc/internal/server"
	"party-calc/internal/server/handlers"
	"party-calc/internal/service"
)

func main() {
	logger.IntializeLogger()
	cfg := config.LoadConfig()

	db := database.New(cfg.DatabaseConfig)

	personsRepo := repository.NewPersonRepository(db)
	eventsRepo := repository.NewEventRepository(db)
	persEventsRepo := repository.NewPersEventsRepository(db, personsRepo, eventsRepo)

	personService := service.NewPersonService(personsRepo)
	eventService := service.NewEventService(eventsRepo)
	persEventService := service.NewPersEventsService(persEventsRepo)
	calcService := service.NewCalcService(persEventsRepo, persEventService)

	personHandler := handlers.NewPersonHandler(personService)
	eventHandler := handlers.NewEventHandler(eventService)
	persEventHandler := handlers.NewPersEventsHandler(persEventService)
	calcHandler := handlers.NewCalcHandler(calcService)

	srv := server.NewServer(cfg.ServerConfig, personHandler, eventHandler, persEventHandler, calcHandler)
	srv.Start()
}
