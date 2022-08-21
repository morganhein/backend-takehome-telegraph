package telegraph

type Service interface {
	// Equipment
	ListEquipment() ([]Equipment, error)
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
