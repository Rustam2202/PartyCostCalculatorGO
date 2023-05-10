package util

import "github.com/spf13/viper"

type Config struct {
	Mode     string `mode`
	Address  string `address`
	Port     int    `port`
	Language string `language`
}

var vp *viper.Viper

func LoadConfig()(Config,error){
	vp=viper.New()
	// var config Config

	return Config{},nil
}