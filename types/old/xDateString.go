package facebook

import (
	"strconv"
	"strings"
	"time"
)

const (
	DateFormat     string = "2006-01-02"
)

type DateString time.Time

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

	*d = DateString(_t)
	return nil
}

func (d *DateString) ValuePtr() *time.Time {
	if d == nil {
		return nil
	}

	_d := time.Time(*d)
	return &_d
}
