package account

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

// MfaVerifier abstracts TOTP / backup-code verification for the MFA step of
// login (/auth/login/mfa) and backs enrollment-time confirmation where the
// service validates the submitted code against the pending secret.
//
// The real implementation (totpVerifier) computes TOTP over the decrypted
// secret obtained through Repository, and does guarded single-use backup-code
// redemption inside the caller's transaction.
//
// stubMfaVerifier FAILS CLOSED: it is retained purely as NewService's
// nil-input safety net so tests that build a bare Service struct literal keep
// working. It is no longer reachable from production wiring (main.go passes
// NewMfaVerifier's result).
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

// stubMfaVerifier is the fail-closed nil-safety net — everything fails
// closed. NOT used by production wiring (which passes the real verifier);
// retained so struct-literal test services compile and keep their old
// fail-closed default.
type stubMfaVerifier struct{}

// VerifyTOTP implements MfaVerifier — always fails closed.
func (stubMfaVerifier) VerifyTOTP(context.Context, uuid.UUID, string) (bool, error) {
	return false, nil
}

// VerifyBackupCode implements MfaVerifier — always fails closed.
func (stubMfaVerifier) VerifyBackupCode(context.Context, pgx.Tx, uuid.UUID, string) (bool, error) {
	return false, nil
}

// ⚠️ TIER 0 FENCED SUB-AREA — DRAFT-FOR-PAIRING ⚠️
// The totpVerifier below is the TOTP generation/verification and guarded-
// redemption core of the MFA slice (account task #6). tasks.md's Tier-0-sub-
// area authorship KPI requires this be human-authored or human-paired before
// merge: the bodies here are the agent's full draft, kept as the review
// harness (all suites pass against them), and must go through a pairing
// session with Anhar (techplan D12) BEFORE the code-review stage. The tests
// are authored to survive a rewrite — they assert end-state invariants, not
// these implementations' internals.

// totpVerifier is the real MfaVerifier. TOTP validation reads the decrypted
// base32 secret through the Repository port; backup-code redemption runs the
// single guarded UPDATE inside the caller's transaction. Crypto stays at the
// adapter boundary per entity.go doctrine (D3): this type never sees
// ciphertext and holds no key material.
type totpVerifier struct {
	repo Repository
}

// VerifyTOTP reports whether code is a valid TOTP for userID's current
// enrollment, computed over the decrypted secret for the standard
// Google-Authenticator window (±1 × 30 s, RFC 6238). Pure read — no DB
// writes. A user with no mfa_totp_secrets row, or one that is not enabled
// (enabled_at IS NULL — mutually exclusive with enrollment being complete
// under INV-account-07), fails closed: (false, nil), never an error.
func (v *totpVerifier) VerifyTOTP(ctx context.Context, userID uuid.UUID, code string) (bool, error) {
	secret, enabledAt, found, err := v.repo.GetMFATOTPSecretForVerify(ctx, userID)
	if err != nil {
		return false, fmt.Errorf("account: load mfa secret for totp verify: %w", err)
	}
	if !found || enabledAt == nil {
		// No enrollment, or MFA not active — nothing to verify against.
		return false, nil
	}

	valid, err := totp.ValidateCustom(code, secret, time.Now(), totp.ValidateOpts{
		Period:    totpPeriod,
		Skew:      totpSkew,
		Digits:    otp.DigitsSix,
		Algorithm: otp.AlgorithmSHA1,
	})
	if err != nil {
		// A stored secret that fails base32 decode is a data-integrity
		// anomaly, not a credential failure — surface it (the caller maps to
		// generic upstream handling, fail-safe direction).
		return false, fmt.Errorf("account: validate totp: %w", err)
	}
	return valid, nil
}

// VerifyBackupCode checks code against unused backup codes and, on a match,
// marks that code's used_at via a guarded single-use UPDATE inside tx
// (INV-account-06: NULL → timestamp at most once, only while the user's
// mfa_totp_secrets.enabled_at IS NOT NULL — implicit invalidation after MFA
// disable, all enforced in the one joined UPDATE at the repository). Returns
// false without writing when the code does not match or is no longer valid.
//
// Input is normalized (lowercase, non-alphanumerics stripped) before hashing
// so the same helper the generator used stays in lockstep (D2). The caller
// (LoginMfa) owns the transaction; this method runs the redemption inside it.
func (v *totpVerifier) VerifyBackupCode(ctx context.Context, tx pgx.Tx, userID uuid.UUID, code string) (bool, error) {
	redeemed, err := v.repo.RedeemMFABackupCode(ctx, tx, userID, sha256Hex(normalizeBackupCode(code)))
	if err != nil {
		return false, fmt.Errorf("account: redeem mfa backup code: %w", err)
	}
	return redeemed, nil
}

// NewMfaVerifier constructs the real MfaVerifier backed by the given
// repository. NOTE vs techplan §10: the earlier sketch listed a keys param,
// but the repository adapter owns all ciphertext work per D3, so this type
// holds no key material — dropped here to avoid a dead dependency (flagged
// for the Tier 0 pairing session).
//
// Registry wiring (task-05) passes the production adapter; NewService's
// nil→stub semantics is unchanged, so struct-literal test services keep their
// fail-closed default.
func NewMfaVerifier(repo Repository) MfaVerifier {
	return &totpVerifier{repo: repo}
}

// Compile-time assertion: totpVerifier satisfies MfaVerifier.
var _ MfaVerifier = (*totpVerifier)(nil)
