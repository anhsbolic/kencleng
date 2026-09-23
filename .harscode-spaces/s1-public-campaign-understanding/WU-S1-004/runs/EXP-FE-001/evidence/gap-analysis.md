# Stage 2 — Frontend Public Campaign Detail Gap Analysis

> Phase/Stage: Exploration / Stage 2 — Gap Analysis  
> Work Unit / Run: `WU-S1-004` / `EXP-FE-001`  
> Author: Codex CLI agent  
> Specialization: Frontend Public Campaign Detail, generated API types, MSW boundary, UI/UX authority, rendered states  
> Created: 2026-09-23  
> Model / Reasoning / Session: `gpt-5.6-terra` / high / Fresh session  
> Target revision: `70ef5c6af1a9` (`validation-03-orchestrator-slice-1`)  
> Workflow revision: not exposed by this Run

## Scope and authority read

This evidence concerns only Slice 1's Public Campaign Detail frontend work. Product/MVP authority is `docs/product/mvp-delivery-slices.md` §4; reconciled delivery authority is Campaign feature 02/03 plus `INV-campaign-14`; API authority is `api/openapi/campaign.yaml`; UI authority is canonical `docs/ui-ux/**`; frontend architecture/rules are `frontend/AGENTS.md` and `docs/project/kencleng-frontend-tech-stack.md`.

`docs/project/kencleng-integration-map.md` §6 maps this surface to `getPublicCampaignDetail` and `getPublicCampaignMediaContent`. It confirms the embedded public steward projection and records live controlled-media/proxy enforcement as a later backend/topology/integration concern.

## Area 1 — Product, reconciled Campaign behavior, and public API contract

### Current state

- Slice 1 exposes a public detail only for an internally `published` fundraising Campaign. It uses the standalone, closed `PublicCampaignDetail` allowlist: `id`, `title`, organizer-text `purpose` and `story`, `steward`, `lifecycle`, `funding`, `media`, and `donation_action` (`docs/spec/4-campaign/invariants.md` `INV-campaign-14`; `api/openapi/campaign.yaml` `PublicCampaignDetail`).
- `getPublicCampaignDetail` is a public `GET /campaigns/{campaignId}` with exactly `200`, `404`, and `503`; optional authorization cannot alter the public response. Both `404` and `503` carry `Cache-Control: private, no-store` (`api/openapi/campaign.yaml` `getPublicCampaignDetail`, `PublicCampaignNotFound`, `PublicCampaignUnavailable`).
- The response communicates source and availability as data: organizer text is plain string plus `source: organizer`; funding is a tagged union; media is `available | absent | unavailable`; and Slice 1's `donation_action` is required but only `availability: unavailable`, `reason: donation_flow_not_available`. Amounts and percentage are strings, not numbers.
- The controlled media-byte operation is a separate public `GET`, and `content_url` is an opaque same-origin `/api/campaigns/.../content` reference. It must not imply public object storage or redirect semantics.

### Requirement

`docs/product/mvp-delivery-slices.md` §4 requires truthful campaign/steward/funding/media/lifecycle/source context, unknown/pending state, and a truthful next-action area. It forbids an active Donate action before the real Donation Flow exists. The reconciled feature specs require the exact public projection and distinguish non-public/not-found from eligible dependency unavailability (`docs/spec/4-campaign/features/02-campaign-detail-listing.md` §§Summary, functional requirements, error cases; `03-campaign-media.md`).

### Gap

The contract is `CONTRACT_READY`, but no frontend route or presentation exists to consume any public-detail field or to represent the mutually distinct contract states. The eventual surface therefore needs contract-shaped coverage before it can claim Slice-1 understanding; this is not a request to extend public payloads or derive hidden eligibility locally.

### Sniffing

- **Risk:** Collapsing tagged unavailable/absent/error states or interpreting strings as client-computed money can make a public trust claim false. Rendering any inherited historical `Campaign`/`Organization` fields would breach the deliberate public allowlist.
- **Edge cases:** `collected_amount: "0.00"` is an available fact; percentage may be `null` when `state: not_computable`; funding can be `above_target`; media available requires at least one item whereas absent/unavailable require none; captions may be `null`.
- **Miscontext:** Probe 01 describes public continuity after closure, but active Slice 1 contract/specs deliberately expose only `fundraising`; closure/result is Slice 3. A frontend must not use the approved future direction to display a closed-public experience now.
- **Misleading signals:** The generated aggregate still contains historical/internal campaign, donation, report, attachment-list, and lifecycle operations. Their presence does not place those shapes or flows in Slice 1. Only the two mapped public operationIds are active surface inputs.
- **Inconsistency:** No active authority conflict was found. The apparent Probe-01 closed-state tension is resolved by current MVP sequencing and `INV-campaign-14`, which explicitly defer historical closed states until later reconciliation.

### Code/contract anchors

- `api/openapi/campaign.yaml` — `getPublicCampaignDetail`, `getPublicCampaignMediaContent`, `PublicCampaignDetail` through `PublicCampaignDonationAction`, public response components.
- `frontend/lib/api/generated/openapi.ts` — `components["schemas"]["PublicCampaignDetail"]`, discriminated public unions, `operations["getPublicCampaignDetail"]`, `operations["getPublicCampaignMediaContent"]`; generated artifact is the TypeScript shape source.
- `docs/spec/4-campaign/invariants.md` — `INV-campaign-14`; security and truth boundary to preserve in frontend presentation/testing.

## Area 2 — Public-detail design authority, media truth, and interaction states

### Current state

- Canonical design authority establishes **Sunlit Editorial / Evidence-Led Optimism**, evidence-first hierarchy, organizer attribution, non-gamified funding progress, calm unavailable states, and responsive preservation of trust-critical information (`docs/ui-ux/product-design-principles.md` §§1–8, 13–15; `docs/ui-ux/design-guidelines.md` §§13–14, 21, 23).
- The page map classifies Public Campaign Detail as a public `Detail` surface for campaign, steward, funding context, story, and available donation action (`docs/ui-ux/page-map.md` §1). The Detail pattern establishes identity/state → facts → trust/context → supporting detail → viewer/state-appropriate actions (`docs/ui-ux/patterns.md` §3).
- Campaign imagery may frame real campaign media but cannot upgrade it into beneficiary identity, distribution, verification, or impact. The missing-media treatment must be recognizable as a placeholder, not synthetic campaign evidence. Exact reusable placeholder asset implementation and final Indonesian provenance terminology remain `OPEN` (`docs/ui-ux/design-guidelines.md` §§17–18, 26; `docs/ui-ux/asset-governance.md` §5).

### Requirement

The product scope requires campaign identity/purpose, public-safe steward context, truthful media, lifecycle, funding, organizer provenance, unknown/pending information, and a truthful next-action area. Design requires distinction between platform facts, organizer-provided information, system state, and unavailable information; it also requires no color-only meaning, and priority preservation at narrow viewports.

### Gap

No Public Campaign Detail hierarchy, visual treatment, presentation copy, state composition, or responsive rendition exists. The canonical material supplies a sufficient direction for the ordinary composition and visual grammar, but final Indonesian provenance/action wording and any reusable/expressive placeholder asset retain their stated Human Design gate.

### Sniffing

- **Risk:** Strong editorial imagery, badges, or progress treatment can imply verification, beneficiary identity, execution, or impact not present in the API. Hiding source or unavailable states would undermine the MVP's trust premise.
- **Edge cases:** Long organizer story/title, an empty media collection with `state: absent`, unavailable media/funding, nullable caption, zero or above-target funding, and narrow viewports must retain source, financial labels, and state meaning without clipping or replacing it with a decorative substitute.
- **Miscontext:** The page map says “available donation action,” while Slice 1's contract makes that action explicitly unavailable. The page map is a surface inventory, not permission for an active CTA; the active product/contract boundary wins.
- **Misleading signals:** Selected visual references establish character and hierarchy, not pixel geometry, API fields, business state, or a real campaign image asset. Existing landing-page CSS also cannot establish detail-page semantics.
- **Inconsistency:** There is no conflict between the design authority and contract once “action” is interpreted as a truthful next-action area, not a functional donation control. Final public Indonesian labels remain intentionally unresolved rather than contradictory.

### Code/authority anchors

- `docs/ui-ux/product-design-principles.md` — confidence before conversion, truth classes, unknown state, funding-as-evidence, responsiveness, readiness decision boundary.
- `docs/ui-ux/design-guidelines.md` — provenance/truth grammar, funding progress, campaign media/placeholder, responsive and accessibility rules, current open decisions.
- `docs/ui-ux/patterns.md` — Detail, Loading, Error and Recovery, Money Presentation, Responsive Transformation.
- `docs/ui-ux/page-map.md` — Guest/Public Visitor → Public campaign detail.

## Area 3 — Live frontend surface and architecture readiness

### Current state

- `frontend/` is the clean-start baseline. Only `frontend/app/page.tsx` exists as an editorial home page; its sole test is `frontend/app/page.test.tsx`. There is no campaign route segment, route-level loading/error file, feature component, API function, query hook, mock directory, or browser test.
- `frontend/app/layout.tsx` already loads the canonical font pairing (`Instrument Sans`, `Newsreader`) and sets `lang="id"`; `frontend/app/globals.css` contains landing-page CSS and a narrow responsive baseline. These are implementation facts, not stronger visual authority than `docs/ui-ux/design-guidelines.md`.
- The approved architecture is Next.js App Router + TypeScript + Tailwind CSS, generated types, optional TanStack Query, Vitest/RTL/MSW, and on-demand Playwright. Route-local composition is the default narrow owner; components do not need a premature shared layer (`frontend/AGENTS.md` §§1, 3, 5, 10–11; `docs/project/kencleng-frontend-tech-stack.md` §§4–8, 15–20).

### Requirement

The Work Unit scope requires a public campaign route/surface with loading, success, missing-media, unavailable/failure, not-public/not-found, responsive, and accessibility states. The task explicitly excludes backend work, donation flow, broad discovery, and resurrecting the retired product UI.

### Gap

The complete Slice-1 frontend capability is absent from the live application. No compatibility burden from an older campaign UI exists in this clean tree; future work must establish new route-local/domain ownership from current authority rather than porting historical structures.

### Sniffing

- **Risk:** A large new route can accidentally promote early abstractions into `components/ui`/`components/shared`, or make the entire route a Client Component solely to fetch/render state. Both increase blast radius or browser surface without current evidence.
- **Edge cases:** Dynamic route identity must be reachable/bookmarkable; asynchronous replacement of loading content needs semantic/focus-aware treatment; realistic content lengths and financial strings need responsive coverage rather than only the current landing-page viewport rules.
- **Miscontext:** The currently polished home page and its test demonstrate scaffold/readiness only. They do not demonstrate Public Campaign Detail, contract handling, mock fidelity, or material UI acceptance.
- **Misleading signals:** Dependencies for TanStack Query, MSW, Playwright, Zustand, forms, and Zod are installed, but no current source uses them. Availability is not evidence that any of their patterns already apply.
- **Inconsistency:** No current implementation contradicts the task because the campaign route does not exist. The baseline's local CSS conventions are subordinate to current design authority, as frontend rules explicitly state.

### Code anchors

- `frontend/app/page.tsx` — current `/` baseline only; confirms no campaign surface.
- `frontend/app/layout.tsx` — global language/font setup relevant to rendered route context.
- `frontend/app/globals.css` — existing global rules; later changes must distinguish global baseline from route-specific styling.
- `frontend/AGENTS.md` — ownership, state, design, mock-boundary, rendered-acceptance rules.
- `docs/project/kencleng-frontend-tech-stack.md` — App Router, state, component, and verification architecture.

## Area 4 — Generated types, network boundary, MSW, and live-integration dependency

### Current state

- `frontend/lib/api/generated/openapi.ts` is present and generated from the bundled OpenAPI via `npm run generate:api-types`. It contains the exact two public operation types and their response unions.
- Repository search found no production `fetch`, typed request function, TanStack Query hook, MSW server/worker/handler, request base configuration, generated-type import, or mock-only production branch. `frontend/next.config.ts` is empty and contains no `/api` rewrite/proxy.
- The integration map specifies a contract-parallel boundary: production data access uses real operation paths; MSW intercepts the network request while backend is unavailable. It also records same-origin `content_url`, private storage, no-store behavior, and `/api` routing as later controlled-media/topology integration work, not frontend-created semantics.

### Requirement

The Work Unit requires a typed production data boundary from generated API types and contract-faithful MSW at the network boundary. `frontend/AGENTS.md` §2 explicitly prohibits alternate mock JSON returned from a production service according to mode/environment. `docs/kencleng-agentic-workflow.md` §§8 and 10 define contract-parallel work and `FRONTEND_MOCK_VERIFIED` without equating it to live integration.

### Gap

Neither production request path nor MSW boundary exists. The frontend cannot yet demonstrate that its rendered route uses the reconciled public operation or that mock responses exercise the same request/response shape that real integration will use. The live same-origin `/api` route remains intentionally unimplemented/deferred outside this Work Unit's scope.

### Sniffing

- **Risk:** Handwritten local models, mock service branches, direct object URLs, endpoint/path drift, or future proxy assumptions can make mock success differ from real API behavior. A media mock that bypasses the controlled path would conceal the revocability/security boundary.
- **Edge cases:** Contract mocks need distinct successful detail, public-safe `404`, dependency `503`, funding unavailable, media absent, media unavailable, and available media with controlled `content_url`; image byte delivery has JPEG/PNG-only success semantics. A malformed/non-public ID remains indistinguishable from absent at the public contract level.
- **Miscontext:** No rewrite in the frontend currently means the live `content_url` path cannot be treated as a frontend defect to solve in this Work Unit. The integration map assigns proxy/topology enforcement to later work; mock-parallel evidence must not claim it is live.
- **Misleading signals:** Generated types alone do not execute HTTP, validate runtime payloads, supply handlers, enforce cache headers, or provide the `/api` topology. Similarly, MSW being installed does not mean interception exists.
- **Inconsistency:** No active contract conflict was found. The actual API operation path is `/campaigns/{campaignId}`, while returned media `content_url` is deliberately `/api/...`; the latter is an opaque same-origin public-facing path and must not be normalized into a presumed direct backend/storage URL.

### Code/contract anchors

- `frontend/package.json` — `generate:api-types`; installed MSW/testing capability.
- `frontend/lib/api/generated/openapi.ts` — sole generated public detail and operation declarations; no imports outside the generated artifact at this revision.
- `frontend/next.config.ts` — no current rewrite/proxy contract.
- `docs/project/kencleng-integration-map.md` §5–6 — active surface mapping and contract-parallel rule.
- `docs/kencleng-agentic-workflow.md` §§8–10 — mock-first milestone boundary.

## Area 5 — Observable state, hostile-content, accessibility, and rendered evidence

### Current state

- Current verification setup has Vitest + RTL with jsdom and a Playwright configuration, but only the home-page test exists and `tests/browser/` does not exist. `npm run verify` is lint plus unit/component tests; `npm run test:browser` is available but on-demand.
- Frontend and project authority require material UI to receive rendered feedback and manual human browser acceptance before delivery. `FRONTEND_MOCK_VERIFIED` additionally requires scoped frontend verification and contract-faithful mocks; it is explicitly not live backend/proxy integration.
- The contract declares organizer content to be plain text with no HTML/Markdown. Its threat model still requires safe framework rendering and hostile-looking plain-text evidence downstream (`docs/spec/4-campaign/threat-model.md`, Public Campaign detail row).

### Requirement

Slice 1 completion requires correct loading, missing-media, success, not-public/not-found, failure, and responsive states without internal data leakage (`docs/product/mvp-delivery-slices.md` §4). The Work Unit further calls for accessibility behavior, hostile-looking organizer plain-text rendering evidence, and human rendered acceptance. The relevant React guidance requires network-layer MSW over hook mocking, observable state assertions, semantic control/name behavior, non-color-only state, responsive stress content, and no raw backend error rendering.

### Gap

There are no Slice-1 state tests, MSW tests, rendered checks, a hostile-text test, browser regression, or human acceptance record. Current baseline verification cannot establish any Slice-1 state transition, public-data boundary, responsive result, or accessibility behavior.

### Sniffing

- **Risk:** Rendering `error.message`/Problem Details directly can leak backend internals; rendering organizer text as markup can introduce stored XSS or false provenance; only testing happy-path static markup misses the exact trust/error states this slice promises.
- **Edge cases:** Hostile-looking organizer text including markup-like characters and long unbroken strings; `404` versus retryable `503`; no media versus unavailable media; long stewardship name/story; zero/large/over-target decimal strings; narrow mobile/wide desktop; keyboard-visible focus and image alternative text.
- **Miscontext:** A generic “empty state” is not the contract response for this detail route. `404`, `503`, `media.absent`, and `media.unavailable` have distinct meanings and cannot be rendered as a single no-content presentation.
- **Misleading signals:** Green lint/type checks, a passed home-page test, or a visual screenshot do not prove network contract use, error state behavior, hierarchy under stress content, keyboard access, or human acceptance. Playwright availability does not make it mandatory by itself.
- **Inconsistency:** The product completion list names missing-media and unavailable/failure separately; the generic UX pattern's four async states must be applied with the Campaign contract's finer discriminants rather than used to erase them. No authority requires a committed browser test before mock verification, but human rendered acceptance remains mandatory for material UI.

### Verification anchors

- `frontend/vitest.config.ts`, `frontend/vitest.setup.ts`, `frontend/playwright.config.ts`, `frontend/package.json` — current executable test capability.
- `frontend/AGENTS.md` §§10–12 — human rendered acceptance, test/browser boundary, and security presentation.
- `docs/kencleng-agentic-workflow.md` §§10, 15 — mock-verification conditions and hostile-content/rendered evidence expectations.
- `../harscode-workspace/best-practices/react/component-test-mocking-discipline.md` — MSW network-layer and observable-state expectations.
- `../harscode-workspace/best-practices/react/loading-empty-error-state-conventions.md` — safe error and not-found distinction.
- `../harscode-workspace/best-practices/react/accessibility-fundamentals.md`, `responsive-layout-robustness.md`, `visual-verification.md` — accessibility, realistic content/responsive, and representative rendered verification lenses.

## Stage-2 findings and boundary status

1. **Finding — full feature absence, contract readiness present.** The clean frontend baseline has no Slice-1 route/data/mock/test implementation, while the reconciled public contract and generated types are available. This is delivery work, not contract redefinition.
2. **Finding — truth states are first-class rendering inputs.** The eventual frontend must preserve distinct public `404`, dependency `503`, funding availability, media availability/absence/unavailability, and unavailable donation action. No frontend-local lifecycle/eligibility/funding calculation is authorized.
3. **Finding — mock fidelity and topology are separate.** MSW must intercept the future production request at the network boundary. Same-origin controlled-media routing, private storage, origin recheck, and no-store preservation are real-integration dependencies owned outside this frontend Work Unit and must remain visible as deferred, not simulated as verified.
4. **Finding — final Human Design gates remain.** Final Indonesian provenance/action wording and a precedent-setting reusable/expressive placeholder asset require the stated design authority/human review. The surface's ordinary hierarchy and truthful state treatment otherwise have adequate authority to proceed to planning.
5. **Finding — verification must cover hostile text and rendered states.** Unit/component evidence alone cannot establish the material UI. The required future evidence includes contract-shaped MSW state coverage, safe plain-text organizer rendering, accessibility/responsive conditions, rendered inspection, and Human rendered acceptance. Playwright is on-demand, not currently mandated.

## Deliberately not decided in Stage 2

- Route path/file decomposition, Server-versus-Client data ownership, component/API/hook shapes, query strategy, retry mechanics, exact visual geometry/copy, placeholder implementation, and test/file layout.
- Whether a browser regression is justified; this requires Stage-3 trade-off evaluation rather than an automatic dependency of the available Playwright capability.
- Any backend, proxy, object-storage, cache, API-contract, Product Authority, or Design Authority change.

