package consumer

import (
	"context"
	"party-calc/internal/logger"
	"party-calc/internal/server/grpc/proto"

	pm "github.com/golang/protobuf/proto"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

func (r *KafkaConsumer) RunPersonCreateReader(ctx context.Context) {
	cfg := *r.cfg
	cfg.Topic = "person-create"
	reader := kafka.NewReader(cfg)
	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			break
		}
		req := proto.PersonCreateRequest{}
		err = pm.Unmarshal(msg.Value, &req)
		if err != nil {
			logger.Logger.Error("Failed to unmarshal message: ", zap.Error(err))
			continue
		}
		_, err = r.services.PersonService.NewPerson(ctx, req.Name)
		if err != nil {
			logger.Logger.Error("Failed to add Peson to db: ", zap.Error(err))
			continue
		}
		err = reader.CommitMessages(ctx, msg)
		if err != nil {
			logger.Logger.Error("Failed to commit message: ", zap.Error(err))
			continue
		}
	}
}

func (r *KafkaConsumer) RunPersonUpdateReader(ctx context.Context) {
	cfg := *r.cfg
	cfg.Topic = "person-update"
	reader := kafka.NewReader(cfg)
	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			break
		}
		req := proto.PersonUpdateRequest{}
		err = pm.Unmarshal(msg.Value, &req)
		if err != nil {
			logger.Logger.Error("Failed to unmarshal message: ", zap.Error(err))
			continue
		}
		err = r.services.PersonService.UpdatePerson(ctx, req.Id, req.Name)
		if err != nil {
			logger.Logger.Error("Failed to update Person in db: ", zap.Error(err))
			continue
		}
		err = reader.CommitMessages(ctx, msg)
		if err != nil {
			logger.Logger.Error("Failed to commit message: ", zap.Error(err))
			continue
		}
	}
}

func (r *KafkaConsumer) RunPersonDeleteReader(ctx context.Context) {
	cfg := *r.cfg
	cfg.Topic = "person-delete"
	reader := kafka.NewReader(cfg)
	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			break
		}
		req := proto.Id{}
		err = pm.Unmarshal(msg.Value, &req)
		if err != nil {
			logger.Logger.Error("Failed to unmarshal message: ", zap.Error(err))
			continue
		}
		err = r.services.PersonService.DeletePersonById(ctx, req.Id)
		if err != nil {
			logger.Logger.Error("Failed to delete Person from db: ", zap.Error(err))
			continue
		}
		err = reader.CommitMessages(ctx, msg)
		if err != nil {
			logger.Logger.Error("Failed to commit message: ", zap.Error(err))
			continue
		}
	}
}
