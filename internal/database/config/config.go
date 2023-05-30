package config

import (
	"flag"
	"party-calc/internal/logger"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type DatabaseConfig struct {
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Dbname   string `mapstructure:"dbname"`
	}
}

func (cfg *DatabaseConfig) LoadConfig() {
	confPath := flag.String("dbconfig", "./internal/database/config/", "path to config file")
	flag.Parse()

	viper.Reset()
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
