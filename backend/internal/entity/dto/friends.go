package dto

type Friends struct {
	Id    int64   `json:"id,omitempty" csv:"id"     db:"id"      goqu:"skipinsert"`
	Peer1 *string `json:"peer_1"       csv:"peer_1" db:"peer_1"`
	Peer2 *string `json:"peer_2"       csv:"peer_2" db:"peer_2"`
}
