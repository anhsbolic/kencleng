# Playbook — Backend Scaffold

> File: `backend/.agents/docs/scaffold-backend.md`
> Scope: one-time, run once before Task #1 of any domain (see
> `docs/spec/domains/<domain>/tasks.md`) if `backend/cmd`,
> `backend/internal`, and `backend/migrations` don't exist yet. If they
> already exist, this playbook doesn't apply — stop and say so instead of
> re-running it.
> Tier: this is infrastructure, not a feature — no domain invariant or
> threat model applies. Still goes through a human checkpoint before merge
> (see Step 5) because it sets the wiring pattern every later task builds
> on.

## Step 0 — Verify environment values first

Before writing any Go code, confirm these are true in the actual repo
state (not just this doc) — they were corrected as part of setting up
this playbook and it's easy for one file to drift from another:

- [ ] `docker-compose.yml`: Postgres host port `5435`, MinIO API host port
      `9087`, MinIO Console host port `9088`, Caddy host port `8080`.
- [ ] `.env.example` exists at the repo root with **no trailing space in
      the filename** (`.env.example`, not `.env.example ` — this was a
      real bug found and fixed; if it reappears, something re-created the
      file incorrectly).
- [ ] `.env.example` has `APP_PORT=8090`, `DATABASE_URL` pointing at
      `localhost:5435` (not `postgres:5432` — the backend runs natively,
      outside the Podman Compose network, so container DNS names don't
      resolve here), and `MINIO_ENDPOINT=localhost:9087`.
- [ ] `Caddyfile` proxies `/api/*` to `host.containers.internal:8090` and
      everything else to `host.containers.internal:3000`.

If any of these don't match, stop and flag it — don't silently code
around a stale value, since the DB/MinIO connection steps below depend on
reading the right one.

Then, locally (not committed): copy `.env.example` to `.env` and fill in
the blank secrets (`ENCRYPTION_KEY`, `HMAC_KEY` via
`openssl rand -base64 32`; Google OAuth values can stay blank until Task
#2, Google OAuth Login/Register, is actually being built).

## Step 1 — Folder structure

Create the skeleton exactly as documented in `backend/AGENTS.md` §1 — no
extra top-level folders beyond what's already specified there:

```
backend/
├── cmd/server/main.go
├── internal/
│   ├── domain/            # empty for now — first domain package added in Task #1
│   ├── transport/http/
│   └── platform/
│       ├── db/
│       ├── crypto/
│       ├── auth/
│       ├── ratelimit/
│       ├── storage/
│       └── scheduler/
└── migrations/
```

Every package gets a doc comment at the top of its first file, per
`backend/AGENTS.md` §2 ("every exported function/type gets a doc
comment").

## Step 2 — `main.go` startup order

This is wiring only — no domain logic, no routes beyond a health check.
Follow this order (each step should fail loudly and exit non-zero if it
can't complete, not continue in a half-initialized state):

1. Load `.env` via `godotenv` — fatal if missing when `APP_ENV=development`;
   non-fatal otherwise (env vars are expected to be injected externally in
   other environments).
2. Validate required env vars are present and well-formed. In particular:
   `ENCRYPTION_KEY` and `HMAC_KEY` must not be empty — exit with a clear
   message if they are, rather than starting without encryption.
3. Open a `pgx` connection pool using `DATABASE_URL`, then call `Ping()`
   before proceeding — fail fast if Postgres isn't reachable, don't let a
   bad connection surface later as a confusing query error.
4. Initialize the MinIO client using `MINIO_ENDPOINT` + credentials, and
   verify both `MINIO_BUCKET_PUBLIC` and `MINIO_BUCKET_PRIVATE` are
   reachable.
5. Load the ES256 key pair from `JWT_PRIVATE_KEY_PATH` /
   `JWT_PUBLIC_KEY_PATH`.
6. Wire `internal/platform/db`, `internal/platform/crypto`,
   `internal/platform/auth` using the above. Do **not** wire any
   `internal/domain/*` package yet — none exist until Task #1.
7. Register the router (`net/http`, Go 1.22+ pattern routing) with a
   single route for now: `GET /healthz` returning `200` + a small JSON
   body (e.g. `{"status":"ok"}`) — this is the only way to confirm the
   process is actually listening before any real endpoint exists.
8. Start `http.Server` on `APP_PORT`, with graceful shutdown via
   `signal.NotifyContext` + `Shutdown(ctx)` on `SIGINT`/`SIGTERM`.

## Step 3 — `golang-migrate` wiring

- Confirm `backend/Makefile`'s `migrate-up`/`migrate-down` targets work
  against an empty `migrations/` folder (they should no-op cleanly, not
  error).
- Don't create any actual migration file in this playbook — the first
  migration belongs to Task #1 (Register & Email Verification), since
  that's the first task that needs a table.
- Flag, don't resolve: `docs/kencleng-agentic-workflow.md` §12 notes a
  known unresolved risk — parallel task sessions generating migrations
  concurrently can collide on `golang-migrate`'s sequential numbering.
  This playbook doesn't fix that; it's still open for whoever picks up
  Task #6 or #8 (the parallel-eligible tasks in `tasks.md`).

## Step 4 — Verify

```bash
go run ./cmd/server              # should start cleanly, no panics
curl localhost:8090/healthz      # expect 200 + {"status":"ok"}
go vet ./...
```

`make verify` at this stage will mostly pass trivially (no tests exist
yet) — that's expected, not a gap to fill in this playbook. Don't write
placeholder tests just to have something for `make verify` to run.

## Step 5 — Human checkpoint

Per `docs/kencleng-agentic-workflow.md` §10, this doesn't need the full
Tier 0/1 review (no invariant, no locking strategy, no authz here), but
still needs a look before merge because every later task inherits this
wiring pattern:

- [ ] Startup order actually matches Step 2 (fail-fast on each stage, not
      silently continuing)
- [ ] No domain package was accidentally started in this session — scope
      stayed at platform/transport/cmd only
- [ ] `.env.example` / `docker-compose.yml` / `Caddyfile` values actually
      match what `main.go` reads (Step 0's checklist re-verified after
      the code exists, not just before)

## Related docs

- `backend/AGENTS.md` — project layout, style conventions
- `docs/project/kencleng-backend-tech-stack.md` — full stack rationale
- `docs/project/kencleng-repo-setup.md` §5 — port mapping, `.env.example`
  source of truth