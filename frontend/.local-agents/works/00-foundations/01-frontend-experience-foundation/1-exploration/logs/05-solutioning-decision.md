# Representative public foundation — Stage 3 decision record

> Phase/Stage: Exploration / Stage 3 — solutioning  
> Author: Codex  
> Created: 2026-09-16  
> Model: GPT-5  
> Reasoning: not exposed  
> Session: not exposed  
> Target revision: `a30ee75`  
> Workflow revision: candidate baseline `4199c6db1b26ef1920ba670f222aff0c6d0f9e59`

## Decision

Establish `/` as a deliberately small, static, text-led public calibration surface. It will contain a minimal public header, an editorial platform-level hero, a same-page primary action that leads to an early transparency/clarity explanation, and a compact closing invitation/route-local action treatment. It will establish the approved typography, warm-neutral surfaces, border/spacing-led grouping, restrained Sun/Berry action emphasis, and desktop/mobile composition in actual rendered code.

The content boundary is equally deliberate: use only platform-level, non-transactional explanatory copy derived from the approved design/public-surface authorities. Do not render invented campaigns, organizers, amounts, targets, progress values, outcomes, rankings, curation, verification badges, donation/payment actions, or claim-like documentary media. Any copy describing truth classes must frame them as clear information distinction rather than asserting unavailable current campaign facts. The page is a foundation calibration, not a mocked campaign marketplace.

## Chosen ownership and implementation direction

- Keep the first composition route-local under `app/` (with route-local supporting files only if they make the root composition readable). Do not create `components/ui/`, `components/shared/`, or a registry entry: one route cannot demonstrate a broad semantic contract.
- Extend `app/globals.css` with only the global visual values and font-role hooks actually consumed by this surface. Their canonical meanings/values remain `design-guidelines.md`; do not reconstruct a legacy token taxonomy.
- Use the approved Newsreader display / Instrument Sans UI-body role separation. The Techplan/Build must validate the chosen font-loading mechanism in the real build; system fallbacks alone would not establish the required role pairing.
- Use semantic document landmarks, visible text labels, ordinary focus behavior, and an in-page anchor rather than a non-functional navigation model. A plain text rendering of the product name is identification, not an attempt to establish an OPEN wordmark/logo.
- Use no campaign or expressive production asset in this slice. If visual judgement shows a meaningful Level 3 treatment is necessary, pause for an asset brief/lifecycle decision rather than substituting stock, generated beneficiary-like, or reference imagery.
- Keep the route Server Component unless a narrow interaction genuinely requires a Client Component. No API client, mock, query hook, form, URL state, Zustand store, or browser test is justified by the selected static scope.

## Verification direction

- Build should run lint, relevant component/behavior tests only if behavior beyond static navigation is introduced, and production build.
- Build must inspect the rendered route at representative desktop and mobile widths with realistic text wrapping. Check hierarchy, anchor reachability, focus/semantic landmarks, clipping, horizontal overflow, action clarity, and preservation of the early transparency explanation.
- The acceptance gate remains human real-browser review at desktop and mobile. It must judge whether the surface coherently expresses Sunlit Editorial / Evidence-Led Optimism without confusing design language for proof.
- Do not add Playwright by default. Reconsider it only if a stable responsive, focus/navigation, or browser-specific regression emerges; the Techplan must then record its benefit, omitted-risk, and authoritative phase.

## Options considered

### A. Complete public landing page with campaign discovery and donation CTA — rejected

This would show more of the selected public reference, but it would require live Campaign semantics/data or fabricated campaign, progress, curation, steward, and trust context. The feature explicitly excludes completing the landing page and resolving Campaign semantics. A visually polished mock would create false product authority.

### B. Minimal coherent static platform/trust slice — chosen

It directly satisfies the calibration purpose while allowing the visual system, public hierarchy, CTA restraint, responsive composition, and truth-first language to be judged. Its consequence is intentional deferral of functional discovery, donation, and campaign-data states to their owning feature/API work.

### C. Generic starter design system before the route — rejected

Creating Button/Card/Badge/Progress/Shell primitives first would recreate a retired-style component agenda without real consumers. It would establish broad contracts and future change impact before this first route demonstrates semantic need.

### D. Use selected-direction screenshots, generic stock imagery, or synthetic documentary imagery as hero material — rejected

The references are not production assets, and asset governance forbids visual material that could be mistaken for campaign evidence. Their use would also bypass the required asset lifecycle and, for a brand-defining treatment, human approval.

### E. Defer implementation until final logo, illustration family, and motion tokens are resolved — rejected

These are deliberately OPEN visual details, not blockers to the approved foundation. A textual identity treatment, no expressive asset, and no motion preserve the approved direction without silently deciding them.

## Consequences and deferred items

- **Deferred product scope:** campaign discovery data/listing, campaign detail, donation flow, public authentication, curation/ranking, real progress, organizer context, outcome/evidence display, and any API integration.
- **Deferred design/asset scope:** final logo/wordmark, production illustration family, real campaign-photography workflow, campaign placeholder system, exact provenance terminology, and motion tokens. These retain their documented OPEN status.
- **Risk to carry into Build/Test Focus:** This is a new global visual precedent. Verify that global font/color rules do not cause a generic-SaaS or decorative-trust treatment, and that narrow layouts retain the transparency explanation and reachable action without horizontal overflow. Treat the absence of campaign data as intentional; do not paper over it with illustrative facts.
- **No API contract gap to resolve now:** the selected slice intentionally does not consume Campaign data. If Build needs a campaign-specific claim or value to make the design work, it must stop and surface the owning spec/API gap.

## Decision sources

- `docs/spec/0-foundations/features/01-frontend-experience-foundation.md` — Goal, Requirements, Explicit non-goals, Acceptance criteria, Verification expectations.
- `docs/ui-ux/product-design-principles.md` — §§1–2, §5, §§7–8, §§12–15.
- `docs/ui-ux/brand-product-ui-brief.md` — §§4–8.
- `docs/ui-ux/design-guidelines.md` — §§3–13, §§20–23, §§26–29.
- `docs/ui-ux/asset-governance.md` — §§1–8.
- `docs/project/kencleng-frontend-tech-stack.md` — §§4, 6–10, 15–21.
- `frontend/components/README.md` — §§1–13.
