package telegraph

import "github.com/morganhein/backend-takehome-telegraph/pkg/time"

type Waybill struct {
	ID                   int           `csv:"id"`
	EquipmentID          string        `csv:"equipment_id"`
	WaybillDate          time.DateTime `csv:"waybill_date"`
	WaybillNumber        int           `csv:"waybill_number"`
	CreatedDate          time.DateTime `csv:"created_date"`
	BillingRoadMarkName  string        `csv:"billing_road_mark_name"`
	WaybillSourceCode    string        `csv:"waybill_source_code"`
	LoadEmptyStatus      string        `csv:"load_empty_status"`
	OriginMarkName       string        `csv:"origin_mark_name"`
	DestinationMarkName  string        `csv:"destination_mark_name"`
	SendingRoadMark      string        `csv:"sending_road_mark"`
	BillOfLadingNumber   string        `csv:"bill_of_lading_number"`
	BillOfLadingDate     time.DateTime `csv:"bill_of_lading_date"`
	EquipmentWeight      int           `csv:"equipment_weight"`
	TareWeight           int           `csv:"tare_weight"`
	AllowableWeight      int           `csv:"allowable_weight"`
	DunnageWeight        int           `csv:"dunnage_weight"`
	EquipmentWeightCode  string        `csv:"equipment_weight_code"`
	CommodityCode        int           `csv:"commodity_code"`
	CommodityDescription string        `csv:"commodity_description"`
	OriginID             string        `csv:"origin_id"`
	DestinationID        int           `csv:"destination_id"`
	Routes               string        `csv:"routes"`
	Parties              string        `csv:"parties"`
}

type WaybillStorage interface {
	CreateWaybill(e Waybill) error
	ListWaybills() ([]Waybill, error)
}

func (s serviceImpl) ListWaybills() ([]Waybill, error) {
	return s.db.ListWaybills()
}
