package entity

type XP struct {
	Id       int64  `json:"id,omitempty" csv:"id"`
	CheckId  *int64 `json:"check_id"     csv:"check_id"`
	XpAmount int32  `json:"xp_amount"    csv:"xp_amount"`
}
