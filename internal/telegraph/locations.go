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

type Route []Location

type LocationStorage interface {
	CreateLocation(e Location) error
	ListLocation() ([]Location, error)
	GetLocation(ID string) (*Location, error)
}

func (s serviceImpl) ListLocations() ([]Location, error) {
	return s.db.ListLocation()
}

func (s serviceImpl) GetRoute(fromID, toID string) (Route, error) {
	from, err := s.db.GetLocation(fromID)
	if err != nil {
		return nil, err
	}
	to, err := s.db.GetLocation(toID)
	if err != nil {
		return nil, err
	}
	return Route{*from, *to}, nil
}
