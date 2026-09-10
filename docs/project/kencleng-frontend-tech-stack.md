# Kencleng — Frontend Tech Stack & Architecture

> Intended path: `docs/project/kencleng-frontend-tech-stack.md`
>
> Status: Draft v2
>
> Purpose: Define Kencleng-specific frontend technology and architectural decisions.
>
> Generic React/frontend engineering practices live in Harscode. This document maps those practices to Kencleng's actual stack and project structure.

---

# A. Technology Stack

Kencleng frontend uses:

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

Primary goals:

* clear ownership of state and data;
* strong alignment with backend contracts;
* accessible and responsive production UI;
* minimal accidental client-side business authority;
* component architecture based on semantic ownership;
* rendered verification for observable UI changes;
* design decisions that remain explicit and reusable.

---

# B. Architectural Boundaries

The frontend is **not the authority for business truth**.

This does not mean the frontend contains no logic.

The frontend legitimately owns:

* presentation logic;
* formatting;
* interaction state;
* form lifecycle;
* navigation state;
* UI composition;
* derived presentation values;
* loading/error/success presentation;
* accessibility behavior;
* responsive behavior;
* product-design implementation.

The backend/domain specification remains authoritative for:

* financial correctness;
* eligibility;
* authorization;
* lifecycle transitions;
* verification semantics;
* business validation;
* irreversible domain effects.

Client-side validation exists to improve UX.

It does not make an operation business-valid.

If the UI appears to require reproducing a business decision that the API does not expose, treat that as a contract/design gap rather than recreating backend logic in React.

---

# C. API Contract

`api/openapi.yaml` is the API contract.

Request/response types are generated with:

```text
openapi-typescript
```

Production structure:

```text
lib/api/
→ pure typed API functions

lib/hooks/
→ TanStack Query hooks consuming lib/api/

generated OpenAPI types
→ API shape authority for TypeScript
```

Do not create parallel handwritten API models when the OpenAPI schema already defines them.

Fetch functions remain intentionally hand-written rather than generating a complete client SDK.

This preserves explicit request behavior while keeping payload types aligned with the API contract.

---

# D. State Ownership

Before creating state, determine whether the value needs an independent owner.

Decision order:

```text
Can it be derived?
        ↓ no
Is external/API data authoritative?
        ↓ no
Is it navigation/shareable state?
        ↓ no
Does it belong to a form lifecycle?
        ↓ no
Is it ephemeral UI state?
        ↓ no
Does genuinely client-owned state need
to cross/outlive natural component boundaries?
        ↓
shared client state
```

---

## 1. Derived values

If a value can be deterministically calculated from existing state:

```text
derive it
```

Do not create another state owner and synchronize it through effects.

Bad:

```tsx
const [isComplete, setIsComplete] = useState(false);

useEffect(() => {
  setIsComplete(progress >= 100);
}, [progress]);
```

Prefer:

```tsx
const isComplete = progress >= 100;
```

Derivation reduces synchronization bugs and competing sources of truth.

---

## 2. Server/API state

Anything whose authority comes from the backend stays with the existing server/data-fetching owner.

For Client Components, this normally means:

```text
TanStack Query
```

Do not mirror API responses into Zustand merely to make access convenient.

Bad:

```text
API
→ TanStack Query
→ copy into Zustand
→ components
```

Prefer:

```text
API
→ TanStack Query
→ consumers
```

When Server Components or server-side lookup legitimately own retrieval, preserve that authority rather than introducing unnecessary client synchronization.

---

## 3. Navigation state

State that should meaningfully support:

* sharing;
* bookmarking;
* refresh;
* browser back/forward;

should generally be URL-owned when appropriate.

Examples may include:

```text
search
filters
pagination
selected public tab
```

Do not put sensitive values into the URL merely for persistence.

Not every toggle belongs in search params.

URL ownership is about navigation semantics, not storage convenience.

---

## 4. Form state

Form-lifecycle state belongs with:

```text
React Hook Form
+
Zod
```

Examples:

* field values;
* touched/dirty state;
* client validation;
* submission form state.

Do not duplicate form fields into Zustand or parallel local state without a concrete lifetime requirement.

Server validation remains authoritative.

---

## 5. Ephemeral UI state

Keep local UI interactions at the narrowest meaningful owner.

Examples:

```text
dialog open
password visible
disclosure expanded
temporary selection
hover/focus-driven presentation
```

Use local React state when the relevant consumers share a natural component boundary.

Do not introduce global state merely to avoid passing a small amount of composition state.

---

## 6. Shared client-owned state

Use Zustand only when state is genuinely client-owned **and** its lifetime or consumers exceed a natural local owner.

Current valid examples may include:

* in-memory authentication/session presentation state;
* other cross-tree client state that cannot naturally live with one component owner.

Do not create a Zustand store for every product domain.

The old rule:

```text
"stores are split per domain"
```

must not be interpreted as:

```text
"every domain gets a store"
```

Instead:

> **When genuinely shared client-owned state needs a store, use Zustand and partition it according to meaningful ownership.**

Zustand must not become a client database mirroring server state.

---

# E. Component Ownership

Production components use **semantic-owner-first placement**.

Canonical rules live in:

```text
frontend/components/README.md
```

High-level mapping:

```text
app/<route>/
→ route-specific composition

components/features/<domain>/
→ domain-semantic components

components/shared/
→ genuinely cross-domain semantic components

components/ui/
→ generic design-system primitives
```

---

## Route-local first when semantics are local

A component used only by one route may remain route-local indefinitely.

Route-local does not mean unfinished.

Do not promote it merely because reusable code is considered cleaner.

---

## Feature/domain components

Use:

```text
components/features/<domain>/
```

when the component represents a stable concept meaningful inside one domain.

Cross-route reuse inside the same domain is not evidence that the component belongs in `shared/`.

---

## Shared components

Use:

```text
components/shared/
```

only when the **same semantic concept and behavioral contract** genuinely spans unrelated domains.

Repeated markup is evidence to investigate.

It is not proof of a shared abstraction.

---

## UI primitives

Use:

```text
components/ui/
```

for generic design-system primitives with no Kencleng business awareness.

They may encode:

* visual tokens;
* generic interaction states;
* accessibility behavior;
* stable visual variants.

They must not encode campaign/donation/organization-specific business semantics.

---

# F. Shared Component Governance

`components/ui/` and `components/shared/` form internal reusable contracts.

Their registry and change policy live in:

```text
frontend/components/README.md
```

Before materially changing one:

```text
classify change
→ discover actual consumers
→ understand blast radius
→ inspect relevant wrappers
→ implement
→ verify component
→ verify representative downstream consumers
```

Do not maintain manually curated consumer lists.

Repository/tooling should discover current consumers when the change happens.

The broader the semantic scope, the greater the expected impact discipline.

A small primitive may have greater change risk than a large route-local component.

---

# G. UI/UX Knowledge Architecture

Frontend implementation consumes several concern-specific design authorities.

```text
product-design-principles.md
→ WHY the experience behaves this way

patterns.md
→ recurring UX behavior

design-guidelines.md
→ canonical visual system

brand-and-visual-assets.md
→ expressive assets and brand identity

prototype-reference.md
→ route-specific prototype authority

design-reference-usage.md
→ how reference output becomes production UI

frontend/components/README.md
→ production component contracts
```

No single visual artifact owns every concern.

Domain/API truth remains authoritative over all UI representation.

---

# H. Design Readiness

Do not assume every feature arrives with complete UI design.

Before implementing a materially new surface, classify it:

```text
READY
PARTIAL
OPEN
```

## READY

Existing UX patterns, design system, assets, and/or approved precedent sufficiently define the experience.

Proceed with implementation.

## PARTIAL

The overall intent is established, but small local decisions remain.

Use senior frontend/product-design judgment within existing principles and patterns.

Surface assumptions when they may create precedent.

## OPEN

Meaningful interaction, information architecture, brand, or product-design decisions remain unresolved.

Perform design exploration before treating production implementation as canonical.

Do not design a materially OPEN product flow accidentally inside JSX.

Exact definitions live in:

```text
docs/ui-ux/product-design-principles.md
```

---

# I. Visual Asset Readiness

A frontend implementation must detect missing visual assets rather than silently substituting generic filler.

When an expressive asset is required:

```text
asset already canonical
→ reuse

current harness can generate suitable candidate
→ brief → generate → review → approve as required

current harness cannot generate
→ produce asset brief + generation prompt
→ human/tool handoff
```

Tool capability determines the production path.

It does not determine whether the design need exists.

Temporary assets must remain explicitly provisional.

Logo, wordmark, primary hero identity, and other brand-defining assets require human approval before becoming canonical.

See:

```text
docs/ui-ux/brand-and-visual-assets.md
```

---

# J. Visual Implementation

Kencleng uses:

```text
Tailwind CSS v4
```

with CSS-first configuration.

Canonical tokens live in:

```text
frontend/app/globals.css
```

and are exposed through:

```css
@theme inline
```

There is no `tailwind.config.js` theme authority.

Prefer existing semantic tokens and production UI components over arbitrary raw values.

A repeated legitimate visual requirement may justify a new system-level token.

One isolated design value usually does not.

---

# K. Action Hierarchy v2

Target visual direction:

```text
Primary
→ filled Kencleng green

Secondary
→ neutral / outlined

Accent amber
→ restrained expressive emphasis,
  not universal Secondary CTA
```

This changes the earlier visual contract where Secondary was amber-filled.

Migration must happen through:

```text
shared Button impact analysis
→ consumer discovery
→ representative rendered verification
```

rather than silently changing the primitive.

---

# L. Prototype Translation

`design-reference/` contains frozen prototype output.

Treat it as:

```text
route-specific design evidence / precedent
```

not production starter code.

Production implementation should preserve:

* hierarchy;
* composition;
* intended states;
* interaction intent;
* visual character;
* responsive intent.

Re-evaluate:

* component boundaries;
* local prototype state;
* API/mock data;
* exact styling values;
* temporary assets.

Current canonical product/design/component authorities override stale prototype implementation details.

Do not modify `design-reference/`.

---

# M. Forms

Forms use:

```text
react-hook-form
+
Zod
```

for client-side form lifecycle and immediate validation feedback.

Validation schemas should reflect relevant documented product constraints.

Do not invent business-validity rules solely in frontend schemas.

Distinguish:

```text
field validation
```

from:

```text
request/business failure returned by API
```

Preserve entered information on recoverable submission failures where practical.

---

# N. User-Controlled Content

User-controlled Markdown or HTML must be rendered using a safe rendering/sanitization path.

Current expected approach:

```text
react-markdown
+
rehype-sanitize
```

Do not manually convert user content to HTML and pass it into:

```text
dangerouslySetInnerHTML
```

without an established safe contract.

Rendering security belongs partly to the frontend even though the frontend is not business authority.

---

# O. Authentication and Authorization Presentation

Coarse session-level route protection and detailed role-aware UI have different responsibilities.

Route-level session handling may redirect unauthenticated users where appropriate.

Role-aware rendering occurs where actual role data is available.

However:

> Hiding or disabling an action in the frontend is not authorization.

Backend authorization remains mandatory and authoritative.

Frontend role gates exist to provide correct UX, not security authority.

---

# P. Automated Testing

Current unit/component stack:

```text
Vitest
React Testing Library
MSW
```

Use tests for:

* hooks and utilities;
* component behavior;
* validation behavior;
* relevant state transitions;
* accessibility-oriented interaction behavior;
* API-connected component behavior through MSW.

Avoid tests that primarily assert internal component structure.

User-controlled Markdown/HTML rendering requires hostile-content coverage, not only happy-path rendering.

---

# Q. Rendered Browser Verification

Automated component tests do not prove:

* layout;
* spatial hierarchy;
* responsive behavior;
* clipping/overflow;
* fixed/sticky interactions;
* real-browser interaction;
* reference alignment.

Therefore any material rendered UI change requires inspection in a real browser or equivalent layout-capable environment.

Use the smallest representative set of:

* states;
* content conditions;
* viewports;

capable of exposing likely failures.

Build may use:

```text
edit
→ render
→ inspect
→ fix
```

as implementation feedback.

When the Harscode workflow uses a separate Testing phase, final rendered verification belongs to Testing and must be independent of Build's claims.

---

# R. Browser Automation Capability

The original v1 decision deferred E2E/browser tooling until a real frontend vertical slice existed.

That condition has now been met.

The project should therefore distinguish:

```text
browser verification capability
```

from:

```text
large permanent E2E suite
```

## Recommended v2 decision

Add **Playwright** as the standard browser automation / rendered-verification capability.

Use it for tasks such as:

* opening production pages in a real browser;
* exercising meaningful interactions;
* checking representative viewport changes;
* inspecting loading/error/success states where practical;
* capturing screenshots when useful;
* reproducing visual/layout regressions.

This does **not** mean every rendered check becomes a committed Playwright E2E test.

Promote a browser check into a permanent automated regression test when the scenario is sufficiently valuable and repeatable.

Good candidates include:

* critical cross-component user flows;
* regressions with meaningful product/security impact;
* repeated responsive failures;
* integration behavior poorly covered by component tests;
* stable end-to-end contracts whose setup cost is justified.

Do not create a comprehensive E2E matrix merely because Playwright exists.

---

# S. Verification Layers

Think of frontend confidence as complementary layers:

```text
static / mechanical
→ TypeScript, lint, build

unit / component
→ Vitest + RTL + MSW

rendered / browser
→ real browser / Playwright

subjective or unresolved design intent
→ human product/design review
```

No layer replaces the others.

A screenshot cannot prove request correctness.

A Vitest suite cannot prove layout correctness.

A browser run cannot decide unresolved brand/product intent.

---

# T. PWA Scope

Current v1 PWA scope remains:

> **app-shell caching only**

Cacheable:

* JS;
* CSS;
* fonts;
* manifest/static shell assets.

Do not introduce offline financial writes or donation queues.

Financial transactions require current confirmation and server authority.

When stale cached data is shown, communicate freshness when stale information could materially affect interpretation.

No custom install prompt is required until product evidence demonstrates a need.

---

# U. Local Verification Commands

Current baseline:

```bash
npm run dev
npm run build
npm run lint
npm run test
npm run verify
```

`npm run verify` currently covers:

```text
lint
+
unit/component tests
```

Browser verification should have an explicit command once Playwright capability is introduced.

Do not document a browser command before the repository actually provides it.

---

# V. Deployment

Deployment/hosting remains an explicit architectural decision to resolve separately.

Possible topology decisions should consider the existing same-origin FE/BE setup and operational requirements.

Do not let frontend implementation assume Vercel-specific behavior unless deployment strategy establishes it.

---

# W. Related Documents

Frontend engineering:

* `frontend/AGENTS.md`
* `frontend/components/README.md`
* Harscode React/frontend best practices

Product design:

* `docs/ui-ux/product-design-principles.md`
* `docs/ui-ux/patterns.md`
* `docs/ui-ux/design-guidelines.md`
* `docs/ui-ux/brand-and-visual-assets.md`
* `docs/ui-ux/page-map.md`
* `docs/ui-ux/prototype-reference.md`
* `docs/ui-ux/design-reference-usage.md`

Contracts:

* `docs/spec/<domain>/...`
* `api/openapi.yaml`

Workflow:

* Harscode `workflow/` — per-feature lifecycle
* `docs/kencleng-agentic-workflow.md` — Kencleng-specific orchestration overlay

Generated task artifacts remain in the Kencleng repository rather than Harscode.
