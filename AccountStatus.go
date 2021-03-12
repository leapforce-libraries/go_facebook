package facebook

import (
	"encoding/json"
)

type AccountStatus string

func (a *AccountStatus) UnmarshalJSON(b []byte) error {
	var u uint32

	err := json.Unmarshal(b, &u)
	if err != nil {
		return err
	}

	statusMap := make(map[uint32]string)
	statusMap[1] = "ACTIVE"
	statusMap[2] = "DISABLED"
	statusMap[3] = "UNSETTLED"
	statusMap[7] = "PENDING_RISK_REVIEW"
	statusMap[8] = "PENDING_SETTLEMENT"
	statusMap[9] = "IN_GRACE_PERIOD"
	statusMap[100] = "PENDING_CLOSURE"
	statusMap[101] = "CLOSED"
	statusMap[201] = "ANY_ACTIVE"
	statusMap[202] = "ANY_CLOSED"

	status, ok := statusMap[uint32(u)]
	if !ok {
		status = "?"
	}

	*a = AccountStatus(status)
	return nil
}

func (a *AccountStatus) ValuePtr() *string {
	if a == nil {
		return nil
	}

	_a := string(*a)
	return &_a
}
