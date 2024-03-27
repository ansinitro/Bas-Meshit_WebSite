package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSignInHandler(t *testing.T) {
	// Open the database connection before running tests
	if err := OpenDB(); err != nil {
		t.Fatalf("Failed to open database connection: %v", err)
	}
	defer func() {
		// Close the database connection after tests
		if err := CloseDB(); err != nil {
			t.Fatalf("Failed to close database connection: %v", err)
		}
	}()

	// Create a request with the required form values in the body
	body := strings.NewReader(`{"email": "15asktt@gmail.com", "passwd": "ansi_4321"}`)
	req, err := http.NewRequest("POST", "/signin", body)
	if err != nil {
		t.Fatal(err)
	}

	// Set the Content-Type header to JSON
	req.Header.Set("Content-Type", "application/json")

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Call the signInHandler function with the ResponseRecorder and the request
	http.HandlerFunc(signInHandler).ServeHTTP(rr, req)

	// Check the status code
	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status %d; got %d", http.StatusOK, rr.Code)
	}
}
