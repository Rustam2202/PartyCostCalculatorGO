package consumer

import (
	"context"
	"party-calc/internal/logger"
	"party-calc/internal/server/grpc/proto"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	pm "google.golang.org/protobuf/proto"
)

func (r *KafkaConsumer) personCreateServe(ctx context.Context, msg kafka.Message) error {
	req := proto.PersonCreateRequest{}
	if err := pm.Unmarshal(msg.Value, &req); err != nil {
		logger.Logger.Error("Failed to unmarshal message: ", zap.Error(err))
		return err
	}
	if _, err := r.services.PersonService.NewPerson(ctx, req.Name); err != nil {
		logger.Logger.Error("Failed to add Peson to db: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *KafkaConsumer) PersonUpdateServe(ctx context.Context, msg kafka.Message) error {
	req := proto.PersonUpdateRequest{}
	if err := pm.Unmarshal(msg.Value, &req); err != nil {
		logger.Logger.Error("Failed to unmarshal message: ", zap.Error(err))
		return err
	}
	if err := r.services.PersonService.UpdatePerson(ctx, req.Id, req.Name); err != nil {
		logger.Logger.Error("Failed to update Person in db: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *KafkaConsumer) personDeleteServe(ctx context.Context, msg kafka.Message) error {
	req := proto.Id{}
	if err := pm.Unmarshal(msg.Value, &req); err != nil {
		logger.Logger.Error("Failed to unmarshal message: ", zap.Error(err))
		return err
	}
	if err := r.services.PersonService.DeletePersonById(ctx, req.Id); err != nil {
		logger.Logger.Error("Failed to delete Person from db: ", zap.Error(err))
		return err
	}
	return nil
}
