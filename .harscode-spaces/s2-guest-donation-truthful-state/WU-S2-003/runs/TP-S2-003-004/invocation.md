# Run Invocation — `TP-S2-003-004`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-03 by Orchestration Operator. Human-facing prose: Bahasa Indonesia; retain canonical Harscode terms and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-003`
- `RUN_ID`: `TP-S2-003-004`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-004`
- `ARTIFACT_TARGET`: `RUN_PATH/techplan.md` and `RUN_PATH/handoff.md` — refreshed Draft Techplan and phase handoff for this occurrence; preserve prior Run artifacts.
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Planner
- `SPECIALIZATION`: None established
- `PARTICIPANT_ID`: `P-S2-003-TP-004-1`
- `PARTICIPANT_PROFILE_ID`: `KC-PLANNER`; profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — material Techplan refresh after exact owning-source acceptance and generated/frontend counterpart convergence; assign a fresh Planner context.
- `TARGET_REVISION`: Kencleng HEAD `bb69cd002b3f1a1056837affcd77bb2b001007b0` plus current working tree; re-ground live source on dispatch.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; ordinary applicable guidance remains current-effective.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: The plan spans monetary admission, concurrency/exact-once settlement, cross-domain Campaign/Donation contracts, security gates and protected paths. The registered repository-work/reasoning model at high effort is sufficient for synthesis; escalate only if execution demonstrates a capability gap, not for missing authority/context.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Fresh Techplan synthesis/refresh of the existing unapproved Draft after WU-S2-006 source and frontend counterpart convergence. Produce a new Run-local Draft and handoff; applicable independent Techplan Review and the Human whole-plan gate remain later.

## Current-effective inputs / `PRIOR_ARTIFACTS`

- Current Draft predecessor: `WU-S2-003/runs/TP-S2-003-003/techplan.md` (SHA-256 `af497b595af9066842333c0950ac2fc8e08825be14c001a908d6626fe3222a72`), its handoff and launch record; preserve valid decisions and predecessor provenance. It is not whole-plan approved.
- Completed prior Exploration evidence: every durable file in `WU-S2-003/runs/EXP-S2-003-001/evidence/` (`stage-2-gap-analysis.md`, `stage-3-solutioning.md`); read once under fresh-session canonical guidance.
- Completed independent finding source `WU-S2-003/runs/RV-S2-003-001/review-findings.md`, handoff and launch record; preserve its retry-order resolution and any still-active obligations.
- WU-S2-006 current manifest, Events, Work Graph, Control Surface and Outcome; Approved `WU-S2-006/runs/TP-S2-006-008/techplan.md`; BLD-S2-006-005, RV-S2-006-006 and BLD-S2-006-006 reports/invocations as needed for precise provenance.
- Exact source-acceptance receipt: seven current accepted Campaign/Donation spec and authored split API source hashes recorded in `WU-S2-006/runs/RV-S2-006-006/invocation.md` and WU-S2-006 manifest. All seven source hashes were rechecked on 2026-10-03 while preparing this Invocation. Acceptance is limited to those bytes.
- BLD-S2-006-006 counterpart evidence: regenerated `api/openapi.yaml`, generated frontend types, public Campaign fixtures and focused fixture test; report `WU-S2-006/runs/BLD-S2-006-006/report.md`. The backend response DTO and old nine-key exact-wire assertion remain unimplemented and are routed to this WU.
- Current project authorities: root `AGENTS.md`; `backend/AGENTS.md`; `docs/kencleng-agentic-workflow.md` applicable backend/delivery sections; Product/MVP scope and Slice 2 delivery authority; monetary standard; accepted Campaign/Donation invariants and features; authored `api/openapi/campaign.yaml`, `api/openapi/donation.yaml`, referenced `common.yaml`, aggregate/generated correspondence; current backend architecture and relevant live code/migration anchors.
- Parent `outcome.md`, `work-graph.md`, current development tracker and project `Authority Map` for status/dependency ownership.
- Canonical Harscode `workflow/2-1-techplan-synthesis-prompt.md`, triggered `workflow/2-techplan/template.md`, `rules.md`, `guardrails.md`, `workflow/orchestrated-run-overlay.md`, plus `orchestration/AGENTS.md` and `orchestration/run-contract.md`.

## Task and completion

Refresh the WU-S2-003 execution-grade Techplan against current accepted sources and live backend authority after source convergence. Preserve still-valid predecessor decisions and identify actual changed requirements, interface/data contracts, implementation scope, risks, protected-write gates, verification ownership, dependencies and Open Items. Ensure the plan carries the Campaign public response `max_donation_amount` projection and the corresponding exact-wire contract/test update with its backend response DTO/mapping. Include the accepted generic capacity no-fit `422 ValidationError` behavior, keeping its predicate distinct from closed/ineligible `409`, and retain idempotent retry and settlement invariants. Recheck live code/spec anchors; do not treat old plan text as current fact.

Do not reopen settled Product/API decisions or change accepted source requirements. Do not invent transaction/locking design; respect Tier-0 fencing for ledger/transaction paths and the explicit protected-write gates in root/backend instructions. Record what needs a Human authorization gate in the plan. Preserve required independent Testing/runtime/security evidence and the explicit owner-confirmed internal rollout posture.

Produce only the refreshed `RUN_PATH/techplan.md` and a compact structured `RUN_PATH/handoff.md` for this occurrence. Keep Draft status. Declare actual delta/materiality and whether independent Techplan Review is required. Do not create the Human report before applicable Review/resolution converges; do not approve the plan, change stable manifests/status projections, author product/spec/API/code, run Build, or claim runtime readiness.

## Execution envelope

- Read current source and planning authorities as required by the canonical prompt; no production, Product, spec, API, tracker, or other Work Unit writes.
- Authorized writes: this Run directory's `techplan.md` and `handoff.md` only.
- No tests, validators, generators, database/runtime actions, protected writes, risk acceptance, or downstream dispatch in this Run.
- Stop and report a material authority contradiction or missing owner decision rather than filling it in.

## Human-assisted dispatch

- Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`
- Fresh Planner / `KC-PLANNER` Session; configured `gpt-6-luna` / `high`.
- Canonical kickoff: `Jalankan Planner Run TP-S2-003-004 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-004/invocation.md dan canonical ../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md dengan orchestrated-run overlay. Berhenti setelah refreshed Draft Techplan dan phase handoff.`
