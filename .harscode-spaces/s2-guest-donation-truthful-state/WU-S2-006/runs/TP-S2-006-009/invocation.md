# Run Invocation — `TP-S2-006-009`

Status: `NOT DISPATCHED — WITHDRAWN BEFORE LAUNCH` — approval Status was reconciled directly by the Orchestrator on 2026-10-02 under current Harscode deterministic-reconciliation guidance; see parent `events.md` entry “TP006008 approval Status reconciled; source Build prepared”. Do not launch this package.
Prepared: 2026-10-02. Bahasa Indonesia untuk human-facing prose.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-006`
- `RUN_ID`: `TP-S2-006-009`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-009`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Planner
- `SPECIALIZATION`: Propagate explicit Human whole-Techplan approval into Status only
- `PARTICIPANT_ID`: `P-S2-006-TP-009-1`
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
- Exact approval receipt in parent `events.md`, heading “2026-10-02 — Human approved TP006008; Status propagation prepared”.
- Source `WU-S2-006/runs/TP-S2-006-008/techplan.md`, pre-change SHA-256 c1a8806a1c754b849c0b8457e688d9a50aa0d024dc3d4fe2f3c4c5a7d35c3324.
- Matching full report `WU-S2-006/runs/TP-S2-006-008/report-techplan.md`, SHA-256 06d6259f4e50960dbae04951b0b786c15503ec6140f27b9f805149a8a5b6eb6c; TP006008 handoff for mechanical resolution and full report provenance.
- `WU-S2-006/runs/RV-S2-006-005/review-findings.md`; parent completion event verifies mechanical-only TP006008 anchor/provenance correction and continued independent Review applicability.
- WU006 manifest, parent Outcome/Work Graph/Control Surface and Authority Map for current coordination boundaries.

WU paths above relative to parent Slice-2 Space.

## Task / completion

Verify explicit Human approval identifies TP006008 source and matching TP006008 report, then verify both exact hashes above. If evidence/identity/hash differs, stop without editing and report discrepancy.

If verified, change only TP006008 frontmatter Status from Draft / In Review to Approved. This narrowly authorized lifecycle metadata edit is the sole exception to preserving earlier Run artifacts: preserve every other byte, including every other provenance and file-table field. Do not create a successor plan or regenerate report. Record before/after hashes and normalized-Status byte equality in Run-local handoff.md (and provenance record if required by overlay).

No substantive plan/Open Item changes, new owner vote, independent re-review, decomposition invocation, source authoring or Build dispatch. Approval accepts this planning baseline; concrete Product/spec/API acceptance/counterparts, Tier-0 permission, DB application, delivery refresh/runtime and residual-risk gates remain separately applicable. Decomposition Skip recommendation remains; no split accepted.

## Execution envelope

- `PREAUTHORIZED`: read approval/source/report/review/current guidance; write only TP006008 Status and this Run's handoff/provenance.
- `ORCHESTRATOR_DECISION`: verify actual status-only delta, reconcile projections, then prepare next justified owning-source phase.
- `HUMAN_REQUIRED`: genuinely new material revision/decision and applicable concrete source/protected/DB/runtime/risk gates.

Do not edit report, other prior artifacts, Product/spec/API/generated/fixture/production/test files, registry or orchestration projections. No tests/validators/generators/services/migrations/browser/runtime checks.

## Human-assisted dispatch package — withdrawn; do not launch

Cwd `/home/anhar-solehudin/kencleng-workspace/kencleng`; fresh Planner / KC-PLANNER, gpt-6-luna / low.

Kickoff: `Jalankan Planner Run TP-S2-006-009 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-009/invocation.md dan canonical ../harscode-workspace/workflow/2-1-techplan-synthesis-prompt.md dengan orchestrated-run overlay. Berhenti setelah phase handoff.`

Report completion atau exact scoped discrepancy kepada Orchestrator; jangan otomatis lanjut Build.
