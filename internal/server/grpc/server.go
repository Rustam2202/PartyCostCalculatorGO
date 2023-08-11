package grpc

import (
	"net"
	"party-calc/internal/logger"
	pb "party-calc/internal/server/grpc/proto"

	"party-calc/internal/server/grpc/server/handlers/calculation"
	"party-calc/internal/server/grpc/server/handlers/event"
	"party-calc/internal/server/grpc/server/handlers/person"
	personevent "party-calc/internal/server/grpc/server/handlers/person_event"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	personHandler       *person.PersonHandler
	eventHandler        *event.EventHandler
	personsEventHandler *personevent.PersonEventHandler
	calcHandler         *calculation.CalcHandler
}

func NewServer(
	ph *person.PersonHandler,
	eh *event.EventHandler,
	peh *personevent.PersonEventHandler,
	ch *calculation.CalcHandler,
) *Server {
	return &Server{
		personHandler:       ph,
		eventHandler:        eh,
		personsEventHandler: peh,
		calcHandler:         ch,
	}
}

func (s *Server) Start() {
	lis, err := net.Listen("tcp", "localhost:50051")
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
	}
}
