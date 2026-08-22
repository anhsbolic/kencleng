# AGENTS.md — backend/

This file adds **Go-specific** conventions on top of the root
`AGENTS.md`. Read the root file first — the golden rules there
(error handling, money as `decimal`, parameterized SQL, no secrets in
logs, PII encryption pattern, explicit authz checks) apply here without
repetition.

This file is scoped to `backend/`. If you're working in `frontend/`,
see `frontend/AGENTS.md` instead.

## 1. Project layout

```
backend/
├── cmd/server/main.go        # entry point only — wiring, not logic
├── internal/
│   ├── domain/<domain>/      # entity.go, repository.go, service.go per domain
│   ├── transport/http/       # handlers, routing, middleware
│   └── platform/             # db, crypto, auth, ratelimit, storage, scheduler
└── migrations/                # golang-migrate, plain .sql
```

- One package per domain under `internal/domain/`: `account`,
  `organisasi`, `campaign`, `donation`, `disbursement`, `notification`
  — matching the domain boundaries in `docs/spec/domains/`.
- A domain package should not import another domain package directly.
  Cross-domain coordination happens at the `transport/http` handler
  level or through an explicit interface, not by reaching into another
  domain's internals.
- `internal/platform/` is shared infrastructure with no business
  rules of its own — it shouldn't need to know what a "campaign" or
  "donation" is.

## 2. Style conventions

- Standard `net/http` with Go 1.22+ pattern routing (`mux.HandleFunc("POST /donations", ...)`)
  — no third-party router unless a concrete need shows up.
- Repository layer uses `goqu` query builder, never raw string
  concatenation for SQL (see root `AGENTS.md` golden rules).
- Every exported function/type gets a doc comment — this codebase will
  be read by people studying it, not just running it.
- Errors: wrap with `fmt.Errorf("...: %w", err)` to preserve the chain;
  never discard the original error.
- Table-driven tests (`[]struct{ name string; ... }`) are the default
  test shape for anything with more than 2-3 cases.

## 3. Testing conventions

- Unit tests live next to the code (`service_test.go` beside
  `service.go`).
- Anything touching `internal/domain/donation` or
  `internal/domain/disbursement` must include a `-race` run in its
  test file's build tags/CI step — these are Tier 0/1 areas per the
  root workflow doc.
- Integration tests (real Postgres via `testcontainers-go`) live under
  `internal/*/integration_test.go` with a build tag (`//go:build integration`)
  so `go test ./...` (fast, unit-only) and `go test -tags=integration ./...`
  (slower, full) stay separate.

## 4. Local commands

```bash
go run ./cmd/server          # run standalone (needs Postgres/MinIO reachable)
                            # dev (APP_ENV=development) also serves:
                            #   GET /docs        — Swagger UI (loads /openapi.yaml)
                            #   GET /openapi.yaml — spec, server rewritten to :APP_PORT
                            # and writes verification tokens to the dev outbox
                            #   (path logged at startup; default $TMPDIR/kencleng-dev-outbox.log)
go test ./...                # unit tests
go test -race ./...          # concurrency check
make verify                  # full gate — see root docs/kencleng-agentic-workflow.md §7
make migrate-up / migrate-down
```

## 5. Dev tooling (manual testing)

Two read-only dev affordances are wired in `cmd/server/main.go`, gated
on `APP_ENV=development`. They surface no secrets in structured logs and
are not wired in non-dev environments (where `FakeSender` is used and
no docs routes are needed).

- **Swagger UI** — `GET /docs` serves a single-page Swagger UI (CDN
  assets) that loads `GET /openapi.yaml` on the same origin (no CORS).
  The served spec has its `servers.url` rewritten from `/api` to
  `http://localhost:<APP_PORT>` so "Try it out" hits the backend's real
  routes directly. The source spec (`../api/openapi.yaml`) is read from
  disk at runtime and never modified on disk (AGENTS.md §4).
- **Dev email outbox** — when `APP_ENV=development`, the
  `notification.DevSender` writes each simulated email (recipient +
  verification token) to a dev outbox file (default
  `$TMPDIR/kencleng-dev-outbox.log`, overridable via `DEV_OUTBOX_PATH`,
  mode 0600). The path is logged once at startup. This is the dev
  stand-in for an SMTP inbox — tokens stay out of `log.Printf` output
  (the "no tokens in logs" golden rule holds: the outbox file is a
  simulated inbox, not a log stream). In every other environment,
  `FakeSender` is used (token never surfaces).

**Known infra gap (not fixable from a `backend/` session):** the root
`Caddyfile` uses `handle /api/*` (not `handle_path`), so it does NOT
strip the `/api` prefix — `:8080/api/*` 404s against the backend's
routes. Manual Swagger sidesteps this by going direct to `:APP_PORT`.
Fixing the Caddyfile is a root-level session (root AGENTS.md §7
directory boundary), not a backend one.

## 6. One-off operational playbooks

For one-time setup work that isn't a feature (initial project scaffold,
security-tooling config) — check `.agents/docs/README.md` for an index
of available playbooks before improvising the approach yourself. These
are read on-demand, not loaded into every session's context — they cover
things that are relevant once, not on every task.

## 7. Related docs

- `docs/kencleng-agentic-workflow.md` — process, tiering, testing
  stages this file's conventions support.
- `docs/spec/domains/*-invariants.md` — the actual correctness rules
  for each domain package.
- Root `AGENTS.md` — golden rules, fencing, Definition of Done.