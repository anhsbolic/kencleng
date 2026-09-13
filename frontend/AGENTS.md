# AGENTS.md — frontend/

Frontend-specific rules on top of root `AGENTS.md`.

Read root `AGENTS.md` first. This file is intentionally compact: detailed rationale belongs in the canonical frontend architecture/design/component documents it points to.

Scope: `frontend/`.

Do not modify backend production code from a frontend-scoped Build.

## 1. Project map

```text
frontend/
├── app/                    # App Router; route-local composition may live here
├── components/
│   ├── ui/                 # generic design-system primitives
│   ├── features/<domain>/  # domain-semantic components
│   └── shared/             # genuinely cross-domain semantic components
├── lib/
│   ├── api/                # typed API functions + generated OpenAPI types
│   ├── hooks/              # TanStack Query hooks
│   └── stores/             # genuinely shared client-owned Zustand state only
├── mocks/                  # MSW handlers
├── tests/browser/          # optional committed Playwright browser regressions
└── public/
```

Component placement is **semantic-owner-first**, not reusable-first.

Canonical component governance: `components/README.md`.

## 2. Business and API authority

Frontend is not business authority.

Do not recreate backend decisions for money, eligibility, permissions, verification, lifecycle transitions, or financial validity.

Client validation, presentation logic, formatting, interaction state, and derived presentation values are legitimate frontend responsibilities.

If required product information is missing from the API/spec, surface the contract gap instead of inventing business behavior in React.

API request/response shape comes from generated OpenAPI types based on `../api/openapi.yaml`. Do not maintain parallel handwritten API models for shapes the contract already owns.

## 3. State ownership

Before creating state, use this order:

```text
derivable
→ derive it

server/API authoritative
→ existing server/TanStack Query owner

navigation/share/bookmark/back-forward state
→ URL when appropriate and non-sensitive

form lifecycle
→ React Hook Form

ephemeral UI interaction
→ narrowest local React owner

genuinely shared client-owned state
→ Zustand
```

Hard rules:

- do not mirror server data into Zustand;
- do not create a store merely because a domain exists;
- do not synchronize deterministic projections through effects;
- do not put sensitive values in the URL for convenience.

Detailed architecture: `../docs/project/kencleng-frontend-tech-stack.md`.

## 4. Forms and user-controlled content

Forms use React Hook Form + Zod for UX validation. Server validation remains authoritative.

Distinguish field validation from request/business failure.

User-controlled Markdown/HTML must use the established safe rendering/sanitization path. Do not manually convert user content into unsafe `dangerouslySetInnerHTML`.

Any such rendering requires hostile-content test coverage.

## 5. Component ownership and shared-change impact

Use the narrowest truthful owner:

```text
route-specific composition → app/<route>/
domain concept             → components/features/<domain>/
cross-domain semantic      → components/shared/
generic primitive          → components/ui/
```

Do not extract by line count or promote by usage count alone. Repetition is evidence to evaluate an abstraction, not proof that one is needed.

Before materially changing anything in `components/ui/` or `components/shared/`:

```text
classify change
→ discover actual consumers/wrappers
→ identify representative risk cases
→ implement
→ verify component contract
→ verify representative downstream consumers
```

Compilation alone does not prove visual, behavioral, semantic, responsive, or accessibility compatibility.

Read `components/README.md` before introducing or materially changing a broad reusable contract.

## 6. Product design readiness

For material UI work classify design readiness according to `../docs/ui-ux/product-design-principles.md`:

```text
READY   → implement established intent
PARTIAL → resolve small gaps using established principles/patterns
OPEN    → use Exploration to resolve material product/design intent before canonical Build
```

Do not silently invent consequential product or interaction intent while coding.

Read the smallest relevant authority set:

- `../docs/ui-ux/product-design-principles.md` — product/design authority and readiness;
- `../docs/ui-ux/patterns.md` — reusable UX behavior;
- `../docs/ui-ux/design-guidelines.md` — visual system;
- `../docs/ui-ux/brand-and-visual-assets.md` — asset/brand authority;
- `../docs/ui-ux/page-map.md` — route/persona inventory.

## 7. Visual assets

Standard library icons are appropriate for ordinary utility actions.

Do not silently replace a materially important expressive/brand asset need with generic iconography, random gradients, stock-like imagery, or generic AI decoration.

When a required asset is missing, follow `../docs/ui-ux/brand-and-visual-assets.md`:

- reuse canonical asset when one exists;
- generate a candidate when the current harness can do so adequately;
- otherwise produce an asset brief + ready-to-use generation prompt for human/tool handoff;
- keep temporary assets explicitly provisional;
- human approval is required for brand-defining assets.

Tool limitation must not silently become design limitation.

## 8. Visual system

Kencleng uses Tailwind CSS v4 CSS-first tokens from `app/globals.css` via `@theme inline`.

There is no `tailwind.config.js` design-token authority.

Prefer canonical tokens and existing primitives. Do not push feature-specific styling into global tokens/primitive variants merely to avoid local composition.

Target action hierarchy:

```text
Primary   → filled green
Secondary → neutral / outlined
Accent    → restrained warm emphasis
```

## 9. Prototype/reference translation

`../design-reference/` is frozen read-only prototype/reference output.

Before prototype-derived implementation read:

- `../docs/ui-ux/prototype-reference.md`;
- `../docs/ui-ux/design-reference-usage.md`.

Preserve route-specific hierarchy, composition, states, interaction intent, and responsive intent. Translate through current domain/API truth, UX patterns, visual system, asset system, and component architecture.

Do not wholesale-copy prototype code. Prototype component boundaries, mock data, local state, exact CSS, and provisional assets are not production authority.

## 10. Rendered iteration and human acceptance

Material rendered UI or spatial-interaction changes need real rendered feedback; compilation, Vitest, RTL, and MSW do not prove hierarchy, responsiveness, clipping/overflow, or whether an interaction feels understandable in context.

During Build, agent-driven `edit → render → inspect → fix` is valid implementation feedback when the selected client/environment supports it. Use the smallest representative set of realistic states, content conditions, and viewports capable of exposing likely failures.

Agent rendered feedback is **not final product acceptance**.

Before merge/delivery of material frontend UI, a human must manually exercise the rendered result in a real browser at a representative scope appropriate to the change. Human acceptance should focus on:

- visual hierarchy and clarity;
- interaction comprehension;
- responsive usability;
- product/design intent;
- asset appropriateness;
- obvious loading/empty/error/success behavior in context.

Keep the human acceptance scope proportional to the change. Screenshots may support discussion/evidence, but they are not a ritual and do not replace actual interaction where interaction matters.

## 11. Testing and Playwright boundary

Automated baseline:

```text
Vitest
React Testing Library
MSW
```

Tests should verify observable behavior rather than internal structure.

Current local commands:

```bash
npm run dev
npm run build
npm run lint
npm run test
npm run verify
```

`npm run verify` intentionally remains the fast lint + unit/component baseline.

Playwright is installed as a **separate, on-demand browser automation capability**. It is not part of the default Harscode phase contract for Kencleng frontend work.

Available Playwright commands:

```bash
npm run browser:install
npm run test:browser
npm run test:browser:headed
npm run test:browser:ui
```

`npm run browser:install` installs the pinned Chromium browser used by Playwright and is normally needed once per Playwright/browser-version change on a machine.

Do not add, expand, or run Playwright merely because a task reached Build, Code Review, Testing, or PR. A human decides when a browser behavior is valuable enough to automate and explicitly requests the relevant Playwright work (or the task itself is browser-regression/automation maintenance).

Committed tests under `tests/browser/` are automation assets scoped by the behavior they protect: smoke, a feature flow, a known regression, or a deliberate broader journey. Their scope is not derived from a Harscode task folder.

Playwright does not replace required human rendered acceptance.

Do not claim any verification that was not actually run.

## 12. Security presentation

- frontend role-aware rendering does not provide authorization security;
- backend authorization remains authoritative;
- never expose secrets, raw tokens, or PII through logs/debug UI;
- follow root `AGENTS.md` for security/fencing rules.

## 13. Workflow and Codex execution authority

Harscode owns the generic feature lifecycle and generic frontend engineering practice.

When manually invoking a phase, start from the current canonical Harscode `workflow/*-prompt.md` entrypoint. Fill its variables and add only narrow task/project context that is not already owned there. Do not create a parallel frontend copy of the Harscode lifecycle prompts.

`../docs/kencleng-agentic-workflow.md` owns Kencleng-specific sequencing, risk/human authority, integration states, and domain delivery.

For Codex frontend work, `../docs/project/codex-frontend-execution-profile.md` owns project-specific model/reasoning/client/capability routing. That execution profile does not redefine Harscode phases or product/design truth.

Historical feature/task docs may contain numeric section references to older revisions of the Kencleng workflow. Treat the **current named rule/source owner** as authoritative rather than inferring policy from an old section number.

## 14. Source routing

```text
business/domain behavior  → ../docs/spec/<domain-dir>/
API shape                 → ../api/openapi.yaml
frontend architecture     → ../docs/project/kencleng-frontend-tech-stack.md
Codex execution profile   → ../docs/project/codex-frontend-execution-profile.md
product-design authority  → ../docs/ui-ux/product-design-principles.md
UX behavior               → ../docs/ui-ux/patterns.md
visual system             → ../docs/ui-ux/design-guidelines.md
brand/assets              → ../docs/ui-ux/brand-and-visual-assets.md
route/persona inventory   → ../docs/ui-ux/page-map.md
prototype authority       → ../docs/ui-ux/prototype-reference.md
component contracts       → components/README.md
project status            → ../docs/project/kencleng-development-tracker.md
```

If authorities genuinely conflict on the same concern, surface the contradiction.

## 15. Output style

Default narration is terse. Final deliverables remain complete where Harscode requires complete risk notes, review findings, testing/build reports, or PR descriptions.
