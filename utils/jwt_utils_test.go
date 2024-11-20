package utils

import (
	"auth-service/config"
	"testing"
)

func init() {
	config.SecretKey = "testsecret"
}

func TestGenerateToken(t *testing.T) {
	token, err := GenerateToken("user@example.com")
	if err != nil {
		t.Errorf("failed to generate token: %v", err)
	}

	if token == "" {
		t.Error("expected token, got empty string")
	}
}
