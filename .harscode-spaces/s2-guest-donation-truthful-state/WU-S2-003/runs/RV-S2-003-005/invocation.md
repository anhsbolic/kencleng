# Run Invocation — `RV-S2-003-005`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-04 by Orchestration Operator under Anhar's instruction to continue with an independent migration design review. Human-facing prose: Bahasa Indonesia; retain canonical Harscode terms/enums and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-003`
- `RUN_ID`: `RV-S2-003-005`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-005`
- `ARTIFACT_TARGET`: `none` — write only Run-owned `review-findings-1.md` and `launch-record.md`.
- `WORK_UNIT_PATH` / compatibility `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Independent pre-implementation review of the minimum Donation/D1 schema and migration design proposal
- `PARTICIPANT_ID`: `P-S2-003-RV-005-1`
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; profiles SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`; Registry SHA-256 `804856c5defad721678be4bac145fc2a1070307cccd720326cc598b8159d7abe`.
- `SESSION_TRANSITION`: `FRESH` — independent review after BLD-S2-003-002 ended `STALLED`; do not reuse Build Participant/Session or Orchestrator implementation-reasoning context as review evidence.
- `TARGET_REVISION`: Kencleng HEAD `19d53315ac2847a03405339c0e86df8a97850761` plus current durable working-tree state; read proposal by pinned hash and verify current live authorities/schema anchors.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; ordinary current-effective review guidance.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: The review target is a bounded written schema/migration proposal against settled campaign/donation requirements and a small set of current PostgreSQL anchors. `gpt-6-luna` declares reasoning and repository-work capability; `high` is sufficient for an independent adversarial assessment. The `gpt-6-sol` cross-cutting escalation is not needed for this bounded review, and its approval gate is not bypassed.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Independent migration design review required by TP-S2-003-006 §10 before minimum Donation/D1 schema writes. This is a pre-implementation review of the exact proposal section in a prior Run report, not Review/approval of a code diff, not a re-review of the whole Techplan, and not final WU Testing. Use canonical Review guidance for independent adversarial passes, adapted to this explicit design-artifact target; do not edit the proposal, Techplan, source, tests, SQL, parent projections, or other Run artifacts.

## Exact review target and current-effective inputs

- Sole design target: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/BLD-S2-003-002/report.md`, SHA-256 `dc8f5475b567f692760bb61e7287af7d93dfb4c4a77f7ea8ae9cfd2577a8a4af`, heading `Usulan konkret untuk migration design review dan pairing berikutnya`, especially the schema table and proposed D1 transaction sequence. If the content hash differs, stop and ask Orchestrator to reconcile the new target; do not silently review changed bytes.
- Why this gate applies: Approved `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-006/techplan.md` §10, table anchors for `backend/migrations/000011_create_public_campaigns.up.sql` (Campaign/D1 persistence), requires the minimum accepted D1 schema only after migration design review. The plan SHA-256 is `b71951c86b2f74f6c9f1a164bcf43447c58ee655874deaf09275e6ff3b2316f9`.
- Build outcome: `BLD-S2-003-002/report.md`, SHA-256 `dc8f5475b567f692760bb61e7287af7d93dfb4c4a77f7ea8ae9cfd2577a8a4af`, ended `STALLED` before source/test/migration writes because this prerequisite had no current-effective review receipt. Its proposal is not approved design.
- D1 source decisions: Approved Techplan §§8–10, especially D9 Campaign-first PostgreSQL `READ COMMITTED` plus explicit Campaign-row `SELECT ... FOR UPDATE` before Donation/idempotency rows; accepted D1 close winner/reason, retry ordering, whole-IDR/capacity and exact-once Funding semantics. These remain authority; do not replace them with an alternative.
- Existing schema/migration context: read root `AGENTS.md`, `backend/AGENTS.md`, the exact current `backend/migrations/000011_create_public_campaigns.up.sql` and `000012_add_campaign_max_donation_amount.{up,down}.sql`, current Campaign/Donation domain and repository anchors relevant to the proposed data shape, `docs/project/kencleng-monetary-data-standard.md`, accepted Campaign/Donation invariants/features, and routed current Harscode PostgreSQL migration-safety best practice. Open only relevant sources/sections.
- Independent scope evidence: `RV-S2-003-004` SHA-256 `4ecd9310a4d2c98acb157b5ddc4ab395796b5c076ad1c5ca27f145d198e20eef` and `TST-S2-003-001` SHA-256 `6397951f4b2b232bffd0f776c110d78691bc2ddef21f761e86af329b820d7294` both cover only the eight-file public cap projection and existing `000012`, not the proposed Donation/D1 schema. Do not imply they satisfy this gate.
- Canonical guidance: `../harscode-workspace/workflow/4-code-review-prompt.md`, `workflow/4-code-review/guidelines.md`, `workflow/4-code-review/checklist.md`, `workflow/orchestrated-run-overlay.md`, `workflow/AGENTS.md`, `../harscode-workspace/orchestration/run-contract.md`, `../harscode-workspace/best-practices/AGENTS.md` and its relevant PostgreSQL/migration references. This review uses independent Reviewer context and a pinned artifact target; no production current diff exists to review.

## Review task and boundaries

Assess whether the proposal is technically safe, complete enough, and consistent with TP-S2-003-006 §10 to authorize designing/writing the minimum Donation/D1 schema in a later Build. Analyze, without assuming the proposal is correct:

- migration sequencing after existing unapplied `000012`, additive compatibility with existing rows/defaults, migration locking/validation impact, down-migration data-loss consequences, constraints/indexes/FKs and deletion behavior;
- separation of nullable Campaign `max_amount` from `target_amount`, persisted winning `closed_reason`, and whether the proposed Donation identity/status/exact whole-IDR/idempotency payload-equivalence data can support accepted retry/terminal semantics;
- whether computing reserved capacity from settled Funding plus full accepted-pending rows is coherent with accepted atomic D1 and Campaign-first serialization, including query/constraint/index implications and whether any second counter is justified;
- compatibility with existing `NUMERIC(19,2)` Funding and accepted whole-IDR rules without floats, hidden rounding, invented balance/backfill, or destructive defaults;
- boundaries preventing this design review from choosing unsettled owner authority, security controls, retention or delivery behavior.

Explicit exclusions and gates:

- Do not include Open Item 7 Organization verification/overdue fact source or representative-membership schema/update path; those fields/handlers remain authority-gated.
- Do not invent O3/O4/O5 guest-email, credential key, abuse, retention, proxy/topology or risk-acceptance schema decisions. Identify dependencies, then defer them to their settled authority routes.
- Do not review or change the whole Techplan, product/API semantics, or D1 mechanism decisions. If proposal cannot fit them without material change, report the exact conflict and recommend Planner/Human reconciliation before Build.
- Do not write or propose executable migration SQL as if approved, modify source/tests/specs/Techplan, apply any migration, connect to or mutate a database, or claim PostgreSQL execution/reversibility evidence.
- This is not approval to write Tier-0 code. Anhar's existing exact authorization remains limited to `backend/internal/domain/campaign/donation_coordinator_db.go` and `backend/internal/domain/donation/ledger.go`; active Human pairing remains mandatory if a future Build writes those files.

Review report must distinguish `verified`, `assumed`, `deferred`, and `not tested`. A positive design verdict can satisfy only this migration-design-review precondition for a later schema Build, subject to any Human decision the report identifies; it does not approve migration application, code, Tier-0 implementation, broad Review/Testing, or any milestone. If material unknowns remain, return `Request changes`/blocked recommendation with exact resolution owner; do not infer consent.

## Required artifacts and human-assisted dispatch

Write `RUN_PATH/review-findings-1.md` and `RUN_PATH/launch-record.md`. Use four independent passes (Safety, Quality, Stack-Specific Best Practices, Consistency) against the same pinned design target; the report must include a focused verdict and one structured `## Phase handoff`. Do not run tests or database/migration commands unless a specific review question requires a read-only targeted check; record exact evidence if used. No downstream Run is to be created or dispatched.

Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`. Fresh Reviewer / `KC-REVIEWER` Session; configured `gpt-6-luna` / `high`.

Kickoff: `Lakukan independent migration design review Run RV-S2-003-005 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-005/invocation.md dan canonical ../harscode-workspace/workflow/4-code-review-prompt.md beserta orchestrated-run overlay. Target hanya proposal schema/D1 yang dipin di report BLD-S2-003-002; baca current authority dan migration anchors sesuai Invocation. Uji desain secara adversarial pada kompatibilitas data, whole-IDR, reservation/idempotency, lock-order fit, constraint/FK/index, up/down risk dan migration locking. Bedakan desain proposal dari authority yang sudah approved. Jangan mengedit Techplan/source/tests/SQL, menjalankan/apply migration, memutasi database, atau menganggap prior Review/Testing cap sebagai review D1 schema. Tulis review-findings-1.md dan launch-record dengan satu structured phase handoff. Berhenti dan rute material gap ke Orchestrator/Human.`

This Invocation prepares a Run only; Anhar dispatches the fresh Reviewer. No Participant has been dispatched by the Orchestrator.
