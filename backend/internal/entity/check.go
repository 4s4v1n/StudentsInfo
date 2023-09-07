package entity

import (
	"encoding/json"
	"github.com/sav1nbrave4code/APG3/internal/entity/utils"
	"time"
)

type Check struct {
	Id   int64     `json:"id,omitempty"      csv:"id"`
	Peer *string   `json:"peer"              csv:"peer"`
	Task *string   `json:"task"              csv:"task"`
	Date time.Time `json:"date,time.RFC3339" csv:"date"`
}

func (c *Check) UnmarshalJSON(b []byte) error {
	type checkAlias Check
	alias := &struct {
		*checkAlias
		Date string `json:"date"`
	}{
		checkAlias: (*checkAlias)(c),
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return err
	}
	t, err := utils.ParseDate(alias.Date)
	if err != nil {
		return err
	}
	c.Date = t
	return nil
}
