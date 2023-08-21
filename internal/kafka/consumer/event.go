package consumer

import (
	"context"
	k "party-calc/internal/kafka"
	"party-calc/internal/logger"
	"party-calc/internal/server/grpc/proto"
	"time"

	pm "google.golang.org/protobuf/proto"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

func (r *KafkaConsumer) RunEventCreateReader(ctx context.Context) {
	cfg := *r.cfg
	cfg.Topic = k.EventCreate
	reader := kafka.NewReader(cfg)
	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			break
		}
		req := proto.EventCreateRequest{}
		err = pm.Unmarshal(msg.Value, &req)
		if err != nil {
			logger.Logger.Error("Failed to unmarshal message: ", zap.Error(err))
			continue
		}
		date, err := time.Parse("2006-01-02", req.Date)
		if err != nil {
			logger.Logger.Error("Failed to parse date: ", zap.Error(err))
			continue
		}
		_, err = r.services.EventService.NewEvent(ctx, req.Name, date)
		if err != nil {
			logger.Logger.Error("Failed to add Event to db: ", zap.Error(err))
			continue
		}
		err = reader.CommitMessages(ctx, msg)
		if err != nil {
			logger.Logger.Error("Failed to commit message: ", zap.Error(err))
			continue
		}
	}
}

func (r *KafkaConsumer) RunEventUpdateReader(ctx context.Context) {
	cfg := *r.cfg
	cfg.Topic = k.EventUpdate
	reader := kafka.NewReader(cfg)
	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			break
		}
		req := proto.EventUpdate{}
		err = pm.Unmarshal(msg.Value, &req)
		if err != nil {
			logger.Logger.Error("Failed to unmarshal message: ", zap.Error(err))
			continue
		}
		date, err := time.Parse("2006-01-02", req.Date)
		if err != nil {
			logger.Logger.Error("Failed to parse date: ", zap.Error(err))
			continue
		}
		err = r.services.EventService.UpdateEvent(ctx, req.Id, req.Name, date)
		if err != nil {
			logger.Logger.Error("Failed to update Event in db: ", zap.Error(err))
			continue
		}
		err = reader.CommitMessages(ctx, msg)
		if err != nil {
			logger.Logger.Error("Failed to commit message: ", zap.Error(err))
			continue
		}
	}
}

func (r *KafkaConsumer) RunEventDeleteReader(ctx context.Context) {
	cfg := *r.cfg
	cfg.Topic = k.EventDelete
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
		err = r.services.EventService.DeleteEventById(ctx, req.Id)
		if err != nil {
			logger.Logger.Error("Failed to delete Event from db: ", zap.Error(err))
			continue
		}
		err = reader.CommitMessages(ctx, msg)
		if err != nil {
			logger.Logger.Error("Failed to commit message: ", zap.Error(err))
			continue
		}
	}
}
