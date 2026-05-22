package service

import (
	"testing"
)

func TestPasswordService(t *testing.T) {
	svc := NewPasswordService()
	password := "securepassword"

	hash, err := svc.HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	if hash == password {
		t.Fatal("Hash should not be equal to password")
	}

	err = svc.ComparePassword(password, hash)
	if err != nil {
		t.Errorf("Password comparison failed: %v", err)
	}

	err = svc.ComparePassword("wrongpassword", hash)
	if err == nil {
		t.Error("Password comparison should have failed for wrong password")
	}
}
