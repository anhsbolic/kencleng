# Campaign/Donation Ordering Brief — `OIR-S2-002-006`

## Provenance

- Work Unit / Run: `WU-S2-002` / `OIR-S2-002-006`
- Phase / Role: Exploration / Explorer — Campaign/Donation threshold and settlement ordering facilitation
- Participant: `P-S2-002-OIR-006-1`
- Created: 2026-09-29
- Model / reasoning: Invocation selects `gpt-6-luna` / `high`; runtime model was not independently exposed.
- Session: not exposed
- Target revision observed: `934b093cea30230743dee951ae1e600762019f30`
- Workflow revision: `7a4dbf2c065bd8fd02c86c24073d7309046bff30`
- Source evidence: `evidence/stage-2-gap-analysis.md`
- Owner decision received: Anhar Solehudin, 2026-09-29, in this Run's Stage 3 gate.

## Outcome

**Ordering direction decided for current Slice 2.** Anhar explicitly adopted the recommended Campaign/Donation contract ordering below as a delivery/contract decision. It preserves the settled Product/MVP O6 policy and does not change Product/MVP authority.

## Owner decision D1 — Eligibility, accepted settlement, and close ordering

1. **Submission eligibility is atomic against Campaign close.** If submission wins while the Campaign is `published`, the Donation is accepted. If close wins first, the new submission is rejected.
2. **Accepted Donations remain settleable.** A Donation accepted while eligible may settle at its full amount even if the Campaign later becomes `closed`.
3. **Success and funding form one exact-once atomic outcome.** Successful settlement and the full funding increment commit together, exactly once.
4. **Post-close settlement does not reverse or rewrite closure.** It does not reopen the Campaign or change the winning close reason. Funding may continue to increase past `max_amount` as accepted pending Donations settle.

Decision scope is Campaign/Donation contract ordering within current Slice 2. The decision does not select a database isolation/locking mechanism, add a new endpoint, or authorize broader Slice 3 closure/public-result behavior.

## Evidence and reconciliation consequence

Product/MVP and O6 already require full settlement for Donations accepted while Campaign is eligible, rejection of new post-close submissions, and permit total funding to exceed the threshold further (`docs/product/mvp-scope.md` Stage D; `docs/product/mvp-delivery-slices.md` §5; `OIR-S2-002-001/resolution-brief.md` O6). The Approved Techplan leaves only delivery ordering to the Campaign/Donation owners (`TP-S2-002-007/techplan.md` §13, item 8). The Authority Map names Anhar as the current Slice-2 Campaign and Donation delivery/domain owner (`.harscode-spaces/authority-map.md` lines 9–10).

The older Donation invariant/feature sequence conditions the funding update on `Campaign.status = 'published'` and does not specify the zero-row path after an accepted Donation becomes pending while the Campaign closes (`docs/spec/5-donation/invariants.md` `INV-donation-08`; `docs/spec/5-donation/features/01-submit-donation-settlement.md` Settlement). That detail must be adapted or replaced in an authorized contract-reconciliation route to express D1. The current authored split OpenAPI also does not yet express accepted-pending settlement and close ordering (`api/openapi/donation.yaml` submit operation; `api/openapi/campaign.yaml` force-close operation).

The delivery consequence is that a close transition and a Donation acceptance need a single observable order; after acceptance, a later close cannot invalidate the Donation's full settlement/funding outcome. Concurrent accepted settlements must all contribute once, while only the first valid close transition establishes Campaign closure/reason. A later increment may increase funding without changing that closure.

## Alternatives considered

Only the adopted direction preserves the approved Product/MVP policy. Treating a prior close as grounds to reject/omit funding for an already accepted Donation would contradict D1 and O6, and could break atomic success/funding coupling. Clamping or partially accepting amounts would also contradict the settled threshold policy. No alternate Product semantics were presented as viable.

## Findings, status, and remaining concerns

- **F1 — ordering contract decision:** resolved at the owner-decision level by D1. Historical invariant/feature and OpenAPI wording remains unreconciled, so this does not establish that authored contract sources are ready or that `CONTRACT_READY` has been reached.
- **F2 — absent live write path:** deferred to later Build/Testing. No Donation/settlement/threshold-close implementation was present at the inspected revision; D1 has not been proven at runtime.
- **O1 shared currency standard:** remains open and separate, as recorded by `OIR-S2-002-005/amount-contract-brief.md`.
- **O2/O3 email-retention conflict:** remains open and separate; not investigated or resolved by this Run.
- Full Slice 3 closure/public-result behavior, broader product policy, money rounding/tax, and implementation locking choice remain outside this decision.

## Verification

- **Performed:** targeted evidence review and pinned-input hash verification, recorded in `evidence/stage-2-gap-analysis.md`; explicit owner direction captured above.
- **Not performed:** code or spec/API changes, tests/race tests, API lint/bundle validation, runtime checks, database/migration checks, or implementation proof.
- No `CONTRACT_READY`, implementation correctness, or Testing milestone is claimed.
