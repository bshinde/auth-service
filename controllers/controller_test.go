package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestRoutes(t *testing.T) {
	router := mux.NewRouter()
	RegisterRoutes(router)

	req := httptest.NewRequest("GET", "/protected", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)
	if w.Code != http.StatusNotFound {
		t.Errorf("expected status 404 for unregistered route, got %d", w.Code)
	}
}
