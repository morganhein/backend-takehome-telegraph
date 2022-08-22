package telegraph

type Service interface {
	// Equipment
	ListEquipment() ([]Equipment, error)
	// Events
	ListEvents(sightingDate string) ([]Event, error)
	// Locations
	ListLocations() ([]Location, error)
	GetRoute(fromID int, toID int) (Route, error)
	// Waybills
	ListWaybills() ([]Waybill, error)
	// Association can be equipment, events, or locations, or empty
	GetWaybill(ID int, association string) (*Waybill, error)
}

type Storage interface {
	EquipmentStorage
	EventStorage
	LocationStorage
	WaybillStorage
}

type serviceImpl struct {
	db Storage
}

func New(storage Storage) Service {
	return &serviceImpl{
		db: storage,
	}
}
