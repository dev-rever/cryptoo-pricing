package db

import (
	"context"
	"log"

	"github.com/dev-rever/cryptoo-pricing/config"
	"github.com/jackc/pgx/v5"
)

func ProvideDB(ctx context.Context) (*pgx.Conn, error) {
	if conn, err := pgx.Connect(ctx, config.GetDBUrl()); err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
		return nil, err
	} else {
		log.Printf("Connect database [%s] success\n", conn.Config().Database)
		return conn, nil
	}
}
