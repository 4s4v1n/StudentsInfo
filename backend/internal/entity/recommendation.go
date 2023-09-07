package entity

type Recommendation struct {
	Id              int64   `json:"id,omitempty"     csv:"id"`
	Peer            *string `json:"peer"             csv:"peer"`
	RecommendedPeer *string `json:"recommended_peer" csv:"recommended_peer"`
}
