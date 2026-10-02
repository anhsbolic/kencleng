# Run Invocation — `TP-S2-006-006`

Status: `READY_FOR_HUMAN_DISPATCH`
Prepared: 2026-10-02. Bahasa Indonesia untuk human-facing prose.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-006`
- `RUN_ID`: `TP-S2-006-006`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-006`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Planner
- `SPECIALIZATION`: Propagate explicit Human whole-Techplan approval into Status only
- `PARTICIPANT_ID`: `P-S2-006-TP-006-1`
- `PARTICIPANT_PROFILE_ID`: `KC-PLANNER`; `.harscode-spaces/participant-profiles/profiles.md`, SHA-256 e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32.
- `SESSION_TRANSITION`: `FRESH` — new occurrence after terminal report handoff and Human approval.
- `TARGET_REVISION`: Kencleng HEAD 7fd8b473b239b20bda3990ab29c51440d321a796 plus current working tree.
- `WORKFLOW_REVISION`: Harscode pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8; use current applicable guidance at dispatch.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `low`
- `MODEL_APPROVAL`: Not required by Human-owned registry.
- `MODEL_ROUTING_RATIONALE`: Lowest sufficient available model/effort for bounded hash verification and one-field metadata propagation; no material synthesis.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Planner-owned approval Status propagation, canonical 2-1 Techplan guidance with orchestrated-run overlay.

## Current-effective inputs

- Canonical `../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md`; applicable Techplan template/rules/guardrails, `workflow/orchestrated-run-overlay.md` and scoped instructions.
- Exact approval receipt in parent `events.md`, heading “2026-10-02 — Human approved TP006004; Status propagation prepared”.
- Source `WU-S2-006/runs/TP-S2-006-004/techplan.md`, pre-change SHA-256 b925c527c101c6695a64a2c3a116ac206cd3f9a6162f1d5c0d82c2fc80c3e9fe.
- Matching full report `WU-S2-006/runs/TP-S2-006-005/report-techplan.md`, SHA-256 c81e4e8d5dfe9ed5e47bce891f27c5dca11763cfacde2d66fd893dcd2cc48347; TP006005 handoff for unchanged source and historical provenance clarification.
- `WU-S2-006/runs/RV-S2-006-001/review-findings.md`; parent completion event verifies mechanical-only TP006004 correction and continued independent Review applicability.
- WU006 manifest, parent Outcome/Work Graph/Control Surface and Authority Map for current coordination boundaries.

WU paths above relative to parent Slice-2 Space.

## Task / completion

Verify explicit Human approval identifies TP006004 source and matching TP006005 report, then verify both exact hashes above. If evidence/identity/hash differs, stop without editing and report discrepancy.

If verified, change only TP006004 frontmatter Status from Draft / In Review to Approved. This narrowly authorized lifecycle metadata edit is the sole exception to preserving earlier Run artifacts: preserve every other byte, including historical effort/file-table metadata. Do not create a successor plan or regenerate report. Record before/after hashes and normalized-Status byte equality in Run-local handoff.md (and provenance record if required by overlay).

No substantive plan/Open Item changes, new owner vote, independent re-review, decomposition invocation, source authoring or Build dispatch. Approval accepts this planning baseline; concrete Product/spec/API acceptance/counterparts, Tier-0 permission, DB application, delivery refresh/runtime and residual-risk gates remain separately applicable. Decomposition Skip recommendation remains; no split accepted.

## Execution envelope

- `PREAUTHORIZED`: read approval/source/report/review/current guidance; write only TP006004 Status and this Run's handoff/provenance.
- `ORCHESTRATOR_DECISION`: verify actual status-only delta, reconcile projections, then prepare next justified owning-source phase.
- `HUMAN_REQUIRED`: genuinely new material revision/decision and applicable concrete source/protected/DB/runtime/risk gates.

Do not edit report, other prior artifacts, Product/spec/API/generated/fixture/production/test files, registry or orchestration projections. No tests/validators/generators/services/migrations/browser/runtime checks.

## Human-assisted dispatch

Cwd `/home/anhar-solehudin/kencleng-workspace/kencleng`; fresh Planner / KC-PLANNER, gpt-6-luna / low.

Kickoff: `Jalankan Planner Run TP-S2-006-006 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-006/invocation.md dan canonical ../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md dengan orchestrated-run overlay. Berhenti setelah phase handoff.`

Report completion atau exact scoped discrepancy kepada Orchestrator; jangan otomatis lanjut Build.
