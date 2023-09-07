package service_types

import "time"

type TransferredPointsRow struct {
	Peer1        string `json:"peer_1"        db:"peer_1"`
	Peer2        string `json:"peer_2"        db:"peer_2"`
	PointsAmount int32  `json:"points_amount" db:"points_amount"`
}

type XpTaskRow struct {
	Peer string `json:"peer" db:"peer"`
	Task string `json:"task" db:"task"`
	Xp   int32  `json:"xp"   db:"xp"`
}

type PeersDontLeaveRow struct {
	Peer string `json:"peer" db:"peer"`
}

type SuccessFailureChecksRow struct {
	Success float64 `json:"success" db:"success"`
	Failure float64 `json:"failure" db:"failure"`
}

type PointsChangeRow struct {
	Peer         string `json:"peer"          db:"peer"`
	PointsChange int32  `json:"points_change" db:"points_change"`
}

type OftenTaskPerDayRow struct {
	Day  time.Time `json:"day,time.RFC3339" db:"day"`
	Task string    `json:"task"             db:"task"`
}

type LastP2PDurationRow struct {
	Duration string `json:"duration" db:"duration"`
}

type ListLastExPeerRow struct {
	Peer string    `json:"peer"             db:"peer"`
	Day  time.Time `json:"day,time.RFC3339" db:"day"`
}

type PeersForP2PRow struct {
	Peer            string `json:"peer"             db:"peer"`
	RecommendedPeer string `json:"recommended_peer" db:"recommended_peer"`
}

type StatisticBlockRow struct {
	StartedBlock1      float64 `json:"started_block_1"       db:"started_block_1"`
	StartedBlock2      float64 `json:"started_block_2"       db:"started_block_2"`
	StartedBothBlocks  float64 `json:"started_both_blocks"   db:"started_both_blocks"`
	DidntStartAnyBlock float64 `json:"didnt_start_any_block" db:"didnt_start_any_block"`
}

type MostFriendlyRow struct {
	Peer         string `json:"peer"          db:"peer"`
	FriendsCount int32  `json:"friends_count" db:"friends_count"`
}

type SuccessAtBirthdayRow struct {
	SuccessFulChecks   float64 `json:"successful_checks"   db:"successful_checks"`
	UnSuccessFulChecks float64 `json:"unsuccessful_checks" db:"unsuccessful_checks"`
}

type PeerXpSumRow struct {
	Peer string `json:"peer" db:"peer"`
	Xp   int32  `json:"xp"   db:"xp"`
}

type PassOneTwoRow struct {
	Peer string `json:"peer" db:"peer"`
}

type PreviousTasksRow struct {
	Task      string `json:"task"       db:"task"`
	PrevCount int32  `json:"prev_count" db:"prev_count"`
}

type SuccessfulDaysRow struct {
	Day string `json:"day" db:"day"`
}

type PeerMostTasksRow struct {
	Peer  string `json:"peer"  db:"peer"`
	Tasks int32  `json:"tasks" db:"tasks"`
}

type PeerMostXpRow struct {
	Peer string `json:"peer" db:"peer"`
	Xp   int32  `json:"xp"   db:"xp"`
}

type MaxTimeDateRow struct {
	Peer string `json:"peer" db:"peer"`
}

type TimePeerByTimeRow struct {
	Peer string `json:"peer" db:"peer"`
}

type EnterPeerByDayRow struct {
	Peer string `json:"peer" db:"peer"`
}

type LastFeastCameRow struct {
	Peer string `json:"peer" db:"peer"`
}

type MoreThenTimePeerRow struct {
	Peer string `json:"peer" db:"peer"`
}

type EarlyEntriesRow struct {
	Month        string  `json:"month"         db:"month"`
	EarlyEntries float64 `json:"early_entries" db:"early_entries"`
}
