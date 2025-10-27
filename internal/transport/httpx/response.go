package httpx

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, status int, msg string) {
	JSON(w, status, map[string]string{"error": msg})
}

func ParamID(r *http.Request) (uint, error) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	return uint(id), err
}

func Decode[T any](r *http.Request, v *T) error {
	return json.NewDecoder(r.Body).Decode(v)
}
