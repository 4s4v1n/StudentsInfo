package entity

import (
	"encoding/json"
	"github.com/sav1nbrave4code/APG3/internal/entity/utils"
	"time"
)

type TimeTracking struct {
	Id    int64     `json:"id,omitempty"      csv:"id"`
	Peer  *string   `json:"peer"              csv:"peer"`
	Date  time.Time `json:"date,time.RFC3339" csv:"date"`
	Time  string    `json:"time"              csv:"time"`
	State int32     `json:"state"             csv:"state"`
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
