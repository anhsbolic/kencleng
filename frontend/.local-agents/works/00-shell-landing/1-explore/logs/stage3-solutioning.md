# Stage 3 — Solutioning

> Task: 00-shell-landing
> Date: 2026-08-26

## Decisions

### 1. Organization name/verified badge on the `/` highlighted-campaign card

**Problem:** Tier 1 prototype's `CampaignBrowseCard` shows org name + a
"verified" badge on every card; `GET /campaigns` (`CampaignListItem`) has
neither — only `organization_id` (UUID). Spec-authored: the composite
(campaign+org+progress) shape was built for the detail endpoint
specifically to avoid N+1 "on the platform's primary conversion surface,"
never extended to the list endpoint.

**Decision: Option A** — omit org name/badge on this card, ship against
the real contract as-is. TypeScript enforces this by construction
(`CampaignListItem[]` has no `organization` field). Visually deviates
from the Tier 1 prototype, but `prototype-reference.md` explicitly
pre-authorizes exactly this: "deviating only where the feature spec
requires different behavior than what was mocked... real data shape
follows the OpenAPI contract, not the mock." Flag as an open item for
`campaign` domain's own techplan (same treatment as the "what does
'highlighted' mean" question already deferred there) — don't invent org
data or reach for a per-card N+1 fetch (rejected: multiplies requests per
card, contradicts the backend's own N+1-avoidance rationale by analogy).
Backend contract enrichment (Option C) is the correct long-term fix but
out of scope for a frontend-only task.

**Component shape:** since `/campaign` (Tier 1, later task) reads the
same `CampaignListItem` shape, place the card at
`components/features/campaign/campaign-card.tsx`, not inlined in the
`(public)` page — avoids a near-certain duplicate when `/campaign`'s task
starts; not scope creep since this task must build the card regardless.

### 2. Typography-scale tokens don't exist yet

**Decision: Option A** — add the full type scale (`display`/`h1`–`h4`/
`body-lg`/`body`/`body-sm`/`caption`) to `globals.css` now, following the
exact pattern already used for color/radius/shadow (CSS custom properties
on `:root`, re-exposed via `@theme inline`). Rejected: scoping to inline
arbitrary Tailwind values just for this page — would leave every
subsequent page reinventing the same numbers, recreating the drift
problem the CSS-variable approach exists to prevent. Matches Incremental
Growth Rule #1 ("a domain task builds shared infra only when that task
concretely needs it... adds it as part of its own scope") — `globals.css`
isn't a Shell, doesn't need its own playbook, just needs the tokens added
as part of this task.

### 3. `HowItWorks` and `TrustStrip` sections (in prototype, not in spec docs)

**`HowItWorks`** ("Cara kerja," 3-step explainer): static copy, no data
dependency, factually describes the actual documented product flow.
**Decision: keep** — legitimate microcopy carried over per
`ai-prototype-to-production-translation.md`.

**`TrustStrip`** (hardcoded stats: "120+ organisasi," "Rp 2,4M+ dana
tersalurkan," "8.500+ donatur"): user-facing factual claims with no
aggregate/stats endpoint anywhere in the API contract to source real
numbers from.

**Decision: Option A — omit for v1.** Showing fabricated numbers as fact
on a platform whose brand direction explicitly flags "handles money" as
trust-sensitive is a real risk, not just a completeness gap. Not a
deviation from a resolved decision (page-map.md never mentions this
section) — declining to add a new one the spec doesn't ask for and the
API can't back honestly. Flag as an open item for whenever a real stats
endpoint exists.

### 4. Campaign photo placeholder

Already directionally settled by task instructions (Known Issue #1 must
not carry over) plus Stage 2's root-cause: the shared `image-slot`
element renders upload-dropzone chrome, and no `Campaign` schema (list or
detail) has any media field today — not a "swap the placeholder graphic"
fix, there's no real photo URL to eventually bind to either.

**Decision:** plain, non-interactive placeholder (background fill +
neutral icon, no "browse files" text, no dropzone border), sized to match
the reference's card image area. Purely decorative until `campaign`
domain's contract adds a media field — swap to `next/image` later as a
localized change.

### 5. Public Shell: build directly, or write a one-time playbook first?

`phase0-shared-infra.md`'s Incremental Growth Rule names Shells as the
one explicit exception to "just build it as part of the task": gets its
own short one-time playbook when its triggering domain starts (mirroring
Step 5's shape for `(dashboard)`), not folded silently into the first
feature task. `(public)` is named by that same rule as due now.

**Decision: Option B** — write a short one-time playbook (e.g.
`frontend/.agents/docs/scaffold-public-shell.md`, mirroring
`phase0-shared-infra.md` Step 5: nav item list, focus-management
requirements reusing `useFocusTrap`, verify steps, human checkpoint)
before/alongside implementation. Rejected: building directly with no
playbook — repeats the exact anti-pattern (silent fold-in) the rule was
written to call out, and leaves no reusable checklist for whoever next
touches the Public Shell (e.g. `/campaign`'s nav needs, or the deferred
footer).

## Settled findings (no trade-off needed)

- **Responsive switching:** the prototype's `m` boolean is a design-tool
  preview artifact (two side-by-side canvases), not a real mechanism —
  implement with an actual CSS breakpoint switch (Tailwind `md:`
  variants), matching the Auth Shell's modal/full-page split.
- **Progress percentage:** render `CampaignProgress.percentage` exactly
  as returned (already capped server-side); never recompute
  `collected_amount / target_amount` client-side (AGENTS.md §2).
- **`days_remaining` nullability:** omit the "X hari lagi" segment when
  null, even though the public listing's `status = 'published'`-only
  filter means it shouldn't occur in practice — cheap guard against a
  malformed fixture.
- **List/Browse states:** implement loading (skeleton cards)/empty (icon
  + message, no CTA for Guest)/error (generic banner + retry)/success per
  `patterns.md` §A.1/§B, even though the mock is a fixed 3-item array —
  these states aren't conditional on "is the backend real yet."
- **Data layer:** `lib/api/campaign.ts` (thin `apiFetch` wrapper, typed
  off already-generated `schema.d.ts`, generic error on non-OK) +
  `lib/hooks/use-campaigns.ts` (TanStack Query) + a new `GET /campaigns`
  handler in `mocks/handlers.ts` returning 3 fixture items — mirrors the
  existing `account.ts`/`use-account-me.ts` convention exactly.

## Implementation scope

### New files

- `frontend/.agents/docs/scaffold-public-shell.md` — one-time playbook
  (Decision 5)
- `app/(public)/_components/public-shell-client.tsx` — hamburger/drawer
  client component, mirrors `DashboardShellClient` shape, uses
  `useFocusTrap`
- `app/(public)/page.tsx` — real landing page (replaces `create-next-app`
  scaffold)
- `components/ui/badge.tsx`, `components/ui/progress-bar.tsx`
- `components/features/campaign/campaign-card.tsx` — the browse card
  (no org name/badge per Decision 1)
- `components/features/landing/hero.tsx`,
  `components/features/landing/how-it-works.tsx`,
  `components/features/landing/highlighted-campaigns.tsx` (or similar
  decomposition mirroring the reference's `Hero`/`HowItWorks`/`Featured`)
- `lib/api/campaign.ts` — `getCampaigns()` fetch function
- `lib/hooks/use-campaigns.ts` — TanStack Query hook

### Modified files

- `app/(public)/layout.tsx` — real Public Shell (desktop nav +
  `PublicShellClient` for mobile)
- `app/globals.css` — add typography-scale tokens (Decision 2)
- `mocks/handlers.ts` — add `GET /campaigns` handler (3 fixtures, no
  `organization` field)
- `AGENTS.md` / `docs/ui-ux/prototype-reference.md` — stale
  `design-reference/` path fix (housekeeping, not blocking)

### Not built (deferred)

- `TrustStrip` (Decision 3) — flagged as open item pending a real stats
  endpoint
- Footer — explicitly deferred by page-map.md, not required for `/`
- Real campaign photo — no media field in the API contract yet
- Org name/verified badge on the home card — flagged as open item for
  `campaign` domain's techplan (Decision 1)

## Housekeeping (flagged, not fixed in this task)

- `AGENTS.md` §3 / `prototype-reference.md` cite `design-reference/` at
  repo root; actual location is `docs/design-reference/`.
- `design-reference-usage.md`'s extraction script misses
  `landing-page.html`'s actual component source (loaded via a gzip-
  compressed manifest blob keyed by an external script `src`, not the
  inline babel script the doc's regex targets) — silently produces a
  near-empty JSX file rather than erroring.
