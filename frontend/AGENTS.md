# AGENTS.md — frontend/

This file adds Kencleng frontend-specific rules on top of root `AGENTS.md`. Read root `AGENTS.md` first.

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

Component placement is semantic-owner-first, not reusable-first. Read `components/README.md` before introducing or materially changing broad reusable components.

---

## 2. Business Authority

Frontend is not business authority. Do not recreate backend decisions for money, eligibility, permissions, verification, lifecycle transitions, or financial validity.

Client validation and presentation derivation are allowed. If required product information is missing from the API/spec, flag the contract gap instead of inventing client-side business behavior.

---

## 3. State Ownership

Before creating state:

```text
1. Derivable? → derive
2. API/server authoritative? → existing server/TanStack Query owner
3. Navigation/share/bookmark/back-forward state? → URL when appropriate and non-sensitive
4. Form-lifecycle state? → React Hook Form
5. Ephemeral UI interaction? → narrowest local owner
6. Genuinely shared client-owned state? → Zustand
```

Do not mirror server data into Zustand, create a store merely because a domain exists, synchronize deterministic projections through effects, or move sensitive state into the URL for convenience.

---

## 4. API and Forms

API request/response types come from generated OpenAPI types in `lib/api/`. Do not hand-write parallel API types already defined by `../api/openapi.yaml`.

Forms use `react-hook-form + zod` for UX validation. Server validation remains authoritative. Distinguish field validation from request/business failures.

---

## 5. Component Ownership

Use the narrowest truthful semantic owner:

```text
route-specific composition → app/<route>/
domain concept → components/features/<domain>/
genuinely cross-domain semantic concept → components/shared/
generic UI primitive → components/ui/
```

Do not extract by line count or promote by usage count alone. Repetition is evidence to evaluate shared semantics, not proof of abstraction. Prefer explicit composition/stable variants over configuration-heavy generic components.

Canonical governance: `components/README.md`.

---

## 6. Shared Component Changes

Before materially changing anything under `components/ui/` or `components/shared/`:

```text
classify the change
→ discover current consumers
→ identify representative risk cases
→ implement
→ verify the component
→ verify representative downstream consumers
```

Do not maintain or trust a stale manual consumer list. A change that compiles can still be visually, behaviorally, semantically, or accessibly breaking.

---

## 7. Product Design Before UI Implementation

For material UI work, determine design readiness:

```text
READY → implement established intent
PARTIAL → extend established patterns using product/design judgment
OPEN → perform design exploration before canonical implementation
```

Do not silently invent material product or interaction intent while coding.

Read as needed:

```text
../docs/ui-ux/README.md
../docs/ui-ux/product-design-principles.md
../docs/ui-ux/brand-product-ui-brief.md
../docs/ui-ux/patterns.md
../docs/ui-ux/asset-governance.md
../docs/ui-ux/page-map.md
```

`brand-product-ui-brief.md` owns the approved **Sunlit Editorial / Evidence-Led Optimism** direction. It defines upstream design intent, not exact tokens or component implementation.

The removed legacy `design-guidelines.md`, prototype-reference system, and `docs/design-reference/` exports are not current authority and must not be recreated as precedent from Git history.

Concrete production visual-system decisions such as exact colors, typography, spacing/radius tokens, motion values, and reusable visual primitives remain intentionally OPEN until deliberately derived from the approved brief.

---

## 8. Visual Assets

Do not silently fill important visual gaps with arbitrary library icons, stock-like imagery, random gradients, generic AI decoration, or synthetic documentary-style people.

Standard library icons are appropriate for standard utility actions.

If an expressive/brand asset is materially required:

```text
canonical asset exists → reuse it
current harness can generate → brief → generate candidate → review → required approval
current harness cannot generate → produce asset brief + ready-to-use generation prompt
```

Logo/wordmark/core illustration language and other brand-defining assets require human approval before becoming canonical.

Illustration must never masquerade as real campaign evidence. Photography or imagery must not imply beneficiary identity, distribution, verification, or impact without supporting product truth.

See `../docs/ui-ux/asset-governance.md`.

---

## 9. Visual Direction

Do not derive production styling from the old green design system or legacy prototype exports.

Current upstream visual direction:

- light, warm, editorial foundation;
- restrained clear-bright yellow as brand energy, not semantic proof;
- mature warm neutrals;
- expressive human/editorial public surfaces;
- disciplined authenticated/product surfaces;
- real editorial-documentary photography;
- mature human illustration for explanatory/brand roles;
- information/trust structure as the primary signature;
- Progress as Evidence as a secondary product-language signature.

Do not invent exact hex values, final fonts, or token contracts from this summary. Those decisions require deliberate visual-system derivation from `brand-product-ui-brief.md`.

Avoid defaulting to generic SaaS composition, endless rounded cards, glassmorphism, excessive gradients, fintech-blue/charity-green trust shorthand, badge theatre, or manipulative donation urgency.

---

## 10. Visual References

Selected visual-direction evidence lives under:

```text
../docs/ui-ux/visual-references/selected-direction/
```

Use it to understand thesis, visual character, information hierarchy, public-vs-product expressive intensity, and the approved Evidence Journal concept.

Do not treat visual references as:

- product/domain truth;
- component architecture;
- exact token values;
- route contracts;
- pixel-perfect screenshots to clone.

If a visual reference conflicts with domain/API truth, domain/API truth wins. If it conflicts with an explicitly derived later visual-system authority, the later concern owner wins.

---

## 11. Rendered Verification

If a change materially affects rendered UI or spatial interaction, inspect the result in a real browser or equivalent layout-capable environment.

Automated unit/component tests do not prove layout, hierarchy, responsive behavior, clipping, overflow, or real-browser spatial interaction.

Use representative states + viewports + realistic content conditions rather than exhaustive screenshot matrices.

During Build, `edit → render → inspect → fix` is valid implementation feedback. When Harscode provides a separate Testing phase, final rendered verification belongs there and must independently verify the observable result.

---

## 12. Testing

Current automated baseline:

```text
Vitest
React Testing Library
MSW
```

Tests should verify observable behavior rather than implementation structure. Any user-controlled Markdown/HTML rendering must include hostile-content sanitization coverage.

Browser verification is separate from the unit/component layer. When Playwright/browser automation is actually available, use it where appropriate; do not invent commands or claim checks that did not run.

---

## 13. Security Presentation

Never render manually converted user-controlled HTML through unsafe `dangerouslySetInnerHTML`.

Role-aware UI does not provide authorization security. Backend authorization remains authoritative. Never expose secrets, tokens, or PII through logs or accidental UI/debug output.

---

## 14. Local Commands

```bash
npm run dev
npm run build
npm run lint
npm run test
npm run verify
```

Do not claim verification that was not actually run. If rendered UI changed, `npm run verify` alone is not proof of rendered correctness.

---

## 15. Workflow Authority

Generic per-feature lifecycle lives in the Harscode workflow. Do not duplicate or redefine that lifecycle here.

Kencleng-specific sequencing, risk tiers, frontend/backend coordination, and domain delivery rules live in `../docs/kencleng-agentic-workflow.md` as a project orchestration overlay.

One-off setup/playbook work remains under `.agents/docs/` and is read only when relevant.

---

## 16. Source-of-Truth Routing

```text
business/domain semantics → ../docs/spec/<domain>/
API shape → ../api/openapi.yaml
frontend architecture → ../docs/project/kencleng-frontend-tech-stack.md
product-design principles → ../docs/ui-ux/product-design-principles.md
brand/product UI direction → ../docs/ui-ux/brand-product-ui-brief.md
UX behavior → ../docs/ui-ux/patterns.md
asset governance → ../docs/ui-ux/asset-governance.md
route/persona inventory → ../docs/ui-ux/page-map.md
selected visual-direction evidence → ../docs/ui-ux/visual-references/selected-direction/
component contracts → components/README.md
```

If authorities genuinely conflict, surface the contradiction. Do not silently choose whichever source makes implementation easiest.

---

## 17. Output Style

Default explanatory/process narration: terse. Final deliverables remain complete. Do not compress risk notes, review findings, testing/build reports, PR descriptions, or sections whose Harscode contract requires completeness.
