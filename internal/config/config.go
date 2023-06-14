package config

import (
	"flag"
	"party-calc/internal/database"
	"party-calc/internal/logger"
	"party-calc/internal/server"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	ServerConfig   server.ServerConfig
	DatabaseConfig database.DatabaseConfig
	// RoundRate float64
}

func LoadConfig() Config {
	var cfg Config
	path := flag.String("confpath", "./", "path to config file")
	flag.Parse()

	viper.Reset()
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(*path)

	err := viper.ReadInConfig()
	if err != nil {
		logger.Logger.Fatal("Can't read configs: ", zap.Error(err))
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		logger.Logger.Fatal("Can't unmarshal configs: ", zap.Error(err))
	}

	return cfg
}
