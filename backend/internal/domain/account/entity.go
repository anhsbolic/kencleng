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

// RefreshToken is a long-lived session token enabling access-token renewal.
// TokenHash is the SHA-256 hex digest of the plain refresh token (the plain
// value exists only in the response/cookie path, never at rest). FamilyID
// groups tokens from one login lineage for rotation/reuse detection
// (INV-account-03/04 — rotation and reuse detection are live as of the
// login/session slice). RevokedAt and ReplacedByID start nil; ReplacedByID
// transitions NULL → child-id at most once via RotateRefreshToken's guarded
// UPDATE (a token with ReplacedByID set can no longer be rotated); RevokedAt
// is set directly on logout (RevokeRefreshTokenByHash) or wholesale across
// the whole family when reuse is detected (RevokeRefreshTokenFamily,
// INV-account-04).
type RefreshToken struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	FamilyID     uuid.UUID
	TokenHash    string
	ExpiresAt    time.Time
	RevokedAt    *time.Time
	ReplacedByID *uuid.UUID
	CreatedAt    time.Time
}

// LoginAttempt records one credential-verification outcome for the
// persistent lockout mechanism (Fitur 2C). Stage distinguishes the password
// step ("password") from the MFA step ("mfa"): both use the same threshold
// (≥5 failures in a trailing 15-minute window) but different keys —
// password-stage lockout counts by IdentifierHash because identity is not
// reliably known yet, MFA-stage lockout counts by UserID (already verified
// via a valid mfa_pending_token). UserID stays nil when identity was never
// established (wrong-email attempts); it carries the known user otherwise.
// Rows are append-only — nothing ever updates or deletes them.
type LoginAttempt struct {
	ID             uuid.UUID
	IdentifierHash string
	UserID         *uuid.UUID // nil when identity unknown at write time
	Stage          string     // "password" | "mfa"
	Success        bool
	AttemptedAt    time.Time
}

// LoginUserView is the read model assembled at login time to populate
// LoginResponse.user (openapi components.schemas.User). Email is decrypted
// plaintext — this is the resource owner's own profile view, the one place
// the repository adapter decrypts primary_email on read (every other flow
// looks up by *_hash and never needs plaintext back). Roles come from
// user_roles (empty until account task #8 ships its assignment API);
// AuthProviders aggregates the distinct provider_types across the user's
// auth_identities rows; MFAEnabled reflects an enabled mfa_totp_secrets row.
type LoginUserView struct {
	ID            uuid.UUID
	Name          string
	Email         string // decrypted plaintext; never logged
	EmailVerified bool   // any email_password identity has verified_at set
	Roles         []string
	AuthProviders []string
	MFAEnabled    bool
	CreatedAt     time.Time
}

// UserLog is an append-only audit entry for account events. ActionType is a
// package constant (e.g. "account_linking"); the DB-level immutability
// constraint (REVOKE UPDATE/DELETE, INV-account-11) and the full action_type
// vocabulary are owned by the user-logs task (#08).
type UserLog struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	ActionType string
	CreatedAt  time.Time
}
