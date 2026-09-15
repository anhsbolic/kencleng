# Stage 2 — Live `/` route, components, and tokens

## Current state

`app/(public)/page.tsx` composes three route-owned feature sections in order:
static `Hero`, client-fetched `HighlightedCampaigns`, and static `HowItWorks`.
`app/(public)/layout.tsx` supplies the public shell and a single mounted auth
modal. Desktop navigation is visible at `md` and above; `PublicShellClient`
supplies a mobile menu/drawer, focus trap, and modal-opening auth controls
below that breakpoint.

The live nav data does **not** point to the page-map’s `/campaign` route:
`publicNavItems` maps `Jelajahi Kampanye` to the local `#kampanye` anchor and
documents that `/campaign` does not yet exist.

The home route already has responsive structural behavior: hero actions stack
on small screens; campaign and explainer grids change from one to three
columns at `md`; the mobile nav replaces desktop nav. The route consumes the
CSS-first token system in `app/globals.css`, including the specified fonts,
type scale, color/surface, radius, and elevation tokens. No live visual
inspection has yet been performed in this exploration phase.

`HighlightedCampaigns` uses the public `GET /campaigns` query and has shaped
loading skeletons, generic retryable error, empty-without-CTA, populated, and
background-revalidation states. It maps only fields supplied by
`CampaignListItem` into the campaign card. Cards are non-interactive,
intentionally have an icon placeholder rather than campaign media, and omit
organization identity/status because the list contract cannot supply it.

The hero and explainer are static copy. The hero says its campaigns are from
verified organizations and marks the badge `Organisasi terverifikasi`; the
explainer promises campaign verification, safe donation with named methods,
a Rp10,000 minimum and no hidden fees, and periodic fund-use reports. The
donation feature source currently specifies an amount minimum of Rp5,000,
and its payment methods are simulated/recorded only; no source found in this
area establishes the Rp10,000 amount or the “no hidden fees” claim.

## Requirement

- Feature spec, **Requirements**: address shell, brand, hierarchy, CTA,
  representative content/card treatment, responsiveness, and component use
  only where live gaps require it; retain proportional scope and do not invent
  unresolved Campaign behavior.
- Feature spec, **Explicit non-goals** and **Acceptance criteria**: no full
  landing page; no Campaign backend/live integration; representative desktop
  and mobile must become judgeable; broad component changes require consumer
  analysis.
- `page-map.md`, **Shell & Benchmark Notes**: the resolved public-shell
  campaign navigation target is `/campaign`; its same section defers a footer.
- `product-design-principles.md`, **A.1 Trust before persuasion** and **A.7
  Do not imply what the product cannot prove**: donation persuasion must have
  real trust context and cannot create unsupported trust/impact/ranking claims.
- `patterns.md`, **B.1 List / Browse**: collection loading must use shaped
  skeletons; empty/error states require an explanation, and retry/CTA only
  when meaningful/authorized; raw backend errors and unsupported sort/filter
  semantics are prohibited.
- `design-guidelines.md`, **C. Implementation and Token Authority**, **G–Q**,
  and **Z. Responsive Visual Behavior**: use CSS-first canonical tokens,
  avoid container/card excess, preserve public-surface hierarchy, maintain
  action hierarchy, and reflow rather than merely compress desktop UI.
- Donation spec `docs/spec/5-donation/features/01-submit-donation-settlement.md`,
  **Request** and **Behavior / Submission**: actual amount floor is `≥ 5000`;
  payment-method values are recorded only in v1.

## Gap

1. The resolved navigation authority says `Jelajahi Kampanye` links to
   `/campaign`, while the actual live shell deliberately links it to
   `#kampanye` because that route is absent. This is a live source-versus-
   authority contradiction that affects the public navigation contract.
2. Static `HowItWorks` copy says a Rp10,000 minimum, but the current donation
   feature source says Rp5,000. It also makes “no hidden fees” and periodic-
   report claims for which this investigation has not located an owning
   authority. Those statements cannot be treated as harmless decorative copy:
   they communicate monetary and accountability product truth.
3. Existing content provides a functional shell/list/explainer slice but no
   recorded desktop/mobile rendered evidence. Whether its assembled hierarchy,
   placeholder campaign cards, and static trust language meet the foundation
   task’s human-judgement threshold remains unestablished.

## Sniffing

- **Risk:** The nav mismatch can leave a public primary navigation item unable
  to perform the defined cross-route action once users expect browse behavior.
  Static monetary/trust copy is higher impact: a false minimum or unsupported
  fee/reporting promise can mislead donors before any backend request occurs.
  Any change to `Button`, `Badge`, or `ProgressBar` would have shared-
  primitive blast radius beyond `/`.
- **Edge cases:** Campaign data can be loading, absent, failed, stale, have a
  missing id/title/progress field, a `null` `days_remaining`, or percentage
  outside 0–100. Current code handles these presentation cases; cards remain
  non-navigable even when populated, so an anchor CTA arrives at a collection
  that provides no detail action. On mobile the drawer and stacked action row
  are structurally present, but rendered tap/overflow/focus behavior still
  needs browser evidence.
- **Miscontext:** Code comments describe `/` as a previously planned landing
  scope and treat the anchor as temporary. The active feature spec instead
  makes the current `/` rendered calibration its subject; historical techplan
  comments are not the authority for deciding acceptance.
- **Misleading signals:** Tests and comments demonstrate that individual
  campaign cards intentionally avoid a verification badge, while the hero and
  explainer make broad verification/accountability claims. Passing component
  tests therefore does not establish that the whole public message is
  truthful, coherent, or judged in context. Token usage and responsive class
  names likewise are not rendered evidence.
- **Inconsistency:** `page-map.md` names `/campaign` for `Jelajahi Kampanye`;
  `nav-items.ts` substitutes `#kampanye`. The payment minimum conflicts
  directly: `HowItWorks` says Rp10,000, while the donation feature spec says
  Rp5,000. No resolution is selected in Stage 2.

## Code anchors

- `frontend/app/(public)/page.tsx` — `Home`: exact representative slice and
  section ordering.
- `frontend/app/(public)/layout.tsx` — `PublicLayout`: public shell, wordmark
  treatment, desktop nav, and auth-modal mounting boundary.
- `frontend/app/(public)/_components/nav-items.ts` — `publicNavItems`: source
  of the `/campaign` versus anchor contradiction.
- `frontend/app/(public)/_components/public-shell-client.tsx` —
  `PublicShellClient`: mobile nav, drawer and auth-trigger behavior to verify
  at mobile scope.
- `frontend/components/features/landing/hero.tsx` — `Hero`: primary public
  CTA and static trust copy.
- `frontend/components/features/landing/highlighted-campaigns.tsx` —
  `HighlightedCampaigns`: data-state behavior, list hierarchy, and current
  campaign-card composition.
- `frontend/components/features/landing/how-it-works.tsx` — `STEPS` and
  `HowItWorks`: static claims that require product-truth routing.
- `frontend/components/features/campaign/campaign-card.tsx` — `CampaignCard`
  and `CampaignCardSkeleton`: card information limits and placeholder/media
  behavior.
- `frontend/lib/hooks/use-campaigns.ts` — `useCampaigns`; `frontend/lib/api/
  campaign.ts` — `getCampaigns`: query/cache and contract boundary.
- `frontend/app/globals.css` — `:root` and `@theme inline`: established
  CSS-first visual-token authority in code.
- `frontend/components/ui/{button,badge,progress-bar}.tsx` — shared primitive
  contracts whose modification would trigger component-governance analysis.
