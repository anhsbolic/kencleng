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

// Compile-time assertion that RepositoryDB implements Repository.
var _ Repository = (*RepositoryDB)(nil)
