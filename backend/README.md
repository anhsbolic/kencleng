# Kencleng Backend

> Status: clean delivery baseline

Neutral Go backend engineering substrate for Kencleng Pilot #3. No product domain, authentication model, API operation, database migration, or background job is considered delivered at baseline.

Product meaning is owned by `../docs/product/product-intent.md`. Backend architecture is owned by `../docs/project/kencleng-backend-tech-stack.md`.

## Local commands

Run these commands from `backend/`. Create the local configuration with
`cp .env.example .env`, then fill `DATABASE_URL`, `GOOGLE_CLIENT_ID`, and
`GOOGLE_CLIENT_SECRET`. Use a restricted application DB role. The development
browser origin is `http://localhost:8080`; register the exact Google callback
`http://localhost:8080/api/auth/google/callback`. Production requires an HTTPS
origin. The server reads `.env` from its working directory; keep the file here
and do not commit it.

Backend secret verification distinguishes repository exposure from legitimate
ignored local configuration. `make verify` scans backend Git history, staged
and unstaged tracked changes, and untracked files that Git does not ignore.
Ignored local `.env` / `.env.local` files are therefore outside the repository
leak gate unless they are force-tracked; a force-tracked secret remains visible
to the staged scan. Do not weaken Gitleaks rules or commit local credentials to
make the gate pass.

`C1_AUTH_TEST_DATABASE_URL` is separate disposable integration-test configuration.
Tests do not load `.env` automatically, so export that variable explicitly when
running integration tests. A configured URL alone does not establish restricted
runtime/SCRAM or real-provider verification.

```bash
go run ./cmd/server
go test ./...
go test -race ./...
make verify
make migrate-up
make migrate-down
```

The baseline server exposes only `GET /healthz` and connects to PostgreSQL. Add product behavior only after the relevant delivery requirement is sufficiently clear.
