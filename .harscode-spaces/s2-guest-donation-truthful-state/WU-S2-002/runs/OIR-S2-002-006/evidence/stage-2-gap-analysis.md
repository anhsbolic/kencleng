# Stage 2 — Campaign/Donation Ordering Evidence

## Provenance

- Work Unit / Run: `WU-S2-002` / `OIR-S2-002-006`
- Phase / Stage / Role: Exploration / Stage 2 — Gap Analysis / Explorer
- Specialization: Campaign/Donation threshold and settlement ordering facilitation
- Participant: `P-S2-002-OIR-006-1`
- Created: 2026-09-29
- Model / reasoning: Invocation selects `gpt-6-luna` / `high`; runtime model was not independently exposed.
- Session: not exposed
- Target revision observed: `934b093cea30230743dee951ae1e600762019f30`
- Workflow revision: `7a4dbf2c065bd8fd02c86c24073d7309046bff30`
- Invocation: `../invocation.md`

## Area 1 — Product/MVP requirement and approved O6 direction

### Current state

The current Product/MVP source says `max_amount` is a closure threshold, not a hard cap. A Donation submitted while Campaign is eligible may be accepted in full; accepted Donations still pending when Campaign closes remain eligible to settle in full; new submissions after closure are rejected. Funding may therefore exceed the threshold further. Slice 3 owns the broader closure/public-result experience. (`docs/product/mvp-scope.md` §Stage B / Stage D, especially lines 132–136; `docs/product/mvp-delivery-slices.md` §5, lines 186–198.)

The O6 resolution records the same direction: crossing Donation is accepted at full amount, accepted pending Donation remains settleable after threshold closure, and new post-close submission is rejected. (`OIR-S2-002-001/resolution-brief.md`, O6.) Approved Techplan `TP-S2-002-007` carries this as Q8/R7/D8 and expressly leaves only contract ordering to Campaign/Donation owners (§§3–5, §13 item 8). The current Authority Map names Anhar as Campaign and Donation delivery/domain owner for these Slice-2 details (lines 9–10).

The pinned current-effective inputs matched their invocation SHA-256 values at dispatch: Approved Techplan, O6 resolution, Authority Map, and OIR-005 amount-contract brief. OIR-005 keeps the shared-currency standard open and separate; this Run does not need to resolve or alter it.

### Requirement

Preserve full-amount acceptance and settlement for donations accepted while eligible, including those still pending after closure; reject later submissions; preserve atomic success/funding coupling, exact-once contribution, and concurrency safety. Do not add Slice 3 closure/public-result behavior, reopen O6, or invent money rounding/tax behavior. Sources: Product/MVP above; `TP-S2-002-007` §§2–4 (Q8, R3, R7) and §13.

### Gap

Product semantics are settled, but no contract-level ordering rule yet specifies how submit eligibility, accepted-pending settlement, funding updates, threshold close, and concurrent close/settlement remain consistent. The gap is a delivery/domain contract decision, not a missing Product decision on the stated O6 behavior.

### Sniffing

- **Risk:** A race that accepts a new Donation after closure, drops an accepted Donation's funding contribution, double-counts it, or commits Donation success without its funding reflection changes money state and donor-visible truth.
- **Edge cases:** A Donation itself crosses the threshold; multiple accepted pending Donations settle after the first crossing; settlement and close overlap; a submit races closure; a Campaign has no configured `max_amount` (the existing historical spec says “when set,” but this Run does not define new null-threshold behavior).
- **Miscontext:** The older domain artifacts may be mistaken for approved current semantics because they contain detailed SQL and test checklists. Product/MVP precedence makes them evidence to reconcile.
- **Misleading signals:** The presence of Q8/R7/D8 may look like complete delivery detail, while Techplan §13 explicitly leaves ordering unresolved.
- **Inconsistency:** Current Product/MVP/O6 are consistent with each other. Any lower-level `published`-only settlement rule must be reconciled against accepted-pending full settlement.

### Code/authority anchors

- `docs/product/mvp-scope.md` §Stage D, lines 128–136 — threshold and accepted-pending product rule.
- `docs/product/mvp-delivery-slices.md` §5, lines 186–198 — correctness floor, threshold rule, Slice 3 boundary.
- `TP-S2-002-007/techplan.md` §13, item 8 — open contract-ordering item and owner route.
- `.harscode-spaces/authority-map.md` lines 9–10 — current scoped Campaign/Donation owner.
- `OIR-S2-002-001/resolution-brief.md` O6 — human-settled policy provenance.

## Area 2 — Campaign/Donation delivery specs and authored OpenAPI

### Current state

The historical Donation invariant `INV-donation-02` requires Campaign `published` at submission and says the check is atomic, but the prose does not define its serialization with close. `INV-donation-08` models the successful funding update as `UPDATE campaigns ... collected_amount = collected_amount + :amount WHERE id = :id AND status = 'published' RETURNING collected_amount`, then closes at threshold in the same transaction. `INV-donation-09` makes the pending-to-terminal transition idempotent. `INV-donation-10` prevents amount clamping and notes full amount acceptance at submit, but it does not define settlement after closure. (`docs/spec/5-donation/invariants.md` lines 39–48 and 129–177.)

The historical feature spec first changes Donation `pending → success`, then attempts the funding increment only while Campaign remains `published`; it does not define what to do if that update affects zero rows. The same file says threshold crossing closes in the same transaction, but does not state how an already-closed Campaign affects an accepted pending Donation's success/funding pair or which closed reason remains authoritative. (`docs/spec/5-donation/features/01-submit-donation-settlement.md` §§Submission, Settlement, Concurrency & correctness notes; especially lines 77–91 and 102–114.)

Campaign `INV-campaign-13` and Feature 09 say the max-amount, deadline, and Admin force-close triggers share a `status = 'published'` guard; whichever close update wins records the one closure, and later close triggers no-op / return their own conflict. The spec locates the max trigger in the Donation success transaction, but does not reconcile the close guard with later accepted-pending funding contributions. (`docs/spec/4-campaign/invariants.md` lines 212–234; `docs/spec/4-campaign/features/09-closure.md` §§Behavior and Concurrency.)

Authored Donation OpenAPI says submit requires `Campaign.status = published` and returns `409` otherwise; Campaign OpenAPI describes the force-close operation and status/closed-reason fields. Neither split contract specifies the accepted-pending settlement rule, atomic serialization, funding contribution after close, or settlement-versus-close precedence. (`api/openapi/donation.yaml` `/campaigns/{campaignId}/donations`; `api/openapi/campaign.yaml` `/campaigns/{campaignId}/force-close` and Campaign schema.)

### Requirement

The active Product/MVP and approved Techplan require accepted pending Donations to settle at full amount after close, while accepting no new submissions after closure. Successful settlement and funding reflection must be one atomic outcome and count exactly once; concurrent donations cannot lose or corrupt funding increments. The contract must preserve the single winning closure state without importing broader Slice 3 result behavior. (`mvp-delivery-slices.md` §5; `TP-S2-002-007` R3/R7 and §13 item 8.)

### Gap

The historical settlement predicate requires `status = 'published'` for funding. If a Campaign closes after Donation acceptance but before settlement, the specified funding update returns no row; the spec does not say whether success is rolled back, funding is still incremented on a closed Campaign, or another outcome occurs. That leaves an explicit mismatch with accepted-pending full settlement and atomic success/funding. It also leaves submit-vs-close and concurrent accepted settlements vs threshold close without a complete ordering contract. Historical rules are not authoritative where they conflict with Product/MVP.

### Sniffing

- **Risk:** The status-first / published-only funding sequence can create a Donation-success-without-funding outcome if a closed Campaign yields zero updated rows and the transaction path does not fail atomically. Concurrent settlement/close could also lose contributions or overwrite the winning close reason.
- **Edge cases:** Threshold-crossing Donation; two or more accepted pending Donations settling after closure; settlement racing max-threshold/deadline/Admin close; new submit racing close; duplicate settlement; zero-row campaign update after the Donation transition.
- **Miscontext:** `INV-donation-08`'s `status = 'published'` funding predicate is historical delivery detail, not the approved rule for accepted pending Donations after close. The current contract must preserve the higher-level full-settlement policy.
- **Misleading signals:** A single conditional SQL update and shared close guard appear to establish concurrency safety, but they do not define the zero-row result or connect accepted pending state to funding once the Campaign is closed.
- **Inconsistency:** Product/MVP permits post-close settlement of accepted Donations; the historical Donation spec's `published`-only funding update cannot express that case. Campaign's “first close wins” rule has no stated interaction with that later funding increment.

### Code/contract anchors

- `docs/spec/5-donation/invariants.md` `INV-donation-02`, `08`, `09`, `10` — submit guard, funding update, idempotence, full amount.
- `docs/spec/5-donation/features/01-submit-donation-settlement.md` §§Submission / Settlement / Concurrency — exact historical sequence and undefined zero-row path.
- `docs/spec/4-campaign/invariants.md` `INV-campaign-13` — three close triggers and one-winner guard.
- `docs/spec/4-campaign/features/09-closure.md` max trigger + concurrency notes — threshold close source and trigger arbitration.
- `api/openapi/donation.yaml` `/campaigns/{campaignId}/donations` — submit contract currently only says `published` and `409` post-close.
- `api/openapi/campaign.yaml` `/campaigns/{campaignId}/force-close`, `CampaignStatus`, `ClosedReason` — close contract/state vocabulary, no Donation settlement ordering.

## Area 3 — Live backend evidence

### Current state

After reading `backend/AGENTS.md`, targeted repository search found only `backend/internal/domain/campaign/` and no Donation package, Donation migration, settlement worker, or Campaign threshold-close write path. The existing Campaign service/repository is the Slice-1 public detail/media read model. `RepositoryDB.FindPublicDetail` filters `c.status = 'published'` to serve public detail; it is not a Donation submit eligibility transaction. `campaign_public` projection explicitly reports donation flow unavailable. Migration `000011_create_public_campaigns.up.sql` describes itself as Slice 1 and says it does not add Donation workflow; it stores target/collected amounts but no `max_amount` field.

### Requirement

The backend must enforce submit-time Campaign eligibility and the O6 accepted-pending/full-settlement rule, atomically couple successful Donation state and funding, and preserve exact-once/concurrent increments (Product/MVP §5; Techplan R3/R7). Backend evidence could establish those behaviors only if the corresponding live write path exists.

### Gap

No live Slice-2 Donation/settlement/close implementation is present in the inspected backend, so runtime ordering cannot be confirmed or used to select a contract rule. Existing Slice-1 read behavior does not establish submit eligibility or settlement semantics. This is not evidence that the future implementation violates O6; it means implementation proof is absent and remains downstream Build/Testing work.

### Sniffing

- **Risk:** When the write path is introduced, this is money/concurrency-sensitive: an accepted Donation must not be lost after closure, counted twice, or have a terminal state without its matching funding effect.
- **Edge cases:** All submit/settle/close overlap cases remain unimplemented here and cannot be runtime-verified at this revision.
- **Miscontext:** The Campaign public read model is not the Campaign delivery/Donation write model; do not infer transactional eligibility from its public projection.
- **Misleading signals:** The backend contains a `published` predicate for public visibility and a `closed` status in the Slice-1 schema. Neither is evidence of the future Donation submission/settlement order.
- **Inconsistency:** No implementation contradiction can be established because the relevant write path is absent. This is an implementation-evidence gap, not a claim that the code violates Product.

### Code anchors

- `backend/internal/domain/campaign/repository_db.go` `RepositoryDB.FindPublicDetail` — public read query filters status to `published`; not submit eligibility.
- `backend/internal/domain/campaign/service.go` `GetPublicDetail` / `toPublicDetail` — Slice-1 public projection, with donation action unavailable.
- `backend/internal/domain/campaign/entity.go` `DonationAction` — explicitly describes Slice-1 unavailable donation flow.
- `backend/migrations/000011_create_public_campaigns.up.sql` — Slice-1 Campaign schema; no Donation workflow or `max_amount` storage.
- `backend/AGENTS.md` §§1, 6 — stack routing and authority precedence followed for inspection.

## Findings and progression effect

### F1 — Current contract ordering is unresolved and historical settlement detail conflicts with O6

Product/MVP/O6 are clear. Historical spec says the funding increment is conditional on `Campaign.status = 'published'` and leaves the zero-row outcome undefined. That cannot demonstrate the approved full settlement/funding effect for accepted pending Donations after closure. This is **decision-relevant** and blocks the current Campaign/Donation contract from being considered reconciled/ready. It does **not** block Stage 3 owner facilitation: the current Slice-2 authority owner is named, O6 is settled, and no Product policy change is needed if the owner direction preserves it. If a proposed direction changes the approved full-settlement/rejection semantics, stop and route that change to Product/MVP authority.

### F2 — No live Donation/threshold write path exists at this revision

The backend inspection found only Slice-1 Campaign read behavior. This is **informational for the owner decision** and **deferable for this Run**; it blocks runtime verification/contract readiness downstream, not Stage 3. The next route for runtime proof is a later authorized Build followed by risk-appropriate Testing; no implementation or test was performed here.

## Verification performed / not performed

- **Performed:** targeted read of current Product/MVP, Approved Techplan, O6 resolution, Authority Map, OIR-005 separation, Campaign/Donation invariants/features, authored split OpenAPI, backend scoped instructions, and targeted backend implementation/migration search; verified the four pinned input hashes in the Invocation.
- **Not performed:** tests, race tests, API lint/bundle validation, runtime checks, database migration checks, or any implementation changes.
- No `CONTRACT_READY`, runtime correctness, money-standard, or residual-risk acceptance is claimed.
