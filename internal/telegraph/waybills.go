package telegraph

type Waybill struct {
	Id                   bool   `csv:"id"`
	EquipmentID          string `csv:"equipment_id"`
	WaybillDate          string `csv:"waybill_date"`
	WaybillNumber        int64  `csv:"waybill_number"`
	CreatedDate          string `csv:"created_date"`
	BillingRoadMarkName  string `csv:"billing_road_mark_name"`
	WaybillSourcecode    int64  `csv:"waybill_source_code"`
	LoadEmptyStatus      string `csv:"load_empty_status"`
	OriginMarkName       string `csv:"origin_mark_name"`
	DestinationMarkName  string `csv:"destination_mark_name"`
	SendingRoadMark      string `csv:"sending_road_mark"`
	BillOfLadingNumber   int64  `csv:"bill_of_lading_number"`
	BillOfLadingDate     string `csv:"bill_of_lading_date"`
	EquipmentWeight      int64  `csv:"equipment_weight"`
	TareWeight           int64  `csv:"tare_weight"`
	AllowableWeight      bool   `csv:"allowable_weight"`
	DunnageWeight        bool   `csv:"dunnage_weight"`
	EquipmentWeightCode  string `csv:"equipment_weight_code"`
	CommodityCode        int64  `csv:"commodity_code"`
	CommodityDescription string `csv:"commodity_description"`
	OriginID             string `csv:"origin_id"`
	DestinationID        bool   `csv:"destination_id"`
	Routes               string `csv:"routes"`
	Parties              string `csv:"parties"`
}

type WaybillStorage interface {
	CreateWaybill(e Waybill) error
}
