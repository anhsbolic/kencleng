## What changed

`frontend/app/(public)/_components/` → moved `Hero`, `HighlightedCampaigns`, and `HowItWorks` plus their colocated tests into the owning public route; preserved the `Hero → HighlightedCampaigns → HowItWorks` order and kept `HighlightedCampaigns` as the sole campaign-fetching client leaf.

`HowItWorks` → replaced unsupported minimum-donation, fee, and periodic-report promises with the Techplan's three fixed authority-bounded messages; added positive and prohibited-copy assertions.

`frontend/app/layout.tsx` → changed the document language declaration from `en` to `id` without changing providers, fonts, metadata, or auth boundaries.

`frontend/tests/browser/home.smoke.spec.ts` → added populated `/` coverage at 1280×800 and 390×844 for the core surface, Indonesian document language, horizontal overflow, the `#kampanye` CTA, and mobile drawer visibility/Escape focus return.

Rendered `/` inspection at 1280×800 and 390×844 → confirmed coherent hierarchy, readable fixed copy, responsive single-column mobile composition, populated no-media campaign cards, and no objective clipping/overflow defect requiring a local class correction. The Next.js development indicator was present in inspection captures and is not production page content.

## Tests run

`npm run test -- 'app/(public)/_components/hero.test.tsx' 'app/(public)/_components/highlighted-campaigns.test.tsx' 'app/(public)/_components/how-it-works.test.tsx' 'app/(public)/_components/public-shell-client.test.tsx' 'components/features/campaign/campaign-card.test.tsx'` → focused component/state/interaction regression → passed, 5 files and 17 tests.

`npm run verify` → repository fast lint + unit/component baseline → passed, 40 files and 227 tests. The existing Vite native-config warning was emitted.

`npm run build` → production compile and server/client import-boundary check → first sandboxed attempt could not fetch configured Google Fonts; rerun with network access passed and generated all routes.

`npm run test:browser` → requested browser regression → first runnable suite had both new home tests pass while the pre-existing login smoke transiently rendered blank at its mobile step; `npx playwright test tests/browser/login.smoke.spec.ts --workers=1` then passed, and a full `npm run test:browser` rerun passed all 3 tests.

Playwright screenshots at 1280×800 and 390×844 after waiting for `h1` → Build-time rendered inspection → passed with no objective responsive/hierarchy correction identified.

`git diff --check` → patch hygiene → passed.

## Verification scope confirmation

No race/concurrency, performance/load, or security-class test was executed in this Build iteration.

## Contract check

- [x] Current build target satisfied in full
- [x] Live-code re-grounding did not invalidate a material contract assumption

## Deferred / not tested here

Human rendered acceptance under R8 remains deferred because automation and agent inspection cannot satisfy the required human decision. Independent Testing may reuse the representative desktop/mobile scope; specialized race/concurrency, performance/load, and security-class testing was not triggered by this presentation-only change.

## Flagged for Techplan / Testing

Human acceptance must explicitly accept or reject hierarchy, CTA clarity, responsive usability, truthful/provisional presentation, and no-media placeholder appropriateness. The active public campaign-navigation authority conflict remains unchanged; `publicNavItems` still uses `#kampanye` as required by this Build target.

## Phase handoff

- Completed: Approved Frontend Experience Foundation Techplan build target completed; human rendered acceptance remains an external gate, not unfinished implementation.
- Artifacts: `frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation/3-build/report.md`
- Open / blocked: Human rendered acceptance and the pre-existing public campaign-navigation authority decision remain open; Build implementation is not blocked.
- Recommended next step: Code Review after initial build
- Session recommendation: FRESH for Code Review/Testing
- Context pointers: parent Techplan + `frontend/app/(public)/page.tsx` + `frontend/app/(public)/_components/{hero,highlighted-campaigns,how-it-works}*` + `frontend/app/layout.tsx` + `frontend/tests/browser/home.smoke.spec.ts` + tests recorded above only
