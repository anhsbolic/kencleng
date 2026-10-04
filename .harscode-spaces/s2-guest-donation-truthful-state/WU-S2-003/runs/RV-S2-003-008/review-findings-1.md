> Phase: Independent migration-design Review
> Author: P-S2-003-RV-008-1 (Reviewer)
> Created: 2026-10-04
> Model: Invocation configured `gpt-6-luna` (active runtime model not independently exposed)
> Reasoning: Invocation configured `high` (active runtime effort not independently exposed)
> Session: Fresh Reviewer Session; identifier not exposed
> Work Unit / Run: `WU-S2-003` / `RV-S2-003-008`
> Target revision: Kencleng HEAD `4e71d3697a479d92b2cad5a79ef5d710f4e8077c` plus current working tree; Approved Techplan candidate SHA-256 `55eb0c94bb922d3311279c462122cb57f9f2181226f33092fff9d68bbd2c41b5`
> Workflow revision: Current-effective Harscode; canonical Review prompt SHA-256 `d152c1419ea8b0e67a08d1ec2cec7c0235cc15b0fccbd032722aea1ed2332f8a`; Review guidelines `513f145a5bc1980589a4967f6a0894a3f59f089ade4fa199ba252b2604a735c0`; checklist `5621854c75f8834ff9c32a8a67a326ab03388e11cceac95f0f2c6bdb0e03e079`; overlay `f806673aa5e7d1d82dae9f433ec54c3562e15ce778b709a141c57d3d22dfaf87`

This is a fresh pre-implementation review of the minimum Donation/D1 schema and migration design only: Approved Techplan §§9 steps 2a–2b and 8, §10 migration/D1 rows, §12 R14/R18/R19/R20, and §13 Item 6. It assesses RV-S2-003-005 F-02/F-03/F-04/F-06 independently. The Approved Techplan hash and invocation-pinned source hashes were verified before review. Prior design findings SHA-256: `fca27d88610017db317f679a68c97af7b49a616d3b7b69d6fa06b45fef359c7f`.

## 1. Safety

No additional finding in this pass. The reviewed design specifies that unavailable settled Funding admits no Donation, accepted-pending amounts participate in reservation, the Campaign-first ordering is preserved, Campaign deletion is restricted once Donation history exists, and post-write down must not erase Donation or close history. Actual transaction behavior, lock impact, and database enforcement remain unverified and are not inferred from the written design.

## 2. Quality

### F-01 — Same-key retry does not define how to return the required status credential

- **Location:** Approved Techplan §8 `Persistence/data shape`, §9 step 2a and R4; accepted `api/openapi/donation.yaml` `Donation` schema and POST `201` response; Donation Feature 01 retry acceptance and Feature 02 status-credential acceptance.
- **Problem:** The design now settles endpoint-wide key uniqueness and complete-request equivalence using a keyed HMAC fingerprint, with no raw request JSON and fingerprint deletion alongside its record. It still does not define the retry response's `status_token`. The accepted `Donation` response requires this field, says it is returned only in the submission response, and the retry rule returns the original Donation. The plan also says never to persist the bearer credential. With only a one-way verifier, an ambiguous retry cannot reproduce a lost token; after hard expiry, returning the original token would not provide usable status access. The design does not say whether retry may issue/rotate a credential or how that can fit the accepted contract and O4 controls.
- **Why it matters:** Same-key retry is the recovery path after an ambiguous outcome. A design that cannot return the contract-required response, or silently returns an expired/unusable credential, can leave a donor unable to access the accepted status flow. Persisting the bearer to solve replay would contradict the plan's credential boundary.
- **Suggested resolution:** Reconcile retry response/credential behavior with the owning API and Security/PII authorities within existing accepted semantics, and make the resulting persistence/replay requirement explicit in the schema design. Preserve the no-bearer-at-rest rule; do not choose credential generation, storage, rotation, or lifetime controls here. O3/O4 remain the owners of their separate open controls.
- **Blocking:** Yes — the idempotency representation is not complete enough to establish the minimum Donation schema's accepted replay behavior.

The other F-02 elements are addressed at design level: R4 names endpoint-wide uniqueness, all accepted canonical intent fields after defaults (including Campaign ID and optional PII), keyed HMAC comparison, separation from settlement replay, and deletion with the bounded record. O3's numeric lifetime remains properly deferred.

## 3. Stack-Specific Best Practices

**RV-S2-003-005 F-03: resolved at design level.** §§9.2b and 10 identify an additive successor after `000012`, leave existing Funding intact, preserve unknown legacy close reasons as NULL, require row/constraint inventory and staged validation where needed, call for PostgreSQL-version/table-size/lock-impact assessment, and distinguish empty pre-write cleanup from a post-write down guard against history loss. This follows the applicable direction in `best-practices/postgresql/migrations-safety.md`. The actual migration, lock behavior, and up/down execution remain deferred to authorized implementation and Testing/Human evidence; this Review does not claim reversibility or PostgreSQL execution evidence.

**RV-S2-003-005 F-04: resolved at design level.** §8 gives the pending aggregate predicate `campaign_id = ? AND status = 'pending'` and a matching partial Campaign index; it also states endpoint-wide unique-key lookup, constrained states and amount, immutable Campaign association, conditional pending-to-terminal writes, and the Campaign-first order. R18 assigns real-Postgres predicate/index and integrity evidence. This is consistent with the targeted guidance in `best-practices/postgresql/indexing-and-query-plans.md`, `transactions-and-locking.md`, and `financial-invariant-enforcement.md`. Query plans, concurrency, and constraints have not been run or independently proven here.

No further Stack-Specific finding. PostgreSQL best-practice clue rows for migrations, indexes, transactions/locking, financial invariants, and HMAC/PII were routed; the applicable migration, index, locking, financial-invariant, and encryption-at-rest guidance was inspected. R4 follows the accepted sensitive-derived-data direction and does not add raw request persistence. No database, SQL, migration, or test command was run.

## 4. Consistency

**RV-S2-003-005 F-06: resolved at design level.** §§8–10 specify a restrictive Donation-to-Campaign FK with no cascade, intentional rejection of `SeedCampaign(replace)` after Donation history exists, nullable `closed_reason` for legacy closed rows whose reason is unknown, and an atomic first winning reason for each new close. R20 checks preservation and no speculative backfill. This matches Campaign INV-campaign-13, Donation INV-donation-08, the current `000011`/`000012` schema anchors, and backend migration conventions. No legacy history is invented. Runtime behavior remains untested.

The resolved Donation amount direction (exact integer Rupiah, zero fractional digits, Rp5,000–Rp1,000,000,000) is consistent with Donation INV-donation-01 and the project monetary standard; it does not claim a project-wide precision or scale. No new monetary inconsistency was found in this target.

## Verification executed during Review

Read-only provenance verification: SHA-256 matched the exact Approved Techplan (`55eb0c94…c41b5`), TP-S2-003-011 report and launch record, prior RV-S2-003-005 findings, all invocation-listed domain/API/migration anchors, participant profile and model registry, and current Harscode Review prompt/guidelines/checklist/overlay plus routed PostgreSQL guidance. `git rev-parse HEAD` matched pinned HEAD `4e71d3697a479d92b2cad5a79ef5d710f4e8077c`. Inspected the pinned Techplan sections, F-02/F-03/F-04/F-06 report, accepted authorities, API response shape, and `000011`/`000012` anchors.

No tests, validators, generators, SQL, migration, database, or runtime action was run. **Verified:** pin/hash identity and written-source alignment reported above. **Assumed:** future migration authoring will implement the stated predicates, constraints, restrictive FK, and down guard without broadening scope. **Deferred:** actual row inventory, supported PostgreSQL lock assessment, constraint/index execution evidence, O3/O4 owner controls, and migration application. **Not tested:** PostgreSQL behavior, query plans, concurrency, up/down behavior, and runtime response/credential replay.

## Verdict

**Request changes.** Blocking finding: F-01 (the remaining F-02 replay-response gap). The revised design answers the identity/equivalence, migration safety, pending aggregate/integrity, Campaign deletion, and close-reason concerns from RV-S2-003-005 at the written-design level, but the status credential required in the same-key POST response is not reconciled with returning the original Donation and never persisting a bearer. This verdict is limited to the positive pre-implementation Donation/D1 migration-design gate. It does not reopen settled product/API decisions, select O3/O4/O5 controls, authorize protected implementation or migration application, or claim PostgreSQL evidence.

## Phase handoff

- **Outcome:** `COMPLETED` — four independent Review passes and verdict recorded for this Run.
- **Result refs:** `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-008/review-findings-1.md`; Approved Techplan candidate SHA-256 `55eb0c94bb922d3311279c462122cb57f9f2181226f33092fff9d68bbd2c41b5`.
- **Findings:** F-01 blocks the minimum Donation/D1 migration-design gate; details and affected contract/design locations are above. RV-S2-003-005 F-03/F-04/F-06 are resolved at design level.
- **Decision requests:** Resolve the same-key retry credential response behavior through the owning API and Security/PII authorities without choosing O3/O4 controls in this Review.
- **Blockers:** Do not treat the minimum Donation/D1 schema design as positively reviewed until F-01 is resolved and receives the required independent review.
- **Open / unverified:** PostgreSQL row/lock/index/constraint/up-down evidence; O3/O4/O5 controls and Human gates; migration application, implementation, and all Testing evidence.
- **Recommended continuation:** Route F-01 for owning-authority reconciliation in the planning/design workflow, then obtain a fresh independent migration-design Review of the exact resulting Approved design. No downstream Run is dispatched by this Participant.
- **Context refs:** Approved Techplan §§8–10, 12 R14/R18/R19/R20, 13 Item 6; RV-S2-003-005 findings; Donation INV-donation-01/02/08/10 and Feature 01/02; `api/openapi/donation.yaml`; PostgreSQL guidance cited above.
