# Kencleng — Frontend Tech Stack & Architecture

> Status: Active v3 — clean-start frontend generation
> Owner: Kencleng project architecture
> Last updated: 2026-09-16
> Purpose: Define Kencleng-specific frontend technology and architectural decisions.
> Reboot context: `docs/project/frontend-reboot-plan.md`

Generic React/frontend engineering practice belongs to Harscode. This document owns Kencleng-specific architecture, selected capabilities, source-of-truth boundaries, and implementation constraints.

## 1. Current posture

Kencleng is beginning a new frontend product implementation after intentionally retiring the previous frontend generation.

The active architecture should therefore be read as:

```text
rules for new implementation
```

not:

```text
instructions to preserve/migrate old components
```

Git history is an archive. Removed routes, components, tokens, local-agent artifacts, and old design/prototype decisions are not implementation precedent unless a current canonical authority explicitly re-establishes the same requirement.

## 2. Selected technology capabilities

The frontend engineering baseline uses:

```text
Next.js App Router
React
TypeScript
Tailwind CSS v4 (CSS-first)
TanStack Query
React Hook Form
Zod
OpenAPI-generated TypeScript types
Vitest
React Testing Library
MSW
Playwright (on-demand browser automation capability)
```

Zustand remains an available dependency/capability for **genuinely shared client-owned state**. Its presence does not mean the project should create global stores or one store per domain.

Phosphor is the canonical utility-icon direction from `docs/ui-ux/design-guidelines.md`. Install/use the appropriate package when real implementation needs icons; the reboot scaffold does not need to pre-create icon abstractions.

PWA/offline behavior is not a prerequisite for the clean reboot baseline. Any future PWA capability must be reintroduced deliberately from a current requirement; do not restore old service-worker behavior merely because the retired frontend had it.

## 3. Authority boundaries

Frontend is not business authority.

Canonical routing:

```text
business/domain behavior
→ docs/spec/<domain-dir>/

API shape
→ api/README.md
→ api/openapi/<domain>.yaml + referenced common components

frontend architecture
→ this document

product/design decision framework
→ docs/ui-ux/product-design-principles.md

selected brand/product UI direction
→ docs/ui-ux/brand-product-ui-brief.md

concrete visual system
→ docs/ui-ux/design-guidelines.md

reusable UX behavior
→ docs/ui-ux/patterns.md

asset lifecycle/truthfulness
→ docs/ui-ux/asset-governance.md

surface/persona intent
→ docs/ui-ux/page-map.md

component ownership/contracts
→ frontend/components/README.md
```

For the same concern, domain/API truth outranks visual representation.

If UI seems to require a business fact or capability that the API/spec does not expose, surface the gap instead of recreating backend logic or inventing product truth in React.

## 4. Data and state ownership

Before creating state, use this order:

```text
derivable
→ derive it

server/API authoritative
→ Server Component ownership or TanStack Query as appropriate

navigation/share/bookmark state
→ URL when appropriate and non-sensitive

form lifecycle
→ React Hook Form

ephemeral UI interaction
→ narrowest local React owner

genuinely shared client-owned state
→ Zustand only when justified
```

Hard rules:

- do not mirror server data into Zustand;
- do not create a store merely because a domain exists;
- do not synchronize deterministic projections through effects;
- do not use URL state for sensitive values merely for persistence;
- do not resurrect old frontend store architecture from Git history as precedent.

## 5. API boundary

Use OpenAPI-generated types for request/response shapes owned by the API contract.

Expected implementation shape when needed:

```text
lib/api/
→ focused typed request functions

lib/hooks/
→ TanStack Query hooks for client-owned server-state consumption

generated schema
→ OpenAPI shape authority in TypeScript
```

These directories do not need to exist before a real task requires them.

Do not create a parallel handwritten model for an OpenAPI-owned shape. Do not generate a large client SDK merely because generation is available; keep network behavior explicit unless a future architectural decision changes that posture.

## 6. Server and Client Component boundary

Prefer Server Components for composition that does not require browser-only state or effects.

Introduce Client Components at the smallest meaningful interactive boundary for concerns such as:

- browser events/state;
- client-side form lifecycle;
- TanStack Query client behavior;
- dialog/drawer interaction;
- browser APIs.

Do not turn route trees into Client Components merely for convenience.

The exact boundary is an implementation decision derived per feature, not copied from retired source files.

## 7. Component ownership

Production components follow semantic-owner-first placement:

```text
route-specific composition
→ app/<route>/...

domain-semantic component
→ components/features/<domain>/...

genuinely cross-domain semantic component
→ components/shared/...

generic visual/interaction primitive
→ components/ui/...
```

These are placement destinations, not required empty folders.

A component can remain route-local indefinitely. Repetition is evidence worth evaluating for abstraction; it is not proof that extraction is required.

Broad `ui`/`shared` contracts use the impact discipline in `frontend/components/README.md`.

## 8. Clean-start component posture

At the reboot baseline there are intentionally **no inherited production `ui`/`shared` contracts**.

Do not recreate historical components such as Button, Badge, ProgressBar, Banner, shell components, or form wrappers simply because they existed previously.

The first new component may establish a broader contract when real usage demonstrates one meaningful boundary. When that happens:

```text
classify semantic owner
→ define the narrowest truthful contract
→ implement
→ verify representative behavior
→ add/update living registry if it belongs to ui/shared
```

## 9. Design readiness

Before material UI implementation classify the task:

```text
READY
PARTIAL
OPEN
```

Definitions and decision authority live in `docs/ui-ux/product-design-principles.md`.

Current global brand direction and concrete visual system are already approved: **Sunlit Editorial / Evidence-Led Optimism**.

Do not reopen that direction inside ordinary engineering work. A feature can still be PARTIAL or OPEN because of feature-specific product/interaction/asset questions.

## 10. Visual implementation

Kencleng uses Tailwind CSS v4 with CSS-first configuration.

New global implementation tokens may live in `app/globals.css` and be exposed through `@theme inline` where useful, but the canonical meaning/value source is `docs/ui-ux/design-guidelines.md`.

The new implementation should derive only tokens needed by real production usage. Do not rebuild the retired token taxonomy for compatibility.

Current visual invariants include:

- Newsreader for approved editorial/display roles;
- Instrument Sans for UI/body/operational roles;
- warm-paper neutral foundation;
- restrained Sun yellow brand energy;
- Berry as selective secondary accent;
- semantic colors independent from brand colors;
- border/spacing-first grouping;
- restrained radius/elevation;
- Phosphor utility-icon baseline;
- provenance/truth presentation grammar;
- funding progress ≠ operational progress ≠ reported outcome;
- public surfaces more expressive, product surfaces more disciplined.

Existing/retired CSS values are not authority.

## 11. Assets

Use `docs/ui-ux/asset-governance.md` for asset classification, truthfulness, approval, and lifecycle.

Do not silently substitute missing meaningful assets with generic filler.

In particular:

- generated/illustrated media must not masquerade as campaign evidence;
- campaign placeholders must clearly remain placeholders;
- Level 4 brand-defining assets require human approval;
- selected-direction references are design evidence, not production assets or screenshot specifications.

## 12. Forms

Use React Hook Form + Zod for client-side form lifecycle and immediate UX validation when a feature needs forms.

Server/domain validation remains authoritative.

Distinguish:

```text
field/client validation
```

from:

```text
request/business failure from the server
```

Preserve user input on recoverable submission failure when practical and consistent with the feature contract.

## 13. Authentication and authorization presentation

Frontend route/session presentation may improve UX, but frontend visibility is never authorization.

Backend authorization remains authoritative.

Do not assume the retired auth provider/store/modal implementation is the required architecture for the new frontend. Derive session/bootstrap/navigation behavior from current account specs, API/security contracts, and the feature being implemented.

## 14. User-controlled content

Any user-controlled Markdown/HTML rendering requires an established safe rendering/sanitization path and hostile-content coverage.

Do not use raw `dangerouslySetInnerHTML` for convenience.

Choose/install the actual rendering packages only when a feature requires this capability; do not retain an unused historical dependency solely as precedent.

## 15. Testing capability

Baseline automated capabilities:

```text
Vitest
React Testing Library
MSW
```

Use them for behavior that benefits from repeatable automation, including utilities, component behavior, validation, meaningful interaction/state transitions, and API-connected component behavior.

Tests should verify observable contracts rather than internal component structure.

Do not create tests merely to preserve historical test counts or retired UI contracts.

## 16. Rendered verification

Material UI work needs rendered feedback because static/unit tests cannot prove hierarchy, layout, responsiveness, clipping, overflow, or interaction comprehension.

During implementation use the smallest representative set of states/content/viewports that can expose likely failures.

Final product acceptance for material UI remains human browser acceptance at proportional desktop/mobile scope.

## 17. Playwright boundary

Playwright is an on-demand real-browser automation capability, not a mandatory phase ritual.

Use or add a browser regression when the behavior is valuable and repeatable enough to justify automation. Examples may include:

- critical cross-component flows;
- interaction/focus behavior poorly represented in jsdom;
- meaningful responsive regressions;
- known browser-specific failures;
- stable end-to-end contracts.

For non-trivial browser automation, the Techplan/verification contract should explain:

```text
why it is worth running
risk if skipped
which phase owns the authoritative run
```

Do not run a broad browser suite in every phase merely because it exists.

## 18. Verification layers

Think of confidence as complementary layers:

```text
static/mechanical
→ TypeScript, lint, build

unit/component
→ Vitest + RTL + MSW where justified

rendered/browser
→ real browser; Playwright when automation is justified

product/design judgement
→ human acceptance
```

Harscode owns generic phase responsibilities and verification ownership between Build, Code Review, and Testing.

## 19. Responsive design

Responsive transformation preserves priority rather than mechanically stacking desktop boxes.

Preserve, in order, the current task, trust/consequence information, essential content, reachable actions, and content dignity. Exact breakpoints and mechanics remain engineering decisions.

## 20. Accessibility

Accessibility behavior is part of production correctness.

Do not rely on color alone for state/meaning. Use semantic HTML, keyboard/focus behavior, labels/names, readable hierarchy, and interaction patterns appropriate to the feature.

Shared accessibility helpers should emerge from real repeated needs; do not preserve old helper APIs automatically.

## 21. Documentation and living contracts

Code owns mechanically discoverable implementation facts. Documentation owns decisions/semantics that cannot safely be reconstructed from code alone.

When a new `components/ui` or `components/shared` contract becomes production precedent, update the living registry in the same change.

Do not document speculative consumers or maintain stale manual consumer lists; discover consumers from current source when changing a broad contract.

## 22. Relationship to frontend reboot

`docs/project/frontend-reboot-plan.md` governs the one-time retirement/reset operation.

After the reboot baseline is frozen, this document becomes the architecture authority for new frontend development and the reboot plan becomes historical project evidence rather than an ongoing implementation instruction.