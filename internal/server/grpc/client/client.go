package main

import (
	"context"
	"fmt"
	"log"
	pb "party-calc/internal/server/grpc/proto"

	"google.golang.org/grpc"
)

type Client struct {
	
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	personClient := pb.NewPersonServiceClient(conn)
	// eventClient := pb.NewEventServiceClient(conn)
	// personEventsClient := pb.NewPersonsEventsServiceClient(conn)
	// calculationClient := pb.NewCalculationClient(conn)

	response, err := personClient.AddPerson(context.Background(), &pb.PersonCreateRequest{Name: "Person"})
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	fmt.Println(response)
}
