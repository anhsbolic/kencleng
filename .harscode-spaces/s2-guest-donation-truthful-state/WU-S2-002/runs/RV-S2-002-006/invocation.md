# Run Invocation — `RV-S2-002-006`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-09-30 by Orchestration Operator. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `RV-S2-002-006`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-006`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Independent review of Planner resolution for `RV-S2-002-005`
- `PARTICIPANT_ID`: `P-S2-002-RV-006-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; pin current Profile content revision at dispatch.
- `SESSION_TRANSITION`: `FRESH` — new Reviewer Run after completed Planner Run `TP-S2-002-011`.
- `TARGET_REVISION`: Kencleng checkout `dfbcf7b9241dbe73e2dc2b4c7707f0f829b6cff5` plus the explicit TP-011 Run artifacts; re-read live authorities at dispatch.
- `WORKFLOW_REVISION`: Harscode checkout `06a38c668b66227c3531431471b32f4f7df3699b`; re-read canonical guidance at dispatch.
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `SELECTED_MODEL`: `gpt-6-luna`; the Human-owned local runtime configuration marks approval not required.
- `REASONING_EFFORT`: `high`.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Required fresh independent review after a material Techplan resolution; canonical Complex gate also applies to cross-domain concurrency/money, payment, and PII/security boundaries.
- `MODEL_ROUTING_RATIONALE`: Use the configured non-escalation model/effort for bounded fidelity review. Escalate only if concrete capability insufficiency is evidenced; missing authority or context must be routed separately.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-011/techplan.md` and `launch-record.md` — current Draft / In Review resolution; review this artifact, not TP-010 or the Approved predecessor.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-005/review-findings.md` — blocking findings being resolved.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-010/techplan.md` — prior material D1 amendment, for lineage comparison.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-007/techplan.md` and matching Human-approved report — current-effective Approved predecessor until a later Human gate.
- All durable Exploration evidence in `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-001/runs/EXP-S2-001-001/evidence/`; broad rereading applies if the canonical Complex gate says review is warranted.
- Relevant focused evidence: OIR-004 `design-review-brief.md`; OIR-005 `amount-contract-brief.md`; OIR-006 decision brief and Stage 2 evidence; current `.harscode-spaces/authority-map.md`, `manifest.md`, `events.md`, tracker, Product/MVP sources.
- Canonical current Harscode guidance: `workflow/2-2-techplan-review-prompt.md`; `workflow/2-techplan/template.md`, `rules.md`, `guardrails.md`; `workflow/orchestrated-run-overlay.md`; `orchestration/run-contract.md`; applicable AGENTS files.
- Re-open current target-repo authority/live code only where a technical claim needs independent verification.

## Task and completion condition

Independently review the whole TP-011 Techplan under the canonical review prompt. Resolve the current template section names at runtime. Check rule/source fidelity and rule-to-testing coverage, decision fidelity, Open Item lifecycle, Test Focus anchors, the current O1 and O2/O3 scoped gates, O7 exact-copy propagation and rejected alternatives, D1 preservation, Tier-0 fencing, split OpenAPI/generated-artifact discipline, and relevant technical claims. Classify material/blocking versus mechanical/non-blocking findings using the canonical prompt. Do not rewrite the Techplan or make authority decisions.

Write only this Run's `review-findings.md` and `launch-record.md`. Do not generate a Planner report, approve the plan, accept residual risk, edit specs/OpenAPI/code/tests, claim `CONTRACT_READY`, or start Build. Stop after the independent review handoff.

## Execution envelope

- `PREAUTHORIZED`: read current cited authority/evidence and write only this Run's review artifacts.
- `ORCHESTRATOR_DECISION`: expand scope or change the workflow route.
- `HUMAN_REQUIRED`: O1 owner/scope attribution, O2/O3 substantive reconciliation, product/design/security/API authority changes, Techplan approval, residual-risk acceptance, protected writes, or milestone acceptance.

## Human-assisted dispatch

Use a fresh Reviewer Session in the Kencleng repository root with the configured model/effort. Start from canonical `workflow/2-2-techplan-review-prompt.md` and this Invocation. Kickoff: `Jalankan independent Techplan review Run RV-S2-002-006 sesuai Invocation durable di .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-006/invocation.md dan canonical Techplan Review prompt Harscode saat ini. Review current TP-S2-002-011 berdasarkan sumber dan evidence durable yang dirujuk; ikuti Complex gate serta format findings canonical. Jangan edit Techplan/authority/spec/API/code/test, jangan buat report, jangan klaim CONTRACT_READY atau mulai Build. Tulis hanya review-findings.md dan launch-record.md Run ini, lalu berhenti.`

After dispatch, report only material issue or completion; Orchestrator reconciles durable state from the Run artifacts.
