package main

import (
	"backend/handler"
	"backend/model"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// test comment

func main() {
	// Initialize handler and insert some data
	db := map[int]*model.Person{
		1: {FirstName: "Daniel", LastName: "Danielsson", Age: 29},
		2: {FirstName: "Kajsa", LastName: "Kajsasson", Age: 34},
		3: {FirstName: "Anna", LastName: "Annasson", Age: 64},
		4: {FirstName: "Bengt", LastName: "Bengtsson", Age: 12},
		5: {FirstName: "Olle", LastName: "Ollesson", Age: 98},
	}
	h := &handler.Handler{DB: db}

	r := mux.NewRouter()

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Routes
	r.HandleFunc("/health", h.HealthCheck).Methods("GET")
	r.HandleFunc("/persons", h.ListPersons).Methods("GET")
	r.HandleFunc("/persons/{id}", h.GetPerson).Methods("GET")
	r.HandleFunc("/persons/{id}", h.DeletePerson).Methods("DELETE")

	log.Println("Listening on :8000")
	log.Fatal(srv.ListenAndServe())
}
