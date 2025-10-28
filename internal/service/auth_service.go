package service

import (
	"context"
	"time"

	"github.com/Pholluxion/task-api/internal/model"
	"github.com/Pholluxion/task-api/internal/store"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type AuthService interface {
	CreateToken(username string) (string, error)
	VerifyToken(tokenString string) (*jwt.Token, error)
	ValidateUser(ctx context.Context, username, password string) (bool, error)
	RegisterUser(ctx context.Context, user *model.User) error
}

type authService struct {
	userStore       store.UserStore
	TokenExpireTime int
	SecretKey       []byte
}

func NewAuthService(userStore store.UserStore, tokenExpireTime int, secretKey []byte) AuthService {
	return &authService{
		userStore:       userStore,
		SecretKey:       secretKey,
		TokenExpireTime: tokenExpireTime,
	}
}

func (s *authService) RegisterUser(ctx context.Context, user *model.User) error {
	return s.userStore.CreateUser(ctx, user)
}

func (s *authService) CreateToken(username string) (string, error) {

	tokenExpireTime := time.Duration(s.TokenExpireTime) * time.Minute

	c := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpireTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	tokenString, err := claims.SignedString(s.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func (s *authService) VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return s.SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *authService) ValidateUser(ctx context.Context, username, password string) (bool, error) {
	user, err := s.userStore.ValidateUser(ctx, username, password)

	if err != nil {
		return false, err
	}

	return user != nil, nil
}
