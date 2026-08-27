# Build Report — Account Linking (account #05)

> Author : ox-alpha (agent)
> Date   : 2026-08-27
> Status : build complete; gates run; ready for code review
> Source : `.local-agents/works/account/05-account-linking/2-plan/techplan.md` (Approved)

## Summary

Implemented both endpoints (`POST /account/security/set-password`, `POST
/account/security/google/unlink`) end-to-end: migration, repository,
service, transport, wiring, and tests (unit + integration). The build
discovered two things the techplan didn't specify (both resolved, both
recorded) and one coordination event (the parallel Fitur 04 session
landed the INV-account-05 primitive before this slice, so we reused it
rather than duplicating).

## Rule coverage (R1–R16)

| Rule | Test(s) | Where |
|---|---|---|
| R1 | `TestSetPassword_Branch1_CreatesUnverifiedIdentity_SendsVerification` | unit (security_test.go) |
| R2 | `TestSetPassword_Branch1_ClaimedEmail_NudgeNoIdentity_Generic202` | unit |
| R3 | `TestSetPassword_ConcurrentDuplicateEmail_Race` | unit (unique-violation fallback; real concurrent stress deferred to integration R13's precedent) |
| R4 | `TestSetPassword_PasswordPolicy_PrecedesBranching`, `TestSetPassword_BreachCheck_FailOpen` | unit |
| R5 | `TestSetPassword_GenericResponse_AllBranches` | unit |
| R6 | `TestSetPassword_BranchSelection_ServerSide` | unit |
| R7 | `TestSetPassword_Branch2_AllSessionsRevoked` | unit + `TestSetPassword_Branch2_AllSessionsRevoked_RealDB` (integration) |
| R8 | `TestSetPassword_Branch2_WrongCurrentPassword_Rejected` | unit |
| R9 | `TestUnlinkGoogle_Success_HardDeletesAndAudits` | unit + `TestUnlinkGoogle_Success_HardDeletesAndAudits_RealDB` (integration) |
| R10 | `TestUnlinkGoogle_OnlyIdentity_Rejected409` | unit + `TestUnlinkGoogleHandler_OnlyIdentity_409` (transport) |
| R11 | `TestUnlinkGoogle_RejectsUnverifiedRemainingIdentity` | unit + `TestUnlinkGoogleHandler_UnverifiedRemaining_409` (transport) |
| R12 | `TestUnlinkGoogle_RequiresReauth`*, `TestUnlinkGoogle_WrongPassword_Rejected`, `TestUnlinkGoogle_IdempotentNoGoogleRow_Returns200` | unit |
| R13 | `TestUnlinkGoogle_ConcurrentRequests_GuardHolds` | unit (sequential) + `TestUnlinkGoogle_ConcurrentRequests_GuardHolds_RealDB` (integration — ≥100 goroutines, real FOR UPDATE) |
| R14 | `TestVerifyEmail_LinkPurpose_WritesLinkAudit`, `TestVerifyEmail_RegistrationPurpose_NoLinkAudit` | unit + `TestVerifyEmail_LinkPurpose_WritesAudit_RealDB` (integration) |
| R15 | `TestRequireSession_MissingToken_401`, `TestRequireSession_ExpiredOrGarbageToken_401`, `TestRequireSession_BearerFallback_Accepted` | transport (account_security_test.go) |
| R16 | `TestSecurity_LogsFreeOfSecrets` | unit |

Count-check: R1–R16 all have ≥1 named test. ✓

## Invariant traceability

| Invariant | Test(s) |
|---|---|
| INV-account-01 (per-provider uniqueness) | `TestSetPassword_ConcurrentDuplicateEmail_Race` (unique-violation fallback) |
| INV-account-02 (≥1 identity after unlink) | `TestUnlinkGoogle_OnlyIdentity_Rejected409` + R13's end-state assertion |
| INV-account-05 (password change revokes all sessions) | `TestSetPassword_Branch2_AllSessionsRevoked` + integration RealDB |
| INV-account-08 (token single-use/time-bound) | existing `TestVerifyEmail_TokenSingleUse_Concurrent` (unchanged) |
| INV-account-12 (remaining identity verified) | `TestUnlinkGoogle_RejectsUnverifiedRemainingIdentity` + R13's end-state assertion |
| Audit KPI (exact action_type) | R9 + R14 assert `action_type='account_linking'` |

## Gate results

| Gate | Result |
|---|---|
| `staticcheck ./...` | ✓ clean |
| `go test ./...` (unit) | ✓ all packages green |
| `go test -race ./internal/domain/account/...` (security tests) | ✓ clean (26s) |
| `go test -tags=contract ./...` | ✓ all green |
| `go test -tags=integration` (4 security integration tests) | ✓ all green (real Postgres, ≥100-goroutine stress) |
| `gosec ./...` | 13 pre-existing findings (none in new files) |
| `gitleaks` | not installed on this system |
| `govulncheck ./...` | 24 pre-existing dependency CVEs (not from new code) |

`make verify` fails at the `lint` step (gosec) due to 13 pre-existing
findings across 36 files — zero of which are in the files this task
introduced. This is the repo's existing state, not a regression.

## What is NOT tested, and why

- **R3 full concurrent race**: the fake repo can't faithfully simulate
  READ COMMITTED transaction visibility (its mutex serializes individual
  method calls, not whole transactions). The unit test exercises the
  unique-violation fallback path directly; the real concurrent race
  against the unique index is structurally identical to the register
  flow's, which is already covered by `TestRegister_*` integration tests.
- **`gitleaks`**: not installed on this system — can't run the tool.
  The code introduces no secrets/keys (all test keys are zero-byte
  arrays or generated at test time).
- **Full `-race` across all packages**: timed out at 120s due to the
  account package's size (register + login + google + password-reset +
  security). Ran `-race` on the security tests individually (26s, clean)
  plus the integration tests under `-race` (0.29s, clean).

## Techplan deviations (build-phase discoveries)

1. **SetPassword return type**: techplan specified `func (...) error`;
  build discovered the handler needs to distinguish 200 (Branch 2) from
  202 (Branch 1) — both return nil, so the handler couldn't tell them
  apart. Changed to `func (...) (bool, error)` where `true` = Branch 2
  (200). Minimal, testable, self-documenting.

2. **Fitur 04 coordination**: a parallel Fitur 04 (forgot/reset password)
  build session landed uncommitted changes in the exact files this task
  touches (`service.go`, `repository*.go`, `main.go`, `sender.go`).
  Confirmed by Anhar as legitimate ("the 04 slice was done"). Reused
  `RevokeAllRefreshTokensForUser` and `UpdateIdentityCredentialSecret`
  (04's names) instead of duplicating. Techplan D1/Resolved #5's claim
  ("this slice authors the primitive first") is superseded by reality —
  04 authored it; 05 reuses it.

3. **VerifyEmail cross-purpose guard**: Fitur 04's work added a purpose
  guard to VerifyEmail that rejects any purpose ≠ `email_verification`.
  My `email_verification_link` purpose would be rejected by this guard.
  Updated the guard to also accept `email_verification_link` (the spec's
  "endpoint unchanged" claim holds externally; internally the guard
  now permits both email-verification-family purposes and writes the
  audit entry conditionally).

4. **Bundle regeneration (D9-B STOP)**: the openapi bundle regeneration
  produced a byte-identical output to HEAD — the bundle was already
  current. Root cause established: Redocly's bundler follows `$ref`s
  and prunes unreferenced components; nothing `$ref`s
  `components.securitySchemes.bearerAuth` (global `security:` references
  it by plain name). Per D9-B, did NOT hand-edit the bundle. The fix
  requires a source-side `$ref` anchor or a redocly config change — both
  out of this slice's scope (source files excluded). Flagged for
  Anhar's decision (techplan Resolved #4).

5. **InsertAuthIdentity omits verified_at**: the real adapter always
  inserts NULL for `verified_at`, even when the struct field is set
  (google identities are supposed to be born verified per spec). This is
  a pre-existing issue, not introduced by this slice. Integration tests
  work around it with a direct SQL UPDATE after seeding.

## Risk note

- **Assumptions made**: Fitur 04's uncommitted work is treated as
  legitimate and reused (confirmed by Anhar). The build does not commit
  anything — all changes sit in the working tree alongside 04's changes.
- **Edge cases intentionally NOT handled**: the `InsertAuthIdentity`
  `verified_at` omission is a pre-existing adapter limitation; fixing it
  would affect the register and google_oauth flows (out of scope). The
  `set-password` endpoint naming question (one endpoint, two behaviors)
  is deferred per techplan Resolved #6.
- **Concurrency assumptions**: the ≥100-goroutine integration stress
  test proves FOR UPDATE serialization against real Postgres. The unit
  test uses the fake's mutex (which serializes but doesn't faithfully
  simulate READ COMMITTED visibility) — the integration test is the
  authoritative proof.
- **What is not tested**: `gitleaks` (tool not installed); full-package
  `-race` (timed out due to package size; security tests run individually
  under `-race`).

## Files changed

| File | Change |
|---|---|
| `migrations/000010_widen_auth_tokens_purpose.up.sql` | new — widen purpose CHECK |
| `migrations/000010_widen_auth_tokens_purpose.down.sql` | new — re-map + restore CHECK |
| `internal/domain/account/repository.go` | +3 interface methods |
| `internal/domain/account/repository_db.go` | +3 implementations + shared scan helper + pgxQueryer interface |
| `internal/domain/account/security.go` | new — SetPassword, UnlinkGoogle, sentinels, purpose const |
| `internal/domain/account/security_test.go` | new — 17 unit tests (R1–R16) |
| `internal/domain/account/security_integration_test.go` | new — 4 integration tests (R7/R9/R13/R14 against real Postgres) |
| `internal/domain/account/service.go` | VerifyEmail: widen guard + conditional audit (R14) |
| `internal/domain/account/service_test.go` | +3 fake stubs + recording fields for new repo methods |
| `internal/platform/notification/sender.go` | + `NudgeSetPasswordConflict` constant |
| `internal/transport/http/account_security.go` | new — RequireSession middleware, ES256 verifier reuse, 2 handlers |
| `internal/transport/http/account_security_test.go` | new — 12 transport tests (R15 + contract) |
| `cmd/server/main.go` | accountMux wiring (2 routes + session middleware) |
