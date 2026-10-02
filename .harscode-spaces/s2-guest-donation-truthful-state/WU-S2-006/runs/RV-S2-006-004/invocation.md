# Run Invocation — `RV-S2-006-004`

Status: `READY_FOR_HUMAN_DISPATCH`
Prepared: 2026-10-02. Human-facing prose Bahasa Indonesia.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-006`
- `RUN_ID`: `RV-S2-006-004`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/RV-S2-006-004`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Independent targeted confirmation after RV006003 source patch
- `PARTICIPANT_ID`: `P-S2-006-RV-004-1`
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; `.harscode-spaces/participant-profiles/profiles.md` SHA-256 e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32.
- `SESSION_TRANSITION`: `FRESH` — independent new Review occurrence following completed Build/Patch.
- `TARGET_REVISION`: Kencleng HEAD 7fd8b473b239b20bda3990ab29c51440d321a796 plus exact four-source patch below; unrelated/prior working-tree changes excluded.
- `WORKFLOW_REVISION`: Harscode pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8; applicable guidance current-effective at dispatch.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `medium`
- `MODEL_APPROVAL`: Not required by Human-owned registry.
- `MODEL_ROUTING_RATIONALE`: Lowest-cost repository model; medium sufficient for bounded independent confirmation of explicit reviewed corrections and adjacent consistency, without new policy/runtime complexity.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Return to requesting Code Review; narrow-patch confirmation under canonical applicability guidance.

## Current-effective inputs / PRIOR_ARTIFACTS

- Canonical `../harscode-workspace/workflow/4-code-review-prompt.md`, applicable guidelines/checklist and `workflow/orchestrated-run-overlay.md`; root AGENTS and KC-REVIEWER profile/relevant authority routing.
- Approved spine `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-004/techplan.md`, SHA-256 b6c9d10efd1d1b0b06cb197a35c25728dea45fd60c49a9f03ca5c52e4389ae78. No decomposed task file.
- Requesting Review `WU-S2-006/runs/RV-S2-006-003/review-findings.md` and `patch-plan.md`: C-01 blocking, Q-01 non-blocking; original independent four-pass evidence preserved.
- Build/Patch `WU-S2-006/runs/BLD-S2-006-003/patch-report-1.md`, SHA-256 c2ee2c37777d0d7ae55dc97c74457c8fab73ddb7e3e1f1b89210871a15c43e34.
- Exact patch `WU-S2-006/runs/BLD-S2-006-003/source-delta.patch`, SHA-256 f089b009eb5acc176279c3e5226ac91dd4005a512580892c32259980f8ae9c5c; baseline mapping below. Actual baseline-to-live diff is Review evidence; report does not substitute for source inspection.
- Accepted Product sources and exact owner receipt named by Approved plan/current parent events. Current WU006/Work Graph/Authority Map for remaining gates. Do not default-load raw Exploration history.

WU-S2 paths relative parent Slice-2 Space.

## Explicit patch scope

Baseline names below relative BLD-S2-006-003 Run. Each snapshot matches the corresponding source reviewed in RV006003; two original checkpoint sources (Campaign features01/02) remained byte-identical to RV006003 snapshots.

| Source | Exact Run-entry snapshot | Current source SHA-256 |
|---|---|---|
| `docs/spec/4-campaign/invariants.md` | `baseline/campaign-invariants.md` | `0d3f250e4d51d98a63c865c6fe0810b1b4fdb5ab063b706c7ab32f7e7bd73342` |
| `docs/spec/4-campaign/features/09-closure.md` | `baseline/campaign-09-closure.md` | `dd0a3c6a897091e887a4e64127322d8e4d28036bd15f6cedc70777f7d329df4d` |
| `docs/spec/5-donation/invariants.md` | `baseline/donation-invariants.md` | `68c7967fba44a3012ee67730bc5b6a2011961859b1b119dbe0709d96d0d953f8` |
| `docs/spec/5-donation/features/01-submit-donation-settlement.md` | `baseline/donation-01-submit-donation-settlement.md` | `3b1f3918752724c4c6448aa001f287ee59461a5fb8f75b82a0495db072755dfe` |

Orchestrator reconstructed baseline-to-current four-source unified diff byte-for-byte and verified unchanged two sibling sources. Actual patch is confined to requesting findings and associated verification text; it restores existing approved Product/Techplan behavior without a new policy/architecture/transport or production change. Canonical guidance allows targeted confirmation for a narrow exact patch rather than an automatic full review replay. This is not Human acceptance or an invented skipped Review.

## Task / completion

Independently confirm C-01/Q-01 against requesting patch plan, approved spine, exact actual delta and surrounding affected sources. Establish whether findings are resolved and whether the correction introduces any material inconsistency. Carry forward original full-review evidence only where still applicable. Do not pre-decide resolution or verdict. If actual delta broadens contract/scope/security/architecture beyond that patch, report applicability/drift and expand/re-route review as justified by canonical guidance.

Write Run-local `review-findings.md` with explicit confirmation outcome for each finding, scope/applicability rationale, any new findings, verification actually executed, verdict and phase handoff. Write `patch-plan.md` only if corrections are required. A satisfactory confirmation routes concrete six-file spec checkpoint to owning Human acceptance; technical Review cannot accept source bytes, complete WU006, or release API/production/runtime gates itself.

## Execution envelope

- `PREAUTHORIZED`: read exact patch/snapshots/live source/approved plan/relevant authorities; write Run-local independent findings/patch plan/provenance.
- `ORCHESTRATOR_DECISION`: reconcile confirmation scope and actual verdict, then findings or owning spec acceptance.
- `HUMAN_REQUIRED`: concrete spec/API source acceptance and genuinely new material choice/protected/DB/runtime/risk gates.

No source/production/test/prior artifact/projection/registry edits. No tests unless explicitly requested; no validators/generators/services/browser/runtime/migration/race/load/security suites. Source/diff reasoning only here. No commits, automatic patch/dispatch, milestones or source acceptance.

## Human-assisted dispatch

Cwd `/home/anhar-solehudin/kencleng-workspace/kencleng`; fresh Reviewer / KC-REVIEWER, gpt-6-luna / medium.

Kickoff: `Jalankan Review Run RV-S2-006-004 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/RV-S2-006-004/invocation.md dan canonical ../harscode-workspace/workflow/4-code-review-prompt.md dengan orchestrated-run overlay. Berhenti setelah review-findings dan phase handoff.`

Laporkan outcome/verdict atau scoped blocker ke Orchestrator; jangan otomatis lanjut acceptance/API Build.
