package main

import (
	"party-calc/internal/config"
	"party-calc/internal/database"
	"party-calc/internal/logger"
	"party-calc/internal/server"
	"party-calc/internal/service"
)

func main() {
	logger.IntializeLogger()
	cfg := config.LoadConfig()

	db := database.New(cfg.DatabaseConfig)

	personRepo := database.NewPersonRepository(db)
	personService := service.NewPersonService(personRepo)
	personHandler := server.NewPersonHandler(personService)

	srv := server.NewServer(cfg.ServerConfig, personHandler)
	srv.Start()
}
