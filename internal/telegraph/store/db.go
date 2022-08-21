package store

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type postgres struct {
	equipment
	events
	locations
	waybills
}

func CreatePostgresStore(connString string) (*postgres, error) {
	c, err := pgxpool.Connect(context.Background(), connString)
	if err != nil {
		return nil, err
	}
	return &postgres{
		// this is a bit repetitive, and could be improved upon later
		equipment: equipment{
			pool: c,
		},
		events: events{
			pool: c,
		},
		locations: locations{
			pool: c,
		},
		waybills: waybills{
			pool: c,
		},
	}, nil
}
