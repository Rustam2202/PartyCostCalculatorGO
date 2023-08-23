package grpc

import (
	"context"
	"fmt"
	"net"
	"party-calc/internal/logger"
	pb "party-calc/internal/server/grpc/proto"
	"sync"

	"party-calc/internal/server/grpc/handlers/calculation"
	"party-calc/internal/server/grpc/handlers/event"
	"party-calc/internal/server/grpc/handlers/person"
	personevent "party-calc/internal/server/grpc/handlers/person_event"

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

func (s *Server) Start(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
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

	go func() {
		logger.Logger.Info("Starting GRPC server ...")
		if err := srv.Serve(lis); err != nil {
			logger.Logger.Error("Failed to start GRPC server", zap.Error(err))
			return
		}
	}()

	<-ctx.Done()
	logger.Logger.Info("Shutting down GRPC server ...")
	srv.GracefulStop()
}
