package account

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"

	"github.com/anhsbolic/kencleng/backend/internal/platform/crypto"
)

// MFA lifecycle sentinels (task #06). Mapped to HTTP by transport/http.
var (
	// ErrMfaAlreadyEnabled is returned by enroll when MFA is already active
	// (enabled_at IS NOT NULL). Maps to 409 with problem type
	// "https://kencleng.dev/errors/mfa-already-enabled" (D8).
	ErrMfaAlreadyEnabled = errors.New("mfa already enabled")
	// ErrInvalidTOTPCode is returned by confirm when the submitted code does
	// not validate against the pending secret. Maps to 422 indistinguishably
	// from ErrMfaNotPending (R7 — self-targeting endpoint, no enumeration
	// signal needed).
	ErrInvalidTOTPCode = errors.New("invalid totp code")
	// ErrMfaNotPending is returned by confirm when there is no pending
	// secret (never enrolled, or already enabled — including the concurrent
	// race loser). Wire-identical to ErrInvalidTOTPCode (R7).
	ErrMfaNotPending = errors.New("mfa enrollment not pending")
)

// TOTP parameter constants. These defaults are what mainstream
// authenticator apps emit (RFC 6238 + Google-Authenticator compatibility);
// deviation buys nothing and breaks scan-and-go UX (D1).
const (
	// backupCodeCount is how many recovery codes confirm issues (feature-spec
	// "exactly 10"); load-bearing literal as a named constant.
	backupCodeCount = 10
	// otpauthIssuer is the issuer label embedded in the otpauth:// URI.
	otpauthIssuer = "Kencleng"
	// backupCodeAlphabet is the char set for randomly generated backup
	// codes (lowercase alphanumeric, ~4.7 bits/char — code is a named
	// constant per repo style, excluded from the random pool to keep codes
	// unambiguous to type).
	backupCodeAlphabet = "abcdefghjkmnpqrstuvwxyz23456789"
	// totpPeriod / totpSkew / totpDigits mirror the Generate/Validate defaults
	// but are pinned as named constants so the two call sites can't diverge.
	totpPeriod = uint(30)
	totpSkew   = uint(1)
)

// actionMfaEnabled / actionMfaDisabled are the user_logs.action_type literals
// for Fitur 9's "MFA enable/disable" scope (D9). user_logs.action_type is
// unconstrained TEXT (migration 000005) so no DDL change is needed; full
// vocabulary ownership/consolidation stays with task #08.
const (
	actionMfaEnabled  = "mfa_enabled"
	actionMfaDisabled = "mfa_disabled"
)

// MfaEnroll starts TOTP enrollment: generates a fresh secret, persists it
// encrypted at rest, and returns the otpauth:// URI for the client to render
// as a QR code. enabled_at stays NULL until confirm (INV-account-07). If MFA
// is already active (enabled_at IS NOT NULL) the upsert's conflict-arm guard
// blocks the overwrite of a live secret and ErrMfaAlreadyEnabled is returned
// (D5 — the 409-when-active guarantee is structural, not a pre-read).
//
// The account-name label is the user's plaintext primary email via the
// existing GetLoginUserView decrypt path (D11); the returned URI embeds the
// secret and is NEVER logged (R15).
func (s *Service) MfaEnroll(ctx context.Context, userID uuid.UUID) (string, error) {
	view, err := s.repo.GetLoginUserView(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("account: load user for mfa enroll: %w", err)
	}
	if view == nil {
		return "", fmt.Errorf("account: user row missing for mfa enroll user_id=%s", userID)
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      otpauthIssuer,
		AccountName: view.Email,
	})
	if err != nil {
		return "", fmt.Errorf("account: generate totp secret: %w", err)
	}
	secret := key.Secret() // base32 plaintext — encrypted at the boundary below

	secretCiphertext, err := crypto.Encrypt([]byte(secret), s.keys)
	if err != nil {
		return "", fmt.Errorf("account: encrypt mfa secret: %w", err)
	}

	inserted, err := s.repo.UpsertPendingMFASecret(ctx, userID, secretCiphertext)
	if err != nil {
		return "", fmt.Errorf("account: store mfa pending secret: %w", err)
	}
	if !inserted {
		// enabled_at was non-null when the upsert landed — MFA already active
		// (or enabled concurrently). Never overwrote the live secret.
		return "", ErrMfaAlreadyEnabled
	}

	log.Printf("account: mfa enrollment started user_id=%s", userID)
	return key.URL(), nil
}

// MfaEnrollConfirm completes enrollment: verifies the submitted TOTP code
// against the pending secret, then in ONE transaction flips enabled_at
// (guarded #NULL-at-most-once) and stores exactly backupCodeCount hashed
// backup codes, plus the mfa_enabled audit entry. The plaintext codes are
// returned once (shown to the user); the server retains only hashes. The
// pending secret is preserved on a wrong code so the user can retry without
// re-scanning (R6); no-pending and already-enabled collapse to the same
// error as a wrong code (R7).
func (s *Service) MfaEnrollConfirm(ctx context.Context, userID uuid.UUID, totpCode string) ([]string, error) {
	secret, enabledAt, found, err := s.repo.GetMFATOTPSecretForVerify(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("account: load mfa pending secret: %w", err)
	}
	if !found || enabledAt != nil {
		// no pending secret (never enrolled / already confirmed / race loser)
		return nil, ErrMfaNotPending
	}

	valid, err := totp.ValidateCustom(totpCode, secret, s.now(), totp.ValidateOpts{
		Period:    totpPeriod,
		Skew:      totpSkew,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
	if err != nil {
		return nil, fmt.Errorf("account: validate mfa totp: %w", err)
	}
	if !valid {
		return nil, ErrInvalidTOTPCode // pending secret untouched (R6)
	}

	plains, err := generateBackupCodes(backupCodeCount)
	if err != nil {
		return nil, err
	}
	codes := make([]MFABackupCode, 0, backupCodeCount)
	now := s.now()
	for _, p := range plains {
		codes = append(codes, MFABackupCode{
			ID:        uuid.New(),
			UserID:    userID,
			CodeHash:  sha256Hex(normalizeBackupCode(p)),
			CreatedAt: now,
		})
	}

	tx, err := s.tx.BeginTx(ctx)
	if err != nil {
		return nil, fmt.Errorf("account: begin mfa confirm tx: %w", err)
	}
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback(ctx)
		}
	}()

	// Guarded enable FIRST (D4-A): losers match zero rows and roll back with
	// no orphan codes under a foreign enable.
	ok, err := s.repo.EnableMFATOTPIfPending(ctx, tx, userID)
	if err != nil {
		return nil, fmt.Errorf("account: enable mfa: %w", err)
	}
	if !ok {
		return nil, ErrMfaNotPending // lost the concurrent confirm race (R8)
	}
	if err := s.repo.InsertMFABackupCodes(ctx, tx, codes); err != nil {
		return nil, fmt.Errorf("account: insert mfa backup codes: %w", err)
	}
	if err := s.repo.InsertUserLog(ctx, tx, &UserLog{
		ID:         uuid.New(),
		UserID:     userID,
		ActionType: actionMfaEnabled,
		CreatedAt:  now,
	}); err != nil {
		return nil, fmt.Errorf("account: write mfa_enabled audit: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, fmt.Errorf("account: commit mfa confirm: %w", err)
	}
	committed = true

	log.Printf("account: mfa enrollment confirmed user_id=%s", userID)
	return plains, nil // shown exactly once
}

// MfaDisable disables MFA by setting enabled_at back to NULL. Backup-code
// rows are deliberately left in place (INV-account-06 implicit invalidation —
// they become unusable via the enabled-check at verification time). The
// email_password branch requires the current password (re-auth guard vs a
// hijacked live session); a Google-only caller reaches this method only after
// the handler has consumed a valid reauth marker (D6), so the service trusts
// that the re-auth already happened. An idempotent repeat disable (0 rows
// affected by the guarded UPDATE) still returns success and writes no
// duplicate audit row (R11).
func (s *Service) MfaDisable(ctx context.Context, userID uuid.UUID, password string) error {
	identities, err := s.repo.FindAuthIdentitiesByUser(ctx, userID)
	if err != nil {
		return fmt.Errorf("account: load identities for mfa disable: %w", err)
	}

	// Server-side provider detection (R14) — never a client-supplied flag.
	var epSecret *string
	for i := range identities {
		if identities[i].ProviderType == providerEmailPassword {
			epSecret = identities[i].CredentialSecret
			break
		}
	}

	if epSecret == nil {
		// Google-only: re-auth was the consumed marker at the handler (D6);
		// nothing to check here. NOTE: this path must not be reachable without
		// a marker — the handler enforces it (R13); the service assumes its
		// precondition was met.
	} else {
		if password == "" {
			return ErrValidation // email_password caller missing password → 422
		}
		// Compare burns comparable CPU on the failure path (timing discipline).
		if err := s.compare(*epSecret, password); err != nil {
			return ErrInvalidCredentials // 401, no state change (R12)
		}
	}

	tx, err := s.tx.BeginTx(ctx)
	if err != nil {
		return fmt.Errorf("account: begin mfa disable tx: %w", err)
	}
	committed := false
	defer func() {
		if !committed {
			_ = tx.Rollback(ctx)
		}
	}()

	disabled, err := s.repo.SetMFADisabledIfEnabled(ctx, tx, userID)
	if err != nil {
		return fmt.Errorf("account: disable mfa: %w", err)
	}
	if disabled {
		if err := s.repo.InsertUserLog(ctx, tx, &UserLog{
			ID:         uuid.New(),
			UserID:     userID,
			ActionType: actionMfaDisabled,
			CreatedAt:  s.now(),
		}); err != nil {
			return fmt.Errorf("account: write mfa_disabled audit: %w", err)
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("account: commit mfa disable: %w", err)
	}
	committed = true

	log.Printf("account: mfa disabled user_id=%s", userID)
	return nil
}

// MfaDisableReauthRequired reports the re-authentication path MFA disable
// needs for a user: true = Google-only (must present a consumed reauth
// marker), false = has an email_password identity (must present a password).
// Server-side determination per R14 — the handler uses it to route the reauth
// gate (consume marker vs rely on submitted password) without the marker
// state ever crossing into the domain service (D6).
func (s *Service) MfaDisableReauthRequired(ctx context.Context, userID uuid.UUID) (bool, error) {
	identities, err := s.repo.FindAuthIdentitiesByUser(ctx, userID)
	if err != nil {
		return false, fmt.Errorf("account: load identities for mfa disable reauth: %w", err)
	}
	for i := range identities {
		if identities[i].ProviderType == providerEmailPassword {
			return false, nil
		}
	}
	return true, nil // no email_password identity ⇒ Google-only
}

// generateBackupCodes returns n random lower-case-alphanumeric backup codes
// from crypto/rand. ~30+ bits of entropy each; brute force is additionally
// capped by the MFA-stage lockout at login. Codes are shown once in
// plaintext; only their hashes are ever stored.
func generateBackupCodes(n int) ([]string, error) {
	out := make([]string, 0, n)
	buf := make([]byte, 1)
	for i := 0; i < n; i++ {
		code := make([]byte, 8)
		for j := range code {
			if _, err := rand.Read(buf); err != nil {
				return nil, fmt.Errorf("account: generate backup code: %w", err)
			}
			code[j] = backupCodeAlphabet[int(buf[0])%len(backupCodeAlphabet)]
		}
		out = append(out, string(code))
	}
	return out, nil
}

// normalizeBackupCode canonicalizes a backup-code input before hashing or
// comparing: lowercase and strip any non-alphanumeric characters (users may
// type codes with spaces/dashes from the recovery sheet). Every consumer
// (confirm generation already stores normalized; verification hashes
// normalized input) shares this single helper so generator and verifier can
// never disagree (D2).
func normalizeBackupCode(s string) string {
	s = strings.ToLower(s)
	var b strings.Builder
	b.Grow(len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
			b.WriteByte(c)
		}
	}
	return b.String()
}
