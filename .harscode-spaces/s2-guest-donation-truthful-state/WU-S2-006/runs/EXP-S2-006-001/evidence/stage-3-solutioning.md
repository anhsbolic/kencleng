# Stage 3 — Solutioning: Monetary Limits & Campaign Capacity

> Work Unit: `WU-S2-006`  
> Run: `EXP-S2-006-001`  
> Phase / Stage: Exploration / Stage 3 — Solutioning  
> Author: Explorer, `P-S2-006-EXP-001-1` (`KC-EXPLORER`)  
> Created: 2026-10-01  
> Model / reasoning: `gpt-6-luna` / `medium` (Invocation)  
> Session: identifier not exposed  
> Target revision: `7fd8b473b239b20bda3990ab29c51440d321a796` plus current working tree  
> Workflow revision: `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`

## Settled direction preserved

The Human's choices in `WU-S2-003/runs/TP-S2-003-003/techplan.md` D16/O1-REP are retained without a new vote:

- an individual Donation limit is configurable per Campaign, with Rp1,000,000,000 as the recorded direction;
- cumulative Campaign Funding must remain representable under the current whole-IDR Campaign capacity;
- reaching capacity closes the Campaign with a new reason;
- a Donation accepted while eligible remains settleable for its full amount, and `max_amount` stays a distinct, overshootable threshold.

The proposed identifier `funding_capacity_reached`, per-Campaign setting/range, configuration authority, guest disclosure, and capacity-vs-pending interaction are not yet accepted source requirements. Product/MVP and API/spec sources remain unchanged by this Exploration Run.

## Decision brief 1 — Meaning and admissible range of Rp1,000,000,000

**Problem.** D16 says the individual maximum is configurable per Campaign and “set to” Rp1,000,000,000, while O1-REP leaves the admissible configuration range undefined. The sources do not say whether Rp1 billion is a hard upper bound or the initial value that can be raised.

**Current context.** Slice 2 accepts whole-IDR donations of at least Rp5,000. No individual upper bound is in the accepted Product or Donation API. The shared monetary standard leaves universal range/precision open. A configured cap below the minimum would make the Campaign unable to accept any valid Donation.

**Options.**

1. Treat Rp1,000,000,000 as the maximum value any Campaign may configure; allow per-Campaign values from Rp5,000 through Rp1,000,000,000 in whole Rupiah, and require an explicit value before the Campaign accepts Donations.
2. Treat Rp1,000,000,000 as an initial/default direction only; allow values above it, with a separately established upper range.

**Recommendation.** Option 1. It gives “maximum” a stable product meaning, keeps configuration within the recorded direction, and prevents an unset or sub-minimum value from producing ambiguous eligibility. This does not determine a UI default; each Campaign would carry an explicit setting.

**Decision ask.** Confirm Option 1, or clarify whether Rp1 billion is only a default and what upper bound applies.

## Decision brief 2 — Configuration authority and edit window

**Problem.** “Configurable per Campaign” does not identify who sets the cap, where it is set, or whether it can change after publication.

**Current context.** MVP scope permits initial Campaign setup to be seeded or operator-assisted when persisted and authorized. Full Owner self-service and Campaign creation UI are outside the baseline MVP. Existing Campaign draft CRUD is historical delivery evidence and does not establish a new cap authority.

**Options.**

1. For MVP, an authorized operator supplies a required per-Campaign cap during persisted setup before publication; it becomes immutable once the Campaign is public. A later change requires a separately authorized lifecycle path.
2. Add the cap to an Owner-managed Campaign draft create/edit surface; permit edits only while draft, consistent with the existing draft-only CRUD contract.

**Recommendation.** Option 1 for MVP. It uses the approved narrow operational posture without adding self-service breadth, while freezing a material donor-facing limit once public. The exact operator identity/authorization mechanism remains delivery-owned and must not be inferred here.

**Decision ask.** Confirm operator-assisted pre-publication configuration locked after publication, or select Owner-managed draft configuration.

## Decision brief 3 — Guest disclosure and closed-Campaign behavior

**Problem.** A guest needs to know the active individual cap, and the contract must distinguish an ineligible Campaign from a donation amount above its cap without leaking internal capacity or closure details.

**Current context.** WU-S2-005 accepted an availability-only `donation_action` snapshot: `available` or generic `campaign_not_eligible` while the detail remains public. POST rechecks eligibility. Its accepted contract does not include the Donation cap. The current Campaign API keeps non-public/closed Campaign detail at its existing `404` behavior. Slice 3 owns persistent public closed-Campaign result experience; this WU must not silently change that future scope. No new public reason is implied by the internal capacity-close decision.

**Options.**

1. Return the configured individual cap through the public Campaign detail contract so the guest can see it before entering/submitting an amount; keep availability and submission errors generic, and do not expose cumulative remaining capacity or `funding_capacity_reached` as a public reason. Preserve the current closed/non-public detail behavior in this Slice-2 reconciliation; persistent public closed-Campaign detail remains Slice 3 work.
2. Keep the GET contract availability-only and disclose the cap only when a submitted amount is rejected; preserve the same current closed/non-public detail behavior and leave persistent public closed-Campaign detail to Slice 3.

**Recommendation.** Option 1. It lets the donor understand an active per-Campaign limit before committing to submission and avoids treating a failed POST as the first source of policy. The additional public field/operation scope must be explicitly reconciled and accepted; no exact field name or placement is selected here. Existing generic errors and closed detail behavior remain in force for this reconciliation; Slice 3 remains responsible for the approved persistent public result after closure.

**Decision ask.** Confirm pre-submit public disclosure of the individual cap with generic ineligible/closed behavior, or choose error-only disclosure.

## Decision brief 4 — Capacity closure with accepted-pending Donations

**Problem.** The Campaign column must hold settled Funding and every already accepted Donation in full. Capacity may be fully committed by accepted `pending` Donations before those Donations settle.

**Current context.** Accepted D1/D1-derived Donation invariants require accepted-pending Donations to settle in full after close, exact-once full Funding, and a stable winning close reason. The backend Draft proposes reserving settled plus accepted-pending amounts but is not an accepted Product/spec source.

**Options.**

1. Count settled Funding plus all accepted-pending obligations against capacity; when no representable capacity remains for a new Donation, close with the capacity reason. If an accepted pending Donation later fails and releases reservation, keep the Campaign closed under the stable winning-reason rule.
2. Stop accepting new Donations when settled Funding plus accepted-pending obligations consume capacity, but keep the Campaign open until successful settlement makes collected Funding reach the ceiling. If a pending Donation fails first, its released capacity can permit new Donations.

**Recommendation.** Option 1. It prevents an accepted Donation from later overflowing Campaign Funding and closes as soon as the chosen cumulative capacity is fully committed. Keeping closure after a failure follows the accepted no-reopen/stable-reason boundary; it can leave unused capacity, which should be stated as a consequence rather than hidden. Option 2 preserves that full-settlement safety but introduces a public Campaign that remains open while temporarily unable to accept new Donations and can resume if a pending Donation fails.

**Decision ask.** Confirm reservation-based capacity closure and no reopening after a pending failure, or direct a different behavior that still guarantees full settlement without overflow.

## Decision brief 5 — Capacity close reason

**Problem.** D16 requires a new Campaign close reason, but the exact stable identifier is still only a Draft proposal.

**Current context.** `api/openapi/campaign.yaml#ClosedReason` currently contains `max_amount_reached`, `deadline_reached`, and `admin_force_closed`. Product treats `max_amount` as a distinct threshold. The proposed `funding_capacity_reached` must not be conflated with `max_amount_reached` or exposed publicly by default.

**Options.**

1. Accept `funding_capacity_reached` as the new internal/API enum reason for closure caused by exhausted cumulative representability capacity.
2. Keep the behavior decision but defer the identifier until Campaign/API source reconciliation selects a different stable enum value.

**Recommendation.** Option 1. It names the actual closure condition and distinguishes it from the target/threshold rule. The enum's external contract use still requires authored API review/acceptance.

**Decision ask.** Confirm `funding_capacity_reached` as the identifier, or provide the preferred identifier.

## Progression and boundaries

These are owner decisions, not implementation choices. The named owner is Anhar Solehudin for scoped Slice-2 Product/MVP, Campaign/Donation, API, and project-wide monetary concerns. There is no active Blocker to making these decisions. Until they are reconciled into owning sources, the affected backend final Techplan approval/Build and frontend amount-limit-dependent approval/Build remain gated; independent work remains governed by the parent Work Graph.

This Explorer does not edit Product/MVP, monetary-standard, spec, API, generated, fixture, or production sources; does not authorize protected Tier-0 implementation; and does not claim runtime proof. After the Human decisions, the source-owning reconciliation route must update/review/accept affected sources and coordinate dependent plan refreshes.

## Human decisions — confirmed

On 2026-10-01, Anhar Solehudin replied “setuju semua rekomendasi” in this Participant interaction. This confirms the five recommendations above as the current Human-selected direction:

1. Rp1,000,000,000 is the hard maximum configurable value; each Campaign has an explicit whole-IDR cap from Rp5,000 through Rp1,000,000,000.
2. For MVP, an authorized operator sets the persisted cap before publication; it is immutable once public. A later change requires a separately authorized lifecycle path. Operator identity/authorization mechanics remain delivery-owned.
3. Guests can see the configured individual cap before entering/submitting an amount. Keep eligibility/submission errors generic; do not expose cumulative remaining capacity or the internal capacity-close reason. Preserve current closed/non-public detail behavior within this Slice-2 reconciliation; Slice 3 retains the approved persistent public closed-Campaign experience.
4. Capacity is consumed by settled Funding plus accepted-pending obligations. Close when those commitments leave no representable capacity for a new Donation. If a pending Donation later fails, retain the closed state/reason, even though this may leave unused capacity.
5. Accept `funding_capacity_reached` as the distinct Campaign close-reason identifier for cumulative representability exhaustion; it remains separate from `max_amount_reached`.

These decisions establish the selected product/domain direction for source reconciliation. They do not by themselves update or constitute acceptance of Product/MVP, spec, API, generated, fixture, or implementation sources.

## Phase handoff

- **Completed:** Stage 1 routing checkpoint; Stage 2 authority/source gap analysis with live producer/consumer anchors; Stage 3 option/trade-off analysis and Human selection of all five recommendations.
- **Artifacts:** `evidence/stage-2-gap-analysis.md`; `evidence/stage-3-solutioning.md`.
- **Human decision:** None needed to close this Explorer Run. The direction is selected; owning-source reconciliation, review, and acceptance remain gates.
- **Open / deferred:** Publish and reconcile the selected direction into applicable Product/MVP, Donation/Campaign specs and authored split API; decide exact contract placement/schema and required counterparts through those sources; perform required independent review/owner acceptance. Backend producer/storage/runtime work and frontend integration/rendered acceptance remain delivery obligations. No source edits, generators, tests, migrations, or runtime checks were performed by this Run.
- **Recommended next step:** Orchestrator routes the selected direction into an authorized source-owning reconciliation and coordinates required source review/acceptance; then start fresh backend/frontend Techplan Runs to refresh affected plans from accepted sources.
- **Session transition:** This orchestrated Exploration Run is complete. Any later source-reconciliation, Techplan, or re-entry work uses a new Run, Participant, and fresh Session/context, reconstructed from this handoff and current authorities.
- **Context pointers:** WU-S2-006 manifest; `.harscode-spaces/authority-map.md`; `docs/product/mvp-scope.md` §§5–7; `docs/project/kencleng-monetary-data-standard.md`; Donation invariants/feature and Campaign invariants/features; `api/openapi/donation.yaml#SubmitDonationRequest`; `api/openapi/campaign.yaml#ClosedReason`, `CampaignCreateRequest`, `CampaignUpdateRequest`, `PublicCampaignDetail`, `PublicCampaignFundingAvailable`; `TP-S2-003-003` D16/O1-REP; WU-S2-005 accepted contract and producer follow-up; WU-S2-004 OI-1; live anchors in the Stage-2 evidence.
