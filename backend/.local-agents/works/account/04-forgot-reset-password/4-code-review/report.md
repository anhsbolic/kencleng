# Code Review Report — Forgot & Reset Password (account #04)

> Ticket      : account domain task #4 — `docs/spec/1-account/features/04-forgot-reset-password.md` (Fitur 2B)
> Reviewed    : 2026-08-26
> Inputs      : `2-plan/techplan.md` (Status: Approved by Anhar), `2-plan/tasks/` decomposition, `3-build/report.md`, working-tree diff (unstaged + untracked)
> Method      : Four-pass review (Safety → Quality → Stack-Specific Best Practices → Consistency) per `harscode-workspace/workflow/4-code-review/guidelines.md`
> Verdict     : **Approve with minor comments** (0 blocking, 3 optional, 1 process gate, 1 accepted residual)

---

## Scope of review

The diff under review is the forgot/reset-password vertical slice built
per the approved techplan:

- **New files**: `internal/transport/http/{auth_password_reset.go, auth_password_reset_test.go}`, `internal/domain/account/{password_reset_test.go, password_reset_integration_test.go}`, `internal/platform/notification/dev_sender_test.go`
- **Edited files**: `internal/domain/account/{repository.go, repository_db.go, repository_db_integration_test.go, service.go, service_test.go}`, `internal/platform/notification/{sender.go, dev_sender.go, sender_test.go}`, `cmd/server/main.go`, `api/openapi/{account.yaml}` + bundled `api/openapi.yaml`
- **Out of review scope** (per directory boundary, root AGENTS.md §7): the frontend diffs and the `05-account-linking` plan artifacts in the same working tree; the `api/openapi.yaml` contract edit (applied under human approval per AGENTS.md §4, Q2 resolved during explore)

The Tier 0 fenced sub-areas (`platform/auth/`, `platform/crypto/`,
`donation/`, `disbursement/`) are untouched — the slice reuses the
existing `RedeemToken` guard, `crypto.HMAC`, and `sha256Hex` as-is, and
introduces no new transaction/locking primitive that would fall under
the donation/disbursement fencing (root AGENTS.md §3).

---

## 1. Safety

**No findings.**

Nil-safety: every nullable return is guarded before dereference.
- `identity` from `FindAuthIdentityByIdentifierHash` checked `if identity != nil` before `issueResetToken` (`service.go` `ForgotPassword`).
- `googleIdentity` checked `if googleIdentity != nil` before `sendNudge` (same method).
- `t` from `FindAuthTokenByHash` checked `if t != nil && !t.ExpiresAt.After(time.Now())` in `ResetPassword`'s disambiguation arm; nil falls through to `ErrTokenNotFound`.
- `hashPassword` seam nil-checked with fallback to `secrets.HashPassword` (`ResetPassword`), mirroring the existing `compare` seam.

Concurrency: `Service` holds no new mutable state. Single-use correctness
rests entirely on the guarded `UPDATE … WHERE used_at IS NULL AND
revoked_at IS NULL AND expires_at > now()` inside `RedeemToken` — the
same atomic CAS predicate already battle-tested by task #3's refresh
rotation. The new `RevokeAllRefreshTokensForUser` and
`UpdateIdentityCredentialSecret` run inside the caller's tx (no separate
step after commit), so a mid-transaction failure rolls both back with
the redeem (INV-account-05). The fake-repo hooks and `captureSender`
use `sync.Mutex`.

Errors: propagated with `%w` throughout; the bare `return
ErrTokenNotFound` on purpose-mismatch matches the `VerifyEmail`
precedent and triggers the deferred rollback. Sender errors are
swallowed by design (post-commit, non-fatal) and logged via
`notificationErrorCategory` — never `%v` verbatim on a sender error that
could embed the recipient/token (per
`go/secrets-and-sensitive-logging.md` §1).

Context: `ctx` threads through all repo/sender calls; handlers use
`r.Context()`. No external call inside an open tx (bcrypt and
`validatePassword`/HIBP both run before `BeginTx`;
`SendPasswordResetEmail` is post-commit).

Resources: both new tx helpers (`issueResetToken`, `ResetPassword`) use
the established `committed`-flag + deferred-rollback pattern; every
early-return path (including the purpose-mismatch returns) triggers the
rollback.

---

## 2. Quality

### Q1 — Dead `expectSuccess` parameter in stress-test `racer` closure (OPTIONAL)

- **Location**: `internal/domain/account/password_reset_integration_test.go:296-309` (`racer` closure in `TestResetPassword_Stress_MixedValidAndReplayed_RealDB`)
- **What's wrong**: The `racer` closure takes `expectSuccess bool` but the body only does `_ = expectSuccess` — the parameter is never read for any assertion. The comment says "classification happens via the success counter," which is the right design, but then the parameter is dead weight. Callers pass `true`/`false` (`:314-315`) that influence nothing.
- **Why it matters**: This is the exact "unused parameters left after refactoring" pattern called out in `workflow/4-code-review/examples.md`. Static analysis catches unused variables, not unused parameters — this one survives `staticcheck` because the `_ =` discard looks intentional. A future reader will assume the parameter drives an assertion and waste time tracing why it doesn't.
- **Suggested fix**: Remove the parameter — `racer := func(token string) { … }`, callers become `go racer(validPlain)` / `go racer("garbage-"+uuid.NewString())`. The success/violation counters already classify outcomes.

### Q2 — `sendPasswordReset` near-duplicates `sendVerification` (OPTIONAL, pre-existing pattern)

- **Location**: `internal/domain/account/service.go:709-716` (`sendPasswordReset`) vs `:695-704` (`sendVerification`)
- **What's wrong**: The two methods are structurally identical — call sender, on error log a sanitized category via `notificationErrorCategory`. Only the log prefix and the sender method differ. `sendNudge` (`:720-727`) is a third near-clone with an extra `nudgeType` field.
- **Why it matters**: Three copies of the same swallow-and-sanitize pattern. If the sanitization logic ever changes (e.g. adding a structured-logger migration), all three must be found and updated independently.
- **Suggested fix**: Not a blocker — this follows the established codebase pattern, and extracting a shared helper (`sendEmail(ctx, sendFn, logPrefix)`) would touch existing merged code outside this slice's scope. Flag for a future cleanup pass, not this PR.

---

## 3. Stack-Specific Best Practices

Matched trigger keywords in `best-practices/index.md` against the diff's
technology (Go: token, error, secrets, logging, nil, rate-limit;
PostgreSQL: transaction, lock; REST API: anti-enumeration, status
codes). Opened: `restapi/anti-enumeration.md`,
`postgresql/transactions-and-locking.md`,
`go/secrets-and-sensitive-logging.md`, `go/jwt-and-token-lifecycle.md`.

### BP1 — Anti-enumeration SMTP timing gap on the no-match branch (OPTIONAL, pre-existing residual)

- **Source**: `restapi/anti-enumeration.md` checklist item: "Response timing on sensitive endpoints doesn't differ significantly between 'found' and 'not found' (avoid an early return that skips an expensive operation on only one branch)."
- **Location**: `internal/domain/account/service.go` `ForgotPassword` no-match branch; tested via `TestForgotPassword_Timing_Branches_RealPostgres` (`password_reset_integration_test.go:415`)
- **What's wrong**: The three forgot branches are shaped for DB-time parity via `dummyWrite`, but the no-match branch skips the email send entirely. In production with a synchronous SMTP sender, the no-match branch would be measurably faster than the match (reset email) and google-only (nudge email) branches. The timing test uses `integrationSilentSender` (instant `return nil`), so it proves DB-time equivalence only — not total wall-clock parity with a real sender.
- **Why it matters**: A remote attacker measuring response latency could distinguish "email not registered" (fast) from "registered" (slow) once real SMTP is wired — an account-existence oracle. The `dummyWrite` device only evens out the DB layer, not the notification layer.
- **Suggested fix**: This is a pre-existing pattern — the register and resend flows have the identical gap and were already accepted in their code reviews. Not a regression introduced by this diff. A codebase-wide fix would be either async/fire-and-forget email send (so the handler returns before SMTP) or a dummy-send on no-op branches. Flag for awareness; do not block this slice on it.

All other checklist items from the four opened files pass cleanly:

- `restapi/anti-enumeration.md`: identical generic 202 regardless of branch ✓; no application-level secret string comparison (token looked up by SHA-256 hash via DB index, not `==`) ✓; error messages don't leak internal details ✓.
- `postgresql/transactions-and-locking.md`: no network call to an external system inside an open DB transaction (bcrypt + HIBP before `BeginTx`, email send post-commit) ✓; no multi-row `FOR UPDATE` ordering concern (single guarded `UPDATE` per tx) ✓; `FindAuthTokenByHash` disambiguation read runs outside the rolled-back tx on a separate connection, no non-repeatable-read pitfall ✓.
- `go/secrets-and-sensitive-logging.md`: sender errors logged via `notificationErrorCategory` (returns "timeout" or "send failed", never `err.Error()`) ✓; handler `%v` only on service-wrapped DB errors — parameterized SQL via goqu `Prepared(true)` carries no PII values, consistent with `ResendVerificationHandler` ✓; no struct logged wholesale ✓.
- `go/jwt-and-token-lifecycle.md`: token purpose separation enforced via cross-purpose guard in both directions (Q1 fix: `VerifyEmail` rejects `password_reset` tokens; `ResetPassword` rejects `email_verification` tokens) ✓; `resetTokenTTL = time.Hour` is risk-proportional (shorter than `tokenTTL = 24h` for the lower-risk email-verification token) ✓; refresh-token mass revoke on reset is the credential-compromise session-invalidation response ✓.

---

## 4. Consistency

### C1 — Inline Problem write deviates from `MapServiceError` pattern (OPTIONAL)

- **Convention violated**: Empty-token boundary rejection should go through `MapServiceError(w, account.ErrTokenNotFound)`, as established in `internal/transport/http/auth_verify_email.go:32-34` (`VerifyEmailHandler`).
- **Location**: `internal/transport/http/auth_password_reset.go:81-86` (`ResetPasswordHandler` empty-token early return); documented as Deviation #2 in `3-build/report.md`.
- **What's wrong**: `ResetPasswordHandler` writes the 404 Problem inline with hardcoded strings:
  ```go
  WriteProblem(w, http.StatusNotFound,
      "https://kencleng.dev/problems/token-not-found",
      "Token Not Found", "The verification token was not found.")
  ```
  while the sibling `VerifyEmailHandler` uses `MapServiceError(w, account.ErrTokenNotFound)`, which maps to the identical strings in `errors.go:86-89`. The inline copy duplicates the type URI / title / detail — if the canonical strings in `MapServiceError` ever change, this copy won't track.
- **Why it matters**: Two sources of truth for the same Problem response. Deviation #2 cites "interface purity (no `account` import in the handler file)," but the handler already calls `MapServiceError` (`errors.go`, same `http` package), which imports `account` — the transitive dependency exists regardless of whether the handler file has a direct import. The purity argument doesn't hold against the maintenance cost of duplicated strings.
- **Suggested fix**: Either (a) import `account` and call `MapServiceError(w, account.ErrTokenNotFound)` — matches `VerifyEmailHandler` exactly; or (b) extract a `writeTokenNotFound(w)` helper in `errors.go` for both handlers to share, keeping the handler file free of a direct `account` import. Either is a one-line change.

All other conventions match the target repo (`backend/AGENTS.md` + root
`AGENTS.md`):

- Error handling: `%w` wrapping throughout ✓ (§2); bare `return ErrTokenNotFound` for purpose-mismatch matches `VerifyEmail` ✓.
- Logging: `log.Printf("account: … (recipient redacted): %s", notificationErrorCategory(err))` for sender errors ✓; `log.Printf("transport: … (recipient redacted): %v", err)` for service errors — matches `ResendVerificationHandler` verbatim ✓; user_id logged, never email/token ✓ (root golden rule).
- Validation: `looksLikeEmail` at handler boundary ✓; `validatePassword` (length ≥8 + HIBP fail-open) in service ✓; empty-token 404 at boundary ✓.
- Naming: `resetTokenTTL` mirrors `tokenTTL`; `purposePasswordReset` mirrors `purposeEmailVerify`; `UpdateIdentityCredentialSecret` mirrors `SetUserVerified`; `RevokeAllRefreshTokensForUser` mirrors `RevokeRefreshTokenFamily`; `SendPasswordResetEmail` mirrors `SendVerificationEmail`; `auth_password_reset.go` mirrors `auth_verify_email.go`; `forgotPasswordService`/`resetPasswordService` narrow ports mirror `loginSessionService` (`auth_login.go:17`) ✓.
- Doc comments on every exported function/type ✓ (§2).
- Table-driven tests (`[]struct{ name string; … }`) for multi-case tests ✓ (§2).
- `//go:build integration` tag on `password_reset_integration_test.go`; unit tests beside source ✓ (§3).

---

## Verdict

**Approve with minor comments.**

No blocking findings. The slice is safety-clean, follows all relevant
best-practice checklists, and is consistent with the codebase's own
conventions except for the one inline-Problem deviation (already
documented as Deviation #2).

Optional follow-ups (none gate merge):
1. **C1** — Replace the inline `WriteProblem` in `ResetPasswordHandler` with `MapServiceError(w, account.ErrTokenNotFound)` (or a shared `writeTokenNotFound` helper) to avoid string duplication. (Consistency)
2. **Q1** — Remove the dead `expectSuccess` parameter from the stress-test `racer` closure. (Quality)
3. **BP1** — The SMTP-timing anti-enumeration gap is a pre-existing codebase-wide residual (same as register/resend); flag for a future async-send pass, not this slice. (Best Practices, accepted residual)

### Process gate (not a code-review finding)

The two deferred heavy-race tests —
`TestResetPassword_TokenSingleUse_Concurrent_RealDB` and
`TestResetPassword_Stress_MixedValidAndReplayed_RealDB` — are gated
behind `KENCLENG_HEAVY_RACE_TESTS=1` and were skipped this session.
The build report and techplan Open Item #5 both acknowledge this: the
Tier 1 KPI for INV-account-08's real-DB stress proof is **not met until
these run clean**. They must be re-run before merge / `make verify`.
This is a process gate, not a code defect — the single-use property is
pinned at unit level (fake CAS mode + wrong-purpose tests), and the
guarded `UPDATE` predicate is byte-for-byte the one already
stress-proven by task #3's `TestRefresh_Stress_MixedValidAndReplayed`.
