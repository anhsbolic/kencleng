# Patch Plan — Round 1

**Build authority required.** This plan intentionally does not modify production code.

## Blocking Finding 1 — Stabilize the required home browser regression

**Target:** `frontend/tests/browser/home.smoke.spec.ts:9-49` (and `frontend/playwright.config.ts:7-9` only if a narrowly scoped per-test readiness timeout cannot express the required wait clearly).

1. Preserve the current behavior scope: desktop and mobile populated home surface, `lang="id"`, no horizontal overflow, hero anchor reaching `#kampanye`, and mobile drawer Escape/focus return.
2. Add an explicit, bounded route-readiness wait for the user-visible home surface before running the existing semantic assertions. Its timeout must accommodate development-server compilation/initial route hydration without exceeding the test's existing 30-second budget or masking an indefinitely unresolved route.
3. Keep the existing semantic assertions (roles, names, heading levels, anchor, overflow, dialog and focus). Do not replace them with arbitrary sleeps, DOM implementation selectors, or an assertion that only a skeleton appears.
4. Retain page-error capture for the populated-route flow and ensure a readiness timeout reports enough context to distinguish a slow startup from a genuine route failure.
5. Run, and record actual output for:
   - `npx playwright test tests/browser/home.smoke.spec.ts --workers=1 --repeat-each=3`
   - `npm run test:browser`
   - the relevant component tests and the repository-prescribed fast verification command when the environment permits.
6. If the test still fails after bounded readiness synchronization, investigate the dev/MSW/provider lifecycle as a separate root cause; do not increase timeouts indefinitely.

## Acceptance criteria

- The focused home browser suite passes repeatedly at both 1280×800 and 390×844.
- The full browser suite passes.
- The test still fails promptly and diagnostically when the route cannot reach its populated state.
- No production source, API contract, mock data shape, or public navigation authority is broadened to address a test-harness issue.
