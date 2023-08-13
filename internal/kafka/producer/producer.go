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
	// for _, broker := range cfg.Brokers {
	// 	config.Brokers = append(config.Brokers, broker)
	// }
	return &KafkaProducer{cfg: &config}
}

func (w *KafkaProducer) WriteMessage(topic string, msg []byte) {
	writer := kafka.NewWriter(*w.cfg)
	m := kafka.Message{
		Topic: topic,
		Value: msg,
	}
	err := writer.WriteMessages(context.Background(), m)
	if err != nil {

	}
}
