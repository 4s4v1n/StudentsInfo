package postgres

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/sav1nbrave4code/APG3/internal/entity"
	"github.com/sav1nbrave4code/APG3/internal/entity/dto"
)

const (
	peersTable = "peers"
)

func (r *repository) AddPeer(ctx context.Context, peer dto.Peer) error {
	insertQuery, _, err := goqu.Insert(peersTable).Rows(peer).ToSQL()
	if err != nil {
		return fmt.Errorf("configure query: %w", err)
	}

	if _, err = r.db.DB.ExecContext(ctx, insertQuery); err != nil {
		return fmt.Errorf("insert data: %w", err)
	}
	return nil
}

func (r *repository) GetPeers(ctx context.Context) ([]entity.Peer, error) {
	selectQuery, _, err := goqu.From(peersTable).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var peersDto []dto.Peer
	if err = r.db.DB.SelectContext(ctx, &peersDto, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	peers := make([]entity.Peer, len(peersDto))
	for i, peerDto := range peersDto {
		peer := entity.Peer{
			Nickname: peerDto.Nickname,
			Birthday: peerDto.Birthday,
		}
		peers[i] = peer
	}
	return peers, nil
}

func (r *repository) UpdatePeer(ctx context.Context, peer dto.Peer) error {
	updateQuery, _, err := goqu.Update(peersTable).Set(
		goqu.Record{
			"birthday": peer.Birthday,
		}).Where(goqu.C("nickname").Eq(peer.Nickname)).
		Returning("nickname").ToSQL()
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

func (r *repository) DeletePeer(ctx context.Context, key string) error {
	deleteQuery, _, err := goqu.Delete(peersTable).
		Where(goqu.C("nickname").Eq(key)).
		Returning("nickname").ToSQL()
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
