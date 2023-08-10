package main

import (
	"context"
	"fmt"
	"log"
	pb "party-calc/internal/server/grpc/proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewPersonServiceClient(conn)
	response, err := client.AddPerson(context.Background(), &pb.PersonCreate{Name: "Person"})
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	fmt.Println(response)
}
