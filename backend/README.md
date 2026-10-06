# Kencleng Backend

> Status: clean delivery baseline

Neutral Go backend engineering substrate for Kencleng Pilot #3. No product domain, authentication model, API operation, database migration, or background job is considered delivered at baseline.

Product meaning is owned by `../docs/product/product-intent.md`. Backend architecture is owned by `../docs/project/kencleng-backend-tech-stack.md`.

## Local commands

```bash
go run ./cmd/server
go test ./...
go test -race ./...
make verify
make migrate-up
make migrate-down
```

The baseline server exposes only `GET /healthz` and connects to PostgreSQL. Add product behavior only after the relevant delivery requirement is sufficiently clear.
