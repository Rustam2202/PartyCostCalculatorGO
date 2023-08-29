package consumer

import (
	"context"
	"party-calc/internal/logger"
	"party-calc/internal/server/grpc/proto"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	pm "google.golang.org/protobuf/proto"
)

func (r *KafkaConsumer) persnEventCreateServe(ctx context.Context, msg kafka.Message) error {
	req := proto.PersonEventCreateRequest{}
	err := pm.Unmarshal(msg.Value, &req)
	if err != nil {
		logger.Logger.Error("Failed to unmarshal message: ", zap.Error(err))
		return err
	}
	_, err = r.services.PersonEventService.AddPersonToPersEvent(
		ctx, req.PersonId, req.EventId, req.Spent, int(req.Factor))
	if err != nil {
		logger.Logger.Error("Failed to add PersonEvent to db: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *KafkaConsumer) personEventUpdateServe(ctx context.Context, msg kafka.Message) error {
	req := proto.PersonEventUpdateRequest{}
	err := pm.Unmarshal(msg.Value, &req)
	if err != nil {
		logger.Logger.Error("Failed to unmarshal message: ", zap.Error(err))
		return err
	}
	err = r.services.PersonEventService.Update(
		ctx, req.Id, req.PersonId, req.EventId, req.Spent, int(req.Factor))
	if err != nil {
		logger.Logger.Error("Failed to update PersonEvent in db: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *KafkaConsumer) personEventDeleteServe(ctx context.Context, msg kafka.Message) error {
	req := proto.Id{}
	err := pm.Unmarshal(msg.Value, &req)
	if err != nil {
		logger.Logger.Error("Failed to unmarshal message: ", zap.Error(err))
		return err
	}
	err = r.services.PersonEventService.Delete(ctx, req.Id)
	if err != nil {
		logger.Logger.Error("Failed to delete PersonEvent from db: ", zap.Error(err))
		return err
	}
	return nil
}
