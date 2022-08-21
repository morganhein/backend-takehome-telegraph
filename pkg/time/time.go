package time

import (
	"errors"
	"fmt"
	"time"
)

type DateTime struct {
	*time.Time
}

var format = "2006-01-02 15:04:05"

func (date DateTime) MarshalCSV() (string, error) {
	return date.Time.Format(format), nil
}

func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	if csv == "" {
		date.Time = &time.Time{}
		return nil
	}
	t, err := time.Parse(format, csv)
	if err != nil {
		return err
	}
	date.Time = &t
	return nil
}

func (date DateTime) MarshalJson() (string, error) {
	return date.MarshalJson()
}

func (date *DateTime) UnmarshalJson(input []byte) (err error) {
	if len(input) == 0 {
		date.Time = &time.Time{}
		return nil
	}
	err = date.Time.UnmarshalJSON(input)
	return err
}

func (date *DateTime) Scan(src any) error {
	fmt.Println(src)
	t, ok := src.(time.Time)
	if !ok {
		return errors.New("did not recieve a date from db")
	}
	date.Time = &t
	return nil
}
