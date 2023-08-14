package consumer

import (
	"context"
	"party-calc/internal/service"

	k "party-calc/internal/kafka"

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
	config.GroupID = "group"
	return &KafkaConsumer{services: s, cfg: &config}
}

func (r *KafkaConsumer) RunKafkaConsumer() {
	ctx := context.Background()
	go r.RunPersonCreateReader(ctx)
	go r.RunPersonUpdateReader(ctx)
	go r.RunPersonDeleteReader(ctx)
	go r.RunEventCreateReader(ctx)
	go r.RunEventUpdateReader(ctx)
	go r.RunEventDeleteReader(ctx)
	go r.RunPersonEventCreateReader(ctx)
	go r.RunPersonEventUpdateReader(ctx)
	go r.RunPersonEventDeleteReader(ctx)
}
