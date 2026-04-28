package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword generates a bcrypt hash of the password for secure storage
// Uses default bcrypt cost factor for security
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword compares a plain text password with its bcrypt hash
// Returns error if password doesn't match the hash
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
