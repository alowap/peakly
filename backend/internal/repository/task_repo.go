package repository

import (
	"context"
	"database/sql"
	"fmt"

	"peakly/internal/model"
)

type taskRepo struct {
	db *sql.DB
}

func NewTaskRepo(db *sql.DB) *taskRepo {
	return &taskRepo{db: db}
}

func (r *taskRepo) List(ctx context.Context) ([]*model.Task, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, title, priority, worked_hours, available_hours, created_at, done_at FROM tasks")
	if err != nil {
		return nil, fmt.Errorf("listing tasks: %w", err)
	}
	defer rows.Close()

	var tasks []*model.Task
	for rows.Next() {
		t := &model.Task{}
		if err := rows.Scan(&t.ID, &t.Title, &t.Priority, &t.WorkedHours, &t.AvailableHours, &t.CreatedAt, &t.DoneAt); err != nil {
			return nil, fmt.Errorf("scanning task: %w", err)
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *taskRepo) Create(ctx context.Context, task *model.Task) error {
	return r.db.QueryRowContext(ctx,
		"INSERT INTO tasks (title) VALUES ($1) RETURNING id",
		task.Title,
	).Scan(&task.ID)
}

func (r *taskRepo) GetByID(ctx context.Context, id string) (*model.Task, error) {
	t := &model.Task{}
	err := r.db.QueryRowContext(ctx,
		"SELECT id, title, priority, worked_hours, available_hours, created_at, done_at FROM tasks WHERE id = $1",
		id,
	).Scan(&t.ID, &t.Title, &t.Priority, &t.WorkedHours, &t.AvailableHours, &t.CreatedAt, &t.DoneAt)
	if err != nil {
		return nil, fmt.Errorf("getting task %s: %w", id, err)
	}
	return t, nil
}
