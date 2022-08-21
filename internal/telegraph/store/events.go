package store

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/morganhein/backend-takehome-telegraph/internal/telegraph"
)

var (
	EventsTable = "events"
)

type events struct {
	pool *pgxpool.Pool
}

func (ev events) CreateEvent(e telegraph.Event) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	stmnt, args, err := psql.Insert(EventsTable).
		Columns("id", "equipment_id", "sighting_date", "sighting_event_code", "reporting_railroad_scac",
			"posting_date", "from_mark_id", "load_empty_status", "sighting_claim_code", "sighting_event_code_text",
			"train_id", "train_alpha_code", "location_id", "waybill_id").
		Values(e.ID, e.EquipmentID, e.SightingDate, e.SightingEventCode, e.ReportingRailroadScac,
			e.PostingDate, e.FromMarkID, e.LoadEmptyStatus, e.SightingClaimCode, e.SightingEventCodeText,
			e.TrainID, e.TrainAlphaCode, e.LocationID, e.WaybillID).
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
