# Run Invocation — `TP-S2-002-017`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-01 by Orchestration Operator. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `TP-S2-002-017`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-017`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Planner
- `SPECIALIZATION`: Reconcile the durable Human approval of TP-015 in Techplan status metadata
- `PARTICIPANT_ID`: `P-S2-002-TP-017-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-PLANNER`; profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — new Planner metadata-reconciliation occurrence after Human approval and report Run TP-S2-002-016; do not reuse prior Participant/Session context.
- `TARGET_REVISION`: Kencleng checkout `650e73c5d646c29c0ddf1931618f02685d15f7b7` plus the approval event and referenced TP-015/TP-016 artifacts; re-read live sources at dispatch.
- `WORKFLOW_REVISION`: Harscode checkout `33b03a3f62cc3aacba6534b8a011465613c64b09`; use current guidance at dispatch unless a material change affects this assignment.
- `MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `low`
- `MODEL_APPROVAL`: Not required by `.harscode-spaces/.local-config.yaml` (`approval_required: false`).
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Required Planner-owned reconciliation of the approved Techplan's `Status` metadata after verifying the explicit Human decision and its matching review report.
- `MODEL_ROUTING_RATIONALE`: This is a narrow, reversible status/provenance reconciliation in one Techplan artifact. Configured `gpt-6-luna` has sufficient repository-work capability; `low` effort is proportionate. No stronger model or additional decision is needed.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/techplan.md` — exact material Techplan approved by Human; expected frontmatter Status is `Draft / In Review`. Change only this Status field after verification.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-016/report-techplan.md` — exact Human-facing report reviewed alongside TP-015; it identifies TP-015 as its source.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-010/review-findings.md` and `launch-record.md` — clean independent Complex re-review after resolution of RV-009 findings.
- `.harscode-spaces/s2-guest-donation-truthful-state/events.md` — durable explicit Human approval event recorded 2026-10-01. Verify that it identifies the exact TP-015 and matching TP-016 report.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-016/launch-record.md` — report-generation provenance and exact inputs.
- Current Harscode guidance: `../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md`, `workflow/2-techplan/template.md`, `rules.md`, `guardrails.md`, `workflow/orchestrated-run-overlay.md`, and applicable `workflow/AGENTS.md`.
- Current orchestration sources: WU-S2-002 manifest, `work-graph.md`, `control-surface.md`, `outcome.md`, and project tracker.

## Task and completion condition

First verify the approval event is explicit and identifies the exact TP-015 Techplan and its matching TP-016 report. Verify the current source and report hashes match the values recorded in the approval event:

- TP-015 `techplan.md`: `c3bcf0d3430fab389a44f430be964ce0d5d39d15ece0bfa060e36f6b60b7c826`
- TP-016 `report-techplan.md`: `38915941b5658052092f2ff668c516aa356a7e890c1e9252bc2daa791b402fb5`

If either approval identity or hash does not match, stop and report the discrepancy without editing. If verification passes, change only TP-015 frontmatter `Status` from `Draft / In Review` to `Approved`, then write this Run's `launch-record.md`. Record before/after file hashes and verify the source is byte-identical when the Status value is normalized. Preserve every other byte of TP-015.

Do not edit the report, substantive Techplan content, Open Items, Product/spec/API authority, any other Techplan, task snapshots, implementation/tests, or orchestration projections. Do not claim `CONTRACT_READY`, accept Security/PII residual risk, or start Build. Stop after the metadata-only phase handoff; Orchestrator reconciles projections separately, then prepares the post-approval task-snapshot reconciliation Run before Task 02 Build.

## Execution envelope

- `PREAUTHORIZED`: verify the assigned approval evidence and write only TP-015 frontmatter `Status` plus this Run's `launch-record.md`.
- `ORCHESTRATOR_DECISION`: reconcile durable projections and prepare the applicable post-approval task-snapshot reconciliation after this Run completes.
- `HUMAN_REQUIRED`: the approval is recorded in `events.md`; any later Techplan revision, new Product/security/domain decision or residual-risk acceptance, protected Tier-0 authorization, or `CONTRACT_READY`/delivery milestone acceptance.

## Human-assisted dispatch

Use a fresh Planner Session in the Kencleng repository root with model `gpt-6-luna` and reasoning effort `low`. Start from `../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md` and this Invocation; follow `workflow/orchestrated-run-overlay.md` identity/path semantics. Kickoff:

`Catat approval Human dengan menyelaraskan hanya field Status pada Techplan TP-S2-002-015 dari Draft / In Review menjadi Approved. Verifikasi event approval durable di .harscode-spaces/s2-guest-donation-truthful-state/events.md mengidentifikasi TP-015 dan report TP-016 yang cocok, serta hash keduanya cocok dengan Invocation ini. Ikuti current canonical Techplan guidance Harscode dan Invocation TP-S2-002-017. Jika bukti tidak cocok, berhenti tanpa edit. Jika cocok, ubah hanya field Status TP-015 dan tulis launch-record Run ini, verifikasi seluruh byte lain tetap sama. Jangan ubah report/isi Techplan/Open Items/authority/task files/projection; jangan klaim CONTRACT_READY atau mulai Build. Berhenti setelah phase handoff.`

After dispatch, verify the Planner's launch record confirms approval/hash checks and a Status-only change. Then reconcile projections and prepare a fresh post-approval task-snapshot Run only for affected child files, preserving the accepted split/dependency/manifest unless a material topology or dependency change is found.
