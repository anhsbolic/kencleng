# AGENTS.md — backend/

Backend-specific rules on top of root `AGENTS.md`.

Read root `AGENTS.md` first. Its Product/MVP routing, money, SQL, error, PII, authorization, fencing, workflow, and evidence rules apply here without repetition.

Scope: `backend/`.

## 1. Project layout

```text
backend/
├── cmd/server/main.go        # entry point/wiring only
├── internal/
│   ├── domain/<domain>/      # domain entities/repositories/services
│   ├── transport/http/       # handlers, routing, middleware
│   └── platform/             # db, crypto, auth, ratelimit, storage, scheduler
└── migrations/               # golang-migrate SQL
```

- Domain packages are created under `internal/domain/` as implemented.
- Their delivery specifications live in numbered directories under `docs/spec/`, for example `docs/spec/1-account/`. See `docs/spec/README.md`.
- Numbered spec directories express semantic/domain organization, not current MVP delivery order.
- Do not assume literal path mirroring between numbered spec directories and unnumbered backend packages.
- Domain packages should not reach directly into another domain's internals. Cross-domain coordination uses explicit boundaries/interfaces or the appropriate orchestration layer.
- `internal/platform/` is shared infrastructure; do not put product business rules there.

## 2. Style conventions

- Use standard `net/http` with Go 1.22+ pattern routing unless a demonstrated need justifies another router.
- Repository SQL uses `goqu`; never concatenate/interpolate user-controlled SQL.
- Exported functions/types get doc comments.
- Preserve error chains with `%w`; do not discard underlying errors.
- Table-driven tests are the default when behavior has multiple cases.

## 3. Testing conventions

- Unit tests live next to the code.
- Changes affecting concurrency-sensitive donation/disbursement behavior require appropriate `go test -race` evidence; exact scrutiny follows the applicable threat model, root fencing, and Kencleng risk tier.
- Real-Postgres integration tests use the repository's established `testcontainers-go`/integration-tag pattern where applicable.
- Do not claim any verification command ran unless it actually ran.

Harscode owns when test classes belong in Build vs Testing. Kencleng adds risk-specific evidence requirements; do not redefine a universal phase sequence here.

## 4. Local commands

```bash
go run ./cmd/server
go test ./...
go test -race ./...
make verify
make migrate-up
make migrate-down
```

`make verify` is defined by the actual `backend/Makefile`. Do not duplicate its target sequence into policy docs when the Makefile can answer it mechanically.

## 5. Development tooling

When `APP_ENV=development`, the current server wiring provides development affordances including Swagger/OpenAPI serving and the simulated email outbox.

Treat the dev outbox as simulated inbox content, not permission to log raw verification tokens through structured application logs.

The root Caddy proxy currently has a known `/api` prefix-handling caveat documented in `docs/project/kencleng-repo-setup.md`. Fixing root proxy configuration is a root-scoped infrastructure task, not something a backend-only feature Build should silently modify.

## 6. Source routing

Start from the active product slice, not from historical domain order:

- `docs/product/README.md` + `docs/product/product-overview.md` — durable product/business truth;
- `docs/product/mvp-scope.md` — current MVP inclusion/exclusion boundary;
- `docs/product/mvp-delivery-slices.md` — current slice sequencing and slice boundaries;
- `docs/spec/<domain-dir>/...` — reconciled delivery/domain invariants, threats, and behavior for the active slice; see `docs/spec/README.md`;
- `api/README.md` + `api/openapi/<domain>.yaml` — reconciled API contract shape for the active slice;
- `docs/project/kencleng-backend-tech-stack.md` — backend architecture;
- `docs/project/kencleng-development-tracker.md` — current cross-domain/slice project status;
- `docs/kencleng-agentic-workflow.md` — Kencleng slice sequencing, risk/human/integration orchestration;
- Harscode — generic lifecycle and engineering best-practices;
- root `AGENTS.md` — repository-wide precedence, protected paths, and write boundaries.

Existing backend code, migrations, tests, and old domain specs are implementation/delivery evidence. For an active slice, preserve useful work where it still fits, but do not let existing implementation silently define higher-level product scope or semantics.

If authorities genuinely conflict on the same concern, surface the contradiction instead of silently choosing one.