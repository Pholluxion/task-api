package transport

import (
	"errors"
	"net/http"

	"github.com/Pholluxion/task-api/internal/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	jwtService *utils.JWTService
}

func NewAuthHandler(jwtUtils *utils.JWTService) *AuthHandler {
	return &AuthHandler{
		jwtService: jwtUtils,
	}
}

func (h *AuthHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username, password, ok := ctx.Request.BasicAuth()
		if !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Authorization header"})
			return
		}

		isValid, err := validateCredentials(username, password)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !isValid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		token, err := h.jwtService.CreateToken(username)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		ctx.JSON(http.StatusOK, map[string]string{"token": token})
	}

}

func validateCredentials(username, password string) (bool, error) {
	if username == "admin" && password == "password" {
		return true, nil
	}
	return false, errors.New("invalid credentials")
}
