package account

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/anhsbolic/kencleng/backend/internal/platform/crypto"
	"github.com/anhsbolic/kencleng/backend/internal/platform/notification"
	"github.com/anhsbolic/kencleng/backend/internal/platform/secrets"
)

// purposeEmailVerifyLink is the auth_tokens.purpose value issued by
// SetPassword Branch 1 (adding an email_password identity to a
// Google-only account). It distinguishes the token from registration's
// email_verification so VerifyEmail can write the user_logs audit entry
// truthfully when the redeemed token came from the linking flow (techplan
// D7; migration 000010 widens the CHECK to admit it).
const purposeEmailVerifyLink = "email_verification_link"

// Sentinel errors for unlink guard outcomes (task #04 uses these; defined
// here so both security flows share one error vocabulary in one file).
var (
	// ErrOnlyIdentity indicates unlink would leave the user with zero
	// identities (INV-account-02). Maps to 409 with problem type
	// "https://kencleng.dev/errors/only-identity".
	ErrOnlyIdentity = errors.New("only identity")
	// ErrRemainingUnverified indicates unlink is blocked because the
	// remaining identity is not yet verified (INV-account-12, stricter
	// than INV-account-02). Maps to 409 with the distinct problem type
	// "https://kencleng.dev/errors/unverified-remaining-identity".
	ErrRemainingUnverified = errors.New("remaining identity not verified")
)

// SetPassword implements the two server-side branches of
// POST /account/security/set-password. Branch selection is by whether the
// caller's user_id currently has an email_password identity — never a
// client-supplied flag (R6). Policy validation runs BEFORE the branch
// lookup (R4) and bcrypt runs on every branch (R5 timing parity), mirroring
// Register's anti-enumeration discipline.
//
// Returns (false, nil) for all generic-202 Branch-1 outcomes (created /
// claimed / race-loser) — the handler writes an identical 202 body.
// Returns (true, nil) for Branch-2 success — the handler writes 200.
// Returns (false, ErrInvalidCredentials) on wrong current_password (401).
// Returns (false, ErrValidation) on password-policy failure (422).
func (s *Service) SetPassword(ctx context.Context, userID uuid.UUID, email, currentPassword, newPassword string) (bool, error) {
	// R4: validate password policy BEFORE any enumeration-sensitive
	// branch lookup so a validation failure cannot leak whether the
	// caller has an email_password identity.
	if err := s.validatePassword(ctx, newPassword); err != nil {
		return false, err
	}

	// R5 (CPU time): always run bcrypt. Branch 1 stores the result as
	// credential_secret; Branch 2 also stores it. Never skip this.
	passwordHash, err := secrets.HashPassword(newPassword)
	if err != nil {
		return false, fmt.Errorf("account: hash password: %w", err)
	}

	// R6: server-side branch selection. FindAuthIdentitiesByUser is a
	// single pool read returning the caller's identities; the result
	// set is tiny (1-2 rows) and Branch 2 reuses it for the credential
	// comparison, avoiding a second lookup.
	identities, err := s.repo.FindAuthIdentitiesByUser(ctx, userID)
	if err != nil {
		return false, fmt.Errorf("account: lookup identities for set-password: %w", err)
	}

	hasEmailPassword := false
	for _, id := range identities {
		if id.ProviderType == providerEmailPassword {
			hasEmailPassword = true
			break
		}
	}

	if !hasEmailPassword {
		return s.setPasswordBranch1(ctx, userID, email, passwordHash)
	}
	return s.setPasswordBranch2(ctx, userID, currentPassword, passwordHash, identities)
}

// setPasswordBranch1 adds a new unverified email_password identity to the
// caller's account (the Google-only case). Mirrors Register's
// anti-enumeration pattern: pre-check by HMAC lookup → if claimed, send a
// conflict nudge + dummyWrite for DB-time parity → else insert identity +
// token in one tx → send verification email after commit. The
// unique-violation fallback covers the race the pre-check misses.
// Returns (false, nil) for every outcome (handler writes generic 202).
func (s *Service) setPasswordBranch1(ctx context.Context, userID uuid.UUID, email, passwordHash string) (bool, error) {
	identifierHash := crypto.HMAC([]byte(email), s.keys)

	// Pre-check: is the email already claimed by any user's
	// email_password identity?
	existing, err := s.repo.FindAuthIdentityByIdentifierHash(ctx, providerEmailPassword, identifierHash)
	if err != nil {
		return false, fmt.Errorf("account: lookup email_password identity for set-password: %w", err)
	}
	if existing != nil {
		// Conflict: claimed by someone. No identity/token created.
		// dummyWrite for DB-time parity with the creation branch;
		// conflict nudge instead of verification email.
		if err := s.dummyWrite(ctx); err != nil {
			return false, fmt.Errorf("account: timing write: %w", err)
		}
		s.sendNudge(ctx, email, notification.NudgeSetPasswordConflict)
		return false, nil
	}

	// Create unverified identity + verification token in one tx.
	plainToken, tokenHash, err := generateToken()
	if err != nil {
		return false, fmt.Errorf("account: generate token: %w", err)
	}

	tx, err := s.tx.BeginTx(ctx)
	if err != nil {
		return false, fmt.Errorf("account: begin tx: %w", err)
	}
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback(ctx)
		}
	}()

	identity := &AuthIdentity{
		ID:               uuid.New(),
		UserID:           userID,
		ProviderType:     providerEmailPassword,
		Identifier:       email,
		CredentialSecret: &passwordHash,
	}
	if err := s.repo.InsertAuthIdentity(ctx, tx, identity); err != nil {
		// R3: concurrent duplicate — the unique index on
		// (provider_type, identifier_hash) fired. Map to clean no-op
		// (return false, nil → generic 202), same as the pre-check
		// conflict. The tx rolls back via the deferred Rollback.
		if isUniqueViolation(err) {
			s.sendNudge(ctx, email, notification.NudgeSetPasswordConflict)
			return false, nil
		}
		return false, fmt.Errorf("account: insert auth_identity: %w", err)
	}

	token := &AuthToken{
		ID:        uuid.New(),
		UserID:    userID,
		Purpose:   purposeEmailVerifyLink,
		TokenHash: tokenHash,
		ExpiresAt: time.Now().Add(tokenTTL),
		CreatedAt: time.Now(),
	}
	if err := s.repo.InsertAuthToken(ctx, tx, token); err != nil {
		return false, fmt.Errorf("account: insert auth_token: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return false, fmt.Errorf("account: commit set-password: %w", err)
	}
	committed = true

	// Send the verification email AFTER commit — never inside the tx.
	s.sendVerification(ctx, email, plainToken)
	log.Printf("account: set-password branch 1 identity created user_id=%s", userID)
	return false, nil
}

// setPasswordBranch2 changes the caller's existing email_password
// credential in place. Requires current_password confirmation (re-auth
// guard vs a stolen-but-still-valid access token's 15-min window). The
// credential update and user-wide session revocation (INV-account-05)
// run in ONE transaction — a failure anywhere rolls both back, so no
// half-applied state (rotated secret without revoked sessions, or vice
// versa) can ever commit. Returns (true, nil) on success (handler → 200).
func (s *Service) setPasswordBranch2(ctx context.Context, userID uuid.UUID, currentPassword, newPasswordHash string, identities []AuthIdentity) (bool, error) {
	// Locate the email_password identity for credential comparison.
	var epIdentity *AuthIdentity
	for i := range identities {
		if identities[i].ProviderType == providerEmailPassword {
			epIdentity = &identities[i]
			break
		}
	}
	if epIdentity == nil || epIdentity.CredentialSecret == nil {
		// Branch selection guarantees existence; reaching here means
		// the identity has no credential_secret (data corruption or
		// a race with an incomplete insert). Fail closed.
		return false, fmt.Errorf("account: email_password identity missing credential_secret for user_id=%s", userID)
	}

	// Re-auth: compare current_password against the stored bcrypt hash.
	// The s.compare seam (secrets.ComparePassword by default) burns
	// comparable CPU even on the failure path.
	if err := s.compare(*epIdentity.CredentialSecret, currentPassword); err != nil {
		return false, ErrInvalidCredentials
	}

	// Atomic: update credential + revoke all sessions in one tx.
	tx, err := s.tx.BeginTx(ctx)
	if err != nil {
		return false, fmt.Errorf("account: begin tx: %w", err)
	}
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback(ctx)
		}
	}()

	if err := s.repo.UpdateIdentityCredentialSecret(ctx, tx, userID, providerEmailPassword, newPasswordHash); err != nil {
		return false, fmt.Errorf("account: update credential_secret: %w", err)
	}
	if err := s.repo.RevokeAllRefreshTokensForUser(ctx, tx, userID); err != nil {
		return false, fmt.Errorf("account: revoke all refresh tokens: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return false, fmt.Errorf("account: commit set-password branch 2: %w", err)
	}
	committed = true

	log.Printf("account: set-password branch 2 password changed user_id=%s", userID)
	return true, nil
}

// UnlinkGoogle removes ALL of the caller's google identities, guarded by
// INV-account-02 (≥1 identity remains) and INV-account-12 (≥1 VERIFIED
// non-google identity remains), with password re-authentication as the
// last gate before mutation. The whole check-then-delete sequence runs
// inside one transaction with FOR UPDATE row locks so concurrent unlinks
// serialize: the loser blocks, then classifies post-commit state and maps
// to idempotent 200 (google row already gone).
//
// Evaluation order (techplan R12 / evaluation-order micro-decision):
// guards FIRST (409s / idempotent no-op reachable by a passwordless
// Google-only caller), password LAST (only after guards pass).
func (s *Service) UnlinkGoogle(ctx context.Context, userID uuid.UUID, password string) error {
	tx, err := s.tx.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("account: begin tx: %w", err)
	}
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback(ctx)
		}
	}()

	// Acquire FOR UPDATE row locks on the caller's identities — this is
	// the serialization point for concurrent unlinks. Under READ COMMITTED
	// a concurrent winner's DELETE is visible to the loser at lock
	// acquisition, so the loser classifies post-commit state.
	rows, err := s.repo.FindAuthIdentitiesByUserForUpdate(ctx, tx, userID)
	if err != nil {
		return fmt.Errorf("account: lookup identities for unlink: %w", err)
	}

	// Partition into google + others.
	var googleIDs []uuid.UUID
	var verifiedEmailPassword *AuthIdentity
	for i := range rows {
		switch rows[i].ProviderType {
		case providerGoogle:
			googleIDs = append(googleIDs, rows[i].ID)
		case providerEmailPassword:
			if rows[i].VerifiedAt != nil && verifiedEmailPassword == nil {
				verifiedEmailPassword = &rows[i]
			}
		}
	}

	// No google identity → idempotent success (concurrent loser or
	// double-submit). Own-data visibility, zero leak.
	if len(googleIDs) == 0 {
		return nil
	}

	// INV-account-02: at least one OTHER identity must remain.
	hasOther := false
	for _, r := range rows {
		if r.ProviderType != providerGoogle {
			hasOther = true
			break
		}
	}
	if !hasOther {
		return ErrOnlyIdentity
	}

	// INV-account-12: at least one OTHER identity must be VERIFIED.
	// verifiedEmailPassword is non-nil iff a verified email_password
	// identity exists. (Google identities are always verified at creation,
	// but removing Google requires a verified NON-google identity — the
	// user must have a verified channel to recover the account through.)
	if verifiedEmailPassword == nil {
		return ErrRemainingUnverified
	}

	// Re-auth: compare password against the verified email_password
	// identity's stored bcrypt hash. This is the last gate before
	// mutation — a Google-only caller (no password) was already rejected
	// by the guards above. The s.compare seam burns comparable CPU even
	// on the failure path.
	if verifiedEmailPassword.CredentialSecret == nil {
		return fmt.Errorf("account: verified email_password identity missing credential_secret for user_id=%s", userID)
	}
	if err := s.compare(*verifiedEmailPassword.CredentialSecret, password); err != nil {
		return ErrInvalidCredentials
	}

	// Hard-delete ALL google identities (D3 — the endpoint means
	// "remove Google as a login method"; multi-google users are
	// reachable via intent=link).
	if err := s.repo.DeleteAuthIdentitiesByIDs(ctx, tx, googleIDs); err != nil {
		return fmt.Errorf("account: delete google identities: %w", err)
	}

	// Audit entry — commits atomically with the delete in the same tx.
	entry := &UserLog{
		ID:         uuid.New(),
		UserID:     userID,
		ActionType: actionAccountLinking,
		CreatedAt:  time.Now(),
	}
	if err := s.repo.InsertUserLog(ctx, tx, entry); err != nil {
		return fmt.Errorf("account: insert unlink audit log: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("account: commit unlink: %w", err)
	}
	committed = true

	log.Printf("account: google unlinked user_id=%s", userID)
	return nil
}
