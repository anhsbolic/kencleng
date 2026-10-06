# Kencleng

Kencleng is a sandbox donation/crowdfunding project and a realistic full-stack learning environment.

This branch is the **Pilot #3 clean delivery baseline candidate**. It preserves approved product/design truth and neutral engineering substrate while intentionally resetting active release scope, delivery sequencing, detailed contracts, feature implementation, and historical orchestration state.

## Current authority

- `docs/product/product-intent.md` — whole-product direction, actors, responsibilities, and trust commitments.
- `docs/ui-ux/README.md` — reusable product-design authority.
- `docs/project/` — repository and stack architecture.
- root/scoped `AGENTS.md` — agent routing, boundaries, and hard rules.

Release goal, delivery slices, detailed requirements, API operations, and implementation state are intentionally **not selected yet**. Git history is evidence/archive, not active authority by inertia.

## Engineering substrate

- Backend: Go + PostgreSQL neutral server scaffold.
- Frontend: Next.js/React/TypeScript/Tailwind scaffold with testing and data/form capabilities available on demand.
- API: OpenAPI validation/bundling tooling with an empty baseline contract.
- Local infrastructure: PostgreSQL, MinIO capability, and Caddy proxy through Compose.

Availability of a technology or service does not mean an active delivery item must use it.

## Local development

```bash
cp .env.example .env
make up-podman
cd backend && go run ./cmd/server
cd frontend && npm install && npm run dev
```

Use `make verify` for a broad repository verification sweep when appropriate. Never claim checks that were not actually run.

This baseline does not authorize Pilot #3 engineering by itself. Product selection and engineering entry remain explicit Human decisions.
