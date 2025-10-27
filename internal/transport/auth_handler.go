package transport

import (
	"errors"
	"net/http"

	"github.com/Pholluxion/task-api/internal/transport/httpx"
	"github.com/Pholluxion/task-api/internal/utils"
)

type AuthHandler struct {
	jwtUtils *utils.JWTUtils
}

func NewAuthHandler(jwtUtils *utils.JWTUtils) *AuthHandler {
	return &AuthHandler{
		jwtUtils: jwtUtils,
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

	token, err := h.jwtUtils.CreateToken(username)

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
