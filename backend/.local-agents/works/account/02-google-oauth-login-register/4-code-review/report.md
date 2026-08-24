# Code Review — 02-google-oauth-login-register

> Ticket    : 02-google-oauth-login-register
> Feature   : Google OAuth login/register (account domain, Fitur 1B)
> Reviewer  : GLM 5.2 (max)
> Date      : 2026-08-24
> Diff ref  : commit `ef428be` (backend scope; `frontend/` excluded by AGENTS.md §7)
> Sources   : `2-plan/` (techplan + task manifest), `3-build/` (report + working-tree changes)
> Conventions: `backend/AGENTS.md`, `backend/README.md`, root `AGENTS.md`

---

## Review scope

Files actually read for this review (the working-tree diff is empty —
the changes were committed in `ef428be`; the review is against that
commit's `backend/` subset, excluding the unrelated `frontend/` changes
in the same commit):

New (12): `migrations/000004_*`, `migrations/000005_*`,
`internal/platform/googleoauth/{client.go,client_test.go,helpers_test.go}`,
`internal/domain/account/{google_oauth.go,google_oauth_test.go}`,
`internal/transport/http/{cookie.go,auth_google.go,auth_google_test.go}`.

Modified (8): `go.mod`/`go.sum`, `internal/domain/account/{entity.go,repository.go,repository_db.go,service.go,service_test.go}`,
`cmd/server/main.go`, `internal/transport/http/auth_verify_email_test.go`, `.env`.

Convention files read first: `backend/AGENTS.md` (§1-7), `backend/README.md`,
root `AGENTS.md` (§1-8), existing handlers (`auth_register.go`,
`auth_verify_email.go`, `errors.go`, `middleware.go`), and the
`api/openapi/account.yaml` contract for the two new endpoints.

---

## 1. Safety

> Looked specifically for: nil/null dereference on legitimately-absent
> data, race conditions in concurrent code, errors silently swallowed,
> missing cancellation/timeout propagation, resources not released on
> early-return paths.

No blocking safety findings. All five safety checklist categories pass:

- **Nil dereference**: every `*uuid.UUID` and `*time.Time` dereference is
  nil-checked before use — `st.UserID` in `callbackLink`
  (`google_oauth.go:412`), the reauth branch (`google_oauth.go:279`),
  and `IssueTokens` guards `s.authKeys == nil` / `s.authKeys.Private == nil`
  (`google_oauth.go:484`). `VerifyIDToken` defensively handles
  `claims.Audience`/`claims.ExpiresAt` being empty/nil before copying
  (`client.go:233-240`).
- **Concurrency**: JWKS cache is mutex-guarded
  (`client.go:114` `sync.Mutex` around `jwksCache`/`jwksFetchedAt`);
  reauth markers use `sync.Map` (`auth_google.go:101`); `Client` is
  constructed once at startup and reused (AGENTS.md §2 — no per-request
  `http.Client`). Proven under `-race` by
  `TestGoogleCallback_ConcurrentDuplicateRegistration_Race` (12 goroutines).
- **Error swallowing**: every error is either returned (wrapped with
  `%w`) or logged with a sanitized category before a clean failure
  result — no empty-catch / discarded-`err` paths.
- **Context/cancellation**: all outbound HTTP calls build the request
  with `http.NewRequestWithContext(ctx, ...)` (`client.go:171,282`); the
  `http.Client` carries a 10s timeout as a backstop (R22).
- **Resource leaks**: every `resp.Body` is `defer ...Close()` and drained
  on non-200 (`client.go:182,185,291,294`). Every `BeginTx` has a
  `committed`-guarded deferred `Rollback` on every return path
  (`google_oauth.go:357-362,435-440,528-533`); `registerGoogleUser`
  additionally rolls back explicitly before the winner re-lookup
  (`google_oauth.go:382`) and is careful not to double-roll (the defer
  checks `committed`, which is only set after a successful `Commit` —
  the explicit `Rollback` path falls through to `return` without setting
  it, so the deferred no-op is harmless).

### Finding S1 — Concurrent link race surfaces as 500 instead of clean error

| Field | Value |
|---|---|
| Location | `internal/domain/account/google_oauth.go:450` (`callbackLink`, the `InsertAuthIdentity` call) |
| Why it matters | `registerGoogleUser:378-394` explicitly handles `isUniqueViolation(insertErr)` to convert a concurrent-duplicate insert into a clean `google_email_conflict` / winner-re-lookup path (R15). `callbackLink` performs the same `InsertAuthIdentity` against the same unique index `(provider_type, identifier_hash)` but does **not** check `isUniqueViolation` — a concurrent link operation for the same Google email (two browser tabs, same user, or two sessions racing) yields a raw unique-violation error that bubbles up to `GoogleCallbackHandler` and is mapped to a 500 Problem Details instead of a 302 with `google_link_conflict`. **Not security-critical**: the unique index still prevents data corruption, and the user can retry successfully. The defect is the *error category*: an internal-500 where a clean, user-recoverable 302 was intended. The report's §7 flags concurrent-duplicate handling as R15-covered, but R15's test (`TestGoogleCallback_ConcurrentDuplicateRegistration_Race`) exercises the login path, not link. |
| Suggested fix | Mirror the `registerGoogleUser` pattern: check `isUniqueViolation(err)` after the link insert; on hit, roll back the tx, re-lookup the winner identity, and either return `google_link_conflict` (if the winner belongs to a different user) or treat as idempotent success (if the winner is the same session user). At minimum, return `s.failResult(true, intentLink, errGoogleLinkConflict), nil` so the error category is correct. Add a `-race` test mirroring the login duplicate-race test for the link path. |

---

## 2. Quality

> Looked for: duplicated logic, unclear signatures, naming vs. behavior,
> missing observability at decision points.

### Finding Q1 — Duplicated TTL constants across domain and transport

| Field | Value |
|---|---|
| Location | `internal/domain/account/google_oauth.go:48-52` vs `internal/transport/http/auth_google.go:22-24` |
| Why it matters | `stateCookieTTL` / `stateCookieMaxAge`, `accessTokenTTL` / `accessTokenCookieTTL`, `refreshTokenTTL` / `refreshTokenCookieTTL` are defined in parallel in two packages, each with a comment saying they "must match" the other side. There is no compile-time assertion that they do. If one side changes (e.g. someone tightens the access-token TTL in task #3 and edits only the domain constant), the cookie MaxAge and the JWT `exp` silently drift — the browser would keep delivering an access cookie whose JWT the verifier rejects, or vice versa. The pairing is load-bearing for correctness, not just style. |
| Suggested fix | Either define the three lifetimes once (in the domain package, exported) and reference them from transport — or add a compile-time assertion in transport: `var _ = [0]struct{}{}[stateCookieMaxAge - stateCookieTTL]` (fails to compile if they diverge). Same for the access/refresh pair. |

### Finding Q2 — `failResult` condition relies on operator precedence

| Field | Value |
|---|---|
| Location | `internal/domain/account/google_oauth.go:193` |
| Why it matters | `if intentKnown && intent == intentLink || intentKnown && intent == intentReauth` is correct (`&&` binds tighter than `||` in Go) but reads as a run-on; a future edit that drops one `intentKnown` thinking it's implicit will introduce a bug. Pure readability, no behavior defect. |
| Suggested fix | `if intentKnown && (intent == intentLink || intent == intentReauth)` — factored form, same semantics, harder to misedit. |

### Finding Q3 — `sessionToken` Bearer extraction is case-sensitive

| Field | Value |
|---|---|
| Location | `internal/transport/http/auth_google.go:87` |
| Why it matters | `auth[:len(prefix)] == prefix` only matches the canonical `Bearer ` (capital B). RFC 6750 §2.1 specifies the authentication scheme is case-insensitive (`Bearer`/`bearer`/`BEARER` all valid). Most clients send the canonical form, so this is a conformance nit rather than a live bug — but non-browser clients/tests using lowercase would be silently rejected from the link/reauth flows, falling through to 401. |
| Suggested fix | Compare case-insensitively: `if len(auth) > len(prefix) && strings.EqualFold(auth[:len(prefix)], prefix) { ... }`. Note this matches the handler's test fallback path (§6.4) where `Authorization: Bearer` is used for non-browser clients. |

### Finding Q4 — Error code not URL-encoded in redirect target

| Field | Value |
|---|---|
| Location | `internal/domain/account/google_oauth.go:198` (`failResult`: `base + "?error=" + code`) |
| Why it matters | All current error-code constants are `[a-z_]+` (safe in a query string), so this works today. It would break if a future code ever contained `&`, `=`, `+`, `%`, or a space — the frontend would receive a mangled `?error=` value. Minor robustness gap; the code set is pinned by techplan §14 Active item 2 (frontend sign-off pending), so it's effectively closed — but the function shouldn't rely on that closure for its own correctness. |
| Suggested fix | `url.Values{"error": {code}}.Encode()` (or `url.QueryEscape(code)`) when constructing the redirect URL. Cheap, future-proof. |

---

## 3. Stack-Specific Best Practices

> Matched `best-practices/index.md` trigger keywords against the diff's
> technology: Go (authn, JWT, HTTP client, goroutines, secrets), REST API
> (CSRF/cookies, anti-enumeration), PostgreSQL (audit log, migrations).
> Opened only the matching files:
> `go/jwt-and-token-lifecycle.md`, `go/http-client-and-transport.md`,
> `go/secrets-and-sensitive-logging.md`, `go/goroutine-lifecycle.md`,
> `restapi/csrf-and-cookie-security.md`, `restapi/anti-enumeration.md`,
> `postgresql/audit-log-design.md`, `postgresql/migrations-safety.md`.

### Finding B1 — `user_logs` lacks DB-level `REVOKE UPDATE, DELETE`

| Field | Value |
|---|---|
| Source | `postgresql/audit-log-design.md` — checklist: "UPDATE/DELETE on the audit table is REVOKEd at the DB level for the application role." |
| Location | `migrations/000005_create_user_logs.up.sql` |
| Why it matters | The audit-log best practice is that immutability must be enforced at the DB level, not just by application convention — the app layer can be bypassed by direct DB access or a future code path that forgets the rule. Migration `000005` creates `user_logs` with no `REVOKE` and no trigger rejecting modification. **Properly deferred, not a gap introduced by this change**: techplan §14 Resolved item 2 and report §7 explicitly defer the `REVOKE UPDATE, DELETE` constraint (INV-account-11) to task #08, which owns the full `user_logs` design (all `action_type` values, trigger points across #05/#06/#08). Recording here for completeness and so a future reviewer of task #08 sees the carry-forward. |
| Suggested fix | None in this ticket. Task #08 must add `REVOKE UPDATE, DELETE ON user_logs FROM kencleng_app` (or equivalent) via a later migration. The deferral is tracked; do not let it close silently. |

### All other best-practice checklists pass

- **`go/jwt-and-token-lifecycle.md`**: separate keys/algorithms per
  token purpose — ES256 for the app's own access tokens
  (`google_oauth.go:494`), RS256-only for Google id_token verification
  (`client.go:210` `jwt.WithValidMethods([]string{"RS256"})`), and the
  refresh token is a random opaque value (not a JWT) whose SHA-256 hash
  is stored (`google_oauth.go:499-511`). Refresh rotation + reuse
  detection deliberately deferred to task #3 (techplan §2 out-of-scope;
  primitives only here). Token storage in cookies is a deliberate
  XSS-vs-CSRF trade-off (techplan §14 Resolved item 7). Expiry is
  proportional to risk: access 15min, refresh 30d, state cookie 10min,
  reauth marker 5min. ✓
- **`go/http-client-and-transport.md`**: `http.Client{Timeout: 10s}`
  constructed once in `NewClient` and reused across all calls
  (`client.go:131`) — no per-request client. `http.NewRequestWithContext`
  on both token-exchange and JWKS fetch. Response bodies are fully read
  (drained via `io.Copy(io.Discard, io.LimitReader(...))` on non-200)
  and `defer Close()`d on every path. ✓
- **`go/secrets-and-sensitive-logging.md`**: Google-client errors are
  reduced to a sanitized category via `categorizeError` before any
  `log.Printf` (`client.go:179,288`); raw `err` is never logged in the
  OAuth path. The service layer logs `category=sanitized` for exchange
  failures rather than the raw error string (`google_oauth.go:252`).
  No `%+v` struct dumps anywhere in the diff. ✓
- **`go/goroutine-lifecycle.md`**: the `init()` sweeper goroutine in
  `auth_google.go:124-142` is a process-lifetime goroutine (ticker,
  `for range t.C`) — same shape as the rate-limiter sweeper in
  `middleware.go:29`. It does not capture per-request state and is not
  a leak. The report §6.3 / §7 call out the pattern inheritance. ✓
- **`restapi/csrf-and-cookie-security.md`**: state cookie is
  `SameSite=Lax` (`cookie.go:30`) — correct, because the redirect back
  from Google is a cross-site navigation and `Strict` would drop the
  cookie. The `state` parameter itself is the second CSRF mitigation
  (double-submit-equivalent). Refresh cookie is `SameSite=Strict`
  (`cookie.go:89`). Both have `HttpOnly` + `Secure` (non-dev). The
  refresh *endpoint* (where the refresh cookie is consumed) is out of
  scope for this ticket — task #3 owns it and must add the custom-header
  check on top of `SameSite=Strict`. ✓
- **`restapi/anti-enumeration.md`**: `subtle.ConstantTimeCompare` for
  both state (`google_oauth.go:243`) and nonce (`client.go:229`) — no
  `==` on secret-shaped values. The six callback error codes are a
  fixed vocabulary surfaced via `?error={code}`, identical regardless
  of internal branch — no internal detail leaks to the redirect target.
  The login redirect handler's R9 (no-auto-merge) path returns the same
  shape as a normal error, preventing enumeration of "email already
  claimed by email_password." ✓
- **`postgresql/migrations-safety.md`**: both migrations are additive
  (new tables), each with a matching down migration (`DROP TABLE IF
  EXISTS`). No existing column altered, no backfill, no data migration.
  Reversibility verified by the report §5 migration round-trip
  (`migrate up` → `down 2` → `up`). ✓

---

## 4. Consistency

> Read first: `backend/AGENTS.md` (§1-7), `backend/README.md`, root
> `AGENTS.md` (§1-8). Then compared against existing handlers
> (`auth_register.go`, `auth_verify_email.go`, `errors.go`,
> `middleware.go`), the repository adapter, and the service.

### Finding C1 — Handler signature deviation (narrow interface vs concrete `*account.Service`)

| Field | Value |
|---|---|
| Location | `internal/transport/http/auth_google.go:34-37` (`googleOAuthService` interface) vs `auth_register.go:22` / `auth_verify_email.go:21` (`*account.Service` concrete) |
| Convention | Sibling handlers in the same package take the concrete `*account.Service` directly. Established in `auth_register.go:22` (`RegisterHandler(svc *account.Service)`) and `auth_verify_email.go:21` (`VerifyEmailHandler(svc *account.Service)`). |
| Why it matters | This is a real style deviation from the established pattern in the same file neighborhood. The report (§6.3) flags it explicitly and justifies it: with the concrete `*account.Service`, handler tests would need the real service wired to a `*network-reaching* googleoauth.Client` whose endpoints are hardcoded in `NewClient` — making the success-path cookie assertions untestable. The narrow interface (`googleOAuthService`, two methods) lets tests inject a stub. `*account.Service` satisfies the interface implicitly; `main.go` wiring is unchanged in shape. **Justified deviation, not a violation**: the rationale is concrete (testability of cookie attributes, which R24's exact-attribute assertions depend on), the alternative (concrete type + network-reaching test client) is worse, and task #3 may later extract a shared helper that regularizes the pattern. Recording here so a future consistency pass can see it was deliberate, not accidental. |
| Suggested fix | None in this ticket. If task #3 extracts a shared session-verification helper, consider whether the interface-vs-concrete question should be settled package-wide at that point. |

### Finding C2 — Sweeper goroutine placement (`init()` vs function-body)

| Field | Value |
|---|---|
| Location | `internal/transport/http/auth_google.go:124-142` (`init()`) vs `internal/transport/http/middleware.go:29-43` (inside `RateLimit`) |
| Convention | The existing rate-limiter sweeper goroutine starts inside `RateLimit()` — i.e., only when the middleware is actually wired. The reauth-marker sweeper starts in `init()`, i.e., unconditionally on package import. |
| Why it matters | The `init()` approach starts a goroutine even in tests that never exercise the Google handlers (e.g. `auth_verify_email_test.go` imports the package and pays for the goroutine). Harmless in practice — the goroutine is lightweight, the ticker fires once a minute, and the process exits after tests — but it's a small inconsistency with the sibling pattern. Not a violation of any written AGENTS.md rule; AGENTS.md §2 only mandates Go 1.22+ pattern routing and `net/http`, not goroutine placement. |
| Suggested fix | None required. If re-touched, consider lazily starting the sweeper on first `SetReauthMarker` call (sync.Once) to match `RateLimit`'s wiring-gated shape — but the current form is defensible for a long-lived server. |

### All other conventions followed

- **Error wrapping** (AGENTS.md §2): every error path uses
  `fmt.Errorf("...: %w", err)` — chain preserved, nothing discarded.
  ✓
- **SQL via goqu** (AGENTS.md §2, root golden rule): `InsertRefreshToken`
  and `InsertUserLog` in `repository_db.go:156-198` use
  `pgDialect.Insert(...).Rows(...).Prepared(true)` — parameterized, no
  string concatenation. ✓
- **Doc comments on exported functions/types** (AGENTS.md §2): every
  exported function and type in the new files has a doc comment
  (`Client`, `NewClient`, `AuthURL`, `ExchangeCode`, `VerifyIDToken`,
  `ErrNonceMismatch`, `TokenResponse`, `Claims`, `GoogleTokenVerifier`,
  `SetReauthMarker`, `CheckReauthMarker`, `GoogleRedirectHandler`,
  `GoogleCallbackHandler`, `CallbackResult`, `IssueTokens`,
  `GoogleRedirect`, `GoogleCallback`). ✓
- **Go 1.22+ pattern routing** (AGENTS.md §2): `main.go:139-142`
  registers `GET /auth/google/redirect` and `GET /auth/google/callback`
  via `mux.HandleFunc("GET /...", ...)`. ✓
- **Problem Details for non-redirect errors** (AGENTS.md §2): the
  redirect handler's 400/401/500 paths use `WriteProblem`
  (`auth_google.go:167,184,193`); callback's 500 path uses `WriteProblem`
  (`auth_google.go:227`). Redirect-leg failures carry only the stable
  `?error={code}` vocabulary — no internals leak. ✓
- **Explicit authz check at handler boundary** (root AGENTS.md §2):
  the link/reauth session check is visible and testable at the handler
  (`auth_google.go:162-171`), not hidden behind a query filter. Named
  test `TestGoogleRedirect_LinkReauthRequireAuth` proves it. ✓
- **PII encryption pattern** (root AGENTS.md §2): the new Google
  identity inserts reuse the existing `InsertAuthIdentity` adapter,
  which encrypts `identifier` (ciphertext) + HMAC `identifier_hash` at
  the storage boundary (`repository_db.go:88-126`). No new PII pattern
  invented. ✓
- **No secrets/tokens in logs** (root AGENTS.md §2): verified by
  `TestGoogleOAuth_LogsNeverCarrySecrets` (svc) and
  `TestGoogleCallback_LogsNeverCarryTokens` (handler) — log capture
  assertions scan for state/nonce/code/id_token/access/refresh values.
  ✓
- **Sentinel error naming**: `ErrInvalidIntent`, `ErrMissingSession`
  follow the existing `ErrValidation`/`ErrTokenExpired`/`ErrTokenNotFound`
  pattern (`service.go:30-41` vs `google_oauth.go:66-69`). ✓
- **Constant naming**: unexported camelCase (`intentLogin`,
  `errStateMismatch`, `actionAccountLinking`, `stateCookieTTL`)
  consistent with `providerEmailPassword`, `purposeEmailVerify`,
  `tokenTTL`. ✓
- **Logging format**: `log.Printf("account: ...")` and
  `log.Printf("transport: ...")` prefixes match existing handlers
  and service. ✓
- **Migration naming**: `000004_create_refresh_tokens`,
  `000005_create_user_logs` — sequential, matches `000001`/`000002`/
  `000003` shape. ✓
- **Domain package isolation** (AGENTS.md §1): `internal/domain/account`
  imports only `platform/*` packages, no other domain. ✓
- **`service_test.go` changes are additive only** (root AGENTS.md §4):
  the fakeRepo grew two new methods (`InsertRefreshToken`,
  `InsertUserLog`) and two new recorded slices — **no existing
  assertions changed**. Verified by reading the diff. ✓
- **`-race` test run** (AGENTS.md §3): the account domain test suite
  is run with `-race` (report §5; `TestGoogleCallback_ConcurrentDuplicateRegistration_Race`
  is the named race test). ✓
- **Fenced paths untouched** (root AGENTS.md §3): `platform/auth/` and
  `platform/crypto/` are consumed read-only (`auth.Load`, `crypto.HMAC`,
  `crypto.Encrypt`); neither package's files were modified. Inline ES256
  verification in the handler avoids `platform/auth/` entirely. ✓
- **Spec authority respected** (root AGENTS.md §1, §4): no
  `docs/spec/*.md` or approved test was edited to match code; the
  techplan's §14 open items are carried forward faithfully in the
  report §8. ✓
- **OpenAPI contract** (`api/openapi/account.yaml:308-363`): the two
  endpoints, the `intent` enum `[login, link, reauth]`, the `code`/`state`
  callback params, and the 302-redirect-with-error-query-param shape all
  match the spec. The six `?error={code}` literals are a backend-pinned
  vocabulary pending frontend sign-off (techplan §14 Active item 2) —
  not a drift from the spec, which is silent on the exact code strings.
  ✓

---

## Verdict

**Approve with minor comments.**

No blocking findings. The implementation is security-correct on the
three top-severity properties (no-auto-merge R9/R10, CSRF/replay via
constant-time state+nonce, fenced paths untouched) and consistent with
the repo's own conventions on every written AGENTS.md rule. The
remaining items are quality/robustness improvements, none of which
change the security posture.

### Optional — should address before PR merge

- **S1** (Safety): `callbackLink` should handle `isUniqueViolation` the
  same way `registerGoogleUser` does, so a concurrent link race returns
  `google_link_conflict` (302) instead of a 500 Problem Details. Not
  security-critical (unique index prevents corruption; retry recovers),
  but the error category is wrong and R15's test coverage doesn't reach
  the link path.
- **Q1** (Quality): de-duplicate the three TTL constant pairs across
  domain/transport, or add a compile-time assertion that they match —
  the pairing is load-bearing for cookie/JWT-expiry consistency.

### Optional — nice-to-have

- **Q2**: factor `failResult`'s `intentKnown && ... || intentKnown && ...`
  condition for readability.
- **Q3**: case-insensitive Bearer prefix match in `sessionToken`
  (RFC 6750 conformance).
- **Q4**: `url.QueryEscape` the error code in `failResult`'s redirect
  construction (future-proof against non-`[a-z_]` codes).

### Acknowledged gaps (properly flagged in build report §7, not findings)

- `email_verified` id_token claim not enforced — assumption pending
  human sign-off (§7 assumptions).
- `make verify` (staticcheck/gosec/govulncheck/gitleaks) not run locally
  — required before merge (§5 verification results, ⚠️ row).
- Integration-tagged Postgres round-trip tests for `InsertRefreshToken`/
  `InsertUserLog` not yet written — should be added to the
  `//go:build integration` suite before PR (§7 "What is not tested").
- `user_logs` `REVOKE UPDATE, DELETE` constraint deferred to task #08
  (B1 above) — tracked, do not let it close silently.
- Dual-model code review (DeepSeek V4 Pro in parallel with this GLM
  pass) still owed per the Complex-tier model-routing rule — this report
  is the GLM half only.
