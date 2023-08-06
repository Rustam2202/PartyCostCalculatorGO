package grpc

import (
	"context"
	"log"
	pb "party-calc/internal/grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"

)

const (
	address = "localhost:50051"
)

func client() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPersonClient(conn)
	c.AddPerson(context.Background(),&pb.Name{})

	proto.Marshal()
}
