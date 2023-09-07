package dto

import (
	"encoding/json"
	"time"

	"github.com/sav1nbrave4code/APG3/internal/entity/utils"
)

type TimeTracking struct {
	Id    int64     `json:"id,omitempty" csv:"id"    db:"id"      goqu:"skipinsert"`
	Peer  *string   `json:"peer"         csv:"peer"  db:"peer"`
	Date  time.Time `json:"date"         csv:"date"  db:"date"`
	Time  string    `json:"time"         csv:"time"  db:"time"`
	State int32     `json:"state"        csv:"state" db:"state"`
}

func (tt *TimeTracking) UnmarshalJSON(b []byte) error {
	type timeTrackingAlias TimeTracking
	alias := &struct {
		*timeTrackingAlias
		Date string `json:"date"`
	}{
		timeTrackingAlias: (*timeTrackingAlias)(tt),
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return err
	}
	t, err := utils.ParseDate(alias.Date)
	if err != nil {
		return err
	}
	tt.Date = t
	return nil
}
