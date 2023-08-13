package consumer

import (
	"context"
	"party-calc/internal/server/grpc/proto"
	"party-calc/internal/service"

	k "party-calc/internal/kafka"

	pm "github.com/golang/protobuf/proto"
	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	services *service.Services
	cfg      *kafka.ReaderConfig
}

func NewKafkaConsumer(cfg k.KafkaConfig, s *service.Services) *KafkaConsumer {
	config := kafka.ReaderConfig{}
	for _, broker := range cfg.Brokers {
		config.Brokers = append(config.Brokers, broker)
	}
	return &KafkaConsumer{services: s, cfg: &config}
}

func (r *KafkaConsumer) RunPersonCreateReader(ctx context.Context) {
	cfg := *r.cfg
	cfg.Topic = "person-create"
	reader := kafka.NewReader(cfg)
	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			break
		}
		personCreate := proto.PersonCreateRequest{}
		pm.Unmarshal(msg.Value, &personCreate)
		r.services.PersonService.NewPerson(ctx, personCreate.Name)
	}
}

func (r *KafkaConsumer) RunKafkaConsumer() {
	ctx := context.Background()
	go r.RunPersonCreateReader(ctx)
}
