# Kencleng — Frontend Tech Stack & Architecture

> Status: Current architecture authority until intentionally replaced by the approved frontend reset
> Purpose: Define Kencleng-specific frontend technology and architecture decisions.
> Boundary: This document does **not** own brand direction or concrete visual-system design.

Generic React/frontend engineering practices live in Harscode. Product-design authority lives under `docs/ui-ux/`.

---

# A. Technology Stack

Current frontend stack:

```text
Next.js App Router
React
TypeScript
Tailwind CSS v4
TanStack Query
Zustand
React Hook Form
Zod
OpenAPI-generated TypeScript types
Vitest
React Testing Library
MSW
```

Playwright is the target browser-automation/rendered-verification capability once repository wiring is present.

The planned frontend reset may deliberately re-evaluate implementation architecture. Until an approved replacement exists, do not infer design authority from the current codebase.

---

# B. Architectural Boundaries

Frontend is not the authority for business truth.

Frontend legitimately owns:

- presentation logic;
- formatting;
- interaction state;
- form lifecycle;
- navigation state;
- UI composition;
- derived presentation values;
- loading/error/success presentation;
- accessibility behavior;
- responsive behavior;
- implementation of approved product-design intent.

Backend/domain authority owns:

- financial correctness;
- eligibility;
- authorization;
- lifecycle transitions;
- verification semantics;
- business validation;
- irreversible domain effects.

If UI appears to require reproducing a business decision the API does not expose, treat it as a contract gap rather than rebuilding backend logic in React.

---

# C. API Contract

`api/openapi.yaml` owns API shape.

Current intended structure:

```text
lib/api/
→ typed API functions

lib/hooks/
→ TanStack Query hooks consuming lib/api/

generated OpenAPI types
→ TypeScript API-shape authority
```

Do not create parallel handwritten API models when OpenAPI already defines them.

---

# D. State Ownership

Before creating state:

```text
Can it be derived?
→ derive

Is API/server state authoritative?
→ server-data owner

Is it meaningful navigation/shareable state?
→ URL when appropriate and non-sensitive

Does it belong to a form lifecycle?
→ form owner

Is it ephemeral UI state?
→ narrowest local owner

Does genuinely client-owned state cross/outlive natural boundaries?
→ approved shared client state
```

Do not mirror API responses into shared client state merely for convenience.

Do not synchronize deterministic projections through effects.

---

# E. Component Ownership

Component placement is semantic-owner-first rather than reusable-first.

Conceptually:

```text
route-specific composition
→ route owner

domain-semantic concept
→ feature/domain owner

genuinely cross-domain semantic concept
→ shared owner

generic UI primitive
→ UI primitive owner
```

Do not extract by line count. Repetition is evidence to evaluate abstraction, not proof of abstraction.

Canonical component governance lives in `frontend/components/README.md`.

---

# F. Shared Component Change Impact

Broad shared components/primitives have potentially large blast radius even when implementation is small.

Before material changes:

```text
classify change
→ discover consumers
→ identify representative risks
→ implement
→ verify component contract
→ verify representative downstream consumers
```

Compilation alone does not prove compatibility.

---

# G. Product Design Authority

Frontend implementation must use:

```text
docs/ui-ux/product-design-principles.md
docs/ui-ux/brand-product-ui-brief.md
docs/ui-ux/patterns.md
docs/ui-ux/asset-governance.md
docs/ui-ux/page-map.md
docs/ui-ux/visual-references/selected-direction/
```

The approved upstream direction is:

```text
Sunlit Editorial
+
Evidence-Led Optimism
```

Selected visual references demonstrate direction and relationships. They are not source code, route contracts, or screenshots to reproduce pixel-for-pixel.

The former legacy prototype/reference generation is superseded and removed from active design authority.

---

# H. Concrete Visual-System Status

There is intentionally no current canonical production specification for exact:

- brand color values;
- font families;
- radius scale;
- spacing tokens;
- shadows/elevation;
- icon family;
- motion values.

Old frontend values — including previous green-primary assumptions and previous font pairing — are not design authority merely because they are present in implementation history.

A future concrete visual-system work item must deliberately derive these decisions from the approved Brand + Product UI Brief.

Do not invent a full design system opportunistically inside unrelated feature Build work.

**Engineering follow-up — intentionally deferred.**

---

# I. Tailwind / Styling Architecture

Current implementation uses Tailwind CSS v4 with CSS-first configuration.

Current token plumbing may use CSS custom properties and `@theme inline`, but existing values are implementation state rather than approved future visual-system truth.

When the frontend reset establishes the concrete production visual system, it should define one coherent token authority rather than maintain parallel theme sources.

Do not use this document to infer exact palette or typography values.

---

# J. Visual Asset Handling

Frontend must detect meaningful visual asset gaps rather than silently substituting generic filler.

Use `docs/ui-ux/asset-governance.md` for:

- utility vs semantic vs expressive vs brand-defining assets;
- photography truthfulness;
- illustration boundaries;
- placeholder governance;
- approval status;
- shared-asset impact analysis.

Brand-defining assets require human approval before canonical use.

---

# K. Forms

Current form approach uses React Hook Form + Zod for client form lifecycle and immediate validation.

Client schemas may reflect documented product constraints, but server validation remains authoritative.

Distinguish field validation from request/business failure and preserve entered information on recoverable failures where practical.

---

# L. User-Controlled Content

User-controlled Markdown/HTML requires an established safe rendering/sanitization path.

Current expected approach may use `react-markdown` + `rehype-sanitize` where appropriate.

Do not manually convert untrusted user content to HTML and pass it unsafely into `dangerouslySetInnerHTML`.

Rendering security is a frontend responsibility even though frontend is not business authority.

---

# M. Authentication and Authorization Presentation

Session routing and role-aware presentation have UX responsibilities, but hiding/disabling frontend controls is not authorization.

Backend authorization remains mandatory and authoritative.

---

# N. Automated Testing

Current unit/component stack:

```text
Vitest
React Testing Library
MSW
```

Use tests for observable behavior such as utilities, hooks, component interactions, validation, meaningful state transitions, accessibility behavior, and API-connected UI through mocks.

Avoid tests whose primary value is asserting implementation structure.

User-controlled content rendering requires hostile-content coverage.

---

# O. Rendered Browser Verification

Component/unit tests do not prove:

- layout;
- spatial hierarchy;
- responsive behavior;
- clipping/overflow;
- fixed/sticky behavior;
- real-browser spatial interaction;
- visual direction fidelity.

Material rendered UI changes require representative real-browser or equivalent layout-capable verification.

Use realistic states, content conditions, and viewports rather than exhaustive screenshot matrices.

---

# P. Browser Automation

Playwright is the selected target capability for browser/rendered verification once repository wiring exists.

Browser-verification capability is distinct from a large permanent E2E suite.

Use permanent E2E coverage only where regression value justifies maintenance cost.

Do not document or claim a Playwright command before the repository actually provides and runs it.

---

# Q. Verification Layers

Frontend confidence is layered:

```text
static / mechanical
→ TypeScript, lint, build

unit / component
→ Vitest + RTL + MSW

rendered / browser
→ real browser / Playwright when available

unresolved product/design intent
→ human/product-design review
```

No layer replaces the others.

---

# R. Mock-First Development

Contract-driven MSW mocks may allow frontend work before live backend availability.

Mocks must follow intended API contracts and must not invent backend product concepts.

Mock verification does not prove real backend integration.

Integration differences must be investigated rather than patched as unreviewed “wiring-only” changes.

---

# S. PWA Scope

Existing v1 decision is app-shell/static-asset caching only unless deliberately superseded.

Do not introduce offline financial writes or donation queues. Financial transactions require current server authority/confirmation.

Stale data must be communicated when freshness materially affects money, status, permissions, or actionable deadlines.

---

# T. Local Verification

Use commands that actually exist in the repository. Current historical baseline includes:

```bash
npm run dev
npm run build
npm run lint
npm run test
npm run verify
```

The upcoming frontend reset must re-document commands if tooling changes.

Do not claim verification that was not run.

---

# U. Deployment

Deployment/hosting remains a separate architectural decision. Do not assume vendor-specific behavior unless deployment strategy establishes it.

---

# V. Frontend Reset Boundary

The approved product-brand direction does not itself select future:

- component architecture;
- folder structure;
- styling mechanics;
- state library;
- data-fetching implementation;
- testing architecture.

Those are engineering decisions to be revalidated during the planned frontend reset.

Design invariants come from `docs/ui-ux/`; engineering implementation must serve them rather than treat the old frontend as a visual baseline.

---

# W. Related Documents

- `docs/ui-ux/README.md`
- `docs/ui-ux/product-design-principles.md`
- `docs/ui-ux/brand-product-ui-brief.md`
- `docs/ui-ux/patterns.md`
- `docs/ui-ux/asset-governance.md`
- `docs/ui-ux/page-map.md`
- `frontend/components/README.md`
- root `AGENTS.md`
- `docs/kencleng-agentic-workflow.md`
