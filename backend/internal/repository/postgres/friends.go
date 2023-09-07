package postgres

import (
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/sav1nbrave4code/APG3/internal/entity"
	"github.com/sav1nbrave4code/APG3/internal/entity/dto"
)

const (
	friendsTable = "friends"
)

func (r *repository) AddFriends(ctx context.Context, friends dto.Friends) error {
	insertQuery, _, err := goqu.Insert(friendsTable).Rows(friends).ToSQL()
	if err != nil {
		return fmt.Errorf("configure query: %w", err)
	}

	if _, err = r.db.DB.ExecContext(ctx, insertQuery); err != nil {
		return fmt.Errorf("insert data: %w", err)
	}
	return nil
}

func (r *repository) GetFriends(ctx context.Context) ([]entity.Friends, error) {
	selectQuery, _, err := goqu.From(friendsTable).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var friendsSetDto []dto.Friends
	if err = r.db.DB.SelectContext(ctx, &friendsSetDto, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	friendsSet := make([]entity.Friends, len(friendsSetDto))
	for i, friendsDto := range friendsSetDto {
		friends := entity.Friends{
			Id:    friendsDto.Id,
			Peer1: friendsDto.Peer1,
			Peer2: friendsDto.Peer2,
		}
		friendsSet[i] = friends
	}
	return friendsSet, nil
}

func (r *repository) UpdateFriends(ctx context.Context, friends dto.Friends) error {
	updateQuery, _, err := goqu.Update(friendsTable).Set(goqu.Record{
		"peer_1": friends.Peer1,
		"peer_2": friends.Peer2,
	}).Where(goqu.C("id").Eq(friends.Id)).Returning("id").ToSQL()
	if err != nil {
		return fmt.Errorf("configure query: %w", err)
	}

	id := ""
	row := r.db.DB.QueryRowxContext(ctx, updateQuery)
	if err = row.Scan(&id); err != nil {
		return fmt.Errorf("update data: %w", err)
	}
	return nil
}

func (r *repository) DeleteFriends(ctx context.Context, key string) error {
	deleteQuery, _, err := goqu.Delete(friendsTable).
		Where(goqu.C("id").Eq(key)).
		Returning("id").ToSQL()
	if err != nil {
		return fmt.Errorf("configure query: %w", err)
	}

	id := ""
	row := r.db.DB.QueryRowxContext(ctx, deleteQuery)
	if err = row.Scan(&id); err != nil {
		return fmt.Errorf("delete data: %w", err)
	}
	return nil
}
