package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignUpController(t *testing.T) {
	payload := []byte(`{"email":"user@example.com", "password":"password"}`)
	req := httptest.NewRequest("POST", "/signup", bytes.NewBuffer(payload))
	w := httptest.NewRecorder()

	SignUp(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected status 201, got %d", w.Code)
	}
}
