package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type JWTUtils struct {
	SecretKey []byte
}

func NewJWTUtils(secret string) *JWTUtils {
	return &JWTUtils{
		SecretKey: []byte(secret),
	}
}

func (j *JWTUtils) CreateToken(username string) (string, error) {

	c := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	tokenString, err := claims.SignedString(j.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func (j *JWTUtils) VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return j.SecretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return token, nil
}
