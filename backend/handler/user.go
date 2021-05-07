package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.DB)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(mux.Vars(r)["id"], 0, 64)
	if err != nil || id == 0 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, `{"message":"%v"}`, http.StatusText(http.StatusUnprocessableEntity))
		return
	}

	// It would probably make more sense to use a map if we want to lookup by id.
	// But this is more how the structure would look like if we used a real database
	for _, user := range h.DB {
		if user.ID == int(id) {
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, `{"message":"%v"}`, http.StatusText(http.StatusNotFound))
}
