package database

import (
	"database/sql"
	"fmt"

	"party-calc/internal/logger"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type DataBase struct {
	DB *sql.DB
}

func New(cfg DatabaseConfig) *DataBase {
	psqlconn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname)
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		logger.Logger.Error("Can't open database: ", zap.Error(err))
		return nil
	}
	return &DataBase{DB: db}
}
