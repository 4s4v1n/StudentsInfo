package entity

import (
	"encoding/json"
	"time"

	"github.com/sav1nbrave4code/APG3/internal/entity/utils"
)

type Peer struct {
	Nickname string    `json:"nickname"              csv:"nickname"`
	Birthday time.Time `json:"birthday,time.RFC3339,time" csv:"birthday"`
}

func (p *Peer) UnmarshalJSON(b []byte) error {
	type peerAlias Peer
	alias := &struct {
		*peerAlias
		Birthday string `json:"birthday"`
	}{
		peerAlias: (*peerAlias)(p),
	}
	if err := json.Unmarshal(b, &alias); err != nil {
		return err
	}
	t, err := utils.ParseDate(alias.Birthday)
	if err != nil {
		return err
	}
	p.Birthday = t
	return nil
}
