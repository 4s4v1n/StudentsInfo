package postgres

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/sav1nbrave4code/APG3/internal/entity"
	"github.com/sav1nbrave4code/APG3/internal/entity/dto"
)

const (
	transferredPointsTable = "transferred_points"
)

func (r *repository) AddTransferredPoints(ctx context.Context, transfer dto.TransferredPoints) error {
	insertQuery, _, err := goqu.Insert(transferredPointsTable).Rows(transfer).ToSQL()
	if err != nil {
		return fmt.Errorf("configure query: %w", err)
	}

	if _, err = r.db.DB.ExecContext(ctx, insertQuery); err != nil {
		return fmt.Errorf("insert data: %w", err)
	}
	return nil
}

func (r *repository) GetTransferredPoints(ctx context.Context) ([]entity.TransferredPoints, error) {
	selectQuery, _, err := goqu.From(transferredPointsTable).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var transferredPointsDto []dto.TransferredPoints
	if err = r.db.DB.SelectContext(ctx, &transferredPointsDto, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	transferredPoints := make([]entity.TransferredPoints, len(transferredPointsDto))
	for i, transferDto := range transferredPointsDto {
		transfer := entity.TransferredPoints{
			Id:           transferDto.Id,
			CheckingPeer: transferDto.CheckingPeer,
			CheckedPeer:  transferDto.CheckedPeer,
			PointsAmount: transferDto.PointsAmount,
		}
		transferredPoints[i] = transfer
	}
	return transferredPoints, nil
}

func (r *repository) UpdateTransferredPoints(ctx context.Context, transfer dto.TransferredPoints) error {
	updateQuery, _, err := goqu.Update(transferredPointsTable).Set(goqu.Record{
		"checking_peer": transfer.CheckingPeer,
		"checked_peer":  transfer.CheckedPeer,
		"points_amount": transfer.PointsAmount,
	}).Where(goqu.C("id").Eq(transfer.Id)).Returning("id").ToSQL()
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

func (r *repository) DeleteTransferredPoints(ctx context.Context, key string) error {
	deleteQuery, _, err := goqu.Delete(transferredPointsTable).
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
