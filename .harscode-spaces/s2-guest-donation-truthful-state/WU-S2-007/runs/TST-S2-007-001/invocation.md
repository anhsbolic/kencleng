# Run Invocation — `TST-S2-007-001`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-04 after Anhar reported `RV-S2-007-002` complete. Review verdict is `Approve with minor comments`; its sole comment BP-1 is non-blocking and no patch plan is required. This is a fresh independent Testing Run; it is prepared but not dispatched.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-007`
- `RUN_ID`: `TST-S2-007-001`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-007/runs/TST-S2-007-001`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-007`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Tester
- `SPECIALIZATION`: Independent API/client contract and observable consumer verification for Donation POST Funding-unavailable reconciliation
- `PARTICIPANT_ID`: `P-S2-007-TST-001-1`
- `PARTICIPANT_PROFILE_ID`: `KC-TESTER`; `.harscode-spaces/participant-profiles/profiles.md`, SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — independent Testing context after Build and Code Review; do not reuse Planner, Implementer, or Reviewer Sessions.
- `TARGET_REVISION`: Kencleng HEAD `4e71d3697a479d92b2cad5a79ef5d710f4e8077c` plus the current working tree. Recheck pinned changed paths before testing and stop/report if scope or hashes materially drift. Unrelated pre-existing WU-S2-003/spec working-tree changes are outside this Run.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; current-effective guidance must be re-read at dispatch.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `medium`
- `MODEL_APPROVAL`: Not required by Human-owned local registry.
- `MODEL_REGISTRY_SOURCE`: `.harscode-spaces/.local-config.yaml`, SHA-256 `ddffdeb4ce8edec181ca36bf32bedf19c513a76e359bc2642985eda6b32f400d`; Luna supports reasoning/repository work at `medium` and is not approval-gated.
- `MODEL_ROUTING_RATIONALE`: The approved scope is a bounded contract/client behavior change with explicit observable assertions and repository-owned final commands. A fresh `gpt-6-luna` / `medium` context is sufficient for independent API validation, generated-surface comparison, and frontend final verification. No backend runtime, migration, concurrency, or protected implementation is in scope; Sol is not required.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Canonical independent Testing: sweep existing evidence, cover each approved Techplan rule and applicable Test Focus Pointer, execute Testing-owned verification and target-repository final checks, fresh-read the whole Approved Techplan, then write one testing report and structured phase handoff. Do not edit production/test implementation.
- `ARTIFACT_TARGETS`: `runs/TST-S2-007-001/testing-report-1.md`; `runs/TST-S2-007-001/patch-plan-1.md` only if a genuine code change is required. Do not copy the Techplan or reviewed diff into this Run.

## Current-effective inputs / `PRIOR_ARTIFACTS`

- Canonical `../harscode-workspace/workflow/5-testing-prompt.md`, SHA-256 `b2885ef6208a25f7541aca0f3a6f034a666aff484733abc852f835d10a03ef8e`; `workflow/5-testing/guidelines.md`, SHA-256 `6c71cc0aa074dd0107379720254ee56860ed270d97f2f606e4d956d7aaddad02`; `workflow/5-testing/checklist.md`, SHA-256 `42126b98abf68817347ba0b5d49ffe5c06892d0741779e179ec0e789d01f4bd6`; and `workflow/orchestrated-run-overlay.md`, SHA-256 `f806673aa5e7d1d82dae9f433ec54c3562e15ce778b709a141c57d3d22dfaf87`. Re-read current versions at dispatch; these hashes are preparation provenance, not a request to ignore current-effective guidance.
- Approved Techplan: `techplan.md`, SHA-256 `ff0cb704151495e9f5539d11c24e808e8618d1c98cdb0d680957df8d51948dc3`.
- Latest Build report: `runs/BLD-S2-007-001/report.md`, SHA-256 `2677f8c39e0723c8a365823edb51ca7c0c35523a38e2db144d4f19732bc4d235`; Build Invocation: `runs/BLD-S2-007-001/invocation.md`, SHA-256 `b02be1e2f3022775c11280ec7a9778c5a6f6c4944e1ca8dae50b5879138cbc6e`.
- Latest independent Review: `runs/RV-S2-007-002/review-findings-1.md`, SHA-256 `23931fe93aa41ca8c8bec067173860f03da7c5f883d4cb67ba9985f29e6a4433`; `runs/RV-S2-007-002/handoff.md`, SHA-256 `8311f0f00bc9ca25b14ab976d568eda865faddeec2e2bc21648cb614069f4232`. Verdict `Approve with minor comments`; BP-1 is a non-blocking suggestion to type the generic Problem fixture against the generated schema. Testing must inspect the corresponding observable assertion and assess whether the untyped fixture creates an actual coverage problem; no patch is pre-authorized or assumed necessary.
- Target authorities: root `AGENTS.md`, SHA-256 `4ec81e122cfeb77b2d13e3b90cc7f4a73e829f97d0627d889793aca847a87716`; `frontend/AGENTS.md`, SHA-256 `67355c8ce4384082453b48409039315696849e1831e3c9b143ee53bd91cd86ba`; `api/README.md`, SHA-256 `0ec88871d43b32b2b06a82ba345f694b079bfefaef3c2974333b43fad16b434f`.
- Current Donation policy authorities: `docs/spec/5-donation/invariants.md`, SHA-256 `4bf7ba48f4980f635ce76a4eb3440be5cf12666fdd99e06dcc8ac3b78fbde39a`; `docs/spec/5-donation/features/01-submit-donation-settlement.md`, SHA-256 `b3f0f99c420c807bb14cccf75a4adcaf43dfdadccae6b727a1c460ef4ed6a7e9`.

## Exact changed paths and expected current hashes

Recheck these five Build-owned paths against this table before verification. They must remain the approved Build output; stop and report drift if any changed materially after Review.

| Changed path | Expected SHA-256 |
|---|---|
| `api/openapi/donation.yaml` | `9f7c31065c3c7ffa77491541102afc0ce84085f50a1e2aec317336c9656d5c1d` |
| `api/openapi.yaml` | `c37038ecb5088aebcc9156e495d8b14bfa7cc405debeb940e16d7b17786d1838` |
| `frontend/lib/api/generated/openapi.ts` | `f267598c563b4bf73442f120c81fdf6ec1b7f48cd29095f78f3622dbd7094e20` |
| `frontend/lib/api/donation.ts` | `5c3cdca3b71d806ece0c15bd1dbd8a0a162895ec0c56a4b63c3271f7fac6d654` |
| `frontend/app/donations/donation-flow.test.tsx` | `e7b66525258a8601746583274a7ae9d794f0795254aba794e29a46f153fb6a38` |

Read-only comparison anchors expected unchanged: `api/openapi/common.yaml` SHA-256 `46bd8fba98d7befe65e87b17f431b0f0a2fa815589db28eb6ded6b00bd9418f8`; `frontend/app/campaigns/[campaignId]/donate/donation-client.tsx` SHA-256 `91b571ae8032a22576b40b2acd8e6f2b3d8d4986bb14fe88e40946c1593ad30b`; `frontend/mocks/handlers/donation.ts` SHA-256 `775baf8a9d0f43215b0d43aaa1e8d6d4babef57704573a3d298e25304e980238`.

## Testing task and verification envelope

1. Read the current Approved Techplan end-to-end, latest Build report first, Review finding, and target repository commands. Follow canonical Testing `SWEEP, DON'T REDO`: use Build-reported checks as claims, spot-check its focused evidence, and independently execute Testing-owned rows.
2. Cover R1–R5 in Techplan §4/§12. Run `cd api && npm run validate` and independently inspect the authored split Donation POST response. Confirm generic `application/problem+json`/shared `Problem`, no internal reason or unsupported `Retry-After`, and no unrelated contract drift. Record warning/error totals and any warning delta only if it can be compared reliably.
3. Independently inspect the Donation POST 503 correspondence in `api/openapi.yaml` and `frontend/lib/api/generated/openapi.ts` against authored source; do not regenerate or hand-edit generated files. Check that generated changes are limited to this operation response.
4. Run the Testing-owned target-repo final command `cd frontend && npm run verify` (`frontend/package.json` defines it as lint plus test). Inspect the focused flow test assertions for: (a) only this operation's 503 maps to generic `request-failure`, clears retry intent, and exposes no ambiguous retry action; (b) transport and other 5xx remain ambiguous and retry with the same key and payload; (c) existing 201/409/422 paths remain; and (d) mock response stays generic and makes no backend/runtime claim. Testing must independently report actual command results; Build's focused 16-test result is context, not Testing evidence.
5. Address Review BP-1 explicitly: determine whether the fixture shape and observable MSW assertions actually verify the required Problem response. Treat schema typing as a non-blocking maintenance follow-up unless the independent checks reveal a real coverage/contract defect. Do not edit the fixture in Testing; any required code change needs a narrow patch plan for Build authority.
6. Test Focus Pointer: the only row is marked `N/A for specialized testing in WU-S2-007` because this Run changes no backend admission/runtime path; R1 covers the API shape. State that no specialized atomicity/concurrency/runtime security verification applies here and preserve those as downstream obligations. Do not claim real Funding availability/admission, backend non-disclosure, concurrency, performance/load, migration, or runtime verification. Do not scan other Exploration material. Open only `runs/EXP-S2-007-001/evidence/stage-2-gap-analysis.md#Area 1 — Accepted Donation policy` if the current Techplan/pointer or a concrete concern makes that N/A classification material or ambiguous.
7. No browser check is required: no UI/copy/state changed, and Techplan/target authority assign final verification to API validation and `frontend verify`. No database/migration check applies. Include applicable compatibility evidence from R2/R4; do not imply external-consumer acceptance.
8. Freshly read the entire current Approved Techplan after executing checks and report any contradiction/gap. If changed-path hashes or assignment scope drift, stop and hand off drift rather than testing an unreviewed change.

Write `runs/TST-S2-007-001/testing-report-1.md` with canonical Sections 0 Sweep Summary, 0a Test Focus Pointer Execution, 1 Test Coverage, 2 Error Verification, 3 Final Verification, 4 New Recurring Bug Patterns, Verdict, and exactly one structured `## Phase handoff` per the overlay. Verdict options: `Pass`, `Pass with flagged follow-ups`, or `Fail — send back to Build`. Distinguish blocking failures from non-blocking BP-1 follow-up and from Human-owned API source acceptance. If a code fix is needed, also write one narrow `patch-plan-1.md` and stop at Testing; do not make the fix.

## Execution envelope

- `PREAUTHORIZED`: read the pinned/current effective inputs; verify hashes/scope; run `cd api && npm run validate` and `cd frontend && npm run verify`; inspect the authored/generated response, observable tests, and whole Techplan; write Testing report and a patch plan only if a real code change is required.
- `HUMAN_REQUIRED`: accept exact API source/counterpart bytes, close WU-S2-007, release WU-S2-003's HARD dependency, change Product/API policy, expand scope, edit implementation/test code, or claim backend/runtime/delivery completion.
- Do not perform API-owner acceptance, change the working tree, dispatch another phase, or imply that Testing alone completes WU-S2-007.

## Human-assisted dispatch

Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`. Fresh Tester / `KC-TESTER`, `gpt-6-luna` / `medium`.

Kickoff: `Jalankan independent Testing Run TST-S2-007-001 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-007/runs/TST-S2-007-001/invocation.md dan canonical ../harscode-workspace/workflow/5-testing-prompt.md beserta orchestrated-run overlay. Baca latest Build report lebih dulu, lalu seluruh Approved Techplan ff0cb704151495e9f5539d11c24e808e8618d1c98cdb0d680957df8d51948dc3, target-repo testing authority, dan Review findings. Verifikasi current hashes/scope terhadap Invocation. Jalankan Testing-owned checks: cd api && npm run validate; cd frontend && npm run verify. Periksa correspondence authored/bundle/types, seluruh R1–R5, error/retry semantics, Review BP-1 fixture, dan Test Focus Pointer yang memang N/A untuk backend/runtime specialized testing. Jangan mengubah code/test atau menjalankan generator. Tulis testing-report-1.md dengan semua section dan satu structured Phase handoff; patch-plan-1.md hanya jika ada genuine change required. Berhenti setelah Testing; jangan menerima exact API source atau dispatch fase berikutnya.`
