# Run Invocation — `EXP-S2-008-001`

Status: `PARKED` — Human HOLD on 2026-10-05; prepared, undispatched; no Participant/Session created.

Prepared: 2026-10-05. Fresh Exploration Run for the bounded WU-S2-008 source-reconciliation work unit. Human dispatch is pending; no Participant or Session is assigned yet.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-008`
- `RUN_ID`: `EXP-S2-008-001`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-008/runs/EXP-S2-008-001`
- `ARTIFACT_TARGET`: `none` — Stage 1 creates no durable artifact; later Exploration evidence, if reached after Human confirmation, is Run-owned under `RUN_PATH/evidence/`.
- `WORK_UNIT_PATH` / compatibility `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-008`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Explorer
- `SPECIALIZATION`: Donation retry credential source/counterpart reconciliation
- `PARTICIPANT_ID`: Not assigned; Human will dispatch a new Participant.
- `PARTICIPANT_PROFILE_ID`: `KC-EXPLORER`
- `SESSION_TRANSITION`: `FRESH` — new Work Unit and independent source-reconciliation context.
- `SESSION`: Not assigned; dispatch creates a fresh Participant Session.
- `TARGET_REVISION`: Observed Kencleng HEAD `7e731f92eb052e615bd64e84902f5c6fbe61cdbd` plus working tree at HOLD capture. The prior prepared `6e78c49…` pointer was inherited from RV10 and corrected before dispatch reliance. At any explicit Human resume, verify current git/candidate/provenance and re-ground rather than treating this snapshot as current.
- `WORKFLOW_REVISION`: Current-effective Harscode workflow at dispatch; canonical prompt `workflow/1-exploration-kickoff-prompt.md`, `workflow/1-exploration/guidelines.md`, `workflow/1-exploration/sniffing-checklist.md`, `workflow/orchestrated-run-overlay.md`, and `workflow/AGENTS.md`.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned `.harscode-spaces/.local-config.yaml`.
- `MODEL_ROUTING_RATIONALE`: A cross-source Exploration must trace Donation policy and authored API wording to generated/internal consumers and exact acceptance boundaries. `gpt-6-luna` is the lowest-cost registered repository-capable model; `high` is selected for cross-boundary evidence synthesis. Escalate only if the phase exposes a concrete capability gap after context and authority routing are checked.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Canonical Harscode Exploration kickoff. Execute Stage 1 only and hard-stop for Human confirmation before Stage 2; later stages require Human confirmation per the canonical prompt.

## Task and current-effective inputs

- `TASK`: Explore the bounded reconciliation need recorded in WU-S2-008 `manifest.md` and WU-S2-003 `techplan.candidate.md` §13 Active item 9. Apply the canonical kickoff with normal variables/context; do not add a solution-steering prompt.
- Candidate target and OI9 statement: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/techplan.candidate.md`, SHA-256 `e895a1da8b90e7f88c449651a9a46add59e1a1d610cce3cc7b739d0c12c30315`.
- Current Review boundary: WU-S2-003 `runs/RV-S2-003-010/review-findings-1.md` SHA-256 `c70564f6219554f210f45f9906213c2050369faafe6d7a2cf35bf670c4a07171`; `launch-record.md` SHA-256 `68d89f04775365c4dc17138e1c676b1be8e5f4a2bccc9f7fdf3de69f22960141`. RV10 is candidate Review evidence only; it is not source acceptance or migration-design verdict.
- Settled Human decision: parent `events.md`, 2026-10-04 event, and candidate §13 item 9. Do not reopen or expand this decision.
- Exact prior API acceptance boundary: WU-S2-007 `handoff-to-WU-S2-003.md`, SHA-256 `cf468693dc384ccad5a4156c43d433a3ff725a6169869c3f171aeff0a27325eb`; accepted hashes cover Funding-unavailable 503 behavior only.
- WU-S2-006 acceptance boundary: WU-S2-006 manifest and its RV-S2-006-006 acceptance snapshot; the seven hashes do not cover retry credential issuance.
- Applicable authorities: root `AGENTS.md`; relevant sections of `docs/kencleng-agentic-workflow.md`; `api/README.md`; current Donation Feature 01, Feature 02, and invariants; authored Donation OpenAPI plus referenced common components; `frontend/AGENTS.md` only if current impact reaches frontend counterparts; WU-S2-003 parent state and current tracker.

## Execution envelope

- Stage 1 is read-only routing/checkpoint work and stops for Human confirmation before Stage 2.
- No authored source, API, generated artifact, frontend consumer, test, migration, database, runtime, browser, or security changes/actions are authorized by this Invocation.
- No source acceptance or exact hash receipt is authorized by Exploration. Any later source change requires its own applicable Harscode Run and explicit reviewed file boundary.
- Do not dispatch this Run. Human performs dispatch and any later Stage 2 confirmation.

## Canonical kickoff

Use `../harscode-workspace/workflow/1-exploration-kickoff-prompt.md` with its normal variables/context and the task above. Also load the orchestrated-run overlay. Follow its Stage 1 hard stop exactly. No task-specific conclusion or source-change plan is prescribed.
