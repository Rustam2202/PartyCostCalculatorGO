package config

import (
	"flag"
	"party-calc/internal/language"
	"party-calc/internal/logger"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Config struct {
	Server struct {
		Host string
		Port int
	}
	DataBase struct {
		Host     string
		Port     int
		User     string
		Password string
		Dbname   string
	}
	Language  language.Language
	RoundRate float64
}

var Cfg Config

func LoadConfig() {
	confPath := flag.String("config", "", "path to config file")
	flag.Parse()
	if *confPath == "" {
		*confPath = "../../" 
		//*confPath="${workspaceFolder}/"
	}

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(*confPath)

	err := viper.ReadInConfig()
	if err != nil {
		logger.Logger.Fatal("Can't read configs: ", zap.Error(err))
	}

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		logger.Logger.Fatal("Can't unmarshal configs: ", zap.Error(err))
	}
}
