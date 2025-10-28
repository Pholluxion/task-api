package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTUtil interface {
	GenerateToken(username string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
}

type jwtUtil struct {
	secretKey       []byte
	tokenExpireTime int
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewJWTUtil(secretKey string, tokenExpireTime int) JWTUtil {
	return &jwtUtil{
		secretKey:       []byte(secretKey),
		tokenExpireTime: tokenExpireTime,
	}
}

func (u *jwtUtil) GenerateToken(username string) (string, error) {

	tokenExpireTime := time.Duration(u.tokenExpireTime) * time.Minute

	c := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpireTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	tokenString, err := claims.SignedString(u.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func (u *jwtUtil) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return u.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
