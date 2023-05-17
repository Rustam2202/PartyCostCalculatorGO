package config

import (
	"flag"
	"party-calc/internal/language"
	"party-calc/internal/logger"

	"github.com/spf13/viper"
)

type Config struct {
	Host     string
	Port     int
	Language language.Language
}

var Cfg Config

func LoadConfig() {
	confPath := flag.String("config", "", "path to config file")
	flag.Parse()

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(*confPath)
	
	err := viper.ReadInConfig()
	if err != nil {
		logger.Logger.Fatal("Can't read configurations")
	}

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		logger.Logger.Fatal("Can't read configurations")
	}
}
