package store

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4"
	"github.com/morganhein/backend-takehome-telegraph/internal/telegraph"
)

var (
	EquipmentTable = "equipment"
)

type equipment struct {
	conn *pgx.Conn
}

func (eq equipment) CreateEquipment(e telegraph.Equipment) error {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	stmnt, args, err := psql.Insert(EquipmentTable).
		Columns("id", "customer", "fleet", "equipment_id", "equipment_status", "date_added", "date_removed").
		Values(e.ID, e.Customer, e.Fleet, e.EquipmentID, e.EquipmentStatus, e.DateAdded.Time, e.DateRemoved.Time).
		ToSql()
	if err != nil {
		return err
	}
	_, err = eq.conn.Exec(context.Background(), stmnt, args...)
	if err != nil {
		return err
	}
	return nil
}
