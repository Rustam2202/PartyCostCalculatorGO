package logger

import "go.uber.org/zap"

var Logger *zap.Logger

func IntializeLogger(cfg *zap.Config) error {
	var err error
	Logger, err = cfg.Build()
	if err != nil {
		return err
	}
	return nil
	//Logger, _ = zap.NewProduction()
}
