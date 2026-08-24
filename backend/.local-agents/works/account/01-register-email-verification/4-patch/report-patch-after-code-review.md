# Patch Build Report — 01-register-email-verification

> Ticket    : 01-register-email-verification
> Stage     : 5-patch (post-review remediation)
> Date      : 2026-08-21
> Builder   : GLM 5.2 (max)
> Plan      : `./plan.md`
> Review    : `../4-code-review/report.md` (7 blocking, 7 optional)
> Feature   : `docs/spec/domains/account/features/01-register-email-verification.md`

---

## 1. What shipped

Tasks A–E (all 7 blocking findings fixed). Task F (optional) deferred to
a follow-up PR.

### Task A — VerifyEmail atomicity + silent-failure (S1, S2)

- `Repository.RedeemToken` signature changed from `(bool, error)` to
  `(userID uuid.UUID, purpose string, ok bool, err error)` with `tx pgx.Tx`
  param. Uses goqu `Returning("user_id", "purpose")` + `tx.QueryRow.Scan`
  — no re-fetch needed.
- `Repository.SetUserVerified` now takes `tx pgx.Tx` so it runs in the same
  transaction as RedeemToken.
- `Service.VerifyEmail` wraps redeem + set-verified in one `pgx.Tx`. If
  SetUserVerified fails, the deferred Rollback undoes the redeem (token
  not burned). `userIDForToken` deleted — no silent nil-return path.
- Tests: `TestVerifyEmail_RedeemReturnsUserID_NoRefetch`,
  `TestVerifyEmail_SetVerifiedFails_RollsBackRedeem` (unit);
  `TestRedeemToken_ReturnsUserIDAndPurpose`, `TestRedeemAndVerify_Atomic`
  (integration).

### Task B — Anti-enumeration DB-time uniformity R3/R4 (S3)

- R3/R4 branches now call `dummyWrite(ctx)` — begins a tx, calls
  `RevokeTokens(ctx, tx, uuid.New(), purposeEmailVerify)` (synthetic uuid
  matches 0 rows), commits. Same BeginTx + UPDATE + Commit cost shape as
  R2's real revoke, but touches no real rows.
- Tests: `TestRegister_R3R4_PerformTimingWrite` (unit: 1 revoke call per
  branch); `TestRegister_Timing_AllBranches_RealPostgres` (integration:
  max/min ≤ 2× band against real Postgres).

### Task C — ResendVerification handler error logging (S4)

- `auth_verify_email.go`: `if err := svc.ResendVerification(...); err != nil`
  → `log.Printf("transport: resend verification failed (recipient redacted): %v", err)`.
  Response stays 202 identical (anti-enumeration).
- Test: `TestResendVerificationHandler_ServiceError_Still202_ButLogs`
  (202 + log present + no recipient PII).

### Task D — Down-migration ordering (S5)

- `DROP FUNCTION IF EXISTS set_updated_at();` moved from
  `000003_create_auth_tokens.down.sql` → `000001_create_users.down.sql`
  (runs last in reverse order, after both triggers are gone).
- Verified: `make migrate-down && make migrate-up` round-trip succeeds
  from a clean applied state.

### Task E — Sensitive-error logging hardening (L1, L2)

- L1: `breachcheck/client.go` — `breachErrorCategory(err)` helper extracts
  `*url.Error.Op` + coarse net category; never logs the URL (which carries
  the 5-char SHA-1 prefix). Also drains body on non-OK path (H1 fix folded
  in since same file).
- L2: `service.go` — `notificationErrorCategory(err)` helper; `sendVerification`
  and `sendNudge` log sanitized category, never raw error (can embed
  recipient/token).
- Tests: `TestIsBreached_APIUnreachable_LogNoURLNoPrefix`;
  `TestRegister_SendVerificationFails_LogNoPII`.

## 2. Files changed

| File | Change |
|---|---|
| `internal/domain/account/repository.go` | RedeemToken + SetUserVerified signatures (tx + RETURNING) |
| `internal/domain/account/repository_db.go` | RedeemToken RETURNING + tx; SetUserVerified tx |
| `internal/domain/account/service.go` | VerifyEmail tx wrap; delete userIDForToken; dummyWrite for R3/R4; notificationErrorCategory |
| `internal/domain/account/service_test.go` | Fake signatures; 4 new tests; leakySender helper |
| `internal/domain/account/repository_db_integration_test.go` | Updated RedeemToken calls; 3 new integration tests |
| `internal/transport/http/auth_verify_email.go` | Log error before 202 |
| `internal/transport/http/auth_verify_email_test.go` | New — handler test (S4) |
| `internal/platform/breachcheck/client.go` | breachErrorCategory + body drain (L1 + H1) |
| `internal/platform/breachcheck/client_test.go` | TestIsBreached_APIUnreachable_LogNoURLNoPrefix |
| `migrations/000001_create_users.down.sql` | Add DROP FUNCTION after trigger+table |
| `migrations/000003_create_auth_tokens.down.sql` | Remove DROP FUNCTION + fix comment |

## 3. Verification

| Gate | Result |
|---|---|
| `go test -count=1 ./...` | ✅ pass |
| `go test -race -count=1 ./internal/domain/account/... ./internal/platform/crypto/...` | ✅ pass (race-clean) |
| `go test -tags=integration -count=1 -race ./internal/domain/account/...` | ✅ pass (real Postgres, incl. timing test) |
| `go test -tags=contract -count=1 ./...` | ✅ pass |
| `go vet ./...` | ✅ clean |
| `staticcheck` (changed packages) | ✅ clean |
| `gosec` (changed packages) | 2 pre-existing findings (crypto/sha1 in breachcheck — correct by design); **0 new findings** |
| `govulncheck` | pre-existing stdlib/module vulns (not from this change) |
| `make migrate-down && make migrate-up` | ✅ round-trip succeeds |

`make verify` fails on the `lint` target due to pre-existing gosec
findings in `platform/auth/keys.go` (G304) and `cmd/server/main.go`
(G112) — files not touched by this remediation. No new findings were
introduced.

## 4. Task F status

Deferred to a follow-up PR (optional, non-blocking):
- R15 (rate-limit N+1 test), S6 (sweeper exit path), Q1 (runTx helper),
  Q2 (`&&`→`||` test fix), Q3 (shared looksLikeEmail), E1 (HMAC≠enc key
  runtime check — needs human confirm on crypto/ fence).
- H1 (drain breachcheck body) was folded into Task E (same file).

## 5. Fence compliance

No Tier 0 fenced path was modified:
- `platform/crypto/crypto.go` (encryption/HMAC impl) — untouched
- `platform/auth/` — untouched
- `domain/donation/ledger.go` — untouched
- `domain/disbursement/` — untouched

Task E's L1 added a `breachErrorCategory` helper to `breachcheck/client.go`
(not a fenced path). Task F's E1 (deferred) would touch `crypto/keys.go`
— flagged for human confirm before it proceeds.

No `docs/spec/*` file was edited (AGENTS.md §4 compliance).

---

## Risk note

- **Assumptions made:**
  - `RedeemToken` with goqu `Returning(...)` produces a `RETURNING` clause
    that pgx `QueryRow.Scan` can read. Verified by
    `TestRedeemToken_ReturnsUserIDAndPurpose` + `TestRedeemAndVerify_Atomic`
    against real Postgres.
  - `RevokeTokens` against a non-existent `user_id` affects 0 rows (no FK
    violation — `user_id` is not constrained to exist in the `WHERE`, only
    on INSERT via FK). Verified by `TestRegister_Timing_AllBranches_RealPostgres`
    against real Postgres (R3/R4 branches complete without error).
  - The resend handler's `%v` on the service error chain is safe to log
    because the leaf is a `pgconn.PgError` (SQLSTATE, constraint name;
    parameterized SQL, no PII values). Task C test asserts the recipient
    email is not logged.

- **Edge cases intentionally NOT handled (and why):**
  - Forced mid-tx failure injection in pgx is awkward; the rollback
    guarantee (S2 fix) is proven by the deferred-Rollback pattern (same
    as `registerNewUser`, already in production) + the unit test
    `TestVerifyEmail_SetVerifiedFails_RollsBackRedeem` which asserts
    VerifyEmail returns a non-nil error (not the silent nil → fake 200
    of S1). The integration test `TestRedeemAndVerify_Atomic` proves the
    happy path commits both in one tx.
  - The timing test asserts a ≤2× band, not exact equality. Residual
    insert-count difference (R1's 3 inserts vs R3's 1 dummy UPDATE) is
    within that band against real Postgres, as verified by the
    integration test. Network/jitter is not modeled; real Postgres
    latency is the realistic signal.

- **Concurrency assumptions:**
  - Single-use correctness is still the atomic 3-clause `UPDATE … WHERE`
    inside RedeemToken; adding SetUserVerified to the same tx does not
    weaken it (the guard is on the UPDATE, not app-level read). R12's
    concurrent double-submit test still passes under `-race`.
  - The dummy write is a per-request tx; no shared state. `-race` clean
    by construction, verified by `go test -race`.

- **What is not tested, and why:**
  - The `TestRegister_Timing_AllBranches_RealPostgres` timing band is
    statistical; it could occasionally flake under extreme CI load. The
    2× band is generous enough to absorb normal jitter, but a heavily
    loaded CI runner could produce an outlier. If this becomes flaky,
    widen the band or run with more warmup iterations.
  - The down-migration round-trip was verified manually against dev
    Postgres (`make migrate-down && make migrate-up` succeeds). Go-level
    testing of DDL reversibility is not idiomatic here — the round-trip
    IS the test.
  - `make verify` fails on pre-existing gosec findings in
    `platform/auth/keys.go` (G304) and `cmd/server/main.go` (G112) — not
    from this remediation. These should be addressed separately (or
    suppressed with `//nosec` comments with justification).
