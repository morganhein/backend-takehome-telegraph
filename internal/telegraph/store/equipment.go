package store

import (
	"github.com/Masterminds/squirrel"
	"github.com/Masterminds/structable"
	"github.com/morganhein/backend-takehome-telegraph/internal/telegraph"
)

var (
	EquipmentTable = "equipment"
)

type equipment struct {
	db squirrel.DBProxyBeginner
}

func (eq equipment) CreateEquipment(e telegraph.Equipment) error {
	r := structable.New(eq.db, "postgres").Bind(EquipmentTable, e)
	return r.Insert()
}
