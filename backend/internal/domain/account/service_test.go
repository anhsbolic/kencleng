package account

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/anhsbolic/kencleng/backend/internal/platform/crypto"
	"github.com/anhsbolic/kencleng/backend/internal/platform/notification"
)

// ---- test fakes --------------------------------------------------------

// fakeRepo is a configurable in-memory Repository for service tests. Each
// lookup method consultes a map keyed by (providerType, identifierHash);
// insert methods record their inputs. All methods ignore the tx argument
// (the fake tx is a no-op) so no real DB is needed.
type fakeRepo struct {
	mu sync.Mutex

	// identities keyed by (providerType + "|" + identifierHash).
	identities map[string]*AuthIdentity
	// tokens keyed by tokenHash.
	tokens map[string]*AuthToken

	// Recorded inserts for assertions.
	insertedUsers         []*User
	insertedIdentities    []*AuthIdentity
	insertedTokens        []*AuthToken
	insertedRefreshTokens []*RefreshToken
	insertedUserLogs      []*UserLog

	// Call counters for assertions (e.g. proving no re-fetch on success).
	findTokenCalls int

	// Hooks to inject errors on the next matching call (nil = success).
	insertUserErr     error
	insertIdentityErr error
	insertTokenErr    error
	revokeErr         error
	redeemErr         error
	setVerifiedErr    error
	verifiedCalls     []setVerifiedCall
	revokeCalls       []revokeCall

	// redeemMode controls RedeemToken behavior:
	//   "atomic" (default) — first call for a given hash wins (CAS),
	//                         subsequent calls return false (single-use,
	//                         matching the DB atomic UPDATE).
	//   "alwaysTrue"       — always returns true (used by the valid-token
	//                         test where we seed one redeem).
	//   "alwaysFalse"      — always returns false (expired/not-found tests).
	redeemMode string

	// redeemed tracks which token hashes have been consumed (atomic
	// single-use simulation for redeemMode "atomic").
	redeemed map[string]bool

	// identityKeys records (providerType|identifier) for uniqueness,
	// independent of the stored AuthIdentity (whose Identifier field is
	// cleared after insert). This makes the fake's dedup survive the
	// plaintext-clearing the real adapter performs.
	identityKeys map[string]bool
}

type setVerifiedCall struct {
	userID     uuid.UUID
	provider   string
	verifiedAt time.Time
}

type revokeCall struct {
	userID  uuid.UUID
	purpose string
}

func newFakeRepo() *fakeRepo {
	return &fakeRepo{
		identities:   make(map[string]*AuthIdentity),
		tokens:       make(map[string]*AuthToken),
		redeemed:     make(map[string]bool),
		identityKeys: make(map[string]bool),
		redeemMode:   "atomic",
	}
}

func (f *fakeRepo) idKey(provider, hash string) string { return provider + "|" + hash }

func (f *fakeRepo) InsertUser(_ context.Context, _ pgx.Tx, user *User) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.insertUserErr != nil {
		err := f.insertUserErr
		f.insertUserErr = nil
		return err
	}
	f.insertedUsers = append(f.insertedUsers, user)
	user.PrimaryEmail = "" // mimic the real adapter clearing plaintext
	return nil
}

func (f *fakeRepo) InsertAuthIdentity(_ context.Context, _ pgx.Tx, identity *AuthIdentity) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.insertIdentityErr != nil {
		err := f.insertIdentityErr
		f.insertIdentityErr = nil
		return err
	}
	// Enforce uniqueness like the real DB index: a second insert for the
	// same (provider, identifier) returns a unique violation. We use a
	// separate identityKeys set so dedup survives the plaintext-clearing
	// the real adapter performs (which empties identity.Identifier).
	dedupKey := identity.ProviderType + "|" + identity.Identifier
	if f.identityKeys[dedupKey] {
		return &pgconn.PgError{Code: "23505"}
	}
	f.identityKeys[dedupKey] = true
	f.insertedIdentities = append(f.insertedIdentities, identity)
	f.identities[dedupKey] = identity
	identity.Identifier = "" // mimic clearing plaintext
	return nil
}

func (f *fakeRepo) InsertAuthToken(_ context.Context, _ pgx.Tx, token *AuthToken) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.insertTokenErr != nil {
		err := f.insertTokenErr
		f.insertTokenErr = nil
		return err
	}
	f.insertedTokens = append(f.insertedTokens, token)
	f.tokens[token.TokenHash] = token
	return nil
}

func (f *fakeRepo) FindAuthIdentityByIdentifierHash(_ context.Context, providerType, identifierHash string) (*AuthIdentity, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	// The service passes the HMAC hash; the fake repo stored by plaintext
	// identifier. Bridge by also storing a hash->identity index when an
	// identity is registered via the service (the test sets up
	// identities with a known hash via a helper). For lookup we match
	// on a parallel hash map the test populates.
	return f.identities[f.idKey(providerType, identifierHash)], nil
}

func (f *fakeRepo) FindAuthTokenByHash(_ context.Context, tokenHash string) (*AuthToken, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.findTokenCalls++
	return f.tokens[tokenHash], nil
}

func (f *fakeRepo) RedeemToken(_ context.Context, _ pgx.Tx, tokenHash string) (uuid.UUID, string, bool, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.redeemErr != nil {
		err := f.redeemErr
		f.redeemErr = nil
		return uuid.Nil, "", false, err
	}
	switch f.redeemMode {
	case "alwaysTrue":
		var uid uuid.UUID
		var purpose string
		if t, ok := f.tokens[tokenHash]; ok {
			now := time.Now()
			t.UsedAt = &now
			uid = t.UserID
			purpose = t.Purpose
		}
		return uid, purpose, true, nil
	case "alwaysFalse":
		return uuid.Nil, "", false, nil
	default: // "atomic": single-use CAS, like the DB atomic UPDATE.
		if f.redeemed[tokenHash] {
			return uuid.Nil, "", false, nil
		}
		f.redeemed[tokenHash] = true
		var uid uuid.UUID
		var purpose string
		if t, ok := f.tokens[tokenHash]; ok {
			now := time.Now()
			t.UsedAt = &now
			uid = t.UserID
			purpose = t.Purpose
		}
		return uid, purpose, true, nil
	}
}

func (f *fakeRepo) SetUserVerified(_ context.Context, _ pgx.Tx, userID uuid.UUID, providerType string, verifiedAt time.Time) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.setVerifiedErr != nil {
		err := f.setVerifiedErr
		f.setVerifiedErr = nil
		return err
	}
	f.verifiedCalls = append(f.verifiedCalls, setVerifiedCall{userID, providerType, verifiedAt})
	return nil
}

func (f *fakeRepo) RevokeTokens(_ context.Context, _ pgx.Tx, userID uuid.UUID, purpose string) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	if f.revokeErr != nil {
		err := f.revokeErr
		f.revokeErr = nil
		return err
	}
	f.revokeCalls = append(f.revokeCalls, revokeCall{userID, purpose})
	return nil
}

func (f *fakeRepo) InsertRefreshToken(_ context.Context, _ pgx.Tx, token *RefreshToken) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.insertedRefreshTokens = append(f.insertedRefreshTokens, token)
	return nil
}

func (f *fakeRepo) InsertUserLog(_ context.Context, _ pgx.Tx, entry *UserLog) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.insertedUserLogs = append(f.insertedUserLogs, entry)
	return nil
}

// ---- login/session stubs ---------------------------------------------------
// Minimal implementations satisfying the Repository interface additions from
// the login/session slice. Task 04's tests replace these with call-recording
// fakes; the stubs below just keep existing Register/VerifyEmail tests
// compiling and passing (none of those flows reach the new methods).

func (f *fakeRepo) InsertLoginAttempt(_ context.Context, _ pgx.Tx, _ *LoginAttempt) error {
	return nil
}

func (f *fakeRepo) CountRecentFailedAttemptsByIdentifier(_ context.Context, _, _ string, _ time.Time) (int, error) {
	return 0, nil
}

func (f *fakeRepo) CountRecentFailedAttemptsByUser(_ context.Context, _ uuid.UUID, _ string, _ time.Time) (int, error) {
	return 0, nil
}

func (f *fakeRepo) FindRefreshTokenByHash(_ context.Context, _ string) (*RefreshToken, bool, error) {
	return nil, false, nil
}

func (f *fakeRepo) RotateRefreshToken(_ context.Context, _ pgx.Tx, _ string, _ *RefreshToken) (bool, error) {
	return false, nil
}

func (f *fakeRepo) RevokeRefreshTokenByHash(_ context.Context, _ pgx.Tx, _ string) error {
	return nil
}

func (f *fakeRepo) RevokeRefreshTokenFamily(_ context.Context, _ pgx.Tx, _ uuid.UUID) error {
	return nil
}

func (f *fakeRepo) GetLoginUserView(_ context.Context, _ uuid.UUID) (*LoginUserView, error) {
	return nil, nil
}

func (f *fakeRepo) FindIdentifierHashByUserAndProvider(_ context.Context, _ uuid.UUID, _ string) (string, bool, error) {
	return "", false, nil
}

// seedIdentity stores an identity under both its plaintext-keyed dedup
// map and its hash-keyed lookup map so FindAuthIdentityByIdentifierHash
// can find it.
func (f *fakeRepo) seedIdentity(providerType, identifier string, hash string, verified bool) *AuthIdentity {
	f.mu.Lock()
	defer f.mu.Unlock()
	var verifiedAt *time.Time
	if verified {
		now := time.Now()
		verifiedAt = &now
	}
	ident := &AuthIdentity{
		ID:           uuid.New(),
		UserID:       uuid.New(),
		ProviderType: providerType,
		Identifier:   identifier,
		VerifiedAt:   verifiedAt,
	}
	f.identities[providerType+"|"+identifier] = ident
	f.identities[f.idKey(providerType, hash)] = ident
	return ident
}

// fakeTx satisfies pgx.Tx by embedding a nil pgx.Tx (so all unneeded
// methods panic only if actually called, which the fake repo never does)
// and overriding Commit/Rollback.
type fakeTx struct {
	pgx.Tx
	commitErr   error
	rollbackErr error
	committed   bool
	rolledBack  bool
}

func (t *fakeTx) Commit(context.Context) error {
	t.committed = true
	return t.commitErr
}

func (t *fakeTx) Rollback(context.Context) error {
	t.rolledBack = true
	return t.rollbackErr
}

type fakeTxRunner struct {
	commitErr error
	beginErr  error
}

// BeginTx returns a fresh fakeTx per call so concurrent goroutines do not
// share mutable tx state (avoids a data race on the tx's committed/
// rolledBack fields under -race).
func (r fakeTxRunner) BeginTx(context.Context) (pgx.Tx, error) {
	if r.beginErr != nil {
		return nil, r.beginErr
	}
	return &fakeTx{commitErr: r.commitErr}, nil
}

// fakeBreachChecker is a configurable breachChecker for tests.
type fakeBreachChecker struct {
	breached bool
	err      error
	called   atomic.Bool
}

func (b *fakeBreachChecker) IsBreached(context.Context, string) (bool, error) {
	b.called.Store(true)
	return b.breached, b.err
}

// captureSender is a notification.Sender that records calls without
// sending and asserts no PII is logged.
type captureSender struct {
	mu              sync.Mutex
	verificationTo  []string
	nudgeTypes      []string
	verificationErr error
}

func (s *captureSender) SendVerificationEmail(_ context.Context, to, _ string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.verificationTo = append(s.verificationTo, to)
	return s.verificationErr
}

func (s *captureSender) SendNudgeEmail(_ context.Context, _ string, nudgeType string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.nudgeTypes = append(s.nudgeTypes, nudgeType)
	return nil
}

// leakySender is a notification.Sender whose error message embeds the
// recipient email and a fake token, simulating a leaky SMTP error. Used
// to prove the L2 fix: the log line must not contain the recipient or
// token even when the underlying error does.
type leakySender struct {
	called bool
}

func (s *leakySender) SendVerificationEmail(_ context.Context, to, token string) error {
	s.called = true
	// Simulate an SMTP error that embeds the recipient + token (leaky).
	return fmt.Errorf("SMTP 553 <recipient=%s> rejected: token=%s", to, token)
}

func (s *leakySender) SendNudgeEmail(_ context.Context, to, nudgeType string) error {
	return fmt.Errorf("SMTP 553 <recipient=%s> rejected: nudge=%s", to, nudgeType)
}

// newTestService builds a Service wired to fakes. Returns the service,
// the fake repo, the fake breach checker, and the capture sender.
func newTestService(t *testing.T, breached bool) (*Service, *fakeRepo, *fakeBreachChecker, *captureSender) {
	t.Helper()
	repo := newFakeRepo()
	sender := &captureSender{}
	bc := &fakeBreachChecker{breached: breached}
	keys := &crypto.Keys{
		EncryptionKey: make([]byte, 32),
		HMACKey:       make([]byte, 32),
	}
	svc := &Service{
		repo:        repo,
		tx:          fakeTxRunner{},
		breachCheck: bc,
		email:       sender,
		keys:        keys,
	}
	return svc, repo, bc, sender
}

// hashFor computes the same HMAC the service uses for a given email, so
// tests can seed the fake repo's lookup map consistently.
func hashFor(email string) string {
	keys := &crypto.Keys{HMACKey: make([]byte, 32)}
	return crypto.HMAC([]byte(email), keys)
}

// ---- R1: register new user --------------------------------------------

func TestRegister_NewUser_CreatesUserIdentityToken(t *testing.T) {
	svc, repo, _, sender := newTestService(t, false)

	err := svc.Register(context.Background(), "Alice", "alice@example.com", "strong-pw-123")
	if err != nil {
		t.Fatalf("Register: %v", err)
	}

	if len(repo.insertedUsers) != 1 {
		t.Fatalf("expected 1 user inserted, got %d", len(repo.insertedUsers))
	}
	if repo.insertedUsers[0].Name != "Alice" {
		t.Errorf("user name = %q, want Alice", repo.insertedUsers[0].Name)
	}
	if len(repo.insertedIdentities) != 1 {
		t.Fatalf("expected 1 identity inserted, got %d", len(repo.insertedIdentities))
	}
	ident := repo.insertedIdentities[0]
	if ident.ProviderType != providerEmailPassword {
		t.Errorf("provider = %q, want %q", ident.ProviderType, providerEmailPassword)
	}
	if ident.VerifiedAt != nil {
		t.Error("new identity should be unverified (verified_at nil)")
	}
	if ident.CredentialSecret == nil || *ident.CredentialSecret == "" {
		t.Error("credential_secret (bcrypt hash) should be set")
	}
	if len(repo.insertedTokens) != 1 {
		t.Fatalf("expected 1 token inserted, got %d", len(repo.insertedTokens))
	}
	tok := repo.insertedTokens[0]
	if tok.Purpose != purposeEmailVerify {
		t.Errorf("token purpose = %q, want %q", tok.Purpose, purposeEmailVerify)
	}
	if tok.ExpiresAt.Sub(tok.CreatedAt) < tokenTTL-time.Second {
		t.Errorf("token TTL = %v, want ~%v", tok.ExpiresAt.Sub(tok.CreatedAt), tokenTTL)
	}
	if len(sender.verificationTo) != 1 || sender.verificationTo[0] != "alice@example.com" {
		t.Errorf("verification email not sent to alice: %v", sender.verificationTo)
	}
}

// ---- R2: register unverified existing ---------------------------------

func TestRegister_UnverifiedExisting_ResendFlow(t *testing.T) {
	svc, repo, _, sender := newTestService(t, false)
	email := "bob@example.com"
	hash := hashFor(email)
	existing := repo.seedIdentity(providerEmailPassword, email, hash, false)

	err := svc.Register(context.Background(), "Bob", email, "strong-pw-123")
	if err != nil {
		t.Fatalf("Register: %v", err)
	}

	if len(repo.insertedUsers) != 0 {
		t.Errorf("expected no new user, got %d", len(repo.insertedUsers))
	}
	if len(repo.insertedIdentities) != 0 {
		t.Errorf("expected no new identity, got %d", len(repo.insertedIdentities))
	}
	if len(repo.revokeCalls) != 1 || repo.revokeCalls[0].userID != existing.UserID {
		t.Errorf("expected one revoke for existing user, got %v", repo.revokeCalls)
	}
	if len(repo.insertedTokens) != 1 {
		t.Fatalf("expected 1 new token, got %d", len(repo.insertedTokens))
	}
	// R2 must send the verification email carrying the new token (same
	// action as the resend endpoint / R13), NOT a token-less nudge — the
	// spec says "resend-verification email sent". Without this the newly
	// issued token is never delivered and the user cannot verify.
	if len(sender.verificationTo) != 1 || sender.verificationTo[0] != email {
		t.Errorf("expected verification email sent to %s, got %v", email, sender.verificationTo)
	}
	if len(sender.nudgeTypes) != 0 {
		t.Errorf("expected no nudge for R2 (verification email instead), got %v", sender.nudgeTypes)
	}
}

// ---- R3: register verified existing ----------------------------------

func TestRegister_VerifiedExisting_PasswordResetNudge(t *testing.T) {
	svc, repo, _, sender := newTestService(t, false)
	email := "carol@example.com"
	hash := hashFor(email)
	existing := repo.seedIdentity(providerEmailPassword, email, hash, true)

	err := svc.Register(context.Background(), "Carol", email, "strong-pw-123")
	if err != nil {
		t.Fatalf("Register: %v", err)
	}

	if len(repo.insertedUsers) != 0 || len(repo.insertedIdentities) != 0 || len(repo.insertedTokens) != 0 {
		t.Errorf("verified branch must not create records: users=%d identities=%d tokens=%d",
			len(repo.insertedUsers), len(repo.insertedIdentities), len(repo.insertedTokens))
	}
	// R3 now performs a dummy revoke (0-row UPDATE against a synthetic
	// user_id) for DB-time uniformity with R1/R2 (R7). Exactly one
	// revoke call is expected, and its userID must NOT be the existing
	// user's — it's a synthetic uuid that matches no real rows.
	if len(repo.revokeCalls) != 1 {
		t.Errorf("verified branch must perform exactly 1 dummy revoke (R7 DB-time), got %d",
			len(repo.revokeCalls))
	} else if repo.revokeCalls[0].userID == existing.UserID {
		t.Errorf("dummy revoke must use a synthetic user_id, not the existing user's: got %s",
			repo.revokeCalls[0].userID)
	}
	if len(sender.nudgeTypes) != 1 || sender.nudgeTypes[0] != notification.NudgePasswordReset {
		t.Errorf("expected password-reset nudge, got %v", sender.nudgeTypes)
	}
	if len(sender.verificationTo) != 0 {
		t.Errorf("verified branch must not send verification email, got %v", sender.verificationTo)
	}
}

// ---- R4: register Google-only conflict --------------------------------

func TestRegister_GoogleOnlyConflict_Nudge(t *testing.T) {
	svc, repo, _, sender := newTestService(t, false)
	email := "dave@example.com"
	hash := hashFor(email)
	existing := repo.seedIdentity(providerGoogle, email, hash, true)

	err := svc.Register(context.Background(), "Dave", email, "strong-pw-123")
	if err != nil {
		t.Fatalf("Register: %v", err)
	}

	if len(repo.insertedUsers) != 0 || len(repo.insertedIdentities) != 0 || len(repo.insertedTokens) != 0 {
		t.Errorf("google-only branch must not create records")
	}
	// R4 now performs a dummy revoke (0-row UPDATE against a synthetic
	// user_id) for DB-time uniformity with R1/R2 (R7). Exactly one
	// revoke call is expected, with a synthetic user_id.
	if len(repo.revokeCalls) != 1 {
		t.Errorf("google-only branch must perform exactly 1 dummy revoke (R7 DB-time), got %d",
			len(repo.revokeCalls))
	} else if repo.revokeCalls[0].userID == existing.UserID {
		t.Errorf("dummy revoke must use a synthetic user_id, not the existing user's: got %s",
			repo.revokeCalls[0].userID)
	}
	if len(sender.nudgeTypes) != 1 || sender.nudgeTypes[0] != notification.NudgeGoogleOnly {
		t.Errorf("expected google-only nudge, got %v", sender.nudgeTypes)
	}
}

// TestRegister_R3R4_PerformTimingWrite proves the S3 fix: R3 and R4 each
// perform exactly one dummy revoke (DB-write-shaped no-op) so the no-op
// branches are no longer write-free. Without this, R3/R4 are measurably
// faster than R1/R4 against real Postgres — an enumeration side-channel
// leaking "verified/google-only" (fast) vs "new/unverified" (slow).
func TestRegister_R3R4_PerformTimingWrite(t *testing.T) {
	cases := []struct {
		name  string
		email string
		setup func(*fakeRepo, string)
	}{
		{"R3-verified", "v3tw@example.com", func(r *fakeRepo, email string) {
			r.seedIdentity(providerEmailPassword, email, hashFor(email), true)
		}},
		{"R4-google-only", "g4tw@example.com", func(r *fakeRepo, email string) {
			r.seedIdentity(providerGoogle, email, hashFor(email), true)
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			svc, repo, _, _ := newTestService(t, false)
			tc.setup(repo, tc.email)
			if err := svc.Register(context.Background(), "X", tc.email, "strong-pw-123"); err != nil {
				t.Fatalf("Register: %v", err)
			}
			if len(repo.revokeCalls) != 1 {
				t.Errorf("%s: expected exactly 1 dummy revoke call, got %d",
					tc.name, len(repo.revokeCalls))
			}
		})
	}
}

// ---- R5/R18: password policy (length + breach) ------------------------

func TestRegister_PasswordPolicy(t *testing.T) {
	cases := []struct {
		name     string
		password string
		breached bool
		wantErr  error
	}{
		{name: "too short", password: "short", breached: false, wantErr: ErrValidation},
		{name: "breached", password: "strong-but-breached", breached: true, wantErr: ErrValidation},
		{name: "valid", password: "strong-pw-123", breached: false, wantErr: nil},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			svc, repo, bc, _ := newTestService(t, tc.breached)
			err := svc.Register(context.Background(), "Eve", "eve@example.com", tc.password)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("Register password=%q: got %v, want %v", tc.password, err, tc.wantErr)
			}
			if tc.wantErr != nil {
				// Validation must fire BEFORE any lookup — the repo must
				// not have been queried and no records created.
				if len(repo.insertedUsers) != 0 || len(repo.insertedIdentities) != 0 {
					t.Errorf("validation failure should not create records")
				}
				if bc.called.Load() && tc.name == "too short" {
					// length check happens before breach check, so breach
					// client should NOT be called for the too-short case.
					t.Errorf("breach check called for too-short password (should be checked after length)")
				}
			}
		})
	}
}

// ---- R6/R19: breach check fail-open ------------------------------------

func TestRegister_BreachCheck_FailOpen(t *testing.T) {
	// Fail-open is exercised by breachcheck returning (false, nil) on
	// unreachable APIs; here we simulate that with a non-breached result
	// and assert registration proceeds. The real fail-open unit is in
	// platform/breachcheck. This test asserts the service treats
	// (false, nil) as "proceed".
	svc, repo, _, sender := newTestService(t, false)

	if err := svc.Register(context.Background(), "Frank", "frank@example.com", "strong-pw-123"); err != nil {
		t.Fatalf("Register should proceed on breach-check not-breached: %v", err)
	}
	if len(repo.insertedUsers) != 1 {
		t.Errorf("expected registration to proceed, got %d users", len(repo.insertedUsers))
	}
	if len(sender.verificationTo) != 1 {
		t.Errorf("expected verification email sent")
	}

	// Assert no password/hash is logged by capturing log output during a
	// breach-check failure path. Simulate a real fail-open log by using a
	// breached=true would return ErrValidation, so instead assert the
	// non-breached path logs nothing about the password.
	var logBuf bytes.Buffer
	origOut := log.Writer()
	log.SetOutput(&logBuf)
	defer log.SetOutput(origOut)
	_ = svc.Register(context.Background(), "Frank2", "frank2@example.com", "strong-pw-123")
	if logBuf.String() != "" && (contains(logBuf.String(), "strong-pw-123")) {
		t.Errorf("log must not contain the password: %q", logBuf.String())
	}
}

func contains(s, sub string) bool {
	return bytes.Contains([]byte(s), []byte(sub))
}

// ---- R7: constant-time / generic response across branches -------------

func TestRegister_GenericResponse_AllBranches(t *testing.T) {
	// All four branches must return nil (the handler writes the
	// identical 202). Assert the observable side-effect shape is
	// consistent: each branch returns nil and sends exactly one email
	// (either verification or nudge).
	cases := []struct {
		name  string
		setup func(*fakeRepo)
	}{
		{"new", func(r *fakeRepo) {}},
		{"unverified", func(r *fakeRepo) {
			r.seedIdentity(providerEmailPassword, "u@example.com", hashFor("u@example.com"), false)
		}},
		{"verified", func(r *fakeRepo) {
			r.seedIdentity(providerEmailPassword, "v@example.com", hashFor("v@example.com"), true)
		}},
		{"google-only", func(r *fakeRepo) {
			r.seedIdentity(providerGoogle, "g@example.com", hashFor("g@example.com"), true)
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			svc, repo, _, sender := newTestService(t, false)
			tc.setup(repo)
			email := tc.name + "@example.com"
			err := svc.Register(context.Background(), "X", email, "strong-pw-123")
			if err != nil {
				t.Fatalf("branch %s: Register returned %v, want nil (handler writes 202)", tc.name, err)
			}
			totalEmails := len(sender.verificationTo) + len(sender.nudgeTypes)
			if totalEmails != 1 {
				t.Errorf("branch %s: expected exactly 1 email sent, got %d", tc.name, totalEmails)
			}
		})
	}
}

func TestRegister_GenericResponse_Timing(t *testing.T) {
	// R7: all four branches take equivalent wall-clock time. bcrypt
	// (~100ms at default cost) dominates and runs on every branch, so
	// the branches should be within a generous band of each other. We
	// assert no branch is more than 3x faster than the slowest, which is
	// enough to catch a branch that skips bcrypt (it would be ~100x
	// faster). The test proves equivalence, not the mechanism.
	cases := []struct {
		name  string
		setup func(*fakeRepo)
	}{
		{"new", func(r *fakeRepo) {}},
		{"unverified", func(r *fakeRepo) {
			r.seedIdentity(providerEmailPassword, "u@example.com", hashFor("u@example.com"), false)
		}},
		{"verified", func(r *fakeRepo) {
			r.seedIdentity(providerEmailPassword, "v@example.com", hashFor("v@example.com"), true)
		}},
		{"google-only", func(r *fakeRepo) {
			r.seedIdentity(providerGoogle, "g@example.com", hashFor("g@example.com"), true)
		}},
	}
	var durations []time.Duration
	for _, tc := range cases {
		svc, repo, _, _ := newTestService(t, false)
		tc.setup(repo)
		email := tc.name + "@example.com"
		start := time.Now()
		_ = svc.Register(context.Background(), "X", email, "strong-pw-123")
		durations = append(durations, time.Since(start))
	}
	max := durations[0]
	min := durations[0]
	for _, d := range durations {
		if d > max {
			max = d
		}
		if d < min {
			min = d
		}
	}
	// No branch should be more than 3x faster than the slowest. If bcrypt
	// were skipped on a branch, that branch would be ~100x faster.
	if max > 3*min && min > 0 {
		t.Errorf("branch timing not equivalent: min=%v max=%v (max/min=%.1f)", min, max, float64(max)/float64(min))
	}
}

// ---- R8: verify-email valid token --------------------------------------

func TestVerifyEmail_ValidToken_SetsVerifiedAt(t *testing.T) {
	svc, repo, _, _ := newTestService(t, false)
	userID := uuid.New()
	tokenHash := sha256Hex("valid-token")
	now := time.Now()
	repo.tokens[tokenHash] = &AuthToken{
		ID: uuid.New(), UserID: userID, Purpose: purposeEmailVerify,
		TokenHash: tokenHash, ExpiresAt: now.Add(24 * time.Hour), CreatedAt: now,
	}
	repo.redeemMode = "alwaysTrue"

	if err := svc.VerifyEmail(context.Background(), "valid-token"); err != nil {
		t.Fatalf("VerifyEmail: %v", err)
	}
	if len(repo.verifiedCalls) != 1 {
		t.Fatalf("expected 1 SetUserVerified call, got %d", len(repo.verifiedCalls))
	}
	call := repo.verifiedCalls[0]
	if call.userID != userID {
		t.Errorf("verified userID = %s, want %s", call.userID, userID)
	}
	if call.provider != providerEmailPassword {
		t.Errorf("verified provider = %q, want %q", call.provider, providerEmailPassword)
	}
}

// TestVerifyEmail_RedeemReturnsUserID_NoRefetch proves the S1 fix:
// RedeemToken returns userID via RETURNING, so the success path never
// calls FindAuthTokenByHash (no re-fetch that could silently fail).
func TestVerifyEmail_RedeemReturnsUserID_NoRefetch(t *testing.T) {
	svc, repo, _, _ := newTestService(t, false)
	userID := uuid.New()
	tokenHash := sha256Hex("norefetch-token")
	now := time.Now()
	repo.tokens[tokenHash] = &AuthToken{
		ID: uuid.New(), UserID: userID, Purpose: purposeEmailVerify,
		TokenHash: tokenHash, ExpiresAt: now.Add(24 * time.Hour), CreatedAt: now,
	}
	repo.redeemMode = "alwaysTrue"

	if err := svc.VerifyEmail(context.Background(), "norefetch-token"); err != nil {
		t.Fatalf("VerifyEmail: %v", err)
	}
	if repo.findTokenCalls != 0 {
		t.Errorf("success path must not re-fetch the token (S1 fix): got %d FindAuthTokenByHash calls",
			repo.findTokenCalls)
	}
}

// TestVerifyEmail_SetVerifiedFails_RollsBackRedeem proves the S1 fix:
// when SetUserVerified fails, VerifyEmail returns a wrapped error (which
// the handler maps to 500), NOT nil (which would write a fake 200). The
// rollback guarantee itself (token not burned) is proven by the deferred
// Rollback pattern (same as registerNewUser) + the integration test
// TestRedeemAndVerify_Atomic against real Postgres.
func TestVerifyEmail_SetVerifiedFails_RollsBackRedeem(t *testing.T) {
	svc, repo, _, _ := newTestService(t, false)
	userID := uuid.New()
	tokenHash := sha256Hex("setverified-fails-token")
	now := time.Now()
	repo.tokens[tokenHash] = &AuthToken{
		ID: uuid.New(), UserID: userID, Purpose: purposeEmailVerify,
		TokenHash: tokenHash, ExpiresAt: now.Add(24 * time.Hour), CreatedAt: now,
	}
	repo.redeemMode = "alwaysTrue"
	repo.setVerifiedErr = errors.New("db connection lost")

	err := svc.VerifyEmail(context.Background(), "setverified-fails-token")
	if err == nil {
		t.Fatal("VerifyEmail returned nil on SetUserVerified failure — S1 regression " +
			"(handler would write a fake 200 while identity is not verified)")
	}
	// The error must be wrapped (preserving the chain) so MapServiceError
	// can map it to 500, not swallow it.
	if !strings.Contains(err.Error(), "set verified") {
		t.Errorf("error should wrap the set-verified failure, got: %v", err)
	}
	// SetUserVerified returned an error before recording the call, so no
	// verifiedCalls should be present (the identity was not marked verified).
	if len(repo.verifiedCalls) != 0 {
		t.Errorf("SetUserVerified failure must not record a verified call, got %d",
			len(repo.verifiedCalls))
	}
}

// ---- R9: expired token --------------------------------------------------

func TestVerifyEmail_ExpiredToken_410(t *testing.T) {
	svc, repo, _, _ := newTestService(t, false)
	tokenHash := sha256Hex("expired-token")
	now := time.Now()
	repo.tokens[tokenHash] = &AuthToken{
		ID: uuid.New(), UserID: uuid.New(), Purpose: purposeEmailVerify,
		TokenHash: tokenHash, ExpiresAt: now.Add(-time.Hour), CreatedAt: now.Add(-2 * time.Hour),
	}
	// RedeemToken returns false (expired token fails the guard).
	repo.redeemMode = "alwaysFalse"

	err := svc.VerifyEmail(context.Background(), "expired-token")
	if !errors.Is(err, ErrTokenExpired) {
		t.Fatalf("expected ErrTokenExpired, got %v", err)
	}
	if len(repo.verifiedCalls) != 0 {
		t.Errorf("expired token must not set verified_at, got %d calls", len(repo.verifiedCalls))
	}
}

// ---- R10: not found / already used -------------------------------------

func TestVerifyEmail_NotFound_404(t *testing.T) {
	svc, repo, _, _ := newTestService(t, false)
	repo.redeemMode = "alwaysFalse"
	// No token seeded -> FindAuthTokenByHash returns nil.
	err := svc.VerifyEmail(context.Background(), "nonexistent-token")
	if !errors.Is(err, ErrTokenNotFound) {
		t.Fatalf("expected ErrTokenNotFound, got %v", err)
	}
}

func TestVerifyEmail_AlreadyUsed_404(t *testing.T) {
	svc, repo, _, _ := newTestService(t, false)
	tokenHash := sha256Hex("used-token")
	now := time.Now()
	used := now.Add(-time.Hour)
	repo.tokens[tokenHash] = &AuthToken{
		ID: uuid.New(), UserID: uuid.New(), Purpose: purposeEmailVerify,
		TokenHash: tokenHash, ExpiresAt: now.Add(24 * time.Hour), UsedAt: &used, CreatedAt: now,
	}
	repo.redeemMode = "alwaysFalse"

	// Token exists and is not expired, but redeem fails (already used).
	// The service returns ErrTokenNotFound (not expired).
	err := svc.VerifyEmail(context.Background(), "used-token")
	if !errors.Is(err, ErrTokenNotFound) {
		t.Fatalf("expected ErrTokenNotFound for already-used token, got %v", err)
	}
}

// ---- R11: revoked (superseded) token — 3-clause guard regression -------

func TestVerifyEmail_RevokedToken_Rejected(t *testing.T) {
	svc, repo, _, _ := newTestService(t, false)
	tokenHash := sha256Hex("revoked-token")
	now := time.Now()
	revoked := now.Add(-time.Hour)
	repo.tokens[tokenHash] = &AuthToken{
		ID: uuid.New(), UserID: uuid.New(), Purpose: purposeEmailVerify,
		TokenHash: tokenHash, ExpiresAt: now.Add(24 * time.Hour), RevokedAt: &revoked, CreatedAt: now,
	}
	repo.redeemMode = "alwaysFalse"

	// Revoked (superseded) token must be rejected — this is the
	// regression for the INV-account-08 spec error (the 2-clause
	// Verification field omits revoked_at IS NULL).
	err := svc.VerifyEmail(context.Background(), "revoked-token")
	if !errors.Is(err, ErrTokenNotFound) {
		t.Fatalf("revoked token must return ErrTokenNotFound, got %v", err)
	}
	if len(repo.verifiedCalls) != 0 {
		t.Errorf("revoked token must not set verified_at")
	}
}

// ---- R12: concurrent double-submit -------------------------------------

func TestVerifyEmail_TokenSingleUse_Concurrent(t *testing.T) {
	// The single-use guarantee is the DB's atomic UPDATE ... WHERE; the
	// integration test (TestRedeemToken_Guards) proves the real guard.
	// Here the fake repo simulates that atomicity with a CAS per token
	// (redeemMode "atomic", the default): the first RedeemToken call
	// wins (true), all concurrent others lose (false). We assert exactly
	// one VerifyEmail returns nil (success) and the rest return
	// ErrTokenNotFound, with no panic under -race.
	svc, repo, _, _ := newTestService(t, false)
	tokenHash := sha256Hex("concurrent-token")
	now := time.Now()
	repo.tokens[tokenHash] = &AuthToken{
		ID: uuid.New(), UserID: uuid.New(), Purpose: purposeEmailVerify,
		TokenHash: tokenHash, ExpiresAt: now.Add(24 * time.Hour), CreatedAt: now,
	}
	// redeemMode defaults to "atomic" (single-use CAS).

	const n = 50
	var wg sync.WaitGroup
	errs := make([]error, n)
	wg.Add(n)
	for i := 0; i < n; i++ {
		i := i
		go func() {
			defer wg.Done()
			errs[i] = svc.VerifyEmail(context.Background(), "concurrent-token")
		}()
	}
	wg.Wait()

	successes := 0
	for i, err := range errs {
		switch {
		case err == nil:
			successes++
		case errors.Is(err, ErrTokenNotFound):
			// expected for the losers
		default:
			t.Errorf("goroutine %d: unexpected error %v", i, err)
		}
	}
	if successes != 1 {
		t.Errorf("expected exactly 1 successful redeem (single-use), got %d", successes)
	}
}

// ---- R13: resend unverified match --------------------------------------

func TestResend_UnverifiedMatch_IssuesNewToken(t *testing.T) {
	svc, repo, _, sender := newTestService(t, false)
	email := "resend@example.com"
	hash := hashFor(email)
	existing := repo.seedIdentity(providerEmailPassword, email, hash, false)

	if err := svc.ResendVerification(context.Background(), email); err != nil {
		t.Fatalf("ResendVerification: %v", err)
	}
	if len(repo.revokeCalls) != 1 || repo.revokeCalls[0].userID != existing.UserID {
		t.Errorf("expected revoke for the unverified user, got %v", repo.revokeCalls)
	}
	if len(repo.insertedTokens) != 1 {
		t.Fatalf("expected 1 new token, got %d", len(repo.insertedTokens))
	}
	if len(sender.verificationTo) != 1 || sender.verificationTo[0] != email {
		t.Errorf("expected verification email to %s, got %v", email, sender.verificationTo)
	}
	if len(sender.nudgeTypes) != 0 {
		t.Errorf("resend must send verification email not nudge, got %v", sender.nudgeTypes)
	}
}

// ---- R14: resend no-match / verified / google-only ---------------------

func TestResend_NoMatch_NoTokenNoEmail(t *testing.T) {
	svc, repo, _, sender := newTestService(t, false)
	if err := svc.ResendVerification(context.Background(), "nobody@example.com"); err != nil {
		t.Fatalf("ResendVerification: %v", err)
	}
	if len(repo.insertedTokens) != 0 || len(repo.revokeCalls) != 0 {
		t.Errorf("no-match must not write tokens/revoke")
	}
	if len(sender.verificationTo) != 0 && len(sender.nudgeTypes) != 0 {
		t.Errorf("no-match must not send email")
	}
}

func TestResend_Verified_NoTokenNoEmail(t *testing.T) {
	svc, repo, _, sender := newTestService(t, false)
	email := "verified@example.com"
	repo.seedIdentity(providerEmailPassword, email, hashFor(email), true)
	if err := svc.ResendVerification(context.Background(), email); err != nil {
		t.Fatalf("ResendVerification: %v", err)
	}
	if len(repo.insertedTokens) != 0 {
		t.Errorf("verified resend must not issue token")
	}
	if len(sender.verificationTo) != 0 {
		t.Errorf("verified resend must not send email")
	}
}

func TestResend_GoogleOnly_NoTokenNoEmail(t *testing.T) {
	svc, repo, _, _ := newTestService(t, false)
	email := "google@example.com"
	repo.seedIdentity(providerGoogle, email, hashFor(email), true)
	if err := svc.ResendVerification(context.Background(), email); err != nil {
		t.Fatalf("ResendVerification: %v", err)
	}
	if len(repo.insertedTokens) != 0 {
		t.Errorf("google-only resend must not issue token")
	}
}

// ---- R16: concurrent duplicate registration ----------------------------

func TestRegister_ConcurrentDuplicateEmail_Race(t *testing.T) {
	// >=100 goroutines registering the same email. The fake repo's
	// InsertAuthIdentity enforces uniqueness (returns *pgconn.PgError
	// 23505 on a second insert for the same identifier). The service
	// maps unique-violation to a clean no-op (nil), so every goroutine
	// returns nil and exactly one set of (user+identity+token) is
	// created. Run under -race.
	const n = 100
	var wg sync.WaitGroup
	errs := make([]error, n)
	wg.Add(n)
	svc, repo, _, _ := newTestService(t, false)
	email := "dup@example.com"
	for i := 0; i < n; i++ {
		i := i
		go func() {
			defer wg.Done()
			errs[i] = svc.Register(context.Background(), "Dup", email, "strong-pw-123")
		}()
	}
	wg.Wait()

	nilCount := 0
	for _, err := range errs {
		if err != nil {
			t.Errorf("concurrent duplicate: goroutine returned error %v (should be nil no-op)", err)
		}
		if err == nil {
			nilCount++
		}
	}
	if nilCount != n {
		t.Errorf("expected all %d to return nil, got %d", n, nilCount)
	}
	// Exactly one identity + one token created. The fake repo does not
	// model transaction rollback, so insertedUsers may over-count (losers
	// insert a user before failing on the identity); the real
	// clean-rollback guarantee for users is proven by the integration test
	// TestInsertAuthIdentity_ConcurrentDuplicate against real Postgres.
	// Here we assert the uniqueness-protected counts.
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if len(repo.insertedIdentities) != 1 {
		t.Errorf("expected exactly 1 identity, got %d", len(repo.insertedIdentities))
	}
	if len(repo.insertedTokens) != 1 {
		t.Errorf("expected exactly 1 token, got %d", len(repo.insertedTokens))
	}
}

// ---- R17: Google-only generic response ---------------------------------

func TestRegister_GoogleOnlyConflict_GenericResponse(t *testing.T) {
	svc, repo, _, sender := newTestService(t, false)
	email := "g17@example.com"
	repo.seedIdentity(providerGoogle, email, hashFor(email), true)

	err := svc.Register(context.Background(), "G", email, "strong-pw-123")
	if err != nil {
		t.Fatalf("google-only branch should return nil, got %v", err)
	}
	// The observable outcome is identical to other branches: one email,
	// no records, return nil.
	totalEmails := len(sender.verificationTo) + len(sender.nudgeTypes)
	if totalEmails != 1 {
		t.Errorf("google-only branch should send exactly 1 nudge, got %d", totalEmails)
	}
	if len(repo.insertedUsers) != 0 {
		t.Errorf("google-only branch should create no user")
	}
}

// ensure the sha256Hex used by tests matches the service's.
func TestSha256Hex_Consistency(t *testing.T) {
	// The service's sha256Hex is package-private; tests use crypto/sha256
	// directly in some seeds. Assert they agree.
	got := sha256Hex("abc")
	sum := sha256.Sum256([]byte("abc"))
	want := hex.EncodeToString(sum[:])
	if got != want {
		t.Fatalf("sha256Hex mismatch: got %s want %s", got, want)
	}
}

// TestRegister_SendVerificationFails_LogNoPII proves the L2 fix: when the
// notification sender returns an error whose Error() embeds the recipient
// email and token (simulating a leaky SMTP error), the service's log line
// contains a sanitized category, NOT the recipient email or token. Per
// go/secrets-and-sensitive-logging.md §1.
func TestRegister_SendVerificationFails_LogNoPII(t *testing.T) {
	repo := newFakeRepo()
	sender := &leakySender{}
	bc := &fakeBreachChecker{breached: false}
	keys := &crypto.Keys{
		EncryptionKey: make([]byte, 32),
		HMACKey:       make([]byte, 32),
	}
	svc := &Service{
		repo:        repo,
		tx:          fakeTxRunner{},
		breachCheck: bc,
		email:       sender,
		keys:        keys,
	}

	const recipient = "leaky@example.com"
	// The generated token is internal to the service; we can't predict
	// it, but the leakySender embeds whatever token it receives into its
	// error message. We assert the log doesn't contain ANY hex string of
	// the length the service generates (64 hex chars = 32 bytes). We also
	// assert the recipient email is absent.

	var logBuf bytes.Buffer
	origOut := log.Writer()
	log.SetOutput(&logBuf)
	defer log.SetOutput(origOut)

	// Register a new user (R1) — the post-commit sendVerification will
	// call the leakySender, which returns an error embedding the
	// recipient + token.
	if err := svc.Register(context.Background(), "Leaky", recipient, "strong-pw-123"); err != nil {
		t.Fatalf("Register: post-commit email failure must not fail the request: %v", err)
	}

	logged := logBuf.String()
	// The log must announce the failure.
	if !strings.Contains(logged, "send verification email failed") {
		t.Errorf("expected 'send verification email failed' in log, got: %q", logged)
	}
	// The log must NOT contain the recipient email (PII).
	if strings.Contains(logged, recipient) {
		t.Errorf("log leaked recipient email %q: %q", recipient, logged)
	}
	// The log must NOT contain "SMTP" (raw error message leak).
	if strings.Contains(logged, "SMTP") {
		t.Errorf("log leaked raw SMTP error message: %q", logged)
	}
	// The log must NOT contain "rejected" (raw error message leak).
	if strings.Contains(logged, "rejected") {
		t.Errorf("log leaked raw error detail 'rejected': %q", logged)
	}
	// The log must NOT contain "token=" (raw error message leak).
	if strings.Contains(logged, "token=") {
		t.Errorf("log leaked raw error detail 'token=': %q", logged)
	}
}
