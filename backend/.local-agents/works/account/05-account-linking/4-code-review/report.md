# Code Review Report — Account Linking (account #05)

> Ticket      : account domain task #5 — `docs/spec/1-account/features/05-account-linking.md`
> Reviewed    : 2026-08-27
> Inputs      : `2-plan/techplan.md` (Status: Approved by Anhar), `2-plan/tasks/` decomposition, `3-build/report.md`, working-tree diff (unstaged + untracked)
> Method      : Four-pass review (Safety → Quality → Stack-Specific Best Practices → Consistency) per `harscode-workspace/workflow/4-code-review/guidelines.md`
> Verdict     : **Request changes** (2 blocking, 3 optional)

---

## Scope of review

The diff under review is the account-linking vertical slice built per
the approved techplan:

- **New files**: `internal/domain/account/{security.go, security_test.go, security_integration_test.go}`, `internal/transport/http/{account_security.go, account_security_test.go}`, `migrations/000010_widen_auth_tokens_purpose.{up,down}.sql`
- **Edited files**: `internal/domain/account/{repository.go, repository_db.go, service.go, service_test.go}`, `internal/platform/notification/sender.go`, `cmd/server/main.go`
- **Out of review scope** (per directory boundary, root AGENTS.md §7): the frontend diffs and the `04-forgot-reset-password` uncommitted changes co-existing in the same working tree (treated as legitimate per the build report's "Fitur 04 coordination" note).

The Tier 0 fenced sub-areas (`platform/auth/`, `platform/crypto/`,
`donation/`, `disbursement/`) are untouched — the slice reuses
`GoogleTokenVerifier`, `crypto.HMAC`, and the existing `committed`-flag
deferred-rollback pattern; it introduces no new transaction/locking
primitive under donation/disbursement fencing (root AGENTS.md §3).

---

## 1. Safety

### S1 — Wrong error category: internal failures returned as 401 invalid-credentials (BLOCKING)

- **Location**: `internal/transport/http/account_security.go:116-129` (`SetPasswordHandler`) and `:179-196` (`UnlinkGoogleHandler`)
- **What's wrong**: `SetPassword`/`UnlinkGoogle` return only three sentinel kinds (`ErrValidation`, `ErrInvalidCredentials`, and the unlink-only `ErrOnlyIdentity`/`ErrRemainingUnverified`). **Every other** error path is a wrapped `fmt.Errorf("account: …: %w", err)` — DB lookup failure, `BeginTx`/`Commit` failure, `InsertAuthIdentity` non-unique-non-23505, `RevokeAllRefreshTokensForUser`, `DeleteAuthIdentitiesByIDs`, `InsertUserLog`, and the "missing credential_secret" data-corruption branch. Both handlers funnel all of those into `WriteProblem(..., 401, problemTypeInvalidCredentials, …, "Email atau password salah.")`. A DB outage or a commit failure is surfaced to the client as "wrong password."
- **Why it matters**: This is the "wrong error category (e.g. internal error vs invalid request)" recurring pattern from `workflow/4-code-review/examples.md` — both outcomes "fail," so it is easy not to notice which category fired. Concrete harm: (a) the client misclassifies a server-side outage as a credentials problem and retries the password; (b) real outages get masked behind 401s in monitoring dashboards keyed on status code; (c) on `SetPassword` Branch 1, a 401 is distinguishable from the generic 202 — an anti-enumeration leak (the existence of an internal failure on the "created" branch breaks the generic-response guarantee for that one request). No test exercises an internal-error path today (the `stubSecurityService` only injects sentinels), which is why this slipped through.
- **Suggested fix**: In `SetPasswordHandler`, after the `isErrValidation` branch, add `if errors.Is(err, account.ErrInvalidCredentials) { → 401 }` and route everything else through `MapServiceError(w, err)` (whose `default` is 500 + server-side `log.Printf`). In `UnlinkGoogleHandler`, keep the two `ErrOnlyIdentity`/`ErrRemainingUnverified` cases, add an explicit `case errors.Is(err, account.ErrInvalidCredentials): → 401`, and make `default: MapServiceError(w, err)`. Add a transport test that injects a non-sentinel wrapped error (e.g. `fmt.Errorf("boom: %w", errors.New("db down"))`) through the stub and asserts `503`/`500` + the generic `internal` problem type — proves the dispatch, not just that *an* error fired.

### S2 — Resource/tx correctness on the failure paths: clean

Nil-safety: every nullable dereference is guarded.
- `epIdentity == nil || epIdentity.CredentialSecret == nil` checked before `*epIdentity.CredentialSecret` in `setPasswordBranch2` (`security.go:191`).
- `verifiedEmailPassword.CredentialSecret == nil` checked before `*verifiedEmailPassword.CredentialSecret` in `UnlinkGoogle` (`security.go:311`).
- `scanAuthIdentities` uses `sql.NullString`/`sql.NullTime` for the nullable `credential_secret`/`verified_at` columns (`repository_db.go`).

Concurrency: `Service` holds no new mutable state. `UnlinkGoogle`'s `FOR UPDATE` is a single acquisition point keyed on `user_id` (no multi-row ordering hazard); the `committed`-flag + deferred-`Rollback` pattern is used in all three tx-owning methods (`setPasswordBranch1`, `setPasswordBranch2`, `UnlinkGoogle`), and every early-return path triggers the rollback. `sendVerification`/`sendNudge` run **post-commit** in Branch 1 — no external call inside an open tx (mirrors `Register`).

Context: `ctx` threads through every repo/sender call; handlers use `r.Context()`. No `context.Background()` mid-chain.

Resources: `rows.Close()` is deferred in `scanAuthIdentities`; the FOR UPDATE rows, the INSERT/UPDATE/DELETE execs, and the audit `InsertUserLog` all run on the caller's tx and roll back together on failure.

---

## 2. Quality

### Q1 — Ambiguous `(bool, error)` return from `SetPassword` (OPTIONAL)

- **Location**: `internal/domain/account/security.go:51`
- **What's wrong**: The bool means "Branch 2 ran → handler writes 200," but a caller reading `changed, err := svc.SetPassword(...)` cannot tell that from "password changed" or "success" without the doc comment. The build report's "techplan deviation #1" justifies the need to distinguish 200 from 202, but the signal is not self-documenting from the signature.
- **Why it matters**: A future caller (or a future reader of the handler) can easily misread `changed=true` as a generic success flag and miss the Branch-2-only semantics. This is the "unclear function signatures" item from `guidelines.md` §2.
- **Suggested fix**: Either a one-field result struct (`type SetPasswordOutcome struct{ Branch2Ran bool }`) or a named return, both self-documenting. The doc comment already covers it; non-blocking.

### Q2 — Duplicated boundary-validation literals (OPTIONAL)

- **Location**: `internal/transport/http/account_security.go:107-108` vs `auth_register.go:40-41`
- **What's wrong**: The `len(req.Password) < 8` check and the `"password is not allowed"` validation message are copy-duplicated from the register handler. If the password policy changes (length, or the breach-list rejection message), both must be found and updated independently.
- **Why it matters**: No tool flags "this is basically the same as that other function" (`workflow/4-code-review/examples.md`). The new code is consistent with the existing register handler — just not DRY.
- **Suggested fix**: Extract a shared `validatePasswordFields(req)` helper (or a `minPasswordLength` named constant — see `go/error-handling.md` §2) in `errors.go` for both handlers to call. Non-blocking; touches merged register code, so flag for a cleanup pass rather than this PR if scope discipline (root AGENTS.md §7) bites.

### Q3 — `purposeEmailVerifyLink` widening comment slightly loose on the down migration (OPTIONAL)

- **Location**: `migrations/000010_widen_auth_tokens_purpose.down.sql:4-7`
- **What's wrong**: The comment claims "token redemption is purpose-blind (RedeemToken guards on the hash's validity, not the purpose value)." Post-this-feature, redemption is NOT purpose-blind: `VerifyEmail`'s guard checks `purpose != purposeEmailVerify && purpose != purposeEmailVerifyLink`, and the R14 audit is conditional on `purpose == purposeEmailVerifyLink`. After the down remaps link→email_verification, a not-yet-redeemed link token would redeem as a registration token and skip the audit.
- **Why it matters**: The safety claim is technically inaccurate. The outcome is still safe **if code + migration roll back together** (rolled-back code has no audit logic and the 2-value guard accepts the remapped token), but a reader who trusts the comment literally could roll back only the migration while keeping the new code and observe a silent behavioral change.
- **Suggested fix**: Tighten the comment to state the actual safety contract: "safe when the migration rolls back alongside the feature code; a standalone rollback leaves link-purpose tokens redeeming as registration tokens (identity still verifies, R14 audit silently skipped)." Non-blocking.

---

## 3. Stack-Specific Best Practices

Matched trigger keywords in `best-practices/index.md` against the
diff's technology (Go: token, error, secrets, logging, nil, jwt;
PostgreSQL: transaction, lock, migration; REST API: anti-enumeration,
status codes). Cross-checked the Security Concern Map (authn, authz,
secrets-and-keys, input-validation). Opened:
`restapi/anti-enumeration.md`,
`postgresql/transactions-and-locking.md`,
`postgresql/migrations-safety.md`,
`go/jwt-and-token-lifecycle.md`,
`go/authorization-and-idor.md`,
`go/error-handling.md`,
`go/secrets-and-sensitive-logging.md`.

### BP1 — Anti-enumeration: 401-on-internal leak on Branch 1 (BLOCKING, same root as S1)

- **Source**: `restapi/anti-enumeration.md` — "Response timing on sensitive endpoints doesn't differ significantly between 'found' and 'not found' (avoid an early return that skips an expensive operation on only one branch)" and "Error messages don't leak internal details that aid enumeration."
- **Location**: `internal/transport/http/account_security.go:116-129` (`SetPasswordHandler` default branch)
- **What's wrong**: Branch 1 is anti-enumeration-shaped at the service layer (generic `(false, nil)` across created/claimed/race-loser, `dummyWrite` for DB-time parity, bcrypt always runs per R5). But the handler's default→401 means an internal failure on the "created" sub-branch returns 401 instead of 202 — distinguishable from the normal 202. Fixing S1 (route non-sentinel errors through `MapServiceError` → 500) doesn't fully close the enumeration gap (a 500 is also distinguishable from 202), but it at least returns the correct category; the established codebase pattern (`RegisterHandler` via `MapServiceError`) accepts this residual on register too, so consistency holds. The 401 category is strictly worse because it is both wrong-category AND a credentials-shaped response on an endpoint whose only legitimate credential error is Branch 2's wrong-current-password.
- **Suggested fix**: same as S1.

### BP2 — Down-migration "purpose-blind" comment (OPTIONAL, same as Q3)

- **Source**: `postgresql/migrations-safety.md` — "The `down` migration is genuinely reversible and has been tested, not just written symmetrically on paper."
- **Location**: `migrations/000010_widen_auth_tokens_purpose.down.sql:4-7`
- **What's wrong**: See Q3. The down is structurally reversible (re-map then restore CHECK), but the comment's safety rationale is overstated. Reversibility holds only under code+migration co-rollback.
- **Suggested fix**: tighten the comment per Q3.

All other checklist items from the opened files pass cleanly:

- `restapi/anti-enumeration.md`: identical generic 202 across Branch-1 outcomes ✓; no application-level secret string comparison (token looked up by SHA-256 hash via DB index; password via bcrypt) ✓; submitted email never echoed in any body or validation error ✓.
- `postgresql/transactions-and-locking.md`: no network call to an external system inside an open DB transaction (bcrypt + `validatePassword` before `BeginTx`; `sendVerification`/`sendNudge` post-commit) ✓; `FOR UPDATE` is a single acquisition point keyed on `user_id` (no multi-row ordering hazard, no deadlock surface) ✓; READ COMMITTED "loser classifies post-commit state" reasoning in `UnlinkGoogle` is correct and explicitly documented ✓.
- `postgresql/migrations-safety.md`: additive CHECK widening on a small table, no new columns, no NOT NULL introduced ✓ (one comment nuance — BP2).
- `go/jwt-and-token-lifecycle.md`: no new signing key introduced — `RequireSession` reuses the existing ES256 access-token verifier (same token purpose, same key, correct) ✓; `purposeEmailVerifyLink` is a distinct purpose in the single-use `auth_tokens` table, not a new JWT type ✓; refresh-token mass revoke on Branch 2 is the credential-compromise session-invalidation response ✓.
- `go/authorization-and-idor.md`: `userID` comes from `RequireSession`'s context, never from URL/body — no IDOR surface and no resource-ID scope check needed ✓; no dedicated IDOR test required (no resource ID from input).
- `go/error-handling.md`: no `return a, b` staleness — all `return false, <call>` use the constant `false`, never a stale computed value ✓; `purposeEmailVerifyLink` is a named const (§2 satisfied) ✓; the `< 8` literal is the named-const opportunity noted in Q2.
- `go/secrets-and-sensitive-logging.md`: all new log lines are `log.Printf("account: … user_id=%s", userID)` — user_id only, never email/password/token ✓ (matches `service.go:356`/`473`/`680`); handler `%v` only on service-wrapped DB errors — parameterized SQL via goqu `Prepared(true)` carries no PII values ✓; no struct logged wholesale ✓.

---

## 4. Consistency

### C1 — Error-dispatch convention violated: handlers bypass `MapServiceError` (BLOCKING, same root as S1/BP1)

- **Convention violated**: The repo's error-mapping convention is embodied by `MapServiceError` (`internal/transport/http/errors.go:76-109`), whose `default` is 500 + `log.Printf("transport: unhandled service error: %v", err)`. The established usage precedent is `RegisterHandler` (`auth_register.go:60`: `MapServiceError(w, err)`) and `ResetPasswordHandler` (`auth_password_reset.go:89`: `MapServiceError(w, err) // 422 / 410 / 404 per sentinel mapping`). The techplan's own task-05 decomposition (`2-plan/tasks/task-05-transport-wiring.md`, Files table) explicitly specified "+ two 409 cases in `MapServiceError`" for `errors.go` — the build did NOT do this; it inlined the 409s in the handler and routed everything else to 401.
- **Location**: `internal/transport/http/account_security.go:116-129` (`SetPasswordHandler`) and `:179-196` (`UnlinkGoogleHandler`)
- **What's wrong**: Two sources of truth for the credentials-error response (the inline copy in `account_security.go` vs the canonical mapping in `errors.go:98-101`), AND the inline `default → 401` misclassifies every internal error. This is the same defect as S1/BP1, cited here against the convention source and the techplan's own task contract.
- **Suggested fix**: Either (a) add the two new 409 sentinels (`ErrOnlyIdentity`, `ErrRemainingUnverified`) to `MapServiceError` per the original task-05 plan and call `MapServiceError(w, err)` in both handlers (matches `RegisterHandler`/`ResetPasswordHandler` exactly, removes the inline 409 strings too), or (b) keep the inline 409s but add an explicit `case errors.Is(err, account.ErrInvalidCredentials): → 401` and make `default: MapServiceError(w, err)` (→ 500). Option (a) is preferred — it is what the techplan specified and collapses C1 + S1 + BP1 into a single fix.

### C2 — Other conventions: clean

All other conventions match the target repo (`backend/AGENTS.md` +
root `AGENTS.md`):

- Error handling: `%w` wrapping throughout ✓ (§2); bare `return ErrOnlyIdentity`/`ErrRemainingUnverified`/`ErrInvalidCredentials` for guard outcomes matches the `ErrTokenNotFound` precedent ✓.
- Logging: `log.Printf("account: … user_id=%s", userID)` for service decision points ✓; user_id logged, never email/password/token ✓ (root golden rule).
- Validation: `len(req.Password) < 8` at handler boundary ✓; `validatePassword` (length + breach-list fail-open) in service ✓; `len(req.Password) < 1` required-check on unlink ✓.
- Naming: `purposeEmailVerifyLink` mirrors `purposeEmailVerify`/`purposePasswordReset`; `actionAccountLinking` reused from `google_oauth.go`; `FindAuthIdentitiesByUser`/`FindAuthIdentitiesByUserForUpdate`/`DeleteAuthIdentitiesByIDs` mirror the existing `Find*`/`Insert*` vocabulary; `NudgeSetPasswordConflict` mirrors `NudgePasswordReset`/`NudgeGoogleOnly`; `account_security.go` mirrors `auth_*.go` ✓.
- Audit-entry construction (`UserLog{ID,UserID,ActionType:actionAccountLinking,CreatedAt}`): byte-for-byte consistent with `google_oauth.go:455-460` ✓.
- Doc comments on every exported function/type ✓ (§2).
- Table-driven tests (`[]struct{ name string; … }`) for `TestSetPassword_GenericResponse_AllBranches` and `TestRequireSession_ExpiredOrGarbageToken_401` ✓ (§2).
- `//go:build integration` tag on `security_integration_test.go`; unit tests beside source (`security_test.go` beside `security.go`) ✓ (§3).
- Money/decimal: N/A (no money in this slice) ✓.

---

## Verdict

**Request changes.**

### Blocking (must fix before merge)

1. **S1 / BP1 / C1** — Fix the error dispatch in `SetPasswordHandler` and `UnlinkGoogleHandler`. Preferred: add the two 409 sentinels to `MapServiceError` (per the techplan's own task-05 contract) and route every non-validation error through `MapServiceError(w, err)`; keep the explicit `ErrInvalidCredentials → 401` mapping. Add a transport test injecting a non-sentinel wrapped error and asserting 500 + the `internal` problem type — proves the dispatch, not just that *an* error fired.

### Optional (non-blocking, do not gate merge)

2. **Q1** — Replace `(bool, error)` with a typed result struct or named return for self-documenting Branch-2 signaling.
3. **Q2** — Extract a shared `validatePasswordFields` helper / `minPasswordLength` constant to drop the `< 8` + `"password is not allowed"` duplication with `RegisterHandler`. (Watch scope discipline — touches merged code.)
4. **Q3 / BP2** — Tighten the down-migration comment to state the actual safety contract (safe under code+migration co-rollback; standalone rollback silently skips the R14 audit).

### Process gate (not a code-review finding)

Carried forward from the build report: `gitleaks` is not installed on
the build machine and was not run; full-package `-race` timed out at
120s due to the account package's size (security tests run individually
under `-race`, clean; integration tests under `-race`, clean). The
`gitleaks` gap and the full `-race` run should be re-run in an
environment where they can complete before merge / `make verify`.
