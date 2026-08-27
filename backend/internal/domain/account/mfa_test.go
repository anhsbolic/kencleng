package account

import (
	"bytes"
	"context"
	"log"
	"net/url"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"

	"github.com/anhsbolic/kencleng/backend/internal/platform/crypto"
	"github.com/anhsbolic/kencleng/backend/internal/platform/secrets"
)

// newMFAService builds a Service wired to the fake repo with a pinned clock,
// real bcrypt comparison, and a seeded login view (so MfaEnroll can read the
// otpauth label email). Returns the service, the repo, and keys.
func newMFAService(t *testing.T) (*Service, *fakeRepo, *crypto.Keys) {
	t.Helper()
	repo := newFakeRepo()
	repo.seedView("alice@example.com")
	keys := &crypto.Keys{
		EncryptionKey: make([]byte, 32),
		HMACKey:       make([]byte, 32),
	}
	svc := &Service{
		repo:        repo,
		tx:          fakeTxRunner{},
		breachCheck: &fakeBreachChecker{},
		email:       &captureSender{},
		keys:        keys,
		now:         func() time.Time { return baseTime() },
		compare:     secrets.ComparePassword,
	}
	return svc, repo, keys
}

// codeFor computes the TOTP code for a secret at the same instant newMFAService
// pins its clock to, so ValidateCustom passes deterministically.
func codeFor(secret string) string {
	code, err := totp.GenerateCode(secret, baseTime())
	if err != nil {
		panic("codeFor: " + err.Error())
	}
	return code
}

// ---- R1: enroll happy path ----------------------------------------------

func TestMfaEnroll_StoresEncryptedSecret_ReturnsOtpauthURI(t *testing.T) {
	svc, repo, keys := newMFAService(t)
	uid := uuid.New()

	uri, err := svc.MfaEnroll(context.Background(), uid)
	if err != nil {
		t.Fatalf("MfaEnroll: %v", err)
	}
	if uri == "" || len(uri) < 20 {
		t.Fatalf("otpauth URI too short: %q", uri)
	}
	u, err := url.Parse(uri)
	if err != nil || u.Scheme != "otpauth" || u.Host != "totp" {
		t.Fatalf("not a valid otpauth://totp URI: %q (err=%v)", uri, err)
	}
	q := u.Query()
	if q.Get("issuer") != otpauthIssuer {
		t.Errorf("issuer = %q, want %q", q.Get("issuer"), otpauthIssuer)
	}
	if q.Get("secret") == "" {
		t.Error("otpauth URI missing secret")
	}

	// enabled_at must stay NULL (pending).
	if repo.mfaEnabledAt[uid] != nil {
		t.Error("enrollment must be pending (enabled_at nil) after enroll")
	}

	// The persisted ciphertext must decrypt to the same base32 the URI carries.
	ct, ok := repo.upsertedMFASecrets[uid]
	if !ok {
		t.Fatalf("pending secret not persisted for user")
	}
	plain, err := crypto.Decrypt(ct, keys)
	if err != nil {
		t.Fatalf("decrypt stored secret: %v", err)
	}
	if string(plain) != q.Get("secret") {
		t.Errorf("stored secret %q != URI secret %q", plain, q.Get("secret"))
	}
}

// ---- R2: enroll restart overwrites pending, stays NULL -------------------

func TestMfaEnroll_RestartOverwritesPendingSecret(t *testing.T) {
	svc, repo, _ := newMFAService(t)
	uid := uuid.New()

	uri1, err := svc.MfaEnroll(context.Background(), uid)
	if err != nil {
		t.Fatalf("first enroll: %v", err)
	}
	uri2, err := svc.MfaEnroll(context.Background(), uid)
	if err != nil {
		t.Fatalf("second enroll: %v", err)
	}
	if uri1 == uri2 {
		t.Error("restart must issue a fresh secret (URIs identical)")
	}
	if repo.mfaEnabledAt[uid] != nil {
		t.Error("second enroll must keep enabled_at NULL")
	}
}

// ---- R3: enroll rejected when MFA active (409 sentinel) ------------------

func TestMfaEnroll_RejectsWhenAlreadyEnabled(t *testing.T) {
	svc, repo, _ := newMFAService(t)
	uid := uuid.New()
	repo.seedMFAEnabled(uid, "JBSWY3DPEHPK3PXP")

	_, err := svc.MfaEnroll(context.Background(), uid)
	if err != ErrMfaAlreadyEnabled {
		t.Fatalf("err = %v, want ErrMfaAlreadyEnabled", err)
	}

	// Live secret must never be overwritten: the upsert must not have written
	// anything for this enabled user.
	if _, wrote := repo.upsertedMFASecrets[uid]; wrote {
		t.Error("enabled secret was written/overwritten by rejected enroll")
	}
}

// ---- R4: write-time guard (upsert returns blocked on enabled) ------------

func TestMfaEnroll_ConcurrentWithEnable_NeverOverwritesLiveSecret(t *testing.T) {
	// Unit-level probe of the guard: with MFA already enabled, the upsert
	// returns inserted=false and MfaEnroll surfaces ErrMfaAlreadyEnabled
	// without touching state. (True interleaving contention lives in the
	// integration race suite, task-06.)
	svc, repo, _ := newMFAService(t)
	uid := uuid.New()
	repo.seedMFAEnabled(uid, "JBSWY3DPEHPK3PXP")

	_, err := svc.MfaEnroll(context.Background(), uid)
	if err != ErrMfaAlreadyEnabled {
		t.Fatalf("err = %v, want ErrMfaAlreadyEnabled", err)
	}
	if _, wrote := repo.upsertedMFASecrets[uid]; wrote {
		t.Error("upsert wrote a secret despite MFA enabled (live secret overwritten)")
	}
}

// ---- R5: confirm enables + issues 10 codes + audits -----------------------

func TestMfaConfirm_EnablesAndIssuesTenCodes_Audits(t *testing.T) {
	svc, repo, _ := newMFAService(t)
	uid := uuid.New()
	repo.seedMFAEnrollment(uid, "JBSWY3DPEHPK3PXP")

	codes, err := svc.MfaEnrollConfirm(context.Background(), uid, codeFor("JBSWY3DPEHPK3PXP"))
	if err != nil {
		t.Fatalf("MfaEnrollConfirm: %v", err)
	}
	if len(codes) != backupCodeCount {
		t.Fatalf("backup codes = %d, want %d", len(codes), backupCodeCount)
	}
	if repo.mfaEnabledAt[uid] == nil {
		t.Error("enabled_at must be set after confirm")
	}
	if len(repo.insertedMFACodes) != backupCodeCount {
		t.Fatalf("inserted backup codes = %d, want %d", len(repo.insertedMFACodes), backupCodeCount)
	}
	// Codes must be stored as hashes, never plaintext.
	storedHashes := map[string]bool{}
	for _, c := range repo.insertedMFACodes {
		if c.CodeHash == "" {
			t.Error("backup code stored without a hash")
		}
		storedHashes[c.CodeHash] = true
	}
	for _, p := range codes {
		if storedHashes[sha256Hex(normalizeBackupCode(p))] == false {
			t.Errorf("plaintext code %q not represented by a stored hash", p)
		}
		if p == "" {
			t.Error("empty backup code returned")
		}
	}
	// Audit entry.
	if got := repo.insertedUserLogs; len(got) != 1 || got[0].ActionType != actionMfaEnabled {
		t.Fatalf("mfa_enabled audit not written exactly once: %+v", got)
	}
}

// INV-account-07: no code path may enable without a verified TOTP in-flow.
func TestMfaEnroll_NoHalfEnabledState(t *testing.T) {
	svc, repo, _ := newMFAService(t)
	uid := uuid.New()

	// Enroll alone (no confirm) must leave enabled_at NULL.
	_, err := svc.MfaEnroll(context.Background(), uid)
	if err != nil {
		t.Fatalf("MfaEnroll: %v", err)
	}
	if repo.mfaEnabledAt[uid] != nil {
		t.Error("enroll without confirm must keep enabled_at NULL (no half-enabled state)")
	}
}

// ---- R6: wrong code preserves the pending secret ---------------------------

func TestMfaConfirm_WrongCode_PreservesPendingSecret(t *testing.T) {
	svc, repo, _ := newMFAService(t)
	uid := uuid.New()
	repo.seedMFAEnrollment(uid, "JBSWY3DPEHPK3PXP")

	_, err := svc.MfaEnrollConfirm(context.Background(), uid, "000000")
	if err != ErrInvalidTOTPCode {
		t.Fatalf("err = %v, want ErrInvalidTOTPCode", err)
	}
	if repo.mfaEnabledAt[uid] != nil {
		t.Error("wrong code must not enable MFA")
	}
	if len(repo.insertedMFACodes) != 0 {
		t.Error("wrong code must not insert backup codes")
	}

	// Retry with the correct code succeeds WITHOUT re-enrolling.
	codes, err := svc.MfaEnrollConfirm(context.Background(), uid, codeFor("JBSWY3DPEHPK3PXP"))
	if err != nil {
		t.Fatalf("retry after wrong code: %v", err)
	}
	if len(codes) != backupCodeCount {
		t.Errorf("retry codes = %d, want %d", len(codes), backupCodeCount)
	}
}

// ---- R7: no-pending confirm is indistinguishable from wrong code ---------

func TestMfaConfirm_NoPending_IndistinguishableFromWrongCode(t *testing.T) {
	svc, _, _ := newMFAService(t)
	uid := uuid.New()
	// No pending enrollment for this user.
	_, err := svc.MfaEnrollConfirm(context.Background(), uid, codeFor("JBSWY3DPEHPK3PXP"))
	if err != ErrInvalidTOTPCode && err != ErrMfaNotPending {
		// Both collapse to the same 422 on the wire; accept either internal
		// sentinel but reject everything else.
		t.Fatalf("err = %v, want a 422-mapped sentinel (wrong-code-class)", err)
	}
}

// ---- R8: concurrent confirm — exactly one winner, exactly 10 codes --------

func TestMfaConfirm_Concurrent_ExactlyOneWinner_TenCodesTotal(t *testing.T) {
	if testing.Short() {
		t.Skip("concurrency stress skipped in -short")
	}
	svc, repo, _ := newMFAService(t)
	uid := uuid.New()
	repo.seedMFAEnrollment(uid, "JBSWY3DPEHPK3PXP")
	code := codeFor("JBSWY3DPEHPK3PXP")

	const attempts = 100
	winners := 0
	var mu sync.Mutex
	other := 0
	var wg sync.WaitGroup
	for i := 0; i < attempts; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			codes, err := svc.MfaEnrollConfirm(context.Background(), uid, code)
			mu.Lock()
			defer mu.Unlock()
			if err == nil {
				winners++
				if len(codes) != backupCodeCount {
					t.Errorf("winner returned %d codes, want %d", len(codes), backupCodeCount)
				}
			} else if err == ErrMfaNotPending || err == ErrInvalidTOTPCode {
				other++
			} else {
				t.Errorf("unexpected confirm error: %v", err)
			}
		}()
	}
	wg.Wait()

	if winners != 1 {
		t.Errorf("winners = %d, want exactly 1", winners)
	}
	if other != attempts-1 {
		t.Errorf("losers = %d, want %d", other, attempts-1)
	}
	// INV-account-07 count property: exactly 10 codes, never 20.
	if len(repo.insertedMFACodes) != backupCodeCount {
		t.Errorf("total inserted codes = %d, want %d", len(repo.insertedMFACodes), backupCodeCount)
	}
}

// ---- R11: disable email_password success + idempotency ---------------------

func TestMfaDisable_Success_EmailPassword_Audits(t *testing.T) {
	svc, repo, _ := newMFAService(t)
	uid := uuid.New()
	repo.seedMFAEnabled(uid, "JBSWY3DPEHPK3PXP")
	credHash, _ := secrets.HashPassword("correct-horse-battery")
	repo.seedEPIdentityForUser(uid, &credHash)

	err := svc.MfaDisable(context.Background(), uid, "correct-horse-battery")
	if err != nil {
		t.Fatalf("MfaDisable: %v", err)
	}
	if repo.mfaEnabledAt[uid] != nil {
		t.Error("enabled_at must be NULL after disable")
	}
	if len(repo.insertedUserLogs) != 1 || repo.insertedUserLogs[0].ActionType != actionMfaDisabled {
		t.Fatalf("mfa_disabled audit not written exactly once: %+v", repo.insertedUserLogs)
	}
}

func TestMfaDisable_RepeatAfterDisable_Idempotent(t *testing.T) {
	svc, repo, _ := newMFAService(t)
	uid := uuid.New()
	repo.seedMFAEnabled(uid, "JBSWY3DPEHPK3PXP")
	credHash, _ := secrets.HashPassword("correct-horse-battery")
	repo.seedEPIdentityForUser(uid, &credHash)

	if err := svc.MfaDisable(context.Background(), uid, "correct-horse-battery"); err != nil {
		t.Fatalf("first disable: %v", err)
	}
	// Repeat — still succeeds, and writes no duplicate audit row (R11).
	if err := svc.MfaDisable(context.Background(), uid, "correct-horse-battery"); err != nil {
		t.Fatalf("repeat disable must stay idempotent-success: %v", err)
	}
	var mfaDisabledLogs int
	for _, l := range repo.insertedUserLogs {
		if l.ActionType == actionMfaDisabled {
			mfaDisabledLogs++
		}
	}
	if mfaDisabledLogs != 1 {
		t.Errorf("mfa_disabled audit rows = %d, want 1 (no duplicate on idempotent repeat)", mfaDisabledLogs)
	}
}

// ---- R12: disable re-auth failures (email_password) ------------------------

func TestMfaDisable_RequiresReauth_EmailPassword(t *testing.T) {
	svc, repo, _ := newMFAService(t)
	uid := uuid.New()
	repo.seedMFAEnabled(uid, "JBSWY3DPEHPK3PXP")
	credHash, _ := secrets.HashPassword("correct-horse-battery")
	repo.seedEPIdentityForUser(uid, &credHash)

	// Wrong password → ErrInvalidCredentials, no state change.
	err := svc.MfaDisable(context.Background(), uid, "wrong-password")
	if err != ErrInvalidCredentials {
		t.Fatalf("wrong password err = %v, want ErrInvalidCredentials", err)
	}
	if repo.mfaEnabledAt[uid] == nil {
		t.Error("wrong password must not disable MFA")
	}
	repo.insertedUserLogs = nil
	// Empty password → ErrValidation (422 required).
	err = svc.MfaDisable(context.Background(), uid, "")
	if err != ErrValidation {
		t.Fatalf("empty password err = %v, want ErrValidation", err)
	}
	if len(repo.insertedUserLogs) != 0 {
		t.Error("failed disable must not audit")
	}
}

// ---- R13: disable google-only (marker consumed at handler; service trusts)
// (marker mechanics are transport-side; here we prove the service path) ------

func TestMfaDisable_GoogleOnlyPath_Succeeds(t *testing.T) {
	svc, repo, _ := newMFAService(t)
	uid := uuid.New()
	repo.seedMFAEnabled(uid, "JBSWY3DPEHPK3PXP")
	// Google-only: no email_password identity seeded.

	err := svc.MfaDisable(context.Background(), uid, "")
	if err != nil {
		t.Fatalf("google-only disable (marker precondition met) must succeed: %v", err)
	}
	if repo.mfaEnabledAt[uid] != nil {
		t.Error("enabled_at must be NULL after google-only disable")
	}
}

// ---- R14: server-side provider detection -----------------------------------

func TestMfaDisable_BranchSelection_ServerSide(t *testing.T) {
	svc, repo, _ := newMFAService(t)

	// Google-only (no email_password) → marker required.
	gUid := uuid.New()
	goog, err := svc.MfaDisableReauthRequired(context.Background(), gUid)
	if err != nil {
		t.Fatalf("reauth check: %v", err)
	}
	if !goog {
		t.Error("user with no email_password identity must be Google-only (marker required)")
	}

	// Has email_password → password path.
	epUid := uuid.New()
	repo.seedEPIdentityForUser(epUid, nil)
	goog2, err := svc.MfaDisableReauthRequired(context.Background(), epUid)
	if err != nil {
		t.Fatalf("reauth check 2: %v", err)
	}
	if goog2 {
		t.Error("user with email_password identity must be password path, not Google-only")
	}
}

// ---- R10: /auth/login/mfa works end-to-end with the real verifier ----------

func TestLoginMfa_WithRealVerifier_CompletesAndFails(t *testing.T) {
	h := newLoginHarness(t)
	h.svc.mfa = NewMfaVerifier(h.repo)
	h.repo.view = &LoginUserView{Name: "Real Verifier User"}
	uid := h.verifyPendingUser
	secret := "JBSWY3DPEHPK3PXP"
	h.repo.seedMFAEnabled(uid, secret)
	ctx := context.Background()

	// Valid TOTP computes against the decrypted secret and completes login.
	validCode, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		t.Fatalf("generate valid code: %v", err)
	}
	res, err := h.svc.LoginMfa(ctx, "pending-token", validCode, "")
	if err != nil {
		t.Fatalf("valid TOTP login: %v", err)
	}
	if res.Status != "ok" || res.AccessToken == "" {
		t.Fatalf("completion shape wrong: %+v", res)
	}
	if succ := h.repo.attemptsBySuccess(true); len(succ) != 1 || succ[0].Stage != stageMFA {
		t.Fatalf("expected one mfa success row, got %+v", h.repo.attempts)
	}

	// Invalid TOTP (a code from a different window, > skew) fails and records
	// a failed mfa attempt — lockout bookkeeping unchanged.
	h.repo.attempts = nil
	stale, err := totp.GenerateCode(secret, time.Now().Add(-5*time.Minute))
	if err != nil {
		t.Fatalf("generate stale code: %v", err)
	}
	_, err = h.svc.LoginMfa(ctx, "pending-token", stale, "")
	if err != ErrInvalidCredentials {
		t.Fatalf("invalid TOTP err = %v, want ErrInvalidCredentials", err)
	}
	if fails := h.repo.attemptsBySuccess(false); len(fails) != 1 || fails[0].Stage != stageMFA {
		t.Fatalf("expected one failed mfa row, got %+v", h.repo.attempts)
	}
}

// Backup-code redemption through the real verifier, including the single-use
// gate (INV-account-06): a consumed code cannot redeem again.
func TestMfaBackupCode_SingleUseGuarded(t *testing.T) {
	h := newLoginHarness(t)
	h.svc.mfa = NewMfaVerifier(h.repo)
	h.repo.view = &LoginUserView{Name: "Backup User"}
	uid := h.verifyPendingUser
	secret := "JBSWY3DPEHPK3PXP"
	h.repo.seedMFAEnabled(uid, secret)
	h.repo.seedMFABackupCode(uid, "abcd1234")
	ctx := context.Background()
	tx, err := h.svc.tx.BeginTx(ctx)
	if err != nil {
		t.Fatalf("begin tx: %v", err)
	}

	ok, err := h.svc.mfa.VerifyBackupCode(ctx, tx, uid, "abcd1234")
	if err != nil || !ok {
		t.Fatalf("first redeem: ok=%t err=%v, want true/nil", ok, err)
	}
	// Replay — must fail (used_at only ever set once).
	ok, err = h.svc.mfa.VerifyBackupCode(ctx, tx, uid, "abcd1234")
	if err != nil || ok {
		t.Fatalf("replay redeem: ok=%t err=%v, want false/nil", ok, err)
	}
	// Normalization: user typing with dashes/uppercase still matches the
	// stored lowercase-alphanumeric hash.
	h.repo.seedMFABackupCode(uid, "wxyz9876")
	res, err := h.svc.LoginMfa(ctx, "pending-token", "", "WXYZ-9876")
	if err != nil || res.Status != "ok" {
		t.Fatalf("normalized backup-code login: status=%s err=%v, want ok/nil", res.Status, err)
	}
}

// Disabling MFA implicitly invalidates unused backup codes (INV-account-06)
// without deleting them — the enabled-gate rejects them at the verifier.
func TestMfaDisable_OldBackupCodesUnusable(t *testing.T) {
	h := newLoginHarness(t)
	h.svc.mfa = NewMfaVerifier(h.repo)
	uid := h.verifyPendingUser
	h.repo.seedMFAEnabled(uid, "JBSWY3DPEHPK3PXP")
	h.repo.seedMFABackupCode(uid, "abcd1234")
	ctx := context.Background()

	// Disable (rows left in place, enabled_at → NULL).
	tx, _ := h.svc.tx.BeginTx(ctx)
	if disabled, err := h.repo.SetMFADisabledIfEnabled(ctx, tx, uid); err != nil || !disabled {
		t.Fatalf("disable: disabled=%t err=%v", disabled, err)
	}

	// The never-used code is now unusable — no write (used_at stays NULL).
	ok, err := h.svc.mfa.VerifyBackupCode(ctx, tx, uid, "abcd1234")
	if err != nil || ok {
		t.Fatal("old backup code must be unusable after disable (implicit invalidation)")
	}
}

// ---- R15: no secrets/tokens in logs -----------------------------------------

func TestMfa_LogsFreeOfSecrets(t *testing.T) {
	buf := &bytes.Buffer{}
	prev := log.Writer()
	log.SetOutput(buf)
	t.Cleanup(func() { log.SetOutput(prev) })

	svc, repo, _ := newMFAService(t)
	uid := uuid.New()
	repo.seedView("alice@example.com")

	uri, err := svc.MfaEnroll(context.Background(), uid)
	if err != nil {
		t.Fatalf("enroll: %v", err)
	}
	// Extract the base32 secret from the URI to seed the same secret for
	// confirm, then exercise the confirm path too.
	q, err := url.Parse(uri)
	if err != nil {
		t.Fatalf("parse uri: %v", err)
	}
	secret := q.Query().Get("secret")
	repo.seedMFAEnrollment(uid, secret)
	if _, err := svc.MfaEnrollConfirm(context.Background(), uid, codeFor(secret)); err != nil {
		t.Fatalf("confirm: %v", err)
	}

	// The opaque base32 secret and the secret-embedding URI are the ONLY
	// things R15 forbids in logs. (Codes/TOTP inputs are never passed to
	// log.Printf by this service at all — the audit/outcome log lines carry
	// only user_id.)
	out := buf.String()
	if strings.Contains(out, secret) || strings.Contains(out, uri) {
		t.Errorf("log lines leaked MFA secret or otpauth URI:\n%s", out)
	}
}
