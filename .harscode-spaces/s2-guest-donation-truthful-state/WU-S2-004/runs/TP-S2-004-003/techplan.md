# Tech Plan: Slice 2 Guest Donation Frontend Flow

> Phase             : Techplan
> Ticket            : WU-S2-004
> Author            : P-S2-004-TP-003-1 (Planner)
> Participant ID    : P-S2-004-TP-003-1
> Profile           : KC-PLANNER
> Role              : Planner
> Model             : gpt-6-luna
> Reasoning         : medium
> Created           : 2026-10-03
> Updated           : 2026-10-03
> Target revision   : bb69cd002b3f1a1056837affcd77bb2b001007b0 plus current working tree
> Workflow revision : pilot/orchestrator-v0.1@63ec4e0fd4f45a9820939ff8e568031236ce98f4 (ordinary applicable guidance current-effective)
> Status            : Approved
> Approach          : Frontend-only guest flow against accepted Campaign/Donation contracts and contract-faithful MSW; no runtime readiness claim.
> Refs              : predecessor `TP-S2-004-002/techplan.md` (SHA-256 `7b1a1f9ce7e68362681d6e2fc27ebc729243517bfca984f7984d1488dfba4535`); finding `RV-S2-004-001-F01`; current Product/MVP, Donation spec, split OpenAPI and WU-S2-006 accepted source/counterpart evidence.

---

## 1. Background

WU-S2-004 delivers the public guest Donation experience from an eligible Campaign through submission and status understanding. The prior Draft was synthesized before WU-S2-006 accepted the per-Campaign cap and capacity error contract. Current accepted sources now require a visible cap before amount entry, explicit IDR in all public detail states, and a generic `422 ValidationError` for an eligible capacity no-fit request. Generated OpenAPI types and public Campaign fixtures have converged to that source; no display/form behavior was changed by the counterpart Build.

## 2. Scope

**In scope:**
- Frontend production work only for Public Campaign Detail donation entry and guest donation submission/result/status surfaces, using current accepted Campaign and Donation contracts.
- Public disclosure of required `max_donation_amount` before amount entry; show its explicit IDR currency on every public detail state where it is available, including funding-unavailable state.
- QRIS-only sandbox interaction; other familiar methods remain visibly unavailable; truthful pending/terminal state; optional guest fields; safe guest status credential handoff; contract-faithful MSW and proportionate frontend verification.

**Out of scope (explicit):**
- Backend DTO/exact-wire response implementation, cap enforcement, eligibility/capacity predicates, idempotency persistence, simulator, settlement, credential controls, email verification/delivery, or database work. Backend response DTO and exact-wire assertion remain coupled to WU-S2-003.
- Changes to Product, spec, API, generated types, shared sources, tracker, or other Work Units. WU-S2-006 accepted the exact current source bytes; frontend Build consumes them.
- Public closed-Campaign detail, Slice 3 identity/result/accountability behavior, account history/claim, donor list, non-QRIS processing, real payment rails, or any implication of external settlement.
- Runtime/security acceptance, real backend integration, residual-risk acceptance, or milestone completion from mocks.

## 3. Requirements

| ID | Requirement | Source / evidence |
|---|---|---|
| Q1 | A guest may enter the Donation flow only from a currently available Campaign affordance; the server rechecks eligibility and the GET projection is a snapshot, not authorization. | Campaign feature 02, INV-campaign-14; `api/openapi/campaign.yaml`; WU-S2-005 accepted baseline. |
| Q2 | Display the Campaign's effective `max_donation_amount` before amount entry. It is an amount/currency object with required `currency_code: IDR`; preserve the explicit currency when Funding projection is unavailable. | Accepted Campaign feature 02 / INV-campaign-02 and Donation INV-donation-01; campaign OpenAPI `PublicCampaignDetail`, `MaxDonationAmount`; generated type and fixture. |
| Q3 | Amount UX accepts whole IDR from Rp5.000 upward, including Rp5.001, and does not invent a client-owned maximum/capacity rule. The accepted Campaign cap may be presented as context; server response remains authoritative. | Product MVP scope §5; Donation feature 01 / INV-donation-01; accepted Campaign contract. |
| Q4 | Eligible amount that exceeds the individual Campaign cap or does not fit remaining finite capacity is represented by generic field-level `422 ValidationError` on `amount`; capacity no-fit does not expose remaining capacity/close reason. Closed/ineligible behavior stays `409`; do not conflate these outcomes. | Donation feature 01 and INV-donation-01; donation OpenAPI POST; common `ValidationError`; accepted WU-S2-006 source receipt DEC-API-01/02. |
| Q5 | Submission uses the accepted `Idempotency-Key`; same key/payload retry denotes the original intent, changed payload reuse is a conflict, duplicate activation is prevented, and a new key follows an intentional new Donation action after failure. | Product MVP scope §5; Donation feature 01 / INV-donation-03; Donation OpenAPI. |
| Q6 | Only QRIS submits the sandbox request. Pending says “Menunggu hasil simulasi”; terminal copy uses “Hasil simulasi donasi: berhasil/gagal”; no ETA, real payment instruction, client-selected result, or automatic pending resubmission. | Product MVP scope §5; Donation feature 01 and INV-donation-02/04. |
| Q7 | Name is optional and not public by default. Email is optional and opt-in only for status notices; use exact accepted disclosure/helper copy and make no claim of verification or delivery from the frontend response alone. | Donation feature 01 exact Design direction; INV-donation-07/08; O3 boundary. |
| Q8 | Status link uses the returned bearer credential via URL fragment, frontend handoff, visible URL cleanup, and `X-Donation-Status-Credential`; only status is shown. Missing/wrong/expired credential and absent Donation share the generic public failure. | Donation feature 02; INV-donation-05/06; donation OpenAPI status endpoint. |
| Q9 | Errors distinguish local field validation from request/business failures, preserve input for recoverable failures, avoid internals/PII, and offer only contract-supported recovery. | Frontend architecture §§12, 18; `patterns.md` §§4, 12; accepted API contracts. |
| Q10 | Responsive and accessible presentation preserves task, cap/currency, trust/consequence information, status meaning, and reachable actions; color alone does not encode state. | Product design principles §§1–5, 8–9, 13–15; design guidelines §§21–24, 26; frontend AGENTS. |
| Q11 | MSW remains at the network boundary with generated-contract-shaped fixtures. Mock completion proves frontend behavior only; it does not establish backend integration, security controls, settlement, or runtime readiness. | WU-S2-004 manifest; frontend AGENTS §§2, 10–12; API README. |

## 4. Rules & Validation

- **R1** — Given Campaign detail returns `donation_action.available`, when the guest selects donation, then the UI enters the guest flow using the Campaign ID. Given generic unavailable action, then no donation activation is offered. The client does not derive eligibility from visibility, lifecycle, Funding, or cap.
- **R2** — Given a public Campaign detail, then its cap and explicit IDR are visible before amount entry. Given a valid detail with Funding unavailable, then the cap still retains and displays IDR. The UI does not label the individual donation cap as the Campaign `max_amount` threshold or remaining capacity.
- **R3** — Given whole-Rupiah input, then minimum Rp5.000 and Rp5.001 are locally accepted; fractional, malformed, or below-minimum values are field errors. No client-invented maximum is imposed; any displayed cap is the current detail snapshot, while POST validation is authoritative.
- **R4** — Given a POST `422` for amount, then show safe amount-level feedback without capacity or close-reason disclosure. Given closed/ineligible `409`, then show the distinct safe request-level outcome and do not relabel it as an amount error. Preserve input where recoverable.
- **R5** — Given a guest intent, then duplicate activation is suppressed and the same idempotency key/payload is retained for an ambiguous retry. Changed-payload reuse is surfaced as a conflict; after a known failed result, a deliberate new Donation uses a new key.
- **R6** — Given the donation screen, then display QRIS, GoPay, ShopeePay, and bank transfer. QRIS is active and interactive for the sandbox simulation; GoPay, ShopeePay, and bank transfer are visibly unavailable and non-interactive. The request submits only QRIS. Given a submission accepted, then show pending as “Menunggu hasil simulasi” and show the backend-provided result when later available. Pending never causes a timer promise, real-payment instruction, or automatic resubmission.
- **R7** — Given optional guest fields, then name may be omitted and is not presented as public. Email is sent only with explicit opt-in and exact accepted disclosure/helper; the UI does not claim ownership verification or notice delivery without evidence from an owning contract.
- **R8** — Given the status credential, then hand it from URL fragment to the status request header and clean the visible URL before continuing navigation/rendered steady state. Status view shows only status. Invalid/missing/expired/absent cases share generic visible failure copy.
- **R9** — Given local field or server/request failure, then render safe, distinct feedback, retain recoverable user input, and provide only an applicable retry or explicit new-intent action; never display stack traces, raw SQL, internal paths, secrets, or PII payloads.
- **R10** — Given representative desktop/mobile and keyboard use, then cap/currency, current task, consequence, state labels, and actions remain understandable and reachable; state meaning is not conveyed by color alone.
- **R11** — Given the frontend flow under MSW, then observable fixtures match generated request/response shapes and production data access still targets the real API. Passing mock/UI checks is reported only as frontend mock evidence.

## 5. Decision Log

| ID | Decision / option | Status | Rationale / consequence |
|---|---|---|---|
| D1 | Consume accepted `donation_action` and derive route navigation from Campaign ID; keep POST eligibility authoritative. | Chosen | Accepted by WU-S2-005; preserves frontend-only ownership and prevents stale GET state from becoming authorization. |
| D2 | Show required per-Campaign cap with explicit IDR before amount entry; keep it distinct from `max_amount` and capacity. | Chosen | Exact accepted WU-S2-006 source and generated/fixture counterpart. No remaining-capacity computation or disclosure. |
| D3 | Treat generic eligible-capacity no-fit `422` as amount/request validation, and closed/ineligible `409` as distinct request outcome. | Chosen | Accepted Donation OpenAPI preserves separate server predicates and generic behavior. Client must not infer the hidden capacity reason. |
| D4 | Route-local feature composition and narrow API functions; form state React Hook Form/Zod, server state appropriate to current TanStack Query precedent, ephemeral state local. | Chosen | Current clean-start frontend architecture; no global store or handwritten API mirror. Exact component split remains Build-level detail. |
| D5 | Keep status credential in fragment only for handoff, immediately clean visible URL, send credential in header, and avoid persistence/logging. | Chosen | Accepted interface and security boundary; UI behavior can be checked in browser, but does not prove infrastructure protections. |
| D6 | Do not add display/form behavior to WU-S2-006 counterpart scope; add cap presentation to this WU-S2-004 plan. | Chosen | BLD-S2-006-006 explicitly regenerated types/fixtures only and left display/form unchanged. |
| D7 | Rejected: infer remaining capacity or expose a new capacity-specific reason/copy. | Rejected | Product/API explicitly prohibit remaining-capacity disclosure; amount no-fit uses generic `422`. |
| D8 | Rejected: copy DTO/exact-wire backend response work into frontend Build. | Rejected | Backend response DTO and exact-wire assertion belong to WU-S2-003; keep frontend-only scope and Work Graph coupling. |
| D9 | Rejected: treat cap as global project monetary policy or as Campaign `max_amount`. | Rejected | Accepted specs define a per-Campaign individual donation cap distinct from fundraising threshold and shared monetary representation. |

## 6. Backward Compatibility

- No persisted data or migration is owned by this Work Unit.
- Consume required `PublicCampaignDetail.max_donation_amount` from the accepted source/generated type. Public fixture counterpart now includes the object in all fixture states, including Funding unavailable. Do not soften the generated contract or hand-edit generated types.
- Consume accepted POST `422` generic `ValidationError` and existing `409` closed/ineligible behavior without inventing a new response contract. Existing QRIS-only, status-only, fragment/header, and uniform `404` contracts remain in force.
- Backend response DTO/exact-wire implementation is a coupled WU-S2-003 dependency. API/spec/generated sources are inputs and remain untouched in frontend Build.

## 7. Edge Cases & Risks

| ID | Risk / edge case | Likelihood | Severity | Mitigation / accepted exposure |
|---|---|---:|---:|---|
| E1 | Missing, stale, or unavailable Campaign cap projection makes amount entry misleading or unusable. | Medium | High | Treat required contract absence as safe request failure, not a guessed default. Fixture/types prove shape only; WU-S2-003 exact-wire response remains dependency. |
| E2 | Donor interprets cap, target threshold, or remaining capacity as the same number. | Medium | High | Label effective per-donation cap in context, keep currency explicit, never calculate or disclose remaining capacity. |
| E3 | Capacity no-fit leaks capacity/close reason or is confused with closed Campaign. | Medium | High | Keep generic amount `422` separate from `409`; no client-side capacity inference. |
| E4 | Ambiguous request, double activation, or key rotation creates duplicate/incorrect intent. | Medium | High | Preserve same key/payload while ambiguous; new key only for deliberate new intent after known failure. Backend durability remains WU-S2-003. |
| E5 | Status bearer credential leaks through URL/history/referrer/log/cache or client persistence. | Low/Medium | Critical | Fragment handoff, visible URL cleanup, header-only lookup, no logging/persistence; browser evidence is frontend-only and does not establish server/infrastructure controls. |
| E6 | Distinct status lookup errors disclose Donation existence/credential validity. | Low/Medium | High | One generic visible failure; runtime status/body/header/cache/timing parity remains O4/O5. |
| E7 | Pending/success is read as real provider payment or settlement. | Medium | High | Exact simulation wording, no ETA/payment instruction; require Human rendered acceptance. |
| E8 | Email opt-in suggests verified ownership or sent notification. | Medium | High | Exact accepted disclosure/helper; only opt-in request fields; no unsupported verification/delivery claim. O3 remains open. |
| E9 | Mobile/assistive presentation hides cap, currency, method availability, consequence, or state. | Medium | Medium/High | Responsive and semantic coverage plus material Human rendered acceptance. |
| E10 | MSW result is mistaken for real backend/security/financial evidence. | Medium | High | Report milestone and evidence strictly as frontend mock verification; integration and Tier-1 runtime gates remain downstream. |

## 8. Interface Contract

**Persistence/data shape:** None authored by frontend. API-owned shapes use generated OpenAPI types. No handwritten mirror for cap or Donation request/response.

**API/event/external interface:**
- `GET /campaigns/{campaignId}` returns `PublicCampaignDetail`, including required `max_donation_amount: { amount, currency_code: "IDR" }`, `funding` tagged projection, and `donation_action` union. The cap is a point-in-time disclosure, not reservation or authorization.
- `POST /campaigns/{campaignId}/donations` uses `Idempotency-Key`, `SubmitDonationRequest`, and accepts QRIS only. An eligible amount no-fit response is generic shared `422 ValidationError` on `amount`; closed/ineligible remains `409`. Submission success is persisted `pending` and includes submission-only `status_token`.
- `GET /donations/{donationId}/status` sends the fragment-carried credential via `X-Donation-Status-Credential`; response is status only. Missing/wrong/expired/absent cases use the accepted public `404` contract.
- MSW represents these network contracts at the boundary. Production adapters continue to target real endpoints; no environment-specific mock response branch.

**Cross-layer/business boundary:** Frontend owns presentation, form lifecycle, field UX validation, navigation and safe request recovery. Backend owns cap/capacity validation, eligibility, idempotency, simulation outcome, persisted status, exact-once funding, email verification/delivery, credential verification/expiry and response parity. The client does not calculate remaining capacity or decide Campaign closure.

## 9. Architecture / Plan

1. Extend Campaign Detail success composition to present the accepted cap and currency before amount entry, using the generated public projection. Preserve explicit cap currency when Funding is unavailable. Keep `donation_action` as the sole entry affordance; unavailable stays non-actionable.
2. Route to a guest Donation composition using Campaign ID. Follow the current App Router layout and current Campaign/MSW initialization boundaries; derive exact route-local file locations from live conventions during Build.
3. Build amount and optional guest fields with React Hook Form + Zod. Use the displayed cap as truthful context; do not invent client policy for capacity or treat the detail snapshot as authority. Distinguish field-level `422`, request-level `409`, transport failure, and local validation.
4. Submit typed generated request with stable idempotency key for an intent. Suppress duplicate activation; ambiguous retries preserve key and payload; only explicit deliberate new intent after known failure gets a new key.
5. Render backend-owned pending/terminal state with exact approved simulation labels and only QRIS interactive. No timer, real payment instructions, or simulated client-selected outcome.
6. For status route, read credential from fragment, clean visible URL, send it in the designated header, and render status-only data or generic failure. Keep token out of persistent client state and logs.
7. Add/adjust contract-faithful MSW handlers/fixtures and observable tests for rules. Keep final rendered review, independent Testing, integration and backend/security evidence distinct.

## 10. Implementation Details

| Anchor | Why relevant | Intended change / precedent |
|---|---|---|
| `frontend/app/campaigns/[campaignId]/campaign-detail-view.tsx` — `CampaignSuccess` | Existing public detail/action composition. | Present `max_donation_amount` with explicit IDR before guest amount entry; preserve unavailable action behavior and current semantic owner. |
| `frontend/app/campaigns/[campaignId]/campaign-detail-client.tsx` — Campaign query/view composition | Current client query lifecycle. | Keep Campaign server state authoritative and pass typed projection to view; do not derive eligibility from unrelated fields. |
| `frontend/mocks/fixtures/public-campaign.ts` — public fixture variants | Already includes required cap object across fixture states after BLD-S2-006-006. | Reopen live fixture; add/adjust only behavior-specific cases if needed, retaining explicit cap/currency. |
| `frontend/lib/api/generated/openapi.ts` — Campaign/Donation `paths` and schemas | Generated contract correspondence. | Import generated types; never hand-edit. Current hash at planning dispatch: `288296d6e65a7500349126e066b3a4215915a647b0c46358f60e954d41262dc4`. |
| `frontend/lib/api/client.ts` — `apiRequest`; `frontend/lib/api/public-campaign.ts` — `getPublicCampaignDetail` | Current API adapter and Campaign precedent. | Extend narrowly for typed POST/header where necessary; generic safe failure handling, no mock branch. |
| `frontend/app/campaigns/[campaignId]/layout.tsx`; `frontend/mocks/handlers/public-campaign.ts`; `frontend/mocks/browser.ts` | Current mock initialization and Campaign network boundary. | Ensure new route requests are intercepted at MSW boundary before request; preserve real production target. |
| `frontend/app/campaigns/[campaignId]/campaign-detail-client.test.tsx`; `frontend/lib/api/public-campaign.test.ts` | Current observable coverage and fixture correspondence. | Update/add behavior assertions for cap disclosure, currencies, action states and errors; add focused tests for new flow at semantic owner. |
| `api/openapi/campaign.yaml` — `PublicCampaignDetail`, `MaxDonationAmount`; `api/openapi/donation.yaml` — submit/status operations; `api/openapi/common.yaml` — `ValidationError` | Accepted external contract authority. | Read-only inputs; current bytes match WU-S2-006 accepted hashes. |
| `docs/ui-ux/patterns.md` §§4, 7, 12–15; `design-guidelines.md` §§4, 13, 21–24, 26; `page-map.md` §1 | Form, status, money, error and public-surface presentation rules. | Derive ordinary route-local layout under Sunlit Editorial / Evidence-Led Optimism. |
| `docs/spec/4-campaign/features/02-campaign-detail-listing.md`; `docs/spec/5-donation/features/01-submit-donation-settlement.md`, `02-donation-status-check.md` | Accepted public projection and guest behavior. | Implement frontend responsibilities only; reopen live source and generated types during Build. |

Design readiness: **PARTIAL**. Product, state, wording, cap disclosure requirement and design system are established. Ordinary layout and amount-entry interaction can be derived from current patterns; no material Product or Brand decision is open. Rendered Human acceptance remains required after implementation.

## 11. Files Changed / Files NOT Changed

| File / area | Change type | Description |
|---|---|---|
| `frontend/app/campaigns/[campaignId]/` | Modify | Campaign cap/currency disclosure and accepted donation entry affordance. |
| `frontend/app/donations/` or current equivalent | Add | Guest form, result and status route composition, location derived from live routing convention. |
| `frontend/lib/api/` | Modify/Add | Narrow typed Donation submit/status adapters using generated types. |
| `frontend/lib/hooks/` | Add only if justified | TanStack Query owner for server state where useful; no mirrored global state. |
| `frontend/mocks/` | Modify/Add | Contract-faithful Campaign/Donation handlers and fixtures at network boundary. |
| Frontend observable tests/configured browser evidence | Modify/Add | Coverage corresponding to §4 and §12. |
| `frontend/components/ui/`, `frontend/components/shared/` | Avoid unless justified | Clean-start registry has no inherited broad contract; follow governance if a real shared abstraction is required. |

| File / area intentionally untouched | Why |
|---|---|
| `backend/` including WU-S2-003 response DTO/exact-wire work | Frontend-only ownership; response mapping remains coupled to WU-S2-003. |
| `api/`, `docs/product/`, `docs/spec/`, generated types | Accepted inputs; changing them requires owning-source review/coordination. |
| Donation ledger/transaction/crypto/auth and disbursement protected areas | Out of scope and root AGENTS Tier-0 fences apply. |
| Tracker, manifests, parent Events/Work Graph/Control Surface | Coordination records are not Planner Run outputs. |

## 12. Testing Checklist

| Rule | Verification / evidence | Primary owner | Why this is worth running / risk if skipped |
|---|---|---|---|
| R1 | RTL/MSW observable action-available/unavailable cases; assert only accepted union controls entry; request boundary retains POST authority. | Testing | Prevents unsupported entry or client eligibility claims. |
| R2 | RTL checks cap amount and IDR appear before amount control in regular and Funding-unavailable fixtures; check distinction from `max_amount`/remaining capacity. | Testing | Missing/ambiguous currency or cap meaning can cause material financial misunderstanding. |
| R3 | Table-driven whole-IDR cases including Rp4.999, Rp5.000, Rp5.001 and fractional/malformed values; assert no unsupported client upper bound and inspect request amount shape. | Testing | Prevents rejecting valid intent or implying capacity policy. |
| R4 | MSW/RTL covers generic amount `422` separately from closed/ineligible `409`, with no capacity/close reason disclosure and recoverable input preservation. | Testing | Confusing predicates can mislead donor and leak hidden capacity state. |
| R5 | Observable same-key/same-payload retry, changed payload conflict, duplicate activation, ambiguous outcome and deliberate new intent after known failure. | Testing | Duplicate financial intent is high severity; button-only testing misses key lifecycle. |
| R6 | RTL checks that QRIS, GoPay, ShopeePay, and bank transfer are all displayed; QRIS is active/interactive, while GoPay, ShopeePay, and bank transfer are visibly unavailable and non-interactive; the request remains QRIS-only. Check exact pending/terminal labels and absence of timer/payment instruction/auto-submit. Human rendered acceptance judges simulation comprehension. | Testing / Human | Assertions prevent omitted or accidentally actionable display-only methods and protect wording; Human judges whether hierarchy implies real settlement. |
| R7 | RTL checks omitted optional fields and opt-in payload plus exact label/helper and no success claim for verification/delivery. | Testing | Prevents accidental PII/notification opt-in or unsupported fulfillment promise; O3 controls remain external. |
| R8 | Focused real-browser check for fragment handoff, visible URL cleanup, header request and status route; RTL checks status-only view and uniform generic failure cases. | Testing | jsdom cannot establish browser address/history behavior; token leakage risk is critical. This is frontend-only evidence. |
| R9 | RTL covers field/request distinction, recoverable input preservation, safe error text and applicable retry behavior. | Testing | Avoids lost intent and leakage of internals/PII. |
| R10 | Human rendered acceptance on representative desktop/mobile and keyboard path; automated semantic role/name/focus assertions for repeatable behavior. | Human / Testing | Unit tests do not prove hierarchy, clipping, responsive reachability or comprehension. |
| R11 | Focused API/component tests assert MSW fixture correspondence to generated types and no production mock branch; independent Testing reports exact scope. | Testing | Confirms frontend mock behavior without implying backend/runtime correctness. |

For browser automation, a narrowly scoped Playwright check is justified for R8 because fragment/address behavior and browser navigation are not credibly established by component tests; Testing owns authoritative execution. Risk if skipped is an unnoticed bearer credential left in visible URL/history or omitted from the intended header. Browser evidence does not prove referrer, logs, cache, key/expiry, abuse, or server parity controls. Later commands must come from live frontend scripts; no test/validator/generator/runtime command is authorized in this synthesis Run.

### Test Focus Pointer

| Area | Why sensitive | Evidence anchor from Exploration | Still relevant post-synthesis? |
|---|---|---|---|
| Guest bearer status credential fragment, browser URL cleanup and status-only access | Credential leaks expose private Donation status; mocks cannot prove infrastructure exposure controls or server parity. | `WU-S2-004/runs/EXP-S2-004-001/evidence/stage-2-gap-analysis.md#area-2--authored-api-contract`; `#area-4--frontend-architecture-and-live-implementation`; `stage-3-solutioning.md#guest-status-credential-and-proof-boundary` | Yes — R8 browser evidence; O4/O5 server/security proof stays downstream. |
| Optional guest email and PII/verification/terminal notice boundary | Email is sensitive and must not imply ownership verification or delivery. | `WU-S2-004/runs/EXP-S2-004-001/evidence/stage-2-gap-analysis.md#area-1--product-donation-domain-and-approved-delivery-boundary`; `stage-3-solutioning.md#optional-email` | Yes — R7 frontend behavior; O3 controls/evidence remain downstream. |
| Idempotency, stale Campaign affordance, cap and capacity no-fit | Duplicate intent or hidden predicate confusion has financial and trust impact; backend ordering/concurrency is not frontend proof. | `WU-S2-004/runs/EXP-S2-004-001/evidence/stage-2-gap-analysis.md#area-1--product-donation-domain-and-approved-delivery-boundary`; `#area-4--frontend-architecture-and-live-implementation`; refreshed source delta from WU-S2-006 accepted sources | Yes — R1/R3/R4/R5 frontend observables; backend concurrency/exact-wire under WU-S2-003. |

## 13. Open Items

### Active — needs external input or verification

1. **O2 — Simulator timing/scenario mechanics:** frontend displays backend-returned state only; no timing promise or client-selected outcome. Backend implementation/evidence remains outside this Work Unit and does not block frontend mock flow unless a required response field/interaction changes.
2. **O3 — Optional email verification and delivery controls:** exact accepted disclosure/request opt-in are known; verification interaction, bounded terminalization, retry, retention/deletion-race controls and residual-risk evidence remain unresolved with owning backend/Security/PII path. No end-to-end email fulfillment claim.
3. **O4/O5 — Status credential and anti-enumeration runtime evidence:** frontend handoff/cleanup and generic UI require frontend browser/mock evidence. Credential generation/strength, key/comparison, expiry enforcement, browser/infrastructure exposure, abuse controls, response/body/header/cache/timing parity and residual-risk evidence remain downstream Security/API/Testing work. No residual risk is accepted.
4. **O6 — WU-S2-003 exact-wire public Campaign response:** generated types and fixture now carry the required cap shape; backend `publicCampaignDetailResponse` mapping and exact-wire assertion remain assigned to WU-S2-003. Integration evidence depends on that work; this plan does not move it into frontend ownership.
5. **O7 — Rendered Human acceptance:** after implementation, a Human must inspect and exercise representative responsive guest flow, cap/currency hierarchy, simulation/status truth and accessible interaction. This is a delivery gate, not a planning blocker.

### Resolved — retained as decision history

1. ~~**Campaign action contract gap from Exploration**~~ **RESOLVED —** WU-S2-005 accepted the `donation_action` union; WU-S2-006 accepted current monetary/capacity source bytes; generated types/fixtures now correspond. GET remains a snapshot and POST rechecks.
2. ~~**O1-REP cap/capacity product/API direction**~~ **RESOLVED —** Human decisions were reconciled into the seven accepted Campaign/Donation source files. Per-Campaign cap, public disclosure, IDR encoding, capacity no-fit generic `422`, and distinct closed/ineligible `409` are current source contracts. Exact source acceptance is recorded in WU-S2-006 `RV-S2-006-006/invocation.md` and manifest; generated/frontend fixture correspondence is evidenced by `BLD-S2-006-006/report.md`.
3. ~~**Design readiness**~~ **RESOLVED —** PARTIAL. Product meaning, page purpose and visual/interaction principles are established; routine composition is derivable. No material new design decision or high-fidelity Figma gate is evidenced. Human rendered acceptance remains a later verification requirement.
4. ~~**Backend implementation placement**~~ **RESOLVED —** Campaign public response DTO/exact-wire implementation is coupled to WU-S2-003; WU-S2-004 remains frontend-only.
