package function_manager

import (
	"context"
	"github.com/sav1nbrave4code/APG3/internal/entity/service_types"
	"time"
)

type Service interface {
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
