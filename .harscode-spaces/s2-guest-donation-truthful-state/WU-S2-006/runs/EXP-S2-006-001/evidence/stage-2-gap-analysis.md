# Stage 2 — Gap Analysis: Monetary Limits & Campaign Capacity

> Work Unit: `WU-S2-006`  
> Run: `EXP-S2-006-001`  
> Phase / Stage: Exploration / Stage 2 — Gap Analysis  
> Author: Explorer, `P-S2-006-EXP-001-1` (`KC-EXPLORER`)  
> Created: 2026-10-01  
> Updated: 2026-10-01 — live-source verification after Human Stage-2 confirmation  
> Model / reasoning: `gpt-6-luna` / `medium` (Invocation)  
> Session: identifier not exposed  
> Target revision: `7fd8b473b239b20bda3990ab29c51440d321a796` plus current working tree  
> Workflow revision: `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`

## Scope and source posture

Task source: `WU-S2-006/manifest.md`, especially Outcome, Scope, Boundaries, Completion condition, and Current-effective inputs. Product/MVP behavior is routed through `docs/product/README.md` to `mvp-scope.md` and `mvp-delivery-slices.md`; cross-feature monetary representation through `docs/project/kencleng-monetary-data-standard.md`; domain behavior through Donation/Campaign specs; shared contract through split OpenAPI sources and required generated/consumer counterparts.

The current owner attribution is explicit: Anhar Solehudin owns this bounded Slice-2 Product/MVP monetary/closure policy, project-wide monetary representation, and current-slice Campaign/Donation/API areas (Authority Map, effective 2026-10-01 for the scoped Product/MVP area). Attribution identifies the decision owner; it does not accept unresolved details or source changes. The Human direction recorded in `TP-S2-003-003` D16/O1-REP remains a candidate until source reconciliation: configurable per-Campaign individual limit oriented at Rp1,000,000,000, cumulative ceiling at current Campaign whole-IDR representability, and closure with a new reason at capacity. The proposed identifier `funding_capacity_reached` is not an existing accepted enum.

No implementation tests, generators, services, migrations, browser, or runtime checks were run, consistent with the Invocation.

## Area 1 — Product/MVP behavior and monetary-standard boundary

**Current state.** `docs/product/mvp-scope.md` §§4–7 and `docs/product/mvp-delivery-slices.md` §§5–6 require whole-IDR input (minimum Rp5.000, Rp1 increments), exact-decimal handling, truthful eligibility, and accepted-pending full settlement. `max_amount` is expressly a closure threshold, not a hard cap; already accepted Donations may take funding above it. Slice 3 owns broader closure/public-result experience. Neither Product source currently states the D16 individual Donation maximum, a cumulative representability ceiling, or a capacity-triggered close reason. The whole-product authority is `docs/product/product-overview.md`; the scoped MVP behavior belongs in the applicable MVP authority.

`docs/project/kencleng-monetary-data-standard.md` approves major-unit decimal strings plus explicit currency code on wire, exact decimal calculation/persistence, and no `float`/`float64`. It deliberately leaves universal precision/scale, range, and currency-specific fractions unresolved; it explicitly says Campaign `NUMERIC(19,2)` is domain precedent, not a project-wide standard.

**Requirement / gap.** Reconcile the bounded Slice-2 policy into Product/MVP while retaining the existing over-threshold/full-settlement behavior. D16/O1-REP cannot be promoted to project-wide monetary policy or silently conflate the new capacity with `max_amount`. Product wording for per-Campaign configuration, cumulative capacity, and capacity closure is currently absent.

**Sniffing lenses.**
- Risk: contradicting threshold semantics could reject or truncate an already accepted full Donation, changing the approved trust loop.
- Edge cases: whole-Rupiah boundary at the Campaign representability ceiling; nullable `max_amount`; funding that has crossed threshold while accepted pending Donations remain.
- Miscontext: a “limit” here has at least two meanings—individual admission and cumulative storage capacity—while existing Product `max_amount` is a third, intentionally overshootable threshold.
- Misleading signals: the shared exact-decimal direction does not itself imply a numeric bound or make every Donation representable in Campaign funding.
- Inconsistency: current Product/MVP permits threshold overshoot but says nothing about a separate capacity close; no explicit Product contradiction yet exists, but downstream prose could create one if the meanings are merged.

**Progression effect — decision-relevant.** Product/MVP source reconciliation requires the named Product owner. Stage 3 can proceed safely to make the missing policy/interface decisions reviewable; it cannot itself claim owner acceptance or update Product sources.

## Area 2 — Donation and Campaign domain specifications

**Current state.** Donation sources are under `docs/spec/5-donation/` (not `3-donation`). `invariants.md#inv-donation-01` fixes the minimum/whole-IDR/exact-decimal direction and leaves concrete range/precision/storage questions open. `#inv-donation-02` and `#inv-donation-08`, plus `features/01-submit-donation-settlement.md`, preserve atomic eligibility ordering, accepted-pending full settlement, exact-once full funding, and stable winning close reason; none currently records a max individual amount or capacity rule. Campaign `docs/spec/4-campaign/invariants.md#inv-campaign-13` assigns lifecycle and winning `closed_reason` to Campaign and captures the narrow D1 threshold boundary. `features/09-closure.md` references D1 but says its broader historical closure lifecycle is not reconciled here. Draft CRUD `features/01-campaign-creation-draft-crud.md` and `INV-campaign-02` cover draft editing and `max_amount`, not a Donation-specific configuration value.

**Requirement / gap.** Reconcile the accepted product decision into Donation amount acceptance and Campaign lifecycle/configuration semantics, preserving the accepted-pending and exact-once obligations. The current specs have no individual cap/configuration authority, no cumulative reservation/capacity invariant, and no capacity-specific winning close reason. The exact editable/configuration authority and relationship to draft-only editing are open.

**Sniffing lenses.**
- Risk: without a cross-domain capacity invariant, an accepted Donation could later fail full settlement or overflow public Funding representation.
- Edge cases: amount exactly at/above the individual cap; cumulative capacity with settled plus pending amounts; failed pending Donation and unused capacity; threshold and capacity closures racing; previously accepted pending amount after a different close wins.
- Miscontext: Campaign close ownership does not mean Donation can independently define a new Campaign close reason; D1 presently covers threshold/eligibility only.
- Misleading signals: `INV-campaign-13` is accepted but narrow; the reference does not establish capacity behavior or select transaction/locking details.
- Inconsistency: `features/09-closure.md` contains older implementation-shaped threshold prose; its header explicitly limits it where it conflicts with D1, so it must not silently override the accepted invariant.

**Progression effect — decision-relevant.** Domain spec reconciliation is required after Product policy is settled. Stage 3 may define the decision questions and source route; spec acceptance remains a separate owning-source gate. No implementation mechanism is selected in this Exploration.

## Area 3 — Shared API contract and source boundaries

**Current state.** `api/README.md` identifies split `api/openapi/donation.yaml` and `api/openapi/campaign.yaml` as authored sources; `api/openapi.yaml` is generated, as is `frontend/lib/api/generated/openapi.ts`. Donation `SubmitDonationRequest.amount` in `donation.yaml` has no maximum and explicitly says so. Campaign `ClosedReason` currently contains `max_amount_reached`, `deadline_reached`, and `admin_force_closed`; `CampaignCreateRequest`/`CampaignUpdateRequest` expose `target_amount` and nullable `max_amount` only. `PublicCampaignFundingAvailable` permits up to 17 integral digits plus two fractional digits. These authored definitions do not express a whole-IDR capacity close or per-Campaign Donation maximum/configuration.

`PublicCampaignDetail.donation_action` and Campaign Feature 02 expose only availability (`available`) or generic `campaign_not_eligible` at the GET snapshot. The WU-S2-005 contract is accepted and completed per parent `events.md` (2026-10-01); its producer/runtime predicate evidence remains a Delivery obligation. Its accepted scope does not tell a guest the configured maximum. The GET action contract must not be silently broadened under this work unit.

**Requirement / gap.** The accepted split sources need bounded contract reconciliation for the new amount constraint, Campaign-owned configuration and closure semantics, and any guest-facing limit/error semantics that Product chooses. A public GET field or operation change may require separate owning API scope/review/acceptance; do not assume WU-S2-005 acceptance includes such a change. Generated aggregate/type correspondence follows authored acceptance.

**Sniffing lenses.**
- Risk: schema mismatch can allow clients to submit values backend rejects without explaining the active Campaign limit, or imply a close state that no Campaign contract represents.
- Edge cases: at-cap validation and generic problem behavior; nullable/absent configuration; whether current cap is public and where; stable close-reason compatibility for existing clients.
- Miscontext: `amount: string` and exact-decimal descriptions only establish representation, not a maximum. The public action is a point-in-time availability signal, not a detailed eligibility or amount contract.
- Misleading signals: generated `openapi.ts` already types Donation amount as `string`; that type does not enforce the missing bound. Campaign's existing `max_amount` field is a different threshold.
- Inconsistency: Donation OpenAPI states “No maximum”; D16 proposes one. Campaign enum and request schemas likewise lack D16 terms. This is a real source-reconciliation gap, not a runtime finding.

**Progression effect — decision-relevant.** A known API owner exists. Stage 3 may formulate the bounded contract decision after the owner-source/product questions are made explicit. Authored API acceptance, any scope expansion, review, and generated counterparts remain required gates; no contract is accepted by this finding.

## Area 4 — Generated and known consumer surfaces

**Current state.** `frontend/lib/api/generated/openapi.ts` currently reflects the unbounded Donation amount description and three-value `ClosedReason` union. `frontend/mocks/fixtures/public-campaign.ts` exercises public Funding zero, above-target, and maximum representable decimal strings, but its public projection has no Donation cap field. The existing Campaign public projection intentionally separates narrow public-safe fields from internal Campaign data. `TP-S2-003-003` O1-REP specifically flags the guest-disclosure question and warns against silently extending WU-S2-005's availability-only contract. WU-S2-005 handoff/testing confirms authored schema correspondence was checked for that prior action contract; this evidence does not cover a new amount/capacity contract.

**Requirement / gap.** Once authored changes are accepted, enumerate and reconcile the exact generated bundle/type and fixture counterparts. Before that, the public consumer path has no current value from which the UI could explain a configured per-Campaign maximum. Exact consumer and disclosure surface is a decision question, not an implementation assumption.

**Sniffing lenses.**
- Risk: stale generated unions/types or mocks could let downstream planning/UI represent obsolete limit/close behavior.
- Edge cases: maximum whole-IDR Campaign funding is `99,999,999,999,999,999` while the schema accepts a two-decimal representation through `.99`; fixture validity at the precise whole-IDR boundary is not proven here.
- Miscontext: generated presence is not independent authority and fixtures do not prove API runtime behavior.
- Misleading signals: current max-range fixtures show representability coverage but neither enforce a cap nor test capacity reservations.
- Inconsistency: accepted WU-S2-005 correspondence is specific to its files/scope; treating it as evidence for D16 consumer readiness would be overreach.

**Progression effect — needs further evidence / decision-relevant.** Stage 3 can safely identify the consumer/disclosure question and route owner input. Exact consumer impact should be rechecked after the authored API shape is decided; generation/validation is deferred to authorized source reconciliation and was not run during Exploration.

**Live consumer anchors verified after Stage-2 confirmation.** `frontend/app/campaigns/[campaignId]/campaign-detail-view.tsx#CampaignSuccess` still renders a fixed “Alur donasi ... belum tersedia” panel and does not branch on `campaign.donation_action`; this is the current Slice-1 implementation posture, not an accepted cap/disclosure behavior. `frontend/lib/api/public-campaign.ts#getPublicCampaignDetail` returns the generated Campaign detail type without interpreting eligibility. The contract-faithful fixture in `frontend/mocks/fixtures/public-campaign.ts` now represents available and `campaign_not_eligible` branches, but does not include a configured Donation maximum. These are anchors for later delivery re-grounding, not a request to expand this shared-source Run into frontend Build.

## Area 5 — Prior decisions, ownership, and coordination dependencies

**Current state.** `TP-S2-003-003/techplan.md` D16/O1-REP and its `handoff.md` preserve the Human candidate direction and explicitly list unresolved configuration surface/range/editability, guest disclosure, Product/spec/API changes, and proposed `funding_capacity_reached`. Parent `events.md` records that scoped Product/MVP ownership was subsequently attributed to Anhar on 2026-10-01; it also records WU-S2-005 completion/acceptance separately and says producer/runtime remains downstream. WU-S2-006 manifest keeps source acceptance and material decisions as gates. Existing working-tree modifications are present in relevant API/spec/generated/fixture paths; they are live evidence but do not establish which current source edits are accepted for this new policy.

**Requirement / gap.** Preserve the candidate choices without re-voting them merely because this is a new Run, while facilitating only genuinely unresolved material details with the named owner. Completion needs Product/spec/API acceptance and required counterpart review, plus a clear delivery handoff; it does not imply runtime proof or whole-backend/frontend plan approval.

**Sniffing lenses.**
- Risk: treating owner attribution, a Planner Draft, or WU-S2-005 action acceptance as blanket D16 source acceptance could bypass required authority/review gates.
- Edge cases: overlapping closed triggers must preserve the first winning reason; independent WU-S2-003 and frontend planning remain runnable, while affected final approval/build stays gated on this reconciliation.
- Miscontext: an older `TP-S2-003-003` handoff calls Product endorsement pending because attribution had not yet occurred; parent events now map that owner, so that historical phrase must not cause an attribution re-request. Its unresolved source questions remain current.
- Misleading signals: WU-S2-005 is `DONE` for its accepted action contract, but its completion does not resolve donation maximum or capacity semantics.
- Inconsistency: no current owner-map gap remains for the scoped Product/API concerns; there is a live difference between candidate D16 and current accepted sources, intentionally routed for reconciliation.

**Progression effect — decision-relevant.** No owner-attribution blocker remains. Stage 3 can proceed with the named owner to frame unresolved policy/interface decisions. Source edits/acceptance and subsequent plan refresh/re-review remain outside this Stage-2 artifact and are not implied.

## Area 6 — Live producer/storage implementation boundary

**Current state.** The checked-in Slice-1 Campaign schema at `backend/migrations/000011_create_public_campaigns.up.sql#campaigns` stores `target_amount` and `collected_amount` as `NUMERIC(19,2)` and has no Campaign `max_amount` or `closed_reason` columns. `backend/internal/domain/campaign/entity.go#DetailRecord` contains only the public read-model funding values; its `DonationAction` comment still describes the Slice-1 unavailable action. `backend/internal/domain/campaign/service.go#toPublicDetail` hardcodes that unavailable action, while `backend/internal/transport/http/campaign_public.go#toPublicCampaignDetailResponse` copies its values to the wire. No `backend/internal/domain/donation/` package currently exists. This matches the explicitly outstanding WU-S2-005 producer/runtime and WU-S2-003 implementation obligations; the source reconciliation does not itself implement them.

**Requirement / gap.** The accepted future behavior needs source-level policy and contract decisions before backend planning can select how the individual cap/configuration, representability ceiling, and close reason are carried into producer/storage behavior. The existing Campaign column is the concrete capacity evidence behind the D16 candidate, but the project monetary standard explicitly prevents promoting its precision/scale into a universal standard. The current producer does not implement WU-S2-005's accepted `donation_action` semantics, and no Donation delivery implementation is present in this working tree.

**Sniffing lenses.**
- Risk: a future accepted Donation could outgrow the existing Campaign accumulator unless policy/contract and producer behavior are reconciled before implementation; this is the concrete basis for the WU-S2-003 representability finding.
- Edge cases: the largest whole-IDR value that fits this column; full settlement after a threshold close; capacity consumed by pending obligations; capacity released after failed settlement; and a close already won by another trigger.
- Miscontext: D16's draft wording refers to current Campaign storage capacity; the approved project-wide monetary standard says the column is precedent for this domain, not a universal storage decision.
- Misleading signals: current Campaign entity and API have `DonationAction` fields/types, but the service hardcodes the old Slice-1 action and does not evaluate the accepted WU-S2-005 predicate.
- Inconsistency: the accepted Campaign contract and updated spec now expose an available/unavailable eligibility union, while the live Campaign service/test fixture still emits `donation_flow_not_available`; producer fidelity remains a separate, already-routed Delivery obligation.

**Progression effect — decision-relevant / needs further delivery evidence.** This does not block Stage 3 policy/interface solutioning or source reconciliation. It does block any claim that source acceptance establishes producer readiness; WU-S2-003 must refresh its plan and retain the producer/storage and runtime proof obligations after WU-S2-006 sources converge. No test or runtime check was run here; the test file was read only as a code anchor.

## Findings and progression summary

1. **F-01 — Product policy absent from accepted Product/MVP text.** The current threshold rule is explicit; individual cap, cumulative capacity and capacity-close behavior are absent. **Decision-relevant**; blocks final source reconciliation/acceptance for this scope, but Stage 3 can proceed with Anhar as the mapped owner. Next: facilitate the bounded Product/MVP decisions and record the direction before source changes.
2. **F-02 — Donation/Campaign specs do not encode D16.** Existing invariants preserve D1/full settlement and stable close reason but have no accepted capacity rule or new reason. **Decision-relevant**; blocks affected final Techplan approval/Build until specs are reconciled and accepted. Stage 3 can proceed; next: derive exact spec decisions from settled Product policy and route acceptance to the domain owner.
3. **F-03 — Authored API omits cap/configuration and capacity closure.** Donation amount has no maximum; Campaign lacks the per-Campaign Donation setting and capacity reason; guest disclosure is unresolved. **Decision-relevant**; blocks affected contract readiness and dependent final approval/Build. Stage 3 can proceed; next: settle public disclosure and contract scope with the known API/Product owner, then reconcile and review authored sources.
4. **F-04 — Generated/fixture and consumer proof is scoped to the previous contract.** Existing generated types and fixtures do not represent D16; WU-S2-005 evidence only proves its accepted availability action contract. **Needs further evidence** after authored shape is decided; does not block Stage 3. Next: re-derive impacted counterparts/consumers, then validate/generate in the owning reconciliation workflow.
5. **F-05 — Capacity and threshold semantics must remain distinct under D1.** Candidate capacity closure must coexist with accepted-pending full settlement, exact-once Funding, and stable winning reason. **Decision-relevant**; blocks safe delivery planning until owner-approved specification makes interaction explicit. Stage 3 can proceed; next: frame the precedence/edge-case decisions without choosing a transaction mechanism or weakening accepted obligations.
6. **F-06 — Current consumer does not expose the accepted Campaign action union or a configured Donation maximum.** The live public Campaign page still shows its Slice-1 “donation flow unavailable” panel regardless of the API action value; fixtures model the accepted availability branches but no limit. **Needs further delivery evidence / decision-relevant** for frontend refresh and guest disclosure; it does not block Stage 3 or source reconciliation. Next: settle Product/API disclosure and field-scope questions, then refresh affected frontend planning against accepted sources.
7. **F-07 — Current producer/storage remains Slice-1 only.** Campaign funding is stored in `NUMERIC(19,2)`, public detail hardcodes the Slice-1 action, and no Donation domain package exists. **Decision-relevant**; blocks implementation/readiness claims, not Stage 3. Next: after source convergence, the backend Planner reopens these live anchors, refreshes the affected plan and retains producer/runtime evidence as downstream obligations.

There is no active inability to perform Stage 3: the relevant Product and API owner is named, sources and concrete gaps are discoverable, and no missing runtime evidence is needed to facilitate the policy decisions. Affected source acceptance/final Techplan approval/Build are not ready until the decision and reconciliation gates are complete. Independent planning/work already identified by the Work Graph remains governed by those dependencies.

## Context pointers for Stage 3

- Task and gates: `WU-S2-006/manifest.md`; `.harscode-spaces/authority-map.md`; parent `events.md` and `work-graph.md`.
- Candidate policy and exact unresolved items: `WU-S2-003/runs/TP-S2-003-003/techplan.md` D16/O1-REP and `handoff.md`.
- Product/monetary authorities: `docs/product/mvp-scope.md` §§4–7; `docs/product/mvp-delivery-slices.md` §§5–6; `docs/project/kencleng-monetary-data-standard.md`.
- Domain anchors: `docs/spec/5-donation/invariants.md#inv-donation-01`, `#inv-donation-02`, `#inv-donation-08`; `features/01-submit-donation-settlement.md`; `docs/spec/4-campaign/invariants.md#inv-campaign-02`, `#inv-campaign-13`; `features/01-campaign-creation-draft-crud.md`; `features/02-campaign-detail-listing.md`; `features/09-closure.md`.
- Contract/consumer anchors: `api/openapi/donation.yaml#/components/schemas/SubmitDonationRequest`; `api/openapi/campaign.yaml#/components/schemas/ClosedReason`, `CampaignCreateRequest`, `CampaignUpdateRequest`, `PublicCampaignDetail`, `PublicCampaignFundingAvailable`; `api/README.md`; `frontend/lib/api/generated/openapi.ts`; `frontend/mocks/fixtures/public-campaign.ts`.
- WU-S2-005 boundary: `WU-S2-005/runs/TST-S2-005-001/handoff.md` and parent event recording its accepted action scope; its producer/runtime follow-up remains separate.
- Live delivery anchors verified in Stage 2: `backend/migrations/000011_create_public_campaigns.up.sql#campaigns`; `backend/internal/domain/campaign/entity.go#DetailRecord` / `DonationAction`; `backend/internal/domain/campaign/service.go#toPublicDetail`; `backend/internal/transport/http/campaign_public.go#toPublicCampaignDetailResponse`; `frontend/app/campaigns/[campaignId]/campaign-detail-view.tsx#CampaignSuccess`; `frontend/lib/api/public-campaign.ts#getPublicCampaignDetail`.
