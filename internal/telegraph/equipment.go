package telegraph

import (
	"github.com/morganhein/backend-takehome-telegraph/pkg/time"
)

type Equipment struct {
	ID              int           `csv:"id"`
	Customer        string        `csv:"customer"`
	Fleet           string        `csv:"fleet"`
	EquipmentID     string        `csv:"equipment_id"`
	EquipmentStatus string        `csv:"equipment_status"`
	DateAdded       time.DateTime `csv:"date_added"`
	DateRemoved     time.DateTime `csv:"date_removed"`
}

type EquipmentStorage interface {
	CreateEquipment(e Equipment) error
}
