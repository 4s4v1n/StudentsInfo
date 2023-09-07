package repository

import (
	"context"
	"io"
	"time"

	"github.com/sav1nbrave4code/APG3/internal/entity/service_types"

	"github.com/sav1nbrave4code/APG3/internal/entity"
	"github.com/sav1nbrave4code/APG3/internal/entity/dto"
)

type Repository interface {
	TableManager
	DataManager
	FunctionManager
}

type TableManager interface {
	AddPeer(ctx context.Context, item dto.Peer) error
	GetPeers(ctx context.Context) ([]entity.Peer, error)
	UpdatePeer(ctx context.Context, item dto.Peer) error
	DeletePeer(ctx context.Context, key string) error

	AddTask(ctx context.Context, item dto.Task) error
	GetTask(ctx context.Context) ([]entity.Task, error)
	UpdateTask(ctx context.Context, item dto.Task) error
	DeleteTask(ctx context.Context, key string) error

	AddCheck(ctx context.Context, item dto.Check) error
	GetChecks(ctx context.Context) ([]entity.Check, error)
	UpdateCheck(ctx context.Context, item dto.Check) error
	DeleteCheck(ctx context.Context, key string) error

	AddP2P(ctx context.Context, item service_types.P2PInsert) error
	GetP2Ps(ctx context.Context) ([]entity.P2P, error)
	UpdateP2P(ctx context.Context, item dto.P2P) error
	DeleteP2P(ctx context.Context, key string) error

	AddVerter(ctx context.Context, item service_types.VerterInsert) error
	GetVerters(ctx context.Context) ([]entity.Verter, error)
	UpdateVerter(ctx context.Context, item dto.Verter) error
	DeleteVerter(ctx context.Context, key string) error

	AddTransferredPoints(ctx context.Context, item dto.TransferredPoints) error
	GetTransferredPoints(ctx context.Context) ([]entity.TransferredPoints, error)
	UpdateTransferredPoints(ctx context.Context, item dto.TransferredPoints) error
	DeleteTransferredPoints(ctx context.Context, key string) error

	AddFriends(ctx context.Context, item dto.Friends) error
	GetFriends(ctx context.Context) ([]entity.Friends, error)
	UpdateFriends(ctx context.Context, item dto.Friends) error
	DeleteFriends(ctx context.Context, key string) error

	AddRecommendation(ctx context.Context, item dto.Recommendation) error
	GetRecommendations(ctx context.Context) ([]entity.Recommendation, error)
	UpdateRecommendation(ctx context.Context, item dto.Recommendation) error
	DeleteRecommendation(ctx context.Context, key string) error

	AddXP(ctx context.Context, item dto.XP) error
	GetXPs(ctx context.Context) ([]entity.XP, error)
	UpdateXP(ctx context.Context, item dto.XP) error
	DeleteXP(ctx context.Context, key string) error

	AddTimeTracking(ctx context.Context, item dto.TimeTracking) error
	GetTimeTracking(ctx context.Context) ([]entity.TimeTracking, error)
	UpdateTimeTracking(ctx context.Context, item dto.TimeTracking) error
	DeleteTimeTracking(ctx context.Context, key string) error
}

type DataManager interface {
	Import(ctx context.Context, table string, body io.Reader) error
	Export(ctx context.Context, table string) ([]byte, error)
	RawQuery(ctx context.Context, query string) ([]map[string]interface{}, error)
}

type FunctionManager interface {
	FncTransferredPoints(ctx context.Context) ([]service_types.TransferredPointsRow, error)
	FncXpTask(ctx context.Context) ([]service_types.XpTaskRow, error)
	FncPeersDontLeave(ctx context.Context, date time.Time) ([]service_types.PeersDontLeaveRow, error)
	FncSuccessFailureChecks(ctx context.Context) ([]service_types.SuccessFailureChecksRow, error)
	FncPointsChangeV1(ctx context.Context) ([]service_types.PointsChangeRow, error)
	FncPointsChangeV2(ctx context.Context) ([]service_types.PointsChangeRow, error)
	FncOftenTaskPerDay(ctx context.Context) ([]service_types.OftenTaskPerDayRow, error)
	FncLastP2PDuration(ctx context.Context) ([]service_types.LastP2PDurationRow, error)
	FncListLastExPeer(ctx context.Context, ex string) ([]service_types.ListLastExPeerRow, error)
	FncPeersForP2P(ctx context.Context) ([]service_types.PeersForP2PRow, error)
	FncStatisticBlock(ctx context.Context, block1 string, block2 string) ([]service_types.StatisticBlockRow, error)
	FncMostFriendly(ctx context.Context, n int) ([]service_types.MostFriendlyRow, error)
	FncSuccessAtBirthDay(ctx context.Context) ([]service_types.SuccessAtBirthdayRow, error)
	FncPeerXpSum(ctx context.Context) ([]service_types.PeerXpSumRow, error)
	FncPassOneTwo(ctx context.Context, task1 string, task2 string, task3 string) ([]service_types.PassOneTwoRow, error)
	FncPreviousTasks(ctx context.Context) ([]service_types.PreviousTasksRow, error)
	FncSuccessfulDays(ctx context.Context, n int) ([]service_types.SuccessfulDaysRow, error)
	FncPeerMostTasks(ctx context.Context) ([]service_types.PeerMostTasksRow, error)
	FncPeerMostXp(ctx context.Context) ([]service_types.PeerMostXpRow, error)
	FncMaxTimeDate(ctx context.Context, date time.Time) ([]service_types.MaxTimeDateRow, error)
	FncTimePeerByTime(ctx context.Context, time string, n int) ([]service_types.TimePeerByTimeRow, error)
	FncEnterPeerByDay(ctx context.Context, n int, m int) ([]service_types.EnterPeerByDayRow, error)
	FncLastFeastCame(ctx context.Context, date time.Time) ([]service_types.LastFeastCameRow, error)
	FncMoreThenTimePeer(ctx context.Context, time string) ([]service_types.MoreThenTimePeerRow, error)
	FncEarlyEntries(ctx context.Context) ([]service_types.EarlyEntriesRow, error)
}
