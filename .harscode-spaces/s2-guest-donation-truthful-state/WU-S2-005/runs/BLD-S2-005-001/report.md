# Build Report — `BLD-S2-005-001`

- **Phase:** Build
- **Work Unit / Run:** `WU-S2-005` / `BLD-S2-005-001`
- **Author:** Implementer, `P-S2-005-BL-001-1` (`KC-IMPLEMENTER`)
- **Created:** 2026-10-01
- **Model / Reasoning:** `gpt-6-luna` / `medium` (dispatch record; runtime identity not independently exposed)
- **Session:** Fresh session per Invocation; session identifier not provided
- **Target revision:** `7fd8b473b239b20bda3990ab29c51440d321a796` plus pre-existing working-tree changes
- **Workflow revision:** `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`

## What changed

- Campaign detail acceptance and `INV-campaign-14` now describe the required `donation_action` union: `available` without a reason, or `unavailable` with only `campaign_not_eligible` when the detail stays public but the backend Donation submission-eligibility predicate fails at the GET snapshot. The text preserves public visibility, anti-enumeration, and independent POST rechecking.
- Authored `api/openapi/campaign.yaml` now defines the two closed object variants and documents GET snapshot semantics. The old `donation_flow_not_available` value was removed.
- Regenerated `api/openapi.yaml` and `frontend/lib/api/generated/openapi.ts` from the authored source.
- Updated the public Campaign fixture with one available and one unavailable action example.

## Tests run

- `cd api && npm run validate` (before edit and after edit) → OpenAPI validation succeeded with zero errors. Baseline and current lint outputs each contain 122 warning anchors; comparison against HEAD found no added or removed warning anchors.
- `cd api && npm run bundle` → generated aggregate successfully.
- `cd frontend && npm run generate:api-types` → generated TypeScript successfully from the new bundle.
- Scoped correspondence inspection → generated type is the two-variant union; fixture examples use only the accepted variants; no old reason remains in the affected contract/fixture surfaces.
- `git diff --check -- <six scoped files>` → passed.

## Verification scope confirmation

No race/concurrency, performance/load, or security-class test was executed in this Build iteration. No broad Testing-owned suite was run. Generation and schema validation are Build-owned contract checks; independent Testing and runtime predicate evidence remain deferred.

## Contract check

- [x] Current Build target satisfied in full for the authorized spec/API/fixture reconciliation and generated counterparts.
- [x] Live-source review did not invalidate the owner-selected wire contract. It confirmed that current Campaign GET production still filters on `status=published` and hardcodes the old unavailable action; this Run makes no producer/runtime claim and did not modify backend code.

## Deferred / not tested here

- The downstream Campaign producer owner must trace the GET-time action value to the same authoritative Donation submission-eligibility predicate used by POST. Current `RepositoryDB.FindPublicDetail` filters `status=published`, and `Service.toPublicDetail` still hardcodes `donation_flow_not_available`; no source/runtime implementation proof exists in this contract-only Run.
- Independent Code Review, final Campaign/API owner acceptance, public projection/anti-enumeration/runtime checks, POST recheck behavior, and broader Slice 2 Testing remain outstanding under their assigned gates.
- No external-consumer compatibility is claimed.

## Flagged for Techplan / Testing

- Downstream producer implementation must not invent an ineligible condition or infer a public reason from internal lifecycle/close labels. If the existing authoritative predicate cannot be evaluated for the GET snapshot while preserving public visibility, return that exact gap to Campaign authority before producer implementation.

## Phase handoff

- **Completed:** Approved TP-S2-005-002 contract-reconciliation target for this Build Run.
- **Artifacts:** this report; authored and generated files listed above.
- **Human decision:** none required during Build; final authored Campaign/API acceptance remains a later Human gate.
- **Open / deferred:** independent Code Review and downstream predicate source-fidelity/runtime proof.
- **Recommended next step:** fresh independent Code Review of the authored/generated/fixture diff.
- **Session transition:** use a new Review Run, Participant, and fresh Session because this Build Run is complete and the next phase is independent.
- **Context pointers:** `TP-S2-005-002/techplan.md`; `docs/spec/4-campaign/features/02-campaign-detail-listing.md`; `docs/spec/4-campaign/invariants.md`; `api/openapi/campaign.yaml`; regenerated `api/openapi.yaml` and `frontend/lib/api/generated/openapi.ts`; `frontend/mocks/fixtures/public-campaign.ts`.
