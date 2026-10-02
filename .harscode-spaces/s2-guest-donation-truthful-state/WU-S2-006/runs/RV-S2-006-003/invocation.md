# Run Invocation — `RV-S2-006-003`

Status: `READY_FOR_HUMAN_DISPATCH`
Prepared: 2026-10-02. Bahasa Indonesia untuk human-facing prose.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-006`
- `RUN_ID`: `RV-S2-006-003`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/RV-S2-006-003`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Independent review of Campaign/Donation specification amendments
- `PARTICIPANT_ID`: `P-S2-006-RV-003-1`
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; `.harscode-spaces/participant-profiles/profiles.md`, SHA-256 e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32.
- `SESSION_TRANSITION`: `FRESH` — independent Reviewer after completed BLD006002 occurrence.
- `TARGET_REVISION`: Kencleng HEAD 7fd8b473b239b20bda3990ab29c51440d321a796 plus exact six-file Run-only delta below; earlier/unrelated working-tree changes excluded.
- `WORKFLOW_REVISION`: Harscode pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8; ordinary guidance current-effective at dispatch.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned registry.
- `MODEL_ROUTING_RATIONALE`: Lowest-cost available repository model; high effort proportionate cross-domain invariant, lifecycle, financial and public-contract fidelity across six spec files. No demonstrated reason for stronger-model escalation.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Canonical independent four-pass Code Review of material spec-source checkpoint.

## Current-effective inputs / PRIOR_ARTIFACTS

- Canonical `../harscode-workspace/workflow/4-code-review-prompt.md`, triggered review guidelines/checklist and `workflow/orchestrated-run-overlay.md`; applicable root AGENTS, KC-REVIEWER profile and routed Product/spec authorities.
- Approved spine `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-004/techplan.md`, SHA-256 b6c9d10efd1d1b0b06cb197a35c25728dea45fd60c49a9f03ca5c52e4389ae78. No decomposed current task file.
- Accepted Product inputs `docs/product/mvp-scope.md` SHA-256 ba2972bc8f91d092e477df170d987b1d124964d9cc36c025d2a8da3ed12709af and `docs/product/mvp-delivery-slices.md` SHA-256 4c69a030e7fedc9c62bf30f85c00e81f9806c45b2ed5471c1d5126762be8091f. Concrete Product acceptance receipt in parent events; acceptance does not extend to these spec amendments.
- Build `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/BLD-S2-006-002/report.md`, SHA-256 e0a7688cf36d1f4d0d9fefbcb2f211632814c536be3576ea614b0ea202c71d8b; orientation, not substitute for actual delta/source.
- Exact Run-only diff `BLD-S2-006-002/source-delta.patch`, SHA-256 d0bc1537eb5336a66d23128bc63fd8259bd5607dd863ed2eae76871719439560. Baseline snapshots under `BLD-S2-006-002/baseline/<repo-relative source path>`; compare each snapshot to live current source. Paths BLD-S2-* relative WU006/runs.
- Parent Work Graph/WU006 manifest/events/Authority Map for current gates; shared monetary standard through root routing when relevant. Do not default-load raw Exploration history.

## Explicit diff / changed scope

| Source | Run-entry snapshot SHA-256 | Current source SHA-256 |
|---|---|---|
| `docs/spec/4-campaign/invariants.md` | `bec2ef663229249f3d032ede08b6833e08091e774d7b5b0261ee9f40deb2b9b6` | `d4bd69b41484fc7b921d9f234d787cc01d095a61e4ccea33d4ada31ac0a57888` |
| `docs/spec/4-campaign/features/01-campaign-creation-draft-crud.md` | `0ecf1a1e82212a10e4202fc588f3ca88eabb496bc08ff306f87813764f4c8535` | `53431422b714adebd6ff9f07476c4c930746516f89f98d4c7ea122b7f35c8d05` |
| `docs/spec/4-campaign/features/02-campaign-detail-listing.md` | `55d0ede37e5fe70e07f5e2f2832f3bd03ce94a929e3995538fece7e681f62349` | `2db5dd950553e97756417fd2f41e29356835c43a795c80c3be1d615fc3cd45a7` |
| `docs/spec/4-campaign/features/09-closure.md` | `40df308338f3885a5e4f8531c7743e71af47b967ef2077da5ad53c03cb2272c6` | `19afdf7e5b71f7ffabda5880a8d683f486a39ff0ef63302f96eaa70820795a62` |
| `docs/spec/5-donation/invariants.md` | `436a1421ec12bf0a0618066b16468226221cf9cfddf2bea6cb8d69b26cbb97e8` | `30743644c0d5b9ebe620212442599afc5f9f5bb0323fc2acef97186f7c9902b4` |
| `docs/spec/5-donation/features/01-submit-donation-settlement.md` | `36cf19f24491284346b6a268213dca2204330a4f9d5f348f18b466b20faa842b` | `5050dd29cafc0bd4a906863b0fe32c2ac333117eb97701d6991cbc7a29d28f21` |

Orchestrator reconstructed unified diff from all six snapshots and live sources: exact patch match; before/after hashes matched report. Verify current scope at Review entry. Full HEAD diff contains prior accepted dirty amendments and is not this Run's scope. Read enough surrounding current source to judge consistency; do not review unrelated earlier changes as newly authored work.

## Task / completion

Run all four canonical passes independently against the exact current six-file amendment and the Approved plan/current accepted higher-level sources. Judge this explicit specification checkpoint; full target/API/generated/consumer/runtime stages remain pending. Do not pre-decide verdict, silently fix sources, or redefine the execution spine. If intent/authority is materially missing or conflicting, report the exact issue and owning route under canonical guidance.

Write Run-local `review-findings.md` with canonical verdict and phase handoff. Write `patch-plan.md` only when changes are required. Findings need location/evidence/impact/resolution and blocking status. Technical Review cannot grant Campaign/Donation owner acceptance. Corrections route fresh Build/Patch; a satisfactory checkpoint is ready for applicable owning Human spec-source acceptance before dependent API authoring. Record remaining verification ownership without claiming runtime or counterpart proof.

## Execution envelope

- `PREAUTHORIZED`: read exact delta/snapshots/current sources/approved plan/relevant authority; write Run-local findings/patch plan/provenance only.
- `ORCHESTRATOR_DECISION`: reconcile verdict and route findings or owning spec acceptance, then remaining approved source stages.
- `HUMAN_REQUIRED`: concrete spec/API acceptance, genuine new material decisions, protected/DB/runtime/risk gates.

No source/production/test/prior Run/projection/registry edits. No tests unless explicitly requested; no validators/generators/services/browser/runtime/migrations/race/load/security suites. Review reasoning and source/diff checks only here. No commits, automatic patch/dispatch or milestones.

## Human-assisted dispatch

Cwd `/home/anhar-solehudin/kencleng-workspace/kencleng`; fresh Reviewer / KC-REVIEWER, gpt-6-luna / high.

Kickoff: `Jalankan Review Run RV-S2-006-003 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/RV-S2-006-003/invocation.md dan canonical ../harscode-workspace/workflow/4-code-review-prompt.md dengan orchestrated-run overlay. Berhenti setelah review-findings dan phase handoff.`

Laporkan verdict/completion atau exact scoped blocker kepada Orchestrator; jangan otomatis lanjut patch/acceptance/API Build.
