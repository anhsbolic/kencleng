# Patch Build Report — Manual Testing Tooling + R2 Fix

> Ticket    : 01-register-email-verification
> Stage     : 4-patch (manual-testing enablement + correctness fix)
> Date      : 2026-08-22
> Builder   : GLM 5.2 (opencode)
> Plan      : `./plan-manual-testing.md`
> Feature   : `docs/spec/1-account/features/01-register-email-verification.md`
> Prior patch : `./report-code-review.md` (Tasks A–E)

---

## 1. What shipped

### F1 — Embedded Swagger UI (same-origin, no CORS)

- New `GET /docs` — single-page Swagger UI (assets from CDN
  `swagger-ui-dist@5`) that loads the spec from `/openapi.yaml` on the
  same origin → **no CORS** needed.
- New `GET /openapi.yaml` — reads the bundled `../api/openapi.yaml`
  from disk (dev-only; `go:embed` can't cross the module boundary) and
  rewrites `servers: url: /api` → `http://localhost:8090` in the served
  copy so "Try it out" hits the backend's real `/auth/*` routes
  directly, bypassing the broken Caddy `/api`-no-strip path. The source
  spec file on disk is **never modified**.
- Both routes registered public, outside the rate-limit middleware.

### F2 — Dev-only `DevSender` (token-retrievable, no secrets in logs)

- New `DevSender` writes the simulated email — recipient + token — to
  a dev outbox file (`os.TempDir()/kencleng-dev-outbox.log`, overridable
  via `DEV_OUTBOX_PATH`), mode 0600, append under a mutex. Gated on
  `APP_ENV=development`; `FakeSender` remains the default in every
  other environment.
- Tokens stay out of structured `log.Printf` output — the outbox file
  is a simulated inbox (the dev stand-in for an SMTP mailbox), not a
  log stream, so the "no tokens in logs" golden rule is honored.
- The startup log prints the outbox path once so the developer knows
  where to find verification tokens.

### F3 — R2 branch delivers the token (correctness fix)

- `Service.Register` R2 branch (unverified existing) now captures the
  plaintext token returned by `issueNewVerificationToken` and calls
  `sendVerification(ctx, email, plainToken)` — the verification email
  carrying the new token, identical to R13 resend. Previously it
  discarded the plaintext and sent a token-less
  `NudgeResendVerification` nudge, so the new token was issued but
  never delivered (the user could not verify).
- `issueNewVerificationToken` doc comment updated to reflect that both
  R2 and R13 send the verification email with the new token.
- Removed the now-dead `NudgeResendVerification` constant (only
  consumer was the buggy R2 branch); R3/R4 nudges unchanged.
- `TestRegister_UnverifiedExisting_ResendFlow` updated to assert the
  verification email is sent to the recipient + no nudge (was asserting
  the wrong nudge behavior, hiding the bug).

## 2. Files changed

| File | Change |
|---|---|
| `internal/transport/http/swagger.go` (new) | `SwaggerHandler` (CDN Swagger UI page) + `OpenAPIHandler` (serve spec with server URL rewritten to `:8090`) |
| `internal/platform/notification/dev_sender.go` (new) | `DevSender` — appends recipient+token to a dev outbox file (0600, mutex); implements `Sender` |
| `internal/platform/notification/sender.go` | Removed dead `NudgeResendVerification` constant |
| `internal/domain/account/service.go` | R2 branch: capture `plainToken` + `sendVerification` (was discard + `sendNudge`); `issueNewVerificationToken` doc comment |
| `internal/domain/account/service_test.go` | `TestRegister_UnverifiedExisting_ResendFlow`: assert verification email sent + no nudge |
| `cmd/server/main.go` | Wired `GET /docs` + `GET /openapi.yaml`; `newEmailSender` picks `DevSender` in dev / `FakeSender` otherwise; logs outbox path |

## 3. Verification

| Gate | Result |
|---|---|
| `go build ./...` | ✅ OK |
| `go vet ./...` | ✅ clean |
| `go test -count=1 ./...` | ✅ all pass |
| `go test -race -count=1 ./internal/domain/account/...` | ✅ race-clean (54s) |
| `go test -count=1 ./internal/platform/notification/... ./internal/transport/http/...` | ✅ pass |
| End-to-end (curl, live server on :8090) | ✅ see below |

End-to-end flow verified against live Postgres + MinIO (podman):

| Flow | Result |
|---|---|
| R1 register (fresh email) → `202` + token in outbox | ✅ |
| R8 verify (valid token) → `200 "Email verified."` | ✅ |
| R10 re-verify (used token) → `404` | ✅ |
| R13 resend (unverified) → `202` + new token in outbox | ✅ |
| R14 resend (verified) → `202` + no new outbox line (no-op) | ✅ |
| **R2 register (unverified existing)** → `202` + **new token in outbox** (was a token-less nudge) | ✅ fixed |
| R2-issued token → verify → `200` | ✅ |
| R3 register (verified existing) → `202` + `password_reset` nudge (unchanged) | ✅ |
| Breach check (common password `supersecret123`) → `422` | ✅ HIBP reachable |

`make verify` not run — it still fails on the pre-existing gosec findings
in `platform/auth/keys.go` (G304) and `cmd/server/main.go` (G112), not
from this change. `staticcheck`/`gosec` on changed packages not run
locally (tools not installed); `go vet` clean.

## 4. Fence compliance

No Tier 0 fenced path modified (AGENTS.md §3): `platform/crypto/`,
`platform/auth/`, `domain/donation/ledger.go`, `domain/disbursement/`
all untouched. No `docs/spec/*` file edited (AGENTS.md §4); the bundled
`api/openapi.yaml` is read at runtime, never modified on disk.

## 5. How to run (manual Swagger testing)

```bash
# 1. Infra (already up via podman): Postgres :5435, MinIO :9087
# 2. Apply migrations (if not already)
set -a; . .env; set +a && make migrate-up
# 3. Run the server (reads .env; boots on :8090)
go run ./cmd/server
# 4. Open Swagger UI in a browser:
#    http://localhost:8090/docs
# 5. Verification tokens are written to the dev outbox (path logged
#    at startup, default /tmp/kencleng-dev-outbox.log). Grab the token
#    from there to test POST /auth/verify-email.
```

---

## Risk note

- **Assumptions made:**
  - `OpenAPIHandler` reads the spec from `../api/openapi.yaml` relative
    to the server's working directory (`backend/`). This is dev-only
    and breaks if the cwd changes — acceptable for a dev docs endpoint,
    not a production path. Verified by `GET /openapi.yaml` returning
    the spec with the server URL rewritten.
  - The server-URL rewrite uses `bytes.Replace(raw, "url: /api", "url:
    http://localhost:8090", 1)` on the served copy. Safe because the
    bundled spec has exactly one `url: /api` (verified by `grep`). The
    source file on disk is untouched.
  - `DevSender` writes recipient + token to the outbox file (mode 0600).
    This is the faithful dev simulation of an inbox — a real email lands
    in a mailbox the recipient can read. It is gated on
    `APP_ENV=development`; in every other environment `FakeSender` is
    used (token never surfaces). Verified by `newEmailSender` branch +
    the startup log line.
  - F3's `sendNudge` → `sendVerification` change does not alter R7
    anti-enumeration timing: both are post-commit calls, and
    `FakeSender`/`DevSender` do no network I/O (log/file only), so the
    DB-work shape (BeginTx + UPDATE/INSERT + Commit) is what determines
    branch wall-clock time — unchanged.

- **Edge cases intentionally NOT handled (and why):**
  - The Caddy `/api` prefix is not stripped (`handle` not `handle_path`)
    → `:8080/api/*` 404s against the backend. This is a pre-existing
    root-level infra issue, outside the `backend/` session scope
    (AGENTS.md §7). Manual Swagger sidesteps it by pointing at `:8090`.
  - `DevSender` has no automated test — it is dev-only tooling with
    trivial append-to-file logic (mutex + `os.OpenFile` append, 0600).
    The end-to-end curl flow (token appears in outbox, verify succeeds)
    is the operational proof.
  - `staticcheck`/`gosec` not run on changed packages (not installed
    locally). `go vet` is clean. `make verify` should be run before
    merge (it also fails on the pre-existing G304/G112, not from this
    change).

- **Concurrency assumptions:**
  - `DevSender.append` holds a process-local mutex; the outbox file is
    opened with `O_APPEND`, which is atomic for small writes on local
    filesystems. `-race` clean across the account suite (which exercises
    R2 + the concurrent paths).
  - The R2 fix does not touch the single-use token guard or the
    `issueNewVerificationToken` transaction shape; R12's concurrent
    double-submit and R16's 100-goroutine race tests still pass under
    `-race`.

- **What is not tested, and why:**
  - No new unit test for the `/docs` or `/openapi.yaml` handlers — they
    are thin dev-only read handlers (serve a constant HTML string /
    stream a file with one byte replacement). The end-to-end curl
    checks (`GET /docs` → HTML, `GET /openapi.yaml` → spec with
    rewritten server) are the proof.
  - `make verify` not run (pre-existing gosec G304/G112 in unrelated
    files block the `lint` target). No new findings expected in the
    changed packages, but this should be confirmed before merge.
