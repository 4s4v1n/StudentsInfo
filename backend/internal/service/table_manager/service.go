package table_manager

import (
	"context"
	"fmt"
	"github.com/sav1nbrave4code/APG3/internal/entity"
	"github.com/sav1nbrave4code/APG3/internal/entity/dto"
	"github.com/sav1nbrave4code/APG3/internal/entity/service_types"
	"github.com/sav1nbrave4code/APG3/internal/repository"
)

type service struct {
	repo repository.Repository
}

func New(repo repository.Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) AddPeer(ctx context.Context, peer entity.Peer) error {
	peerDto := dto.Peer{
		Nickname: peer.Nickname,
		Birthday: peer.Birthday,
	}
	if err := s.repo.AddPeer(ctx, peerDto); err != nil {
		return fmt.Errorf("add peer: %w", err)
	}
	return nil
}

func (s *service) GetPeers(ctx context.Context) ([]entity.Peer, error) {
	items, err := s.repo.GetPeers(ctx)
	if err != nil {
		return nil, fmt.Errorf("get peers: %w", err)
	}
	return items, nil
}

func (s *service) UpdatePeer(ctx context.Context, peer entity.Peer) error {
	peerDto := dto.Peer{
		Nickname: peer.Nickname,
		Birthday: peer.Birthday,
	}
	if err := s.repo.UpdatePeer(ctx, peerDto); err != nil {
		return fmt.Errorf("update peer: %w", err)
	}
	return nil
}

func (s *service) DeletePeer(ctx context.Context, key string) error {
	if err := s.repo.DeletePeer(ctx, key); err != nil {
		return fmt.Errorf("delete peer: %w", err)
	}

	return nil
}

func (s *service) AddTask(ctx context.Context, task entity.Task) error {
	taskDto := dto.Task{
		Title:      task.Title,
		ParentTask: task.ParentTask,
		MaxXp:      task.MaxXp,
	}
	if err := s.repo.AddTask(ctx, taskDto); err != nil {
		return fmt.Errorf("add task: %w", err)
	}
	return nil
}

func (s *service) GetTasks(ctx context.Context) ([]entity.Task, error) {
	items, err := s.repo.GetTask(ctx)
	if err != nil {
		return nil, fmt.Errorf("get tasks: %w", err)
	}
	return items, nil
}

func (s *service) UpdateTask(ctx context.Context, task entity.Task) error {
	taskDto := dto.Task{
		Title:      task.Title,
		ParentTask: task.ParentTask,
		MaxXp:      task.MaxXp,
	}
	if err := s.repo.UpdateTask(ctx, taskDto); err != nil {
		return fmt.Errorf("update task: %w", err)
	}
	return nil
}

func (s *service) DeleteTask(ctx context.Context, key string) error {
	if err := s.repo.DeleteTask(ctx, key); err != nil {
		return fmt.Errorf("delete task: %w", err)
	}

	return nil
}

func (s *service) AddCheck(ctx context.Context, check entity.Check) error {
	checkDto := dto.Check{
		Peer: check.Peer,
		Task: check.Task,
		Date: check.Date,
	}
	if err := s.repo.AddCheck(ctx, checkDto); err != nil {
		return fmt.Errorf("add check: %w", err)
	}
	return nil
}

func (s *service) GetChecks(ctx context.Context) ([]entity.Check, error) {
	items, err := s.repo.GetChecks(ctx)
	if err != nil {
		return nil, fmt.Errorf("get checks: %w", err)
	}
	return items, nil
}

func (s *service) UpdateCheck(ctx context.Context, check entity.Check) error {
	checkDto := dto.Check{
		Id:   check.Id,
		Peer: check.Peer,
		Task: check.Task,
		Date: check.Date,
	}
	if err := s.repo.UpdateCheck(ctx, checkDto); err != nil {
		return fmt.Errorf("update check: %w", err)
	}
	return nil
}

func (s *service) DeleteCheck(ctx context.Context, key string) error {
	if err := s.repo.DeleteCheck(ctx, key); err != nil {
		return fmt.Errorf("delete check: %w", err)
	}

	return nil
}

func (s *service) AddP2P(ctx context.Context, p2p service_types.P2PInsert) error {
	if err := s.repo.AddP2P(ctx, p2p); err != nil {
		return fmt.Errorf("add p2p: %w", err)
	}
	return nil
}

func (s *service) GetP2Ps(ctx context.Context) ([]entity.P2P, error) {
	items, err := s.repo.GetP2Ps(ctx)
	if err != nil {
		return nil, fmt.Errorf("get p2ps: %w", err)
	}
	return items, nil
}

func (s *service) UpdateP2P(ctx context.Context, p2p entity.P2P) error {
	p2pDto := dto.P2P{
		Id:           p2p.Id,
		CheckId:      p2p.CheckId,
		CheckingPeer: p2p.CheckingPeer,
		State:        p2p.State,
		Time:         p2p.Time,
	}
	if err := s.repo.UpdateP2P(ctx, p2pDto); err != nil {
		return fmt.Errorf("update p2p: %w", err)
	}
	return nil
}

func (s *service) DeleteP2P(ctx context.Context, key string) error {
	if err := s.repo.DeleteP2P(ctx, key); err != nil {
		return fmt.Errorf("delete p2p: %w", err)
	}

	return nil
}

func (s *service) AddVerter(ctx context.Context, verter service_types.VerterInsert) error {
	if err := s.repo.AddVerter(ctx, verter); err != nil {
		return fmt.Errorf("add verter: %w", err)
	}
	return nil
}

func (s *service) GetVerters(ctx context.Context) ([]entity.Verter, error) {
	items, err := s.repo.GetVerters(ctx)
	if err != nil {
		return nil, fmt.Errorf("get verters: %w", err)
	}
	return items, nil
}

func (s *service) UpdateVerter(ctx context.Context, verter entity.Verter) error {
	verterDto := dto.Verter{
		Id:      verter.Id,
		CheckId: verter.CheckId,
		State:   verter.State,
		Time:    verter.Time,
	}
	if err := s.repo.UpdateVerter(ctx, verterDto); err != nil {
		return fmt.Errorf("update verter: %w", err)
	}
	return nil
}

func (s *service) DeleteVerter(ctx context.Context, key string) error {
	if err := s.repo.DeleteVerter(ctx, key); err != nil {
		return fmt.Errorf("delete verter: %w", err)
	}

	return nil
}

func (s *service) AddTransferredPoints(ctx context.Context, transfer entity.TransferredPoints) error {
	transferDto := dto.TransferredPoints{
		CheckingPeer: transfer.CheckingPeer,
		CheckedPeer:  transfer.CheckedPeer,
		PointsAmount: transfer.PointsAmount,
	}
	if err := s.repo.AddTransferredPoints(ctx, transferDto); err != nil {
		return fmt.Errorf("add transferred_points: %w", err)
	}
	return nil
}

func (s *service) GetTransferredPoints(ctx context.Context) ([]entity.TransferredPoints, error) {
	items, err := s.repo.GetTransferredPoints(ctx)
	if err != nil {
		return nil, fmt.Errorf("get transferred_points: %w", err)
	}
	return items, nil
}

func (s *service) UpdateTransferredPoints(ctx context.Context, transfer entity.TransferredPoints) error {
	transferDto := dto.TransferredPoints{
		Id:           transfer.Id,
		CheckingPeer: transfer.CheckingPeer,
		CheckedPeer:  transfer.CheckedPeer,
		PointsAmount: transfer.PointsAmount,
	}
	if err := s.repo.UpdateTransferredPoints(ctx, transferDto); err != nil {
		return fmt.Errorf("update transferred_points: %w", err)
	}
	return nil
}

func (s *service) DeleteTransferredPoints(ctx context.Context, key string) error {
	if err := s.repo.DeleteTransferredPoints(ctx, key); err != nil {
		return fmt.Errorf("delete transferred_points: %w", err)
	}

	return nil
}

func (s *service) AddFriends(ctx context.Context, friends entity.Friends) error {
	friendsDto := dto.Friends{
		Peer1: friends.Peer1,
		Peer2: friends.Peer2,
	}
	if err := s.repo.AddFriends(ctx, friendsDto); err != nil {
		return fmt.Errorf("add friends: %w", err)
	}
	return nil
}

func (s *service) GetFriends(ctx context.Context) ([]entity.Friends, error) {
	items, err := s.repo.GetFriends(ctx)
	if err != nil {
		return nil, fmt.Errorf("get friends: %w", err)
	}
	return items, nil
}

func (s *service) UpdateFriends(ctx context.Context, friends entity.Friends) error {
	friendsDto := dto.Friends{
		Id:    friends.Id,
		Peer1: friends.Peer1,
		Peer2: friends.Peer2,
	}
	if err := s.repo.UpdateFriends(ctx, friendsDto); err != nil {
		return fmt.Errorf("update friends: %w", err)
	}
	return nil
}

func (s *service) DeleteFriends(ctx context.Context, key string) error {
	if err := s.repo.DeleteFriends(ctx, key); err != nil {
		return fmt.Errorf("delete friends: %w", err)
	}

	return nil
}

func (s *service) AddRecommendation(ctx context.Context, recommendation entity.Recommendation) error {
	recommendationDto := dto.Recommendation{
		Peer:            recommendation.Peer,
		RecommendedPeer: recommendation.RecommendedPeer,
	}
	if err := s.repo.AddRecommendation(ctx, recommendationDto); err != nil {
		return fmt.Errorf("add recommendation: %w", err)
	}
	return nil
}

func (s *service) GetRecommendations(ctx context.Context) ([]entity.Recommendation, error) {
	items, err := s.repo.GetRecommendations(ctx)
	if err != nil {
		return nil, fmt.Errorf("get recommendations: %w", err)
	}
	return items, nil
}

func (s *service) UpdateRecommendation(ctx context.Context, recommendation entity.Recommendation) error {
	recommendationDto := dto.Recommendation{
		Id:              recommendation.Id,
		Peer:            recommendation.Peer,
		RecommendedPeer: recommendation.RecommendedPeer,
	}
	if err := s.repo.UpdateRecommendation(ctx, recommendationDto); err != nil {
		return fmt.Errorf("update recommendation: %w", err)
	}
	return nil
}

func (s *service) DeleteRecommendation(ctx context.Context, key string) error {
	if err := s.repo.DeleteRecommendation(ctx, key); err != nil {
		return fmt.Errorf("delete recommendation: %w", err)
	}

	return nil
}

func (s *service) AddXP(ctx context.Context, xp entity.XP) error {
	xpDto := dto.XP{
		CheckId:  xp.CheckId,
		XpAmount: xp.XpAmount,
	}
	if err := s.repo.AddXP(ctx, xpDto); err != nil {
		return fmt.Errorf("add xp: %w", err)
	}
	return nil
}

func (s *service) GetXPs(ctx context.Context) ([]entity.XP, error) {
	items, err := s.repo.GetXPs(ctx)
	if err != nil {
		return nil, fmt.Errorf("get xps: %w", err)
	}
	return items, nil
}

func (s *service) UpdateXP(ctx context.Context, xp entity.XP) error {
	xpDto := dto.XP{
		Id:       xp.Id,
		CheckId:  xp.CheckId,
		XpAmount: xp.XpAmount,
	}
	if err := s.repo.UpdateXP(ctx, xpDto); err != nil {
		return fmt.Errorf("update xp: %w", err)
	}
	return nil
}

func (s *service) DeleteXP(ctx context.Context, key string) error {
	if err := s.repo.DeleteXP(ctx, key); err != nil {
		return fmt.Errorf("delete xp: %w", err)
	}

	return nil
}

func (s *service) AddTimeTracking(ctx context.Context, track entity.TimeTracking) error {
	trackDto := dto.TimeTracking{
		Id:    track.Id,
		Peer:  track.Peer,
		Date:  track.Date,
		Time:  track.Time,
		State: track.State,
	}
	if err := s.repo.AddTimeTracking(ctx, trackDto); err != nil {
		return fmt.Errorf("add time_tracking: %w", err)
	}
	return nil
}

func (s *service) GetTimeTracking(ctx context.Context) ([]entity.TimeTracking, error) {
	items, err := s.repo.GetTimeTracking(ctx)
	if err != nil {
		return nil, fmt.Errorf("get time_trackings: %w", err)
	}
	return items, nil
}

func (s *service) UpdateTimeTracking(ctx context.Context, track entity.TimeTracking) error {
	trackDto := dto.TimeTracking{
		Id:    track.Id,
		Peer:  track.Peer,
		Date:  track.Date,
		Time:  track.Time,
		State: track.State,
	}
	if err := s.repo.UpdateTimeTracking(ctx, trackDto); err != nil {
		return fmt.Errorf("update time_tracking: %w", err)
	}
	return nil
}

func (s *service) DeleteTimeTracking(ctx context.Context, key string) error {
	if err := s.repo.DeleteTimeTracking(ctx, key); err != nil {
		return fmt.Errorf("delete time_tracking: %w", err)
	}

	return nil
}
