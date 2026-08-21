package account

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// Repository is the persistence port for the account domain.
// Implementations must parameterize all SQL via goqu — never string
// concatenation (AGENTS.md golden rule). PII columns (primary_email,
// identifier) are encrypted at insert time by the adapter; *_hash
// columns hold the HMAC for lookup.
//
// Insert methods take a caller-supplied pgx.Tx so the service can wrap
// multi-row writes in a single transaction (R16: concurrent duplicate
// registration must roll back cleanly — no orphaned users row).
type Repository interface {
	// InsertUser inserts a new User. The adapter encrypts
	// user.PrimaryEmail into the primary_email column and computes the
	// HMAC into primary_email_hash, then clears user.PrimaryEmail so
	// plaintext is not retained past the insert. Called within the
	// caller's tx.
	InsertUser(ctx context.Context, tx pgx.Tx, user *User) error

	// InsertAuthIdentity inserts a new AuthIdentity. The adapter
	// encrypts identity.Identifier and computes the HMAC, then clears
	// identity.Identifier. Called within tx.
	InsertAuthIdentity(ctx context.Context, tx pgx.Tx, identity *AuthIdentity) error

	// InsertAuthToken inserts a new single-use token. Called within tx.
	InsertAuthToken(ctx context.Context, tx pgx.Tx, token *AuthToken) error

	// FindAuthIdentityByIdentifierHash looks up an identity by
	// (providerType, identifierHash). Returns (nil, nil) if not found.
	// The returned identity has its non-encrypted fields populated;
	// Identifier is left empty (decryption is not needed for the
	// current flows, which look up by hash).
	FindAuthIdentityByIdentifierHash(ctx context.Context, providerType, identifierHash string) (*AuthIdentity, error)

	// FindAuthTokenByHash looks up a token by its SHA-256 hash.
	// Returns (nil, nil) if not found.
	FindAuthTokenByHash(ctx context.Context, tokenHash string) (*AuthToken, error)

	// RedeemToken atomically marks a token used iff it is currently
	// valid: used_at IS NULL AND revoked_at IS NULL AND expires_at > now()
	// (full 3-clause predicate per INV-account-08 Statement). On success
	// (exactly 1 row affected) it returns the token's user_id and purpose
	// via RETURNING — no second round-trip is needed. Returns ok=false if
	// 0 rows were affected (not-found / already-used / revoked / expired);
	// the caller disambiguates expired vs other via FindAuthTokenByHash.
	//
	// Takes the caller's tx so the subsequent SetUserVerified can run in
	// the same transaction (S2 fix: redeem + set-verified are atomic — a
	// set-verified failure rolls back the redeem, so the token is not
	// burned without the identity being verified).
	//
	// The 3-clause guard is non-negotiable. The invariant's
	// Verification field omits the revoked_at IS NULL clause (2-clause)
	// — that is a documented spec error (techplan §14 Open Item #2).
	// Use the Statement's 3-clause version. Do not edit the spec.
	RedeemToken(ctx context.Context, tx pgx.Tx, tokenHash string) (userID uuid.UUID, purpose string, ok bool, err error)

	// SetUserVerified sets auth_identities.verified_at = verifiedAt for
	// the single identity matching (userID, providerType). Called by
	// VerifyEmail after a successful RedeemToken. The token carries
	// user_id (not identity_id), so verification is keyed on
	// (user_id, provider_type) — for the email_verification flow that is
	// (userID, "email_password").
	//
	// Takes the caller's tx so it runs in the same transaction as
	// RedeemToken (S2 fix: atomic redeem + set-verified).
	SetUserVerified(ctx context.Context, tx pgx.Tx, userID uuid.UUID, providerType string, verifiedAt time.Time) error

	// RevokeTokens sets revoked_at = now() for all unused, unrevoked
	// tokens of (userID, purpose). Called by resend before issuing a new
	// token (R13). Takes the caller's tx so revoke + insert of the new
	// token are atomic within a single transaction.
	RevokeTokens(ctx context.Context, tx pgx.Tx, userID uuid.UUID, purpose string) error
}
