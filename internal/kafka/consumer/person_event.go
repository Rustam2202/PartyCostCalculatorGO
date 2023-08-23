package consumer

import (
	"context"
	k "party-calc/internal/kafka"
	"party-calc/internal/logger"
	"party-calc/internal/server/grpc/proto"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	pm "google.golang.org/protobuf/proto"
)

func (r *KafkaConsumer) RunPersonEventCreateReader(ctx context.Context) {
	for {
		func() {
			defer func() {
				if r := recover(); r != nil {
					logger.Logger.Error("Panic occurred: ", zap.Any("panic in event delete consumer", r))
				}
			}()
			cfg := *r.cfg
			cfg.Topic = k.PersonEventCreate
			reader := kafka.NewReader(cfg)
			for {
				msg, err := reader.ReadMessage(ctx)
				if err != nil {
					logger.Logger.Error("Failed to read message: ", zap.Error(err))
					continue
				}
				req := proto.PersonEventCreateRequest{}
				err = pm.Unmarshal(msg.Value, &req)
				if err != nil {
					logger.Logger.Error("Failed to unmarshal message: ", zap.Error(err))
					continue
				}
				_, err = r.services.PersonEventService.AddPersonToPersEvent(
					ctx, req.PersonId, req.EventId, req.Spent, int(req.Factor))
				if err != nil {
					logger.Logger.Error("Failed to add PersonEvent to db: ", zap.Error(err))
					continue
				}
				err = reader.CommitMessages(ctx, msg)
				if err != nil {
					logger.Logger.Error("Failed to commit message: ", zap.Error(err))
					continue
				}
			}
		}()
	}
}

func (r *KafkaConsumer) RunPersonEventUpdateReader(ctx context.Context) {
	for {
		func() {
			defer func() {
				if r := recover(); r != nil {
					logger.Logger.Error("Panic occurred: ", zap.Any("panic in event delete consumer", r))
				}
			}()
			cfg := *r.cfg
			cfg.Topic = k.PersonEventUpdate
			reader := kafka.NewReader(cfg)
			for {
				msg, err := reader.ReadMessage(ctx)
				if err != nil {
					logger.Logger.Error("Failed to read message: ", zap.Error(err))
					continue
				}
				req := proto.PersonEventUpdateRequest{}
				err = pm.Unmarshal(msg.Value, &req)
				if err != nil {
					logger.Logger.Error("Failed to unmarshal message: ", zap.Error(err))
					continue
				}
				err = r.services.PersonEventService.Update(
					ctx, req.Id, req.PersonId, req.EventId, req.Spent, int(req.Factor))
				if err != nil {
					logger.Logger.Error("Failed to update PersonEvent in db: ", zap.Error(err))
					continue
				}
				err = reader.CommitMessages(ctx, msg)
				if err != nil {
					logger.Logger.Error("Failed to commit message: ", zap.Error(err))
					continue
				}
			}
		}()
	}
}

func (r *KafkaConsumer) RunPersonEventDeleteReader(ctx context.Context) {
	for {
		func() {
			defer func() {
				if r := recover(); r != nil {
					logger.Logger.Error("Panic occurred: ", zap.Any("panic in event delete consumer", r))
				}
			}()
			cfg := *r.cfg
			cfg.Topic = k.PersonEventDelete
			reader := kafka.NewReader(cfg)
			for {
				msg, err := reader.ReadMessage(ctx)
				if err != nil {
					logger.Logger.Error("Failed to read message: ", zap.Error(err))
					continue
				}
				req := proto.Id{}
				err = pm.Unmarshal(msg.Value, &req)
				if err != nil {
					logger.Logger.Error("Failed to unmarshal message: ", zap.Error(err))
					continue
				}
				err = r.services.PersonEventService.Delete(ctx, req.Id)
				if err != nil {
					logger.Logger.Error("Failed to delete PersonEvent from db: ", zap.Error(err))
					continue
				}
				err = reader.CommitMessages(ctx, msg)
				if err != nil {
					logger.Logger.Error("Failed to commit message: ", zap.Error(err))
					continue
				}
			}
		}()
	}
}
