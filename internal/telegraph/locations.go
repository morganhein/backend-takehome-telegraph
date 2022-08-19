package telegraph

type Location struct {
	ID              int64  `csv:"id"`
	Customer        string `csv:"customer"`
	Fleet           string `csv:"fleet"`
	EquipmentID     string `csv:"equipment_id"`
	EquipmentStatus bool   `csv:"equipment_status"`
	DateAdded       string `csv:"date_added"`
	DateRemoved     string `csv:"date_removed"`
}

type LocationStorage interface {
	CreateLocation(e Location) error
}
