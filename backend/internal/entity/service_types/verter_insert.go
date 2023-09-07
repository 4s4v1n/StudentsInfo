package service_types

type VerterInsert struct {
	Peer  string `json:"peer"  db:"nickname"`
	Task  string `json:"task"  db:"task_name"`
	State string `json:"state" db:"verter_state" goqu:"enum"`
	Time  string `json:"time"  db:"verter_time"`
}
