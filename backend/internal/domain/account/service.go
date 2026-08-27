package account

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/anhsbolic/kencleng/backend/internal/platform/auth"
	"github.com/anhsbolic/kencleng/backend/internal/platform/breachcheck"
	"github.com/anhsbolic/kencleng/backend/internal/platform/crypto"
	"github.com/anhsbolic/kencleng/backend/internal/platform/googleoauth"
	"github.com/anhsbolic/kencleng/backend/internal/platform/notification"
	"github.com/anhsbolic/kencleng/backend/internal/platform/secrets"
)

// Sentinel errors — mapped to HTTP status by transport/http (Task 05).
// Defined here so the service owns its error vocabulary; the transport
// layer only translates.
var (
	// ErrValidation indicates a request failed validation (password
	// policy, breach-list hit). Maps to 422.
	ErrValidation = errors.New("validation error")
	// ErrTokenExpired indicates a verification token exists but has
	// expired. Maps to 410.
	ErrTokenExpired = errors.New("token expired")
	// ErrTokenNotFound indicates a verification token was not found,
	// already used, or revoked (the 3-clause guard rejected it for any
	// reason other than expiry). Maps to 404.
	ErrTokenNotFound = errors.New("token not found")
)

// tokenTTL is the validity window for an email_verification token.
const tokenTTL = 24 * time.Hour

// resetTokenTTL is the validity window for a password_reset token
// (Fitur 2B). Deliberately NOT derived from tokenTTL: a reset link is a
// higher-risk credential (direct account takeover if stolen) and gets a
// much shorter window per the feature spec.
const resetTokenTTL = time.Hour

// emailVerification is the auth_identities provider_type / auth_tokens
// purpose for the email/password verification flow.
const (
	providerEmailPassword = "email_password"
	providerGoogle        = "google"
	purposeEmailVerify    = "email_verification"
	purposePasswordReset  = "password_reset"
)

// TxRunner abstracts db.BeginTx so the service is unit-testable without a
// real pgxpool. The real implementation, poolRunner, wraps *pgxpool.Pool.
type TxRunner interface {
	BeginTx(ctx context.Context) (pgx.Tx, error)
}

// breachChecker is the subset of breachcheck.Client the service depends
// on. *breachcheck.Client satisfies it; tests inject a fake so no network
// is hit and the fail-open path is exercisable deterministically.
type breachChecker interface {
	IsBreached(ctx context.Context, password string) (bool, error)
}

// poolRunner adapts *pgxpool.Pool to TxRunner.
type poolRunner struct{ pool *pgxpool.Pool }

// BeginTx begins a transaction on the underlying pool.
func (p poolRunner) BeginTx(ctx context.Context) (pgx.Tx, error) {
	return p.pool.BeginTx(ctx, pgx.TxOptions{})
}

// Service implements the account domain's registration, verification,
// resend, Google OAuth flows, and (since the login/session slice) the
// credential login / MFA completion / refresh rotation / logout flows.
// It is safe for concurrent use: it holds no mutable state of its own, and
// every injected dependency is goroutine-safe — session correctness lives
// in Postgres behind guarded statements, not in process memory.
type Service struct {
	repo        Repository
	tx          TxRunner
	breachCheck breachChecker
	email       notification.Sender
	keys        *crypto.Keys

	googleOAuth googleOAuthClient
	authKeys    *auth.Keys
	frontendURL string

	// ---- login/session seams (task #03 slice) ----

	// mfa verifies TOTP/backup codes at /auth/login/mfa. nil →
	// stubMfaVerifier (fails closed until account task #6 lands).
	mfa MfaVerifier
	// now is the clock seam; lockout windows and token TTLs derive from it.
	now func() time.Time
	// compare is the password-comparison seam (R18 timing discipline:
	// tests inject a recorder to prove every login branch burns comparable
	// CPU work). nil → secrets.ComparePassword.
	compare func(hashedPassword, password string) error
	// hashPassword is the password-hashing seam, mirroring compare:
	// production always leaves it nil (real bcrypt via
	// secrets.HashPassword); concurrency-heavy integration tests inject a
	// cheap stub so the invariant under test is DB row contention, not
	// crypto cost. nil → secrets.HashPassword.
	hashPassword func(password string) (string, error)
	// mintAccess signs an ES256 access token with purpose="access"
	// (wraps platform/auth MintAccessToken). Required for any flow that
	// issues sessions; calls fail with a clear error if left nil.
	mintAccess func(userID uuid.UUID, now time.Time) (string, error)
	// mintMFAPending signs the short-lived mfa_pending step-up token
	// (wraps platform/auth MintMFAPending). Required for Login's MFA branch.
	mintMFAPending func(userID uuid.UUID, now time.Time) (string, error)
	// verifyPending validates an mfa_pending_token and returns its subject
	// (wraps platform/auth VerifyMFAPending). Required for LoginMfa; fails
	// clearly if nil.
	verifyPending func(token string, now time.Time) (uuid.UUID, error)
}

// NewService constructs an account Service. db is used only to begin
// transactions; reads/standalone writes go through repo.
//
// Parameters beyond the original six: 7–8 are the Google OAuth additions
// (techplan §10), and 9–14 are the login/session seams — mfa verifier
// (nil → fail-closed stub until task #6), clock, password comparer
// (both nil → sensible defaults), and three token closures built over
// platform/auth/token.go (Tier 0 paired output). Deviations from the
// techplan's enumerated parameter count are flagged in the ticket report.
func NewService(repo Repository, db *pgxpool.Pool, bc *breachcheck.Client, sender notification.Sender, keys *crypto.Keys, googleOauth *googleoauth.Client, authKeys *auth.Keys, frontendURL string, mfa MfaVerifier, nowFn func() time.Time, compareFn func(hashedPassword, password string) error, mintAccess func(userID uuid.UUID, now time.Time) (string, error), mintMFAPending func(userID uuid.UUID, now time.Time) (string, error), verifyPending func(token string, now time.Time) (uuid.UUID, error)) *Service {
	if mfa == nil {
		mfa = stubMfaVerifier{}
	}
	if nowFn == nil {
		nowFn = time.Now
	}
	if compareFn == nil {
		compareFn = secrets.ComparePassword
	}
	return &Service{
		repo:        repo,
		tx:          poolRunner{pool: db},
		breachCheck: bc,
		email:       sender,
		keys:        keys,

		googleOAuth: googleOauth,
		authKeys:    authKeys,
		frontendURL: strings.TrimRight(frontendURL, "/"),

		mfa:            mfa,
		now:            nowFn,
		compare:        compareFn,
		hashPassword:   secrets.HashPassword,
		mintAccess:     mintAccess,
		mintMFAPending: mintMFAPending,
		verifyPending:  verifyPending,
	}
}

// Register orchestrates R1-R7, R16-R19. The four internal branches
// (new / unverified-existing / verified-existing / Google-only-conflict)
// all return nil so the handler (Task 05) writes an identical 202
// generic response. Branch wall-clock time is shaped to be equivalent
// (R7): bcrypt runs on every branch, and every branch performs at least
// one DB round-trip.
//
// PII handling: email is never logged. Only the fact ("register
// attempt"/"register completed") and, after creation, user_id are
// logged.
func (s *Service) Register(ctx context.Context, name, email, password string) error {
	// R5/R18: validate password policy BEFORE any enumeration-sensitive
	// branch lookup so a validation failure cannot leak whether the
	// email is registered.
	if err := s.validatePassword(ctx, password); err != nil {
		return err
	}

	// R7 (CPU time): always run bcrypt. The new-user branch stores the
	// result as credential_secret; the other three branches discard it.
	// Never skip this on no-op branches, never replace with a sleep.
	passwordHash, err := secrets.HashPassword(password)
	if err != nil {
		return fmt.Errorf("account: hash password: %w", err)
	}

	// Compute the lookup hash for the email_password identity.
	identifierHash := crypto.HMAC([]byte(email), s.keys)

	// R1-R4 branch dispatch on the existing email_password identity.
	identity, err := s.repo.FindAuthIdentityByIdentifierHash(ctx, providerEmailPassword, identifierHash)
	if err != nil {
		return fmt.Errorf("account: lookup email_password identity: %w", err)
	}

	switch {
	case identity == nil:
		// No email_password identity. Check for a Google-only conflict
		// (R4, R17) before creating a new user.
		google, err := s.repo.FindAuthIdentityByIdentifierHash(ctx, providerGoogle, identifierHash)
		if err != nil {
			return fmt.Errorf("account: lookup google identity: %w", err)
		}
		if google != nil {
			// R4 / R17: Google-only conflict. No new user, no token.
			// DB-write-shaped no-op for timing equivalence with R1/R2
			// (R7 DB-time half): a revoke against a non-existent user_id
			// affects 0 rows but has the same UPDATE+commit cost shape
			// as R2's real revoke. Then the nudge after.
			if err := s.dummyWrite(ctx); err != nil {
				return fmt.Errorf("account: timing write: %w", err)
			}
			s.sendNudge(ctx, email, notification.NudgeGoogleOnly)
			return nil
		}
		// R1: new user. Insert user + identity + token in one tx.
		return s.registerNewUser(ctx, name, email, passwordHash)

	case identity.VerifiedAt == nil:
		// R2: unverified existing. Same internal action as
		// verify-email/resend (R13): revoke old tokens, issue a new one,
		// send the verification email with the new token. No new
		// user/identity. The new token must be delivered (the resend
		// email carries it) so the user can complete verification.
		plainToken, err := s.issueNewVerificationToken(ctx, identity.UserID)
		if err != nil {
			return err
		}
		s.sendVerification(ctx, email, plainToken)
		return nil

	default:
		// R3: verified existing. No new record. DB-write-shaped no-op for
		// timing equivalence with R1/R2 (R7 DB-time half): a revoke
		// against a non-existent user_id affects 0 rows but has the same
		// UPDATE+commit cost shape as R2's real revoke. Then the
		// password-reset nudge.
		if err := s.dummyWrite(ctx); err != nil {
			return fmt.Errorf("account: timing write: %w", err)
		}
		s.sendNudge(ctx, email, notification.NudgePasswordReset)
		return nil
	}
}

// dummyWrite performs a DB-write-shaped no-op for anti-enumeration timing
// equivalence (R7 DB-time half). It begins a tx, calls RevokeTokens with a
// synthetic (non-existent) user_id so the UPDATE affects 0 rows, and
// commits — same BeginTx + UPDATE + Commit round-trip shape as R2's real
// revoke, but touching no real rows. R3/R4 call this so their wall-clock
// DB time matches R1/R2; without it, the no-op branches are measurably
// faster against real Postgres, leaking "verified/google-only" (fast) vs
// "new/unverified" (slow) — an enumeration side-channel.
func (s *Service) dummyWrite(ctx context.Context) error {
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
	if err := s.repo.RevokeTokens(ctx, tx, uuid.New(), purposeEmailVerify); err != nil {
		return fmt.Errorf("account: dummy revoke: %w", err)
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("account: commit dummy: %w", err)
	}
	committed = true
	return nil
}

// registerNewUser inserts a new User + AuthIdentity + AuthToken in a
// single transaction (R16: on a concurrent duplicate, the
// auth_identities unique index fails and the whole tx rolls back
// cleanly — no orphaned users row). Sends the verification email after
// commit. R1.
func (s *Service) registerNewUser(ctx context.Context, name, email, passwordHash string) error {
	userID := uuid.New()
	identityID := uuid.New()
	plainToken, tokenHash, err := generateToken()
	if err != nil {
		return fmt.Errorf("account: generate token: %w", err)
	}

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

	user := &User{
		ID:           userID,
		Name:         name,
		PrimaryEmail: email,
	}
	if err := s.repo.InsertUser(ctx, tx, user); err != nil {
		return fmt.Errorf("account: insert user: %w", err)
	}

	identity := &AuthIdentity{
		ID:               identityID,
		UserID:           userID,
		ProviderType:     providerEmailPassword,
		Identifier:       email,
		CredentialSecret: &passwordHash,
	}
	if err := s.repo.InsertAuthIdentity(ctx, tx, identity); err != nil {
		// R16: concurrent duplicate registration. The unique index on
		// (provider_type, identifier_hash) failed — map to a clean
		// no-op (return nil → 202 generic), indistinguishable from a
		// normal R1 to the caller. The tx rolls back via the deferred
		// Rollback, leaving no orphaned users row.
		if isUniqueViolation(err) {
			log.Printf("account: concurrent duplicate registration suppressed (provider=%s)", providerEmailPassword)
			return nil
		}
		return fmt.Errorf("account: insert auth_identity: %w", err)
	}

	token := &AuthToken{
		ID:        uuid.New(),
		UserID:    userID,
		Purpose:   purposeEmailVerify,
		TokenHash: tokenHash,
		ExpiresAt: time.Now().Add(tokenTTL),
		CreatedAt: time.Now(),
	}
	if err := s.repo.InsertAuthToken(ctx, tx, token); err != nil {
		return fmt.Errorf("account: insert auth_token: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("account: commit register: %w", err)
	}
	committed = true

	// Send the verification email AFTER commit — never inside the tx.
	// If the send fails, the DB state is correct and the user can resend.
	s.sendVerification(ctx, email, plainToken)
	log.Printf("account: registration completed user_id=%s", userID)
	return nil
}

// issueNewVerificationToken revokes the user's old email_verification
// tokens and issues a new one in a single transaction (R13), returning
// the plain token so the caller can send the verification email with it
// after commit. Shared by R2 (register unverified existing) and R13
// (resend endpoint) — both send the verification email carrying the
// new token.
func (s *Service) issueNewVerificationToken(ctx context.Context, userID uuid.UUID) (string, error) {
	plainToken, tokenHash, err := generateToken()
	if err != nil {
		return "", fmt.Errorf("account: generate token: %w", err)
	}

	tx, err := s.tx.BeginTx(ctx)
	if err != nil {
		return "", fmt.Errorf("account: begin tx: %w", err)
	}
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback(ctx)
		}
	}()

	if err := s.repo.RevokeTokens(ctx, tx, userID, purposeEmailVerify); err != nil {
		return "", fmt.Errorf("account: revoke tokens: %w", err)
	}
	token := &AuthToken{
		ID:        uuid.New(),
		UserID:    userID,
		Purpose:   purposeEmailVerify,
		TokenHash: tokenHash,
		ExpiresAt: time.Now().Add(tokenTTL),
		CreatedAt: time.Now(),
	}
	if err := s.repo.InsertAuthToken(ctx, tx, token); err != nil {
		return "", fmt.Errorf("account: insert auth_token: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return "", fmt.Errorf("account: commit resend: %w", err)
	}
	committed = true

	return plainToken, nil
}

// VerifyEmail orchestrates R8-R12. Redeem + set-verified run in a single
// transaction (S2 fix): a SetUserVerified failure rolls back the redeem,
// so the token is not burned without the identity being verified. The
// redeemed user_id comes back from RedeemToken via RETURNING — no
// re-fetch that could silently fail (S1 fix: no userIDForToken path).
// Single-use correctness is still the atomic 3-clause UPDATE ... WHERE
// inside RedeemToken — no application-level locking.
func (s *Service) VerifyEmail(ctx context.Context, token string) error {
	tokenHash := sha256Hex(token)

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

	userID, purpose, ok, err := s.repo.RedeemToken(ctx, tx, tokenHash)
	if err != nil {
		return fmt.Errorf("account: redeem token: %w", err)
	}
	if ok {
		// Cross-purpose guard (Q1 fix): RedeemToken's guard is
		// purpose-agnostic by design; the consuming flow owns the
		// purpose check. A password_reset token presented here must NOT
		// verify an email — returning an error rolls the redeem back via
		// the deferred Rollback, so the token is not consumed.
		// Both email_verification (registration) and
		// email_verification_link (set-password Branch 1) are valid
		// here; the link purpose additionally triggers an audit entry
		// after SetUserVerified (R14).
		if purpose != purposeEmailVerify && purpose != purposeEmailVerifyLink {
			return ErrTokenNotFound
		}
		// R8: valid token redeemed. Set the user's email_password
		// identity verified_at. The userID comes from RedeemToken's
		// RETURNING — no re-fetch. If this fails, the deferred
		// Rollback undoes the redeem (token not burned, user can retry).
		if err := s.repo.SetUserVerified(ctx, tx, userID, providerEmailPassword, time.Now()); err != nil {
			return fmt.Errorf("account: set verified: %w", err)
		}
		// R14: when the redeemed token came from the set-password
		// Branch-1 flow (purpose=email_verification_link), write the
		// audit entry in the SAME transaction — the identity is now
		// verified, which is the moment the spec says to audit ("on
		// successful verification, not at initial creation").
		// Registration-purpose redemptions write no such row.
		if purpose == purposeEmailVerifyLink {
			entry := &UserLog{
				ID:         uuid.New(),
				UserID:     userID,
				ActionType: actionAccountLinking,
				CreatedAt:  time.Now(),
			}
			if err := s.repo.InsertUserLog(ctx, tx, entry); err != nil {
				return fmt.Errorf("account: insert link audit log: %w", err)
			}
		}
		if err := tx.Commit(ctx); err != nil {
			return fmt.Errorf("account: commit verify: %w", err)
		}
		committed = true

		log.Printf("account: email verified user_id=%s (token redacted)", userID)
		return nil
	}

	// !ok: disambiguate expired (R9) vs not-found/used/revoked (R10/R11).
	// This is a read after the tx rolled back (nothing to undo on the
	// ok==false path), so it runs outside the redeem tx.
	t, err := s.repo.FindAuthTokenByHash(ctx, tokenHash)
	if err != nil {
		return fmt.Errorf("account: find token: %w", err)
	}
	if t != nil && !t.ExpiresAt.After(time.Now()) {
		return ErrTokenExpired
	}
	return ErrTokenNotFound
}

// ResendVerification orchestrates R13-R14. Rate limiting (R15) is
// transport middleware (Task 05). Returns nil for both the match and
// no-match branches so the handler writes an identical 202 generic.
func (s *Service) ResendVerification(ctx context.Context, email string) error {
	identifierHash := crypto.HMAC([]byte(email), s.keys)
	identity, err := s.repo.FindAuthIdentityByIdentifierHash(ctx, providerEmailPassword, identifierHash)
	if err != nil {
		return fmt.Errorf("account: lookup identity for resend: %w", err)
	}
	if identity != nil && identity.VerifiedAt == nil {
		// R13: unverified match — revoke + new token, then send the
		// verification email with the new token after commit.
		plainToken, err := s.issueNewVerificationToken(ctx, identity.UserID)
		if err != nil {
			return err
		}
		s.sendVerification(ctx, email, plainToken)
		return nil
	}
	// R14: no match / verified / google-only — no token, no email, no-op.
	return nil
}

// ForgotPassword orchestrates the Fitur 2B request half: three internal
// branches (registered email_password / Google-only / no match) that all
// return nil so the handler writes an identical 202 generic response.
//
// Branch wall-clock time is shaped for anti-enumeration equivalence
// (same reasoning as Register R7): the registered branch does lookup +
// BeginTx + INSERT + Commit; each no-op branch does its lookups plus a
// dummyWrite so its DB-time shape matches. Unlike the verification resend
// flow, a new reset token is issued WITHOUT revoking prior outstanding
// ones — resolved spec Assumption A keeps every token guarded
// independently by INV-account-08's single-use predicate.
//
// The plain token leaves the process exactly once, via the post-commit
// SendPasswordResetEmail call, and is never logged.
func (s *Service) ForgotPassword(ctx context.Context, email string) error {
	identifierHash := crypto.HMAC([]byte(email), s.keys)

	identity, err := s.repo.FindAuthIdentityByIdentifierHash(ctx, providerEmailPassword, identifierHash)
	if err != nil {
		return fmt.Errorf("account: lookup email_password identity for forgot: %w", err)
	}
	if identity != nil {
		plainToken, err := s.issueResetToken(ctx, identity.UserID)
		if err != nil {
			return err
		}
		s.sendPasswordReset(ctx, email, plainToken)
		log.Printf("account: password reset requested user_id=%s", identity.UserID)
		return nil
	}

	googleIdentity, err := s.repo.FindAuthIdentityByIdentifierHash(ctx, providerGoogle, identifierHash)
	if err != nil {
		return fmt.Errorf("account: lookup google identity for forgot: %w", err)
	}

	// Both no-op branches take the same DB-write-shaped path (dummyWrite)
	// so their timing matches the real branch above; only the email
	// content differs between them, and neither leaks into the API
	// response.
	if err := s.dummyWrite(ctx); err != nil {
		return fmt.Errorf("account: timing write: %w", err)
	}
	if googleIdentity != nil {
		s.sendNudge(ctx, email, notification.NudgeGoogleOnly)
	}
	return nil
}

// issueResetToken inserts a fresh password_reset AuthToken in a single
// transaction, returning the plain token so the caller can send it by
// email after commit.
//
// Deliberately NOT reusing issueNewVerificationToken: that helper revokes
// the user's prior tokens of the same purpose before inserting, which is
// correct for verification resend but would violate resolved Assumption A
// here — repeated forgot-password requests must not invalidate
// previously-issued, still-unexpired reset tokens.
func (s *Service) issueResetToken(ctx context.Context, userID uuid.UUID) (string, error) {
	plainToken, tokenHash, err := generateToken()
	if err != nil {
		return "", fmt.Errorf("account: generate token: %w", err)
	}

	tx, err := s.tx.BeginTx(ctx)
	if err != nil {
		return "", fmt.Errorf("account: begin tx: %w", err)
	}
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback(ctx)
		}
	}()

	token := &AuthToken{
		ID:        uuid.New(),
		UserID:    userID,
		Purpose:   purposePasswordReset,
		TokenHash: tokenHash,
		ExpiresAt: time.Now().Add(resetTokenTTL),
		CreatedAt: time.Now(),
	}
	if err := s.repo.InsertAuthToken(ctx, tx, token); err != nil {
		return "", fmt.Errorf("account: insert auth_token: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return "", fmt.Errorf("account: commit reset token: %w", err)
	}
	committed = true

	return plainToken, nil
}

// ResetPassword orchestrates the Fitur 2B redemption half. Order of
// operations is load-bearing (spec Assumption B):
//
//  1. validatePassword runs BEFORE any DB work — a fixable input mistake
//     must never burn the single-use token (422 with used_at still NULL).
//  2. The bcrypt hash runs before BeginTx — ~100ms of CPU must never hold
//     row locks open.
//  3. Redeem → purpose check → credential update → mass session revoke
//     all run inside ONE transaction: INV-account-05 requires the
//     credential update and RevokeAllRefreshTokensForUser to commit or
//     roll back together, and the same boundary makes any earlier failure
//     un-redeem the token automatically (deferred Rollback).
//
// On !ok the disambiguation read runs after the rolled-back tx (nothing
// to undo), mirroring VerifyEmail: expired → ErrTokenExpired (410);
// not-found / already-used / revoked → ErrTokenNotFound (404). Under a
// concurrent double-submit the loser observes used_at already set and
// maps to 404 — exactly the outcome INV-account-08's race test asserts.
func (s *Service) ResetPassword(ctx context.Context, token, newPassword string) error {
	if err := s.validatePassword(ctx, newPassword); err != nil {
		return err
	}

	// Nil-safe like the rest of the seam family: NewService wires
	// secrets.HashPassword, but struct-literal test constructions may
	// leave it unset.
	hashFn := s.hashPassword
	if hashFn == nil {
		hashFn = secrets.HashPassword
	}
	passwordHash, err := hashFn(newPassword)
	if err != nil {
		return fmt.Errorf("account: hash password: %w", err)
	}

	tokenHash := sha256Hex(token)

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

	userID, purpose, ok, err := s.repo.RedeemToken(ctx, tx, tokenHash)
	if err != nil {
		return fmt.Errorf("account: redeem token: %w", err)
	}
	if ok {
		// Cross-purpose guard (Q1 mirror): a token minted for email
		// verification must not reset a password. Returning here rolls
		// back the redeem — the token stays unconsumed.
		if purpose != purposePasswordReset {
			return ErrTokenNotFound
		}

		if err := s.repo.UpdateIdentityCredentialSecret(ctx, tx, userID, providerEmailPassword, passwordHash); err != nil {
			return fmt.Errorf("account: update credential_secret: %w", err)
		}
		if err := s.repo.RevokeAllRefreshTokensForUser(ctx, tx, userID); err != nil {
			return fmt.Errorf("account: revoke sessions: %w", err)
		}

		if err := tx.Commit(ctx); err != nil {
			return fmt.Errorf("account: commit reset: %w", err)
		}
		committed = true

		log.Printf("account: password reset completed user_id=%s (token redacted)", userID)
		return nil
	}

	t, err := s.repo.FindAuthTokenByHash(ctx, tokenHash)
	if err != nil {
		return fmt.Errorf("account: find token: %w", err)
	}
	if t != nil && !t.ExpiresAt.After(time.Now()) {
		return ErrTokenExpired
	}
	return ErrTokenNotFound
}

// validatePassword enforces the password policy: length >= 8 and not in
// the breach list (fail-open). R5/R6/R18/R19. Checked before any branch
// lookup so it cannot leak email state.
func (s *Service) validatePassword(ctx context.Context, password string) error {
	if len(password) < 8 {
		return ErrValidation
	}
	breached, err := s.breachCheck.IsBreached(ctx, password)
	if err != nil {
		// IsBreached fails open (returns false, nil) on API unreachable;
		// a non-nil error here is a real failure, not a fail-open.
		return fmt.Errorf("account: breach check: %w", err)
	}
	if breached {
		return ErrValidation
	}
	return nil
}

// sendVerification sends the verification email with the plain token.
// The plain token leaves the process exactly once: here, after commit.
// It is never logged.
func (s *Service) sendVerification(ctx context.Context, email, plainToken string) {
	if err := s.email.SendVerificationEmail(ctx, email, plainToken); err != nil {
		// Post-commit email failure is non-fatal: the DB state is
		// correct and the user can resend. Log a sanitized category,
		// not the raw error — a real SMTP error can embed the recipient
		// or token. Per go/secrets-and-sensitive-logging.md §1.
		log.Printf("account: send verification email failed (recipient redacted): %s",
			notificationErrorCategory(err))
	}
}

// sendPasswordReset sends the password-reset email with the plain token.
// The plain token leaves the process exactly once: here, after commit.
// It is never logged.
func (s *Service) sendPasswordReset(ctx context.Context, email, plainToken string) {
	if err := s.email.SendPasswordResetEmail(ctx, email, plainToken); err != nil {
		// Same sanitization as sendVerification: a real SMTP error can
		// embed the recipient or token. Per go/secrets-and-sensitive-logging.md §1.
		log.Printf("account: send password reset email failed (recipient redacted): %s",
			notificationErrorCategory(err))
	}
}

// sendNudge sends a nudge email. The recipient address is passed to the
// sender but never logged by it (FakeSender redacts).
func (s *Service) sendNudge(ctx context.Context, email, nudgeType string) {
	if err := s.email.SendNudgeEmail(ctx, email, nudgeType); err != nil {
		// Same sanitization as sendVerification: a real SMTP error can
		// embed the recipient. nudgeType is a package constant, safe.
		log.Printf("account: send nudge email failed type=%s (recipient redacted): %s",
			nudgeType, notificationErrorCategory(err))
	}
}

// notificationErrorCategory reduces a notification-sender error to a
// safe, PII-free category string for logging. It never returns
// err.Error() verbatim — a real SMTP error can embed the recipient email
// or token. Per go/secrets-and-sensitive-logging.md §1.
func notificationErrorCategory(err error) string {
	// A minimal timeout check for errors that implement the timeout
	// interface (net package convention). Anything else gets a coarse
	// category — never the raw message.
	var t interface{ Timeout() bool }
	if errors.As(err, &t) && t.Timeout() {
		return "timeout"
	}
	return "send failed"
}

// generateToken produces a user-facing plain token and its stored
// SHA-256 hex hash. 32 random bytes from crypto/rand, hex-encoded for
// the email link; the SHA-256 of that plain token is stored (never the
// plain token itself).
func generateToken() (plainToken, tokenHash string, err error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", "", fmt.Errorf("generate token: read rand: %w", err)
	}
	plain := hex.EncodeToString(b)
	return plain, sha256Hex(plain), nil
}

// sha256Hex returns the lowercase SHA-256 hex digest of s.
func sha256Hex(s string) string {
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}

// isUniqueViolation reports whether err is a PostgreSQL unique-constraint
// violation (SQLSTATE 23505). Used to detect concurrent duplicate
// registration (R16) without string-matching the error message.
func isUniqueViolation(err error) bool {
	var pgErr *pgconn.PgError
	return errors.As(err, &pgErr) && pgErr.Code == "23505"
}
