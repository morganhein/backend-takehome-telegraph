package store

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"

	sq "github.com/Masterminds/squirrel"
	"github.com/morganhein/backend-takehome-telegraph/internal/telegraph"
)

var (
	EquipmentTable = "equipment"
)

func (p postgres) CreateEquipment(e telegraph.Equipment) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	stmnt, args, err := psql.Insert(EquipmentTable).
		Columns("id", "customer", "fleet", "equipment_id", "equipment_status", "date_added", "date_removed").
		Values(e.ID, e.Customer, e.Fleet, e.EquipmentID, e.EquipmentStatus, e.DateAdded.Time, e.DateRemoved.Time).
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

func (p postgres) ListEquipment() ([]telegraph.Equipment, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	stmnt, _, err := psql.Select("*").
		From(EquipmentTable).ToSql()
	if err != nil {
		return nil, err
	}

	var e []telegraph.Equipment
	err = pgxscan.Select(context.Background(), p.pool, &e, stmnt)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (p postgres) GetEquipmentByEquipmentID(equipmentID string) ([]telegraph.Equipment, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	stmnt, _, err := psql.Select("*").
		From(EquipmentTable).
		Where("equipment_id = ?", equipmentID).
		ToSql()
	if err != nil {
		return nil, err
	}

	var e []telegraph.Equipment
	err = pgxscan.Select(context.Background(), p.pool, &e, stmnt)
	if err != nil {
		return nil, err
	}
	return e, nil
}
