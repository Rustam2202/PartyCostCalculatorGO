package consumer

import (
	"context"
	"party-calc/internal/service"

	"github.com/segmentio/kafka-go"
	k "party-calc/internal/kafka"
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
		r.services.PersonService.NewPerson(ctx, string(msg.Value))
	}
}

func (r *KafkaConsumer) RunKafkaConsumer() {
	ctx := context.Background()
	go r.RunPersonCreateReader(ctx)
}
