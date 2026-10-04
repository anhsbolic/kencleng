# Run Invocation — `RV-S2-007-002`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-04 after Anhar reported BLD-S2-007-001 complete. The Build report, current five-file diff, changed-file hashes, and completed checks were reconciled. This is an independent Code Review; it is prepared but not dispatched.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-007`
- `RUN_ID`: `RV-S2-007-002`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-007/runs/RV-S2-007-002`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-007`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Independent safety, quality, stack-specific, and repository-consistency review of Donation POST 503 reconciliation
- `PARTICIPANT_ID`: `P-S2-007-RV-002-1`
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; `.harscode-spaces/participant-profiles/profiles.md`, SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — independent Reviewer context after Build completion; do not reuse Planner or Implementer Sessions.
- `TARGET_REVISION`: Kencleng HEAD `4e71d3697a479d92b2cad5a79ef5d710f4e8077c` plus current working tree. Recheck every pinned current file before review and stop if hashes or scope have materially drifted.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; current-effective guidance, re-read at dispatch.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned local registry.
- `MODEL_REGISTRY_SOURCE`: `.harscode-spaces/.local-config.yaml`, SHA-256 `ddffdeb4ce8edec181ca36bf32bedf19c513a76e359bc2642985eda6b32f400d`; Luna supports reasoning/repository work at `high` and is not approval-gated.
- `MODEL_ROUTING_RATIONALE`: Review spans the authored OpenAPI contract, generated response surfaces, and consumer classification that governs whether retry intent is retained. A fresh `gpt-6-luna` / `high` context is sufficient for a four-pass independent review of this bounded cross-boundary diff. Sol is approval-gated and not required; no code, runtime, or protected crypto/transaction implementation is under review.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Canonical one-shot four-pass Code Review: Safety → Quality → Stack-Specific Best Practices → Consistency. Do not edit production code. Write review findings, optional patch plan only if changes are required, and this Run's structured phase handoff.
- `ARTIFACT_TARGETS`: `runs/RV-S2-007-002/review-findings-1.md`; optional `runs/RV-S2-007-002/patch-plan-1.md` only if code changes are required; `runs/RV-S2-007-002/handoff.md`. Do not copy the reviewed diff into the Run.

## Current-effective inputs / `PRIOR_ARTIFACTS`

- Approved Techplan spine: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-007/techplan.md`, SHA-256 `ff0cb704151495e9f5539d11c24e808e8618d1c98cdb0d680957df8d51948dc3`.
- Build report: `runs/BLD-S2-007-001/report.md`, SHA-256 `2677f8c39e0723c8a365823edb51ca7c0c35523a38e2db144d4f19732bc4d235`.
- Build Invocation: `runs/BLD-S2-007-001/invocation.md`, SHA-256 `b02be1e2f3022775c11280ec7a9778c5a6f6c4944e1ca8dae50b5879138cbc6e`.
- WU-S2-007 manifest, SHA-256 `c2fa1fdfccb5aa2581eddcde492994b91d02f225f551610a631a45a4e2fcca2c`.
- Target authorities: root `AGENTS.md` SHA-256 `4ec81e122cfeb77b2d13e3b90cc7f4a73e829f97d0627d889793aca847a87716`; `frontend/AGENTS.md` SHA-256 `67355c8ce4384082453b48409039315696849e1831e3c9b143ee53bd91cd86ba`; `api/README.md` SHA-256 `0ec88871d43b32b2b06a82ba345f694b079bfefaef3c2974333b43fad16b434f`.
- Current Donation OI9 authorities: `docs/spec/5-donation/invariants.md` SHA-256 `4bf7ba48f4980f635ce76a4eb3440be5cf12666fdd99e06dcc8ac3b78fbde39a`; `docs/spec/5-donation/features/01-submit-donation-settlement.md` SHA-256 `b3f0f99c420c807bb14cccf75a4adcaf43dfdadccae6b727a1c460ef4ed6a7e9`.
- Canonical `../harscode-workspace/workflow/4-code-review-prompt.md`, `workflow/4-code-review/guidelines.md`, `workflow/4-code-review/checklist.md`, `workflow/orchestrated-run-overlay.md`, and relevant `best-practices/AGENTS.md` routing authority; re-read current versions at dispatch.

## Exact review scope and pinned current files

Review the actual current diff for precisely these five Build-owned production/contract/test files against the pinned pre-Build snapshot below. The Build report is context, not a substitute for inspecting the diff and code.

| Changed path | Pre-Build SHA-256 (from Build Invocation) | Current SHA-256 |
|---|---|---|
| `api/openapi/donation.yaml` | `609411688477132267847db52ec6c7d78b5bfd50df97d7af4e85c049917ad2ca` | `9f7c31065c3c7ffa77491541102afc0ce84085f50a1e2aec317336c9656d5c1d` |
| `api/openapi.yaml` | `a3a67da1294085d15f2bf0f7aef9deb94bd7527e3bede1593570f9323e1e605d` | `c37038ecb5088aebcc9156e495d8b14bfa7cc405debeb940e16d7b17786d1838` |
| `frontend/lib/api/generated/openapi.ts` | `288296d6e65a7500349126e066b3a4215915a647b0c46358f60e954d41262dc4` | `f267598c563b4bf73442f120c81fdf6ec1b7f48cd29095f78f3622dbd7094e20` |
| `frontend/lib/api/donation.ts` | `0bda1df8cb0b2f25c27af258eab2d30bb05c72a1704a826d9e3191c286c3d7ab` | `5c3cdca3b71d806ece0c15bd1dbd8a0a162895ec0c56a4b63c3271f7fac6d654` |
| `frontend/app/donations/donation-flow.test.tsx` | `82bdd24998c9566aa17948aa123341c002b3e3f232185c523ecf74ffb7f2ba86` | `e7b66525258a8601746583274a7ae9d794f0795254aba794e29a46f153fb6a38` |

Read-only comparison anchors that should remain unchanged: `api/openapi/common.yaml` SHA-256 `46bd8fba98d7befe65e87b17f431b0f0a2fa815589db28eb6ded6b00bd9418f8`; `frontend/app/campaigns/[campaignId]/donate/donation-client.tsx` SHA-256 `91b571ae8032a22576b40b2acd8e6f2b3d8d4986bb14fe88e40946c1593ad30b`; `frontend/mocks/handlers/donation.ts` SHA-256 `775baf8a9d0f43215b0d43aaa1e8d6d4babef57704573a3d298e25304e980238`.

Pre-existing unrelated working-tree files include WU-S2-003 orchestration artifacts and Campaign/Donation spec edits predating this Build. They are outside this five-file Code Review scope; do not attribute them to BLD-S2-007-001. If review evidence shows a changed in-scope path no longer matches its pinned hash, or the actual five-file delta differs materially from the Build report/Approved Techplan, stop and report drift before issuing a verdict.

## Review task and verification posture

Run the canonical four passes, in order, against the same exact diff:

1. Safety — check fail-closed 503 contract semantics, no disclosure, response classification ordering, and preservation of ambiguous retries.
2. Quality — maintainability and clarity of the narrow response/client/test change.
3. Stack-Specific Best Practices — use the Harscode best-practices clue map and open only matching API/OpenAPI, TypeScript/fetch, and testing guidance; cite a matching source for any finding. If no clue matches, state that explicitly.
4. Consistency — compare with the exact Kencleng root/frontend instructions, API README, current Donation contract, generated-client workflow, and local test conventions.

The Techplan's contract is authoritative: only the documented Donation POST status 503 is definitive non-admission (`request-failure`); transport and other 5xx outcomes stay ambiguous with same-key/same-payload retry. Confirm the generated bundle/types correspond to the authored split source, generic Problem fields only, no internal detail or `Retry-After`, no shared `Problem` change, no caller UI/copy change, and no misleading runtime claim in the test.

Review is reasoning-first. Do not rerun bundle/type generation or full tests by default. Run a narrowly targeted reproduction only if needed to resolve a concrete review question and record its command/question/result. Do not fix code. For each finding record path/location, problem, significance, suggested resolution, and blocking/non-blocking classification. If code changes are needed, create one narrow patch plan and recommend Build/Patch; otherwise Testing is the normal next phase.

## Execution envelope

- `PREAUTHORIZED`: inspect the exact five-file diff and pinned read-only comparison anchors; read routed review/best-practice/target-repo guidance; run a targeted repro only if needed; write review findings and structured handoff; write patch plan only if required.
- `HUMAN_REQUIRED`: production edits, scope expansion, API-owner acceptance of exact bytes, WU-S2-007 completion, WU-S2-003 dependency release, or any new material Product/API/security decision.
- Do not modify production or test files, run the broad Testing suite, claim independent final Testing, accept exact API bytes, or dispatch another phase.

## Human-assisted dispatch

Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`. Fresh Reviewer / `KC-REVIEWER`, `gpt-6-luna` / `high`.

Kickoff: `Jalankan independent four-pass Code Review Run RV-S2-007-002 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-007/runs/RV-S2-007-002/invocation.md dan canonical ../harscode-workspace/workflow/4-code-review-prompt.md beserta orchestrated-run overlay. Review tepat lima file yang dipin terhadap Approved Techplan hash ff0cb704151495e9f5539d11c24e808e8618d1c98cdb0d680957df8d51948dc3, dengan urutan Safety, Quality, Stack-Specific Best Practices, Consistency. Cocokkan current hashes ke tabel invocation dan baca current diff sendiri; Build report hanya konteks. Gunakan best-practices/AGENTS.md sebagai router, buka hanya guidance yang match. Pertahankan batas bahwa hanya Donation POST 503 merupakan definitive non-admission; transport dan 5xx lain tetap ambiguous same-key/same-payload retry. Jangan edit production/test. Hindari rerun full suite; targeted reproduction hanya jika diperlukan untuk pertanyaan review konkret. Tulis review-findings-1.md dan satu structured handoff; tulis patch-plan-1.md hanya jika ada perubahan kode yang diperlukan. Berhenti setelah Review; jangan otomatis dispatch Testing.`
