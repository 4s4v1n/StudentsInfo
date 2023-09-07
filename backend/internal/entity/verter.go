package entity

type Verter struct {
	Id      int64  `json:"id,omitempty"      csv:"id"`
	CheckId *int64 `json:"check_id"          csv:"check_id"`
	State   string `json:"state"             csv:"state"`
	Time    string `json:"time,time.RFC3339" csv:"time"`
}
