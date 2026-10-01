# Build/Patch Report — `BLD-S2-002-007`

> Phase: Build/Patch  
> Work Unit: `WU-S2-002`  
> Run: `BLD-S2-002-007`  
> Role / specialization: Implementer / narrow patch for Review finding F-1  
> Participant: `P-S2-002-BL-007-1`  
> Author: Codex Implementer  
> Created: 2026-10-01  
> Model / reasoning: `gpt-6-luna` / `high` (Invocation configuration; runtime selection not independently exposed)  
> Session: Fresh Implementer Session; Session ID not exposed  
> Target revision: `055aa1283ff5f1ab9c08c1422981a9cbccf98e6c` plus current working-tree changes, including BLD-006's authored/generated API diff  
> Workflow revision: `33b03a3f62cc3aacba6534b8a011465613c64b09`

## What changed

- `api/openapi/donation.yaml` → added `required: [amount, currency_code]` to `DonationListItem`, `ClaimableDonation`, and `MyDonation`, resolving RV-013 F-1 while preserving all field schemas, descriptions, and other response behavior.
- `api/openapi.yaml` and `frontend/lib/api/generated/openapi.ts` → regenerated from the authored split source. The generated bundle retains the existing BLD-006 contract changes and now marks both monetary fields required in all three projection types. No unrelated authored or generated source was changed by this patch.

## Tests run

- `cd api && npm run validate` (run from `api/`) → authored OpenAPI validation → PASS, zero errors and 124 warnings across the aggregate, consistent with the documented historical warning backlog.
- `cd api && npm run bundle` → generated OpenAPI bundle → PASS.
- `cd frontend && npm run generate:api-types` → generated frontend API types → PASS; verified `amount: string` and `currency_code: string` are required for all three corresponding projection types.
- `git diff --check -- api/openapi/donation.yaml api/openapi.yaml frontend/lib/api/generated/openapi.ts` → scoped whitespace/conflict check → PASS.

## Verification scope confirmation

No race/concurrency, performance/load, or security-class test was executed in this Build/Patch iteration. No broad Testing-owned suite or runtime check was run. Schema validation and generated types do not prove runtime behavior, settlement atomicity, exact-once funding, security/PII controls, or empirical status-response parity.

## Contract check

- [x] RV-013 F-1 is addressed in the three specified authored response schemas.
- [x] Generated bundle and TypeScript make both monetary fields required for each corresponding projection type.
- [x] Live schemas and generated counterparts were reopened before editing; generated outputs were regenerated and inspected.
- [x] Patch scope is limited to the three required-field declarations and their generated consequences; existing BLD-006 working-tree changes were preserved.

## Deferred / not tested here

- Independent targeted Review confirmation of F-1 and generated correspondence.
- Runtime correctness, concurrency, performance/load, security/PII controls, empirical response parity, rendered acceptance, residual-risk acceptance, and other TP-015 gates remain outside this contract-only patch.

## Flagged for Techplan / Testing

None. No material authority gap or scope change was found.

## Phase handoff

- Completed: F-1's three monetary response projections now require `amount` and `currency_code`; generated outputs are refreshed and scoped checks passed.
- Artifacts: this report and `launch-record.md` under Run `BLD-S2-002-007`.
- Human decision: none required for this narrow correction.
- Open / deferred: independent targeted Review confirmation and the downstream evidence/owner gates listed above.
- Recommended next step: fresh targeted Code Review confirmation of F-1 and generated correspondence; do not restart a full review loop automatically.
- Session transition: Orchestration should dispatch a new targeted Review Run with a new Participant and fresh Session because this Build/Patch re-entry has completed.
- Context pointers: TP-015; Task 02; RV-013 `review-findings-1.md` and `patch-plan-1.md`; `api/openapi/donation.yaml`; regenerated `api/openapi.yaml` and `frontend/lib/api/generated/openapi.ts`.
