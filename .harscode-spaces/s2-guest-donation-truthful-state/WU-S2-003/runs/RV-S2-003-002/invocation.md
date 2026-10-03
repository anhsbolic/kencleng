# Run Invocation — `RV-S2-003-002`

Status: `SUPERSEDED_BEFORE_DISPATCH`

Superseded on 2026-10-03 after Human approved a material additional planning route for minimum persisted representative/Organization eligibility data. This Invocation remains preserved as evidence of the prepared package; its exact TP-S2-003-005 review target is no longer current. Do not dispatch. A new independent Review package must pin the refreshed Draft produced by TP-S2-003-006.

Prepared: 2026-10-03 by Orchestration Operator. Human-facing prose: Bahasa Indonesia; preserve canonical Harscode terms and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-003`
- `RUN_ID`: `RV-S2-003-002`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-002`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Independent Techplan fidelity review
- `PARTICIPANT_ID`: `P-S2-003-RV-002-1`
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; `.harscode-spaces/participant-profiles/profiles.md` SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — independent Reviewer context from Planner TP-S2-003-005 and prior Review RV-S2-003-001; separate Run, Participant, and Session.
- `TARGET_REVISION`: Kencleng HEAD `bb69cd002b3f1a1056837affcd77bb2b001007b0` plus current working tree. Exact review target is separately pinned below; re-ground live source when dispatched.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; observed provenance, ordinary guidance remains current-effective.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: This plan crosses Campaign and Donation contracts, money, D1 concurrency/exact-once settlement, credentials/PII, backend authorization and protected implementation boundaries. An independent `KC-REVIEWER` at high effort is suitable for the fidelity checks; identify authority gaps from evidence rather than resolving them by model choice.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Independent Techplan Review of the completed WU-S2-003 Draft. Reviewer independently determines whether the canonical Complex gate applies and performs only the review warranted by that gate.

## Canonical route and current inputs

- Canonical phase prompt: `../harscode-workspace/workflow/2-2-techplan-review-prompt.md` plus `../harscode-workspace/workflow/orchestrated-run-overlay.md`.
- Canonical current Techplan Review prompt calls for current `../harscode-workspace/workflow/2-techplan/template.md`, `rules.md`, and `guardrails.md`; all durable WU-S2-003 Exploration artifacts when the Complex gate applies; diagram guidance only if the plan contains a diagram; and target-repository sources for technical-fact spot checks.
- Assigned review target: `WU-S2-003/runs/TP-S2-003-005/techplan.md`, SHA-256 `be0ba57dc87475b892c24dfeb61381f8f2ef9e48f3460a48175c0ac35d35dfa1`; terminal handoff `WU-S2-003/runs/TP-S2-003-005/handoff.md`, SHA-256 `5258c0e9e8d7fa39081160e102c3c43132144de2b987f0052317d88eb5dbae6b`. Target remains Draft / In Review, unapproved. This exact content identity is the review assignment; if the target materially changes before review, stop and reconcile target identity before claiming review of it.
- Complete WU-S2-003 Exploration corpus: all durable evidence under `WU-S2-003/runs/EXP-S2-003-001/evidence/` and its terminal handoff. Read the full corpus if the Step-0 Complex gate applies; do not infer it only from the Planner's Draft.
- Current authorities and evidence: root `AGENTS.md`; `backend/AGENTS.md`; relevant `docs/kencleng-agentic-workflow.md` delivery/risk sections; `docs/product/mvp-scope.md`; `docs/product/mvp-delivery-slices.md` Slice 2; monetary standard; accepted Campaign and Donation invariants/features; split authored `api/openapi/campaign.yaml`, `donation.yaml`, and referenced `common.yaml`; `api/README.md`; backend architecture and live code; WU-S2-003 manifest and parent Outcome/Work Graph/Control Surface/Events; WU-S2-006 manifest and its source/counterpart evidence.
- Exact WU-S2-006 source acceptance: seven reviewed source hashes recorded in `WU-S2-006/runs/RV-S2-006-006/invocation.md`; Anhar accepted only those exact revisions. Reconfirm current hashes against that snapshot. BLD-S2-006-006 report records the generated API/types and frontend fixture/test counterpart reconciliation. Do not infer runtime proof from these results.
- New scoped authority issue in the assigned Draft: its Open Item 7 states that accepted Campaign draft `POST`/`PATCH` authorization requires representative membership plus fresh Organization `verified` and `has_overdue_report` checks, while the live backend has no such membership/eligibility source or Campaign write handlers. Independently assess whether the item correctly preserves the accepted contract and bounds the blocker. Product/MVP currently permits seeded/operator-assisted setup while excluding full Organization self-service by default; the Organization invariants/API are Draft/historical evidence and must not be treated as accepted MVP authority. Do not choose a source, implement a persistence design, weaken predicates, or let this scoped issue erase the rest of the Donation plan.

## Review assignment

Independently assess the exact assigned Techplan against current authority and durable evidence. Follow the canonical review prompt's gate and checks: rule fidelity and Testing Checklist coverage; decision fidelity; diagram validity only if present; Open Items lifecycle; 2–3 non-obvious technical-fact/guardrail spot-checks against live sources; and Test Focus Pointer completeness against Exploration evidence when Complex review applies. Identify material contradictions or unsupported claims. Do not introduce alternative product/API semantics, re-litigate settled choices, or direct the result toward approval or rejection.

The material refresh adds WU003 ownership of accepted Campaign draft create/PATCH behavior (default, omission preservation, publication freeze, response shape and exact-wire work) and identifies Open Item 7 on the authority/source required for representative and Organization eligibility checks. Verify both against current sources and the prior settled routing decision. The review can assess the complete Draft while this issue remains open; it does not resolve the Human/source-owner question or authorize a Build.

## Execution envelope

- `PREAUTHORIZED`: Read the assigned Draft, complete applicable Exploration evidence, and routed current authorities; write only this Run's `review-findings.md`.
- `ORCHESTRATOR_DECISION`: Reconcile findings and route a fresh Planner resolution/re-review/report sequence as applicable.
- `HUMAN_REQUIRED`: Whole-Techplan approval; any Product/domain/API/security/authority decision; protected writes; risk acceptance; delivery/runtime milestone.
- No writes to the Techplan, Product, design, specs, API, code, tests, prior artifacts, manifests, tracker, registry, or other Work Unit. Do not generate `report-techplan.md`, approve, build, run tests/validators/generators/runtime/browser checks, or claim any delivery milestone. No Participant is dispatched by this preparation.
- Review output: one canonical `RUN_PATH/review-findings.md` with provenance and exactly one structured `## Phase handoff` per the active orchestrated overlay. Add `patch-plan.md` only if the canonical Review result actually requires one.

## Human-assisted dispatch

Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`.
Use a fresh independent Reviewer / `KC-REVIEWER` Session, `gpt-6-luna` / `high`.

Kickoff: `Jalankan independent Techplan Review Run RV-S2-003-002 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-002/invocation.md dan canonical ../harscode-workspace/workflow/2-2-techplan-review-prompt.md dengan orchestrated-run overlay. Tulis review-findings.md pada RUN_PATH dan berhenti setelah phase handoff.`

Laporkan completion atau scoped blocker ke Orchestrator. Tidak ada resolusi Planner, Human report/approval, Build, atau downstream dispatch otomatis.
