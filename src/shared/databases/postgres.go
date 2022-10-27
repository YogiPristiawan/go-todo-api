package databases

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB_CONN_URL = os.Getenv("DB_CONN_URL")

func NewPgConn() (pool *pgxpool.Pool) {
	pool, err := pgxpool.New(context.Background(), DB_CONN_URL)
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}
	return pool
}
