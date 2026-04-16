package service

import (
	"context"
	"peakly/internal/model"
)

type TaskService interface {
	List(ctx context.Context) ([]*model.Task, error)
	Create(ctx context.Context, title string) (*model.Task, error)
	GetByID(ctx context.Context, id string) (*model.Task, error)
}
