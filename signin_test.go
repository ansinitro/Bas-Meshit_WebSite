package main

import (
	"testing"
)

func TestIsValidEmail_ValidEmail(t *testing.T) {
	email := "test@example.com"
	valid := IsValidEmail(email)
	if !valid {
		t.Errorf("Expected %s to be valid", email)
	}
}

func TestIsValidEmail_InvalidEmail(t *testing.T) {
	email := "invalid-email"
	valid := IsValidEmail(email)
	if valid {
		t.Errorf("Expected %s to be invalid", email)
	}
}

func TestIsValidEmail_EmptyEmail(t *testing.T) {
	email := ""
	valid := IsValidEmail(email)
	if valid {
		t.Error("Expected empty email to be invalid")
	}
}
