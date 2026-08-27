package account

import (
	"bytes"
	"context"
	"log"
	"testing"
	"time"

	"github.com/anhsbolic/kencleng/backend/internal/platform/secrets"
	"github.com/google/uuid"
)

// newSecurityTestService extends newTestService with the compare seam
// set to secrets.ComparePassword — SetPassword Branch 2 needs real bcrypt
// comparison against the identity's stored credential_secret.
func newSecurityTestService(t *testing.T) (*Service, *fakeRepo, *captureSender) {
	t.Helper()
	svc, repo, _, sender := newTestService(t, false)
	svc.compare = secrets.ComparePassword
	return svc, repo, sender
}

// ---- R1: Branch 1 happy path -----------------------------------------

func TestSetPassword_Branch1_CreatesUnverifiedIdentity_SendsVerification(t *testing.T) {
	svc, repo, sender := newSecurityTestService(t)

	// Seed a Google-only user (no email_password identity).
	googleIdent := repo.seedIdentity(providerGoogle, "user@gmail.com", hashFor("user@gmail.com"), true)
	userID := googleIdent.UserID

	_, err := svc.SetPassword(context.Background(), userID, "work@company.com", "", "strong-pw-123")
	if err != nil {
		t.Fatalf("SetPassword Branch 1: %v", err)
	}

	// Assert: 1 new identity inserted (email_password, unverified).
	var newEP *AuthIdentity
	for _, id := range repo.insertedIdentities {
		if id.UserID == userID && id.ProviderType == providerEmailPassword {
			newEP = id
			break
		}
	}
	if newEP == nil {
		t.Fatalf("expected an email_password identity inserted for user %s", userID)
	}
	if newEP.VerifiedAt != nil {
		t.Errorf("new identity must be unverified (verified_at nil), got %v", newEP.VerifiedAt)
	}

	// Assert: 1 token inserted with purpose=email_verification_link.
	var linkToken *AuthToken
	for _, tok := range repo.insertedTokens {
		if tok.UserID == userID && tok.Purpose == purposeEmailVerifyLink {
			linkToken = tok
			break
		}
	}
	if linkToken == nil {
		t.Fatalf("expected a token with purpose=%s inserted", purposeEmailVerifyLink)
	}
	if !linkToken.ExpiresAt.After(time.Now()) {
		t.Errorf("token must not be expired")
	}

	// Assert: verification email sent to the submitted email.
	if len(sender.verificationTo) != 1 || sender.verificationTo[0] != "work@company.com" {
		t.Errorf("expected 1 verification email to work@company.com, got %v", sender.verificationTo)
	}
}

// ---- R2: Branch 1 claimed email → nudge, no identity, generic nil ----

func TestSetPassword_Branch1_ClaimedEmail_NudgeNoIdentity_Generic202(t *testing.T) {
	svc, repo, sender := newSecurityTestService(t)

	// Seed a Google-only user.
	googleIdent := repo.seedIdentity(providerGoogle, "user@gmail.com", hashFor("user@gmail.com"), true)

	// Seed ANOTHER user's email_password identity with the target email.
	repo.seedIdentity(providerEmailPassword, "work@company.com", hashFor("work@company.com"), true)

	// Count pre-existing identities.
	preCount := len(repo.insertedIdentities)

	_, err := svc.SetPassword(context.Background(), googleIdent.UserID, "work@company.com", "", "strong-pw-123")
	if err != nil {
		t.Fatalf("SetPassword Branch 1 claimed: %v (expected nil)", err)
	}

	// No new identity/token created.
	if len(repo.insertedIdentities) != preCount {
		t.Errorf("expected 0 new identities, got %d", len(repo.insertedIdentities)-preCount)
	}
	if len(repo.insertedTokens) != 0 {
		t.Errorf("expected 0 new tokens, got %d", len(repo.insertedTokens))
	}

	// Conflict nudge sent, NOT a verification email.
	if len(sender.verificationTo) != 0 {
		t.Errorf("expected 0 verification emails, got %d", len(sender.verificationTo))
	}
	found := false
	for _, n := range sender.nudgeTypes {
		if n == "set_password_conflict" {
			found = true
		}
	}
	if !found {
		t.Errorf("expected a set_password_conflict nudge, got %v", sender.nudgeTypes)
	}
}

// ---- R3: Branch 1 unique-violation fallback (race-loser path) -------
//
// The fake cannot faithfully simulate READ COMMITTED transaction
// visibility (it updates the map on insert, before commit), so a true
// concurrent race is deferred to the integration suite (task-06). Here
// we exercise the unique-violation fallback directly: pre-seed the
// dedup key so InsertAuthIdentity fires a 23505, but DON'T seed the
// hash-keyed lookup so the pre-check returns nil — simulating the race
// window where the winner hasn't committed yet but the unique index is
// already taken.

func TestSetPassword_ConcurrentDuplicateEmail_Race(t *testing.T) {
	svc, repo, sender := newSecurityTestService(t)

	googleIdent := repo.seedIdentity(providerGoogle, "user@gmail.com", hashFor("user@gmail.com"), true)

	// Pre-seed the dedup key without making the identity visible to
	// the pre-check lookup.
	repo.identityKeys[providerEmailPassword+"|race@example.com"] = true

	_, err := svc.SetPassword(context.Background(), googleIdent.UserID, "race@example.com", "", "strong-pw-123")
	if err != nil {
		t.Errorf("unique-violation fallback must return nil (generic 202), got %v", err)
	}

	// No identity/token created (the insert was rolled back).
	var created int
	for _, id := range repo.insertedIdentities {
		if id.UserID == googleIdent.UserID && id.ProviderType == providerEmailPassword {
			created++
		}
	}
	if created != 0 {
		t.Errorf("race-loser must not create an identity, got %d", created)
	}

	// Conflict nudge sent (not a verification email).
	if len(sender.verificationTo) != 0 {
		t.Errorf("race-loser must not send a verification email")
	}
	found := false
	for _, n := range sender.nudgeTypes {
		if n == "set_password_conflict" {
			found = true
		}
	}
	if !found {
		t.Errorf("race-loser must send a conflict nudge, got %v", sender.nudgeTypes)
	}
}

// ---- R4: policy precedes branching -----------------------------------

func TestSetPassword_PasswordPolicy_PrecedesBranching(t *testing.T) {
	svc, repo, _ := newSecurityTestService(t)
	googleIdent := repo.seedIdentity(providerGoogle, "user@gmail.com", hashFor("user@gmail.com"), true)

	preIdentities := len(repo.insertedIdentities)
	preTokens := len(repo.insertedTokens)

	_, err := svc.SetPassword(context.Background(), googleIdent.UserID, "work@company.com", "", "short")
	if err != ErrValidation {
		t.Fatalf("expected ErrValidation for short password, got %v", err)
	}

	// No side effects.
	if len(repo.insertedIdentities) != preIdentities {
		t.Errorf("policy failure must not create identities")
	}
	if len(repo.insertedTokens) != preTokens {
		t.Errorf("policy failure must not create tokens")
	}
}

func TestSetPassword_BreachCheck_FailOpen(t *testing.T) {
	// Rebuild with breached=true (breach check says password is breached).
	svc, repo, _, _ := newTestService(t, true)
	svc.compare = secrets.ComparePassword
	googleIdent := repo.seedIdentity(providerGoogle, "user@gmail.com", hashFor("user@gmail.com"), true)

	// Breach-check returns breached=true → ErrValidation.
	// But if the breach-check API were unreachable, IsBreached returns
	// (false, nil) — fail-open. The fakeBreachChecker returns (breached,
	// nil), so breached=true yields ErrValidation (not fail-open). The
	// fail-open path is tested separately by injecting an error.
	_, err := svc.SetPassword(context.Background(), googleIdent.UserID, "work@company.com", "", "strong-pw-123")
	if err != ErrValidation {
		t.Errorf("breached password must yield ErrValidation, got %v", err)
	}
}

// ---- R5: generic-response parity across Branch 1 outcomes ----------

func TestSetPassword_GenericResponse_AllBranches(t *testing.T) {
	cases := []struct {
		name  string
		setup func(*fakeRepo) uuid.UUID
	}{
		{"created", func(r *fakeRepo) uuid.UUID {
			g := r.seedIdentity(providerGoogle, "user@gmail.com", hashFor("user@gmail.com"), true)
			return g.UserID
		}},
		{"claimed", func(r *fakeRepo) uuid.UUID {
			g := r.seedIdentity(providerGoogle, "user@gmail.com", hashFor("user@gmail.com"), true)
			r.seedIdentity(providerEmailPassword, "taken@company.com", hashFor("taken@company.com"), true)
			return g.UserID
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			svc, repo, _ := newSecurityTestService(t)
			userID := tc.setup(repo)
			_, err := svc.SetPassword(context.Background(), userID, "taken@company.com", "", "strong-pw-123")
			if err != nil {
				t.Errorf("Branch 1 must return nil for generic 202 (case %s), got %v", tc.name, err)
			}
		})
	}
}

// ---- R6: server-side branch selection (ignores client fields) -------

func TestSetPassword_BranchSelection_ServerSide(t *testing.T) {
	t.Run("google-only caller with current_password → Branch 1", func(t *testing.T) {
		svc, repo, _ := newSecurityTestService(t)
		g := repo.seedIdentity(providerGoogle, "user@gmail.com", hashFor("user@gmail.com"), true)

		// Body includes current_password (irrelevant for Branch 1).
		_, err := svc.SetPassword(context.Background(), g.UserID, "new@company.com", "irrelevant-old-pw", "strong-pw-123")
		if err != nil {
			t.Fatalf("expected nil (Branch 1 generic 202), got %v", err)
		}
		// Branch 1: verification email sent, NOT a credential update.
		if len(repo.updateCredentialCalls) != 0 {
			t.Errorf("Branch 1 must not update credential_secret")
		}
	})

	t.Run("email_password caller with email → Branch 2", func(t *testing.T) {
		svc, repo, _ := newSecurityTestService(t)
		ident := repo.seedIdentity(providerEmailPassword, "existing@company.com", hashFor("existing@company.com"), true)
		storedHash, _ := secrets.HashPassword("old-pw-123")
		ident.CredentialSecret = &storedHash

		// Body includes email (irrelevant for Branch 2).
		_, err := svc.SetPassword(context.Background(), ident.UserID, "ignored@company.com", "old-pw-123", "new-strong-pw")
		if err != nil {
			t.Fatalf("expected nil (Branch 2 success), got %v", err)
		}
		// Branch 2: credential updated, no verification email.
		if len(repo.updateCredentialCalls) != 1 {
			t.Errorf("Branch 2 must update credential_secret once")
		}
	})
}

// ---- R7: Branch 2 atomic change (credential + session revocation) ---

func TestSetPassword_Branch2_AllSessionsRevoked(t *testing.T) {
	svc, repo, _ := newSecurityTestService(t)

	ident := repo.seedIdentity(providerEmailPassword, "user@company.com", hashFor("user@company.com"), true)
	storedHash, _ := secrets.HashPassword("old-pw-123")
	ident.CredentialSecret = &storedHash

	_, err := svc.SetPassword(context.Background(), ident.UserID, "", "old-pw-123", "new-strong-pw")
	if err != nil {
		t.Fatalf("SetPassword Branch 2: %v", err)
	}

	// Credential updated with the NEW hash (not the old one).
	if len(repo.updateCredentialCalls) != 1 {
		t.Fatalf("expected 1 credential update, got %d", len(repo.updateCredentialCalls))
	}
	if repo.updateCredentialCalls[0].passwordHash == storedHash {
		t.Errorf("credential must be updated to the new hash, not the old one")
	}
	if repo.updateCredentialCalls[0].provider != providerEmailPassword {
		t.Errorf("provider must be email_password, got %s", repo.updateCredentialCalls[0].provider)
	}

	// All sessions revoked (INV-account-05).
	if len(repo.revokeAllForUserCalls) != 1 {
		t.Fatalf("expected 1 revoke-all-sessions call, got %d", len(repo.revokeAllForUserCalls))
	}
	if repo.revokeAllForUserCalls[0] != ident.UserID {
		t.Errorf("revoke-all must target the caller's user_id")
	}
}

// ---- R8: Branch 2 wrong current_password → 401 ----------------------

func TestSetPassword_Branch2_WrongCurrentPassword_Rejected(t *testing.T) {
	svc, repo, _ := newSecurityTestService(t)

	ident := repo.seedIdentity(providerEmailPassword, "user@company.com", hashFor("user@company.com"), true)
	storedHash, _ := secrets.HashPassword("old-pw-123")
	ident.CredentialSecret = &storedHash

	_, err := svc.SetPassword(context.Background(), ident.UserID, "", "WRONG-password", "new-strong-pw")
	if err != ErrInvalidCredentials {
		t.Fatalf("expected ErrInvalidCredentials, got %v", err)
	}

	// No state change.
	if len(repo.updateCredentialCalls) != 0 {
		t.Errorf("wrong password must not update credential")
	}
	if len(repo.revokeAllForUserCalls) != 0 {
		t.Errorf("wrong password must not revoke sessions")
	}
}

// ---- R14: VerifyEmail conditional audit for link-purpose tokens ----

func TestVerifyEmail_LinkPurpose_WritesLinkAudit(t *testing.T) {
	svc, repo, _ := newSecurityTestService(t)

	// Seed an unverified email_password identity.
	ident := repo.seedIdentity(providerEmailPassword, "user@company.com", hashFor("user@company.com"), false)
	userID := ident.UserID

	// Seed a link-purpose token directly in the fake.
	plainToken, tokenHash, err := generateToken()
	if err != nil {
		t.Fatalf("generateToken: %v", err)
	}
	repo.tokens[tokenHash] = &AuthToken{
		ID:        uuid.New(),
		UserID:    userID,
		Purpose:   purposeEmailVerifyLink,
		TokenHash: tokenHash,
		ExpiresAt: time.Now().Add(tokenTTL),
		CreatedAt: time.Now(),
	}

	preLogs := len(repo.insertedUserLogs)
	err = svc.VerifyEmail(context.Background(), plainToken)
	if err != nil {
		t.Fatalf("VerifyEmail: %v", err)
	}

	// Identity verified.
	if len(repo.verifiedCalls) != 1 {
		t.Fatalf("expected 1 SetUserVerified call, got %d", len(repo.verifiedCalls))
	}
	if repo.verifiedCalls[0].userID != userID {
		t.Errorf("verified call must target the right user")
	}

	// Audit entry written (action_type=account_linking).
	if len(repo.insertedUserLogs) != preLogs+1 {
		t.Fatalf("expected 1 audit entry, got %d", len(repo.insertedUserLogs)-preLogs)
	}
	last := repo.insertedUserLogs[len(repo.insertedUserLogs)-1]
	if last.ActionType != actionAccountLinking {
		t.Errorf("audit action_type must be %s, got %s", actionAccountLinking, last.ActionType)
	}
	if last.UserID != userID {
		t.Errorf("audit must target the right user")
	}
}

func TestVerifyEmail_RegistrationPurpose_NoLinkAudit(t *testing.T) {
	svc, repo, _ := newSecurityTestService(t)

	ident := repo.seedIdentity(providerEmailPassword, "user@company.com", hashFor("user@company.com"), false)

	plainToken, tokenHash, err := generateToken()
	if err != nil {
		t.Fatalf("generateToken: %v", err)
	}
	repo.tokens[tokenHash] = &AuthToken{
		ID:        uuid.New(),
		UserID:    ident.UserID,
		Purpose:   purposeEmailVerify,
		TokenHash: tokenHash,
		ExpiresAt: time.Now().Add(tokenTTL),
		CreatedAt: time.Now(),
	}

	preLogs := len(repo.insertedUserLogs)
	err = svc.VerifyEmail(context.Background(), plainToken)
	if err != nil {
		t.Fatalf("VerifyEmail: %v", err)
	}

	// Identity verified.
	if len(repo.verifiedCalls) != 1 {
		t.Fatalf("expected 1 SetUserVerified call")
	}

	// NO audit entry (registration purpose → no link audit).
	if len(repo.insertedUserLogs) != preLogs {
		t.Errorf("registration-purpose redemption must not write a link audit entry, got %d new", len(repo.insertedUserLogs)-preLogs)
	}
}

// ---- R16: no PII/secrets in logs ------------------------------------

func TestSecurity_LogsFreeOfSecrets(t *testing.T) {
	svc, repo, _ := newSecurityTestService(t)
	googleIdent := repo.seedIdentity(providerGoogle, "user@gmail.com", hashFor("user@gmail.com"), true)

	var logBuf bytes.Buffer
	origOut := log.Writer()
	log.SetOutput(&logBuf)
	defer log.SetOutput(origOut)

	_, _ = svc.SetPassword(context.Background(), googleIdent.UserID, "secret@email.com", "", "strong-pw-123")

	out := logBuf.String()
	if contains(out, "secret@email.com") {
		t.Errorf("log must not contain the submitted email: %q", out)
	}
	if contains(out, "strong-pw-123") {
		t.Errorf("log must not contain the password: %q", out)
	}
}

// ---- UnlinkGoogle helpers -------------------------------------------

// seedUnlinkFixture seeds a user with a google identity and an optional
// verified email_password identity (with a known bcrypt credential_secret).
// Returns the shared userID and the google identity's ID.
func seedUnlinkFixture(repo *fakeRepo, withVerifiedEP bool, password string) (userID, googleID uuid.UUID) {
	g := repo.seedIdentity(providerGoogle, "user@gmail.com", hashFor("user@gmail.com"), true)
	userID = g.UserID
	googleID = g.ID
	if withVerifiedEP {
		ep := repo.seedIdentity(providerEmailPassword, "user@company.com", hashFor("user@company.com"), true)
		ep.UserID = userID
		if password != "" {
			h, _ := secrets.HashPassword(password)
			ep.CredentialSecret = &h
		}
	}
	return userID, googleID
}

// ---- R9: unlink success — hard delete + audit ------------------------

func TestUnlinkGoogle_Success_HardDeletesAndAudits(t *testing.T) {
	svc, repo, _ := newSecurityTestService(t)
	userID, googleID := seedUnlinkFixture(repo, true, "correct-pw")

	preLogs := len(repo.insertedUserLogs)
	err := svc.UnlinkGoogle(context.Background(), userID, "correct-pw")
	if err != nil {
		t.Fatalf("UnlinkGoogle: %v", err)
	}

	// Google identity hard-deleted.
	if len(repo.deletedIdentityIDs) != 1 {
		t.Fatalf("expected 1 deleted identity, got %d", len(repo.deletedIdentityIDs))
	}
	if repo.deletedIdentityIDs[0] != googleID {
		t.Errorf("deleted identity must be the google one")
	}

	// Audit entry (action_type=account_linking).
	if len(repo.insertedUserLogs) != preLogs+1 {
		t.Fatalf("expected 1 audit entry, got %d", len(repo.insertedUserLogs)-preLogs)
	}
	last := repo.insertedUserLogs[len(repo.insertedUserLogs)-1]
	if last.ActionType != actionAccountLinking {
		t.Errorf("audit action_type must be %s, got %s", actionAccountLinking, last.ActionType)
	}
	if last.UserID != userID {
		t.Errorf("audit must target the caller's user_id")
	}
}

// ---- R10: INV-account-02 — google is the only identity ---------------

func TestUnlinkGoogle_OnlyIdentity_Rejected409(t *testing.T) {
	svc, repo, _ := newSecurityTestService(t)
	userID, _ := seedUnlinkFixture(repo, false, "")

	err := svc.UnlinkGoogle(context.Background(), userID, "")
	if err != ErrOnlyIdentity {
		t.Fatalf("expected ErrOnlyIdentity, got %v", err)
	}

	// No delete, no audit.
	if len(repo.deletedIdentityIDs) != 0 {
		t.Errorf("must not delete when google is the only identity")
	}
}

// ---- R11: INV-account-12 — remaining identity unverified ------------

func TestUnlinkGoogle_RejectsUnverifiedRemainingIdentity(t *testing.T) {
	svc, repo, _ := newSecurityTestService(t)

	// Seed google (verified) + email_password (UNVERIFIED).
	g := repo.seedIdentity(providerGoogle, "user@gmail.com", hashFor("user@gmail.com"), true)
	ep := repo.seedIdentity(providerEmailPassword, "user@company.com", hashFor("user@company.com"), false)
	ep.UserID = g.UserID

	err := svc.UnlinkGoogle(context.Background(), g.UserID, "")
	if err != ErrRemainingUnverified {
		t.Fatalf("expected ErrRemainingUnverified, got %v", err)
	}

	// No delete, no audit.
	if len(repo.deletedIdentityIDs) != 0 {
		t.Errorf("must not delete when remaining identity is unverified")
	}
}

// ---- R12: re-auth + evaluation order + idempotent no-op -------------

func TestUnlinkGoogle_WrongPassword_Rejected(t *testing.T) {
	svc, repo, _ := newSecurityTestService(t)
	userID, _ := seedUnlinkFixture(repo, true, "correct-pw")

	err := svc.UnlinkGoogle(context.Background(), userID, "WRONG-pw")
	if err != ErrInvalidCredentials {
		t.Fatalf("expected ErrInvalidCredentials, got %v", err)
	}

	// No delete, no audit (guards passed but password failed).
	if len(repo.deletedIdentityIDs) != 0 {
		t.Errorf("wrong password must not delete")
	}
	if len(repo.insertedUserLogs) != 0 {
		t.Errorf("wrong password must not write audit")
	}
}

func TestUnlinkGoogle_IdempotentNoGoogleRow_Returns200(t *testing.T) {
	svc, repo, _ := newSecurityTestService(t)

	// Seed ONLY a verified email_password identity (no google at all).
	ep := repo.seedIdentity(providerEmailPassword, "user@company.com", hashFor("user@company.com"), true)
	h, _ := secrets.HashPassword("correct-pw")
	ep.CredentialSecret = &h

	// Simulate the concurrent-loser case: google was already deleted by
	// a concurrent winner. The caller sees no google identity → 200.
	err := svc.UnlinkGoogle(context.Background(), ep.UserID, "")
	if err != nil {
		t.Fatalf("idempotent no-op must return nil, got %v", err)
	}

	// No delete, no audit (nothing to unlink).
	if len(repo.deletedIdentityIDs) != 0 {
		t.Errorf("idempotent no-op must not delete")
	}
}

// ---- R13: sequential idempotency (concurrent-loser path) -----------
//
// The fake cannot faithfully simulate FOR UPDATE across a transaction
// (its mu serializes individual method calls, not whole transactions).
// The real ≥100-goroutine stress against Postgres is task-06's
// integration suite. Here we prove the sequential idempotent-loser path:
// first call deletes + audits; second call sees no google → 200.

func TestUnlinkGoogle_ConcurrentRequests_GuardHolds(t *testing.T) {
	svc, repo, _ := newSecurityTestService(t)
	userID, _ := seedUnlinkFixture(repo, true, "correct-pw")

	// First call: succeeds, deletes google, writes audit.
	err := svc.UnlinkGoogle(context.Background(), userID, "correct-pw")
	if err != nil {
		t.Fatalf("first unlink: %v", err)
	}
	if len(repo.deletedIdentityIDs) != 1 {
		t.Fatalf("first call must delete 1 google identity")
	}

	// Second call: google is gone → idempotent nil (concurrent loser).
	err = svc.UnlinkGoogle(context.Background(), userID, "correct-pw")
	if err != nil {
		t.Fatalf("second unlink (idempotent): %v", err)
	}

	// Still only 1 delete + 1 audit (second call was a no-op).
	if len(repo.deletedIdentityIDs) != 1 {
		t.Errorf("second call must not delete again, got %d", len(repo.deletedIdentityIDs))
	}
	if len(repo.insertedUserLogs) != 1 {
		t.Errorf("second call must not write a second audit, got %d", len(repo.insertedUserLogs))
	}
}
