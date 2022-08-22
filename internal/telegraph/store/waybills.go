package store

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/morganhein/backend-takehome-telegraph/internal/telegraph"
)

var (
	WaybillsTable = "waybills"
)

func (p postgres) CreateWaybill(e telegraph.Waybill) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	stmnt, args, err := psql.Insert(WaybillsTable).
		Columns("id", "equipment_id", "waybill_date", "waybill_number", "created_date",
			"billing_road_mark_name", "waybill_source_code", "load_empty_status", "origin_mark_name",
			"destination_mark_name", "sending_road_mark", "bill_of_lading_number", "bill_of_lading_date",
			"equipment_weight", "tare_weight", "allowable_weight", "dunnage_weight", "equipment_weight_code",
			"commodity_code", "commodity_description", "origin_id", "destination_id", "routes", "parties").
		Values(e.ID, e.EquipmentID, e.WaybillDate, e.WaybillNumber, e.CreatedDate,
			e.BillingRoadMarkName, e.WaybillSourceCode, e.LoadEmptyStatus, e.OriginMarkName,
			e.DestinationMarkName, e.SendingRoadMark, e.BillOfLadingNumber, e.BillOfLadingDate,
			e.EquipmentWeight, e.TareWeight, e.AllowableWeight, e.DunnageWeight, e.EquipmentWeightCode,
			e.CommodityCode, e.CommodityDescription, e.OriginID, e.DestinationID, e.Routes, e.Parties).
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

func (p postgres) ListWaybills() ([]telegraph.Waybill, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	stmnt, _, err := psql.Select("*").
		From(WaybillsTable).ToSql()
	if err != nil {
		return nil, err
	}

	var e []telegraph.Waybill
	err = pgxscan.Select(context.Background(), p.pool, &e, stmnt)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (p postgres) GetWaybill(ID string, association string) (*telegraph.Waybill, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	stmnt, args, err := psql.Select("*").
		From(fmt.Sprintf("%v as w", WaybillsTable)).
		Where(sq.Eq{"w.id": ID}).ToSql()

	if err != nil {
		return nil, err
	}
	var e telegraph.Waybill
	err = pgxscan.Get(context.Background(), p.pool, &e, stmnt, args...)
	if err != nil {
		return nil, err
	}

	if association == "equipment" {
		eq, err := p.GetEquipmentByEquipmentID(e.EquipmentID)
		if err != nil {
			return nil, err
		}
		e.Equipment = eq
	}
	if association == "locations" {
		var locations []telegraph.Location
		loc, err := p.GetLocation(e.OriginID)
		if err != nil {
			return nil, err
		}
		locations = append(locations, *loc)
		loc, err = p.GetLocation(e.DestinationID)
		if err != nil {
			return nil, err
		}
		locations = append(locations, *loc)
	}
	if association == "events" {
		ev, err := p.GetEventsByWaybillID(e.ID)
		if err != nil {
			return nil, err
		}
		e.Event = ev
	}

	return &e, nil
}
