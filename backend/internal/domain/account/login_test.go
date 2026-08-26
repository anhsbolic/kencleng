package account

import (
	"bytes"
	"context"
	"errors"
	"log"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"

	"github.com/anhsbolic/kencleng/backend/internal/platform/auth"
	"github.com/anhsbolic/kencleng/backend/internal/platform/crypto"
	"github.com/anhsbolic/kencleng/backend/internal/platform/secrets"
)

// ---- loginFakeRepo: *fakeRepo + call-recording overrides for the
// login/session methods. Embedding keeps Register/VerifyEmail behavior from
// service_test.go available; login flows only touch the overridden half. ----

type loginFakeRepo struct {
	*fakeRepo
	mu sync.Mutex

	attempts []LoginAttempt

	countByIdentifier int
	countByUser       int

	refresh map[string]*RefreshToken // token_hash → row

	view          *LoginUserView
	idHashForUser string // backfill answer for FindIdentifierHashByUserAndProvider
}

func newLoginFakeRepo() *loginFakeRepo {
	return &loginFakeRepo{
		fakeRepo: newFakeRepo(),
		refresh:  make(map[string]*RefreshToken),
	}
}

func (f *loginFakeRepo) InsertLoginAttempt(_ context.Context, _ pgx.Tx, a *LoginAttempt) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	copy := *a
	f.attempts = append(f.attempts, copy)
	return nil
}

func (f *loginFakeRepo) CountRecentFailedAttemptsByIdentifier(_ context.Context, _, _ string, _ time.Time) (int, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.countByIdentifier, nil
}

func (f *loginFakeRepo) CountRecentFailedAttemptsByUser(_ context.Context, _ uuid.UUID, _ string, _ time.Time) (int, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	return f.countByUser, nil
}

func (f *loginFakeRepo) FindRefreshTokenByHash(_ context.Context, hash string) (*RefreshToken, bool, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	tok, ok := f.refresh[hash]
	if !ok {
		return nil, false, nil
	}
	cp := *tok
	return &cp, true, nil
}

func (f *loginFakeRepo) RotateRefreshToken(_ context.Context, _ pgx.Tx, oldTokenHash string, child *RefreshToken) (bool, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	parent, ok := f.refresh[oldTokenHash]
	if !ok || parent.RevokedAt != nil || parent.ReplacedByID != nil || !parent.ExpiresAt.After(time.Now()) {
		return false, nil
	}
	mark := child.ID
	parent.ReplacedByID = &mark
	child.UserID = parent.UserID
	child.FamilyID = parent.FamilyID
	stored := *child
	f.refresh[child.TokenHash] = &stored
	return true, nil
}

func (f *loginFakeRepo) RevokeRefreshTokenByHash(_ context.Context, _ pgx.Tx, hash string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if tok, ok := f.refresh[hash]; ok && tok.RevokedAt == nil {
		now := time.Now()
		tok.RevokedAt = &now
	}
	return nil
}

func (f *loginFakeRepo) RevokeRefreshTokenFamily(_ context.Context, _ pgx.Tx, familyID uuid.UUID) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	for _, tok := range f.refresh {
		if tok.FamilyID == familyID && tok.RevokedAt == nil {
			now := time.Now()
			tok.RevokedAt = &now
		}
	}
	return nil
}

func (f *loginFakeRepo) GetLoginUserView(_ context.Context, userID uuid.UUID) (*LoginUserView, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.view == nil {
		return nil, nil
	}
	cp := *f.view
	cp.ID = userID
	return &cp, nil
}

func (f *loginFakeRepo) FindIdentifierHashByUserAndProvider(_ context.Context, _ uuid.UUID, _ string) (string, bool, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.idHashForUser == "" {
		return "", false, nil
	}
	return f.idHashForUser, true, nil
}

func (f *loginFakeRepo) attemptsBySuccess(success bool) []LoginAttempt {
	var out []LoginAttempt
	for _, a := range f.attempts {
		if a.Success == success {
			out = append(out, a)
		}
	}
	return out
}

// ---- harness ----------------------------------------------------------------

type loginHarness struct {
	svc     *Service
	repo    *loginFakeRepo
	now     time.Time // read directly by tests; svc.now closure reads the same var
	compare struct {
		calls int
	}
	verifyPendingErr  error
	verifyPendingUser uuid.UUID
	logs              *bytes.Buffer
	restoreLog        func()
}

func newLoginHarness(t *testing.T) *loginHarness {
	t.Helper()
	keys := &crypto.Keys{EncryptionKey: make([]byte, 32), HMACKey: make([]byte, 32)}
	repo := newLoginFakeRepo()

	h := &loginHarness{repo: repo, now: baseTime()}
	h.verifyPendingUser = uuid.New()

	// Capture log output so R19 leak assertions can inspect it.
	buf := &bytes.Buffer{}
	h.logs = buf
	prev := log.Writer()
	log.SetOutput(buf)
	h.restoreLog = func() { log.SetOutput(prev) }
	t.Cleanup(h.restoreLog)

	realCredHash, err := secrets.HashPassword("correct-horse-battery")
	if err != nil {
		t.Fatalf("hash test credential: %v", err)
	}

	svc := &Service{
		repo:        repo,
		tx:          fakeTxRunner{},
		breachCheck: &fakeBreachChecker{},
		email:       &captureSender{},
		keys:        keys,
		mfa:         stubMfaVerifier{},
		now: func() time.Time {
			return h.now
		},
		compare: func(hashedPassword, password string) error {
			h.compare.calls++
			if hashedPassword == realCredHash && password == "correct-horse-battery" {
				return nil
			}
			return secrets.ComparePassword(hashedPassword, password) // mismatch path burns real bcrypt too
		},
		mintAccess: func(userID uuid.UUID, now time.Time) (string, error) {
			return "access-jwt-marker", nil
		},
		mintMFAPending: func(userID uuid.UUID, now time.Time) (string, error) {
			return "pending-jwt-marker", nil
		},
		verifyPending: func(token string, now time.Time) (uuid.UUID, error) {
			if h.verifyPendingErr != nil {
				return uuid.Nil, h.verifyPendingErr
			}
			return h.verifyPendingUser, nil
		},
	}
	h.svc = svc
	return h
}

// seedIdentity wires an email_password identity into both the register-era
// fake (for FindAuthIdentityByIdentifierHash) and the login view.
func (h *loginHarness) seedIdentity(t *testing.T, email string, credHash string) string {
	t.Helper()
	idHash := crypto.HMAC([]byte(email), h.svc.keys)
	ident := h.repo.seedIdentity(providerEmailPassword, email, idHash, false)
	ident.CredentialSecret = &credHash
	h.repo.view = &LoginUserView{Name: "View User"}
	h.repo.idHashForUser = idHash
	return idHash
}

func baseTime() time.Time {
	return time.Date(2026, 8, 26, 12, 0, 0, 0, time.UTC)
}

// ---- Login ------------------------------------------------------------------

// TestLogin_Success_NoMfa proves R1 end-to-end at the service layer.
func TestLogin_Success_NoMfa(t *testing.T) {
	h := newLoginHarness(t)
	credHash, err := secrets.HashPassword("correct-horse-battery")
	if err != nil {
		t.Fatalf("hash: %v", err)
	}
	idHash := h.seedIdentity(t, "user@example.com", credHash)

	res, err := h.svc.Login(context.Background(), "user@example.com", "correct-horse-battery")
	if err != nil {
		t.Fatalf("Login: %v", err)
	}

	if res.Status != "ok" || res.AccessToken == "" || res.RefreshTokenPlain == "" {
		t.Fatalf("unexpected ok-shape result: %+v", res)
	}
	if !res.AccessTokenExpiresAt.Equal(h.now.Add(auth.AccessTokenTTL)) {
		t.Errorf("access expiry = %v, want now+15m", res.AccessTokenExpiresAt)
	}
	if len(res.User.Name) == 0 {
		t.Error("user view not carried on result")
	}

	// Attempt bookkeeping: exactly one success row for this identifier.
	successes := h.repo.attemptsBySuccess(true)
	if len(successes) != 1 {
		t.Fatalf("success attempts = %d, want 1", len(successes))
	}
	a := successes[0]
	if a.Stage != stagePassword || a.IdentifierHash != idHash || a.UserID == nil {
		t.Errorf("attempt row wrong: %+v", a)
	}

	// Refresh row persisted hash-only with fresh family + 30d TTL.
	if len(h.repo.insertedRefreshTokens) != 1 {
		t.Fatalf("refresh rows = %d, want 1", len(h.repo.insertedRefreshTokens))
	}
	rt := h.repo.insertedRefreshTokens[0]
	if rt.ExpiresAt != h.now.Add(refreshTokenTTL) {
		t.Errorf("refresh expiry = %v, want now+30d", rt.ExpiresAt)
	}
	if sha256Hex(res.RefreshTokenPlain) != rt.TokenHash {
		t.Error("stored refresh hash does not match plain value")
	}
}

// TestLogin_MfaRequired_NoTokensIssuedYet proves R2's strict ordering:
// credentials passed ⇒ attempt recorded; MFA enrolled ⇒ NO tokens, NO
// refresh row — only the pending carrier.
func TestLogin_MfaRequired_NoTokensIssuedYet(t *testing.T) {
	h := newLoginHarness(t)
	credHash, _ := secrets.HashPassword("correct-horse-battery")
	h.seedIdentity(t, "mfa-user@example.com", credHash)
	h.repo.view.MFAEnabled = true

	res, err := h.svc.Login(context.Background(), "mfa-user@example.com", "correct-horse-battery")
	if err != nil {
		t.Fatalf("Login: %v", err)
	}
	if res.Status != "mfa_required" || res.MFAPendingToken == "" {
		t.Fatalf("unexpected mfa_required shape: %+v", res)
	}
	if res.AccessToken != "" || res.RefreshTokenPlain != "" || res.User != nil {
		t.Errorf("tokens/user leaked into mfa_required result: %+v", res)
	}
	if len(h.repo.insertedRefreshTokens) != 0 {
		t.Errorf("refresh rows created despite pending MFA: %d", len(h.repo.insertedRefreshTokens))
	}
	if fails := h.repo.attemptsBySuccess(false); len(fails) != 0 {
		t.Errorf("failed attempts recorded: %d", len(fails))
	}
	if successes := h.repo.attemptsBySuccess(true); len(successes) != 1 {
		t.Errorf("password success attempt not recorded")
	}
}

// TestLogin_GenericErrorMessage proves R3: wrong-password and unknown-email
// produce the SAME sentinel (transport renders byte-identical bodies), and
// both record a failed attempt (unknown email → NULL user_id).
func TestLogin_GenericErrorMessage(t *testing.T) {
	h := newLoginHarness(t)
	credHash, _ := secrets.HashPassword("correct-horse-battery")
	idHash := h.seedIdentity(t, "known@example.com", credHash)

	_, errWrongPw := h.svc.Login(context.Background(), "known@example.com", "wrong-password-123")
	_, errUnknownEmail := h.svc.Login(context.Background(), "ghost@example.com", "whatever-pw")

	if !errors.Is(errWrongPw, ErrInvalidCredentials) || !errors.Is(errUnknownEmail, ErrInvalidCredentials) {
		t.Fatalf("sentinels differ: wrongPw=%v unknown=%v", errWrongPw, errUnknownEmail)
	}

	fails := h.repo.attemptsBySuccess(false)
	if len(fails) != 2 {
		t.Fatalf("failed attempts = %d, want 2", len(fails))
	}
	knownRow, ghostRow := fails[0], fails[1]
	if knownRow.UserID == nil || *knownRow.UserID == uuid.Nil {
		t.Error("known-identity failure should carry user_id")
	}
	if ghostRow.UserID != nil {
		t.Error("unknown-email failure must carry NULL user_id")
	}
	if ghostRow.IdentifierHash == idHash {
		t.Error("unknown-email attempt should carry its own (different) identifier hash")
	}
}

// TestLogin_Lockout_5Failed15Min proves R4: threshold trips BEFORE any
// credential work, no attempt row is written for the rejected attempt, and
// 4 failures still pass through to verification.
func TestLogin_Lockout_5Failed15Min(t *testing.T) {
	h := newLoginHarness(t)
	credHash, _ := secrets.HashPassword("correct-horse-battery")
	idHash := h.seedIdentity(t, "locked@example.com", credHash)
	email := "locked@example.com"
	password := "wrong-password-123"

	// Below threshold: proceeds (and fails credentials).
	h.repo.countByIdentifier = maxFailedAttempts - 1
	before := h.compare.calls
	if _, err := h.svc.Login(context.Background(), email, password); !errors.Is(err, ErrInvalidCredentials) {
		t.Fatalf("below-threshold: err = %v, want ErrInvalidCredentials", err)
	}
	if h.compare.calls != before+1 {
		t.Errorf("below threshold: compare calls = %d, want +1", h.compare.calls-before)
	}

	// At threshold: locked out BEFORE bcrypt runs (R4 ordering).
	h.repo.countByIdentifier = maxFailedAttempts
	before = h.compare.calls
	beforeAttempts := len(h.repo.attempts)
	_, err := h.svc.Login(context.Background(), email, password)
	if !errors.Is(err, ErrLockedOut) {
		t.Fatalf("at threshold: err = %v, want ErrLockedOut", err)
	}
	if h.compare.calls != before {
		t.Errorf("lockout ran compare %d times; credential check must be skipped entirely", h.compare.calls-before)
	}
	if len(h.repo.attempts) != beforeAttempts {
		t.Error("lockout-rejected attempt wrote a row (spec violation)")
	}
	_ = idHash
}

// TestLogin_UnverifiedIdentity_Succeeds proves R5: verified_at plays no role
// in login.
func TestLogin_UnverifiedIdentity_Succeeds(t *testing.T) {
	h := newLoginHarness(t)
	credHash, _ := secrets.HashPassword("correct-horse-battery")
	// seedIdentity leaves VerifiedAt=false (nil) — deliberately unverified.
	h.seedIdentity(t, "unverified@example.com", credHash)

	if _, err := h.svc.Login(context.Background(), "unverified@example.com", "correct-horse-battery"); err != nil {
		t.Errorf("unverified identity login failed: %v", err)
	}
}

// TestLogin_TimingShape_NoEarlyReturn proves R18 structurally: every branch
// that reaches credential evaluation burns exactly one compare, including
// unknown identifiers; lockout skips it entirely (by design).
func TestLogin_TimingShape_NoEarlyReturn(t *testing.T) {
	h := newLoginHarness(t)
	credHash, _ := secrets.HashPassword("correct-horse-battery")
	h.seedIdentity(t, "timing-known@example.com", credHash)

	before := h.compare.calls
	if _, err := h.svc.Login(context.Background(), "timing-known@example.com", "wrong-password-123"); !errors.Is(err, ErrInvalidCredentials) {
		t.Fatalf("known wrong-password: err = %v", err)
	}
	if h.compare.calls != before+1 {
		t.Errorf("known-identity wrong-password: compare calls delta = %d, want 1", h.compare.calls-before)
	}

	before = h.compare.calls
	if _, err := h.svc.Login(context.Background(), "timing-ghost@example.com", "anything-at-all"); !errors.Is(err, ErrInvalidCredentials) {
		t.Fatalf("unknown identity: err = %v", err)
	}
	if h.compare.calls != before+1 {
		t.Errorf("unknown-identity: compare calls delta = %d, want 1 (dummy burn missing)", h.compare.calls-before)
	}
}

// TestLogin_LoggingNeverLeaksCredentials proves the domain half of R19:
// issued tokens and the submitted password never appear in log output.
func TestLogin_LoggingNeverLeaksCredentials(t *testing.T) {
	h := newLoginHarness(t)
	credHash, _ := secrets.HashPassword("correct-horse-battery")
	h.seedIdentity(t, "leakcheck@example.com", credHash)

	const markerPassword = "correct-horse-battery"
	res, err := h.svc.Login(context.Background(), "leakcheck@example.com", markerPassword)
	if err != nil {
		t.Fatalf("Login: %v", err)
	}

	logged := h.logs.String()
	for _, secret := range []string{markerPassword, res.RefreshTokenPlain, res.AccessToken} {
		if strings.Contains(logged, secret) {
			t.Errorf("log output contains secret %q", secret)
		}
	}
}

// ---- LoginMfa ---------------------------------------------------------------

// TestLoginMfa_Lockout_5Failed15Min proves R7: MFA-stage lockout keyed by
// user_id fires BEFORE code verification and writes no attempt row.
func TestLoginMfa_Lockout_5Failed15Min(t *testing.T) {
	h := newLoginHarness(t)
	h.repo.countByUser = maxFailedAttempts

	_, err := h.svc.LoginMfa(context.Background(), "pending-token", "123456", "")
	if !errors.Is(err, ErrLockedOut) {
		t.Fatalf("err = %v, want ErrLockedOut", err)
	}
	if len(h.repo.attempts) != 0 {
		t.Errorf("lockout-rejected MFA attempt wrote rows: %d", len(h.repo.attempts))
	}
}

// TestLoginMfa_InvalidPendingToken proves R10: expired/malformed/wrong-key
// pending tokens fail with no writes at all.
func TestLoginMfa_InvalidPendingToken(t *testing.T) {
	h := newLoginHarness(t)
	h.verifyPendingErr = errors.New("invalid")

	_, err := h.svc.LoginMfa(context.Background(), "bad-token", "123456", "")
	if !errors.Is(err, ErrMfaPendingInvalid) {
		t.Fatalf("err = %v, want ErrMfaPendingInvalid", err)
	}
	if len(h.repo.attempts) != 0 {
		t.Errorf("invalid-pending wrote attempt rows: %d", len(h.repo.attempts))
	}
	if len(h.repo.insertedRefreshTokens) != 0 {
		t.Errorf("invalid-pending issued refresh rows: %d", len(h.repo.insertedRefreshTokens))
	}
}

// recordingVerifier captures tx usage for backup-code assertions.
type recordingVerifier struct {
	totpCalls   int
	backupCalls int
	lastTx      pgx.Tx
	result      bool
}

func (r *recordingVerifier) VerifyTOTP(context.Context, uuid.UUID, string) (bool, error) {
	r.totpCalls++
	return r.result, nil
}

func (r *recordingVerifier) VerifyBackupCode(_ context.Context, tx pgx.Tx, _ uuid.UUID, _ string) (bool, error) {
	r.backupCalls++
	r.lastTx = tx
	return r.result, nil
}

// TestLoginMfa_WrongCode proves R11: valid pending + wrong code ⇒ failed
// MFA-stage attempt row, generic error, no tokens.
func TestLoginMfa_WrongCode(t *testing.T) {
	h := newLoginHarness(t)
	rec := &recordingVerifier{result: false}
	h.svc.mfa = rec

	_, err := h.svc.LoginMfa(context.Background(), "pending-token", "000000", "")
	if !errors.Is(err, ErrInvalidCredentials) {
		t.Fatalf("err = %v, want ErrInvalidCredentials", err)
	}
	if rec.totpCalls != 1 {
		t.Errorf("totp verifier calls = %d, want 1", rec.totpCalls)
	}
	fails := h.repo.attemptsBySuccess(false)
	if len(fails) != 1 || fails[0].Stage != stageMFA || fails[0].UserID == nil {
		t.Errorf("expected one mfa failure row with user_id, got %+v", h.repo.attempts)
	}
	if len(h.repo.insertedRefreshTokens) != 0 {
		t.Error("tokens issued despite wrong code")
	}
}

// TestLoginMfa_TotpSuccess_CompletesLogin proves R8: correct TOTP ⇒ mfa
// success attempt row + issuance identical to the no-MFA branch.
func TestLoginMfa_TotpSuccess_CompletesLogin(t *testing.T) {
	h := newLoginHarness(t)
	h.svc.mfa = &recordingVerifier{result: true}
	h.repo.view = &LoginUserView{Name: "MFA User"}

	res, err := h.svc.LoginMfa(context.Background(), "pending-token", "654321", "")
	if err != nil {
		t.Fatalf("LoginMfa: %v", err)
	}
	if res.Status != "ok" || res.AccessToken == "" || res.RefreshTokenPlain == "" || res.User == nil {
		t.Fatalf("completion shape wrong: %+v", res)
	}
	successes := h.repo.attemptsBySuccess(true)
	if len(successes) != 1 || successes[0].Stage != stageMFA {
		t.Errorf("expected one mfa success row, got %+v", h.repo.attempts)
	}
	if len(h.repo.insertedRefreshTokens) != 1 {
		t.Errorf("refresh rows = %d, want 1", len(h.repo.insertedRefreshTokens))
	}
}

// TestLoginMfa_BackupCode_CompletesLogin proves R9's flow shape: backup path
// redeems inside a tx and completes like TOTP (single-use invariant itself
// is proven against the real store in task #6 / task-06 suite).
func TestLoginMfa_BackupCode_CompletesLogin(t *testing.T) {
	h := newLoginHarness(t)
	rec := &recordingVerifier{result: true}
	h.svc.mfa = rec
	h.repo.view = &LoginUserView{Name: "Backup User"}

	res, err := h.svc.LoginMfa(context.Background(), "pending-token", "", "ABCD-1234")
	if err != nil {
		t.Fatalf("LoginMfa: %v", err)
	}
	if res.Status != "ok" {
		t.Fatalf("status = %q, want ok", res.Status)
	}
	if rec.backupCalls != 1 || rec.lastTx == nil {
		t.Errorf("backup verifier calls=%d lastTx nil=%t; redemption must run inside the caller tx",
			rec.backupCalls, rec.lastTx == nil)
	}
	successes := h.repo.attemptsBySuccess(true)
	if len(successes) != 1 || successes[0].Stage != stageMFA {
		t.Errorf("expected mfa success row, got %+v", h.repo.attempts)
	}
}

// TestLoginMfa_DefensiveBothOrNeitherCodes proves the service-level boundary:
// neither or both second factors is invalid input (handler validates first;
// service defends).
func TestLoginMfa_DefensiveBothOrNeitherCodes(t *testing.T) {
	h := newLoginHarness(t)

	if _, err := h.svc.LoginMfa(context.Background(), "p", "", ""); !errors.Is(err, ErrValidation) {
		t.Errorf("neither code: err = %v, want ErrValidation", err)
	}
	if _, err := h.svc.LoginMfa(context.Background(), "p", "111111", "ABCD"); !errors.Is(err, ErrValidation) {
		t.Errorf("both codes: err = %v, want ErrValidation", err)
	}
}

// ---- Refresh ----------------------------------------------------------------

func seedLiveToken(h *loginHarness, plain string, familyID uuid.UUID) *RefreshToken {
	hash := sha256Hex(plain)
	tok := &RefreshToken{
		ID: uuid.New(), UserID: uuid.New(), FamilyID: familyID,
		TokenHash: hash, ExpiresAt: h.now.Add(refreshTokenTTL), CreatedAt: h.now,
	}
	h.repo.refresh[hash] = tok
	return tok
}

// TestRefresh_Rotates_IssuesChild_SameFamily proves R12: guarded rotation
// marks the parent once, lands the child in the same family, and returns a
// fresh pair.
func TestRefresh_Rotates_IssuesChild_SameFamily(t *testing.T) {
	h := newLoginHarness(t)
	family := uuid.New()
	oldPlain := "old-refresh-plain"
	parent := seedLiveToken(h, oldPlain, family)

	res, err := h.svc.Refresh(context.Background(), oldPlain)
	if err != nil {
		t.Fatalf("Refresh: %v", err)
	}
	if res.RefreshTokenPlain == "" || res.RefreshTokenPlain == oldPlain {
		t.Fatalf("new plain token missing or unchanged")
	}
	if res.AccessToken == "" || !res.AccessTokenExpiresAt.Equal(h.now.Add(auth.AccessTokenTTL)) {
		t.Errorf("access pair wrong: %+v", res)
	}

	child, ok := h.repo.refresh[sha256Hex(res.RefreshTokenPlain)]
	if !ok {
		t.Fatal("child token not persisted")
	}
	if child.FamilyID != family {
		t.Errorf("child family = %s, want %s (rotation must preserve lineage)", child.FamilyID, family)
	}
	if parent.ReplacedByID == nil || *parent.ReplacedByID != child.ID {
		t.Errorf("parent replaced_by_id = %v, want child id %s", parent.ReplacedByID, child.ID)
	}
}

// TestRefresh_MissingOrExpiredCookie proves R13: empty input and expired
// tokens both collapse to ErrInvalidCredentials; expired tokens also take
// their family down (compromise assumed).
func TestRefresh_MissingOrExpiredCookie(t *testing.T) {
	h := newLoginHarness(t)
	family := uuid.New()
	expiredPlain := "expired-refresh-plain"
	tok := seedLiveToken(h, expiredPlain, family)
	tok.ExpiresAt = h.now.Add(-time.Minute)

	if _, err := h.svc.Refresh(context.Background(), ""); !errors.Is(err, ErrInvalidCredentials) {
		t.Errorf("empty plain: err = %v, want ErrInvalidCredentials", err)
	}
	if _, err := h.svc.Refresh(context.Background(), expiredPlain); !errors.Is(err, ErrInvalidCredentials) {
		t.Errorf("expired plain: err = %v, want ErrInvalidCredentials", err)
	}
	after, ok := h.repo.refresh[sha256Hex(expiredPlain)]
	if !ok || after.RevokedAt == nil {
		t.Error("expired token's family was not revoked on rejection")
	}
}

// TestRefresh_ReuseDetection_FamilyRevoked proves R14 / INV-account-04 with
// the full A→B→C chain: replaying A after two legitimate rotations revokes
// A, B, AND C, and C is dead afterward.
func TestRefresh_ReuseDetection_FamilyRevoked(t *testing.T) {
	h := newLoginHarness(t)
	family := uuid.New()

	plainA := "chain-a-plain"
	seedLiveToken(h, plainA, family)

	resB, err := h.svc.Refresh(context.Background(), plainA) // A → B
	if err != nil {
		t.Fatalf("rotate A: %v", err)
	}
	resC, err := h.svc.Refresh(context.Background(), resB.RefreshTokenPlain) // B → C
	if err != nil {
		t.Fatalf("rotate B: %v", err)
	}

	// Replay A — the stolen-token signature.
	if _, err := h.svc.Refresh(context.Background(), plainA); !errors.Is(err, ErrInvalidCredentials) {
		t.Fatalf("replay A: err = %v, want ErrInvalidCredentials", err)
	}

	for name, plain := range map[string]string{"A": plainA, "B": resB.RefreshTokenPlain, "C": resC.RefreshTokenPlain} {
		tok, ok := h.repo.refresh[sha256Hex(plain)]
		if !ok {
			t.Fatalf("%s missing from store", name)
		}
		if tok.RevokedAt == nil {
			t.Errorf("%s survived reuse detection (INV-account-04 violation)", name)
		}
	}

	// C — the last legitimately-issued token — must also be dead.
	if _, err := h.svc.Refresh(context.Background(), resC.RefreshTokenPlain); !errors.Is(err, ErrInvalidCredentials) {
		t.Errorf("refresh with revoked C: err = %v, want ErrInvalidCredentials", err)
	}
}

// TestRefresh_ConcurrentRequests_ExactlyOneWins simulates R15 at unit level
// (the authoritative proof runs against real Postgres in task-06): two
// goroutines race the same token; exactly one succeeds, and because the
// loser is treated as reuse (Assumption D), the whole family ends revoked.
func TestRefresh_ConcurrentRequests_ExactlyOneWins(t *testing.T) {
	h := newLoginHarness(t)
	family := uuid.New()
	plain := "race-refresh-plain"
	seedLiveToken(h, plain, family)

	const racers = 8
	results := make(chan error, racers)
	var wg sync.WaitGroup
	for i := 0; i < racers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := h.svc.Refresh(context.Background(), plain)
			results <- err
		}()
	}
	wg.Wait()
	close(results)

	wins := 0
	for err := range results {
		if err == nil {
			wins++
		} else if !errors.Is(err, ErrInvalidCredentials) {
			t.Fatalf("racer got unexpected error class: %v", err)
		}
	}
	if wins != 1 {
		t.Errorf("winners = %d, want exactly 1 (INV-account-03)", wins)
	}

	// Family fully revoked: every descendant of the lineage is dead.
	live := 0
	for _, tok := range h.repo.refresh {
		if tok.FamilyID == family && tok.RevokedAt == nil {
			live++
		}
	}
	if live != 0 {
		t.Errorf("%d live tokens remain in family after race (want 0 — loser ≡ attacker)", live)
	}
}

// ---- Logout -----------------------------------------------------------------

// TestLogout_RevokesAndClears proves R16's happy half: present cookie ⇒
// revoked_at set, nil error.
func TestLogout_RevokesAndClears(t *testing.T) {
	h := newLoginHarness(t)
	plain := "logout-refresh-plain"
	seedLiveToken(h, plain, uuid.New())

	if err := h.svc.Logout(context.Background(), plain); err != nil {
		t.Fatalf("Logout: %v", err)
	}
	tok := h.repo.refresh[sha256Hex(plain)]
	if tok.RevokedAt == nil {
		t.Error("revoked_at not set by logout")
	}
}

// TestLogout_NoCookie_Still204 proves R16's idempotent half: absent token ⇒
// nil error, nothing written.
func TestLogout_NoCookie_Still204(t *testing.T) {
	h := newLoginHarness(t)
	before := len(h.repo.attempts)

	if err := h.svc.Logout(context.Background(), ""); err != nil {
		t.Fatalf("Logout(empty): %v", err)
	}
	if err := h.svc.Logout(context.Background(), "never-issued-plain"); err != nil {
		t.Fatalf("Logout(never-issued): %v", err)
	}
	if len(h.repo.attempts) != before {
		t.Error("no-op logout wrote rows")
	}
}

// ---- writeAttempt fail-open (Q1 follow-up) ----------------------------------

// failingAttemptRepo wraps loginFakeRepo and forces InsertLoginAttempt to
// fail. Used to prove writeAttempt's fail-open contract: a lost audit row
// must NOT block a valid login nor mask a credential rejection.
type failingAttemptRepo struct {
	*loginFakeRepo
}

func (f *failingAttemptRepo) InsertLoginAttempt(_ context.Context, _ pgx.Tx, _ *LoginAttempt) error {
	return errors.New("simulated audit-db outage")
}

// TestWriteAttempt_FailOpen_ValidLoginStillSucceeds proves that a
// login_attempts write failure does not block a valid credential login
// (the audit row is bookkeeping, not a state machine — a lost row can
// only undercount toward lockout, never lock spuriously).
func TestWriteAttempt_FailOpen_ValidLoginStillSucceeds(t *testing.T) {
	h := newLoginHarness(t)
	credHash, _ := secrets.HashPassword("correct-horse-battery")
	h.seedIdentity(t, "failopen@example.com", credHash)

	h.svc.repo = &failingAttemptRepo{loginFakeRepo: h.repo}

	res, err := h.svc.Login(context.Background(), "failopen@example.com", "correct-horse-battery")
	if err != nil {
		t.Fatalf("valid login blocked by audit-write failure (must be fail-open): %v", err)
	}
	if res.Status != "ok" || res.AccessToken == "" {
		t.Errorf("login result wrong despite fail-open: %+v", res)
	}
}

// TestWriteAttempt_FailOpen_InvalidCredentialsStillRejected proves the
// fail-open path does not mask a credential failure: wrong password +
// audit-write error ⇒ still ErrInvalidCredentials, not nil.
func TestWriteAttempt_FailOpen_InvalidCredentialsStillRejected(t *testing.T) {
	h := newLoginHarness(t)
	credHash, _ := secrets.HashPassword("correct-horse-battery")
	h.seedIdentity(t, "failopen-reject@example.com", credHash)
	h.svc.repo = &failingAttemptRepo{loginFakeRepo: h.repo}

	_, err := h.svc.Login(context.Background(), "failopen-reject@example.com", "wrong-password")
	if !errors.Is(err, ErrInvalidCredentials) {
		t.Fatalf("fail-open masked credential rejection: err = %v, want ErrInvalidCredentials", err)
	}
}
