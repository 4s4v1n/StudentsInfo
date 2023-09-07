package postgres

import (
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/sav1nbrave4code/APG3/internal/entity"
	"github.com/sav1nbrave4code/APG3/internal/entity/dto"
)

const (
	timeTrackingTable = "time_tracking"
)

func (r *repository) AddTimeTracking(ctx context.Context, track dto.TimeTracking) error {
	insertQuery, _, err := goqu.Insert(timeTrackingTable).Rows(track).ToSQL()
	if err != nil {
		return fmt.Errorf("configure query: %w", err)
	}

	if _, err = r.db.DB.ExecContext(ctx, insertQuery); err != nil {
		return fmt.Errorf("insert data: %w", err)
	}
	return nil
}

func (r *repository) GetTimeTracking(ctx context.Context) ([]entity.TimeTracking, error) {
	selectQuery, _, err := goqu.From(timeTrackingTable).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var tracksDto []dto.TimeTracking
	if err = r.db.DB.SelectContext(ctx, &tracksDto, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	tracks := make([]entity.TimeTracking, len(tracksDto))
	for i, trackDto := range tracksDto {
		track := entity.TimeTracking{
			Id:    trackDto.Id,
			Peer:  trackDto.Peer,
			Date:  trackDto.Date,
			Time:  trackDto.Time,
			State: trackDto.State,
		}
		tracks[i] = track
	}
	return tracks, nil
}

func (r *repository) UpdateTimeTracking(ctx context.Context, track dto.TimeTracking) error {
	updateQuery, _, err := goqu.Update(timeTrackingTable).Set(goqu.Record{
		"peer":  track.Peer,
		"date":  track.Date,
		"time":  track.Time,
		"state": track.State,
	}).Where(goqu.C("id").Eq(track.Id)).Returning("id").ToSQL()
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

func (r *repository) DeleteTimeTracking(ctx context.Context, key string) error {
	deleteQuery, _, err := goqu.Delete(timeTrackingTable).
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
