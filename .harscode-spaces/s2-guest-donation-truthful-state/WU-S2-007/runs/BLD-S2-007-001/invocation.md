# Run Invocation — `BLD-S2-007-001`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-04 after Anhar approved the exact WU-S2-007 Techplan Draft and its current-effective `Status` was reconciled to `Approved`. This is the initial Build occurrence; it has not been dispatched.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-007`
- `RUN_ID`: `BLD-S2-007-001`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-007/runs/BLD-S2-007-001`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-007`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Implementer
- `SPECIALIZATION`: Donation POST Funding-unavailable API contract, generated counterpart, and internal consumer reconciliation
- `PARTICIPANT_ID`: `P-S2-007-BLD-001-1`
- `PARTICIPANT_PROFILE_ID`: `KC-IMPLEMENTER`; `.harscode-spaces/participant-profiles/profiles.md`, SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — initial Build after exact Human Techplan approval; create a fresh Implementer context and do not reuse Planner/Reviewer sessions.
- `TARGET_REVISION`: Kencleng HEAD `4e71d3697a479d92b2cad5a79ef5d710f4e8077c` plus current working tree. Reverify all pinned source/anchor hashes at dispatch; stop and reconcile unexplained drift.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; current-effective guidance, re-read at dispatch.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `medium`
- `MODEL_APPROVAL`: Not required by Human-owned local registry.
- `MODEL_REGISTRY_SOURCE`: `.harscode-spaces/.local-config.yaml`, SHA-256 `ddffdeb4ce8edec181ca36bf32bedf19c513a76e359bc2642985eda6b32f400d`; Luna has coding/repository-work capabilities, supports `medium`, and is not approval-gated.
- `MODEL_ROUTING_RATIONALE`: The Human-approved plan fixes the contract and implementation direction. The work follows established OpenAPI generation and frontend test patterns across a bounded set of authored/generated/client files; focused validators, generators, and a flow test provide direct evidence. `gpt-6-luna` / `medium` is the lowest sufficient non-gated capability for this implementation. If live evidence exposes a material semantic/authority conflict or a new cross-cutting design choice, stop and route it rather than escalating the model to invent a resolution.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Canonical initial Build for the entire Approved Techplan. Decomposition was Skip; do not invent a narrower task boundary.
- `ARTIFACT_TARGET`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-007/runs/BLD-S2-007-001/report.md` — sole Run-owned phase report. Stable source target is the Approved Techplan referenced below.

## Current-effective inputs / `PRIOR_ARTIFACTS`

- Approved stable Techplan: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-007/techplan.md`, SHA-256 `ff0cb704151495e9f5539d11c24e808e8618d1c98cdb0d680957df8d51948dc3`. Anhar approved its exact pre-status Draft hash `d9b483180a20dd2b4df618654a13164333816e552d3cc87f2e9787aa891f6928` against the report below; current change from that preimage is Status-only. This is the full Build target; no task decomposition exists.
- Human approval report: `report-techplan.md`, SHA-256 `e022e41909c5aa024b989e278cc9d752b0ae45b751deb039fa33da54411186e6`.
- Independent Techplan Review: `runs/RV-S2-007-001/review-findings-1.md`, SHA-256 `70361a6188bd5ac89c9d9d1001f322e8eea6430f1ef8ba1b3708ac5ede2746f7`; no blockers, one mechanical/non-blocking `(D2)` pointer correction, resolved in TP-S2-007-002 without semantic change. Re-review was not required.
- Techplan resolution/report handoff: `runs/TP-S2-007-002/handoff.md`, SHA-256 `9aa8456b319ddbcad9bea878d2346686dbac3a38fe4bbc7b90c9fa5e05eb8b91`.
- WU-S2-007 manifest, SHA-256 `f66de255960f603e99107eddf76d1b7233429749c1086b0444c56ddd3c7f078e`.
- Root `AGENTS.md`, `frontend/AGENTS.md`, `api/README.md`, `docs/spec/5-donation/invariants.md`, and `docs/spec/5-donation/features/01-submit-donation-settlement.md`; current hashes are pinned in the source baseline table below where applicable.
- Canonical `../harscode-workspace/workflow/3-build-prompt.md`, `workflow/3-build/guidelines.md`, `workflow/3-build/checklist.md`, `workflow/orchestrated-run-overlay.md`, and `orchestration/run-contract.md`; follow current guidance at dispatch.
- Current source of API generation: `api/openapi/donation.yaml` plus referenced shared `api/openapi/common.yaml`; use `api/README.md` for bundle/type generation commands.

## Pinned source/anchor baseline

Recheck these exact bytes and live symbols before editing. If any mismatch is unexplained, stop and report it before making changes.

| Path | SHA-256 | Build relevance |
|---|---|---|
| `api/openapi/donation.yaml` | `609411688477132267847db52ec6c7d78b5bfd50df97d7af4e85c049917ad2ca` | Authored Donation POST operation; add accepted 503 response/description. |
| `api/openapi/common.yaml` | `46bd8fba98d7befe65e87b17f431b0f0a2fa815589db28eb6ded6b00bd9418f8` | Shared `Problem` schema, read-only. |
| `api/openapi.yaml` | `a3a67da1294085d15f2bf0f7aef9deb94bd7527e3bede1593570f9323e1e605d` | Generated bundle baseline; regenerate, never hand-edit. |
| `frontend/lib/api/generated/openapi.ts` | `288296d6e65a7500349126e066b3a4215915a647b0c46358f60e954d41262dc4` | Generated response types baseline; regenerate, never hand-edit. |
| `frontend/lib/api/donation.ts` | `0bda1df8cb0b2f25c27af258eab2d30bb05c72a1704a826d9e3191c286c3d7ab` | `submitDonation` classification baseline. |
| `frontend/app/campaigns/[campaignId]/donate/donation-client.tsx` | `91b571ae8032a22576b40b2acd8e6f2b3d8d4986bb14fe88e40946c1593ad30b` | Caller semantics; inspect, no production change intended. |
| `frontend/mocks/handlers/donation.ts` | `775baf8a9d0f43215b0d43aaa1e8d6d4babef57704573a3d298e25304e980238` | Shared mock; inspect, no change intended unless test proves reusable fixture is necessary and scope is reconciled. |
| `frontend/app/donations/donation-flow.test.tsx` | `82bdd24998c9566aa17948aa123341c002b3e3f232185c523ecf74ffb7f2ba86` | Focused observable 503 and ambiguous-retry evidence. |
| `docs/spec/5-donation/invariants.md` | `4bf7ba48f4980f635ce76a4eb3440be5cf12666fdd99e06dcc8ac3b78fbde39a` | Accepted OI9 policy; read-only. |
| `docs/spec/5-donation/features/01-submit-donation-settlement.md` | `b3f0f99c420c807bb14cccf75a4adcaf43dfdadccae6b727a1c460ef4ed6a7e9` | Accepted OI9 feature contract; read-only. |

Other preparation baseline hashes: `AGENTS.md` `4ec81e122cfeb77b2d13e3b90cc7f4a73e829f97d0627d889793aca847a87716`; `frontend/AGENTS.md` `67355c8ce4384082453b48409039315696849e1831e3c9b143ee53bd91cd86ba`; `api/README.md` `0ec88871d43b32b2b06a82ba345f694b079bfefaef3c2974333b43fad16b434f`; Implementer profile `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`; local model registry `ddffdeb4ce8edec181ca36bf32bedf19c513a76e359bc2642985eda6b32f400d`.

## Task and stop gate

Execute the entire Approved Techplan `§§1–13`, reopening current code/spec anchors before edits. Implement only:

1. Add the accepted generic, fail-closed `503 Problem` response to the authored Donation POST in `api/openapi/donation.yaml`. Reuse the existing Problem schema; disclose no internal reason and add no `Retry-After`.
2. Regenerate `api/openapi.yaml` with `cd api && npm run bundle`, and `frontend/lib/api/generated/openapi.ts` with `cd frontend && npm run generate:api-types`. Never hand-edit generated outputs. Run `cd api && npm run validate` and inspect generated diffs for this operation only; stop on unrelated generated contract change or schema contradiction.
3. In `frontend/lib/api/donation.ts`, map only this Donation POST's contract-defined status `503` to existing `request-failure` before the generic `>=500` branch. Preserve ambiguous outcomes for transport failures and all other 5xx. Do not change `DonationClient.send` UI text/state or retry behavior for other statuses.
4. Add focused test-scoped MSW evidence in `frontend/app/donations/donation-flow.test.tsx` for generic Problem 503/no ambiguous retry affordance and preserve/assert same-key/same-payload retry for transport or another 5xx. Run the focused flow test (`cd frontend && npm run test -- app/donations/donation-flow.test.tsx`) to verify the authored test is executable. Do not broaden to the full Testing-owned suite unless target-repo authority or an implementation failure makes it necessary; explain any broad run.
5. Write `runs/BLD-S2-007-001/report.md` with exact changed-file/source/generated hashes, actual commands/results, deferred Testing evidence, verification-scope confirmation, and one structured canonical `## Phase handoff`.

Stop and report if live source or accepted policy conflicts with the Approved plan, if another accepted meaning makes status-only 503 classification insufficient, if source/generated outputs require unapproved shared API or UX changes, or if exact in-scope diffs cannot be isolated. Do not invent a discriminator or branch-specific user copy.

## Execution envelope

- `PREAUTHORIZED`: read/re-ground the pinned paths; edit `api/openapi/donation.yaml`, `frontend/lib/api/donation.ts`, and `frontend/app/donations/donation-flow.test.tsx` within Approved scope; regenerate the two committed generated outputs; run only the focused Build checks above; write the Run report.
- `HUMAN_REQUIRED`: any new Product/domain/security/API/UX decision, shared-schema or other material scope expansion, changes to backend/runtime or protected paths, acceptance of exact authored/counterpart bytes, WU-S2-007 completion, or WU-S2-003 dependency release.
- Explicitly out of scope: Donation specs, `api/openapi/common.yaml`, `DonationClient.send` production behavior, shared mock handlers/fixtures absent a reconciled need, backend code, migrations/database, browser/runtime work, public rollout, and WU-S2-006 accepted artifacts.
- Testing owns independent final checks including `cd frontend && npm run verify`. Build must confirm explicitly in its report that no concurrency, performance/load, or security-class test was run.

## Human-assisted dispatch

Working directory: `/home/anhar-solehudin/kencleng-workspace/kencleng`. Fresh Implementer / `KC-IMPLEMENTER`, `gpt-6-luna` / `medium`.

Kickoff: `Jalankan initial Build Run BLD-S2-007-001 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-007/runs/BLD-S2-007-001/invocation.md dan canonical ../harscode-workspace/workflow/3-build-prompt.md beserta orchestrated-run overlay. Implement full Approved Techplan hash ff0cb704151495e9f5539d11c24e808e8618d1c98cdb0d680957df8d51948dc3; re-ground every live anchor and verify pinned baselines first. Scope: authored Donation POST 503 response; regenerate API bundle/types using api bundle, validate, generate:api-types commands; update submitDonation to classify only this 503 as existing request-failure; add focused test-scoped MSW no-retry evidence while preserving same-key/payload retry for ambiguous outcomes; run focused flow test. Stop on material conflict or scope expansion. Do not change Donation specs, shared Problem schema, caller UI behavior, shared mock unless reconciled necessary, backend/runtime, migrations, or WU-S2-006 history. Write report.md with hashes, real verification outcomes, deferred Testing checks, no-heavy-check confirmation, and one canonical phase handoff. Stop after Build report; do not route Review or Testing automatically.`
