package utils

import "go.uber.org/zap"

var Logger *zap.Logger

func IntializeLogger() {
	Logger, _ = zap.NewProduction()
}
