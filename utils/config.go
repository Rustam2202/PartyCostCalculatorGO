package utils

import "github.com/spf13/viper"

type Config struct {
	Mode     string `json:"mode"`
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Language string `json:"language"`
}

var v *viper.Viper

func LoadConfig() (Config, error) {
	v = viper.New()
	v.AddConfigPath("./")
	v.SetConfigFile("config")
	v.SetConfigType("yaml")
	// var config Config

	return Config{}, nil
}
