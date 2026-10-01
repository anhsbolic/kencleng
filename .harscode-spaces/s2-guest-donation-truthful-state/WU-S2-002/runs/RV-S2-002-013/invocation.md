# Run Invocation — `RV-S2-002-013`

Status: `READY_FOR_HUMAN_DISPATCH`

Prepared: 2026-10-01 by Orchestration Operator. Human-facing language: Bahasa Indonesia; retain canonical Harscode terms/enums and code/API/schema identifiers.

## Identity and assignment

- `WORK_UNIT_ID`: `WU-S2-002`
- `RUN_ID`: `RV-S2-002-013`
- `RUN_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-013`
- `WORK_UNIT_PATH`: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002`
- `ROLE`: Reviewer
- `SPECIALIZATION`: Independent four-pass review of Task 02 authored Donation OpenAPI reconciliation and generated outputs
- `PARTICIPANT_ID`: `P-S2-002-RV-013-1` (ephemeral for this Run)
- `PARTICIPANT_PROFILE_ID`: `KC-REVIEWER`; current Profile content SHA-256 `e545651805b302727bcb89087484b767dc7046d470567f6679c8bb8362771d32`.
- `SESSION_TRANSITION`: `FRESH` — new independent Reviewer Run after completed Build/Patch BLD-S2-002-006.
- `SESSION_TRANSITION_REASON`: Review independence; use a new Reviewer Participant and Session grounded only in the current Approved spine/task, exact API diff, Build handoff, and relevant live authorities.
- `TARGET_REVISION`: Kencleng checkout `650e73c5d646c29c0ddf1931618f02685d15f7b7` plus BLD-006's authored/generated API diff; inspect exact named paths at dispatch.
- `WORKFLOW_REVISION`: Harscode checkout `33b03a3f62cc3aacba6534b8a011465613c64b09`; re-read canonical Review and best-practice routing guidance at dispatch.
- `HARSCODE_WORKSPACE_ROOT`: `../harscode-workspace`
- `RUNTIME_HARNESS`: `codex-cli`
- `SELECTED_MODEL`: `gpt-6-luna`; Human-owned local configuration marks approval not required.
- `REASONING_EFFORT`: `high`.
- `COMMUNICATION_PROFILE_PATH`: `docs/project/communication-profile.md`
- `PHASE_ROUTE`: Required full independent Code Review after material Task 02 contract authoring. Run all four passes in order against exactly the BLD-006 diff. If fixes are required, write a narrow patch plan for a fresh Build/Patch Run; if approved, return to the orchestrator for Testing/owner-acceptance routing. Review does not accept residual risk or claim `CONTRACT_READY`.
- `RISK_TIER`: Tier 1 — contract changes cover monetary wire values, simulated settlement, idempotency, guest PII/status credentials, anti-enumeration, and Campaign ordering.
- `MODEL_ROUTING_RATIONALE`: Independent cross-file/API review of the authored Donation source, shared idempotency component, and regenerated bundle/types. `gpt-6-luna` is the Human-configured low-cost Reviewer-capable model; `high` is the established sufficient effort for Tier-1 Slice 2 review. Escalate only if the reviewer demonstrates concrete capability insufficiency, not to supply missing authority.

## Current-effective inputs and exact diff

- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TP-S2-002-015/techplan.md` — current-effective Human-approved `Approved` spine and the contract/risk/verification/`CONTRACT_READY` authority.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/TPD-S2-002-002/tasks/02-donation-openapi-reconciliation.md` and `tasks/manifest.md` — accepted Task 02 execution contract and hard dependency.
- `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/BLD-S2-002-006/report.md` and `launch-record.md` — Build handoff, actual generated outputs, exact verification and limitations; independently inspect the source diff.
- Review exactly these BLD-006 changed paths by reading the actual diff and current files:
  - `api/openapi/donation.yaml` — authored Donation submit/status operations and schemas.
  - `api/openapi/common.yaml` — shared `Idempotency-Key` contract component.
  - `api/openapi.yaml` — mechanically bundled aggregate generated from split sources.
  - `frontend/lib/api/generated/openapi.ts` — generated frontend API types.
- Do not include the earlier current Task 01 spec status/provenance changes, orchestration projections, or other pre-existing working-tree edits in this Review scope. They are already reconciled independently. Do not use the Build report as a substitute for reviewing the actual four-path diff.
- Current domain/product authority: root `AGENTS.md`; `docs/product/README.md`, `docs/product/mvp-scope.md`, `docs/product/mvp-delivery-slices.md` §§5–6; the now-`agreed` five files under `docs/spec/5-donation/`; `docs/project/kencleng-monetary-data-standard.md`; and applicable sections of `docs/kencleng-agentic-workflow.md`.
- API and generated-code conventions: `api/README.md`; `frontend/AGENTS.md`; `frontend/package.json` only if needed to understand generated types.
- Canonical current Harscode: `workflow/4-code-review-prompt.md`, `workflow/4-code-review/guidelines.md`, `workflow/4-code-review/checklist.md`, `workflow/orchestrated-run-overlay.md`, `workflow/context-management.md`, `workflow/AGENTS.md`, `orchestration/run-contract.md`, and `orchestration/AGENTS.md`.
- Stack-specific routing: read `best-practices/AGENTS.md`; scan only relevant trigger/security rows in `best-practices/index.md` and its Security Concern Map. Likely matches to confirm against the actual diff include `restapi/anti-enumeration.md`, `restapi/idempotency-and-versioning.md`, and `restapi/openapi-spec-first-drift.md`. Apply `go/decimal-and-money.md` only if its implementation guidance materially applies to these contract changes; do not force implementation-only practices onto OpenAPI.

## Four-pass Review assignment

Run all four canonical passes against the same exact diff, in this order:

1. **Safety:** look for unsafe or misleading credential, anti-enumeration, cache, PII, money, simulation, idempotency, threshold-ordering, or evidence claims. Distinguish API contract shape from unproven runtime controls and empirical parity. Compare any specialized security/concurrency verification expectation with TP-015's Test Focus Pointer and report Techplan drift separately if newly warranted.
2. **Quality:** review clarity and maintainability of descriptions, schemas, response references/examples, shared component wording, and generated output. Report only material issues.
3. **Stack-Specific Best Practices:** use the routed Harscode clue/security map and matching files for only relevant API/idempotency/anti-enumeration/generated-contract concerns. Cite the matching source for findings; state explicitly when a trigger has no finding or no match.
4. **Consistency:** compare actual authored/generated diff against TP-015 and Task 02, current `agreed` Donation specs, Product/MVP, monetary standard, root `AGENTS.md`, API README, and relevant frontend generated-type instructions. In particular check: major-unit decimal string + currency code and exact-decimal semantics; whole-IDR bounds; QRIS-only sandbox and server-owned terminal status; request idempotency separate from settlement replay; D1/D15; fragment handoff with optional/missing header semantics, difficult-to-guess credential, HMAC/24h/status-only; uniform absent/missing/wrong/expired `404` body/headers/cache including `private, no-store`; O11/D19 without invented bounds/mechanics; bounded O8 scope; generated bundle/type correspondence; and no unsupported claims of runtime proof, residual-risk acceptance, or `CONTRACT_READY`.

## Verification posture and output

Review is primarily independent reasoning against the actual four-path diff and current contract. Build reported `cd api && npm run validate` PASS with 124 warnings, `cd api && npm run bundle` PASS, `cd frontend && npm run generate:api-types` PASS, and path-scoped `git diff --check` PASS. These Build checks are context only; do not replay the full matrix. Run a targeted read-only command only to resolve a concrete uncertainty. Do not add/run tests, alter source/spec/generated files, or update projections in Review.

Write `review-findings-1.md` under this Run using the canonical four-pass structure and compact provenance. Write `patch-plan-1.md` only if a code/API/spec correction is required. For each finding include location, problem, impact, suggested correction, and blocking/non-blocking status. Do not promote minor comments into a patch loop.

Verdict must be `Approve`, `Approve with minor comments`, or `Request changes`. If approved, recommend the next phase per current Testing ownership and identify remaining API owner/Human acceptance needed before `CONTRACT_READY`. If changes are needed, identify blocking findings and return only their narrow correction scope to a fresh Build/Patch Run. Do not claim tests/runtime/security evidence, residual-risk acceptance, `CONTRACT_READY`, or a milestone.

## Execution envelope

- `PREAUTHORIZED`: inspect the exact four-path diff/current files, named contract/task/authority/guidance sources; write Review findings and a patch plan only if needed; run a targeted read-only command only for a concrete review uncertainty.
- `ORCHESTRATOR_DECISION`: widen review scope, invoke broad verification, or change route.
- `HUMAN_REQUIRED`: edit API/spec/generated source, choose new Product/API/Security policy, accept residual risk, or promote `CONTRACT_READY`/delivery milestone.

## Human-assisted dispatch

Use a fresh Reviewer Session from the Kencleng repository root. Start from `../harscode-workspace/workflow/4-code-review-prompt.md` and this Invocation. Kickoff:

`Jalankan independent four-pass Code Review Run RV-S2-002-013 sesuai Invocation durable di .harscode-spaces/s2-guest-donation-truthful-state/WU-S2-002/runs/RV-S2-002-013/invocation.md dan canonical Review guidance Harscode saat ini. Review hanya actual BLD-006 diff pada api/openapi/donation.yaml, api/openapi/common.yaml, api/openapi.yaml, dan frontend/lib/api/generated/openapi.ts; baca full changed sources saat ini dan verifikasi bundle/types correspond dengan authored split sources. Jalankan Safety, Quality, Stack-Specific Best Practices, lalu Consistency. Ground pada Approved TP-015, Task 02 TPD-002, current agreed Task 01 specs, Product/MVP, monetary standard, root AGENTS.md, api/README.md, dan frontend/AGENTS.md yang relevan. Route best practices dari Harscode AGENTS/index/security map secara selektif; cite temuan. Build report adalah konteks, bukan pengganti inspeksi diff. Jangan review perubahan spec-status/projection sebelumnya, jangan edit source atau projection, dan jangan jalankan tests/broad matrix; lakukan command read-only hanya untuk ketidakpastian konkret. Tulis review-findings-1.md dengan empat pass/verdict/handoff, patch-plan-1.md hanya bila koreksi diperlukan. Jangan klaim runtime/security evidence, residual-risk acceptance, CONTRACT_READY, atau milestone. Berhenti setelah handoff.`

After dispatch, reconcile from the durable findings and any patch plan. A clean Review does not alone earn `CONTRACT_READY`; preserve contract-time owner acceptance and downstream evidence gates.
