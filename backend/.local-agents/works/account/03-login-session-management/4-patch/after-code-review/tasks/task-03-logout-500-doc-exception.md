# Task 03: Document logout 500-on-infra exception or align contract

> Back-reference : `4-code-review/report.md` §2 (Q3)
> Priority       : Optional (non-blocking, doc-only)
> Model          : DeepSeek V4 Pro — doc wording precision
> Depends on     : none

## Objective

R16 and techplan §3 frame `POST /auth/logout` as "always 204 / none
documented." The idempotent cases (no cookie, not-found, already-revoked)
do return 204. But `LogoutHandler` maps a genuine DB error (e.g. the
guarded `RevokeRefreshTokenByHash` UPDATE fails) to a Problem 500 via
`MapServiceError`. That's defensible — masking infra failures as 204 would
hide real problems — but the literal contract wording disagrees. Close the
gap with a one-line doc note, or align the contract.

**Recommended choice: document the exception.** Returning 500 on infra
failure is the correct behavior (don't mask real errors); the contract
wording just needs to acknowledge it.

## ⚠️ Human decision required

This task defaults to **documenting the exception** (code unchanged, doc
added). If Anhar prefers to align the contract wording instead (e.g. amend
the openapi/logout response definition to list 500), that's a spec edit
under AGENTS.md §4 authority — requires explicit human approval and should
not be agent-applied silently.

## Files

| File | Change |
|---|---|
| `backend/internal/transport/http/auth_login.go` | Add one-line doc note to `LogoutHandler` (option a) |

## Option (a) — document the exception (RECOMMENDED)

### Doc change (`auth_login.go` `LogoutHandler`)

Current doc:

```go
// LogoutHandler handles POST /auth/logout idempotently: revokes the
// presented refresh token when present, ALWAYS clears the cookie, ALWAYS
// answers 204 — an absent or already-dead cookie is not an error condition
// (R16).
```

Replace with:

```go
// LogoutHandler handles POST /auth/logout idempotently: revokes the
// presented refresh token when present, ALWAYS clears the cookie, and
// answers 204 for every idempotent case — an absent or already-dead
// cookie is not an error condition (R16). The sole exception is a
// genuine infrastructure failure (e.g. the revoke UPDATE fails at the DB
// level): that surfaces as a 500 Problem Details response rather than a
// masked 204, so real outages are not hidden behind a success code.
```

No code change. No test change (the existing `TestLogout_*` tests cover the
idempotent paths; a 500-on-DB-failure test would require a failing repo
injection similar to task-01's `failingAttemptRepo` — optional, not
required for this doc-only fix).

## Option (b) — align contract wording (ALTERNATIVE, requires human spec edit)

If Anhar prefers the openapi/feature-spec to explicitly list the 500 case,
this becomes a spec edit:

1. `api/openapi/account.yaml` — add a `500` response to the `POST /auth/logout` operation (reference the existing generic `Problem` schema).
2. `docs/spec/1-account/features/03-login-session-management.md` — update the R16 wording to acknowledge the infra-failure exception.
3. Regenerate the bundle (`api/openapi.yaml`) via the redocly script per `api/README.md`.

**This option must not be agent-applied without explicit human approval**
(AGENTS.md §4: spec edits go through the human; an agent must not edit a
`docs/spec/*.md` to make it match code). Flag it for Anhar and stop.

## Verification

```bash
# Option a: doc-only — just confirm it builds and existing tests pass
go build ./...
go test ./internal/transport/http/ -run TestLogout -v

# Gate
make verify
```

## Out of scope

- Changing the logout error-handling behavior (500-on-infra is correct)
- The Tier 0 paired rewrite pass (separate human gate)
- task-01, task-02 (independent)
