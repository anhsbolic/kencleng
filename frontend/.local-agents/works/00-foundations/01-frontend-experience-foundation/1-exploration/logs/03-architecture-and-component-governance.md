# Architecture and component governance — Stage 2 evidence

> Phase/Stage: Exploration / Stage 2 — gap analysis  
> Author: Codex  
> Created: 2026-09-16  
> Model: GPT-5  
> Reasoning: not exposed  
> Session: not exposed  
> Target revision: `a30ee75`  
> Workflow revision: candidate baseline `4199c6db1b26ef1920ba670f222aff0c6d0f9e59`

## Area: first-generation architecture, ownership, and reusable-contract boundary

### Current state

- `app/page.tsx` and `app/layout.tsx` are currently Server Component-compatible route composition; neither declares browser state/effects or a client boundary.
- The source inventory contains no production component directories beyond `components/README.md`; the living registry explicitly records no new-generation `ui` or `shared` contracts.
- `app/globals.css` is the only global styling file and has Tailwind CSS v4 imported with no product tokens. `postcss.config.mjs` enables the Tailwind v4 PostCSS plugin.
- `tsconfig.json` is strict and supplies the `@/*` alias. No API functions, query hooks, Zustand stores, generated schema, mocks, or client data behavior are present.
- The architecture/tooling documents retain capabilities rather than prescribe their use: Next.js App Router, Tailwind, React Query, RHF/Zod, Zustand, OpenAPI generation, Vitest/RTL/MSW, and on-demand Playwright.

### Requirement

- The feature spec **Requirements** requires semantic-owner-first placement, only genuinely needed global tokens/reusable primitives, and no component-tree restoration or speculative variant/compatibility systems.
- `kencleng-frontend-tech-stack.md` §§3–8 sets frontend’s non-business-authority boundary, state-owner order, Server Component preference, smallest Client Component boundary, and semantic-owner-first component locations. Its §8 prohibits restoring historic buttons/badges/progress/shell/form-wrapper contracts without demonstrated new usage.
- The same architecture document §10 permits only usage-justified global tokens in `app/globals.css`; §§15–21 require observable behavior tests where meaningful, rendered/human acceptance for material UI, responsive/accessibility correctness, and registry updates with every new broad `ui`/`shared` contract.
- `components/README.md` §§1–13 makes the narrowest semantic owner the default, requires cross-domain semantic proof before `shared/`, and requires a same-change living-registry update for any new `ui` or `shared` contract.

### Gap

No product architecture exists yet beyond the clean Server Component scaffold. The task requires a first precedent but the live code provides neither a route-level production composition nor a demonstrated basis for any broad component contract, client state owner, API/data boundary, or token set. Those choices must remain evidence-led rather than inferred from installed dependencies or retired abstractions.

### Sniffing

- **Risk:** A speculative generic shell/button/card/progress abstraction would become the first shared precedent and enlarge later change impact without real consumers. Conversely, a whole-route client boundary would unnecessarily move static composition into browser code.
- **Edge cases:** A public calibration slice may be static/provisional, but any future interactive navigation/menu must preserve keyboard/focus behavior and avoid putting deterministic or sensitive state in effects, stores, or URLs. Global token changes reach every future route.
- **Miscontext:** Empty component directories and missing data layers are intentional clean-start posture, not omissions that need filling for architectural completeness.
- **Misleading signals:** Installed React Query/RHF/Zod/Zustand/MSW/Playwright and Tailwind configuration show retained capability only. They do not establish a requirement for client data fetching, forms, global stores, mocks, browser tests, or a design system.
- **Inconsistency:** None found. The feature’s minimal-primitive direction, architecture’s clean-start rule, and component registry’s empty state agree.

### Code anchors

- `app/page.tsx` — `Home`: current route-local composition boundary and later owner-placement reference.
- `app/layout.tsx` — `RootLayout`: current global document/composition boundary.
- `app/globals.css` — Tailwind import/global rules: only present global token/style boundary.
- `components/README.md` — §§1–13, especially §5 registry and §§7–10 impact/ownership discipline: required documentation and verification routing if a broad contract becomes justified.
- `docs/project/kencleng-frontend-tech-stack.md` — §§3–10, §§15–21: canonical architecture, state, client/server, styling, verification, and accessibility constraints.
- `tsconfig.json`, `postcss.config.mjs`: live TypeScript alias/strictness and CSS pipeline anchors.
