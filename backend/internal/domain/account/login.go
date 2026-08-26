package account

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/anhsbolic/kencleng/backend/internal/platform/auth"
	platformcrypto "github.com/anhsbolic/kencleng/backend/internal/platform/crypto"
	"github.com/anhsbolic/kencleng/backend/internal/platform/secrets"
)

// Sentinel errors for the login/session flows — mapped to HTTP status +
// Problem Details by transport/http (task-05). The service owns its error
// vocabulary; the transport layer only translates.
var (
	// ErrInvalidCredentials covers wrong email, wrong password, wrong
	// TOTP/backup code, and every refresh-token rejection class (missing /
	// expired / revoked / reuse-detected). One error on purpose: the wire
	// response must be byte-identical across all of them (anti-enumeration,
	// R3/R4; refresh classes are indistinguishable per contract).
	ErrInvalidCredentials = errors.New("invalid credentials")
	// ErrLockedOut means the persistent lockout threshold (≥5 failed
	// attempts in a trailing 15-minute window) tripped BEFORE credential
	// verification. Transport maps to 429 with the same generic body as the
	// 401 above — only the status code differs (R4).
	ErrLockedOut = errors.New("locked out")
	// ErrMfaPendingInvalid means the mfa_pending_token was expired,
	// malformed, or signed by the wrong key. No attempt row is recorded —
	// identity is not reliably known (R10). Maps to 401.
	ErrMfaPendingInvalid = errors.New("mfa pending token invalid")
)

const (
	// Lockout stages (login_attempts.stage; see migration 000006).
	stagePassword = "password"
	stageMFA      = "mfa"

	// Persistent lockout parameters (Fitur 2C): ≥5 failed attempts within a
	// trailing 15-minute window trip the lockout for that stage's key.
	lockoutWindow     = 15 * time.Minute
	maxFailedAttempts = 5
)

// dummyBcryptHash is burned on login branches where no real credential
// exists (unknown identifier), so wall-clock time is comparable to the
// wrong-password branch — extending register's R7 anti-enumeration timing
// discipline to login (R18). Generated lazily once per process; bcrypt of a
// constant cannot realistically fail, hence the panic on the impossible path.
var dummyBcryptHash = sync.OnceValue(func() string {
	h, err := secrets.HashPassword("kencleng-timing-parity-dummy-credential")
	if err != nil {
		panic("account: generate dummy bcrypt hash: " + err.Error())
	}
	return h
})

// LoginResult carries the outcome of Login/LoginMfa. Exactly one shape is
// populated per Status:
//
//   - "ok":           AccessToken/RefreshTokenPlain/AccessTokenExpiresAt/User set;
//     RefreshTokenPlain exists only to reach the cookie/body
//     write path and must never be logged or stored.
//   - "mfa_required": MFAPendingToken set; NO session tokens issued yet
//     (strict ordering — tokens come only after MFA
//     verification succeeds, R2).
type LoginResult struct {
	Status               string
	AccessToken          string
	RefreshTokenPlain    string
	AccessTokenExpiresAt time.Time
	User                 *LoginUserView
	MFAPendingToken      string
}

// RefreshResult carries the outcome of Refresh. The rotated plain token
// travels only to the Set-Cookie path (rotate-on-use).
type RefreshResult struct {
	AccessToken          string
	RefreshTokenPlain    string
	AccessTokenExpiresAt time.Time
}

// Login verifies email_password credentials and either completes the session
// outright (no MFA enrolled → R1) or returns an mfa_pending_token for the
// MFA step (R2). Ordering is contractual:
//
//	lockout check → identity lookup → bcrypt compare (always runs, R18)
//	→ attempt row → MFA branch → token issuance ONLY after both factors.
//
// Unverified identities log in fine (R5): verification is enforced by other
// domains at the point of the restricted action, not here.
func (s *Service) Login(ctx context.Context, email, password string) (LoginResult, error) {
	identifierHash := platformcrypto.HMAC([]byte(email), s.keys)

	// 1. Password-stage lockout keyed by identifier_hash — identity is not
	// reliably known yet, so this is the only key available. Rejected
	// attempts write NO attempt row: the credential was never checked (R4).
	since := s.now().Add(-lockoutWindow)
	failed, err := s.repo.CountRecentFailedAttemptsByIdentifier(ctx, identifierHash, stagePassword, since)
	if err != nil {
		return LoginResult{}, fmt.Errorf("account: count failed logins: %w", err)
	}
	if failed >= maxFailedAttempts {
		log.Printf("account: login locked out (stage=password)")
		return LoginResult{}, ErrLockedOut
	}

	// 2. Identity lookup by HMAC hash.
	identity, err := s.repo.FindAuthIdentityByIdentifierHash(ctx, providerEmailPassword, identifierHash)
	if err != nil {
		return LoginResult{}, fmt.Errorf("account: lookup email_password identity for login: %w", err)
	}

	// 3. Bcrypt-shaped work runs on EVERY branch (R18): unknown identifier
	// burns a compare against a fixed dummy hash so wrong-email and
	// wrong-password are wall-clock indistinguishable.
	matched := false
	if identity != nil && identity.CredentialSecret != nil {
		matched = s.compare(*identity.CredentialSecret, password) == nil
	} else {
		_ = s.compare(dummyBcryptHash(), password)
	}

	attemptUserID := (*uuid.UUID)(nil)
	if identity != nil {
		id := identity.UserID
		attemptUserID = &id
	}
	if err := s.insertAttempt(ctx, stagePassword, matched, identifierHash, attemptUserID); err != nil {
		return LoginResult{}, err
	}
	if !matched {
		return LoginResult{}, ErrInvalidCredentials
	}

	// 4. MFA branch: enabled_at IS NOT NULL decides (INV-account-07 makes
	// half-enabled state impossible, so this predicate is trustworthy).
	view, err := s.repo.GetLoginUserView(ctx, identity.UserID)
	if err != nil {
		return LoginResult{}, fmt.Errorf("account: load login user view: %w", err)
	}
	if view == nil {
		// Identity without a user row is an integrity anomaly, not a
		// credential failure — surface as internal error, not 401.
		return LoginResult{}, fmt.Errorf("account: user row missing for identity user_id=%s", identity.UserID)
	}
	if view.MFAEnabled {
		pending, err := s.mintPending(identity.UserID)
		if err != nil {
			return LoginResult{}, err
		}
		log.Printf("account: login requires mfa user_id=%s", identity.UserID)
		return LoginResult{Status: "mfa_required", MFAPendingToken: pending}, nil
	}

	result, err := s.issueSessionTokens(ctx, view)
	if err != nil {
		return LoginResult{}, err
	}
	log.Printf("account: login completed user_id=%s", identity.UserID)
	return result, nil
}

// LoginMfa completes a login whose password step already succeeded,
// proving it with mfa_pending_token + a second factor. Ordering (R7–R11):
//
//	pending-token verify (fail ⇒ 401, NO attempt row — identity unknown)
//	→ MFA-stage lockout keyed by user_id (fail ⇒ 429, NO attempt row)
//	→ factor verification → attempt row → token issuance exactly like
//	the no-MFA branch of Login.
func (s *Service) LoginMfa(ctx context.Context, pendingToken, totpCode, backupCode string) (LoginResult, error) {
	// Defensive boundary: exactly one second factor must be presented.
	// The handler validates too; duplicated here so the service contract
	// holds regardless of caller.
	if (totpCode == "") == (backupCode == "") {
		return LoginResult{}, ErrValidation
	}

	// 1. Pending-token verification — cryptographic gate before anything
	// touches the DB (R10).
	userID, err := s.verifyPendingFor(pendingToken)
	if err != nil {
		return LoginResult{}, ErrMfaPendingInvalid
	}

	// 2. MFA-stage lockout keyed by user_id (identity now established),
	// strictly BEFORE any code verification (R7); rejected attempts write
	// no attempt row.
	since := s.now().Add(-lockoutWindow)
	failed, err := s.repo.CountRecentFailedAttemptsByUser(ctx, userID, stageMFA, since)
	if err != nil {
		return LoginResult{}, fmt.Errorf("account: count failed mfa attempts: %w", err)
	}
	if failed >= maxFailedAttempts {
		log.Printf("account: login locked out (stage=mfa)")
		return LoginResult{}, ErrLockedOut
	}

	// 3. Second-factor verification. Backup codes redeem through a tx
	// because a match marks used_at (single-use, INV-account-06).
	var verified bool
	if backupCode != "" {
		tx, txErr := s.tx.BeginTx(ctx)
		if txErr != nil {
			return LoginResult{}, fmt.Errorf("account: begin backup-code tx: %w", txErr)
		}
		verified, err = s.mfa.VerifyBackupCode(ctx, tx, userID, backupCode)
		if err != nil {
			_ = tx.Rollback(ctx)
			return LoginResult{}, fmt.Errorf("account: verify backup code: %w", err)
		}
		if verified {
			if commitErr := tx.Commit(ctx); commitErr != nil {
				return LoginResult{}, fmt.Errorf("account: commit backup-code redemption: %w", commitErr)
			}
		} else {
			_ = tx.Rollback(ctx) // nothing was marked
		}
	} else {
		verified, err = s.mfa.VerifyTOTP(ctx, userID, totpCode)
		if err != nil {
			return LoginResult{}, fmt.Errorf("account: verify totp: %w", err)
		}
	}

	// 4. Attempt row records the outcome either way (R11).
	if err := s.insertAttemptWithUser(ctx, stageMFA, verified, userID); err != nil {
		return LoginResult{}, err
	}
	if !verified {
		return LoginResult{}, ErrInvalidCredentials
	}

	// 5. Completion mirrors the no-MFA branch of Login exactly (R8).
	view, err := s.repo.GetLoginUserView(ctx, userID)
	if err != nil {
		return LoginResult{}, fmt.Errorf("account: load login user view: %w", err)
	}
	if view == nil {
		return LoginResult{}, fmt.Errorf("account: user row missing for mfa user_id=%s", userID)
	}
	result, err := s.issueSessionTokens(ctx, view)
	if err != nil {
		return LoginResult{}, err
	}
	log.Printf("account: mfa login completed user_id=%s", userID)
	return result, nil
}

// Refresh rotates a single-use refresh token into a new session pair
// (INV-account-03) and detects reuse (INV-account-04). All four rejection
// classes — not-found, expired, revoked, already-rotated — collapse into
// one indistinguishable ErrInvalidCredentials, and every case where the
// family is known ends with the WHOLE family revoked, including the
// concurrent-race loser (spec Assumption D: race-loser ≡ attacker).
//
// ⚠️ Tier 0 fenced sub-area: the reuse/race-loser branch below is part of
// the human-paired rewrite scope (techplan Resolved #13).
func (s *Service) Refresh(ctx context.Context, refreshTokenPlain string) (RefreshResult, error) {
	if refreshTokenPlain == "" {
		return RefreshResult{}, ErrInvalidCredentials
	}
	now := s.now()
	tokenHash := sha256Hex(refreshTokenPlain)

	current, found, err := s.repo.FindRefreshTokenByHash(ctx, tokenHash)
	if err != nil {
		return RefreshResult{}, fmt.Errorf("account: find refresh token: %w", err)
	}

	// Classification: anything not currently-live goes down the same path.
	live := found &&
		current.RevokedAt == nil &&
		current.ReplacedByID == nil &&
		current.ExpiresAt.After(now)
	if !live {
		if found {
			// Known lineage — compromise assumed (reuse), kill the family
			// INCLUDING already-rotated descendants (INV-account-04).
			if revokeErr := s.revokeFamily(ctx, current.FamilyID); revokeErr != nil {
				return RefreshResult{}, revokeErr
			}
			log.Printf("account: refresh reuse detected — family revoked user_id=%s", current.UserID)
		}
		return RefreshResult{}, ErrInvalidCredentials
	}

	// Prepare the child + new access token BEFORE opening the rotation tx:
	// minting after commit would leave a rotated-cookie client with a 500
	// (fail toward re-login); minting before just wastes ~50µs of signing
	// on race losers.
	newPlain, err := randomHex(32)
	if err != nil {
		return RefreshResult{}, fmt.Errorf("account: generate refresh token: %w", err)
	}
	accessToken, expiresAt, mintErr := s.mintAccessFor(current.UserID, now)
	if mintErr != nil {
		return RefreshResult{}, mintErr
	}

	child := &RefreshToken{
		ID:        uuid.New(),
		TokenHash: sha256Hex(newPlain),
		ExpiresAt: now.Add(refreshTokenTTL),
		CreatedAt: now,
	}

	tx, err := s.tx.BeginTx(ctx)
	if err != nil {
		return RefreshResult{}, fmt.Errorf("account: begin rotate tx: %w", err)
	}
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback(ctx)
		}
	}()

	rotated, err := s.repo.RotateRefreshToken(ctx, tx, tokenHash, child)
	if err != nil {
		return RefreshResult{}, fmt.Errorf("account: rotate refresh token: %w", err)
	}
	if !rotated {
		// Lost the guarded-update race: another concurrent request won this
		// token milliseconds ago. Spec Assumption D defines this as
		// identical to theft — revoke the whole family, force full re-login.
		//
		// The tx MUST be rolled back explicitly before proceeding: BeginTx
		// already checked a pooled connection out, and skipping the deferred
		// rollback here would leak that connection in transaction state —
		// under sustained concurrency this starves the pool into a permanent
		// BeginTx deadlock (found by the ≥100-goroutine stress harness).
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return RefreshResult{}, fmt.Errorf("account: rollback lost-race tx: %w", rbErr)
		}
		committed = true // rolled back; nothing further for the defer to do
		if revokeErr := s.revokeFamily(ctx, current.FamilyID); revokeErr != nil {
			return RefreshResult{}, revokeErr
		}
		log.Printf("account: concurrent refresh lost race — family revoked user_id=%s", current.UserID)
		return RefreshResult{}, ErrInvalidCredentials
	}

	if err := tx.Commit(ctx); err != nil {
		return RefreshResult{}, fmt.Errorf("account: commit rotate: %w", err)
	}
	committed = true

	log.Printf("account: refresh rotated user_id=%s", current.UserID)
	return RefreshResult{
		AccessToken:          accessToken,
		RefreshTokenPlain:    newPlain,
		AccessTokenExpiresAt: expiresAt,
	}, nil
}

// Logout revokes the presented refresh token if present and unrevoked.
// Idempotent by contract: absent cookie or already-revoked token still
// yields success — the transport always clears the cookie and answers 204
// (R16).
func (s *Service) Logout(ctx context.Context, refreshTokenPlain string) error {
	if refreshTokenPlain == "" {
		return nil
	}
	token, found, err := s.repo.FindRefreshTokenByHash(ctx, sha256Hex(refreshTokenPlain))
	if err != nil {
		return fmt.Errorf("account: find refresh token for logout: %w", err)
	}
	if !found || token.RevokedAt != nil {
		return nil // nothing to revoke — idempotent no-op, not an error
	}

	tx, err := s.tx.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("account: begin logout tx: %w", err)
	}
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback(ctx)
		}
	}()
	if err := s.repo.RevokeRefreshTokenByHash(ctx, tx, token.TokenHash); err != nil {
		return fmt.Errorf("account: revoke refresh token: %w", err)
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("account: commit logout: %w", err)
	}
	committed = true

	log.Printf("account: logout completed user_id=%s", token.UserID)
	return nil
}

// ---- internal helpers ------------------------------------------------------

// insertAttempt records a password-stage outcome (identifier-keyed; user_id
// attached only when identity was established).
func (s *Service) insertAttempt(ctx context.Context, stage string, success bool, identifierHash string, userID *uuid.UUID) error {
	return s.writeAttempt(ctx, &LoginAttempt{
		ID:             uuid.New(),
		IdentifierHash: identifierHash,
		UserID:         userID,
		Stage:          stage,
		Success:        success,
		AttemptedAt:    s.now(),
	})
}

// insertAttemptWithUser records an MFA-stage outcome for a known user. Per
// spec Assumption C the NOT NULL identifier_hash column is backfilled from
// the user's own email_password identity (looked up via user_id) purely for
// schema consistency — it is never consulted by the MFA-stage lockout query.
//
// Deviation from Assumption C's literal wording (flagged): a Google-only
// user with MFA enrolled has no email_password identity at all, so the
// backfill falls back to a deterministic synthetic hash derived from the
// user ID rather than failing the login. The column stays non-empty and
// user-attributable; the lockout query is unaffected either way.
func (s *Service) insertAttemptWithUser(ctx context.Context, stage string, success bool, userID uuid.UUID) error {
	identifierHash := sha256Hex("mfa-stage:" + userID.String()) // synthetic fallback
	if hash, found, err := s.repo.FindIdentifierHashByUserAndProvider(ctx, userID, providerEmailPassword); err != nil {
		return fmt.Errorf("account: backfill identifier_hash for mfa attempt: %w", err)
	} else if found {
		identifierHash = hash
	}
	return s.writeAttempt(ctx, &LoginAttempt{
		ID:             uuid.New(),
		IdentifierHash: identifierHash,
		UserID:         &userID,
		Stage:          stage,
		Success:        success,
		AttemptedAt:    s.now(),
	})
}

// writeAttempt persists one attempt row in its own short transaction. The
// attempt row is bookkeeping, not a state machine: if the write fails after
// the credential decision, the login result stands (fail-open), the failure
// is logged with stage + success + the wrapped error, and the gap is
// observable — acceptable for a non-financial audit table (a lost attempt
// row can only undercount toward the lockout threshold, never lock anyone
// out spuriously). The logged error comes from pgx/goqu and holds no user
// credentials or tokens; only non-sensitive stage/success metadata are
// included in the log line (R19 log hygiene).
func (s *Service) writeAttempt(ctx context.Context, attempt *LoginAttempt) error {
	tx, err := s.tx.BeginTx(ctx)
	if err != nil {
		log.Printf("account: begin attempt tx failed (fail-open): stage=%s success=%t: %v",
			attempt.Stage, attempt.Success, err)
		return nil
	}
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback(ctx)
		}
	}()
	if err := s.repo.InsertLoginAttempt(ctx, tx, attempt); err != nil {
		log.Printf("account: insert login attempt failed (fail-open): stage=%s success=%t: %v",
			attempt.Stage, attempt.Success, err)
		return nil
	}
	if err := tx.Commit(ctx); err != nil {
		log.Printf("account: commit attempt tx failed (fail-open): stage=%s success=%t: %v",
			attempt.Stage, attempt.Success, err)
		return nil
	}
	committed = true
	return nil
}

// issueSessionTokens mints the access token, creates a first-generation
// refresh row (fresh family_id), loads the fresh-enough user view, and
// packages the ok-shaped LoginResult. Shared by Login's no-MFA branch and
// LoginMfa completion (R8: "completes exactly like R1's issuance half").
func (s *Service) issueSessionTokens(ctx context.Context, view *LoginUserView) (LoginResult, error) {
	now := s.now()
	access, expiresAt, err := s.mintAccessFor(view.ID, now)
	if err != nil {
		return LoginResult{}, err
	}

	plain, err := randomHex(32)
	if err != nil {
		return LoginResult{}, fmt.Errorf("account: generate refresh token: %w", err)
	}
	rt := &RefreshToken{
		ID:        uuid.New(),
		UserID:    view.ID,
		FamilyID:  uuid.New(), // fresh lineage — INV-account-03 chains start here
		TokenHash: sha256Hex(plain),
		ExpiresAt: now.Add(refreshTokenTTL),
		CreatedAt: now,
	}
	tx, err := s.tx.BeginTx(ctx)
	if err != nil {
		return LoginResult{}, fmt.Errorf("account: begin session tx: %w", err)
	}
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback(ctx)
		}
	}()
	if err := s.repo.InsertRefreshToken(ctx, tx, rt); err != nil {
		return LoginResult{}, fmt.Errorf("account: insert refresh token: %w", err)
	}
	if err := tx.Commit(ctx); err != nil {
		return LoginResult{}, fmt.Errorf("account: commit session: %w", err)
	}
	committed = true

	return LoginResult{
		Status:               "ok",
		AccessToken:          access,
		RefreshTokenPlain:    plain,
		AccessTokenExpiresAt: expiresAt,
		User:                 view,
	}, nil
}

// mintAccessFor wraps the injected closure with uniform error wrapping and
// expiry computation (expiry derives from the SAME clock that stamped the
// token, so transport-visible ExpiresAt can never disagree with the JWT).
func (s *Service) mintAccessFor(userID uuid.UUID, now time.Time) (string, time.Time, error) {
	if s.mintAccess == nil {
		return "", time.Time{}, errors.New("account: access token minting not configured")
	}
	token, err := s.mintAccess(userID, now)
	if err != nil {
		return "", time.Time{}, fmt.Errorf("account: mint access token: %w", err)
	}
	return token, now.Add(auth.AccessTokenTTL), nil
}

// mintPending wraps the injected pending-token closure.
func (s *Service) mintPending(userID uuid.UUID) (string, error) {
	if s.mintMFAPending == nil {
		return "", errors.New("account: mfa_pending token minting not configured")
	}
	token, err := s.mintMFAPending(userID, s.now())
	if err != nil {
		return "", fmt.Errorf("account: mint mfa_pending token: %w", err)
	}
	return token, nil
}

// verifyPendingFor wraps the injected pending-verifier closure: validates
// signature/expiry/purpose under the dedicated secret and returns the
// subject userID. Every failure collapses to a generic error — callers map
// it to ErrMfaPendingInvalid.
func (s *Service) verifyPendingFor(token string) (uuid.UUID, error) {
	if s.verifyPending == nil {
		return uuid.Nil, errors.New("account: mfa_pending verification not configured")
	}
	return s.verifyPending(token, s.now())
}

// revokeFamily kills every token in the lineage inside its own transaction
// (best-effort idempotent: rows already revoked are untouched by the guard).
func (s *Service) revokeFamily(ctx context.Context, familyID uuid.UUID) error {
	tx, err := s.tx.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("account: begin family-revoke tx: %w", err)
	}
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback(ctx)
		}
	}()
	if err := s.repo.RevokeRefreshTokenFamily(ctx, tx, familyID); err != nil {
		return fmt.Errorf("account: revoke refresh family: %w", err)
	}
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("account: commit family-revoke: %w", err)
	}
	committed = true
	return nil
}
