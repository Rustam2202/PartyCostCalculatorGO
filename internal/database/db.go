package database

import (
	"context"
	//	"database/sql"
	"fmt"

	"party-calc/internal/logger"

	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type DataBase struct {
	//DB    *sql.DB
	DBPGX *pgx.Conn
}

// func New(cfg DatabaseConfig) *DataBase {
// 	psqlconn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
// 		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname)
// 	db, err := sql.Open("postgres", psqlconn)
// 	if err != nil {
// 		logger.Logger.Error("Can't open database: ", zap.Error(err))
// 		return nil
// 	}
// 	return &DataBase{DB: db}
// }

func NewPGX(cfg DatabaseConfig) *DataBase {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname)
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		logger.Logger.Error("Can't open database: ", zap.Error(err))
		return nil
	}
	//defer conn.Close(context.Background())

	return &DataBase{DBPGX: conn}
}
