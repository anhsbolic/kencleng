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
- Engineering Standards index → `../docs/project/standards/README.md`;
- frontend architecture → `../docs/project/standards/frontend.md`;
- shared identity/authentication policy → `../docs/project/standards/identity-authentication.md`;
- reusable component governance → `components/README.md`;
- shared API contract → `../api/` only after an active contract exists;
- generic lifecycle / engineering practice → current Harscode authority.

Release scope and delivery order are not defined at baseline. Do not infer them from historical routes, specs, contracts, code, or Git history.

If product meaning is missing, surface the product decision gap. If product meaning is clear but engineering/contract support is missing, surface the appropriate engineering gap rather than inventing backend behavior in React.

## Setup and how to run

- Follow `README.md` and `../docs/project/kencleng-repo-setup.md`. Use Node.js 24.22.1 via root `.nvmrc` and `nvm use`.
- Work from `frontend/`; `nvm use` discovers the root `.nvmrc`. `make setup` preserves/creates `.env.local` from `.env.example` and installs dependencies with `npm ci`. Avoid incidental lockfile upgrades.
- Start with `make nvm-run` from `frontend/` to source nvm, select the project `.nvmrc`, then run the dev server. It uses `NVM_DIR` or `~/.nvm` and only activates the child runner environment. `make run`, `make dev`, or `npm run dev` uses the already-active Node.js version. The development server binds to `127.0.0.1:3000` and uses Next.js Fast Refresh.
- Container mode: `make docker-run COMPOSE=podman-compose` from `frontend/` (or `make docker-run` for Docker). It runs only the frontend service with Webpack/polling, named dependency/cache volumes, and server-side upstream `http://backend:8090`. Keep the Dockerfile Node version aligned with `.nvmrc`.
- `make check-hot-reload` opt-in verifies both running apps in a real browser and temporarily edits/restores the FE page plus backend health handler. Read `../docs/project/local-development-modes.md`; do not include this source-mutating check in routine parallel UI tests or run it while someone edits its target files.
- Browser entry is `http://localhost:3000`. Browser API calls use relative `/api/...`; the development rewrite removes `/api` and forwards to `BACKEND_INTERNAL_URL` (default `http://127.0.0.1:8090`).
- `BACKEND_INTERNAL_URL` belongs to the Next.js server, without `NEXT_PUBLIC_`. Configure it in `.env.local`; restart `next dev` after changing the proxy destination.
- Reserve `/api/*` for the Go backend in development. Do not create conflicting Next.js API routes or reintroduce Caddy/CORS workarounds without an actual requirement.
- Proxy rewrites are development-only. `next build`/`next start` do not establish a production API gateway; that deployment decision remains separate.
- With backend/database running, `curl --fail http://localhost:3000/api/healthz` checks forwarding. Verify Fast Refresh in a real browser when runtime access permits.

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
