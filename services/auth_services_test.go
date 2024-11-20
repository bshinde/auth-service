package services

import "testing"

func TestSignUp(t *testing.T) {
	err := SignUp("user@example.com", "password")
	if err != nil {
		t.Errorf("failed to sign up: %v", err)
	}

	err = SignUp("user@example.com", "password")
	if err == nil {
		t.Error("expected error for duplicate signup, got nil")
	}
}

func TestSignIn(t *testing.T) {
	SignUp("user@example.com", "password")

	_, err := SignIn("user@example.com", "password")
	if err != nil {
		t.Errorf("failed to sign in: %v", err)
	}

	_, err = SignIn("user@example.com", "wrongpassword")
	if err == nil {
		t.Error("expected error for incorrect password, got nil")
	}

	_, err = SignIn("nonexistent@example.com", "password")
	if err == nil {
		t.Error("expected error for nonexistent user, got nil")
	}
}
