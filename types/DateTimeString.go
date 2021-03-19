package facebook

import (
	"strconv"
	"strings"
	"time"
)

const (
	DateTimeFormat string = "2006-01-02T15:04:05-0700"
)

type DateTimeString time.Time

func (d *DateTimeString) UnmarshalJSON(b []byte) error {
	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	if strings.Trim(unquoted, " ") == "" {
		d = nil
		return nil
	}

	_t, err := time.Parse(DateTimeFormat, unquoted)
	if err != nil {
		return err
	}

	*d = DateTimeString(_t)
	return nil
}

func (d *DateTimeString) ValuePtr() *time.Time {
	if d == nil {
		return nil
	}

	_d := time.Time(*d)
	return &_d
}
