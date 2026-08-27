# Patch Plan — Account Linking (account #05), after code review

> Source      : `4-code-review/report.md` (Verdict: Request changes — 2 blocking, 3 optional)
> Scope       : address the blocking error-dispatch defect (S1 / BP1 / C1, one root cause) and the optional follow-ups flagged by the review
> Tier        : **Simple** (one localized correctness fix + tests; no new domain logic, no schema change, no contract change)
> Boundaries  : backend only; Tier 0 fences (`platform/auth/`, `platform/crypto/`, `donation/`, `disbursement/`) untouched

## What the review found

The two new handlers in `internal/transport/http/account_security.go`
inline their error dispatch and route **every** non-`ErrValidation`
error (in `SetPasswordHandler`) / non-409-sentinel error (in
`UnlinkGoogleHandler`) to `401 invalid-credentials`. That misclassifies
internal failures (DB outage, `BeginTx`/`Commit` failure,
`InsertAuthIdentity` non-unique-non-23505, `RevokeAllRefreshTokensForUser`,
`DeleteAuthIdentitiesByIDs`, `InsertUserLog`, the "missing
credential_secret" data-corruption branch) as a credentials problem —
wrong category, breaks anti-enumeration on Branch 1, and bypasses the
repo's `MapServiceError` convention established in `errors.go:76-109`
and used by `RegisterHandler` / `ResetPasswordHandler`. The techplan's
own task-05 Files table specified "+ two 409 cases in `MapServiceError`"
for `errors.go`; the build skipped that and inlined the 409s in the
handlers instead.

## Decision

Fix per the original techplan contract: extend `MapServiceError` with
the two new 409 sentinels and route every non-validation error through
`MapServiceError` from both handlers. This collapses S1 + BP1 + C1 into
a single change and removes the duplicated inline 409 strings as a
bonus. Keep an explicit `ErrInvalidCredentials → 401` mapping in each
handler (it is the only legitimate credentials-shaped failure on these
endpoints and must stay visible/testable at the handler boundary per
the root AGENTS.md golden rule "every authorization check is explicit").

The optional follow-ups (Q1 typed result, Q2 shared validation helper,
Q3/BP2 down-migration comment) are recorded as a separate, clearly-
marked optional task — they do not gate merge.

## Out of scope

- The `gitleaks` + full-`-race` process gates carried forward from the
  build report — those are environmental, not code defects; they belong
  to the testing phase (`5-testing`), not this patch.
- The `InsertAuthIdentity` `verified_at` omission (build report
  deviation #5) — pre-existing, out of slice scope.
- The openapi bundle regeneration gap (deviation #4) — source-side fix,
  out of slice scope.

## Verification gate for this patch

```bash
go test ./internal/transport/http/...                       # new dispatch tests
go test -race ./internal/domain/account/...                 # security tests still clean
go test -tags=integration ./internal/domain/account/...     # 4 integration tests still clean
staticcheck ./...
```

(`make verify` full run is the testing phase's job — this patch only
needs to prove the dispatch change is correct and does not regress the
existing security/integration suites.)

## Risk note

- **Assumptions made**: `MapServiceError`'s `default` 500 + server-side
  `log.Printf("transport: unhandled service error: %v", err)` is the
  correct sink for internal errors on these endpoints (matches
  `RegisterHandler`). The wrapped `%w` errors carry SQLSTATE / constraint
  names from goqu `Prepared(true)` parameterized SQL — no PII values, so
  the `%v` log line is safe per `go/secrets-and-sensitive-logging.md`.
- **Edge cases intentionally NOT handled**: the residual anti-enumeration
  gap (a 500 on Branch 1 is still distinguishable from 202) is NOT closed
  by this patch — it is the same accepted residual as register/resend
  (codebase-wide), flagged for a future async-send pass, not this slice.
- **Concurrency assumptions**: unchanged — no concurrency code touched.
- **What is not tested**: the `gitleaks` and full-package `-race` runs
  (environmental gaps from the build report, not patch-introduced).
