package facebook

import (
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/civil"
)

const (
	DateFormat string = "2006-01-02"
)

type DateString civil.Date

func (d *DateString) UnmarshalJSON(b []byte) error {
	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	if strings.Trim(unquoted, " ") == "" {
		d = nil
		return nil
	}

	_t, err := time.Parse(DateFormat, unquoted)
	if err != nil {
		return err
	}

	*d = DateString(civil.DateOf(_t))
	return nil
}

func (d *DateString) ValuePtr() *civil.Date {
	if d == nil {
		return nil
	}

	_d := civil.Date(*d)
	return &_d
}

func (d DateString) Value() civil.Date {
	return civil.Date(d)
}
