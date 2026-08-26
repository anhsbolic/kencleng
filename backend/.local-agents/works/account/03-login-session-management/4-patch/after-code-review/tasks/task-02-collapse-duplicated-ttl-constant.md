# Task 02: Collapse duplicated access-token TTL constant

> Back-reference : `4-code-review/report.md` §2 (Q2)
> Priority       : Optional (non-blocking, cheap)
> Model          : DeepSeek V4 Pro — mechanical constant consolidation
> Depends on     : none

## Objective

The 15-minute access-token TTL exists as two separate constants in two
packages: `accessTokenTTL` in `internal/domain/account/google_oauth.go:50`
and `auth.AccessTokenTTL` in `internal/platform/auth/token.go:42`. The login
slice mints the JWT `exp` from `auth.AccessTokenTTL` but the test
(`login_test.go:250`) asserts against the account package's `accessTokenTTL`.
If one drifts, the wire `expires_at` silently disagrees with the JWT's real
`exp` — and the test would assert the wrong value. Collapse to one source of
truth.

## Files

| File | Change |
|---|---|
| `backend/internal/domain/account/google_oauth.go` | Remove local `accessTokenTTL` const; reference `auth.AccessTokenTTL` instead |
| `backend/internal/domain/account/google_oauth_test.go` | Update any test references from `accessTokenTTL` to `auth.AccessTokenTTL` |
| `backend/internal/domain/account/login_test.go` | Update `accessTokenTTL` references to `auth.AccessTokenTTL` (lines 250, 605) |
| `backend/internal/domain/account/login_integration_test.go` | Update if it references `accessTokenTTL` (check; integration tests may use `refreshTokenTTL` only) |

## Approach

`auth.AccessTokenTTL` is the canonical home (it lives in `platform/auth/`,
the package that actually mints the token — the TTL is a property of the
token primitive, not the account domain). The account package already
imports `platform/auth` in `login.go` and `service.go`.

### Step 1 — remove the duplicate

In `google_oauth.go`, remove:

```go
accessTokenTTL  = 15 * time.Minute
```

Keep `refreshTokenTTL` in place for now (it has no counterpart in
`platform/auth/` — refresh tokens are opaque random values, not JWTs; the
TTL is an account-domain concern). If a future task moves refresh TTL to a
shared location, do it then — this task is scope-disciplined to the access
constant only.

### Step 2 — update references

Replace all `accessTokenTTL` references in the account package with
`auth.AccessTokenTTL`. Known locations (verify with grep before editing):

```
google_oauth.go:492     — ExpiresAt: jwt.NewNumericDate(now.Add(accessTokenTTL))
google_oauth_test.go:501 — time.Until(rc.ExpiresAt.Time) > accessTokenTTL+5*time.Second
google_oauth_test.go:502 — t.Errorf("exp should be ~%s away...", accessTokenTTL)
login_test.go:250       — res.AccessTokenExpiresAt.Equal(h.now.Add(accessTokenTTL))
login_test.go:605       — res.AccessTokenExpiresAt.Equal(h.now.Add(accessTokenTTL))
```

Each becomes `auth.AccessTokenTTL`. The import `platform/auth` is already
present in files that need it; `google_oauth_test.go` may need the import
added if it doesn't already import `platform/auth` — check and add if
missing.

### Step 3 — verify no remaining references

```bash
# Should return zero matches after the edit
grep -rn "accessTokenTTL" backend/internal/domain/account/
```

## Verification

```bash
# Unit tests — the TTL assertions must still pass against the single constant
go test ./internal/domain/account/ -v

# Race + full suite
go test -race ./...

# Gate
make verify
```

The test assertions (`h.now.Add(auth.AccessTokenTTL)`) will pass unchanged
because the value is identical (15m) — this task is about removing the
second source of truth, not changing behavior.

## Out of scope

- Moving `refreshTokenTTL` to a shared location (no counterpart in `platform/auth/`; account-domain concern; defer)
- Any TTL value change (15m stays 15m)
- The Tier 0 paired rewrite pass (separate human gate)
