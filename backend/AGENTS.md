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
go test ./...                # unit tests
go test -race ./...          # concurrency check
make verify                  # full gate — see root docs/kencleng-agentic-workflow.md §7
make migrate-up / migrate-down
```

## 5. One-off operational playbooks

For one-time setup work that isn't a feature (initial project scaffold,
security-tooling config) — check `.agents/docs/README.md` for an index
of available playbooks before improvising the approach yourself. These
are read on-demand, not loaded into every session's context — they cover
things that are relevant once, not on every task.

## 6. Related docs

- `docs/kencleng-agentic-workflow.md` — process, tiering, testing
  stages this file's conventions support.
- `docs/spec/domains/*-invariants.md` — the actual correctness rules
  for each domain package.
- Root `AGENTS.md` — golden rules, fencing, Definition of Done.