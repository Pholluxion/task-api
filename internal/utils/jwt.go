package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func CreateToken(username string) (string, error) {

	var JWT_SECRET = GetString("JWT_SECRET", "")
	var secretKey = []byte(JWT_SECRET)

	c := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	fmt.Println("Secret Key:", string(secretKey))

	tokenString, err := claims.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		var JWT_SECRET = GetString("JWT_SECRET", "")
		var secretKey = []byte(JWT_SECRET)
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return token, nil
}
