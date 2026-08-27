# Testing Report — MFA TOTP (account #06)

> Location : `.local-agents/works/account/06-mfa-totp/3-build/testing-report.md`
> Date     : 2026-08-27
> Scope    : what was tested, why some runs were slow/hung, issues found, and what remains unrun

## TL;DR

Unit and transport tests are green and fast (~7 s). The expensive `-race` full account
suite passed in ~5.3 min. The **Postgres-backed integration suite was NOT validated to
green** in this session — a real defect in one integration test (an un-closed transaction)
leaked a DB pool connection, which (combined with a long concurrent-confirm test) caused
repeated hangs/timeouts. The user directed that integration/race runs be skipped; the
integration tests are written and compile-checked, with the leaking-tx bug fixed, but remain
**unrun**. `make verify` (staticcheck/gosec/gitleaks/govulncheck/contract) was not run.

## Why testing "took long" this session

### 1. The `-race` full account suite is inherently slow (~319 s)
`go test -race ./internal/domain/account/...` took **319 s (5.3 min)**. Cause: the account
package is a heavy-bcrypt unit suite — registration/login tests run real bcrypt
(`secrets.HashPassword`/`ComparePassword`, ~10⁴–10⁵ rounds each, ~100 ms/call), the
register timing-parity tests deliberately burn comparable bcrypt on every branch, and the
race detector roughly doubles all of it. This is pre-existing weight, not added by Task 06;
the isolated MFA `-race` suite (only `-run TestMfa*`) is fast at **17 s**.

### 2. Integration tests hung / timed out (the actual time-sink)
Running `go test -tags=integration -run 'TestMfa...'` against the local Postgres
(`:5435`) **hung past a 240 s timeout**, and even a scoped run was aborted. Two concrete
root causes:

- **A real pool-leak bug I introduced in `TestMfaBackupCode_SingleUseAndDisabledInvalid_RealDB`**:
  it called `tx, _ := svc.tx.BeginTx(ctx)` and then `_ = tx` — never committing or rolling
  back. Each leaked `pgx.Tx` holds a pooled connection in transaction state; once the pool
  is exhausted, any subsequent `BeginTx` blocks forever. This is exactly how a whole test
  binary stalls. **Fixed before this report** (commit the redeem tx; rollback the no-op
  probe), but the fix is compile-checked, not re-run.
- **The concurrent-confirm test (R8, 100 goroutines against real Postgres)** is genuinely
  long and, on a depleted pool, blocks indefinitely. It also never completed in-session.

### 3. Chained `go build && go vet && go test` commands timed out at 180 s
A single bash command chaining build+vet+two `go test` invocations exceeded 180 s because a
leaked `account.test` process from the earlier hung integration run was still alive,
contending for resources and DB connections. Killing the orphan process restored fast unit
runs (~7 s).

## What WAS validated

| Check | Command | Result | Notes |
|---|---|---|---|
| Build | `go build ./...` | ✅ | all packages |
| Vet | `go vet ./...` | ✅ | no diagnostics |
| Unit (domain account) | `go test ./internal/domain/account/...` | ✅ 6.8 s | incl. all new MFA tests |
| Unit (transport) | `go test ./internal/transport/http/...` | ✅ 0.004 s | incl. MFA handler + marker tests |
| Unit (all) | `go test ./...` | ✅ ~10 s | 8 packages |
| MFA race | `go test -race -run 'TestMfa' ./internal/domain/account/...` | ✅ 17 s | no races found |
| Full race (earlier) | `go test -race ./internal/domain/account/...` | ✅ 319 s | passed before the integration file was added |
| Integration compile | `go vet -tags=integration ./internal/domain/account/...` | ✅ | MFA integration tests compile |

## New MFA tests delivered (unit; all pass)

`internal/domain/account/mfa_test.go` — R1 enroll stores encrypted secret + returns
otpauth URI; R2 restart overwrites pending, stays NULL; R3 409 when active; R4 write-time
guard; R5 confirm enables + exactly 10 hashed codes + `mfa_enabled` audit; INV-account-07
`NoHalfEnabledState`; R6 wrong-code preserves pending (retry without rescan); R7 no-pending ≡
wrong-code; R8 concurrent confirm exactly-one-winner/exactly-10-codes (100 goroutines,
`-short`-skippable); R9 single-use guarded + replay rejects + normalization; R10
`LoginMfa` real-verifier completes/fails with unchanged lockout bookkeeping;
`TestMfaDisable_OldBackupCodesUnusable`; R11 disable success/audit/idempotent-repeat;
R12 reauth failures; R13 Google-only path; R14 server-side provider detection; R15 log
free of secrets.

`internal/transport/http/account_security_mfa_test.go` — R3 wire 409; confirm 200 + R7
byte-identical 422; missing-code 422; disable email-password success/401/422; Google-only
marker consume + replay-401 + password-does-not-bypass (R14); all three handlers 401
without a session (R16).

## Integration tests delivered but NOT run to green

`internal/domain/account/mfa_integration_test.go` (build tag `integration`) covers the
DB-level truths: R1/R3 encrypted-at-rest + 409-leaves-secret; NoHalfEnabled (INV-account-07);
R5 confirm one-tx enable+codes+audit; **R8 ≥100 concurrent confirms exactly-one-winner,
COUNT(codes)=10**; R9 joined-redemption exactly-once + disabled-owner rejection; R11/R12
disable email-password + idempotency + wrong-password.

**Why not validated:** the pool-leak bug above + the long concurrent-DB test caused the runs
to hang and time out; per the user's direction the integration/race runs were skipped.
These must be executed against a reachable Postgres (`DATABASE_URL=… go test -tags=integration -race ./internal/domain/account/... -run 'TestMfa'`) before merge — the pooled
`tx` leak is fixed but unverified.

## Issues found (and dispositions)

1. **Integration test tx leak (pool exhaustion → hang)** — *fixed* (commit/rollback now __explicit__). **Remains unverified** by execution. This is the single biggest cause of the session's perceived slowness/timouts.
2. **Nil seam in shared `integrationService`** — the base helper leaves `now`/`compare`/`mfa` unset; my first integration run panic'd (`MfaEnrollConfirm` → `s.now()` nil deref). Resolved by a local `integrationMFAService` wrapper that seals the seams. (Documented for the pairing/review session; the base helper was not modified to avoid affecting other suites.)
3. **`pgx` import / type slips in the integration file** — resolved (vet-clean).
4. **Pre-existing `securitySchemes` gap in the openapi bundle** — found, unrelated to tests; carried (see build report follow-up #3).
5. **`make verify` not run** — requires staticcheck/gosec on path + a running breach/DB environment; not executed in-session. Should be the first post-merge gate.

## Recommended next actions
1. Book the **Tier 0 pairing** session for `totpVerifier` (D12), then code-review.
2. Run the integration suite against Postgres and confirm the tx-leak fix holds.
3. Run `make verify` (lint/security) after integration is green.
4. Decide the `securitySchemes` `index.yaml` source fix (separate, pre-existing).
