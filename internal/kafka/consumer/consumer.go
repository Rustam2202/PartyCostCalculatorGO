package consumer

import (
	"context"
	k "party-calc/internal/kafka"
	"party-calc/internal/service"
	"sync"

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

func (r *KafkaConsumer) RunKafkaConsumer(ctx context.Context, wg *sync.WaitGroup) {
	TopicsServe := map[string]func(context.Context, kafka.Message) error{
		k.PersonCreate:      r.personCreateServe,
		k.PersonUpdate:      r.PersonUpdateServe,
		k.PersonDelete:      r.personDeleteServe,
		k.EventCreate:       r.eventCreateServe,
		k.EventUpdate:       r.eventUpdateServe,
		k.EventDelete:       r.eventDeleteServe,
		k.PersonEventCreate: r.persnEventCreateServe,
		k.PersonEventUpdate: r.personEventUpdateServe,
		k.PersonEventDelete: r.personEventDeleteServe,
	}
	for k, v := range TopicsServe {
		wg.Add(1)
		go r.RunReader(ctx, wg, k, v)
	}

	//wg.Add(1)
	// go r.RunReader(ctx, wg, k.PersonCreate, r.personCreateServe)
	// wg.Add(1)
	// go r.RunReader(ctx, wg, k.PersonUpdate, r.PersonUpdateServe)
	// wg.Add(1)
	// r.RunReader(ctx, wg, k.PersonDelete, r.personDeleteServe)

}
