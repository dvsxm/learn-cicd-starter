package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRespondWithJSON(t *testing.T) {
	rr := httptest.NewRecorder()
	payload := map[string]string{"foo": "bar"}

	respondWithJSON(rr, http.StatusCreated, payload)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected status %v, got %v", http.StatusCreated, rr.Code)
	}

	if rr.Header().Get("Content-Type") != "application/json" {
		t.Errorf("expected Content-Type application/json, got %v", rr.Header().Get("Content-Type"))
	}

	var actual map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &actual); err != nil {
		t.Fatalf("couldn't unmarshal response: %v", err)
	}

	if actual["foo"] != "bar" {
		t.Errorf("expected payload foo=bar, got %v", actual["foo"])
	}
}

func TestRespondWithError(t *testing.T) {
	rr := httptest.NewRecorder()
	msg := "something went wrong"

	respondWithError(rr, http.StatusBadRequest, msg, errors.New("actual error"))

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status %v, got %v", http.StatusBadRequest, rr.Code)
	}

	var actual map[string]string
	if err := json.Unmarshal(rr.Body.Bytes(), &actual); err != nil {
		t.Fatalf("couldn't unmarshal response: %v", err)
	}

	if actual["error"] != msg {
		t.Errorf("expected error message %v, got %v", msg, actual["error"])
	}
}
