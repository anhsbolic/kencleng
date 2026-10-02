# Run Invocation — `BLD-S2-006-006`

Status: `READY_FOR_HUMAN_DISPATCH`
Prepared: 2026-10-02. Human-facing prose Bahasa Indonesia.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-006`
- `RUN_ID`: `BLD-S2-006-006`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/BLD-S2-006-006`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Implementer
- `SPECIALIZATION`: Coordinated internal generated API and frontend fixture counterpart reconciliation
- `PARTICIPANT_ID`: `P-S2-006-BLD-006-1`
- `PARTICIPANT_PROFILE_ID`: `KC-IMPLEMENTER`; `.harscode-spaces/participant-profiles/profiles.md`, SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — new counterpart Build after completed source Review, confirmed MVP1 distribution posture, and explicit exact-source acceptance.
- `TARGET_REVISION`: Kencleng HEAD `bd40ff5d292cba8731761c73c727243af4f2dccf` plus current working tree. Re-ground the accepted split sources and generated targets before edits. Stop on any changed accepted source hash or unexplained target drift.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`; current-effective and not semantically pinned. Re-ground applicable guidance at dispatch.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `medium`
- `MODEL_APPROVAL`: Not required by Human-owned registry.
- `MODEL_REGISTRY_SOURCE`: `.harscode-spaces/.local-config.yaml`, SHA-256 `ddffdeb4ce8edec181ca36bf32bedf19c513a76e359bc2642985eda6b32f400d`; `gpt-6-luna` has coding/repository-work capabilities, supports `medium`, cost tier `low`, `approval_required: false`.
- `MODEL_ROUTING_RATIONALE`: Low-cost coding-capable model at medium effort is sufficient for deterministic bundle/type generation and bounded fixture reconciliation under an accepted source contract. Stop and route if generated output reveals a contract contradiction or an unplanned consumer/semantic decision.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Initial counterpart Build after accepted authored-source Build and independent Review. This Run does not implement the feature's backend/frontend delivery behavior.
- `ARTIFACT_TARGET`: Generated `api/openapi.yaml`, generated `frontend/lib/api/generated/openapi.ts`, and contract-facing public Campaign fixtures/mock consumers that must match the accepted PublicCampaignDetail response. Do not hand-edit generated outputs.

## Current-effective inputs / PRIOR_ARTIFACTS

- Canonical `../harscode-workspace/workflow/3-build-prompt.md`, `workflow/3-build/guidelines.md`, `workflow/3-build/checklist.md`, `workflow/orchestrated-run-overlay.md`, and `orchestration/run-contract.md`.
- Root `AGENTS.md`, `frontend/AGENTS.md`, `api/README.md`, `.harscode-spaces/participant-profiles/profiles.md` / `KC-IMPLEMENTER`, current Product/MVP and monetary authorities, WU-S2-006 manifest, Events, Work Graph, Control Surface, and the current accepted Campaign/Donation source specifications.
- Approved Techplan `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-008/techplan.md`, SHA-256 `93c09bf7629500cd8fb80fd59b6af464b169484419d722a269b782b78bbbf438`, especially §§6, 9, 11, and 13.
- Accepted authored-source Review `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/RV-S2-006-006/review-findings.md`, SHA-256 `df2edd50f7cdc796dadae817ff5a13fe7d8977e87d7ebdf4e18ebdaa3c43bb02`, verdict `Approve with minor comments`; RV-006-01 is resolved for current MVP1 by the owner receipt in parent `events.md`.
- Build source report `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/BLD-S2-006-005/report.md`, SHA-256 `b9f0670fcc6bbb41a49ac6f09c4fb44bdaead9f0ebb486851b3c4e89b947d806`.
- Owner acceptance and current MVP1 distribution/rollout posture are recorded in WU-S2-006 manifest, parent Events, and `docs/project/kencleng-development-tracker.md`. The exact accepted source identities are repeated below. Do not reopen historical O8 or infer universal future external-client compatibility.

## Accepted authored-source snapshot

All seven live hashes were verified against the exact accepted `RV-S2-006-006/invocation.md` snapshot before this preparation.

| Accepted source | SHA-256 |
|---|---|
| `docs/spec/4-campaign/invariants.md` | `b46797a183a12f818989ca16553751a001cb358abfda02f6570af4a02c98c27f` |
| `docs/spec/4-campaign/features/01-campaign-creation-draft-crud.md` | `06dd64a2f4df04193ecd5d89dfc64f599df142d93997f6bf7222ee68c02e255c` |
| `docs/spec/4-campaign/features/02-campaign-detail-listing.md` | `88758edd31f5d0d867bc02b8a5b7d425d571da0097bbc4d96cf91e4978444aff` |
| `docs/spec/5-donation/invariants.md` | `adf13ce0eb23ff6a29bcf54fa6c63ab7cd1b685bf77531198831d414328628b0` |
| `docs/spec/5-donation/features/01-submit-donation-settlement.md` | `b65511829e97276f6df3a109061a6d8f5506cefd6b6bee2e08a7b016ed8c42a3` |
| `api/openapi/campaign.yaml` | `12a20e4be31b32df8ee73794400ce73f2ae15f2c8b48b9107a377cc83fba7226` |
| `api/openapi/donation.yaml` | `609411688477132267847db52ec6c7d78b5bfd50df97d7af4e85c049917ad2ca` |

## Task and stop gate

Execute the counterpart portion of approved TP-S2-006-008 §9 using the accepted Campaign/Donation split sources:

1. Follow `api/README.md`: regenerate the aggregate bundle from split sources using `cd api && npm run bundle`; do not hand-edit `api/openapi.yaml`.
2. Regenerate frontend OpenAPI types with `cd frontend && npm run generate:api-types`; do not hand-edit the generated TypeScript file.
3. Update the public Campaign fixture/mock data and only the contract-facing frontend fixtures/consumers that need the newly required `{ amount, currency_code }` response field. Keep mock responses aligned with the accepted public schema across funding availability states. Do not add cap-disclosure presentation or other UI behavior in this Run; that behavior belongs to the refreshed WU-S2-004 delivery plan.
4. Run the focused frontend public Campaign API/client fixture tests affected by the generated type/fixture delta, plus the target-repo checks necessary to show generated correspondence. Record actual commands/results and avoid broad suites.

Re-ground on live generated/API/fixture consumers before editing. Keep the accepted seven authored sources byte-identical. Do not change domain/spec/API source semantics, `common.yaml`, backend production code, database/migrations, or cross-domain behavior. No user-facing display, form, or campaign detail UI behavior changes are in this Run.

The current backend `publicCampaignDetailResponse` does not yet contain `max_donation_amount`; the existing exact-wire test therefore cannot be changed into a passing backend counterpart without its response DTO/mapping. Do not edit that test or backend production code in this WU006 Run. Orchestrator has routed the exact-wire assertion update together with the WU-S2-003 response implementation under its fresh Techplan, and it must be completed before backend runtime verification. This Run must report that explicit downstream assignment. WU-S2-004's visible public cap disclosure remains a required frontend delivery-plan item before frontend runtime/rendered acceptance.

Stop and report if generated output or fixture correspondence appears to require changing accepted source semantics, if accepted hashes drift, or if a consumer update requires Product/Design/API choice. No WU-S2-006 completion, runtime/delivery milestone, or source revision acceptance is implied.

Write `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/BLD-S2-006-006/report.md` with canonical provenance, generated/fixture changes, exact snapshots, actual checks, known deferred cross-stack items, and exactly one structured `## Phase handoff` containing: `Outcome`, `Result refs`, `Findings`, `Decision requests`, `Blockers`, `Open / unverified`, `Recommended continuation`, and `Context refs`.

## Execution envelope

- `PREAUTHORIZED`: inspect accepted current source/generated/frontend fixture consumers; regenerate API bundle and frontend types; update contract-facing fixtures/mock data; run focused frontend fixture/client verification; write this Run report.
- `ORCHESTRATOR_DECISION`: inspect the output correspondence and determine when WU006 counterpart convergence is sufficient to start fresh WU-S2-003 and WU-S2-004 Techplan refreshes.
- `HUMAN_REQUIRED`: new Product/Design/API decision, changing accepted source semantics, protected Tier-0 work, manual DB/index application, or residual-risk acceptance.

## Human-assisted dispatch

Cwd `/home/anhar-solehudin/kencleng-workspace/kencleng`; fresh Implementer / KC-IMPLEMENTER, `gpt-6-luna` / `medium`.

Kickoff: `Jalankan coordinated counterpart Build Run BLD-S2-006-006 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/BLD-S2-006-006/invocation.md dan canonical ../harscode-workspace/workflow/3-build-prompt.md dengan orchestrated-run overlay. Berhenti setelah report dan phase handoff; jangan mengubah tujuh accepted source files.`

Laporkan report/handoff atau exact scoped discrepancy kepada Orchestrator. Run ini tidak mengimplementasikan response DTO, melakukan delivery/runtime progression, atau menyatakan WU006 complete.
