package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func ProvideDB(ctx context.Context) (*pgx.Conn, error) {
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	if conn, err := pgx.Connect(ctx, dbURL); err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
		return nil, err
	} else {
		log.Printf("Connect database [%s] success\n", conn.Config().Database)
		return conn, nil
	}
}
