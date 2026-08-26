# Build Report — account/04-forgot-reset-password

> Date: 2026-08-26
> Techplan: `../2-plan/techplan.md` (Draft)
> Task files executed: `../2-plan/tasks/task-01..06`

## What was built

Per-task summary (full detail lives in the task files; only deltas and
coverage claims here):

### Task 01 — repository foundation
- `Repository.UpdateIdentityCredentialSecret(ctx, tx, userID, providerType, passwordHash)` and
  `Repository.RevokeAllRefreshTokensForUser(ctx, tx, userID)` added
  (`repository.go`) with goqu implementations in `repository_db.go`,
  keyed/shaped exactly per D1. Guard semantics proven against real
  Postgres inside the integration suite (see task 06).

### Task 02 — notification platform
- `Sender.SendPasswordResetEmail(ctx, to, token)` added; `FakeSender`
  logs category-only; `DevSender` appends a third `[password_reset]`
  outbox line type (0600 preserved). Unit tests: Fake no-PII-in-log,
  Dev outbox line + mode assertion (`dev_sender_test.go` is new — the
  file previously had no test coverage at all).
- Compile ripple fixed through all `Sender`/`Repository` fakes in the
  account test suite.

### Task 03 — service core (`service.go`)
- Consts `resetTokenTTL = time.Hour`, `purposePasswordReset`.
- `issueResetToken`: insert-only tx helper (no revoke — Assumption A).
- `ForgotPassword`: 3-branch dispatch, dummyWrite timing shaping,
  post-commit send.
- `ResetPassword`: validate → hash → single tx (redeem → purpose check →
  credential update → mass session revoke → commit) → expired-vs-other
  disambiguation after rollback.
- **Q1 fix**: `VerifyEmail` now checks the redeemed purpose and rejects
  non-`email_verification` tokens via rollback (`ErrTokenNotFound`).

### Task 04 — transport
- New narrow service ports `forgotPasswordService` / `resetPasswordService`
  (see Deviation #1); handlers `ForgotPasswordHandler` /
  `ResetPasswordHandler` in new `auth_password_reset.go`; two route lines
  in `cmd/server/main.go` inheriting the mount-time rate limiter.

### Task 05 — contract completion
- `"429": $ref TooManyRequests` added to `/auth/reset-password` in
  `api/openapi/account.yaml`; bundle regenerated via redocly; diff
  verified to contain ONLY that addition. Pre-existing validation
  failure (security-defined + `$ref`-sibling warnings) confirmed present
  on HEAD before my change — untouched.

### Task 06 — integration & race suite (`password_reset_integration_test.go`)
| Test | Proves | Result |
|---|---|---|
| `TestResetPassword_AllSessionsRevoked_Atomic_RealDB` | INV-account-05: sessions across families incl. rotated rows revoked atomically | PASS |
| `TestResetPassword_FailureBetweenWrites_RollsBackBoth_RealDB` | R18 rollback arm: injected revoke failure rolls back redeem+credential; token stays usable | PASS |
| `TestResetPassword_ExpiredToken_NoStateChange_RealDB` | R9 zero-mutation | PASS |
| `TestResetPassword_TokenSingleUse_Concurrent_RealDB` | INV-account-08 race (100 goroutines) | **SKIPPED** (see Deferral) |
| `TestResetPassword_Stress_MixedValidAndReplayed_RealDB` | ≥100-goroutine stress mix | **SKIPPED** (see Deferral) |
| `TestForgotPassword_Timing_Branches_RealPostgres` | R5 anti-enumeration timing band vs real DB | PASS |

Unit suite (`password_reset_test.go`): R1–R4, R8, R10, R12, R13,
R14, breach fail-open — all passing under `-race`. Handler tests
(`auth_password_reset_test.go`): generic-202 byte equality, error
swallowing w/o PII leak, 400/422 mapping, empty-token 404 pre-service,
rate-limit 429 on both endpoints, status passthrough — all passing.

## Gate results

| Gate | Result |
|---|---|
| `go build ./...` / `go vet ./...` (+integration tag) | clean |
| `go test ./...` | ok |
| `go test -race ./...` | ok |
| `go test -tags=integration -race ./internal/domain/account/...` | ok (2 env-gated skips) |
| `go test -tags=contract ./...` | ok |
| staticcheck | clean |
| gosec | 13 findings — verified identical on HEAD (all pre-existing; none in this slice's files) |
| govulncheck | 24 affected — verified identical on HEAD (pre-existing) |
| gitleaks | NOT INSTALLED on this machine — could not run |
| `make verify` full chain | not run as one command (frontend half out of scope; backend stages run individually above) |

## Rule coverage map (techplan §4 → tests)

| Rule | Test(s) |
|---|---|
| R1 | `TestForgotPassword_Match_IssuesTokenAndSendsEmail` |
| R2 | `TestForgotPassword_GoogleOnly_NoticeNoToken` |
| R3 | `TestForgotPassword_NoMatch_NothingSent` |
| R4 | `TestForgotPassword_Repeat_DoesNotRevokePriorTokens` |
| R5 | `TestForgotPassword_GenericResponse_AllBranches` + handler `Generic202_AllBranches` + `TestForgotPassword_Timing_Branches_RealPostgres` |
| R6 | `TestForgotPassword_RateLimited`, `TestResetPassword_RateLimited` |
| R7 | `TestResetPassword_HappyPath_UpdatesAndRevokes` + `TestResetPassword_AllSessionsRevoked_Atomic_RealDB` + handler success test |
| R8 | `TestResetPassword_PasswordPolicy_TokenNotConsumed` |
| R9 | `TestResetPassword_TokenStateMapping/expired` + `_ExpiredToken_NoStateChange_RealDB` |
| R10 | `TestResetPassword_TokenStateMapping/{not-found,already-used}` |
| R11 | unit fake-based concurrent precedent + **real-DB version DEFERRED** |
| R12 | `TestVerifyEmail_RejectsResetPurposeToken` |
| R13 | `TestResetPassword_WrongPurpose_RejectedNoMutation` |
| R14 | `TestPasswordReset_SendFails_LogsNoPIIOrToken` + handler internal-error test |
| R15 | `TestResetPassword_EmptyToken_404_NoServiceCall` |
| R16 | malformed-JSON handler tests (both endpoints) |
| R17 | `TestForgotPassword_Handler_MalformedBodyAndEmail` |
| R18 | `TestResetPassword_FailureBetweenWrites_RollsBackBoth_RealDB` |

## Deviations from techplan (flagged, none silent)

1. **Handlers take narrow interfaces** (`forgotPasswordService`,
   `resetPasswordService`) instead of `*account.Service`. The techplan's
   §8 sketch implied concrete-type handlers like Register/VerifyEmail;
   the newest precedent in-tree is `LoginHandler(loginSessionService)`,
   which I followed — it also lets handler tests stub the service instead
   of fighting a nil pgxpool. Compile-time assertions pin
   `*account.Service` to both ports.
2. **Empty-token early-reject writes the Problem inline** rather than via
   `MapServiceError(account.ErrTokenNotFound)` — same type URI/title/
   detail strings, kept for interface purity (no `account` import in the
   handler file).
3. **`hashPassword` seam added to Service** (nil-default
   `secrets.HashPassword`), mirroring the existing `compare` seam —
   introduced mid-build so concurrency tests don't pay ~1s/hash × 220
   racers under `-race`. Production behavior unchanged (NewService wires
   the real bcrypt; ResetPassword also nil-falls-back defensively).
4. **Breach fail-open unit test** follows the register precedent —
   simulates fail-open as `(false, nil)` (the real fail-open unit lives
   in platform/breachcheck), not by injecting an error into the checker.

## Risk note (AGENTS.md §5 structure)

- **Assumptions made**: forgot-password's generic message text matches
  the openapi example ("Kalau email terdaftar, instruksi sudah dikirim.");
  reset success message uses the contract example verbatim; timing-band
  for forgot set at ≤3× (vs register's ≤2×) because branches differ by one
  identity lookup — measured worst ratio was well inside it.
- **Edge cases intentionally NOT handled (and why)**:
  - Access-token revocation on reset — INV-account-05 scopes refresh
    tokens; access dies on its own ≤15-min TTL (techplan §7).
  - Rate-limit values untouched (`AUTH_RATE_*` env-driven; read, not tuned).
  - Problem-URI prefix split left as-is (repo-wide decision pending).
- **Concurrency assumptions**: correctness rests entirely on the guarded
  `UPDATE ... WHERE used_at IS NULL AND revoked_at IS NULL AND expires_at
  > now()` inside one tx — proven by `AllSessionsRevoked_Atomic` +
  `FailureBetweenWrites_RollsBackBoth` against real Postgres, and by the
  same predicate's battle-testing in task #3's refresh rotation.
  **NOT proven this session**: see deferral below.
- **What is not tested, and why**:
  - `TestResetPassword_TokenSingleUse_Concurrent_RealDB` and
    `TestResetPassword_Stress_MixedValidAndReplayed_RealDB` are gated
    behind `KENCLENG_HEAVY_RACE_TESTS=1` — deferred per Anhar's call
    during this session after repeated multi-minute runs (bcrypt cost ×
    `-race`). Recorded as techplan Open Item #5. **The Tier 1 KPI for
    INV-account-08's real-DB stress proof is therefore unmet until these
    two run clean; must happen before merge.** The single-use property
    itself IS pinned at unit level (fake CAS mode + wrong-purpose tests),
    and the guarded UPDATE is byte-for-byte the predicate already
    stress-proven by task #3's `TestRefresh_Stress_MixedValidAndReplayed`.
  - gitleaks stage could not run (tool absent).

## Feature spec fulfillment

Fulfills `docs/spec/1-account/features/04-forgot-reset-password.md`
(Fitur 2B). Spec file untouched (§4 authority separation); the one
contract edit (429 documentation) was human-approved during explore.
