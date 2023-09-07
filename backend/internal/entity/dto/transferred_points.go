package dto

type TransferredPoints struct {
	Id           int64   `json:"id,omitempty"  csv:"id"            db:"id"             goqu:"skipinsert"`
	CheckingPeer *string `json:"checking_peer" csv:"checking_peer" db:"checking_peer"`
	CheckedPeer  *string `json:"checked_peer"  csv:"checked_peer"  db:"checked_peer"`
	PointsAmount int64   `json:"points_amount" csv:"points_amount" db:"points_amount"`
}
