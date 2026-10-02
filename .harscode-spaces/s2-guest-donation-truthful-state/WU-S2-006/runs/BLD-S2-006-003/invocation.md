# Run Invocation — `BLD-S2-006-003`

Status: `READY_FOR_HUMAN_DISPATCH`
Prepared: 2026-10-02. Human-facing prose Bahasa Indonesia.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-006`
- `RUN_ID`: `BLD-S2-006-003`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/BLD-S2-006-003`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Implementer
- `SPECIALIZATION`: Apply independent Review patch plan to Campaign/Donation specifications
- `PARTICIPANT_ID`: `P-S2-006-BLD-003-1`
- `PARTICIPANT_PROFILE_ID`: `KC-IMPLEMENTER`; `.harscode-spaces/participant-profiles/profiles.md` SHA-256 e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32.
- `SESSION_TRANSITION`: `FRESH` — new Build/Patch occurrence after terminal Review; no prior session reuse.
- `TARGET_REVISION`: Kencleng HEAD 7fd8b473b239b20bda3990ab29c51440d321a796 plus current working tree; preserve all earlier/unrelated changes.
- `WORKFLOW_REVISION`: Harscode pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8; current applicable guidance at dispatch.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `medium`
- `MODEL_APPROVAL`: Not required by Human-owned registry.
- `MODEL_ROUTING_RATIONALE`: Lowest-cost repository model; medium sufficient for bounded explicit four-file source correction to existing approved rules and review patch plan. No new policy/architecture or stronger-model escalation justified.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Canonical Build/Patch re-entry requested by Code Review RV006003.

## Current-effective inputs / PRIOR_ARTIFACTS

- Canonical `../harscode-workspace/workflow/3-build-prompt.md`, applicable Build guidelines/checklist and `workflow/orchestrated-run-overlay.md`; root AGENTS, KC-IMPLEMENTER and routed source authorities.
- Approved spine `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-004/techplan.md`, SHA-256 b6c9d10efd1d1b0b06cb197a35c25728dea45fd60c49a9f03ca5c52e4389ae78; no decomposed task file.
- Specific requesting patch plan `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/RV-S2-006-003/patch-plan.md`, SHA-256 3cd6a4743b3866cc56f736cf7178b671e74b8cc9eac22d13a8e88b429eb23767; findings at same Run `review-findings.md`, SHA-256 bf4f5aed6f9b8506221a537487dd1f7a2d2807b0e74ad6667800b8a1bd69d74e. Verdict Request changes, blocking C-01 and non-blocking Q-01.
- Predecessor BLD006002/report.md, its Run-entry baseline snapshots and source-delta.patch for exact existing amendment; parent current state/event records remain coordination authority.
- Accepted Product `docs/product/mvp-scope.md` and `docs/product/mvp-delivery-slices.md`; exact acceptance receipt in parent events. Product and approved plan remain unchanged; no new Human decision needed to implement these findings.

## Assigned patch scope / current baseline

- `docs/spec/4-campaign/invariants.md`: `d4bd69b41484fc7b921d9f234d787cc01d095a61e4ccea33d4ada31ac0a57888`
- `docs/spec/4-campaign/features/09-closure.md`: `19afdf7e5b71f7ffabda5880a8d683f486a39ff0ef63302f96eaa70820795a62`
- `docs/spec/5-donation/invariants.md`: `30743644c0d5b9ebe620212442599afc5f9f5bb0323fc2acef97186f7c9902b4`
- `docs/spec/5-donation/features/01-submit-donation-settlement.md`: `5050dd29cafc0bd4a906863b0fe32c2ac333117eb97701d6991cbc7a29d28f21`

All six original spec checkpoint hashes were verified unchanged after Review. Only the four sources above are patch-authorized by the specific patch plan. Preserve the two other reviewed source files and accepted Product/API/production/prior evidence.

## Task / completion

Execute RV006003's specific patch plan C-01 and Q-01 against approved parent plan and live relevant sources. Do not re-explore settled behavior or invent new policy/mechanism/transport detail. This is correction of the same pending spec-source checkpoint, not acceptance of it or completion of the whole reconciliation target.

Capture exact Run-entry snapshots for changed sources before writing. Produce Run-only source-delta.patch and before/after hashes so subsequent independent Review can distinguish this correction from earlier dirty amendments. Write Run-local `patch-report-1.md` with canonical Build/Patch report and phase handoff, completed findings versus any remaining issue, actual source checks, and remaining source/API/counterpart/runtime gates. Do not overwrite predecessor report/patch evidence.

Return to requesting Code Review after the patch. Orchestrator determines targeted confirmation versus broader re-review from actual delta/current guidance; no automatic waiver or full new review loop claimed here. Source-owner acceptance remains after applicable independent confirmation; do not proceed to API or Testing on unresolved spec mismatch.

## Execution envelope

- `PREAUTHORIZED`: read approved plan/requesting patch plan/current authority; apply only assigned four-source patch and write Run-local snapshots/diff/report/provenance.
- `ORCHESTRATOR_DECISION`: inspect actual delta and route back to independent Review, then owning spec acceptance.
- `HUMAN_REQUIRED`: genuinely new material decision or concrete source acceptance and later protected/DB/runtime/risk gates.

No Product/API/generated/fixture/production/migration/test changes, registry/projection/earlier Run edits. No tests unless explicitly requested; focused source traceability/diff checks only. No validators/generators/services/browser/runtime/race/load/security suites, commits, auto-dispatch or milestones.

## Human-assisted dispatch

Cwd `/home/anhar-solehudin/kencleng-workspace/kencleng`; fresh Implementer / KC-IMPLEMENTER, gpt-6-luna / medium.

Kickoff: `Jalankan Build/Patch Run BLD-S2-006-003 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/BLD-S2-006-003/invocation.md dan canonical ../harscode-workspace/workflow/3-build-prompt.md dengan orchestrated-run overlay. Berhenti setelah patch report dan phase handoff.`

Laporkan completion atau exact scoped blocker ke Orchestrator. Jangan otomatis lanjut Review/acceptance/API Build.
