package time

import "time"

type DateTime struct {
	time.Time
}

var format = "2006-01-02 15:04:05"

func (date *DateTime) MarshalCSV() (string, error) {
	return date.Time.Format(format), nil
}

func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	if csv == "" {
		date.Time = time.Time{}
		return nil
	}
	date.Time, err = time.Parse(format, csv)
	return err
}

func (date *DateTime) MarshalJson() (string, error) {
	return date.MarshalJson()
}

func (date *DateTime) UnmarshalJson(input []byte) (err error) {
	if len(input) == 0 {
		date.Time = time.Time{}
		return nil
	}
	err = date.Time.UnmarshalJSON(input)
	return err
}
