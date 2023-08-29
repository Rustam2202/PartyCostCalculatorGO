package consumer

import (
	"context"
	"party-calc/internal/logger"
	"party-calc/internal/server/grpc/proto"
	"time"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	pm "google.golang.org/protobuf/proto"
)

func (r *KafkaConsumer) eventCreateServe(ctx context.Context, msg kafka.Message) error {
	req := proto.EventCreateRequest{}
	err := pm.Unmarshal(msg.Value, &req)
	if err != nil {
		logger.Logger.Error("Failed to unmarshal message: ", zap.Error(err))
		return err
	}
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		logger.Logger.Error("Failed to parse date: ", zap.Error(err))
		return err
	}
	_, err = r.services.EventService.NewEvent(ctx, req.Name, date)
	if err != nil {
		logger.Logger.Error("Failed to add Event to db: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *KafkaConsumer) eventUpdateServe(ctx context.Context, msg kafka.Message) error {
	req := proto.EventUpdateRequest{}
	err := pm.Unmarshal(msg.Value, &req)
	if err != nil {
		logger.Logger.Error("Failed to unmarshal message: ", zap.Error(err))
		return err
	}
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		logger.Logger.Error("Failed to parse date: ", zap.Error(err))
		return err
	}
	err = r.services.EventService.UpdateEvent(ctx, req.Id, req.Name, date)
	if err != nil {
		logger.Logger.Error("Failed to update Event in db: ", zap.Error(err))
		return err
	}
	return nil
}

func (r *KafkaConsumer) eventDeleteServe(ctx context.Context, msg kafka.Message) error {
	req := proto.Id{}
	err := pm.Unmarshal(msg.Value, &req)
	if err != nil {
		logger.Logger.Error("Failed to unmarshal message: ", zap.Error(err))
		return err
	}
	err = r.services.EventService.DeleteEventById(ctx, req.Id)
	if err != nil {
		logger.Logger.Error("Failed to delete Event from db: ", zap.Error(err))
		return err
	}
	return nil
}
