package producer

import (
	"context"
	k "party-calc/internal/kafka"

	"github.com/segmentio/kafka-go"
)

type KafkaProducer struct {
	Writer *kafka.Writer
	cfg    *kafka.WriterConfig
}

func NewKafkaProducer(cfg k.KafkaConfig) *KafkaProducer {
	config := kafka.WriterConfig{Brokers: cfg.Brokers}
	return &KafkaProducer{cfg: &config}
}

func (w *KafkaProducer) WriteMessage(ctx context.Context, topic string, msg []byte) error {
	writer := kafka.NewWriter(*w.cfg)
	m := kafka.Message{
		Topic: topic,
		Value: msg,
	}
	err := writer.WriteMessages(ctx, m)
	if err != nil {
		return err
	}
	return nil
}
