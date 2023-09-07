package dto

type XP struct {
	Id       int64  `json:"id,omitempty" csv:"id"        db:"id"         goqu:"skipinsert"`
	CheckId  *int64 `json:"check_id"     csv:"check_id"  db:"check_id"`
	XpAmount int32  `json:"xp_amount"    csv:"xp_amount" db:"xp_amount"`
}
