package handler

import (
	"backend/model"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

var mockDB = []*model.User{
	{ID: 1, Name: "Daniel1337", Email: "daniel@mail.se"},
	{ID: 2, Name: "AnnaPanna", Email: "anna@mail.se"},
	{ID: 3, Name: "Trollfar", Email: "troll@mail.se"},
	{ID: 4, Name: "Kakburken", Email: "kaka@mail.se"},
}

// https://github.com/gorilla/mux#testing-handlers

func TestListUsers(t *testing.T) {
	// Setup
	h := &Handler{DB: mockDB}

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.ListUsers)

	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
	}

	// The amount of users should match our mocked database
	var users []*model.User
	err = json.Unmarshal(rr.Body.Bytes(), &users)
	if err != nil {
		t.Fatal(err)
	}

	if len(users) != len(mockDB) {
		t.Errorf("unexpected amount of users: got %v want %v", len(users), len(mockDB))
	}
}

func TestGetUser(t *testing.T) {
	// Setup
	h := &Handler{DB: mockDB}

	tt := []struct {
		id   string
		code int
		data string
	}{
		{"2", http.StatusOK, `{"id":2,"name":"AnnaPanna","email":"anna@mail.se"}`},
		{"4", http.StatusOK, `{"id":4,"name":"Kakburken","email":"kaka@mail.se"}`},
		{"73", http.StatusNotFound, `{"message":"Not Found"}`},
		{"6144", http.StatusNotFound, `{"message":"Not Found"}`},
		{"ab4a", http.StatusUnprocessableEntity, `{"message":"Unprocessable Entity"}`},
		{"-355", http.StatusUnprocessableEntity, `{"message":"Unprocessable Entity"}`},
	}

	for _, tc := range tt {
		path := fmt.Sprintf("/users/%s", tc.id)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/users/{id}", h.GetUser)
		router.ServeHTTP(rr, req)

		// Check status code
		if status := rr.Code; status != tc.code {
			t.Errorf("wrong status code: got %v want %v",
				status, tc.code)
		}

		// Check return data
		if strings.TrimSuffix(rr.Body.String(), "\n") != tc.data {
			t.Errorf("unexpected body: got %v want %v",
				rr.Body.String(), tc.data)
		}
	}
}
