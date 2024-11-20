package models

import "testing"

func TestCreateUser(t *testing.T) {
	err := CreateUser("user@example.com", "hashedpassword")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	err = CreateUser("user@example.com", "hashedpassword")
	if err == nil {
		t.Errorf("expected error for duplicate user, got nil")
	}
}

func TestGetUser(t *testing.T) {
	CreateUser("user@example.com", "hashedpassword")

	_, err := GetUser("user@example.com")
	if err != nil {
		t.Errorf("expected user, got error: %v", err)
	}

	_, err = GetUser("nonexistent@example.com")
	if err == nil {
		t.Errorf("expected error for nonexistent user, got nil")
	}
}
