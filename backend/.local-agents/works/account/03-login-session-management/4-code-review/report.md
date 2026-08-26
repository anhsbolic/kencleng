# Code Review Report — Login & Session Management (account #03)

> Ticket      : account domain task #3 — `docs/spec/1-account/features/03-login-session-management.md`
> Reviewed    : 2026-08-26
> Inputs      : `2-plan/techplan.md` (Status: Approved), `2-plan/tasks/` decomposition, `3-build/report.md`, working-tree diff (unstaged + untracked)
> Method      : Four-pass review (Safety → Quality → Stack-Specific Best Practices → Consistency) per `harscode-workspace/workflow/4-code-review/guidelines.md`
> Verdict     : **Request changes** (1 blocking, 2 optional, 1 process gate, 1 accepted residual)

---

## Scope of review

The diff under review is the full login/session vertical slice built per
the approved techplan:

- **New files**: `migrations/000006..000009` (8 SQL), `internal/domain/account/{login.go, login_test.go, mfa_verifier.go, login_integration_test.go}`, `internal/platform/auth/{token.go, token_test.go}`, `internal/transport/http/{auth_login.go, auth_login_test.go}`
- **Edited files**: `entity.go`, `repository.go`, `repository_db.go`, `repository_db_integration_test.go`, `service.go`, `service_test.go`, `internal/transport/http/{cookie.go, errors.go}`, `cmd/server/main.go`, `.env.example`; two pre-existing staticcheck touch-ups (`google_oauth_test.go`, `googleoauth/helpers_test.go`)
- **Out of review scope** (per directory boundary, root AGENTS.md §7): the frontend diffs in the same working tree; the `docs/spec/*` and `api/openapi.yaml` edits (applied separately under human approval per AGENTS.md §4)

The Tier 0 fenced sub-area (`platform/auth/token.go`, `repository_db.go`
rotation methods, `login.go` reuse/race-loser branch) is agent-drafted and
under review here, but the review does **not** substitute for the mandatory
human paired rewrite pass (techplan Resolved #13; see Process Gate below).

---

## 1. Safety

**No findings.**

Nil-safety: every nullable return is guarded.
- `identity` / `identity.CredentialSecret` checked before deref (`login.go:124`).
- `view` nil-checked in both `Login` and `LoginMfa` (`login.go:148`, `:245`).
- `current` only touched when `found` is true, relying on `&&` short-circuit (`login.go:278-292`).
- mint/verify closures nil-checked in their wrappers (`mintAccessFor`, `mintPending`, `verifyPendingFor`).

Concurrency: `Service` holds no mutable state; `dummyBcryptHash` uses `sync.OnceValue`; rotation correctness lives entirely behind the guarded `UPDATE … WHERE replaced_by_id IS NULL AND revoked_at IS NULL AND expires_at > now()` (`repository_db.go` `RotateRefreshToken`) — the sole writer of `replaced_by_id`. The race-loser connection-leak (found by the ≥100-goroutine stress harness) is correctly fixed: explicit `tx.Rollback` + `committed = true` before opening the family-revoke tx (`login.go:339-347`).

Errors: propagated with `%w`, never swallowed. The `_ = tx.Rollback(ctx` ignores are safe no-ops per `go/defer-pitfalls.md` §3 (Rollback on a committed/rolled-back tx is a documented no-op).

Context: `ctx` threads through all I/O. The mint closures are CPU-only crypto (no ctx needed). `rows.Close()` / `roleRows.Close()` deferred in `GetLoginUserView`.

---

## 2. Quality

### Q1 — `writeAttempt` doc describes fail-open; code is fail-closed; neither is tested (BLOCKING)

- **Location**: `internal/domain/account/login.go:449-468` (doc) vs `:135-137` / `:233-235` (call sites)
- **What's wrong**: The `writeAttempt` doc states "if the write fails after the credential decision, the login result stands, the failure is logged" (fail-open). The build report deviation #5 frames undercount-on-lost-row as the accepted trade-off, implying fail-open. The code instead returns the write error → handler maps it to a 500, failing a valid login on any transient audit-DB hiccup (fail-closed). `loginFakeRepo.InsertLoginAttempt` always returns nil, so the behavior is untested either way.
- **Why it matters**: A reviewer/operator cannot tell which was intended. If fail-open was the accepted decision (the doc's reasoning: "a lost attempt row can only undercount toward the lockout threshold, never lock anyone out spuriously"), an audit-DB outage takes down all logins under the current code. If fail-closed is the actual choice, the doc and report mislead future readers.
- **Suggested fix**: Pick one and make doc/code/test agree. Recommended: **fail-open** (implement the documented behavior + add a test). See task-01 in `tasks/`.

### Q2 — Duplicated access-token TTL constant across two packages (OPTIONAL)

- **Location**: `internal/domain/account/google_oauth.go:50` (`accessTokenTTL = 15*time.Minute`) vs `internal/platform/auth/token.go:42` (`AccessTokenTTL = 15*time.Minute`)
- **What's wrong**: Two sources of truth for the same load-bearing lifetime. The login slice mints the JWT `exp` and the response `ExpiresAt` from `auth.AccessTokenTTL`, while `login_test.go` asserts against the account package's `accessTokenTTL`. If one drifts, the wire `expires_at` silently disagrees with the JWT's real `exp` (and the test would assert the wrong value).
- **Suggested fix**: Have the account package reference `auth.AccessTokenTTL` instead of keeping its own. See task-02 in `tasks/`.

### Q3 — `LogoutHandler` can return 500 where the contract says "always 204" (OPTIONAL)

- **Location**: `internal/transport/http/auth_login.go:170-177`
- **What's wrong**: R16 / techplan §3 frame logout as "always 204 / none documented." The idempotent cases (no cookie, not-found, already-revoked) do return 204, but a genuine DB error yields a Problem 500. Defensible (don't mask infra failures), but the literal contract wording disagrees.
- **Suggested fix**: Add a one-line handler doc noting infra failures are the deliberate exception, or align the contract wording. See task-03 in `tasks/`.

---

## 3. Stack-Specific Best Practices

Matched trigger keywords in `best-practices/index.md` against the diff's
technology: **Go**, **jwt/token/refresh-rotation**, **mfa**, **rate-limit/lockout**,
**secrets/keys**, **anti-enumeration**, **csrf/cookie**, **postgresql transactions/migrations/encryption**,
**error-wrapping**, **defer**, **concurrency/integration testing**.

Files opened and checklists applied:

| Best-practices file | Result |
|---|---|
| `go/jwt-and-token-lifecycle.md` | Clean — separate keys/algorithms (ES256 access vs HS256 mfa_pending); rotate-on-use; family-wide reuse detection; deliberate XSS-vs-CSRF storage split (access in body / refresh HttpOnly cookie); proportional TTLs (15m / 5m / 30d) |
| `go/secrets-and-key-management.md` | Clean — one-key-one-purpose upheld; `.env` gitignored (verified); `.env.example` ships empty default. (Key-rotation plan not documented — out of scope for this slice) |
| `go/secrets-and-sensitive-logging.md` | Clean — log lines carry `user_id` + category only; R19 marker leak-sweep test exists (`login_test.go:418`) |
| `restapi/anti-enumeration.md` | Clean — byte-identical 401 bodies across wrong-email/wrong-password; 429 detail byte-identical to 401; dummy bcrypt burn on unknown identifier; lockout keyed by `identifier_hash` (ghost emails lock identically, no existence leak via 429) |
| `postgresql/transactions-and-locking.md` | Clean — no external call inside any tx; stale `FindRefreshTokenByHash` read not trusted (guarded `UPDATE … RETURNING` is the atomic arbiter); single-statement guard avoids multi-row lock-ordering |
| `postgresql/migrations-safety.md` | Clean — additive `CREATE TABLE` pairs, zero `ALTER`/backfill, reversible `DROP TABLE IF EXISTS`, round-trip verified per build report |
| `postgresql/encryption-at-rest.md` | Clean — ciphertext + separate HMAC lookup; decrypt-on-read by ID only (`GetLoginUserView`); HMAC key ≠ encryption key |
| `go/defer-pitfalls.md` | Clean — no defer-in-loop; `committed` flag pattern correctly scopes rollback; LIFO order correct |
| `go/error-wrapping.md` | Clean — `%w` throughout; `errors.Is` at transport layer, no string matching |
| `go/context-propagation.md` | Clean — no `context.Background()` mid-chain; ctx not stored in struct |
| `go/testing-concurrency.md` | Clean — `-race` gate; ≥100-goroutine stress harness; invariant-asserting concurrent tests |
| `go/integration-testing-setup.md` | Clean — `//go:build integration` tag on both integration files (verified); reserved for constraint/transaction proofs a fake repo can't prove |
| `restapi/csrf-and-cookie-security.md` | **One accepted residual** — see below |

### CSRF — accepted residual risk (non-blocking, documented)

- **Location**: `internal/transport/http/cookie.go` (`writeRefreshCookie` / `clearRefreshCookie`) + `auth_login.go` `RefreshHandler` / `LogoutHandler`
- **Checklist item unsatisfied**: `restapi/csrf-and-cookie-security.md` item 2 — "There is a second CSRF mitigation (custom header check or double-submit token) on top of `SameSite`." Only `SameSite=Strict` + `HttpOnly` + `Secure`-conditional are set; no custom-header or double-submit layer.
- **Status**: **Explicitly accepted, human-approved** (techplan Resolved #7; threat-model residual-risk entry #6). Rationale: no untrusted same-site subdomains exist in the v1 sandbox topology; `SameSite=Strict` alone is sufficient for v1. Revisit trigger = frontend API client landing (centralized React API client is the natural place to add a custom header at near-zero cost). No action required in this slice beyond keeping the recorded acceptance tracked.

---

## 4. Consistency

**No findings.** Checked against `backend/AGENTS.md` (§1-§7) and `backend/README.md`:

- **Error handling**: `fmt.Errorf("…: %w", err)` wrapping throughout (AGENTS.md §2); `MapServiceError` uses `errors.Is`, no string matching — matches existing `errors.go` pattern.
- **Logging**: stdlib `log.Printf("account: … user_id=%s")` matches existing account-package style (no structured logger in this repo).
- **Validation**: handler-boundary validation via `WriteValidationError` / `write400InvalidJSON` follows the `auth_register.go` precedent.
- **SQL**: all new SQL is `goqu`-built with `Prepared(true)` (golden rule + AGENTS.md §2); nullable scanning uses `sql.NullTime` / `sql.NullString`, the established pattern already in `repository_db.go:217-260`.
- **PII**: decrypt-on-read via `crypto.Decrypt` with `r.keys`; lookups via `*_hash` — matches the encryption pattern in the tech-stack doc.
- **Doc comments**: present on every new exported symbol (AGENTS.md §2).
- **Routing**: `authMux.HandleFunc("POST /auth/…", …)` (Go 1.22 pattern routing, AGENTS.md §2); integration tests gated behind `//go:build integration` (AGENTS.md §3); naming `<thing>_integration_test.go` consistent with existing `repository_db_integration_test.go`.
- **Cross-domain imports**: none — `login.go` imports only `platform/*` (AGENTS.md §1).
- **Staticcheck touch-ups**: the two fixes in `google_oauth_test.go` (S1024) and `googleoauth/helpers_test.go` (U1000) are justified `make verify` gate fixes, not unrelated feature drift.

---

## Verdict

**Request changes.**

### Blocking (must fix before commit)

| # | Finding | Fix |
|---|---|---|
| Q1 | `writeAttempt` doc/code/test disagreement on fail-open vs fail-closed | task-01 |

### Optional (non-blocking, cheap)

| # | Finding | Fix |
|---|---|---|
| Q2 | Duplicated access-token TTL constant | task-02 |
| Q3 | Logout 500-on-infra vs "always 204" contract wording | task-03 |

### Process gate (not a code change)

The build report correctly self-flags **"BLOCKED on Tier 0 paired pass before any commit."** The human paired rewrite/review must cover exactly these agent-drafted Tier 0 regions before any commit (root AGENTS.md §3 fence; techplan Resolved #13):

1. `internal/platform/auth/token.go` — both token purposes' mint/verify
2. `internal/domain/account/repository_db.go` — rotation methods (`RotateRefreshToken`, `RevokeRefreshTokenByHash`, `RevokeRefreshTokenFamily`)
3. `internal/domain/account/login.go` — reuse/race-loser branch of `Refresh` (incl. the connection-leak fix)

The Safety / Quality / Best-Practices passes above review agent-drafted Tier 0 code and are **not** a substitute for that gate.

### Accepted residual (keep tracked)

| # | Item | Revisit trigger |
|---|---|---|
| — | CSRF second layer on cookie-authenticated refresh/logout | Frontend API client landing (centralized React fetch wrapper is the natural place to add a custom header) |
