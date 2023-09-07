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
	verterTable              = "verter"
	callVerterInsertTemplate = `CALL insert_verter($1, $2, $3, $4)`
)

func (r *repository) AddVerter(ctx context.Context, verter service_types.VerterInsert) error {
	if _, err := r.db.DB.ExecContext(ctx, callVerterInsertTemplate,
		verter.Peer, verter.Task, verter.State, verter.Time); err != nil {
		return fmt.Errorf("insert data: %w", err)
	}
	return nil
}

func (r *repository) GetVerters(ctx context.Context) ([]entity.Verter, error) {
	selectQuery, _, err := goqu.From(verterTable).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var vertersDto []dto.Verter
	if err = r.db.DB.SelectContext(ctx, &vertersDto, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	verters := make([]entity.Verter, len(vertersDto))
	for i, verterDto := range vertersDto {
		verter := entity.Verter{
			Id:      verterDto.Id,
			CheckId: verterDto.CheckId,
			State:   verterDto.State,
			Time:    verterDto.Time,
		}
		verters[i] = verter
	}
	return verters, nil
}

func (r *repository) UpdateVerter(ctx context.Context, verter dto.Verter) error {
	updateQuery, _, err := goqu.Update(verterTable).Set(
		goqu.Record{
			"check_id": verter.CheckId,
			"state":    verter.State,
			"time":     verter.Time,
		}).Where(goqu.C("id").Eq(verter.Id)).
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

func (r *repository) DeleteVerter(ctx context.Context, key string) error {
	deleteQuery, _, err := goqu.Delete(verterTable).
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
