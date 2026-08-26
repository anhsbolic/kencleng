package account

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// MfaVerifier abstracts TOTP / backup-code verification for the MFA step of
// login (/auth/login/mfa). The real implementation — TOTP computation over
// the encrypted secret and guarded single-use backup-code redemption — is
// owned by account task #6 (mfa_totp_secrets / mfa_backup_codes logic).
//
// Until task #6 lands, stubMfaVerifier FAILS CLOSED: no code ever verifies,
// so an attacker cannot reach token issuance by submitting codes, and the
// endpoint's lockout/attempt bookkeeping is still fully exercised.
//
// TODO(task-06-mfa-totp): replace the stub with the real verifier wired in
// cmd/server/main.go. This is a feature flag by construction — the seam is
// live code, not commented-out scaffolding.
type MfaVerifier interface {
	// VerifyTOTP reports whether code is a valid TOTP for userID's current
	// enrollment window(s). No DB writes.
	VerifyTOTP(ctx context.Context, userID uuid.UUID, code string) (bool, error)

	// VerifyBackupCode checks code against unused backup codes and, on a
	// match, marks that code's used_at via a guarded single-use UPDATE
	// inside tx (INV-account-06: NULL → timestamp at most once, only while
	// the user's mfa_totp_secrets.enabled_at IS NOT NULL — implicit
	// invalidation after MFA disable). Returns false without writing when
	// the code does not match or is no longer valid.
	VerifyBackupCode(ctx context.Context, tx pgx.Tx, userID uuid.UUID, code string) (bool, error)
}

// stubMfaVerifier is the pre-task-#6 placeholder: everything fails closed.
type stubMfaVerifier struct{}

// VerifyTOTP implements MfaVerifier — always fails closed until task #6.
func (stubMfaVerifier) VerifyTOTP(context.Context, uuid.UUID, string) (bool, error) {
	return false, nil
}

// VerifyBackupCode implements MfaVerifier — always fails closed until task #6.
func (stubMfaVerifier) VerifyBackupCode(context.Context, pgx.Tx, uuid.UUID, string) (bool, error) {
	return false, nil
}
