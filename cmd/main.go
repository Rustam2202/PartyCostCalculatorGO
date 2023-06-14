package main

import (
	"fmt"
	"party-calc/internal/config"
	"party-calc/internal/logger"
	"party-calc/internal/server"
)

func main() {
	logger.IntializeLogger()
	cfg := config.LoadConfig()
	fmt.Println(cfg)
	server.StartServer()
}
