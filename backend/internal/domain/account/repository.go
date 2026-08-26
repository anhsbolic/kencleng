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

	// InsertRefreshToken inserts a new session refresh_token row.
	// TokenHash is already a SHA-256 hex digest (no encryption needed).
	// Called within tx. Rotation/reuse columns stay nil at issue time.
	InsertRefreshToken(ctx context.Context, tx pgx.Tx, token *RefreshToken) error

	// InsertLoginAttempt appends one credential-verification outcome to
	// login_attempts (Fitur 2C). attempt.UserID may be nil when identity
	// was never established (wrong-email attempts) — the column stays
	// NULL. Called within tx so the attempt row commits atomically with
	// the caller's surrounding login bookkeeping. Rows are append-only;
	// nothing ever updates them.
	InsertLoginAttempt(ctx context.Context, tx pgx.Tx, attempt *LoginAttempt) error

	// CountRecentFailedAttemptsByIdentifier returns how many failed login
	// attempts exist for identifierHash with the given stage ("password"
	// or "mfa") whose attempted_at is strictly after since. This is the
	// password-stage lockout hot query — identity is not reliably known
	// at check time, so the count is keyed by identifier_hash. Callers
	// derive `since` from their own clock seam (deterministic tests); the
	// adapter never calls time.Now() here.
	CountRecentFailedAttemptsByIdentifier(ctx context.Context, identifierHash, stage string, since time.Time) (int, error)

	// CountRecentFailedAttemptsByUser is the MFA-stage counterpart of
	// CountRecentFailedAttemptsByIdentifier: same threshold semantics,
	// but keyed by userID because identity was already established via a
	// validated mfa_pending_token by the time this query runs.
	CountRecentFailedAttemptsByUser(ctx context.Context, userID uuid.UUID, stage string, since time.Time) (int, error)

	// FindRefreshTokenByHash looks up a refresh token by its SHA-256 hex
	// digest. Returns ok=false when no row matches (the plain token was
	// never issued by this deployment, or garbage input).
	FindRefreshTokenByHash(ctx context.Context, tokenHash string) (token *RefreshToken, ok bool, err error)

	// RotateRefreshToken implements INV-account-03's exactly-once rotation
	// inside ONE transaction on the caller's tx:
	//
	//  1. Guarded parent mark: UPDATE refresh_tokens SET replaced_by_id =
	//     child.ID WHERE token_hash = oldTokenHash AND replaced_by_id IS
	//     NULL AND revoked_at IS NULL AND expires_at > now(). The 3-clause
	//     guard makes this the only writer of replaced_by_id and the sole
	//     arbiter of concurrent refresh races — exactly one caller can win.
	//  2. On success (RETURNING user_id, family_id from the parent), the
	//     child row is inserted into the SAME transaction with the parent's
	//     user_id/family_id. A child-insert failure rolls back the whole tx,
	//     leaving no parent-marked-without-child state (which would brick
	//     the family via reuse detection on the client's next refresh).
	//
	// The caller pre-sets child.ID, child.TokenHash (SHA-256 of the new
	// plain token), child.ExpiresAt, and child.CreatedAt; UserID and
	// FamilyID are populated from the parent's RETURNING values. Returns
	// rotated=false (transaction untouched — caller decides whether to
	// roll back) when the guard matched zero rows: not found, already
	// rotated, revoked, or expired. Per spec Assumption D, all four cases
	// are treated identically downstream (reuse ⇒ family revocation).
	RotateRefreshToken(ctx context.Context, tx pgx.Tx, oldTokenHash string, child *RefreshToken) (rotated bool, err error)

	// RevokeRefreshTokenByHash sets revoked_at = now() for the single row
	// matching tokenHash if it is not already revoked (logout path). The
	// revoked_at IS NULL guard keeps repeated logouts idempotent.
	RevokeRefreshTokenByHash(ctx context.Context, tx pgx.Tx, tokenHash string) error

	// RevokeRefreshTokenFamily sets revoked_at = now() for EVERY row in
	// the family that is not yet revoked — deliberately including rows
	// whose replaced_by_id is already set, per INV-account-04: reuse of a
	// rotated-out token means the whole lineage is compromised, and a
	// token that was legitimately rotated further down the chain must die
	// with it. Called on reuse detection (and on the race-loser branch,
	// which spec Assumption D defines as equivalent).
	RevokeRefreshTokenFamily(ctx context.Context, tx pgx.Tx, familyID uuid.UUID) error

	// FindIdentifierHashByUserAndProvider returns the identifier_hash of
	// the single identity matching (userID, providerType). Used by the MFA
	// login step to backfill login_attempts.identifier_hash from the user's
	// own email_password identity (spec Assumption C — schema consistency
	// only; the MFA-stage lockout query keys on user_id). Returns found=false
	// when no such identity exists.
	FindIdentifierHashByUserAndProvider(ctx context.Context, userID uuid.UUID, providerType string) (identifierHash string, found bool, err error)

	// GetLoginUserView assembles the LoginResponse.user read model for
	// userID (techplan §8): profile fields from users (with primary_email
	// DECRYPTED — the one decrypt-on-read path in this repository),
	// EmailVerified from any verified email_password auth_identity,
	// AuthProviders from the distinct provider_types across identities,
	// Roles from user_roles (empty until account task #8 ships assignment),
	// MFAEnabled from an enabled mfa_totp_secrets row. Returns (nil, nil)
	// if the user does not exist.
	GetLoginUserView(ctx context.Context, userID uuid.UUID) (*LoginUserView, error)

	// InsertUserLog appends an audit entry to user_logs. Called within
	// tx so the audit write is atomic with the action it records (e.g.
	// attaching a Google AuthIdentity on link intent). Append-only by
	// convention; DB-level enforcement arrives with task #08.
	InsertUserLog(ctx context.Context, tx pgx.Tx, entry *UserLog) error
}
