# Run Invocation — `TPD-S2-002-002`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-01 by Orchestration Operator. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `TPD-S2-002-002`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-002`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Planner
- `SPECIALIZATION`: Post-approval task-snapshot reconciliation for the material successor Techplan TP-015
- `PARTICIPANT_ID`: `P-S2-002-PLD-002-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-PLANNER`; profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — new Planner occurrence after status reconciliation Run TP-S2-002-017; do not reuse earlier Planner/Reviewer Participant or Session context.
- `TARGET_REVISION`: Kencleng checkout `650e73c5d646c29c0ddf1931618f02685d15f7b7` plus TP-015 approval/status reconciliation and current durable working-tree decisions/projections; re-read live sources at dispatch.
- `WORKFLOW_REVISION`: Harscode checkout `33b03a3f62cc3aacba6534b8a011465613c64b09`; re-read current canonical decomposition guidance at dispatch.
- `MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by `.harscode-spaces/.local-config.yaml` (`approval_required: false`).
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Required fresh post-approval reconciliation of child task snapshots affected by TP-015. Preserve the already Human-accepted Task 01 → Task 02 split, dependency, and manifest unless review of the Approved spine proves a material topology/dependency change.
- `MODEL_ROUTING_RATIONALE`: The Planner must preserve a previously accepted dependency graph while faithfully refreshing two execution snapshots and shared authority/ID references against a material cross-domain plan. Configured `gpt-6-luna` with `high` effort is suitable; missing authority or a real topology change is routed to the Human/Orchestrator, not guessed.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/techplan.md` — current-effective Human-approved `Approved` Techplan; complete authoritative spine and sole parent plan for these refreshed snapshots.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-016/report-techplan.md` — Human-facing approval report; use only for approval/review provenance, not as a second contract.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-017/launch-record.md` and `.harscode-spaces/s2-guest-donation-truthful-state/events.md` — verified explicit Human approval, status-only reconciliation, and current-effective pointer provenance.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-001/tasks/manifest.md`, `01-donation-domain-spec-reconciliation.md`, and `02-donation-openapi-reconciliation.md` — existing Human-accepted decomposition shape and task snapshots to reconcile; preserve original Run files unchanged.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-001/launch-record.md` and `.harscode-spaces/s2-guest-donation-truthful-state/events.md` — original Step 0/split rationale, accepted Task 01 → Task 02 order/dependency, and Human split acceptance.
- Task 01 completion/review history: `BLD-S2-002-001/report.md`, `BLD-S2-002-002/patch-report-1.md`, `RV-S2-002-007/review-findings.md`, and `RV-S2-002-008/review-confirmation.md`. Task 01 Human acceptance of the actual Donation spec drafts remains a separate parallel work item.
- Task 02 history: `BLD-S2-002-003/report.md`; no authored API diff was produced there. Reconcile its task snapshot with TP-015, but do not edit/validate OpenAPI in this Run.
- Current Work Unit state: WU-S2-002 manifest, `../work-graph.md`, `../events.md`, `../control-surface.md`, `../outcome.md`, `docs/project/kencleng-development-tracker.md`, and `.harscode-spaces/authority-map.md`.
- Relevant target routing: root `AGENTS.md`, `backend/AGENTS.md` only when needed, `docs/spec/README.md`, `api/README.md`, current Product/MVP and monetary authority sources cited by TP-015. Do not elevate stale child snapshots above TP-015.
- Current Harscode guidance at dispatch: canonical `../harscode-workspace/workflow/2-3-techplan-decomposition-prompt.md`; read `workflow/2-techplan/rules.md` §10, `workflow/orchestrated-run-overlay.md`, and applicable `workflow/AGENTS.md`. Apply the prompt's Step 0 and decomposition invariants to this narrowly scoped post-approval snapshot reconciliation.
- Profile/runtime routing sources: `.harscode-spaces/participant-profiles/profiles.md#kc-planner` and `.harscode-spaces/.local-config.yaml`.

## Task and completion condition

Reconcile the task snapshots produced under TPD-001 to current-effective Approved TP-015 after the required Human approval/status reconciliation. This is content reconciliation under the existing accepted split, not a new request to redesign or reapprove the task topology.

First state the canonical decomposition Step 0 gate and the execution/review value of preserving the accepted domain-spec → authored-OpenAPI boundary. Preserve exactly:

- Task 01: Donation domain-spec reconciliation; no hard task dependency.
- Task 02: authored Donation OpenAPI reconciliation; hard dependency on Task 01.
- Existing task purpose, scope axis, order, and accepted manifest topology, unless a material contradiction in current Approved TP-015 makes preservation impossible. If a material topology/dependency change is genuinely required, stop and route that split for Human review; do not silently change it.

Reconcile only affected child snapshots into this Run's `RUN_PATH/tasks/` and write the updated manifest last. The current children are stale relative to the approved spine. Task 01 must carry any applicable O1 shared amount representation and O11 terminal-notice/terminalization semantics, current TP-015 parent path/status, and up-to-date parent requirement/rule/decision/risk/verification/Open Item references. Task 02 must additionally carry bounded O8 replacement clearance and the settled O4 fragment/frontend-handoff/URL-cleanup/one-way-HMAC direction plus O5 uniform `404`, identical body/header/cache behavior and `Cache-Control: private, no-store`; distinguish authored contract detail from downstream security/runtime evidence. Re-evaluate precise changed references against TP-015 in full; do not assume these examples enumerate every stale passage.

Keep TP-015 as the complete authority. Preserve all unaffected task instructions, the existing accepted dependency graph, and the Human-accepted split; do not add tasks or make a second source of truth. Update the refreshed task and manifest parent pointers/status/IDs to TP-015's current plan only where required for accuracy. Preserve the accepted task topology/dependency and the existing task manifest's execution map. Keep Task 01 Human acceptance of the current domain-spec drafts as an independent parallel item. Keep Task 02 hard dependency and the already-recorded Task 01 Build/review history visible; do not claim that current spec acceptance or `CONTRACT_READY` is complete.

If a material contract/security/verification decision needed to scope either task is absent from Approved TP-015, stop and report the exact gap; do not amend the Techplan or solve the gap in a child task. Do not regenerate report TP-016, edit source specs/OpenAPI/code/tests, modify prior Run files, or change orchestration projections in this Run. Do not start Build, claim `CONTRACT_READY`, or reopen Human split approval solely because snapshot content changed.

Write the affected task snapshot files and manifest under `RUN_PATH/tasks/`, plus this Run's `launch-record.md`, as required by the canonical decomposition prompt and overlay. If any child or manifest is confirmed unaffected, preserve it by not emitting a redundant copy and state why in the launch record. Stop after phase handoff for Orchestrator route reconciliation.

## Execution envelope

- `PREAUTHORIZED`: read the assigned current sources and write this Run's affected task snapshots, manifest, and `launch-record.md` only.
- `ORCHESTRATOR_DECISION`: reconcile projections after handoff and prepare the next justified Build/Review route once dependencies/gates permit.
- `HUMAN_REQUIRED`: any changed split/dependency, new Product/security/domain decision, residual-risk acceptance, protected Tier-0 authorization, spec/API owner acceptance, or `CONTRACT_READY`/delivery milestone acceptance.

## Human-assisted dispatch

Use a fresh Planner Session in the Kencleng repository root with model `gpt-6-luna` and reasoning effort `high`. Start from `../harscode-workspace/workflow/2-3-techplan-decomposition-prompt.md` and this Invocation; apply `workflow/orchestrated-run-overlay.md` path semantics (`RUN_PATH/tasks/`). Kickoff:

`Jalankan post-approval task-snapshot reconciliation Run TPD-S2-002-002 sesuai Invocation durable di .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-002/invocation.md dan canonical Techplan Decomposition prompt Harscode saat ini. Rekonstruksi dari current-effective Approved TP-015 dan task split yang sudah Human-accepted pada TPD-001. Terapkan Step 0, pertahankan Task 01 → Task 02 topology/dependency/manifest kecuali TP-015 membuktikan perubahan material, dan segarkan hanya child snapshots yang terdampak beserta manifest references yang stale. Task 01 Human acceptance tetap parallel. Tulis task files terdampak dan manifest di RUN_PATH/tasks/ serta launch-record Run ini saja. Jangan edit Techplan, authority, source specs/OpenAPI, code/test, task history, atau projection; jangan minta approval split ulang hanya karena konten snapshot; jangan mulai Build atau klaim CONTRACT_READY. Berhenti setelah phase handoff.`

After dispatch, verify that the accepted topology/dependency remains unchanged and that each refreshed snapshot points to TP-015. Then reconcile projections. Do not prepare Task 02 Build until the hard Task 01 dependency and all relevant Human/owner gates are satisfied; Task 01 Human acceptance remains independently actionable.
