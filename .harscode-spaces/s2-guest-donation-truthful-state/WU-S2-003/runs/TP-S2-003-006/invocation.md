# Run Invocation — `TP-S2-003-006`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-03 by Orchestration Operator after Human approved the bounded source/persistence route. Human-facing prose: Bahasa Indonesia; retain canonical Harscode terms and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-003`
- `RUN_ID`: `TP-S2-003-006`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-006`
- `ARTIFACT_TARGET`: `RUN_PATH/techplan.md` and `RUN_PATH/handoff.md` — refreshed Draft Techplan and phase handoff; preserve all prior Run artifacts.
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Planner
- `SPECIALIZATION`: None established
- `PARTICIPANT_ID`: `P-S2-003-TP-006-1`
- `PARTICIPANT_PROFILE_ID`: `KC-PLANNER`; `.harscode-spaces/participant-profiles/profiles.md` SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — new Planner occurrence from TP-S2-003-005; do not reuse a prior Planner Session.
- `TARGET_REVISION`: Kencleng HEAD `bb69cd002b3f1a1056837affcd77bb2b001007b0` plus current working tree; re-ground live sources on dispatch.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; observed provenance, ordinary applicable guidance remains current-effective.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: The plan spans monetary admission, D1 concurrency, Campaign/Donation contracts, credentials/PII, authorization and protected boundaries. `gpt-6-luna` at high effort is configured for repository planning; identify any evidence gap from current authority instead of filling it with assumptions.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Fresh Techplan revision of the latest unapproved WU-S2-003 Draft after Anhar approved the bounded route for minimum persisted representative/Organization eligibility data using MVP-permitted seeded/operator-assisted setup. Preserve Draft status; independent Review and whole-plan Human approval remain later gates.

## Current-effective inputs / `PRIOR_ARTIFACTS`

- Latest completed Draft predecessor: `WU-S2-003/runs/TP-S2-003-005/techplan.md`, SHA-256 `be0ba57dc87475b892c24dfeb61381f8f2ef9e48f3460a48175c0ac35d35dfa1`; handoff `WU-S2-003/runs/TP-S2-003-005/handoff.md`, SHA-256 `5258c0e9e8d7fa39081160e102c3c43132144de2b987f0052317d88eb5dbae6b`. Draft remains In Review/unapproved.
- TP-S2-003-005 invocation `WU-S2-003/runs/TP-S2-003-005/invocation.md`; its prior Draft predecessors and review; complete WU-S2-003 Exploration corpus under `runs/EXP-S2-003-001/evidence/` and terminal handoff.
- Human decision recorded in parent Events on 2026-10-03: assign to WU-S2-003 the planning/delivery ownership for the minimum persisted source/integration needed by accepted Campaign draft writes to evaluate representative membership and Organization `verified` / `has_overdue_report`. Populate initial MVP data using the already-approved seeded/operator-assisted setup posture; full Organization self-service and representative-management UI remain out of scope. Do not change the accepted Owner/Staff authorization and fresh Organization-state predicates. This routes delivery ownership and minimum persisted prerequisite only; it does not decide which human/operator institution is entitled to declare verification/overdue truth, authorize protected writes, accept risk, approve the plan, or authorize Build.
- WU-S2-003 current manifest, parent Outcome/Work Graph/Control Surface/Events, and current development tracker.
- WU-S2-006 manifest and Events; Approved `WU-S2-006/runs/TP-S2-006-008/techplan.md`; exact seven accepted Campaign/Donation source hashes in `WU-S2-006/runs/RV-S2-006-006/invocation.md`; BLD-S2-006-005, RV-S2-006-006, and BLD-S2-006-006 evidence. Acceptance remains limited to those exact source bytes; confirm current source hashes before relying on them.
- Current authorities: root `AGENTS.md`; `backend/AGENTS.md`; relevant `docs/kencleng-agentic-workflow.md`; `docs/product/mvp-scope.md`; Slice 2 in `docs/product/mvp-delivery-slices.md`; monetary standard; accepted Campaign/Donation invariants and features; split authored Campaign/Donation OpenAPI plus referenced `common.yaml`; `api/README.md`; backend architecture and live code.
- The Organization invariants/API are Draft/historical evidence, not accepted MVP owner/source authority. Product/MVP allows seeded/operator-assisted initial setup under its constraints and excludes full Organization self-service/representative management by default. Use the current Product/MVP authority and routed decision; do not promote draft Organization design into accepted authority.
- Canonical Harscode `workflow/2-1-techplan-synthesis-prompt.md`, `workflow/2-techplan/template.md`, `rules.md`, `guardrails.md`, `workflow/orchestrated-run-overlay.md`, `orchestration/AGENTS.md`, and `orchestration/run-contract.md`.

## Task and completion

Revise the unapproved WU-S2-003 execution-grade Draft to incorporate the Human-routed minimum persisted representative membership and Organization eligibility source/integration needed by the accepted Campaign draft create/PATCH operations. Reassess scope, ownership, migration/data-shape, security/authorization, risks, dependencies, Rules & Validation, Testing Checklist, and Open Items against current Product/MVP authority, accepted Campaign sources, backend architecture, and live implementation. Keep initial data setup seeded/operator-assisted as permitted for MVP; do not create full Organization self-service or representative-management UI. Preserve the accepted Owner/Staff authorization requirement and fresh `verified` / `has_overdue_report` checks unchanged.

Keep explicit any unresolved Product/Security/data-authority question about who may establish or update the Organization verification/overdue facts and how that operation is governed. Do not infer that authority from the WU003 ownership route, broad Organization Draft/API, JWT claims, or an unowned operator role. Bound the resulting blocker to the affected source and Campaign draft handlers, and state which unaffected Donation planning/build areas remain gated by their separate prerequisites. Preserve all prior accepted monetary, capacity, retry, D1, settlement, security, and exact-wire obligations.

Produce only `RUN_PATH/techplan.md` and compact structured `RUN_PATH/handoff.md`. Keep Draft / In Review, state the material delta and independent Review recommendation. Do not generate a Human report before applicable Review/resolution converges; do not approve, alter stable manifests/projections, modify Product/spec/API/code/tests, run Build, or claim runtime readiness.

## Execution envelope

- Read current sources and planning authorities required by the canonical prompt. Reconfirm all seven accepted-source hashes against the WU006 snapshot.
- Authorized writes: only `RUN_PATH/techplan.md` and `RUN_PATH/handoff.md`.
- No tests, validators, generators, DB/runtime actions, protected writes, risk acceptance, or downstream dispatch in this Run.
- Stop and report source drift, authority contradiction, or a material issue beyond the approved bounded route.

## Human-assisted dispatch

- Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`
- Fresh Planner / `KC-PLANNER` Session; configured `gpt-6-luna` / `high`.
- Canonical kickoff: `Jalankan Planner Run TP-S2-003-006 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-006/invocation.md dan canonical ../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md dengan orchestrated-run overlay. Berhenti setelah refreshed Draft Techplan dan phase handoff.`
