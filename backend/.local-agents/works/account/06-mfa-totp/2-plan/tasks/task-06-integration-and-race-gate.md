# Task 06: Integration & race gate — DB-truth suites, stress harnesses, full verify

> Back-reference : `.local-agents/works/account/06-mfa-totp/2-plan/techplan.md` (Status: Approved by Anhar) — sections 12 (Testing Checklist — the authoritative list), 9 (step 10), 7 (top-3 High risks), tasks.md KPI block
> Depends on    : task-05 (all code exists; gate runs against the whole slice)
> Model         : DeepSeek V4 Pro (model-routing Testing/Complex row verbatim: rule-ID count-check reliability once rules ≥15)

## Objective

Prove the three High-severity risk rows under real Postgres contention, close every checkbox in techplan §12, and run the full gate. This is the slice's merge-readiness instrument — its output feeds both the pairing harness and the PR description.

## Files

| File | Change |
|---|---|
| `backend/internal/domain/account/mfa_integration_test.go` | new (`//go:build integration`) — testcontainers suite |

## Scope of proof (mapping §12 → here)

Integration halves already named by unit tasks:
- R1/R3: row-level truth post-enroll/post-409 (byte-for-byte `secret_encrypted` stability on the 409 path)
- R4 `TestMfaEnroll_ConcurrentWithEnable_NeverOverwritesLiveSecret` — **≥100 goroutines, `-race`**, interleaving enroll × {confirm, disable}; invariant asserted = zero overwrites of enabled rows, not "didn't crash" (testing-concurrency checklist)
- R5: committed-row assertion of exactly-10 codes + audit row in-tx; forced mid-tx failure probe leaves neither change applied (§7 risk row 3)
- R8 `TestMfaConfirm_Concurrent_ExactlyOneWinner_TenCodesTotal` — ≥100 confirms, one winner, end-state `COUNT(*)=10`, losers all observed failure
- R9: joined redemption UPDATE under real contention — replay rejects; disabled-owner codes reject with ZERO `used_at` writes
- R11/R13 integration halves as checklist rows specify

Gate sweep (whole slice):
- Full §12 checklist walked item-by-item with rule IDs counted against §4 (count-check re-run at this tier is mandatory)
- `make verify` exits 0 (lint → unit → race → contract → security-A → integration) — backend AGENTS §4
- Coverage ≥80% of new/changed lines per tasks.md KPI
- Audit-value spot-check: exact `mfa_enabled` / `mfa_disabled` literals asserted in DB

## Implementation constraints

- Harness patterns follow task #03/#05 precedents (`login_integration_test.go`, `security_integration_test.go`) including pool-exhaustion lessons from that era's notes
- Deterministic clock injection for lockout-window interplay cases (mirrors existing seams; no sleeps as synchronization)
- No test asserts implementation detail over invariant (per testing-concurrency §1)

## Rules enabled (proven here at integration/race level)

R1–R5, R8–R13 integration halves + full-slice rule-ID closure.

## Verification

- Build report drafted per house convention (what passed, what was deferred intentionally, coverage number, race-harness goroutine counts) — this text becomes raw material for stage 5-testing and the PR body
