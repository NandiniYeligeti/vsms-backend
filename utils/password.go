package utils

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword generates a bcrypt hash from the given plaintext password.
func HashPassword(password string) (string, error) {
	if password == "" {
		return "", errors.New("password cannot be empty")
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// CheckPassword compares a bcrypt hashed password with its possible plaintext equivalent.
func CheckPassword(hash, password string) error {
	if hash == "" || password == "" {
		return errors.New("invalid password or hash")
	}
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
