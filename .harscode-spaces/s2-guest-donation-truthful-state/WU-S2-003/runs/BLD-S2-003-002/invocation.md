# Run Invocation — `BLD-S2-003-002`

Status: `DISPATCHED — HUMAN REPORTED`

Prepared: 2026-10-03 by Orchestration Operator after TST-S2-003-001 and explicit Human Tier-0 authorization. Human-facing prose: Bahasa Indonesia; retain canonical Harscode terms/enums and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-003`
- `RUN_ID`: `BLD-S2-003-002`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/BLD-S2-003-002`
- `ARTIFACT_TARGET`: `none` — backend source changes follow the Approved Techplan; Run-owned evidence is `RUN_PATH/report.md` and `RUN_PATH/launch-record.md`.
- `WORK_UNIT_PATH` / compatibility `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Implementer
- `SPECIALIZATION`: Human-paired Tier-0 D1 transaction/locking continuation within approved Slice-2 Donation backend delivery
- `PARTICIPANT_ID`: `P-S2-003-BLD-002-1`
- `PARTICIPANT_PROFILE_ID`: `KC-IMPLEMENTER`; profiles SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`; Registry SHA-256 `804856c5defad721678be4bac145fc2a1070307cccd720326cc598b8159d7abe`.
- `SESSION_TRANSITION`: `FRESH` — new Build/Patch occurrence after BLD-S2-003-001 ended `STALLED` and its bounded slice completed independent Review and Testing. Do not reuse prior Participant or Session.
- `TARGET_REVISION`: Kencleng HEAD `19d53315ac2847a03405339c0e86df8a97850761` plus current durable working-tree changes; inspect live sources and diff before editing.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; ordinary current-effective phase guidance.
- `SELECTED_MODEL`: `gpt-6-sol`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Anhar explicitly approved `gpt-6-sol` / `high` for this Participant Run on 2026-10-03. Local runtime registry marks `gpt-6-sol` `approval_required: true`; the explicit Run-scoped Human approval is recorded in parent `events.md`.
- `MODEL_ROUTING_RATIONALE`: This re-entry implements the settled Campaign-first D1 admission/close ordering and atomic exact-once Funding path across Campaign and Donation while touching two Tier-0 transaction/locking files. Human authorization and active Human pairing satisfy separate governance boundaries but do not lower the implementation reasoning need. The local registry's strong-reasoning, architecture, cross-cutting-analysis and repository-work capabilities at `high` are the best fit. Anhar approved this model/effort before dispatch. The separate orchestration-pairing model remains at its configured default.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Fresh Build/Patch re-entry against the complete Approved Techplan `TP-S2-003-006`. `TPD-S2-003-001` canonical Step 0 was `NO`; no task files exist and no split-review gate applies. Continue from the exact prior Build stop while reconstructing from durable artifacts and current live code. Do not invent a D1-only task boundary or claim the whole spine complete if blocked/deferred scope remains.

## Current-effective inputs / `PRIOR_ARTIFACTS`

- Sole execution spine: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-006/techplan.md`, `Approved`, SHA-256 `b71951c86b2f74f6c9f1a164bcf43447c58ee655874deaf09275e6ff3b2316f9`. Its §13.5 stale Draft/In Review sentence is documented in the WU003 manifest; the Approved header and approval provenance are controlling. Do not edit the Techplan.
- Approval provenance: Human-approved exact Draft preimage hash `ca49da69784b40d762d96c4b523f4603019c4c2568e3b8445c1e97a415898095`, after report TP-S2-003-007; approved artifact hash is above. This plan approval did not authorize Tier-0 paths by itself.
- Prior Build: `runs/BLD-S2-003-001/report.md`, SHA-256 `6800f6433b32339cc9829aa97b2987e19d6bc2882cda16dcec51dd9237a4b32a`; outcome `STALLED` after the public Campaign cap projection/exact-wire slice and unapplied migration. Reconstruct the exact stop and named D1 blocker; do not inherit its Session state.
- Independent Review: `runs/RV-S2-003-004/review-findings-1.md`, SHA-256 `4ecd9310a4d2c98acb157b5ddc4ab395796b5c076ad1c5ca27f145d198e20eef`; `Approve`, no findings, only for the eight-file cap projection/migration diff.
- Independent Testing: `runs/TST-S2-003-001/testing-report-1.md`, SHA-256 `6397951f4b2b232bffd0f776c110d78691bc2ddef21f761e86af329b820d7294`; `Pass with flagged follow-ups`, only for the public Campaign cap-projection slice. Cached and uncached focused Campaign/domain HTTP commands passed. PostgreSQL migration, full-spine behavior, concurrency, security/runtime and other evidence remain open.
- Exact Human authorization: Anhar explicitly authorized implementation of the Approved D1 behavior on `backend/internal/domain/campaign/donation_coordinator_db.go` and `backend/internal/domain/donation/ledger.go` on 2026-10-03 and stated readiness for Human-paired process. This authorization is limited to those two paths and D1 scope described by TP-S2-003-006. It does not authorize equivalent transaction/locking implementation in other files, migration application, Open Item 7 work, crypto/auth/disbursement changes, or risk acceptance.
- Current coordination truth: WU-S2-003 manifest, parent events/work-graph/outcome/control-surface and `docs/project/kencleng-development-tracker.md`.
- Before editing, read root `AGENTS.md` and `backend/AGENTS.md`; relevant Product/MVP boundaries, approved Techplan §§8–13, accepted Campaign/Donation specs/invariants/threat-model, authored OpenAPI sources and referenced shared components, monetary standard, backend architecture, applicable Harscode Build rules, and current live Go/migration/test code at the plan anchors. Do not load raw Exploration logs by default.
- Canonical guidance: `../harscode-workspace/workflow/3-build-prompt.md`, `workflow/3-build/guidelines.md`, `workflow/3-build/checklist.md`, `workflow/orchestrated-run-overlay.md`, `workflow/AGENTS.md`, `workflow/context-management.md`, `../harscode-workspace/orchestration/AGENTS.md`, `orchestration/run-contract.md`, and `orchestration/local-runtime-config.md`. Project risk/Human gate routing: root `AGENTS.md` §§3, 7, 9 and `docs/kencleng-agentic-workflow.md` §4.

## Build target and execution envelope

Execute the complete Approved TP-S2-003-006 spine, with focus on safely completing the D1 ordering and atomic Funding work that blocked BLD-S2-003-001. Re-ground every relevant live-code assumption first. Preserve the approved Campaign-first transaction/lock order, close winner/reason semantics, idempotency ordering, accepted-pending capacity reservation and exact-once successful Funding reflection. Implement only the approved plan and current authorities; do not redesign product or concurrency policy. Continue other whole-plan work only where it is coherent and unblocked; stop/report material gates without inventing partial task slices.

### Exact Tier-0 authorization and Human-paired process

- Anhar has explicitly authorized D1 implementation writes to exactly `backend/internal/domain/campaign/donation_coordinator_db.go` and `backend/internal/domain/donation/ledger.go`, within approved TP-S2-003-006 behavior.
- These writes remain **Human-paired**. Anhar must be actively present for the implementation of both files: discuss/re-ground the concrete transaction/locking change against the approved D1 plan, co-review the proposed code as it is developed, and inspect the resulting diff. Participant must pause and ask for the Human pair if the required presence/review is unavailable; authorization is not permission for an unobserved autonomous Tier-0 write.
- Do not write another file that implements or relocates balance transaction/locking logic. If live code demonstrates that the authorized change requires another protected path, stop before that write and return the exact path/scope for new Human authorization.
- Record in `report.md` and `launch-record.md` how the required pairing was performed; do not claim Tier-0 completion solely from the authorization message.

### Execution envelope

- `PREAUTHORIZED`: read current authority/live code; make ordinary backend Build changes within the complete approved Techplan; make D1 Tier-0 changes only in the two explicitly authorized files and only in the actively Human-paired process above; edit/add focused backend tests as Techplan allows; write only this Run's `report.md` and `launch-record.md`; run focused Build-loop verification appropriate to actual edits.
- `HUMAN_REQUIRED` — Open Item 7: do not create/update Organization `verified` / `has_overdue_report` truth source or implement/enable affected Campaign draft create/PATCH handlers until authority and governance are resolved. The gate does not block unrelated Donation work.
- `HUMAN_REQUIRED`: do not manually apply migrations/indexes; modify protected crypto/auth/disbursement state-machine paths; decide O3/O4/O5 owners' controls; accept residual risk; claim runtime/topology or Human-rendered evidence. Stop if implementation requires any such action or live evidence materially contradicts the approved plan.
- Keep all frontend/API/spec/Product authority files outside this backend Build scope. Do not edit Techplan, tracker, parent projections, other Runs, or other WUs from the Participant Run.
- Run focused Build-loop checks for code/tests changed. Any new/changed automated test must be executed sufficiently to show it runs and asserts intended behavior. Do not run broad race/performance/security-class or full PostgreSQL Testing-owned suites merely for confidence; route assigned independent evidence to Testing. Record exact commands/results and state that concurrency/race, performance/load and security-class tests were not run unless they were.
- Preserve existing unapplied migration `000012_add_campaign_max_donation_amount`; do not apply any migration or modify local/shared database state. Surface migration sequence/schema needs through code/artifact evidence only.
- Do not claim `BACKEND_VERIFIED`, whole WU completion, runtime/integration readiness, or risk acceptance. If the complete plan is not done, give `STALLED`/partial outcome with exact remaining work and gates.

## Required Build report and handoff

Write `RUN_PATH/report.md` and `RUN_PATH/launch-record.md` using current canonical formats. Report the full approved-plan target status, files/behavior changed, focused checks, explicit verification boundaries, Human-pair evidence, any blockers and safe remaining work. Include exactly one structured `## Phase handoff` with `Outcome`, `Result refs`, `Findings`, `Decision requests`, `Blockers`, `Open / unverified`, `Recommended continuation`, and `Context refs`. Do not create or dispatch downstream Runs.

## Human-assisted dispatch

Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`. Fresh Implementer / `KC-IMPLEMENTER` Session. Selected model: `gpt-6-sol` / `high`; Anhar approved and reports dispatch. Participant must confirm active Human pairing before writing either Tier-0 file and stop if it is unavailable.

Kickoff: `Jalankan Build/Patch Run BLD-S2-003-002 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/BLD-S2-003-002/invocation.md dan canonical ../harscode-workspace/workflow/3-build-prompt.md beserta orchestrated-run overlay. Jalankan seluruh Approved Techplan TP-S2-003-006 sebagai target karena TPD-S2-003-001 Step 0 = NO dan tidak ada task split. Lanjutkan dari blocker D1 pada Run sebelumnya dengan fresh Participant/Session; re-ground seluruh live authority/code. Anhar telah mengotorisasi D1 Tier-0 writes hanya pada dua file bernama di Invocation dan harus hadir dalam Human-paired implementation/review. Ikuti Open Item 7, migration, security/runtime dan seluruh scoped gates. Lakukan focused Build-loop verification, tulis report lengkap dan launch-record ke RUN_PATH. Jangan mengarang task boundary, melakukan autonomous Tier-0 write, apply migration, mengedit authority/projections, atau mengklaim milestone; berhenti tepat pada gate/kontradiksi.`

Dispatch status: Anhar reported on 2026-10-03 that `BLD-S2-003-002` was dispatched after approving the selected model/effort. Participant Session/runtime identifiers and implementation/pairing progress are not yet exposed; reconstruct those only from the Participant's launch record/report or later Human evidence.

This Invocation records the assignment and authorization. It is not evidence that implementation completed, that the Human-paired process was performed, or that any check ran.
