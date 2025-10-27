package transport

import (
	"errors"
	"net/http"

	"github.com/Pholluxion/task-api/internal/transport/httpx"
	"github.com/Pholluxion/task-api/internal/utils"
)

type AuthHandler struct {
	jwtService *utils.JWTService
}

func NewAuthHandler(jwtUtils *utils.JWTService) *AuthHandler {
	return &AuthHandler{
		jwtService: jwtUtils,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	username, password, ok := r.BasicAuth()
	if !ok {
		httpx.Error(w, http.StatusBadRequest, "Invalid Authorization header")
		return
	}

	isValid, err := validateCredentials(username, password)

	if err != nil {
		httpx.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	if !isValid {
		httpx.Error(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	token, err := h.jwtService.CreateToken(username)

	if err != nil {
		httpx.Error(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	httpx.JSON(w, http.StatusOK, map[string]string{"token": token})
}

func validateCredentials(username, password string) (bool, error) {
	if username == "admin" && password == "password" {
		return true, nil
	}
	return false, errors.New("invalid credentials")
}
