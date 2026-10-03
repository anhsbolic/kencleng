# Run Invocation — `BLD-S2-003-001`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-03 by Orchestration Operator after completed TPD-S2-003-001. Human-facing prose: Bahasa Indonesia; retain canonical Harscode terms/enums and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-003`
- `RUN_ID`: `BLD-S2-003-001`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/BLD-S2-003-001`
- `ARTIFACT_TARGET`: `none` — backend source changes are governed by the approved Techplan; Run-owned evidence is `RUN_PATH/report.md` and `RUN_PATH/launch-record.md`.
- `WORK_UNIT_PATH` / compatibility `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Implementer
- `SPECIALIZATION`: Approved Slice-2 Donation backend delivery
- `PARTICIPANT_ID`: `P-S2-003-BLD-001-1`
- `PARTICIPANT_PROFILE_ID`: `KC-IMPLEMENTER`; `profiles.md` SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`; Registry SHA-256 `804856c5defad721678be4bac145fc2a1070307cccd720326cc598b8159d7abe`.
- `SESSION_TRANSITION`: `FRESH` — new Build phase after Planner decomposition completed; do not reuse TPD/Techplan Participant or Session context.
- `TARGET_REVISION`: Kencleng HEAD `19d53315ac2847a03405339c0e86df8a97850761` plus current durable working-tree changes; re-ground approved authority and live code before editing.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; ordinary applicable phase guidance remains current-effective.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: The Approved spine is complex backend repository work, but Product, contract, D1 ordering, monetary behavior and the scoped authority boundaries are settled. The Human-declared coding/repository-work capability at high effort is sufficient for implementation and focused verification. Do not escalate to approval-gated `gpt-6-sol` for missing authority, stale context or environment constraints; stop and route those causes. Escalation is appropriate only if a concrete capability insufficiency remains after those causes are ruled out.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Initial Build/Patch of the complete Approved Techplan `TP-S2-003-006`. TPD-S2-003-001 completed canonical Step 0 with `NO`; no task files exist and no split-review gate applies. Do not invent an ad-hoc task boundary.

## Current-effective inputs / `PRIOR_ARTIFACTS`

- Sole execution spine: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-006/techplan.md`, `Approved`, SHA-256 `b71951c86b2f74f6c9f1a164bcf43447c58ee655874deaf09275e6ff3b2316f9`.
- Approval provenance: report-only `TP-S2-003-007/report-techplan.md`, SHA-256 `5fa5ef51309d8112f7725c0c185d4e9f8e94d353d6229853433ff45dc0c8e5ed`; Human approved the exact pre-reconciliation Draft hash `ca49da69784b40d762d96c4b523f4603019c4c2568e3b8445c1e97a415898095`; status-only reconciliation verified the approved plan hash above. Whole-plan approval preserves, and does not clear, all scoped gates.
- Completed decomposition evidence: `TPD-S2-003-001/launch-record.md`, SHA-256 `0be60c5bc9c11bf4526116234abb5069e461db540a457e7928190b54d637c97b`; exact Approved Techplan assessed; STEP 0 `NO`; no task files/manifest were generated. Configured model/effort `gpt-6-luna` / `high`; active runtime/session values were not independently exposed.
- Independent Review: `RV-S2-003-003/review-findings.md`, SHA-256 `b5cd4d47e714b8499ef0ae8d33062283763812d0a18cc7999f1f14d6f4e2a79d`; clean against pre-reconciliation Techplan hash `d98b3d37e3c00962075428c3e14f87d5ee59dfba46b6ae36ce35f78dcebbce7a`. Do not claim it reviewed the current approved bytes.
- Current coordination state: WU-S2-003 manifest, parent Events/Work Graph/Outcome/Control Surface and `docs/project/kencleng-development-tracker.md`; these record Open Item 7's precise boundary, protected Tier-0 gates and downstream evidence ownership.
- Read root `AGENTS.md` and `backend/AGENTS.md` before edits. Reopen current Product/MVP scope and sequencing, approved Techplan anchors, applicable accepted Campaign/Donation specs and OpenAPI sources, monetary standard, backend architecture, project workflow risk/human gates, and current live backend code/migrations/tests as the Plan requires. Read only relevant sections/anchors; do not load raw Exploration logs by default.
- Canonical guidance: `../harscode-workspace/workflow/3-build-prompt.md`, `workflow/3-build/guidelines.md`, `workflow/3-build/checklist.md`, `workflow/orchestrated-run-overlay.md`, `workflow/AGENTS.md`, `workflow/context-management.md`, `../harscode-workspace/orchestration/AGENTS.md`, and `orchestration/run-contract.md`.

## Build target and completion condition

Execute the complete Approved TP-S2-003-006 spine, not a new scope proposal. Reopen live code/spec at the plan's code anchors before editing. Implement only behavior and paths authorized by the plan and current authorities: guest Donation submission/status and simulator-owned lifecycle; Campaign-owned availability, D1 ordering/close and funding-capacity coordination; accepted monetary cap projection and exact-wire counterpart; accepted backend Campaign draft create/PATCH and minimum persisted eligibility integration only when Open Item 7 is resolved; required persistence/migrations, error/SQL/security controls, and tests/evidence within the Plan's gates.

The Build target is the whole Approved spine because decomposition found no useful independent task. Apply the plan's gates to each action. Do not silently replace a blocked plan element with a new task boundary or claim the full target complete if required work remains blocked. If the approved cross-domain design cannot be implemented safely without a currently gated action, stop at that boundary and give the Orchestrator an exact blocker, affected scope, evidence, remaining safe work if any, and next owner. Continue only with work that remains coherent and valid under the complete spine and does not create an unsafe partial seam.

## Execution envelope

- `PREAUTHORIZED`: read the assigned durable authorities and live code; execute approved backend changes within `backend/` and the exact paths required by TP-S2-003-006; write only this Run's `report.md` and `launch-record.md`; run focused Build-loop verification appropriate to actual edits and the approved Testing Checklist. Tests authored/changed as part of this implementation must be run enough to show they execute and exercise their intended state. Do not edit `frontend/`, `api/`, Product/spec authority, Techplan, other Runs or Work Units, or orchestration projections.
- `HUMAN_REQUIRED` — Open Item 7: do not create/update the Organization `verified` / `has_overdue_report` truth source or implement/enable affected Campaign draft create/PATCH handlers until authority and governance are resolved. Seeded/operator-assisted setup is a data posture, not permission for an unowned setter path. The blocker is scoped to these source updates/handlers, not unrelated Donation work.
- `HUMAN_REQUIRED` — Tier-0: do not write `backend/internal/domain/donation/ledger.go`, `backend/internal/domain/campaign/donation_coordinator_db.go`, or any file implementing balance transaction/locking logic without exact file/scope Human authorization and the required Human-paired process. Plan approval and D1 mechanism selection are not write authorization. Do not route around a protected implementation by duplicating equivalent transaction/lock behavior in an unprotected file.
- `HUMAN_REQUIRED`: do not manually apply migrations/indexes; do not modify protected crypto/auth/disbursement state-machine files; do not resolve O3/O4/O5 security, provider, retention, abuse, proxy/parity or residual-risk owner decisions; do not claim their acceptance or runtime proof. Stop/report if the implementation requires any such gated action or if live code invalidates a material approved assumption.
- Build verification must follow `3-build-prompt.md` and the Approved Testing Checklist. Run focused checks needed to make edits credible, including tests added/changed; leave independent PostgreSQL/concurrency, broad race/performance/security, topology/runtime, residual-risk, and Human rendered evidence to their assigned owners unless the plan/repo expressly requires a focused check in this Build. Record exact commands/results and explicitly state what was deferred or not tested. Never claim checks not run.
- Do not mark `BACKEND_VERIFIED`, complete WU-S2-003, accept risk, self-approve, apply migrations, or infer integration/runtime readiness.

## Required Build report and handoff

Write `RUN_PATH/report.md` and `RUN_PATH/launch-record.md`. The report must follow the canonical Build format, include compact known provenance and exactly one structured `## Phase handoff` with fields: Outcome, Result refs, Findings, Decision requests, Blockers, Open / unverified, Recommended continuation, Context refs. It must state changed behavior/files, actual checks, verification scope confirmation for race/concurrency/performance/security-class tests, deferred/not-tested items, plan-target completion status, every blocker and safe unaffected work, and why the Run ended. Do not create or dispatch downstream Runs.

## Human-assisted dispatch

Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`. Fresh Implementer / `KC-IMPLEMENTER` Session; configured `gpt-6-luna` / `high`.

Kickoff: `Jalankan Build/Patch Run BLD-S2-003-001 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/BLD-S2-003-001/invocation.md dan canonical ../harscode-workspace/workflow/3-build-prompt.md beserta orchestrated-run overlay. Jalankan seluruh Approved Techplan sebagai target karena TPD-S2-003-001 Step 0 = NO dan tidak ada task split. Re-ground current authority/live code, patuhi seluruh execution envelope dan scoped gates, lakukan focused Build-loop verification sesuai canonical guidance, tulis report lengkap dengan structured phase handoff dan launch-record ke RUN_PATH. Jangan mengarang task boundary, melewati gate, menerapkan migration, atau mengklaim milestone; berhenti dan laporkan tepat bila kontradiksi/gate menghalangi safe progress.`

Human reports completion or the exact blocker to the Orchestrator. This Invocation prepares a Run only; it does not itself dispatch a Participant.
