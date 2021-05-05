package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// https://github.com/gorilla/mux#testing-handlers

func TestHealthCheck(t *testing.T) {
	// Setup
	h := &Handler{DB: mockDB}

	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.HealthCheck)

	handler.ServeHTTP(rr, req)

	// Check status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check response body
	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
