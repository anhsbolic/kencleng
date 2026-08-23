# Kencleng — Repo Structure & Setup

> File: `docs/kencleng-repo-setup.md`
> Status: **Agreed** — monorepo, no CI/CD (local docker-compose only
> for local deployment/verification).
> Repo: https://github.com/anhsbolic/kencleng
> Last updated: 2026-08-21 (rev — `docs/ui-ux/` added, replaces
> `docs/wireframes/`; `design-reference/` added as a new top-level
> directory; see §2.2)

## 1. Decisions

- **Monorepo**, not separate backend/frontend repos — reason: single
  source of truth for the API contract (`api/openapi.yaml`) used by
  both sides, easy to clone/fork for open-source purposes, and the
  agentic workflow often needs to touch backend+frontend+spec within
  a single work slice.
- **No CI/CD.** Verification (build, test, lint, security scan) runs
  locally via `make verify` / `docker-compose`, not through an
  automated pipeline like GitHub Actions. This is consistent with the
  "lowest complexity first" principle — CI/CD gets added later if
  there's a concrete need (e.g. if collaboration with other people
  starts).
- **Docker Compose** as the only local orchestration mechanism —
  Postgres, MinIO, backend, frontend all run through a single
  `docker-compose.yml` at the root.

## 2. Directory structure

```
kencleng/
├── docs/
│   ├── project/                  # existing docs (narrative, for humans to read)
│   │   ├── kencleng-erd.md
│   │   ├── kencleng-backend-tech-stack.md
│   │   ├── kencleng-frontend-tech-stack.md
│   │   ├── kencleng-actors-entities.md
│   │   ├── kencleng-business-process-overview.md
│   │   ├── kencleng-phase0-detail.md ... phase3-detail.md
│   │   └── kencleng-roadmap-next-steps.md
│   ├── ui-ux/                     # frontend UX doc set — NEW 2026-08-20, replaces docs/wireframes/
│   │   ├── design-guidelines.md   # visual tokens (moved from docs/project/)
│   │   ├── page-map.md            # per-persona page inventory (moved from docs/project/, evolved)
│   │   ├── patterns.md            # reusable page-shape + state-handling + shared component behavior
│   │   ├── prototype-reference.md # Tier 1/Tier 2 map of which routes have a design-reference/ prototype — NEW 2026-08-21
│   │   └── design-reference-usage.md # how to extract & use design-reference/ exports during FE build — NEW 2026-08-21
│   ├── spec/                     # executable spec, for the agent — see kencleng-agentic-workflow.md
│   │   ├── README.md              # structure & blank templates for each doc type below (cross-domain, stays at spec/ root)
│   │   ├── account/               # domain-first: everything for one domain lives together
│   │   │   ├── invariants.md      # once per domain, stable
│   │   │   ├── threat-model.md    # once per domain, revised on domain-level changes
│   │   │   ├── tasks.md           # once per domain — task list, tier, delivery KPI, parallel/serial grouping (§12 step 5-6)
│   │   │   └── features/          # one file per endpoint/vertical-slice, grows over time — NN-<fitur>.md, NN = task # from tasks.md
│   │   │       ├── 01-register-email-verification.md
│   │   │       └── ...
│   │   ├── notification/
│   │   │   ├── invariants.md
│   │   │   ├── threat-model.md
│   │   │   └── features/
│   │   ├── organization/
│   │   ├── campaign/
│   │   ├── donation/
│   │   └── disbursement/          # same 3-item shape (invariants.md, threat-model.md, features/) for each
│   └── kencleng-agentic-workflow.md  # process reference doc (lives at docs/ root, not project/ or spec/ — this is process, not product spec)
│
├── design-reference/              # Claude Design-exported UI prototype code — NEW 2026-08-21.
│   ├── README.md                  # frozen reference, NOT the real frontend app — see docs/ui-ux/prototype-reference.md
│   └── *.html                     # per-page standalone exports; read-only for agents, see AGENTS.md §3
│
├── api/
│   └── openapi.yaml               # API contract, source of truth — used by backend (hand-written types) & frontend (openapi-typescript)
│
├── backend/
│   ├── go.mod                     # Go module root — repo root does NOT have a go.mod
│   ├── go.sum
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   ├── internal/
│   │   ├── domain/                # flat package per domain — account/, organization/, campaign/, donation/, disbursement/, notification/
│   │   │   └── <domain>/
│   │   │       ├── entity.go
│   │   │       ├── repository.go
│   │   │       └── service.go
│   │   ├── transport/
│   │   │   └── http/              # handlers, routing (Go 1.22+ pattern routing), middleware
│   │   └── platform/              # shared infra: db, crypto, auth, ratelimit, storage (MinIO client), scheduler
│   ├── migrations/                # golang-migrate, run manually via CLI
│   └── Makefile                   # local backend targets: test, verify, migrate — see §4
│
├── frontend/
│   ├── package.json               # Node/Next.js module root — repo root does NOT have a package.json
│   ├── app/                       # Next.js App Router
│   ├── components/
│   ├── lib/
│   │   └── api/                   # types generated by openapi-typescript from ../../api/openapi.yaml
│   ├── public/
│   │   └── manifest.json          # PWA manifest
│   └── package.json scripts       # test, lint — see §4
│
├── AGENTS.md                       # golden rules & DoD for agents — applies across backend+frontend
├── docker-compose.yml              # Postgres, MinIO, Caddy, backend, frontend — LOCAL DEV ONLY (see §3.1)
├── Caddyfile                       # same-origin reverse proxy — see §5.4
├── .env.example                    # env var template — see §5.2
├── Makefile                        # root — calls make in backend/ and frontend/, single verify entry point
└── README.md                       # public-facing: what Kencleng is, how to run it locally
```

### 2.1 `docs/spec/` layout — domain-first **[RESOLVED — 2026-08-05]**

`docs/spec/<domain>/` groups all three document types (`invariants.md`,
`threat-model.md`, `features/`) under one folder per domain, instead of
a type-first split (`spec/domains/`, `spec/threat-model/`,
`spec/features/` each holding files for all 6 domains mixed together).
Reasons:

- **Mirrors `backend/internal/domain/<domain>/`** — same domain name,
  same shape, so navigating spec vs. code for a given domain uses the
  same mental model.
- **`features/` is the fastest-growing folder** (new file per
  endpoint) — keeping it type-first would mean one flat folder mixing
  feature specs from all 6 domains as the project grows; domain-first
  keeps each domain's features isolated and easy to scope a work
  session around.
- `docs/spec/README.md` stays at `spec/` root — it's the shared
  template/reference doc, not owned by any single domain.

Only cost: seeing "all invariants across all domains at once" needs
opening 6 folders instead of 1 — acceptable given there are only 6
domains total, and cross-domain invariant references (§5.1 of
`kencleng-agentic-workflow.md`) just point to
`docs/spec/<domain>/invariants.md#INV-<domain>-NN` instead of the old
path shape.

### 2.2 `docs/ui-ux/` and `design-reference/` **[RESOLVED — 2026-08-20/21]**

`docs/wireframes/` (gray-box HTML/SVG per page) is retired — it went
stale within a month of being drawn, ahead of the 2026-08-20
spec-first pass, and per-page wireframes don't get cheaper to maintain
as the domain count grows. Replaced by `docs/ui-ux/`, a doc set that
validates the *pattern* (reusable page shape + state handling) rather
than every individual page:

- `design-guidelines.md` — visual tokens (moved from `docs/project/`)
- `page-map.md` — per-persona page inventory (moved from
  `docs/project/`, evolved to reference patterns instead of wireframes)
- `patterns.md` — the reusable page-shape/state catalog that replaces
  wireframes
- `prototype-reference.md` — since a handful of representative pages
  *were* prototyped (via Claude Design, exported into
  `design-reference/`) to validate `patterns.md`/`design-guidelines.md`
  actually look right, this doc maps which routes have a near-final
  prototype ("Tier 1") vs which must be derived from `patterns.md`
  alone ("Tier 2"), plus known issues found during prototyping that
  must not be carried into implementation
- `design-reference-usage.md` — how to extract and use the
  `design-reference/` exports (they're self-bootstrapping bundles, not
  plain HTML) without copying architecture that contradicts
  `kencleng-frontend-tech-stack.md`

`design-reference/` itself is a **new top-level directory**, not under
`docs/` — kept separate from `frontend/` (the real app) specifically so
there's no ambiguity about which tree is the actual implementation.
It's frozen, read-only reference output (see `AGENTS.md` §3) — an
agent may read it for visual/structural precedent but must never
write to it or copy it wholesale into `frontend/`.

## 3.1 docker-compose scope — local dev only

`docker-compose.yml` in this repo represents the **local development
topology**, not production. If Kencleng ever gets deployed to a real
server later (not now — there's no demonstrated need for that yet),
that's a separate config (`docker-compose.prod.yml` or whatever setup
fits the hosting platform), not a silent modification of this same
file. This is called out explicitly so there's no mistaken assumption
that this file is "production-ready" as-is.

Why the reverse proxy is still needed in local dev (not just at deploy
time later): the refresh-token cookie uses `SameSite=Strict`, which
only works correctly if the browser sees the backend and frontend as
the **same origin**. Without the proxy, the dev environment would need
looser CORS/cookie settings than what's actually used later — meaning
the security assumptions tested in dev would differ from what actually
applies. The proxy in local compose keeps the dev environment honest
to the same topology as the real target.

## 3. Placement notes (why here, not there)

| Item | Location | Reason |
|---|---|---|
| `api/openapi.yaml` | Root-level `api/`, alongside `backend/` and `frontend/` | Shared contract, not owned by either side, and not narrative documentation (different category from `docs/project/`) |
| `AGENTS.md` | Root | Golden rules apply to all code, not backend/frontend-specific |
| `docker-compose.yml` | Root | Orchestrates all services at once, natural at the repo's top level |
| `kencleng-agentic-workflow.md` | `docs/` (not `docs/project/` or `docs/spec/`) | This is a process/meta document, not a product spec and not product domain documentation |
| `go.mod` | `backend/go.mod` | The repo root shouldn't be a Go module and a Node project at the same time — avoids tooling ambiguity |
| `package.json` | `frontend/package.json` | Same reason — the repo root stays neutral, not owned by a single toolchain |
| `docs/ui-ux/` | `docs/` (sibling to `project/` and `spec/`) | Frontend UX-specific docs (visual tokens, page inventory, page patterns) — grouped separately from `docs/project/`'s general narrative docs since they're referenced together as a set during frontend work, and separately from `docs/spec/` since they're not per-domain |
| `design-reference/` | repo root (sibling to `docs/`, `backend/`, `frontend/`, `api/`) | Frozen Claude Design export — visual/structural precedent for specific pages, not a working codebase. Kept separate from `frontend/` so there's no ambiguity about which tree is the real app. |

## 4. Local verification (in place of CI/CD)

Since there's no CI/CD, `make verify` at the root is the single gate
that must pass before code is considered done (whether written by an
agent or by hand). The root `Makefile` is just a forwarder:

```makefile
verify:
	cd backend && make verify
	cd frontend && npm run verify

up:
	docker-compose up -d

down:
	docker-compose down
```

The detailed contents of `make verify` on each side (lint, unit tests,
race detector, security scan, etc.) follow the stage order agreed on
in `kencleng-agentic-workflow.md` §7-8 — this document doesn't repeat
that, it just maps out where files live.

## 5. Setup decisions — details

### 5.1 Decision summary

| Item | Decision |
|---|---|
| Postgres | version 16 |
| MinIO | single-node single-drive, buckets created automatically via init container (`mc`) |
| Bucket names | `kencleng-public`, `kencleng-private` |
| Port mapping (host) | Caddy `8080` (public entrypoint), Postgres `5435`, MinIO API `9087`, MinIO Console `9088`, backend `8090` (native, proxied — not exposed directly), frontend `3000` (native, proxied — not exposed directly) |
| Reverse proxy | Caddy (simplest config for our needs) |
| Volumes | named volumes, survive `docker-compose down`, only removed with `down -v` |
| App access | `http://localhost:8080` (via Caddy) — not `localhost:3000`/`:8090` directly, so the same-origin assumption is valid from the start |

**Why non-default ports**: Postgres/MinIO's host-side ports are intentionally
moved off their defaults (`5432`/`9000`/`9001`) to avoid colliding with other
local Postgres/MinIO instances running in the same Podman/Docker environment
— container-internal ports are unaffected, only the host-side mapping
changes. **Backend's native port is `8090`, not `8080`** — `8080` is already
taken by Caddy's host-side mapping, and since backend runs as a native host
process (not a container), it shares the host's port namespace with Caddy
and would collide with it at `8080`.

### 5.2 `.env.example`

```env
APP_ENV=development

# Backend HTTP listen port (native process on host — must not collide with
# Caddy's host-mapped port 8080, since Caddy proxies to this port via
# host.containers.internal, not through the same host port)
APP_PORT=8090

# Database
# NOTE: host + host-mapped port (5435), NOT the container-internal DNS name
# "postgres" — the backend runs natively on the host, outside the Podman
# Compose network, so container service names don't resolve here.
DATABASE_URL=postgres://kencleng:kencleng@localhost:5435/kencleng?sslmode=disable

# MinIO / S3-compatible storage
# Same reasoning as DATABASE_URL above — host-mapped port (9087), not the
# container-internal name/port.
MINIO_ENDPOINT=localhost:9087
MINIO_ACCESS_KEY=kencleng
MINIO_SECRET_KEY=kencleng123
MINIO_BUCKET_PUBLIC=kencleng-public
MINIO_BUCKET_PRIVATE=kencleng-private
MINIO_USE_SSL=false

# Auth — JWT (ES256) & session
JWT_PRIVATE_KEY_PATH=./keys/es256-private.pem
JWT_PUBLIC_KEY_PATH=./keys/es256-public.pem
ACCESS_TOKEN_TTL=15m
REFRESH_TOKEN_TTL=720h

# PII encryption-at-rest — generate with: openssl rand -base64 32
ENCRYPTION_KEY=
HMAC_KEY=

# Google OAuth
GOOGLE_CLIENT_ID=
GOOGLE_CLIENT_SECRET=
GOOGLE_REDIRECT_URI=http://localhost:8080/auth/google/callback

# Rate limiting
LOGIN_LOCKOUT_THRESHOLD=5
LOGIN_LOCKOUT_WINDOW=15m
```

> Note: the real `.env` (with actual values) is never committed —
> only `.env.example` goes into the repo, as a reference for which
> variables need to be filled in.

### 5.3 `docker-compose.yml`

```yaml
services:
  postgres:
    image: postgres:16
    environment:
      POSTGRES_USER: kencleng
      POSTGRES_PASSWORD: kencleng
      POSTGRES_DB: kencleng
    volumes:
      - kencleng_pgdata:/var/lib/postgresql/data
    ports:
      - "5432:5432"

  minio:
    image: minio/minio:latest
    command: server /data --console-address ":9001"
    environment:
      MINIO_ROOT_USER: kencleng
      MINIO_ROOT_PASSWORD: kencleng123
    volumes:
      - kencleng_miniodata:/data
    ports:
      - "9000:9000"
      - "9001:9001"

  minio-init:
    image: minio/mc:latest
    depends_on:
      - minio
    entrypoint: >
      /bin/sh -c "
      until mc alias set local http://minio:9000 kencleng kencleng123; do sleep 1; done;
      mc mb -p local/kencleng-public;
      mc mb -p local/kencleng-private;
      mc anonymous set download local/kencleng-public;
      "

  backend:
    build: ./backend
    env_file: .env
    depends_on:
      - postgres
      - minio

  frontend:
    build: ./frontend
    env_file: .env
    depends_on:
      - backend

  caddy:
    image: caddy:2-alpine
    volumes:
      - ./Caddyfile:/etc/caddy/Caddyfile
    ports:
      - "80:80"
    depends_on:
      - backend
      - frontend

volumes:
  kencleng_pgdata:
  kencleng_miniodata:
```

### 5.4 `Caddyfile`

```
localhost:80 {
	handle /api/* {
		reverse_proxy backend:8080
	}
	handle {
		reverse_proxy frontend:3000
	}
}
```

### 5.5 `backend/Makefile` & `frontend/package.json` scripts (skeleton)

The detailed contents of each target follow the stage order agreed on
in `kencleng-agentic-workflow.md` §7-8 (lint → unit → race → property →
contract → security layer A → integration). Initial skeleton:

```makefile
# backend/Makefile
verify: lint test-unit test-race test-contract security
lint:
	staticcheck ./...
	gosec ./...
test-unit:
	go test ./...
test-race:
	go test -race ./...
test-contract:
	go test -tags=contract ./...
security:
	gitleaks detect
	govulncheck ./...
migrate-up:
	migrate -path migrations -database "$$DATABASE_URL" up
migrate-down:
	migrate -path migrations -database "$$DATABASE_URL" down 1
```

```json
// frontend/package.json — "scripts" section
{
  "scripts": {
    "dev": "next dev",
    "build": "next build",
    "lint": "eslint .",
    "test": "vitest run",
    "verify": "npm run lint && npm run test"
  }
}
```

## 6. Not yet decided / discussed separately

- [ ] `docs/spec/README.md` — blank templates for domain-invariant,
      threat-model, feature-spec (already covered separately)
- [ ] Full configuration for security gate layer A per tool (config
      files for `gosec`, `.gitleaks.toml`, etc.) — configuration
      detail, no longer a repo-structure question
- [ ] Generate ES256 keypair (`keys/es256-private.pem` /
      `es256-public.pem`) — one-time manual step during initial
      setup, documented in `README.md`