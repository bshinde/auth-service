package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Set environment variable for testing
	os.Setenv("JWT_SECRET", "testsecret")

	// Call the LoadConfig function
	LoadConfig()

	// Verify that SecretKey matches the test value
	if SecretKey != "testsecret" {
		t.Errorf("expected JWT_SECRET to be 'testsecret', got '%s'", SecretKey)
	}
}
