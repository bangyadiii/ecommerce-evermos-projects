package password

import (
	"golang.org/x/crypto/bcrypt"
)

// use byte slices instead of string
func HashPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}

// use byte slices instead of string
func ComparePassword(hashedPassword []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
