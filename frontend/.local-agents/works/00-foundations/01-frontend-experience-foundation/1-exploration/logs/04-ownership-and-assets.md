# Stage 2 — Component ownership and asset boundary

## Current state

The public-route composition is route-local in `app/(public)`, but its three
home-only sections currently live in `components/features/landing/`. Repository
references show `Hero`, `HighlightedCampaigns`, and `HowItWorks` are imported
only by `app/(public)/page.tsx`; no evidence establishes cross-route or
domain-semantic reuse. `CampaignCard` is different: it is correctly
campaign-semantic and its documentation anticipates reuse by future campaign
browse surfaces.

`/` consumes existing `Badge`, `Button`, and `ProgressBar` UI primitives.
These primitives are registered as Foundation/Evolving contracts. Current
consumer discovery finds `Button` across account forms and dashboard shell,
`Badge` in the hero and dashboard notification treatment, and `ProgressBar`
only in campaign cards. No foundation-specific primitive variant has been
introduced.

Campaign cards use a `lucide` `ImageOff` icon inside a neutral fixed-ratio-like
placeholder region because the list schema currently has no media field. It
does not impersonate campaign photography, but it is visually a plain gray
missing-media treatment. There is no canonical campaign-placeholder asset in
the current frontend tree.

## Requirement

- Feature spec, **Requirements** and **Explicit non-goals**: maintain existing
  component ownership/impact rules; do not create global tokens, abstractions,
  or shared variants merely to polish this representative slice.
- `frontend/AGENTS.md`, **Component ownership and shared-change impact**:
  choose the narrowest truthful owner; route-specific composition belongs
  under `app/<route>/`; broad `ui`/`shared` changes require consumer discovery
  and representative verification.
- `frontend/components/README.md`, **A. Component Layers**, **G–L. Component
  contract/change discipline**: feature components require meaningful domain
  semantics, while generic UI contracts have the broadest blast radius and may
  not encode campaign-specific treatment.
- `brand-and-visual-assets.md`, **D. Truth and Representation** and **E.
  Placeholder System**: campaign imagery/placeholder treatment cannot
  fabricate evidence; a permanent missing-media treatment should become
  recognizable, intentional, layout-stable, and brand-compatible.

## Gap

1. The three `features/landing` components are currently used only for the
   `/` composition and do not express a business-domain concept. Their present
   location conflicts with the semantic-owner-first default for route-specific
   composition. Whether any section has demonstrated reuse deserving a
   broader owner is not established.
2. The campaign card’s no-media state preserves truth but does not meet the
   asset authority’s stated permanent placeholder ambition beyond a neutral
   icon and gray surface. The absence of a canonical asset/system is a
   foundation design/asset gap, not permission to substitute synthetic
   campaign imagery.
3. Any attempt to globally change `Badge`, `Button`, or `ProgressBar` for
   home-page polish would expand beyond the representative slice and requires
   separate consumer-impact analysis; current evidence does not establish
   that such a broad change is necessary.

## Sniffing

- **Risk:** Moving or broadening home sections without semantic evidence can
  establish a false shared pattern; altering primitives risks account forms,
  dashboard navigation, and future campaign surfaces. A generic or synthetic
  card image risks being understood as documentary evidence for a real
  campaign.
- **Edge cases:** The placeholder must remain legible/stable for every list
  state and long title, and must not become an upload affordance. A future
  media-capable campaign schema needs a distinct loading/missing/error asset
  behavior, not a silent reuse of today’s no-data icon.
- **Miscontext:** `components/features/` being available does not make any
  presentational landing section a domain feature. Likewise, a standard icon
  is permitted for utility treatment but is not evidence that a brand-level
  placeholder system exists.
- **Misleading signals:** The current `CampaignCard` test confirms that no
  upload affordance is present, which proves only the negative behavioral
  boundary. It does not prove that the placeholder is a judged, canonical, or
  brand-coherent asset. Existing primitive tests do not prove their use is the
  correct foundation expression in a combined page.
- **Inconsistency:** `frontend/AGENTS.md`/component governance prescribe
  route-local ownership for route-only composition, while all three home-only
  sections are located under `features/landing`. The asset guidance says avoid
  a permanent gray/broken-image-like missing-media treatment; the live card is
  intentionally neutral but not yet a defined placeholder system. No
  reclassification or asset direction is chosen in Stage 2.

## Code anchors

- `frontend/app/(public)/page.tsx` — sole current importer of the three
  landing sections; starting point for ownership verification.
- `frontend/components/features/landing/{hero,highlighted-campaigns,how-it-works}.tsx`
  — route-specific section candidates whose current owner needs a deliberate
  assessment.
- `frontend/components/features/campaign/campaign-card.tsx` — campaign-
  semantic card and missing-media presentation boundary.
- `frontend/components/ui/button.tsx`, `badge.tsx`, `progress-bar.tsx` —
  current broad contracts; their direct consumers were discovered before any
  potential change.
- `frontend/components/README.md` — A, Current registry, G–L, and T–U:
  owner selection, consumer-impact, primitive, and visual-asset governance.
- `docs/ui-ux/brand-and-visual-assets.md` — D–E and K–N: truth-preserving
  placeholder and asset-status/change authority.
