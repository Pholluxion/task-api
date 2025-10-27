package store

import (
	"context"

	"github.com/Pholluxion/task-api/internal/model"
	"gorm.io/gorm"
)

type TaskStore interface {
	GetAll(ctx context.Context) ([]model.Task, error)
	GetByID(ctx context.Context, id uint) (*model.Task, error)
	Create(ctx context.Context, task *model.Task) (*model.Task, error)
	Update(ctx context.Context, id uint, task model.Task) (*model.Task, error)
	Delete(ctx context.Context, id uint) error
}

type taskStore struct {
	db *gorm.DB
}

func New(db *gorm.DB) TaskStore {
	return &taskStore{db: db}
}

func (s *taskStore) GetAll(ctx context.Context) ([]model.Task, error) {
	tasks, err := gorm.G[model.Task](s.db).Find(ctx)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (s *taskStore) GetByID(ctx context.Context, id uint) (*model.Task, error) {
	task, err := gorm.G[model.Task](s.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (s *taskStore) Create(ctx context.Context, task *model.Task) (*model.Task, error) {
	if err := gorm.G[model.Task](s.db).Create(ctx, task); err != nil {
		return nil, err
	}
	return task, nil
}

func (s *taskStore) Update(ctx context.Context, id uint, updated model.Task) (*model.Task, error) {
	res, err := gorm.G[model.Task](s.db).Where("id = ?", id).Select("*").Updates(ctx, updated)

	if err != nil {
		return nil, err
	}

	if res == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return &updated, nil
}

func (s *taskStore) Delete(ctx context.Context, id uint) error {
	if _, err := gorm.G[model.Task](s.db).Where("id = ?", id).Delete(ctx); err != nil {
		return err
	}
	return nil
}
