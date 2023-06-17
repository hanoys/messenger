package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateConnectionPool(ctx context.Context, uri string) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(ctx, uri)
	if err != nil {
		return nil, err
	}

	if err = dbpool.Ping(ctx); err != nil {
		dbpool.Close()
		return nil, err
	}

	return dbpool, nil
}
