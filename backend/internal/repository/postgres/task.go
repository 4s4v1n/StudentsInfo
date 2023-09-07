package postgres

import (
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/sav1nbrave4code/APG3/internal/entity"
	"github.com/sav1nbrave4code/APG3/internal/entity/dto"
)

const (
	tasksTable = "tasks"
)

func (r *repository) AddTask(ctx context.Context, task dto.Task) error {
	if task.ParentTask != nil && *task.ParentTask == "" {
		task.ParentTask = nil
	}

	insertQuery, _, err := goqu.Insert(tasksTable).Rows(task).ToSQL()
	if err != nil {
		return fmt.Errorf("configure query: %w", err)
	}

	if _, err = r.db.DB.ExecContext(ctx, insertQuery); err != nil {
		return fmt.Errorf("insert data: %w", err)
	}
	return nil
}

func (r *repository) GetTask(ctx context.Context) ([]entity.Task, error) {
	selectQuery, _, err := goqu.From(tasksTable).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("configure query: %w", err)
	}

	var tasksDto []dto.Task
	if err = r.db.DB.SelectContext(ctx, &tasksDto, selectQuery); err != nil {
		return nil, fmt.Errorf("select data: %w", err)
	}

	tasks := make([]entity.Task, len(tasksDto))
	for i, taskDto := range tasksDto {
		task := entity.Task{
			Title:      taskDto.Title,
			ParentTask: taskDto.ParentTask,
			MaxXp:      taskDto.MaxXp,
		}
		tasks[i] = task
	}
	return tasks, nil
}

func (r *repository) UpdateTask(ctx context.Context, task dto.Task) error {
	updateQuery, _, err := goqu.Update(tasksTable).Set(goqu.Record{
		"parent_task": task.ParentTask,
		"max_xp":      task.MaxXp,
	}).Where(goqu.C("title").Eq(task.Title)).Returning("title").ToSQL()
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

func (r *repository) DeleteTask(ctx context.Context, key string) error {
	deleteQuery, _, err := goqu.Delete(tasksTable).
		Where(goqu.C("title").Eq(key)).
		Returning("title").ToSQL()
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
