// Package hasher provides functions for hashing string and comparing string with hash
package hasher

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
)

// HashPassword hashes provided string with usage SHA256
// Error is returned if provided string is empty
// If password is not empty returned hashed password and nil
func HashPassword(password string) (string, error) {
	if len(password) == 0 {
		return "", errors.New("password can`t be empty")
	}
	sha256Hasher := sha256.New()
	sha256Hasher.Write([]byte(password))
	return hex.EncodeToString(sha256Hasher.Sum(nil)), nil
}

// CheckPasswordHash checks if provided password and hash are the same
// return true if hashed password have the same hash as provided hash
// return false otherwise
func CheckPasswordHash(password, hash string) bool {
	hashedPassword, error := HashPassword(password)
	if error != nil {
		return false
	}
	return hashedPassword == hash
}
