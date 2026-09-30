# Run Invocation — `BLD-S2-002-003`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-09-30 by Orchestration Operator. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `BLD-S2-002-003`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-003`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `ROLE`: Implementer
- `SPECIALIZATION`: Task 02 authored Donation OpenAPI reconciliation
- `PARTICIPANT_ID`: `P-S2-002-BL-003-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-IMPLEMENTER`; current Profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — new Implementer Run after Task 01's completed targeted independent review confirmation `RV-S2-002-008`.
- `SESSION_TRANSITION_REASON`: Fresh task transition — start the next accepted decomposition task with a new Participant and Session, re-grounded on current-effective domain specs and the Approved spine.
- `TARGET_REVISION`: Kencleng checkout `6891341a050982e14174ab5af132a200f24e71d9` plus completed Task 01 source diff and Run artifacts; reopen all current API/spec/authority sources at dispatch.
- `WORKFLOW_REVISION`: Harscode checkout `b179360088fbf8a7dec6c285cd8c0b8992cbb4aa`; re-read canonical Build guidance at dispatch.
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `SELECTED_MODEL`: `gpt-6-luna`; Human-owned local configuration marks approval not required.
- `REASONING_EFFORT`: `high`.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: First Build for Human-accepted Task 02 after Task 01's build, full Code Review, F-01 patch, and targeted Review confirmation. Advance authored API details supported by current authority; keep only dependent fields/details gated.
- `RISK_TIER`: Tier 1 — API contract covers payment simulation, money, settlement state, guest PII, status credentials, anti-enumeration, and Campaign ordering. No runtime behavior or protected implementation is authorized.
- `MODEL_ROUTING_RATIONALE`: Task-scoped contract reconciliation across one authored domain OpenAPI file, optional referenced shared components, path registry, and generated projections. The configured low-cost Implementer-capable model at high effort is sufficient; escalate only if concrete capability insufficiency is evidenced, not for missing authority or context.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-011/techplan.md` — current-effective Human-approved `Approved` spine; full authority for Q1–Q11, R1–R11, D1–D16, RISK-1–RISK-10, §12, and O1–O5/O8/O11 boundaries.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-001/tasks/02-donation-openapi-reconciliation.md` — Human-accepted Task 02 execution contract; hard dependency on Task 01 is satisfied by completed Build and Review confirmation.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-001/tasks/manifest.md` — accepted dependency/sequence; Task 02 is now the next task.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-001/report.md` and `runs/BLD-S2-002-002/patch-report-1.md` — Task 01 Build/Patch handoffs.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-007/review-findings.md` and `runs/RV-S2-002-008/review-confirmation.md` — Task 01 independent review and targeted F-01 closure; domain spec outputs remain `draft`, not owner/Human-agreed.
- Current authority sources: root `AGENTS.md`; `docs/product/README.md`, `docs/product/mvp-scope.md`, `docs/product/mvp-delivery-slices.md` §§5–6; `docs/ui-ux/README.md` and only relevant Design patterns; `docs/spec/README.md`; current Donation domain specs under `docs/spec/5-donation/`; `docs/spec/4-campaign/invariants.md#inv-campaign-13`; `.harscode-spaces/authority-map.md`; `.harscode-spaces/s2-guest-donation-truthful-state/events.md` for recorded O1/O11 owner gates; applicable sections of `docs/kencleng-agentic-workflow.md`.
- API authoring/verification authority: `api/README.md`; `api/openapi/donation.yaml`; only referenced portions of `api/openapi/common.yaml`; `api/openapi/index.yaml`; `frontend/AGENTS.md` and `frontend/package.json` if generated frontend API types are produced. Do not load `api/openapi.yaml` in full unless aggregate/generated correspondence requires it.
- Canonical current Harscode guidance: `workflow/3-build-prompt.md`; `workflow/3-build/guidelines.md`; `workflow/3-build/checklist.md`; `workflow/orchestrated-run-overlay.md`; `workflow/context-management.md`; `orchestration/run-contract.md`; `orchestration/AGENTS.md`; `workflow/AGENTS.md`.

## Task and completion condition

Execute only accepted Task 02: reconcile authored Donation OpenAPI for Slice 2 guest submission and status lookup against the Approved TP-011 spine, current authority, and Task 01 outputs. Re-ground on the live authored source and classify historical contract details by current authority; preserve settled requirements/decisions, including exact-decimal/no-float and whole-IDR product rules, QRIS-only sandbox simulation, backend-owned terminal status, request idempotency distinct from settlement replay, generic status failure behavior, D1, and exact O7 guidance where it applies.

Advance only contract fields whose meaning and shape are sufficiently authorized. Keep O1 amount encoding/currency/storage/precision details and O11/O2–O3 verified-email cap versus terminal notice details gated. Preserve O2 simulator timing/scenario, O3–O5 email/security/token/parity/abuse/risk controls, and conditional O8 compatibility boundary. Do not invent schemas/defaults to close gates; stop only affected fields and record what authority/evidence is needed. Before removing or replacing any historical operation, obtain the conditional O8 consumer/distribution evidence; otherwise do not remove/replace it in this Run. Do not accept Security/PII residual risk.

Authorized source writes: `api/openapi/donation.yaml`; `api/openapi/common.yaml` only for a genuinely required, authority-supported shared component; and `api/openapi/index.yaml` mechanically if paths are added/removed with O8/owner gates satisfied. If authored source changes, regenerate `api/openapi.yaml` using the documented bundle command and regenerate frontend API types using `cd frontend && npm run generate:api-types`; never hand-edit generated files. Follow the applicable frontend guidance when writing generated types. Do not change Product/MVP, Design, domain specs, Techplan, runtime code, tests, migrations, protected Tier-0 paths, or orchestration projections.

Run `cd api && npm run validate` as required by Task 02. If source changes, follow `api/README.md` for bundle/type generation and inspect authored source, path index, generated bundle, and generated types for consistency. Do not add or run product/runtime tests in this contract-only Build. The handoff must distinguish actual schema validation/generation from owner review, Security evidence, rendered Human acceptance, and downstream runtime Testing; no API validation/generation implies `CONTRACT_READY`.

Write this Run's `report.md` and `launch-record.md` using canonical Build format. Trace active contract decisions/gates to the source files and report exact checks/commands/results. Keep donation domain specs `draft` pending applicable owner/Human acceptance. No `CONTRACT_READY`, delivery milestone, residual-risk acceptance, API consumer assumption, or runtime correctness claim is allowed from this Run.

## Execution envelope

- `PREAUTHORIZED`: read named current authorities and authored API sources; edit only the Task 02 source paths above; run the Task-mandated API validation and documented generators for authored changes; write only this Run's report and launch record.
- `ORCHESTRATOR_DECISION`: widen target files/scope or change the phase route.
- `HUMAN_REQUIRED`: any unresolved Product/Design/Security/API authority decision, O1/O11 resolution, residual-risk acceptance, removal/replacement behind O8 without evidence, protected-path write, spec agreement, or milestone acceptance.

## Human-assisted dispatch

Use a fresh Implementer Session in the Kencleng repository root with the configured model/effort. Start from `../harscode-workspace/workflow/3-build-prompt.md` and this Invocation. Kickoff: `Jalankan Build Run BLD-S2-002-003 sebagai Implementer sesuai Invocation durable di .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-003/invocation.md dan canonical Build prompt Harscode saat ini. Reconstruct dari Approved TP-011, task 02 yang diterima, current authority, serta hasil Task 01 yang sudah melalui review loop RV-007/RV-008; buka ulang authored Donation OpenAPI dan komponen yang dirujuk. Majukan hanya field contract yang sudah cukup berwenang; jaga O1, O11/O2-O3, O4/O5, dan conditional O8 sebagai gate pada detail terdampak. Jangan hapus/ganti operasi historis tanpa bukti O8; jangan ubah Product/Design/spec/Techplan/runtime/test/Tier-0/orchestration. Jalankan `cd api && npm run validate`; bila source berubah, ikuti workflow bundle dan generated types tanpa hand-edit. Tulis report.md dan launch-record.md dengan batas verifikasi, jangan klaim CONTRACT_READY atau terima residual risk, lalu berhenti untuk handoff independent review.`

After dispatch, Human reports a material issue or completion. The Orchestrator reconciles from this Run's durable artifacts; no automatic Participant dispatch or background monitoring is authorized.
