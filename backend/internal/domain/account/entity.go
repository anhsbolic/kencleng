// Package account implements the account domain: user identity,
// authentication identities (email/password, google, phone OTP), and
// single-use verification tokens.
//
// Entity shape. Entities are domain objects carrying plaintext PII
// (PrimaryEmail, Identifier). Ciphertext and HMAC hashes are storage
// concerns owned by the repository adapter (RepositoryDB), which
// encrypts at the storage boundary per the PII pattern in
// docs/project/kencleng-backend-tech-stack.md. The service therefore
// never handles raw ciphertext. On the write path the caller sets the
// plaintext field; the adapter encrypts it into the BYTEA column and
// computes the HMAC into the *_hash column, then clears the plaintext
// field so it is not retained past the insert. On reads, the adapter
// populates the non-encrypted fields; the plaintext field is left
// empty unless a caller explicitly needs decryption (the current
// register/verify/resend flows look up by hash and do not read
// plaintext back, so decryption is not on the hot path).
package account

import (
	"time"

	"github.com/google/uuid"
)

// User is the top-level account entity. PrimaryEmail is the plaintext
// email (set by the caller on insert; the adapter encrypts it into the
// primary_email BYTEA column and the HMAC into primary_email_hash).
type User struct {
	ID           uuid.UUID
	Name         string
	PrimaryEmail string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// AuthIdentity is a provider-scoped credential binding for a User.
// ProviderType is one of "email_password", "google", "phone_otp".
// Identifier is the plaintext identifier (set by the caller on insert;
// the adapter encrypts it). CredentialSecret holds the bcrypt hash for
// email_password and is nil for google.
type AuthIdentity struct {
	ID               uuid.UUID
	UserID           uuid.UUID
	ProviderType     string
	Identifier       string
	CredentialSecret *string
	VerifiedAt       *time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

// AuthToken is a single-use, time-bound verification or reset token.
// TokenHash is the SHA-256 hex digest of the plain token (never
// encrypted — it is already a hash). Redemption is guarded by the
// 3-clause predicate
// (used_at IS NULL AND revoked_at IS NULL AND expires_at > now())
// per INV-account-08 — see Repository.RedeemToken.
type AuthToken struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Purpose   string // "email_verification" | "password_reset"
	TokenHash string
	ExpiresAt time.Time
	UsedAt    *time.Time
	RevokedAt *time.Time
	CreatedAt time.Time
}
