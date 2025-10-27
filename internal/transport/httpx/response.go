package httpx

import (
	"net/http"
	"strconv"
)

func ParamID(r *http.Request) (uint, error) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	return uint(id), err
}
