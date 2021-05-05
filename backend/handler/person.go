package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) ListPersons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")

	json.NewEncoder(w).Encode(h.DB)
}

func (h *Handler) GetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")

	id, err := strconv.ParseUint(mux.Vars(r)["id"], 0, 64)
	if err != nil || id == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, `{"message":"%v"}`, http.StatusText(http.StatusUnprocessableEntity))
		return
	}

	value, ok := h.DB[int(id)]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, `{"message":"%v"}`, http.StatusText(http.StatusNotFound))
		return
	}

	json.NewEncoder(w).Encode(value)
}

func (h *Handler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=utf8")

	id, err := strconv.ParseUint(mux.Vars(r)["id"], 0, 64)
	if err != nil || id == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, `{"message":"%v"}`, http.StatusText(http.StatusUnprocessableEntity))
		return
	}

	delete(h.DB, int(id))
	w.WriteHeader(http.StatusNoContent)
}
