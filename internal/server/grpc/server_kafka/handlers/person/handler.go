package person

import (
	"party-calc/internal/kafka/producer"
	"party-calc/internal/server/grpc/proto"
	"party-calc/internal/service"
)

type PersonHandler struct {
	proto.PersonServiceServer
	service  *service.PersonService
	producer *producer.KafkaProducer
}

func NewPersonHandler(s *service.PersonService, p *producer.KafkaProducer) *PersonHandler {
	return &PersonHandler{service: s, producer: p}
}
