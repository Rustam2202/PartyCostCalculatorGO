package event

import (
	"party-calc/internal/kafka/producer"
	"party-calc/internal/server/grpc/proto"
	"party-calc/internal/service"
)

type EventHandler struct {
	proto.EventServiceServer
	service  *service.EventService
	producer *producer.KafkaProducer
}

func NewEventHandler(s *service.EventService, p *producer.KafkaProducer) *EventHandler {
	return &EventHandler{service: s, producer: p}
}
