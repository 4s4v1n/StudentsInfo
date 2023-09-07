package table_manager

import (
	"context"
	"github.com/sav1nbrave4code/APG3/internal/entity"
	"github.com/sav1nbrave4code/APG3/internal/entity/service_types"
)

type Service interface {
	AddPeer(ctx context.Context, item entity.Peer) error
	GetPeers(ctx context.Context) ([]entity.Peer, error)
	UpdatePeer(ctx context.Context, item entity.Peer) error
	DeletePeer(ctx context.Context, key string) error

	AddTask(ctx context.Context, item entity.Task) error
	GetTasks(ctx context.Context) ([]entity.Task, error)
	UpdateTask(ctx context.Context, item entity.Task) error
	DeleteTask(ctx context.Context, key string) error

	AddCheck(ctx context.Context, item entity.Check) error
	GetChecks(ctx context.Context) ([]entity.Check, error)
	UpdateCheck(ctx context.Context, item entity.Check) error
	DeleteCheck(ctx context.Context, key string) error

	AddP2P(ctx context.Context, item service_types.P2PInsert) error
	GetP2Ps(ctx context.Context) ([]entity.P2P, error)
	UpdateP2P(ctx context.Context, item entity.P2P) error
	DeleteP2P(ctx context.Context, key string) error

	AddVerter(ctx context.Context, item service_types.VerterInsert) error
	GetVerters(ctx context.Context) ([]entity.Verter, error)
	UpdateVerter(ctx context.Context, item entity.Verter) error
	DeleteVerter(ctx context.Context, key string) error

	AddTransferredPoints(ctx context.Context, item entity.TransferredPoints) error
	GetTransferredPoints(ctx context.Context) ([]entity.TransferredPoints, error)
	UpdateTransferredPoints(ctx context.Context, item entity.TransferredPoints) error
	DeleteTransferredPoints(ctx context.Context, key string) error

	AddFriends(ctx context.Context, item entity.Friends) error
	GetFriends(ctx context.Context) ([]entity.Friends, error)
	UpdateFriends(ctx context.Context, item entity.Friends) error
	DeleteFriends(ctx context.Context, key string) error

	AddRecommendation(ctx context.Context, item entity.Recommendation) error
	GetRecommendations(ctx context.Context) ([]entity.Recommendation, error)
	UpdateRecommendation(ctx context.Context, item entity.Recommendation) error
	DeleteRecommendation(ctx context.Context, key string) error

	AddXP(ctx context.Context, item entity.XP) error
	GetXPs(ctx context.Context) ([]entity.XP, error)
	UpdateXP(ctx context.Context, item entity.XP) error
	DeleteXP(ctx context.Context, key string) error

	AddTimeTracking(ctx context.Context, item entity.TimeTracking) error
	GetTimeTracking(ctx context.Context) ([]entity.TimeTracking, error)
	UpdateTimeTracking(ctx context.Context, item entity.TimeTracking) error
	DeleteTimeTracking(ctx context.Context, key string) error
}
