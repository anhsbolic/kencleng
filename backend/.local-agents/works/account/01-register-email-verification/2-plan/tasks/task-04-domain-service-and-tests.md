# Task 04 — Domain Account: Service + Tests (Business Logic)

> Ticket    : 01-register-email-verification
> Sub-task  : 4 of 5
> Axis      : Dependency/sequence chain (primary) + layer (vertical slice)
> Status    : Blocked on Task 02 + Task 03 (and transitively on the Tier 0 crypto prerequisite)
> Back-ref  : `../2-plan/techplan.md` (originating contract techplan — cross-check high-level decisions there whenever needed)

---

## 1. Scope

The business-logic layer of the account domain — the highest-risk task
in this feature. The service orchestrates Register (4 anti-enumeration
branches with constant-time shaping), VerifyEmail (single-use token
redemption with the full INV-account-08 guard), and ResendVerification,
plus the unit + concurrency test suite that proves R1-R19.

This is where the security-critical correctness lives. Per AGENTS.md
§3, **this task does NOT touch the Tier 0 fenced paths** — it consumes
`platform/crypto` via the repository (Task 03) and `platform/secrets`,
`platform/breachcheck`, `platform/notification` via their public APIs
(Task 02). It does not implement JWT, TOTP, or transaction/locking
primitives itself; it uses `pgx.Tx` for transaction boundaries and the
repository's atomic `RedeemToken` for the single-use guard.

**In scope:**
- `service.go` — `Service` struct + `Register`, `VerifyEmail`,
  `ResendVerification` methods + `generateToken` helper
- `service_test.go` — table-driven + concurrency tests covering
  R1-R19; `-race` required (Tier 1 feature per AGENTS.md §3)

**Out of scope:**
- Entity / repository (Task 03)
- HTTP handlers / middleware / wiring (Task 05)
- Token *storage* (Task 03's `InsertAuthToken`) — this task generates
  the plain token + hash and hands them to the repository

## 2. Dependencies

- **Hard deps:**
  - Task 02 (`platform/secrets`, `platform/breachcheck`, `platform/notification`)
  - Task 03 (`domain/account/entity.go`, `repository.go`,
    `repository_db.go` — the `Repository` interface and entities)
- **Soft deps:** none
- **Blocks:** Task 05 (handlers call the service)

## 3. Files

| File | Change Type |
|---|---|
| `backend/internal/domain/account/service.go` | New |
| `backend/internal/domain/account/service_test.go` | New |

## 4. Implementation detail

### `backend/internal/domain/account/service.go` (new)

```go
package account

import (
    "context"
    "crypto/rand"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "time"

    "github.com/google/uuid"
    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgxpool"

    "kencleng/internal/platform/breachcheck"
    "kencleng/internal/platform/crypto"
    "kencleng/internal/platform/notification"
    "kencleng/internal/platform/secrets"
)

// Sentinel errors — mapped to HTTP status by transport/http/errors.go
// (Task 05). Defined here so the service owns its error vocabulary;
// the transport layer only translates.
var (
    ErrValidation   = errors.New("validation error")
    ErrTokenExpired = errors.New("token expired")
    ErrTokenNotFound = errors.New("token not found")
)

// Service implements the account domain's registration, verification,
// and resend flows. It is safe for concurrent use.
type Service struct {
    repo         Repository
    db           *pgxpool.Pool
    hasher       *secrets.Hasher // or just the package funcs — see Task 02
    breachCheck  *breachcheck.Client
    emailSender  notification.Sender
    cryptoKeys   *crypto.Keys
    tokenTTL     time.Duration // 24h for email_verification
}

func NewService(repo Repository, db *pgxpool.Pool, hasher *secrets.Hasher, bc *breachcheck.Client, sender notification.Sender, keys *crypto.Keys) *Service {
    return &Service{
        repo: repo, db: db, hasher: hasher,
        breachCheck: bc, emailSender: sender, cryptoKeys: keys,
        tokenTTL: 24 * time.Hour,
    }
}
```

#### `Register(ctx, name, email, password string) error`

Orchestrates R1-R7, R16-R19. The four branches (new / unverified-existing
/ verified-existing / Google-only-conflict) must take equivalent
wall-clock time and return an identical 202-generic outcome to the
caller (the handler in Task 05 writes the actual 202; the service
returns `nil` on all four branches).

```
Register(ctx, name, email, password string) error
  1. Validate password length >= 8  → ErrValidation if fails (R5, R18)
  2. Breach check (k-anonymity, fail-open) → ErrValidation if breached (R6, R19)
  3. ALWAYS run hasher.HashPassword(password)  [constant-time — result
     used by the new-user branch, discarded by the other three] (R7)
  4. Compute identifierHash = crypto.HMAC(email, keys.HMACKey)
  5. Lookup auth_identity by (email_password, identifierHash):
     - found, unverified → resend branch:
         * BEGIN tx
         * RevokeTokens(userID, "email_verification")
         * InsertAuthToken(new token, 24h)
         * COMMIT
         * emailSender.SendNudgeEmail(..., NudgeResendVerification)  [after commit]
         (R2)
     - found, verified → nudge branch:
         * no DB write (DB-write-shaped no-op — see "DB-time uniformity" below)
         * emailSender.SendNudgeEmail(..., NudgePasswordReset)  (R3)
     - not found:
         * check google identity by (google, identifierHash):
           - found → Google-only nudge branch:
               * no DB write (DB-write-shaped no-op)
               * emailSender.SendNudgeEmail(..., NudgeGoogleOnly)  (R4, R17)
           - not found → new-user branch:
               * BEGIN tx
               * InsertUser (encrypts primary_email)
               * InsertAuthIdentity (encrypts identifier; credential_secret = bcrypt hash)
               * InsertAuthToken (token_hash = SHA-256(plainToken), 24h)
               * COMMIT
               * emailSender.SendVerificationEmail(..., plainToken)  [after commit]  (R1)
  6. All branches return nil (handler writes 202 generic) (R7)
```

**Constant-time — two halves (techplan §5 Decision 8, §7 risk row 7):**

1. **CPU time.** `hasher.HashPassword(password)` (bcrypt, default cost
   ~100ms) runs on **every** branch. The new-user branch stores the
   result as `credential_secret`; the other three branches discard it.
   This eliminates the CPU-time side-channel. Do NOT skip bcrypt on
   no-op branches. Do NOT replace with an artificial sleep.
2. **DB time.** All branches perform DB-write-shaped work per feature
   spec Assumption B. The new-user and unverified-existing branches do
   real writes; the verified-existing and Google-only branches perform
   a DB-write-shaped no-op (e.g. a read-only round-trip that the
   observer cannot distinguish from a write at the API boundary). The
   exact mechanism is left to the implementer — the test
   `TestRegister_GenericResponse_Timing` proves equivalence, not the
   mechanism. Document the chosen mechanism in the task's risk note.

**Email send after commit (techplan §7 risk row 2, §13 row 1).** All
DB writes (insert user + auth_identity + token, or revoke + insert
token) commit BEFORE the email send. Sending inside the tx holds locks
during network I/O — forbidden. If the post-commit email fails, the DB
state is correct and the user can use resend. The
`notification.Sender` contract (Task 02) makes this explicit.

**Concurrent duplicate (R16).** Two goroutines registering the same
email under `email_password`: both compute the same `identifierHash`;
one's `InsertAuthIdentity` succeeds, the other's fails on
`ux_auth_identities_provider_identifier`. Both ran inside their own
transactions, so the loser rolls back cleanly — no orphaned `users`
row. Map the unique-violation to a no-op (return nil → 202 generic),
NOT to an error response (the duplicate is indistinguishable from a
normal R1 to the caller). Detect via `errors.As` on `*pgconn.PgError`,
code `23505`.

**No PII in logs (AGENTS.md golden rule, §13 row 5/6).** Log the fact
("registration attempt", "registration completed") and the outcome,
plus `user_id` *after* creation. Never `log.Printf("register email=%s", email)`.
Never `fmt.Sprintf("%+v", user)`. Breach-check failures log
"breach check API unreachable" + status, never the password or hash.

#### `VerifyEmail(ctx, token string) error`

Orchestrates R8-R12.

```
VerifyEmail(ctx, token string) error
  1. tokenHash = sha256(hex(token))  [or sha256(raw bytes) — pick and be consistent]
  2. ok, err := repo.RedeemToken(ctx, tokenHash)
       — atomic UPDATE ... WHERE used_at IS NULL AND revoked_at IS NULL
         AND expires_at > now()  [full 3-clause — INV-account-08 Statement]
       — returns true iff 1 row affected
  3. if err: wrap + return (handler maps to 500)
  4. if ok:
       * fetch the token row (RedeemToken returned true but we need
         user_id + identity to set verified_at — or have RedeemToken
         return the row; coordinate with Task 03)
       * repo.SetVerifiedAt(identityID, now)  (R8)
       * return nil  [handler writes 200]
  5. if !ok:
       * t, _ := repo.FindAuthTokenByHash(ctx, tokenHash)
       * if t != nil && t.ExpiresAt <= now(): return ErrTokenExpired  [handler 410]  (R9)
       * else: return ErrTokenNotFound  [handler 404]  (R10 already-used, R11 revoked, R10 non-existent)
```

**Single-use is the DB predicate, not a service-level check.** The
concurrent double-submit (R12) is resolved by the atomic UPDATE: two
concurrent calls, exactly one affects a row (200), the other affects
zero (404). The service does NOT do a read-then-write — that would
race. `RedeemToken` is the source of truth.

**The 3-clause guard is non-negotiable.** The invariant's Verification
field omits `revoked_at IS NULL` (2-clause) — that is a documented
spec error (techplan §14 Open Item #2). The implementation uses the
Statement's 3-clause version. Do not "fix" the spec — the agent is
forbidden from editing `docs/spec/*` (AGENTS.md §4).

#### `ResendVerification(ctx, email string) error`

Orchestrates R13-R15 (rate limit itself is Task 05 middleware).

```
ResendVerification(ctx, email string) error
  1. identifierHash = crypto.HMAC(email, keys.HMACKey)
  2. identity, err := repo.FindAuthIdentityByIdentifierHash("email_password", identifierHash)
  3. if identity != nil && identity.VerifiedAt == nil:  [unverified match]
       * BEGIN tx
       * repo.RevokeTokens(identity.UserID, "email_verification")
       * plainToken, tokenHash := generateToken()
       * repo.InsertAuthToken(&AuthToken{..., TokenHash: tokenHash, ExpiresAt: now+24h, ...})
       * COMMIT
       * emailSender.SendVerificationEmail(..., plainToken)  [after commit]  (R13)
  4. else:  [no match / verified / google-only]
       * no token, no email  (R14)
       * (DB-write-shaped no-op for uniformity, mirroring Register's approach)
  5. return nil  [handler writes 202 generic — identical for both branches]
```

#### `generateToken() (plainToken, tokenHash string, err error)`

Techplan §5 Decision 6 — in-service, ~10 lines, YAGNI a separate
package until a second consumer appears.

```go
func generateToken() (string, string, error) {
    b := make([]byte, 32)
    if _, err := rand.Read(b); err != nil {
        return "", "", fmt.Errorf("generate token: %w", err)
    }
    plain := hex.EncodeToString(b)            // user-facing
    sum := sha256.Sum256([]byte(plain))
    return plain, hex.EncodeToString(sum[:]), nil // stored
}
```

- 32 bytes from `crypto/rand` (not `math/rand`).
- Plain token is hex-encoded for the email link; **SHA-256 of the
  plain token** is stored (never the plain token itself).
- The plain token leaves the process exactly once: in the
  `SendVerificationEmail` call, after commit. It is never logged.

## 5. Rules covered (this task's slice — all end-user rules)

| Rule | How this task satisfies it | Named test |
|---|---|---|
| R1 | new-user branch: insert user+identity+token, send verification email, return nil (→ 202 generic) | `TestRegister_NewUser_CreatesUserIdentityToken` |
| R2 | unverified-existing branch: revoke+new token+resend-verification nudge, return nil | `TestRegister_UnverifiedExisting_ResendFlow` |
| R3 | verified-existing branch: password-reset nudge, return nil | `TestRegister_VerifiedExisting_PasswordResetNudge` |
| R4 | Google-only conflict branch: Google-only nudge, return nil | `TestRegister_GoogleOnlyConflict_Nudge` |
| R5 | password length <8 → ErrValidation BEFORE branch lookup | `TestRegister_PasswordPolicy` (length half) |
| R6 | breach API unreachable → proceed without check, logged | `TestRegister_BreachCheck_FailOpen` |
| R7 | all 4 branches return identical nil + equivalent wall-clock | `TestRegister_GenericResponse_AllBranches`, `TestRegister_GenericResponse_Timing` |
| R8 | valid token → SetVerifiedAt, return nil (→ 200) | `TestVerifyEmail_ValidToken_SetsVerifiedAt` |
| R9 | expired token → ErrTokenExpired (→ 410) | `TestVerifyEmail_ExpiredToken_410` |
| R10 | non-existent / already-used → ErrTokenNotFound (→ 404) | `TestVerifyEmail_NotFound_404`, `TestVerifyEmail_AlreadyUsed_404` |
| R11 | revoked (superseded) → ErrTokenNotFound (→ 404) — 3-clause guard | `TestVerifyEmail_RevokedToken_Rejected` |
| R12 | concurrent double-submit → exactly one 200, other 404 | `TestVerifyEmail_TokenSingleUse_Concurrent` |
| R13 | resend unverified match → revoke+new token+email, return nil | `TestResend_UnverifiedMatch_IssuesNewToken` |
| R14 | resend no-match/verified/google-only → no token, no email, return nil | `TestResend_NoMatch_NoTokenNoEmail`, `TestResend_Verified_NoTokenNoEmail`, `TestResend_GoogleOnly_NoTokenNoEmail` |
| R16 | concurrent duplicate registration → exactly one succeeds, other clean no-op | `TestRegister_ConcurrentDuplicateEmail_Race` (≥100 goroutines) |
| R17 | Google-only branch response identical to others | `TestRegister_GoogleOnlyConflict_GenericResponse` |
| R18 | password <8 or breached → ErrValidation with field-level info | `TestRegister_PasswordPolicy` (breach half) |
| R19 | breach API unreachable → registration proceeds | `TestRegister_BreachCheck_FailOpen` (same as R6) |

R15 (rate limit) is Task 05 — listed here for completeness; this task
does not test it.

## 6. Testing checklist (this task's slice)

Every R1-R19 line from techplan §12 that names a test belongs here:

- [ ] R1: `TestRegister_NewUser_CreatesUserIdentityToken` — asserts
      User + AuthIdentity (verified_at=nil) + AuthToken (24h) created,
      verification email sent, service returns nil.
- [ ] R2: `TestRegister_UnverifiedExisting_ResendFlow` — no new
      User/Identity; old tokens revoked; new token issued; resend
      nudge sent; returns nil.
- [ ] R3: `TestRegister_VerifiedExisting_PasswordResetNudge` — no new
      record; password-reset nudge sent; returns nil.
- [ ] R4: `TestRegister_GoogleOnlyConflict_Nudge` — no new User;
      Google-only nudge sent; returns nil.
- [ ] R5/R18: `TestRegister_PasswordPolicy` — table-driven: <8 chars
      → ErrValidation; breached → ErrValidation; check fires before
      any repo lookup (mock the repo and assert it was NOT called on
      failure).
- [ ] R6/R19: `TestRegister_BreachCheck_FailOpen` — breach client
      returns unreachable; registration proceeds; log captured
      contains no password/hash.
- [ ] R7: `TestRegister_GenericResponse_AllBranches` — all 4 branches
      return nil with identical observable side-effect shape (from the
      handler's perspective); `TestRegister_GenericResponse_Timing` —
      all 4 branches within an equivalent wall-clock band (the test
      proves equivalence, not the mechanism — see §4 constant-time
      note).
- [ ] R8: `TestVerifyEmail_ValidToken_SetsVerifiedAt` — returns nil;
      `verified_at` set on the identity.
- [ ] R9: `TestVerifyEmail_ExpiredToken_410` — returns ErrTokenExpired;
      no state change.
- [ ] R10: `TestVerifyEmail_NotFound_404`, `TestVerifyEmail_AlreadyUsed_404`
      — returns ErrTokenNotFound; no state change.
- [ ] R11: `TestVerifyEmail_RevokedToken_Rejected` — revoked token
      returns ErrTokenNotFound (proves the 3-clause guard; regression
      for the INV-account-08 spec error).
- [ ] R12: `TestVerifyEmail_TokenSingleUse_Concurrent` — same token
      submitted twice concurrently → exactly one nil, one
      ErrTokenNotFound. Run under `-race`.
- [ ] R13: `TestResend_UnverifiedMatch_IssuesNewToken` — old tokens
      revoked; new token issued; verification email sent; returns nil.
- [ ] R14: `TestResend_NoMatch_NoTokenNoEmail`,
      `TestResend_Verified_NoTokenNoEmail`,
      `TestResend_GoogleOnly_NoTokenNoEmail` — no token, no email,
      returns nil (identical to R13 from the handler's view).
- [ ] R16: `TestRegister_ConcurrentDuplicateEmail_Race` — ≥100
      goroutines registering the same email → exactly one full
      creation, others clean no-ops (no orphaned rows, no error
      response). Run under `-race`. Per `docs/spec/domains/account/tasks.md` KPI.
- [ ] R17: `TestRegister_GoogleOnlyConflict_GenericResponse` —
      Google-only branch's observable outcome identical to R1-R3.
- [ ] `go test -race ./internal/domain/account/...` clean (Tier 1
      feature — AGENTS.md §3).
- [ ] `make verify` passes (lint, unit, race, contract, security,
      integration).

## 7. Common mistakes to avoid (techplan §13 slice)

| Mistake | Fix |
|---|---|
| Email send inside DB transaction | Send after `tx.Commit()`. If email fails post-commit, user uses resend. |
| 2-clause token guard (missing `revoked_at IS NULL`) | Full 3-clause predicate per INV-account-08 Statement. |
| Early return on "email not found" branch (no bcrypt) | Always run `HashPassword` on all 4 branches; discard on no-ops. |
| Logging email in plaintext | Log the fact + outcome + user_id (after creation), never the email. |
| Logging breach-check error verbatim | Log sanitized summary + status code. |
| String-matching DB errors | `errors.As` on `*pgconn.PgError`, code `23505` for unique-violation. |
| Not wrapping errors with `%w` | `fmt.Errorf("...: %w", err)`. |
| `users` + `auth_identities` in separate transactions | Single `pgx.Tx` wrapping all inserts; rollback is clean (R16). |

## 8. Risk note

- Assumptions made: the DB-time uniformity mechanism for the
  verified-existing and Google-only branches is left to the
  implementer (feature spec Assumption B); the chosen mechanism MUST
  be documented here once chosen, and `TestRegister_GenericResponse_Timing`
  is the proof. The repo's `RedeemToken` returns `(bool, error)`; if
  the service needs the token row (for `user_id`/`identity_id`) on a
  successful redeem, coordinate with Task 03 to either return the row
  or do a follow-up fetch inside the same transaction.
- Edge cases intentionally NOT handled: HMAC key rotation (v1 out of
  scope per tech stack); rate limiting (Task 05); real SMTP delivery
  (Task 02 FakeSender); account linking / set-password (task #5).
- Concurrency assumptions: `Service` is safe for concurrent use
  (stateless beyond its injected clients, all of which are
  goroutine-safe). Single-use token correctness is delegated to the
  repository's atomic `RedeemToken` — the service never does a
  read-then-write on tokens. R16's clean-rollback guarantee depends on
  the single-transaction insert pattern in §4.
- What is not tested, and why: the exact wall-clock time of bcrypt is
  not asserted (machine-dependent) — only branch-equivalence is. Live
  breach-check HTTP is not tested (mocked — see Task 02). The
  end-to-end HTTP layer is Task 05; here the service is tested via
  direct calls with a mock/fake `Repository` and the Task 02 fakes.
