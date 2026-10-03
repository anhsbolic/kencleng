# Run Invocation — `TPD-S2-003-001`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-03 by Orchestration Operator after Human approved TP-S2-003-006. Human-facing prose: Bahasa Indonesia; retain canonical Harscode terms/enums and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-003`
- `RUN_ID`: `TPD-S2-003-001`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TPD-S2-003-001`
- `ARTIFACT_TARGET`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/tasks/` — stable task-file and manifest set, written only if canonical Step 0 concludes decomposition is useful.
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Planner
- `SPECIALIZATION`: Post-approval Techplan decomposition gate
- `PARTICIPANT_ID`: `P-S2-003-PLD-001-1`
- `PARTICIPANT_PROFILE_ID`: `KC-PLANNER`; `.harscode-spaces/participant-profiles/profiles.md` SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — new Planner Run after Human approval; do not reuse TP-S2-003-006/007 Participant or Session context.
- `TARGET_REVISION`: Kencleng HEAD `19d53315ac2847a03405339c0e86df8a97850761` plus current durable working-tree changes; re-ground exact approved plan and current code/authority at dispatch.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; ordinary applicable guidance remains current-effective.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: This is a bounded structural assessment of a complex backend plan with separate Donation, Campaign eligibility, concurrency, protected-write, and evidence boundaries. The available non-escalation model at high effort is sufficient to apply canonical Step 0 and preserve the approved spine. If a material authority or contract gap appears, stop and route it; do not escalate models to invent missing authority.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Optional post-Approval decomposition gate. The approved Techplan handoff says `Consider`; apply canonical Step 0 and decide whether a genuine execution/review/context benefit exists. Do not presuppose YES/NO, task count, task boundary, dependency, or splitting axis.

## Current-effective inputs / `PRIOR_ARTIFACTS`

- Current-effective approved spine: `WU-S2-003/runs/TP-S2-003-006/techplan.md`, status `Approved`, SHA-256 `b71951c86b2f74f6c9f1a164bcf43447c58ee655874deaf09275e6ff3b2316f9`.
- Human approval evidence: current report `WU-S2-003/runs/TP-S2-003-007/report-techplan.md`, SHA-256 `5fa5ef51309d8112f7725c0c185d4e9f8e94d353d6229853433ff45dc0c8e5ed`; launch record `.../TP-S2-003-007/launch-record.md`, SHA-256 `e73f7cba59f1c7a69bc94c8135e091be87421b9a333b6168e82b7f7bf87f8476`; approval and status-only hash verification are recorded in current parent Events and WU003 manifest.
- Review provenance: `WU-S2-003/runs/RV-S2-003-003/review-findings.md`, SHA-256 `b5cd4d47e714b8499ef0ae8d33062283763812d0a18cc7999f1f14d6f4e2a79d`, clean with no findings against its exact pre-reconciliation TP003006 hash `d98b3d37e3c00962075428c3e14f87d5ee59dfba46b6ae36ce35f78dcebbce7a`. Do not claim it reviewed the current Approved hash.
- Current Human-scoped gate: Anhar chose to carry Open Item 7. Authority/governance for setting or updating Organization `verified` / `has_overdue_report` remains unresolved. It blocks only those source updates and affected Campaign create/PATCH handlers; whole-plan approval is not authorization to cross that gate.
- Current Work Unit/project state: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/manifest.md`, parent Work Graph/Events/Control Surface/Outcome, and `docs/project/kencleng-development-tracker.md`.
- Root `AGENTS.md`, `backend/AGENTS.md`, relevant `docs/kencleng-agentic-workflow.md`, current Product/MVP/spec/API/architecture authority referenced by the approved Techplan, and current live code only where a proposed task boundary depends on current module ownership.
- Canonical current Harscode: `workflow/2-3-techplan-decomposition-prompt.md` in full; `workflow/2-techplan/rules.md` §10; `workflow/orchestrated-run-overlay.md`; `workflow/AGENTS.md`; `orchestration/AGENTS.md`; `orchestration/run-contract.md`.

## Task and completion condition

Apply the canonical post-approval Techplan decomposition prompt to the exact Approved TP-S2-003-006. Explicitly answer STEP 0 before generating anything.

- If decomposition is not genuinely useful, record STEP 0 `NO` and its reason in this Run's `launch-record.md`; create no task files and recommend Build from the Approved Techplan subject to its scoped gates.
- If decomposition is genuinely useful, state the concrete execution/review/context signal, choose and justify the least-surprising axis, and write complete task file(s) plus a manifest under `ARTIFACT_TARGET`. The Approved Techplan remains the complete authoritative spine. If tasks are generated, Human split review is required before Build.
- If a material Product/domain/security/authority/interface/data/verification decision needed to define a safe task is absent, stop and report the gap; do not patch the Techplan or invent a child-task answer.

Preserve Open Item 7's exact boundary: it blocks only Organization eligibility source updates and affected Campaign draft create/PATCH handlers until authority/governance is resolved. Do not assume whole-plan approval resolved it. Keep protected-path Human authorization, migration application, O3/O4/O5, DTO/exact-wire, PostgreSQL/concurrency, runtime/security, and residual-risk boundaries explicit wherever they affect task readiness. Do not invent dependencies or claim delivery readiness.

## Execution envelope

- `PREAUTHORIZED`: read the assigned approved spine, current relevant sources and canonical guidance; write this Run's `launch-record.md`; write task files/manifest under `ARTIFACT_TARGET` only if Step 0 is `YES`.
- Do not edit the Approved Techplan/report, Product/spec/API/code/tests, WU/parent projections, prior Runs, or any artifact outside this Run's launch record and conditional task target.
- Do not dispatch Build, apply migrations, write protected Tier-0 paths, accept residual risk, or claim `BACKEND_VERIFIED`/`CONTRACT_READY`.
- If current Techplan hash/status, approval provenance, current authority, or task-target state differs materially at dispatch, stop and report the discrepancy rather than write stale task artifacts.

## Human-assisted dispatch

- Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`
- Fresh Planner / `KC-PLANNER` Session; configured `gpt-6-luna` / `high`.
- Canonical kickoff: `Jalankan post-approval Techplan decomposition Run TPD-S2-003-001 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TPD-S2-003-001/invocation.md dan canonical ../harscode-workspace/workflow/2-3-techplan-decomposition-prompt.md. Terapkan STEP 0 secara eksplisit pada exact Approved TP-S2-003-006. Jangan mengasumsikan decomposition wajib atau mengarahkan hasil/axis/task count. Jika tidak berguna, catat NO dan berhenti tanpa task files. Jika berguna, hasilkan task file(s) dan manifest sesuai prompt di stable ARTIFACT_TARGET; pertahankan Approved spine dan scoped gate Open Item 7. Human split review diperlukan bila task set dibuat. Tulis launch-record dan task artifacts yang diizinkan saja; jangan edit source authority, code/test, Techplan, atau projections; jangan mulai Build atau klaim milestone. Berhenti setelah phase handoff.`
