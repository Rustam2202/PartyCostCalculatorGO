package main

import (
	"context"
	"party-calc/internal/service-grpc/proto"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/protocol"
)

type Writer struct {
	Writer *kafka.Writer
}

func NewKafkaWriter() *Writer {
	w := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: "partycalc",
	}
	return &Writer{Writer: w}
}

func (w *Writer) WriteMessages(ctx context.Context) error {
	per := proto.PersonData{
		Id:     &proto.Id{Id: 3},
		Name:   &proto.Name{Name: "Person 3"},
		Events: []*proto.EventData{},
	}

	//id := proto.Id{Id: 3}
	msgid, err := protocol.Marshal(3, per.String())
	if err != nil {
		return err
	}
	//	msg, err := protocol.Marshal(3,per)
	if err != nil {
		return err
	}
	err = w.Writer.WriteMessages(ctx, kafka.Message{Value: msgid})
	if err != nil {
		return err
	}
	return nil
}

func main() {

	writer := NewKafkaWriter()
	ctx := context.Background()

	writer.WriteMessages(ctx)

	// conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "topic_test", 0)
	// if err != nil {
	// }
	// conn.SetWriteDeadline(time.Now().Add(time.Second * 10))
	// conn.WriteMessages(kafka.Message{
	// 	Value: []byte(`Hello from VSCode`),
	// })

}
