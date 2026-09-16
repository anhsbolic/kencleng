# Verification, assets, and baseline health — Stage 2 evidence

> Phase/Stage: Exploration / Stage 2 — gap analysis  
> Author: Codex  
> Created: 2026-09-16  
> Model: GPT-5  
> Reasoning: not exposed  
> Session: not exposed  
> Target revision: `a30ee75`  
> Workflow revision: candidate baseline `4199c6db1b26ef1920ba670f222aff0c6d0f9e59`

## Area: live verification harness, asset inventory, and baseline health

### Current state

- `package.json` exposes `lint`, `test`, `verify`, `build`, and separate Playwright commands. `verify` is lint plus Vitest; no test file currently exists.
- `vitest.config.ts` is a jsdom/RTL-ready configuration with `passWithNoTests: true` and intentionally excludes `tests/browser/**`. `vitest.setup.ts` installs the jest-dom matchers.
- `playwright.config.ts` configures Chromium against a local Next dev server, but the repository has no `tests/browser/` directory or committed browser scenario.
- The live static inventory has no `public/` directory or production imagery/icon/font assets. The only source files are the bootstrap app files and component-governance README.
- Commands executed during this exploration:
  - `npm run verify` — passed. ESLint passed; Vitest exited 0 with no test files. Vite emitted a future config-loader compatibility warning for `vitest.config.ts`; it did not fail the command.
  - `npm run build` — passed when rerun outside the sandbox. The sandboxed attempt failed because Turbopack’s CSS processing could not bind an internal port (`Operation not permitted`), not because of a project compile/type error. The successful build statically prerendered `/` and `/_not-found`.

### Requirement

- Feature spec **Verification expectations** requires proportional static/build checks, focused unit/component checks where meaningful behavior exists, representative desktop/mobile real-browser inspection, and human rendered acceptance. It makes Playwright conditional on a documented repeatable browser-risk justification.
- Feature spec **Requirements** and `asset-governance.md` §§1–12 require Level 3/4 assets to follow classification/truth/approval rules; no generic filler or synthetic documentary imagery may stand in for campaign evidence.
- `kencleng-frontend-tech-stack.md` §§15–20 requires observable-contract tests only where warranted, real rendered feedback for material UI, on-demand rather than ritual Playwright, and accessibility/responsive correctness beyond static checks.

### Gap

The baseline verification harness is healthy but intentionally has no behavior/browser coverage or assets for the representative experience. It cannot provide the task-required desktop/mobile visual evidence or human acceptance until material UI exists. Any expressive/brand asset requirement remains unfulfilled and must be classified/approved rather than filled incidentally.

### Sniffing

- **Risk:** Passing lint/build with zero tests cannot establish visual hierarchy, responsive behavior, accessibility interaction, or truthfulness. A future unclassified asset could create a false campaign-evidence or brand-approval signal.
- **Edge cases:** Browser tests may be unnecessary for a static slice but become justified for stable responsive or keyboard/navigation regressions; mobile and desktop inspection remain mandatory regardless. No-image and asset-loading presentation is currently unrepresented.
- **Miscontext:** The successful scaffold verification is baseline engineering health, not evidence that a production public experience or acceptance criterion is met. A configured Playwright runner is not a browser test.
- **Misleading signals:** `passWithNoTests` yields a green test command while no behavior is protected. The first sandbox build failure looked like a CSS/build fault, but the escalated rerun confirms it was sandbox port policy.
- **Inconsistency:** None found. The no-test/no-asset baseline aligns with reboot-plan §10 and the task’s requirement to introduce only justified implementation/verification.

### Code anchors

- `package.json` — scripts: canonical live commands for future verification planning.
- `vitest.config.ts` — `test.passWithNoTests`, browser exclusion: explains the green baseline with no tests.
- `playwright.config.ts` — `testDir`, `webServer`, Chromium settings: browser automation capability/constraint if later justified.
- `docs/ui-ux/asset-governance.md` — §§1–12: asset class, truth, lifecycle, and approval requirements.
- `docs/spec/0-foundations/features/01-frontend-experience-foundation.md` — Verification expectations: task-specific required evidence boundary.
