# Task 06 — Integration & race suite: INV-05 property, INV-08 stress, timing parity

> Back-reference (contract): `../techplan.md` — sections 1–8 are the source of truth. Techplan wins over this file on any apparent conflict.
> Splitting axis: dependency/sequence chain (see `manifest.md`). Depends on Task 03 (service flows exist); benefits from Task 04 for end-to-end shape but does not hard-require it except where noted.

## Scope

**In scope:**
- New integration test file `internal/domain/account/password_reset_integration_test.go` (`//go:build integration`)
- The three proofs that fakes structurally cannot deliver: INV-account-05 atomicity, INV-account-08 single-use stress, forgot-password branch timing parity
- Full-gate run of the whole domain suite

**Out of scope (this task):**
- Unit tests (Tasks 01/03/04 own those)
- Any production-code change — if a bug surfaces here, fix lands in the owning task's files and this suite re-runs

## Dependencies

- Task 03 merged (flows callable against real DB)
- Real Postgres reachable via `DATABASE_URL`; tests skip cleanly when unset (existing convention, login_integration_test.go:42–45)
- R6 rate-limit proof lives in Task 04 (handler level) — not duplicated here

## Why these must be integration tests (techplan Area-7 finding)

The INV-05 atomicity claim ("revoke happens in the same transaction as credential update") and the INV-08 single-use guard are properties of Postgres behavior under contention/crash boundaries — a fake repository cannot distinguish atomic from sequential execution. Tier 1 gate makes these mandatory, not optional (tasks.md KPI: named invariant test traceable to each referenced `INV-account-NN`, `-race` clean, ≥100 concurrent goroutines, 0 invariant violations).

## Rules this task must prove (verbatim from techplan §4)

- **R7**: valid token + passing password → ONE tx: credential updated, used_at set, EVERY refresh row for the user revoked.
- **R9**: expired token → no state change.
- **R11**: N≥100 concurrent resets submitting the same valid token → exactly one 200-equivalent success, others rejected, exactly one credential update, zero anomalies.
- **R5**: forgot-password branches indistinguishable by wall-clock DB time.
- **R18 (INV-05 property)**: credential-update-committed ⟺ all-sessions-revoked-committed; injected failure between writes rolls back BOTH, leaving token redeemable and sessions alive.

## Test plan (redistributed from techplan §12 — this task's rows)

File header comment follows house style (see login_integration_test.go:1–17): build tag, what real Postgres proves that fakes can't, run command.

- [ ] `TestResetPassword_AllSessionsRevoked_Atomic` — seed user with ≥2 ACTIVE refresh tokens across ≥2 different families (one rotated-out mid-chain, mirroring real device usage); complete password reset; assert every row revoked AND credential updated, single transaction boundary respected. This is the named invariant test for INV-account-05.
- [ ] `TestResetPassword_TokenSingleUse_Concurrent` — spawn goroutines double-submitting the SAME valid token; assert exactly one success path, all others rejected with token-consumed semantics. Named invariant test for INV-account-08 (unit-level sibling exists in Task 03's fake form; this is the real-contender version).
- [ ] `TestResetPassword_Stress_MixedValidAndReplayed` — ≥100 goroutines mixing valid submits and replays of consumed/expired tokens (model: `TestRefresh_Stress_MixedValidAndReplayed`, login_integration_test.go:207); assert 0 invariant violations across the run.
- [ ] `TestResetPassword_FailureBetweenWrites_RollsBackBoth` — inject failure into `UpdateIdentityCredentialSecret` or `RevokeAllRefreshTokensForUser` (test-only repo wrapper that errors on call N); assert rollback left: token still redeemable (used_at NULL), sessions still alive, error surfaced. This is R18's rollback arm.
- [ ] `TestForgotPassword_Timing_Branches_RealPostgres` — measure wall-clock across the three forgot branches against real Postgres; assert no significant found/not-found distinction (model: `TestRegister_Timing_AllBranches_RealPostgres`, repository_db_integration_test.go:1117). Proves the `dummyWrite` shaping actually works at the DB layer.
- [ ] Expired-token no-state-change check rides the existing helpers (`TestRedeemToken_Guards` precedent, repository_db_integration_test.go:248) — assert R9 leaves zero mutated rows.

Run command (matches AGENTS.md §3 / tasks.md KPI):
```
go test -tags=integration -race ./internal/domain/account/...
```

## Harness conventions (from existing integration files)

- Build the service over a REAL pool with trivial injected seams: `integrationTestKeys(t)`, real `RepositoryDB`, `poolRunner{pool}`, fake breach checker returning false, silent sender, stub MFA (login_integration_test.go:40–60 pattern)
- `t.Skip("DATABASE_URL not set...")` when env missing
- Seed helpers create users/identities/tokens directly through the repo with known hashes — reuse whatever seeding helpers tasks #03's suite already established (TBD — verify their exact names at build time; do not duplicate)

## Common mistakes that apply here (techplan §13)

| Mistake | Fix |
|---|---|
| Asserting "didn't crash" instead of invariants in concurrency tests | R11 asserts exact winner count + single credential update |
| Timing test comparing HTTP bodies only | DB-time channel needs real Postgres wall-clock measurement (R5) |
| Writing a second unit test duplicating what a named integration test proves | Spot-check discipline (integration-testing best-practice): don't re-prove the same rule the same way |
| Skipping `-race` because the suite is slow | Non-negotiable for this tier — tasks.md KPI |

## Gate

Full gate per techplan §12 hygiene: `go test ./...`, `go test -race ./...`, `go test -tags=integration -race ./internal/domain/account/...` — all green, 0 skipped-in-CI surprises (DATABASE_URL-dependent skips acceptable locally only).
