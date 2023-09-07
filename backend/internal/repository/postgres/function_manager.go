package postgres

import (
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/sav1nbrave4code/APG3/internal/entity/service_types"
	"time"
)

const (
	fncTransferredPoints    = `fnc_transferred_points`
	fncXpTask               = `fnc_xp_task`
	fncPeersDontLeave       = `fnc_peers_dont_leave`
	fncSuccessFailureChecks = `fnc_success_failure_checks`
	fncPointsChangeV1       = `fnc_points_change_v1`
	fncPointsChangeV2       = `fnc_points_change_v2`
	fncOftenTaskPerDay      = `fnc_often_task_per_day`
	fncLastP2PDuration      = `fnc_last_p2p_duration`
	fncListLastExPeer       = `fnc_list_last_ex_peer`
	fncPeersForP2P          = `fnc_peers_for_p2p`
	fncStatisticBlock       = `fnc_statistic_block`
	fncMostFriendly         = `fnc_most_friendly`
	fncSuccessAtBirthday    = `fnc_success_at_birthday`
	fncPeerXpSum            = `fnc_peer_xp_sum`
	fncPassOneTwo           = `fnc_pass_one_two`
	fncPreviousTasks        = `fnc_previous_tasks`
	fncSuccessfulDays       = `fnc_successful_days`
	fncPeerMostTasks        = `fnc_peer_most_tasks`
	fncPeerMostXp           = `fnc_peer_most_xp`
	fncMaxTimeDate          = `fnc_max_time_date`
	fncTimePeerByTime       = `fnc_time_peer_by_time`
	fncEnterPeerByDay       = `fnc_enter_peer_by_day`
	fncLastFeastCame        = `fnc_last_feast_came`
	fncMoreThenTimePeer     = `fnc_more_then_time_peer`
	fncEarlyEntries         = `fnc_early_entries`
)

func (r *repository) FncTransferredPoints(ctx context.Context) ([]service_types.TransferredPointsRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncTransferredPoints)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var transfer []service_types.TransferredPointsRow
	if err = r.db.DB.SelectContext(ctx, &transfer, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return transfer, nil
}

func (r *repository) FncXpTask(ctx context.Context) ([]service_types.XpTaskRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncXpTask)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var xpTask []service_types.XpTaskRow
	if err = r.db.DB.SelectContext(ctx, &xpTask, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return xpTask, nil
}

func (r *repository) FncPeersDontLeave(ctx context.Context, date time.Time) ([]service_types.PeersDontLeaveRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncPeersDontLeave, date)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var dontLeave []service_types.PeersDontLeaveRow
	if err = r.db.DB.SelectContext(ctx, &dontLeave, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return dontLeave, nil
}

func (r *repository) FncSuccessFailureChecks(ctx context.Context) ([]service_types.SuccessFailureChecksRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncSuccessFailureChecks)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var successFailure []service_types.SuccessFailureChecksRow
	if err = r.db.DB.SelectContext(ctx, &successFailure, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return successFailure, nil
}

func (r *repository) FncPointsChangeV1(ctx context.Context) ([]service_types.PointsChangeRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncPointsChangeV1)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var change []service_types.PointsChangeRow
	if err = r.db.DB.SelectContext(ctx, &change, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return change, nil
}

func (r *repository) FncPointsChangeV2(ctx context.Context) ([]service_types.PointsChangeRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncPointsChangeV2)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var change []service_types.PointsChangeRow
	if err = r.db.DB.SelectContext(ctx, &change, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return change, nil
}

func (r *repository) FncOftenTaskPerDay(ctx context.Context) ([]service_types.OftenTaskPerDayRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncOftenTaskPerDay)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var oftenTasks []service_types.OftenTaskPerDayRow
	if err = r.db.DB.SelectContext(ctx, &oftenTasks, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return oftenTasks, nil
}

func (r *repository) FncLastP2PDuration(ctx context.Context) ([]service_types.LastP2PDurationRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncLastP2PDuration)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var duration []service_types.LastP2PDurationRow
	if err = r.db.DB.SelectContext(ctx, &duration, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return duration, nil
}

func (r *repository) FncListLastExPeer(ctx context.Context, ex string) ([]service_types.ListLastExPeerRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncListLastExPeer, ex)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var lastEx []service_types.ListLastExPeerRow
	if err = r.db.DB.SelectContext(ctx, &lastEx, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return lastEx, nil
}

func (r *repository) FncPeersForP2P(ctx context.Context) ([]service_types.PeersForP2PRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncPeersForP2P)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var peers []service_types.PeersForP2PRow
	if err = r.db.DB.SelectContext(ctx, &peers, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return peers, nil
}

func (r *repository) FncStatisticBlock(ctx context.Context, block1 string, block2 string) ([]service_types.StatisticBlockRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncStatisticBlock, block1, block2)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var statistic []service_types.StatisticBlockRow
	if err = r.db.DB.SelectContext(ctx, &statistic, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return statistic, nil
}

func (r *repository) FncMostFriendly(ctx context.Context, n int) ([]service_types.MostFriendlyRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncMostFriendly, n)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var friendly []service_types.MostFriendlyRow
	if err = r.db.DB.SelectContext(ctx, &friendly, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return friendly, nil
}

func (r *repository) FncSuccessAtBirthDay(ctx context.Context) ([]service_types.SuccessAtBirthdayRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncSuccessAtBirthday)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var success []service_types.SuccessAtBirthdayRow
	if err = r.db.DB.SelectContext(ctx, &success, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return success, nil
}

func (r *repository) FncPeerXpSum(ctx context.Context) ([]service_types.PeerXpSumRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncPeerXpSum)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var sum []service_types.PeerXpSumRow
	if err = r.db.DB.SelectContext(ctx, &sum, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return sum, nil
}

func (r *repository) FncPassOneTwo(ctx context.Context, task1 string, task2 string, task3 string) ([]service_types.PassOneTwoRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncPassOneTwo, task1, task2, task3)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var pass []service_types.PassOneTwoRow
	if err = r.db.DB.SelectContext(ctx, &pass, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return pass, nil
}

func (r *repository) FncPreviousTasks(ctx context.Context) ([]service_types.PreviousTasksRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncPreviousTasks)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var tasks []service_types.PreviousTasksRow
	if err = r.db.DB.SelectContext(ctx, &tasks, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return tasks, nil
}

func (r *repository) FncSuccessfulDays(ctx context.Context, n int) ([]service_types.SuccessfulDaysRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncSuccessfulDays, n)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var days []service_types.SuccessfulDaysRow
	if err = r.db.DB.SelectContext(ctx, &days, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return days, nil
}

func (r *repository) FncPeerMostTasks(ctx context.Context) ([]service_types.PeerMostTasksRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncPeerMostTasks)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var peer []service_types.PeerMostTasksRow
	if err = r.db.DB.SelectContext(ctx, &peer, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return peer, nil
}

func (r *repository) FncPeerMostXp(ctx context.Context) ([]service_types.PeerMostXpRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncPeerMostXp)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var peer []service_types.PeerMostXpRow
	if err = r.db.DB.SelectContext(ctx, &peer, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return peer, nil
}

func (r *repository) FncMaxTimeDate(ctx context.Context, date time.Time) ([]service_types.MaxTimeDateRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncMaxTimeDate, date)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var peer []service_types.MaxTimeDateRow
	if err = r.db.DB.SelectContext(ctx, &peer, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return peer, nil
}

func (r *repository) FncTimePeerByTime(ctx context.Context, time string, n int) ([]service_types.TimePeerByTimeRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncTimePeerByTime, time, n)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var peer []service_types.TimePeerByTimeRow
	if err = r.db.DB.SelectContext(ctx, &peer, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return peer, nil
}

func (r *repository) FncEnterPeerByDay(ctx context.Context, n int, m int) ([]service_types.EnterPeerByDayRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncEnterPeerByDay, n, m)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var peer []service_types.EnterPeerByDayRow
	if err = r.db.DB.SelectContext(ctx, &peer, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return peer, nil
}

func (r *repository) FncLastFeastCame(ctx context.Context, date time.Time) ([]service_types.LastFeastCameRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncLastFeastCame, date)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var feast []service_types.LastFeastCameRow
	if err = r.db.DB.SelectContext(ctx, &feast, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return feast, nil
}

func (r *repository) FncMoreThenTimePeer(ctx context.Context, time string) ([]service_types.MoreThenTimePeerRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncMoreThenTimePeer, time)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var peer []service_types.MoreThenTimePeerRow
	if err = r.db.DB.SelectContext(ctx, &peer, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return peer, nil
}

func (r *repository) FncEarlyEntries(ctx context.Context) ([]service_types.EarlyEntriesRow, error) {
	selectQuery, _, err := goqu.From(goqu.Func(fncEarlyEntries)).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var entries []service_types.EarlyEntriesRow
	if err = r.db.DB.SelectContext(ctx, &entries, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	return entries, nil
}
