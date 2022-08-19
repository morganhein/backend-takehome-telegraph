package telegraph

type Equipment struct {
	Id              int64  `csv:"id"`
	Customer        string `csv:"customer"`
	Fleet           string `csv:"fleet"`
	EquipmentId     string `csv:"equipment_id"`
	EquipmentStatus bool   `csv:"equipment_status"`
	DateAdded       string `csv:"date_added"`
	DateRemoved     string `csv:"date_removed"`
}
