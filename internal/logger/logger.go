package logger

import (
	"log"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func IntializeLogger(cfg LoggerConfig) {
	var err error
	Logger, err = zap.NewProductionConfig().Build()
	if err != nil {
		log.Fatal(err)
	}
}
