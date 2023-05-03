package postgres

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateConnectionPool(ctx context.Context, uri string) *pgxpool.Pool {
	dbpool, err := pgxpool.New(ctx, uri)
	if err != nil {
		log.Fatalf("unable to establish connection with database: %v", err)
	}

	return dbpool
}
