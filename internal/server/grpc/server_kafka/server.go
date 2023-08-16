package serverkafka

import (
	"fmt"
	"net"
	"party-calc/internal/logger"
	pb "party-calc/internal/server/grpc/proto"

	"party-calc/internal/server/grpc/server_kafka/handlers/calculation"
	"party-calc/internal/server/grpc/server_kafka/handlers/event"
	"party-calc/internal/server/grpc/server_kafka/handlers/person"
	personevent "party-calc/internal/server/grpc/server_kafka/handlers/person_event"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	cfg                 *ServerGrpcKafkaConfig
	personHandler       *person.PersonHandler
	eventHandler        *event.EventHandler
	personsEventHandler *personevent.PersonEventHandler
	calcHandler         *calculation.CalcHandler
}

func NewServer(
	cfg *ServerGrpcKafkaConfig,
	handlers *GRPCKafkaHandlers,
) *Server {
	return &Server{
		cfg:                 cfg,
		personHandler:       handlers.PersonHandler,
		eventHandler:        handlers.EventHandler,
		personsEventHandler: handlers.PersEventsHandler,
		calcHandler:         handlers.CalcHandler,
	}
}

func (s *Server) Start() {
	lis, err := net.Listen(s.cfg.Network, fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port))
	if err != nil {
		logger.Logger.Error("Failed to listen", zap.Error(err))
		return
	}

	srv := grpc.NewServer()
	reflection.Register(srv)

	pb.RegisterPersonServiceServer(srv, s.personHandler)
	pb.RegisterEventServiceServer(srv, s.eventHandler)
	pb.RegisterPersonsEventsServiceServer(srv, s.personsEventHandler)
	pb.RegisterCalculationServer(srv, s.calcHandler)

	if err := srv.Serve(lis); err != nil {
		logger.Logger.Error("Failed to serve", zap.Error(err))
		return
	} else {
		logger.Logger.Info("GRPC server with Kafka started")
	}
}
