# Kencleng — Frontend Tech Stack & Architecture

> Status: Active clean-baseline architecture
> Purpose: Define Kencleng-specific frontend technology and architectural constraints without preselecting routes or delivery scope.
> Cross-cutting Engineering Standards: [index](README.md), including [Identity & Authentication](identity-authentication.md).

## Selected capabilities

The frontend baseline provides:

```text
Next.js App Router
React
TypeScript
Tailwind CSS v4
TanStack Query
React Hook Form
Zod
Zustand
OpenAPI-generated TypeScript types
Vitest
React Testing Library
MSW
Playwright
```

Availability is not obligation. A dependency or directory should be used only when the active work has a real need for it.

## Authority boundary

Frontend is not product/business authority.

```text
docs/product/product-intent.md
+ applicable docs/ui-ux authority
        ↓
Human-selected delivery outcome / sufficiently clear requirement
        ↓
contract and implementation decisions at the appropriate boundary
```

If the UI needs product meaning that is not established, surface the product gap. If product meaning is clear but data/contract support is missing, surface the engineering/contract gap. Do not invent backend semantics in React.

## State ownership

Prefer the narrowest truthful owner:

```text
derivable value                 → derive it
server/API authoritative data   → Server Component or TanStack Query as appropriate
navigation/share state          → URL when appropriate and non-sensitive
form lifecycle                  → React Hook Form
local interaction state         → local React state
genuinely shared client state   → Zustand only when justified
```

Do not mirror server data into a global client store or create one store per product domain by convention.

## Components

Organize components by semantic ownership. Route-local composition stays route-local; broader primitives/shared contracts should emerge only from demonstrated use. `frontend/components/README.md` owns reusable component governance.

## API boundary

Development uses the Next.js server rewrite `/api/*` to the Go backend, removing
the `/api` prefix. Browser entry is `http://localhost:3000`; the server-only
`BACKEND_INTERNAL_URL` defaults to `http://127.0.0.1:8090` on the host and is
`http://backend:8090` in app containers. This rewrite is active only during
`next dev`. Production API routing is a separate deployment decision.

Setup/run, pinned runtimes and provider selection are owned by
[repository setup](../development/kencleng-repo-setup.md) and
[development modes](../development/local-development-modes.md). The current local gateway is
Next.js; Caddy is retained only in migration/troubleshooting notes.

When an active OpenAPI contract exists, use generated types for the shapes it owns. Keep request functions focused and do not maintain a parallel handwritten model for the same contract.

MSW may provide contract-faithful network mocks when parallel frontend work is justified, but production code must not contain a hidden mock-data branch presented as real backend behavior.

## Product design

`docs/design/README.md` routes to the approved reusable product-design standards. Illustrative design examples and selected visual references guide expression; they do not establish release scope, route commitments, or completed product state.

Material UI work needs proportional rendered verification. Browser automation is an available capability, not a ritual for every change.

For project-wide risk-based testing, 80% affected testable-scope unit coverage, targeted Vitest/RTL and optional Playwright policy, follow [Testing & Verification Standards (T07)](testing-verification.md). This architecture note does not require running every available test tool.
