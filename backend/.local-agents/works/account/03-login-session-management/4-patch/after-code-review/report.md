# Patch Report — Login & Session Management (account #03)

> Ticket      : account domain task #3 — `docs/spec/1-account/features/03-login-session-management.md`
> Executed    : 2026-08-26
> Phase       : 4-patch (code-review follow-up)
> Source      : `4-code-review/report.md` (Request changes — Q1 blocking, Q2/Q3 optional)
> Decision    : Option A for all three tasks (fail-open; TTL dedup; doc-only logout exception)
> Status      : Patch complete — all verification green

---

## Execution summary

| Task | Finding | Priority | Result |
|---|---|---|---|
| 01 | `writeAttempt` fail-open vs fail-closed doc/code/test disagreement | **Blocking** | ✅ Implemented fail-open (Option A); 2 new tests pass |
| 02 | Duplicated `accessTokenTTL` constant across two packages | Optional | ✅ Collapsed to `auth.AccessTokenTTL` |
| 03 | `LogoutHandler` 500-on-infra vs "always 204" contract wording | Optional | ✅ Documented the exception (Option A, doc-only) |

## Files changed

| File | Change type | Description |
|---|---|---|
| `internal/domain/account/login.go` | Edit (untracked from build) | `writeAttempt` — 3 error-return paths changed to log + `return nil` (fail-open); doc comment updated |
| `internal/domain/account/login_test.go` | Edit (untracked from build) | +`failingAttemptRepo` type; +2 fail-open tests; +`platform/auth` import |
| `internal/domain/account/google_oauth.go` | Edit (tracked) | Removed `accessTokenTTL` const; `IssueTokens` uses `auth.AccessTokenTTL`; +`platform/auth` import; const-block doc updated |
| `internal/domain/account/google_oauth_test.go` | Edit (tracked) | `accessTokenTTL` → `auth.AccessTokenTTL` (2 refs) |
| `internal/platform/auth/token.go` | Edit (untracked from build) | `AccessTokenTTL` doc comment updated — single source of truth |
| `internal/transport/http/auth_login.go` | Edit (untracked from build) | `LogoutHandler` doc — documents 500-on-infra exception |

## Patch details

### Task 01 — `writeAttempt` fail-open (Q1, blocking)

**Problem:** The `writeAttempt` doc and build-report deviation #5 described
fail-open behavior (a lost audit-row write is logged, credential decision
stands), but the code returned the write error → 500 (fail-closed). Neither
path was tested (`loginFakeRepo.InsertLoginAttempt` always returned nil).

**Fix (Option A — fail-open, Anhar-approved):** All three DB error paths in
`writeAttempt` now log the error and return `nil`:

```go
// login.go:443-477 — writeAttempt (patched)
func (s *Service) writeAttempt(ctx context.Context, attempt *LoginAttempt) error {
    tx, err := s.tx.BeginTx(ctx)
    if err != nil {
        log.Printf("account: begin attempt tx failed (fail-open): stage=%s success=%t: %v",
            attempt.Stage, attempt.Success, err)
        return nil
    }
    committed := false
    defer func() {
        if !committed {
            _ = tx.Rollback(ctx)
        }
    }()
    if err := s.repo.InsertLoginAttempt(ctx, tx, attempt); err != nil {
        log.Printf("account: insert login attempt failed (fail-open): stage=%s success=%t: %v",
            attempt.Stage, attempt.Success, err)
        return nil
    }
    if err := tx.Commit(ctx); err != nil {
        log.Printf("account: commit attempt tx failed (fail-open): stage=%s success=%t: %v",
            attempt.Stage, attempt.Success, err)
        return nil
    }
    committed = true
    return nil
}
```

**Log hygiene (R19):** the logged error comes from pgx/goqu — it holds SQL
detail but no user credentials or tokens. Only non-sensitive `stage` +
`success` metadata are included in the log line. The `attempt.IdentifierHash`
and `attempt.UserID` are NOT logged.

**Tests added** (`login_test.go:767-813`):

```go
type failingAttemptRepo struct {
    *loginFakeRepo
}
func (f *failingAttemptRepo) InsertLoginAttempt(...) error {
    return errors.New("simulated audit-db outage")
}

func TestWriteAttempt_FailOpen_ValidLoginStillSucceeds(t *testing.T) { ... }
func TestWriteAttempt_FailOpen_InvalidCredentialsStillRejected(t *testing.T) { ... }
```

- `ValidLoginStillSucceeds`: valid creds + audit-write failure → login
  succeeds (Status "ok", tokens issued).
- `InvalidCredentialsStillRejected`: wrong password + audit-write failure →
  still `ErrInvalidCredentials` (fail-open doesn't mask credential
  rejection).

### Task 02 — TTL constant dedup (Q2, optional)

**Problem:** `accessTokenTTL` (account package, `google_oauth.go:50`) and
`auth.AccessTokenTTL` (`platform/auth/token.go:42`) were both
`15 * time.Minute` — two sources of truth for the same load-bearing
lifetime.

**Fix:** Removed the account-package constant; all references now use
`auth.AccessTokenTTL`:

| File | Line(s) | Change |
|---|---|---|
| `google_oauth.go` | const block (48-52) | Removed `accessTokenTTL = 15 * time.Minute` |
| `google_oauth.go` | imports | Added `platform/auth` |
| `google_oauth.go` | `IssueTokens` (494) | `accessTokenTTL` → `auth.AccessTokenTTL` |
| `google_oauth_test.go` | 501-502 | `accessTokenTTL` → `auth.AccessTokenTTL` |
| `login_test.go` | 251, 606 | `accessTokenTTL` → `auth.AccessTokenTTL` |
| `login_test.go` | imports | Added `platform/auth` |
| `token.go` | 40-42 | Doc comment updated — "single source of truth" |

`refreshTokenTTL` stays in the account package (no `platform/auth/`
counterpart — refresh tokens are opaque random values, not JWTs).

### Task 03 — Logout 500-on-infra doc (Q3, optional)

**Problem:** R16 / techplan §3 frame logout as "always 204," but
`LogoutHandler` maps a genuine DB error to 500 via `MapServiceError`.

**Fix (Option A — doc-only, no spec edit):** `LogoutHandler` doc comment
updated to document the exception:

```go
// LogoutHandler handles POST /auth/logout idempotently: revokes the
// presented refresh token when present, ALWAYS clears the cookie, and
// answers 204 for every idempotent case — an absent or already-dead
// cookie is not an error condition (R16). The sole exception is a
// genuine infrastructure failure (e.g. the revoke UPDATE fails at the DB
// level): that surfaces as a 500 Problem Details response rather than a
// masked 204, so real outages are not hidden behind a success code.
```

No code change. No spec/openapi edit (Option B would require human spec
approval per AGENTS.md §4).

## Verification results

```
$ go build ./...
(no output — clean)

$ go test ./internal/domain/account/ -v -run 'TestWriteAttempt|TestLogin_|TestLoginMfa_|TestRefresh_|TestLogout_'
=== RUN   TestLogin_Success_NoMfa              --- PASS (0.18s)
=== RUN   TestLogin_MfaRequired_NoTokensIssuedYet --- PASS (0.18s)
=== RUN   TestLogin_GenericErrorMessage        --- PASS (0.30s)
=== RUN   TestLogin_Lockout_5Failed15Min       --- PASS (0.18s)
=== RUN   TestLogin_UnverifiedIdentity_Succeeds --- PASS (0.18s)
=== RUN   TestLogin_TimingShape_NoEarlyReturn  --- PASS (0.24s)
=== RUN   TestLogin_LoggingNeverLeaksCredentials --- PASS (0.18s)
=== RUN   TestLoginMfa_Lockout_5Failed15Min    --- PASS (0.06s)
=== RUN   TestLoginMfa_InvalidPendingToken     --- PASS (0.06s)
=== RUN   TestLoginMfa_WrongCode               --- PASS (0.06s)
=== RUN   TestLoginMfa_TotpSuccess_CompletesLogin --- PASS (0.06s)
=== RUN   TestLoginMfa_BackupCode_CompletesLogin --- PASS (0.06s)
=== RUN   TestLoginMfa_DefensiveBothOrNeitherCodes --- PASS (0.06s)
=== RUN   TestRefresh_Rotates_IssuesChild_SameFamily --- PASS (0.06s)
=== RUN   TestRefresh_MissingOrExpiredCookie   --- PASS (0.06s)
=== RUN   TestRefresh_ReuseDetection_FamilyRevoked --- PASS (0.06s)
=== RUN   TestRefresh_ConcurrentRequests_ExactlyOneWins --- PASS (0.06s)
=== RUN   TestLogout_RevokesAndClears          --- PASS (0.06s)
=== RUN   TestLogout_NoCookie_Still204         --- PASS (0.06s)
=== RUN   TestWriteAttempt_FailOpen_ValidLoginStillSucceeds --- PASS (0.18s)
=== RUN   TestWriteAttempt_FailOpen_InvalidCredentialsStillRejected --- PASS (0.18s)
PASS
ok  github.com/anhsbolic/kencleng/backend/internal/domain/account  2.502s

$ go test -race ./...
ok  github.com/anhsbolic/kencleng/backend/internal/domain/account   124.194s
ok  github.com/anhsbolic/kencleng/backend/internal/platform/auth     1.051s
ok  github.com/anhsbolic/kencleng/backend/internal/platform/breachcheck (cached)
ok  github.com/anhsbolic/kencleng/backend/internal/platform/crypto    (cached)
ok  github.com/anhsbolic/kencleng/backend/internal/platform/googleoauth (cached)
ok  github.com/anhsbolic/kencleng/backend/internal/platform/notification (cached)
ok  github.com/anhsbolic/kencleng/backend/internal/platform/secrets   (cached)
ok  github.com/anhsbolic/kencleng/backend/internal/transport/http     1.046s
PASS

$ staticcheck ./...
(no output — clean)
```

All 21 login/session unit tests pass (19 pre-existing + 2 new fail-open
tests). Full suite green under `-race`. staticcheck clean.

## Risk note (root AGENTS.md §5)

- **Assumptions made:** Fail-open is the correct trade-off for the
  `login_attempts` audit table (a lost row can only undercount toward
  lockout, never lock spuriously); the doc and build-report deviation #5
  already described this as the accepted stance — this patch makes the code
  match. Named tests: `TestWriteAttempt_FailOpen_ValidLoginStillSucceeds`,
  `TestWriteAttempt_FailOpen_InvalidCredentialsStillRejected`.
- **Edge cases intentionally NOT handled:** A persistent audit-DB outage
  will silently lose all attempt rows for its duration — lockout thresholds
  undercount during the outage. Accepted (same trade-off as the original
  doc's framing); the outage itself is observable via the logged errors.
- **Concurrency assumptions:** Unchanged from build phase — `writeAttempt`
  holds no shared state; the fail-open change is per-call and does not
  introduce new concurrency surfaces.
- **What is not tested, and why:** The fail-open path under real Postgres
  (integration level) — the unit test with `failingAttemptRepo` proves the
  contract at the service layer, which is where the fail-open decision
  lives. Integration-level audit-row failure injection would test pgx
  behavior, not the fail-open decision itself.

## Outstanding gates (not this patch's scope)

1. **Tier 0 paired rewrite pass** (techplan Resolved #13) — still pending.
   Must cover `platform/auth/token.go`, `repository_db.go` rotation
   methods, `login.go` reuse/race-loser branch BEFORE any commit.
2. **CSRF second layer** — accepted residual risk (techplan Resolved #7).
   Revisit trigger: frontend API client landing.
3. **`make verify` gosec stage** — still red on pre-existing 13-finding
   baseline (not this slice's contribution; documented in `3-build/report.md`).

## Feature spec reference

Fulfills `docs/spec/1-account/features/03-login-session-management.md`,
subject to the Tier 0 paired pass and the code-review follow-up applied
herein.
