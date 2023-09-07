package service_types

type ManagerPaths struct {
	PeersPath             string `json:"peers_path"`
	TasksPath             string `json:"tasks_path"`
	RecommendationsPath   string `json:"recommendations_path"`
	FriendsPath           string `json:"friends_path"`
	ChecksPath            string `json:"checks_path"`
	P2PPath               string `json:"p2p_path"`
	VerterPath            string `json:"verter_path"`
	XPPath                string `json:"xp_path"`
	TransferredPointsPath string `json:"transferred_points_path"`
	TimeTrackingPath      string `json:"time_tracking_path"`
}
