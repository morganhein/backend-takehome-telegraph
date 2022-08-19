package telegraph

type Events struct {
	ID                    int64  `csv:"id"`
	EquipmentID           string `csv:"equipment_id"`
	SightingDate          string `csv:"sighting_date"`
	SightingEventCode     int64  `csv:"sighting_event_code"`
	ReportingRailroadScac string `csv:"reporting_railroad_scac"`
	PostingDate           string `csv:"posting_date"`
	FromMarkID            string `csv:"from_mark_id"`
	LoadEmptyStatus       string `csv:"load_empty_status"`
	SightingClaimCode     string `csv:"sighting_claim_code"`
	SightingEventCodeText string `csv:"sighting_event_code_text"`
	TrainID               string `csv:"train_id"`
	TrainAlphaCode        string `csv:"train_alpha_code"`
	LocationID            int64  `csv:"location_id"`
	WaybillID             int64  `csv:"waybill_id"`
}
