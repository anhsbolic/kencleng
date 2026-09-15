# Stage 2 — Design and route authority

## Current state

The agreed foundation feature deliberately leaves the minimum representative
`/` slice to Exploration. The current public-route authority resolves the
guest public shell for `/`: logo at left; `Beranda` (`/`) and `Jelajahi
Kampanye` (`/campaign`) links; `Masuk`/`Daftar` actions at right; and a mobile
hamburger/drawer with the established dashboard-drawer focus-management
behavior. A footer is explicitly deferred.

For the home content, the page map calls for browsing highlighted campaigns,
but says “highlighted” is mock-scope only: `GET /campaigns` has neither a
featured flag nor a sorting parameter. Its eventual product meaning is open
to the Campaign domain and must not be inferred here.

The design authority describes the target character as warm, trustworthy,
transparent, and calm, prioritizing trustworthiness then clarity over warmth
and delight. Its design-readiness guidance identifies a major landing-page
story or new brand-defining visual direction as OPEN; the feature spec itself
assigns Exploration responsibility for choosing the minimum slice.

## Requirement

- Feature spec, **Feature surface**, **Requirements**, and **Acceptance
  criteria**: `/` is a calibration surface; its minimum coherent slice is to
  be determined from current authorities; it must support human desktop/mobile
  judgement without completing the landing page or inventing Campaign rules.
- `docs/ui-ux/page-map.md`, **Shell & Benchmark Notes** and **Guest**: the
  resolved public-shell contents and the explicitly provisional highlighted-
  campaign mock scope govern route/navigation truth.
- `docs/ui-ux/product-design-principles.md`, **B. Design Readiness**, **C.
  Agent Design Responsibility**, and **D. Design Exploration Contract**:
  material unsettled landing/brand decisions require exploration and human
  approval rather than silent implementation invention.

## Gap

The authorities establish shell/navigation and constraints around campaign
content, but do not prescribe a complete home-page information hierarchy or
brand-expression composition. The live route and its existing visual system
must therefore be inspected to determine which limited rendered slice can be
judged while keeping campaign content and any brand-defining asset claims
explicitly provisional.

## Sniffing

- **Risk:** A polished home surface can accidentally assert a campaign ranking,
  recommendation, verification, or impact claim that has no Campaign-domain
  authority. This reaches public trust, not merely route styling.
- **Edge cases:** `GET /campaigns` may be empty, loading, or fail; the
  authority does not allow an invented “featured” predicate to distinguish
  these from a populated fixed mock set. The footer is also intentionally not
  required, so its absence is not evidence of a broken foundation.
- **Miscontext:** “Highlighted campaigns” may be read as a product-backed
  collection. The governing note says it is only a mock fixture convention.
- **Misleading signals:** A logo-looking text treatment or a prominent green
  campaign card can look like canonical brand identity, verification, or
  recommendation without having that authority.
- **Inconsistency:** The feature spec calls for an exploration-determined
  representative slice, while the page map defines only the shell and a
  provisional campaign-list framing. This is not a conflict: the narrower page
  map explicitly leaves home-story composition unresolved; it must remain a
  Stage-3 decision or a human product decision where material.

## Code and authority anchors

- `docs/spec/0-foundations/features/01-frontend-experience-foundation.md` —
  Feature surface, Requirements, Explicit non-goals, Acceptance criteria:
  task boundary and required evidence.
- `docs/ui-ux/page-map.md` — Shell & Benchmark Notes; Guest `/` row and its
  highlighted-campaign footnote: public shell contract and campaign-content
  limitation.
- `docs/ui-ux/product-design-principles.md` — B–D: readiness classification,
  decision authority, and the design-exploration questions later planning must
  answer.
- `docs/ui-ux/design-guidelines.md` — A–C, G–Q, W–Z: visual hierarchy/token,
  action, imagery, and responsive constraints to assess against live UI.
- `docs/ui-ux/brand-and-visual-assets.md` — G. Landing and Marketing Imagery;
  H. Logo and Brand Identity: required treatment if live `/` depends on an
  expressive hero or logo asset.
