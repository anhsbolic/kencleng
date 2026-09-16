# Scaffold and root route — Stage 2 evidence

> Phase/Stage: Exploration / Stage 2 — gap analysis  
> Author: Codex  
> Created: 2026-09-16  
> Model: GPT-5  
> Reasoning: not exposed  
> Session: not exposed  
> Target revision: `a30ee75`  
> Workflow revision: candidate baseline `4199c6db1b26ef1920ba670f222aff0c6d0f9e59`

## Area: clean scaffold and the `/` calibration surface

### Current state

- The repository is past the reboot gate: `docs/project/frontend-reboot-plan.md` §12 marks it **READY FOR DEVELOPMENT** and §10 defines the minimal baseline.
- `app/page.tsx` exports `Home`, which renders only a `main`, `h1` (`Kencleng`), and reboot-baseline paragraph.
- `app/layout.tsx` exports `RootLayout` with `lang="id"`, imports the global stylesheet, and supplies baseline-only metadata. It contains no public shell or navigation.
- `app/globals.css` imports Tailwind v4 and supplies only full-height document and zero body-margin rules.
- The live inventory has no `components/ui`, `components/shared`, `components/features`, `lib`, `mocks`, or browser-test implementation folders. `package.json` retains the expected scaffold/tooling capabilities, including build/lint/Vitest/Playwright commands.

### Requirement

- The feature spec, **Feature surface** and **Goal**, makes `/` the first representative production calibration surface and requires a minimum coherent rendered slice rather than a complete landing page.
- Its **Requirements** require enough actual foundation for public shell/navigation, visual/CTA/content treatment, responsive behavior, and only justified reusable primitives; its **Acceptance criteria** require human judgement at desktop and mobile sizes.
- `frontend-reboot-plan.md` §§5–6 and §10–13 establish that the current shell is intentionally plain, has no product/design-precedent status, and must not restore retired implementation by inertia.

### Gap

The live `/` route is the intentional technical reboot placeholder, not a representative public production slice. It therefore contains none of the foundation evidence the feature expects a human to judge (public shell, canonical typography/surfaces/hierarchy, representative content/CTA treatment, or responsive behavior).

### Sniffing

- **Risk:** This is the first new-generation visual precedent. Any accidental reuse of retired structures/tokens would propagate broadly, despite the current small code surface.
- **Edge cases:** The placeholder has no navigation, content states, assets, or interactive controls; future rendered evidence must cover narrow/mobile width as well as desktop because the acceptance criterion explicitly requires both.
- **Miscontext:** The baseline’s plain `Home` and metadata can look like an incomplete product implementation, but reboot-plan §10 explicitly says they are not product/design precedent.
- **Misleading signals:** Tailwind, React Query, RHF/Zod, MSW, and Playwright are installed capabilities, not evidence that the root route currently uses or needs them.
- **Inconsistency:** None found. The feature’s clean-baseline assumption matches the live source and reboot gate.

### Code anchors

- `app/page.tsx` — `Home`: current `/` implementation and primary later Build/render verification target.
- `app/layout.tsx` — `metadata`, `RootLayout`: document-language and global-shell boundary; currently baseline only.
- `app/globals.css` — global document rules: current styling foundation and future visual-token insertion boundary.
- `package.json` — `scripts`, dependencies/devDependencies: live verification and retained-capability source; availability must not be mistaken for required architecture.
- `docs/project/frontend-reboot-plan.md` — §§5–6, §10–13: authoritative current-baseline and legacy-retirement evidence.
