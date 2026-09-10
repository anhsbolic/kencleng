# AGENTS.md — frontend/

This file adds Kencleng frontend-specific rules on top of root `AGENTS.md`.

Read root `AGENTS.md` first.

Scope:

```text
frontend/
```

Do not modify `backend/` from a frontend-scoped implementation session.

---

## 1. Project Map

```text
frontend/
├── app/                    # Next.js App Router; route-local composition may live here
├── components/
│   ├── ui/                 # generic design-system primitives
│   ├── features/<domain>/  # domain-semantic components
│   └── shared/             # genuinely cross-domain semantic components
├── lib/
│   ├── api/                # typed fetch functions + generated OpenAPI types
│   ├── hooks/              # TanStack Query hooks
│   └── stores/             # only genuinely shared client-owned Zustand state
├── mocks/                  # MSW handlers
└── public/
```

Component placement is **semantic-owner-first**, not reusable-first.

Read:

```text
components/README.md
```

before introducing or materially changing broad reusable components.

---

## 2. Business Authority

Frontend is **not business authority**.

Do not recreate backend decisions for:

* money;
* eligibility;
* permissions;
* verification;
* lifecycle transitions;
* financial validity.

Client validation and presentation derivation are allowed.

Backend/domain truth remains authoritative.

If required product information is missing from the API/spec, flag the contract gap instead of inventing client-side business behavior.

---

## 3. State Ownership

Before creating state:

```text
1. Derivable?
   → derive

2. API/server authoritative?
   → existing server/TanStack Query owner

3. Navigation/share/bookmark/back-forward state?
   → URL when appropriate and non-sensitive

4. Form-lifecycle state?
   → React Hook Form

5. Ephemeral UI interaction?
   → narrowest local owner

6. Genuinely shared client-owned state?
   → Zustand
```

Do not mirror server data into Zustand.

Do not create a Zustand store simply because a domain exists.

Do not synchronize deterministic projections through effects.

Do not move sensitive state into the URL for convenience.

---

## 4. API and Forms

API request/response types come from the generated OpenAPI types in `lib/api/`.

Do not hand-write parallel API types already defined by:

```text
../api/openapi.yaml
```

Forms use:

```text
react-hook-form
+
zod
```

for UX validation.

Server validation remains authoritative.

Distinguish field validation from request/business failures.

---

## 5. Component Ownership

Use the narrowest truthful semantic owner:

```text
route-specific composition
→ app/<route>/

domain concept
→ components/features/<domain>/

genuinely cross-domain semantic concept
→ components/shared/

generic UI primitive
→ components/ui/
```

Do not extract by line count.

Do not promote by usage count alone.

Repeated code is evidence to evaluate shared semantics, not proof of abstraction.

Prefer explicit composition / stable variants over configuration-heavy generic components.

Canonical governance:

```text
components/README.md
```

---

## 6. Shared Component Changes

Before materially changing anything under:

```text
components/ui/
components/shared/
```

you must:

```text
classify the change
→ discover current consumers
→ identify representative risk cases
→ implement
→ verify the component
→ verify representative downstream consumers
```

Do not maintain or trust a stale manual consumer list.

A change that compiles can still be visually, behaviorally, semantically, or accessibly breaking.

UI primitives have potentially high blast radius even when their implementation is small.

---

## 7. Product Design Before UI Implementation

For material UI work, determine design readiness:

```text
READY
→ implement established intent

PARTIAL
→ extend established patterns using product/design judgment

OPEN
→ perform design exploration before canonical implementation
```

Do not silently invent material product or interaction intent while coding.

Read as needed:

```text
../docs/ui-ux/product-design-principles.md
../docs/ui-ux/patterns.md
../docs/ui-ux/design-guidelines.md
../docs/ui-ux/brand-and-visual-assets.md
../docs/ui-ux/page-map.md
```

---

## 8. Visual Assets

Do not silently fill important visual gaps with:

* arbitrary library icons;
* stock-like imagery;
* random gradients;
* generic AI decoration.

Standard library icons are appropriate for standard utility actions.

If an expressive/brand asset is materially required:

```text
canonical asset exists
→ reuse it

current harness can generate
→ brief → generate candidate → review → required approval

current harness cannot generate
→ produce asset brief + ready-to-use generation prompt
```

Do not permanently downgrade the design because the current harness lacks image-generation capability.

Logo/wordmark/brand-defining assets require human approval before becoming canonical.

See:

```text
../docs/ui-ux/brand-and-visual-assets.md
```

---

## 9. Visual System

Read:

```text
../docs/ui-ux/design-guidelines.md
```

before introducing new visual-system behavior.

Kencleng uses Tailwind CSS v4 CSS-first tokens from:

```text
app/globals.css
```

via:

```css
@theme inline
```

There is no `tailwind.config.js` design-token authority.

Prefer canonical tokens and existing primitives.

Do not create feature-specific global tokens or primitive variants merely to avoid local composition.

Target action hierarchy:

```text
Primary   → filled green
Secondary → neutral / outlined
Accent    → restrained warm emphasis
```

---

## 10. Prototype / Design Reference

Before implementing a prototype-derived surface, read:

```text
../docs/ui-ux/prototype-reference.md
../docs/ui-ux/design-reference-usage.md
```

`../design-reference/` is frozen Tier-0 read-only reference output.

Never modify it.

Never wholesale-copy prototype code into production.

Preserve:

* hierarchy;
* composition;
* states;
* interaction intent;
* responsive intent.

Translate through current:

* domain/API truth;
* UX patterns;
* visual system;
* asset system;
* component architecture.

Prototype component boundaries, mock data, local state, exact CSS, and provisional assets are not production authority.

---

## 11. Rendered Verification

If a change materially affects rendered UI or spatial interaction, inspect the result in a real browser or equivalent layout-capable environment.

Automated unit/component tests do not prove:

* layout;
* hierarchy;
* responsive behavior;
* clipping;
* overflow;
* real-browser spatial interaction.

Use representative:

```text
states
+
viewports
+
realistic content conditions
```

rather than exhaustive screenshot matrices.

During Build:

```text
edit → render → inspect → fix
```

is valid implementation feedback.

When Harscode provides a separate Testing phase, final rendered verification belongs to Testing and must independently verify the observable result.

Screenshots are useful evidence when they help; they are not a ritual.

---

## 12. Testing

Current automated baseline:

```text
Vitest
React Testing Library
MSW
```

Tests should verify observable behavior rather than implementation structure.

Any user-controlled Markdown/HTML rendering must include hostile-content sanitization coverage.

Browser verification is separate from the unit/component test layer.

When Playwright/browser automation is available, use it for real rendered verification and promote checks into permanent E2E tests only when the regression value justifies maintenance.

---

## 13. Security Presentation

Never render manually converted user-controlled HTML through unsafe `dangerouslySetInnerHTML`.

Role-aware UI does not provide authorization security.

Backend authorization remains authoritative.

Never expose secrets, tokens, or PII through logs or accidental UI/debug output.

Follow root `AGENTS.md` security rules.

---

## 14. Local Commands

```bash
npm run dev
npm run build
npm run lint
npm run test
npm run verify
```

Do not claim verification that was not actually run.

If rendered UI changed, `npm run verify` alone is not proof of rendered correctness.

---

## 15. Workflow Authority

Generic per-feature lifecycle lives in:

```text
../../harscode-workspace/workflow/
```

Default lifecycle:

```text
Exploration + Techplan
→ Build / patch loop
→ Code Review
→ Testing
→ Pull Request
```

Do not duplicate or redefine that lifecycle here.

Kencleng-specific sequencing, risk tiers, frontend/backend coordination, and domain delivery rules live in:

```text
../docs/kencleng-agentic-workflow.md
```

as a **project orchestration overlay**, not a competing generic feature workflow.

One-off setup/playbook work remains under:

```text
.agents/docs/
```

and is read only when relevant.

---

## 16. Source-of-Truth Routing

Use the authority that owns the question.

```text
business/domain semantics
→ ../docs/spec/<domain>/

API shape
→ ../api/openapi.yaml

frontend architecture
→ ../docs/project/kencleng-frontend-tech-stack.md

product design principles
→ ../docs/ui-ux/product-design-principles.md

UX behavior
→ ../docs/ui-ux/patterns.md

visual system
→ ../docs/ui-ux/design-guidelines.md

brand/assets
→ ../docs/ui-ux/brand-and-visual-assets.md

route/persona inventory
→ ../docs/ui-ux/page-map.md

prototype authority
→ ../docs/ui-ux/prototype-reference.md

component contracts
→ components/README.md
```

If authorities genuinely conflict, surface the contradiction.

Do not silently pick whichever source makes implementation easiest.

---

## 17. Output Style

Default explanatory/process narration:

```text
terse
```

Final deliverables remain complete.

Do not compress:

* risk notes;
* review findings;
* testing/build reports;
* PR descriptions;
* sections whose Harscode workflow contract requires completeness.
