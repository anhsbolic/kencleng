# Run Invocation — `TP-S2-002-011`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-09-29 by Orchestration Operator. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `TP-S2-002-011`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-011`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `ROLE`: Planner
- `SPECIALIZATION`: Resolve independent review findings in material Slice-2 Techplan amendment
- `PARTICIPANT_ID`: `P-S2-002-TP-011-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-PLANNER`; pin the current Profile revision at dispatch.
- `SESSION_TRANSITION`: `FRESH` — new Planner Run after completed independent Reviewer Run `RV-S2-002-005`.
- `TARGET_REVISION`: Kencleng checkout `dfbcf7b9241dbe73e2dc2b4c7707f0f829b6cff5` plus explicitly referenced TP-010/RV-005 current Run artifacts; re-read changed sources at dispatch.
- `WORKFLOW_REVISION`: Harscode checkout `06a38c668b66227c3531431471b32f4f7df3699b`; re-read canonical guidance at dispatch.
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `SELECTED_MODEL`: `gpt-6-luna`; Human-owned local config marks approval not required.
- `REASONING_EFFORT`: `high`.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Planner resolution of blocking findings from the required independent Techplan review. This is a new Run, not a continuation of TP-010.
- `MODEL_ROUTING_RATIONALE`: Current configured non-escalation model/effort is the minimum route for the bounded document reconciliation. Escalate only on evidenced capability insufficiency, not missing authority/context.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-005/review-findings.md` — three blocking findings and required handoff.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-010/techplan.md` and `launch-record.md` — current Draft / In Review D1 amendment; preserve provenance and do not overwrite.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-007/techplan.md` and matching Human-approved report — current-effective Approved predecessor until a later Human gate.
- `OIR-S2-002-004/design-review-brief.md` — completed O7 decisions and remaining rendered-acceptance boundary.
- `.harscode-spaces/s2-guest-donation-truthful-state/events.md` (2026-09-28 O1/O2-O3 entries), `WU-S2-002/manifest.md`, and current tracker — current scoped `AUTHORITY_SYNC` / `HUMAN_DECISION` state.
- `.harscode-spaces/authority-map.md` — current authority owners and missing cross-feature money-standard owner.
- `OIR-S2-002-006/campaign-donation-ordering-brief.md` — D1; carry forward without reopening.
- `docs/product/mvp-scope.md`, `docs/product/mvp-delivery-slices.md`, root/scoped `AGENTS.md`, `docs/kencleng-agentic-workflow.md`.
- Current Harscode canonical `workflow/2-1-techplan-synthesis-prompt.md`, Techplan template/rules/guardrails/report template, review prompt, `workflow/orchestrated-run-overlay.md`, and applicable orchestration guidance.

## Task and completion condition

Create a new Techplan amendment under this Run that resolves only the findings in `RV-S2-002-005` while preserving decision provenance and the current authority boundaries:

1. Reconcile O7 as completed using the exact existing Design outcomes; retain rendered visual acceptance and genuinely open O2/O3 delivery dependencies as deferred work. Do not ask for O7 to be decided again.
2. Explicitly preserve the O2/O3 conflict between independent pending-email cap route B and mandatory terminal notice for verified opt-in email as a scoped `HUMAN_DECISION`. Record what it blocks and what remains safe. Do not choose a cap, alternate retention mechanism, terminal bound, timeout result, or residual-risk acceptance.
3. Preserve O1's current Slice-2 amount direction and identify the shared cross-feature currency-standard owner/scope gap as a separate scoped `AUTHORITY_SYNC`. Do not assign an owner or choose wire/storage representation, supported currencies, range/scale, or a global standard.
4. Carry D1 from OIR-006 and the full prior plan spine forward without changing its meaning; retain all other Active/conditional Open Items and the Slice-2 boundary.

Recheck current sources at dispatch. Declare whether the reconciliation changes material scope, architecture/ownership, business/security/interface semantics, or verification strategy, and route independent re-review accordingly. Do not generate the Human report until the applicable review/resolution path converges.

## Execution envelope

- `PREAUTHORIZED`: read current cited authorities and write only this Run's new Techplan and launch record; update no source authority, prior Run, spec/API/code/test, or projection.
- `ORCHESTRATOR_DECISION`: expand beyond the findings or create new Work Units/dependency edges.
- `HUMAN_REQUIRED`: resolve the O2/O3 substantive conflict, attribute O1 cross-feature authority, change Product/Design/Security/API authority, accept residual risk, approve the Techplan, or touch protected paths.

## Human-assisted dispatch

Use a fresh Planner Session with the configured model/effort and working directory at the Kencleng repository root. Start from canonical `workflow/2-1-techplan-synthesis-prompt.md` and this Invocation. Kickoff: `Jalankan Planner resolution Run TP-S2-002-011 sesuai Invocation durable di .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-011/invocation.md dan canonical Techplan synthesis workflow Harscode saat ini. Reconstruct dari artefak yang ditunjuk; selesaikan finding RV-S2-002-005 hanya dalam authority yang sudah ada; jangan menyelesaikan atau menebak O1 AUTHORITY_SYNC maupun O2/O3 HUMAN_DECISION; tulis hanya Techplan dan launch record Run ini, lalu berhenti.`

After dispatch, report only material issue or phase completion; Orchestrator reconciles durable state from the Run artifacts.
