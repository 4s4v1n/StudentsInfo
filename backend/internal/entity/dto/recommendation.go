package dto

type Recommendation struct {
	Id              int64   `json:"id,omitempty"     csv:"id"               db:"id"                goqu:"skipinsert"`
	Peer            *string `json:"peer"             csv:"peer"             db:"peer"`
	RecommendedPeer *string `json:"recommended_peer" csv:"recommended_peer" db:"recommended_peer"`
}
