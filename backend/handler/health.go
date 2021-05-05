package handler

import (
	"io"
	"net/http"
)

func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")

	io.WriteString(w, `{"alive": true}`)
}
