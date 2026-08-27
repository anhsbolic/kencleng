//go:build integration

// package account — account-linking (Fitur 4) integration tests against
// a real Postgres. Proves what fakes structurally cannot:
//
//   - R7: Branch 2's credential update + user-wide session revocation
//     commit atomically against real Postgres;
//   - R9: unlink hard-deletes the google identity and commits the audit
//     entry in the same transaction;
//   - R13: ≥100 concurrent unlinks against the same user under FOR UPDATE
//     — exactly one deletes, the rest land in idempotent 200, and the
//     INV-account-02/12 guard holds at every observed end-state;
//   - R14: link-purpose token redemption writes the audit entry in-tx.
//
// Run: go test -tags=integration -race ./internal/domain/account/...
package account

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/anhsbolic/kencleng/backend/internal/platform/secrets"
)

// integrationSecurityService returns a real Service with compare set to
// secrets.ComparePassword (Branch 2 / unlink re-auth needs real bcrypt
// comparison against the identity's stored credential_secret).
func integrationSecurityService(t *testing.T) (*Service, *RepositoryDB, *pgxpool.Pool) {
	t.Helper()
	svc, repo, pool := integrationService(t)
	svc.compare = secrets.ComparePassword
	return svc, repo, pool
}

// seedGoogleIdentity inserts a verified google identity for userID.
// InsertAuthIdentity always omits verified_at (NULL on insert), so a
// direct UPDATE is needed — google identities are born verified per spec
// (R14), but the adapter doesn't persist that field on insert.
func seedGoogleIdentity(t *testing.T, repo *RepositoryDB, pool *pgxpool.Pool, userID uuid.UUID, email string) {
	t.Helper()
	ctx := context.Background()
	now := time.Now()
	ident := &AuthIdentity{
		ID:           uuid.New(),
		UserID:       userID,
		ProviderType: "google",
		Identifier:   email,
		VerifiedAt:   &now,
	}
	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		t.Fatalf("begin tx: %v", err)
	}
	if err := repo.InsertAuthIdentity(ctx, tx, ident); err != nil {
		_ = tx.Rollback(ctx)
		t.Fatalf("insert google identity: %v", err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatalf("commit: %v", err)
	}
	// Manually set verified_at (adapter omits it on insert).
	if _, err := pool.Exec(ctx,
		`UPDATE auth_identities SET verified_at = $2 WHERE id = $1`,
		ident.ID, now); err != nil {
		t.Fatalf("set google verified_at: %v", err)
	}
}

// seedVerifiedEmailPasswordIdentity inserts a verified email_password
// identity with a known bcrypt hash for userID. Same verified_at
// workaround as seedGoogleIdentity.
func seedVerifiedEmailPasswordIdentity(t *testing.T, repo *RepositoryDB, pool *pgxpool.Pool, userID uuid.UUID, email, password string) {
	t.Helper()
	ctx := context.Background()
	hash, err := secrets.HashPassword(password)
	if err != nil {
		t.Fatalf("hash: %v", err)
	}
	now := time.Now()
	ident := &AuthIdentity{
		ID:               uuid.New(),
		UserID:           userID,
		ProviderType:     "email_password",
		Identifier:       email,
		CredentialSecret: &hash,
		VerifiedAt:       &now,
	}
	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		t.Fatalf("begin tx: %v", err)
	}
	if err := repo.InsertAuthIdentity(ctx, tx, ident); err != nil {
		_ = tx.Rollback(ctx)
		t.Fatalf("insert email_password identity: %v", err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatalf("commit: %v", err)
	}
	// Manually set verified_at (adapter omits it on insert).
	if _, err := pool.Exec(ctx,
		`UPDATE auth_identities SET verified_at = $2 WHERE id = $1`,
		ident.ID, now); err != nil {
		t.Fatalf("set email_password verified_at: %v", err)
	}
}

// seedRefreshToken inserts a non-revoked refresh token for userID.
func seedRefreshToken(t *testing.T, repo *RepositoryDB, pool *pgxpool.Pool, userID uuid.UUID) {
	t.Helper()
	ctx := context.Background()
	rt := &RefreshToken{
		ID:        uuid.New(),
		UserID:    userID,
		FamilyID:  uuid.New(),
		TokenHash: sha256Hex(uuid.NewString()),
		ExpiresAt: time.Now().Add(24 * time.Hour),
		CreatedAt: time.Now(),
	}
	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		t.Fatalf("begin tx: %v", err)
	}
	if err := repo.InsertRefreshToken(ctx, tx, rt); err != nil {
		_ = tx.Rollback(ctx)
		t.Fatalf("insert refresh token: %v", err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatalf("commit: %v", err)
	}
}

// countUnrevokedSessionsForUser returns the count of non-revoked
// refresh_tokens rows for userID.
func countUnrevokedSessionsForUser(t *testing.T, pool *pgxpool.Pool, userID uuid.UUID) int {
	t.Helper()
	ctx := context.Background()
	var n int
	if err := pool.QueryRow(ctx,
		`SELECT COUNT(*) FROM refresh_tokens WHERE user_id = $1 AND revoked_at IS NULL`,
		userID).Scan(&n); err != nil {
		t.Fatalf("count sessions: %v", err)
	}
	return n
}

// googleIdentityCount returns the count of google auth_identities for userID.
func googleIdentityCount(t *testing.T, pool *pgxpool.Pool, userID uuid.UUID) int {
	t.Helper()
	ctx := context.Background()
	var n int
	if err := pool.QueryRow(ctx,
		`SELECT COUNT(*) FROM auth_identities WHERE user_id = $1 AND provider_type = 'google'`,
		userID).Scan(&n); err != nil {
		t.Fatalf("count google identities: %v", err)
	}
	return n
}

// userLogCountForAction returns the count of user_logs rows for userID
// with the given action_type.
func userLogCountForAction(t *testing.T, pool *pgxpool.Pool, userID uuid.UUID, action string) int {
	t.Helper()
	ctx := context.Background()
	var n int
	if err := pool.QueryRow(ctx,
		`SELECT COUNT(*) FROM user_logs WHERE user_id = $1 AND action_type = $2`,
		userID, action).Scan(&n); err != nil {
		t.Fatalf("count user_logs: %v", err)
	}
	return n
}

// ---- R7: Branch 2 atomic change (credential + session revocation) ----

func TestSetPassword_Branch2_AllSessionsRevoked_RealDB(t *testing.T) {
	svc, repo, pool := integrationSecurityService(t)
	ctx := context.Background()

	user := newTestUser(t, repo, pool)
	email := fmt.Sprintf("ep-%s@example.com", uuid.NewString())
	seedVerifiedEmailPasswordIdentity(t, repo, pool, user.ID, email, "old-pw-123")
	seedRefreshToken(t, repo, pool, user.ID)
	seedRefreshToken(t, repo, pool, user.ID)

	// Pre-condition: 2 unrevoked sessions.
	if got := countUnrevokedSessionsForUser(t, pool, user.ID); got != 2 {
		t.Fatalf("pre: expected 2 sessions, got %d", got)
	}

	_, err := svc.SetPassword(ctx, user.ID, "", "old-pw-123", "new-strong-pw")
	if err != nil {
		t.Fatalf("SetPassword Branch 2: %v", err)
	}

	// All sessions revoked.
	if got := countUnrevokedSessionsForUser(t, pool, user.ID); got != 0 {
		t.Errorf("post: expected 0 unrevoked sessions, got %d", got)
	}

	// Credential changed (the old password no longer matches).
	identities, err := repo.FindAuthIdentitiesByUser(ctx, user.ID)
	if err != nil {
		t.Fatalf("find identities: %v", err)
	}
	for _, id := range identities {
		if id.ProviderType == "email_password" && id.CredentialSecret != nil {
			if err := secrets.ComparePassword(*id.CredentialSecret, "new-strong-pw"); err != nil {
				t.Errorf("credential must match the new password")
			}
			if err := secrets.ComparePassword(*id.CredentialSecret, "old-pw-123"); err == nil {
				t.Errorf("credential must NOT match the old password")
			}
		}
	}
}

// ---- R9: unlink hard-deletes + audits atomically --------------------

func TestUnlinkGoogle_Success_HardDeletesAndAudits_RealDB(t *testing.T) {
	svc, repo, pool := integrationSecurityService(t)
	ctx := context.Background()

	user := newTestUser(t, repo, pool)
	seedGoogleIdentity(t, repo, pool, user.ID, fmt.Sprintf("g-%s@gmail.com", uuid.NewString()))
	seedVerifiedEmailPasswordIdentity(t, repo, pool, user.ID, fmt.Sprintf("ep-%s@example.com", uuid.NewString()), "correct-pw")

	// Pre: 1 google identity.
	if got := googleIdentityCount(t, pool, user.ID); got != 1 {
		t.Fatalf("pre: expected 1 google identity, got %d", got)
	}

	err := svc.UnlinkGoogle(ctx, user.ID, "correct-pw")
	if err != nil {
		t.Fatalf("UnlinkGoogle: %v", err)
	}

	// Post: 0 google identities (hard-deleted).
	if got := googleIdentityCount(t, pool, user.ID); got != 0 {
		t.Errorf("post: expected 0 google identities, got %d", got)
	}

	// Audit entry committed.
	if got := userLogCountForAction(t, pool, user.ID, actionAccountLinking); got < 1 {
		t.Errorf("expected ≥1 audit entry with action_type=%s, got %d", actionAccountLinking, got)
	}
}

// ---- R13: ≥100 concurrent unlinks under FOR UPDATE --------------------

func TestUnlinkGoogle_ConcurrentRequests_GuardHolds_RealDB(t *testing.T) {
	svc, repo, pool := integrationSecurityService(t)
	ctx := context.Background()

	user := newTestUser(t, repo, pool)
	seedGoogleIdentity(t, repo, pool, user.ID, fmt.Sprintf("g-%s@gmail.com", uuid.NewString()))
	seedVerifiedEmailPasswordIdentity(t, repo, pool, user.ID, fmt.Sprintf("ep-%s@example.com", uuid.NewString()), "correct-pw")

	const goroutines = 100
	var wg sync.WaitGroup
	var successCount atomic.Int64
	var errCount atomic.Int64

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := svc.UnlinkGoogle(ctx, user.ID, "correct-pw")
			if err == nil {
				successCount.Add(1)
			} else {
				errCount.Add(1)
			}
		}()
	}
	wg.Wait()

	// Invariant: google identity is deleted (count = 0).
	if got := googleIdentityCount(t, pool, user.ID); got != 0 {
		t.Errorf("post: expected 0 google identities after %d concurrent unlinks, got %d", goroutines, got)
	}

	// Invariant: user still has ≥1 identity (INV-account-02 holds).
	identities, err := repo.FindAuthIdentitiesByUser(ctx, user.ID)
	if err != nil {
		t.Fatalf("find identities: %v", err)
	}
	if len(identities) < 1 {
		t.Errorf("INV-account-02 violated: user has 0 identities after unlink")
	}

	// Invariant: user has ≥1 VERIFIED identity (INV-account-12 holds).
	hasVerified := false
	for _, id := range identities {
		if id.VerifiedAt != nil {
			hasVerified = true
			break
		}
	}
	if !hasVerified {
		t.Errorf("INV-account-12 violated: no verified identity remains after unlink")
	}

	// Invariant: all 100 goroutines succeeded (idempotent nil for
	// losers, nil for the winner). No spurious 409s.
	if got := errCount.Load(); got != 0 {
		t.Errorf("expected 0 errors from %d concurrent unlinks, got %d", goroutines, got)
	}
	if got := successCount.Load(); got != goroutines {
		t.Errorf("expected %d successes (idempotent), got %d", goroutines, got)
	}

	// Exactly 1 audit entry (the winner; idempotent no-ops write none).
	if got := userLogCountForAction(t, pool, user.ID, actionAccountLinking); got != 1 {
		t.Errorf("expected exactly 1 audit entry, got %d", got)
	}
}

// ---- R14: link-purpose redemption writes audit ------------------------

func TestVerifyEmail_LinkPurpose_WritesAudit_RealDB(t *testing.T) {
	svc, repo, pool := integrationSecurityService(t)
	ctx := context.Background()

	user := newTestUser(t, repo, pool)
	email := fmt.Sprintf("link-%s@example.com", uuid.NewString())

	// Seed an UNVERIFIED email_password identity for the user.
	ident := &AuthIdentity{
		ID:           uuid.New(),
		UserID:       user.ID,
		ProviderType: "email_password",
		Identifier:   email,
	}
	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		t.Fatalf("begin tx: %v", err)
	}
	if err := repo.InsertAuthIdentity(ctx, tx, ident); err != nil {
		_ = tx.Rollback(ctx)
		t.Fatalf("insert identity: %v", err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatalf("commit: %v", err)
	}

	// Seed a link-purpose token.
	plainToken, tokenHash, err := generateToken()
	if err != nil {
		t.Fatalf("generate token: %v", err)
	}
	token := &AuthToken{
		ID:        uuid.New(),
		UserID:    user.ID,
		Purpose:   purposeEmailVerifyLink,
		TokenHash: tokenHash,
		ExpiresAt: time.Now().Add(tokenTTL),
		CreatedAt: time.Now(),
	}
	tx2, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		t.Fatalf("begin tx2: %v", err)
	}
	if err := repo.InsertAuthToken(ctx, tx2, token); err != nil {
		_ = tx2.Rollback(ctx)
		t.Fatalf("insert token: %v", err)
	}
	if err := tx2.Commit(ctx); err != nil {
		t.Fatalf("commit2: %v", err)
	}

	preLogs := userLogCountForAction(t, pool, user.ID, actionAccountLinking)

	// Redeem the link-purpose token.
	err = svc.VerifyEmail(ctx, plainToken)
	if err != nil {
		t.Fatalf("VerifyEmail: %v", err)
	}

	// Identity now verified.
	identities, err := repo.FindAuthIdentitiesByUser(ctx, user.ID)
	if err != nil {
		t.Fatalf("find identities: %v", err)
	}
	for _, id := range identities {
		if id.ProviderType == "email_password" && id.VerifiedAt == nil {
			t.Errorf("identity must be verified after link-purpose redemption")
		}
	}

	// Audit entry committed in the same tx.
	postLogs := userLogCountForAction(t, pool, user.ID, actionAccountLinking)
	if postLogs != preLogs+1 {
		t.Errorf("expected %d→%d audit entries, got %d→%d", preLogs, preLogs+1, preLogs, postLogs)
	}
}
