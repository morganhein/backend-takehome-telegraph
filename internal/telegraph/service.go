package telegraph

type Service interface {
	// Equipment
	ListEquipment() ([]Equipment, error)
	//ListEvents
	ListEvents() ([]Event, error)
	//ListLocations
	ListLocations() ([]Location, error)
	//ListWaybills
	ListWaybills() ([]Waybill, error)
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
