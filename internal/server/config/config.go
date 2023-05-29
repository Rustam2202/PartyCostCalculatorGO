package config

import (
	"flag"
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

func (cfg *ServerConfig) LoadConfig() {
	confPath := flag.String("serverconfig", "./internal/server/config", "path to config file")
	flag.Parse()

	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.AddConfigPath(*confPath)

	err := viper.ReadInConfig()
	if err != nil {
		logger.Logger.Fatal("Can't read configs: ", zap.Error(err))
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		logger.Logger.Fatal("Can't unmarshal configs: ", zap.Error(err))
	}
}
