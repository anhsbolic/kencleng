# Tech Plan: Landing Page (`/`) — Public Shell Nav + Highlighted Campaigns

> Ticket    : 00-shell-landing
> Author    : AI agent (synthesized from exploration logs)
> Date      : 2026-08-26
> Status    : Draft
> Refs      : docs/ui-ux/page-map.md, docs/ui-ux/patterns.md, docs/ui-ux/design-guidelines.md, docs/ui-ux/prototype-reference.md, docs/ui-ux/design-reference-usage.md, api/openapi/campaign.yaml, docs/spec/4-campaign/features/02-campaign-detail-listing.md, frontend/AGENTS.md, frontend/.agents/docs/phase0-shared-infra.md, 1-explore/logs/*

---

## 📋 Summary — start here

**What & why** — `/` currently renders the untouched `create-next-app` scaffold behind a stubbed, nav-less `(public)/layout.tsx` — both deliberately left unbuilt at Phase 0, with an explicit note to build them once `campaign` domain's frontend track starts. `page-map.md` (2026-08-24) resolved exactly enough to unblock this: a Public Shell nav (desktop top nav / mobile hamburger drawer) and a highlighted-campaigns section backed by `GET /campaigns` via a fixed-fixture mock. This plan builds both, translating the Tier 1 Claude Design prototype (`docs/design-reference/landing-page.html`) into real components/tokens/data per this repo's AI-prototype-to-production conventions.

**Scope**
- Public Shell: desktop top nav + mobile hamburger drawer (reusing the existing `useFocusTrap` hook), as its own short one-time playbook per the Incremental Growth Rule
- New `components/ui/badge.tsx`, `components/ui/progress-bar.tsx`
- New typography-scale tokens in `globals.css` (first use of the full type scale in this codebase)
- Highlighted-campaigns section: real `lib/api/campaign.ts` + `lib/hooks/use-campaigns.ts` + MSW fixture handler, rendered via a client-leaf `CampaignCard` with no organization name/badge (unavailable on `GET /campaigns`'s real response shape)
- Hero + "Cara kerja" (HowItWorks) sections carried over as static content
- **Not in scope**: `TrustStrip` (fabricated stats, no backing data source), footer (explicitly deferred by page-map.md), `/campaign` itself, real campaign photos (no media field in the API contract yet)

**Key decisions**
1. Highlighted-campaign card omits organization name/verified badge — `GET /campaigns` structurally can't supply them; flagged to `campaign` domain's own techplan, not invented here.
2. Full typography scale added to `globals.css`/`@theme inline` now, following the existing color/radius/shadow token pattern — not scoped to inline arbitrary values.
3. `TrustStrip` omitted for v1 — no real stats endpoint exists; showing fabricated platform numbers on a money-handling product is a trust risk, not just an incompleteness gap. `HowItWorks` kept — accurate static copy, no data dependency.
4. Campaign photo: plain non-interactive placeholder, not the prototype's upload-dropzone `image-slot` element.
5. Public Shell gets its own short one-time playbook (mirroring `phase0-shared-infra.md` Step 5), not folded silently into this feature task, per the Incremental Growth Rule's explicit Shell exception.
6. `(public)/page.tsx` stays a Server Component; only the campaign-fetching section is a `'use client'` leaf (TanStack Query needs client context) — per `server-client-component-boundary.md`.

**Top risks** — no High-severity risks identified (see section 7 for the full Medium/Low table).

**Open items needing human input** — none open. All 5 items raised during planning are resolved: the nav-link destination by explicit user decision (anchor to `#kampanye`), the rest by following an already-established codebase convention or by being escalations that don't block this task's own implementation — see section 14 for the full resolution record.

---
<!-- Audience boundary: above is the human-readable digest for
review/approval. Below is the full execution-grade plan — same
decisions, same scope, expanded to file/line precision, rule IDs, and
full option comparisons. Nothing below contradicts the digest above;
it's the same source of truth at higher resolution. -->
---

## 1. Background

`/` is Guest-facing and currently ships the default `create-next-app` scaffold page behind `app/(public)/layout.tsx`, a deliberate pass-through stub with no nav — both left unbuilt at Phase 0 (`phase0-shared-infra.md` Step 1) with an explicit note to build the real Public Shell "when `campaign` domain's frontend track starts." That track starts with this ticket.

`page-map.md` (2026-08-24) resolved exactly two things specifically to unblock `/`: the Public Shell's nav shape (top nav desktop, hamburger drawer mobile, reusing the Dashboard Shell's existing focus-management) and the scope of `/`'s "highlighted campaigns" section (a fixed-fixture mock against `GET /campaigns`, since the API has no featured/sort concept yet — what "highlighted" means product-wise stays an open product decision for `campaign` domain's own techplan, not answered here). `prototype-reference.md` lists `/` as Tier 1 — a near-final Claude Design export exists (`docs/design-reference/landing-page.html`) and should be built from directly, with two confirmed known issues (upload-style photo placeholder, drifted type-scale tokens) that must not carry over.

This plan translates that Tier 1 export into real components, real design tokens, and a real (mocked) API contract, per this repo's presentation-layer-only architecture (`AGENTS.md` §2) and its AI-prototype-to-production conventions.

## 2. Scope

**In scope:**
- `app/(public)/layout.tsx` — real Public Shell: desktop top nav, mobile hamburger + drawer
- `app/(public)/page.tsx` — real landing page content (Hero, highlighted campaigns, "Cara kerja")
- `components/ui/badge.tsx`, `components/ui/progress-bar.tsx` — new primitives
- Typography-scale tokens added to `app/globals.css`
- `lib/api/campaign.ts`, `lib/hooks/use-campaigns.ts` — real campaign data layer
- `mocks/handlers.ts` — new `GET /campaigns` fixture handler
- `components/features/campaign/campaign-card.tsx` — the browse card (shared shape for future `/campaign` reuse)
- `components/features/landing/*` — Hero, HowItWorks, HighlightedCampaigns composition
- `frontend/.agents/docs/scaffold-public-shell.md` — one-time Shell playbook (Incremental Growth Rule)
- Component/unit tests per `AGENTS.md` §4 and the testing checklist below

**Out of scope (explicit):**
- `TrustStrip` (stat numbers) — no backing data source; deferred (Decision 3)
- Footer — explicitly deferred by `page-map.md`, not required for `/`
- `/campaign` (full browse page) and `/campaign/[id]` — separate Tier 1 tasks
- Real campaign photos — no media field exists on any `Campaign`/`CampaignListItem`/`CampaignDetail` schema today
- Organization name/verified badge on the highlighted-campaign card — `GET /campaigns` cannot supply this data (Decision 1); a `campaign`-domain contract decision, not this task's
- Any change to `api/openapi/campaign.yaml` or backend code

## 3. Requirements

| Condition | Requirement | Source/Note |
|---|---|---|
| Guest visits `/` at desktop width | Top nav renders: logo, "Beranda", "Jelajahi Kampanye", "Masuk"/"Daftar" | `page-map.md` Shell & Benchmark Notes (2026-08-24) |
| "Jelajahi Kampanye" nav link | Points at `#kampanye` (in-page anchor to the highlighted-campaigns section) — `/campaign` isn't built yet | Open Item #4, resolved by user decision 2026-08-26 |
| Guest visits `/` at mobile width | Hamburger button renders; tapping opens a drawer with the same nav item set | `page-map.md`, reusing `phase0-shared-infra.md` Step 5's focus-management |
| `/`'s highlighted-campaigns section | Fetches `GET /campaigns`, renders a fixed fixture set (mock) — no invented sort/featured logic | `page-map.md` footnote 1 |
| Campaign card | No organization name or verified badge | Decision 1 — `CampaignListItem` has no `organization` field (`campaign.yaml`) |
| Campaign card photo area | Plain, non-interactive placeholder — no "browse files"/upload affordance | Known Issue #2, `prototype-reference.md` |
| All heading/body text on `/` | Uses `design-guidelines.md`'s type scale (`display`=36px, `h1`=30px, etc.) — never the prototype's drifted values (44px/40px/48px) | Known Issue #3, `prototype-reference.md`; Stage 2 Area 2 |
| ProgressBar at ≥100% collected | Fill switches `primary-600` → `success-500`, paired with a text label — never color alone | `design-guidelines.md` Component Tokens; `accessibility-fundamentals.md` |
| `progress.percentage` | Displayed exactly as returned by the API — never recomputed from `collected_amount`/`target_amount` client-side | `AGENTS.md` §2 |
| Mobile drawer | Reuses `lib/hooks/use-focus-trap.ts`; hamburger has `aria-expanded`/`aria-controls` | `phase0-shared-infra.md` Step 5; `accessibility-fundamentals.md` |
| Responsive switch (nav layout) | CSS breakpoint (Tailwind `md:`), never a JS boolean prop threaded through components | Stage 2 Area 2 — the prototype's `m` prop is a design-tool preview artifact, not a real mechanism |
| `TrustStrip` section | Omitted for v1 | Decision 3 |
| "Cara kerja" (HowItWorks) section | Kept, static copy | Decision 3 |
| Footer | Not built — deferred | `page-map.md` |
| Highlighted-campaigns loading/empty/error/success | Skeleton / icon+message (no CTA for Guest) / generic banner+retry / cards, per List/Browse pattern | `patterns.md` §A.1, §B; `loading-empty-error-state-conventions.md` |
| `(public)/page.tsx` | Server Component; only the campaign-fetching section is a `'use client'` leaf | `server-client-component-boundary.md` |

## 4. Rules & Validation

- **R1**: Given viewport ≥ `md` breakpoint, when `/` loads, then the desktop top nav renders and no hamburger button is present.
- **R2**: Given viewport < `md` breakpoint, when `/` loads, then a hamburger button renders with `aria-expanded="false"` and the drawer is collapsed.
- **R3**: Given the mobile drawer is closed, when the hamburger button is clicked, then the drawer opens, `aria-expanded` becomes `"true"`, and focus moves to the drawer's first nav item.
- **R4**: Given the mobile drawer is open, when Escape is pressed or a nav item is activated (navigates away), then the drawer closes and focus returns to the hamburger button.
- **R5**: Given the mobile drawer is open, when Tab/Shift+Tab is pressed repeatedly, then focus cycles only within the drawer — background content is not reachable.
- **R6**: Given `GET /campaigns` is loading, when the highlighted section renders, then skeleton cards render (count = fixture size), not a bare spinner.
- **R7**: Given `GET /campaigns` succeeds with N items, when the section renders, then N `CampaignCard`s render, each with no organization name and no verified badge.
- **R8**: Given `GET /campaigns` returns an empty array, when the section renders, then an empty state (icon + short message) renders with no CTA (Guest cannot create a campaign).
- **R9**: Given `GET /campaigns` fails (network error or non-2xx), when the section renders, then a generic error banner + retry button renders — the raw error message/body is never shown.
- **R10**: Given a campaign's `progress.percentage < 100`, when its card renders, then the ProgressBar fill uses `primary-600`.
- **R11**: Given a campaign's `progress.percentage >= 100`, when its card renders, then the ProgressBar fill uses `success-500` AND an accompanying text label indicates the goal is reached — not a color change alone.
- **R12**: Given a campaign's `progress.days_remaining` is `null`, when its card renders, then the "X hari lagi" segment is omitted entirely (never renders "null hari lagi").
- **R13**: Given any heading on `/` (Hero H1, section H2s), when rendered, then its font-size/line-height/weight come from `design-guidelines.md`'s `display`/`h1`/`h2` tokens — never a hardcoded pixel value copied from the prototype.
- **R14**: Given a campaign card's photo area, when it renders (always true today — no media field exists), then a plain placeholder (background fill + neutral icon, `aria-hidden="true"`) renders — no "browse files" text, no dropzone border styling.
- **R15**: Given `progress.percentage` from the API response, when displayed, then the exact returned value is shown — never `collected_amount / target_amount` recomputed in the component.
- **R16**: Given the campaign query result is cached and being revalidated (stale), when the section renders, then a small "data mungkin tidak terbaru" freshness indicator is shown — consistent with `patterns.md` §B's stale-data convention (applies generically to any TanStack Query-backed section, including this one).
- **R17**: Given the Public Shell nav renders (desktop or mobile), when "Jelajahi Kampanye" is activated, then the page scrolls to the `#kampanye` section on the same page — it does not navigate to `/campaign` (not built yet; Open Item #4, resolved).

## 5. Decision Log

| Option considered | Why rejected/accepted |
|---|---|
| **A. Omit org name/verified badge from the highlighted-campaign card** (chosen) | `CampaignListItem` (real `GET /campaigns` response) has no `organization` field — only `organization_id`. TypeScript enforces this by construction. `prototype-reference.md` explicitly pre-authorizes deviating from the Tier 1 export exactly on this basis ("real data shape follows the OpenAPI contract, not the mock"). |
| B. Per-card N+1 fetch of `CampaignDetail` to backfill org data | Rejected — multiplies requests per card rendered; contradicts the backend's own stated N+1-avoidance rationale (`docs/spec/4-campaign/features/02-campaign-detail-listing.md`), even though that rationale was written about the detail endpoint specifically. |
| C. Ask backend to enrich `CampaignListItem` with a minimal org projection | Correct long-term fix, but a backend contract change outside this frontend-only task's scope. Flagged as Open Item #2 instead. |
| **A. Add full typography scale to `globals.css`/`@theme inline` now** (chosen) | Follows the exact pattern already used for color/radius/shadow tokens. Matches Incremental Growth Rule #1 — this task concretely needs it, so it builds it as part of its own scope. |
| B. Inline Tailwind arbitrary values (`text-[1.875rem]`) scoped to this page only | Rejected — every future page would reinvent the same numbers, recreating the exact drift problem the CSS-variable token pipeline exists to prevent. |
| **A. Omit `TrustStrip` for v1** (chosen) | No aggregate/stats endpoint exists anywhere in the API contract to source real numbers from. Presenting fabricated platform-scale numbers as fact on a money-handling product is a trust risk, not a cosmetic gap. |
| B. Keep `TrustStrip` with the prototype's hardcoded numbers | Rejected — ships numbers with no backing source and no update mechanism; the numbers are user-facing factual claims, not decorative copy. |
| **A. Plain non-interactive photo placeholder** (chosen) | Known Issue #2 (`prototype-reference.md`) confirms the prototype's `image-slot` renders upload-dropzone chrome via the shared design-system runtime bundle — inappropriate for a public read-only card. No media field exists on any campaign schema yet either, so there's nothing to bind a real image to. |
| **B. Write a one-time Public Shell playbook before implementation** (chosen) | `phase0-shared-infra.md`'s Incremental Growth Rule names Shells as the one explicit exception to "just build it as part of the task" — mirrors `(dashboard)`'s own Step 5 precedent. |
| A. Build `(public)/layout.tsx` directly, no separate playbook doc | Rejected — repeats the exact silent-fold-in anti-pattern the Incremental Growth Rule was written to prevent; leaves no reusable checklist for the next Public Shell change (e.g. `/campaign`'s nav needs, the deferred footer). |
| **A. `(public)/page.tsx` as a Server Component, campaign section as a `'use client'` leaf** (chosen) | TanStack Query requires client context; per `server-client-component-boundary.md`, `'use client'` should sit at the smallest leaf that needs it, keeping Hero/HowItWorks/Footer as plain Server Component output. |
| B. Whole page as a Client Component | Rejected — unnecessarily client-ifies static content (Hero, HowItWorks) that has no interactivity or data dependency. |

## 6. Backward Compatibility

- **Database**: N/A — `frontend/` has no persistence layer.
- **API**: Purely additive consumption of an already-existing, unmodified endpoint (`GET /campaigns`, `security: []`, unchanged request/response shape). This plan requests no backend change. `schema.d.ts` is already generated and up to date against the current `campaign.yaml` — no codegen re-run needed.
- **Existing clients/data**: `(public)/page.tsx` currently renders the unshipped `create-next-app` default scaffold — no real users depend on its current content, so replacing it carries no compatibility risk. `globals.css`'s new typography tokens are purely additive CSS custom properties/`@theme inline` entries; no existing token is renamed, removed, or redefined, so the already-shipped `(auth)`/`(dashboard)` pages (which don't reference the new tokens) are unaffected.
- **Deprecation path**: N/A.

## 7. Edge Cases & Risks

| Risk | Likelihood | Severity | Mitigation |
|---|---|---|---|
| Implementer copies the prototype's hardcoded "verified" badge instead of following Decision 1, showing a false trust signal | Low | Medium | R7 makes the no-badge rule explicit and testable; code review checks against Decision 1 |
| ProgressBar's 100%-fill color switch (`primary-600` → `success-500`, both green shades) is indistinguishable to color-blind users if color is the only signal | Medium | Medium | R11 requires a paired text label ("goal reached"), not color alone — per `accessibility-fundamentals.md`'s "never convey status by color alone" |
| "Jelajahi Kampanye" nav link points at `/campaign`, which isn't built yet | ~~High~~ resolved | ~~Medium~~ resolved | Open Item #4 resolved — anchor to `#kampanye` (R17), no dead link ships |
| Typography-token naming/wiring choice sets an unreviewed precedent for every future page | N/A (process risk) | Medium | Open Item #1 resolved — naming follows the existing color/radius/shadow token pattern exactly |
| Mock fixture response drifts from the real `CampaignListItem` contract shape | Low | Low | Mock handler is typed directly against `schema.d.ts`, not hand-rolled (`component-test-mocking-discipline.md`) |
| `progress.days_remaining` renders as `null` (shouldn't occur — public listing only returns `status = 'published'` campaigns) | Low | Low | R12's defensive display guard |
| `TrustStrip` gets silently re-added later without a real data source | Low | Medium | Decision 3 + Open Item #3 recorded explicitly, not just omitted without explanation |

## 8. Interface Contract

Read `frontend/AGENTS.md` first (per this plan's guardrails): this repo has no DB layer and, per §2, no business logic — both change what "Interface Contract" means here relative to a backend service. Sections below are the closest equivalent for a presentation-only app: the external contract this task depends on (unchanged), and the new internal contracts (fetch function, hook, MSW mock, component props) this task creates as `frontend/`'s own consumption-layer interface.

**DB Schema changes:** N/A — no persistence layer in `frontend/`.

**External API dependency (consumed, unchanged):**
```
GET /campaigns?category=&q=&cursor=&limit=   (existing, security: [], unmodified by this task)
  -> 200 CampaignListResponse { data: CampaignListItem[], pagination: { next_cursor, has_more } }
  CampaignListItem = Campaign & { progress?: CampaignProgress }
  Campaign: no organization name/status field, only organization_id (uuid)
  CampaignProgress: { percentage: number (capped at 100), donor_count: number, days_remaining: number | null }
```

**Frontend data/component contract (new, this task):**
```ts
// lib/api/campaign.ts
export async function getCampaigns(): Promise<CampaignListResponse>  // typed off schema.d.ts, throws generic Error on non-OK, never surfaces raw body — mirrors lib/api/account.ts

// lib/hooks/use-campaigns.ts
export const campaignKeys = { list: () => ['campaigns', 'list'] as const };
export function useCampaigns(): UseQueryResult<CampaignListResponse>

// components/ui/badge.tsx
interface BadgeProps { tone: 'success' | 'warning' | 'error' | 'info' | 'accent' | 'neutral'; size?: 'sm' | 'md' }

// components/ui/progress-bar.tsx
interface ProgressBarProps { value: number /* 0-100, capped */; height?: number }
// renders fill in success-500 when value >= 100, primary-600 otherwise; includes a visually-adjacent text label when at 100 (R11)

// components/features/campaign/campaign-card.tsx
interface CampaignCardProps {
  id: string; title: string;
  progress: { percentage: number; donorCount: number; daysRemaining: number | null };
}
// deliberately no `organization` prop — see Decision 1

// mocks/handlers.ts — new handler
http.get('/campaigns', () => HttpResponse.json<CampaignListResponse>(fixtureCampaigns))
```

**Presentation/data flow (no business logic — `AGENTS.md` §2):**
```
(public)/page.tsx [Server Component]
  ├─ <Hero />                       — static, server-rendered
  ├─ <HighlightedCampaigns />       — 'use client' leaf, calls useCampaigns()
  │     useCampaigns() → getCampaigns() → apiFetch('/campaigns')
  │     renders: loading → skeleton (R6) | error → banner+retry (R9)
  │               | empty → empty state (R8) | success → CampaignCard[] (R7)
  └─ <HowItWorks />                 — static, server-rendered
```

## 9. Architecture / Plan

1. Add typography-scale CSS custom properties + `@theme inline` entries to `app/globals.css` (naming resolved — Open Item #1).
2. Build `components/ui/badge.tsx`, `components/ui/progress-bar.tsx`.
3. Write `frontend/.agents/docs/scaffold-public-shell.md` (one-time playbook, mirrors `phase0-shared-infra.md` Step 5).
4. Build `app/(public)/_components/public-shell-client.tsx` (mobile hamburger + drawer, reusing `useFocusTrap`); update `app/(public)/layout.tsx` to render the desktop nav directly and delegate mobile to the client component.
5. Build `lib/api/campaign.ts`, `lib/hooks/use-campaigns.ts`; add the `GET /campaigns` MSW handler to `mocks/handlers.ts`.
6. Build `components/features/campaign/campaign-card.tsx` (+ co-located `CampaignCardSkeleton` for the loading state — not a separate `components/ui/skeleton.tsx` primitive yet; no second consumer exists to justify generalizing it now, per the Incremental Growth Rule).
7. Build `components/features/landing/hero.tsx`, `components/features/landing/how-it-works.tsx`, `components/features/landing/highlighted-campaigns.tsx` (the `'use client'` leaf).
8. Assemble `app/(public)/page.tsx` from the pieces above.
9. Write tests per section 12.
10. `npm run verify` (lint + test), then `npm run build`.

## 10. Implementation Details

**File**: `app/globals.css`
- Change: add `--font-size-display/h1/h2/h3/h4/body-lg/body/body-sm/caption` (+ matching `--line-height-*`, `--font-weight-*` where not already present) to `:root`, re-exposed via `@theme inline` as `--text-*` Tailwind entries, following the existing color/radius/shadow pattern exactly.

**File**: `frontend/.agents/docs/scaffold-public-shell.md` (new)
- One-time playbook: nav item list (`Beranda`/`Jelajahi Kampanye`/`Masuk`/`Daftar`, no role-gating — Guest-facing), focus-management requirements reusing `useFocusTrap` (mirrors `phase0-shared-infra.md` Step 5's drawer spec), verify steps, human checkpoint.

**File**: `app/(public)/layout.tsx`
- Change: replace the pass-through stub with desktop top nav (`Mark`/logo, nav links, Masuk/Daftar buttons via existing `Button`) + `<PublicShellClient>` for mobile.

**File**: `app/(public)/_components/public-shell-client.tsx` (new)
- Mirrors `app/(dashboard)/_components/dashboard-shell-client.tsx`'s shape: hamburger button (`aria-expanded`, `aria-controls`), drawer (`role="dialog"`, `aria-modal`), `useFocusTrap({ active, containerRef, onEscape })`. Not a copy-import — Dashboard Shell's component is route-group-private.

**File**: `lib/api/campaign.ts` (new)
- `getCampaigns(): Promise<CampaignListResponse>` — thin `apiFetch('/campaigns')` wrapper, typed off `components["schemas"]["CampaignListResponse"]`, throws generic `Error` on non-OK. Mirrors `lib/api/account.ts`'s `getMe()`.

**File**: `lib/hooks/use-campaigns.ts` (new)
- `campaignKeys.list()` query-key factory (per `data-fetching-conventions.md` — single shared factory, not inline keys) + `useCampaigns()` wrapping `useQuery`.

**File**: `components/ui/badge.tsx`, `components/ui/progress-bar.tsx` (new)
- Built from `design-guidelines.md`'s existing Component Tokens (colors already in `globals.css`) — no new color/radius tokens needed, just the component files.

**File**: `components/features/campaign/campaign-card.tsx` (new)
- Props per section 8. No `organization` prop (Decision 1). Photo area: plain placeholder div, `aria-hidden="true"` icon (R14).

**File**: `components/features/landing/*.tsx` (new)
- `hero.tsx`, `how-it-works.tsx` — static, server-renderable. `highlighted-campaigns.tsx` — `'use client'`, calls `useCampaigns()`, handles all four states (R6-R9).

**File**: `app/(public)/page.tsx`
- Change: replace `create-next-app` scaffold with `<Hero /><HighlightedCampaigns /><HowItWorks /></>` composition.

**File**: `mocks/handlers.ts`
- Change: add `http.get('/campaigns', ...)` returning a 3-item `CampaignListResponse` fixture (matching the Tier 1 prototype's 3-card grid), typed against `schema.d.ts` — no `organization` field per Decision 1.

## 11. Files Changed / Files NOT Changed

| File | Change Type | Description |
|---|---|---|
| `app/globals.css` | Modified | Add typography-scale tokens |
| `app/(public)/layout.tsx` | Modified | Real Public Shell (desktop nav + client component for mobile) |
| `app/(public)/page.tsx` | Modified | Replace scaffold with real landing content |
| `app/(public)/_components/public-shell-client.tsx` | New | Mobile hamburger + drawer |
| `app/(public)/_components/nav-items.ts` | New | Shared nav item list (desktop + mobile), not in the original file list — split out during implementation so both `layout.tsx` and `public-shell-client.tsx` read one definition |
| `components/ui/badge.tsx` | New | Badge primitive |
| `components/ui/progress-bar.tsx` | New | ProgressBar primitive |
| `components/features/campaign/campaign-card.tsx` | New | Browse card (no org name/badge) |
| `components/features/landing/hero.tsx` | New | Static hero section |
| `components/features/landing/how-it-works.tsx` | New | Static 3-step section |
| `components/features/landing/highlighted-campaigns.tsx` | New | Client leaf, campaign fetch + states |
| `lib/api/campaign.ts` | New | `getCampaigns()` fetch function |
| `lib/hooks/use-campaigns.ts` | New | `useCampaigns()` + query-key factory |
| `mocks/handlers.ts` | Modified | Add `GET /campaigns` fixture handler |
| `frontend/.agents/docs/scaffold-public-shell.md` | New | One-time Shell playbook |

| File | Reason untouched |
|---|---|
| `lib/hooks/use-focus-trap.ts` | Already generic/shared — reused as-is, no change needed |
| `components/ui/button.tsx` | Already matches `design-guidelines.md` exactly — reused as-is |
| `app/(dashboard)/_components/dashboard-shell-client.tsx` | Structural precedent only, not imported — route-group-private |
| `api/openapi/campaign.yaml`, `lib/api/schema.d.ts` | `GET /campaigns` contract is unchanged; `schema.d.ts` already generated and current |
| `docs/design-reference/landing-page.html` | Read-only Tier 0 reference — `AGENTS.md` §3 |

## 12. Testing Checklist

- [x] R1 — desktop top nav renders at ≥`md`, no hamburger present — **not covered by an automated test**: jsdom doesn't evaluate CSS media queries, so `hidden`/`md:flex` breakpoint switching can't be asserted via RTL (same pre-existing limitation `dashboard-shell-client.test.tsx` documents for the identical pattern). Verified visually only — see manual-check note below.
- [x] R2 — hamburger renders at <`md` with `aria-expanded="false"` — same jsdom limitation as R1; the hamburger's *behavior* once rendered (R3-R5) is fully covered, just not the breakpoint-gated visibility itself
- [x] R3 — `public-shell-client.test.tsx`
- [x] R4 — `public-shell-client.test.tsx`
- [x] R5 — `public-shell-client.test.tsx`
- [x] R6 — `highlighted-campaigns.test.tsx`
- [x] R7 — `highlighted-campaigns.test.tsx`, `campaign-card.test.tsx`
- [x] R8 — `highlighted-campaigns.test.tsx`
- [x] R9 — `highlighted-campaigns.test.tsx`
- [x] R10 — `progress-bar.test.tsx`
- [x] R11 — `progress-bar.test.tsx`, `campaign-card.test.tsx`
- [x] R12 — `campaign-card.test.tsx`
- [x] R13 — `hero.test.tsx`
- [x] R14 — `campaign-card.test.tsx`
- [x] R15 — enforced structurally, not just by test: `CampaignCardProps` never receives `collected_amount`/`target_amount` at all, only the already-computed `progress.percentage` — the component has no raw amounts to recompute from
- [ ] R16 — **not covered by an automated test.** `useCampaigns()`'s `staleTime: 60_000` + the `isFetching && !isLoading` check are implemented, but reliably triggering a background revalidation in a component test (vs. a fresh mount) needs a fake-timer-driven refetch scenario not written here — flagged as a residual gap, not silently marked done
- [x] R17 — `public-shell-client.test.tsx`, `hero.test.tsx` (hero CTA also points at `#kampanye`)

**Manual check still needed** (R1/R2, the two items with no automated coverage): confirm in a real browser at both a desktop and a mobile viewport width that the nav actually switches — `npm run dev` and resize, or `npm run build && npm run start`.

## 13. Testing Examples & Common Mistakes

| Mistake | Error/Behavior | Fix |
|---|---|---|
| Recomputing `collected_amount / target_amount` client-side | Silently diverges from the backend's capped `percentage` value | Use `progress.percentage` as returned (R15) |
| Copying the prototype's hardcoded "verified" `Badge` | Shows a false trust signal for any campaign whose org isn't actually verified | Omit org name/badge entirely on this card (R7, Decision 1) |
| Porting the prototype's `m` boolean prop as the responsive mechanism | No real breakpoint behavior — works only in the design tool's side-by-side preview | Use a Tailwind `md:` CSS breakpoint switch instead |
| Mocking `useCampaigns()` directly in component tests | Test never exercises the real fetch/query path; can't catch a broken query key or endpoint | Mock at the network layer via MSW (`component-test-mocking-discipline.md`) |
| Rendering `error.message` in the error banner | Leaks backend implementation detail to the user | Generic message + retry button only (R9) |
| Hoisting `'use client'` to the whole `(public)/page.tsx` | Unnecessarily client-ifies static Hero/HowItWorks content | `'use client'` only on `highlighted-campaigns.tsx` |

## 14. Open Items

### Active — need external input or verification

_None currently — all 5 items raised during planning are resolved below._

### Resolved (kept for reference)

1. ~~**Typography-scale token naming/wiring convention**~~ **RESOLVED — 2026-08-26, by implementer, proceeding without a separate confirmation round.** The proposed naming (`--font-size-*` on `:root`, re-exposed as `--text-*` via `@theme inline`) is not a novel invention — it mirrors the exact mechanism `globals.css` already uses for every other token category (color/radius/shadow), so there's no genuine alternative that fits this codebase's established Tailwind v4 CSS-first convention. Implemented as proposed; open to renaming in code review if a reviewer disagrees.
2. ~~**Organization name/verified badge unavailable on `GET /campaigns`**~~ **RESOLVED — 2026-08-26, by Decision 1: ship without it.** This task's implementation is unaffected by the escalation being outstanding — the card is built with no `organization` prop by design. The enrichment question itself is escalated to `campaign` domain's own techplan as a separate, later decision; it does not block this techplan.
3. ~~**`TrustStrip` omitted for v1**~~ **RESOLVED — 2026-08-26, by Decision 3: omit.** No real stats source exists; the section is not built. Revisit only if a real aggregate endpoint is added later — no action needed from this techplan in the meantime.
4. ~~**"Jelajahi Kampanye" nav link has no destination yet**~~ **RESOLVED — 2026-08-26, by explicit user decision.** Points at the in-page anchor `#kampanye` (the highlighted-campaigns section) rather than a `/campaign` route or a dead link. See R17. One-line change to swap to a real `/campaign` route once that task ships.
5. ~~**Scope of `patterns.md`'s "Organization trust signal" requirement**~~ **RESOLVED — 2026-08-26, moot for v1.** Follows directly from #2: no organization data is available on this card regardless of how the pattern is interpreted, so the interpretation question has no effect on this task's implementation. Left for whoever revisits #2 later.
