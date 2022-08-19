package telegraph

type Locations struct {
	Id              int64  `csv:"id"`
	Customer        string `csv:"customer"`
	Fleet           string `csv:"fleet"`
	Equipmentid     string `csv:"equipment_id"`
	Equipmentstatus bool   `csv:"equipment_status"`
	Dateadded       string `csv:"date_added"`
	Dateremoved     string `csv:"date_removed"`
}
