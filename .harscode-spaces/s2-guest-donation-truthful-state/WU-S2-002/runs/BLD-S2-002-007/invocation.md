# Run Invocation — `BLD-S2-002-007`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-01 by Orchestration Operator. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and technical identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `BLD-S2-002-007`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-007`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `ROLE`: Implementer
- `SPECIALIZATION`: Narrow Build/Patch for blocking Review finding F-1 in the Task 02 Donation OpenAPI contract
- `PARTICIPANT_ID`: `P-S2-002-BL-007-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-IMPLEMENTER`; use the current Human-owned Participant Profile at dispatch.
- `SESSION_TRANSITION`: `FRESH` — new Build/Patch Run after the completed RV-013 Review occurrence.
- `SESSION_TRANSITION_REASON`: Orchestrated Review-to-Build re-entry requires a new Run, Participant, and fresh Session; re-ground only on the Approved spine, current Task 02, RV-013's specific patch plan, and live affected schemas/generated artifacts.
- `TARGET_REVISION`: Kencleng `055aa1283ff5f1ab9c08c1422981a9cbccf98e6c` plus current working-tree changes including BLD-006's uncommitted authored/generated API diff; inspect live sources at dispatch.
- `WORKFLOW_REVISION`: Harscode checkout `33b03a3f62cc3aacba6534b8a011465613c64b09`; resolve canonical Build guidance at dispatch.
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `SELECTED_MODEL`: `gpt-6-luna`; Human-owned local configuration marks approval not required.
- `REASONING_EFFORT`: `high`.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Required narrow Build/Patch re-entry for Review finding F-1 only. After the patch, return to the requesting Review phase for targeted independent confirmation of F-1 and generated correspondence; do not automatically restart a full four-pass review loop.
- `RISK_TIER`: Tier 1 — this contract surface carries monetary wire values, though the correction only makes existing amount/currency fields mandatory in three response projections and regenerates derived artifacts.
- `MODEL_ROUTING_RATIONALE`: This is a bounded three-schema required-field correction plus documented bundle/type regeneration. `gpt-6-luna` is the Human-configured low-cost Implementer-capable model; `high` is the established sufficient effort for this Task 02 API contract. Escalate only if concrete capability insufficiency is demonstrated, not to resolve missing authority.

## Current-effective inputs

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/techplan.md` — current-effective Human-approved `Approved` spine; material planning authority.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-002/tasks/02-donation-openapi-reconciliation.md` and `tasks/manifest.md` — accepted Task 02 execution boundary and verification.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-013/review-findings-1.md` — blocking finding F-1 and exact schema/type impact.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-013/patch-plan-1.md` — sole patch scope and acceptance evidence for this re-entry.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-006/report.md` — prior authored source, generated outputs, verification, and limitations; inspect current diff rather than treating report as a source substitute.
- Current live authority: root `AGENTS.md`; `docs/project/kencleng-monetary-data-standard.md`; relevant agreed Donation specs under `docs/spec/5-donation/`; `api/README.md`; `frontend/AGENTS.md`.
- Canonical current Harscode: `../harscode-workspace/workflow/3-build-prompt.md`, `workflow/3-build/guidelines.md`, `workflow/3-build/checklist.md`, `workflow/orchestrated-run-overlay.md`, `workflow/context-management.md`, `workflow/AGENTS.md`, `orchestration/run-contract.md`, and `orchestration/AGENTS.md`.

## Patch target and boundaries

- Resolve RV-013 F-1 only: in `api/openapi/donation.yaml`, make both `amount` and `currency_code` required in the `DonationListItem`, `ClaimableDonation`, and `MyDonation` response schemas.
- Reopen those live schemas and their existing generated counterparts before editing. Preserve field schemas, descriptions, other response behavior, and all settled contract decisions.
- Regenerate `api/openapi.yaml` and `frontend/lib/api/generated/openapi.ts` using the documented workflow; never hand-edit generated outputs.
- Keep `api/openapi/common.yaml`, `api/openapi/index.yaml`, specs, Product/MVP, Design, Techplan, runtime code, tests, migrations, Tier-0 paths, and orchestration projections unchanged in this Run.
- Preserve existing unrelated working-tree edits and the larger BLD-006 diff. The patch delta must contain only the three required-field changes and their mechanical generated consequences.
- Do not choose monetary range/precision/storage parameters, expand the scope beyond these three projections, claim runtime/security proof or residual-risk acceptance, or declare `CONTRACT_READY`.

## Verification and output

- Run `cd api && npm run validate` after editing; report actual warning/error count.
- Follow `api/README.md` to run `cd api && npm run bundle` and `cd frontend && npm run generate:api-types`; inspect that the corresponding generated fields are required and no unrelated generated change was introduced.
- Run a scoped `git diff --check` for the authored Donation source and the two generated artifacts. Do not add or run product/runtime tests or the broad Testing matrix in this contract-only patch.
- Write only this Run's `patch-report-1.md` and `launch-record.md`. The report must use the canonical Build/Patch structure and state that no concurrency, performance/load, or security-class verification was run.
- Handoff to Orchestration for a fresh, targeted independent Review confirmation of F-1; do not claim that confirmation has already happened.

## Execution envelope

- `PREAUTHORIZED`: read the named current-effective artifacts and live affected schemas/generated files; edit `api/openapi/donation.yaml` only for the three required lists; regenerate `api/openapi.yaml` and `frontend/lib/api/generated/openapi.ts`; run the named validation/generation commands and scoped diff check; write this Run's two artifacts.
- `ORCHESTRATOR_DECISION`: widen the correction scope or change the post-patch route.
- `HUMAN_REQUIRED`: choose new monetary/product/API policy, edit protected files, accept residual risk, or promote `CONTRACT_READY` / a delivery milestone.

## Human-assisted dispatch

Use a fresh Implementer Session from the Kencleng repository root. Start from `../harscode-workspace/workflow/3-build-prompt.md` and this Invocation. Kickoff:

`Jalankan Build/Patch Run BLD-S2-002-007 sebagai Implementer sesuai Invocation durable di .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-007/invocation.md dan canonical Build guidance Harscode saat ini. Reconstruct dari Approved TP-015, Task 02 TPD-002, RV-013 review finding F-1 dan patch-plan-1, serta live affected Donation schemas dan generated outputs; baca ulang sumber aktual sebelum edit. Lakukan hanya koreksi F-1: wajibkan amount dan currency_code pada DonationListItem, ClaimableDonation, dan MyDonation di api/openapi/donation.yaml, lalu regenerate api/openapi.yaml dan frontend/lib/api/generated/openapi.ts sesuai api/README.md. Jangan sentuh common/index/spec/Product/Design/Techplan/runtime/tests/migrations/Tier-0/projections atau perubahan working tree lain. Jalankan npm run validate dari api/, bundle, generate:api-types, dan scoped git diff --check; laporkan hasil aktual dan batas verifikasi. Tulis patch-report-1.md dan launch-record.md, lalu handoff untuk targeted independent Review confirmation; jangan klaim review closure, residual-risk acceptance, runtime proof, CONTRACT_READY, atau milestone.`
