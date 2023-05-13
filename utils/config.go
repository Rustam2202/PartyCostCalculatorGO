package utils

import (
	"fmt"
	"party-calc/internal/language"

	"github.com/spf13/viper"
)

type Config struct {
	Host     string
	Port     int
	Language language.Language
}

var Cfg Config

func LoadConfig() {
	//	viper.AddConfigPath(".")
	//	viper.SetConfigFile("config")
	//	viper.SetConfigType("yml")
	viper.SetConfigFile("./config.yaml")

	err := viper.ReadInConfig()
	if err != nil {
		Logger.Error("Can't read configurations")
		panic(err)
	}


	err = viper.Unmarshal(&Cfg)
	if err != nil {
		Logger.Error("Can't read configurations")
		panic(err)
	}

	fmt.Println(Cfg.Language)

}
