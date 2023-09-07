package entity

type TransferredPoints struct {
	Id           int64   `json:"id,omitempty"  csv:"id"`
	CheckingPeer *string `json:"checking_peer" csv:"checking_peer"`
	CheckedPeer  *string `json:"checked_peer"  csv:"checked_peer"`
	PointsAmount int64   `json:"points_amount" csv:"points_amount"`
}
