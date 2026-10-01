# Patch Plan — `RV-S2-002-013`

> Phase: Code Review patch plan  
> Work Unit: `WU-S2-002`  
> Run: `RV-S2-002-013`  
> Author: Codex Reviewer  
> Created: 2026-10-01  
> Model / reasoning: `gpt-6-luna` / `high` (Invocation configuration; runtime selection not independently exposed)  
> Target revision: `650e73c5d646c29c0ddf1931618f02685d15f7b7` plus BLD-006's four-path diff  
> Workflow revision: `33b03a3f62cc3aacba6534b8a011465613c64b09`

## Scope

Resolve blocking finding F-1 from `review-findings-1.md` only.

## Required correction

In `api/openapi/donation.yaml`, require both `amount` and `currency_code` in the `DonationListItem`, `ClaimableDonation`, and `MyDonation` response schemas. This makes the monetary pair mandatory wherever those projections are returned and aligns the generated consumer types with the approved project-wide representation.

Regenerate `api/openapi.yaml` and `frontend/lib/api/generated/openapi.ts` using the documented `api/README.md` workflow. Do not hand-edit generated files. Inspect the resulting four-path diff and keep it scoped to this finding.

## Acceptance evidence for the patch

- The three authored response schemas require both monetary fields.
- Generated TypeScript makes both fields required for the three corresponding projection types.
- The bundle and generated types match the authored source; no unrelated diff is introduced.
- Run only the normal authored-source validation and required generation commands owned by Build; downstream broad/runtime Testing remains separate.
