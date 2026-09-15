## What changed

`frontend/app/(public)/_components/highlighted-campaigns.test.tsx` → added observable R4 coverage for a populated campaign collection undergoing background revalidation. The test seeds the real default populated fixture, holds the second MSW response pending, verifies both existing cards and `Data mungkin tidak terbaru.` remain visible, then releases the successful cached-shape response and verifies the notice clears.

Test support in the same file → exposed the existing per-test `QueryClient` construction through a small helper so the test can invalidate the established campaign query key. No production file, query behavior, API/mock contract, browser runtime interface, or navigation behavior changed.

## Tests run

`npm run test -- 'app/(public)/_components/highlighted-campaigns.test.tsx'` → focused R4 component/state regression → passed, 1 file and 6 tests.

`npm run verify` → repository fast lint + unit/component baseline → passed, 40 files and 228 tests. The existing Vite native-config warning was emitted.

`npm run test:browser` → prescribed unchanged browser regression → first run passed all 3 home tests but the pre-existing login smoke hit its known second-navigation service-worker timing race; immediate full rerun passed all 4 browser tests.

`git diff --check` → patch hygiene → passed.

## Verification scope confirmation

No race/concurrency, performance/load, or security-class test was executed in this Build iteration.

## Contract check

- [x] Current build target satisfied in full
- [x] Live-code re-grounding did not invalidate a material contract assumption

## Deferred / not tested here

R8 human rendered acceptance remains deferred because automation cannot supply the required human decision. No production build or new rendered inspection was run because this patch changes only a component test and does not change runtime output, layout, imports, or production behavior.

## Flagged for Techplan / Testing

The previously documented development service-worker overlap remains observable: the first full browser run failed only the unrelated login smoke's second navigation, then the unchanged full suite passed on rerun. Future stabilization of that login test or the shared browser harness requires its own authorized patch; it was not broadened into this R4-only Testing patch.

## Phase handoff

- Completed: Testing Patch Plan 1 completed; R4 now covers the stale/background-revalidation state through settlement.
- Artifacts: `frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation/3-build/patch-report-2.md`
- Open / blocked: No patch blocker. R8 human rendered acceptance remains an external gate; the unrelated login browser timing race remains flagged.
- Recommended next step: Return to the requesting Testing phase after this patch.
- Session recommendation: CONTINUE for another Build iteration while focused; FRESH for Code Review/Testing.
- Context pointers: parent Techplan + `5-testing/patch-plan-1.md` + `frontend/app/(public)/_components/highlighted-campaigns.test.tsx` + tests recorded above only
