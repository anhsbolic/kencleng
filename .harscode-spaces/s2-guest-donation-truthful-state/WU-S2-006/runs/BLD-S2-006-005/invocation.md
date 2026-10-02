# Run Invocation — `BLD-S2-006-005`

Status: `READY_FOR_HUMAN_DISPATCH`
Prepared: 2026-10-02. Human-facing prose Bahasa Indonesia.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-006`
- `RUN_ID`: `BLD-S2-006-005`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/BLD-S2-006-005`
- `WORK_UNIT_PATH` / `TASK_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006`
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `ROLE`: Implementer
- `SPECIALIZATION`: Reconcile remaining approved monetary/capacity feature specs and authored split OpenAPI sources
- `PARTICIPANT_ID`: `P-S2-006-BLD-005-1`
- `PARTICIPANT_PROFILE_ID`: `KC-IMPLEMENTER`; `.harscode-spaces/participant-profiles/profiles.md`, SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — new source-authoring Build after the Human-approved TP-S2-006-008 plan gate.
- `TARGET_REVISION`: Kencleng HEAD `7fd8b473b239b20bda3990ab29c51440d321a796` plus current working tree; verify live source immediately before edits.
- `WORKFLOW_REVISION`: Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`; use current applicable guidance at dispatch.
- `SELECTED_MODEL`: `gpt-6-luna`
- `REASONING_EFFORT`: `high`
- `MODEL_APPROVAL`: Not required by Human-owned registry.
- `MODEL_ROUTING_RATIONALE`: Lowest-cost available coding model; high effort is justified by cross-domain monetary/error/object-schema correspondence across accepted specs and authored OpenAPI. No stronger-model escalation is supported by current evidence.
- `COMMUNICATION_LANGUAGE`: Bahasa Indonesia
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Initial Build against the whole unsplit approved WU-S2-006 spine; authored sources only, stopping before independent Review and owning source acceptance.

## Current-effective inputs / PRIOR_ARTIFACTS

- Canonical `../harscode-workspace/workflow/3-build-prompt.md`, `workflow/3-build/guidelines.md`, `workflow/3-build/checklist.md`, and `workflow/orchestrated-run-overlay.md`.
- Root `AGENTS.md`, `.harscode-spaces/participant-profiles/profiles.md` / `KC-IMPLEMENTER`, API source-routing instructions in `api/README.md`, current Product/MVP, monetary standard, applicable Campaign/Donation specs, Authority Map, WU-S2-006 manifest, parent events, Work Graph and current accepted WU-S2-002/WU-S2-005 baselines.
- Current approved spine `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-008/techplan.md`, SHA-256 `93c09bf7629500cd8fb80fd59b6af464b169484419d722a269b782b78bbbf438`; its Status-only preimage at approval was `c1a8806a1c754b849c0b8457e688d9a50aa0d024dc3d4fe2f3c4c5a7d35c3324`, independently reconstructed and matched. Whole-plan approval receipt and unchanged matching report SHA-256 `06d6259f4e50960dbae04951b0b786c15503ec6140f27b9f805149a8a5b6eb6c` are in parent `events.md`, “2026-10-02 — Human approved TP006008; Status propagation prepared”.
- Existing exact owner decisions: DEC-API-01/02 in parent `events.md`, “2026-10-02 — Owner settled API transport/encoding; Planner propagation prepared”; no decision repeat is needed.
- Prior concrete accepted Product/MVP and six Campaign/Donation source snapshots are listed in `WU-S2-006/manifest.md`. They establish the prior source acceptance only; any newly changed source bytes require their own applicable Review and owning acceptance.

### Run-entry source coordinates

These are last observed snapshots, not permission to overwrite concurrent changes. Re-read and record exact Run-entry state before each write; stop on unexplained relevant drift.

| Source | SHA-256 at preparation |
|---|---|
| `docs/spec/4-campaign/invariants.md` | `3ca67e64aadd0d210c4de5a789ccc60172ea1a4e7f25de599dc6793db5e38931` |
| `docs/spec/4-campaign/features/01-campaign-creation-draft-crud.md` | `59945e6437b0356936fe92f0778f7529d4155bc3036db95476df473d1d59aaa0` |
| `docs/spec/4-campaign/features/02-campaign-detail-listing.md` | `4ad59a8c1ec3e1a5c188f5a63dee72e2661852decd455903e35fb46aab775fae` |
| `docs/spec/4-campaign/features/09-closure.md` | `31f08f4f3f6e61ee295f3049c8ff7103f3312e33df9ed9fb8ff720ab840a4740` |
| `docs/spec/5-donation/invariants.md` | `af3cbe68819a045a4db3479d704f228043de7476cc81bbc678a0fe13c1b5de9f` |
| `docs/spec/5-donation/features/01-submit-donation-settlement.md` | `85cda7bf84bdcbe7013fe7e9e9a26cab061219988f69009f7841203a2f153174` |
| `api/openapi/campaign.yaml` | `65464f9aa177579a4a1536dc1d599160e3eab18eb1ab687feb9cdcf04c0f3134` |
| `api/openapi/donation.yaml` | `873834d59695fa8ae6011f6958101a496581d3f9c03a84ed28ab6f2263e5ea88` |
| `api/openapi/common.yaml` | `46bd8fba98d7befe65e87b17f431b0f0a2fa815589db28eb6ded6b00bd9418f8` |

## Task and stop gate

Execute Techplan TP-S2-006-008 §§1–13 against current source: author only the affected Campaign/Donation feature-spec clarifications and split authored OpenAPI changes required for the accepted per-Campaign cap, closed `{amount: decimal-string, currency_code: IDR}` shape, generic eligible capacity no-fit `422 ValidationError` on `amount`, preserved closed/ineligible `409` and idempotency semantics, and `funding_capacity_reached` close reason. Campaign feature wording should receive the concrete closed-object encoding where applicable. Do not alter the already accepted Product/MVP amendments unless live evidence reveals a contradiction; if so, stop and route it to the owning authority.

Keep source semantics bounded by the Approved Techplan and current authorities. Preserve Slice-2 closed/non-public `404` and accepted availability-only Campaign action; no capacity disclosure, new reason in public projection, or silent change to WU-S2-005. Use `common.yaml` only if the approved shared `ValidationError` component needs a concrete schema/description change; explain necessity and blast radius. Do not author aggregate/generated API, frontend types, fixtures, known-consumer updates, migrations, production/backend/frontend code, or tests.

Run the focused `cd api && npm run validate` source validation required by the approved source plan and `api/README.md`. Record actual diagnostics and compare warning coordinates/rules to the known 124-warning baseline; touched source must have zero errors and no new warning coordinates. This is schema validation, not runtime or independent Testing evidence. Do not run tests or generators in this Run.

Write the Build report with canonical provenance, changed sources, exact before/after hashes and Run-only deltas, validation actually run, limitations, and one structured `## Phase handoff`. Stop after authored-source Build. Newly changed spec/API bytes require independent Review and the relevant spec/API owner acceptance before counterparts or WU-S2-006 completion. No whole-plan re-approval unless the material meaning changes; any material plan/source-authority contradiction routes back to the Orchestrator.

## Execution envelope

- `PREAUTHORIZED`: read current authority and sources; author bounded affected `docs/spec/` and `api/openapi/` bytes under the Approved plan; run the focused OpenAPI source validator; produce snapshots, Run-only diff and Build report/handoff.
- `ORCHESTRATOR_DECISION`: inspect actual delta, decide applicable Review/acceptance route and later counterpart ordering within the approved scope.
- `HUMAN_REQUIRED`: new Product/API/domain choice, concrete acceptance of changed authoritative spec/API bytes, protected Tier-0 write, manual DB/index application, or residual-risk acceptance.

## Human-assisted dispatch

Cwd `/home/anhar-solehudin/kencleng-workspace/kencleng`; fresh Implementer / KC-IMPLEMENTER, `gpt-6-luna` / `high`.

Kickoff: `Jalankan Build Run BLD-S2-006-005 sesuai .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/BLD-S2-006-005/invocation.md dan canonical ../harscode-workspace/workflow/3-build-prompt.md dengan orchestrated-run overlay. Berhenti setelah report dan phase handoff pada gate yang berlaku.`

Laporkan report/handoff atau discrepancy material kepada Orchestrator. Run ini tidak menerima source bytes, memulai counterpart generation, atau menyatakan WU006 complete.
