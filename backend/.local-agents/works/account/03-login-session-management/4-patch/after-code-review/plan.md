# Patch Plan — Login & Session Management (account #03)

> Ticket      : account domain task #3 — `docs/spec/1-account/features/03-login-session-management.md`
> Phase       : 4-patch (code-review follow-up)
> Date        : 2026-08-26
> Source      : `4-code-review/report.md` (Request changes)
> Decision    : Option A for task-01 (fail-open), as approved by Anhar

---

## Objective

Apply the three findings from the code review (`4-code-review/report.md`):

1. **Q1 (blocking)** — `writeAttempt` doc described fail-open but code was
   fail-closed, untested either way. Implement **fail-open** (Option A,
   Anhar-approved): log the audit-write error and return `nil` so the
   credential decision stands. Add two tests proving the contract.
2. **Q2 (optional)** — Duplicated `accessTokenTTL` constant in two packages.
   Collapse to a single source of truth: `auth.AccessTokenTTL`.
3. **Q3 (optional)** — `LogoutHandler` doc said "always 204" but the handler
   can return 500 on infra failure. Document the exception (Option A,
   doc-only, no contract/spec edit).

## Files changed

| File | Task | Change |
|---|---|---|
| `internal/domain/account/login.go` | Q1 | `writeAttempt` body — 3 error returns → log + `return nil`; doc comment updated to state fail-open explicitly |
| `internal/domain/account/login_test.go` | Q1 | +`failingAttemptRepo` type; +`TestWriteAttempt_FailOpen_ValidLoginStillSucceeds`; +`TestWriteAttempt_FailOpen_InvalidCredentialsStillRejected`; import `platform/auth` added (for Q2) |
| `internal/domain/account/google_oauth.go` | Q2 | Removed `accessTokenTTL` const; `IssueTokens` now uses `auth.AccessTokenTTL`; import `platform/auth` added; const-block doc updated |
| `internal/domain/account/google_oauth_test.go` | Q2 | `accessTokenTTL` → `auth.AccessTokenTTL` (2 refs); pre-existing `rt.ExpiresAt.Sub` → `time.Until` staticcheck fix preserved |
| `internal/platform/auth/token.go` | Q2 | `AccessTokenTTL` doc comment updated to mark it as the single source of truth |
| `internal/transport/http/auth_login.go` | Q3 | `LogoutHandler` doc comment — documents the 500-on-infra-failure exception to "always 204" |

## Approach

### Q1 — fail-open (Option A)

The `writeAttempt` helper persists a `login_attempts` row in its own short
transaction. The audit row is bookkeeping, not a state machine — a lost row
can only undercount toward the lockout threshold, never lock anyone out
spuriously. Fail-open means: if any of the three DB operations (BeginTx,
InsertLoginAttempt, Commit) fails, log the error (stage + success + wrapped
pgx/goqu error — no credentials/tokens per R19) and return `nil`, so the
credential decision from the caller (`Login` or `LoginMfa`) stands.

Two tests prove the contract using `failingAttemptRepo` (wraps
`loginFakeRepo`, forces `InsertLoginAttempt` to error):
- Valid credentials + audit-write failure → login still succeeds (Status
  "ok", tokens issued).
- Invalid credentials + audit-write failure → still `ErrInvalidCredentials`
  (fail-open does not mask the credential rejection).

### Q2 — TTL constant dedup

`accessTokenTTL` (account package, `google_oauth.go:50`) and
`auth.AccessTokenTTL` (`platform/auth/token.go:42`) were both `15 *
time.Minute`. `platform/auth/` is the canonical home (it mints the JWT —
the TTL is a property of the token primitive). The account package now
references `auth.AccessTokenTTL` everywhere. `refreshTokenTTL` stays in the
account package (no counterpart in `platform/auth/` — refresh tokens are
opaque random values, not JWTs).

### Q3 — logout doc exception (Option A)

Doc-only change to `LogoutHandler`: the "always 204" framing now
acknowledges the sole exception — a genuine infra failure (e.g. the revoke
UPDATE fails at the DB level) surfaces as 500 Problem Details, not a masked
204. No spec/openapi edit (that would be Option B, requiring human spec
approval per AGENTS.md §4).

## Verification plan

```bash
go build ./...                    # compile
go test ./internal/domain/account/ -run 'TestWriteAttempt|TestLogin_|TestLoginMfa_|TestRefresh_|TestLogout_' -v  # targeted
go test -race ./...               # full suite + race
staticcheck ./...                 # lint
```

## Out of scope

- Tier 0 paired rewrite pass (separate human gate — techplan Resolved #13)
- Integration tests (already green from build phase; no integration-level
  change in this patch)
- `make verify` full gate (gosec still red on pre-existing baseline — not
  this slice's contribution; documented in `3-build/report.md`)
- Any spec/openapi edit (Q3 Option B would require human approval)
