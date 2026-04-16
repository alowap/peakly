package service

import (
	"context"
	"fmt"
	"peakly/internal/model"
)

type taskRepository interface {
	List(ctx context.Context) ([]*model.Task, error)
	Create(ctx context.Context, task *model.Task) error
	GetByID(ctx context.Context, id string) (*model.Task, error)
}

type taskService struct {
	repo taskRepository
}

// 3. constructor
func NewTaskService(repo taskRepository) *taskService {
	return &taskService{repo: repo}
}

// 4. implementations
func (s *taskService) List(ctx context.Context) ([]*model.Task, error) {
	return s.repo.List(ctx)
}

func (s *taskService) Create(ctx context.Context, title string) (*model.Task, error) {
	task := &model.Task{
		Title: title,
	}
	if err := s.repo.Create(ctx, task); err != nil {
		return nil, fmt.Errorf("creating task: %w", err)
	}
	return task, nil
}

func (s *taskService) GetByID(ctx context.Context, id string) (*model.Task, error) {
	return s.repo.GetByID(ctx, id)
}
