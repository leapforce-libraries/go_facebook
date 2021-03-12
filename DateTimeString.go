package facebook

import (
	"strconv"
	"strings"
	"time"
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

	_t, err := time.Parse(DateFormat, unquoted)
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
