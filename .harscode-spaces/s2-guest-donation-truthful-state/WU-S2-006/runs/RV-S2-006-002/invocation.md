# Run Invocation — `RV-S2-006-002`

Status: `READY_FOR_HUMAN_DISPATCH`
Prepared: 2026-10-02. Bahasa Indonesia untuk human-facing prose.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-006`
- `RUN_ID`: `RV-S2-006-002`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/RV-S2-006-002`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Independent review of Product/MVP source amendment checkpoint
- `PARTICIPANT_ID`: `P-S2-006-RV-002-1`
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; `.harscode-spaces/participant-profiles/profiles.md`, SHA-256 e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32.
- `SESSION_TRANSITION`: `FRESH` — independent Reviewer after completed Build occurrence.
- `TARGET_REVISION`: Kencleng HEAD 7fd8b473b239b20bda3990ab29c51440d321a796 plus exact current two-file diff below; unrelated working-tree changes excluded.
- `WORKFLOW_REVISION`: Harscode pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8; guidance current-effective at dispatch.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `medium`
- `MODEL_APPROVAL`: Not required by Human-owned registry.
- `MODEL_ROUTING_RATIONALE`: Lowest sufficient available model; medium effort for cohesive two-document Product-policy fidelity review without implementation/runtime complexity. No demonstrated reason for stronger-model escalation.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Canonical independent four-pass Code Review of material source amendments.

## Current-effective inputs / PRIOR_ARTIFACTS

- Canonical `../harscode-workspace/workflow/4-code-review-prompt.md`, triggered review guidelines/checklist and `workflow/orchestrated-run-overlay.md`; root AGENTS and KC-REVIEWER profile/applicable Product authority routing.
- Approved spine `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-004/techplan.md`, SHA-256 b6c9d10efd1d1b0b06cb197a35c25728dea45fd60c49a9f03ca5c52e4389ae78. No decomposition/current task file.
- Build checkpoint `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/BLD-S2-006-001/report.md`, SHA-256 dae5931bef13dd7913b5b645126ef9e2403866efd3e2d8f2704dec119c03e46e; orientation only, actual diff owns Review evidence.
- Explicit diff: `git diff HEAD -- docs/product/mvp-scope.md docs/product/mvp-delivery-slices.md`. Both files clean at Build entry, baseline hashes independently checked against HEAD. No other source changes in this checkpoint.
- `docs/product/mvp-scope.md`: baseline e015a828f7b6997030dbfba894a70bb1dff23c92136de7463b983f21ee7a2813; current ba2972bc8f91d092e477df170d987b1d124964d9cc36c025d2a8da3ed12709af.
- `docs/product/mvp-delivery-slices.md`: baseline 72f36b5d2a523a4bd3fba7454b6646cef08d9275ecdf1154555ff3c293307654; current 4c69a030e7fedc9c62bf30f85c00e81f9806c45b2ed5471c1d5126762be8091f.
- WU006 manifest, parent events/Work Graph/Control Surface/Authority Map for checkpoint and authority boundaries. Do not default-load raw Exploration or historical Runs.

## Task / completion

Review the exact current two-file Product/MVP amendments through all four canonical passes against the Approved spine and applicable target authority. Judge this explicit checkpoint, not completion of future Campaign/Donation spec/API/counterpart work. Verify source/diff identity; if material unexpected drift changes assigned scope, surface it before adopting another target.

Write Run-local `review-findings.md` including canonical verdict and phase handoff. Write `patch-plan.md` only if changes are required. Record findings with location, evidence, impact, resolution and blocking status; do not edit sources or invent findings. Report applicable next gate based on actual verdict: corrections return to fresh Build/Patch; a satisfactory Product checkpoint is ready for Human concrete source acceptance. Technical Review verdict cannot grant Product/MVP acceptance or WU006 completion.

## Execution envelope

- `PREAUTHORIZED`: read exact diff/plan/current relevant authority; write Run-local independent findings/patch plan/provenance.
- `ORCHESTRATOR_DECISION`: route findings or owning Product acceptance, then remaining approved source work.
- `HUMAN_REQUIRED`: concrete Product/MVP acceptance and all later applicable source/protected/runtime/risk gates.

No source/production/Participant history/projection/registry changes. No tests unless explicitly requested; no validators/generators/browser/runtime/services/migrations/race/load/security suites. Review primarily reasoning. Planned runtime/compatibility evidence remains downstream; do not claim it ran.

## Human-assisted dispatch

Cwd `/home/anhar-solehudin/kencleng-workspace/kencleng`; fresh Reviewer / KC-REVIEWER, gpt-6-luna / medium.

Kickoff: `Jalankan Review Run RV-S2-006-002 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/RV-S2-006-002/invocation.md dan canonical ../harscode-workspace/workflow/4-code-review-prompt.md dengan orchestrated-run overlay. Berhenti setelah review-findings dan phase handoff.`

Laporkan verdict/completion atau scoped blocker ke Orchestrator; jangan otomatis lanjut patch/acceptance/source Build.
