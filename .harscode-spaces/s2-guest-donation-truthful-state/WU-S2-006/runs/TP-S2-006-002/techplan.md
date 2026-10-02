# Tech Plan: Slice 2 Donation Limits & Campaign Capacity Source Reconciliation

> Phase             : Techplan
> Ticket            : WU-S2-006 / TP-S2-006-002
> Author            : P-S2-006-TP-002-1 (Planner)
> Participant ID    : P-S2-006-TP-002-1
> Profile           : KC-PLANNER
> Role              : Planner
> Model             : Invocation configures `gpt-6-luna`; active runtime model not independently exposed
> Reasoning         : Invocation configures `medium`; active runtime effort not independently exposed
> Session           : not exposed
> Created           : 2026-10-01
> Updated           : 2026-10-01
> Target revision   : `7fd8b473b239b20bda3990ab29c51440d321a796` plus current working tree; source anchors re-read
> Workflow revision : Harscode `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`
> Status            : Draft / In Review
> Approach          : Proposal material contract/configuration compatibility and source ownership under settled D-01–D-04; owner decision and source acceptance remain separate gates.
> Refs              : WU-S2-006 manifest; EXP-S2-006-001 evidence + handoff; parent `events.md` receipts; TP-S2-006-001 Draft; TP-S2-003-003 D16/O1-REP; accepted WU-S2-005 and Donation baseline WU-S2-002.

---

## 1. Background

RV-S2-003-001 found that a Donation valid under the current contract might not fit fully in Campaign `collected_amount NUMERIC(19,2)`. The Human settled four policy directions in the current parent event receipt: a per-Campaign whole-IDR cap (Rp5.000–Rp1.000.000.000, default Rp1.000.000.000), Owner/Staff draft configuration frozen on publication, pre-entry disclosure plus POST recheck, and reservation-based capacity close with a Slice-3 public-result handoff. A later owner receipt also accepts `funding_capacity_reached` as the distinct capacity close reason. Those receipts select policy direction and identifier; they do not accept concrete source edits.

The predecessor Draft left OI-1 too broad for an authoring agent to execute safely. This revision proposes exact request/response locations, omission/default treatment, old-row handling, and over-cap error shape. These remain an owner decision until explicitly accepted; this plan is not source acceptance or Build authorization.

## 2. Scope

**In scope:**

- Reconcile Product/MVP, Campaign and Donation specifications, and authored split APIs for the settled cap, capacity, closure, and Slice-3 boundary.
- Decide and document request/response compatibility and old-row default treatment from the concrete proposal in D6 / OI-1 before source authoring.
- Identify generated/bundle/type/fixture counterparts and source review/acceptance gates.
- After accepted source convergence, route separate fresh backend and frontend Techplan refreshes under the Work Graph.

**Out of scope (explicit):**

- Backend/frontend implementation, migrations, DB application, generators, tests, or runtime/browser checks in this planning Run.
- Transaction/locking choices or edits to Tier-0 ledger/locking code; any protected implementation.
- A universal monetary precision/range policy, unrelated Donation behavior, real payment settlement, or public exposure of internal remaining capacity/reason.
- Activating closed-Campaign detail visibility in Slice 2 or completing Slice-3 delivery.
- Accepting source changes, residual risk, this whole Techplan, or delivery milestones.

## 3. Requirements

| ID | Requirement | Source / evidence |
|---|---|---|
| Q1 | Each Campaign has a whole-IDR individual Donation cap in Rp5.000–Rp1.000.000.000, default Rp1.000.000.000; Owner/Staff can configure it in draft and it is frozen after publication. | Parent `events.md`, “Human clarified EXP006 D-01–D-04”; `WU-S2-006/manifest.md`. |
| Q2 | Public Campaign detail discloses the active cap before amount entry; GET is a snapshot and POST independently rechecks eligibility and cap. | Same parent receipt; accepted WU-S2-005 contract; `docs/spec/4-campaign/invariants.md#inv-campaign-14`. |
| Q3 | Capacity admission accounts for settled Funding plus accepted-pending obligations against Campaign whole-IDR ceiling `99,999,999,999,999,999`; exhaustion closes with the distinct `funding_capacity_reached` reason. | `TP-S2-003-003` D16/O1-REP; `RV-S2-003-001`; Campaign migration anchor; current parent OI-2 owner receipt. |
| Q4 | Accepted pending Donations settle their full amount exactly once after closure. Failure contributes no Funding and does not reopen a capacity-closed Campaign. `max_amount` remains a distinct, overshootable threshold. | Accepted Donation INV-02/08 and Campaign INV-13; parent D-03 receipt. |
| Q5 | Slice 3 owns public closed-Campaign truth: preserve identity, remove donation action, and do not mark Funding final while accepted Donations remain pending. Current Slice-2 closed/non-public behavior remains 404. | Parent D-04 receipt; WU-S2-006 manifest; current Campaign INV-14. |
| Q6 | Monetary wire/data values follow the shared major-unit decimal-string/exact-decimal convention without `float`/`float64`; this feature does not set universal precision or scale. | `docs/project/kencleng-monetary-data-standard.md`; root `AGENTS.md`; `api/README.md`. |

## 4. Rules & Validation

- **R1 — Configuration boundary:** For a draft Campaign, an authorized Owner/Staff can set a cap from Rp5.000 through Rp1.000.000.000 in whole Rupiah; omission resolves to the default Rp1.000.000.000. After publication, Slice-2 configuration cannot change it.
- **R2 — Disclosure and authoritative submit:** Public detail includes the active cap before amount entry and preserves WU-S2-005 `donation_action` snapshot semantics. Donation POST independently checks current eligibility and cap; an over-cap response is generic and does not reveal cumulative remaining capacity or internal close reason.
- **R3 — Capacity admission:** With settled Funding `F`, accepted pending obligations `P`, new amount `A`, and cumulative Campaign ceiling `C`, accept only if `F + P + A ≤ C`; close for capacity when no additional valid amount can be represented.
- **R4 — Close reason and threshold distinction:** The first winning close reason remains stable. Capacity uses `funding_capacity_reached`, distinct from `max_amount_reached`; the latter remains an overshootable threshold for already accepted full-value Donations.
- **R5 — Pending settlement/failure:** An accepted pending Donation can settle in full once after close without reopening Campaign or replacing its reason. Failure contributes no Funding and cannot reopen a capacity-closed Campaign even if it releases reservation.
- **R6 — Slice-3 public closure applicability:** Later Slice-3 behavior preserves public identity, removes donation action, and keeps Funding non-final while accepted Donations are pending. This does not alter current Slice-2 closed/non-public 404 behavior.
- **R7 — Exact representation:** Configured limits, reservations, persisted Funding, wire values, generated contracts, and fixtures preserve exact whole-IDR values through the accepted capacity boundary, without binary floating point or silent truncation.

## 5. Decision Log

| ID | Decision / option | Status | Rationale / consequence |
|---|---|---|---|
| D1 | Per-Campaign hard cap: Rp5.000–Rp1.000.000.000, default Rp1.000.000.000; Owner/Staff draft configuration frozen on publication. | Chosen — Human, 2026-10-01; source acceptance pending | Current durable receipt supersedes historical brief mismatch. |
| D2 | Disclose cap on public detail and recheck independently on POST. | Chosen — Human, 2026-10-01; API amendment pending | Enables informed amount entry; GET is not authorization. |
| D3 | Reserve settled Funding plus accepted pending; close at capacity; no reopening after later failure. | Chosen — Human, 2026-10-01; source acceptance pending | Guarantees accepted Donations remain settleable in full; unused capacity after failure is an explicit consequence. |
| D4 | Capacity-close public result belongs to Slice 3; retain current Slice-2 404 behavior. | Chosen — Human, 2026-10-01; Slice-3 source handoff pending | Keeps current scope bounded while preserving future public truth. |
| D5 | Capacity close reason is `funding_capacity_reached`. | Resolved — Anhar owner receipt, 2026-10-01; authored source acceptance pending | Distinct from `max_amount_reached`; propagate to Campaign/API sources, do not repeat the identifier decision. |
| D6 | Proposal: name the per-Campaign field `max_donation_amount`; accept it as optional in create/patch requests for legacy-client compatibility, with omission resolving to persisted Rp1.000.000.000; expose it as a required field on Campaign and public detail responses; backfill existing rows to the same default; over-cap POST returns the shared generic `422 ValidationError` at `amount`. | Active — awaiting explicit owner choice | Fits existing Campaign draft create/patch operations, avoids a new configuration endpoint, preserves omitted-field clients, makes public disclosure unambiguous, and uses the existing shared per-field validation envelope. Since `max_amount` already has a separate threshold meaning, sources must keep the new field semantically distinct. Adding a required response property and altering existing row values need counterpart/client compatibility review. |
| D6-alt | Add a separate cap-setting endpoint or use another field name/error status / preserve null for existing Campaigns. | Open alternative pending D6 decision | Requires a separately accepted operation/semantic; null cap would leave published eligibility undefined and therefore needs another owner-selected fallback rule. |
| D7 | Keep shared monetary standard format-only; no new platform-wide precision/scale. | Chosen — existing shared authority | Feature limit and Campaign storage ceiling are scoped to this behavior. |
| D8 | Preserve WU-S2-005 action union and privacy; do not expose cumulative reserve or capacity reason publicly. | Chosen — accepted predecessor contract plus current owner direction | Only the limit disclosure and current source changes need fresh acceptance. |

## 6. Backward Compatibility

- **Existing rows:** D6 proposes defaulting every existing Campaign row, including published rows, to Rp1.000.000.000 before cap-dependent behavior is enabled. This avoids an unset published cap. Source and migration owners must accept exact backfill semantics; no migration/application is authorized here.
- **Requests:** D6 makes the new create/patch request property optional with server default so old clients that omit it continue to create/update drafts. An explicitly supplied cap is validated and draft-only. Existing client/schema compatibility must be reviewed against actual generated and external consumers.
- **Responses:** D6 proposes a required `max_donation_amount` in Campaign and public detail projections. This is additive for tolerant clients but may affect strict decoders because the public schema is closed (`additionalProperties: false`); generated types/fixtures and known consumers require coordinated update and compatibility review before acceptance.
- **Donation submit:** Global amount schema remains valid as a money representation; campaign-specific over-cap is proposed as shared `422 ValidationError` on `amount`. Reusing generic field validation avoids a public capacity signal, but owner must accept the status/type/message contract.
- **Migration/deprecation:** Do not widen `NUMERIC(19,2)` by assumption. Backfill, rollout ordering, rollback, and manual DB application belong to later delivery plans and Human gates.

## 7. Edge Cases & Risks

| ID | Risk / edge case | Likelihood | Severity | Mitigation / accepted exposure |
|---|---|---:|---:|---|
| E1 | Accepted Donation cannot settle fully because Funding plus reservations exceed storage capacity. | Medium | Critical | Align source admission semantics; downstream PostgreSQL concurrency proof required. |
| E2 | Concurrent admissions both observe capacity and over-reserve. | Medium | Critical | Preserve specialized Testing obligation; mechanism remains delivery-owned and protected where applicable. |
| E3 | Pending failure releases capacity after close and reopens Campaign. | Medium | High | Stable winning close and no-reopen rule; unused capacity is documented. |
| E4 | Capacity is confused with `max_amount` or internal reason/remaining capacity leaks publicly. | Medium | High | Keep distinct rule; inspect all projections/errors and Slice-3 handoff. |
| E5 | Cap missing, out of range, mutable after publication, or inconsistent between config and POST. | Medium | High | D6 proposes default/backfill and specific fields; source acceptance plus producer/POST tests prove parity. |
| E6 | Slice-2 begins returning closed details because Slice-3 truth is carried forward. | Low | High | Keep applicability explicit; current 404 remains until Slice-3 source acceptance. |
| E7 | Amount/Funding rounds, truncates, or fails at the representability boundary. | Medium | High | Exact-decimal representations and boundary evidence across storage, API, generated types, and consumers. |
| E8 | D6 is mistaken for accepted source contract before owner receipt. | Medium | High | Keep OI-1 active and D6 explicitly proposed until a durable owner decision; no source authoring before the decision. |

## 8. Interface Contract

**Persistence/data shape:** Current Campaign funding uses `NUMERIC(19,2)` in `backend/migrations/000011_create_public_campaigns.up.sql#campaigns`, whose maximum whole-IDR value is `99,999,999,999,999,999`. Persist an effective per-Campaign cap in the accepted range/default and make it immutable after publication. D6 proposes defaulting existing rows and new omitted request values to Rp1.000.000.000; exact column/type/backfill remains source-owner work after OI-1.

**API/event/external interface:** Current authored operations are `POST /organizations/{organizationId}/campaigns`, `PATCH /campaigns/{campaignId}`, public `GET /campaigns/{campaignId}`, and `POST /campaigns/{campaignId}/donations`. D6 proposes `max_donation_amount` on create/patch (optional with default), Campaign response, and required public detail; over-cap submit uses shared `422 ValidationError` with field `amount`. Add `funding_capacity_reached` to Campaign `ClosedReason`. Keep `donation_action` unchanged, generic public errors, and closed/non-public 404 in Slice 2. No capacity/reason disclosure in public projection. These are proposed exact shapes until OI-1 and owning-source acceptance.

**Cross-layer/business boundary:** Campaign owns cap configuration, lifecycle and winning close reason. Donation owns accepted status and full-value contribution. Admission must atomically reserve capacity; accepted pending Donations settle once at full value after close. GET cap/action is a snapshot; POST is authoritative. This plan selects no locking/transaction implementation.

## 9. Architecture / Plan

1. Resolve OI-1 against D6 and record the Human receipt durably; revise the proposal if directed. Do not start owning-source edits while the contract choice remains open.
2. Reconcile Product/MVP wording and Campaign/Donation invariants/features against Q1–Q6, including default/backfill and exact compatibility decision.
3. Reconcile split Campaign and Donation OpenAPI, shared Problem usage if needed, generated bundle/types, and public fixture/consumers. Keep Campaign close reason internal and preserve current Slice-2 closed-detail behavior.
4. Independently review material source changes and required contract/counterpart correspondence; obtain named-owner acceptance before WU-S2-006 is complete.
5. After source convergence, start fresh backend WU-S2-003 and frontend WU-S2-004 Techplan Runs; preserve independent implementation, review, Testing, and runtime gates.

There is no independently operable migration/runbook in this source reconciliation Run. Migration rollout/application is delivery planning after sources settle.

## 10. Implementation Details

| Anchor | Why relevant | Intended change / precedent |
|---|---|---|
| `docs/product/mvp-scope.md` §§5–7, 12–13; `docs/product/mvp-delivery-slices.md` §§5–6 | Own approved guest behavior, slice boundary, and closure result | Reconcile cap/default/config/freeze and capacity semantics; preserve threshold distinction and Slice-3 applicability. |
| `docs/project/kencleng-monetary-data-standard.md` “Approved representation direction” / “Deliberately unresolved” | Shared representation authority; does not set feature cap | Preserve exact decimal direction; do not add universal range/scale. |
| `docs/spec/4-campaign/invariants.md` INV-campaign-02/13/14; Campaign features 01/02/09 | Campaign owns draft configuration, close reason, public projection | Reconcile cap validation/freeze, capacity close, required disclosure and later Slice-3 behavior. |
| `docs/spec/5-donation/invariants.md` INV-donation-01/02/08; `features/01-submit-donation-settlement.md` | Donation owns amount acceptance and full exact-once settlement | Add cap and capacity interaction without weakening pending settlement, idempotency, or threshold overshoot. |
| `api/openapi/campaign.yaml` `CampaignCreateRequest`, `CampaignUpdateRequest`, `Campaign`, `PublicCampaignDetail`, `ClosedReason` | Live authored configuration, response, public allowlist, enum | Resolve OI-1 then add accepted fields/defaults and distinct reason; preserve closed public projection boundary. |
| `api/openapi/donation.yaml` `SubmitDonationRequest` POST operation; `api/openapi/common.yaml` `ValidationError` / `ValidationProblem` | Live submit amount contract and shared per-field error shape | Express cap validation and accepted generic error without capacity leakage. |
| `api/README.md`; `api/openapi/index.yaml`; `api/openapi.yaml` | Authored split source, routing, generated aggregate workflow | Modify only affected authored split sources; update aggregate/generated counterparts through documented workflow, not by hand. |
| `backend/migrations/000011_create_public_campaigns.up.sql#campaigns` | Current representability ceiling | Read-only evidence in this Run; later delivery plans persisted cap/backfill and DB proof. |
| `backend/internal/domain/campaign/entity.go` `DetailRecord`/`DonationAction`; `service.go` `toPublicDetail`; `backend/internal/transport/http/campaign_public.go` `toPublicCampaignDetailResponse` | Live producer path for accepted public detail/action | Later backend planning includes cap projection and submit parity; no code change here. |
| `frontend/lib/api/generated/openapi.ts`; `frontend/mocks/fixtures/public-campaign.ts`; `frontend/lib/api/public-campaign.ts#getPublicCampaignDetail`; `frontend/app/campaigns/[campaignId]/campaign-detail-view.tsx` | Generated and rendered consumers | Update counterparts after authored API acceptance, then refresh frontend plan; no frontend production edits here. |
| `WU-S2-003/runs/TP-S2-003-003/techplan.md` D16/O1-REP; `RV-S2-003-001/review-findings.md` | Source of funding representability finding | Preserve capacity rationale and full settlement boundary. |
| `WU-S2-005/manifest.md`; `WU-S2-005/runs/TST-S2-005-001/handoff.md` | Accepted availability-only contract and its runtime boundary | Treat as baseline; any public response amendment requires fresh review/acceptance. |

## 11. Files Changed / Files NOT Changed

| File / area | Change type | Description |
|---|---|---|
| `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-002/techplan.md` | New, Run-local | Successor Draft with exact D6 proposal, resolved OI-2, and source/verification ownership. |
| `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/TP-S2-006-002/handoff.md` | New, Run-local | Phase completion, remaining owner decision, recommendation, and next gate. |

| File / area intentionally untouched | Why |
|---|---|
| Product/MVP, monetary standard, Campaign/Donation specs, authored/generated API, fixtures, production code | Invocation reserves source authoring/acceptance for a later owning-source route. |
| Prior Exploration, prior Techplan/handoffs, Work Unit records, events, and projections | Immutable/current orchestration inputs; this Run records no parent state transition. |
| Tier-0 ledger/locking, disbursement state machine, crypto/auth files | Root `AGENTS.md` fencing; no implementation authorization. |

## 12. Testing Checklist

| Rule | Verification / evidence | Primary owner | Why this is worth running / risk if skipped |
|---|---|---|---|
| R1 | Human review of Product/spec/API default, range, draft authority/freeze; later PostgreSQL/runtime tests cover request omission, old-row backfill, range boundaries, and post-publication rejection. | Human for source policy; Testing for runtime | Missing or mutable cap makes guest eligibility undefined and can change a published donor promise. |
| R2 | OpenAPI/bundle/generated/fixture correspondence and later public GET contract test; POST boundary tests for cap and independent eligibility recheck/generic `422`. | Testing | Stale client disclosure or GET-as-authorization can mislead donors or admit an over-cap request. |
| R3 | PostgreSQL integration/concurrency evidence asserts settled plus accepted-pending reservations never exceed ceiling and no accepted Donation becomes un-settleable. | Testing | In-memory tests do not establish database ordering or concurrent reservation safety; failure risks financial inconsistency. |
| R4 | Contract/domain tests assert `funding_capacity_reached` is distinct from `max_amount_reached` and the winning close reason is stable. | Testing | Conflation or replacement falsifies lifecycle history; identifier is owner-resolved but source acceptance remains required. |
| R5 | PostgreSQL integration/replay evidence covers accepted-pending success after close, full exact-once Funding, failure releasing reservation without reopen, rollback, and stable reason. | Testing | Regression risks lost/duplicate money, overflow, or false Campaign reopening. |
| R6 | Human Product/spec review of Slice-3 wording; when Slice 3 is delivered, rendered and contract acceptance for identity/action/non-final Funding. | Human | Product truth of post-close identity and pending Funding cannot be inferred from automation; this is downstream, not current Slice-2 runtime proof. |
| R7 | Split API validation, bundle/type generation correspondence, boundary fixtures, and exact-decimal persistence/projection evidence. | Testing | Precision or scale drift can truncate values or make valid accepted Donations impossible to settle. |

### Test Focus Pointer

| Area | Why sensitive | Evidence anchor from Exploration | Still relevant post-synthesis? |
|---|---|---|---|
| Concurrent Campaign capacity reservation, Donation acceptance, settlement/Funding atomicity and close ordering | Interleavings can over-reserve, accept an un-settleable Donation, lose/duplicate Funding, or replace the winning close reason; implementation mechanism/proof remain downstream. | `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-006/runs/EXP-S2-006-001/evidence/stage-2-gap-analysis.md#area-2--donation-and-campaign-domain-specifications`; `#f-05--capacity-and-threshold-semantics-must-remain-distinct-under-d1` | Yes — source reconciliation does not establish runtime atomicity or concurrency safety. |

## 13. Open Items

### Active — needs external input or verification

1. **OI-1 — Exact contract and compatibility proposal (D6).** Owner: Anhar Solehudin for scoped Slice-2 Product/MVP, Campaign/Donation and API authority (`.harscode-spaces/authority-map.md`). Decide whether to accept `max_donation_amount` optional on create/patch with omission defaulting to Rp1.000.000.000, required on Campaign/public-detail responses, defaulting existing rows to Rp1.000.000.000, and generic shared `422 ValidationError` on `amount` for over-cap POST; otherwise provide directed edits. Evidence: current API schemas and shared validation response. This blocks concrete source authoring; it is not a blanket approval request.
2. **OI-3 — Slice-3 source handoff/applicability.** Owner: Anhar as scoped Product/Campaign authority. Carry D-04 into Slice-3 owning Product/spec/API sources when that slice is reconciled; preserve Slice-2 closed/non-public 404. No repeat vote on D-04 and no current closed-detail activation.
3. **OI-4 — Owning-source acceptance and counterpart proof.** Named source owners must review/accept material Product/spec/API changes and required generated/bundle/type/fixture correspondence. WU-S2-005 acceptance is the prior baseline, not acceptance of this amendment. Orchestrator coordinates the route.
4. **OI-5 — Delivery plan refresh and runtime obligations.** Following source convergence, fresh backend WU-S2-003 and frontend WU-S2-004 Techplan Runs must consume accepted sources. Backend capacity/concurrency/database and frontend rendered disclosure/input evidence remain their delivery/testing scope. Work Graph owns dependency topology.

### Resolved — retained as decision history

1. ~~**OI-2 — Capacity close enum**~~ **RESOLVED —** Anhar accepts `funding_capacity_reached` as the distinct capacity close reason in the 2026-10-01 owner receipt. Propagate it to Campaign/API authoring; concrete source acceptance and review remain OI-4.
2. ~~**Rp1b individual cap direction**~~ **RESOLVED —** Hard cap Rp1.000.000.000; per-Campaign whole-IDR range Rp5.000–Rp1.000.000.000; default Rp1.000.000.000; Owner/Staff draft configuration and freeze after publication. Concrete request/response compatibility remains OI-1.
3. ~~**Guest disclosure direction**~~ **RESOLVED —** Pre-entry public-detail disclosure and independent POST recheck; exact field/error shape remains OI-1.
4. ~~**Reservation/closure behavior**~~ **RESOLVED —** Count settled plus accepted pending; close at capacity; do not reopen if a pending Donation later fails.
5. ~~**Public closure meaning**~~ **RESOLVED —** D-04 assigns identity retention, donation-action removal, and non-final Funding during accepted pending to Slice 3; does not change current Slice-2 closed detail behavior.
6. ~~**Historical five-brief mismatch**~~ **RESOLVED —** Current parent receipt plus EXP006 handoff own the clarified directions; historical pending/proposal status is not revived.
