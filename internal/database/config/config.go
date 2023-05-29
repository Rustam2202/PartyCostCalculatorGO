package config

import (
	"flag"
	"party-calc/internal/logger"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		Dbname   string
	}
}

func (cfg *Config) LoadConfig() {
	confPath := flag.String("dbconfig", "./internal/database/config", "path to config file")
	flag.Parse()

	viper.SetConfigType("yaml")
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
