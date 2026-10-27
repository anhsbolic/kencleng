# AGENTS.md — backend/

Backend-specific rules on top of root `AGENTS.md`.

Scope: `backend/`.

## Baseline posture

This backend is a neutral engineering substrate. No historical domain, route, migration, security mechanism, or implementation is current merely because it exists in Git history.

## Authority routing

- product/business meaning → `../docs/product/product-intent.md`;
- Engineering Standards index → `../docs/project/standards/README.md`;
- backend architecture → `../docs/project/standards/backend.md`;
- shared identity/authentication policy → `../docs/project/standards/identity-authentication.md`;
- reusable product-design concerns when relevant → `../docs/ui-ux/README.md`;
- API contract → `../api/` only after an active contract exists;
- generic lifecycle / engineering practice → current Harscode authority.

Release scope and delivery order are intentionally open. Do not infer them from historical code/specs/contracts.

## Setup and how to run

- Follow `README.md` and `../docs/project/kencleng-repo-setup.md`. Work from `backend/`; activate Go 1.27.3 with `gvm use "$(cat ../.go-version)"`.
- From `backend/`, `make setup` preserves/creates `.env` from `../.env.example` and installs pinned Air into ignored `bin/`; it does not start PostgreSQL.
- Start the database from root with `make COMPOSE=podman-compose up` or `make up` for Docker Engine/Desktop. The command waits for PostgreSQL readiness.
- `make up` also checks database collation compatibility. Keep the pinned Trixie image family for the existing data; do not run `REFRESH COLLATION VERSION`, reindex, or reset volumes as an automatic setup repair.
- Start with `make gvm-run` from `backend/` to source gvm, select `../.go-version`, then run Air. It uses `GVM_ROOT` or `~/.gvm` and only activates the child runner environment. `make run` (`make dev` alias) uses the already-active Go version. Air rebuilds and restarts after Go/config changes. `go run ./cmd/server` is an alternative without automatic reload.
- Container mode: `make docker-run COMPOSE=podman-compose` from `backend/` (or `make docker-run` for Docker) delegates infra preparation to root and runs only the backend app service. `make docker-stop` with the same provider stops that service. See `../docs/project/local-development-modes.md`; keep Dockerfile Go/Air versions aligned with project pins.
- Container settings bind `0.0.0.0:8090` internally, publish only loopback on the host, and connect to `postgres:5432`. Explicit runtime env works without `.env`; malformed/unreadable dotenv remains fatal. Preserve exported-env precedence.
- `.env` is loaded from the working directory; run backend commands from `backend/`. The host setup uses `APP_HOST=127.0.0.1`, `APP_PORT=8090`, and PostgreSQL at `127.0.0.1:5435`. Preserve existing credentials when updating local configuration.
- After a Human-requested reset to an empty local database, compare `.env` connection settings with the current Compose user/database; reset preserves local env files. Exported environment takes precedence over `.env`. `env -u DATABASE_URL make run` from `backend/` starts one runner using the local file without clearing the parent shell's environment. Do not change loading precedence or recreate historical database roles to hide a stale-config error.
- Direct health check: `curl --fail http://127.0.0.1:8090/healthz`; with frontend running, use `http://localhost:3000/api/healthz` to check the development proxy too.
- Keep backend routes unprefixed by the development gateway's `/api`; Next.js strips that prefix. Do not add permissive CORS just to bypass local proxy setup.
- `make verify` requires its lint/security tools separately; compile or `[no test files]` output is not runtime or behavioral evidence. Do not run migrations/resets as generic setup.

- Optional SeaweedFS is prepared through root `make storage-up` with the same explicit Compose provider. Endpoint host: `http://127.0.0.1:9087`; container: `http://seaweedfs:8333`. No S3 client is implemented yet. Consult `../docs/project/local-object-storage-options.md` before integration; presign against the public endpoint without hostname replacement.

## Structure

```text
backend/
├── cmd/server/main.go
├── internal/
│   ├── domain/          # created only when real delivery needs a product domain
│   └── platform/db/     # neutral database connectivity
└── migrations/          # created by active delivery needs
```

Do not create empty architecture layers merely to make the repository look complete.

## Engineering rules

- Use `net/http` unless demonstrated need justifies otherwise.
- Preserve error chains with `%w`; do not swallow expected failures.
- Parameterize SQL; never interpolate user-controlled values into SQL.
- Do not use `float64` for money.
- Do not log secrets, raw tokens, or unnecessary PII.
- Backend authorization, when introduced, is explicit; frontend gates are not security authority.
- Choose security mechanisms from the active risk/data semantics, not historical implementation inertia.
- Simulation/fixtures must not be presented as real settlement, verification, provenance, or external evidence.

## Testing and evidence

Use unit, integration, race, contract, and security checks proportionally to the active work. Never claim a verification command ran unless it actually ran.

## Write boundaries

A backend-scoped implementation should not casually modify frontend production code. Shared docs/API changes require explicit coordination with the concern they own.
