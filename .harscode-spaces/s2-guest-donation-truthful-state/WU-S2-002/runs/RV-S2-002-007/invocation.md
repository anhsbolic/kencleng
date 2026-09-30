# Run Invocation — `RV-S2-002-007`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-09-30 by Orchestration Operator. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `RV-S2-002-007`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-007`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Independent four-pass review of Task 01 Donation domain-spec reconciliation
- `PARTICIPANT_ID`: `P-S2-002-RV-007-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; pinned current Profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — independent Reviewer Run after completed Implementer Run `BLD-S2-002-001`.
- `SESSION_TRANSITION_REASON`: `INDEPENDENCE` — review the actual Task 01 spec diff from a separate Participant and Session context.
- `TARGET_REVISION`: Kencleng checkout `6891341a050982e14174ab5af132a200f24e71d9` plus completed Run artifacts and the current Task 01 source diff; inspect the exact files below at dispatch.
- `WORKFLOW_REVISION`: Harscode checkout `b179360088fbf8a7dec6c285cd8c0b8992cbb4aa`; re-read canonical Review guidance at dispatch.
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `SELECTED_MODEL`: `gpt-6-luna`; Human-owned local configuration marks approval not required.
- `REASONING_EFFORT`: `high`.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Required fresh independent Code Review after Tier-1 Build of Task 01. Review the actual current source diff; no production edits.
- `RISK_TIER`: Tier 1 — payment, money-state, concurrency, guest privacy, and security requirements are in scope. Independent review and applicable owner/Human review are required before affected specs are treated as `agreed`.
- `MODEL_ROUTING_RATIONALE`: Bounded document-diff review across a known approved spine, accepted task, and target-repo conventions. The configured low-cost Reviewer-capable model at high effort is sufficient; escalate only if concrete capability insufficiency is evidenced, not for missing authority or context.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-011/techplan.md` — current-effective Human-approved `Approved` spine; review against its Q1–Q11, R1–R11, D1–D16, RISK-1–RISK-10, §12 verification, and O1–O5/O8/O11 boundaries.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-001/tasks/01-donation-domain-spec-reconciliation.md` — Human-accepted Task 01 execution contract.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-001/tasks/manifest.md` — accepted split; Task 02 is downstream and is outside this Review.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-001/report.md` and `launch-record.md` — Implementer handoff/context, not a substitute for inspecting the diff.
- Actual Task 01 source diff, limited to:
  - `docs/spec/5-donation/invariants.md`
  - `docs/spec/5-donation/threat-model.md`
  - `docs/spec/5-donation/tasks.md`
  - `docs/spec/5-donation/features/01-submit-donation-settlement.md`
  - `docs/spec/5-donation/features/02-donation-status-check.md`
  - `docs/spec/5-donation/features/03-public-donor-list.md` through `06-guest-email-reveal.md`
  - `docs/spec/4-campaign/invariants.md`
  - `docs/spec/4-campaign/features/09-closure.md`
- Target-repo consistency authorities: root `AGENTS.md`, `docs/spec/README.md`, and the applicable domain-spec / owner-review section of `docs/kencleng-agentic-workflow.md`.
- Canonical current Harscode guidance: `workflow/4-code-review-prompt.md`; `workflow/4-code-review/guidelines.md`; `workflow/4-code-review/checklist.md`; `workflow/orchestrated-run-overlay.md`; `workflow/context-management.md`; `orchestration/run-contract.md`; `orchestration/AGENTS.md`; `workflow/AGENTS.md`.
- For Stack-Specific Pass 3, route through `best-practices/AGENTS.md` and targeted rows in `best-practices/index.md`; open only a matching practice for this documentation-only diff. If no trigger matches, state the pass is clean/no-op.

## Task and completion condition

Run all four canonical review passes against the same explicit current Task 01 source diff, in order: Safety, Quality, Stack-Specific Best Practices, Consistency. Inspect the actual diff independently against TP-011 and the accepted Task 01; do not rely on the Build report as proof. Check authority fidelity, domain-first spec structure/status, D1 and exact-money/no-float requirements, idempotent submit versus settlement replay, backend-owned terminal state, guest privacy, threat-model risk acceptance, O1/O11 and O2–O5/O8 scoping, Q11 copy propagation, deferred features, and Campaign scope boundaries. Do not invent findings to fill a pass.

Write only `review-findings.md` and `launch-record.md` in this Run, using the canonical four-pass output, provenance, findings, verdict, and handoff sections. Add `patch-plan.md` only if a finding requires production changes. Do not edit production/spec authority. Do not run broad tests; targeted read-only commands/reproductions are allowed only to resolve a concrete review question and must be recorded. No runtime/testing evidence is inferred from this documentation diff.

If changes are required, give location, problem, impact, suggested resolution, and blocking/non-blocking classification; route the patch through a new Build Run. If review approves, the next route is applicable Donation/Campaign, Security/PII, API, and Human/domain-owner review before affected specs may be treated as `agreed`. Task 02 may be considered only after Task 01 review/acceptance and must preserve its own scoped gates. Do not claim `CONTRACT_READY`, accept residual risk, execute Task 02, or change orchestration projections from the Reviewer Run.

## Execution envelope

- `PREAUTHORIZED`: read the named current-effective inputs, actual scoped diff, and applicable Review guidance; write only this Run's findings and launch record; use focused read-only verification if needed for a concrete finding.
- `ORCHESTRATOR_DECISION`: widen the review scope or change the phase route.
- `HUMAN_REQUIRED`: owner/Human agreement of specs, any Product/Design/Security/API decision, residual-risk acceptance, protected-path writes, milestone acceptance, or any source write.

## Human-assisted dispatch

Use a fresh Reviewer Session in the Kencleng repository root with the configured model/effort. Start from `../harscode-workspace/workflow/4-code-review-prompt.md` and this Invocation. Kickoff: `Jalankan independent Code Review Run RV-S2-002-007 sesuai Invocation durable di .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-007/invocation.md dan canonical Code Review prompt Harscode saat ini. Jalankan empat pass berurutan terhadap diff Task 01 yang benar-benar berubah: Safety, Quality, Stack-Specific Best Practices, Consistency. Cocokkan dengan Approved TP-011, task yang diterima Human, serta authority/review guidance repo. Tulis hanya review-findings.md dan launch-record.md milik Run ini, plus patch-plan.md hanya bila perlu. Jangan edit source/spec/authority, jangan jalankan broad test, jangan ubah orchestration projection, jangan klaim CONTRACT_READY atau menyetujui spec sebagai Human/owner; berhenti setelah phase handoff.`

After dispatch, Human reports a material issue or completion. The Orchestrator reconciles from this Run's durable artifacts; no automatic Participant dispatch or background monitoring is authorized.
