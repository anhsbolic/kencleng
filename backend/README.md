# Kencleng Backend
 
Go backend for Kencleng. See the [root README](../README.md) for what
Kencleng is and how to run the full stack via Docker Compose.
 
This README is for working **inside this folder specifically** — e.g.
running the backend standalone, running tests, running migrations —
without necessarily spinning up the whole stack.
 
## Stack
 
Go (`net/http` stdlib, Go 1.22+ pattern routing), `pgx`, `goqu`,
`golang-migrate`, MinIO client. Full rationale in
`../docs/project/kencleng-backend-tech-stack.md`.
 
## Layout
 
```
backend/
├── cmd/server/main.go     # entry point
├── internal/
│   ├── domain/            # business logic, one package per domain
│   ├── transport/http/    # handlers, routing, middleware
│   └── platform/          # shared infra: db, crypto, auth, storage, scheduler
└── migrations/            # golang-migrate .sql files
```
 
See `AGENTS.md` in this folder for Go-specific coding conventions.
 
## Prerequisites
 
- Go 1.22+
- Postgres and MinIO reachable — either via the root
  `docker-compose up -d` (recommended), or your own local instances
- A `.env` file at the repo root (or in this folder) with the
  variables listed in `../.env.example`
## Running standalone
 
```bash
go run ./cmd/server
```
 
This assumes Postgres/MinIO are already reachable at the addresses in
your `.env` — it doesn't start them. Use the root `docker-compose up`
for that.
 
## Testing
 
```bash
go test ./...                    # unit tests, fast
go test -race ./...              # concurrency check — required for
                                  # anything touching donation/disbursement
go test -tags=integration ./...  # integration tests against real Postgres
```
 
## Migrations
 
```bash
make migrate-up      # apply all pending migrations
make migrate-down     # roll back one migration
```
 
## Verification
 
```bash
make verify
```
 
Runs lint (`staticcheck`, `gosec`), unit tests, race detector, contract
tests, and security scans (`gitleaks`, `govulncheck`) in that order.
See `../docs/kencleng-agentic-workflow.md` §7-8 for what each stage
checks and why.