# Stage 3 — Frontend Public Campaign Detail Direction

> Phase/Stage: Exploration / Stage 3 — Solutioning  
> Work Unit / Run: `WU-S1-004` / `EXP-FE-001`  
> Author: Codex CLI agent  
> Specialization: Frontend Public Campaign Detail, generated API types, MSW boundary, UI/UX authority, rendered states  
> Created: 2026-09-23  
> Model / Reasoning / Session: `gpt-5.6-terra` / high / Fresh session  
> Target revision: `70ef5c6af1a9` (`validation-03-orchestrator-slice-1`)  
> Workflow revision: not exposed by this Run

## Decision

Use one dynamic public App Router route at `app/campaigns/[campaignId]/`, with route-local detail composition and a smallest-possible Client Component boundary for TanStack Query-driven public server state. The route consumes one focused typed request function for `getPublicCampaignDetail`; the function performs the real same-origin `GET /api/campaigns/{campaignId}` request using the generated operation/schema types. MSW intercepts that exact request in browser mock runtime and in component tests; it is never selected from inside the production request function.

This is the narrowest direction that can render and independently exercise the required asynchronous states before the backend/topology work is available. It keeps the query/data boundary, network path, and response model unchanged when MSW is removed for real integration.

The operation path `/campaigns/{campaignId}` is combined with the canonical server base URL `/api` from `api/openapi/index.yaml`. The returned media `content_url` remains opaque and is rendered only as the contract-supplied same-origin reference; frontend code does not construct storage URLs or normalize it to another origin.

## Options evaluated

| Option | Result | Rationale |
|---|---|---|
| **A. Client query boundary + real same-origin request + browser/node MSW** | **Chosen** | A small Client Component can expose required loading and retryable states. Browser MSW intercepts the same `/api/campaigns/{campaignId}` request used by production, while test MSW shares handlers at the network layer. TanStack Query owns server state rather than duplicating it in local/global storage. |
| B. Server Component directly fetches the public detail | Rejected for this Work Unit | This is viable after real integration, but browser MSW does not intercept a server-side request. It would need additional server-side mock bootstrapping or a different mock route, weakening the single mock-to-real path and adding framework wiring before the backend exists. |
| C. Route-local fixture/conditional mock service returns campaign JSON | Rejected | A mode/environment branch in the production data layer makes mock and real behavior diverge and violates `frontend/AGENTS.md` §2 plus the integration-map contract-parallel rule. |
| D. Add a Next.js API proxy/route handler or a local media store | Rejected | This would duplicate the topology/controlled-media responsibilities of the backend/topology work units and could falsely claim revocability, proxy behavior, or cache enforcement. |

## Settled delivery direction

### 1. Route, ownership, and data state

- Add the public route at `app/campaigns/[campaignId]/`. Keep page composition and detail-specific visual sections route-local. Do not introduce a broad `components/ui`, `components/shared`, Zustand store, Organization request, or discovery surface for this one page.
- Keep the page/server shell as a Server Component where possible; introduce a leaf Client Component only for query lifecycle, retry, and rendering that needs browser data state. Do not hoist `'use client'` to global layout.
- Use a stable campaign-detail query-key factory keyed by the route `campaignId`. Keep returned API data in the query owner; do not copy it into Zustand, a context mirror, or effect-synchronized local state.
- The route parameter is a navigation identity, not a frontend eligibility decision. Encode it as one path segment and let the public contract's response distinguish a successful public Campaign from the deliberately non-disclosing `404`. Do not pre-classify a malformed/non-public identifier with a frontend business rule.
- Do not add polling, optimistic updates, mutation invalidation, donation submission, or an active Donate control. Any cached/revalidating data must not be represented as a newly confirmed financial/system fact.

### 2. Typed request and error boundary

- Add a single low-level frontend request boundary under `frontend/lib/api/`; the focused public-campaign function calls it and imports `operations`/`components` types from `frontend/lib/api/generated/openapi.ts`. No parallel response interface or hand-rolled campaign model is permitted.
- The focused function requests only `GET /api/campaigns/{campaignId}`. It returns the generated `PublicCampaignDetail` success projection and centralizes transport classification for `404`, documented `503`, and unexpected/network failure. It must not turn response bodies, raw Problem Details, lifecycle fields, or HTTP status into invented business facts.
- UI mapping is limited to the settled transport/contract semantics: `404` is the public-safe not-found/not-public presentation; documented `503` is eligible detail unavailable; all other failures are generic safe request failures. Raw `error.message`, Problem `detail`, stack text, and internal headers are not rendered.
- Use the contract's decimal strings and backend-authored relationship as display input. A presentation formatter may group/label the exact string but must not convert with floating-point arithmetic, recompute percentage/progress, cap an above-target value, or infer eligibility.

### 3. Contract-shaped presentation

The presentation reads only the public allowlist and maps its discriminated fields directly:

| Contract input | Surface responsibility | Explicit non-responsibility |
|---|---|---|
| `title`, `purpose`, `story`, `steward`, `lifecycle` | Detail hierarchy; organizer attribution remains visible for organizer text; system/lifecycle meaning is labeled without raw internal state. | No Organization detail fetch, verification badge, raw status/reason, or inferred steward trust claim. |
| `funding` tagged union | Show labeled target/collected/progress relationship when `available`; show a calm explicit unavailable state when not. | No `Number`/float conversion, percentage calculation, donor count, gamified urgency, or outcome/impact implication. |
| `media` tagged union | Render supplied controlled `content_url` and `alt_text` when available; show truthful route-local missing/unavailable treatment for the two no-item states. | No direct object URL, synthetic beneficiary/campaign evidence, or treating unavailable as absent. |
| `donation_action` | Render a truthful non-activating next-action area because Slice 1 always supplies unavailable. | No Donate link, disabled fake flow, donation eligibility reconstruction, or future Slice-2 route. |

Ordinary hierarchy follows the canonical Detail pattern: identity/current state → financial and steward context → organizer story/provenance → media/supporting context → truthful next-action area. It uses the established funding/provenance grammar, semantic labels plus color/structure, and responsive preservation of trust-critical content. Exact Indonesian provenance/action wording remains provisional until Human Design review; the contract's machine terms are not user-facing copy.

The missing-media treatment stays route-local and structural (warm-neutral, explicit placeholder semantics, no fabricated people/documentary scene). It does not create a reusable expressive asset; any move to a reusable or brand-defining asset reopens the stated Human Design gate.

### 4. MSW contract-parallel boundary

- Define shared contract-faithful MSW handlers for the public detail request, using fixtures that `satisfy` the generated public types. Define the controlled media-content handler at the opaque `content_url` path as well, returning only a JPEG/PNG-like mock response when an available-media fixture causes the browser to request it.
- Use those handlers through a browser worker in the mock development/runtime entry and through an MSW node server in Vitest. Handler registration may be environment/test-runtime-specific; the production request function must have no mock conditional, alternate endpoint, or fixture return path.
- Cover the minimum contract states through handlers: success with available media, success with `media.absent`, success with `media.unavailable`, success with unavailable funding, public `404`, documented `503`, and generic network failure. These are test/runtime fixtures for settled states, not new business states.
- Do not claim that MSW proves controlled storage, origin parent/member recheck, `Cache-Control` preservation, response/timing anti-enumeration parity, or proxy configuration. Those remain backend/topology/real-integration evidence.

### 5. Verification contract for `FRONTEND_MOCK_VERIFIED`

| Evidence | Minimum representative conditions | Boundary proved |
|---|---|---|
| Type/lint/unit-component | Generated-type fixture compilation; real typed request/query under MSW; success with available/absent/unavailable media; funding unavailable; loading; `404`; `503`/network failure; unavailable non-CTA; decimal/above-target/zero display; hostile-looking plain organizer strings. Assertions use roles, labels, text, `alt`, and safe user-visible error copy. | Contract use, state distinction, safe escaped text, accessible semantics, no fake action. |
| Rendered inspection | Wide success with media; narrow stress content with absent or unavailable media and long story/financial values; `404`; retryable unavailable/failure. Inspect hierarchy, labels, wrapping, focus visibility, image alternative text, action absence, overflow, and state distinction. | Spatial/responsive/rendered behavior not proven by jsdom. |
| Human rendered acceptance | A human exercises the representative real-browser mock route at proportional desktop and mobile scope, including success and failure/not-public conditions. | Material UI/product-design acceptance; required before delivery/milestone promotion. |
| Deferred real integration | Real `/api` proxy, backend public projection, controlled media bytes/retraction/cache headers, and response/timing parity after WU-S1-003/WU-S1-005 rendezvous. | `INTEGRATED_VERIFIED`, not `FRONTEND_MOCK_VERIFIED`. |

### 6. Browser automation decision

Do **not** add a committed Playwright regression as part of the initial narrow direction. The route has no supported user mutation or complex browser-only interaction, and the required state fidelity is better covered by shared MSW component tests plus representative human rendered acceptance. This keeps the test suite proportional and avoids treating Playwright availability as a phase ritual.

Re-evaluate a browser test if any of these appear during Build/Testing: browser-worker startup or dynamic-route interception is unreliable; an image/content URL regression cannot be represented in jsdom; keyboard/focus or responsive behavior becomes repeatable and materially risky; or a real browser defect is found. Such a change needs an explicit verification rationale in the Techplan rather than automatic inclusion.

## Consequences and carried risks

- The initial route is intentionally client-query-backed to preserve one MSW-to-real network path. It carries a small client bundle/query-provider cost; no server-fetch-only variant is selected while mock-parallel delivery is the active goal.
- Presentation can format exact API decimal strings but must not calculate money/progress. Tests need values such as zero, large, and above-target strings to prevent a convenience conversion from silently returning.
- The route renders the public contract's public `404` with no existence/state distinction. It must not introduce an authenticated/retry diagnostic that defeats this public boundary.
- The final Indonesian labels and a reusable placeholder asset are not settled by this decision. Build may use authority-consistent provisional wording/treatment only for mock evidence; promotion requires the recorded Human Design acceptance.
- No Frontend mock evidence proves backend-controlled media withdrawal or topology security. These risks remain visible in the integration map and must be revisited before integrated completion.

## Implementation scope for Techplan synthesis

Expected production scope is limited to the dynamic public route, its narrow route-local presentation/styles, a query provider only if needed for this route, one typed public API boundary and query hook/key, and MSW browser/node handler bootstrap. Expected test scope is the corresponding behavior tests and mock fixtures. No backend, OpenAPI, topology, Product/MVP, public Organization, Donation, Account, shared-component, or broad design-system change is authorized by this direction.

