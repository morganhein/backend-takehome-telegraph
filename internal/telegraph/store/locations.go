package store

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/morganhein/backend-takehome-telegraph/internal/telegraph"
)

var (
	LocationsTable = "locations"
)

func (p postgres) CreateLocation(e telegraph.Location) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	stmnt, args, err := psql.Insert(LocationsTable).
		Columns("id", "city", "city_long", "station", "fsac", "scac", "splc", "state", "time_zone", "longitude", "latitude", "country").
		Values(e.ID, e.City, e.CityLong, e.Station, e.Fsac, e.Scac, e.Splc, e.State, e.TimeZone, e.Longitude, e.Latitude, e.Country).
		ToSql()
	if err != nil {
		return err
	}
	_, err = p.pool.Exec(context.Background(), stmnt, args...)
	if err != nil {
		return err
	}
	return nil
}

func (p postgres) ListLocation() ([]telegraph.Location, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	stmnt, _, err := psql.Select("*").
		From(LocationsTable).ToSql()
	if err != nil {
		return nil, err
	}

	var e []telegraph.Location
	err = pgxscan.Select(context.Background(), p.pool, &e, stmnt)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (p postgres) GetLocation(ID string) (*telegraph.Location, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	stmnt, args, err := psql.Select("*").
		From(LocationsTable).
		Where("id = ?", ID).
		ToSql()
	if err != nil {
		return nil, err
	}
	var e telegraph.Location
	err = pgxscan.Get(context.Background(), p.pool, &e, stmnt, args...)
	if err != nil {
		return nil, err
	}
	return &e, nil
}
