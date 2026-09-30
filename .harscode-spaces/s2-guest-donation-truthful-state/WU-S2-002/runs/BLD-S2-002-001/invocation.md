# Run Invocation — `BLD-S2-002-001`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-09-30 by Orchestration Operator. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `BLD-S2-002-001`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-001`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `ROLE`: Implementer
- `SPECIALIZATION`: Donation domain spec reconciliation — Task 01
- `PARTICIPANT_ID`: `P-S2-002-BL-001-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-IMPLEMENTER`; current Profile file SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — new Implementer Run after Human acceptance of the decomposition split and completed Planner Run `TPD-S2-002-001`.
- `TARGET_REVISION`: Kencleng checkout `6891341a050982e14174ab5af132a200f24e71d9` plus current durable working-tree artifacts; re-open task and authority sources at dispatch. The checkout has pre-existing orchestration changes; preserve them.
- `WORKFLOW_REVISION`: Harscode checkout `b179360088fbf8a7dec6c285cd8c0b8992cbb4aa`; re-read canonical Build guidance at dispatch.
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `SELECTED_MODEL`: `gpt-6-luna`; Human-owned local runtime configuration marks approval not required.
- `REASONING_EFFORT`: `high`.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Initial Build/Patch for accepted decomposition Task 01. Execute only the task file plus the complete Approved TP-011 spine. Task 02 is not in this Run.
- `RISK_TIER`: Tier 1 — this spec work carries payment, money-state, concurrency, guest-privacy, and security requirements derived from the threat surface. This is not authorization to change Tier-0 implementation. Use independent review and applicable Human/domain-owner review before treating affected specs as `agreed`.
- `MODEL_ROUTING_RATIONALE`: Task 01 reconciles multiple domain-spec artifacts against the approved product and delivery spine, with scoped security/payment/concurrency boundaries. The configured non-escalation model at high effort is sufficient; escalate only for demonstrated capability insufficiency, not missing authority or context.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-011/techplan.md` — current-effective Human-approved `Approved` spine; preserve all requirements, rules, decisions, risks, verification obligations, and Open Items.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-001/tasks/01-donation-domain-spec-reconciliation.md` — accepted current task slice; source of this Run's exact scope, anchors, verification, and NOT-in-task boundaries.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-001/tasks/manifest.md` — accepted task sequence; Task 02 is downstream and is not an input to execute here.
- `.harscode-spaces/s2-guest-donation-truthful-state/events.md` — explicit Human acceptance of this split and prior owner decisions; do not ask for D1/O7 again.
- Current authority sources routed by root `AGENTS.md`: `docs/product/README.md`, applicable Product/MVP and Design authorities; `docs/spec/README.md`; root `AGENTS.md`; applicable section(s) of `docs/kencleng-agentic-workflow.md`.
- Canonical Harscode entrypoint `../harscode-workspace/workflow/3-build-prompt.md`; read `workflow/3-build/guidelines.md`, `workflow/3-build/checklist.md`, `workflow/orchestrated-run-overlay.md`, and applicable workflow/orchestration AGENTS guidance.
- `.harscode-spaces/participant-profiles/profiles.md#kc-implementer` and `.harscode-spaces/.local-config.yaml` — current Profile/runtime routing.

## Task and completion condition

Execute only accepted Task 01: reconcile current Slice 2 Donation domain specifications against the Approved Techplan and current Product/MVP and Design authority. Reopen the live spec sources at the task's anchors. Use `docs/spec/README.md`'s domain-first structure and classify historical details `KEEP`, `ADAPT`, `REPLACE`, or `DEFER` with traceable authority/evidence. Reconcile the applicable Donation invariants, threat model, tasks, and feature acceptance; make only the narrow Campaign threshold/eligibility cross-reference allowed by Task 01 if it is needed to preserve D1/R7. Preserve the Approved Techplan as the full authority spine.

Keep O1 `AUTHORITY_SYNC`, O11 `HUMAN_DECISION`, O2–O5 owner/security details, and conditional O8 explicit and scoped. Do not choose amount encoding/precision, simulator timing/outcome policy, email cap/retention resolution, token/anti-enumeration controls or residual-risk acceptance, or API shapes. Carry settled D1/O6, O9, and O7 exactly; do not select transaction/locking mechanisms or import Slice 3 behavior. If live authority contradicts the spine or a material detail is missing, stop only the affected work and report the gap instead of resolving it here.

Authorized target writes are limited to `docs/spec/5-donation/` files named or conditionally authorized by Task 01 and, only when strictly necessary within its approved boundary, the narrow Campaign reference in `docs/spec/4-campaign/invariants.md` or `docs/spec/4-campaign/features/09-closure.md`. Do not modify Product/MVP or Design authority, Techplan/decomposition artifacts, OpenAPI, code, tests, migrations, generated outputs, or orchestration projections. Do not modify any protected/Tier-0 path. Preserve or use `draft` status where Human/project-authority review is still required; do not mark a spec `agreed` without the applicable recorded owner acceptance.

Write this Run's `report.md` and `launch-record.md` under `RUN_PATH` following the canonical Build report format. Perform focused spec-to-Q/R traceability and diff checks needed for a credible handoff. Do not add or run tests for this documentation-only task; no runtime, race/concurrency, performance/load, or security-class testing belongs in this Run. The report must explicitly distinguish checks actually performed from deferred owner review and downstream runtime Testing, and confirm the heavyweight verification boundary.

Do not execute Task 02, claim `CONTRACT_READY`, accept residual risk, or promote a delivery milestone. After a complete Build handoff, Orchestrator will prepare an independent Code Review Run; applicable Human/domain-owner review remains required before affected specs are treated as `agreed`.

## Execution envelope

- `PREAUTHORIZED`: read the current inputs and authority sources; update only the Task 01 spec paths above; write only this Run's `report.md` and `launch-record.md`; perform scoped manual traceability/diff checks.
- `ORCHESTRATOR_DECISION`: widen task scope, add a file outside the task boundary, or alter the phase route.
- `HUMAN_REQUIRED`: any new Product/Design/Security/API/domain decision, changed approval scope, residual-risk acceptance, protected-path write, or Human/project-owner agreement of a spec where required.

## Human-assisted dispatch

Use a fresh Implementer Session in the Kencleng repository root with the configured model/effort. Start from `../harscode-workspace/workflow/3-build-prompt.md` and this Invocation. Kickoff: `Jalankan Build/Patch Run BLD-S2-002-001 sebagai Implementer sesuai Invocation durable di .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-001/invocation.md. Gunakan Approved TP-S2-002-011 sebagai spine dan hanya Task 01 yang telah diterima Human; buka ulang live spec serta authority yang dirujuk. Rekonsiliasi hanya scope docs/spec yang diizinkan Task 01; jaga O1/O11 dan O2–O5/O8 sebagai gate pada detail terdampak. Jangan ubah API/code/test/Product/Design/Techplan/orchestration, jangan mengarang keputusan atau menerima risk, jangan tandai spec agreed tanpa review owner, jangan klaim CONTRACT_READY. Tulis report.md dan launch-record.md sesuai canonical Build prompt, lakukan hanya traceability/diff check yang relevan, lalu berhenti untuk independent Code Review.`

After dispatch, Human reports a material problem or completion. The Orchestrator reconciles from this Run's durable artifacts; no automatic Participant dispatch or background monitoring is authorized.
