# Area 7 — Test conventions

> Stage 2 gap analysis. Files: `internal/domain/account/*_test.go`,
> `login_integration_test.go`, `repository_db_integration_test.go`.

## Current state

- **Unit tests beside code** with fake repo/sender/breachchecker/clock
  injected into `Service` struct literal (integration helper shows the
  full seam list, login_integration_test.go:54+); table-driven style;
  concurrency tests exist in unit form against fakes
  (`TestVerifyEmail_TokenSingleUse_Concurrent`,
  `TestRefresh_ConcurrentRequests_ExactlyOneWins`).
- **Integration tests** under `//go:build integration`, real Postgres via
  `DATABASE_URL`, skip when unset, `-race` mandated; stress harness
  precedent: `TestRefresh_Stress_MixedValidAndReplayed`
  (tasks.md KPI: ≥100 goroutines).
- **Invariant traceability**: tests named for what they prove
  (`TestRedeemToken_Guards`, `TestRevokeTokens_OnlyUnusedUnrevoked`,
  `TestRedeemAndVerify_Atomic`); real-Postgres timing precedent
  `TestRegister_Timing_AllBranches_RealPostgres`.

## Requirement

Feature-spec threat table + tasks.md KPI demand: generic-response-
all-branches test, single-use concurrent race test (real DB), INV-05
atomic-revoke property test, password-policy + breach-fail-open tests,
rate-limit test — each traceable to its invariant ID.

## Gap

None of these exist yet; all required harness pieces (fake seams,
integration helpers, stress pattern) do.

## Sniffing findings

1. **Risk** — the INV-05 atomicity claim ("revoke happens in the same
   transaction as credential update") is **only provable with real
   Postgres**; fake-repo unit tests cannot distinguish atomic from
   sequential. Integration coverage is mandatory for the Tier 1 gate.
2. **Misleading signal** — the threat model table names specific future
   tests (`TestForgotPassword_GenericResponse_AllBranches`, etc.) that
   read like they exist; they are prescriptive targets, zero written.
3. **Edge case** — forgot-password's three branches do *different DB work*
   (identity lookups ×2 vs token insert), so the dummyWrite-style timing
   shaping question resurfaces; register's real-Postgres timing test is
   the template for proving it.
