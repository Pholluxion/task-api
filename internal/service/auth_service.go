package service

import (
	"context"

	"github.com/Pholluxion/task-api/internal/model"
	"github.com/Pholluxion/task-api/internal/store"
)

type AuthService interface {
	ValidateUser(ctx context.Context, username, password string) (bool, error)
	RegisterUser(ctx context.Context, user *model.User) error
}

type authService struct {
	userStore store.UserStore
}

func NewAuthService(userStore store.UserStore) AuthService {
	return &authService{userStore: userStore}
}

func (s *authService) RegisterUser(ctx context.Context, user *model.User) error {
	return s.userStore.CreateUser(ctx, user)
}

func (s *authService) ValidateUser(ctx context.Context, username, password string) (bool, error) {
	user, err := s.userStore.ValidateUser(ctx, username, password)

	if err != nil {
		return false, err
	}

	return user != nil, nil
}
