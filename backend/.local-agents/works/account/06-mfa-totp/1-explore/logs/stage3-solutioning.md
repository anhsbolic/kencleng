# Stage 3 — Solutioning

> Feature: 06-mfa-totp
> Date: 2026-08-27

## Decision Log

### D1: TOTP Library Choice

**Options:**
- **A) `pquerna/otp`** — most popular Go TOTP library, RFC 6238 compliant, supports base32 secret generation, configurable period/digits, and `ValidateCustom` for tolerance control
- **B) Hand-rolled TOTP** — implement RFC 6238 from scratch using `crypto/hmac` + `encoding/base32`

**Decision: A (`pquerna/otp`)**. Hand-rolling TOTP is unnecessary risk for zero benefit. The library is well-maintained, the feature spec references it explicitly ("`pquerna/otp` usage"), and the Tier 0 fence means the human reviews this code anyway. The library's `totp.Generate` and `totp.ValidateCustom` cover generation and verification with configurable tolerance.

### D2: Backup Code Hashing — SHA-256 vs bcrypt

**Options:**
- **A) SHA-256** — same pattern as `auth_tokens.token_hash`. Fast, deterministic, sufficient entropy since backup codes are random (not user-chosen)
- **B) bcrypt** — slower, designed for passwords. Overkill for random high-entropy codes

**Decision: A (SHA-256)**. Backup codes are ~47 bits of entropy each (8 alphanumeric chars from `crypto/rand`). They're not user-chosen passwords susceptible to dictionary attacks. SHA-256 matches the existing `auth_tokens.token_hash` pattern. The `code_hash` column is TEXT (hex-encoded SHA-256), consistent with `token_hash`.

### D3: Reauth Marker Consumption — Transport vs Service Layer

**Options:**
- **A) Handler checks + consumes marker, then calls service** — keeps the reauth marker (a transport-layer concern from task #02) out of the domain service. The handler calls `CheckReauthMarker` → `ConsumeReauthMarker` → `svc.MfaDisable(ctx, userID, "")`. Service only sees "reauth passed."
- **B) Inject reauth checker into service** — service receives an interface `ReauthChecker` that wraps the marker check. More testable, but pulls a transport-layer concept into the domain.
- **C) Service receives a `reauthFn func(uuid.UUID) bool` closure** — similar to B but lighter, matches the existing seam pattern (`mintAccess`, `compare`, etc.)

**Decision: A (Handler checks, service is agnostic)**. The reauth marker is a transport-layer artifact (in-memory `sync.Map` in `auth_google.go`). The service should not know about it. The handler determines the re-auth method based on the user's auth provider:
- Has `email_password` identity → require `password` in body → service verifies bcrypt
- No `email_password` identity (Google-only) → handler checks + consumes reauth marker → calls service with empty password (service trusts that reauth already passed)

This keeps the service clean: `MfaDisable(ctx, userID, password string)` where non-empty password means "verify this password" and empty means "caller already re-authed." The handler is responsible for the reauth gate.

**New function needed:** `ConsumeReauthMarker(userID) bool` — atomically checks and deletes. Current `CheckReauthMarker` only checks, doesn't consume.

### D4: Repository Method Shape

Keep repository methods minimal and composable:
- `UpsertMFATOTPSecret(ctx, tx, userID, secretEncrypted []byte)` — INSERT ON CONFLICT UPDATE
- `EnableMFATOTP(ctx, tx, userID) error` — UPDATE SET enabled_at = now() WHERE enabled_at IS NULL
- `DisableMFATOTP(ctx, tx, userID) error` — UPDATE SET enabled_at = NULL
- `GetMFATOTPSecret(ctx, userID) (*MFATOTPSecret, error)` — SELECT for verifier
- `InsertMFABackupCodes(ctx, tx, codes []MFABackupCode) error` — batch insert
- `RedeemMFABackupCode(ctx, tx, userID, codeHash string) (bool, error)` — single-use UPDATE

New entity type:
```go
type MFATOTPSecret struct {
    UserID          uuid.UUID
    SecretEncrypted []byte
    EnabledAt       *time.Time
}

type MFABackupCode struct {
    ID        uuid.UUID
    UserID    uuid.UUID
    CodeHash  string
}
```

### D5: MfaDisable Service Signature — Provider Detection

**Options:**
- **A) Service detects provider internally** — `MfaDisable(ctx, userID, password)` loads identities, branches on provider type
- **B) Handler passes a flag** — leaks transport concern into service
- **C) Two separate service methods** — cleaner but doubles the surface

**Decision: A (Service detects internally)**. The service already has `FindAuthIdentitiesByUser` available. It can check whether the user has an `email_password` identity and branch accordingly. This matches the pattern in `SetPassword` (which also branches server-side). The handler just passes whatever it has (password or empty string).

### D6: ConsumeReauthMarker — New Function vs Modify Existing

**Decision:** Add a new `ConsumeReauthMarker(userID uuid.UUID) bool` function alongside `CheckReauthMarker`. `CheckReauthMarker` stays read-only (used by the handler to check before consuming). `ConsumeReauthMarker` does check + delete atomically. Both live in `auth_google.go` since that's where the marker store lives.

### D7: Audit Log + Notification for MFA Enable/Disable

**Decision:** Follow the established pattern from `UnlinkGoogle` and `VerifyEmail`:
- Write `user_logs` entry in the same transaction as the state change
- Action types: `"mfa_enabled"` (on confirm success), `"mfa_disabled"` (on disable success)
- Notification trigger: deferred (cross-domain dependency on `notification` domain, same as `05-account-linking.md`). Log the intent but don't block on the notification domain being built.

### D8: Error Vocabulary

New sentinel errors in the service:
```go
var (
    ErrMfaAlreadyEnabled  = errors.New("mfa already enabled")
    ErrMfaNotPending      = errors.New("mfa enrollment not pending")
    ErrInvalidTOTPCode    = errors.New("invalid totp code")
)
```

Transport mapping:
- `ErrMfaAlreadyEnabled` → 409, problem type `https://kencleng.dev/errors/mfa-already-enabled`
- `ErrInvalidTOTPCode` / `ErrMfaNotPending` → 422, problem type `https://kencleng.dev/errors/validation`
- `ErrInvalidCredentials` (wrong password on disable) → 401

---

## Implementation Approach

### File Changes Summary

| File | Change |
|---|---|
| `internal/domain/account/mfa_verifier.go` | Replace `stubMfaVerifier` with real `totpVerifier` struct. Add `NewMfaVerifier(repo, keys)` constructor. |
| `internal/domain/account/repository.go` | Add 6 new methods to `Repository` interface |
| `internal/domain/account/repository_db.go` | Implement the 6 new methods using goqu |
| `internal/domain/account/entity.go` | Add `MFATOTPSecret` and `MFABackupCode` entity types |
| `internal/domain/account/service.go` | Add `MfaEnroll`, `MfaEnrollConfirm`, `MfaDisable` methods. Add new sentinel errors. Add `actionMfaEnabled`/`actionMfaDisabled` constants. |
| `internal/transport/http/account_security.go` | Extend `securityService` interface. Add 3 handlers. |
| `internal/transport/http/auth_google.go` | Add `ConsumeReauthMarker` function |
| `cmd/server/main.go` | Wire real `MfaVerifier`, register 3 new routes |
| `go.mod` | Add `github.com/pquerna/otp` dependency |

### Tier 0 Boundary

The Tier 0 fenced sub-area is the TOTP secret generation/encryption and verification logic — specifically:
- `totpVerifier.VerifyTOTP` — decrypts secret, computes TOTP, validates
- `totpVerifier.VerifyBackupCode` — hashes code, queries, guards `used_at` update
- `NewMfaVerifier` constructor
- The `MFATOTPSecret` entity type

These files must be human-authored or human-rewritten, not agent-generated wholesale. The rest (repository methods, handlers, routing, service orchestration) is standard Tier 1.

### Task Breakdown (Implementation Order)

1. **Entity types** — `MFATOTPSecret`, `MFABackupCode` in `entity.go`
2. **Repository methods** — add to `repository.go` interface + `repository_db.go` implementations
3. **MfaVerifier implementation** — replace stub with real `totpVerifier` (Tier 0)
4. **Service methods** — `MfaEnroll`, `MfaEnrollConfirm`, `MfaDisable`
5. **ConsumeReauthMarker** — add to `auth_google.go`
6. **Transport handlers** — 3 new handlers + extend `securityService` interface
7. **Wiring** — `main.go` routes + real verifier injection
8. **Tests** — unit tests for verifier, service, handlers; integration tests for DB layer

---

## Risk Note

- **Assumptions made:**
  - `pquerna/otp` is the intended TOTP library (feature spec references it by name)
  - Backup codes use SHA-256 hashing (consistent with `auth_tokens.token_hash` pattern)
  - Reauth marker consumption happens at the transport layer, not in the service
  - The `notification` domain integration (user-facing MFA enable/disable emails) is deferred — same cross-domain pattern as task #05

- **Edge cases intentionally NOT handled (and why):**
  - TOTP clock drift beyond ±1 step: standard tolerance, not over-engineering for sandbox
  - Backup code brute-force at login: handled by MFA-stage lockout (task #03's `login_attempts` mechanism), not by this task
  - Re-enrollment after disable: old backup codes remain in DB (implicit invalidation per INV-account-06, no cleanup job)

- **Concurrency assumptions:**
  - `RedeemMFABackupCode` uses `used_at IS NULL` guard under READ COMMITTED — same proven pattern as `RedeemToken` (INV-account-08)
  - `EnableMFATOTP` uses `enabled_at IS NULL` guard — prevents double-enable race
  - `ConsumeReauthMarker` uses `sync.Map.LoadAndDelete` for atomic check+delete

- **What is not tested, and why:**
  - Actual TOTP code generation from a known secret (deterministic test with fixed time) — will be tested in the verifier unit tests during implementation
  - Integration with the `notification` domain — deferred until that domain exists
