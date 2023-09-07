package entity

type Friends struct {
	Id    int64   `json:"id,omitempty" csv:"id"`
	Peer1 *string `json:"peer_1"       csv:"peer_1"`
	Peer2 *string `json:"peer_2"       csv:"peer_2"`
}
