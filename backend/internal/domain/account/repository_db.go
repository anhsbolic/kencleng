package account

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres" // register postgres dialect
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/anhsbolic/kencleng/backend/internal/platform/crypto"
)

// pgDialect is the goqu dialect wrapper for PostgreSQL. The blank import
// of the postgres dialect package registers the dialect at init time, so
// this package-level var is safe (it only stores the dialect name; the
// dialect options are looked up at ToSQL time).
var pgDialect = goqu.Dialect("postgres")

// RepositoryDB is the goqu + pgx adapter for the account Repository
// port. It uses *pgxpool.Pool for read/standalone queries and the
// caller-supplied pgx.Tx for inserts so the service controls
// transaction boundaries. PII encryption of primary_email and
// identifier is delegated to platform/crypto at the storage boundary.
type RepositoryDB struct {
	db   *pgxpool.Pool
	keys *crypto.Keys
}

// NewRepositoryDB constructs a RepositoryDB backed by the given pool.
// keys provides the AES-GCM encryption key and HMAC key used to encrypt
// and hash PII columns on insert.
func NewRepositoryDB(db *pgxpool.Pool, keys *crypto.Keys) *RepositoryDB {
	return &RepositoryDB{db: db, keys: keys}
}

// InsertUser inserts a new User. The adapter encrypts user.PrimaryEmail
// into the primary_email BYTEA column and computes the HMAC into
// primary_email_hash, then clears user.PrimaryEmail so plaintext is not
// retained past the insert. Called within the caller's tx.
//
// On a unique violation (concurrent duplicate, R16) the wrapped
// *pgconn.PgError (code 23505) is returned up the chain so the service
// can detect it via errors.As and map it to a clean no-op — the
// repository does not swallow it.
func (r *RepositoryDB) InsertUser(ctx context.Context, tx pgx.Tx, user *User) error {
	emailCt, err := crypto.Encrypt([]byte(user.PrimaryEmail), r.keys)
	if err != nil {
		return fmt.Errorf("account: encrypt primary_email: %w", err)
	}
	emailHash := crypto.HMAC([]byte(user.PrimaryEmail), r.keys)

	sqlStr, args, err := pgDialect.Insert("users").
		Rows(goqu.Record{
			"id":                 user.ID,
			"name":               user.Name,
			"primary_email":      emailCt,
			"primary_email_hash": emailHash,
			// created_at/updated_at default to now() in the schema.
		}).
		Prepared(true).
		ToSQL()
	if err != nil {
		return fmt.Errorf("account: build insert users: %w", err)
	}
	if _, err := tx.Exec(ctx, sqlStr, args...); err != nil {
		return fmt.Errorf("account: insert users: %w", err)
	}
	// Plaintext must not be retained past the insert.
	user.PrimaryEmail = ""
	return nil
}

// InsertAuthIdentity inserts a new AuthIdentity. The adapter encrypts
// identity.Identifier and computes the HMAC, then clears
// identity.Identifier. Called within tx.
//
// On a unique violation (concurrent duplicate registration, R16) the
// wrapped *pgconn.PgError (code 23505) is returned up the chain so the
// service can detect it via errors.As and map it to a clean no-op
// rollback — the repository does not swallow it. The caller-supplied tx
// rolls back, leaving no orphaned users row.
func (r *RepositoryDB) InsertAuthIdentity(ctx context.Context, tx pgx.Tx, identity *AuthIdentity) error {
	identifierCt, err := crypto.Encrypt([]byte(identity.Identifier), r.keys)
	if err != nil {
		return fmt.Errorf("account: encrypt identifier: %w", err)
	}
	identifierHash := crypto.HMAC([]byte(identity.Identifier), r.keys)

	rec := goqu.Record{
		"id":              identity.ID,
		"user_id":         identity.UserID,
		"provider_type":   identity.ProviderType,
		"identifier":      identifierCt,
		"identifier_hash": identifierHash,
	}
	// credential_secret is nullable; include it only when set to avoid
	// the nil-pointer-in-interface pitfall (nil *string assigned to an
	// interface{} is non-nil). A new google identity has no credential.
	if identity.CredentialSecret != nil {
		rec["credential_secret"] = *identity.CredentialSecret
	}
	// verified_at: a new identity is unverified (NULL). Omit so the DB
	// default applies; the service sets verified_at later via
	// SetVerifiedAt after a successful token redeem.
	// created_at/updated_at default to now() in the schema.

	sqlStr, args, err := pgDialect.Insert("auth_identities").
		Rows(rec).
		Prepared(true).
		ToSQL()
	if err != nil {
		return fmt.Errorf("account: build insert auth_identities: %w", err)
	}
	if _, err := tx.Exec(ctx, sqlStr, args...); err != nil {
		return fmt.Errorf("account: insert auth_identities: %w", err)
	}
	// Plaintext must not be retained past the insert.
	identity.Identifier = ""
	return nil
}

// InsertAuthToken inserts a new single-use token. Tokens are not
// encrypted (token_hash is already a SHA-256 hex digest), so this
// method does not depend on the crypto prerequisite. Called within tx.
func (r *RepositoryDB) InsertAuthToken(ctx context.Context, tx pgx.Tx, token *AuthToken) error {
	sqlStr, args, err := pgDialect.Insert("auth_tokens").
		Rows(goqu.Record{
			"id":         token.ID,
			"user_id":    token.UserID,
			"purpose":    token.Purpose,
			"token_hash": token.TokenHash,
			"expires_at": token.ExpiresAt,
			"created_at": token.CreatedAt,
			// used_at and revoked_at default to NULL for a new token.
		}).
		Prepared(true).
		ToSQL()
	if err != nil {
		return fmt.Errorf("account: build insert auth_tokens: %w", err)
	}
	if _, err := tx.Exec(ctx, sqlStr, args...); err != nil {
		return fmt.Errorf("account: insert auth_tokens: %w", err)
	}
	return nil
}

// InsertRefreshToken inserts a new session refresh_token row. Same shape
// as InsertAuthToken: token_hash is already a hash, no encryption, called
// within the caller's tx. Rotation/reuse columns stay nil at issue time.
func (r *RepositoryDB) InsertRefreshToken(ctx context.Context, tx pgx.Tx, token *RefreshToken) error {
	sqlStr, args, err := pgDialect.Insert("refresh_tokens").
		Rows(goqu.Record{
			"id":         token.ID,
			"user_id":    token.UserID,
			"family_id":  token.FamilyID,
			"token_hash": token.TokenHash,
			"expires_at": token.ExpiresAt,
			"created_at": token.CreatedAt,
			// revoked_at and replaced_by_id default to NULL for a new token.
		}).
		Prepared(true).
		ToSQL()
	if err != nil {
		return fmt.Errorf("account: build insert refresh_tokens: %w", err)
	}
	if _, err := tx.Exec(ctx, sqlStr, args...); err != nil {
		return fmt.Errorf("account: insert refresh_tokens: %w", err)
	}
	return nil
}

// InsertUserLog appends an audit entry to user_logs. Called within the
// caller's tx so the audit write commits atomically with the action it
// records.
func (r *RepositoryDB) InsertUserLog(ctx context.Context, tx pgx.Tx, entry *UserLog) error {
	sqlStr, args, err := pgDialect.Insert("user_logs").
		Rows(goqu.Record{
			"id":          entry.ID,
			"user_id":     entry.UserID,
			"action_type": entry.ActionType,
			"created_at":  entry.CreatedAt,
		}).
		Prepared(true).
		ToSQL()
	if err != nil {
		return fmt.Errorf("account: build insert user_logs: %w", err)
	}
	if _, err := tx.Exec(ctx, sqlStr, args...); err != nil {
		return fmt.Errorf("account: insert user_logs: %w", err)
	}
	return nil
}

// FindAuthIdentityByIdentifierHash looks up an identity by
// (providerType, identifierHash). Returns (nil, nil) if not found.
// The returned identity has its non-encrypted fields populated;
// Identifier is left empty (decryption is not needed for the current
// flows, which look up by hash).
func (r *RepositoryDB) FindAuthIdentityByIdentifierHash(ctx context.Context, providerType, identifierHash string) (*AuthIdentity, error) {
	sqlStr, args, err := pgDialect.From("auth_identities").
		Select("id", "user_id", "provider_type", "credential_secret", "verified_at", "created_at", "updated_at").
		Where(goqu.Ex{"provider_type": providerType, "identifier_hash": identifierHash}).
		Prepared(true).
		ToSQL()
	if err != nil {
		return nil, fmt.Errorf("account: build select auth_identities: %w", err)
	}

	var (
		identity   AuthIdentity
		credSecret sql.NullString
		verifiedAt sql.NullTime
	)
	if err := r.db.QueryRow(ctx, sqlStr, args...).Scan(
		&identity.ID,
		&identity.UserID,
		&identity.ProviderType,
		&credSecret,
		&verifiedAt,
		&identity.CreatedAt,
		&identity.UpdatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("account: select auth_identities: %w", err)
	}
	if credSecret.Valid {
		s := credSecret.String
		identity.CredentialSecret = &s
	}
	if verifiedAt.Valid {
		t := verifiedAt.Time
		identity.VerifiedAt = &t
	}
	return &identity, nil
}

// FindAuthTokenByHash looks up a token by its SHA-256 hash.
// Returns (nil, nil) if not found.
func (r *RepositoryDB) FindAuthTokenByHash(ctx context.Context, tokenHash string) (*AuthToken, error) {
	sqlStr, args, err := pgDialect.From("auth_tokens").
		Select("id", "user_id", "purpose", "token_hash", "expires_at", "used_at", "revoked_at", "created_at").
		Where(goqu.Ex{"token_hash": tokenHash}).
		Prepared(true).
		ToSQL()
	if err != nil {
		return nil, fmt.Errorf("account: build select auth_tokens: %w", err)
	}

	var (
		token     AuthToken
		usedAt    sql.NullTime
		revokedAt sql.NullTime
	)
	if err := r.db.QueryRow(ctx, sqlStr, args...).Scan(
		&token.ID,
		&token.UserID,
		&token.Purpose,
		&token.TokenHash,
		&token.ExpiresAt,
		&usedAt,
		&revokedAt,
		&token.CreatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("account: select auth_tokens: %w", err)
	}
	if usedAt.Valid {
		t := usedAt.Time
		token.UsedAt = &t
	}
	if revokedAt.Valid {
		t := revokedAt.Time
		token.RevokedAt = &t
	}
	return &token, nil
}

// RedeemToken atomically marks a token used iff it is currently valid:
// used_at IS NULL AND revoked_at IS NULL AND expires_at > now() (full
// 3-clause predicate per INV-account-08 Statement). On success (exactly
// 1 row affected) it returns the token's user_id and purpose via
// RETURNING — no second round-trip is needed. Returns ok=false if 0 rows
// were affected (not-found / already-used / revoked / expired). Single-
// use correctness is this atomic UPDATE ... WHERE — no application-level
// locking.
//
// Runs inside the caller's tx so the subsequent SetUserVerified is in
// the same transaction (S2 fix: redeem + set-verified are atomic).
func (r *RepositoryDB) RedeemToken(ctx context.Context, tx pgx.Tx, tokenHash string) (uuid.UUID, string, bool, error) {
	sqlStr, args, err := pgDialect.Update("auth_tokens").
		Set(goqu.Record{"used_at": time.Now()}).
		Where(
			goqu.Ex{"token_hash": tokenHash},
			goqu.L("used_at IS NULL"),
			goqu.L("revoked_at IS NULL"),
			goqu.L("expires_at > now()"),
		).
		Returning("user_id", "purpose").
		Prepared(true).
		ToSQL()
	if err != nil {
		return uuid.Nil, "", false, fmt.Errorf("account: build redeem auth_tokens: %w", err)
	}
	var (
		userID  uuid.UUID
		purpose string
	)
	if err := tx.QueryRow(ctx, sqlStr, args...).Scan(&userID, &purpose); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return uuid.Nil, "", false, nil // 0 rows affected — not redeemed
		}
		return uuid.Nil, "", false, fmt.Errorf("account: redeem auth_tokens: %w", err)
	}
	return userID, purpose, true, nil
}

// SetUserVerified sets auth_identities.verified_at = verifiedAt for the
// single identity matching (userID, providerType). Called by VerifyEmail
// after a successful RedeemToken. The token carries user_id, so
// verification is keyed on (user_id, provider_type).
//
// Runs inside the caller's tx so it is atomic with RedeemToken (S2 fix).
func (r *RepositoryDB) SetUserVerified(ctx context.Context, tx pgx.Tx, userID uuid.UUID, providerType string, verifiedAt time.Time) error {
	sqlStr, args, err := pgDialect.Update("auth_identities").
		Set(goqu.Record{"verified_at": verifiedAt}).
		Where(goqu.Ex{"user_id": userID, "provider_type": providerType}).
		Prepared(true).
		ToSQL()
	if err != nil {
		return fmt.Errorf("account: build set verified_at: %w", err)
	}
	if _, err := tx.Exec(ctx, sqlStr, args...); err != nil {
		return fmt.Errorf("account: set verified_at: %w", err)
	}
	return nil
}

// RevokeTokens sets revoked_at = now() for all unused, unrevoked tokens
// of (userID, purpose). Called by resend before issuing a new token
// (R13). Uses the caller's tx so revoke + insert of the new token are
// atomic in a single transaction. Already-used tokens keep their
// used_at; already-revoked tokens are not touched (the revoked_at IS
// NULL guard).
func (r *RepositoryDB) RevokeTokens(ctx context.Context, tx pgx.Tx, userID uuid.UUID, purpose string) error {
	sqlStr, args, err := pgDialect.Update("auth_tokens").
		Set(goqu.Record{"revoked_at": time.Now()}).
		Where(
			goqu.Ex{"user_id": userID, "purpose": purpose},
			goqu.L("used_at IS NULL"),
			goqu.L("revoked_at IS NULL"),
		).
		Prepared(true).
		ToSQL()
	if err != nil {
		return fmt.Errorf("account: build revoke auth_tokens: %w", err)
	}
	if _, err := tx.Exec(ctx, sqlStr, args...); err != nil {
		return fmt.Errorf("account: revoke auth_tokens: %w", err)
	}
	return nil
}

// InsertLoginAttempt appends one credential-verification outcome to
// login_attempts. attempt.UserID may be nil (wrong-email attempts) — the
// column is omitted so the DB NULL default applies, avoiding the
// nil-pointer-in-interface pitfall noted in InsertAuthIdentity.
// attempt.AttemptedAt is written explicitly so tests can seed historical
// rows for window-boundary assertions. Called within tx.
func (r *RepositoryDB) InsertLoginAttempt(ctx context.Context, tx pgx.Tx, attempt *LoginAttempt) error {
	rec := goqu.Record{
		"id":              attempt.ID,
		"identifier_hash": attempt.IdentifierHash,
		"stage":           attempt.Stage,
		"success":         attempt.Success,
		"attempted_at":    attempt.AttemptedAt,
	}
	if attempt.UserID != nil {
		rec["user_id"] = *attempt.UserID
	}

	sqlStr, args, err := pgDialect.Insert("login_attempts").
		Rows(rec).
		Prepared(true).
		ToSQL()
	if err != nil {
		return fmt.Errorf("account: build insert login_attempts: %w", err)
	}
	if _, err := tx.Exec(ctx, sqlStr, args...); err != nil {
		return fmt.Errorf("account: insert login_attempts: %w", err)
	}
	return nil
}

// CountRecentFailedAttemptsByIdentifier returns the number of failed login
// attempts for identifierHash at the given stage with attempted_at strictly
// after since. The cutoff is a bound parameter computed by the caller from
// its own clock — no time.Now() hides in this adapter, so lockout windows
// are deterministically testable.
func (r *RepositoryDB) CountRecentFailedAttemptsByIdentifier(ctx context.Context, identifierHash, stage string, since time.Time) (int, error) {
	sqlStr, args, err := pgDialect.From("login_attempts").
		Select(goqu.COUNT("*")).
		Where(
			goqu.Ex{"identifier_hash": identifierHash, "stage": stage, "success": false},
			goqu.L("attempted_at > ?", since),
		).
		Prepared(true).
		ToSQL()
	if err != nil {
		return 0, fmt.Errorf("account: build count login_attempts by identifier: %w", err)
	}
	var n int
	if err := r.db.QueryRow(ctx, sqlStr, args...).Scan(&n); err != nil {
		return 0, fmt.Errorf("account: count login_attempts by identifier: %w", err)
	}
	return n, nil
}

// CountRecentFailedAttemptsByUser is the MFA-stage counterpart keyed on
// user_id + stage; see CountRecentFailedAttemptsByIdentifier for the
// cutoff semantics.
func (r *RepositoryDB) CountRecentFailedAttemptsByUser(ctx context.Context, userID uuid.UUID, stage string, since time.Time) (int, error) {
	sqlStr, args, err := pgDialect.From("login_attempts").
		Select(goqu.COUNT("*")).
		Where(
			goqu.Ex{"user_id": userID, "stage": stage, "success": false},
			goqu.L("attempted_at > ?", since),
		).
		Prepared(true).
		ToSQL()
	if err != nil {
		return 0, fmt.Errorf("account: build count login_attempts by user: %w", err)
	}
	var n int
	if err := r.db.QueryRow(ctx, sqlStr, args...).Scan(&n); err != nil {
		return 0, fmt.Errorf("account: count login_attempts by user: %w", err)
	}
	return n, nil
}

// FindRefreshTokenByHash looks up a refresh token row by its SHA-256 hex
// digest. Returns ok=false when no row matches. All columns are populated,
// including nullable rotation state (RevokedAt/ReplacedByID), so callers
// can run the full reuse-detection classification without a second query.
func (r *RepositoryDB) FindRefreshTokenByHash(ctx context.Context, tokenHash string) (*RefreshToken, bool, error) {
	sqlStr, args, err := pgDialect.From("refresh_tokens").
		Select("id", "user_id", "family_id", "token_hash", "expires_at", "revoked_at", "replaced_by_id", "created_at").
		Where(goqu.Ex{"token_hash": tokenHash}).
		Prepared(true).
		ToSQL()
	if err != nil {
		return nil, false, fmt.Errorf("account: build select refresh_tokens: %w", err)
	}

	var (
		token        RefreshToken
		revokedAt    sql.NullTime
		replacedByID sql.NullString
	)
	if err := r.db.QueryRow(ctx, sqlStr, args...).Scan(
		&token.ID,
		&token.UserID,
		&token.FamilyID,
		&token.TokenHash,
		&token.ExpiresAt,
		&revokedAt,
		&replacedByID,
		&token.CreatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, false, nil
		}
		return nil, false, fmt.Errorf("account: select refresh_tokens: %w", err)
	}
	if revokedAt.Valid {
		t := revokedAt.Time
		token.RevokedAt = &t
	}
	if replacedByID.Valid {
		id, parseErr := uuid.Parse(replacedByID.String)
		if parseErr != nil {
			return nil, false, fmt.Errorf("account: parse refresh_tokens.replaced_by_id: %w", parseErr)
		}
		token.ReplacedByID = &id
	}
	return &token, true, nil
}

// RotateRefreshToken implements INV-account-03's exactly-once rotation in
// ONE transaction on the caller's tx: guarded parent mark → child insert.
// See Repository.RotateRefreshToken for the full contract. On rotated=false
// nothing has been written and the tx is left untouched — the caller rolls
// back and classifies (not-found / already-rotated / revoked / expired are
// deliberately indistinguishable per spec Assumption D).
func (r *RepositoryDB) RotateRefreshToken(ctx context.Context, tx pgx.Tx, oldTokenHash string, child *RefreshToken) (bool, error) {
	markSQL, markArgs, err := pgDialect.Update("refresh_tokens").
		Set(goqu.Record{"replaced_by_id": child.ID}).
		Where(
			goqu.Ex{"token_hash": oldTokenHash},
			goqu.L("replaced_by_id IS NULL"),
			goqu.L("revoked_at IS NULL"),
			goqu.L("expires_at > now()"), // DB clock, same convention as RedeemToken
		).
		Returning("user_id", "family_id").
		Prepared(true).
		ToSQL()
	if err != nil {
		return false, fmt.Errorf("account: build rotate refresh_tokens mark: %w", err)
	}

	var parentUserID, familyID uuid.UUID
	if err := tx.QueryRow(ctx, markSQL, markArgs...).Scan(&parentUserID, &familyID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			// Guard matched zero rows: not found / already rotated /
			// revoked / expired. Nothing was modified in the tx.
			return false, nil
		}
		return false, fmt.Errorf("account: rotate refresh_tokens mark: %w", err)
	}

	child.UserID = parentUserID
	child.FamilyID = familyID

	insertSQL, insertArgs, err := pgDialect.Insert("refresh_tokens").
		Rows(goqu.Record{
			"id":         child.ID,
			"user_id":    child.UserID,
			"family_id":  child.FamilyID,
			"token_hash": child.TokenHash,
			"expires_at": child.ExpiresAt,
			"created_at": child.CreatedAt,
			// revoked_at / replaced_by_id default to NULL.
		}).
		Prepared(true).
		ToSQL()
	if err != nil {
		return false, fmt.Errorf("account: build rotate refresh_tokens child insert: %w", err)
	}
	if _, err := tx.Exec(ctx, insertSQL, insertArgs...); err != nil {
		// The whole tx fails here — the deferred rollback in the service
		// undoes the parent mark too, leaving no marked-without-child state.
		return false, fmt.Errorf("account: insert rotated refresh_token child: %w", err)
	}
	return true, nil
}

// RevokeRefreshTokenByHash sets revoked_at = now() for the single matching
// row if it is not already revoked. Idempotent by construction: repeated
// calls match zero rows under the revoked_at IS NULL guard and change
// nothing.
func (r *RepositoryDB) RevokeRefreshTokenByHash(ctx context.Context, tx pgx.Tx, tokenHash string) error {
	sqlStr, args, err := pgDialect.Update("refresh_tokens").
		Set(goqu.Record{"revoked_at": time.Now()}).
		Where(
			goqu.Ex{"token_hash": tokenHash},
			goqu.L("revoked_at IS NULL"),
		).
		Prepared(true).
		ToSQL()
	if err != nil {
		return fmt.Errorf("account: build revoke refresh_tokens by hash: %w", err)
	}
	if _, err := tx.Exec(ctx, sqlStr, args...); err != nil {
		return fmt.Errorf("account: revoke refresh_tokens by hash: %w", err)
	}
	return nil
}

// RevokeRefreshTokenFamily sets revoked_at = now() for every unrevoked row
// sharing familyID — INCLUDING already-rotated rows (no replaced_by_id
// guard): INV-account-04 means a reused token condemns its entire lineage.
func (r *RepositoryDB) RevokeRefreshTokenFamily(ctx context.Context, tx pgx.Tx, familyID uuid.UUID) error {
	sqlStr, args, err := pgDialect.Update("refresh_tokens").
		Set(goqu.Record{"revoked_at": time.Now()}).
		Where(
			goqu.Ex{"family_id": familyID},
			goqu.L("revoked_at IS NULL"),
		).
		Prepared(true).
		ToSQL()
	if err != nil {
		return fmt.Errorf("account: build revoke refresh_tokens family: %w", err)
	}
	if _, err := tx.Exec(ctx, sqlStr, args...); err != nil {
		return fmt.Errorf("account: revoke refresh_tokens family: %w", err)
	}
	return nil
}

// UpdateIdentityCredentialSecret sets credential_secret = passwordHash for
// the single identity matching (userID, providerType). Keyed on
// (user_id, provider_type) — mirroring SetUserVerified — because the reset
// token carries user_id and INV-account-01 guarantees at most one identity
// per (user_id, provider_type).
//
// Runs inside the caller's tx so it is atomic with RedeemToken and
// RevokeAllRefreshTokensForUser (INV-account-05): a failure here rolls back
// the redeem, leaving the token unconsumed (spec Assumption B).
func (r *RepositoryDB) UpdateIdentityCredentialSecret(ctx context.Context, tx pgx.Tx, userID uuid.UUID, providerType string, passwordHash string) error {
	sqlStr, args, err := pgDialect.Update("auth_identities").
		Set(goqu.Record{"credential_secret": passwordHash}).
		Where(goqu.Ex{"user_id": userID, "provider_type": providerType}).
		Prepared(true).
		ToSQL()
	if err != nil {
		return fmt.Errorf("account: build update credential_secret: %w", err)
	}
	if _, err := tx.Exec(ctx, sqlStr, args...); err != nil {
		return fmt.Errorf("account: update credential_secret: %w", err)
	}
	return nil
}

// RevokeAllRefreshTokensForUser sets revoked_at = now() for EVERY
// refresh_tokens row matching userID that is not already revoked — all
// families, including already-rotated rows (no replaced_by_id guard),
// because INV-account-05 scopes the mass revoke to the user. The
// revoked_at IS NULL guard keeps repeat calls idempotent.
//
// Runs inside the caller's tx so it commits or rolls back together with
// RedeemToken and UpdateIdentityCredentialSecret — never as a separate
// best-effort step after the credential update.
func (r *RepositoryDB) RevokeAllRefreshTokensForUser(ctx context.Context, tx pgx.Tx, userID uuid.UUID) error {
	sqlStr, args, err := pgDialect.Update("refresh_tokens").
		Set(goqu.Record{"revoked_at": time.Now()}).
		Where(
			goqu.Ex{"user_id": userID},
			goqu.L("revoked_at IS NULL"),
		).
		Prepared(true).
		ToSQL()
	if err != nil {
		return fmt.Errorf("account: build revoke refresh_tokens by user: %w", err)
	}
	if _, err := tx.Exec(ctx, sqlStr, args...); err != nil {
		return fmt.Errorf("account: revoke refresh_tokens by user: %w", err)
	}
	return nil
}

// FindIdentifierHashByUserAndProvider returns the identifier_hash of the
// single identity matching (userID, providerType); found=false when absent.
// See Repository.FindIdentifierHashByUserAndProvider.
func (r *RepositoryDB) FindIdentifierHashByUserAndProvider(ctx context.Context, userID uuid.UUID, providerType string) (string, bool, error) {
	sqlStr, args, err := pgDialect.From("auth_identities").
		Select("identifier_hash").
		Where(goqu.Ex{"user_id": userID, "provider_type": providerType}).
		Prepared(true).
		ToSQL()
	if err != nil {
		return "", false, fmt.Errorf("account: build select identity hash by user: %w", err)
	}
	var hash string
	if err := r.db.QueryRow(ctx, sqlStr, args...).Scan(&hash); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", false, nil
		}
		return "", false, fmt.Errorf("account: select identity hash by user: %w", err)
	}
	return hash, true, nil
}

// GetLoginUserView assembles the LoginResponse.user read model. It runs as
// four simple prepared queries rather than one wide join: each sub-answer
// (identity aggregation, roles, mfa flag) is independently indexable and
// keeps the decrypt-on-read step isolated to the profile row. primary_email
// is decrypted here — the plaintext lives only in the returned view and is
// never logged (R19 discipline is enforced upstream, but the view doc also
// warns). Returns (nil, nil) when the user does not exist.
func (r *RepositoryDB) GetLoginUserView(ctx context.Context, userID uuid.UUID) (*LoginUserView, error) {
	view := &LoginUserView{
		ID:            userID,
		Roles:         []string{},
		AuthProviders: []string{},
	}

	// 1. Profile row: name + ciphertext email + created_at.
	profileSQL, profileArgs, err := pgDialect.From("users").
		Select("name", "primary_email", "created_at").
		Where(goqu.Ex{"id": userID}).
		Prepared(true).
		ToSQL()
	if err != nil {
		return nil, fmt.Errorf("account: build select users for login view: %w", err)
	}
	var emailCiphertext []byte
	if err := r.db.QueryRow(ctx, profileSQL, profileArgs...).Scan(
		&view.Name, &emailCiphertext, &view.CreatedAt,
	); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("account: select users for login view: %w", err)
	}
	plaintext, err := crypto.Decrypt(emailCiphertext, r.keys)
	if err != nil {
		return nil, fmt.Errorf("account: decrypt primary_email for login view: %w", err)
	}
	view.Email = string(plaintext)

	// 2. Identity aggregation: distinct providers + verified flag.
	identSQL, identArgs, err := pgDialect.From("auth_identities").
		Select("provider_type", "verified_at").
		Where(goqu.Ex{"user_id": userID}).
		Prepared(true).
		ToSQL()
	if err != nil {
		return nil, fmt.Errorf("account: build select auth_identities for login view: %w", err)
	}
	rows, err := r.db.Query(ctx, identSQL, identArgs...)
	if err != nil {
		return nil, fmt.Errorf("account: select auth_identities for login view: %w", err)
	}
	defer rows.Close()
	seenProviders := make(map[string]bool)
	for rows.Next() {
		var providerType string
		var verifiedAt sql.NullTime
		if err := rows.Scan(&providerType, &verifiedAt); err != nil {
			return nil, fmt.Errorf("account: scan auth_identities for login view: %w", err)
		}
		if !seenProviders[providerType] {
			seenProviders[providerType] = true
			view.AuthProviders = append(view.AuthProviders, providerType)
		}
		if providerType == "email_password" && verifiedAt.Valid {
			view.EmailVerified = true
		}
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("account: iterate auth_identities for login view: %w", err)
	}

	// 3. Roles (empty until account task #8's assignment API writes rows).
	rolesSQL, rolesArgs, err := pgDialect.From("user_roles").
		Select("role").
		Where(goqu.Ex{"user_id": userID}).
		Order(goqu.I("role").Asc()).
		Prepared(true).
		ToSQL()
	if err != nil {
		return nil, fmt.Errorf("account: build select user_roles for login view: %w", err)
	}
	roleRows, err := r.db.Query(ctx, rolesSQL, rolesArgs...)
	if err != nil {
		return nil, fmt.Errorf("account: select user_roles for login view: %w", err)
	}
	defer roleRows.Close()
	for roleRows.Next() {
		var role string
		if err := roleRows.Scan(&role); err != nil {
			return nil, fmt.Errorf("account: scan user_roles for login view: %w", err)
		}
		view.Roles = append(view.Roles, role)
	}
	if err := roleRows.Err(); err != nil {
		return nil, fmt.Errorf("account: iterate user_roles for login view: %w", err)
	}

	// 4. MFA enabled flag: an mfa_totp_secrets row with enabled_at set
	// (INV-account-07 guarantees enabled_at is only set after verified
	// enrollment; the login branch keys off exactly this predicate).
	mfaSQL, mfaArgs, err := pgDialect.From("mfa_totp_secrets").
		Select(goqu.COUNT("*")).
		Where(goqu.Ex{"user_id": userID}, goqu.L("enabled_at IS NOT NULL")).
		Prepared(true).
		ToSQL()
	if err != nil {
		return nil, fmt.Errorf("account: build select mfa_totp_secrets for login view: %w", err)
	}
	var enabledCount int
	if err := r.db.QueryRow(ctx, mfaSQL, mfaArgs...).Scan(&enabledCount); err != nil {
		return nil, fmt.Errorf("account: select mfa_totp_secrets for login view: %w", err)
	}
	view.MFAEnabled = enabledCount > 0

	return view, nil
}

// Compile-time assertion that RepositoryDB implements Repository.
var _ Repository = (*RepositoryDB)(nil)
