package main

import (
	"backend/handler"
	"backend/model"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func contentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Initialize handler and insert some data
	db := []*model.User{
		{ID: 1, Name: "Daniel1337", Email: "daniel@mail.se"},
		{ID: 2, Name: "AnnaPanna", Email: "anna@mail.se"},
		{ID: 3, Name: "Trollfar", Email: "troll@mail.se"},
		{ID: 4, Name: "Kakburken", Email: "kaka@mail.se"},
	}
	h := &handler.Handler{DB: db}

	r := mux.NewRouter()

	// Middleware
	r.Use(contentTypeMiddleware)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Routes
	r.HandleFunc("/users", h.ListUsers).Methods("GET")
	r.HandleFunc("/users/{id}", h.GetUser).Methods("GET")

	log.Println("Listening on :8000")
	log.Fatal(srv.ListenAndServe())
}
