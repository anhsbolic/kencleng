# Run Invocation — `BLD-S2-004-001`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-03 by Orchestration Operator. Human-facing prose: Bahasa Indonesia; retain canonical Harscode terms and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-004`
- `RUN_ID`: `BLD-S2-004-001`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/BLD-S2-004-001`
- `ARTIFACT_TARGET`: `none` — implementation changes belong to current scoped frontend sources; Run-owned execution evidence is `RUN_PATH/report.md` and `RUN_PATH/launch-record.md`.
- `WORK_UNIT_PATH` / compatibility `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Implementer
- `SPECIALIZATION`: None; use `KC-IMPLEMENTER` for frontend Build/Patch.
- `PARTICIPANT_ID`: `P-S2-004-BLD-001-1`
- `PARTICIPANT_PROFILE_ID`: `KC-IMPLEMENTER`; `profiles.md` SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`; Registry SHA-256 `804856c5defad721678be4bac145fc2a1070307cccd720326cc598b8159d7abe`.
- `SESSION_TRANSITION`: `FRESH` — new Build phase and independent execution context after Planner and Human approval.
- `TARGET_REVISION`: Kencleng HEAD `bb69cd002b3f1a1056837affcd77bb2b001007b0` plus current working tree; re-ground approved plan and live source anchors before editing.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; current ordinary phase guidance remains authoritative.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: Frontend implementation spans the guest donation flow, contract-faithful mocks, and sensitive status-credential handling. The Human-declared low-cost `gpt-6-luna` covers coding/repository work; high effort is selected as sufficient for the approved multi-surface implementation and privacy-sensitive behavior. Do not escalate to approval-gated `gpt-6-sol` unless execution establishes a material capability gap; missing context or authority is not a model gap.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Initial Build/Patch of the exact Human-approved frontend Techplan `TP-S2-004-003`. No decomposition/task file exists. One Build Run for the whole approved WU004 spine.

## Approval and current-effective inputs

- Human approval receipt: 2026-10-03, Anhar: “aku setuju dengan techplan nya bro. lanjut”. The report presented for decision is `TP-S2-004-004/report-techplan.md`, SHA-256 `ada63615525ab061ceec4ada72414e07fbc8d13e6636a1d91203a19f6b9fbac2`.
- Approved Techplan: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-004/runs/TP-S2-004-003/techplan.md`, current SHA-256 `e6a93bc65a760084955f63a469b238fd70a80c20ff38cc1d22ba1be000ad99eb`; `Status: Approved`. Approval candidate before status-only propagation SHA-256 `187f1f1b827eab39dbed7e5c46c3e62727fffaddc2b942991602184048171c03`; Orchestrator reconstructed that exact preimage and verified only the `Status` marker changed. Approval applies to that plan's semantic content; plan content must otherwise remain unchanged.
- Approval report and resolution handoff remain derivation/provenance: report above; `TP-S2-004-003/handoff.md` SHA-256 `d82c760d7f2f73cb8d4f8941e5670a3148876dcf7cd291b57e71f05d9f71e112`. `RV-S2-004-001` reviewed predecessor TP004002, not this target; F01 is resolved non-materially in TP004003. Do not claim direct Review of TP004003.
- Current WU-S2-006 manifest and Parent Work Graph satisfy the hard source/counterpart readiness dependency; the exact accepted API/generated types and frontend contract fixture counterparts are in place. Backend response DTO/wire assertion remains separately assigned to WU003.
- Read root `AGENTS.md` and `frontend/AGENTS.md` before edits. Re-read current approved scope/sequencing in `docs/product/mvp-scope.md` and `docs/product/mvp-delivery-slices.md`; use `docs/ui-ux/README.md` routing, then only the applicable current design/pattern/page authorities. Reopen the relevant sections of `docs/project/kencleng-frontend-tech-stack.md`, `api/README.md`, accepted `api/openapi/campaign.yaml`, `api/openapi/donation.yaml` plus required shared components, and the relevant accepted Campaign/Donation feature specs. Treat the plan's code anchors as coordinates and reopen live code/types/tests. Do not load raw Exploration logs by default.
- Phase guidance: `../harscode-workspace/workflow/3-build-prompt.md`, `workflow/3-build/guidelines.md`, `workflow/3-build/checklist.md`, `workflow/orchestrated-run-overlay.md`, `workflow/AGENTS.md`, and `../harscode-workspace/orchestration/run-contract.md`. Use `../harscode-workspace/orchestration/AGENTS.md` for applicable orchestration rules.

## Risk and execution envelope

- Kencleng project risk tier for this Run: **Tier 1**. The frontend receives and transports a guest status credential and handles optional PII/email opt-in; the Techplan also governs high-impact amount/currency disclosure. Implement within UI scope; do not implement server credential controls, money/cap enforcement, or security authority. Tier 1 evidence, independent Review/Testing, and Human review remain required before merge.
- `PREAUTHORIZED`: read and edit only `frontend/` sources needed to execute approved `TP-S2-004-003`; create/update frontend-scoped component, route, adapter, MSW, fixture, and observable-test files required by §11; run focused Build-loop verification supported by live frontend scripts; write only `RUN_PATH/report.md` and `RUN_PATH/launch-record.md` as Run-owned evidence. Read allowed shared authority and contract sources.
- `HUMAN_REQUIRED`: any material product/design/API/security/architecture decision, scope expansion, protected/Tier-0 write, migration/manual DB/index action, or residual-risk acceptance. Stop and report if live code invalidates a material approved assumption; do not silently redesign.
- Out of scope: all `backend/`, authored or generated `api/`, Product/spec/Design authority sources, `docs/`, manifests/events/Work Graph/Control Surface/tracker, and other Work Units. No frontend shared/UI abstraction unless live need justifies it and the applicable `frontend/components/README.md` governance is followed. No settlement, payment rail, mock production branch, credential-server control, eligibility/capacity enforcement, or real backend integration.
- Verification: inspect live frontend package scripts/config and plan anchors before choosing commands. Run only fast Build-loop checks needed to make this implementation credible; tests added/changed in this Run must be executed enough to prove they run and cover intended behavior. The plan assigns final checklist evidence R1–R11 primarily to independent Testing/Human; do not claim final Testing completion, rendered Human acceptance, or browser/security-class evidence. Do not run broad suites, race/performance/load/security sweeps, migration/runtime/integration tests, or browser automation merely for extra confidence. Report exact commands/results and explicitly confirm the heavy verification boundary.
- Do not mark `FRONTEND_MOCK_VERIFIED`, complete WU-S2-004, approve your own work, or claim a delivery/runtime milestone.

## Task and completion condition

Execute the whole approved Techplan: public Campaign cap/IDR disclosure and accepted donation entry; guest form and QRIS-only sandbox request; stable idempotency/recovery; truthful result/status view; safe URL-fragment-to-header status credential handling; contract-faithful MSW handlers/fixtures and focused observable tests. Use generated API types; keep production access pointed at real API endpoints. Display GoPay, ShopeePay, and bank transfer as unavailable/non-interactive, while only QRIS can initiate a request.

Preserve all settled contract boundaries: cap is disclosure, not remaining capacity or admission authority; backend owns eligibility, amount/capacity, idempotency persistence, outcomes, status and credential validity; frontend does not infer hidden reasons. Do not imply real payment or settlement. Preserve exact accepted error semantics and safe user-facing failures.

Before editing, re-ground on live sources. Resolve ordinary route/local presentation details from current project patterns and Sunlit Editorial / Evidence-Led Optimism authority. If any material behavior, ownership, API shape, threat control, or verification premise differs from the Approved Techplan, stop with a bounded handoff and do not choose a substitute.

Completion for this Run: implement the approved frontend target, execute and report focused Build-owned checks actually run, and write the complete canonical `RUN_PATH/report.md` plus `launch-record.md` with structured phase handoff. Independent Code Review, Testing, rendered Human acceptance, real integration, and runtime/security evidence remain subsequent gates.

## Required dispatch prompt

Jalankan Harscode Build/Patch Run `BLD-S2-004-001` sesuai Invocation ini, canonical `../harscode-workspace/workflow/3-build-prompt.md`, dan `workflow/orchestrated-run-overlay.md`. Eksekusi penuh Techplan `TP-S2-004-003` yang sudah Approved; baca panduan Build, authority frontend yang diroute, lalu re-ground semua code/type/test anchors pada kondisi live sebelum mengubah file. Implementasikan frontend-only sesuai execution envelope dan jalankan focused Build-loop checks dari scripts yang benar-benar tersedia. Tulis report lengkap beserta phase handoff dan launch record pada Run path. Berhenti dan laporkan bila ada kontradiksi material; jangan melintasi frontend, menerima risk, menyatakan milestone, atau menganggap Build ini lulus Review/Testing/rendered acceptance.
