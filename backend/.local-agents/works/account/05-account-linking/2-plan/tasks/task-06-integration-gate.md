# Task 06: Integration & race suite + full gate

> Back-reference : `.local-agents/works/account/05-account-linking/2-plan/techplan.md` (Status: Approved) — sections 4 (all), 9 (step 7), 12 (Testing Checklist), 13; tasks.md KPI table (`docs/spec/1-account/tasks.md`)
> Depends on    : tasks 01–05 (everything under test exists)
> Model         : DeepSeek V4 Pro (model-routing Testing/Complex row: rule-ID count-check needs more reliability than Flash once rule count ≥15)

## Objective

Close the DB-level truths a fake repository cannot prove (constraint-backed race behavior, transaction atomicity, lock serialization) via testcontainers integration tests, then run the full repo gate and confirm every KPI from `docs/spec/1-account/tasks.md`.

## Files

| File | Change |
|---|---|
| `backend/internal/domain/account/security_integration_test.go` | New (`//go:build integration`) |

## Integration coverage (DB truths only — unit-proven logic is NOT re-tested here per integration-testing-setup step 0 "spot-check, don't rewrite")

- **R1** committed-row assertions: unverified identity + token with `purpose='email_verification_link'` in one tx
- **R3** concurrent duplicate email through the real unique index `(provider_type, identifier_hash)` — exactly one winner, loser clean
- **R7** Branch 2 atomicity against real Postgres: multi-session fixture (≥2 refresh families), secret rotated + ALL sessions revoked atomically; forced mid-tx failure leaves NEITHER applied; `identifier_hash` untouched
- **R9** unlink DB truth: google rows hard-deleted, exact `user_logs.action_type='account_linking'` row committed atomically
- **R13** the stress harness re-run against real Postgres locks (≥100 goroutines, end-state invariants)
- **R14** link-purpose redemption writes the audit row in-tx; registration-purpose does not; `TestVerifyEmail_TokenSingleUse_Concurrent` still green untouched
- Migration 000010 interplay: token insert with new purpose succeeds post-migration (task-01's probe already covered DDL; here only app-path inserts)

## Gate checklist (techplan §12 carried forward — run, don't re-derive)

- [ ] R1–R16 all have their named test passing somewhere (unit suite: tasks 03–05; integration: this task). Count-check R1–R16 explicitly.
- [ ] `make verify` exits 0 (lint → unit → race → contract → security layer A → integration)
- [ ] `go test -race ./...` clean (backend AGENTS §3 requirement for Tier 0/1-adjacent work)
- [ ] Coverage ≥80% of new/changed lines in touched packages (tasks.md KPI)
- [ ] Security layer A (gosec/gitleaks/govulncheck): 0 findings or explicit accepted-risk note
- [ ] Handler↔spec match spot-check incl. both 409 error shapes (openapi-spec-first-drift checklist)
- [ ] Named invariant traceability per KPI: INV-account-01 (`TestSetPassword_ConcurrentDuplicateEmail_Race`), INV-account-02+12 (`TestUnlinkGoogle_*Guard*`, `*_OnlyIdentity_*`, `*_RejectsUnverifiedRemainingIdentity`), INV-account-05 (`TestSetPassword_Branch2_AllSessionsRevoked`), INV-account-08 (`TestVerifyEmail_TokenSingleUse_Concurrent`), audit KPI (`action_type` value asserted in R9/R14 tests)

## Report obligation

Produce the build report (`3-build/report.md` pattern) with a rule-coverage table and an explicit "what is not tested, and why" section — claims without named tests count as unverified (root AGENTS §5).

## Common mistakes

- Rewriting unit-level proofs as integration duplicates → pure cost, no signal; integrate only what fakes can't prove
- Skipping `-race` on the integration pass because "the unit pass had it" → both required for this domain
- Marking gate rows done from memory → run each command, paste exit status into the report

## Verification

Full `make verify` exit 0 is the task's definition of done.
