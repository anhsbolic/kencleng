# Patch Report — Account Linking (account #05), after code review

> Date    : 2026-08-27
> Status  : patch complete; all gates green
> Source  : `4-patch/after-code-review/plan.md` + `tasks/{manifest,task-01-error-dispatch-fix,task-02-optional-cleanup}.md`
> Review  : `4-code-review/report.md` (Verdict: Request changes — 2 blocking, 3 optional)

## Summary

Resolved the blocking finding from the code review — the wrong-error-
category defect (S1 / BP1 / C1, one root cause) — by extending
`MapServiceError` with the two new 409 sentinels and routing both
security handlers' non-validation errors through it. Also landed the
one zero-risk optional item (Q3 down-migration comment tightening).
Q1 and Q2 were **deferred** per their task file's own scope-discipline
guidance (rationale below).

## Findings disposition

| Finding | Pass | Severity | Disposition |
|---|---|---|---|
| S1 — internal errors returned as 401 invalid-credentials | Safety | Blocking | **Fixed** (task-01) |
| BP1 — anti-enumeration: 401-on-internal leak on Branch 1 | Best Practices | Blocking | **Fixed** (task-01; same root as S1) |
| C1 — handlers bypass `MapServiceError`, duplicating inline Problem strings | Consistency | Blocking | **Fixed** (task-01; same root as S1) |
| Q1 — ambiguous `(bool, error)` return from `SetPassword` | Quality | Optional | **Deferred** |
| Q2 — duplicated boundary-validation literals with `RegisterHandler` | Quality | Optional | **Deferred** |
| Q3 / BP2 — down-migration "purpose-blind" comment overstated | Quality/BP | Optional | **Fixed** |

## Changes

### Task-01: error dispatch fix (BLOCKING — all three findings in one change)

1. **`internal/transport/http/errors.go`** — `MapServiceError` gained
   two new case branches:
   - `account.ErrOnlyIdentity` → `409`, type `https://kencleng.dev/errors/only-identity`
   - `account.ErrRemainingUnverified` → `409`, type `https://kencleng.dev/errors/unverified-remaining-identity`

   Titles + Indonesian details moved **verbatim** from
   `UnlinkGoogleHandler` — restoring exactly what the techplan's task-05
   Files table originally specified ("+ two 409 cases in
   `MapServiceError`") and collapsing the duplicate-string maintenance
   risk C1 flagged.

2. **`internal/transport/http/account_security.go`** — both handlers
   simplified:
   - `SetPasswordHandler`: `isErrValidation` branch (422) kept explicit,
     everything else now a single `MapServiceError(w, err)` call
     (`ErrInvalidCredentials` → 401 via the existing mapping at
     `errors.go:98-101`; wrapped internal errors → the existing 500
     generic default).
   - `UnlinkGoogleHandler`: the entire inline 409/401 `switch` replaced
     by one `MapServiceError(w, err)` call.
   - All inline duplicate Problem strings removed. The handler file no
     longer imports `errors` or the `account` package at all.
   - Comments cite the review finding IDs so the next reader knows why
     delegation (not a bare `default → 401`) is load-bearing here.

3. **`internal/transport/http/account_security_test.go`** — two new
   tests proving the dispatch category, per the review's "prove the
   dispatch, not just that an error fired" requirement:
   - `TestSetPasswordHandler_InternalError_500` — stub returns a wrapped
     non-sentinel error (`fmt.Errorf("...: %w", errors.New("db down"))`);
     asserts `500` + problem type `https://kencleng.dev/problems/internal`
     + detail not equal to the credentials-shaped generic message.
   - `TestUnlinkGoogleHandler_InternalError_500` — same shape for unlink.

### Task-02 (partial): Q3 down-migration comment

4. **`migrations/000010_widen_auth_tokens_purpose.down.sql`** — comment
   rewritten to state the actual safety contract instead of the
   incorrect "token redemption is purpose-blind" claim: safe when the
   migration rolls back alongside the feature code (rolled-back code has
   no audit logic; its 2-value guard accepts the remapped token); a
   standalone migration-only rollback would silently skip the R14 audit
   for link-purpose tokens redeeming as registration tokens — explicitly
   marked "do not roll back this migration alone." SQL unchanged;
   comment-only edit.

## Deferred items (with rationale)

- **Q1 (typed `SetPasswordOutcome` result)** — mechanical rename
  cascading through service → `securityService` interface → transport
  stub → ~17 unit-test call sites, for zero behavior gain; the doc
  comment on `SetPassword` already documents the bool semantics. Pure
  readability churn on an unmerged-but-stable slice; better taken in a
  dedicated cleanup pass than mixed into this correctness patch.
- **Q2 (shared password-validation helper)** — requires editing
  `auth_register.go`, which is already-merged register-slice code
  outside this story's scope boundary. Root AGENTS.md §7 (scope
  discipline): the minimal-risk move is deferring to a later cleanup
  pass rather than touching a second slice's files inside this patch.
  Neither item gates merge per the code-review verdict.

## Gate results

| Gate | Result |
|---|---|
| `go build ./...` | ✓ clean |
| `go test ./internal/transport/http/...` | ✓ ok (new 500-dispatch tests pass) |
| `go test -run <dispatch+401+409 tests>` (verbose) | ✓ 6/6 PASS (`_InternalError_500` × 2, `_WrongPassword_401` × 2, `_OnlyIdentity_409`, `_UnverifiedRemaining_409`) |
| `go test -race ./internal/domain/account/...` | ✓ ok (182.8s, clean) |
| `staticcheck ./internal/transport/http/` | ✓ exit 0 |

Not re-run here (unchanged by a transport-only dispatch fix):

- `go test -tags=integration ./internal/domain/account/...` — the patch
  touches neither repository adapters nor any DB path; the build report's
  4 green integration tests are structurally unaffected.
- Full `make verify` — belongs to the testing phase (`5-testing`) /
  final PR gate, including the environmental gaps carried forward
  (`gitleaks` not installed; full-package `-race` timeout).

## Risk note

- **Assumptions made**: delegating to `MapServiceError` preserves
  byte-identical responses for every previously-tested error category —
  verified by the pre-existing handler-contract tests
  (`TestSetPasswordHandler_WrongPassword_401`,
  `TestUnlinkGoogleHandler_OnlyIdentity_409`,
  `TestUnlinkGoogleHandler_UnverifiedRemaining_409`,
  `TestUnlinkGoogleHandler_WrongPassword_401`,
  `TestUnlinkGoogleHandler_Success_200`,
  `TestSetPasswordHandler_Branch{1,2}_*`,
  `TestRequireSession_*`) staying green without modification. The moved
  409 titles/details were copied verbatim, confirmed by those tests'
  exact string assertions passing.
- **Edge cases intentionally NOT handled**: the residual
  anti-enumeration gap remains open by design — a 500 on SetPassword
  Branch 1 is still distinguishable from the generic 202. This is the
  same accepted residual as register/resend (codebase-wide async-send
  future work), documented in the review's BP1 and the patch plan; the
  fix removes the strictly-worse credentials-shaped response, not the
  whole gap.
- **Concurrency assumptions**: none changed — no concurrency code,
  transaction boundaries, locks, or schema touched. Single-threaded
  handler logic only.
- **What is not tested, and why**: integration suite not re-run (no DB
  path changed); full-package `-race` under `go test -race ./...` (the
  account package alone takes ~183s, consistent with the build report's
  documented 120s-timeout environment limitation — security package ran
  clean with `-race` here); `gitleaks` (tool unavailable in this
  environment, carried forward as a process gate). Every claim of "this
  edge case is handled" above names its proving test.

## Feature spec reference

Fulfills the post-review corrective work for
`docs/spec/1-account/features/05-account-linking.md` — specifically
restores techplan §10/task-05 contract fidelity (sentinel→Problem
mappings centralized per `openapi-spec-first-drift` discipline) without
altering any R1–R16 rule coverage or acceptance criterion from the
original feature spec.
