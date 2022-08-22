package telegraph

type Location struct {
	ID        int     `csv:"id"`
	City      string  `csv:"city"`
	CityLong  string  `csv:"city_long"`
	Station   string  `csv:"station"`
	Fsac      int     `csv:"fsac"`
	Scac      string  `csv:"scac"`
	Splc      int64   `csv:"splc"`
	State     string  `csv:"state"`
	TimeZone  string  `csv:"time_zone"`
	Longitude float64 `csv:"longitude"`
	Latitude  float64 `csv:"latitude"`
	Country   string  `csv:"country"`
}

type LocationStorage interface {
	CreateLocation(e Location) error
	ListLocation() ([]Location, error)
}

func (s serviceImpl) ListLocations() ([]Location, error) {
	return s.db.ListLocation()
}
