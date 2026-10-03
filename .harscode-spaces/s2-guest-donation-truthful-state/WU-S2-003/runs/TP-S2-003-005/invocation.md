# Run Invocation — `TP-S2-003-005`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-03 by Orchestration Operator. Human-facing prose: Bahasa Indonesia; retain canonical Harscode terms and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-003`
- `RUN_ID`: `TP-S2-003-005`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-005`
- `ARTIFACT_TARGET`: `RUN_PATH/techplan.md` and `RUN_PATH/handoff.md` — refreshed Draft Techplan and phase handoff for this Run; preserve all prior Run artifacts.
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Planner
- `SPECIALIZATION`: None established
- `PARTICIPANT_ID`: `P-S2-003-TP-005-1`
- `PARTICIPANT_PROFILE_ID`: `KC-PLANNER`; `.harscode-spaces/participant-profiles/profiles.md` SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — new Planner occurrence to revise the still-unapproved Draft after a Human scope-routing decision; do not reuse a prior Planner Session.
- `TARGET_REVISION`: Kencleng HEAD `bb69cd002b3f1a1056837affcd77bb2b001007b0` plus the current working tree; re-ground live source on dispatch.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; observed provenance, ordinary applicable guidance remains current-effective.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: The existing execution plan spans monetary admission, concurrency/exact-once settlement, Campaign/Donation interfaces, security gates, and protected paths. The registered repository-work/reasoning model at high effort is sufficient for the bounded Draft revision; escalate only if a demonstrated capability gap remains after current authority and evidence are available.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Fresh Techplan revision of the latest unapproved WU-S2-003 Draft after the Human routed Campaign draft cap create/PATCH backend ownership to WU003. Preserve Draft status; independent Review and whole-plan Human approval remain later gates.

## Current-effective inputs / `PRIOR_ARTIFACTS`

- Current completed Draft predecessor: `WU-S2-003/runs/TP-S2-003-004/techplan.md`, SHA-256 `f9a5fec40374abe0f89b3ae2015d936f6ca65e654b4ecf031b2d67d745397899`; handoff `4c0062a3290fbf1a168425099dca65e9404c5e9f407104a8491eda92b0886cfb`. This plan remains Draft / In Review and unapproved. Preserve its valid decisions, exact accepted-source semantics, and provenance; update the plan under the existing project Run-local Draft convention.
- Prior Draft predecessor `WU-S2-003/runs/TP-S2-003-003/techplan.md` and its handoff/launch record; completed Review `WU-S2-003/runs/RV-S2-003-001/review-findings.md`, handoff, and launch record; completed Exploration artifacts under `WU-S2-003/runs/EXP-S2-003-001/evidence/` and terminal handoff.
- Human routing decision recorded in the Parent Events for 2026-10-03: Campaign draft cap create/PATCH backend ownership belongs to WU-S2-003. This is a Work Unit routing/scope decision only; it does not amend Product/API semantics, approve this Techplan, authorize protected writes, or accept delivery/runtime risk.
- WU-S2-003 current manifest and parent `outcome.md`, `work-graph.md`, `control-surface.md`, and development tracker for current dependencies/status.
- WU-S2-006 manifest and Events; Approved `WU-S2-006/runs/TP-S2-006-008/techplan.md`; exact seven accepted Campaign/Donation spec and authored split API hashes in `WU-S2-006/runs/RV-S2-006-006/invocation.md`; BLD-S2-006-005, RV-S2-006-006, and BLD-S2-006-006 evidence. Acceptance remains limited to those exact source bytes. BLD006006 report is the current generated/frontend counterpart evidence.
- Current applicable authorities: root `AGENTS.md`; `backend/AGENTS.md`; relevant `docs/kencleng-agentic-workflow.md` delivery/risk sections; `docs/product/mvp-scope.md`; `docs/product/mvp-delivery-slices.md` Slice 2; current monetary standard; accepted Campaign and Donation invariants/features; split authored `api/openapi/campaign.yaml`, `donation.yaml`, and referenced `common.yaml`; `api/README.md`; backend architecture and live code anchors. Do not change these sources in this Run.
- Canonical Harscode `workflow/2-1-techplan-synthesis-prompt.md`, `workflow/2-techplan/template.md`, `rules.md`, `guardrails.md`, `workflow/orchestrated-run-overlay.md`, `orchestration/AGENTS.md`, and `orchestration/run-contract.md`.

## Task and completion

Revise the WU-S2-003 execution-grade Draft to reflect WU003's ownership of the backend Campaign draft create/PATCH capability required to configure the accepted per-Campaign donation cap. Reconcile the accepted create-omission default, PATCH-omission preservation, and publication-freeze semantics from current Product/spec/API sources. Reassess the affected requirements, implementation ownership, code/API anchors, risks, Rules & Validation, Testing Checklist, dependencies, and Open Items against live backend behavior; determine the minimum coherent backend work required by the accepted Campaign write contract without silently broadening unrelated Campaign lifecycle or frontend scope. Keep the public response DTO/mapping and exact-wire test coupled in WU003.

Preserve all accepted Product/API semantics and the existing donation/capacity, retry, D1, settlement, security, and monetary decisions. Do not introduce a new decision or pretend that the scope assignment resolves protected-path permissions, missing O3/O4/O5 controls, runtime evidence, or risk acceptance. Keep backend/frontend production work separate. If current authority conflicts or a material decision beyond this routing is required, record the bounded issue and stop for Orchestrator/Human routing rather than inventing a resolution.

Produce only `RUN_PATH/techplan.md` and a compact structured `RUN_PATH/handoff.md`. Keep Draft / In Review. Declare the material delta and independent Review recommendation. Do not generate a Human report before applicable Review/resolution converges; do not approve, change stable manifests/projections, modify Product/spec/API/code/tests, run Build, or claim runtime readiness.

## Execution envelope

- Read current sources and planning authorities as required by the canonical prompt. The exact seven accepted source hashes are already tracked in WU006; confirm they remain current before relying on them.
- Authorized writes: only `RUN_PATH/techplan.md` and `RUN_PATH/handoff.md`.
- No tests, validators, generators, database/runtime actions, protected writes, risk acceptance, or downstream dispatch in this Run.
- Stop and report material source drift, authority contradiction, or missing owner decision not covered by the settled WU003 route.

## Human-assisted dispatch

- Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`
- Fresh Planner / `KC-PLANNER` Session; configured `gpt-6-luna` / `high`.
- Canonical kickoff: `Jalankan Planner Run TP-S2-003-005 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-005/invocation.md dan canonical ../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md dengan orchestrated-run overlay. Berhenti setelah refreshed Draft Techplan dan phase handoff.`
