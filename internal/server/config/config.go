package config

import (
	"party-calc/internal/logger"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type ServerConfig struct {
	Server struct {
		Host string
		Port int
	}
}

func (cfg *ServerConfig) LoadConfig(path string) {
	viper.Reset()
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		logger.Logger.Fatal("Can't read configs: ", zap.Error(err))
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		logger.Logger.Fatal("Can't unmarshal configs: ", zap.Error(err))
	}
}
