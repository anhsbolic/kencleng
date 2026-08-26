# Manifest — 05-account-linking task decomposition

> Generated : 2026-08-26
> Source    : `.local-agents/works/account/05-account-linking/2-plan/techplan.md` (Status: Approved)
> Note      : snapshot at generation time — this file does NOT track progress status; status belongs to the PR/ticket domain.

## Task list

| # | File | Title |
|---|---|---|
| 1 | `task-01-migration-and-contract-bundle.md` | Migration 000010 (purpose CHECK widen, reversible) + stale `api/openapi.yaml` bundle regeneration |
| 2 | `task-02-repository-foundation.md` | Four new repository operations (identities-by-user, locked finder variant, delete-by-ids, credential-secret update, revoke-all-refresh-for-user) |
| 3 | `task-03-set-password-service.md` | `SetPassword` service: server-side branch selection, Branch 1 anti-enumeration flow, Branch 2 atomic change+revoke; `VerifyEmail` conditional audit delta; nudge constant |
| 4 | `task-04-unlink-service.md` | `UnlinkGoogle` service: FOR UPDATE guard classification, re-auth ordering, hard delete + audit; ≥100-goroutine concurrency stress harness |
| 5 | `task-05-transport-wiring.md` | Session middleware (`requireSession` + inline ES256 verifier), two handlers, 409 sentinel mappings, `main.go` route group |
| 6 | `task-06-integration-gate.md` | testcontainers DB-truth suite + full `make verify` gate + KPI checklist + build report |

## Splitting axis

**Dependency/sequence chain** (layer-aligned). Rationale: every stage genuinely waits on the previous — migration/reversibility before app-path inserts of the new purpose value; repository methods before both service flows; both flows before handlers; everything before the integration gate. A component/module split is impossible here (one domain package, one shared `security.go` file forces tasks 03–04 serial anyway), and a risk axis adds nothing the sequence doesn't already isolate (the high-risk concurrency work is its own task boundary).

## Dependency graph

```
01 ──► 02 ──► 03 ──► 05 ──► 06
              ▲
              04 ─────────────┘
(03 → 04 are serial by choice: shared security.go file;
 both depend only on 02. 05 depends on 03+04. 06 depends on all.)
```

No parallel execution recommended despite 03/04 being logically independent — same-file ownership.

## Model routing per task

Overall tier: **Complex** (16 rules ≥15 threshold; touches auth/session surface). Per-task picks from `best-practices/model-routing.md`:

| Task | Model | Why |
|---|---|---|
| 01 | DeepSeek V4 Pro | Precision DDL + reversibility subtlety; matches task-03's migrations precedent |
| 02 | DeepSeek V4 Pro | Rule-table-heavy precision: goqu parameterization golden rule, nullable scans |
| 03 | GLM 5.2 (max) | Multi-step branching + timing-parity reasoning (GLM's strength); reward-hacking caveat → covered by mandatory dual-model review |
| 04 | GLM 5.2 (max) | Concurrency-invariant design = highest-risk reasoning in the slice |
| 05 | DeepSeek V4 Pro | Contract-shape precision: exact status codes, problem-type URIs, verbatim Indonesian details |
| 06 | DeepSeek V4 Pro | Testing/Complex row verbatim: rule-ID count-check reliability once rules ≥15 |

Cross-task non-negotiables from model-routing: this is Complex-tier work, so **code review must run GLM 5.2 (max) + DeepSeek V4 Pro in parallel with manual diffing** (no exception), and any Claude fallback follows that doc's mapping table.

## Back-reference

Originating contract techplan: `.local-agents/works/account/05-account-linking/2-plan/techplan.md` — "Tech Plan: Account Linking (account #05)". Every task file carries this back-reference; executors cross-check decisions there (D1–D9, R1–R16, §8 flow blocks are authoritative over any paraphrase).

## Review note

These derived task files remain subject to `workflow/4-code-review/checklist.md` like all derived-section content.
