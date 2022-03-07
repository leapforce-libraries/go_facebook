package facebook

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type AccountId int64

func (a *AccountId) UnmarshalJSON(b []byte) error {
	prefix := "act_"

	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return err
	}

	unquoted = strings.Trim(unquoted, " ")

	if unquoted == "" {
		a = nil
		return errors.New("error parsing to AccountId")
	}

	if !strings.HasPrefix(unquoted, prefix) {
		a = nil
		return fmt.Errorf("AccountId does not have prefix '%s'", prefix)
	}
	unquoted = strings.TrimLeft(unquoted, prefix)

	i, err := strconv.ParseInt(unquoted, 10, 64)
	if err != nil {
		return err
	}

	*a = AccountId(i)
	return nil
}

func (a *AccountId) Value() int64 {
	if a == nil {
		return 0
	}

	return int64(*a)
}
