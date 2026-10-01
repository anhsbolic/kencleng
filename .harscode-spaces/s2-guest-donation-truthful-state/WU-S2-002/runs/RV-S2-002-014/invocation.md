# Run Invocation — `RV-S2-002-014`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-01 by Orchestration Operator. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `RV-S2-002-014`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-014`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Targeted independent confirmation of blocking finding F-1 from RV-S2-002-013
- `PARTICIPANT_ID`: `P-S2-002-RV-014-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; current Profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — new Reviewer Run after completed Build/Patch BLD-S2-002-007.
- `SESSION_TRANSITION_REASON`: Independent confirmation of the exact Review-requested schema patch using a new Reviewer Participant and fresh Session/context.
- `TARGET_REVISION`: Kencleng checkout `055aa1283ff5f1ab9c08c1422981a9cbccf98e6c` plus current BLD-006/BLD-007 authored and generated API working-tree changes; inspect the live assigned schemas and types at dispatch.
- `WORKFLOW_REVISION`: Harscode checkout `33b03a3f62cc3aacba6534b8a011465613c64b09`; re-read current Code Review guidance at dispatch.
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `SELECTED_MODEL`: `gpt-6-luna`; Human-owned local configuration marks approval not required.
- `REASONING_EFFORT`: `high`.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Narrow targeted Review re-entry requested by RV-S2-002-013 after Build/Patch implemented its exact F-1 patch plan. Confirm the three authored required-field declarations and corresponding generated bundle/type contract. Do not repeat the full four-pass review unless the patch materially changed behavior/contract beyond F-1 or introduced a new material concern.
- `RISK_TIER`: Tier 1 — confirm the explicit currency pairing for existing monetary projections; the patch itself only changes requiredness and mechanically generated outputs.
- `MODEL_ROUTING_RATIONALE`: The assigned question is a bounded comparison against one accepted finding/patch plan and its generated correspondence. `gpt-6-luna` at `high` is the configured Reviewer-capable minimum sufficient route; escalate only for demonstrated capability insufficiency, not missing authority/context.

## Current-effective inputs and confirmation scope

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-013/review-findings-1.md` — blocking F-1, its impact, and authority.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-013/patch-plan-1.md` — exact requested correction and acceptance evidence.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-007/patch-report-1.md` and `launch-record.md` — patch handoff and reported verification; inspect live contract/types independently.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/techplan.md` Q12 and current `docs/project/kencleng-monetary-data-standard.md` — governing major-unit amount plus explicit currency-code direction.
- Current Task 02: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-002/tasks/02-donation-openapi-reconciliation.md`.
- Inspect only these affected authored schemas in `api/openapi/donation.yaml`: `DonationListItem`, `ClaimableDonation`, and `MyDonation`.
- Inspect their generated counterparts in `api/openapi.yaml` and `frontend/lib/api/generated/openapi.ts`. Confirm each generated type requires both `amount` and `currency_code` (rather than optional properties) and that source, bundle, and types correspond.
- Ignore unrelated BLD-006 API contract changes and pre-existing spec/orchestration changes; do not expand review scope unless this patch appears to introduce a material collateral change.
- Target-repo convention: root `AGENTS.md` and `api/README.md` (split Donation source is authored; aggregate and frontend types are generated).
- Canonical current Harscode: `../harscode-workspace/workflow/4-code-review-prompt.md`, `workflow/4-code-review/guidelines.md`, `workflow/4-code-review/checklist.md`, `workflow/orchestrated-run-overlay.md`, `workflow/context-management.md`, `workflow/AGENTS.md`, `orchestration/run-contract.md`, and `orchestration/AGENTS.md`.

## Review question and completion condition

Independently confirm whether BLD-S2-002-007 resolves F-1 exactly:

- `amount` and `currency_code` are both required in each of the three assigned authored response schemas.
- Generated aggregate and frontend TypeScript counterparts require the same fields and correspond to authored source.
- The narrow correction preserves the existing field definitions and introduces no material collateral change within the reviewed patch scope.
- No currency range/precision/storage policy or runtime behavior is inferred; runtime, security, residual-risk, and `CONTRACT_READY` gates remain separate.

This is a read-only targeted confirmation of one patch plan, not a repeat of all four review passes. Do not edit production/spec/generated files or projections. Do not run tests or runtime/security checks. A focused read-only command is allowed only to resolve a concrete uncertainty about generated field requiredness or source/bundle/type correspondence.

Write only this Run's `review-confirmation.md` and `launch-record.md`. If F-1 is resolved with no new material issue, state the original BLD-006 Review loop is complete for the reviewed contract diff and hand off to Orchestration for the next applicable API-owner/Testing gates. If F-1 remains unresolved or a material collateral issue is found, describe the exact anchor and write `patch-plan-1.md` only when another source correction is required. Do not claim product/runtime/security proof, residual-risk acceptance, owner acceptance not explicitly recorded, or `CONTRACT_READY`.

## Execution envelope

- `PREAUTHORIZED`: read F-1 and its patch plan, BLD-007 handoff, the named Techplan/Task/monetary authority, the three live schemas and generated counterparts, and targeted Review guidance; perform focused read-only comparison; write this Run's confirmation and launch record (and a narrow patch plan only if required).
- `ORCHESTRATOR_DECISION`: broaden Review scope or change the route.
- `HUMAN_REQUIRED`: edit source, choose new monetary/API policy, accept residual risk, claim owner acceptance not given, or promote `CONTRACT_READY` / a milestone.

## Human-assisted dispatch

Use a fresh Reviewer Session from the Kencleng repository root with the configured model/effort. Start from `../harscode-workspace/workflow/4-code-review-prompt.md` and this Invocation. Kickoff:

`Jalankan targeted independent Review Run RV-S2-002-014 sesuai Invocation durable di .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-014/invocation.md dan canonical Code Review guidance Harscode saat ini. Konfirmasi hanya blocking F-1 dari RV-013 terhadap patch-plan-1 dan BLD-007 patch-report-1: inspeksi source live DonationListItem, ClaimableDonation, dan MyDonation; pastikan amount serta currency_code required pada ketiganya; inspeksi aggregate api/openapi.yaml dan frontend/lib/api/generated/openapi.ts untuk correspondence dan required TypeScript properties. Cocokkan dengan TP-015 Q12, Task 02, monetary standard, root AGENTS.md, dan api/README.md. Ini targeted confirmation karena patch mengikuti patch plan sempit; jangan ulang full four-pass kecuali menemukan perubahan collateral yang material. Jangan edit source/spec/generated/projection, jangan jalankan tests/runtime/security checks. Tulis hanya review-confirmation.md dan launch-record.md; patch-plan-1 hanya bila koreksi source lain memang diperlukan. Laporkan apakah F-1 resolved dan handoff ke orchestration untuk API-owner/Testing routing; jangan klaim runtime/security proof, residual-risk acceptance, owner acceptance yang belum diberikan, CONTRACT_READY, atau milestone. Berhenti setelah handoff.`
