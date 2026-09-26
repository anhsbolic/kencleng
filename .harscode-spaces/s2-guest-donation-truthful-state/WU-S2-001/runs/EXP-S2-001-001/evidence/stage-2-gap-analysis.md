# Exploration Evidence — Stage 2 Gap Analysis

## Provenance

- Phase/Stage: Exploration / Stage 2 — Gap Analysis
- Work Unit: `WU-S2-001`
- Run: `EXP-S2-001-001`
- Role: Explorer
- Participant: Codex Explorer
- Author: Codex Explorer
- Created: 2026-09-25
- Model / reasoning: `gpt-6-luna` / `medium` (invocation dispatch metadata)
- Session: `01a0d8ea-1404-7521-99b0-5623057b0519` (from the Run launch record)
- Target revision: `ee0d4b072d9f5cf279952fe309049f687c95e30e` (invocation)
- Observed checkout: `7ee281c4acf6c6860ba52830fe3980b8a88e1940`; the target revision is an ancestor and the only commit after it is `S2_WU-S2-001_EXP-S2-001-001 : Invocation`.
- Workflow revision: `b2d7ca4918b520d960139bc392f87619410b27ed`

## Scope and method

Explored the approved Product/MVP boundary, applicable design authority, historical donation delivery specs and split OpenAPI sources, then the live backend/frontend and current public Campaign boundary. Followed the Harscode five sniffing lenses in each area. No implementation was inspected for a proposed solution and no tests or runtime checks were run. This Participant authored only this evidence artifact under `RUN_PATH`; observed Orchestrator launch-state updates are recorded separately and were not edited by this Participant.

Authority order used: Product/MVP and applicable Design authority → reconciled active-slice delivery detail → split API contract → implementation/tests as evidence. Donation domain artifacts dated 2026-08-20 and marked draft are treated as historical/reference evidence because Slice 2 has not yet reconciled them.

## Area 1 — Product/MVP and slice boundary

### Current state and requirement

The approved MVP loop and primary actor are explicit. `docs/product/mvp-scope.md` §§1–7 says a visitor can donate without an account, understand the real sandbox processing/result state, and later find the same public Campaign; money, privacy, eligibility, and security correctness are floors. `docs/product/mvp-delivery-slices.md` §5 makes Slice 2 a coherent guest flow with amount semantics/validation, one deliberately supported sandbox path if sufficient, truthful pending/success/failure, safe guest status revisiting, and explicit non-real-settlement copy. It requires duplicate protection where required, non-forgeable settlement, exactly-once successful funding updates, concurrency safety, safe guest status access, sensitive token/identifier handling, safe logs, and backend-authoritative Campaign eligibility.

The same Slice 2 section excludes account prerequisite/guest claim/history, payment-method breadth merely for completeness, real rails, public donor/social-proof list, and closure/result behavior except what is needed to preserve donation eligibility. It states historical Donation specs/OpenAPI are reference evidence; exact payment breadth, delays, percentages, minimum amount, endpoint shape, and guest-data requirements need revalidation. The tracker §4 reports Donation active generation `NOT_STARTED`; §7 says Slice 1 is `SLICE_FINALIZED` and Slice 2 needs its own lifecycle.

### Gap

Product scope is sufficiently explicit for the baseline journey and correctness floor. Product Authority does not settle low-level amount minimum, payment method, processing timing/probability, token/lookup contract, or guest data required; those remain delivery/contract reconciliation questions. Existing draft values cannot be promoted by detail alone.

### Five lenses

- **Risk:** Incorrectly importing old Account or public-list breadth expands MVP; incorrect sandbox/status language can misrepresent payment settlement; money concurrency or eligibility errors reach Campaign funding and later public truth.
- **Edge cases:** Campaign becomes ineligible after detail load but before submit; a retry after an ambiguous response; pending/failed donations must not count as collected funding; closure/maximum-amount boundaries intersect eligibility.
- **Miscontext:** Historical domain order and detailed old specs can look like current scope, while the approved vertical slice explicitly excludes much of that breadth.
- **Misleading signals:** The tracker records historical Account implementation and complete Donation specs/contracts, but explicitly says those do not prove current relevance or completion.
- **Inconsistency:** No contradiction in the approved baseline outcome found. The exact legacy choices listed above are not settled Product/MVP decisions.

### Anchors

- `docs/product/mvp-scope.md` §§1–7 — actor, loop, guest donation, correctness/security floor.
- `docs/product/mvp-delivery-slices.md` §5 — Slice 2 scope, exclusions, revalidation posture, completion evidence.
- `docs/project/kencleng-development-tracker.md` §§4, 7 — Donation `NOT_STARTED`; Slice 1 final and Slice 2 lifecycle boundary.
- `docs/kencleng-agentic-workflow.md` §§4–5 — risk is assigned before Build; backend/frontend preconditions and authority routing.

## Area 2 — Product Design / UX authority

### Current state and requirement

`docs/ui-ux/README.md` routes stable experience principles, brand direction, visual grammar, interaction patterns, and persona/surface inventory to separate canonical sources. `product-design-principles.md` §§1–6 requires confidence before donation, clear money meaning, distinction between platform fact/report/system state/pending information, and visible unknowns. `patterns.md` §§3–4, 7 requires validation and request failures to remain distinct, duplicate prevention for non-idempotent submission, meaningful consequence/next action for status, and anti-enumeration-compatible failures for sensitive unauthenticated lookup. `page-map.md` §1 identifies public Donation Flow (Form) and Donation Status/Tracking surfaces. `design-guidelines.md` §§3, 10–11 says optimism/color must not upgrade evidence certainty and pending states must stay explicit. `brand-product-ui-brief.md` §§14–15 leaves final truth-state terminology and some source labeling open.

### Gap

The journey categories and truthfulness constraints are available. No Slice-2-specific rendered flow, exact status vocabulary, or source labeling has been reconciled. The open terminology is a design follow-up only if exact language is needed; status meaning itself remains domain/product authority.

### Five lenses

- **Risk:** A success-colored or celebratory treatment could overstate a sandbox result; a status badge without consequence/next action can confuse a donor.
- **Edge cases:** Pending, failed, inaccessible/expired-or-invalid lookup, request failure, stale campaign eligibility, validation error, and successful retry/revisit need distinct supported presentation states.
- **Miscontext:** Existing campaign page is a Slice-1 public detail, not a donation flow; visual convention cannot supply donation or payment semantics.
- **Misleading signals:** Page-map surface names describe intent, not proof that pages/routes exist; visual references are not contracts or business authority.
- **Inconsistency:** No active design rule was found that contradicts the product boundary. Exact final truth-state/source terminology is expressly open.

### Anchors

- `docs/ui-ux/product-design-principles.md` §§1, 3, 5–6 — donation context, money semantics, truth classes, unknown states.
- `docs/ui-ux/patterns.md` §§3–4, 7 — Detail/Form/Status interaction rules and anti-enumeration UX boundary.
- `docs/ui-ux/page-map.md` §1 and §8–9 — guest surfaces and unresolved surface questions.
- `docs/ui-ux/design-guidelines.md` §§3, 13–14 — evidence-aware expression, provenance/pending, and funding progress grammar.
- `docs/ui-ux/brand-product-ui-brief.md` §§5, 12 — status/pending behavior and open final terminology.

## Area 3 — Donation delivery specs and API evidence

### Current state and requirement

Draft donation specs (`docs/spec/5-donation/`, authored 2026-08-20) describe: `POST /campaigns/{campaignId}/donations`; internal asynchronous settlement; pending→success/failed; an `Idempotency-Key`; amount minimum Rp 5,000; six payment-method values; a non-expiring `status_token`; public donor list; Account history/claim; and PII encryption/HMAC. `invariants.md` INV-donation-08/09/11 describe conditional atomic funding increment, idempotent settlement, and submission idempotency. INV-donation-02 requires active `published` Campaign at submit time; INV-donation-05 makes the status token a long-lived read credential; INV-donation-04 uses the established guest-email encryption/HMAC pattern. The threat model calls out settlement never being HTTP-reachable and notes the token itself grants status-read access.

`api/README.md` makes `api/openapi/donation.yaml` the split domain source and `api/openapi/common.yaml` the source of shared components; bundled `api/openapi.yaml` and generated frontend types are derived. The draft Donation OpenAPI has POST submit, GET status, public donor-list and Account operations. The active Slice-1 Campaign spec/API instead has a closed public projection whose `donation_action` is unavailable and has no activation target (`docs/spec/4-campaign/features/02-campaign-detail-listing.md` §§Summary/Behavior; `api/openapi/campaign.yaml` `PublicCampaignDonationAction`). This expresses the Slice-1 boundary, not a permanent Slice-2 rule.

### Gap and authority questions

No reconciled Slice-2 Donation feature spec or API contract was found. The old contract includes capabilities outside Slice 2 (public donor list, Account history/claim, event context) and six payment methods; Product/MVP explicitly excludes several and says one sandbox path may suffice. Exact minimum, payment choice, timing/failure simulation, guest data, status credential semantics, and endpoint shapes require revalidation in their owning delivery/contract authorities.

The old status feature spec says absent donation, wrong token, and missing token must produce identical `401`; the OpenAPI `GET /donations/{donationId}/status` also lists `404` for not found. This is an internal draft spec/contract inconsistency, not a resolved choice. The old threat model records non-expiring token and no status-endpoint rate limit as accepted sandbox residual risks; current MVP policy still requires the applicable security floor and Human/project risk gate, so that old acceptance is not current authorization.

Donation feature 01 says a successful amount crossing `max_amount` triggers Campaign closure in the same transaction; Campaign feature 09/INV-campaign-13 describes that trigger. Slice 2 excludes closure/result behavior except what is required to keep eligibility correct. The exact necessary interaction at this boundary needs explicit reconciliation; the historical hook alone does not establish broader Slice-2 closure scope.

### Five lenses

- **Risk:** Public unauthenticated creation and token-guarded status are sensitive boundaries; settlement forgery, duplicate creation, duplicate increment, lost funding update, PII/token logging, or a stale eligibility check can harm financial truth/privacy.
- **Edge cases:** same-key retry after lost response; two simultaneous successful settlements; pending versus terminal status; missing guest email; crossing the maximum; Campaign closing between read and submit; not-found/wrong/missing status token; amount precision and minimum boundary.
- **Miscontext:** Detailed draft specs/API can appear implemented and approved, but Product/MVP labels them reference evidence; account and donor-list operations are outside baseline Slice 2.
- **Misleading signals:** Donation paths/types remain present in split/generated OpenAPI despite no live Donation module/routes. Schema presence is not runtime availability.
- **Inconsistency:** Status lookup 401-vs-404 discrepancy; historical draft enum breadth versus one sufficient sandbox path; Slice-2 closure exclusion versus old same-transaction closure hook need reconciliation.

### Anchors

- `docs/spec/5-donation/features/01-submit-donation-settlement.md` §§Summary, Critical, Behavior, Concurrency — old submit/settle design and risks.
- `docs/spec/5-donation/features/02-donation-status-check.md` §§Behavior, Validation — uniform 401 expectation.
- `docs/spec/5-donation/invariants.md` INV-donation-02, 04, 05, 08, 09, 11 — eligibility, PII, status token, atomicity, idempotency.
- `docs/spec/5-donation/threat-model.md` “Submit donation”, “Token-based status check”, “Internal settlement process” — threat evidence and historical accepted residual risks.
- `api/openapi/donation.yaml` paths `/campaigns/{campaignId}/donations`, `/donations/{donationId}/status`, and Account paths; `PaymentMethod`, `SubmitDonationRequest`, `Donation` schemas.
- `api/openapi/common.yaml` `IdempotencyKeyHeader`, `Unauthorized`, `NotFound` — shared contract components.
- `docs/spec/4-campaign/features/09-closure.md` and `docs/spec/4-campaign/invariants.md` INV-campaign-13 — old cross-domain close trigger.
- `docs/spec/4-campaign/features/02-campaign-detail-listing.md` and `api/openapi/campaign.yaml` `PublicCampaignDonationAction` — active Slice-1 public action boundary.

## Area 4 — Backend live state and security/correctness boundary

### Current state and requirement

At observed checkout `7ee281c…`, `backend/cmd/server/main.go` registers public Campaign detail/media plus auth/other existing operations; no Donation route/service wiring was found. No `backend/internal/domain/donation/` package or Donation migration was found in the live tree. The current Campaign service constructs `DonationAction{Availability: "unavailable", Reason: "donation_flow_not_available"}`. Campaign public detail resolves only eligible published campaigns and maps a closed public projection. `backend/AGENTS.md` §3 calls for race evidence for future concurrency-sensitive donation/disbursement behavior; root `AGENTS.md` fences future `backend/internal/domain/donation/ledger.go` and transaction/locking implementation from agent writes without explicit human authorization. The Orchestrator launch record contains the Session identity and launch-time state; this Participant did not edit Work Unit/control files.

### Gap

The live backend does not currently implement the Slice-2 submit/status/settlement path, storage, or runtime contract. There is no current Donation runtime code to compare against draft Donation invariants. Future work will touch sensitive money/concurrency/security areas, including a protected Tier-0 ledger/locking path; this Exploration has only read current evidence.

### Five lenses

- **Risk:** Settlement must be internal-only; money updates must be exact once under retries/concurrency; explicit campaign eligibility and guarded state changes protect public truth. Protected ledger/locking code is human-authored/human-paired.
- **Edge cases:** Concurrent submissions/settlements, pending transition replays, close-versus-settle ordering, boundary amounts, and DB failure partway through the compound donation/funding state change have no live Donation implementation to validate.
- **Miscontext:** Draft spec language about transactions/locking is not evidence of a current implementation. Backend presence for Campaign does not imply Donation support.
- **Misleading signals:** `GET /campaigns/{campaignId}` works, but response always marks donation unavailable; old Donation routes are not wired into `ServeMux`.
- **Inconsistency:** No runtime-vs-draft Donation code contradiction can be measured because runtime Donation code is absent. Draft status lookup discrepancy remains as recorded in Area 3.

### Anchors

- `backend/cmd/server/main.go` router registrations around `GET /campaigns/{campaignId}` — live route surface; no donation operation registered.
- `backend/internal/domain/campaign/service.go` `GetPublicDetail` and `toPublicDetail` (around line 96) — current eligibility/projection and unavailable donation action.
- `backend/internal/transport/http/campaign_public.go` `PublicCampaignDetailHandler`, `publicCampaignDetailResponse`, and `toPublicCampaignDetailResponse` — current public wire boundary.
- `backend/internal/domain/campaign/repository_db.go` `FindPublicDetail` — live Campaign eligibility query boundary.
- `backend/AGENTS.md` §3 and root `AGENTS.md` §3 — future concurrency verification and protected donation ledger path.
- Absence evidence: no live `backend/internal/domain/donation/` files, donation router/service references, or donations migration found by repository search.

## Area 5 — Frontend live state and cross-stack surface

### Current state and requirement

Live frontend routes/files are the public Campaign detail only; no Donation form/status route or donation-specific API client/hook was found. `CampaignSuccess` renders a non-activating next-action area explaining that donation is unavailable. `campaign-detail-view.tsx` has current campaign/funding rendering and loading/not-found/unavailable/failure states. Its current fixture carries `donation_action.unavailable`; the Campaign detail client test is explicitly about rendering a non-activating donation context. Frontend generated OpenAPI types still contain historical Donation operations, but no live client flow uses them. `docs/project/kencleng-integration-map.md` § Public Campaign Detail records the current FE/BE public Campaign integration boundary.

### Gap

No Slice-2 donor interaction, request/retry behavior, form validation, pending/result display, or safe status revisit exists in the current frontend. FE presentation must consume the future reconciled API/domain semantics; it must not invent status, money validity, campaign eligibility, authorization, or credential security.

### Five lenses

- **Risk:** Client-only success, client-derived funding, duplicate submissions with fresh keys, or exposing/storing a guest status credential unsafely would make the visible result misleading or disclose donation data.
- **Edge cases:** Client timeout after submit, double activation, invalid amount, server-side campaign rejection after a stale detail page, pending state when revisited, status lookup mismatch, and response/network failure.
- **Miscontext:** A working Campaign page, generated types, and MSW fixtures do not constitute a live donation flow or backend contract implementation.
- **Misleading signals:** The generated TypeScript contains Donation paths/schemas from the old bundle, while the actual public page explicitly says the flow is unavailable and no Donation client/hook exists.
- **Inconsistency:** Current API fixture/runtime both say unavailable, matching the reconciled Slice-1 boundary. No live FE/BE Donation integration exists to compare.

### Anchors

- `frontend/app/campaigns/[campaignId]/campaign-detail-view.tsx` `CampaignSuccess` and `Funding` — current non-activating action and funding display.
- `frontend/app/campaigns/[campaignId]/campaign-detail-client.tsx` `CampaignDetailQuery` — existing Campaign loading/error state flow.
- `frontend/lib/api/public-campaign.ts` `getPublicCampaignDetail` — only relevant current Campaign API client found.
- `frontend/mocks/fixtures/public-campaign.ts` `baseCampaign.donation_action` — fixture confirms no active flow.
- `frontend/app/campaigns/[campaignId]/campaign-detail-client.test.tsx` — current test intent covers non-activating context (not executed in this Exploration).
- `frontend/lib/api/generated/openapi.ts` around path `/campaigns/{campaignId}/donations` — generated legacy type presence only; no runtime support implied.
- `docs/project/kencleng-integration-map.md` Public Campaign Detail row — current integration boundary.

## Cross-cutting evidence routed to later reconciliation

- **Financial precision/atomicity/idempotency:** The relevant portable guidance routes to `go/decimal-and-money.md`, `postgresql/financial-invariant-enforcement.md`, `postgresql/transactions-and-locking.md`, and `restapi/idempotency-and-versioning.md`. The old Donation spec proposes decimal strings, idempotent POST, guarded state transition, and atomic conditional increment; none is currently implemented for Donation.
- **Sensitive status credential:** `restapi/anti-enumeration.md` and `go/secrets-and-sensitive-logging.md` are relevant routed guidance. The old threat model recognizes token sensitivity; no implementation exists to establish comparison, URI/log handling, response/timing parity, or rate limiting.
- **Slice boundary:** Account flows/public donor list/multiple methods are excluded unless evidence proves enabling-critical. Campaign closure details remain included only to the extent needed to preserve donation eligibility; this boundary intersects the old max-amount hook and needs explicit reconciliation.
- **Coordination:** Evidence establishes a BE runtime/contract gap plus a FE journey gap on the existing public Campaign surface. It does not establish the downstream Work Unit split or dependency graph; those are for Orchestrator after this Exploration evidence is reviewed.

## Evidence status

- **Verified by source inspection:** Approved Slice-2 outcome and scope; Slice-1 Campaign public action unavailable; live backend has no Donation route/domain/migration; live frontend has no Donation flow/status page/API client; draft contract and current Campaign contract differ in lifecycle status; draft status feature/API disagree on not-found response.
- **Assumed:** No behavioral assumption beyond the inspected repository sources. The draft Donation behavior is labeled historical evidence, not assumed current truth.
- **Deferred:** Exact Donation amount rules, payment path, delay/failure semantics, guest field policy, status credential/access/error model, and the max-amount/closure interaction; these need delivery/API reconciliation.
- **Not tested:** No automated tests, API validation, build, runtime flow, race test, or integration check was run. Exploration used read/search only.

## Stage 2 completion checkpoint

Stage 2 evidence covers Product/MVP, Design, Donation specs/API, backend, frontend, and current cross-stack boundaries. No solution has been selected. The material delivery/contract questions and the concrete draft inconsistency above remain visible for review. Human confirmation is required before Stage 3 — Solutioning under the canonical Exploration prompt.
