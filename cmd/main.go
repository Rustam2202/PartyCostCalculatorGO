package main

import (
	"party-calc/internal/config"
	"party-calc/internal/database"
	"party-calc/internal/kafka/consumer"
	"party-calc/internal/kafka/producer"
	"party-calc/internal/logger"
	"party-calc/internal/repository"
	grpc "party-calc/internal/server/grpc/server"
	serverkafka "party-calc/internal/server/grpc/server_kafka"
	"party-calc/internal/server/http"
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

	services := service.NewServices(personsRepo, eventsRepo, persEventsRepo)

	httpServer := http.NewServer(cfg.ServerHTTPConfig, http.NewHTTPHandlers(services))
	go httpServer.Start()

	grpcServer := grpc.NewServer(&cfg.ServerGrpcConfig, *grpc.NewGRPCHandlers(services))
	go grpcServer.Start()

	kafkaConsumer := consumer.NewKafkaConsumer(cfg.KafkaConfig, services)
	go kafkaConsumer.RunKafkaConsumer()

	kafkaProducer := producer.NewKafkaProducer(cfg.KafkaConfig)
	grpsKafkaServer := serverkafka.NewServer(&cfg.ServerGrpcKafkaConfig,
		serverkafka.NewGRPCKafkaHandlers(services, kafkaProducer))
	grpsKafkaServer.Start()
}
