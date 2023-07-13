package database

import (
	"context"
	"party-calc/internal/logger"

	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type DBInterface interface {
	Begin(context.Context) (pgx.Tx, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	Ping(context.Context) error
	Prepare(context.Context, string, string) (*pgconn.StatementDescription, error)
	Close(context.Context) error
}

type DataBase struct {
	DBPGX DBInterface
}

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
