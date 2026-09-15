# Code Review Findings — Round 1

**Scope reviewed:** current frontend diff for the Frontend Experience Foundation home-route slice: relocation of route-owned landing sections and tests, authority-bounded explainer copy, document language, and new browser smoke coverage.

**Authority read:** root and `frontend/AGENTS.md`, `frontend/components/README.md`, `frontend/README.md`, the active Techplan, Harscode Code Review guidelines/checklist, and the targeted React best-practice authorities routed from `best-practices/index.md`.

## 1. Safety

No findings.

The sole client leaf remains `HighlightedCampaigns`; it imports no server-only or sensitive configuration, retains the existing network/query boundary, and surfaces a generic retryable failure rather than raw request data. No new concurrency, cancellation, resource-lifecycle, money, PII, authentication, or hostile-content surface was introduced. The Techplan Test Focus Pointer already covers the applicable public-trust and responsive/focus risks, so there is no Safety-driven Techplan drift.

## 2. Quality

No findings.

The relocation removes the obsolete feature-layer owner cleanly, preserves the existing local composition and client boundary, and the static copy remains simple and explicit. No material duplication, misleading naming, dead production code, or unnecessary abstraction was introduced.

## 3. Stack-Specific Best Practices

### Finding 1 — Browser regression is not reliably executable at its intended representative viewport scope

- **Location:** `frontend/tests/browser/home.smoke.spec.ts:33-49`, especially `expectHomeSurface` at lines 9-30; `frontend/playwright.config.ts:7-9` supplies the default 5-second assertion timeout.
- **Problem:** The new required smoke test is flaky under its standard `next dev` harness. A targeted run failed on the mobile step while it remained on the campaign skeleton for more than five seconds; a repeated run (`--repeat-each=3`) failed again when the page had not rendered the `h1` within the same timeout. The remaining drawer scenario passed, and other repeated iterations passed.
- **Why it matters:** A browser regression that intermittently fails before its asserted route state is ready cannot provide dependable responsive or focus-regression evidence. It will also create noisy CI failures that obscure an actual defect.
- **Suggested resolution:** Build should make the browser test wait explicitly for a route-ready, user-visible condition with a timeout appropriate to the configured development-server startup/compile behavior (while retaining the 30-second test budget), then re-run the focused smoke repeatedly and the full browser suite. Keep the semantic assertions and do not paper over real page errors or a permanently unresolved loading state.
- **Blocking:** **Blocking.** The Techplan explicitly requires the route browser regression at desktop and mobile (R1-R3); the current added regression must pass reliably before it can serve as that evidence.
- **Best-practice source:** [`react/testing-automation-boundary.md`](/home/anhar-solehudin/kencleng-workspace/harscode-workspace/best-practices/react/testing-automation-boundary.md) (objective rendered outcomes need executable verification) and [`react/visual-verification.md`](/home/anhar-solehudin/kencleng-workspace/harscode-workspace/best-practices/react/visual-verification.md) (rendered UI verification must use representative, failure-revealing contexts). The targeted companion checks in `accessibility-fundamentals.md`, `server-client-component-boundary.md`, `app-router-routing-conventions.md`, `data-fetching-conventions.md`, `component-composition-and-abstraction.md`, `component-test-mocking-discipline.md`, `loading-empty-error-state-conventions.md`, and `responsive-layout-robustness.md` produced no additional finding.

## 4. Consistency

### Finding 1 — The required Techplan browser evidence is currently flaky

- **Location:** `frontend/tests/browser/home.smoke.spec.ts:33-49`.
- **Problem:** The new test does not consistently establish the populated home route at the two required viewports, as demonstrated by the failed mobile iterations described above.
- **Why it matters:** It does not yet meet the reviewed execution contract's required R1-R3 automated browser evidence.
- **Suggested resolution:** Apply the Build patch plan for the readiness synchronization, then demonstrate stable targeted and full browser-suite results.
- **Blocking:** **Blocking.**
- **Target-repo authority:** [`techplan.md §12 — R1-R3 and Testing Checklist`](/home/anhar-solehudin/kencleng-workspace/kencleng/frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation/2-techplan/techplan.md) requires desktop/mobile route coverage, no overflow, `#kampanye` behavior, and mobile drawer Escape/focus-return evidence. [`frontend/AGENTS.md` browser-test guidance](/home/anhar-solehudin/kencleng-workspace/kencleng/frontend/AGENTS.md) defines committed browser tests as behavior-scoped automation assets.

## Verdict

**Request changes**

Blocking finding: Finding 1 (the non-deterministic required browser regression), reported in both its stack-specific and target-repo-authority dimensions.

## Verification evidence

- `npx vitest run 'app/(public)/_components/hero.test.tsx' 'app/(public)/_components/highlighted-campaigns.test.tsx' 'app/(public)/_components/how-it-works.test.tsx' --pool=forks --maxWorkers=1` — passed: 3 files, 9 tests. Vite emitted the pre-existing native-config warning.
- `npx playwright test tests/browser/home.smoke.spec.ts --workers=1` — failed: populated mobile step timed out waiting for its campaign card; mobile drawer test passed.
- `npx playwright test tests/browser/home.smoke.spec.ts --workers=1 --repeat-each=3` — failed: 5 passed, 1 failed when the mobile iteration timed out before the `h1` rendered.
- `npm run verify` — not completed in this review session: the full Vitest process produced no result after starting and was interrupted after repeated 30-second observation windows. This is recorded as **not verified**, not a passing result.
- Initial browser execution inside the sandbox could not bind `127.0.0.1:3000` (`EPERM`); the targeted browser checks above were rerun with approved local-server permission.

## Phase handoff

- **Completed:** four-pass review + verdict
- **Artifacts:** `frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation/4-code-review/review-findings-1.md`; `frontend/.local-agents/works/00-foundations/01-frontend-experience-foundation/4-code-review/patch-plan-1.md`
- **Open / blocked:** Finding 1 — reliable execution of the required home browser regression.
- **Recommended next step:** Build/Patch
- **Session recommendation:** BUILD authority for patches; FRESH for Testing after the patch
- **Context pointers:** Techplan §4 R1-R3 and §12; `frontend/tests/browser/home.smoke.spec.ts:9-49`; Finding 1 and patch plan 1
