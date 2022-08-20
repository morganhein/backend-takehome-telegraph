package telegraph

type Storage interface {
	EquipmentStorage
	EventStorage
	LocationStorage
	WaybillStorage
}
