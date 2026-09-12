# Kencleng — Repository Structure & Local Setup

> File: `docs/project/kencleng-repo-setup.md`
>
> Status: Current project context
>
> Last updated: 2026-09-12
>
> Purpose: Explain the repository layout and local-development topology. Executable files in the repository remain authoritative for exact commands, ports, and dependency versions.

## 1. Repository decisions

### Monorepo

Kencleng keeps backend, frontend, API contract, project documentation, and design references in one repository.

Benefits include:

- one shared OpenAPI contract;
- coordinated project context;
- easy clone/fork for a sandbox/open-source project;
- explicit cross-stack integration without forcing backend/frontend production writes into the same implementation scope.

Monorepo does **not** mean one agent session should freely modify both applications. Root and scoped `AGENTS.md` files define the write boundaries.

### No CI/CD currently

Verification currently runs locally through the repository's actual Makefile/package scripts.

Do not treat this document as the command source of truth. Inspect:

- root `Makefile`;
- `backend/Makefile`;
- `frontend/package.json`;
- other executable configuration relevant to the task.

CI/CD can be introduced later when collaboration/deployment needs justify it.

### Local infrastructure through Compose

The current `docker-compose.yml` runs local infrastructure/proxy services:

```text
PostgreSQL 16
MinIO
MinIO init
Caddy
```

The backend and frontend run natively on the host, not as Compose services.

Current topology:

```text
browser
  ↓
localhost:8080
  ↓
Caddy container
  ├─ /api/* → host backend :8090
  └─ other  → host frontend :3000

host backend
  ├─ PostgreSQL localhost:5435
  └─ MinIO localhost:9087
```

Inspect `docker-compose.yml`, `Caddyfile`, and `.env.example` before relying on those values; they own the executable topology.

## 2. Repository map

```text
kencleng/
├── AGENTS.md
├── api/
│   ├── openapi.yaml
│   └── openapi/                 # split domain source files where applicable
├── backend/
│   ├── AGENTS.md
│   ├── Makefile
│   ├── cmd/server/
│   ├── internal/
│   └── migrations/
├── frontend/
│   ├── AGENTS.md
│   ├── package.json
│   ├── app/
│   ├── components/
│   ├── lib/
│   ├── mocks/
│   └── public/
├── docs/
│   ├── kencleng-agentic-workflow.md
│   ├── project/
│   │   ├── kencleng-backend-tech-stack.md
│   │   ├── kencleng-frontend-tech-stack.md
│   │   ├── kencleng-development-tracker.md
│   │   └── ...
│   ├── spec/
│   │   ├── README.md
│   │   ├── 1-account/
│   │   ├── 2-notification/
│   │   ├── 3-organization/
│   │   ├── 4-campaign/
│   │   ├── 5-donation/
│   │   └── 6-disbursement/
│   └── ui-ux/
│       ├── product-design-principles.md
│       ├── patterns.md
│       ├── design-guidelines.md
│       ├── brand-and-visual-assets.md
│       ├── page-map.md
│       ├── prototype-reference.md
│       └── design-reference-usage.md
├── design-reference/            # frozen/read-only prototype output
├── docker-compose.yml
├── Caddyfile
└── Makefile
```

This map is conceptual, not an exhaustive file inventory.

## 3. Spec layout

Specs use numbered domain directories:

```text
docs/spec/1-account/
docs/spec/2-notification/
docs/spec/3-organization/
docs/spec/4-campaign/
docs/spec/5-donation/
docs/spec/6-disbursement/
```

The number expresses domain/project order. Backend packages remain unnumbered.

Do not assume a literal path mirror such as:

```text
docs/spec/1-account/
=
backend/internal/domain/1-account/
```

Instead, use `docs/spec/README.md` for spec structure and each architecture document for implementation layout.

## 4. UI/UX and design references

`docs/ui-ux/` contains canonical reusable frontend design knowledge.

`design-reference/` is separate because it contains frozen prototype/reference output rather than production code or canonical design policy.

Agents may inspect `design-reference/` but must not modify it or wholesale-copy its implementation. Use:

- `docs/ui-ux/prototype-reference.md`;
- `docs/ui-ux/design-reference-usage.md`.

The old `docs/project/kencleng-design-guidelines.md` is retained only as a superseded pointer. Current visual authority lives in `docs/ui-ux/`.

## 5. Shared API contract

`api/openapi.yaml` is the API-shape authority shared by backend and frontend.

Do not duplicate request/response shape in narrative documentation when the OpenAPI contract already owns it.

Backend and frontend may use different implementation mechanisms around that same contract.

## 6. Local commands

Root convenience commands currently include:

```bash
make verify
make up
make down
make up-podman
make down-podman
```

The current root `make verify` forwards to:

```text
backend:  make verify
frontend: npm run verify
```

Current frontend scripts include:

```bash
npm run dev
npm run build
npm run lint
npm run test
npm run verify
```

Current backend targets are defined by `backend/Makefile`.

These examples are orientation only. If executable configuration changes, update this document when the conceptual topology changes; do not copy command internals into multiple docs merely to keep them synchronized.

## 7. Local ports and services

At the time of this update, executable configuration exposes:

| Service | Host endpoint |
|---|---|
| Caddy entrypoint | `localhost:8080` |
| Frontend native process | `localhost:3000` |
| Backend native process | `localhost:8090` |
| PostgreSQL | `localhost:5435` |
| MinIO API | `localhost:9087` |
| MinIO console | `localhost:9088` |

Use `.env.example` and current configuration when running the stack.

## 8. Known proxy caveat

The current `Caddyfile` uses:

```caddyfile
handle /api/* {
    reverse_proxy host.containers.internal:8090
}
```

This does not strip the `/api` prefix before proxying. Backend tooling has historically worked around this for manual Swagger testing by calling the backend directly.

Treat a proxy fix as an explicit root-scoped infrastructure task. Do not silently modify Caddy from a backend-only or frontend-only feature Build.

## 9. Authority boundaries

This document owns repository/topology rationale only.

Use:

- root/scoped `AGENTS.md` for hard agent rules and write boundaries;
- `docs/kencleng-agentic-workflow.md` for Kencleng project orchestration;
- Harscode for the generic feature lifecycle;
- `docs/spec/` for domain/feature behavior;
- `api/openapi.yaml` for API shape;
- project tech-stack docs for application architecture;
- `docs/ui-ux/` for frontend product/design truth;
- `docs/project/kencleng-development-tracker.md` for current cross-domain status.

If this document disagrees with an executable repository file about a command or topology detail, inspect the executable file and then update this document rather than treating old prose as stronger evidence.

## 10. Evolution

Update this file when the repository structure or local-development topology materially changes.

Do not use it as:

- a feature progress log;
- a duplicate workflow document;
- a catalog of every file;
- a place to preserve obsolete setup snippets that Git history already retains.