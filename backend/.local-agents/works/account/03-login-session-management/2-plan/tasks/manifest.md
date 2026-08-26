# Manifest — Login & Session Management decomposition

> Generated   : 2026-08-26
> Snapshot    : generation-time index, NOT a progress tracker (status lives in PRs/`tasks.md`)
> Source      : `.local-agents/works/account/03-login-session-management/2-plan/techplan.md` (Status: Approved, 2026-08-26)
> Feature spec: `docs/spec/1-account/features/03-login-session-management.md`

## Task files

| File | Title |
|---|---|
| `task-01-migrations-and-schema.md` | Migrations 000006–000009 (schema-pre-settle) |
| `task-02-domain-repository-foundation.md` | Domain entities + Repository port & adapter |
| `task-03-tier0-token-helpers.md` | Tier 0 — JWT mint/verify primitives (`platform/auth/token.go`) |
| `task-04-login-session-service.md` | Login/session domain services |
| `task-05-transport-and-wiring.md` | HTTP handlers, cookies, error mapping, wiring |
| `task-06-integration-race-suite.md` | Integration + race suite, final gate |

## Splitting axis

**Dependency/sequence chain.** Hard dependencies dominate: migrations → repo methods → services → handlers → wiring → proof suite. Matches the techplan's own execution order (§9) and `tasks.md`'s S1 serial-group rationale (shared tables). No parallel component boundaries exist (single domain package). The one risk-axis concern — isolating the Tier 0 fenced sub-area — is satisfied by the chain itself: Tier 0 work lands in task-03 plus flagged regions of tasks 02/04.

## Dependency graph

```
task-01 ──► task-02 ──► task-04 ──► task-05 ──► task-06
                ▲         ▲
task-03 ────────┘─────────┘   (independent of 01/02; must precede 04)
```

Linear execution `01 → 02 → 03 → 04 → 05 → 06` is the mandated order (`tasks.md` S1 is serial; migration numbering is shared state). task-03 has no technical dependency on 01/02 but nothing downstream of it may start first.

## Model routing

Techplan tier: **Complex** (20 rules ≥ 15; auth contract surface) per `best-practices/model-routing.md`. Decomposition executed under the Complex-tier routing (MiMo V2.5 Pro — Step 0 gate justified).

| Task | Build model | Rationale from routing table |
|---|---|---|
| 01 migrations | DeepSeek V4 Pro | Rule-table-heavy precision work, no diagram judgment |
| 02 repository | DeepSeek V4 Pro | goqu/invariant precision ("ties Claude on SWE-bench Verified" profile) |
| 03 Tier 0 tokens | GLM 5.2 (max) | Multi-step reasoning; **non-negotiable**: human paired rewrite pass regardless of model (Resolved #13) |
| 04 services | GLM 5.2 (max) | Branching flow / state-transition reasoning (GLM's lean-per-routing) |
| 05 transport | DeepSeek V4 Pro | Byte-equal contract precision without diagram work |
| 06 race/integration | DeepSeek V4 Pro | Testing row, Complex tier — Flash unreliable for rule-count ≥15 checks |

Downstream reminder: code review at this tier = **GLM 5.2 (max) + DeepSeek V4 Pro in parallel, diff manually** (non-negotiable dual-model row).

## ⚠️ Tier 0 paired-pass checklist (before ANY commit — Resolved #13)

Agent drafts with heavy doc-comments + exhaustive tests; Anhar's dedicated paired rewrite/review covers exactly:

1. `internal/platform/auth/token.go` — both token purposes' mint/verify (task-03 output)
2. `internal/domain/account/repository_db.go` — rotation methods (`RotateRefreshToken`, `RevokeRefreshTokenFamily`) (task-02 region)
3. `internal/domain/account/login.go` — reuse/race-loser branch in `Refresh` (task-04 region)

The build report must carry an explicit "Tier 0 files awaiting paired rewrite" flag list. Nothing Tier 0 commits without this pass (tasks.md KPI boolean gate).

## Cross-references

- Contract techplan: `.local-agents/works/account/03-login-session-management/2-plan/techplan.md` — every task file back-references it; cross-check high-level decisions there whenever a task seems ambiguous.
- Exploration raw logs: `.local-agents/works/account/03-login-session-management/1-explore/logs/`
- Review checklist: derived content here remains subject to `workflow/4-code-review/checklist.md`
- Rule-ID coverage: R1–R20 distributed as — T3: R6,R17 · T4: R1–R5,R7–R11,R13–R16,R18,R19 · T5: transport halves of R1–R4,R7,R8,R10,R11,R13,R16 + R19 sweep · T6: R12,R14,R15 proofs + R20 round-trip · T1: R20 additive/reversible · T2: R12/R14/R16/R19 halves, R20. Reconcile at task-06 gate.
