package config

import (
	"flag"
	"log"
	"party-calc/internal/database"
	"party-calc/internal/kafka"
	"party-calc/internal/logger"
	"party-calc/internal/server/grpc"
	"party-calc/internal/server/http"

	"github.com/spf13/viper"
)

type Config struct {
	ServerConfig     http.ServerConfig
	ServerGrpcConfig grpc.ServerGrpcConfig
	DatabaseConfig   database.DatabaseConfig
	LoggerConfig     logger.LoggerConfig
	KafkaConfig      kafka.KafkaConfig
}

func LoadConfig() *Config {
	var cfg Config
	path := flag.String("confpath", "./", "path to config file")
	flag.Parse()

	viper.Reset()
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(*path)

	err := viper.ReadInConfig()
	if err != nil {
		// logger.Logger.Fatal("Failed to read configs: ", zap.Error(err))
		log.Fatal(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		//	logger.Logger.Fatal("Failed to unmarshal configs: ", zap.Error(err))
		log.Fatal(err)
	}
	return &cfg
}
