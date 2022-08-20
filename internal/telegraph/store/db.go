package store

import (
	"context"
	"github.com/jackc/pgx/v4"
)

type postgres struct {
	equipment
	events
	locations
	waybills
}

func CreatePostgresStore(connString string) (*postgres, *pgx.Conn, error) {
	c, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, nil, err
	}
	return &postgres{
		// this is a bit repetitive, and could be improved upon later
		equipment: equipment{
			conn: c,
		},
		events: events{
			conn: c,
		},
		locations: locations{
			conn: c,
		},
		waybills: waybills{
			conn: c,
		},
	}, c, nil
}
