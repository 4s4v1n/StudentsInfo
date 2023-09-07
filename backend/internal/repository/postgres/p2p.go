package postgres

import (
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/sav1nbrave4code/APG3/internal/entity"
	"github.com/sav1nbrave4code/APG3/internal/entity/dto"
	"github.com/sav1nbrave4code/APG3/internal/entity/service_types"
)

const (
	p2pTable              = `p2p`
	callP2PInsertTemplate = `CALL insert_p2p($1, $2, $3, $4, $5)`
)

func (r *repository) AddP2P(ctx context.Context, p2p service_types.P2PInsert) error {
	if _, err := r.db.DB.ExecContext(ctx, callP2PInsertTemplate,
		p2p.CheckedPeer, p2p.CheckingPeer, p2p.Task, p2p.State, p2p.Time); err != nil {
		return fmt.Errorf("insert data: %w", err)
	}
	return nil
}

func (r *repository) GetP2Ps(ctx context.Context) ([]entity.P2P, error) {
	selectQuery, _, err := goqu.From(p2pTable).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var p2psDto []dto.P2P
	if err = r.db.DB.SelectContext(ctx, &p2psDto, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	p2ps := make([]entity.P2P, len(p2psDto))
	for i, p2pDto := range p2psDto {
		p2p := entity.P2P{
			Id:           p2pDto.Id,
			CheckId:      p2pDto.CheckId,
			CheckingPeer: p2pDto.CheckingPeer,
			State:        p2pDto.State,
			Time:         p2pDto.Time,
		}
		p2ps[i] = p2p
	}
	return p2ps, nil
}

func (r *repository) UpdateP2P(ctx context.Context, p2p dto.P2P) error {
	updateQuery, _, err := goqu.Update(p2pTable).Set(
		goqu.Record{
			"check_id":      p2p.CheckId,
			"checking_peer": p2p.CheckingPeer,
			"state":         p2p.State,
			"time":          p2p.Time,
		}).Where(goqu.C("id").Eq(p2p.Id)).
		Returning("id").ToSQL()
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

func (r *repository) DeleteP2P(ctx context.Context, key string) error {
	deleteQuery, _, err := goqu.Delete(p2pTable).
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
