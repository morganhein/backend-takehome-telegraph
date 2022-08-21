package store

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"

	sq "github.com/Masterminds/squirrel"
	"github.com/morganhein/backend-takehome-telegraph/internal/telegraph"
)

var (
	LocationsTable = "locations"
)

type locations struct {
	pool *pgxpool.Pool
}

func (ev locations) CreateLocation(e telegraph.Location) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	stmnt, args, err := psql.Insert(LocationsTable).
		Columns("id", "city", "city_long", "station", "fsac", "scac", "splc", "state", "time_zone", "longitude", "latitude", "country").
		Values(e.ID, e.City, e.CityLong, e.Station, e.Fsac, e.Scac, e.Splc, e.State, e.Timezone, e.Longitude, e.Latitude, e.Country).
		ToSql()
	if err != nil {
		return err
	}
	_, err = ev.pool.Exec(context.Background(), stmnt, args...)
	if err != nil {
		return err
	}
	return nil
}
