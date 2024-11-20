package utils

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRespondWithJSON(t *testing.T) {
	// Create a Response Recorder (simulates an HTTP response)
	rr := httptest.NewRecorder()

	// Sample payload
	payload := map[string]string{"message": "success"}

	// Call the RespondWithJSON function
	RespondWithJSON(rr, http.StatusOK, payload)

	// Assert status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Assert Content-Type
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	// Assert response body
	var responseBody map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, payload, responseBody)
}

func TestRespondWithError(t *testing.T) {
	// Create a Response Recorder (simulates an HTTP response)
	rr := httptest.NewRecorder()

	// Sample error message
	errorMessage := "something went wrong"

	// Call the RespondWithError function
	RespondWithError(rr, http.StatusBadRequest, errorMessage)

	// Assert status code
	assert.Equal(t, http.StatusBadRequest, rr.Code)

	// Assert Content-Type
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	// Assert response body
	var responseBody map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, map[string]string{"error": errorMessage}, responseBody)
}
