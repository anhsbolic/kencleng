# Run Invocation — `BLD-S2-002-004`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-01 by Orchestration Operator. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `BLD-S2-002-004`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-004`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `ROLE`: Implementer
- `SPECIALIZATION`: Task 01 post-approval Donation domain-spec reconciliation to TP-015
- `PARTICIPANT_ID`: `P-S2-002-BL-004-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-IMPLEMENTER`; current Profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — new Build/Patch occurrence after TP-S2-002-017 approval reconciliation and TPD-S2-002-002 snapshot reconciliation.
- `SESSION_TRANSITION_REASON`: Fresh Build occurrence; prior Task 01 Build/review loop predates approved TP-015 O1/O11 directions and the refreshed child task. Re-ground on the current spine/task and current live specs.
- `TARGET_REVISION`: Kencleng checkout `650e73c5d646c29c0ddf1931618f02685d15f7b7` plus current durable approval, TPD-002 task snapshots, and working-tree artifacts; reopen live sources at dispatch.
- `WORKFLOW_REVISION`: Harscode checkout `33b03a3f62cc3aacba6534b8a011465613c64b09`; re-read canonical Build guidance at dispatch.
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `SELECTED_MODEL`: `gpt-6-luna`; Human-owned local configuration marks approval not required.
- `REASONING_EFFORT`: `high`.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Fresh Build/Patch re-entry for accepted Task 01 because the existing Donation draft specs predate the materially approved TP-015 O1/O11 directions. After this Build handoff, route the actual diff through independent Code Review. Human/domain-owner review and acceptance of the current drafts remains a parallel, independent gate required before `CONTRACT_READY`; it is not a second split-approval gate. Task 02 may advance only after its Task 01 hard dependency is satisfied and any field-specific authority gates are met.
- `RISK_TIER`: Tier 1 — domain specifications cover money, guest PII, donation state, status credentials, and campaign-funding semantics; this Run changes documentation only and must preserve unresolved owner/evidence gates.
- `MODEL_ROUTING_RATIONALE`: Focused but cross-document reconciliation of approved requirements across Donation invariants, threat model, task list, and active feature acceptance. `gpt-6-luna` at `high` is the configured Implementer-capable minimum sufficient routing; escalate only for demonstrated capability insufficiency, not for missing authority.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/techplan.md` — current-effective Human-approved `Approved` spine; sole material authority for scope, decisions, risks, verification ownership, and Open Items.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-002/tasks/01-donation-domain-spec-reconciliation.md` — current Task 01 snapshot reconciled to TP-015; contains exact execution scope, references, and write boundaries.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-002/tasks/manifest.md` — preserved accepted Task 01 → Task 02 sequence and dependency; this Run is Task 01 re-entry, not Task 02.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-002/launch-record.md` — confirms Step 0, refreshed snapshots, preserved accepted topology/dependency, and no new split approval requirement.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-001/report.md`; `runs/BLD-S2-002-002/patch-report-1.md`; `runs/RV-S2-002-007/review-findings.md`; `runs/RV-S2-002-008/review-confirmation.md` — prior Task 01 implementation/review history. It is evidence/context, not a substitute for reconciling the current approved spine.
- Current live spec anchors named by Task 01: `docs/spec/5-donation/invariants.md`, `threat-model.md`, `tasks.md`, `features/01-submit-donation-settlement.md`, `features/02-donation-status-check.md`, and only any other Donation feature file whose affected language must be corrected. Reopen all files before editing.
- Narrow Campaign reference, only if live inspection proves necessary within Task 01: `docs/spec/4-campaign/invariants.md#inv-campaign-13` and `docs/spec/4-campaign/features/09-closure.md`. Do not expand into Campaign lifecycle work.
- Authorities: root `AGENTS.md`; `docs/product/README.md`, `docs/product/mvp-scope.md`, `docs/product/mvp-delivery-slices.md` §§5–6; `docs/project/kencleng-monetary-data-standard.md`; `docs/ui-ux/README.md` and relevant pattern sections; `docs/spec/README.md`; `.harscode-spaces/authority-map.md`; recorded owner decisions in `.harscode-spaces/s2-guest-donation-truthful-state/events.md`; relevant sections of `docs/kencleng-agentic-workflow.md`.
- Canonical Harscode guidance: `workflow/3-build-prompt.md`, `workflow/3-build/guidelines.md`, `workflow/3-build/checklist.md`, `workflow/orchestrated-run-overlay.md`, `workflow/context-management.md`, `workflow/AGENTS.md`, `orchestration/run-contract.md`, `orchestration/AGENTS.md`.

## Task and completion condition

Execute only refreshed Task 01: reconcile the affected Slice 2 Donation domain specifications to current Approved TP-015 and current Product/MVP, Design, monetary, and domain authority. Preserve prior reviewed work where still correct, and patch only the affected passages needed to align the existing draft specs with the current snapshot.

Carry TP-015's settled shared monetary representation: major-unit decimal string plus explicit currency code for API/wire values and exact-decimal calculation/persistence, while current Slice 2 input remains whole-IDR Rupiah, minimum Rp5.000 and Rp1 increments. Keep only supported concrete currency/range/fraction/scale/storage details open; do not restate the representation direction as unresolved or infer additional precision, currencies, migrations, or rounding.

Carry D19/O11: an opted-in verified email remains eligible until the required terminal status notice is fulfilled, and Donation Delivery must provide bounded, recoverable terminalization. Preserve unresolved O2/O3 mechanics, numeric bound, architecture, timeout meaning, security controls, retention/deletion race evidence, and residual-risk gate. Do not claim any of these are implemented or accepted.

Carry the settled O4/O5 behavior in applicable domain acceptance: fragment-carried status URL with frontend handoff and URL cleanup; one-way HMAC verifier; hard 24-hour expiry from issuance; status-only result; absent Donation and missing/wrong/expired credential share a uniform public `404`, identical Problem Details body/headers/cache behavior, including `Cache-Control: private, no-store`. Keep exact authored API expression with Task 02. Do not imply browser exposure protections, key/comparison controls, expiry enforcement, abuse controls, empirical parity, or Security/PII residual-risk acceptance are proven.

Preserve D1, O7 wording, simulator backend ownership, no-`float64` money, settled email verification/notification direction, Slice 2 exclusions, and all other TP-015 requirements unchanged. Retain explicit `KEEP` / `ADAPT` / `REPLACE` / `DEFER` rationale where already useful. The parent Techplan remains the full authority spine. Keep affected spec statuses `draft`; do not mark `agreed` or claim current Human/domain-owner acceptance.

## Execution envelope

- `PREAUTHORIZED`: Read the specified spine, task, prior Run evidence, live authorities, and live spec anchors; edit only the Task 01 Donation spec targets and a strictly necessary narrow Campaign cross-reference already within its scope; write this Run's `report.md` and `launch-record.md`; perform focused traceability/source reread and `git diff --check` on the exact changed documentation paths.
- `ORCHESTRATOR_DECISION`: Any write outside Task 01's bounded spec scope, materially broader Campaign change, or reroute to planning because the approved spine is insufficient.
- `HUMAN_REQUIRED`: New Product/Design/domain/API/Security decision; residual-risk acceptance; spec `agreed` status; protected-path edit; broader scope or milestone acceptance.

Do not change Product/MVP, Design authority, TP-015, task snapshots/manifest, OpenAPI, runtime code, tests, migrations, generated artifacts, protected Tier-0 paths, or orchestration projections. Do not create a new split or request split approval again. Do not choose unresolved O1/O2/O3/O4/O5 details, accept residual risk, claim runtime/security evidence or `CONTRACT_READY`, or begin Task 02.

Automated tests are outside this documentation-only Build. Do not add or run tests. In the report, record focused source/traceability checks and `git diff --check` if run; explicitly say no tests or runtime/security evidence were run. Keep all independent Testing obligations assigned by TP-015 for later.

Write only this Run's `report.md` and `launch-record.md` as Run evidence. Follow the canonical Build report template, including changed files, actual checks/results, heavyweight verification boundary, deferred/not-tested obligations, and phase handoff. Recommend fresh independent Code Review of this Build diff. After Code Review converges, Human/domain owner review and acceptance of current Donation drafts remains required. No `CONTRACT_READY` or delivery milestone follows automatically.

## Human-assisted dispatch

Use a fresh Implementer Session from the Kencleng repository root with the configured model/effort. Start from `../harscode-workspace/workflow/3-build-prompt.md` and this Invocation. Kickoff:

`Jalankan Build/Patch Run BLD-S2-002-004 sebagai Implementer sesuai Invocation durable di .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-004/invocation.md dan canonical Build prompt Harscode saat ini. Reconstruct dari Approved TP-015 dan Task 01 snapshot TPD-S2-002-002; buka ulang live Donation specs dan authority yang dirujuk. Selaraskan hanya passage spec terdampak dengan arah O1, O11/D19, serta O4/O5 yang sudah settled; pertahankan semua detail dan evidence yang masih deferred, status spec tetap draft, serta scope Task 01. Jangan mengubah OpenAPI, Product/Design/Techplan, code/test, task files, Tier-0, atau projections; jangan mengambil keputusan baru, menerima residual risk, klaim CONTRACT_READY, atau mulai Task 02. Jalankan hanya pemeriksaan traceability/diff dokumentasi yang proporsional, tanpa tests/runtime checks. Tulis report.md dan launch-record.md dengan batas evidence yang jelas; berhenti setelah handoff untuk independent Code Review.`

After dispatch, reconcile from this Run's durable artifacts. Independent Code Review uses a separate fresh Run/Participant/Session. Human/domain-owner acceptance applies to the current spec versions and may proceed independently; do not silently infer it from prior draft review or treat it as the already accepted task-split decision. It remains required for `CONTRACT_READY`.
