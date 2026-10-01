# Run Invocation — `BLD-S2-002-006`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-01 by Orchestration Operator. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `BLD-S2-002-006`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-006`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `ROLE`: Implementer
- `SPECIALIZATION`: Task 02 authored Donation OpenAPI reconciliation
- `PARTICIPANT_ID`: `P-S2-002-BL-006-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-IMPLEMENTER`; current Profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — new Task 02 Implementer Run after Task 01's completed Review loop.
- `SESSION_TRANSITION_REASON`: Task boundary and current-effective plan/task refresh; prior BLD-003 used superseded TP-011/task snapshots and made no authored API changes. Start with a new Participant/Session grounded on Approved TP-015, TPD-002 Task 02, and completed Task 01 output/review.
- `TARGET_REVISION`: Kencleng checkout `650e73c5d646c29c0ddf1931618f02685d15f7b7` plus current Task 01 source diff and orchestration evidence; reopen live authored API/spec/authority files at dispatch.
- `WORKFLOW_REVISION`: Harscode checkout `33b03a3f62cc3aacba6534b8a011465613c64b09`; re-read current Build guidance at dispatch.
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `SELECTED_MODEL`: `gpt-6-luna`; Human-owned local configuration marks approval not required.
- `REASONING_EFFORT`: `high`.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: First current-effective Build for Human-accepted Task 02. The hard Task 01 Build/Review dependency is met, and Anhar's 2026-10-01 acceptance of all five current Task 01 specs is recorded in `events.md` and their headers. Reconcile only authored Donation OpenAPI details supported by Approved TP-015 and current authority. This acceptance does not accept residual risk or earn `CONTRACT_READY`. Return changed API source through independent Code Review; if no source change is warranted, report that outcome without inventing a diff.
- `RISK_TIER`: Tier 1 — API contract covers money, simulated settlement, guest PII/status credentials, anti-enumeration, and Campaign ordering.
- `MODEL_ROUTING_RATIONALE`: This is a bounded, cross-field authored-contract change with required schema validation and possible documented bundle/type regeneration. `gpt-6-luna` is the Human-configured low-cost Implementer-capable model; `high` is the previously established sufficient effort for this same Task 02 contract scope. Escalate only on demonstrated capability insufficiency, not to decide missing authority.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/techplan.md` — current-effective Human-approved `Approved` spine; sole material planning authority.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-002/tasks/02-donation-openapi-reconciliation.md` and `tasks/manifest.md` — current Task 02 scope, hard dependency, accepted ordering, verification, and boundaries.
- Task 01 output: the five live `docs/spec/5-donation/` files, now `agreed` after explicit Human/domain-owner acceptance recorded on 2026-10-01; Build reports `runs/BLD-S2-002-004/report.md` and `runs/BLD-S2-002-005/patch-report-1.md`; full Review findings/patch plan `runs/RV-S2-002-011/review-findings-1.md`; targeted closure `runs/RV-S2-002-012/review-confirmation.md` and `launch-record.md` (F-001 resolved; Task 01 Review loop complete). Agreement does not accept residual risk or claim runtime/security evidence.
- Current authority: root `AGENTS.md`; `docs/product/README.md`, `docs/product/mvp-scope.md`, `docs/product/mvp-delivery-slices.md` §§5–6; relevant `docs/ui-ux/README.md`; `docs/spec/README.md`; `docs/project/kencleng-monetary-data-standard.md`; `.harscode-spaces/authority-map.md`; and applicable sections of `docs/kencleng-agentic-workflow.md`.
- API authoring/verification: `api/README.md`; `api/openapi/donation.yaml`; only referenced portions of `api/openapi/common.yaml`; `api/openapi/index.yaml`; `frontend/AGENTS.md` and `frontend/package.json` if generated API types are produced. Do not load `api/openapi.yaml` in full unless aggregate/generated correspondence requires it.
- Canonical current Harscode: `../harscode-workspace/workflow/3-build-prompt.md`, `workflow/3-build/guidelines.md`, `workflow/3-build/checklist.md`, `workflow/orchestrated-run-overlay.md`, `workflow/context-management.md`, `workflow/AGENTS.md`, `orchestration/run-contract.md`, and `orchestration/AGENTS.md`.

## Build target and boundaries

Execute only Task 02 from the current Approved spine and snapshot. Reopen the live authored API sources and current Donation specs before editing; classify historical detail against authority rather than inheriting it. Express only supported Slice 2 contract behavior, including:

- major-unit decimal-string amounts with explicit currency code and exact-decimal/no-float semantics; whole-IDR input remains minimum Rp5,000 in Rp1 increments;
- QRIS-only sandbox submission and backend-owned persisted status; no real payment rail or donor-selected terminal outcome;
- request idempotency distinct from settlement replay; D1/D15 eligibility, full accepted-pending settlement, atomic exactly-once success/funding, stable close reason, and threshold may be exceeded;
- settled O4 flow: fragment-carried guest status URL, frontend handoff/URL cleanup, difficult-to-guess bearer credential, one-way HMAC verifier, hard 24-hour expiry from issuance, and status-only projection;
- settled O5 public failure: absent Donation or missing/wrong/expired credential gives a uniform `404`, identical Problem Details body/headers/cache behavior, including `Cache-Control: private, no-store`;
- O11/D19 terminal-notice eligibility and bounded/recoverable terminalization direction without selecting a numeric bound, architecture, timeout-as-`failed`, or deletion/race policy;
- bounded O8 clearance permits replacement of the historical operations only within the current repository/Slice 2 scope supported by the Human/API-owner evidence.

Keep concrete currency/range/fraction/storage parameters, O2 simulator timing/mechanism, O3 lifecycle/security implementation controls, O4 credential implementation/evidence, O5 empirical parity/timing/abuse, and residual-risk acceptance open where TP-015 says so. Stop and report any material decision gap; do not amend TP-015 or make a Product/Design/Security/API-owner decision in Build.

Authorized source writes: `api/openapi/donation.yaml`; `api/openapi/common.yaml` only for a required, authority-supported shared component; `api/openapi/index.yaml` mechanically only if paths are added/removed and the bounded O8 condition is satisfied. If authored sources change, use `api/README.md` to regenerate `api/openapi.yaml` and frontend API types (`cd frontend && npm run generate:api-types`); never hand-edit generated outputs. Do not edit Product/MVP, Design, domain specs, Techplan, runtime code, tests, migrations, Tier-0 paths, or orchestration projections.

Run the Task-mandated `cd api && npm run validate`. Do not add/run product or runtime tests in this contract-only Build. Report exact commands/results and distinguish schema validation/generation from owner acceptance, Security evidence, rendered acceptance, downstream Testing, and `CONTRACT_READY`. Keep Task 01 specs `draft`; do not claim Human/domain-owner acceptance, residual-risk acceptance, runtime correctness, `CONTRACT_READY`, or a delivery milestone.

Write only this Run's `report.md` and `launch-record.md`. If API source changes, recommend fresh independent Code Review of the actual authored and generated diff. If no source changes are needed, explain why and stop at the handoff; do not dispatch Review for a nonexistent diff.

## Execution envelope

- `PREAUTHORIZED`: read the named authorities, current Task 02 snapshot, Task 01 output/review evidence, live API/spec sources, and current Build guidance; edit only the authorized authored API source paths; run `cd api && npm run validate` and documented generators when source changes; write this Run's report and launch record.
- `ORCHESTRATOR_DECISION`: widen target files/scope or change phase route.
- `HUMAN_REQUIRED`: resolve any new Product/Design/Security/API authority question, choose deferred monetary or O2/O3/O4/O5 detail, accept residual risk, claim Task 01 draft acceptance, write a protected path, or promote a milestone.

## Human-assisted dispatch

Use a fresh Implementer Session from the Kencleng repository root. Start from `../harscode-workspace/workflow/3-build-prompt.md` and this Invocation. Kickoff:

`Jalankan Build Run BLD-S2-002-006 sebagai Implementer sesuai Invocation durable di .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-006/invocation.md dan canonical Build guidance Harscode saat ini. Rekonstruksi dari Approved TP-015, snapshot Task 02 TPD-002, live authority/authored Donation OpenAPI, dan Task 01 specs yang review loop-nya selesai melalui RV-012 serta telah Human-accepted (`agreed`); buka kembali sumber live sebelum edit. Majukan hanya field/detail contract yang didukung authority, bawa keputusan settled O1/O4/O5/O8/O11/D1 dengan tepat, dan pertahankan deferred parameters/evidence. Edit hanya source path yang diizinkan; jalankan `cd api && npm run validate`, serta bundle dan generated types sesuai README jika source berubah. Jangan ubah Product/Design/spec/Techplan/runtime/tests/migrations/Tier-0/projections; jangan klaim residual-risk acceptance, CONTRACT_READY, atau delivery milestone. Tulis report.md dan launch-record.md dengan hasil aktual lalu berhenti untuk handoff.`

After dispatch, reconcile from this Run's durable artifacts. Task 01 Human/domain-owner acceptance is recorded; do not infer from it Security/PII residual-risk acceptance, runtime proof, or `CONTRACT_READY`.
