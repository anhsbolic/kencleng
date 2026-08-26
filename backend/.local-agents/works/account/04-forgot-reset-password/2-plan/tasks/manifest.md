# Manifest — techplan decomposition: account/04-forgot-reset-password

> Snapshot at generation time (2026-08-26). This file does NOT track
> execution status — status belongs to the PR/ticket domain, not this
> workspace.

## Source contract

- Originating techplan: `../techplan.md` — "Tech Plan: Forgot & Reset Password (account task #4)", Status: Draft
- The contract sections (Summary, 1–8) stay in that single file and remain the source of truth for human review. Task files redistribute only derived-section detail (§9–§13 material + scoped slices of §3–§7).
- Explore-stage raw material: `../../1-explore/logs/` (00–07 area logs + stage3-solutioning).

## Splitting axis

**Dependency/sequence chain** — chosen because the techplan's §9 execution order is a strict compile-green chain (repo methods → service → transport), and this axis carries the least assumption. Two deviations from pure sequence are deliberate and explicit: Task 02 (notification, different package) and Task 05 (contract edit, no code) have no upstream dependencies and may run in parallel.

## Task files

| File | Title |
|---|---|
| `task-01-repository-foundation.md` | Repository: credential update + user-scoped session revoke (+ integration guard tests) |
| `task-02-notification-sender.md` | Notification platform: `SendPasswordResetEmail` on Sender/Fake/Dev |
| `task-03-service-core.md` | Service: `ForgotPassword`, `ResetPassword`, VerifyEmail purpose check + unit suite |
| `task-04-transport-wiring.md` | Transport: two handlers, routing, rate-limit inheritance + handler tests |
| `task-05-contract-completion.md` | openapi: add documented 429 to reset-password + bundle regen |
| `task-06-integration-race-suite.md` | Integration & race: INV-05 atomicity property, INV-08 ≥100-goroutine stress, timing parity |

## Dependency graph

```
task-01 (repository) ──┐
                       ├──> task-03 (service) ──> task-04 (transport)
task-02 (notification)─┘              │
                                      └─────> task-06 (integration/race)
task-05 (contract) — no hard dependency (parallel-safe with all)
```

- task-01 → task-03: service calls the two new repository methods
- task-02 → task-03: service calls the new Sender method
- task-03 → task-04: handlers call the service methods
- task-03 → task-06: integration suite exercises the flows against real Postgres
- task-05: independent; should land within the same merge window as task-04 so spec and handlers ship together (spec-first discipline)

## Model routing per task

Tier justification: **Complex** per `harscode-workspace/best-practices/model-routing.md` — 18 rules (≥15 threshold) + auth-contract surface. Build stage, decomposed → route per sub-task; GLM when the work leans multi-step/stateful reasoning, DeepSeek V4 Pro when it's rule-table-heavy precision without a diagram.

| Task | Model | Why |
|---|---|---|
| task-01 repository | DeepSeek V4 Pro | Precision goqu/guard SQL work, no branching judgment |
| task-02 notification | DeepSeek V4 Pro | Mechanical interface addition + compile ripple; low judgment |
| task-03 service core | GLM 5.2 (max) | The hard one: tx ordering semantics, anti-enumeration branch logic, purpose-check rollback reasoning — multi-step stateful reasoning |
| task-04 transport | DeepSeek V4 Pro | Pattern-cloning from existing handlers; precision over invention |
| task-05 contract | DeepSeek V4 Flash | Pure mechanical spec edit + bundler run (Simple-tier chunk inside a Complex plan); bundle-diff review still human-checked |
| task-06 integration/race | GLM 5.2 (max) | Concurrency harness + invariant property design = subtle multi-step reasoning; note the separate 5-testing phase later routes to DeepSeek V4 Pro for rule-ID count-checking |

Non-negotiable downstream rows (not build tasks, listed for the pipeline): code review at this tier runs **GLM 5.2 (max) + DeepSeek V4 Pro parallel, diffed manually**; testing phase routes **DeepSeek V4 Pro** (rule count ≥15).

## Guardrails carried into every task

- Contract authority: `../techplan.md` §1–8 wins over any task file; contradictions get flagged, never silently resolved.
- No detail compression: each task file holds full detail for its scope; nothing was summarized away.
- Fenced paths untouched by all tasks: `internal/platform/auth/*`, `internal/platform/crypto/*`, `docs/spec/**`.
- Open items from techplan §14 Active (URI prefix split, tasks.md staleness, AGENTS.md outbox wording, rate-limit values TBD-verify) are re-stated inside the tasks they touch.
