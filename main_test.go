package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCurrentTimeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/current-time", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(currentTimeHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status 200 OK, got %v", status)
	}

	if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("expected content type application/json, got %v", contentType)
	}

	var response Response
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal("Error decoding response:", err)
	}

	_, err = time.Parse(time.RFC3339, response.CurrentTime)
	if err != nil {
		t.Errorf("expected valid RFC3339 time format, got %v", response.CurrentTime)
	}
}
