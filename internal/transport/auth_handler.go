package transport

import (
	"net/http"

	"github.com/Pholluxion/task-api/internal/utils"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthHandler struct{}

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {

	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if username == "admin" && password == "password" {
		token, err := utils.CreateToken(username)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"token":"` + token + `"}`))
		return
	}

	http.Error(w, "Unauthorized", http.StatusUnauthorized)
}
