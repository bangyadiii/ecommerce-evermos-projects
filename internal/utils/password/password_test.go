package password_test

import (
	"ecommerce-evermos-projects/internal/utils/password"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	pw := []byte("myPassword123")

	// Test HashPassword function
	hashedPassword, err := password.HashPassword(pw)
	if err != nil {
		t.Errorf("HashPassword failed with error: %v", err)
	}

	// Test hash length
	if len(hashedPassword) == 0 {
		t.Errorf("HashPassword returned an empty hash")
	}
}

func TestComparePassword(t *testing.T) {
	pw := []byte("myPassword123")
	incorrectPassword := []byte("wrongPassword")

	// Generate hash to compare against
	hashedPassword, _ := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)

	// Test ComparePassword function with correct password
	err := password.ComparePassword(hashedPassword, pw)
	if err != nil {
		t.Errorf("ComparePassword failed with error: %v", err)
	}

	// Test ComparePassword function with incorrect password
	err = password.ComparePassword(hashedPassword, incorrectPassword)
	if err == nil {
		t.Errorf("ComparePassword did not return an error with incorrect password")
	}
}
