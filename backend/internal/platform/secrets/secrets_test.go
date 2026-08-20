package secrets

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

// TestHashPassword_ReturnsBcryptHash verifies HashPassword produces a hash
// that bcrypt itself accepts for the original password.
func TestHashPassword_ReturnsBcryptHash(t *testing.T) {
	password := "correct horse battery staple"
	h, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword: unexpected error: %v", err)
	}
	if len(h) == 0 {
		t.Fatal("HashPassword: returned empty hash")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(h), []byte(password)); err != nil {
		t.Fatalf("hash does not match original password: %v", err)
	}
}

// TestComparePassword_MatchesAndRejects verifies ComparePassword succeeds for
// the original password and fails for a wrong one.
func TestComparePassword_MatchesAndRejects(t *testing.T) {
	h, err := HashPassword("supersecret-123")
	if err != nil {
		t.Fatalf("HashPassword: %v", err)
	}
	if err := ComparePassword(h, "supersecret-123"); err != nil {
		t.Fatalf("ComparePassword: expected match, got %v", err)
	}
	if err := ComparePassword(h, "wrong-password"); err == nil {
		t.Fatal("ComparePassword: expected mismatch error, got nil")
	}
}
