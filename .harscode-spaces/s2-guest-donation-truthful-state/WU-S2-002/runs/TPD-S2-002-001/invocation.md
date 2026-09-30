# Run Invocation — `TPD-S2-002-001`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-09-30 by Orchestration Operator. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `TPD-S2-002-001`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-001`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `ROLE`: Planner
- `SPECIALIZATION`: Post-approval Techplan decomposition gate
- `PARTICIPANT_ID`: `P-S2-002-PLD-001-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-PLANNER`; current Profile file SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — new Planner Run after approval/status reconciliation Run `TP-S2-002-013`.
- `TARGET_REVISION`: Kencleng checkout `6891341a050982e14174ab5af132a200f24e71d9` plus the current durable working-tree artifacts; re-open them at dispatch. The checkout has pre-existing changes; do not discard or overwrite them.
- `WORKFLOW_REVISION`: Harscode checkout `b179360088fbf8a7dec6c285cd8c0b8992cbb4aa`; re-read the canonical prompt and guidance at dispatch.
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `SELECTED_MODEL`: `gpt-6-luna`; Human-owned local runtime configuration marks approval not required.
- `REASONING_EFFORT`: `high`.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Optional post-approval Techplan decomposition gate. Apply the canonical Step 0 decision before generating tasks. Current orchestration permits unaffected Slice 2 source/spec work while O1/O11-dependent details remain gated; do not infer that decomposition must pass or fail.
- `MODEL_ROUTING_RATIONALE`: The Planner must assess a bounded execution-context/dependency question against the approved spine and current repo layout. Configured non-escalation model at high effort is sufficient; missing authority or a material contract gap is a route back to Human/Orchestrator, not a reason to escalate models.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-011/techplan.md` — current-effective Human-approved Techplan (`Approved`); sole planning spine for the decomposition gate.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-012/report-techplan.md` and `TP-S2-002-013/launch-record.md` — approval/report and status-reconciliation provenance.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/manifest.md`, `../work-graph.md`, and `../events.md` — current Work Unit and scoped O1/O11 state.
- `docs/spec/README.md`, `api/README.md`, root `AGENTS.md`, and the paths/requirements named in TP-011 §§9–13 — only as needed to judge whether execution boundaries/dependencies are real. Re-open current sources; do not assume historical spec/API details are authority.
- `.harscode-spaces/participant-profiles/profiles.md#kc-planner` and `.harscode-spaces/.local-config.yaml` — current Profile/runtime routing.
- Canonical current Harscode entrypoint `../harscode-workspace/workflow/2-3-techplan-decomposition-prompt.md`; also read `workflow/2-techplan/rules.md` §10, `workflow/orchestrated-run-overlay.md`, and applicable `workflow/AGENTS.md`.

## Task and completion condition

Apply the canonical post-approval Techplan decomposition prompt to TP-011. Explicitly answer STEP 0 first. If decomposition is not genuinely useful, record the gate outcome in this Run's `launch-record.md`, create no task files, and recommend Build from the Approved Techplan. If decomposition is useful, choose and justify the least-surprising axis, create executable task file(s) and a manifest under this Run's `tasks/` directory, and preserve TP-011 as the complete authoritative spine. Do not predetermine the gate result, task count, boundaries, or splitting axis.

Any generated task must remain within approved TP-011, preserve all material requirements, rules, decisions, risks, interface/data contracts, verification obligations, and Open Items, and distinguish tasks that can proceed from O1 `AUTHORITY_SYNC` or O11 `HUMAN_DECISION`-dependent details. Do not invent owner decisions, mechanisms, product meaning, API shapes, risk acceptance, or a `CONTRACT_READY` claim. If the approved spine lacks a material decision needed to define a task safely, stop and report that gap; do not amend TP-011 in this Run.

Use the canonical decomposition artifact requirements. The orchestrated path mapping for this project is: task files and decomposition manifest, if warranted, go under `RUN_PATH/tasks/`; record the Step 0 result and phase handoff in `RUN_PATH/launch-record.md`. Keep the parent Techplan unchanged. Do not create a progress ledger.

Do not edit Product/MVP, Design, authority, Techplan/report, source specs, OpenAPI, runtime code, tests, or other orchestration projections. Do not start Build or claim `CONTRACT_READY`. If task files are generated, stop at the Human split-review gate; Orchestrator will reconcile the outcome before preparing Build dispatch.

## Human-assisted dispatch

Use a fresh Planner Session in the Kencleng repository root with the configured model/effort. Start from `../harscode-workspace/workflow/2-3-techplan-decomposition-prompt.md` and this Invocation. Kickoff: `Jalankan post-approval Techplan decomposition Run TPD-S2-002-001 sesuai Invocation durable di .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-001/invocation.md dan canonical Techplan Decomposition prompt Harscode saat ini. Terapkan STEP 0 secara eksplisit pada Approved TP-S2-002-011; jangan mengasumsikan decomposition wajib. Jika berguna, hasilkan task file dan manifest yang memenuhi prompt di RUN_PATH/tasks/; jika tidak, catat keputusan gate di launch-record.md tanpa task file. Jangan ubah Techplan, authority, spec/API, code/test, atau state orchestration; jangan mulai Build atau klaim CONTRACT_READY. Tulis hanya artifact Run yang diizinkan Invocation ini lalu berhenti.`

After dispatch, Human reports a material problem or completion. If a split is generated, Human reviews it before any Build Run; Orchestrator then reconciles and prepares the justified next route. No automated Participant dispatch or background monitoring is authorized.
