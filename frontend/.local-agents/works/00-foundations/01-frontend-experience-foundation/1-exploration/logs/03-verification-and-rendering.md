# Stage 2 — Verification and rendered-evidence surface

## Current state

The frontend supplies `npm run verify` (lint plus Vitest) and an on-demand
Playwright configuration that starts the Next development server unless a base
URL is supplied. The configured browser is Chromium. The sole committed
browser test is a `/login` smoke test at 1280×800 and 390×844; there is no
browser test for `/`, its public shell, its campaign states, or its mobile
drawer.

Targeted component tests for the current public foundation route passed in
this exploration:

```text
npm run test -- components/features/landing app/(public)/_components/public-shell-client.test.tsx components/features/campaign/campaign-card.test.tsx
5 files passed; 16 tests passed.
```

That run emitted a non-failing Vite future-compatibility warning about native
config loading and the CommonJS/ESM configuration relationship. It did not
run lint, the complete unit suite, a production build, Playwright execution,
or a real rendered `/` inspection. `npx playwright test --list` confirms the
only discovered browser test is the login smoke test.

Existing route/component tests prove selected behavior: shell drawer controls
and auth-modal triggering, hero type-token classes/no fabricated organization
count, campaign loading/populated/empty/error behavior, and card display
boundaries. `app/layout.tsx` declares `lang="en"`, despite this representative
surface’s visible Indonesian copy; no existing test covers document language
or home-page keyboard/focus behavior.

## Requirement

- Feature spec, **Acceptance criteria** and **Verification expectations**:
  `/` must be evaluated at representative desktop and mobile scope; useful
  Playwright/browser automation is requested but does not replace human
  rendered acceptance.
- `frontend/AGENTS.md`, **Rendered iteration and human acceptance** and
  **Testing and Playwright boundary**: material UI requires real rendered
  feedback and human browser acceptance; this task’s explicit browser request
  authorizes relevant Playwright work, but ordinary automated checks do not
  prove hierarchy/responsiveness.
- Harscode `best-practices/react/visual-verification.md`: evidence must name
  surfaces, states, viewports, and interactive conditions; tests/builds alone
  cannot prove visual correctness.
- Harscode `best-practices/react/responsive-layout-robustness.md`: inspect
  real content variation, constrained drawers/actions, and accidental
  horizontal overflow at representative widths.
- Harscode `best-practices/react/accessibility-fundamentals.md`: native
  semantics, accessible icon controls, explicit drawer/modal focus movement,
  non-color-only signals, and keyboard completion require direct coverage.

## Gap

The current checks have no `/` browser coverage and no durable rendered
desktop/mobile inspection evidence. Consequently, the feature’s required
human judgement of the assembled public shell, hero, cards, CTA hierarchy,
mobile drawer, and realistic campaign states cannot be reconstructed from
existing test output. Test coverage is component-local and cannot expose
cross-section spacing/hierarchy, overflow, visual focus, or spatial overlay
problems.

The root document language is English while the home content is Indonesian;
this is a concrete accessibility/localization mismatch to validate and route
under the frontend’s document-language policy, which was not found among the
authorities inspected for this task.

## Sniffing

- **Risk:** A green test suite can mask a visually unusable public entrypoint:
  mobile drawer overlap/focus problems, a clipped CTA/card title, and poor
  hierarchy reach all guest users. The language mismatch can cause assistive
  technologies to select the wrong pronunciation behavior for the public
  content.
- **Edge cases:** Current browser evidence excludes home loading, network
  error/retry, empty campaign list, stale-data notice, long campaign titles,
  mobile open/close/Escape traversal, and horizontal-overflow checks. The
  static normal fixture is not enough to validate responsive growth.
- **Miscontext:** `npm run verify` is described as the normal fast baseline,
  not proof of visual acceptance. The existing login browser test demonstrates
  capability but does not satisfy the explicitly requested `/` coverage.
- **Misleading signals:** Passing 16 targeted tests and token/classname
  assertions can look like completed UI verification even though no rendered
  page was inspected. Playwright being installed can likewise look like `/`
  automation exists; test discovery shows it does not.
- **Inconsistency:** The task explicitly requests useful representative
  Playwright/browser automation and desktop/mobile human review, while the
  committed browser suite covers only `/login`. `lang="en"` conflicts with
  the Indonesian public copy, though the inspected project authorities do not
  state a local document-language rule.

## Code and evidence anchors

- `frontend/package.json` — scripts: actual unit, lint, verify, and browser
  commands.
- `frontend/playwright.config.ts` — browser runtime, automatic dev server,
  failure-only screenshot behavior.
- `frontend/tests/browser/login.smoke.spec.ts` — current desktop/mobile
  browser-test precedent and evidence of missing `/` coverage.
- `frontend/components/features/landing/{hero,highlighted-campaigns,how-it-works}.test.tsx`
  and `frontend/components/features/campaign/campaign-card.test.tsx` —
  component-level behavior presently tested.
- `frontend/app/(public)/_components/public-shell-client.test.tsx` — current
  drawer/auth test coverage; no real viewport/rendered check.
- `frontend/app/layout.tsx` — root `html` language declaration and global
  provider boundary.
- `frontend/AGENTS.md` — §§10–11: rendered/human acceptance and Playwright
  execution rules for later Build/Testing.
