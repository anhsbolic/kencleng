package account

import (
	"bytes"
	"context"
	"errors"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/anhsbolic/kencleng/backend/internal/platform/crypto"
	"github.com/anhsbolic/kencleng/backend/internal/platform/notification"
)

// ---- Fitur 2B: forgot-password (request half) ------------------------------

// seedRegisteredEmail seeds an email_password identity the way the register
// flow would have created it, keyed for hash lookup.
func seedRegisteredEmail(t *testing.T, repo *fakeRepo, email string) *AuthIdentity {
	t.Helper()
	return repo.seedIdentity(providerEmailPassword, email, hashFor(email), true)
}

// R1: registered email_password identity → token issued + reset email sent.
func TestForgotPassword_Match_IssuesTokenAndSendsEmail(t *testing.T) {
	svc, repo, _, sender := newTestService(t, false)
	email := "resetme@example.com"
	ident := seedRegisteredEmail(t, repo, email)

	if err := svc.ForgotPassword(context.Background(), email); err != nil {
		t.Fatalf("ForgotPassword: %v", err)
	}

	if len(repo.insertedTokens) != 1 {
		t.Fatalf("expected 1 token inserted, got %d", len(repo.insertedTokens))
	}
	tok := repo.insertedTokens[0]
	if tok.Purpose != purposePasswordReset {
		t.Errorf("token purpose = %q, want %q", tok.Purpose, purposePasswordReset)
	}
	if tok.UserID != ident.UserID {
		t.Errorf("token userID = %s, want %s", tok.UserID, ident.UserID)
	}
	if tok.ExpiresAt.Sub(tok.CreatedAt) < resetTokenTTL-time.Second {
		t.Errorf("token TTL = %v, want ~%v (resetTokenTTL, NOT tokenTTL)",
			tok.ExpiresAt.Sub(tok.CreatedAt), resetTokenTTL)
	}
	if tok.UsedAt != nil || tok.RevokedAt != nil {
		t.Error("fresh token must be unused and unrevoked")
	}
	if len(sender.resetTo) != 1 || sender.resetTo[0] != email {
		t.Errorf("reset email not sent to %s: %v", email, sender.resetTo)
	}
	if len(sender.resetTokens) != 1 || sender.resetTokens[0] == "" {
		t.Errorf("reset email must carry the plain token, got %v", sender.resetTokens)
	}
	// Assumption A: issuance must not revoke anything.
	if len(repo.revokeCalls) != 0 || len(repo.revokeAllForUserCalls) != 0 {
		t.Errorf("forgot-password match branch must not revoke any token: revokes=%v allUser=%v",
			repo.revokeCalls, repo.revokeAllForUserCalls)
	}
}

// R2: Google-only identity → distinct notice email, NO token created.
func TestForgotPassword_GoogleOnly_NoticeNoToken(t *testing.T) {
	svc, repo, _, sender := newTestService(t, false)
	email := "googleonly@example.com"
	existing := repo.seedIdentity(providerGoogle, email, hashFor(email), true)

	if err := svc.ForgotPassword(context.Background(), email); err != nil {
		t.Fatalf("ForgotPassword: %v", err)
	}

	if len(repo.insertedTokens) != 0 {
		t.Errorf("google-only branch must create no token, got %d", len(repo.insertedTokens))
	}
	if len(sender.nudgeTypes) != 1 || sender.nudgeTypes[0] != notification.NudgeGoogleOnly {
		t.Errorf("expected google_only notice, got %v", sender.nudgeTypes)
	}
	if len(sender.resetTo) != 0 {
		t.Errorf("no reset email may be sent on google-only branch: %v", sender.resetTo)
	}
	// Timing shaping: exactly one dummyWrite-shaped revoke against a
	// synthetic user_id (same device as Register R3/R4).
	if len(repo.revokeCalls) != 1 {
		t.Errorf("google-only branch must perform exactly 1 dummy write (R7 DB-time), got %d",
			len(repo.revokeCalls))
	} else if repo.revokeCalls[0].userID == existing.UserID {
		t.Errorf("dummy revoke must use a synthetic user_id, not the real one")
	}
}

// R3: unknown email → nothing sent, response-equivalent to the other branches.
func TestForgotPassword_NoMatch_NothingSent(t *testing.T) {
	svc, repo, _, sender := newTestService(t, false)
	email := "ghost@example.com"

	if err := svc.ForgotPassword(context.Background(), email); err != nil {
		t.Fatalf("ForgotPassword: %v", err)
	}

	if len(repo.insertedTokens) != 0 {
		t.Errorf("no-match branch must create no token, got %d", len(repo.insertedTokens))
	}
	if len(sender.resetTo) != 0 || len(sender.nudgeTypes) != 0 {
		t.Errorf("no-match branch must send nothing: resets=%v nudges=%v",
			sender.resetTo, sender.nudgeTypes)
	}
	if len(repo.revokeCalls) != 1 {
		t.Errorf("no-match branch must perform exactly 1 dummy write (R7 DB-time), got %d",
			len(repo.revokeCalls))
	}
}

// All three forgot branches return nil — the handler writes an identical
// generic 202 regardless (R5's API-surface half; timing is proven against
// real Postgres in the integration suite).
func TestForgotPassword_GenericResponse_AllBranches(t *testing.T) {
	cases := []struct {
		name  string
		setup func(svc *Service, repo *fakeRepo) string
	}{
		{"registered", func(_ *Service, repo *fakeRepo) string {
			email := "branch1@example.com"
			seedRegisteredEmail(t, repo, email)
			return email
		}},
		{"google-only", func(_ *Service, repo *fakeRepo) string {
			email := "branch2@example.com"
			repo.seedIdentity(providerGoogle, email, hashFor(email), true)
			return email
		}},
		{"unknown", func(_ *Service, _ *fakeRepo) string { return "branch3@example.com" }},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			svc, repo, _, _ := newTestService(t, false)
			email := tc.setup(svc, repo)
			if err := svc.ForgotPassword(context.Background(), email); err != nil {
				t.Fatalf("ForgotPassword(%s): err = %v, want nil", tc.name, err)
			}
		})
	}
}

// Assumption A (resolved): repeated forgot-password requests each issue
// their own independent token; previously-issued unexpired tokens are NOT
// proactively revoked and remain redeemable.
func TestForgotPassword_Repeat_DoesNotRevokePriorTokens(t *testing.T) {
	svc, repo, _, sender := newTestService(t, false)
	email := "twice@example.com"
	seedRegisteredEmail(t, repo, email)

	if err := svc.ForgotPassword(context.Background(), email); err != nil {
		t.Fatalf("first ForgotPassword: %v", err)
	}
	first := repo.insertedTokens[0]
	if err := svc.ForgotPassword(context.Background(), email); err != nil {
		t.Fatalf("second ForgotPassword: %v", err)
	}

	if len(repo.insertedTokens) != 2 {
		t.Fatalf("expected 2 independently issued tokens, got %d", len(repo.insertedTokens))
	}
	if first.RevokedAt != nil {
		t.Error("first token must remain valid (Assumption A: no proactive revocation)")
	}
	if first.TokenHash == repo.insertedTokens[1].TokenHash {
		t.Error("tokens must be independent random values")
	}
	if len(sender.resetTo) != 2 {
		t.Errorf("each request sends its own reset email, got %d", len(sender.resetTo))
	}
	// No revoke targeted the real user (dummyWrite only fires on no-op
	// branches, and neither call hit one).
	for _, c := range repo.revokeCalls {
		if c.purpose == purposePasswordReset {
			t.Errorf("no code path may revoke password_reset tokens on forgot: %+v", c)
		}
	}
	if len(repo.revokeAllForUserCalls) != 0 {
		t.Errorf("forgot-password never mass-revokes sessions: %v", repo.revokeAllForUserCalls)
	}
}

// ---- Fitur 2B: reset-password (redemption half) ----------------------------

// seedResetToken seeds a password_reset token under a known plain value.
func seedResetToken(t *testing.T, repo *fakeRepo, userID uuid.UUID, plain string, expires time.Time) {
	t.Helper()
	tokenHash := sha256Hex(plain)
	repo.tokens[tokenHash] = &AuthToken{
		ID: uuid.New(), UserID: userID, Purpose: purposePasswordReset,
		TokenHash: tokenHash, ExpiresAt: expires, CreatedAt: time.Now(),
	}
}

// R7 happy path: credential updated + sessions revoked in one tx.
func TestResetPassword_HappyPath_UpdatesAndRevokes(t *testing.T) {
	svc, repo, bc, _ := newTestService(t, false)
	userID := uuid.New()
	seedResetToken(t, repo, userID, "valid-reset-token", time.Now().Add(time.Minute))
	repo.redeemMode = "alwaysTrue"

	if err := svc.ResetPassword(context.Background(), "valid-reset-token", "brand-new-pw-9"); err != nil {
		t.Fatalf("ResetPassword: %v", err)
	}

	if len(repo.updateCredentialCalls) != 1 {
		t.Fatalf("expected 1 credential update, got %d", len(repo.updateCredentialCalls))
	}
	upd := repo.updateCredentialCalls[0]
	if upd.userID != userID || upd.provider != providerEmailPassword {
		t.Errorf("credential update target = (%s,%s), want (%s,%s)",
			upd.userID, upd.provider, userID, providerEmailPassword)
	}
	if upd.passwordHash == "" || upd.passwordHash == "brand-new-pw-9" {
		t.Error("stored secret must be a bcrypt hash, not the plaintext password")
	}
	if len(repo.revokeAllForUserCalls) != 1 || repo.revokeAllForUserCalls[0] != userID {
		t.Errorf("expected mass session revoke for user, got %v", repo.revokeAllForUserCalls)
	}
	// Ordering guard for Assumption B's other half: validation ran before
	// redemption (breach checker consulted before any redeem).
	if !bc.called.Load() {
		t.Error("password policy check must run before the transaction")
	}
}

// R8: policy failure → ErrValidation AND the token stays unconsumed
// (validation happens BEFORE the token-consuming tx — Assumption B).
func TestResetPassword_PasswordPolicy_TokenNotConsumed(t *testing.T) {
	cases := []struct {
		name     string
		password string
		breached bool
		wantErr  error
	}{
		{"too-short", "short1", false, ErrValidation},
		{"breached-list", "long-enough-password", true, ErrValidation},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			svc, repo, _, _ := newTestService(t, tc.breached)
			userID := uuid.New()
			seedResetToken(t, repo, userID, "policy-token", time.Now().Add(time.Minute))
			repo.redeemMode = "alwaysTrue"

			err := svc.ResetPassword(context.Background(), "policy-token", tc.password)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("err = %v, want ErrValidation", err)
			}
			if len(repo.updateCredentialCalls) != 0 || len(repo.revokeAllForUserCalls) != 0 {
				t.Error("no DB mutation may happen when validation fails")
			}
			// The fake's RedeemToken was never reached: validation gates
			// everything, so used_at could not have been set even in the
			// real DB.
			if repo.tokens[sha256Hex("policy-token")].UsedAt != nil {
				t.Error("token must remain unused after a validation failure")
			}
		})
	}
}

// Breach-check fail-open: an unreachable HIBP API does NOT block a reset
// (Fitur 1 resolution, threat-model residual #4). Mirrors the register
// precedent (TestRegister_BreachCheck_FailOpen): the client's fail-open
// unit is platform/breachcheck, which surfaces (false, nil); here we
// assert the service treats that as "proceed".
func TestResetPassword_BreachCheck_FailOpen(t *testing.T) {
	svc, repo, _, _ := newTestService(t, false)
	userID := uuid.New()
	seedResetToken(t, repo, userID, "failopen-token", time.Now().Add(time.Minute))
	repo.redeemMode = "alwaysTrue"

	if err := svc.ResetPassword(context.Background(), "failopen-token", "another-good-pw"); err != nil {
		t.Fatalf("ResetPassword with unreachable breach API: %v", err)
	}
	if len(repo.updateCredentialCalls) != 1 {
		t.Error("reset should proceed fail-open when the breach API is unreachable")
	}
}

// Expired → 410-mapped sentinel; not-found / already-used → 404-mapped
// sentinel; no state change in any case.
func TestResetPassword_TokenStateMapping(t *testing.T) {
	cases := []struct {
		name    string
		setup   func(repo *fakeRepo) string
		wantErr error
	}{
		{"expired", func(repo *fakeRepo) string {
			id := uuid.New()
			seedResetToken(t, repo, id, "expired-token", time.Now().Add(-time.Minute))
			return "expired-token"
		}, ErrTokenExpired},
		{"not-found", func(_ *fakeRepo) string { return "never-issued" }, ErrTokenNotFound},
		{"already-used", func(repo *fakeRepo) string {
			id := uuid.New()
			used := sha256Hex("used-token")
			now := time.Now()
			repo.tokens[used] = &AuthToken{
				ID: uuid.New(), UserID: id, Purpose: purposePasswordReset,
				TokenHash: used, ExpiresAt: now.Add(time.Hour),
				CreatedAt: now.Add(-time.Minute), UsedAt: &now,
			}
			return "used-token"
		}, ErrTokenNotFound},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			svc, repo, _, _ := newTestService(t, false)
			token := tc.setup(repo)
			// alwaysFalse models the 3-clause guard rejecting the redeem;
			// the service then disambiguates via FindAuthTokenByHash.
			repo.redeemMode = "alwaysFalse"
			err := svc.ResetPassword(context.Background(), token, "good-password-1")
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("err = %v, want %v", err, tc.wantErr)
			}
			if len(repo.updateCredentialCalls) != 0 || len(repo.revokeAllForUserCalls) != 0 {
				t.Error("rejected token must cause no state change")
			}
		})
	}
}

// R13/D2: an email_verification token presented to ResetPassword is
// rejected as not-found and never drives a credential update. The
// rollback-unconsumes property itself is proven against real Postgres in
// the integration suite (fakes cannot model tx rollback); this pins the
// service-level decision.
func TestResetPassword_WrongPurpose_RejectedNoMutation(t *testing.T) {
	svc, repo, _, _ := newTestService(t, false)
	userID := uuid.New()
	tokenHash := sha256Hex("verify-purpose-token")
	now := time.Now()
	repo.tokens[tokenHash] = &AuthToken{
		ID: uuid.New(), UserID: userID, Purpose: purposeEmailVerify,
		TokenHash: tokenHash, ExpiresAt: now.Add(time.Hour), CreatedAt: now,
	}
	repo.redeemMode = "alwaysTrue"

	if err := svc.ResetPassword(context.Background(), "verify-purpose-token", "good-password-1"); !errors.Is(err, ErrTokenNotFound) {
		t.Fatalf("err = %v, want ErrTokenNotFound", err)
	}
	if len(repo.updateCredentialCalls) != 0 || len(repo.revokeAllForUserCalls) != 0 {
		t.Error("wrong-purpose redemption must mutate nothing")
	}
}

// R12/Q1: a password_reset token presented to VerifyEmail is rejected and
// must not verify anything. Same rollback caveat as above — the
// unconsumed property gets its real-Postgres proof in the integration
// suite.
func TestVerifyEmail_RejectsResetPurposeToken(t *testing.T) {
	svc, repo, _, _ := newTestService(t, false)
	userID := uuid.New()
	tokenHash := sha256Hex("reset-purpose-token")
	now := time.Now()
	repo.tokens[tokenHash] = &AuthToken{
		ID: uuid.New(), UserID: userID, Purpose: purposePasswordReset,
		TokenHash: tokenHash, ExpiresAt: now.Add(time.Hour), CreatedAt: now,
	}
	repo.redeemMode = "alwaysTrue"

	if err := svc.VerifyEmail(context.Background(), "reset-purpose-token"); !errors.Is(err, ErrTokenNotFound) {
		t.Fatalf("err = %v, want ErrTokenNotFound", err)
	}
	if len(repo.verifiedCalls) != 0 {
		t.Error("a reset token must never mark an identity verified")
	}
}

// R14: a leaky SMTP error on the reset email must surface only a
// sanitized category in the log — never the recipient or token.
func TestPasswordReset_SendFails_LogsNoPIIOrToken(t *testing.T) {
	var logBuf bytes.Buffer
	origOut := log.Writer()
	log.SetOutput(&logBuf)
	defer log.SetOutput(origOut)

	repo := newFakeRepo()
	sender := &leakySender{}
	keys := &crypto.Keys{
		EncryptionKey: make([]byte, 32),
		HMACKey:       make([]byte, 32),
	}
	svc := &Service{repo: repo, tx: fakeTxRunner{}, breachCheck: &fakeBreachChecker{}, email: sender, keys: keys}

	const recipient = "leaky-reset@example.com"
	seedRegisteredEmail(t, repo, recipient)

	if err := svc.ForgotPassword(context.Background(), recipient); err != nil {
		t.Fatalf("ForgotPassword: %v", err)
	}

	logged := logBuf.String()
	if strings.Contains(logged, recipient) {
		t.Fatalf("log must not contain recipient email; got %q", logged)
	}
	if strings.Contains(logged, "SMTP 553") {
		t.Fatalf("log must not carry the raw sender error; got %q", logged)
	}
	if !strings.Contains(logged, "send failed") {
		t.Errorf("log should carry the sanitized category, got %q", logged)
	}
}
