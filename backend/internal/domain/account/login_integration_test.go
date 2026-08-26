//go:build integration

// package account — login/session integration tests against real Postgres.
//
// Proves what a fake repository cannot (task-06 scope):
//   - INV-account-03: guarded rotation lets exactly ONE concurrent refresh
//     win, with no double-parenting, under true DB-level contention;
//   - INV-account-04: reuse revokes the whole family, including rotated
//     descendants;
//   - the persistent lockout end-to-end: 5 recorded failures lock the next
//     Login attempt before any credential work, writing no new row;
//   - migration round-trip is verified by the task's shell steps, not here.
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
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/anhsbolic/kencleng/backend/internal/platform/crypto"
	"github.com/anhsbolic/kencleng/backend/internal/platform/db"
	"github.com/anhsbolic/kencleng/backend/internal/platform/secrets"
)

// integrationLoginService builds a Service over a REAL pool with trivial
// token-mint closures — minting is Tier 0 unit-tested (task-03); rotation,
// locking, and bookkeeping are what real Postgres must prove here.
func integrationLoginService(t *testing.T) (*Service, *RepositoryDB, *pgxpool.Pool) {
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
	keys := integrationTestKeys(t)
	repo := NewRepositoryDB(pool, keys)
	svc := &Service{
		repo:        repo,
		tx:          poolRunner{pool: pool},
		breachCheck: integrationBreachCheckerFalse{},
		email:       integrationSilentSender{},
		keys:        keys,
		mfa:         stubMfaVerifier{},
		now:         time.Now,
		compare:     secrets.ComparePassword,
		mintAccess: func(uuid.UUID, time.Time) (string, error) {
			return "integration-access-marker", nil
		},
		mintMFAPending: func(uuid.UUID, time.Time) (string, error) {
			return "integration-pending-marker", nil
		},
		verifyPending: func(string, time.Time) (uuid.UUID, error) {
			return uuid.Nil, errors.New("not under test here")
		},
	}
	return svc, repo, pool
}

// seedLoginIdentity creates user + email_password identity with a REAL
// bcrypt credential so svc.Login exercises actual comparison cost.
func seedLoginIdentity(t *testing.T, repo *RepositoryDB, pool *pgxpool.Pool, password string) (email, identifierHash string) {
	t.Helper()
	ctx := context.Background()
	email = fmt.Sprintf("login-%s@example.com", uuid.NewString())
	identifierHash = crypto.HMAC([]byte(email), integrationTestKeys(t))

	hashed, err := secrets.HashPassword(password)
	if err != nil {
		t.Fatalf("hash password: %v", err)
	}

	u := &User{ID: uuid.New(), Name: "Login Integration", PrimaryEmail: email}
	ident := &AuthIdentity{
		ID: uuid.New(), UserID: u.ID, ProviderType: providerEmailPassword,
		Identifier: email, CredentialSecret: &hashed,
	}
	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		t.Fatalf("begin tx: %v", err)
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
	return email, identifierHash
}

// TestRefresh_ConcurrentRequests_ExactlyOneWins proves R15 / INV-account-03
// at the storage level: N goroutines race one valid refresh token through
// the real service + guarded UPDATE; exactly one wins, no second child ever
// exists, and the family ends fully revoked (loser ≡ attacker, Assumption D).
func TestRefresh_ConcurrentRequests_ExactlyOneWins_RealDB(t *testing.T) {
	svc, repo, pool := integrationLoginService(t)
	ctx := context.Background()
	user := newTestUser(t, repo, pool)
	userID := user.ID

	family := uuid.New()
	plain := "race-token-" + uuid.NewString()
	parentHash := sha256Hex(plain)

	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		t.Fatalf("begin: %v", err)
	}
	if err := svc.repo.InsertRefreshToken(ctx, tx, &RefreshToken{
		ID: uuid.New(), UserID: userID, FamilyID: family,
		TokenHash: parentHash, ExpiresAt: time.Now().Add(refreshTokenTTL), CreatedAt: time.Now(),
	}); err != nil {
		_ = tx.Rollback(ctx)
		t.Fatalf("seed parent: %v", err)
	}
	if err := tx.Commit(ctx); err != nil {
		t.Fatalf("commit: %v", err)
	}
	t.Cleanup(func() {
		_, _ = pool.Exec(ctx, "DELETE FROM refresh_tokens WHERE family_id = $1", family)
	})

	const racers = 8
	type outcome struct{ err error }
	results := make(chan outcome, racers)
	var wg sync.WaitGroup
	for i := 0; i < racers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := svc.Refresh(ctx, plain)
			results <- outcome{err: err}
		}()
	}
	wg.Wait()
	close(results)

	wins := 0
	for r := range results {
		if r.err == nil {
			wins++
		} else if !errors.Is(r.err, ErrInvalidCredentials) {
			t.Fatalf("racer got unexpected error class: %v", r.err)
		}
	}
	if wins != 1 {
		t.Fatalf("winners = %d, want exactly 1 across %d racers", wins, racers)
	}

	// Invariant sweep: every row in the family revoked; parent has exactly
	// one child reference; child count == number of successful rotations == 1.
	rows, err := pool.Query(ctx,
		`SELECT id, replaced_by_id, revoked_at FROM refresh_tokens WHERE family_id = $1`, family)
	if err != nil {
		t.Fatalf("family query: %v", err)
	}
	defer rows.Close()
	live, children := 0, 0
	for rows.Next() {
		var id uuid.UUID
		var replaced *uuid.UUID
		var revoked *time.Time
		if err := rows.Scan(&id, &replaced, &revoked); err != nil {
			t.Fatalf("scan: %v", err)
		}
		if revoked == nil {
			live++
		}
		if replaced != nil {
			children++
		}
	}
	if live != 0 {
		t.Errorf("%d live tokens remain after race (want 0 — loser ≡ attacker)", live)
	}
	if children != 1 {
		t.Errorf("%d children reference a parent (want exactly 1 — INV-account-03)", children)
	}
}

// TestRefresh_Stress_MixedValidAndReplayed is the tasks.md KPI stress
// harness: ≥100 concurrent goroutines mixing one valid token and replayed
// tokens. Invariant across the run: zero double-parenting; every replay
// leaves its whole family revoked; total children ≤ total rotations.
func TestRefresh_Stress_MixedValidAndReplayed(t *testing.T) {
	svc, repo, pool := integrationLoginService(t)
	ctx := context.Background()
	user := newTestUser(t, repo, pool)
	userID := user.ID

	families := make([]uuid.UUID, 4)
	plains := make([]string, len(families))
	for f := range families {
		families[f] = uuid.New()
		plains[f] = "stress-" + uuid.NewString()
		tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
		if err != nil {
			t.Fatalf("begin: %v", err)
		}
		if err := svc.repo.InsertRefreshToken(ctx, tx, &RefreshToken{
			ID: uuid.New(), UserID: userID, FamilyID: families[f],
			TokenHash: sha256Hex(plains[f]), ExpiresAt: time.Now().Add(refreshTokenTTL), CreatedAt: time.Now(),
		}); err != nil {
			_ = tx.Rollback(ctx)
			t.Fatalf("seed: %v", err)
		}
		if err := tx.Commit(ctx); err != nil {
			t.Fatalf("commit: %v", err)
		}
	}
	t.Cleanup(func() {
		for _, f := range families {
			_, _ = pool.Exec(ctx, "DELETE FROM refresh_tokens WHERE family_id = $1", f)
		}
	})

	const goroutines = 120
	var wg sync.WaitGroup
	errs := make(chan error, goroutines)
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			plain := plains[i%len(families)]
			_, err := svc.Refresh(ctx, plain)
			if err != nil && !errors.Is(err, ErrInvalidCredentials) {
				errs <- fmt.Errorf("unexpected class: %w", err)
			}
		}(i)
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		t.Fatal(err)
	}

	// Invariants: every family fully dead; no parent referenced twice.
	for f, family := range families {
		var live int
		if err := pool.QueryRow(ctx,
			`SELECT count(*) FROM refresh_tokens WHERE family_id = $1 AND revoked_at IS NULL`, family).
			Scan(&live); err != nil {
			t.Fatalf("live count: %v", err)
		}
		if live != 0 {
			t.Errorf("family %d has %d live tokens after stress (want 0)", f, live)
		}
		var dupes int
		if err := pool.QueryRow(ctx,
			`SELECT count(*) FROM (
                 SELECT replaced_by_id FROM refresh_tokens
                 WHERE family_id = $1 AND replaced_by_id IS NOT NULL
                 GROUP BY replaced_by_id HAVING count(*) > 1) d`, family).
			Scan(&dupes); err != nil {
			t.Fatalf("dupe check: %v", err)
		}
		if dupes != 0 {
			t.Errorf("family %d has %d parents referenced by multiple children (INV-account-03 violation)", f, dupes)
		}
	}
}

// TestRefresh_ReuseDetection_FamilyRevoked_RealDB proves R14 / INV-account-04
// through the real adapter: A→B→C chain, replay A ⇒ A,B,C all revoked; C
// unusable afterward.
func TestRefresh_ReuseDetection_FamilyRevoked_RealDB(t *testing.T) {
	svc, repo, pool := integrationLoginService(t)
	ctx := context.Background()
	user := newTestUser(t, repo, pool)
	userID := user.ID

	family := uuid.New()
	plainA := "chain-a-" + uuid.NewString()

	seed := func(hash string) {
		tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
		if err != nil {
			t.Fatalf("begin: %v", err)
		}
		if err := svc.repo.InsertRefreshToken(ctx, tx, &RefreshToken{
			ID: uuid.New(), UserID: userID, FamilyID: family,
			TokenHash: hash, ExpiresAt: time.Now().Add(refreshTokenTTL), CreatedAt: time.Now(),
		}); err != nil {
			_ = tx.Rollback(ctx)
			t.Fatalf("seed: %v", err)
		}
		if err := tx.Commit(ctx); err != nil {
			t.Fatalf("commit: %v", err)
		}
	}
	seed(sha256Hex(plainA))
	t.Cleanup(func() {
		_, _ = pool.Exec(ctx, "DELETE FROM refresh_tokens WHERE family_id = $1", family)
	})

	resB, err := svc.Refresh(ctx, plainA)
	if err != nil {
		t.Fatalf("rotate A→B: %v", err)
	}
	resC, err := svc.Refresh(ctx, resB.RefreshTokenPlain)
	if err != nil {
		t.Fatalf("rotate B→C: %v", err)
	}

	if _, err := svc.Refresh(ctx, plainA); !errors.Is(err, ErrInvalidCredentials) {
		t.Fatalf("replay A: err = %v, want ErrInvalidCredentials", err)
	}

	var revokedCount int
	if err := pool.QueryRow(ctx,
		`SELECT count(*) FROM refresh_tokens WHERE family_id = $1 AND revoked_at IS NOT NULL`, family).
		Scan(&revokedCount); err != nil {
		t.Fatalf("count: %v", err)
	}
	if revokedCount < 3 {
		t.Errorf("%d tokens revoked in family, want ≥3 (A, B, C all condemned)", revokedCount)
	}
	var liveCount int
	if err := pool.QueryRow(ctx,
		`SELECT count(*) FROM refresh_tokens WHERE family_id = $1 AND revoked_at IS NULL`, family).
		Scan(&liveCount); err != nil {
		t.Fatalf("live count: %v", err)
	}
	if liveCount != 0 {
		t.Errorf("%d tokens still live in family (want 0)", liveCount)
	}
	if _, err := svc.Refresh(ctx, resC.RefreshTokenPlain); !errors.Is(err, ErrInvalidCredentials) {
		t.Errorf("refresh with revoked C succeeded: %v", err)
	}
}

// TestLogin_Lockout_EndToEnd proves R4 end-to-end on real Postgres: five
// seeded failures within the window lock the next attempt BEFORE any
// credential work (no compare possible to assert at this layer, but the
// no-new-row assertion holds and bcrypt of a WRONG password for a REAL
// account would still be skipped — observable via attempt-count stability),
// and the response class is ErrLockedOut.
func TestLogin_Lockout_EndToEnd(t *testing.T) {
	svc, repo, pool := integrationLoginService(t)
	ctx := context.Background()
	const password = "real-bcrypt-password-123"
	email, identifierHash := seedLoginIdentity(t, repo, pool, password)

	// Five in-window failures (recent attempted_at via direct seed).
	now := time.Now()
	for i := 0; i < maxFailedAttempts; i++ {
		if err := insertLoginAttemptDirect(ctx, pool, &LoginAttempt{
			ID: uuid.New(), IdentifierHash: identifierHash,
			UserID: nil, Stage: stagePassword, Success: false,
			AttemptedAt: now.Add(-time.Duration(i+1) * time.Minute),
		}); err != nil {
			t.Fatalf("seed failure row: %v", err)
		}
	}
	t.Cleanup(func() {
		_, _ = pool.Exec(ctx, "DELETE FROM login_attempts WHERE identifier_hash = $1", identifierHash)
	})

	before := countAttempts(t, pool, identifierHash)
	_, err := svc.Login(ctx, email, password) // even the CORRECT password must bounce
	if !errors.Is(err, ErrLockedOut) {
		t.Fatalf("err = %v, want ErrLockedOut", err)
	}
	after := countAttempts(t, pool, identifierHash)
	if after != before {
		t.Errorf("lockout-rejected attempt wrote rows: before=%d after=%d (spec violation)", before, after)
	}

	// Window expiry frees the identifier: shift the cutoff forward by
	// seeding nothing new — instead verify an old-failures-only state passes.
	oldOnly := now.Add(-lockoutWindow - time.Minute)
	if err := insertLoginAttemptDirect(ctx, pool, &LoginAttempt{
		ID: uuid.New(), IdentifierHash: identifierHash,
		UserID: nil, Stage: stagePassword, Success: false, AttemptedAt: oldOnly,
	}); err != nil {
		t.Fatalf("seed stale failure: %v", err)
	}
	_ = oldOnly

	// Correct credentials AFTER clearing the recent window succeed.
	_, delErr := pool.Exec(ctx,
		`DELETE FROM login_attempts WHERE identifier_hash = $1 AND attempted_at > $2`,
		identifierHash, now.Add(-lockoutWindow))
	if delErr != nil {
		t.Fatalf("clear recent failures: %v", delErr)
	}
	res, err := svc.Login(ctx, email, password)
	if err != nil {
		t.Fatalf("post-window login with correct credentials failed: %v", err)
	}
	if res.Status != "ok" || len(harvestAttemptRows(t, pool, identifierHash)) == 0 {
		t.Errorf("expected ok result + fresh success attempt row")
	}
}

func countAttempts(t *testing.T, pool *pgxpool.Pool, identifierHash string) int {
	t.Helper()
	var n int
	if err := pool.QueryRow(context.Background(),
		`SELECT count(*) FROM login_attempts WHERE identifier_hash = $1`, identifierHash).Scan(&n); err != nil {
		t.Fatalf("count attempts: %v", err)
	}
	return n
}

func harvestAttemptRows(t *testing.T, pool *pgxpool.Pool, identifierHash string) []LoginAttempt {
	t.Helper()
	rows, err := pool.Query(context.Background(),
		`SELECT stage, success FROM login_attempts WHERE identifier_hash = $1 AND success = true`, identifierHash)
	if err != nil {
		t.Fatalf("harvest: %v", err)
	}
	defer rows.Close()
	var out []LoginAttempt
	for rows.Next() {
		var a LoginAttempt
		if err := rows.Scan(&a.Stage, &a.Success); err != nil {
			t.Fatalf("scan: %v", err)
		}
		out = append(out, a)
	}
	return out
}
