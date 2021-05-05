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

var mockDB = map[int]*model.Person{
	1: {FirstName: "Daniel", LastName: "Danielsson", Age: 29},
	2: {FirstName: "Kajsa", LastName: "Kajsasson", Age: 34},
	3: {FirstName: "Anna", LastName: "Annasson", Age: 64},
	4: {FirstName: "Bengt", LastName: "Bengtsson", Age: 12},
	5: {FirstName: "Olle", LastName: "Ollesson", Age: 98},
}

// https://github.com/gorilla/mux#testing-handlers

func TestListPersons(t *testing.T) {
	// Setup
	h := &Handler{DB: mockDB}

	req, err := http.NewRequest("GET", "/persons", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.ListPersons)

	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
	}

	// The amount of people should match our mocked database
	var persons map[int]*model.Person
	err = json.Unmarshal(rr.Body.Bytes(), &persons)
	if err != nil {
		t.Fatal(err)
	}

	if len(persons) != len(mockDB) {
		t.Errorf("unexpected amount of people: got %v want %v", len(persons), len(mockDB))
	}
}

func TestGetPerson(t *testing.T) {
	// Setup
	h := &Handler{DB: mockDB}

	tt := []struct {
		id   string
		code int
		data string
	}{
		{"2", http.StatusOK, `{"first_name":"Kajsa","last_name":"Kajsasson","age":34}`},
		{"5", http.StatusOK, `{"first_name":"Olle","last_name":"Ollesson","age":98}`},
		{"73", http.StatusNotFound, `{"message":"Not Found"}`},
		{"6144", http.StatusNotFound, `{"message":"Not Found"}`},
		{"ab4a", http.StatusUnprocessableEntity, `{"message":"Unprocessable Entity"}`},
		{"-355", http.StatusUnprocessableEntity, `{"message":"Unprocessable Entity"}`},
	}

	for _, tc := range tt {
		path := fmt.Sprintf("/persons/%s", tc.id)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/persons/{id}", h.GetPerson)
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

func TestDeletePerson(t *testing.T) {
	// Setup
	h := &Handler{DB: mockDB}

	tt := []struct {
		id   string
		code int
	}{
		{"1", http.StatusNoContent},
		{"2", http.StatusNoContent},
		{"abc", http.StatusUnprocessableEntity},
		{"-500", http.StatusUnprocessableEntity},
	}

	for _, tc := range tt {
		path := fmt.Sprintf("/persons/%s", tc.id)
		req, err := http.NewRequest("DELETE", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		router := mux.NewRouter()
		router.HandleFunc("/persons/{id}", h.DeletePerson)
		router.ServeHTTP(rr, req)

		// Check status code
		if status := rr.Code; status != tc.code {
			t.Errorf("wrong status code: got %v want %v",
				status, tc.code)
		}

		// Check if person has been removed
		// TODO....
	}
}
