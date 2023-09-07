package dto

type P2P struct {
	Id           int64   `json:"id,omitempty"  csv:"id"            db:"id"             goqu:"skipinsert"`
	CheckId      *int64  `json:"check_id"      csv:"check_id"      db:"check_id"`
	CheckingPeer *string `json:"checking_peer" csv:"checking_peer" db:"checking_peer"`
	State        string  `json:"state"         csv:"state"         db:"state"          goqu:"enum"`
	Time         string  `json:"time"          csv:"time"          db:"time"`
}
