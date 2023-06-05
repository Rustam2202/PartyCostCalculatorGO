package config

import (
	"flag"
	srvCfg "party-calc/internal/database/config"
	dbCfg "party-calc/internal/server/config"
)

type Config struct {
	DatabaseCobfig srvCfg.DatabaseConfig
	ServerConfig   dbCfg.ServerConfig
	// RoundRate float64
}

func LoadConfig() Config {
	var cfg Config

	srvCfgPath := flag.String("srvcfg", "../server/config/", "path to server config file")
	dbCfgPath := flag.String("dbcfg", "../database/config/", "path to database config file")
	flag.Parse()

	cfg.DatabaseCobfig.LoadConfig(*dbCfgPath)
	cfg.ServerConfig.LoadConfig(*srvCfgPath)

	return cfg
}
