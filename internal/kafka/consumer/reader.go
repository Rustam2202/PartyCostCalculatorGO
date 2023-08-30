package consumer

import (
	"context"
	"fmt"
	"party-calc/internal/logger"
	"sync"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

func (r *KafkaConsumer) RunReader(ctx context.Context, wg *sync.WaitGroup,
	topic string, serve func(context.Context, kafka.Message) error) {

	defer wg.Done()
	for {
		func() {
			defer func() {
				if r := recover(); r != nil {
					logger.Logger.Error("Panic occurred: ",
						zap.Any(fmt.Sprintf("panic in %s reader", topic), r))
				}
			}()
			cfg := *r.cfg
			cfg.Topic = topic
			reader := kafka.NewReader(cfg)
			defer reader.Close()
			logger.Logger.Info(fmt.Sprintf("%s reader created", topic))
			go func() {
				<-ctx.Done()
				logger.Logger.Info(fmt.Sprintf("%s reader closing ...", topic))
				reader.Close()
			}()
			for {
				select {
				case <-ctx.Done():
					return
				default:
					msg, err := reader.ReadMessage(ctx)
					if err != nil {
						logger.Logger.Error("Failed to read message: ", zap.Error(err))
						continue
					}
					if err = serve(ctx, msg); err != nil {
						continue
					}
					if err = reader.CommitMessages(ctx, msg); err != nil {
						logger.Logger.Error("Failed to commit message: ", zap.Error(err))
						continue
					}
				}
			}
		}()
		if ctx.Err() != nil {
			break
		}
	}
}
