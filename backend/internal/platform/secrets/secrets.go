// Package secrets provides credential hashing primitives.
// It is distinct from platform/crypto, which handles PII encryption at rest.
package secrets

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns a bcrypt hash of the password using the default cost.
// The default cost (~10) is what underpins the constant-time anti-enumeration
// approach in the account service — do not lower it.
func HashPassword(password string) (string, error) {
	h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("secrets: hash password: %w", err)
	}
	return string(h), nil
}

// ComparePassword reports whether the password matches the stored hash.
// Returns a non-nil error (bcrypt.ErrMismatchedHashAndPassword) on mismatch.
func ComparePassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
