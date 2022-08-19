package telegraph

type Equipment struct {
	ID              int64  `csv:"id" stbl:"id, PRIMARY_KEY, AUTO_INCREMENT"`
	Customer        string `csv:"customer" stbl:"customer"`
	Fleet           string `csv:"fleet" stbl:"fleet"`
	EquipmentID     string `csv:"equipment_id" stbl:"equipment_id"`
	EquipmentStatus bool   `csv:"equipment_status" stbl:"equipment_status"`
	DateAdded       string `csv:"date_added" stbl:"date_added"`
	DateRemoved     string `csv:"date_removed" stbl:"date_removed"`
}

type EquipmentStorage interface {
	CreateEquipment(e Equipment) error
}
