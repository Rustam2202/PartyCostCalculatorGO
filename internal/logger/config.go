package logger

import "go.uber.org/zap/zapcore"

type LoggerConfig struct {
	Encoding         string                `yaml:"encoding"`
	Level            string                `yaml:"level"`
	OutputPaths      []string              `yaml:"outputPaths"`
	ErrorOutputPaths []string              `yaml:"errorOutputPaths"`
	EncoderConfig    zapcore.EncoderConfig `yaml:"encoderConfig"`
}
