# Build Report — `BLD-S2-002-006`

> Phase: Build  
> Work Unit: `WU-S2-002`  
> Run: `BLD-S2-002-006`  
> Author: Codex Implementer  
> Role / specialization: Implementer / Task 02 authored Donation OpenAPI reconciliation  
> Participant: `P-S2-002-BL-006-1`  
> Session: Fresh Implementer Session; Session ID not exposed  
> Created: 2026-10-01  
> Model / reasoning: `gpt-6-luna` / `high` (Invocation configuration; runtime selection not independently exposed)  
> Target revision: `650e73c5d646c29c0ddf1931618f02685d15f7b7` plus current Task 01 and working-tree artifacts  
> Workflow revision: `33b03a3f62cc3aacba6534b8a011465613c64b09`

## What changed

- `api/openapi/donation.yaml` → reconciled guest submit and status lookup to Slice 2. Request and Donation money fields pair a major-unit decimal string with `currency_code`; IDR submission is whole Rupiah, minimum 5000, with Rp1 increments and no invented maximum. Exact-decimal calculation/persistence and no-float semantics are stated. QRIS is the only accepted sandbox method; the request cannot choose a terminal result. The contract records same-key/same-payload retry, different-payload rejection, separation from settlement replay, D1/D15 full accepted-pending settlement and exact-once funding, and O11/D19 bounded/recoverable terminal-notice eligibility.
- The status operation now uses `X-Donation-Status-Credential`, consistent with the frontend moving the bearer credential from the fragment into the request after URL cleanup. The contract records the one-way HMAC verifier direction, hard 24-hour expiry from issuance, and status-only response. Missing, incorrect, expired, and absent-donation lookups share one `404` Problem Details example and `Cache-Control: private, no-store` contract.
- Guest email is explicitly opt-in and status-only, with the approved 24-hour verification disclosure and terminal simulation label. Verification, retention, retry, deletion, and bounded terminalization mechanics remain owner-defined where TP-015 leaves them open.
- `api/openapi/common.yaml` → clarified idempotency-key retry semantics for its sole current consumer. `api/openapi.yaml` and `frontend/lib/api/generated/openapi.ts` were regenerated from authored sources. No path was added or removed, so `api/openapi/index.yaml` did not need an update. Historical donor-list and Account claim/history operations remain untouched because bounded O8 evidence authorizes replacement of submit/status only.
- No Product/MVP, Design, domain spec, Techplan, runtime, tests, migrations, or Tier-0 source was changed. Pre-existing unrelated working-tree changes were preserved. The current Task 01 source files were not changed in this Run; this Build does not assert residual-risk acceptance, `CONTRACT_READY`, or a delivery milestone.

## Tests run

- `cd api && npm run validate` (executed with `api/` as working directory) → OpenAPI source lint/validation → PASS, no validation errors; Redocly reported 124 warnings across the aggregate. The known repository warning backlog remains; no warning cleanup was in this Task 02 scope.
- `cd api && npm run bundle` → authored split-source bundle generation → PASS; `api/openapi.yaml` regenerated.
- `cd frontend && npm run generate:api-types` → generated type refresh from the bundle → PASS; frontend API types regenerated.
- `git diff --check -- api/openapi/donation.yaml api/openapi/common.yaml api/openapi.yaml frontend/lib/api/generated/openapi.ts` → whitespace/conflict check → PASS.
- An initial `npm run validate` was mistakenly run from repository root and failed because root has no `package.json`; the prescribed validation was then run from `api/` and passed as recorded above.
- No automated/product/runtime tests were added or run; this is contract-only Build scope.

## Verification scope confirmation

No race/concurrency, performance/load, or security-class test was executed in this Build iteration. No broad Testing-owned suite or runtime check was run. Schema validation and generated artifacts do not prove atomicity, exact-once funding, simulator behavior, credential protection, email lifecycle, 404 parity, timing parity, abuse controls, or rendered acceptance.

## Contract check

- [x] Task 02 authored-source target completed for the settled Slice 2 contract details in scope.
- [x] Live Donation OpenAPI/specs and current authorities were reopened; no material assumption required an unapproved Product, Design, Security, or API-owner decision for the encoded details.
- [x] No path change was made; `index.yaml` remains aligned with the authored path set.
- [x] Bundle and generated types were regenerated from authored sources.

## Deferred / not tested here

- O1 additional currencies and concrete range/fraction/precision/scale/storage parameters; only current whole-IDR submission bounds are expressed. Derived-value rounding and tax behavior remain undecided.
- O2 simulator timing and failure-scenario mechanism.
- O3 verification/security controls, retry/retention/deletion mechanics and races, terminalization bound/architecture/timeout meaning, and lifecycle evidence.
- O4 credential-generation strength evidence, key/comparison/lifecycle controls, browser/referrer/log/cache exposure mitigations, abuse controls, and residual-risk decision.
- O5 empirical equality of response body/headers/cache behavior, timing evidence, and abuse controls.
- Runtime D1/D15 ordering, accepted-pending settlement, exact-once funding, threshold overshoot, and stable close reason; independent Testing owns this evidence.
- Independent Code Review of the authored and generated diff, remaining owner gates, and any required rendered Design acceptance.

## Flagged for Techplan / Testing

No new material authority gap surfaced. The exact status-credential header is documented as the API carrier after fragment handoff; its implementation and protection remain open O4 evidence. The `404` contract is expressed in one response with one body example and cache directive; runtime parity remains O5 evidence. Existing out-of-scope donor-list and Account operations were not removed under the bounded O8 authorization.

## Phase handoff

- Completed: Task 02 authored Donation OpenAPI reconciliation and generated-artifact refresh.
- Artifacts: this `report.md` and `launch-record.md` in Run `BLD-S2-002-006`.
- Human decision: none newly required by the authored details in this Run. No residual risk is accepted here.
- Open / deferred: owner controls and evidence O1–O5 as applicable, independent review, runtime Testing, rendered acceptance where applicable, and the `CONTRACT_READY` gate.
- Recommended next step: fresh independent Code Review of the authored `donation.yaml` / `common.yaml` and regenerated bundle/types diff. Do not infer `CONTRACT_READY` or a delivery milestone from validation/generation.
- Session transition: start a new independent Reviewer Participant/Run with a fresh Session context; any later Build/Patch re-entry requires a new Run and fresh Participant Session.
- Context pointers: Approved `TP-S2-002-015/techplan.md`; Task 02 snapshot `TPD-S2-002-002/tasks/02-donation-openapi-reconciliation.md`; changed authored/generated API paths only.
