# Implementation Report — 01-register-email-verification

> Ticket    : 01-register-email-verification
> Feature   : Register & Email Verification (account domain, first vertical slice)
> Date      : 2026-08-20
> Spec ref  : `docs/spec/domains/account/features/01-register-email-verification.md`
> Techplan  : `.local-agents/works/account/01-register-email-verification/2-plan/techplan.md`
> Tasks     : `.local-agents/works/account/01-register-email-verification/3-tasks/` (5 task files)

---

## 1. Summary

The account domain's first vertical slice: public registration with
`email_password`, email verification via single-use token, and
resend-verification. Delivered three public endpoints, three migrations,
three platform packages, the domain package (entity/repository/service),
transport handlers + middleware, and `main.go` wiring. The defining
constraint — anti-enumeration — is enforced throughout: every
register/resend branch returns an identical `202` generic response, and
all four internal register branches take equivalent wall-clock time
(always-bcrypt + DB-shaped work).

All five task files were executed in dependency order (01 → 02 → 03 → 04
→ 05). The Tier 0 crypto prerequisite was resolved mid-session by an
explicit per-session §3 fence lift (see §6 below).

---

## 2. Files changed

### New files (24)

| File | LoC | Task | Description |
|---|---|---|---|
| `migrations/000001_create_users.up.sql` | 22 | 01 | `set_updated_at()` + `users` + unique index + trigger |
| `migrations/000001_create_users.down.sql` | 4 | 01 | drop trigger + table |
| `migrations/000002_create_auth_identities.up.sql` | 20 | 01 | `auth_identities` + unique `(provider_type, identifier_hash)` + FK + trigger |
| `migrations/000002_create_auth_identities.down.sql` | 2 | 01 | drop trigger + table |
| `migrations/000003_create_auth_tokens.up.sql` | 16 | 01 | `auth_tokens` + unique `token_hash` + partial valid index |
| `migrations/000003_create_auth_tokens.down.sql` | 6 | 01 | drop indexes + table + `set_updated_at()` function |
| `internal/platform/secrets/secrets.go` | 26 | 02 | bcrypt wrapper (`HashPassword`, `ComparePassword`) |
| `internal/platform/secrets/secrets_test.go` | 38 | 02 | round-trip + match/reject tests |
| `internal/platform/breachcheck/client.go` | 70 | 02 | HaveIBeenPwned k-anonymity client, fail-open, explicit timeout |
| `internal/platform/breachcheck/client_test.go` | 120 | 02 | not-breached / breached / 5xx-fail-open / conn-error-fail-open + no-PII-in-log |
| `internal/platform/notification/sender.go` | 45 | 02 | `Sender` interface + `FakeSender` (logged, no SMTP) + nudge constants |
| `internal/platform/notification/sender_test.go` | 60 | 02 | returns nil + no recipient/token in log |
| `internal/platform/crypto/crypto.go` | 96 | 03* | `Encrypt`/`Decrypt` (AES-GCM) + `HMAC` (HMAC-SHA256 hex) |
| `internal/platform/crypto/crypto_test.go` | 196 | 03* | 12 tests: round-trip, tamper, wrong-key, nonce uniqueness, HMAC determinism/format, short-ciphertext, nil-keys, bad-key-size |
| `internal/domain/account/entity.go` | 68 | 03 | `User`, `AuthIdentity`, `AuthToken` (domain plaintext form) |
| `internal/domain/account/repository.go` | 73 | 03 | `Repository` interface (8 methods) |
| `internal/domain/account/repository_db.go` | 313 | 03 | goqu + pgx adapter; goqu `Prepared(true)` → `$1..` placeholders |
| `internal/domain/account/repository_db_integration_test.go` | 450 | 03 | `//go:build integration` — tx round-trip, concurrent duplicate (23505), RedeemToken 3-clause guards, RevokeTokens |
| `internal/domain/account/service.go` | 422 | 04 | `Service`: `Register` (4 branches, constant-time), `VerifyEmail`, `ResendVerification`, `generateToken` |
| `internal/domain/account/service_test.go` | 926 | 04 | R1-R19 unit tests, `-race` clean (22 tests incl. R12 concurrent, R16 100-goroutine) |
| `internal/transport/http/errors.go` | 93 | 05 | RFC 9457 Problem Details: `WriteProblem`, `WriteValidationError`, `MapServiceError` |
| `internal/transport/http/middleware.go` | 72 | 05 | `RateLimit(rps, burst)` per-IP token bucket + idle-key eviction |
| `internal/transport/http/auth_register.go` | 86 | 05 | `POST /auth/register` handler (boundary validation + 202 generic) |
| `internal/transport/http/auth_verify_email.go` | 76 | 05 | `POST /auth/verify-email` + `POST /auth/verify-email/resend` handlers |

\* Task 03's crypto prerequisite (see §6).

### Modified files (7, this feature)

| File | Task | Description |
|---|---|---|
| `go.mod` | 01/02/03/05 | +`goqu/v9 v9.19.0` (direct), +`x/time v0.14.0` (direct), +`x/crypto v0.46.0` (indirect→direct), +`google/uuid v1.6.0` (indirect→direct) |
| `go.sum` | 01/02/03/05 | corresponding hashes |
| `Makefile` | 01 | `migrate-down` changed from `down 1` → `down -all` |
| `cmd/server/main.go` | 05 | wired `account.Service`, registered 3 routes behind `RateLimit`, fail-fast env for RPS/burst |
| `internal/platform/crypto/doc.go` | 03* | updated to document the §3 fence lift (agent-authored under human review) |
| `.env` | 05 | +`AUTH_RATE_RPS=2`, +`AUTH_RATE_BURST=10` (with comments) |
| `docs/spec/domains/account/invariants.md` | open item #2 | INV-account-08 Verification field: added `revoked_at IS NULL AND` (§4 per-session exception) |

### Pre-existing changes (NOT this feature — out of scope, flagged)

| File | Note |
|---|---|
| `api/openapi.yaml` | pre-existing `maxItems: 50` on notification_ids — unrelated to this feature |
| `frontend/AGENTS.md` | pre-existing design-guidelines doc link — unrelated (frontend/ boundary, AGENTS.md §7) |
| `.env.example` | pre-existing — not touched by this session |

---

## 3. API endpoints delivered

| Method | Path | Handler | Success | Errors |
|---|---|---|---|---|
| `POST` | `/auth/register` | `RegisterHandler` | 202 GenericAcceptedMessage (identical for all 4 branches) | 400 invalid JSON, 422 ValidationProblem, 429 rate-limited, 500 internal |
| `POST` | `/auth/verify-email` | `VerifyEmailHandler` | 200 `{message}` | 400 invalid JSON, 404 token-not-found, 410 token-expired, 429, 500 |
| `POST` | `/auth/verify-email/resend` | `ResendVerificationHandler` | 202 GenericAcceptedMessage (identical for match/no-match) | 400 invalid JSON, 422 invalid email, 429 |

All three mounted behind `RateLimit(rps=2, burst=10)` per-IP middleware.
Error responses are RFC 9457 Problem Details (`application/problem+json`),
no internal leakage.

---

## 4. Rule coverage (R1-R19)

| Rule | Named test | Status |
|---|---|---|
| R1 (register new) | `TestRegister_NewUser_CreatesUserIdentityToken` | ✅ |
| R2 (register unverified existing) | `TestRegister_UnverifiedExisting_ResendFlow` | ✅ |
| R3 (register verified existing) | `TestRegister_VerifiedExisting_PasswordResetNudge` | ✅ |
| R4 (register Google-only conflict) | `TestRegister_GoogleOnlyConflict_Nudge` | ✅ |
| R5 (password validation order) | `TestRegister_PasswordPolicy` (length half) | ✅ |
| R6 (breach check fail-open) | `TestRegister_BreachCheck_FailOpen` | ✅ |
| R7 (constant-time anti-enumeration) | `TestRegister_GenericResponse_AllBranches` + `TestRegister_GenericResponse_Timing` | ✅ |
| R8 (verify valid token) | `TestVerifyEmail_ValidToken_SetsVerifiedAt` | ✅ |
| R9 (verify expired) | `TestVerifyEmail_ExpiredToken_410` | ✅ |
| R10 (verify not found / already used) | `TestVerifyEmail_NotFound_404` + `TestVerifyEmail_AlreadyUsed_404` | ✅ |
| R11 (verify revoked — 3-clause guard) | `TestVerifyEmail_RevokedToken_Rejected` | ✅ |
| R12 (verify concurrent double-submit) | `TestVerifyEmail_TokenSingleUse_Concurrent` (`-race`) | ✅ |
| R13 (resend unverified match) | `TestResend_UnverifiedMatch_IssuesNewToken` | ✅ |
| R14 (resend no match / verified / google-only) | `TestResend_NoMatch_NoTokenNoEmail` + `TestResend_Verified_NoTokenNoEmail` + `TestResend_GoogleOnly_NoTokenNoEmail` | ✅ |
| R15 (rate limit) | middleware implemented + handler wiring; `TestResend_RateLimited` not yet written (handler tests deferred) | ⚠️ |
| R16 (concurrent duplicate registration) | `TestRegister_ConcurrentDuplicateEmail_Race` (100 goroutines, `-race`) | ✅ |
| R17 (Google-only generic response) | `TestRegister_GoogleOnlyConflict_GenericResponse` | ✅ |
| R18 (password policy 422) | `TestRegister_PasswordPolicy` (breach half) | ✅ |
| R19 (breach check fail-open test) | `TestRegister_BreachCheck_FailOpen` | ✅ |

Integration tests (against real Postgres, `//go:build integration`):
- `TestInsertUserAndIdentity_InTransaction` — tx round-trip, plaintext cleared ✅
- `TestInsertAuthIdentity_ConcurrentDuplicate` — `*pgconn.PgError` code 23505 ✅
- `TestRedeemToken_Guards` — valid/used/revoked/expired/non-existent (3-clause incl. revoked regression) ✅
- `TestRevokeTokens_OnlyUnusedUnrevoked` — only fresh unused of matching purpose ✅

---

## 5. Verification results

| Gate | Command | Result |
|---|---|---|
| Build | `go build ./...` | ✅ OK |
| Vet | `go vet ./...` | ✅ OK |
| Unit tests (fast) | `go test -count=1 ./...` | ✅ all pass (account 8s, crypto 14ms, breachcheck 24ms, notification 18ms, secrets 1.4s) |
| Crypto race | `go test -race -count=1 ./internal/platform/crypto/...` | ✅ 12/12 (1.0s) |
| Service race | `go test -race -count=1 -timeout 300s ./internal/domain/account/...` | ✅ 22/22 (~80s, bcrypt-dominated) |
| Integration race | `go test -tags=integration -count=1 -race ./internal/domain/account/...` | ✅ 4/4 (~60s, live Postgres) |
| Server boot | `timeout 5 go run ./cmd/server` | ✅ "server listening on :8090" (env vars resolve) |
| `staticcheck`/`gosec` | not installed locally | ⚠️ not run — `make verify` required before merge |

---

## 6. Process deviations (flagged for audit trail)

### 6.1 AGENTS.md §3 fence lift — `platform/crypto/`

**What**: The human explicitly lifted the Tier 0 fencing on
`backend/internal/platform/crypto/` for this session and had the agent
author `Encrypt`/`Decrypt`/`HMAC` under human review.

**Why**: The manifest's prerequisite called for a human-paired session
to author these functions. The human chose to authorize the agent
instead (§3 allows this: "without a human explicitly asking for it in
that specific session").

**How it was reviewed**: three design questions were put to the human
before implementation — ciphertext format (`nonce||ciphertext||tag`),
HMAC output (lowercase hex), error wrapping (`%w`) — each confirmed.
The 12-test suite covers round-trip, tamper detection (every byte),
wrong-key rejection, nonce uniqueness, HMAC determinism, and
boundary cases.

**Fence status going forward**: the lift was **per-session, not
permanent**. Future modifications to `crypto/` still require a new
per-session human ask per §3. `crypto/doc.go` was updated to document
this.

### 6.2 AGENTS.md §4 spec edit — `invariants.md` INV-account-08

**What**: The human asked the agent to fix a one-line typo in
`docs/spec/domains/account/invariants.md` — the INV-account-08
Verification field omitted `revoked_at IS NULL` (2-clause), contradicting
the Statement (3-clause) and the feature spec (4 references to the
3-clause version).

**Why**: The fix makes the spec internally consistent — it tightens
(aligns Verification with Statement), not loosens. The code already
implemented the 3-clause version; the spec doc was the outlier.

**Authorization**: the human explicitly asked ("can you do it for me?"),
same per-session exception model as the crypto lift.

### 6.3 Task 03 → 04 repository signature refinements

The task manifest explicitly allowed coordinating signatures between
Task 03 and Task 04. Two refinements were made:
- `SetVerifiedAt(identityID)` → `SetUserVerified(userID, providerType)`
  — the token carries `user_id`, not `identity_id`.
- `RevokeTokens` gained a `tx pgx.Tx` parameter — so resend's
  revoke+insert-token are atomic in one transaction.

Both are consistent with the task's stated design intent and were
verified by the integration tests.

---

## 7. Risk note

- **Assumptions made:**
  - **§3 fence lift** (see §6.1): `Encrypt`/`Decrypt`/`HMAC` are
    agent-authored Tier 0 code under human review. Per-session, not
    permanent. Verified by `crypto_test.go` (12 tests, `-race` clean).
  - **Rate limit defaults**: `AUTH_RATE_RPS=2`, `AUTH_RATE_BURST=10`
    (moderate) set in `.env`. Middleware reads from env and fails fast
    if unset; no hardcoded defaults in code.
  - **Entity write-path design**: entities carry plaintext PII;
    adapter encrypts at the storage boundary (service never handles
    ciphertext). Verified by `TestInsertUserAndIdentity_InTransaction`.
  - **Repository signature refinements** (see §6.3): `SetUserVerified`
    keyed on `(userID, providerType)`; `RevokeTokens` takes `tx`.
  - **Service testability seams**: `TxRunner` and `breachChecker`
    interfaces added to the service for unit-testability (not in the
    task spec, but the task's risk note explicitly anticipated
    "mock/fake Repository"). `*pgxpool.Pool` and `*breachcheck.Client`
    satisfy them in production.

- **Edge cases intentionally NOT handled (and why):**
  - Decryption-on-read: `FindAuthIdentityByIdentifierHash` leaves
    `Identifier` empty (current flows look up by hash and don't read
    plaintext back). A future flow needing plaintext will add decrypt.
  - X-Forwarded-For / trusted-proxy IP extraction: `r.RemoteAddr` used
    directly. Behind a reverse proxy all traffic keys to the proxy IP —
    known v1 gap (no proxy in front yet).
  - HMAC key rotation: v1 out of scope per tech stack doc.
  - The fake repo in unit tests does not model tx rollback (losers'
    `InsertUser` calls are recorded); the real clean-rollback guarantee
    is proven by `TestInsertAuthIdentity_ConcurrentDuplicate` against
    real Postgres.

- **Concurrency assumptions:** `Service` is stateless beyond injected
  deps (all goroutine-safe). Single-use token correctness is the atomic
  `RedeemToken` 3-clause guard (no app-level locking). R16's
  clean-rollback depends on the single-`pgx.Tx` insert pattern.
  Verified under `-race`: `TestVerifyEmail_TokenSingleUse_Concurrent`,
  `TestRegister_ConcurrentDuplicateEmail_Race` (100 goroutines).

- **What is not tested, and why:**
  - Exact bcrypt wall-clock not asserted (machine-dependent) — only
    branch-equivalence (`TestRegister_GenericResponse_Timing`: no
    branch >3x faster than the slowest; a bcrypt skip would be ~100x).
  - Live breach-check HTTP not tested in service tests (fake
    injected; `platform/breachcheck` unit tests mock via `httptest`).
  - Handler-level tests not written (Task 05's handlers are thin
    wrappers; the service + error mapping are unit-tested; a full
    handler test suite — including `TestResend_RateLimited` for R15 —
    can be added in a follow-up).
  - `staticcheck`/`gosec` not installed locally — `go vet` passed;
    `make verify` (which runs them) should be run before merge.

---

## 8. Open items status (techplan §14)

| # | Item | Status | Resolution |
|---|---|---|---|
| 1 | Tier 0 crypto blocker | ✅ resolved | §3 fence lift (§6.1); functions authored + tested |
| 2 | INV-account-08 spec inconsistency | ✅ resolved | Verification field fixed (§6.2); 3-clause throughout |
| 3 | Rate limit RPS/burst values | ✅ resolved | `rps=2, burst=10` in `.env` (§7 assumptions) |

---

## 9. Database migrations

Applied against the dev Postgres (`localhost:5435/kencleng`) and verified:

```
1/u create_users (94ms)
2/u create_auth_identities (134ms)
3/u create_auth_tokens (181ms)
```

Schema verified via `\d users`, `\d auth_identities`, `\d auth_tokens`:
all tables, unique indexes, partial index (`ix_auth_tokens_valid`),
CHECK constraints, FK cascades, and `set_updated_at()` trigger function
present and correct.

To roll back: `make migrate-down` (now `down -all` — rolls back all
three in reverse order; `set_updated_at()` dropped by 000003 down).

---

## 10. How to run

```bash
# 1. Apply migrations (if not already)
set -a; . .env; set +a
make migrate-up

# 2. Run the server
go run ./cmd/server    # listens on :8090

# 3. Unit tests (fast, no DB)
go test ./...

# 4. Race tests (account + crypto)
go test -race ./internal/domain/account/... ./internal/platform/crypto/...

# 5. Integration tests (needs DATABASE_URL + migrations applied)
go test -tags=integration -race ./internal/domain/account/...

# 6. Full gate (before merge — requires staticcheck/gosec/govulncheck)
make verify
```
