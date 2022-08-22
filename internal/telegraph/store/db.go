package store

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type postgres struct {
	pool *pgxpool.Pool
}

func CreatePostgresStore(connString string) (*postgres, error) {
	c, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	return &postgres{
		pool: c,
	}, nil
}
