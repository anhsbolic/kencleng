# Patch Implementation Plan — Manual Testing Tooling + R2 Fix

> Ticket    : 01-register-email-verification
> Stage     : 4-patch (manual-testing enablement + correctness fix)
> Date      : 2026-08-22
> Author    : GLM 5.2 (opencode)
> Source    : manual-testing session — no prior review finding; R2 bug
>            discovered while exercising the register/verify flow
> Feature   : `docs/spec/1-account/features/01-register-email-verification.md`
> Prior patch : `./plan-code-review.md`, `./report-code-review.md` (Tasks A–E)
> Convention: `AGENTS.md` (root + backend), `backend/README.md`

---

## 0. Status

Two concerns surfaced when attempting to manually test the three auth
endpoints via a browser-based Swagger UI:

1. **No Swagger UI served, no CORS** — the backend exposes no docs page
   and no CORS middleware. A browser Swagger pointed at `:8090` from
   another origin is blocked from reading responses.
2. **Token not retrievable for manual verify** — `FakeSender` (v1 email
   stand-in) deliberately never logs the recipient or token (golden
   rule). The plaintext token exists only during the
   `SendVerificationEmail` call and is stored in the DB only as a
   SHA-256 hash (not recoverable), so there is no way to obtain a
   verification token to exercise `POST /auth/verify-email`.

While verifying the token-retrieval fix, a **correctness bug** was
found in the R2 branch (register unverified existing): it issues a new
token but discards the plaintext and sends a token-less nudge, so the
new token is never delivered. This contradicts the feature spec.

This plan addresses both.

## 1. Findings

### F1 — No Swagger UI / no CORS (dev-tooling gap)

- `internal/transport/http/` has no docs handler; `cmd/server/main.go`
  registers only `/healthz` + the three `/auth/*` routes.
- No CORS middleware exists (`grep` for `cors|Access-Control|Origin`
  → 0 matches in `backend/`).
- The bundled OpenAPI spec (`api/openapi.yaml`) declares
  `servers: url: /api` (relative, expects the Caddy reverse proxy).
- **Pre-existing infra mismatch (flagged, not fixed here):** the
  `Caddyfile` uses `handle /api/*` (not `handle_path`), so Caddy does
  NOT strip the `/api` prefix — a request to `:8080/api/auth/register`
  reaches the backend as `/api/auth/register`, but the backend
  registers `/auth/register` → 404. The Caddyfile is a root-level file
  outside the `backend/` session scope (AGENTS.md §7).

### F2 — Token not retrievable (dev-tooling gap)

- `platform/notification/sender.go` `FakeSender.SendVerificationEmail`
  logs `"verification email queued (recipient redacted)"` and drops
  the token. This is correct for the secure default but blocks manual
  testing — the plaintext token leaves the process exactly once (inside
  the sender call) and is never persisted in recoverable form.

### F3 — R2 branch does not deliver the new token (correctness bug)

- `internal/domain/account/service.go` R2 branch
  (`identity.VerifiedAt == nil`, lines ~153–161) calls
  `issueNewVerificationToken` but **discards** the returned plaintext
  (`if _, err := …`) and calls `sendNudge(ctx, email,
  notification.NudgeResendVerification)` — a token-less nudge.
- Compare R13 `ResendVerification` (lines ~393–401), which correctly
  captures `plainToken` and calls `sendVerification(ctx, email,
  plainToken)`.
- **Spec authority:** `docs/spec/1-account/features/01-register-email-verification.md:29-31`
  — R2 must fire "the same internal action as `verify-email/resend`
  (old unused token revoked, new one issued, resend-verification
  **email sent**)". The resend endpoint's spec (lines 84–85) sends "a
  new verification email" — i.e. the email carrying the token. So R2
  must deliver the token via `sendVerification`, exactly like R13.
- The existing test `TestRegister_UnverifiedExisting_ResendFlow`
  asserted the **wrong** behavior (a `NudgeResendVerification` nudge),
  so the bug was hidden by a test that matched the buggy code.

## 2. Decisions (resolved)

| Decision | Choice | Rationale |
|---|---|---|
| F1 Swagger UI host | Embed in the Go backend (same-origin) | Avoids CORS entirely; no extra container/Caddy changes; stays within `backend/` scope |
| F1 spec serving | `GET /openapi.yaml` reads `../api/openapi.yaml` from disk (dev-only) | `go:embed` cannot reach outside the module tree; dev-only disk read is acceptable. Source spec file never modified (AGENTS.md §4) |
| F1 server URL | Rewrite `url: /api` → `http://localhost:8090` in the served copy | Makes Swagger "Try it out" hit the backend's real `/auth/*` routes directly, bypassing the broken Caddy `/api`-no-strip path (F1 infra mismatch) |
| F1 Swagger assets | CDN (`unpkg.com/swagger-ui-dist@5`) | No Go dependency, no vendored assets; dev-only page |
| F2 token retrieval | Dev-only `DevSender` writing to a dev outbox file | Tokens stay out of structured `log.Printf` (outbox = simulated inbox, not a log stream → golden rule honored). Gated on `APP_ENV=development`; `FakeSender` remains the secure default in every other env, and its existing tests stay untouched |
| F2 outbox path | `os.TempDir()/kencleng-dev-outbox.log`, overridable via `DEV_OUTBOX_PATH`; logged once at startup | Developer knows where to find tokens; file mode 0600 (owner-only) |
| F3 R2 fix | Capture `plainToken`, call `sendVerification` (like R13) | Matches spec; delivers the token so the user can verify. DB-work shape unchanged (both `sendVerification`/`sendNudge` are post-commit, no network in FakeSender/DevSender) → R7 anti-enumeration timing preserved |
| F3 test update | Assert verification email sent + no nudge | Aligns test with spec (tightens, not loosens — AGENTS.md §4 allows when the test was wrong and the spec is the authority) |
| F3 dead constant | Remove `NudgeResendVerification` | Only consumer was the buggy R2 branch; R3/R4 nudges (`NudgePasswordReset`, `NudgeGoogleOnly`) unchanged |

## 3. Execution order & grouping

```
Step 1 (dev tooling, file-disjoint):
  F1  swagger.go (new) + main.go wiring (/docs, /openapi.yaml)
  F2  dev_sender.go (new) + main.go wiring (newEmailSender)

Step 2 (correctness, disjoint from step 1 files):
  F3  service.go R2 branch + service_test.go R2 test + sender.go constant

Step 3 (verify):
  build + vet + unit + race + end-to-end curl (R1/R2/R8/R10/R13/R14)
```

F1/F2 touch `transport/http/`, `notification/`, `cmd/server/`. F3 touches
`domain/account/`, `notification/` (constant removal). The only overlap
is `notification/sender.go` (F3 removes a constant) — file-disjoint
methods from F2's `dev_sender.go` (new file). No conflict.

## 4. Out of scope (flagged, not fixed)

- **Caddy `/api` prefix mismatch** (F1 infra) — root-level `Caddyfile`,
  outside `backend/` session scope (AGENTS.md §7). Manual Swagger
  sidesteps it by going direct to `:8090`. A separate session should
  switch Caddy to `handle_path /api/*` (strips prefix).
- **`make verify` pre-existing gosec findings** — `platform/auth/keys.go`
  (G304) and `cmd/server/main.go` (G112), not from this or prior patch
  work. Address separately or suppress with justified `//nosec`.
- **Real SMTP** — out of scope for v1 (per `kencleng-phase3-detail.md`);
  `DevSender`/`FakeSender` remain the stand-ins.

## 5. Fence compliance

No Tier 0 fenced path (AGENTS.md §3) is modified:
- `platform/crypto/` — untouched
- `platform/auth/` — untouched
- `domain/donation/ledger.go` — untouched
- `domain/disbursement/` — untouched

No `docs/spec/*` file is edited (AGENTS.md §4). The bundled
`api/openapi.yaml` is read at runtime but never modified on disk.
