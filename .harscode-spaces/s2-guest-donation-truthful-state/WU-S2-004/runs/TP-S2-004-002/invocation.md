# Run Invocation — `TP-S2-004-002`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-03 by Orchestration Operator. Human-facing prose: Bahasa Indonesia; retain canonical Harscode terms and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-004`
- `RUN_ID`: `TP-S2-004-002`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/TP-S2-004-002`
- `ARTIFACT_TARGET`: `RUN_PATH/techplan.md` and `RUN_PATH/handoff.md` — refreshed Draft Techplan and phase handoff for this occurrence; preserve prior Run artifacts.
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Planner
- `SPECIALIZATION`: None established
- `PARTICIPANT_ID`: `P-S2-004-TP-002-1`
- `PARTICIPANT_PROFILE_ID`: `KC-PLANNER`; profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — material Techplan refresh after exact owning-source acceptance and generated/frontend counterpart convergence; assign a fresh Planner context.
- `TARGET_REVISION`: Kencleng HEAD `bb69cd002b3f1a1056837affcd77bb2b001007b0` plus current working tree; re-ground live source on dispatch.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; ordinary applicable guidance remains current-effective.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `medium`
- `MODEL_APPROVAL`: Not required by Human-owned `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: This is a frontend plan refresh against an accepted changed contract and established design/architecture sources. Registered repository-work/reasoning capability at medium effort is sufficient for the scoped synthesis; escalate only if execution demonstrates a capability gap, not for missing authority/context.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Fresh Techplan synthesis/refresh of the existing unapproved Draft after WU-S2-006 source and generated/frontend fixture counterpart convergence. Produce a new Run-local Draft and handoff; recommended independent Techplan Review and the Human whole-plan gate remain later.

## Current-effective inputs / `PRIOR_ARTIFACTS`

- Current Draft predecessor: `WU-S2-004/runs/TP-S2-004-001/techplan.md` (SHA-256 `530c4e53b74db5a974e6031a9b9f698c3589c1ce6917d07a7172a4b11d6063b0`), handoff and launch record; preserve valid product/design/architecture decisions and predecessor provenance. It is Draft / In Review, not approved.
- Completed prior Exploration evidence: every durable file in `WU-S2-004/runs/EXP-S2-004-001/evidence/` (`stage-2-gap-analysis.md`, `stage-3-solutioning.md`); read once under fresh-session canonical guidance.
- WU-S2-006 current manifest, Events, Work Graph, Control Surface and Outcome; Approved `WU-S2-006/runs/TP-S2-006-008/techplan.md`; BLD-S2-006-005, RV-S2-006-006 and BLD-S2-006-006 reports/invocations as needed for exact provenance.
- Exact source-acceptance receipt: seven current accepted Campaign/Donation spec and authored split API source hashes recorded in `WU-S2-006/runs/RV-S2-006-006/invocation.md` and WU-S2-006 manifest. All seven source hashes were rechecked on 2026-10-03 while preparing this Invocation. Acceptance is limited to those bytes.
- BLD-S2-006-006 counterpart evidence: generated OpenAPI types and public Campaign fixtures now contain required `max_donation_amount` shape; its report is in `WU-S2-006/runs/BLD-S2-006-006/report.md`. No public display/form behavior was changed by that Run.
- Current Product/MVP and design authorities: root `AGENTS.md`; `frontend/AGENTS.md`; applicable `docs/kencleng-agentic-workflow.md` sections; `docs/product/mvp-scope.md`; `docs/product/mvp-delivery-slices.md`; `docs/ui-ux/README.md`, page map and applicable patterns; current frontend architecture; accepted Campaign/Donation specs and split authored API (`campaign.yaml`, `donation.yaml`, referenced `common.yaml`) plus generated correspondence.
- Parent `outcome.md`, `work-graph.md`, current development tracker and project `Authority Map` for status/dependency ownership.
- Canonical Harscode `workflow/2-1-techplan-synthesis-prompt.md`, triggered `workflow/2-techplan/template.md`, `rules.md`, `guardrails.md`, `workflow/orchestrated-run-overlay.md`, plus `orchestration/AGENTS.md` and `orchestration/run-contract.md`.

## Task and completion

Refresh the WU-S2-004 execution-grade Techplan against the current accepted API/spec sources and generated frontend counterparts. Preserve valid predecessor decisions and derive actual plan changes from the live authority. Include the accepted per-Campaign cap contract in the public guest path: disclose the required `max_donation_amount` before amount entry, preserve explicit IDR currency in every public detail state, and cover the accepted generic `422 ValidationError` for an eligible capacity no-fit while keeping closed/ineligible behavior distinct. Plan visible cap disclosure and any amount-entry interaction in the appropriate frontend surfaces using current product/design authority; do not invent a new policy, expose remaining capacity/close reason, or imply real payment settlement.

Reassess requirements, scope, errors/recovery, fixtures/mocks, design readiness, Testing checklist and rendered Human acceptance against current sources. Preserve frontend-only ownership; do not change backend or shared sources. Keep backend DTO/exact-wire implementation coupled to WU-S2-003 rather than placing it in frontend scope. Record external integration/security/optional-email and Slice-3 boundaries accurately.

Produce only the refreshed `RUN_PATH/techplan.md` and a compact structured `RUN_PATH/handoff.md` for this occurrence. Keep Draft status. Declare actual delta/materiality and whether independent Techplan Review applies (the existing WU004 manifest recommends it after source convergence). Do not create the Human report before applicable Review/resolution converges; do not approve the plan, change stable manifests/status projections, author source/API/backend/frontend production code, run Build, or claim runtime readiness.

## Execution envelope

- Read current source and planning authorities as required by the canonical prompt; no production, Product, spec, API, tracker, backend, or other Work Unit writes.
- Authorized writes: this Run directory's `techplan.md` and `handoff.md` only.
- No tests, validators, generators, database/runtime actions, risk acceptance, or downstream dispatch in this Run.
- Stop and report a material authority contradiction or missing owner decision rather than filling it in.

## Human-assisted dispatch

- Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`
- Fresh Planner / `KC-PLANNER` Session; configured `gpt-6-luna` / `medium`.
- Canonical kickoff: `Jalankan Planner Run TP-S2-004-002 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/TP-S2-004-002/invocation.md dan canonical ../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md dengan orchestrated-run overlay. Berhenti setelah refreshed Draft Techplan dan phase handoff.`
