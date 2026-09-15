# Tech Plan: Frontend Experience Foundation

> Ticket    : 00-foundations/01-frontend-experience-foundation
> Author    : Codex
> Date      : 2026-09-15
> Updated   : 2026-09-15
> Status    : Approved - Anhar
> Approach  : Calibrate the existing `/` slice through local, truth-preserving route work and representative browser evidence; do not expand into a landing-page or Campaign delivery.
> Refs      : `docs/spec/0-foundations/features/01-frontend-experience-foundation.md`; `1-exploration/logs/01-design-and-route-authority.md`; `1-exploration/logs/02-live-home-route-and-components.md`; `1-exploration/logs/03-verification-and-rendering.md`; `1-exploration/logs/04-ownership-and-assets.md`; `1-exploration/logs/05-solutioning-and-decisions.md`

---

## 1. Background

`/` already has a public shell, text-first hero, mock-backed campaign collection,
and explanatory section, but it lacks durable desktop/mobile rendered evidence
for the foundation acceptance criterion. Exploration also found public copy
that conflicts with the active donation specification, route-only sections
located in a broader component layer, an English document-language declaration
for Indonesian UI copy, and no `/` browser test. This work establishes a
limited, reviewable foundation without implying unfinished Campaign behavior or
brand assets are complete.

## 2. Scope

**In scope:**

- Calibrate the existing representative `/` composition using existing tokens
  and local route components; address objective rendered hierarchy/responsive
  defects found during Build without creating a new visual narrative.
- Replace the explanatory section's unsupported or conflicting trust/money/
  reporting language with the fixed authority-bounded copy defined in §10.
- Move the three home-only sections from `components/features/landing/` into
  the public route's `_components/` owner, preserving the existing server/client
  boundary and test coverage.
- Change the application document language to Indonesian (`lang="id"`).
- Add deterministic browser coverage for the populated `/` route at desktop and
  mobile widths, including the hero-to-collection CTA and mobile drawer.
- Retain/extend component tests for loading, populated, empty, and error list
  states as the deterministic evidence for asynchronous exceptional behavior.
- Obtain human rendered acceptance at representative desktop and mobile scope.

**Out of scope (explicit):**

- Completing the landing page, adding a footer, creating `/campaign` or a
  campaign-detail route, changing Campaign APIs/data/mocks, or adding live
  Campaign integration.
- Defining highlighted/featured/sorted campaign semantics; adding organization,
  media, ranking, recommendation, verification-score, or impact fields to the
  list contract or fixtures.
- New global tokens, `ui`/`shared` variants, or changes to `Button`, `Badge`,
  or `ProgressBar` contracts.
- Hero key art, a canonical logo, synthetic campaign photography, or a new
  campaign-placeholder asset system.
- Backend, API, persistence, migration, cron, deployment, or runbook changes.

## 3. Requirements

| ID | Requirement | Source / evidence |
|---|---|---|
| Q1 | The representative `/` slice must let a human judge coherent public brand/experience at desktop and mobile scope, without becoming a complete landing page. | Feature spec — Feature surface, Requirements, Acceptance criteria; Exploration `05-solutioning-and-decisions.md#Chosen direction` |
| Q2 | Public content must not claim unsupported ranking, impact, verification detail, money terms, fees, or reporting behavior. | Feature spec — Requirements; `product-design-principles.md` A.1, A.7; Exploration `02-live-home-route-and-components.md#Gap` |
| Q3 | Campaign collection behavior remains a fixed mock-backed `GET /campaigns` list with its established loading, success, empty, error, and stale states; no Campaign contract is invented. | `page-map.md` Guest footnote; `patterns.md` B.1; Exploration `02-live-home-route-and-components.md#Current state` |
| Q4 | Ownership stays semantic-owner-first; route-only composition is route-local and broad primitives are not changed for page polish. | `frontend/AGENTS.md` §§1,5; `frontend/components/README.md` A, G–L; Exploration `04-ownership-and-assets.md#Gap` |
| Q5 | Material UI requires real rendered desktop/mobile feedback, useful requested browser automation, and human rendered acceptance. | Feature spec — Verification expectations; `frontend/AGENTS.md` §§10–11; Exploration `03-verification-and-rendering.md#Requirement` |
| Q6 | Public mobile navigation maintains native control semantics, reachable drawer behavior, and focus return on Escape; responsive layouts must not introduce horizontal overflow. | `frontend/AGENTS.md` §§10–12; Harscode `best-practices/react/accessibility-fundamentals.md`, `responsive-layout-robustness.md`; Exploration `03-verification-and-rendering.md#Requirement` |

## 4. Rules & Validation

- **R1** — Given a guest opens `/` at 1280×800 or 390×844, when the populated mock collection renders, then the public shell, hero heading, primary browse CTA, campaign section, and supporting explainer are visible without horizontal document overflow.
- **R2** — Given the home hero CTA is activated, when it targets the current in-page campaign collection, then focus/scroll reaches the `#kampanye` section; it must not navigate to an absent Campaign route.
- **R3** — Given the mobile viewport, when the hamburger opens the drawer and the user activates Escape, then the drawer uses its existing dialog/focus-trap behavior and focus returns to the hamburger; mobile navigation and auth controls remain reachable.
- **R4** — Given the campaign request is loading, succeeds, returns an empty list, or fails, when `HighlightedCampaigns` renders, then it preserves its shaped skeleton, safe retryable error, no-CTA empty state, and contract-limited populated cards; no raw backend error, invented list semantics, organization detail, or media is rendered.
- **R5** — Given the static home explanatory copy is rendered, then it does not state a Rp10,000 minimum, no-hidden-fees promise, periodic-report promise, invented trust score/ranking, or unsupported campaign-specific verification detail. It uses only the fixed copy/claims in §10.
- **R6** — Given the application root renders Indonesian public copy, then the document language is `id`.
- **R7** — Given the route-only home sections are relocated, when `/` renders, then its section order and the client-only campaign-query boundary are unchanged; `CampaignCard` remains campaign-owned and shared primitives retain their current contracts.
- **R8** — Given the finished representative slice is reviewed in a real browser at desktop and mobile scope, then a human explicitly accepts or rejects its hierarchy, CTA clarity, responsive usability, truthful/provisional presentation, and placeholder appropriateness. Automated evidence cannot satisfy this rule alone.

## 5. Decision Log

| ID | Decision / option | Status | Rationale / consequence |
|---|---|---|---|
| D1 | Keep the minimum slice: public shell + text-first hero + mock-backed campaign list + compact explainer. | Chosen | Calibrates the required brand/hierarchy/CTA/card/responsive surface without completing the landing page. |
| D1-alt | Complete or substantially redesign the landing page. | Rejected | Violates explicit scope and would require an unapproved landing story and likely broader asset/component work. |
| D2 | Preserve the current in-page `#kampanye` CTA target and non-interactive campaign cards during this task. | Chosen, provisional | The Campaign list/detail routes are absent. The anchor makes the current browse CTA functional without fabricating route behavior. This does not resolve the separate page-map conflict in Open Item 1. |
| D2-alt | Link navigation/cards to `/campaign` or a detail route now. | Rejected | Would create a broken path or expand into Campaign delivery, which is `NOT_STARTED`. |
| D3 | Replace unsupported explainer claims with authority-bounded static copy; remove the amount, fee, and periodic-report promises. | Chosen | The active donation spec sets a `≥ 5000` floor and does not support the current Rp10,000, no-fee, or reporting-frequency claims. Removing them avoids false money/accountability representation. |
| D4 | Keep verification language only as a platform-level published-campaign policy; never add per-card verification data/badges. | Chosen | Verified-org creation and auto-unpublish on status reversal support the policy-level claim, while the list response cannot supply organization identity/status. |
| D5 | Move home-only sections to `app/(public)/_components/`; retain `CampaignCard` in `features/campaign/`. | Chosen | Matches semantic-owner-first governance. No evidence establishes reusable/domain semantics for Hero, HighlightedCampaigns, or HowItWorks; CampaignCard has campaign semantics and anticipated reuse. |
| D6 | Do not alter global tokens or shared/UI primitives. | Chosen | Existing tokens and primitives suffice; a global visual change would have unneeded cross-product blast radius. |
| D6-alt | Add primitive variants or global token changes for home-page polish. | Rejected | No evidence of a cross-product contract need; `Button`, `Badge`, and `ProgressBar` have established consumer/contract obligations. |
| D7 | Retain the neutral icon-based no-media card treatment as provisional; do not create an asset. | Chosen | It preserves truth while a permanent placeholder/key-art/logo direction requires separate asset exploration and applicable human approval. |
| D7-alt | Generate hero key art, a logo, synthetic campaign photos, or a new shared placeholder now. | Rejected | These assets risk asserting documentary/brand truth not approved for this calibration work. |
| D8 | Browser-test the populated route/responsive and drawer flows; keep deterministic list exceptional-state coverage in Vitest/MSW. | Chosen | Existing browser MSW handlers expose only the populated fixture, while component tests already deterministically exercise empty/error states. This avoids adding a test-only runtime interface merely to force browser failures. |

## 6. Backward Compatibility

- **Existing data:** No persistence or generated schema change. Existing MSW campaign fixture shape and `CampaignListItem` mapping remain unchanged.
- **API/contracts/clients:** No API request/response, route, auth, or shared-component contract change. `GET /campaigns` remains the sole data boundary for the collection.
- **UI consumers:** Moving route-only files changes internal import paths only; preserve the `Home` output order and the `HighlightedCampaigns` client boundary. `CampaignCard`, `Button`, `Badge`, and `ProgressBar` retain their current public contracts.
- **Migration/deprecation compatibility:** None. No independently operable script, migration, cron, cleanup, monitoring, or runbook concern exists.

## 7. Edge Cases & Risks

| ID | Risk / edge case | Likelihood | Severity | Mitigation / accepted exposure |
|---|---|---:|---:|---|
| RISK-1 | Unsupported money, fee, reporting, ranking, or verification claims mislead a public donor. | Medium | High | Apply D3/D4; retain contract-limited cards and test the prohibited-copy boundary. |
| RISK-2 | A `/campaign` link is broken or the page-map authority is silently overwritten. | High | Medium | Do not change public nav target in this task; retain Open Item 1 for human/project decision. |
| RISK-3 | Desktop/mobile output looks coherent in component tests but has clipped/overflowing content or unusable drawer focus in a browser. | Medium | Medium | Apply R1–R3 browser checks plus R8 human acceptance. |
| RISK-4 | Moving files accidentally changes the server/client boundary or list-state coverage. | Low | Medium | Preserve `HighlightedCampaigns` as the client leaf and run R4/R7 tests. |
| RISK-5 | Neutral missing-media treatment is mistaken for an approved permanent asset or replaced with fabricated imagery. | Medium | Medium | Apply D7; human acceptance evaluates provisional clarity only. Asset-system work is deferred. |

## 8. Interface Contract

**Persistence/data shape:** No change. Continue to consume generated `CampaignListResponse`/`CampaignListItem`; do not add client models, fields, money calculations, media, organization details, or featured/sort parameters.

**API/event/external interface:** No change. The collection continues calling public `GET /campaigns`; no route, event, browser query parameter, or mock-only runtime interface is added. Browser tests use the current dev/MSW populated fixture; Vitest/MSW remains the exceptional-state seam.

**Cross-layer/business boundary:** Frontend owns presentation, static copy, responsive composition, and client list-state rendering only. Campaign selection meaning, organization identity/status returned by the list, payment fees, report frequency, and all money/business semantics remain domain/API authority. The list contract does not authorize per-card verification representation.

## 9. Architecture / Plan

1. Reopen current authority and affected source anchors; keep the public nav item untouched pending Open Item 1.
2. Move only `Hero`, `HighlightedCampaigns`, and `HowItWorks` plus their colocated tests to `app/(public)/_components/`; update `Home` imports. Keep `HighlightedCampaigns` as the only data-fetching client leaf and retain `CampaignCard` where it is.
3. Apply the fixed explanatory-copy boundary and `lang="id"`; make only local token/class composition adjustments necessary to correct objective rendered failures at the two representative sizes.
4. Add the route browser regression alongside the existing login smoke precedent. It validates only the populated mock route and public-shell interaction that can run deterministically in the current browser harness.
5. Run component/unit, lint/build, and requested browser checks; inspect the rendered route during Build; then obtain a human desktop/mobile acceptance decision before delivery.

## 10. Implementation Details

| Anchor | Why relevant | Intended change / precedent |
|---|---|---|
| `frontend/app/(public)/page.tsx` — `Home` | Defines current section order and route composition. | Update imports after relocation; preserve `Hero → HighlightedCampaigns → HowItWorks` order. Do not add a route or section outside D1. |
| `frontend/components/features/landing/hero.tsx` — `Hero` | Route-only static hero and primary in-page CTA. | Move to `frontend/app/(public)/_components/hero.tsx`. Preserve its `#kampanye` CTA and platform-level verified-organization statement; do not add unsupported trust imagery/claims. Use established token utilities only for objective rendered corrections. |
| `frontend/components/features/landing/highlighted-campaigns.tsx` — `HighlightedCampaigns` | Sole client data-fetching leaf and collection-state boundary. | Move to `frontend/app/(public)/_components/highlighted-campaigns.tsx`; retain `'use client'`, `useCampaigns`, existing loading/error/empty/stale branches, `id="kampanye"`, and `CampaignCard` mapping. Do not introduce a card link, feature filter, organization fields, media, or new query behavior. |
| `frontend/components/features/landing/how-it-works.tsx` — `STEPS`, `HowItWorks` | Static copy contains unsupported/conflicting promises. | Move to `frontend/app/(public)/_components/how-it-works.tsx`. Replace `STEPS` with exactly these authority-bounded messages: (1) **Temukan kampanye yang sedang berjalan** — **Jelajahi kampanye dari organisasi yang telah diverifikasi.** (2) **Pilih metode donasi** — **Transfer bank, kartu debit, e-wallet, atau QRIS tersedia untuk dipilih saat berdonasi.** (3) **Pantau status donasi** — **Setelah donasi dibuat, cek statusnya melalui tautan yang tersedia.** Keep it static; do not add a CTA or claims about fees, minimum amount, safety guarantees, or report publication. |
| `frontend/components/features/landing/{hero,highlighted-campaigns,how-it-works}.test.tsx` | Existing behavior/copy tests must follow ownership relocation. | Move with their components; revise assertions for the fixed copy and retain tests that prohibit fabricated organization counts and preserve list states. |
| `frontend/components/features/campaign/campaign-card.tsx` — `CampaignCard`, `CampaignCardSkeleton` | Campaign-semantic card and provisional no-media boundary. | Do not move or alter. Retain non-interactive, non-upload, no-organization/no-badge, no-media behavior. |
| `frontend/app/layout.tsx` — `RootLayout` | Owns document language. | Change the root `html` declaration from `lang="en"` to `lang="id"`; do not alter providers, fonts, metadata, or auth boundaries. |
| `frontend/app/(public)/_components/nav-items.ts` — `publicNavItems` | Contains the `/campaign` vs `#kampanye` authority contradiction. | Intentionally do not change pending Open Item 1; browser test must validate the current anchor only, not present it as final route resolution. |
| `frontend/app/(public)/_components/public-shell-client.tsx` — `PublicShellClient` | Existing mobile drawer/focus behavior to protect. | No production behavior change planned; cover the browser-level populated-home interaction at mobile width. |
| `frontend/tests/browser/login.smoke.spec.ts` and `frontend/playwright.config.ts` | Established browser-test viewport/overflow/dev-server precedent. | Add `frontend/tests/browser/home.smoke.spec.ts` following its 1280×800/390×844 loop. With current populated MSW fixture, assert core headings/CTA/collection visibility and no horizontal overflow at both widths; at mobile, open the hamburger, assert dialog navigation is visible, press Escape, assert it closes and focus returns to the trigger; activate the hero CTA and assert `#kampanye` is reached. Capture page errors as the login smoke does. |
| `frontend/components/features/campaign/campaign-card.test.tsx`, `frontend/components/ui/{button,badge,progress-bar}.tsx` | Current campaign and broad primitive contracts. | Do not change. Use as regression/ownership boundaries; any later request to alter these primitives requires new consumer-impact analysis outside this plan. |

## 11. Files Changed / Files NOT Changed

| File / area | Change type | Description |
|---|---|---|
| `frontend/app/(public)/page.tsx` | Modify | Update imports after route-local relocation; preserve composition order. |
| `frontend/app/(public)/_components/hero.tsx` | Move/modify | Route-local hero; retain truthful platform-level verification language and in-page CTA. |
| `frontend/app/(public)/_components/highlighted-campaigns.tsx` | Move | Route-local client list section; preserve established API/state behavior. |
| `frontend/app/(public)/_components/how-it-works.tsx` | Move/modify | Route-local explanatory section with fixed authority-bounded copy. |
| `frontend/app/(public)/_components/{hero,highlighted-campaigns,how-it-works}.test.tsx` | Move/modify | Preserve/make copy and state assertions at their new owner. |
| `frontend/app/layout.tsx` | Modify | Set document language to Indonesian. |
| `frontend/tests/browser/home.smoke.spec.ts` | Add | Representative desktop/mobile populated-route and mobile-drawer browser regression. |
| `frontend/components/features/landing/` | Remove after move | Former non-semantic owner; no standalone files remain after relocation. |

| File / area intentionally untouched | Why |
|---|---|
| `frontend/app/(public)/_components/nav-items.ts` | Open Item 1 prevents silently choosing between page-map `/campaign` and current `#kampanye`. |
| `frontend/components/features/campaign/campaign-card.tsx` and test | Campaign-owned contract and provisional truth-preserving placeholder are deliberately retained. |
| `frontend/components/ui/`, `frontend/app/globals.css` | No evidenced cross-product primitive/token deficiency; avoids unnecessary blast radius. |
| `frontend/lib/api/campaign.ts`, `frontend/lib/hooks/use-campaigns.ts`, `frontend/mocks/handlers.ts`, `api/` | No API/data/mock contract change is authorized or needed. |
| `backend/`, migrations, scripts, cron/runbooks | No backend or independently operable concern is in scope. |

## 12. Testing Checklist

- [ ] R1 — Add/run `home.smoke.spec.ts` at 1280×800 and 390×844 against the populated mock route; assert named core sections/CTA and no horizontal overflow. Perform Build-time rendered inspection at the same viewports.
- [ ] R2 — In the home browser smoke, activate `Mulai berdonasi` and assert the `#kampanye` collection is reached/visible without a navigation failure.
- [ ] R3 — In the mobile home browser smoke, assert open drawer dialog/navigation, Escape close, and focus return to the hamburger; retain/run unit drawer focus-trap coverage.
- [ ] R4 — Run relocated `HighlightedCampaigns` Vitest suite covering shaped loading cards, populated contract-limited cards, empty with no CTA, generic retryable error, and retry transition. Run the campaign-card suite to retain no-upload/no-organization behavior.
- [ ] R5 — Add/update `HowItWorks` test assertions for all three fixed messages and explicit absence of the removed Rp10,000/no-hidden-fees/periodic-report promises; retain hero no-fabricated-count test.
- [ ] R6 — Add a focused render/metadata assertion that the root document element has `lang="id"`, or cover it in the home browser smoke with `document.documentElement.lang`.
- [ ] R7 — Run the relocated landing component suites and `npm run build` to prove import/server-client boundary integrity; confirm `CampaignCard` and UI primitive tests remain unchanged/passing.
- [ ] R8 — Human manually exercises rendered `/` at representative desktop and mobile sizes after Build: hierarchy, CTA clarity, responsive usability, truthful provisional presentation, and no-media placeholder appropriateness. Record acceptance/rejection separately; do not mark it accepted from automation alone.
- [ ] Run `npm run verify` and `npm run test:browser`. If Chromium is unavailable, run `npm run browser:install` first; record all commands actually run and the existing Vite configuration warning if it persists.

### Test Focus Pointer

| Area | Why sensitive | Evidence anchor from Exploration | Still relevant post-synthesis? |
|---|---|---|---|
| Public trust/money static copy | Unsupported terms can mislead donors before any backend authority can correct the impression. | `1-exploration/logs/02-live-home-route-and-components.md#Sniffing` | Yes — R5 and R8 protect it; no payment/PII/concurrency test class is introduced because no payment flow or sensitive data boundary changes. |
| Mobile drawer/focus and responsive public shell | Focus loss, overflow, or hidden actions affects public accessibility and cannot be proved by component markup alone. | `1-exploration/logs/03-verification-and-rendering.md#Sniffing` | Yes — R1–R3 browser coverage and R8 human review. |
| Shared UI primitive blast radius | Visual primitive changes could affect account/dashboard consumers. | `1-exploration/logs/04-ownership-and-assets.md#Sniffing` | N/A — no primitive change under D6; if scope expands, reopen consumer discovery and independent test focus. |
| Campaign placeholder/media truth | Fabricated imagery would be a public trust risk. | `1-exploration/logs/04-ownership-and-assets.md#Sniffing` | N/A — D7 intentionally retains current provisional no-media treatment; R4/R8 verify it remains non-fabricating. |

## 13. Open Items

### Active — needs external input or verification

1. **Public campaign navigation authority conflict.** `docs/ui-ux/page-map.md`
   resolves `Jelajahi Kampanye` to `/campaign`, but that route is absent,
   Campaign delivery is `NOT_STARTED`, and live code uses `#kampanye`. A
   human/project owner must either formally allow the anchor as a temporary
   exception for this foundation slice or schedule/change the route authority
   with Campaign delivery. Build must leave `publicNavItems` untouched pending
   that decision.

2. **Human rendered acceptance.** A human must accept/reject the resulting
   desktop/mobile calibration after Build; automation is supporting evidence
   only.

### Resolved — retained as decision history

1. ~~**Minimum representative slice was unspecified.**~~ **RESOLVED — retain
   the public shell, text-first hero, fixed mock-backed campaign collection,
   and compact explainer (D1).** This gives the required calibration surface
   without expanding into a complete landing page.

2. ~~**How should deterministic exceptional campaign-list states be browser-tested?**~~
   **RESOLVED — retain Vitest/MSW as the exceptional-state seam and use
   Playwright for populated-route/responsive/drawer evidence (D8).** No
   test-only runtime mock interface will be introduced.
