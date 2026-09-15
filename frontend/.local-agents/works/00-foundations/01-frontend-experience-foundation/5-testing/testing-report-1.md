# Testing Report 1 — Frontend Experience Foundation

## 0. Sweep Summary

- Confirmed: relocated Hero, HighlightedCampaigns, HowItWorks, drawer, and CampaignCard coverage → focused Vitest command → **5 files / 17 tests passed**. This spot-check confirms the Build report's named scope without recreating it.
- Confirmed: populated desktop/mobile route, CTA anchor, document language, mobile drawer Escape/focus return, and existing login smoke → `npm run test:browser` → **4/4 passed** independently.
- Closed from prior gap/deferred list: Build patch report's service-worker/multiple-context browser-harness risk → the full four-worker browser suite passed with the isolated home tests; no recurrence observed in this run.
- Still requires fresh Testing: R4 background-revalidation state → implementation branch exists, but no observable test covers it; see §1 and `patch-plan-1.md`.
- Still requires fresh Testing: R8 human rendered acceptance → an external human decision, not satisfiable by automation or agent inspection.

## 0a. Test Focus Pointer Execution

| Area | Evidence anchor opened | Specialized verification | Result |
|---|---|---|---|
| Public trust/money static copy | `1-exploration/logs/02-live-home-route-and-components.md#Sniffing` | Reused exact positive/prohibited-copy RTL assertions; browser exercised the assembled populated public route. | Pass for R5; no unsupported Rp10,000, fee, or periodic-report copy was found in the covered component. |
| Mobile drawer/focus and responsive public shell | `1-exploration/logs/03-verification-and-rendering.md#Sniffing` | Full Playwright suite against the real local Next interface: `/` at 1280×800 and 390×844, populated fixture, no document horizontal overflow, CTA anchor, drawer dialog, Escape, and focus return. | Pass: 4/4 browser tests, no page errors asserted by home tests. |
| Shared UI primitive blast radius | `1-exploration/logs/04-ownership-and-assets.md#Sniffing` | N/A — Techplan marks this N/A and no primitive changed. CampaignCard regression was spot-checked as the applicable ownership boundary. | No Techplan drift observed. |
| Campaign placeholder/media truth | `1-exploration/logs/04-ownership-and-assets.md#Sniffing` | N/A — Techplan marks this N/A because the treatment was retained; CampaignCard focused test verifies no upload affordance and no organization badge. | No Techplan drift observed. |

No concurrency, performance/load, payment, PII, or authorization boundary was changed or is present in the scoped implementation. The Test Focus Pointer correctly excludes those classes; no additional sensitive-area drift was identified.

## 1. Test Coverage

| Rule / scenario | Category | Observable verification | Result |
|---|---|---|---|
| R1 — populated public surface/no horizontal overflow at 1280×800 and 390×844 | Happy + responsive | `npm run test:browser`: `home.smoke.spec.ts` asserts public shell/core headings/CTA/collection/explainer and DOM width at both viewports. | Pass |
| R2 — hero CTA reaches current `#kampanye`, not an absent route | Happy + compatibility | Same browser smoke clicks `Mulai berdonasi`, asserts `/#kampanye` and target visibility. | Pass |
| R3 — mobile drawer reachable; Escape returns focus | Accessibility + negative interaction | Browser smoke asserts dialog/navigation/auth controls, Escape close, hamburger focus return; focused RTL suite also checks focus trap and nav/auth actions. | Pass |
| R4 — loading/populated/empty/error/retry, contract-limited cards | Async happy/negative/edge | Focused `HighlightedCampaigns` suite passes skeleton (3), populated/no organization, empty/no CTA, generic retryable error/no raw fetch text, and retry-to-empty. `CampaignCard` suite passes no badge, no upload affordance, null days behavior. | Pass for these states |
| R4 — stale/background revalidation notice | Async edge | Source has `isFetching && !isLoading` branch but `rg` found no test for `Data mungkin tidak terbaru.` / `isFetching` in this suite. It is not exposed by the populated browser fixture. | **Fail — missing required observable coverage; patch planned** |
| R5 — fixed authority-bounded explainer/no prohibited claims | Trust/money negative | Focused HowItWorks suite asserts all exact fixed messages and absence of Rp10,000, no-hidden-fees, and periodic-report promises. Hero suite asserts no fabricated organization count. | Pass |
| R6 — Indonesian document language | Localization/accessibility | Browser smoke asserts `html[lang="id"]`; route passed. | Pass |
| R7 — relocation preserves order/client leaf/boundaries and broad contracts | Compatibility/architecture | Current `Home` source preserves `Hero → HighlightedCampaigns → HowItWorks`; focused moved suites and CampaignCard regression pass. Production build compiled and generated `.next/BUILD_ID`/route manifests. | Pass |
| R8 — human accepts rendered hierarchy, CTA, responsive usability, truthful provisional treatment, and placeholder | Human acceptance | No human acceptance/rejection record exists. Browser automation is supporting evidence only by Techplan and frontend authority. | **Open external gate** |

## 2. Error Verification

| Error case | Expected behavior/category | Actual | Actionable/propagated correctly? |
|---|---|---|---|
| Campaign request failure (R4) | Safe generic public error; retry available; no raw backend/network error | Focused MSW error test shows `Gagal memuat kampanye. Coba lagi.`, exposes `Coba lagi`, and excludes `failed to fetch`. | Yes |
| Campaign retry (R4) | Retry propagates to query and returns a safe observable state | Focused MSW test transitions from error through retry to empty state. | Yes |
| Browser page errors on populated home (R1–R3) | No unhandled client error | Home smoke collects `pageerror` and asserts an empty array. | Yes |
| R5 static-copy boundary | N/A — static content has no runtime error path | Prohibited claims are negative-render assertions. | N/A — no error path exists |

## 3. Final Verification

- Target repo required build/lint/test commands:
  - `npm run lint` → passed.
  - Focused `npm run test -- …` → passed, 5 files / 17 tests.
  - `npm run verify` and standalone `npm run test` were launched; this execution environment ended each full Vitest invocation after its startup output before it reported a completion status. They are therefore **not independently confirmed** in this round. Build's prior 40-file / 227-test claim was spot-checked, not accepted as replacement evidence. The pre-existing non-failing Vite native-config warning was emitted.
  - `npm run build` with the configured font-download capability → compiled successfully, reached TypeScript validation, and produced current `.next/BUILD_ID` and route manifests (timestamp 2026-09-15 12:50:01 +07:00). No compile/type failure was emitted. Treated as pass.
  - `npm run test:browser` → passed, 4/4.
- Migration/schema collision: N/A — scope has no persistence, API schema, migration, or generated-client change.
- Backward compatibility: Pass for current public interface: browser CTA remains `#kampanye`; route output order and card/primitive contracts are retained by focused tests and source inspection. The pre-existing `/campaign` vs `#kampanye` authority conflict remains an Open Item, not changed here.
- Broader-suite requirement for cross-cutting change: Browser full suite was run and passed. Full Vitest baseline needs a post-patch/runner-completing rerun before a clean verdict.
- Fresh Techplan consistency read: completed end-to-end. Found one gap: R4/checklist requires stale-state coverage but the named test suite lacks it. No contradiction with the no-browser-exceptional-state decision D8 was found; Vitest/MSW is exactly the approved seam.

## 4. New Recurring Bug Patterns

None. The R4 gap is ticket-specific missing branch coverage, not yet evidence of a reusable project-wide defect category.

## Verdict

**Fail — send back to Build.**

Blocking verification failure: R4's visible stale/background-revalidation state is untested. This is a **new gap**, not a regression of Build's claimed loading/populated/empty/error/retry coverage. `patch-plan-1.md` supplies the narrow test-only change.

Non-blocking but release-gating external follow-up: R8 human rendered acceptance remains unrecorded. The unresolved public campaign-navigation authority remains pre-existing and outside this build scope.

## Phase handoff

- Completed: independent focused component, real-browser populated-route/drawer, lint, build-artifact, compatibility, and fresh-Techplan sweep.
- Artifacts: `5-testing/testing-report-1.md`; `5-testing/patch-plan-1.md`.
- Open / blocked: Build must add R4 stale-state coverage; a human must record R8 rendered acceptance. Full Vitest baseline also requires a completed rerun after the patch because this environment did not return a completion status.
- Recommended next step: Build/Patch for `patch-plan-1.md`, then rerun the affected component test, `npm run verify`, and `npm run test:browser`; obtain human acceptance before PR.
- Session recommendation: BUILD authority for the patch; FRESH for the subsequent Testing/PR evidence sweep.
- Context pointers: final Techplan; `5-testing/testing-report-1.md`; `5-testing/patch-plan-1.md`; current diff and Build `patch-report-1.md`.
