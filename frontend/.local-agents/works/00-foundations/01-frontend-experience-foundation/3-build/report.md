# Build Report — Frontend Experience Foundation

> Phase: Build  
> Author: Codex  
> Created/Updated: 2026-09-16  
> Model: GPT-5  
> Target revision: `a30ee75307ff14a7053dbdb749dae030fbcbb727`  
> Workflow revision: `4199c6db1b26ef1920ba670f222aff0c6d0f9e59`

## What changed

- `app/page.tsx` → replaced the reboot placeholder with a route-local, static public calibration surface: semantic header, editorial hero, early clarity explanation, valid same-page links, and compact closing action. Copy stays at platform level and presents no Campaign, donation, verification, or outcome claim.
- `app/layout.tsx` → loaded Newsreader and Instrument Sans through `next/font/google`, retained Indonesian document language, and updated root metadata.
- `app/globals.css` → added only values and rules used by this surface: warm neutrals, restrained Sun/Berry accents, type roles, border/spacing-led composition, responsive layout, focus outline, and reduced-motion behavior.
- `app/page.test.tsx` → added observable coverage for public landmarks, the primary heading, and the visible fragment link to the existing clarity section.

## Tests run

- `npm run verify` → fast lint and focused component behavior → passed (1 test). Build owns execution of the authored test to show that the anchor/landmark contract is exercised; Testing owns an independent final repeat.
- `npm run build` → production/static and font integration → passed outside the network-restricted sandbox. The initial sandbox run failed only because Google Fonts could not be reached; the network-enabled rerun compiled, type-checked, and prerendered `/` as static content. This belongs in Build because the new font path crosses Next build processing.
- Chromium inspection of built `/` at 1440×900 and 375×812 → rendered desktop/mobile behavior → passed. `documentElement.scrollWidth` equaled viewport width at both sizes; the primary link changed the hash to `#kejelasan` and brought the section near the viewport top (35px desktop, 27px mobile); keyboard focus showed a solid outline. Both full-page renders were inspected for hierarchy, wrapping, clipping, and action reachability.
- `git diff --check` → patch whitespace integrity → passed.

## Verification scope confirmation

No race/concurrency, performance/load, or security-class test was executed in this Build iteration. No broad Testing-owned suite was executed. `npm run verify` is the repository's fast baseline and was run to execute the new authored test and lint; independent Testing still owns its final repeat.

## Contract check

- [x] Current build target satisfied in full at Build scope; human rendered acceptance remains a separate required gate before task acceptance.
- [x] Live-code re-grounding did not invalidate a material contract assumption.

## Deferred / not tested here

- Independent Testing repeat of `npm run verify` and `npm run build`, per Techplan.
- Human desktop/mobile rendered acceptance, including truthfulness and interaction comprehension, before the task is marked accepted.
- No Playwright regression was committed: the stable fragment contract is covered by the focused component test and the Build browser inspection found no repeatable browser defect.

## Flagged for Techplan / Testing

None.

## Phase handoff

- Completed: Approved Techplan's static `/` calibration surface and Build-owned checks.
- Artifacts: `.local-agents/works/00-foundations/01-frontend-experience-foundation/3-build/report.md`.
- Human decision: rendered desktop/mobile acceptance is required before final task acceptance; no Build-direction decision is needed now.
- Open / deferred: independent Testing and human acceptance as above.
- Recommended next step: Code Review.
- Session transition: start fresh Code Review for independent inspection of the implementation and this report.
- Context pointers: `2-techplan/techplan.md`; `app/page.tsx`, `app/layout.tsx`, `app/globals.css`, `app/page.test.tsx`; this report.
