package main

import (
	"party-calc/internal/config"
	"party-calc/internal/database"
	"party-calc/internal/database/repository"
	"party-calc/internal/logger"
	"party-calc/internal/server"
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
	personHandler := NewPersonHandler(personService)

	srv := server.NewServer(cfg.ServerConfig)
	srv.Start()
}
