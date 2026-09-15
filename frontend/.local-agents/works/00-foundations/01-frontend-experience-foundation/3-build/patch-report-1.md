## What changed

`frontend/tests/browser/home.smoke.spec.ts` → added a bounded 15-second readiness wait for the named populated campaign heading before the existing semantic home assertions. A timeout now reports the viewport flow, URL, loading-status count, alert text, document readiness/visible text, and underlying Playwright wait failure.

`frontend/tests/browser/home.smoke.spec.ts` → isolated desktop and mobile populated-route checks as separate Playwright tests/browser contexts. Investigation showed that a second hard navigation in one context was controlled by the root-scoped app-shell service worker after it replaced the root-scoped MSW worker, causing `/campaigns` to reach Next.js and return 404. Context isolation preserves the approved behavior scope without changing production code, API/mock contracts, navigation, global timeouts, or using arbitrary sleeps.

## Tests run

`npx playwright test tests/browser/home.smoke.spec.ts --workers=1 --repeat-each=3` → focused repeated browser regression → initial diagnostic run failed all three populated-flow repetitions specifically on their second/mobile navigation, reporting the campaign error state and confirming `/campaigns` returned 404 after service-worker takeover; after viewport context isolation, passed 9/9.

`npm run test -- 'app/(public)/_components/hero.test.tsx' 'app/(public)/_components/highlighted-campaigns.test.tsx' 'app/(public)/_components/how-it-works.test.tsx' 'app/(public)/_components/public-shell-client.test.tsx' 'components/features/campaign/campaign-card.test.tsx'` → relevant component/state/interaction regression → passed, 5 files and 17 tests.

`npm run verify` → repository fast lint + unit/component baseline → passed, 40 files and 227 tests. The existing Vite native-config warning was emitted.

`npm run test:browser` → full browser regression → passed, 4/4 tests across the isolated desktop home, mobile home, mobile drawer, and existing login flows.

`git diff --check` → patch hygiene → passed.

## Verification scope confirmation

No race/concurrency, performance/load, or security-class test was executed in this Build iteration.

## Contract check

- [x] Current build target satisfied in full
- [x] Live-code re-grounding did not invalidate a material contract assumption

## Deferred / not tested here

Human rendered acceptance under R8 remains deferred to the required human decision. No production build was rerun because this patch changes only Playwright test orchestration and diagnostics; lint, the full fast baseline, focused repeated browser verification, and the full browser suite cover the patch surface.

## Flagged for Techplan / Testing

The development harness registers both MSW's `mockServiceWorker.js` and the app-shell `/sw.js` at the root scope. A second hard navigation in one browser context can therefore lose mock API interception after `/sw.js` replaces MSW, sending `/campaigns` to Next.js as a 404. This patch isolates the representative viewport cases, but future browser journeys that require multiple hard reloads in one context may need dedicated test-harness lifecycle work.

## Phase handoff

- Completed: Patch Plan Round 1 completed; the required home browser regression is bounded, diagnostic, and stable across repeated desktop/mobile execution.
- Artifacts: `frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation/3-build/patch-report-1.md`
- Open / blocked: No patch blocker. Human rendered acceptance and the previously documented public campaign-navigation decision remain external open items.
- Recommended next step: Return to the requesting Code Review phase after this patch.
- Session recommendation: CONTINUE for another Build iteration while focused; FRESH for Code Review/Testing.
- Context pointers: parent Techplan + `4-code-review/patch-plan-1.md` + `frontend/tests/browser/home.smoke.spec.ts` + tests recorded above only
