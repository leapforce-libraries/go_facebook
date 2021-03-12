package facebook

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type AccountID int64

func (a *AccountID) UnmarshalJSON(b []byte) error {
	prefix := "act_"

	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	unquoted = strings.Trim(unquoted, " ")

	if unquoted == "" {
		a = nil
		return errors.New("Error parsing to AccountID")
	}

	if !strings.HasPrefix(unquoted, prefix) {
		a = nil
		return fmt.Errorf("AccountID does not have prefix '%s'", prefix)
	}
	unquoted = strings.TrimLeft(unquoted, prefix)

	i, err := strconv.ParseInt(unquoted, 10, 64)
	if err != nil {
		return err
	}

	*a = AccountID(i)
	return nil
}

func (a *AccountID) Value() int64 {
	if a == nil {
		return 0
	}

	return int64(*a)
}
