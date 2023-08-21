package consumer

import (
	"context"
	k "party-calc/internal/kafka"
	"party-calc/internal/service"

	"github.com/segmentio/kafka-go"
)

type KafkaConsumer struct {
	services *service.Services
	cfg      *kafka.ReaderConfig
}

func NewKafkaConsumer(cfg k.KafkaConfig, s *service.Services) *KafkaConsumer {
	config := kafka.ReaderConfig{}
	config.Brokers = cfg.Brokers
	config.GroupID = cfg.Group
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
