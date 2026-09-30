# Run Invocation — `BLD-S2-002-002`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-09-30 by Orchestration Operator. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `BLD-S2-002-002`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-002`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `ROLE`: Implementer
- `SPECIALIZATION`: Narrow Build/Patch for blocking Review finding F-01 in Task 01
- `PARTICIPANT_ID`: `P-S2-002-BL-002-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-IMPLEMENTER`; current Profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — new Build/Patch Run after completed independent Code Review Run `RV-S2-002-007`.
- `SESSION_TRANSITION_REASON`: Fresh Build re-entry — the completed Review requested a production/spec correction, so the patch uses a new orchestrated Run and Participant grounded on the accepted patch plan.
- `TARGET_REVISION`: Kencleng checkout `6891341a050982e14174ab5af132a200f24e71d9` plus current Task 01 source diff and completed `BLD-S2-002-001` / `RV-S2-002-007` Run artifacts; reopen the live target paragraph before editing.
- `WORKFLOW_REVISION`: Harscode checkout `b179360088fbf8a7dec6c285cd8c0b8992cbb4aa`; re-read canonical Build/Patch guidance at dispatch.
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `SELECTED_MODEL`: `gpt-6-luna`; Human-owned local configuration marks approval not required.
- `REASONING_EFFORT`: `high`.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Narrow patch re-entry requested by independent Code Review; after the patch, return to Review for targeted confirmation of F-01.
- `RISK_TIER`: Tier 1 — preserve the explicit boundary that Campaign close-ordering mechanism is unselected. No implementation or Tier-0 path is in scope.
- `MODEL_ROUTING_RATIONALE`: One-paragraph documentation correction against an exact accepted patch plan and approved spine. The configured low-cost Implementer-capable model at high effort is sufficient; escalate only if concrete capability insufficiency is evidenced, not to compensate for missing authority or context.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-011/techplan.md` — current-effective Human-approved `Approved` spine; preserve D1/R7 and the unselected ordering mechanism.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-001/tasks/01-donation-domain-spec-reconciliation.md` — Human-accepted Task 01 boundary; its narrow Campaign reference permits only threshold/eligibility clarification and excludes broader Slice 3 closure reconciliation.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-007/review-findings.md` — blocking F-01 and requesting-phase handoff.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-007/patch-plan.md` — sole patch contract; implement only the proposed removal/clarification of the shared-SQL-guard claim in the Campaign feature Summary.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-001/report.md` — prior Build handoff, including actual scoped checks and limitations.
- Live source to reopen: `docs/spec/4-campaign/features/09-closure.md` Summary and its D1 cross-reference; compare with `docs/spec/4-campaign/invariants.md#inv-campaign-13` and Donation `INV-donation-02` / `INV-donation-08`.
- Target-repo instructions/authority: root `AGENTS.md`, `docs/spec/README.md`, and applicable sections of `docs/kencleng-agentic-workflow.md`.
- Canonical current Harscode guidance: `workflow/3-build-prompt.md`; `workflow/3-build/guidelines.md`; `workflow/3-build/checklist.md`; `workflow/orchestrated-run-overlay.md`; `workflow/context-management.md`; `orchestration/run-contract.md`; `orchestration/AGENTS.md`; `workflow/AGENTS.md`.

## Patch target and completion condition

Resolve only F-01: remove the unsupported assertion that threshold, deadline, and force-close share a `WHERE status = 'published'` idempotency guard. Preserve the three trigger names as feature context and clarify that this narrow D1 cross-reference does not select the close-ordering mechanism; direct readers to `INV-campaign-13` for settled D1 behavior and leave mechanism selection to its authorized Campaign delivery work.

Change only `docs/spec/4-campaign/features/09-closure.md`, limited to the Summary paragraph identified by F-01. Do not edit the Campaign invariant, Donation specs, Product/MVP, Techplan, API, runtime code, tests, migrations, generated output, or orchestration projections. Do not introduce a replacement lock/isolation/SQL mechanism or expand into broader Slice 3 closure reconciliation. Keep the affected spec `draft`; no owner approval is implied.

Write this Run's `patch-report-1.md` and `launch-record.md`. Perform a focused reread against F-01, the patch plan, D1, and `INV-campaign-13`, plus `git diff --check`. This documentation correction does not call for tests; do not add or run tests. State explicitly that no race/concurrency, performance/load, or security-class test ran, and preserve all downstream Testing obligations. Do not claim `CONTRACT_READY` or accept residual risk.

After the patch, return to the requesting Review concern with targeted confirmation of F-01; the accepted patch plan says a full four-pass review is not warranted if the patch is narrow and exact. Orchestrator will prepare the required fresh Reviewer Run after this Build completes. Task 02 remains downstream until the Task 01 patch/review route is resolved.

## Execution envelope

- `PREAUTHORIZED`: read the named spine/task/finding/patch plan and live source; edit only the specified Summary paragraph; write only this Run's `patch-report-1.md` and `launch-record.md`; run the focused diff/document check stated above.
- `ORCHESTRATOR_DECISION`: widen file or behavior scope, or choose a different review route.
- `HUMAN_REQUIRED`: any Product/Design/domain/API/security decision, residual-risk acceptance, broader closure change, protected-path write, spec `agreed` status, or milestone acceptance.

## Human-assisted dispatch

Use a fresh Implementer Session in the Kencleng repository root with the configured model/effort. Start from `../harscode-workspace/workflow/3-build-prompt.md` and this Invocation. Kickoff: `Jalankan Build/Patch Run BLD-S2-002-002 sebagai Implementer sesuai Invocation durable di .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-002/invocation.md. Re-ground hanya pada Approved TP-011, Task 01, finding F-01 dan patch-plan RV-007, serta live Summary source. Buat koreksi minimal pada docs/spec/4-campaign/features/09-closure.md agar tidak mengklaim shared SQL guard sebagai mekanisme yang disetujui; pertahankan trigger context, D1, dan batas mekanisme-unselected. Jangan sentuh file lain atau memperluas scope closure Slice 3. Tulis patch-report-1.md dan launch-record.md dengan focused reread serta git diff --check; jangan tambah/jalankan tests. Setelah selesai, handoff kembali untuk targeted Review F-01; jangan klaim CONTRACT_READY atau terima residual risk.`

After dispatch, Human reports a material problem or completion. The Orchestrator reconciles from this Run's durable artifacts; no automatic Participant dispatch or background monitoring is authorized.
