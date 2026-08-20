//go:build integration

// package account — integration tests against a real Postgres.
//
// These tests are opt-in via the `integration` build tag so they are
// excluded from the fast `go test ./...` run. Run them manually:
//
//	migrate -path migrations -database "$DATABASE_URL" up   # apply schema (manual)
//	go test -tags=integration -race ./internal/domain/account/...
//
// The tests assume the schema from migrations/000001..000003 is applied
// against the database named by DATABASE_URL. Each test isolates itself
// with fresh random UUIDs and cleans up its own rows via t.Cleanup
// (deleting the user cascades to auth_identities and auth_tokens via the
// FK ON DELETE CASCADE), so the dev DB is not polluted.
package account

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/anhsbolic/kencleng/backend/internal/platform/crypto"
	"github.com/anhsbolic/kencleng/backend/internal/platform/db"
)

// testKeys returns a fixed, non-secret 32-byte key pair for integration
// tests. The dev DB is not production; these keys are for deterministic
// test encryption only.
func integrationTestKeys(t *testing.T) *crypto.Keys {
	t.Helper()
	return &crypto.Keys{
		EncryptionKey: make([]byte, 32), // 32 zero bytes — valid AES-256 key for tests
		HMACKey:       make([]byte, 32),
	}
}

// integrationEnv returns a connected RepositoryDB and its underlying
// pool, or skips the test if DATABASE_URL is not set. The schema must
// already be applied (run migrations manually beforehand).
func integrationEnv(t *testing.T) (*RepositoryDB, *pgxpool.Pool) {
	t.Helper()
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		t.Skip("DATABASE_URL not set; skipping integration test")
	}
	ctx := context.Background()
	pool, err := db.Open(ctx, databaseURL)
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	t.Cleanup(pool.Close)
	return NewRepositoryDB(pool, integrationTestKeys(t)), pool
}

// newTestUser builds a fresh User with a unique email (random UUID) so
// concurrent tests never collide. Cleanup deletes the user row, which
// cascades to its identities and tokens. InsertUser clears
// u.PrimaryEmail after insert (plaintext not retained); callers needing
// the email should capture it before calling.
func newTestUser(t *testing.T, repo *RepositoryDB, pool *pgxpool.Pool) *User {
	t.Helper()
	ctx := context.Background()
	u := &User{
		ID:           uuid.New(),
		Name:         "Integration Tester",
		PrimaryEmail: fmt.Sprintf("test-%s@example.com", uuid.NewString()),
	}
	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		t.Fatalf("begin tx: %v", err)
	}
	if err := repo.InsertUser(ctx, tx, u); err != nil {
		t.Fatalf("insert user: %v", err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatalf("commit: %v", err)
	}
	t.Cleanup(func() {
		_, _ = pool.Exec(ctx, "DELETE FROM users WHERE id = $1", u.ID)
	})
	return u
}

// TestInsertUserAndIdentity_InTransaction verifies a User + AuthIdentity
// + AuthToken can be inserted in a single transaction and the rows are
// present after commit, with plaintext cleared on the entities.
func TestInsertUserAndIdentity_InTransaction(t *testing.T) {
	repo, pool := integrationEnv(t)
	ctx := context.Background()

	email := fmt.Sprintf("roundtrip-%s@example.com", uuid.NewString())
	user := &User{ID: uuid.New(), Name: "Round Trip", PrimaryEmail: email}
	cred := "bcrypt-hash-placeholder"
	identity := &AuthIdentity{
		ID:               uuid.New(),
		UserID:           user.ID,
		ProviderType:     "email_password",
		Identifier:       email,
		CredentialSecret: &cred,
	}
	plainToken := "secret-token-123"
	token := &AuthToken{
		ID:        uuid.New(),
		UserID:    user.ID,
		Purpose:   "email_verification",
		TokenHash: sha256Hex(plainToken),
		ExpiresAt: time.Now().Add(24 * time.Hour),
		CreatedAt: time.Now(),
	}

	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		t.Fatalf("begin tx: %v", err)
	}
	if err := repo.InsertUser(ctx, tx, user); err != nil {
		_ = tx.Rollback(ctx)
		t.Fatalf("InsertUser: %v", err)
	}
	if err := repo.InsertAuthIdentity(ctx, tx, identity); err != nil {
		_ = tx.Rollback(ctx)
		t.Fatalf("InsertAuthIdentity: %v", err)
	}
	if err := repo.InsertAuthToken(ctx, tx, token); err != nil {
		_ = tx.Rollback(ctx)
		t.Fatalf("InsertAuthToken: %v", err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatalf("commit: %v", err)
	}
	t.Cleanup(func() {
		_, _ = pool.Exec(ctx, "DELETE FROM users WHERE id = $1", user.ID)
	})

	// Plaintext must have been cleared by the inserts.
	if user.PrimaryEmail != "" {
		t.Errorf("user.PrimaryEmail not cleared after insert: %q", user.PrimaryEmail)
	}
	if identity.Identifier != "" {
		t.Errorf("identity.Identifier not cleared after insert: %q", identity.Identifier)
	}

	// The identity should be findable by its HMAC hash.
	// (Recompute since the entity no longer carries the email.)
	idHash := crypto.HMAC([]byte(email), integrationTestKeys(t))
	got, err := repo.FindAuthIdentityByIdentifierHash(ctx, "email_password", idHash)
	if err != nil {
		t.Fatalf("FindAuthIdentityByIdentifierHash: %v", err)
	}
	if got == nil {
		t.Fatal("identity not found after insert")
	}
	if got.UserID != user.ID || got.VerifiedAt != nil {
		t.Errorf("unexpected identity: %+v", got)
	}

	// The token should be findable by its hash.
	gotToken, err := repo.FindAuthTokenByHash(ctx, token.TokenHash)
	if err != nil {
		t.Fatalf("FindAuthTokenByHash: %v", err)
	}
	if gotToken == nil {
		t.Fatal("token not found after insert")
	}
	if gotToken.UsedAt != nil || gotToken.RevokedAt != nil {
		t.Errorf("new token should be unused/unrevoked: %+v", gotToken)
	}
}

// TestInsertAuthIdentity_ConcurrentDuplicate verifies that two concurrent
// inserts of the same (provider_type, identifier_hash) result in exactly
// one success and one wrapped *pgconn.PgError with code 23505 (R16 at the
// storage layer).
func TestInsertAuthIdentity_ConcurrentDuplicate(t *testing.T) {
	repo, pool := integrationEnv(t)
	ctx := context.Background()

	email := fmt.Sprintf("dup-%s@example.com", uuid.NewString())
	emailHash := crypto.HMAC([]byte(email), integrationTestKeys(t))

	makeIdentity := func() (*User, *AuthIdentity) {
		u := &User{ID: uuid.New(), Name: "Dup A", PrimaryEmail: email}
		ident := &AuthIdentity{
			ID:               uuid.New(),
			UserID:           u.ID,
			ProviderType:     "email_password",
			Identifier:       email,
			CredentialSecret: ptrString("h"),
		}
		return u, ident
	}

	doInsert := func() error {
		u, ident := makeIdentity()
		tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
		if err != nil {
			return fmt.Errorf("begin: %w", err)
		}
		if err := repo.InsertUser(ctx, tx, u); err != nil {
			_ = tx.Rollback(ctx)
			return err
		}
		if err := repo.InsertAuthIdentity(ctx, tx, ident); err != nil {
			_ = tx.Rollback(ctx)
			return err
		}
		return tx.Commit(ctx)
	}

	// First insert succeeds.
	if err := doInsert(); err != nil {
		t.Fatalf("first insert failed: %v", err)
	}
	// Clean up the winner's user row (cascades). The loser created nothing.
	// We don't know which UUID won, so clean by the shared email hash.
	t.Cleanup(func() {
		_, _ = pool.Exec(ctx, "DELETE FROM users WHERE primary_email_hash = $1", emailHash)
	})

	// Second concurrent insert for the same identifier should fail with a
	// unique violation detectable via errors.As on *pgconn.PgError.
	err := doInsert()
	if err == nil {
		t.Fatal("second insert of duplicate (provider_type, identifier_hash) succeeded; unique index not enforced")
	}
	var pgErr *pgconn.PgError
	if !errors.As(err, &pgErr) {
		t.Fatalf("error is not *pgconn.PgError (R16 detection broken): %v", err)
	}
	if pgErr.Code != "23505" {
		t.Fatalf("expected SQLSTATE 23505, got %s", pgErr.Code)
	}
}

// TestRedeemToken_Guards exercises the full 3-clause INV-account-08 guard
// via RedeemToken: valid → true; already-used → false; revoked (superseded)
// → false (regression for the spec-error 2-clause version); expired →
// false; non-existent → false.
func TestRedeemToken_Guards(t *testing.T) {
	repo, pool := integrationEnv(t)
	ctx := context.Background()
	user := newTestUser(t, repo, pool)
	t.Cleanup(func() {
		_, _ = pool.Exec(ctx, "DELETE FROM users WHERE id = $1", user.ID)
	})

	now := time.Now()
	cases := []struct {
		name      string
		tokenHash string
		setup     func() *AuthToken
		wantOK    bool
	}{
		{
			name:      "valid",
			tokenHash: sha256Hex("valid-token"),
			setup: func() *AuthToken {
				return &AuthToken{ID: uuid.New(), UserID: user.ID, Purpose: "email_verification",
					TokenHash: sha256Hex("valid-token"), ExpiresAt: now.Add(24 * time.Hour), CreatedAt: now}
			},
			wantOK: true,
		},
		{
			name:      "already used",
			tokenHash: sha256Hex("used-token"),
			setup: func() *AuthToken {
				used := now.Add(-time.Hour)
				return &AuthToken{ID: uuid.New(), UserID: user.ID, Purpose: "email_verification",
					TokenHash: sha256Hex("used-token"), ExpiresAt: now.Add(24 * time.Hour),
					UsedAt: &used, CreatedAt: now}
			},
			wantOK: false,
		},
		{
			name:      "revoked (superseded) — 3-clause guard regression",
			tokenHash: sha256Hex("revoked-token"),
			setup: func() *AuthToken {
				revoked := now.Add(-time.Hour)
				return &AuthToken{ID: uuid.New(), UserID: user.ID, Purpose: "email_verification",
					TokenHash: sha256Hex("revoked-token"), ExpiresAt: now.Add(24 * time.Hour),
					RevokedAt: &revoked, CreatedAt: now}
			},
			wantOK: false,
		},
		{
			name:      "expired",
			tokenHash: sha256Hex("expired-token"),
			setup: func() *AuthToken {
				return &AuthToken{ID: uuid.New(), UserID: user.ID, Purpose: "email_verification",
					TokenHash: sha256Hex("expired-token"), ExpiresAt: now.Add(-time.Hour), CreatedAt: now}
			},
			wantOK: false,
		},
	}

	// Insert each token directly into the DB so we can set used_at/revoked_at
	// to non-NULL (InsertAuthToken always inserts a fresh unused token).
	for _, tc := range cases {
		tok := tc.setup()
		if err := insertTokenDirect(ctx, pool, tok); err != nil {
			t.Fatalf("seed token (%s): %v", tc.name, err)
		}
	}

	// Non-existent case uses a hash that was never inserted.
	const nonExistent = "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ok, err := repo.RedeemToken(ctx, tc.tokenHash)
			if err != nil {
				t.Fatalf("RedeemToken (%s): unexpected error: %v", tc.name, err)
			}
			if ok != tc.wantOK {
				t.Errorf("RedeemToken (%s): got %v, want %v", tc.name, ok, tc.wantOK)
			}
		})
	}

	t.Run("non-existent", func(t *testing.T) {
		ok, err := repo.RedeemToken(ctx, nonExistent)
		if err != nil {
			t.Fatalf("RedeemToken (non-existent): %v", err)
		}
		if ok {
			t.Error("RedeemToken (non-existent): got true, want false")
		}
	})
}

// TestRevokeTokens_OnlyUnusedUnrevoked verifies RevokeTokens sets
// revoked_at only for unused, unrevoked tokens of the given purpose;
// already-used and already-revoked tokens keep their state, and tokens
// of a different purpose are untouched.
func TestRevokeTokens_OnlyUnusedUnrevoked(t *testing.T) {
	repo, pool := integrationEnv(t)
	ctx := context.Background()
	user := newTestUser(t, repo, pool)

	now := time.Now()
	used := now.Add(-time.Hour)
	revoked := now.Add(-2 * time.Hour)

	type seed struct {
		hash      string
		purpose   string
		usedAt    *time.Time
		revokedAt *time.Time
	}
	seeds := []seed{
		{hash: sha256Hex("fresh-ev"), purpose: "email_verification"},
		{hash: sha256Hex("used-ev"), purpose: "email_verification", usedAt: &used},
		{hash: sha256Hex("revoked-ev"), purpose: "email_verification", revokedAt: &revoked},
		{hash: sha256Hex("fresh-pr"), purpose: "password_reset"}, // different purpose — must NOT be revoked
	}
	for _, s := range seeds {
		tok := &AuthToken{
			ID: uuid.New(), UserID: user.ID, Purpose: s.purpose,
			TokenHash: s.hash, ExpiresAt: now.Add(24 * time.Hour),
			UsedAt: s.usedAt, RevokedAt: s.revokedAt, CreatedAt: now,
		}
		if err := insertTokenDirect(ctx, pool, tok); err != nil {
			t.Fatalf("seed token: %v", err)
		}
	}

	// RevokeTokens now takes the caller's tx (resend does revoke+insert
	// atomically). Here we test revoke in isolation: begin, revoke,
	// commit.
	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		t.Fatalf("begin tx: %v", err)
	}
	if err := repo.RevokeTokens(ctx, tx, user.ID, "email_verification"); err != nil {
		_ = tx.Rollback(ctx)
		t.Fatalf("RevokeTokens: %v", err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatalf("commit: %v", err)
	}

	// The fresh EV token must now be revoked.
	fresh, _ := repo.FindAuthTokenByHash(ctx, sha256Hex("fresh-ev"))
	if fresh == nil || fresh.RevokedAt == nil {
		t.Error("fresh email_verification token was not revoked")
	}
	// The used EV token keeps its used_at and must NOT gain a revoked_at
	// (the used_at IS NULL guard excludes it).
	usedTok, _ := repo.FindAuthTokenByHash(ctx, sha256Hex("used-ev"))
	if usedTok == nil || usedTok.RevokedAt != nil {
		t.Error("already-used token should not be touched by RevokeTokens")
	}
	// The already-revoked EV token keeps its earlier revoked_at.
	revTok, _ := repo.FindAuthTokenByHash(ctx, sha256Hex("revoked-ev"))
	if revTok == nil || revTok.RevokedAt == nil {
		t.Error("already-revoked token lost its revoked_at")
	}
	// The password_reset token must NOT be revoked (different purpose).
	prTok, _ := repo.FindAuthTokenByHash(ctx, sha256Hex("fresh-pr"))
	if prTok == nil || prTok.RevokedAt != nil {
		t.Error("password_reset token should not be revoked by an email_verification revoke call")
	}
}

// sha256Hex is defined in service.go (same package); the integration
// tests reuse it rather than redeclaring, to avoid a duplicate-symbol
// build failure under the integration build tag.

// insertTokenDirect inserts an AuthToken with arbitrary used_at/revoked_at
// (InsertAuthToken always inserts a fresh unused token, so tests that need
// an already-used/revoked/expired row seed it here). SQL is built with
// goqu per the AGENTS.md golden rule; nil pointer columns are omitted so
// the DB default (NULL) applies, avoiding the nil-pointer-in-interface
// pitfall.
func insertTokenDirect(ctx context.Context, pool *pgxpool.Pool, tok *AuthToken) error {
	rec := goqu.Record{
		"id":         tok.ID,
		"user_id":    tok.UserID,
		"purpose":    tok.Purpose,
		"token_hash": tok.TokenHash,
		"expires_at": tok.ExpiresAt,
		"created_at": tok.CreatedAt,
	}
	if tok.UsedAt != nil {
		rec["used_at"] = *tok.UsedAt
	}
	if tok.RevokedAt != nil {
		rec["revoked_at"] = *tok.RevokedAt
	}
	sqlStr, args, err := goqu.Dialect("postgres").Insert("auth_tokens").
		Rows(rec).Prepared(true).ToSQL()
	if err != nil {
		return fmt.Errorf("build seed insert: %w", err)
	}
	if _, err := pool.Exec(ctx, sqlStr, args...); err != nil {
		return fmt.Errorf("seed insert: %w", err)
	}
	return nil
}

func ptrString(s string) *string { return &s }
