package personevent

import (
	"party-calc/internal/kafka/producer"
	"party-calc/internal/server/grpc/proto"
	"party-calc/internal/service"
)

type PersonEventHandler struct {
	proto.PersonsEventsServiceServer
	service  *service.PersonsEventsService
	producer *producer.KafkaProducer
}

func NewPersonEventHandler(s *service.PersonsEventsService, p *producer.KafkaProducer) *PersonEventHandler {
	return &PersonEventHandler{service: s, producer: p}
}
