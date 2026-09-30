# Run Invocation — `RV-S2-002-008`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-09-30 by Orchestration Operator. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `RV-S2-002-008`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-008`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Targeted independent confirmation of blocking finding F-01 from `RV-S2-002-007`
- `PARTICIPANT_ID`: `P-S2-002-RV-008-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; current Profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — new Reviewer Run after completed Build/Patch Run `BLD-S2-002-002`.
- `SESSION_TRANSITION_REASON`: Fresh Review re-entry — independent targeted confirmation of the Review-requested patch from a new Participant and Session.
- `TARGET_REVISION`: Kencleng checkout `6891341a050982e14174ab5af132a200f24e71d9` plus the current Task 01 diff and completed `BLD-S2-002-002` / `RV-S2-002-007` artifacts; inspect the exact changed Summary at dispatch.
- `WORKFLOW_REVISION`: Harscode checkout `b179360088fbf8a7dec6c285cd8c0b8992cbb4aa`; re-read current Code Review entrypoint/guidance at dispatch.
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `SELECTED_MODEL`: `gpt-6-luna`; Human-owned local configuration marks approval not required.
- `REASONING_EFFORT`: `high`.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Narrow Review re-entry requested by completed Build/Patch; targeted confirmation only, because the patch plan explicitly says an exact narrow correction does not require another full four-pass review.
- `RISK_TIER`: Tier 1 — confirm the D1 mechanism boundary remains accurate. No new implementation, architecture, or broad closure review is in scope.
- `MODEL_ROUTING_RATIONALE`: The assigned question is a single-summary comparison against one accepted finding and patch plan. The configured low-cost Reviewer-capable model at high effort is sufficient; escalate only for evidenced capability insufficiency, not missing authority or context.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-007/review-findings.md` — F-01 and its blocking impact.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-007/patch-plan.md` — exact accepted correction and explicit bounds.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-002/patch-report-1.md` and `launch-record.md` — Implementer handoff/context; independently inspect the source correction itself.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-011/techplan.md` — current-effective Human-approved spine, especially D1/R7 and the unselected mechanism boundary.
- Accepted task: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-001/tasks/01-donation-domain-spec-reconciliation.md`.
- Actual confirmation scope: only the changed Summary paragraph in `docs/spec/4-campaign/features/09-closure.md`; compare with `docs/spec/4-campaign/invariants.md#inv-campaign-13` and the D1 cross-reference directly above the Summary. Do not review unrelated pre-existing source changes as part of this targeted check.
- Canonical current Review guidance: `../harscode-workspace/workflow/4-code-review-prompt.md`, especially its narrow-patch note; `workflow/4-code-review/guidelines.md`; `workflow/orchestrated-run-overlay.md`; `workflow/context-management.md`; `orchestration/run-contract.md`; `orchestration/AGENTS.md`; `workflow/AGENTS.md`.
- Target-repo convention: `docs/spec/README.md` and root `AGENTS.md`.

## Review question and completion condition

Independently verify whether the changed Campaign Summary resolves F-01 exactly: it must not claim the three close triggers share the `WHERE status = 'published'` mechanism; it should retain trigger context, point to `INV-campaign-13` / the D1 cross-reference, and state that the close-ordering mechanism remains unselected. Confirm it neither contradicts D1 nor expands this narrow reference into broader Slice 3 closure reconciliation.

This is targeted confirmation, not a new four-pass review. Do not infer runtime/concurrency verification. Do not edit production/spec files. Write only `review-confirmation.md` and `launch-record.md` in this Run. Record the evidence anchor, whether F-01 is resolved, and the precise route. If F-01 remains or the patch introduces a new material contradiction, identify it and write a patch plan only if a further production change is required. No tests are indicated or authorized for this read-only confirmation.

If F-01 is resolved, the Task 01 review loop is complete; Orchestrator will reconcile the frontier and determine the Task 02 route, keeping applicable owner/Human review and O1/O11/O2–O5/O8 gates visible. Do not mark specs `agreed`, claim `CONTRACT_READY`, accept residual risk, or edit orchestration projections from this Reviewer Run.

## Execution envelope

- `PREAUTHORIZED`: read only the specified finding, patch plan, spine/task anchors, live changed paragraph, and relevant Review guidance; write only this Run's confirmation and launch record; perform focused read-only comparison.
- `ORCHESTRATOR_DECISION`: broaden the Review scope or change the workflow route.
- `HUMAN_REQUIRED`: any source edit, owner/Human spec agreement, Product/Design/API/Security decision, residual-risk acceptance, or milestone acceptance.

## Human-assisted dispatch

Use a fresh Reviewer Session in the Kencleng repository root with the configured model/effort. Start from `../harscode-workspace/workflow/4-code-review-prompt.md` and this Invocation. Kickoff: `Jalankan targeted independent Review Run RV-S2-002-008 sesuai Invocation durable di .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-008/invocation.md dan canonical Code Review guidance saat ini. Konfirmasi hanya F-01 terhadap paragraph Summary yang diubah, patch plan RV-007, D1/INV-campaign-13, dan batas Task 01. Ini targeted confirmation untuk patch sempit, bukan full four-pass review. Tulis hanya review-confirmation.md dan launch-record.md milik Run ini (patch-plan hanya jika muncul kebutuhan patch baru). Jangan edit source/projection, jangan klaim runtime/concurrency verification, jangan tandai spec agreed atau CONTRACT_READY; berhenti setelah handoff.`

After dispatch, Human reports a material issue or completion. The Orchestrator reconciles from this Run's durable artifacts; no automatic Participant dispatch or background monitoring is authorized.
