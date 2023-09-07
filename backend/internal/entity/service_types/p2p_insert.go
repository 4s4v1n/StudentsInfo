package service_types

type P2PInsert struct {
	CheckedPeer  string `json:"checked_peer"  db:"checked_peer"`
	CheckingPeer string `json:"checking_peer" db:"checking_peer"`
	Task         string `json:"task"          db:"task_name"`
	State        string `json:"state"         db:"state"          goqu:"enum"`
	Time         string `json:"time"          db:"time"`
}
