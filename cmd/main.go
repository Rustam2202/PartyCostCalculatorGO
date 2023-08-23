package main

import (
	"context"
	"os/signal"
	"party-calc/internal/config"
	"party-calc/internal/database"
	"party-calc/internal/kafka/consumer"
	"party-calc/internal/kafka/producer"
	"party-calc/internal/logger"
	"party-calc/internal/repository"
	"party-calc/internal/server/grpc"
	"party-calc/internal/server/http"
	"party-calc/internal/service"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	//_ = cancel
	defer cancel()

	wg := &sync.WaitGroup{}

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
	wg.Add(1)
	go httpServer.Start(ctx, wg)

	kafkaConsumer := consumer.NewKafkaConsumer(cfg.KafkaConfig, services)
	wg.Add(1)
	go kafkaConsumer.RunKafkaConsumer(ctx, wg)

	kafkaProducer := producer.NewKafkaProducer(cfg.KafkaConfig)
	grpcServer := grpc.NewServer(&cfg.ServerGrpcKafkaConfig,
		grpc.NewGRPCKafkaHandlers(services, kafkaProducer))
	wg.Add(1)
	go grpcServer.Start(ctx, wg)

	time.Sleep(2 * time.Second)
	cancel()

	<-ctx.Done()
	//	wg.Wait()
}
