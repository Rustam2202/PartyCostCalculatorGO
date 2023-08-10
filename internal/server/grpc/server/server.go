package main

import (
	"context"
	"log"
	"net"
	"party-calc/internal/server/grpc/proto"
	pb "party-calc/internal/server/grpc/proto"
	"party-calc/internal/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type PersonService struct {
	pb.PersonServiceServer
	service service.PersonService
}

func (s *PersonService) AddPerson(ctx context.Context, pb *proto.PersonCreate) (*proto.Id, error) {
	id, err := s.service.NewPerson(ctx, pb.Name)
	if err != nil {
		return nil, err
	}
	return &proto.Id{Id: id}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	//var pserv pb.PersonServiceServer
	pb.RegisterPersonServiceServer(s, &PersonService{})
	log.Printf("Starting gRPC listener on port " + port)
	s.Serve(lis)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
