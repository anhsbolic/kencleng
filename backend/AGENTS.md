# AGENTS.md — backend/

This file adds Go-specific conventions on top of root `AGENTS.md`.

Read the root file first. Its golden rules for error handling, money, parameterized SQL, secrets, PII, authorization, fencing, and authority routing apply here without repetition.

This file is scoped to `backend/`. For frontend work, use `frontend/AGENTS.md`.

## 1. Project layout

```text
backend/
├── cmd/server/main.go
├── internal/
│   ├── domain/<domain>/
│   ├── transport/http/
│   └── platform/
└── migrations/
```

- Domain packages are created under `internal/domain/` as domains are implemented.
- Their specifications live in numbered directories under `docs/spec/`, for example `docs/spec/1-account/`. See `docs/spec/README.md`.
- Do not assume literal path mirroring between spec directories and backend package directories.
- A domain package should not import another domain package directly. Cross-domain coordination happens at the transport/application boundary or through an explicit interface.
- `internal/platform/` is shared infrastructure and should not own business rules.

## 2. Style conventions

- Use standard `net/http` with Go 1.22+ pattern routing unless a concrete need justifies another router.
- Repository queries use the established `goqu` pattern; never concatenate user-controlled SQL.
- Exported functions/types get doc comments.
- Wrap errors with `fmt.Errorf("...: %w", err)`; preserve the original error chain.
- Table-driven tests are the default for behavior with multiple meaningful cases.

## 3. Testing conventions

- Unit tests live next to the code.
- Changes affecting concurrency-sensitive donation/disbursement code require appropriate `go test -race` evidence. Exact scrutiny follows the applicable threat model, root `AGENTS.md` fencing, and Kencleng risk tiering.
- Real-Postgres integration tests use the established integration-test pattern/build tag so fast unit tests and slower integration tests remain separable.
- Do not claim a command ran unless it actually ran.

## 4. Local commands

```bash
go run ./cmd/server
go test ./...
go test -race ./...
make verify
make migrate-up
make migrate-down
```

`make verify` is the backend gate defined by the actual `backend/Makefile`. Do not duplicate its exact target sequence in docs when the Makefile can answer it more reliably.

## 5. Dev tooling

Development-only Swagger UI / OpenAPI serving and the simulated email outbox are implementation affordances, not authority for product/domain behavior. Keep tokens out of structured logs.

Known root proxy/configuration gaps remain root-scoped work; a backend-only Build must not silently modify root/frontend production scope.

## 6. One-off operational playbooks

Use `.agents/docs/README.md` to discover one-off setup/tooling playbooks when relevant. They do not override canonical specs, architecture docs, root/backend `AGENTS.md`, or Harscode workflow.

## 7. Related docs

- `docs/kencleng-agentic-workflow.md` — Kencleng-specific orchestration, risk tiering, human authority, and integration coordination. Harscode owns generic lifecycle/testing-phase mechanics.
- `docs/spec/<domain-dir>/invariants.md` and `threat-model.md` — domain correctness/security rules; see `docs/spec/README.md` for numbered directory naming.
- `docs/project/kencleng-backend-tech-stack.md` — backend architecture.
- Root `AGENTS.md` — repository-wide hard rules, fencing, and authority routing.
