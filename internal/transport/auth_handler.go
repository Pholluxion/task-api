package transport

import (
	"net/http"

	"github.com/Pholluxion/task-api/internal/model"
	"github.com/Pholluxion/task-api/internal/service"
	"github.com/Pholluxion/task-api/internal/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService service.AuthService
	jwtUtil     utils.JWTUtil
}

func NewAuthHandler(authService service.AuthService, jwtUtil utils.JWTUtil) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		jwtUtil:     jwtUtil,
	}
}

func (h *AuthHandler) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		username, password, ok := ctx.Request.BasicAuth()

		if !ok {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid Authorization header"})
			return
		}

		isValid, err := h.authService.ValidateUser(ctx, username, password)

		if !isValid || err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		token, err := h.jwtUtil.GenerateToken(username)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		ctx.JSON(http.StatusOK, map[string]string{"token": token})
	}

}

func (h *AuthHandler) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var user model.User

		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
			return
		}

		err := h.authService.RegisterUser(ctx, &user)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		ctx.Status(http.StatusCreated)
	}
}
