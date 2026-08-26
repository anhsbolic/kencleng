//go:build integration

// package account — password-reset (Fitur 2B) integration tests against a
// real Postgres. Proves what fakes structurally cannot:
//
//   - INV-account-05: the credential update and the mass session revoke
//     commit atomically — and roll back atomically when the revoke write
//     fails between the two writes (R18's property arm);
//   - INV-account-08: N concurrent resets of the SAME reset token produce
//     exactly one winner under true DB-level contention, incl. a ≥100-
//     goroutine stress mix of valid submits and replays;
//   - forgot-password branch timing stays within an equivalent band against
//     real Postgres (the dummyWrite anti-enumeration device actually works
//     at the DB-time layer);
//   - expired tokens cause zero state change.
//
// Run:
//
//	go test -tags=integration -race ./internal/domain/account/...
package account

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/anhsbolic/kencleng/backend/internal/platform/secrets"
)

// seedResetUser creates a user + verified email_password identity with a
// placeholder credential, ready for a reset flow. Cleanup deletes the user,
// cascading to identities, auth_tokens, and refresh_tokens.
func seedResetUser(t *testing.T, repo *RepositoryDB, pool *pgxpool.Pool) (*User, string) {
	t.Helper()
	ctx := context.Background()
	email := fmt.Sprintf("reset-%s@example.com", uuid.NewString())
	u := &User{ID: uuid.New(), Name: "Reset Integration", PrimaryEmail: email}
	ident := &AuthIdentity{
		ID: uuid.New(), UserID: u.ID, ProviderType: providerEmailPassword,
		Identifier: email, CredentialSecret: ptrString("old-bcrypt-hash"),
	}
	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		t.Fatalf("begin: %v", err)
	}
	if err := repo.InsertUser(ctx, tx, u); err != nil {
		_ = tx.Rollback(ctx)
		t.Fatalf("insert user: %v", err)
	}
	if err := repo.InsertAuthIdentity(ctx, tx, ident); err != nil {
		_ = tx.Rollback(ctx)
		t.Fatalf("insert identity: %v", err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatalf("commit: %v", err)
	}
	t.Cleanup(func() { _, _ = pool.Exec(ctx, "DELETE FROM users WHERE id = $1", u.ID) })
	return u, email
}

// seedResetTokenRow inserts a password_reset auth_token row for userID
// with the given TTL and returns the plain token. Cascades away with the
// user.
func seedResetTokenRow(t *testing.T, repo *RepositoryDB, pool *pgxpool.Pool, userID uuid.UUID, ttl time.Duration) string {
	t.Helper()
	ctx := context.Background()
	plain := "reset-" + uuid.NewString()
	tok := &AuthToken{
		ID: uuid.New(), UserID: userID, Purpose: purposePasswordReset,
		TokenHash: sha256Hex(plain), ExpiresAt: time.Now().Add(ttl), CreatedAt: time.Now(),
	}
	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		t.Fatalf("begin: %v", err)
	}
	if err := repo.InsertAuthToken(ctx, tx, tok); err != nil {
		_ = tx.Rollback(ctx)
		t.Fatalf("seed reset token: %v", err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatalf("commit: %v", err)
	}
	return plain
}

type refreshTokenSeed struct {
	familyID uuid.UUID
	rotated  bool // replaced_by_id set (points at another seeded row)
	revoked  bool // already revoked before the reset runs
}

// seedRefreshTokens inserts refresh_tokens rows for userID and returns
// their hashes. InsertRefreshToken always writes NULL revoked_at /
// replaced_by_id (new-token contract), so rotated/revoked seed states are
// applied with direct follow-up UPDATEs. Cascades away with the user.
func seedRefreshTokens(t *testing.T, repo *RepositoryDB, pool *pgxpool.Pool, userID uuid.UUID, seeds []refreshTokenSeed) []string {
	t.Helper()
	ctx := context.Background()
	var hashes []string
	for _, s := range seeds {
		plain := "sess-" + uuid.NewString()
		hashes = append(hashes, sha256Hex(plain))
		rowID := uuid.New()
		row := &RefreshToken{
			ID: rowID, UserID: userID, FamilyID: s.familyID,
			TokenHash: sha256Hex(plain), ExpiresAt: time.Now().Add(time.Hour), CreatedAt: time.Now(),
		}
		tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
		if err != nil {
			t.Fatalf("begin: %v", err)
		}
		if err := repo.InsertRefreshToken(ctx, tx, row); err != nil {
			_ = tx.Rollback(ctx)
			t.Fatalf("seed refresh token: %v", err)
		}
		if s.rotated {
			if _, err := tx.Exec(ctx,
				`UPDATE refresh_tokens SET replaced_by_id = $1 WHERE id = $2`,
				uuid.New(), rowID); err != nil {
				_ = tx.Rollback(ctx)
				t.Fatalf("mark rotated: %v", err)
			}
		}
		if s.revoked {
			if _, err := tx.Exec(ctx,
				`UPDATE refresh_tokens SET revoked_at = now() WHERE id = $1`, rowID); err != nil {
				_ = tx.Rollback(ctx)
				t.Fatalf("mark revoked: %v", err)
			}
		}
		if err := tx.Commit(ctx); err != nil {
			t.Fatalf("commit: %v", err)
		}
	}
	return hashes
}

// countUnrevokedSessions returns how many refresh_tokens rows for userID
// are still live (revoked_at IS NULL).
func countUnrevokedSessions(t *testing.T, pool *pgxpool.Pool, userID uuid.UUID) int {
	t.Helper()
	var n int
	if err := pool.QueryRow(context.Background(),
		`SELECT count(*) FROM refresh_tokens WHERE user_id = $1 AND revoked_at IS NULL`,
		userID).Scan(&n); err != nil {
		t.Fatalf("count sessions: %v", err)
	}
	return n
}

// credentialSecret reads the current bcrypt hash for the user's
// email_password identity.
func credentialSecret(t *testing.T, pool *pgxpool.Pool, userID uuid.UUID) string {
	t.Helper()
	var secret *string
	if err := pool.QueryRow(context.Background(),
		`SELECT credential_secret FROM auth_identities WHERE user_id = $1 AND provider_type = $2`,
		userID, providerEmailPassword).Scan(&secret); err != nil {
		t.Fatalf("read credential_secret: %v", err)
	}
	if secret == nil {
		return ""
	}
	return *secret
}

// INV-account-05 / R7: one reset revokes EVERY session across ALL families
// — including already-rotated rows — in the same transaction as the
// credential update.
func TestResetPassword_AllSessionsRevoked_Atomic_RealDB(t *testing.T) {
	svc, repo, pool := integrationService(t)
	ctx := context.Background()
	user, _ := seedResetUser(t, repo, pool)

	f1, f2 := uuid.New(), uuid.New()
	seedRefreshTokens(t, repo, pool, user.ID, []refreshTokenSeed{
		{familyID: f1},                // active, family 1
		{familyID: f2},                // active, family 2 (other device)
		{familyID: f2, rotated: true}, // rotated-out parent
		{familyID: f1, revoked: true}, // logged out earlier — must stay revoked
	})
	plain := seedResetTokenRow(t, repo, pool, user.ID, time.Hour)
	// Pre-reset sanity: 2 active + 1 rotated-out = 3 unrevoked (the 4th
	// row was seeded already-revoked).
	if got := countUnrevokedSessions(t, pool, user.ID); got != 3 {
		t.Fatalf("seed sanity: unrevoked sessions = %d, want 3", got)
	}
	oldSecret := credentialSecret(t, pool, user.ID)

	if err := svc.ResetPassword(ctx, plain, "brand-new-pw-9"); err != nil {
		t.Fatalf("ResetPassword: %v", err)
	}

	if got := countUnrevokedSessions(t, pool, user.ID); got != 0 {
		t.Errorf("unrevoked sessions after reset = %d, want 0 (INV-account-05)", got)
	}
	newSecret := credentialSecret(t, pool, user.ID)
	if newSecret == "" || newSecret == oldSecret {
		t.Error("credential_secret must be updated by the reset")
	}
	if !secretsCompare(newSecret, "brand-new-pw-9") {
		t.Error("stored credential must verify against the new password")
	}
}

// stubHashService clones a real-pool service but replaces bcrypt hashing
// with an instant stub. The concurrency invariants under test (INV-05/08)
// are about DB row contention, not crypto cost — real bcrypt at ~1s/hash
// under -race would turn 100+ racer runs into multi-minute CPU burns that
// prove nothing the guarded UPDATE doesn't already prove. Production keeps
// secrets.HashPassword via the NewService default.
func stubHashService(svc *Service) *Service {
	svc.hashPassword = func(password string) (string, error) {
		return "stub-hash:" + password, nil
	}
	return svc
}

// INV-account-08 / R11: 100 concurrent resets of the SAME token — exactly
// one winner, losers see the consumed token as not-found.
//
// DEFERRED (Anhar, 2026-08-26): gated behind KENCLENG_HEAVY_RACE_TESTS=1
// for this session only — see techplan Open Items. This skip is TEMPORARY;
// the Tier 1 gate is not cleared until it runs clean.
func TestResetPassword_TokenSingleUse_Concurrent_RealDB(t *testing.T) {
	if os.Getenv("KENCLENG_HEAVY_RACE_TESTS") == "" {
		t.Skip("heavy concurrency proof deferred (techplan Open Item); set KENCLENG_HEAVY_RACE_TESTS=1 to run")
	}
	svc, repo, pool := integrationService(t)
	svc = stubHashService(svc)
	ctx := context.Background()
	user, _ := seedResetUser(t, repo, pool)
	plain := seedResetTokenRow(t, repo, pool, user.ID, time.Hour)

	const racers = 100
	var wins atomic.Int32
	var badErr atomic.Int32
	var wg sync.WaitGroup
	for i := 0; i < racers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := svc.ResetPassword(ctx, plain, "stress-pw-value")
			switch {
			case err == nil:
				wins.Add(1)
			case errors.Is(err, ErrTokenNotFound):
				// expected loser outcome
			default:
				badErr.Add(1)
				t.Logf("unexpected racer error: %v", err)
			}
		}()
	}
	wg.Wait()

	if w := wins.Load(); w != 1 {
		t.Errorf("winners = %d, want exactly 1 across %d racers", w, racers)
	}
	if b := badErr.Load(); b != 0 {
		t.Errorf("%d racers failed with unexpected error classes", b)
	}
	if got := credentialSecret(t, pool, user.ID); got != "stub-hash:stress-pw-value" {
		t.Errorf("final credential = %q, want the single winner's stub hash", got)
	}
}

// Stress mix (R11): ≥100 goroutines split between racing the one valid
// token and replaying garbage/consumed values. Zero invariant violations
// allowed: at most one success total; garbage never succeeds.
//
// DEFERRED (Anhar, 2026-08-26): same gate as TokenSingleUse above.
func TestResetPassword_Stress_MixedValidAndReplayed_RealDB(t *testing.T) {
	if os.Getenv("KENCLENG_HEAVY_RACE_TESTS") == "" {
		t.Skip("heavy concurrency proof deferred (techplan Open Item); set KENCLENG_HEAVY_RACE_TESTS=1 to run")
	}
	svc, repo, pool := integrationService(t)
	svc = stubHashService(svc)
	ctx := context.Background()
	user, _ := seedResetUser(t, repo, pool)
	validPlain := seedResetTokenRow(t, repo, pool, user.ID, time.Hour)

	var successes atomic.Int32
	var violations atomic.Int32
	var wg sync.WaitGroup

	racer := func(token string, expectSuccess bool) {
		defer wg.Done()
		err := svc.ResetPassword(ctx, token, "mixed-stress-pw")
		switch {
		case err == nil:
			successes.Add(1)
		case errors.Is(err, ErrTokenNotFound), errors.Is(err, ErrTokenExpired):
			// fine
		default:
			violations.Add(1)
			t.Logf("unexpected error class: %v", err)
		}
		_ = expectSuccess // classification happens via the success counter
	}

	const perGroup = 60
	for i := 0; i < perGroup; i++ {
		wg.Add(2)
		go racer(validPlain, true)
		go racer("garbage-"+uuid.NewString(), false)
	}
	wg.Wait()

	if s := successes.Load(); s > 1 {
		t.Errorf("successes = %d, want ≤ 1 (single-use violated)", s)
	}
	if v := violations.Load(); v != 0 {
		t.Errorf("%d unexpected error classes during stress", v)
	}
}

// R18 rollback arm: if the session-revoke write fails AFTER the credential
// update inside the tx, the WHOLE transaction rolls back — token still
// redeemable, sessions still alive, credential untouched.
func TestResetPassword_FailureBetweenWrites_RollsBackBoth_RealDB(t *testing.T) {
	_, realRepo, pool := integrationService(t)
	ctx := context.Background()
	user, _ := seedResetUser(t, realRepo, pool)
	seedRefreshTokens(t, realRepo, pool, user.ID, []refreshTokenSeed{{familyID: uuid.New()}})
	plain := seedResetTokenRow(t, realRepo, pool, user.ID, time.Hour)
	oldSecret := credentialSecret(t, pool, user.ID)

	// Wrap the real repository: delegate everything, fail only on the
	// mass revoke (the write BETWEEN redeem/update and commit).
	failRepo := &failingRevokeRepo{Repository: realRepo}
	svc := &Service{
		repo:        failRepo,
		tx:          poolRunner{pool: pool},
		breachCheck: integrationBreachCheckerFalse{},
		email:       integrationSilentSender{},
		keys:        realRepo.keys,
	}

	if err := svc.ResetPassword(ctx, plain, "never-committed-pw"); err == nil {
		t.Fatal("ResetPassword must fail when the session revoke fails")
	}

	// Token NOT consumed (redeem rolled back — Assumption B's structural twin).
	var usedAt *time.Time
	if err := pool.QueryRow(ctx,
		`SELECT used_at FROM auth_tokens WHERE token_hash = $1`,
		sha256Hex(plain)).Scan(&usedAt); err != nil {
		t.Fatalf("read token: %v", err)
	}
	if usedAt != nil {
		t.Error("token must remain unconsumed after a mid-transaction failure")
	}
	// Sessions NOT revoked.
	if got := countUnrevokedSessions(t, pool, user.ID); got != 1 {
		t.Errorf("sessions alive after rollback = %d, want 1", got)
	}
	// Credential NOT changed.
	if got := credentialSecret(t, pool, user.ID); got != oldSecret {
		t.Errorf("credential changed despite rollback: %q -> %q", oldSecret, got)
	}
}

// failingRevokeRepo delegates everything to the real repository except
// RevokeAllRefreshTokensForUser, which fails on demand.
type failingRevokeRepo struct {
	Repository
}

func (f *failingRevokeRepo) RevokeAllRefreshTokensForUser(context.Context, pgx.Tx, uuid.UUID) error {
	return errors.New("injected revoke failure (test)")
}

// R9: expired token → ErrTokenExpired, zero state change anywhere.
func TestResetPassword_ExpiredToken_NoStateChange_RealDB(t *testing.T) {
	svc, repo, pool := integrationService(t)
	ctx := context.Background()
	user, _ := seedResetUser(t, repo, pool)
	seedRefreshTokens(t, repo, pool, user.ID, []refreshTokenSeed{{familyID: uuid.New()}})
	plain := seedResetTokenRow(t, repo, pool, user.ID, -time.Minute)

	if err := svc.ResetPassword(ctx, plain, "should-not-apply-pw"); !errors.Is(err, ErrTokenExpired) {
		t.Fatalf("err = %v, want ErrTokenExpired", err)
	}
	if got := countUnrevokedSessions(t, pool, user.ID); got != 1 {
		t.Errorf("expired token must change nothing: sessions = %d, want 1", got)
	}
	if got := credentialSecret(t, pool, user.ID); got != "old-bcrypt-hash" {
		t.Errorf("credential changed on expired token: %q", got)
	}
	var usedAt *time.Time
	if err := pool.QueryRow(ctx,
		`SELECT used_at FROM auth_tokens WHERE token_hash = $1`,
		sha256Hex(plain)).Scan(&usedAt); err != nil {
		t.Fatalf("read token: %v", err)
	}
	if usedAt != nil {
		t.Error("expired token must stay unconsumed")
	}
}

// R5 DB-time half: forgot-password's three branches stay within a ≤3×
// wall-clock band against real Postgres (registered does lookup+INSERT tx;
// no-op branches do lookups+dummyWrite tx). Band is slightly wider than
// Register's 2× because the branches' read counts differ by one lookup.
func TestForgotPassword_Timing_Branches_RealPostgres(t *testing.T) {
	svc, repo, pool := integrationService(t)
	ctx := context.Background()

	registeredEmail := fmt.Sprintf("fp-reg-%s@example.com", uuid.NewString())
	regUser := &User{ID: uuid.New(), Name: "FP", PrimaryEmail: registeredEmail}
	regIdent := &AuthIdentity{
		ID: uuid.New(), UserID: regUser.ID, ProviderType: providerEmailPassword,
		Identifier: registeredEmail, CredentialSecret: ptrString("h"),
	}
	seedVerifiedIdentity(t, repo, pool, regUser, regIdent)
	t.Cleanup(func() { _, _ = pool.Exec(ctx, "DELETE FROM users WHERE id = $1", regUser.ID) })

	googleOnlyEmail := fmt.Sprintf("fp-goog-%s@example.com", uuid.NewString())
	googUser := &User{ID: uuid.New(), Name: "FP", PrimaryEmail: googleOnlyEmail}
	googIdent := &AuthIdentity{
		ID: uuid.New(), UserID: googUser.ID, ProviderType: providerGoogle,
		Identifier: googleOnlyEmail,
	}
	seedVerifiedIdentity(t, repo, pool, googUser, googIdent)
	t.Cleanup(func() { _, _ = pool.Exec(ctx, "DELETE FROM users WHERE id = $1", googUser.ID) })

	unknownEmail := fmt.Sprintf("fp-none-%s@example.com", uuid.NewString())

	cases := []struct {
		name  string
		email func() string
	}{
		{"registered", func() string { return registeredEmail }},
		{"google-only", func() string { return googleOnlyEmail }},
		{"unknown", func() string { return unknownEmail }},
	}

	// Warm-up pass primes plan caches so first calls aren't outliers.
	for _, tc := range cases {
		if err := svc.ForgotPassword(ctx, tc.email()); err != nil {
			t.Fatalf("%s warmup: %v", tc.name, err)
		}
	}
	// The registered-branch warmup issued one real token; its user cleanup
	// above cascades it away.

	var durations []time.Duration
	for round := 0; round < 5; round++ {
		for _, tc := range cases {
			start := time.Now()
			if err := svc.ForgotPassword(ctx, tc.email()); err != nil {
				t.Fatalf("%s round %d: %v", tc.name, round, err)
			}
			durations = append(durations, time.Since(start))
		}
	}

	// Compare per-round max/min across branches (each round samples all
	// three back-to-back, damping load noise).
	var maxRatio float64
	for round := 0; round < 5; round++ {
		min, max := durations[round*3], durations[round*3]
		for _, d := range durations[round*3 : round*3+3] {
			if d < min {
				min = d
			}
			if d > max {
				max = d
			}
		}
		ratio := float64(max) / float64(min)
		if ratio > maxRatio {
			maxRatio = ratio
		}
	}
	if maxRatio > 3.0 {
		t.Errorf("forgot-password branch timing not equivalent: max/min ratio %.2f across rounds (band 3.0); durations=%v",
			maxRatio, durations)
	}
	t.Logf("worst per-round max/min ratio: %.2f", maxRatio)
}

// secretsCompare wraps platform/secrets.ComparePassword for terse
// assertions in this file.
func secretsCompare(hashed, password string) bool {
	return secrets.ComparePassword(hashed, password) == nil
}
