# Tech Plan: Slice 2 Guest Donation Frontend Flow

> Phase             : Techplan
> Ticket            : WU-S2-004
> Author            : P-S2-004-TP-001-1
> Participant ID    : P-S2-004-TP-001-1
> Profile           : KC-PLANNER
> Role              : Planner
> Model             : gpt-6-luna
> Reasoning          : medium
> Created           : 2026-10-01
> Target revision   : 7fd8b473b239b20bda3990ab29c51440d321a796 + accepted current working tree
> Workflow revision : pilot/orchestrator-v0.1@95ecf37ba8ae449a5b3b278c27331aca87360bc8
> Status            : Draft / In Review
> Approach          : Implement the guest path against accepted Donation and Campaign contracts with MSW at the network boundary; preserve the `FRONTEND_MOCK_VERIFIED` evidence boundary.
> Refs              : `WU-S2-004/manifest.md`; `runs/EXP-S2-004-001/evidence/{stage-2-gap-analysis.md,stage-3-solutioning.md}`; accepted WU-S2-002 and WU-S2-005 artifacts; current Product, Donation, Campaign, UI/UX, and frontend architecture sources listed below.

---

## 1. Background

The clean-start frontend currently has a public Campaign Detail route and Campaign GET mock path, but no guest Donation form, submission/status adapter, status route, or Donation MSW handlers. Slice 2 requires a guest to move from an eligible public Campaign through a sandbox Donation and understand its persisted status without an Account. The accepted Donation contract and the Campaign action contract provide the current network shapes. Completion is `FRONTEND_MOCK_VERIFIED`; it does not prove backend runtime, security, settlement, email delivery, or real integration.

The live Campaign contract has since been reconciled and accepted through WU-S2-005. The earlier Exploration mismatch is therefore historical, not a current blocker: `donation_action` reports point-in-time backend eligibility, and Donation POST independently rechecks it. A separate backend planning Run has since surfaced a monetary/configuration proposal that still needs owning Product/API source reconciliation; it must not be silently added to the accepted interfaces.

## 2. Scope

**In scope:**
- Update the current public Campaign Detail action to follow accepted `donation_action` availability and navigate to the guest Donation flow only when the API says available.
- Add guest Donation amount/method/optional-field form, submission, truthful pending/terminal result, and temporary status revisit using accepted generated OpenAPI types and Donation contract.
- Keep QRIS as the only simulated method; show GoPay, ShopeePay, and bank transfer as unavailable and non-interactive.
- Implement frontend status credential fragment handoff and visible URL cleanup, status-only rendering, and generic public link failure behavior.
- Use contract-faithful MSW handlers/fixtures at the browser network boundary for scoped mock verification; production API code continues to call the actual contract endpoints.
- Preserve source-first simulation wording, responsive/accessibility behavior, and no public donor PII.
- Plan focused automated, real-browser, and Human rendered evidence needed to reach `FRONTEND_MOCK_VERIFIED`.

**Out of scope (explicit):**
- Any write under `backend/`, backend integration, server eligibility/security proof, D1 concurrency/financial evidence, or real settlement/payment rails.
- Changing Product, domain specs, OpenAPI, generated types, fixtures outside frontend-owned mock fixtures, or accepted Campaign/Donation contracts from this frontend Build.
- Treating MSW as proof of persistence, settlement, verification, secure backend token handling, anti-enumeration parity, email delivery, or real-world evidence.
- Account requirement, donation claim/history, donor list, public guest name/email, campaign-wide email updates, broader Slice 3 closure/result/accountability behavior.
- Implementing an optional-email verification workflow or making a delivery claim without its supported contract/owner evidence; the form may carry the already accepted optional fields and disclosure.
- Choosing or imposing a per-Campaign donation limit, capacity behavior/reason, configuration range, or new Campaign API field from the backend Draft candidate.
- Design-system-wide component extraction absent a demonstrated shared contract.

## 3. Requirements

| ID | Requirement | Source / evidence |
|---|---|---|
| Q1 | Deliver the guest Campaign Detail → Donation → status loop without an Account and within the `FRONTEND_MOCK_VERIFIED` milestone. | `docs/product/mvp-delivery-slices.md` §5; WU-S2-004 `manifest.md`; accepted Campaign action contract WU-S2-005. |
| Q2 | Use the Donation API contract for submission, idempotency, status lookup, and errors; use generated types rather than parallel handwritten contract models. | `api/README.md`; `api/openapi/donation.yaml` + referenced `common.yaml`; `frontend/AGENTS.md` §2; `frontend/lib/api/generated/openapi.ts`. |
| Q3 | Keep Campaign eligibility and Donation state server-owned. Detail availability is an affordance snapshot only; POST is authoritative and may reject stale eligibility. | `api/openapi/campaign.yaml` `PublicCampaignDonationAction`; `api/openapi/donation.yaml`; Campaign feature 02; Donation feature 01. |
| Q4 | Preserve deliberate submission intent: prevent double-click; same idempotency key and payload on ambiguous retry; never rotate key automatically while outcome is ambiguous; failed status permits a new key only after explicit donor action. | Donation feature 01; INV-donation-08/09; shared `IdempotencyKeyHeader`. |
| Q5 | Show only QRIS as an active, clearly labeled sandbox simulation; other familiar methods are visibly unavailable. Never imply real settlement or real payment instructions. | `docs/product/mvp-scope.md` §5; Donation feature 01; `docs/ui-ux/patterns.md` §§14–15. |
| Q6 | Keep `pending` truthful with “Menunggu hasil simulasi”, no estimate/payment instruction/automatic resubmission; terminal result uses “Hasil simulasi donasi: berhasil/gagal”. | Donation feature 01; Donation feature 02; INV-donation-07/08. |
| Q7 | Guest name remains optional and non-public; email remains optional and explicitly opted-in for status-only email. Use approved label and 24-hour verification disclosure, without claiming verification or delivery occurred. | Donation feature 01 exact Design direction; INV-donation-10/11; O3 remains open. |
| Q8 | Status URL credential is handed from fragment to `X-Donation-Status-Credential`, then visible URL is cleaned; response UI reveals status only. Missing/wrong/expired/absent cases use the generic public failure behavior. | Donation feature 02; INV-donation-05/06; authored Donation OpenAPI status operation. |
| Q9 | Follow current Sunlit Editorial / Evidence-Led Optimism intent, semantic status distinctions, accessible form/recovery patterns, and preserve hierarchy/consequences at mobile widths. | `docs/ui-ux/README.md`; `product-design-principles.md`; `patterns.md` §§4,7,12–15; `design-guidelines.md`; design readiness PARTIAL in Exploration. |
| Q10 | Use MSW solely at the network boundary and preserve production requests to real same-origin API paths. Mock evidence is limited to frontend observable behavior. | `frontend/AGENTS.md` §2; frontend architecture §§5,15–18; WU-S2-004 manifest. |
| Q11 | Do not adopt the backend Draft's candidate Rp1,000,000,000 cap, Campaign capacity limit/closure reason, or guest disclosure/API field absent owning-source reconciliation. | WU-S2-003 `TP-S2-003-003/techplan.md` O1-REP/D16; current Product/MVP; accepted WU-S2-005 Campaign contract. |

## 4. Rules & Validation

- **R1 — Campaign action:** Given a public Campaign detail response, when `donation_action.availability` is `available`, then render the truthful Donation entry and navigate with Campaign context; when unavailable, do not offer the Donation action and use only the accepted generic reason. The client must not derive eligibility from `public_state`, progress, or visibility. The Donation submission remains authoritative if the snapshot becomes stale.
- **R2 — Valid amount and contract boundary:** Given an IDR amount, then accept whole Rupiah values from Rp5.000 upward, including Rp5.001, and submit the accepted major-unit decimal string with `currency_code: IDR`. Do not add a monetary upper bound or capacity disclosure from O1-REP until Open Item OI-1 is reconciled. Server validation remains authoritative.
- **R3 — Method truth:** Given the method choices, then QRIS is the only interactive simulated option; GoPay, ShopeePay, and bank transfer are clearly unavailable and cannot submit. Copy and layout must not provide real-payment instructions or imply provider settlement.
- **R4 — Deliberate idempotent intent:** Given a valid form, when submission begins, then repeated activation is suppressed. If the result is ambiguous, any retry uses the same key and same payload; no automatic key rotation or second intent occurs. Same-key changed-payload conflict is shown as a request/business failure while preserving recoverable input. A terminal `failed` state offers a distinct explicit action to start a new Donation with a new key.
- **R5 — Truthful status:** Given `pending`, render “Menunggu hasil simulasi” without a timer, ETA, payment instruction, or automatic resubmission. Given terminal status, label it “Hasil simulasi donasi: berhasil/gagal”; do not characterize it as a real payment or independent verification.
- **R6 — Optional guest fields:** Given optional name/email, permit omission, keep name out of public surfaces, and submit email only with the explicit status-email opt-in. Show the accepted email label/helper exactly. The UI never asserts address ownership verification or notice delivery. No email is required to complete the baseline flow.
- **R7 — Status credential handoff:** Given a status URL with a credential in its fragment, before status lookup, read the fragment credential, send it only in `X-Donation-Status-Credential`, and remove it from the visible URL/history using the supported browser API. Do not put it into query parameters, persistent client state, logs, or rendered text.
- **R8 — Status-only and anti-enumeration UX:** Given a valid credential, render only Donation status and its approved next action. Given missing/wrong/expired credential or absent Donation, show the same generic message “Link status tidak tersedia atau mungkin kedaluwarsa.” Do not distinguish the underlying cause. Backend response/body/header/cache/timing parity remains downstream evidence, not established by the UI or MSW.
- **R9 — Failure and recovery:** Given field validation failure, identify the relevant field; given transport/business failure, show a request-level error distinct from field errors and preserve user input where recoverable. Retry is explicit and respects R4. Errors remain generic and do not expose stack traces, raw server internals, credential, or PII.
- **R10 — Accessible responsive presentation:** Given the flow at supported desktop/mobile sizes and keyboard use, maintain semantic labels, visible focus, status meaning that does not rely on color alone, reachable actions, and priority for current task/trust/consequence information. Human review accepts the rendered hierarchy and comprehension.
- **R11 — Mock evidence boundary:** Given MSW fixtures, they match accepted API types and observable contract states at actual network routes. No mock result is described as runtime settlement/status/email/security evidence; production API adapters contain no mock-mode response branch.
- **R12 — Scope boundary:** No frontend behavior adopts O1-REP's proposed cap/capacity semantics/configuration or adds an input constraint/disclosure. This rule remains gated by OI-1; affected amount guidance is not approval-ready until owning-source reconciliation resolves the candidate policy and any contract dependency.

## 5. Decision Log

| ID | Decision / option | Status | Rationale / consequence |
|---|---|---|---|
| D1 | Use accepted Campaign `donation_action` union as point-in-time affordance; POST rechecks eligibility. | Chosen | WU-S2-005 reconciled and accepted the prior Slice-1 unavailable-only contract. Client inference/hard-coded CTA is rejected. Stale GET may lead to authoritative POST rejection. |
| D2 | Keep WU-S2-004 frontend-only; Campaign/API source changes route through the owning Work Unit/authority. | Chosen | Preserves the manifest boundary and cross-stack coordination. No API/spec/generated edits in frontend Build. |
| D3 | Route-local composition first; introduce domain/shared/UI components only where actual usage establishes a truthful contract. | Chosen | Clean-start architecture and empty broad-contract registry make premature abstractions unnecessary. |
| D4 | Use MSW interception at the network boundary, not production fixture branches. | Chosen | Keeps production adapter contract-faithful while enabling mock milestone evidence; mock evidence remains limited. |
| D5 | Leave the backend D16/O1-REP monetary candidate unadopted pending owning Product/API reconciliation. | Chosen for this Draft; approval blocked for affected direction | The backend plan is Draft and records source decisions/configuration as unresolved. Current accepted Campaign GET contract has no limit field. Do not silently adopt the candidate or extend availability-only action contract. |
| D6-alt | Ignore Campaign action and show CTA based on public visibility/fundraising state. | Rejected | Would recreate eligibility in the client and bypass current API authority; stale-state POST checks do not make this read-side inference valid. |
| D7-alt | Implement status lookup with a Donation ID alone or expose extra Donation data for reassurance. | Rejected | Violates the bearer credential/status-only contract and anti-enumeration boundary. |
| D8-alt | Add a production mock-mode branch to submit/status adapters. | Rejected | Would fork production behavior and weaken contract-parallel verification; MSW owns interception. |

## 6. Backward Compatibility

- **Existing data:** No persistence/schema changes in frontend scope.
- **API/contracts/clients:** Consume accepted Campaign and Donation contracts. Generated OpenAPI types are the shape authority. Do not edit or regenerate shared contract artifacts in this Work Unit. If OI-1 reconciliation changes the required Campaign/Donation field surface, pause affected Build work and coordinate a contract dependency update first.
- **Routes/URLs:** Use current App Router conventions and derive concrete route paths from active route structure during Build. Status credential remains fragment-carried and is removed from the visible URL after handoff. No sensitive token in query/path or durable storage.
- **Migration/deprecation compatibility:** Not applicable; no schema migration or API deprecation is planned.

## 7. Edge Cases & Risks

| ID | Risk / edge case | Likelihood | Severity | Mitigation / accepted exposure |
|---|---|---:|---:|---|
| E1 | Campaign becomes ineligible after detail GET but before POST. | Medium | High | Treat GET action only as affordance; display POST rejection faithfully and do not claim accepted Donation. Backend owns predicate; runtime proof is downstream. |
| E2 | Ambiguous network result, double click, or changed payload with reused key creates duplicate/incorrect intent. | Medium | High | Disable repeated activation; retain same key/payload on explicit retry; no key rotation until known outcome; separate deliberate fresh intent. |
| E3 | `status_token` appears in browser URL/history/log/cache/referrer or client persistence. | Low/Medium | Critical | Fragment handoff, immediate visible URL cleanup, header-only lookup, no logging/persistence/rendering. Frontend browser evidence is necessary but does not prove server/infrastructure controls. |
| E4 | Distinct UI errors reveal Donation existence or credential validity. | Low | High | One generic visible lookup failure; no cause-specific UI. Runtime body/header/cache/timing parity remains O4/O5. |
| E5 | User reads pending/success as provider payment/settlement. | Medium | High | Exact simulation copy, no ETA/payment instructions; Human rendered acceptance for material UI. |
| E6 | Opt-in email suggests verified ownership or delivered notification although verification/delivery controls are unresolved. | Medium | High | Exact accepted disclosure; UI limits itself to captured opt-in request fields and makes no verification/delivery claim. O3 is a downstream contract/evidence dependency for any end-to-end notification fulfillment. |
| E7 | Candidate monetary cap/capacity policy changes validation or requires a guest disclosure absent accepted contract. | High | High | OI-1 blocks final plan approval/affected Build; no guessed client constraint or Campaign API expansion. |
| E8 | MSW success is mistaken for runtime or financial/security proof. | Medium | High | Keep milestone labels/evidence explicit; backend integration, Tier-0 proof, and residual risk remain separate. |
| E9 | Mobile or assistive technology presentation hides method availability, consequences, or status meaning. | Medium | Medium/High | RTL/component evidence plus proportional browser/rendered Human review; semantic and responsive rules in R10. |

## 8. Interface Contract

**Persistence/data shape:** None authored or persisted by frontend. Do not create handwritten mirrors of API-owned request/response types.

**API/event/external interface:**
- Campaign detail: current generated `PublicCampaignDetail`, including required `donation_action` tagged union (`available` or generic `unavailable/campaign_not_eligible`). It is point-in-time display data, not authorization.
- Guest submission: `POST /campaigns/{campaignId}/donations`, shared UUID `Idempotency-Key`, accepted `SubmitDonationRequest` fields (`amount`, `currency_code`, QRIS method, optional `guest_name`, optional `guest_email`, explicit `guest_email_status_opt_in`); accepted Donation response includes persisted `pending` and submission-only `status_token`. Follow live OpenAPI/generated types for exact shape.
- Status: `GET /donations/{donationId}/status`, credential only via `X-Donation-Status-Credential`; response `DonationStatusResponse` exposes status only. Missing/wrong/expired/absent cases map to the common public 404 behavior. Never include token in query.
- MSW intercepts those network requests with contract-faithful response fixtures; production adapters remain pointed at the real endpoints.

**Cross-layer/business boundary:** Client performs UX validation and presentation only. Backend owns amount validation, campaign eligibility, idempotency record, simulator result, persisted Donation state, email verification/delivery, credential verification/expiry, and public 404 parity. The monetary range/configuration candidate in OI-1 is not part of this accepted interface. Exact Campaign action shape is already accepted through WU-S2-005; do not reopen it absent a new owning-source decision.

## 9. Architecture / Plan

1. Extend the existing Campaign detail success composition to render the accepted action branch. Preserve existing loading/error/media handling. Keep the Campaign query authoritative and send Campaign context into the guest form without turning the snapshot into a permission.
2. Add route-local Donation composition and narrow typed API functions using generated types. Use React Hook Form + Zod for form lifecycle and local UX validation. State stays local/form-owned unless a specific server-state need justifies TanStack Query; no Zustand store.
3. On submit, generate/retain the idempotency key for the user intent; suppress repeat activation. Send the same payload/key for an ambiguous retry. Render the returned pending state and status link from the response without inventing simulator timing.
4. On status route entry, read fragment credential, clean the visible URL, request status using the dedicated header, and render status-only response or one generic public error. Do not persist credential.
5. Add MSW handlers and fixtures for accepted Campaign available/unavailable states, submit success/pending, relevant contract errors, status states, and generic status failure. Mock outcomes prove frontend behavior only.
6. Build evidence in phases: focused Build checks for changed artifacts, independent Testing for rule coverage and targeted browser evidence, then Human rendered acceptance of the material guest path before milestone completion. Backend/runtime and real integration gates remain downstream.

## 10. Implementation Details

| Anchor | Why relevant | Intended change / precedent |
|---|---|---|
| `frontend/app/campaigns/[campaignId]/campaign-detail-view.tsx` — `CampaignSuccess` | Current detail success/action composition. | Consume accepted action union; render an action only for `available`, generic non-actionable state otherwise. |
| `frontend/app/campaigns/[campaignId]/campaign-detail-client.tsx` | Campaign query and retry/focus ownership. | Preserve existing GET lifecycle; pass data to view without client-side eligibility derivation. |
| `frontend/app/campaigns/[campaignId]/page.tsx` and route-local layout | Existing dynamic route and MSW startup boundary. | Follow current App Router placement; determine Donation/status route coordinates from current router convention during Build. |
| `frontend/lib/api/client.ts` — `apiRequest` | Current GET-only adapter with generic transport errors. | Extend deliberately to required methods/headers/body while retaining generic safe failures; no mock-mode branch. |
| `frontend/lib/api/public-campaign.ts` — `getPublicCampaignDetail`; `frontend/lib/hooks/use-public-campaign-detail.ts` | Existing Campaign request/hook precedent. | Keep Campaign read owner and generated type mapping. |
| `frontend/lib/api/generated/openapi.ts` — Campaign and Donation `paths`/schemas | Generated API shape source for Campaign action and Donation submit/status. | Import contract types; do not hand-edit. Recheck generated correspondence during Build if accepted source changes first. |
| `frontend/mocks/handlers/public-campaign.ts`; `frontend/mocks/fixtures/public-campaign.ts` | Current Campaign-only network-boundary mocks. | Adapt fixture cases to accepted Campaign action branches; add Donation-specific handler/fixture ownership as needed. |
| `frontend/app/campaigns/[campaignId]/layout.tsx`; `mock-service-worker.tsx` | Current route-local MSW initialization before query. | Ensure Donation routes initialize mocks before requests without introducing production mock logic. |
| `frontend/app/campaigns/[campaignId]/campaign-detail-client.test.tsx`; `frontend/lib/api/public-campaign.test.ts`; `mock-service-worker.test.tsx` | Current observable tests include the old non-activating action. | Replace stale expectation with accepted action behavior and extend observable tests; preserve useful loading/error/focus/hostile-content coverage. |
| `frontend/package.json`, `vitest.config.ts`, `playwright.config.ts` | Actual command/config authority. | Plan relevant commands from current scripts; no execution during this synthesis Run. |
| `docs/ui-ux/patterns.md` §§4,7,12–15; `design-guidelines.md` §§4,13,21–24,26 | Form/status/recovery/money, semantic visual and responsive/accessibility rules. | Use exact product copy and derive ordinary presentation; no new brand decision. |
| `docs/spec/5-donation/features/01-submit-donation-settlement.md`; `02-donation-status-check.md`; `invariants.md` INV-donation-01–11 | Accepted Donation behavior, credential and risk boundaries. | Implement only frontend responsibilities; keep O2–O5 proof/risk boundaries visible. |

## 11. Files Changed / Files NOT Changed

| File / area | Change type | Description |
|---|---|---|
| `frontend/app/campaigns/[campaignId]/` | Modify | Accepted Campaign action and entry navigation; route-local MSW/test adjustments if required. |
| `frontend/app/donations/` (or current equivalent derived from live routing) | Add | Guest form/result/status route composition; exact route path to be selected from current route convention without changing contract. |
| `frontend/lib/api/` | Modify/Add | Typed Donation submit/status adapters and narrowly scoped types only where OpenAPI does not own them. |
| `frontend/lib/hooks/` | Add only if needed | Query/mutation hooks where client-owned server state warrants them. |
| `frontend/mocks/` | Modify/Add | Contract-faithful Campaign and Donation handlers/fixtures at network boundary. |
| `frontend` observable tests/configured browser evidence | Modify/Add | Coverage for rules below using actual project scripts/config. |
| `frontend/components/ui/`, `frontend/components/shared/` | Avoid unless justified | No current broad reusable contract; if a first contract is established, update registry and representative evidence in same change. |

| File / area intentionally untouched | Why |
|---|---|
| `backend/` | Frontend-only Work Unit; runtime/security/integration belong to separate owners and gates. |
| `api/`, `docs/product/`, `docs/spec/`, generated OpenAPI source/types | Accepted authorities are inputs. Changes require owning-source reconciliation and coordinated dependency, not frontend Build edits. |
| Donation ledger/transaction/crypto/auth and disbursement protected areas | Out of scope; root Tier-0 fences apply. |
| Product-wide or brand-defining design assets | No material unresolved design direction or asset requirement identified. |

## 12. Testing Checklist

| Rule | Verification / evidence | Primary owner | Why this is worth running / risk if skipped |
|---|---|---|---|
| R1 | RTL/MSW observable cases for available/unavailable action; verify only accepted field controls entry. Test stale eligibility POST rejection as a request outcome. | Testing | Prevents false action/eligibility claims. Backend predicate proof remains separate. |
| R2 | Table-driven form validation for Rp4.999, Rp5.000, Rp5.001, fractional/invalid values; inspect submitted request exact decimal string/currency. Add no unapproved upper bound. | Testing | Money input errors can reject valid intent or imply unsupported policy. OI-1 must resolve before approving any limit-related direction. |
| R3 | RTL verifies only QRIS submits; unavailable methods cannot activate. Human review of method wording if layout materially changes. | Testing / Human | Prevents implied real payment capability or accidental alternate submission. |
| R4 | Observable tests for repeated activation, same key+payload ambiguous retry, changed payload conflict, and explicit fresh-intent key after failure. | Testing | Duplicate donations are a material user and financial risk; component-only button tests are insufficient for key lifecycle. |
| R5 | RTL covers pending and both terminal labels, no ETA/payment instruction/auto-submit; Human rendered review checks simulation comprehension. | Testing / Human | Automated assertions protect exact wording; human judges whether hierarchy still communicates simulation truth. |
| R6 | RTL covers omitted optional values and email opt-in payload; assert exact approved label/helper and no verification/delivery success claim. | Testing | Prevents accidental PII/notification opt-in and unsupported fulfillment promise. O3 backend controls remain unproven. |
| R7 | Real-browser test (Playwright, narrowly scoped) opens fragment URL, observes header request, then confirms visible URL has no credential and no second credential-bearing navigation. | Testing | jsdom cannot credibly establish browser history/address behavior. Skipping risks exposing a bearer capability. This checks frontend handoff only, not infrastructure leakage controls. |
| R8 | RTL verifies status-only rendering and identical visible generic failure for missing/wrong/expired/absent mock cases; browser check for status route load after cleanup. | Testing | Prevents enumeration through UI and extra PII display. Runtime parity belongs to backend/security Testing. |
| R9 | RTL distinguishes field vs request failures, preserves input on recoverable failures, generic safe text, and explicit retry semantics. | Testing | Recovery mistakes can lose intent or expose internals/PII. |
| R10 | Human rendered acceptance on representative desktop/mobile flow and keyboard/focus path; automated accessible role/name/focus assertions for repeatable behavior. | Human / Testing | Unit tests do not establish hierarchy, clipping, responsive reachability, or interaction comprehension. Material UI requires Human acceptance per frontend authority. |
| R11 | Run focused MSW/API/component checks and inspect adapter target paths; confirm fixtures conform to generated types. Assert no production fixture branch. | Testing | Establishes contract-parallel frontend behavior while avoiding false claims about real service behavior. |
| R12 | Source/contract review before approval: compare O1-REP against current Product/Campaign/Donation authority; verify no cap/capacity/disclosure was coded or implied. | Human / Orchestrator + Testing | A guessed limit can reject valid donations or require a missing disclosure contract; only owner reconciliation can unblock it. |

Suggested current frontend commands for the later authorized phases are `npm run lint`, focused Vitest via the actual configured runner, `npm run verify` for the fast lint+unit/component baseline, and a narrowly scoped `npm run test:browser` only for R7/R8 if the route/browser behavior warrants the planned regression. `npm run build` is appropriate when changed routing/type boundaries make production compilation evidence useful. Final broad verification owner is Testing; Build may run focused checks needed to establish its new tests are executable. Do not run any of these in this synthesis Run.

### Test Focus Pointer

| Area | Why sensitive | Evidence anchor from Exploration | Still relevant post-synthesis? |
|---|---|---|---|
| Bearer status credential fragment, browser URL cleanup, and status-only access | A leaked or mishandled credential grants access to private Donation status; frontend mocks cannot prove exposure controls or server parity. | `WU-S2-004/runs/EXP-S2-004-001/evidence/stage-2-gap-analysis.md#area-2--authored-api-contract`; `#area-4--frontend-architecture-and-live-implementation`; `stage-3-solutioning.md#guest-status-credential-and-proof-boundary` | Yes — R7/R8 require browser evidence; backend O4/O5 remain downstream. |
| Guest email/optional PII and verification/terminal-notice claim boundary | Email is optional sensitive data; UI must not imply verification or delivery while O3 controls remain open. | `WU-S2-004/runs/EXP-S2-004-001/evidence/stage-2-gap-analysis.md#area-1--product-donation-domain-and-approved-delivery-boundary`; `stage-3-solutioning.md#optional-email` | Yes — frontend payload/copy checks remain; O3 backend/privacy evidence deferred and no residual risk is accepted. |
| Donation idempotency and stale Campaign eligibility | Ambiguous retry, duplicate intent, or stale read can cause a duplicate or rejected Donation; actual transaction ordering is backend-owned. | `WU-S2-004/runs/EXP-S2-004-001/evidence/stage-2-gap-analysis.md#area-1--product-donation-domain-and-approved-delivery-boundary`; `#area-4--frontend-architecture-and-live-implementation` | Yes — test client intent/retry observables here; transaction/concurrency proof is downstream WU-S2-003. |

## 13. Open Items

### Active — needs external input or verification

1. **OI-1 — O1-REP Product/API source reconciliation (blocks whole-Techplan approval and affected Build).** Backend `TP-S2-003-003` records a candidate maximum individual Donation of Rp1,000,000,000 configurable per Campaign, cumulative Funding bounded by Campaign whole-IDR storage capacity, and closure at capacity with a proposed new close reason. That Run explicitly marks the source update, admissible configuration range, owner/API reconciliation, and guest disclosure/API question unresolved. Current accepted Campaign GET action is availability-only; do not add a cap, client validation/disclosure, capacity reason, or configuration field from that Draft. Product Authority plus Campaign/Donation/API owners must publish/reconcile the durable policy and interfaces. Until then the no-unapproved-cap path can be designed, but no final plan approval or affected Build may assume numeric-limit behavior. Source: `.harscode-spaces/s2-guest-donation-truthful-state/WU-S2-003/runs/TP-S2-003-003/techplan.md` O1-REP.
2. **OI-2 — O2 simulator timing/scenario mechanics.** Frontend must show backend-returned states with no timing promise and cannot control outcome. Backend implementation/evidence is outside this milestone. This is non-blocking for frontend mock completion unless a required response field/interaction is introduced.
3. **OI-3 — O3 optional email verification and delivery controls.** Exact approved disclosure and request opt-in fields are known, but verification interaction, retry/terminalization controls, retention/deletion evidence, and residual-risk decision remain unresolved. No end-to-end email fulfillment claim; owner evidence is needed before representing a verification subflow. Baseline no-email path and safe opt-in field submission can proceed within accepted contract.
4. **OI-4 — O4/O5 status credential and anti-enumeration runtime evidence.** Frontend fragment/header/cleanup behavior and generic UI must be demonstrated in browser/contract-faithful mocks. Backend/infrastructure key, expiry, exposure, abuse, response/body/header/cache/timing parity, and residual-risk evidence remain with API/Security/PII and independent Testing. This downstream proof is outside `FRONTEND_MOCK_VERIFIED` but required before integrated/security acceptance; no residual risk is accepted.
5. **OI-5 — final rendered Human acceptance.** Required after implementation for responsive hierarchy, action comprehension, simulation/status truth, and accessible interaction. This is a milestone gate, not a prerequisite for drafting this plan.

### Resolved — retained as decision history

1. ~~**F-1 Campaign action contract gap from Exploration**~~ **RESOLVED —** WU-S2-005 accepted a tagged `donation_action` union with available and generic unavailable branches; Campaign feature/INV/API and generated fixture counterparts were reconciled. GET is an affordance snapshot; Donation POST rechecks eligibility. No schema/route decision is reopened here. See WU-S2-005 manifest/accepted Run and current Campaign sources.
2. ~~**Campaign action route recommendation**~~ **RESOLVED —** Keep WU-S2-004 frontend-only and use the accepted Campaign projection; prior Exploration Option 1's coordination route is complete. No Option 2 scope expansion.
3. ~~**Design readiness**~~ **RESOLVED —** PARTIAL: core form/status/error meanings and approved visual direction are established; ordinary route-local composition may be derived. No high-fidelity Figma artifact or new material design decision is required based on current evidence. Human rendered acceptance remains open above.
