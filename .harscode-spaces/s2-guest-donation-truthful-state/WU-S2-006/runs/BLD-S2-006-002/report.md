> Phase: Build
> Work Unit / Run: WU-S2-006 / BLD-S2-006-002
> Author: P-S2-006-BLD-002-1 (KC-IMPLEMENTER)
> Created: 2026-10-02
> Model / Reasoning: Invocation configured `gpt-6-luna` / high; active runtime values not independently exposed
> Session: fresh Participant context per Invocation; session identifier not exposed
> Target revision: `7fd8b473b239b20bda3990ab29c51440d321a796` plus working tree at dispatch
> Workflow revision: Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`
> Source acceptance: Pending independent Review and Campaign/Donation owner acceptance; prior acceptance does not cover this amendment

## What changed

- `docs/spec/4-campaign/invariants.md` → extends Campaign field validation with the per-Campaign whole-IDR Donation cap, create default, PATCH omission preservation, draft-only configuration, publication freeze, and existing-row backfill prerequisite. Extends INV-campaign-13 with reservation against finite funding capacity, the distinct `funding_capacity_reached` close reason, stable closure after a pending failure, and downstream concurrency evidence. Extends INV-campaign-14 with required public cap disclosure and an explicit Slice-3 handoff while preserving closed-Campaign `404` behavior in Slice 2.
- `docs/spec/4-campaign/features/01-campaign-creation-draft-crud.md` → records the cap request/default and response behavior, PATCH preservation, validation boundaries, and backfill/freeze evidence. Corrects stale source-path references to the actual `docs/spec/4-campaign/` location.
- `docs/spec/4-campaign/features/02-campaign-detail-listing.md` → includes required `max_donation_amount` in the public detail projection and requires POST to recheck eligibility, current cap, and remaining capacity. Records D-04 as Slice-3-only applicability; current Slice-2 closed detail remains `404`.
- `docs/spec/4-campaign/features/09-closure.md` → separates the overshootable `max_amount` threshold from finite funding capacity, records reservation/closure/failure outcomes and the `funding_capacity_reached` identifier, and hands closed public-result behavior to Slice 3 without activating it here.
- `docs/spec/5-donation/invariants.md` → extends INV-donation-02 and INV-donation-08 for cap enforcement, accepted-pending reservation, finite capacity, exact full settlement, and no reopen after reservation release.
- `docs/spec/5-donation/features/01-submit-donation-settlement.md` → adds guest-facing cap disclosure and generic over-cap validation, capacity admission, distinction from `max_amount`, and failed-pending/no-reopen behavior.

Each changed source has an exact Run-entry byte snapshot under `baseline/`. The Run-only unified diff is `source-delta.patch`; it was generated against those snapshots, so unrelated pre-existing working-tree edits are excluded.

| Source | Run-entry SHA-256 | Current SHA-256 |
|---|---|---|
| `docs/spec/4-campaign/invariants.md` | `bec2ef663229249f3d032ede08b6833e08091e774d7b5b0261ee9f40deb2b9b6` | `d4bd69b41484fc7b921d9f234d787cc01d095a61e4ccea33d4ada31ac0a57888` |
| `docs/spec/4-campaign/features/01-campaign-creation-draft-crud.md` | `0ecf1a1e82212a10e4202fc588f3ca88eabb496bc08ff306f87813764f4c8535` | `53431422b714adebd6ff9f07476c4c930746516f89f98d4c7ea122b7f35c8d05` |
| `docs/spec/4-campaign/features/02-campaign-detail-listing.md` | `55d0ede37e5fe70e07f5e2f2832f3bd03ce94a929e3995538fece7e681f62349` | `2db5dd950553e97756417fd2f41e29356835c43a795c80c3be1d615fc3cd45a7` |
| `docs/spec/4-campaign/features/09-closure.md` | `40df308338f3885a5e4f8531c7743e71af47b967ef2077da5ad53c03cb2272c6` | `19afdf7e5b71f7ffabda5880a8d683f486a39ff0ef63302f96eaa70820795a62` |
| `docs/spec/5-donation/invariants.md` | `436a1421ec12bf0a0618066b16468226221cf9cfddf2bea6cb8d69b26cbb97e8` | `30743644c0d5b9ebe620212442599afc5f9f5bb0323fc2acef97186f7c9902b4` |
| `docs/spec/5-donation/features/01-submit-donation-settlement.md` | `36cf19f24491284346b6a268213dca2204330a4f9d5f348f18b466b20faa842b` | `5050dd29cafc0bd4a906863b0fe32c2ac333117eb97701d6991cbc7a29d28f21` |

The accepted Product/MVP input hashes remain unchanged at `ba2972bc8f91d092e477df170d987b1d124964d9cc36c025d2a8da3ed12709af` (`mvp-scope.md`) and `4c69a030e7fedc9c62bf30f85c00e81f9806c45b2ed5471c1d5126762be8091f` (`mvp-delivery-slices.md`). The Approved Techplan and Implementer profile hashes also match the Invocation.

Run-only `source-delta.patch` SHA-256: `d0bc1537eb5336a66d23128bc63fd8259bd5607dd863ed2eae76871719439560`.

## Tests run

- No tests run, as explicitly directed by the Run Invocation.
- Source traceability: reopened the accepted Product/MVP sources, Approved Techplan, current Campaign/Donation source anchors, monetary standard, and owner receipt. The accepted Product hashes and all six Run-entry spec hashes match the Invocation coordinates.
- Run-only delta capture: compared all six current spec files to their Run-entry snapshots and wrote `source-delta.patch`; confirmed that only the six assigned sources differ from their snapshots.
- `git diff --check -- docs/spec/4-campaign/invariants.md docs/spec/4-campaign/features/01-campaign-creation-draft-crud.md docs/spec/4-campaign/features/02-campaign-detail-listing.md docs/spec/4-campaign/features/09-closure.md docs/spec/5-donation/invariants.md docs/spec/5-donation/features/01-submit-donation-settlement.md` → passed with no whitespace errors.
- Manual source traceability checked cap range/default/configuration freeze, PATCH omission preservation, required detail disclosure, generic over-cap `422` direction, settled-plus-pending capacity, close-reason distinction/stability, full accepted-pending settlement, and Slice-3-only public closure behavior against the accepted Product direction and Techplan Q1–Q6.

## Verification scope confirmation

No race/concurrency, performance/load, security-class, service, browser, runtime, API validation/generation, migration, or database check was executed. This documentation-only checkpoint does not establish runtime atomicity, storage backfill, API correspondence, or consumer compatibility. No broad Testing-owned suite was run.

## Contract check

- [ ] Current Build target satisfied in full — this checkpoint completes the assigned Campaign/Donation specification stage only; authored API and counterpart stages remain.
- [x] Live-source re-grounding did not invalidate a material Techplan assumption.

## Deferred / not tested here

- Independent Review and explicit Campaign/Donation owner acceptance of the six concrete spec amendments. The prior Product acceptance and prior Donation acceptance cover their named earlier source bytes only.
- Authored `api/openapi/campaign.yaml` and `api/openapi/donation.yaml` reconciliation, including exact wire schema for `max_donation_amount`, generic Problem Details reuse, and the transport behavior when capacity prevents admission. API source authoring is the next dependent stage after this source gate.
- API aggregate/generated/type/fixture/known-consumer correspondence after authored API acceptance; no counterpart or frontend/backend production source was changed.
- Runtime, database, exact-value, compatibility, concurrency, settlement/replay, and rendered evidence owned by later delivery/Testing. No migration or DB application is authorized in this Run.

## Flagged for Techplan / Testing

No material assumption break. The finite-capacity runtime ordering, reservation lifecycle, close-reason stability, and accepted-pending settlement remain explicit Testing obligations. API wire encoding must follow the shared monetary representation standard and the accepted spec; this Run does not choose a new precision/scale policy or disclose capacity state.

## Phase handoff

- Completed: concrete Campaign/Donation specification amendments are reviewable with exact Run-entry snapshots, Run-only delta, and before/after hashes.
- Artifacts: this report at `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/BLD-S2-006-002/report.md`; snapshots under the same Run's `baseline/`; exact delta at `source-delta.patch`.
- Human decision: after independent Review, accept or revise these exact six Campaign/Donation spec source amendments. No source acceptance is recorded by this Build.
- Open / deferred: authored API source reconciliation and then API/generated/fixture/consumer counterparts remain under the still-pending whole Techplan target; backend/frontend delivery and runtime evidence remain downstream.
- Recommended next step: Orchestrator routes independent Review of this exact Run-only six-file delta and the Campaign/Donation owner acceptance gate. After acceptance, continue the approved target at the authored split API stage. Do not auto-dispatch or claim WU-S2-006 completion.
- Session transition: stop this Build occurrence at the owning-spec gate. Any later phase re-entry uses its own Run and fresh Participant Session/context; this Build target is not complete.
- Context pointers: approved `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-004/techplan.md`; current Run Invocation; the six spec sources and `source-delta.patch`; accepted Product source hashes recorded above.
