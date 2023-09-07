package function_manager

import (
	"context"
	"fmt"
	"github.com/sav1nbrave4code/APG3/internal/entity/service_types"
	"github.com/sav1nbrave4code/APG3/internal/repository"
	"time"
)

type service struct {
	repo repository.Repository
}

func New(repo repository.Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) FncTransferredPoints(ctx context.Context) ([]service_types.TransferredPointsRow, error) {
	items, err := s.repo.FncTransferredPoints(ctx)
	if err != nil {
		return nil, fmt.Errorf("get fnc_transferred_points: %w", err)
	}
	return items, nil
}

func (s *service) FncXpTask(ctx context.Context) ([]service_types.XpTaskRow, error) {
	items, err := s.repo.FncXpTask(ctx)
	if err != nil {
		return nil, fmt.Errorf("get fnc_xp_task: %w", err)
	}
	return items, nil
}

func (s *service) FncPeersDontLeave(ctx context.Context, date time.Time) ([]service_types.PeersDontLeaveRow, error) {
	items, err := s.repo.FncPeersDontLeave(ctx, date)
	if err != nil {
		return nil, fmt.Errorf("get fnc_peers_dont_leave: %w", err)
	}
	return items, nil
}

func (s *service) FncSuccessFailureChecks(ctx context.Context) ([]service_types.SuccessFailureChecksRow, error) {
	items, err := s.repo.FncSuccessFailureChecks(ctx)
	if err != nil {
		return nil, fmt.Errorf("get fnc_success_failure_checks: %w", err)
	}
	return items, nil
}

func (s *service) FncPointsChangeV1(ctx context.Context) ([]service_types.PointsChangeRow, error) {
	items, err := s.repo.FncPointsChangeV1(ctx)
	if err != nil {
		return nil, fmt.Errorf("get fnc_points_change_v1: %w", err)
	}
	return items, nil
}

func (s *service) FncPointsChangeV2(ctx context.Context) ([]service_types.PointsChangeRow, error) {
	items, err := s.repo.FncPointsChangeV2(ctx)
	if err != nil {
		return nil, fmt.Errorf("get fnc_points_change_v2: %w", err)
	}
	return items, nil
}

func (s *service) FncOftenTaskPerDay(ctx context.Context) ([]service_types.OftenTaskPerDayRow, error) {
	items, err := s.repo.FncOftenTaskPerDay(ctx)
	if err != nil {
		return nil, fmt.Errorf("get fnc_often_task_per_day: %w", err)
	}
	return items, nil
}

func (s *service) FncLastP2PDuration(ctx context.Context) ([]service_types.LastP2PDurationRow, error) {
	items, err := s.repo.FncLastP2PDuration(ctx)
	if err != nil {
		return nil, fmt.Errorf("get fnc_last_p2p_duration: %w", err)
	}
	return items, nil
}

func (s *service) FncListLastExPeer(ctx context.Context, ex string) ([]service_types.ListLastExPeerRow, error) {
	items, err := s.repo.FncListLastExPeer(ctx, ex)
	if err != nil {
		return nil, fmt.Errorf("get fnc_list_last_ex_peer: %w", err)
	}
	return items, nil
}

func (s *service) FncPeersForP2P(ctx context.Context) ([]service_types.PeersForP2PRow, error) {
	items, err := s.repo.FncPeersForP2P(ctx)
	if err != nil {
		return nil, fmt.Errorf("get fnc_peers_for_p2p: %w", err)
	}
	return items, nil
}

func (s *service) FncStatisticBlock(ctx context.Context, block1 string, block2 string) ([]service_types.StatisticBlockRow, error) {
	items, err := s.repo.FncStatisticBlock(ctx, block1, block2)
	if err != nil {
		return nil, fmt.Errorf("get fnc_statistic_block: %w", err)
	}
	return items, nil
}

func (s *service) FncMostFriendly(ctx context.Context, n int) ([]service_types.MostFriendlyRow, error) {
	items, err := s.repo.FncMostFriendly(ctx, n)
	if err != nil {
		return nil, fmt.Errorf("get fnc_most_friendly: %w", err)
	}
	return items, nil
}

func (s *service) FncSuccessAtBirthDay(ctx context.Context) ([]service_types.SuccessAtBirthdayRow, error) {
	items, err := s.repo.FncSuccessAtBirthDay(ctx)
	if err != nil {
		return nil, fmt.Errorf("get fnc_success_at_birthday: %w", err)
	}
	return items, nil
}

func (s *service) FncPeerXpSum(ctx context.Context) ([]service_types.PeerXpSumRow, error) {
	items, err := s.repo.FncPeerXpSum(ctx)
	if err != nil {
		return nil, fmt.Errorf("get fnc_peer_xp_sum: %w", err)
	}
	return items, nil
}

func (s *service) FncPassOneTwo(ctx context.Context, task1 string, task2 string, task3 string) ([]service_types.PassOneTwoRow, error) {
	items, err := s.repo.FncPassOneTwo(ctx, task1, task2, task3)
	if err != nil {
		return nil, fmt.Errorf("get fnc_pass_one_two: %w", err)
	}
	return items, nil
}

func (s *service) FncPreviousTasks(ctx context.Context) ([]service_types.PreviousTasksRow, error) {
	items, err := s.repo.FncPreviousTasks(ctx)
	if err != nil {
		return nil, fmt.Errorf("get fnc_previous_tasks: %w", err)
	}
	return items, nil
}

func (s *service) FncSuccessfulDays(ctx context.Context, n int) ([]service_types.SuccessfulDaysRow, error) {
	items, err := s.repo.FncSuccessfulDays(ctx, n)
	if err != nil {
		return nil, fmt.Errorf("get fnc_successful_days: %w", err)
	}
	return items, nil
}

func (s *service) FncPeerMostTasks(ctx context.Context) ([]service_types.PeerMostTasksRow, error) {
	items, err := s.repo.FncPeerMostTasks(ctx)
	if err != nil {
		return nil, fmt.Errorf("get fnc_peer_most_tasks: %w", err)
	}
	return items, nil
}

func (s *service) FncPeerMostXp(ctx context.Context) ([]service_types.PeerMostXpRow, error) {
	items, err := s.repo.FncPeerMostXp(ctx)
	if err != nil {
		return nil, fmt.Errorf("get fnc_peer_most_xp: %w", err)
	}
	return items, nil
}

func (s *service) FncMaxTimeDate(ctx context.Context, date time.Time) ([]service_types.MaxTimeDateRow, error) {
	items, err := s.repo.FncMaxTimeDate(ctx, date)
	if err != nil {
		return nil, fmt.Errorf("get fnc_max_time_date: %w", err)
	}
	return items, nil
}

func (s *service) FncTimePeerByTime(ctx context.Context, time string, n int) ([]service_types.TimePeerByTimeRow, error) {
	items, err := s.repo.FncTimePeerByTime(ctx, time, n)
	if err != nil {
		return nil, fmt.Errorf("get fnc_time_peer_by_time: %w", err)
	}
	return items, nil
}

func (s *service) FncEnterPeerByDay(ctx context.Context, n int, m int) ([]service_types.EnterPeerByDayRow, error) {
	items, err := s.repo.FncEnterPeerByDay(ctx, n, m)
	if err != nil {
		return nil, fmt.Errorf("get fnc_enter_peer_by_dat: %w", err)
	}
	return items, nil
}

func (s *service) FncLastFeastCame(ctx context.Context, date time.Time) ([]service_types.LastFeastCameRow, error) {
	items, err := s.repo.FncLastFeastCame(ctx, date)
	if err != nil {
		return nil, fmt.Errorf("get fnc_last_feast_came: %w", err)
	}
	return items, nil
}

func (s *service) FncMoreThenTimePeer(ctx context.Context, time string) ([]service_types.MoreThenTimePeerRow, error) {
	items, err := s.repo.FncMoreThenTimePeer(ctx, time)
	if err != nil {
		return nil, fmt.Errorf("get fnc_more_then_time_peer: %w", err)
	}
	return items, nil
}

func (s *service) FncEarlyEntries(ctx context.Context) ([]service_types.EarlyEntriesRow, error) {
	items, err := s.repo.FncEarlyEntries(ctx)
	if err != nil {
		return nil, fmt.Errorf("get fnc_early_entries: %w", err)
	}
	return items, nil
}
