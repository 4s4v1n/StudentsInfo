package postgres

import (
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/sav1nbrave4code/APG3/internal/entity"
	"github.com/sav1nbrave4code/APG3/internal/entity/dto"
)

const (
	checksTable = "checks"
)

func (r *repository) AddCheck(ctx context.Context, check dto.Check) error {
	insertQuery, _, err := goqu.Insert(checksTable).Rows(check).ToSQL()
	if err != nil {
		return fmt.Errorf("configure query: %w", err)
	}

	if _, err = r.db.DB.ExecContext(ctx, insertQuery); err != nil {
		return fmt.Errorf("insert data: %w", err)
	}
	return nil
}

func (r *repository) GetChecks(ctx context.Context) ([]entity.Check, error) {
	selectQuery, _, err := goqu.From(checksTable).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var checksDto []dto.Check
	if err = r.db.DB.SelectContext(ctx, &checksDto, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	checks := make([]entity.Check, len(checksDto))
	for i, checkDto := range checksDto {
		check := entity.Check{
			Id:   checkDto.Id,
			Peer: checkDto.Peer,
			Task: checkDto.Task,
			Date: checkDto.Date,
		}
		checks[i] = check
	}
	return checks, nil
}

func (r *repository) UpdateCheck(ctx context.Context, check dto.Check) error {
	updateQuery, _, err := goqu.Update(checksTable).Set(goqu.Record{
		"peer": check.Peer,
		"task": check.Task,
		"date": check.Date,
	}).Where(goqu.C("id").Eq(check.Id)).Returning("id").ToSQL()
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

func (r *repository) DeleteCheck(ctx context.Context, key string) error {
	deleteQuery, _, err := goqu.Delete(checksTable).
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
