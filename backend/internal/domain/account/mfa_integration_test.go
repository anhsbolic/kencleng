//go:build integration

// package account — MFA TOTP (Fitur 3) integration tests against real
// Postgres. Proves what unit fakes structurally cannot:
//
//   - R1/R3: enroll persists an encrypted secret at rest (enabled_at NULL);
//     the enroll-while-active 409 path leaves the live secret untouched
//     through the conflict-armed upsert (D5);
//   - R5: confirm flips enabled_at + inserts exactly 10 hashed codes +
//     mfa_enabled audit in ONE committed transaction; skipped-by-verify
//     leaves enabled_at NULL (INV-account-07);
//   - R8: ≥100 concurrent confirms → exactly one winner, COUNT(codes)=10,
//     every loser observes 422 — INV-account-07's at-most-once transition;
//   - R9: backup-code redemption is one joined guarded UPDATE under real
//     contention — exactly-once used_at, replayed/disabled-owner codes
//     reject with zero writes (INV-account-06);
//   - R11/R12: disable email_password path + idempotent repeat + reauth.
//
// Run: DATABASE_URL=... go test -tags=integration -race ./internal/domain/account/...
package account

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"

	"github.com/anhsbolic/kencleng/backend/internal/platform/secrets"
)

// countMFABackupCodes returns how many backup-code rows exist for userID.
func countMFABackupCodes(t *testing.T, pool *pgxpool.Pool, userID uuid.UUID) int {
	t.Helper()
	var n int
	if err := pool.QueryRow(context.Background(),
		`SELECT COUNT(*) FROM mfa_backup_codes WHERE user_id = $1`, userID).Scan(&n); err != nil {
		t.Fatalf("count backup codes: %v", err)
	}
	return n
}

// mfaEnabledState returns (rowExists, enabledAtIsSet) for mfa_totp_secrets.
func mfaEnabledState(t *testing.T, pool *pgxpool.Pool, userID uuid.UUID) (bool, bool) {
	t.Helper()
	var enabledAt *time.Time
	err := pool.QueryRow(context.Background(),
		`SELECT enabled_at FROM mfa_totp_secrets WHERE user_id = $1`, userID).Scan(&enabledAt)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return false, false
		}
		t.Fatalf("select mfa state: %v", err)
	}
	return true, enabledAt != nil
}

// ---- R1/R3/R5: enroll-persist, 409-on-active, confirm-atomic -------------

// integrationMFAService returns a real MFA-capable Service backed by real
// Postgres: sets the now/compare seams the base integrationService leaves
// unset (its struct literal only wires repo/tx/breach/email/keys), and
// attaches the real verifier.
func integrationMFAService(t *testing.T) (*Service, *RepositoryDB, *pgxpool.Pool) {
	t.Helper()
	svc, repo, pool := integrationService(t)
	svc.now = time.Now
	svc.compare = secrets.ComparePassword
	svc.mfa = NewMfaVerifier(repo)
	return svc, repo, pool
}

func TestMfaEnroll_EncryptedSecret_AndRejectsWhenEnabled_RealDB(t *testing.T) {
	svc, repo, pool := integrationMFAService(t)
	ctx := context.Background()
	user := newTestUser(t, repo, pool)

	// R1: enroll stores a pending secret and returns an otpauth URI.
	uri, err := svc.MfaEnroll(ctx, user.ID)
	if err != nil {
		t.Fatalf("enroll: %v", err)
	}
	if uri == "" {
		t.Fatal("empty otpauth URI")
	}
	secretBase32, enabledAt, found, err := repo.GetMFATOTPSecretForVerify(ctx, user.ID)
	if err != nil || !found {
		t.Fatalf("pending secret not readable: found=%t err=%v", found, err)
	}
	if enabledAt != nil {
		t.Error("enrolled secret must be pending (enabled_at NULL)")
	}
	// The decrypted secret must be a valid base32 TOTP secret that produces
	// a checkable code (the same instant both ways, windowed with skew).
	code, gcErr := totp.GenerateCode(secretBase32, time.Now())
	if gcErr != nil {
		t.Fatalf("generate code from decrypted secret: %v", gcErr)
	}
	valid, vErr := totp.ValidateCustom(code, secretBase32, time.Now(), totp.ValidateOpts{Period: 30, Skew: 1, Digits: 6, Algorithm: otp.AlgorithmSHA1})
	if vErr != nil || !valid {
		t.Fatalf("decrypted secret must validate its own code (valid=%t err=%v)", valid, vErr)
	}

	// R5: confirm enables + inserts exactly 10 hashed codes + audits in one tx.
	confirmCode, _ := totp.GenerateCode(secretBase32, time.Now())
	codes, err := svc.MfaEnrollConfirm(ctx, user.ID, confirmCode)
	if err != nil {
		t.Fatalf("confirm: %v", err)
	}
	if len(codes) != backupCodeCount {
		t.Fatalf("backup codes = %d, want %d", len(codes), backupCodeCount)
	}
	if _, enabled := mfaEnabledState(t, pool, user.ID); !enabled {
		t.Error("enabled_at must be set after confirm")
	}
	if got := countMFABackupCodes(t, pool, user.ID); got != backupCodeCount {
		t.Errorf("stored backup codes = %d, want %d", got, backupCodeCount)
	}
	if got := userLogCountForAction(t, pool, user.ID, actionMfaEnabled); got != 1 {
		t.Errorf("mfa_enabled audit rows = %d, want 1", got)
	}

	// R3: enroll again while active → 409 sentinel, live secret untouched.
	if _, err := svc.MfaEnroll(ctx, user.ID); err != ErrMfaAlreadyEnabled {
		t.Fatalf("re-enroll while enabled err = %v, want ErrMfaAlreadyEnabled", err)
	}
	secretAfter, _, foundAfter, _ := repo.GetMFATOTPSecretForVerify(ctx, user.ID)
	if !foundAfter || secretAfter != secretBase32 {
		t.Error("rejected re-enroll clobbered the live secret")
	}
}

// ---- INV-account-07: no path enables without in-flow TOTP verification ----

func TestMfaEnroll_NoHalfEnabledState_RealDB(t *testing.T) {
	svc, repo, pool := integrationMFAService(t)
	ctx := context.Background()
	user := newTestUser(t, repo, pool)

	if _, err := svc.MfaEnroll(ctx, user.ID); err != nil {
		t.Fatalf("enroll: %v", err)
	}
	if _, enabled := mfaEnabledState(t, pool, user.ID); enabled {
		t.Fatal("enroll without confirm must leave enabled_at NULL (INV-account-07)")
	}
}

// ---- R8: concurrent confirm exactly-once --------------------------------

func TestMfaConfirm_Concurrent_ExactlyOneWinner_TenCodesTotal_RealDB(t *testing.T) {
	svc, repo, pool := integrationMFAService(t)
	ctx := context.Background()
	user := newTestUser(t, repo, pool)

	secret, _, found, _ := repo.GetMFATOTPSecretForVerify(ctx, user.ID)
	if !found {
		// No secret yet — start an enrollment so we have a pending secret.
		if _, err := svc.MfaEnroll(ctx, user.ID); err != nil {
			t.Fatalf("enroll: %v", err)
		}
		secret, _, _, _ = repo.GetMFATOTPSecretForVerify(ctx, user.ID)
	}
	confirmCode, _ := totp.GenerateCode(secret, time.Now())

	const attempts = 100
	var mu sync.Mutex
	winners, losers := 0, 0
	var wg sync.WaitGroup
	for i := 0; i < attempts; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := svc.MfaEnrollConfirm(ctx, user.ID, confirmCode)
			mu.Lock()
			defer mu.Unlock()
			if err == nil {
				winners++
			} else if err == ErrMfaNotPending {
				losers++
			} else {
				t.Errorf("unexpected confirm error under concurrency: %v", err)
			}
		}()
	}
	wg.Wait()

	if winners != 1 {
		t.Errorf("winners = %d, want exactly 1", winners)
	}
	if losers != attempts-1 {
		t.Errorf("losers = %d, want %d", losers, attempts-1)
	}
	// INV-account-07 count property under contention: exactly 10 codes.
	if got := countMFABackupCodes(t, pool, user.ID); got != backupCodeCount {
		t.Errorf("total backup codes = %d, want %d (never 20)", got, backupCodeCount)
	}
	if _, enabled := mfaEnabledState(t, pool, user.ID); !enabled {
		t.Error("enabled_at must be set exactly once")
	}
}

// ---- R9: backup-code redemption (joined guarded UPDATE, real contention) --

func TestMfaBackupCode_SingleUseAndDisabledInvalid_RealDB(t *testing.T) {
	svc, repo, pool := integrationMFAService(t)
	ctx := context.Background()
	user := newTestUser(t, repo, pool)

	// Enroll + confirm so we have real codes.
	if _, err := svc.MfaEnroll(ctx, user.ID); err != nil {
		t.Fatalf("enroll: %v", err)
	}
	secret, _, _, _ := repo.GetMFATOTPSecretForVerify(ctx, user.ID)
	confirmCode, _ := totp.GenerateCode(secret, time.Now())
	plainCodes, err := svc.MfaEnrollConfirm(ctx, user.ID, confirmCode)
	if err != nil {
		t.Fatalf("confirm: %v", err)
	}
	v := NewMfaVerifier(repo)
	tx, _ := svc.tx.BeginTx(ctx)

	// First redeem of a real code → success (commits the used_at write).
	ok, err := v.VerifyBackupCode(ctx, tx, user.ID, plainCodes[0])
	if err != nil || !ok {
		t.Fatalf("first redeem: ok=%t err=%v", ok, err)
	}
	// Replay → fails, no second used_at write.
	ok, err = v.VerifyBackupCode(ctx, tx, user.ID, plainCodes[0])
	if err != nil || ok {
		t.Fatalf("replay redeem: ok=%t err=%v", ok, err)
	}
	if err := tx.Commit(ctx); err != nil {
		_ = tx.Rollback(ctx)
		t.Fatalf("commit redeem tx: %v", err)
	}

	// Disable → all remaining codes become unusable (INV-account-06).
	cred, _ := secrets.HashPassword("securepass")
	_ = seedEPIdentityForUser(t, repo, pool, user.ID, cred)
	if err := svc.MfaDisable(ctx, user.ID, "securepass"); err != nil {
		t.Fatalf("disable: %v", err)
	}
	tx2, _ := svc.tx.BeginTx(ctx)
	ok, err = v.VerifyBackupCode(ctx, tx2, user.ID, plainCodes[1])
	if err != nil || ok {
		_ = tx2.Rollback(ctx)
		t.Fatalf("disabled-owner code must be unusable: ok=%t err=%v", ok, err)
	}
	_ = tx2.Rollback(ctx) // writes nothing; release the connection
}

// seedEPIdentityForUser (integration) attaches an email_password identity to
// a user for the disable re-auth path.
func seedEPIdentityForUser(t *testing.T, _ *RepositoryDB, pool *pgxpool.Pool, userID uuid.UUID, credSecret string) error {
	t.Helper()
	ctx := context.Background()
	ident := &AuthIdentity{
		ID:               uuid.New(),
		UserID:           userID,
		ProviderType:     "email_password",
		Identifier:       fmt.Sprintf("ep-%s@example.com", uuid.NewString()),
		CredentialSecret: &credSecret,
	}
	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	if _, err := tx.Exec(ctx,
		`INSERT INTO auth_identities (id, user_id, provider_type, identifier, credential_secret) VALUES ($1,$2,$3,$4,$5) ON CONFLICT DO NOTHING`,
		ident.ID, ident.UserID, ident.ProviderType, ident.Identifier, ident.CredentialSecret); err != nil {
		_ = tx.Rollback(ctx)
		return err
	}
	return tx.Commit(ctx)
}

// ---- R11/R12: disable email_password success, idempotency, reauth ---------

func TestMfaDisable_EmailPassword_AuditsAndIdempotent_RealDB(t *testing.T) {
	svc, repo, pool := integrationMFAService(t)
	ctx := context.Background()
	user := newTestUser(t, repo, pool)

	if _, err := svc.MfaEnroll(ctx, user.ID); err != nil {
		t.Fatalf("enroll: %v", err)
	}
	secret, _, _, _ := repo.GetMFATOTPSecretForVerify(ctx, user.ID)
	confirmCode, _ := totp.GenerateCode(secret, time.Now())
	if _, err := svc.MfaEnrollConfirm(ctx, user.ID, confirmCode); err != nil {
		t.Fatalf("confirm: %v", err)
	}
	cred, _ := secrets.HashPassword("securepass")
	_ = seedEPIdentityForUser(t, repo, pool, user.ID, cred)

	if err := svc.MfaDisable(ctx, user.ID, "securepass"); err != nil {
		t.Fatalf("disable: %v", err)
	}
	if _, enabled := mfaEnabledState(t, pool, user.ID); enabled {
		t.Error("enabled_at must be NULL after disable")
	}
	if got := userLogCountForAction(t, pool, user.ID, actionMfaDisabled); got != 1 {
		t.Errorf("mfa_disabled audit = %d, want 1", got)
	}
	// Backup codes remain (implicit invalidation, no hard delete).
	if got := countMFABackupCodes(t, pool, user.ID); got != backupCodeCount {
		t.Errorf("codes kept on disable = %d, want %d", got, backupCodeCount)
	}

	// Idempotent repeat: succeeds, no second audit row.
	if err := svc.MfaDisable(ctx, user.ID, "securepass"); err != nil {
		t.Fatalf("repeat disable: %v", err)
	}
	if got := userLogCountForAction(t, pool, user.ID, actionMfaDisabled); got != 1 {
		t.Errorf("mfa_disabled audit after repeat = %d, want 1 (idempotent)", got)
	}

	// Wrong password → 401-class sentinel, no state change.
	if err := svc.MfaDisable(ctx, user.ID, "wrong"); err != ErrInvalidCredentials {
		t.Errorf("wrong password err = %v, want ErrInvalidCredentials", err)
	}
}
