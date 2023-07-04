package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func IntializeLogger(cfg LoggerConfig) error {
	var err error
	Logger, err = zap.NewProductionConfig().Build()
	if err != nil {
		return err
	}
	return nil
}
