package main

import (
	"context"
	"fmt"
	"party-calc/internal/service-grpc/proto"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/protocol"
)

type Reader struct {
	Reader *kafka.Reader
}

func NewKafkaReader() *Reader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "partycalc",
		GroupID: "group",
	})

	return &Reader{
		Reader: reader,
	}
}

func main() {
	ctx := context.Background()
	reader := NewKafkaReader()
	msg, err := reader.Reader.ReadMessage(ctx)
	if err != nil {
		fmt.Println(err)
	}
	var per proto.PersonData
	err = protocol.Unmarshal(msg.Value, 3, &per)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(per.Id,per.Name)
}
