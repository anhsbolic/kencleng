# Run Invocation — `RV-S2-002-005`

Status: `DISPATCHED`

Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `RV-S2-002-005`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-005`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Independent review of material Techplan amendment D1
- `PARTICIPANT_ID`: `P-S2-002-RV-005-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; pinned current Profile file at dispatch
- `SESSION_TRANSITION`: `FRESH` — independent phase Run after completed Planner Run `TP-S2-002-010`.
- `TARGET_REVISION`: Kencleng checkout `dfbcf7b9241dbe73e2dc2b4c7707f0f829b6cff5`; `TP-S2-002-010` is a current untracked Run artifact and is reviewed as its explicit input.
- `WORKFLOW_REVISION`: Harscode checkout `06a38c668b66227c3531431471b32f4f7df3699b`.
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `SELECTED_MODEL`: `gpt-6-luna`; requested by the Human-owned runtime configuration. Runtime model identity is not independently exposed.
- `REASONING_EFFORT`: `high` (per configured route; runtime setting not independently exposed).
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Required independent review; amended Techplan is material and crosses Campaign/Donation money ordering, payment, contract, and PII boundaries.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-010/techplan.md` — material Draft / In Review amendment; review this artifact, not the Approved predecessor.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-010/launch-record.md` — Planner handoff and input-pin record.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-007/techplan.md` — previous Approved spine, for change/history comparison only.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/evidence/` — both durable Stage 2 gap analysis and Stage 3 solutioning artifacts; broad reread for Complex-gate fidelity.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/OIR-S2-002-006/` — owner D1 evidence, decision brief, and handoff.
- Current Product/MVP authorities: `docs/product/mvp-scope.md`, `docs/product/mvp-delivery-slices.md`; current Authority Map and `docs/kencleng-agentic-workflow.md`.
- Canonical review authority at current Harscode revision: `workflow/2-2-techplan-review-prompt.md`; `workflow/2-techplan/template.md`, `rules.md`, `guardrails.md`; `workflow/orchestrated-run-overlay.md`; `orchestration/run-contract.md`; applicable root/scoped Kencleng AGENTS.
- Independently inspect relevant current target-repo sources for 2–3 non-obvious technical claims.

## Task and write boundary

Independently review the whole amended Techplan for canonical Complex gate applicability, traceability/fidelity, current Product/MVP and OIR decision preservation, rule-to-testing coverage, Open Item lifecycle, Test Focus Pointer completeness, Tier-0 and API/generated-artifact guardrails, and spot-checked target-repo technical claims. Do not rewrite the Techplan or settle authority questions.

Write only this Run's `review-findings.md` and `launch-record.md`. Do not edit prior Runs, Product/MVP/design/spec/API/code/tests or orchestration projections. Do not generate a Planner report, claim `CONTRACT_READY`, approve the Techplan, accept residual risk, or start Build. Stop at the independent review handoff.

## Execution envelope

- `PREAUTHORIZED`: read the cited source and workflow evidence; write only this Run's review artifacts.
- `ORCHESTRATOR_DECISION`: expand scope or alter the phase route.
- `HUMAN_REQUIRED`: Product/Design/Security/PII/API authority decisions, residual-risk acceptance, Techplan approval, protected writes, or milestone acceptance.

## Canonical kickoff

Use the canonical `workflow/2-2-techplan-review-prompt.md` with these normal variables and this Invocation as the task context. The review's material focus is already stated in the Techplan and its D1 provenance; do not infer new policy. Follow the prompt's Complex gate and finding format. Stop after the review handoff.
