# Tech Plan: Slice 1 Public Campaign Detail Frontend Mock-Parallel

> Phase             : Techplan
> Ticket            : `WU-S1-004`
> Work Unit / Run   : `WU-S1-004` / `TP-FE-001`
> Role              : Planner
> Specialization    : Frontend / Next.js / Public Campaign Detail
> Author            : Codex CLI agent
> Model             : `gpt-5.6-terra`
> Reasoning         : high
> Session           : Continue existing healthy `EXP-FE-001` session
> Created           : 2026-09-23
> Target revision   : `a053b48ef3fef35a073c5fe57e5ee4581f848e07`
> Workflow revision : not exposed by this Run
> Status            : Draft
> Approach          : Route detail publik yang memakai satu jalur request nyata bertipe; MSW hanya mengganti respons pada network boundary sampai integrasi nyata tersedia.
> Refs              : `EXP-FE-001/evidence/gap-analysis.md`; `EXP-FE-001/evidence/solutioning.md`; `WU-S1-002/TP-001/techplan.md`; `WU-S1-002/TST-001/testing-report-1.md`; `docs/project/kencleng-integration-map.md`

---

## 1. Background

Slice 1 perlu memberi pengunjung publik yang skeptis tetapi terbuka cukup konteks untuk memahami satu Campaign publik: tujuan, steward, sumber cerita, lifecycle, funding, media, dan keadaan aksi berikutnya. Kontrak Slice-1 sudah `CONTRACT_READY`, tetapi frontend clean-start belum memiliki route Campaign, request boundary, mock network, atau bukti state yang diperlukan.

Work Unit ini membangun bukti frontend secara contract-parallel saat backend dan topology belum tersedia. Bukti tersebut harus cukup untuk klaim `FRONTEND_MOCK_VERIFIED`, tetapi tidak boleh mengklaim bahwa backend projection, same-origin proxy, private storage, cache/revocation media, atau response/timing parity sudah terintegrasi nyata.

## 2. Scope

**In scope:**

- Route publik App Router `frontend/app/campaigns/[campaignId]/` beserta komposisi detail yang dimiliki route tersebut.
- Satu request boundary `GET /api/campaigns/{campaignId}` yang memakai generated OpenAPI types dan klasifikasi transport aman.
- TanStack Query untuk server state detail, tanpa Zustand atau mirror local/global state.
- MSW browser dan Vitest/node yang mengintersepsi request produksi yang sama; fixture harus memenuhi generated public types.
- Presentasi loading, success, media `available`/`absent`/`unavailable`, funding unavailable, public-safe `404`, retryable `503`, dan generic network failure.
- Presentasi funding/provenance/media/action yang jujur, rendering plain text organizer yang aman, aksesibilitas, responsive/rendered evidence, dan Human rendered acceptance.
- Dokumentasi konfigurasi mock browser dan fixture route yang diperlukan agar pemeriksaan browser dapat direproduksi.

**Out of scope (explicit):**

- Backend, OpenAPI/generated type changes, API proxy/rewrite, Caddy/topology, storage, cache, atau controlled-media enforcement nyata.
- Donation Flow, link/tombol Donate aktif, mutation, polling, optimistic update, donation eligibility calculation, dan Campaign Discovery/Organization detail.
- Fetch Organization tambahan, public internal Campaign fields, atau state/lifecycle pasca fundraising yang belum direkonsiliasi untuk Slice 1.
- Shared design-system/component registry, Zustand store, reusable expressive placeholder asset, dan perubahan root layout/global visual system yang tidak dibutuhkan route ini.
- Klaim `INTEGRATED_VERIFIED`; integrasi nyata tetap rendezvous WU-S1-003, WU-S1-005, lalu WU-S1-006.

## 3. Requirements

| ID | Requirement | Source / evidence |
|---|---|---|
| Q1 | Pengunjung publik dapat memahami campaign/steward/funding/source/media/lifecycle/action truth pada loading, success, missing media, not-found, failure, dan responsive states tanpa kebocoran internal. | `docs/product/mvp-delivery-slices.md` §4; `WU-S1-004/manifest.md` Scope |
| Q2 | Surface hanya memakai public-safe `PublicCampaignDetail` dan endpoint detail publik; `404` tidak boleh menjelaskan apakah ID absent, malformed, atau non-public. | `docs/spec/4-campaign/invariants.md` `INV-campaign-14`; `api/openapi/campaign.yaml` `getPublicCampaignDetail` |
| Q3 | Funding, media, dan donation action harus mempertahankan discriminant serta arti yang ditetapkan kontrak; frontend tidak menghitung ulang funding atau eligibility. | `api/openapi/campaign.yaml` public schemas; `EXP-FE-001/evidence/solutioning.md` §2–3 |
| Q4 | Production request memakai generated types dan real same-origin path; MSW bekerja pada network boundary tanpa service mock-mode branch. | `frontend/AGENTS.md` §2; `docs/project/kencleng-integration-map.md` §5–6 |
| Q5 | Material UI memakai Sunlit Editorial / Evidence-Led Optimism, hierarchy Detail, state yang tidak hanya dibedakan warna, dan media/placeholder yang tidak memalsukan evidence. | `docs/ui-ux/design-guidelines.md` §§13, 17–18, 23, 27; `docs/ui-ux/patterns.md` §§3, 10–14, 19 |
| Q6 | Organizer plain text harus tampil sebagai text aman, tanpa raw Problem Details/error/stack ke pengguna; accessibility dan responsive trust-critical information tetap terjaga. | `docs/spec/4-campaign/threat-model.md` public-detail rows; `frontend/AGENTS.md` §§4, 10–12 |
| Q7 | `FRONTEND_MOCK_VERIFIED` hanya dapat diklaim setelah UI produksi, mock contract-faithful, scoped frontend verification, dan Human rendered acceptance selesai; ini bukan integrasi backend nyata. | `docs/kencleng-agentic-workflow.md` §10 |

## 4. Rules & Validation

- **R1 — Real typed detail request.** Given a `campaignId` route segment, when the detail query runs, then it makes exactly the generated-contract-shaped same-origin request `GET /api/campaigns/{encodeURIComponent(campaignId)}` through the single low-level API request boundary and returns only `components["schemas"]["PublicCampaignDetail"]` on `200`.
- **R2 — Mock-to-real path is singular.** Given mock development/browser runtime or Vitest, when the route requests campaign detail or an available image, then MSW intercepts the same `/api/campaigns/{campaignId}` and opaque `content_url` network paths; the production request function neither chooses fixture data nor an alternate endpoint based on mode/environment.
- **R3 — Safe public error distinction.** Given a detail response of `404`, documented `503`, or another/network failure, when the route renders it, then it respectively shows a public-safe not-found/not-public state, a retryable unavailable state, or a generic retryable request-failure state, and never exposes Problem Details body, `error.message`, stack, headers, or a reason for `404`.
- **R4 — Public projection and funding truth.** Given a successful detail response, when campaign content is presented, then only fields in `PublicCampaignDetail` are used; organizer source stays visible, money/percentage decimal strings are display inputs only, zero and above-target facts remain intact, and no float conversion, percentage/progress recomputation, capping, or eligibility inference occurs.
- **R5 — Media truth.** Given `media.state` of `available`, `absent`, or `unavailable`, when media is rendered, then available items use their supplied opaque same-origin `content_url`, `alt_text`, nullable caption, and organizer source; absent and unavailable have visibly distinct, non-fabricated treatments; the latter is never represented as absence.
- **R6 — No simulated donation action.** Given Slice-1 `donation_action` is unavailable, when the next-action area renders, then it gives truthful non-activating context and contains no Donate link, form, disabled fake control, future-route target, or client-side donation eligibility decision.
- **R7 — Observable accessible async UI.** Given loading, success, not-found, unavailable/failure, and retry transitions, when a keyboard or assistive-technology user uses the route, then semantic headings/landmarks, named retry controls, announced state changes, visible focus, meaningful image alternative text, and text/structure in addition to color communicate the state without focus becoming lost on replacement.
- **R8 — Responsive evidence hierarchy.** Given wide and narrow viewports with long title/story/steward text and large/zero/above-target decimal values, when layout reflows, then campaign identity, source, funding labels/value relationship, state meaning, media status, and next-action context remain readable, reachable, and free of unintended clipping or horizontal overflow.
- **R9 — Honest frontend milestone boundary.** Given frontend verification evidence is collected, when the Work Unit is handed off, then it identifies mock/network/route behavior actually verified and explicitly defers live backend projection, proxy, storage/retraction, cache preservation, and response/timing parity to their owning integration work.

## 5. Decision Log

| ID | Decision / option | Status | Rationale / consequence |
|---|---|---|---|
| D1 | Dynamic route `app/campaigns/[campaignId]/` with route-local composition and smallest leaf Client Component for TanStack Query. | Chosen | Keeps bookmarkable identity in the URL and browser server-state behavior confined to the smallest meaningful boundary. The page shell/layout remains server-capable. |
| D1-alt | Fetch detail directly in a Server Component. | Rejected for this Work Unit | Browser MSW would not intercept server fetches without a second mock path/server bootstrapping, weakening mock-to-real correspondence before backend availability. |
| D2 | One focused typed API function through a centralized low-level request helper. | Chosen | Preserves explicit generated-contract request/response ownership and one transport classification point. It adds no auth/eligibility semantics absent from this public read operation. |
| D2-alt | Route/component `fetch`, handwritten interfaces, or mocked service return branch. | Rejected | Fragments error handling, duplicates an OpenAPI model, or violates the contract-parallel rule. |
| D3 | MSW browser worker plus node server/handlers, enabled only by runtime/test bootstrap outside production request code. | Chosen | Tests and browser mock runtime exercise the same HTTP path. Mock enablement is infrastructure configuration, not a production API behavior branch. |
| D3-alt | Next API proxy/route handler, local media store, or direct fixture rendering. | Rejected | Would duplicate WU-S1-003/WU-S1-005 responsibilities or conceal the real network contract. |
| D4 | Render available campaign media with native `<img src={content_url}>`; mock the controlled path with JPEG/PNG content. | Chosen | The contract-supplied relative URL remains opaque and is fetched in the browser at its controlled same-origin path. Avoiding an image optimizer/proxy prevents frontend from introducing a second server-side media path outside this Work Unit. |
| D5 | Query key factory with one detail key per `campaignId`; disable automatic retries/polling and expose only explicit user retry for retryable states. | Chosen | Prevents ad-hoc cache keys and avoids silently re-requesting consequential funding/status data. No mutation/invalidation behavior is needed. |
| D6 | Funding/progress uses API decimal strings and backend relationship as display input. | Chosen | Preserves money precision and public funding meaning; formatting may group/label text but must not calculate or cap values. |
| D7 | Route-local structural missing-media treatment, no generated/reusable expressive placeholder asset. | Chosen | Satisfies truthful missing-media state while the exact reusable placeholder remains an intentional Human Design decision. |
| D8 | No initial committed Playwright regression. | Chosen | The route has no mutation/complex browser-only flow; RTL + node MSW protect the contract behavior, while rendered inspection and Human acceptance cover spatial/product judgment. Reconsider only if browser-worker/image/focus/responsive defects prove a repeatable browser regression warrants it. |

## 6. Backward Compatibility

- **Existing data:** N/A. This frontend Work Unit introduces no persistence, migration, or data transformation.
- **API/contracts/clients:** No API change. It consumes the current generated `PublicCampaignDetail` type and current public operation only; no existing Campaign UI/consumer exists in the clean-start frontend.
- **Migration/deprecation compatibility:** N/A. The mock browser bootstrap is explicitly opt-in configuration; removing that development setting later exposes the unchanged production request path for real integration. No mock branch is permitted in `lib/api/`.

## 7. Edge Cases & Risks

| ID | Risk / edge case | Likelihood | Severity | Mitigation / accepted exposure |
|---|---|---:|---:|---|
| RISK-1 | Internal/historical Campaign or Organization fields leak into a public surface. | Medium | High | Type fixtures and view consume only `PublicCampaignDetail`; no secondary Organization request; tests assert public observable content and no invented internal/lifecycle claims. |
| RISK-2 | `404` messaging discloses why a Campaign is unavailable, or raw failure data leaks internals. | Medium | High | One safe `404` presentation; `503` and generic failures use fixed user-facing copy and retry only. Never render raw error/Problem fields. |
| RISK-3 | Funding is changed by float conversion/recalculation, zero is treated as no data, or over-target is capped. | Medium | High | Generated string types, formatter-only display boundary, and fixtures/tests for zero, large, `above_target`, and `not_computable`. |
| RISK-4 | Media absent, unavailable, and available are collapsed or image handling bypasses the controlled URL. | Medium | High | Discriminated state rendering; native image uses supplied `content_url`; browser/node handlers include matching content route. Real revocation/cache proof remains deferred. |
| RISK-5 | Organizer text becomes markup/XSS or presentation implies platform verification. | Medium | High | Render as ordinary React text only; no `dangerouslySetInnerHTML`; hostile-looking text fixture verifies literal escaped output and organizer attribution. |
| RISK-6 | Browser worker starts after the first query, producing flaky unmocked requests. | Medium | Medium | Route-local mock bootstrap gates its children until `worker.start()` resolves when the explicit mock flag is enabled; Vitest uses node server lifecycle independently. |
| RISK-7 | Long content/narrow layout hides funding/provenance/action meaning or loses keyboard focus after async replacement. | Medium | Medium | Intrinsic layout/meaningful reflow; semantic live/state treatment; focused component assertions plus wide/narrow rendered inspection with stress fixtures. |
| RISK-8 | Mock evidence is mistaken for live backend/topology integration. | Medium | High | Report exact mock boundary; retain Test Focus pointers and explicit deferred dependencies. Do not alter tracker/control surface from this Work Unit. |
| RISK-9 | Provisional Indonesian provenance/action copy or a route-local placeholder is treated as final/reusable brand precedent. | Medium | Medium | Keep wording/treatment deliberately provisional and structural; Human Design review is an Active Open Item before promotion. |

## 8. Interface Contract

**Persistence/data shape:**

N/A. The frontend stores no campaign data. TanStack Query owns the transient cached server response under a stable `campaignKeys.detail(campaignId)` key; it is not mirrored into Zustand, context, or effect-synchronized state.

**API/event/external interface:**

- `getPublicCampaignDetail` is the sole JSON data call: `GET /api/campaigns/{campaignId}`, derived from API server base `/api` plus `getPublicCampaignDetail` path `/campaigns/{campaignId}`. The request boundary path-encodes a single supplied route segment but does not validate/classify public eligibility locally.
- `200` returns `components["schemas"]["PublicCampaignDetail"]`; only this generated type is imported for the detail response. `404` maps to a typed safe not-found outcome, `503` maps to typed temporary-unavailable outcome, and all other non-`200`/network failures map to a generic safe failure outcome. Problem Details are not parsed into user copy.
- Available media is not manually fetched as JSON. Native image loading uses the `PublicCampaignMediaItem.content_url` as supplied. The MSW content handler accepts that opaque `/api/campaigns/{campaignId}/media/{mediaId}/content` path and returns only contract-compatible JPEG/PNG mock content.
- Browser mock activation uses an explicitly documented public development flag (for example `NEXT_PUBLIC_MSW_ENABLED=true`). When inactive, the bootstrap must render children immediately and the same API function reaches the actual same-origin network. Vitest always uses the node MSW server; it does not alter the API function.

**Cross-layer/business boundary:**

- Backend alone decides public eligibility, the closed allowlist, decimal/progress computation, funding/media availability, parent/member media access, cache headers, and donation eligibility. Frontend formats and presents only the returned public truth.
- `404` is intentionally non-disclosing. The frontend must not add a privileged diagnostic, auth variant, status-specific explanation, or retry state that claims it can determine non-publicness.
- The returned `content_url` is an opaque controlled reference, not storage knowledge. WU-S1-004 can prove that UI requests it under MSW, not that backend origin rechecks, private storage, retraction, `Cache-Control: private, no-store`, proxy routing, or timing parity work in reality.

## 9. Architecture / Plan

The route has one data/control path in both mock and real use:

```text
Server route shell: app/campaigns/[campaignId]/page.tsx
    -> route-local client boundary and QueryClient provider
    -> usePublicCampaignDetail(campaignId)
    -> getPublicCampaignDetail(campaignId)
    -> apiRequest('/api/campaigns/{campaignId}')
    -> browser fetch
       -> MSW worker (explicit mock browser runtime only)
       -> real same-origin backend/proxy (later integration)
```

The server route supplies the navigation identity only. The client boundary owns TanStack Query lifecycle and chooses the visual state from one query result. The view is route-local and derives all presentation from the query data/discriminants rather than storing copies. A route-local mock-worker gate starts before its children only when development mock runtime is explicitly enabled, preventing the first request from escaping interception.

Build sequence:

1. Establish generated-type aliases, centralized public read request/error outcome, key factory, and query hook before visual composition.
2. Add route shell, local query provider/client composition, and contract-shaped visual states. Use CSS module/local styling plus the current global typography/color foundation; do not turn root layout into a Client Component.
3. Add contract-faithful fixture set and shared MSW handlers/browser/node bootstrap. Commit the generated worker asset and document the opt-in browser mock command/configuration plus representative fixture URLs.
4. Add API/component tests through node MSW, then perform Build-time rendered feedback for representative states. Independent Testing repeats the selected checks and rendered inspection; Human performs final rendered acceptance before a milestone is claimed.

## 10. Implementation Details

| Anchor | Why relevant | Intended change / precedent |
|---|---|---|
| `frontend/app/campaigns/[campaignId]/page.tsx` — new route | No Campaign route exists; the dynamic segment is the bookmarkable public identity. | Add a Server Component shell that passes `campaignId` to the route-local client detail boundary. It performs no server data fetch or public eligibility decision. |
| `frontend/app/campaigns/[campaignId]/layout.tsx` and `mock-service-worker.tsx` — new route-local bootstrap | Browser MSW must start before a client query, but mock activation must not live in `lib/api/`. | Add the smallest Client wrapper that gates descendants on `worker.start()` only when the documented mock flag is on; otherwise immediately render children. Keep it within this route tree. |
| `frontend/app/campaigns/[campaignId]/campaign-detail-client.tsx` — new leaf Client Component | TanStack Query/browser state requires client code; whole route/layout must not become client code. | Own a route-local QueryClient/provider and query lifecycle. Use `campaignKeys.detail(campaignId)`, `retry: false`, no polling, and explicit retry only for unavailable/generic failure. |
| `frontend/app/campaigns/[campaignId]/campaign-detail-view.tsx` and `campaign-detail.module.css` — new route-local presentation | The clean tree has no reusable component contract; this is one route's public detail semantics. | Render semantic Detail hierarchy and all R3–R8 states. Use native image for available media and structural route-local placeholder treatments. Preserve Newsreader/Instrument Sans roles, warm-neutral/border-first grammar, non-gamified funding, provenance, and no color-only meaning. |
| `frontend/lib/api/client.ts` — new low-level helper | No current request convention exists; `frontend/AGENTS.md` requires production access to retain real API shape. | Centralize raw `fetch`, basic safe response handling, and transport-error construction. Do not add an auth, CSRF, refresh, mock, or proxy convention that the public read contract/current frontend has not established. |
| `frontend/lib/api/public-campaign.ts` — new focused function | Generated types currently exist but no consumer does. | Export generated-type aliases and `getPublicCampaignDetail`; build only the `/api/campaigns/${encodeURIComponent(campaignId)}` request and status classification in §8. Do not define handwritten response interfaces. |
| `frontend/lib/hooks/use-public-campaign-detail.ts` — new query key/hook | Query ownership/key consistency must be established at the first consumer. | Export `campaignKeys.detail` and one focused `usePublicCampaignDetail` wrapper; preserve API data in query cache only. |
| `frontend/mocks/fixtures/public-campaign.ts` — new fixtures | Browser and node tests need the same contract-faithful state examples. | Export generated-type-checked fixture builders/IDs for available media, media absent, media unavailable, funding unavailable, zero/large/above-target/not-computable values, and hostile organizer plain text. Use stable fixture IDs solely to select mock scenarios, not as product semantics. |
| `frontend/mocks/handlers/public-campaign.ts`, `browser.ts`, `server.ts` — new MSW boundary | The mock path must be network-layer and shared across browser/tests. | Implement detail and controlled-media handlers using exact same-origin paths; export browser worker and Vitest node server. Support public `404`, `503`, and test-overridable generic network failure without fixture branches in production code. |
| `frontend/public/mockServiceWorker.js` — generated MSW worker | Browser interception requires the MSW worker asset. | Generate with the installed MSW CLI and commit the generated worker unchanged; do not hand-author a service worker or storage/proxy substitute. |
| `frontend/vitest.setup.ts` — existing test setup | It currently installs jest-dom only. | Add node-server lifecycle setup/reset/close once the MSW server exists, preserving existing matcher setup and excluding Playwright tests. |
| `frontend/app/campaigns/[campaignId]/campaign-detail-client.test.tsx` and `frontend/lib/api/public-campaign.test.ts` — new tests | No Slice-1 behavior test exists. | Exercise real query/API path with node MSW and observable RTL assertions for R1–R8. Keep test helpers local; do not mock the hook/client or test class names/internal state. |
| `frontend/README.md` — mock runtime section | Human/Testing needs a reproducible browser mock entrypoint. | Document opt-in development flag, worker asset generation/commit expectation, representative fixture route URLs, and the explicit boundary between mock evidence and real integration. |
| `frontend/app/layout.tsx` — `RootLayout` | Current global font/language context is already correct. | Keep server-side; do not add global query/mock provider for a one-route feature. |
| `frontend/lib/api/generated/openapi.ts` — generated public declarations | Canonical TypeScript response shape source. | Read/import only; never hand-edit or regenerate because this Work Unit changes no OpenAPI authority. |

## 11. Files Changed / Files NOT Changed

| File / area | Change type | Description |
|---|---|---|
| `frontend/app/campaigns/[campaignId]/page.tsx` | Add | Public detail route shell. |
| `frontend/app/campaigns/[campaignId]/layout.tsx`, `mock-service-worker.tsx` | Add | Route-local browser MSW initialization gate. |
| `frontend/app/campaigns/[campaignId]/campaign-detail-client.tsx`, `campaign-detail-view.tsx`, `campaign-detail.module.css` | Add | Query lifecycle and route-local detail/state presentation. |
| `frontend/lib/api/client.ts`, `frontend/lib/api/public-campaign.ts` | Add | Central low-level public read boundary and generated-type consumer. |
| `frontend/lib/hooks/use-public-campaign-detail.ts` | Add | Stable detail query key and hook. |
| `frontend/mocks/fixtures/public-campaign.ts`, `frontend/mocks/handlers/public-campaign.ts`, `frontend/mocks/browser.ts`, `frontend/mocks/server.ts` | Add | Shared contract-faithful fixture and MSW network boundary. |
| `frontend/public/mockServiceWorker.js` | Add/generated | MSW browser worker asset. |
| `frontend/vitest.setup.ts` | Modify | Node MSW lifecycle alongside existing jest-dom setup. |
| `frontend/app/campaigns/[campaignId]/campaign-detail-client.test.tsx`, `frontend/lib/api/public-campaign.test.ts` | Add | Focused contract/state behavior tests. |
| `frontend/README.md` | Modify | Reproducible browser mock fixture/runtime instructions and integration-limit wording. |

| File / area intentionally untouched | Why |
|---|---|
| `backend/**`, `Caddyfile`, `docker-compose.yml`, topology/proxy files | Backend projection, private media storage, and same-origin topology are separate Work Units; protected backend paths remain fenced. |
| `api/openapi/**`, `api/openapi.yaml`, `frontend/lib/api/generated/openapi.ts` | Contract is current `CONTRACT_READY`; this Work Unit consumes it and does not change/regenerate it. |
| `frontend/app/layout.tsx`, `frontend/app/globals.css` | Existing root font/language/global foundation suffices; styles belong to the route CSS module and mock provider must not become global. |
| `frontend/components/ui/**`, `frontend/components/shared/**`, `frontend/lib/stores/**` | One surface does not establish a broad shared component or client-state contract. |
| `docs/product/**`, `docs/ui-ux/**`, Campaign specs/invariants/threat model | Upstream product/design/contract authority is sufficient; the remaining Human Design decisions must not be silently rewritten here. |
| `.harscode-spaces/**/manifest.md`, `control-surface.md`, development tracker | Orchestration/human authority owns status transitions; Build must not self-promote `FRONTEND_MOCK_VERIFIED`. |

## 12. Testing Checklist

| Rule | Verification / evidence | Primary owner | Why this is worth running / risk if skipped |
|---|---|---|---|
| R1 | API-boundary test and component test use node MSW to assert the exact `/api/campaigns/{campaignId}` request, generated fixture response, and route-driven query key. Run focused Vitest test while authoring. | Build | Proves the route does not silently use a fixture/direct component path or wrong endpoint; request drift would defeat later integration. |
| R2 | Verify browser/node handlers receive the same detail path and available image `content_url`; inspect `lib/api/` for no mock flag, alternate URL, or fixture import. Test mock-worker gate in browser runtime during rendered feedback. | Testing | Network-layer correspondence is the central mock-parallel contract. Skipping it could leave a green UI that cannot make the real request. |
| R3 | RTL/MSW cases for `404`, `503`, and rejected network request; assert distinct safe headings/copy, named retry only for retryable cases, and absence of injected raw error/Problem text. | Build | Protects the public anti-enumeration presentation and prevents internal-error leakage. |
| R4 | RTL fixtures cover organizer attribution; `0.00`, large values, `above_target`, and `not_computable`; assert displayed supplied values/relationship and no absent funding content is fabricated. | Testing | Money and public projection errors are truth/security failures, not formatting polish. |
| R5 | RTL/MSW cases for available, absent, and unavailable media; assert exact image `src`/`alt`, nullable caption behavior, source treatment, and visibly distinct no-item states. Render available image handler with JPEG/PNG content in a real browser. | Testing | Distinguishes public media semantics and catches an accidental direct-object URL/false-absence mapping. |
| R6 | RTL checks the next-action region has truthful unavailable context and no link/button/form whose accessible name/action invites donation. | Build | A disabled or fake Donate CTA would violate the active Slice boundary despite looking harmless. |
| R7 | RTL assertions use landmarks/headings/roles/accessible names, live loading/error status, retry keyboard activation, focus-visible styling/retained meaningful focus, and literal rendering of hostile markup-like organizer text. | Testing | jsdom catches structural accessibility, unsafe rendering, and lost-control regressions before human review. |
| R8 | During Build and independent Testing, render wide success with available media and narrow stress fixture with long content plus absent/unavailable media; inspect wrapping, labels, reachability, overflow, image alt, focus visibility, and state differentiation. | Testing | Unit tests cannot prove spatial hierarchy/responsive behavior; skipping this risks hiding trust-critical facts on mobile. |
| R9 | Review the final Build/Testing evidence and Human acceptance record against `docs/kencleng-agentic-workflow.md` §10. Require explicit deferred real-integration risks and no tracker/control-surface self-promotion. | Human | The milestone meaning and material UI acceptance require human authority; mock tests cannot prove integration or subjective design correctness. |
| R1–R9 | `cd frontend && npm run lint`, focused/new `npm run test`, then `npm run verify`; independent Testing runs `npm run build`, `git diff --check`, and scoped status review. | Build for fast edit-loop; Testing for final independent evidence | These catch type/lint/test/build integration and accidental scope expansion. A production build is deferred from Build to avoid duplicate broad-suite cost, but Testing must run it because new route/client/browser-boundary modules exercise Next compilation. |

### Test Focus Pointer

| Area | Why sensitive | Evidence anchor from Exploration | Still relevant post-synthesis? |
|---|---|---|---|
| Public projection regression | A public view can accidentally use internal/historical fields or invent semantics. | `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/EXP-FE-001/evidence/gap-analysis.md#area-1--product-reconciled-campaign-behavior-and-public-api-contract` | Yes — R1/R4 and RISK-1 cover frontend contract consumption; backend negative projection mapping remains WU-S1-003 Testing. |
| Existence disclosure / optional-auth variance | A detailed 404 presentation or variant can reveal why public access failed. | `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/EXP-FE-001/evidence/gap-analysis.md#area-1--product-reconciled-campaign-behavior-and-public-api-contract` | Yes — R3 covers the frontend safe presentation. Runtime auth/body/header/timing parity is deferred to backend/integration Testing. |
| Media revocation and cache staleness | A controlled path can be bypassed or proxy/cache semantics falsely assumed. | `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/EXP-FE-001/evidence/gap-analysis.md#area-4--generated-types-network-boundary-msw-and-live-integration-dependency` | N/A for real-enforcement testing — D4/R9 prove only supplied-path rendering under MSW; origin recheck, withdrawal, cache headers, and proxy behavior belong to WU-S1-003/WU-S1-005/WU-S1-006. |
| Money truth | Client conversion/recomputation can misstate financial facts. | `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/EXP-FE-001/evidence/gap-analysis.md#area-1--product-reconciled-campaign-behavior-and-public-api-contract` | Yes — R4/RISK-3 needs focused frontend display tests; backend decimal calculation remains separate. |
| Organizer-controlled content | Unsafe plain-text rendering can create stored XSS or false provenance. | `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/EXP-FE-001/evidence/gap-analysis.md#area-5--observable-state-hostile-content-accessibility-and-rendered-evidence` | Yes — R7/RISK-5 requires hostile-looking text and attribution evidence. |
| Runtime concurrency/performance | This surface has no mutation, polling, shared mutable state, or new hot-path algorithm. | `.harscode-spaces/s1-public-campaign-understanding/WU-S1-004/runs/EXP-FE-001/evidence/solutioning.md#consequences-and-carried-risks` | N/A — D5 deliberately disables automatic retry/polling; no specialized concurrency/load test is proportionate. |

## 13. Open Items

### Active — needs external input or verification

1. **Final Indonesian provenance and unavailable-action wording — Human Design decision; non-blocking for Build, blocking for milestone promotion/delivery.** Build may use clearly provisional, authority-consistent copy for mock evidence. It must not expose machine literals such as `organizer`, `unavailable`, or `donation_flow_not_available` as final user language. Human Design review must approve the promoted wording.
2. **Exact reusable campaign placeholder asset/system — Human Design decision; non-blocking while treatment remains route-local and structural.** Build must not create/promote a reusable or expressive asset. If implementation reveals that a reusable asset is necessary, stop and obtain the required Design decision before creating it.
3. **Human rendered acceptance — required before `FRONTEND_MOCK_VERIFIED`.** A human must exercise representative desktop and mobile mock routes for success/media, not-found, and retryable unavailable/failure after independent Testing convergence.
4. **Live integration evidence — deferred external dependency.** WU-S1-003/WU-S1-005/WU-S1-006 must later prove actual backend projection, controlled byte delivery, same-origin routing, `private, no-store`, storage/retraction, and response/timing parity. This item does not block frontend mock Build but blocks `INTEGRATED_VERIFIED`/slice delivery.

### Resolved — retained as decision history

1. ~~**Route/data ownership direction**~~ **RESOLVED — D1/D5 select a dynamic route with a smallest Client query boundary, stable query key, no Zustand mirror, no polling, and manual retry only.**
2. ~~**Mock-to-real request boundary**~~ **RESOLVED — D2/D3 require a generated-type-backed real same-origin request and MSW interception only at browser/node network boundaries.**
3. ~~**Media delivery representation**~~ **RESOLVED — D4 renders supplied opaque `content_url` through native image loading and mocks its controlled JPEG/PNG path without introducing proxy/storage behavior.**
4. ~~**Initial browser automation choice**~~ **RESOLVED — D8 skips a committed Playwright regression as disproportionate; independent rendered inspection and required Human acceptance remain mandatory, with explicit re-evaluation triggers.**
