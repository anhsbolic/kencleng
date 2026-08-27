# Stage 2 — Area 1: MFA Verifier Interface & Stub

> File: `internal/domain/account/mfa_verifier.go`

## Current State

`MfaVerifier` interface with two methods:

- `VerifyTOTP(ctx, userID, code) (bool, error)` — pure read, no DB writes
- `VerifyBackupCode(ctx, tx, userID, code) (bool, error)` — takes a `pgx.Tx` because matching marks `used_at` (single-use, INV-account-06)

`stubMfaVerifier` — both methods return `(false, nil)`, meaning no code ever verifies. FAILS CLOSED: an attacker cannot bypass MFA by submitting codes because the stub always rejects.

Wired as default in `NewService` (`service.go:139-141`): `if mfa == nil { mfa = stubMfaVerifier{} }`

`LoginMfa` in `login.go:206-230` already consumes the interface correctly: backup codes go through a tx (begin → verify → commit/rollback), TOTP goes through a direct call. The consumer is ready for a real implementation.

`TODO(task-06-mfa-totp)` at line 19 explicitly marks the replacement point.

## Requirement

Task #6 must replace `stubMfaVerifier` with a real implementation that:
1. Generates TOTP secrets during enrollment, encrypts them at rest (`secret_encrypted`)
2. Verifies TOTP codes against the encrypted secret (with ±1 step tolerance)
3. Generates 10 backup codes on enrollment, hashes them at rest
4. Verifies backup codes with single-use `used_at` marking (guarded UPDATE)
5. Checks `enabled_at IS NOT NULL` before accepting backup codes (INV-account-06 implicit invalidation)

## Gap

The entire TOTP/backup-code implementation is missing. The interface is defined, the stub exists, the consumer (`LoginMfa`) is wired — but no real verifier exists.

## Sniffing

- **Risk:** `VerifyBackupCode` takes `pgx.Tx` — the real implementation must execute a guarded `UPDATE ... WHERE used_at IS NULL` inside that tx. If the implementation uses a standalone query (outside tx), concurrent backup-code redemption could race past the single-use guard. The interface contract is correct (tx is passed), but the implementation must honor it.
- **Edge cases:** TOTP has standard ±1 step tolerance (30-second windows). What happens if the user's clock is significantly drifted? The `pquerna/otp` library handles this with configurable `Period` and `Digits`, but the implementation needs to decide the tolerance window. Also: what if `VerifyTOTP` is called for a user with no `mfa_totp_secrets` row? Must return `(false, nil)`, not an error.
- **Miscontext:** The interface says `VerifyTOTP` does "No DB writes" — correct for TOTP. But `VerifyBackupCode` does a write (`used_at` marking). The naming is clear enough, but someone implementing `VerifyTOTP` might accidentally add a write. No such column exists, so this is low risk.
- **Misleading signal:** The `mfa_verifier.go` file looks like it might already have some real logic because it's 47 lines with detailed comments. In reality it's entirely stub — the comments are spec references, not implementation.
- **Inconsistency:** None found. The interface, stub, and consumer are consistent.
