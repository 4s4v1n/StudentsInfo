package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/doug-martin/goqu/v9"
	"github.com/gocarina/gocsv"
	"github.com/sav1nbrave4code/APG3/internal/entity/dto"
)

func (r *repository) Import(ctx context.Context, table string, body io.Reader) error {
	tx, err := r.db.DB.Beginx()
	defer func() {
		_ = tx.Rollback()
	}()
	if err != nil {
		return fmt.Errorf("cannot start transaction: %w", err)
	}

	// truncateQuery, _, err := goqu.Truncate(table).Cascade().ToSQL()
	// if err != nil {
	// 	return err
	// }

	// if _, err = tx.ExecContext(ctx, truncateQuery); err != nil {
	// 	return err
	// }

	switch table {
	case peersTable:
		var data []dto.Peer
		if err = json.NewDecoder(body).Decode(&data); err != nil {
			return err
		}
		insertQuery, _, err := goqu.Insert(peersTable).
			Rows(data).
			OnConflict(goqu.DoNothing()).
			ToSQL()
		if err != nil {
			return fmt.Errorf("configure insert query peers: %w", err)
		}
		if _, err = tx.ExecContext(ctx, insertQuery); err != nil {
			return fmt.Errorf("insert peers: %w", err)
		}
	case tasksTable:
		var data []dto.Task
		if err = json.NewDecoder(body).Decode(&data); err != nil {
			return err
		}
		insertQuery, _, err := goqu.Insert(tasksTable).
			Rows(data).
			OnConflict(goqu.DoNothing()).
			ToSQL()
		if err != nil {
			return fmt.Errorf("configure insert query tasks: %w", err)
		}
		if _, err = tx.ExecContext(ctx, insertQuery); err != nil {
			return fmt.Errorf("insert tasks: %w", err)
		}
	case recommendationsTable:
		var data []dto.Recommendation
		if err = json.NewDecoder(body).Decode(&data); err != nil {
			return err
		}
		insertQuery, _, err := goqu.Insert(recommendationsTable).
			Rows(data).
			OnConflict(goqu.DoNothing()).
			ToSQL()
		if err != nil {
			return fmt.Errorf("configure insert query recommendations: %w", err)
		}
		if _, err = tx.ExecContext(ctx, insertQuery); err != nil {
			return fmt.Errorf("insert recommendations: %w", err)
		}
	case friendsTable:
		var data []dto.Friends
		if err = json.NewDecoder(body).Decode(&data); err != nil {
			return err
		}
		insertQuery, _, err := goqu.Insert(friendsTable).
			Rows(data).
			OnConflict(goqu.DoNothing()).
			ToSQL()
		if err != nil {
			return fmt.Errorf("configure insert query friends: %w", err)
		}
		if _, err = tx.ExecContext(ctx, insertQuery); err != nil {
			return fmt.Errorf("insert friends: %w", err)
		}
	case checksTable:
		var data []dto.Check
		if err = json.NewDecoder(body).Decode(&data); err != nil {
			return err
		}
		insertQuery, _, err := goqu.Insert(checksTable).
			Rows(data).
			OnConflict(goqu.DoNothing()).
			ToSQL()
		if err != nil {
			return fmt.Errorf("configure insert query checks: %w", err)
		}
		if _, err = tx.ExecContext(ctx, insertQuery); err != nil {
			return fmt.Errorf("insert checks: %w", err)
		}
	case p2pTable:
		var data []dto.P2P
		if err = json.NewDecoder(body).Decode(&data); err != nil {
			return err
		}
		insertQuery, _, err := goqu.Insert(p2pTable).
			Rows(data).
			OnConflict(goqu.DoNothing()).
			ToSQL()
		if err != nil {
			return fmt.Errorf("configure insert query p2ps: %w", err)
		}
		if _, err = tx.ExecContext(ctx, insertQuery); err != nil {
			return fmt.Errorf("insert p2ps: %w", err)
		}
	case verterTable:
		var data []dto.Verter
		if err = json.NewDecoder(body).Decode(&data); err != nil {
			return err
		}
		insertQuery, _, err := goqu.Insert(verterTable).
			Rows(data).
			OnConflict(goqu.DoNothing()).
			ToSQL()
		if err != nil {
			return fmt.Errorf("configure insert query verters: %w", err)
		}
		if _, err = tx.ExecContext(ctx, insertQuery); err != nil {
			return fmt.Errorf("insert verters: %w", err)
		}
	case xpTable:
		var data []dto.XP
		if err = json.NewDecoder(body).Decode(&data); err != nil {
			return err
		}
		insertQuery, _, err := goqu.Insert(xpTable).
			Rows(data).
			OnConflict(goqu.DoNothing()).
			ToSQL()
		if err != nil {
			return fmt.Errorf("configure insert query xps: %w", err)
		}
		if _, err = tx.ExecContext(ctx, insertQuery); err != nil {
			return fmt.Errorf("insert xps: %w", err)
		}
	case transferredPointsTable:
		var data []dto.TransferredPoints
		if err = json.NewDecoder(body).Decode(&data); err != nil {
			return err
		}
		insertQuery, _, err := goqu.Insert(transferredPointsTable).
			Rows(data).
			OnConflict(goqu.DoNothing()).
			ToSQL()
		if err != nil {
			return fmt.Errorf("configure insert query transferred points: %w", err)
		}
		if _, err = tx.ExecContext(ctx, insertQuery); err != nil {
			return fmt.Errorf("insert transferred points: %w", err)
		}
	case timeTrackingTable:
		var data []dto.TimeTracking
		if err = json.NewDecoder(body).Decode(&data); err != nil {
			return err
		}
		insertQuery, _, err := goqu.Insert(timeTrackingTable).
			Rows(data).
			OnConflict(goqu.DoNothing()).
			ToSQL()
		if err != nil {
			return fmt.Errorf("configure insert query time tracking: %w", err)
		}
		if _, err = tx.ExecContext(ctx, insertQuery); err != nil {
			return fmt.Errorf("insert time tracking: %w", err)
		}
	default:
		return fmt.Errorf("unknown table: %s", table)
	}

	return tx.Commit()
}

func (r *repository) Export(ctx context.Context, table string) ([]byte, error) {
	var data any
	var err error

	switch table {
	case checksTable:
		data, err = r.GetChecks(ctx)
		if err != nil {
			return nil, err
		}
	case friendsTable:
		data, err = r.GetFriends(ctx)
		if err != nil {
			return nil, err
		}
	case p2pTable:
		data, err = r.GetP2Ps(ctx)
		if err != nil {
			return nil, err
		}
	case peersTable:
		data, err = r.GetPeers(ctx)
		if err != nil {
			return nil, err
		}
	case recommendationsTable:
		data, err = r.GetRecommendations(ctx)
		if err != nil {
			return nil, err
		}
	case tasksTable:
		data, err = r.GetTask(ctx)
		if err != nil {
			return nil, err
		}
	case timeTrackingTable:
		data, err = r.GetTimeTracking(ctx)
		if err != nil {
			return nil, err
		}
	case transferredPointsTable:
		data, err = r.GetTransferredPoints(ctx)
		if err != nil {
			return nil, err
		}
	case verterTable:
		data, err = r.GetVerters(ctx)
		if err != nil {
			return nil, err
		}
	case xpTable:
		data, err = r.GetXPs(ctx)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unknown table: %s", table)
	}

	return gocsv.MarshalBytes(data)
}

func (r *repository) RawQuery(ctx context.Context, query string) ([]map[string]interface{}, error) {
	rows, err := r.db.DB.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := []map[string]interface{}{}
	for rows.Next() {
		res := make(map[string]interface{})
		if err := rows.MapScan(res); err != nil {
			return nil, err
		}
		results = append(results, res)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return results, nil
}
