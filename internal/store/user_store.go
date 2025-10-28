package store

import (
	"context"

	"github.com/Pholluxion/task-api/internal/model"
	"gorm.io/gorm"
)

type UserStore interface {
	CreateUser(ctx context.Context, user *model.User) error
	ValidateUser(ctx context.Context, username, password string) (*model.User, error)
}

type userStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) UserStore {
	return &userStore{db: db}
}

func (s *userStore) CreateUser(ctx context.Context, user *model.User) error {
	return gorm.G[model.User](s.db).Create(ctx, user)
}

func (s *userStore) ValidateUser(ctx context.Context, username, password string) (*model.User, error) {
	user, err := gorm.G[model.User](s.db).Where("username = ? AND password = ?", username, password).First(ctx)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userStore) GetUserByID(ctx context.Context, id uint) (*model.User, error) {
	user, err := gorm.G[model.User](s.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
