# Run Invocation — `TP-S2-002-013`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-09-30 by Orchestration Operator. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `TP-S2-002-013`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-013`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Planner
- `SPECIALIZATION`: Human approval status reconciliation
- `PARTICIPANT_ID`: `P-S2-002-PL-013-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-PLANNER`; current Profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — new Planner Run after Human approval of TP-011 and report generation Run TP-012.
- `TARGET_REVISION`: Kencleng checkout `6891341a050982e14174ab5af132a200f24e71d9` plus referenced current Run artifacts; re-read approval event and source/report at dispatch.
- `WORKFLOW_REVISION`: Harscode checkout `b179360088fbf8a7dec6c285cd8c0b8992cbb4aa`; re-read canonical guidance at dispatch.
- `MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `low`
- `MODEL_APPROVAL`: Not required by `.harscode-spaces/.local-config.yaml` (`approval_required: false`).
- `PHASE_ROUTE`: Required Planner-owned status reconciliation after explicit Human approval of a material Techplan amendment.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-011/techplan.md` — material amendment approved by Human; change only its frontmatter Status field.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-012/report-techplan.md` — exact report reviewed alongside TP-011; SHA-256 `e5140014cc3f56197e47466143c6f9df69621a227b8f3c8900bb613d6d97cefe`.
- `.harscode-spaces/s2-guest-donation-truthful-state/events.md` — durable explicit Human approval event added 2026-09-30; verify it identifies TP-011 and its matching report.
- `../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md`, `workflow/2-techplan/template.md`, `rules.md`, `guardrails.md`, and `workflow/orchestrated-run-overlay.md` — current canonical Planner entrypoint and Techplan status/report rules.
- Applicable `AGENTS.md` files and current `.harscode-spaces` orchestration sources.

## Task and completion condition

Verify that the durable Human approval event identifies the exact current material Techplan TP-011 and its matching report TP-012. Reconcile only TP-011 frontmatter `Status` from `Draft / In Review` to `Approved` and write this Run's `launch-record.md`. Preserve every other byte of TP-011. Record actual source/report hashes in the launch record when checked.

Do not edit the report, alter substantive Techplan content, resolve or edit Open Items, update Product/spec/API authority, change any other Techplan or orchestration projection, claim `CONTRACT_READY`, or start Build. If approval evidence does not match, stop and report the discrepancy without editing. Stop after this metadata-only phase handoff.

## Human-assisted dispatch

Use a fresh Planner Session in the Kencleng repository root with configured model/effort. Start from canonical `workflow/2-1-techplan-synthesis-prompt.md` and this Invocation. Kickoff: `Catat approval Human dengan menyelaraskan hanya field Status pada Techplan TP-S2-002-011 dari Draft / In Review menjadi Approved. Verifikasi event approval durable di .harscode-spaces/s2-guest-donation-truthful-state/events.md dan report TP-S2-002-012/report-techplan.md yang cocok. Ikuti Invocation TP-S2-002-013 dan current canonical Techplan guidance Harscode. Jangan ubah isi substantif Techplan/report atau memutuskan Open Items; jangan ubah orchestration projection, jangan klaim CONTRACT_READY atau mulai Build. Tulis hanya field Status TP-011 dan launch-record Run ini, lalu berhenti.`
