# Patch Plan 1 — Close R4 stale-state coverage gap

## Trigger

Independent Testing found that Techplan R4 requires observable coverage of the
background-revalidation state, but the current `HighlightedCampaigns` suite
only covers loading, populated, empty, error, and retry. The implementation
has a visible `isFetching && !isLoading` branch (`Data mungkin tidak terbaru.`)
with no corresponding test.

This is a **new coverage gap**, not a regression of the Build report's named
coverage. No production behavior defect was observed.

## Build-authority change

1. In `frontend/app/(public)/_components/highlighted-campaigns.test.tsx`, add
   one observable component test that first resolves the populated campaign
   fixture, then triggers a background refetch whose response is deliberately
   held pending.
2. Assert that the existing populated cards remain visible and the user-facing
   stale-data notice `Data mungkin tidak terbaru.` appears while the second
   request is in flight. Resolve the deferred response and assert the notice
   clears, so the test cannot pass solely from an initial loading state.
3. Do not add a test-only browser runtime interface, modify mock/API contracts,
   change query behavior, or change production files. Vitest/MSW is the
   Techplan-approved exceptional-state seam (D8).

## Required rerun after patch

```bash
npm run test -- 'app/(public)/_components/highlighted-campaigns.test.tsx'
npm run verify
npm run test:browser
```

R4's focused suite must pass before Testing reruns its affected sweep. Browser
coverage need not be expanded: it intentionally covers the populated route,
not exceptional query states.

## External (non-code) gate

R8 human rendered acceptance remains required after the patch. A human must
exercise `/` at 1280×800 and 390×844 and explicitly accept or reject
hierarchy, CTA clarity, responsive usability, truthful/provisional
presentation, and the no-media placeholder. Automation and agent inspection
cannot close this gate.
