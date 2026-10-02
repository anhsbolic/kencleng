# Tech Plan: Slice 2 Campaign Donation Entry Contract Reconciliation

> Phase             : Techplan
> Ticket            : WU-S2-005
> Work Unit         : WU-S2-005
> Run               : TP-S2-005-001
> Author            : P-S2-005-TP-001-1 (Planner)
> Participant ID    : P-S2-005-TP-001-1
> Profile           : KC-PLANNER
> Role              : Planner
> Model             : not exposed (Invocation configured `gpt-6-luna`)
> Reasoning         : not exposed (Invocation configured `high`)
> Session           : not exposed
> Created           : 2026-10-01
> Updated           : 2026-10-01
> Target revision   : `7fd8b473b239b20bda3990ab29c51440d321a796` plus current working tree (Invocation)
> Workflow revision : `pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8`
> Status            : Draft / In Review
> Approach          : Reconcile Campaign acceptance and authored/generated API surfaces around an availability-only, backend-authored GET snapshot while preserving POST eligibility authority.
> Refs              : `WU-S2-005/manifest.md`; `EXP-S2-005-001/evidence/stage-2-gap-analysis.md`; `EXP-S2-005-001/evidence/stage-3-solutioning.md`; parent `events.md` decision “EXP-S2-005-001 inspected; owner selected availability-only”; Product Slice 2; current Campaign and Donation authorities.

---

## 1. Background

The approved Slice 2 outcome requires a visitor to enter guest Donation from an eligible Public Campaign Detail. The reconciled Slice 1 Campaign feature, invariant, authored OpenAPI, generated types, fixture, backend projection, and frontend view still describe the donation action as unavailable. The Campaign/API owner has selected an availability-only response direction: Campaign GET evaluates eligibility at read time, the frontend chooses its route from Campaign ID, and Donation POST checks eligibility again. This Run defines the reconciliation work and the proposed contract for its Human/review gates; it does not claim that the spec, API, generated artifacts, or runtime behavior have already changed.

A GET response is a snapshot. `Cache-Control: private, no-store` limits intermediary storage but cannot revoke a response already held by a client. The Donation submission contract and D1 ordering remain authoritative when eligibility changes after GET.

## 2. Scope

**In scope:**

- Reconcile Slice-2 Campaign Detail donation-action acceptance in `docs/spec/4-campaign/features/02-campaign-detail-listing.md` and the applicable public-boundary rule in `docs/spec/4-campaign/invariants.md`.
- Reconcile the authored `GET /campaigns/{campaignId}` response in `api/openapi/campaign.yaml` to the owner-selected availability-only direction.
- Regenerate and reconcile the derived API bundle, generated TypeScript, and contract fixture according to `api/README.md`.
- Preserve Campaign's public-safe projection, existing public not-found/anti-enumeration behavior, no-store response policy, and backend ownership of eligibility.
- Record the GET/POST freshness boundary and keep Donation submission rejection behavior authoritative.

**Out of scope (explicit):**

- Product/MVP changes; Donation submission/D1 semantics; Slice 3 closure, persistent closed-Campaign public identity, or public-result behavior.
- Backend Campaign GET producer or Donation POST implementation, frontend route/UI implementation, migrations, runtime proof, and delivery testing. Those remain in their delivery Work Units; Campaign GET producer work in WU-S2-003 waits for this contract's owner acceptance.
- Adding an action target/URI or coupling Campaign API to a frontend route.
- Changing unrelated Campaign operations, generated API sources by hand, or shared `common.yaml` components unless reconciliation evidence proves a necessary shared contract change.
- Claiming `CONTRACT_READY`, WU-S2-005 completion, residual-risk acceptance, or authorization to begin dependent Build work.

## 3. Requirements

| ID | Requirement | Source / evidence |
|---|---|---|
| Q1 | A visitor must be able to enter the real guest Donation flow from an eligible Public Campaign Detail without creating an Account. | `docs/product/mvp-delivery-slices.md` §5; `docs/product/mvp-scope.md` §§4–5; Stage-2 Gap Analysis, Area 1 |
| Q2 | The Campaign action response is an availability projection only. It contains no navigation target; the frontend derives its route from Campaign ID. | Owner decision in parent `events.md`, 2026-10-01; Stage-3 Solutioning, “Decision framing — Semantik `donation_action`” |
| Q3 | `available` means the backend judged Donation eligible when it produced that GET response; it is not a promise that a later POST will succeed. Exact wire shape and unavailable-reason vocabulary require the Campaign/API owner's approval. | Owner decision in parent `events.md`; `WU-S2-005/manifest.md`; Stage-3 Solutioning, “Dampak yang harus dibawa ke rekonsiliasi” |
| Q4 | Backend Donation submission remains the current eligibility authority and can reject a now-ineligible Campaign; preserve D1 ordering and do not transfer authority to frontend visibility. | `docs/spec/4-campaign/invariants.md#inv-campaign-13`; `docs/spec/5-donation/invariants.md#inv-donation-02`; `api/openapi/donation.yaml` submit operation; `WU-S2-005/manifest.md` |
| Q5 | Public detail remains a closed public-safe projection. Existing anti-enumeration and `Cache-Control: private, no-store` requirements continue to apply. | `docs/spec/4-campaign/invariants.md#inv-campaign-14`; Campaign feature `02-campaign-detail-listing.md`; Stage-2 Gap Analysis, Area 2 |
| Q6 | The authored Campaign OpenAPI remains the source; the aggregate spec, generated TypeScript, and fixtures must correspond to it. | `api/README.md`; `api/openapi/campaign.yaml`; Stage-2 Gap Analysis, Area 3 |
| Q7 | Reconciliation changes remain limited to the Campaign donation-entry contract and owning acceptance criteria; no Product, Donation submission, or Slice 3 behavior is redefined. | `WU-S2-005/manifest.md`; approved Product Slice 2; owner decision event |

## 4. Rules & Validation

- **R1 — Availability-only action shape (proposed for approval):** `PublicCampaignDetail.donation_action` remains required and exposes a finite `availability` value of `available` or `unavailable`; it contains no link, URI, route, or activation target. Proposed shape: `reason` is absent when available and required from a finite public-safe vocabulary when unavailable. The final conditional shape and exact reason vocabulary are an Active owner decision, not yet accepted contract.
- **R2 — Snapshot semantics:** When the response says `available`, it records the backend's eligibility assessment at GET time only. The frontend may use it to present the entry route based on the Campaign ID; it may not use it as authority to accept a Donation. A later POST must reevaluate current eligibility and may reject a stale GET result.
- **R3 — Unavailable truth:** When the backend assesses that Donation is unavailable at GET time, the response expresses that state without a target and without exposing raw internal status, close reason, operational metadata, or other non-public detail. The final finite `reason` vocabulary and which public Campaign states can produce it must be settled in the contract gate before spec/API acceptance.
- **R4 — Public visibility boundary:** Absent, malformed/non-resolvable, and non-public Campaigns retain the same `PublicCampaignNotFound` `404` response regardless of optional Authorization. This reconciliation does not make a non-public Campaign public or redefine closed-Campaign visibility.
- **R5 — Cache and projection:** Public detail remains the standalone allowlisted projection, and successful/error responses retain `Cache-Control: private, no-store` as required by the current Campaign public boundary.
- **R6 — Submission authority:** Donation POST independently checks eligibility against current backend state. A stale `available` GET response never authorizes submission and cannot weaken INV-campaign-13 / INV-donation-02 D1 ordering.
- **R7 — Authored/generated correspondence:** The authored split Campaign API, bundled `api/openapi.yaml`, generated `frontend/lib/api/generated/openapi.ts`, and `frontend/mocks/fixtures/public-campaign.ts` describe the same accepted action contract; generated files are regenerated, not hand-edited.
- **R8 — Scope fidelity:** The reconciliation changes only Campaign Detail action acceptance/invariant/API contract and necessary generated counterparts. It does not amend Product/MVP, Donation submission/D1, broader Slice 3 closure/result behavior, or production implementation.

## 5. Decision Log

| ID | Decision / option | Status | Rationale / consequence |
|---|---|---|---|
| D1 | Availability-only Campaign response; frontend chooses route from Campaign ID; Donation POST reevaluates eligibility. | Chosen — Campaign/API owner, Anhar Solehudin, 2026-10-01 | Recorded in parent `events.md`. No backend-supplied route; GET is a staleable snapshot. Exact schema and reason values remain open. |
| D1-alt | Availability plus backend-provided action target/route. | Rejected — Campaign/API owner direction | Would bind the Campaign response to consumer navigation without eliminating the GET/POST race. Do not reopen absent new authority evidence. |
| D2 | Keep Campaign public visibility/projection/no-store and Donation D1 as separate constraints around the action field. | Chosen — existing approved authority preserved | INV-campaign-14 continues to own public projection/visibility; INV-campaign-13 and INV-donation-02 continue to own the close/submission ordering boundary. This Work Unit does not reconcile Slice 3 public closure behavior. |
| D3 | Treat generated bundle/types/fixture as derived consumers of authored `api/openapi/campaign.yaml`. | Chosen — API source convention | `api/README.md` makes split domain files authored sources and prescribes bundling/type generation; incomplete regeneration would leave consumers inconsistent. |

## 6. Backward Compatibility

- The response already requires `donation_action` with `availability: unavailable` and `reason: donation_flow_not_available`. The proposed contract retains the field and finite state model while adding `available` and revisiting conditional reason presence/vocabulary. Clients that validate enums or require `reason` may need regeneration/adaptation; this is a material authored contract change and must be reviewed/owner-accepted before it is treated as current.
- No endpoint, request, authentication policy, HTTP status, or Donation submission contract is intended to change.
- Current non-public Campaign behavior, including the existing Slice 1 visibility predicate, remains in force. This plan does not silently claim post-closure public detail support.
- The frontend generated types and MSW fixture must follow the accepted authored schema. Do not hand-edit generated files or the aggregate bundle.

## 7. Edge Cases & Risks

| ID | Risk / edge case | Likelihood | Severity | Mitigation / accepted exposure |
|---|---|---:|---:|---|
| RISK-1 | Campaign changes eligibility after GET but before Donation POST. | Medium | High | State the snapshot meaning; POST rechecks eligibility and may reject. Keep user-facing behavior for a rejection with the Donation delivery owner; this plan does not claim that implementation exists. |
| RISK-2 | Adding availability or reason values exposes internal lifecycle/close reasons or widens the public projection. | Medium | High | Use only an explicit public-safe response schema; keep the existing allowlist and anti-enumeration behavior. Owner must accept exact values before contract finalization. |
| RISK-3 | Only some API/generated/fixture surfaces are updated, causing API producer and frontend consumer disagreement. | Medium | Medium | Regenerate bundle/types and update fixture in the same reconciliation delivery; verify source-to-generated correspondence. |
| RISK-4 | A strict consumer treats the new enum/conditional reason as incompatible or continues to assume a required reason on `available`. | Medium | Medium | Review the exact schema at the owner gate; regenerate types; check known in-repository consumers and record compatibility limits. Do not claim compatibility for unknown external consumers. |
| RISK-5 | Slice 1 currently treats closed Campaign states as non-public, while Slice 3 later requires persistent public Campaign identity. | Medium | High | Keep closed-Campaign visibility and Slice 3 result semantics out of this reconciliation; record the scope boundary and route any required change to its owning slice. |

## 8. Interface Contract

**Persistence/data shape:** No persistence or migration change. The existing Campaign donation-action projection is a public response field; its accepted state/reason schema is the contract change.

**API/event/external interface:** Keep `GET /campaigns/{campaignId}` and its `PublicCampaignDetail` response. Proposed `donation_action` shape has `availability: available | unavailable`, no navigation target, and a conditional public-safe `reason` only for unavailable. The exact `reason` enum, conditional requiredness, and Campaign states represented by unavailable remain an approval-gated Open Item. Existing public `404`/`503` and `Cache-Control: private, no-store` contract remains applicable.

**Cross-layer/business boundary:** Backend Campaign GET authors an eligibility snapshot. Frontend chooses local navigation from Campaign ID and treats availability as presentation/entry guidance. Donation POST remains the authority that checks current eligibility. If state changes after GET, the POST may reject. No request field, submission contract, or D1 invariant changes.

## 9. Architecture / Plan

1. At the Human Techplan gate, settle the proposed wire shape and the finite public-safe unavailable-reason vocabulary, including which states may return public detail with `unavailable`. Keep current Slice 1 non-public behavior and Slice 3 closure semantics separate.
2. Reconcile Campaign feature acceptance and `INV-campaign-14` only where needed to express the approved Slice-2 entry action and preserve existing public safety requirements. Do not alter `INV-campaign-13` or Donation `INV-donation-02`.
3. Update the authored `api/openapi/campaign.yaml` schema/operation description to match the accepted Campaign behavior; use shared `common.yaml` components only if a genuinely shared contract is required. Preserve existing public operation auth, response, anti-enumeration, and cache behavior.
4. Regenerate `api/openapi.yaml` and `frontend/lib/api/generated/openapi.ts` using the documented commands, then update the Campaign contract fixture to match the generated response type.
5. Run source validation/generation checks appropriate to the authored/generated change and inspect source-to-bundle/type/fixture correspondence. Independent Testing and Human owner acceptance remain separate gates; this Techplan does not claim them.
6. After contract acceptance, route the backend Campaign GET producer to its authorized delivery scope and unblock the dependent frontend Techplan/Build Work Unit. Keep backend and frontend production writes separate.

No independent migration, script, cron, or runbook lifecycle is introduced.

## 10. Implementation Details

| Anchor | Why relevant | Intended change / precedent |
|---|---|---|
| `docs/spec/4-campaign/features/02-campaign-detail-listing.md` — `Detail`, acceptance criteria 3 and 6 | Current reconciled Slice 1 behavior requires unavailable action | Adapt Campaign-owned acceptance for Slice 2 while preserving public-safe detail and ensuring an active action is backed by the real guest flow. |
| `docs/spec/4-campaign/invariants.md` — `INV-campaign-14` | Owns public eligibility, anti-enumeration, allowlist projection, action meaning, and no-store | Amend only the public-boundary statements needed for the accepted action contract; do not broaden closed-Campaign visibility in this Work Unit. |
| `docs/spec/4-campaign/invariants.md` — `INV-campaign-13` | Owns Campaign close/D1 ordering | Reference as an unchanged boundary; do not alter mechanism or submission ordering. |
| `api/openapi/campaign.yaml` — `getPublicCampaignDetail`, `PublicCampaignDetail`, `PublicCampaignDonationAction` | Authored Campaign response source; current schema permits unavailable only | Reconcile description/schema after owner approval; retain public operation and response boundary. |
| `api/README.md` — “Editing workflow” | Defines authored/generated OpenAPI workflow | Validate split source, bundle generated aggregate, and regenerate frontend types; never hand-edit generated outputs. |
| `backend/internal/domain/campaign/repository_db.go` — `RepositoryDB.FindPublicDetail` | Current repository filters public detail and selects an explicit projection | Backend producer delivery later reopens this anchor to assess the eligible detail source; this reconciliation does not edit it. |
| `backend/internal/domain/campaign/service.go` — `Service.GetPublicDetail` / `toPublicDetail` | Current mapper assigns the unavailable action | Producer delivery later maps accepted GET-time availability; no implementation is authorized by this Techplan. |
| `backend/internal/transport/http/campaign_public.go` — `PublicCampaignDetailHandler` / `toPublicCampaignDetailResponse` | Current response DTO serializes availability/reason | Producer delivery later maps the accepted contract without widening the closed projection. |
| `frontend/app/campaigns/[campaignId]/campaign-detail-view.tsx` — `CampaignSuccess` | Current detail view does not use action availability to enter Donation | Frontend delivery later derives local route from Campaign ID and handles submission-time rejection under its own approved plan. |
| `frontend/lib/api/generated/openapi.ts` — `PublicCampaignDonationAction` | Generated TypeScript consumer contract | Regenerate from bundled authored sources after schema acceptance; do not hand-edit. |
| `frontend/mocks/fixtures/public-campaign.ts` — `donation_action` | Current contract-faithful response fixture | Update to exercise the accepted availability states and match generated types. |

## 11. Files Changed / Files NOT Changed

| File / area | Change type | Description |
|---|---|---|
| `docs/spec/4-campaign/features/02-campaign-detail-listing.md` | Adapt | Campaign Detail acceptance for eligible guest Donation entry. |
| `docs/spec/4-campaign/invariants.md` | Adapt, if needed | Reconcile `INV-campaign-14` public action behavior; keep INV-campaign-13 unchanged. |
| `api/openapi/campaign.yaml` | Adapt | Authored response schema/operation description after owner gate. |
| `api/openapi.yaml` | Regenerate | Derived bundled OpenAPI contract. |
| `frontend/lib/api/generated/openapi.ts` | Regenerate | Derived API TypeScript types. |
| `frontend/mocks/fixtures/public-campaign.ts` | Adapt | Contract-faithful fixture values. |
| `api/openapi/common.yaml` and `api/openapi/index.yaml` | Leave unchanged unless required | No shared component or path change is currently indicated. |
| Backend/frontend production implementation, migrations, tests, Product/MVP, Donation spec/API, `INV-campaign-13`, and Slice 3 behavior | Leave unchanged in this Work Unit | Explicit scope and authority boundaries; implementation is in separate delivery Work Units. |

## 12. Testing Checklist

| Rule | Verification / evidence | Primary owner | Why this is worth running / risk if skipped |
|---|---|---|---|
| R1 | Human/API owner reviews the exact final OpenAPI shape, finite availability values, conditional reason behavior, absence of a target, and reason vocabulary; validate authored spec after acceptance. | Human (contract meaning); Testing (source validation) | This is a material public contract decision. A guessed or invalid schema can misdirect both producer and consumer; exact vocabulary is still open. |
| R2 | Contract scenario review: GET reports available, Campaign becomes ineligible, POST reevaluates and may reject; later Donation implementation tests the observable stale-snapshot case. | Testing | Confirms contract language preserves server authority across the GET/POST boundary. Skipping it risks frontend treating a stale snapshot as authorization. |
| R3 | Validate each final unavailable reason maps only to an approved public-safe condition; inspect schema/acceptance for absence of raw internal status, close reason, or operational fields. | Testing | Public enum/reason leakage could expose internal lifecycle semantics or create false claims. |
| R4 | Contract review verifies same public `404` shape for absent/malformed/non-public Campaigns and no Authorization-dependent visibility; runtime parity remains in the owning backend delivery/testing evidence. | Testing | Protects anti-enumeration and prevents this action change from widening visibility. Runtime proof is not part of this reconciliation Run. |
| R5 | Validate public response/header contract retains closed projection and `Cache-Control: private, no-store` for success and public errors. | Testing | A response/header regression can leak internal data or allow stale visibility after withdrawal. |
| R6 | Trace Campaign acceptance/API text against `INV-campaign-13`, Donation `INV-donation-02`, and Donation POST contract; downstream Donation Testing owns runtime close/order proof. | Testing | Prevents action visibility from becoming a second eligibility authority or weakening submission/close ordering. |
| R7 | Run the repository's authored OpenAPI validation and bundle/type generation; compare generated `PublicCampaignDonationAction` and fixture against the authored schema. | Build (generation loop); Testing (independent final correspondence) | Generated drift can make typed consumers compile against a different response than the authored API. |
| R8 | Review changed-file set and diff against WU manifest boundaries and accepted Product/D1/Slice 3 sources; record Human approval for protected spec/API changes at the owning gate. | Human | Scope/authority drift can silently change Product or a separate delivery contract. |

### Test Focus Pointer

| Area | Why sensitive | Evidence anchor from Exploration | Still relevant post-synthesis? |
|---|---|---|---|
| Public projection, non-public Campaign anti-enumeration, and `no-store` behavior | Changing an action enum/reason must not widen public disclosure or distinguish missing/non-public resources; runtime timing parity is security-sensitive. | `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/EXP-S2-005-001/evidence/stage-2-gap-analysis.md#area-2-campaign-acceptance-criteria-dan-invariants`; `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/EXP-S2-005-001/evidence/stage-2-gap-analysis.md#area-3-authored-dan-generated-api-contract-surfaces` | Yes — preserve the specialized public-boundary checks; full runtime/timing proof stays with backend Testing. |
| Stale GET availability versus later Donation submission | Eligibility can change between detail read and POST; a client-held snapshot is not revocable and cannot authorize submit. | `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/EXP-S2-005-001/evidence/stage-2-gap-analysis.md#area-4-perilaku-backendfrontend-yang-terkait`; `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/EXP-S2-005-001/evidence/stage-3-solutioning.md#batas-keputusan-dan-risiko` | Yes — carry the contract-level scenario; runtime rejection/ordering proof belongs to the Donation/Campaign delivery and Testing owners. |
| Concurrent close/submission/settlement ordering (D1) | Concurrency-sensitive money/lifecycle invariant, but this Work Unit does not change its rule or implementation. | `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/EXP-S2-005-001/evidence/stage-2-gap-analysis.md#area-2-campaign-acceptance-criteria-dan-invariants`; `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-005/runs/EXP-S2-005-001/evidence/stage-2-gap-analysis.md#area-4-perilaku-backendfrontend-yang-terkait` | N/A for WU-S2-005 specialized execution — D1 is unchanged and owned by WU-S2-003/approved Donation planning; retain its evidence obligation there (D2 and Q4). |

## 13. Open Items

### Active — needs external input or verification

1. **Owner approval of exact `PublicCampaignDonationAction` schema:** confirm whether the proposed `availability: available | unavailable` enum and conditional `reason` presence are the accepted shape, and approve the finite unavailable `reason` vocabulary. Owner: Anhar Solehudin, Campaign/API owner. This blocks final contract/spec acceptance.
2. **Unavailable state domain:** determine which public Campaign conditions can return `unavailable` while the detail itself remains public, without changing Slice 1's current non-public predicate or pulling Slice 3 closed-Campaign visibility into this Work Unit. Owner: Campaign/API owner; route a missing Product decision only if the existing approved Slice 2/3 boundary cannot answer it.
3. **Protected authority reconciliation:** approve any material change to Campaign acceptance or `INV-campaign-14` through the applicable Campaign/spec owner gate; keep `INV-campaign-13` and Donation `INV-donation-02` unchanged. Owner: Anhar Solehudin, Campaign/API owner.
4. **Independent review and final owner acceptance:** independent review is recommended before report/gate convergence because the response crosses Campaign domain acceptance, authored API, generated consumers, and the GET/POST authority boundary. Resolve review findings, then obtain final Campaign/API contract acceptance and reconcile all required generated counterparts. These gates are not satisfied by this Draft.

### Resolved — retained as decision history

1. ~~**Action navigation ownership**~~ **RESOLVED — availability-only; frontend derives the route from Campaign ID.** Anhar Solehudin selected this as Campaign/API owner on 2026-10-01; backend-provided target/route is rejected (D1).
2. ~~**Submission authority**~~ **RESOLVED — POST checks current eligibility again.** A GET availability result is a snapshot and can be stale; the owner decision does not amend Donation submission or D1 semantics.
