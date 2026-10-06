# AGENTS.md — frontend/

Frontend-specific rules on top of root `AGENTS.md`.

Scope: `frontend/`.

## Baseline posture

This tree is a neutral engineering scaffold. No route, feature, API contract, product slice, or reusable component is considered delivered merely because tooling is present.

Git history is evidence/archive, not active implementation precedent.

## Authority routing

Start from the concern that is active:

- product/business meaning → `../docs/product/product-intent.md`;
- reusable product-design / UX / visual guidance → `../docs/ui-ux/README.md`;
- frontend architecture → `../docs/project/kencleng-frontend-tech-stack.md`;
- reusable component governance → `components/README.md`;
- shared API contract → `../api/` only after an active contract exists;
- generic lifecycle / engineering practice → current Harscode authority.

Release scope and delivery order are not defined at baseline. Do not infer them from historical routes, specs, contracts, code, or Git history.

If product meaning is missing, surface the product decision gap. If product meaning is clear but engineering/contract support is missing, surface the appropriate engineering gap rather than inventing backend behavior in React.

## Write boundaries

A frontend-scoped implementation should not casually modify backend production code. Shared docs/API changes require explicit coordination with the concern they own.

## State ownership

Prefer:

```text
derivable                     → derive
server/API authoritative      → Server Component or TanStack Query as appropriate
navigation/share state        → URL when appropriate and non-sensitive
form lifecycle                → React Hook Form
local interaction             → local React state
genuinely shared client state → Zustand only when justified
```

Do not mirror server data into global client state or create stores by domain convention.

## Components

Use the narrowest truthful semantic owner. Route-specific composition stays local; feature/domain components remain feature/domain scoped; cross-domain shared or generic UI contracts should emerge only from real usage. Follow `components/README.md` when introducing or materially changing broad reusable contracts.

## Design

Use `../docs/ui-ux/README.md` to locate the smallest relevant authority. Illustrative examples and selected visual references do not define release scope or prove a feature exists.

Do not invent consequential product or trust semantics while coding. Material UI requires proportional rendered verification.

## API and mocks

When an active OpenAPI contract exists, generate/use types for the shapes it owns instead of maintaining a parallel handwritten model.

MSW may be used at the network boundary for contract-faithful parallel work. Do not add a production branch that returns mock data and present it as real backend behavior.

## Evidence

Never claim a check ran when it did not. Compilation alone is not proof of visual, behavioral, accessibility, responsive, or semantic correctness.
