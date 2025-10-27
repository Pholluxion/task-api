package service

import (
	"context"

	"github.com/Pholluxion/task-api/internal/model"
	"github.com/Pholluxion/task-api/internal/store"
)

type TaskService interface {
	GetAll(ctx context.Context) ([]model.Task, error)
	GetByID(ctx context.Context, id uint) (*model.Task, error)
	Create(ctx context.Context, task *model.Task) (*model.Task, error)
	Update(ctx context.Context, id uint, task model.Task) (*model.Task, error)
	Delete(ctx context.Context, id uint) error
}

type taskService struct {
	store store.TaskStore
}

func New(store *store.TaskStore) TaskService {
	return &taskService{store: *store}
}

func (s *taskService) GetAll(ctx context.Context) ([]model.Task, error) {
	return s.store.GetAll(ctx)
}
func (s *taskService) GetByID(ctx context.Context, id uint) (*model.Task, error) {
	return s.store.GetByID(ctx, id)
}

func (s *taskService) Create(ctx context.Context, task *model.Task) (*model.Task, error) {
	return s.store.Create(ctx, task)
}

func (s *taskService) Update(ctx context.Context, id uint, task model.Task) (*model.Task, error) {
	return s.store.Update(ctx, id, task)
}

func (s *taskService) Delete(ctx context.Context, id uint) error {
	return s.store.Delete(ctx, id)
}
