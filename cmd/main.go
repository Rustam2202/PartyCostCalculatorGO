package main

import (
//	"party-calc/internal/config"
	"party-calc/internal/logger"
	"party-calc/internal/server"
)

func main() {
	logger.IntializeLogger()
//	config.LoadConfig()
	server.StartServer()
}
