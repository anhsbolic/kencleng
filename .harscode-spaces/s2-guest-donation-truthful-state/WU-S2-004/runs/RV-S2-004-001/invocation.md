# Run Invocation — `RV-S2-004-001`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-03 by Orchestration Operator. Human-facing prose: Bahasa Indonesia; preserve canonical Harscode terms and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-004`
- `RUN_ID`: `RV-S2-004-001`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/RV-S2-004-001`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Independent Techplan fidelity review
- `PARTICIPANT_ID`: `P-S2-004-RV-001-1`
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; `.harscode-spaces/participant-profiles/profiles.md` SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — independent Reviewer context from Planner TP-S2-004-002; separate Run, Participant, and Session.
- `TARGET_REVISION`: Kencleng HEAD `bb69cd002b3f1a1056837affcd77bb2b001007b0` plus current working tree. Exact review target is separately pinned below; re-ground live source when dispatched.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; observed provenance, ordinary guidance remains current-effective.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: Independent fidelity review crosses public Campaign and Donation interfaces, monetary constraints, payment-state/error semantics, and frontend design/rendered-acceptance boundaries. `gpt-6-luna` has registered repository-work/reasoning capability; high effort is selected for the cross-contract review. If evidence is missing, identify the authority gap rather than escalating models to compensate.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Independent Techplan Review of the completed WU-S2-004 Draft after accepted WU-S2-006 source and generated/frontend fixture convergence. Reviewer determines whether canonical Step-0 Complex gate applies and conducts or stops the review accordingly.

## Canonical route and current inputs

- Canonical phase prompt: `../harscode-workspace/workflow/2-2-techplan-review-prompt.md` plus `../harscode-workspace/workflow/orchestrated-run-overlay.md`.
- Canonical current Techplan Review prompt calls for reading current `../harscode-workspace/workflow/2-techplan/template.md`, `rules.md`, and `guardrails.md`; all durable Exploration artifacts when the Complex gate applies; diagram guidance only if the assigned Techplan contains a diagram; and target-repository sources for technical-fact spot checks.
- Assigned review target: `WU-S2-004/runs/TP-S2-004-002/techplan.md`, SHA-256 `7b1a1f9ce7e68362681d6e2fc27ebc729243517bfca984f7984d1488dfba4535`; terminal handoff `WU-S2-004/runs/TP-S2-004-002/handoff.md`, SHA-256 `81ede34ca30aa24a18630f01a8fa3fdce9e7829fb7645930504b2ebbea367423`. Target is Draft / In Review, not approved. Treat this exact content identity as the review assignment; if the target materially changes before review, stop and reconcile the target identity before claiming review of it.
- Complete WU-S2-004 Exploration corpus: all durable evidence under `WU-S2-004/runs/EXP-S2-004-001/evidence/` and the Exploration terminal handoff. Read all of it if Step-0 Complex gate applies; do not infer evidence solely from the Planner's plan.
- Current authorities for relevant checks: root `AGENTS.md`; `docs/product/mvp-scope.md`; `docs/product/mvp-delivery-slices.md` Slice 2; `docs/ui-ux/README.md`, `docs/ui-ux/page-map.md`, and applicable design guidance; `docs/project/kencleng-frontend-tech-stack.md`; `frontend/AGENTS.md`; accepted Campaign/Donation specs and split authored API sources plus referenced common components; WU-S2-004 manifest and Parent Outcome/Work Graph/Control Surface; WU-S2-006 manifest and the exact source-acceptance/counterpart evidence below.
- WU-S2-006 source/rollout evidence: `WU-S2-006/runs/RV-S2-006-006/invocation.md` and review; `WU-S2-006/runs/BLD-S2-006-006/report.md`; current WU-S2-006 manifest. Seven reviewed Campaign/Donation spec and authored API hashes were explicitly accepted by Anhar. Owner-confirmed current MVP1 posture is no production rollout, no external contract consumer, internal/repository-development consumers only, and coordinated counterpart reconciliation before runtime/delivery. Historical O8 scope is unchanged.
- Current counterpart evidence: generated OpenAPI bundle/types and public Campaign fixtures/focused client test recorded in `WU-S2-006/runs/BLD-S2-006-006/report.md`. Backend response DTO/exact-wire assertion is assigned to WU-S2-003 and is not frontend ownership.
- WU-S2-003 routing context: its current Draft/handoff identifies a separate Human routing question over who owns Campaign draft create/PATCH for the accepted Owner/Staff cap configuration requirement. This scoped backend planning issue does not itself block independent review of WU-S2-004; do not silently assign backend writes to frontend or resolve the Human route in this Run.
- Current WU-S2-004 manifest, predecessor `TP-S2-004-001` Techplan, current accepted WU-S2-002 and WU-S2-005 baselines, and applicable communication profile.

## Review assignment

Independently assess the exact assigned Techplan against current authority and evidence. Follow the canonical review prompt's own gate and checks: rule fidelity and Testing Checklist coverage; decision fidelity; diagram validity only if present; Open Items lifecycle; 2–3 non-obvious technical-fact/guardrail spot-checks against live sources; and Test Focus Pointer completeness against Exploration evidence when Complex review applies. Identify material contradictions or unsupported claims; do not introduce alternative product/API semantics or duplicate settled decisions as open questions.

The Draft's material refreshed delta includes the accepted `max_donation_amount` disclosure before amount entry, explicit IDR in public detail states, and generic eligible-capacity no-fit `422 ValidationError` distinct from closed/ineligible `409`. Independently verify those and other plan claims against their authorities; this context does not prescribe findings or verdict. Preserve the boundary that frontend mocks/rendering do not establish backend runtime/security evidence.

## Execution envelope

- `PREAUTHORIZED`: Read assigned Draft, complete applicable Exploration evidence and routed current authorities; write only this Run's `review-findings.md`.
- `ORCHESTRATOR_DECISION`: Reconcile any Review findings and route a fresh Planner resolution/re-review/report sequence as applicable.
- `HUMAN_REQUIRED`: Whole-Techplan approval; any new Product/API/design decision; protected writes; risk acceptance; rendered Human acceptance.
- No writes to Techplan, Product, design, specs, API, code, tests, prior artifacts, manifests, tracker, registry, or other Work Unit. Do not generate `report-techplan.md`, approve, build, run tests/validators/generators/runtime/browser checks, or claim any delivery milestone. No Participant is dispatched by this preparation.
- Review output: one canonical `RUN_PATH/review-findings.md` with provenance and exactly one structured `## Phase handoff` per the active orchestrated overlay. Add `patch-plan.md` only if the canonical Review result actually requires one.

## Human-assisted dispatch

Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`.
Use a fresh independent Reviewer / `KC-REVIEWER` Session, `gpt-6-luna` / `high`.

Kickoff: `Jalankan independent Techplan Review Run RV-S2-004-001 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/RV-S2-004-001/invocation.md dan canonical ../harscode-workspace/workflow/2-2-techplan-review-prompt.md dengan orchestrated-run overlay. Tulis review-findings.md pada RUN_PATH dan berhenti setelah phase handoff.`

Report completion or the exact scoped blocker to the Orchestrator. No automatic Planner resolution, Human report, approval, Build, or downstream dispatch.
