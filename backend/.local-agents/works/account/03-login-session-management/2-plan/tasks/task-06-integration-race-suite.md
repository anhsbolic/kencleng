# Task 06: Integration + race suite, final gate

> Back-reference : `.local-agents/works/account/03-login-session-management/2-plan/techplan.md` (Status: Approved) — sections 9 (testing strategy), 12 (full checklist), 7 (High risks)
> Depends on    : tasks 01–05 all merged into the working tree
> Model         : DeepSeek V4 Pro (Testing row, Complex tier: rule-ID count-check reliability at ≥15 rules)
> Rules touched : R12, R14, R15 (proof), R20 (round-trip), plus whole-plan gate
> Tier 0        : none authored here — but the suite is the proof layer for Tier 0 rotation logic

## Objective

Prove the invariants a fake repository cannot: rotation single-use under true concurrency (INV-account-03), family-wide reuse revocation (INV-account-04), and migration reversibility — then run the full `make verify` gate.

## Files

| File | Change |
|---|---|
| `backend/internal/domain/account/integration_test.go` | New (`//go:build integration`) |
| `backend/internal/domain/account/race_test.go` | New (`//go:build integration`) |

## Concurrency proof (INV-account-03/04)

Setup: testcontainers/real Postgres via `testcontainers-go` per existing integration pattern; seed a user + fresh `family_id` token chain.

1. **`TestRefresh_ConcurrentRequests_ExactlyOneWins`**: N goroutines (start with 2 for determinism, then stress variant) call `svc.Refresh(samePlainToken)` simultaneously.
   Assert exactly one `nil` error; every other attempt fails; DB state afterward:
   - exactly ONE child row exists for the parent (`replaced_by_id` unique per parent);
   - parent's `replaced_by_id` = that child's id;
   - no orphaned second child anywhere in the family.
2. **Stress variant (tasks.md KPI)**: ≥100 concurrent goroutines on one valid token + mixed replayed-token attackers. Invariants across the run: 0 double-parenting; every replay ⇒ full-family `revoked_at`; final live-token count == number of successful sequential rotations. Run under `-race`.
3. **`TestRefresh_ReuseDetection_FamilyRevoked`** (A→B→C chain): rotate A→B, B→C, then replay A ⇒ assert A, B, C ALL have `revoked_at IS NOT NULL`, and a subsequent refresh with C also fails.
4. **Lockout window boundary** against real clock math: insert failures at now-14min (counts) vs now-16min (doesn't), asserting the strictly-greater cutoff semantics chosen in task-02.

## Migration round-trip (R20)

In the integration pass: `make migrate-up && make migrate-down && make migrate-up` exit 0; spot-check `login_attempts` partial indexes exist post-up.

## Full-slice gate

```bash
go build ./...
go vet ./...
go test ./...                 # unit, fast
go test -race ./...           # concurrency check (backend AGENTS.md §3 mandate)
go test -tags=integration ./... 
make verify                   # full gate: lint, unit, race, contract, security-A, integration
```

Then run the **rule-ID count reconciliation** (techplan §12): confirm R1–R20 each map to ≥1 passing named test across tasks 03–06; report any gap as a blocker, not a note.

## Reporting duties (carry into task report / PR)

1. **"Tier 0 files awaiting paired rewrite" flag list** — per techplan Resolved #13, name: `internal/platform/auth/token.go` (T3), rotation methods in `internal/domain/account/repository_db.go` (T2), reuse/race-loser branch in `internal/domain/account/login.go` (T4). Nothing Tier 0 commits before Anhar's paired pass.
2. Risk note skeleton per root AGENTS.md §5 (assumptions, unhandled edges, concurrency assumptions, untested areas) — claims must cite the specific test above that proves them.
3. Feature-spec reference: fulfills `docs/spec/1-account/features/03-login-session-management.md`.

## Out of scope

Fixing anything this suite exposes in earlier tasks' files beyond direct defects found (loop back to the owning task instead of drive-by edits); performance tuning; the deferred X-Forwarded-For limiter work.
