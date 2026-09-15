# Stage 3 — Solutioning and decisions

## Chosen direction

Calibrate the already-existing `/` as the smallest representative foundation
slice: public shell, restrained text-first hero, the fixed mock-backed campaign
collection, and a compact explanatory section. Keep it deliberately
incomplete: no footer, complete marketing story, campaign detail flow, live
Campaign integration, ranking semantics, or new expressive hero/key-art asset.

The implementation direction is a localized route calibration, not a design
system rewrite:

1. Correct and narrow public claims to current domain authority. The donation
   floor must agree with the current `≥ 5000` contract. Remove or neutralize
   fee and reporting-frequency promises until their owning specifications are
   identified. A broad verified-organization statement may remain only as the
   published-campaign policy supported by the organization-verification and
   auto-unpublish rules; it must not turn into per-card verification data,
   scores, ranking, or impact evidence not available from the list API.
2. Make the representative public hierarchy judgeable with existing tokens and
   local composition: shell → primary browse CTA → campaign collection →
   supporting explanation. The primary CTA may continue to reach the in-page
   collection while no campaign detail route exists; cards must not be made to
   navigate to a non-existent detail route.
3. Return the three route-only landing sections to route-local ownership (or
   document an actual future reuse before retaining a broader owner). Retain
   `CampaignCard` in `features/campaign/`; do not change `Button`, `Badge`, or
   `ProgressBar` unless rendered evidence demonstrates a cross-product
   contract need.
4. Keep the card missing-media treatment explicitly provisional and
   truth-preserving. Do not introduce synthetic campaign photos, generic AI
   decoration, or unapproved key art merely to make the page feel finished.
   An intentional brand-placeholder asset needs a separate asset brief and,
   if brand-defining, human approval.
5. Correct the root document language to Indonesian if repository-wide page
   language confirms that scope; otherwise localize language declaration at
   the narrowest valid App Router boundary. This is an accessibility correction,
   not a visual-language decision.
6. Add narrowly scoped `/` browser evidence at 1280×800 and 390×844: populated
   route, no horizontal overflow, visible core hierarchy, hero CTA reaching
   the collection, and mobile drawer open/close/keyboard behavior. Exercise at
   least one async/list exceptional state where the existing mock harness can
   make it deterministic. Treat this as repeatable evidence, not replacement
   for required human desktop/mobile rendered acceptance.

## Authority conflict requiring human/project routing

The page map resolves `Jelajahi Kampanye` to `/campaign`; current campaign
delivery is `NOT_STARTED`, the route is absent, and live navigation points to
`#kampanye`. Build must not silently make the link broken or rewrite the page
map to match current code. Until a human/project owner decides whether the
foundation task may retain the anchor as a formal temporary exception, this
remains open. The rest of the localized calibration can proceed independently.

## Alternatives considered and rejected

### Complete or substantially redesign the landing page

Rejected. It would answer the foundation task with an unapproved landing
story, likely broaden component/asset scope, and violates the explicit
non-goal of completing the landing page. The selected slice is sufficient to
calibrate the inherited shell, visual hierarchy, CTA treatment, cards, and
responsive behavior once rendered evidence is collected.

### Link navigation/cards to `/campaign` or campaign detail now

Rejected. `/campaign` does not exist and Campaign delivery is `NOT_STARTED`.
Linking to it creates a known broken public path; implementing it expands into
the Campaign route/domain scope. Keep the current card non-navigation boundary
and escalate the nav-authority contradiction rather than inventing a route.

### Add a shared primitive variant/global token system to polish `/`

Rejected. Existing token and primitive contracts already supply the required
hierarchy. No rendered evidence identifies a genuine cross-product deficiency,
while `Button` has many account/dashboard consumers and `Badge`/`ProgressBar`
are Evolving contracts. A local route adjustment is lower blast radius and
matches semantic-owner-first governance.

### Generate hero imagery, a logo, or campaign photography

Rejected for this work unit. The available authorities classify landing key art
and brand-defining identity as an asset/design exploration requiring a brief
and human approval; campaign images cannot imply real beneficiaries or impact.
The chosen text-first direction preserves the present calibration scope and
does not silently turn a provisional decorative output into product identity.

### Preserve every existing marketing claim because a prior landing task wrote it

Rejected. Historical task comments/tests are not product authority. The
Rp10,000 minimum directly conflicts with the active donation feature source;
the fee/reporting claims have no source identified in Stage 2. Public trust and
money language must route to current domain truth.

## Consequences and verification focus

- **Scope consequence:** likely production changes stay within `frontend/app/(public)/`,
  its route-local section files/tests, `app/layout.tsx` or the narrowest
  language boundary, and `tests/browser/`; no backend/API change is implied.
- **Campaign contract consequence:** retain the real `GET /campaigns` type and
  MSW fixture shape; do not add featured/sort/organization/media fields or
  recalculate progress/money in the client.
- **Shared-component consequence:** a proposed primitive change must first
  repeat consumer discovery and select representative account/dashboard
  consumers. Current direction deliberately avoids that branch.
- **Rendered test focus:** normal populated state at both widths, long fixture
  title/card wrapping, horizontal overflow, hero-to-collection CTA, drawer
  focus/Escape/open-close, and at least one list error/empty/retry condition.
- **Human acceptance focus:** warm/trustworthy/calm hierarchy; whether the
  restrained no-media cards look intentional rather than unfinished; clear
  primary vs secondary actions; mobile reachability; and whether provisional
  trust language remains honest in context.

## Context anchors for Techplan

- `docs/spec/0-foundations/features/01-frontend-experience-foundation.md`
  — task boundaries and acceptance requirements.
- `docs/ui-ux/page-map.md` — public shell and unresolved navigation conflict.
- `docs/ui-ux/product-design-principles.md` — trust/no-implied-claims and
  OPEN/PARTIAL decision routing.
- `docs/ui-ux/design-guidelines.md` — visual hierarchy/responsive/tokens.
- `docs/ui-ux/brand-and-visual-assets.md` — truthful placeholders and
  brand-asset approval boundary.
- `docs/spec/5-donation/features/01-submit-donation-settlement.md` — current
  minimum amount and payment-method truth.
- `docs/spec/3-organization/features/04-organization-edit.md` — verified-org
  status reversal auto-unpublishes campaigns.
- `docs/project/kencleng-development-tracker.md` — Campaign is `NOT_STARTED`;
  frontend foundation is the current validation run.
- `frontend/app/(public)/page.tsx`, `layout.tsx`, and `_components/nav-items.ts`
  — representative route/shell/nav anchors.
- `frontend/components/features/landing/` and `features/campaign/campaign-card.tsx`
  — ownership and content-card anchors.
- `frontend/tests/browser/login.smoke.spec.ts` and `playwright.config.ts` —
  established browser-test style and harness.
