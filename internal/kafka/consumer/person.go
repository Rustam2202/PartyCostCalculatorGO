package consumer

import (
	"context"
	k "party-calc/internal/kafka"
	"party-calc/internal/logger"
	"party-calc/internal/server/grpc/proto"
	"sync"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	pm "google.golang.org/protobuf/proto"
)

func (r *KafkaConsumer) RunPersonCreateReader(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		func() {
			defer func() {
				if r := recover(); r != nil {
					logger.Logger.Error("Panic occurred: ", zap.Any("panic in person craete consumer", r))
				}
			}()

			cfg := *r.cfg
			cfg.Topic = k.PersonCreate
			reader := kafka.NewReader(cfg)
			logger.Logger.Info("person-create reader created")
			go func() {
				<-ctx.Done()
				logger.Logger.Info("person-create reader closing ...")
				reader.Close()
			}()
			for {
				msg, err := reader.ReadMessage(ctx)
				if err != nil {
					logger.Logger.Error("Failed to read message: ", zap.Error(err))
				}
				req := proto.PersonCreateRequest{}
				if err = pm.Unmarshal(msg.Value, &req); err != nil {
					logger.Logger.Error("Failed to unmarshal message: ", zap.Error(err))
				}
				if _, err = r.services.PersonService.NewPerson(ctx, req.Name); err != nil {
					logger.Logger.Error("Failed to add Peson to db: ", zap.Error(err))
				}
				if err = reader.CommitMessages(ctx, msg); err != nil {
					logger.Logger.Error("Failed to commit message: ", zap.Error(err))
				}
			}
		}()
	}
}

func (r *KafkaConsumer) StopPersonCreateReader(ctx context.Context) {

}

func (r *KafkaConsumer) RunPersonUpdateReader(ctx context.Context) {
	for {
		func() {
			defer func() {
				if r := recover(); r != nil {
					logger.Logger.Error("Panic occurred: ", zap.Any("panic in person update consumer", r))
				}
			}()
			cfg := *r.cfg
			cfg.Topic = k.PersonUpdate
			reader := kafka.NewReader(cfg)
			for {
				msg, err := reader.ReadMessage(ctx)
				if err != nil {
					logger.Logger.Error("Failed to read message: ", zap.Error(err))
					continue
				}
				req := proto.PersonUpdateRequest{}
				if err = pm.Unmarshal(msg.Value, &req); err != nil {
					logger.Logger.Error("Failed to unmarshal message: ", zap.Error(err))
					continue
				}
				if err = r.services.PersonService.UpdatePerson(ctx, req.Id, req.Name); err != nil {
					logger.Logger.Error("Failed to update Person in db: ", zap.Error(err))
					continue
				}
				if err = reader.CommitMessages(ctx, msg); err != nil {
					logger.Logger.Error("Failed to commit message: ", zap.Error(err))
					continue
				}
			}
		}()
	}
}

func (r *KafkaConsumer) RunPersonDeleteReader(ctx context.Context) {
	for {
		func() {
			defer func() {
				if r := recover(); r != nil {
					logger.Logger.Error("Panic occurred: ", zap.Any("panic in person delete consumer", r))
				}
			}()
			cfg := *r.cfg
			cfg.Topic = k.PersonDelete
			reader := kafka.NewReader(cfg)
			for {
				msg, err := reader.ReadMessage(ctx)
				if err != nil {
					logger.Logger.Error("Failed to read message: ", zap.Error(err))
					continue
				}
				req := proto.Id{}
				if err = pm.Unmarshal(msg.Value, &req); err != nil {
					logger.Logger.Error("Failed to unmarshal message: ", zap.Error(err))
					continue
				}
				if err = r.services.PersonService.DeletePersonById(ctx, req.Id); err != nil {
					logger.Logger.Error("Failed to delete Person from db: ", zap.Error(err))
					continue
				}
				if err = reader.CommitMessages(ctx, msg); err != nil {
					logger.Logger.Error("Failed to commit message: ", zap.Error(err))
					continue
				}
			}
		}()
	}
}
