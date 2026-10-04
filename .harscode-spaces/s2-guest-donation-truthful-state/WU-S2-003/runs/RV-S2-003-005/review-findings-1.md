> Phase: Independent Code Review — migration design review
> Author: P-S2-003-RV-005-1 (Reviewer)
> Created: 2026-10-04
> Model: Invocation configured `gpt-6-luna` (runtime model not independently exposed)
> Reasoning: Invocation configured `high` (runtime effort not independently exposed)
> Session: Fresh Reviewer Session; session identifier not exposed
> Work Unit / Run: `WU-S2-003` / `RV-S2-003-005`
> Target revision: Kencleng HEAD `19d53315ac2847a03405339c0e86df8a97850761` plus current durable working-tree state; proposal SHA-256 verified as `dc8f5475b567f692760bb61e7287af7d93dfb4c4a77f7ea8ae9cfd2577a8a4af`; Techplan SHA-256 verified as `b71951c86b2f74f6c9f1a164bcf43447c58ee655874deaf09275e6ff3b2316f9`
> Workflow revision: `pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4`

Review scope is the pinned section “Usulan konkret untuk migration design review dan pairing berikutnya” in `runs/BLD-S2-003-002/report.md`. This is a pre-implementation schema/D1 design review, not review of a code diff. I verified both pinned hashes before reading the proposal. The prior eight-file cap Review/Testing scope does not include this Donation/D1 schema.

## 1. Safety

### F-01 — Funding-unavailable behavior leaves admission without a safe capacity basis

- **Location:** Proposal, “Money compatibility” and “Reservation”; `backend/migrations/000011_create_public_campaigns.up.sql` (`campaigns_funding_pair_check`); TP-S2-003-006 §§8–10, R5/R11.
- **Problem:** The current Campaign schema permits `target_amount` and `collected_amount` to be NULL together. The proposal says to review treatment of unavailable/invalid existing Funding without inventing a balance or destructive backfill, but does not define how D1 admission computes settled Funding, or whether admission is possible, when that pair is NULL. Treating NULL as zero would invent capacity; admitting without a value would make the finite-ceiling predicate unverifiable.
- **Why it matters:** D1 must serialize admission against close and never admit settled Funding plus accepted-pending amounts beyond IDR `99,999,999,999,999,999` (Donation INV-donation-02; Campaign INV-campaign-13). A schema write made before this case is resolved cannot support a proven safe predicate for every existing Campaign state.
- **Suggested resolution:** Route the unavailable-Funding case to the owning Planner/Human for reconciliation with the accepted Funding-availability contract. State the permitted admission/repair behavior and its authority before schema Build; do not backfill zero or infer a balance in this Review.
- **Blocking:** Yes — blocks minimum Donation/D1 schema Build pending authority resolution.

## 2. Quality

### F-02 — Idempotency identity and payload-equivalence shape are not design-complete

- **Location:** Proposal, “Donation persistence”; TP-S2-003-006 Q5/R4, D4; `api/openapi/donation.yaml` POST `Idempotency-Key` contract; Donation INV-donation-10.
- **Problem:** “Request-idempotency identity and canonical-payload-equivalence data” does not state the uniqueness scope for the key, the fields included in canonical equivalence, or the immutable/terminal replay data needed to return the original Donation after Campaign close. Accepted payload includes amount, currency, payment method and optional guest fields. Email handling and retention controls remain O3; a raw or weakly hashed payload copy could expose sensitive input. INV-donation-10 also explicitly leaves retry-record lifetime/serialization as delivery detail.
- **Why it matters:** A uniqueness constraint with the wrong scope can either collide distinct requests or admit duplicate intents. An equivalence representation that omits an accepted field can incorrectly replay a changed payload; persisting sensitive fields outside the established protection pattern can expose PII.
- **Suggested resolution:** Have the Planner reconcile the minimum identity scope and the exact settled-field equivalence semantics against Q5 and the accepted request contract. Route any representation involving guest email or retention to the existing Security/PII/O3 owner. The schema proposal should then name the required uniqueness/equivalence properties without selecting O3 policy.
- **Blocking:** Yes — the key constraint and persisted retry representation cannot be safely finalized yet.

## 3. Stack-Specific Best Practices

### F-03 — Migration rollback and operational impact are not assessable from this proposal

- **Location:** Proposal, “Migration sequence” and “Reversibility”; proposed successor after `000012`; Harscode `best-practices/postgresql/migrations-safety.md`.
- **Problem:** The proposal gives no actual additive migration shape, existing-row constraint strategy, or down-migration data-loss posture. A down migration that drops the new Donation table necessarily deletes accepted Donation/retry/terminal history after writes; dropping the new Campaign columns also removes winning close reasons and threshold settings. The proposal correctly says no migration is to run, but does not distinguish a pre-write rollback from a post-write rollback or state how locking/constraint validation risk will be assessed for existing tables.
- **Why it matters:** “Additive successor” alone does not establish compatibility or genuine reversibility. The routed best practice calls for additive-first treatment, checked locking impact, and tested two-way reversibility; this review cannot claim any of those execution properties.
- **Suggested resolution:** Before writing SQL, document existing-row compatibility and constraint validation posture, intended lock-impact assessment, and explicit down behavior before versus after data exists. If post-write down is intentionally destructive, route that consequence through the project’s migration-application/Human gate and rely on forward correction for deployed data rather than describing it as reversible. Keep execution and DB evidence for Testing/Human ownership.
- **Blocking:** Yes — migration design review prerequisite is not met for a schema write.

### F-04 — Pending-reservation access path and integrity enforcement are only deferred as questions

- **Location:** Proposal, “Donation persistence” and “Reservation”; Harscode `best-practices/postgresql/indexing-and-query-plans.md`, `transactions-and-locking.md`, and `financial-invariant-enforcement.md`.
- **Problem:** The proposal names the need to review pending-by-Campaign access, terminal-state constraints, uniqueness, FKs and indexes, but supplies no query shapes or resulting integrity requirements. At minimum the design must support summing all accepted-pending Donation amounts for one Campaign under the Campaign-first lock, distinguish pending from terminal contribution, and prevent duplicate request identity. An index cannot be judged without the actual lookup predicates; an application-only amount/status assumption is not a DB integrity design.
- **Why it matters:** Campaign-row serialization is the approved D9 mechanism, but it only makes the aggregate coherent if every relevant writer follows that lock order and the data model can query/enforce the intended set. The PostgreSQL guidance requires consistent lock order, transaction-level financial invariants, and indexes matched to actual filter order.
- **Suggested resolution:** Carry a concise schema invariant/access-pattern table into the Build design: the pending aggregate predicate, key-lookup predicate, terminal transition constraints, and FK deletion policy. Preserve Campaign → Donation lock order from D9; do not add a counter unless the Planner proves an aggregate cannot satisfy the accepted semantics. Determine concrete indexes from those predicates and assess them against the relevant table/write volume before SQL is written.
- **Blocking:** Yes — the target is not yet specific enough to review constraints/indexes as part of the required schema gate.

## 4. Consistency

### F-05 — Monetary storage authority is internally unresolved for this feature

- **Location:** Proposal, “Money compatibility”; TP-S2-003-006 D3 and Q2/R2/R5; `docs/spec/5-donation/invariants.md` INV-donation-01/O1 and Donation feature 01 O1; `docs/project/kencleng-monetary-data-standard.md`.
- **Problem:** The proposal commits only to exact whole-IDR input and says the Donation storage shape must reject fractional/out-of-range values. The approved project standard deliberately leaves concrete precision, range and scale to feature owners. Donation INV-donation-01 and the active feature retain O1 for concrete range/precision/storage parameters, while TP D3 says not to reuse `NUMERIC(19,2)` and describes feature bounds as settled. No approved source specifies the Donation column type/scale and DB-level range/integer enforcement that implement both amount and cumulative-cap bounds.
- **Why it matters:** Choosing a type or check here would either invent an unresolved storage parameter or silently treat D3 as having resolved O1. Without a concrete exact-decimal bound, database constraints cannot show that fractional or out-of-range values are rejected while remaining compatible with the Campaign funding ceiling.
- **Suggested resolution:** Planner/Human reconcile D3 against the still-open O1 in the owning invariant/feature authority, then make the concrete range/type decision there. Update the execution contract through its owning route before Build; do not choose a numeric type in this Review.
- **Blocking:** Yes — the persistence type and monetary checks are part of the minimum schema.

### F-06 — Campaign deletion and close-reason compatibility need an explicit preservation rule

- **Location:** Proposal, “Campaign lifecycle” and “Donation persistence”; `backend/migrations/000011_create_public_campaigns.up.sql`; `backend/internal/domain/campaign/repository_db.go` `SeedCampaign`; Campaign INV-campaign-13.
- **Problem:** The proposal does not say how Donation foreign keys behave if a Campaign is deleted/replaced, nor how the current seed replacement path that deletes a Campaign is expected to behave once Donation history exists. Campaign-owned `closed_reason` must also preserve the first winning reason, while existing Campaign rows may already have `status = 'closed'` with no known reason; “do not backfill an unknown reason” leaves compatibility treatment unstated.
- **Why it matters:** Cascading deletion could erase money-related Donation history. A restrictive FK changes the existing `SeedCampaign(replace)` behavior after donations exist. Conversely, fabricating a close reason for legacy closed rows violates the proposal’s own no-invented-history rule. New close transitions need a persisted reason without requiring false historical values.
- **Suggested resolution:** State a no-cascade preservation policy for Donation-to-Campaign history and its expected seed-replacement behavior; ensure that remains consistent with MVP deletion scope. Define how constraints permit legacy unknown reasons while requiring the reason for newly closed Campaigns, or route an explicit legacy reconciliation to the authority owner. Do not add a speculative history backfill.
- **Blocking:** Yes — FK/deletion and close-reason compatibility are required migration design decisions.

## Verification executed during Review

None. The two pinned artifact hashes were checked with `sha256sum`. Review evidence is read-only inspection of the proposal, approved Techplan, current migration/domain anchors, routed PostgreSQL guidance, and accepted Campaign/Donation authorities. No test, SQL, migration, database, or runtime command was run; no PostgreSQL execution, locking, query-plan, or up/down reversibility claim is made.

## Verdict

**Request changes.** Blocking findings: F-01 through F-06. The proposal is not approved as a basis to write the minimum Donation/D1 schema. This verdict is limited to the migration-design-review prerequisite; it does not reopen accepted D1 ordering or other settled product/API behavior, approve Tier-0 implementation, or authorize migration application. Resolve authority-owned gaps through Planner/Human routes and return with a reconciled design target for a fresh review occurrence.

## Phase handoff

- **Outcome:** `COMPLETED` — four-pass independent review completed against the pinned design proposal.
- **Result refs:** `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-005/review-findings-1.md`; pinned proposal `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/BLD-S2-003-002/report.md` SHA-256 `dc8f5475b567f692760bb61e7287af7d93dfb4c4a77f7ea8ae9cfd2577a8a4af`.
- **Findings:** F-01–F-06 are blocking for the minimum schema design gate; see this report for evidence and owning resolution routes.
- **Decision requests:** Planner/Human reconciliation of Funding-unavailable admission, Donation storage precision/range, and legacy close-reason treatment; Security/PII owner input where idempotency equivalence touches guest email or retention.
- **Blockers:** Donation/D1 schema Build is blocked on the unresolved design/authority items in F-01–F-06. Read-only work outside that schema scope is not assessed here.
- **Open / unverified:** No implementation, database, migration, lock, query-plan, or reversibility evidence; migration application and all D1 PostgreSQL/concurrency Testing remain downstream gates.
- **Recommended continuation:** Orchestrator route the cited authority gaps to their owners, reconcile the design proposal, then dispatch a new independent migration-design Review Run before schema writes. This recommendation does not create or dispatch a Run.
- **Context refs:** `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/RV-S2-003-005/invocation.md`; `runs/TP-S2-003-006/techplan.md` §§8–10, §§12–13; root/backend `AGENTS.md`; `000011`/`000012`; Campaign INV-campaign-13; Donation INV-donation-01/02/08/10; PostgreSQL migration/locking/index/financial-invariant guidance cited above.
