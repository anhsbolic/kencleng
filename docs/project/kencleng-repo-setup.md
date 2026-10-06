# Kencleng — Repository Structure & Local Setup

> Status: Clean delivery baseline candidate
> Purpose: Own repository topology and local-development setup. This document does not choose release scope, delivery order, or product capability priorities.

## Repository posture

Kencleng is a monorepo containing product/design authority, backend and frontend engineering substrates, API-contract tooling, and local infrastructure.

The Pilot #3 delivery baseline intentionally preserves settled engineering substrate while resetting active delivery state. Git history remains evidence and archive; removed delivery specs, contracts, implementations, and workflow artifacts are not active authority merely because they remain in history.

## Repository map

```text
kencleng/
├── docs/
│   ├── product/       # approved upstream product intent
│   ├── ui-ux/         # approved reusable product-design authority
│   └── project/       # repository and stack architecture
├── api/               # OpenAPI tooling; contract starts empty until needed
├── backend/           # neutral Go engineering substrate
├── frontend/          # neutral Next.js engineering substrate
├── AGENTS.md
├── Makefile
├── docker-compose.yml
└── Caddyfile
```

No release goal, slice, delivery spec, API operation, or feature implementation is implied by the existence of these directories.

## Local topology

Local infrastructure is available through Compose:

```text
PostgreSQL 16
MinIO
MinIO initialization
Caddy
```

Backend and frontend run natively on the host. These services are available capabilities, not requirements that every delivery slice must use.

```text
browser
  ↓
localhost:8080
  ↓
Caddy container
  ├─ /api/* → host backend :8090 (the /api prefix is stripped)
  └─ other  → host frontend :3000

host backend
  └─ PostgreSQL localhost:5435

available object storage
  └─ MinIO localhost:9087 (console :9088)
```

Executable configuration remains authoritative for exact ports and commands.

## Commands

Root convenience commands:

```bash
make verify
make up
make down
make up-podman
make down-podman
```

Backend and frontend own their stack-specific commands in `backend/Makefile` and `frontend/package.json`.

## Authority boundary

Use:

- `docs/product/product-intent.md` for whole-product direction and business/product truth;
- `docs/ui-ux/README.md` for reusable product-design authority;
- this directory for engineering architecture and repository topology;
- root/scoped `AGENTS.md` for agent routing and write boundaries;
- Harscode for the active generic workflow when delivery work is authorized.

Release scope, delivery order, detailed delivery requirements, API contracts, and implementation state must be created from current decisions and evidence. Do not reconstruct them from Git history by inertia.
