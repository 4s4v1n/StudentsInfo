package postgres

import (
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/sav1nbrave4code/APG3/internal/entity"
	"github.com/sav1nbrave4code/APG3/internal/entity/dto"
)

const (
	xpTable = "xp"
)

func (r *repository) AddXP(ctx context.Context, xp dto.XP) error {
	insertQuery, _, err := goqu.Insert(xpTable).Rows(xp).ToSQL()
	if err != nil {
		return fmt.Errorf("configure query: %w", err)
	}

	if _, err = r.db.DB.ExecContext(ctx, insertQuery); err != nil {
		return fmt.Errorf("insert data: %w", err)
	}
	return nil
}

func (r *repository) GetXPs(ctx context.Context) ([]entity.XP, error) {
	selectQuery, _, err := goqu.From(xpTable).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var xpsDto []dto.XP
	if err = r.db.DB.SelectContext(ctx, &xpsDto, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	xps := make([]entity.XP, len(xpsDto))
	for i, xpDto := range xpsDto {
		xp := entity.XP{
			Id:       xpDto.Id,
			CheckId:  xpDto.CheckId,
			XpAmount: xpDto.XpAmount,
		}
		xps[i] = xp
	}
	return xps, nil
}

func (r *repository) UpdateXP(ctx context.Context, xp dto.XP) error {
	updateQuery, _, err := goqu.Update(xpTable).Set(
		goqu.Record{
			"check_id":  xp.CheckId,
			"xp_amount": xp.XpAmount,
		}).Where(goqu.C("id").Eq(xp.Id)).
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

func (r *repository) DeleteXP(ctx context.Context, key string) error {
	deleteQuery, _, err := goqu.Delete(xpTable).
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
