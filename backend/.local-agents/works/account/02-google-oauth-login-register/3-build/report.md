# Implementation Report — 02-google-oauth-login-register

> Ticket    : 02-google-oauth-login-register
> Feature   : Google OAuth login/register (account domain, second vertical slice)
> Date      : 2026-08-24
> Spec ref  : `docs/spec/1-account/features/02-google-oauth-login-register.md` (Fitur 1B)
> Techplan  : `.local-agents/works/account/02-google-oauth-login-register/2-plan/techplan.md`
> Tasks     : `.local-agents/works/account/02-google-oauth-login-register/2-plan/tasks/` (4 task files + manifest)

---

## 1. Summary

The account domain's second authentication path: Google OAuth via two shared
endpoints (`GET /auth/google/redirect` + `GET /auth/google/callback`)
handling three intents — `login` (full business rules implemented here),
`link` and `reauth` (mechanics implemented; intent-specific business rules
extend in tasks #05/#06). Includes a new `googleoauth` platform package
(code exchange + JWKS-backed id_token verification), token-issuance
primitives (`IssueTokens` — ES256 access + hashed refresh, so task #3 can
build rotation on top without depending back on this ticket), the
`refresh_tokens`/`user_logs` tables, cookie delivery of session tokens, an
in-memory reauth marker store, and handler-level conditional session
verification for link/reauth.

The defining security properties, all enforced and tested:

- **No-auto-merge** (R9/R10): no code path creates or attaches an identity
  from a Google email claim alone — the takeover defense.
- **CSRF/replay**: constant-time state comparison *before* any Google API
  call; nonce verified with the same discipline; replay (nonce mismatch) is
  distinguishable from forgery (generic verification failure).
- **Fenced paths untouched**: `platform/auth/` and `platform/crypto/` are
  consumed read-only only.

Executed by one agent session in dependency order (01 → 02 → 03 → 04), with
a human review stop between every task. No commits made (human's explicit
instruction); everything is working-tree changes.

---

## 2. Files changed

### New files (12)

| File | LoC | Task | Description |
|---|---|---|---|
| `migrations/000004_create_refresh_tokens.up.sql` | 15 | 01 | minimal `refresh_tokens` + unique `token_hash` + partial active index |
| `migrations/000004_create_refresh_tokens.down.sql` | 1 | 01 | drop table |
| `migrations/000005_create_user_logs.up.sql` | 8 | 01 | minimal append-only `user_logs` |
| `migrations/000005_create_user_logs.down.sql` | 1 | 01 | drop table |
| `internal/platform/googleoauth/client.go` | 352 | 02 | `AuthURL`, `ExchangeCode`, `VerifyIDToken` (RS256-only, iss/aud/exp, constant-time nonce), JWKS cache w/ refresh-on-miss, sanitized logging |
| `internal/platform/googleoauth/client_test.go` | 344 | 02 | forged-token table (sig/iss/aud/exp/HS256-confusion/missing-kid), nonce-mismatch distinction, timeout, body-leak checks, refresh-on-miss fetch counting |
| `internal/platform/googleoauth/helpers_test.go` | 144 | 02 | mock JWKS/token servers, RS256 signer, thread-safe log capture |
| `internal/domain/account/google_oauth.go` | 542 | 03 | `GoogleRedirect`, `GoogleCallback` (3-intent branching), `IssueTokens`, `CallbackResult`, oauth-state encode/decode |
| `internal/domain/account/google_oauth_test.go` | 697 | 03 | R1–R26 service slices incl. named no-auto-merge tests, 12-goroutine duplicate race, log-discipline capture |
| `internal/transport/http/cookie.go` | 91 | 04 | state cookie write/read/clear (R24 attrs) + auth-cookie pair (access Lax / refresh Strict) |
| `internal/transport/http/auth_google.go` | 251 | 04 | both handlers, inline ES256 verifier (R25), reauth marker store (sync.Map + sweeper) |
| `internal/transport/http/auth_google_test.go` | 437 | 04 | handler contract tests incl. named `TestGoogleRedirect_LinkReauthRequireAuth`, cookie-attribute assertions, all six error codes verbatim |

### Modified files (8)

| File | Task | Description |
|---|---|---|
| `go.mod` / `go.sum` | 01→02 | +`github.com/golang-jwt/jwt/v5 v5.3.1` (direct after Task 02 imported it) |
| `internal/domain/account/entity.go` | 03 | +`RefreshToken`, +`UserLog` structs |
| `internal/domain/account/repository.go` | 03 | +`InsertRefreshToken`, +`InsertUserLog` to interface |
| `internal/domain/account/repository_db.go` | 03 | both inserts implemented (goqu, prepared, tx-taking, `InsertAuthToken` pattern) |
| `internal/domain/account/service.go` | 03 | +`googleOAuth`/`authKeys`/`frontendURL` deps; `NewService` signature extended |
| `internal/domain/account/service_test.go` | 03 | fakeRepo grew the two insert methods (+ recorded slices); **no existing assertions changed** |
| `cmd/server/main.go` | 03+04 | requireEnv += GOOGLE_CLIENT_ID/SECRET/REDIRECT_URI, FRONTEND_URL; google client constructed; auth keys now captured; routes registered behind rate limiter; `cookieSecure = appEnv != "development"` |
| `.env` | 04 | `GOOGLE_REDIRECT_URI` corrected :8080→:8090 (matches APP_PORT/.env.example); `FRONTEND_URL=http://localhost:3000` added. **GOOGLE_CLIENT_ID/SECRET left empty — user must supply real values** (server fails fast until then) |

### Pre-existing changes (NOT this feature — flagged)

| File | Note |
|---|---|
| `frontend/AGENTS.md`, `frontend/.claude/*` | modified by another session before/during this one — frontend/ boundary respected, not touched from here (AGENTS.md §7) |

---

## 3. API endpoints delivered

| Method | Path | Handler | Success | Errors |
|---|---|---|---|---|
| `GET` | `/auth/google/redirect?intent={login\|link\|reauth}` | `GoogleRedirectHandler` | 302 → Google consent screen (+state cookie) | 400 invalid intent (R18), 401 link/reauth without session (R2), 500 internal |
| `GET` | `/auth/google/callback?code=...&state=...` | `GoogleCallbackHandler` | 302 → frontend (auth cookies set, or reauth marker) | 302 `?error={code}` — state_mismatch, nonce_mismatch, google_token_invalid, google_unavailable, google_email_conflict, google_link_conflict; 500 on internal failure |

Both mounted under the existing `RateLimit(rps, burst)` middleware via
`mux.Handle("/auth/", ...)`.

Error responses for non-redirect failures are RFC 9457 Problem Details;
redirect failures carry only the stable error-code vocabulary — never raw
errors, tokens, or internals.

---

## 4. Rule coverage (R1–R26)

Full task-level matrix lives in `2-plan/tasks/manifest.md`; named-test
status here. ✅ = test exists and passes.

| Rule | Named test(s) | Status |
|---|---|---|
| R1 (login redirect, no auth) | `TestGoogleRedirect_Login_NoSessionRequired` (svc) + `TestGoogleRedirect_LoginNeedsNoAuth` (handler) | ✅ |
| R2 (link/reauth 401 pre-redirect) | **`TestGoogleRedirect_LinkReauthRequireAuth`** (handler) + `TestGoogleRedirect_LinkReauthWithoutSessionRejectedDefensively` (svc) | ✅ |
| R3 (user_id into cookie) | `TestGoogleRedirect_LinkWithSessionEncodesUserID` | ✅ |
| R4 (state mismatch, no Google call) | `TestGoogleCallback_StateMismatch_NoGoogleCall` (zero-call assertion) | ✅ |
| R5 (nonce mismatch → nonce_mismatch) | `TestVerifyIDToken_NonceMismatchIsDistinguishable` (platform) + `TestGoogleCallback_NonceMismatchIsReplayNotForgery` (domain) | ✅ |
| R6 (timeout → google_unavailable) | `TestExchangeCode_TimeoutReturnsCleanError` + `TestGoogleCallback_GoogleUnavailableMappedCleanly` | ✅ |
| R7 (login existing identity) | `TestGoogleCallback_Login_ExistingGoogleIdentityIssuesTokens` | ✅ |
| R8 (login new user, tx) | `TestGoogleCallback_Login_NewUserCreatesVerifiedIdentity` | ✅ |
| R9 (**no-auto-merge login**) | **`TestGoogleCallback_NoAutoMerge_Login`** | ✅ |
| R10 (**no-auto-merge link**) | **`TestGoogleCallback_NoAutoMerge_Link`** | ✅ |
| R11 (link attach + audit) | `TestGoogleCallback_LinkSuccess_AttachesAndAudits` (action_type=account_linking asserted) | ✅ |
| R12 (reauth marker, no side effects) | `TestGoogleCallback_Reauth_NoSideEffects` (svc) + `TestGoogleCallback_ReauthSetsMarker` (handler) | ✅ |
| R13 (fixed redirect_uri) | `TestExchangeCode_SendsConfiguredParamsAndParsesResponse` | ✅ |
| R14 (verified_at=now at insert) | asserted in new-user + link-success tests | ✅ |
| R15 (concurrent duplicate clean failure) | `TestGoogleCallback_ConcurrentDuplicateRegistration_Race` (12 goroutines, `-race`) + `..._WinnerInvisibleYet` (deterministic) | ✅ |
| R16 (no secrets in logs) | `TestGoogleOAuth_LogsNeverCarrySecrets` (svc) + `TestGoogleCallback_LogsNeverCarryTokens` (handler) | ✅ |
| R17 (sanitized Google-client errors) | `TestExchangeCode_Non200_SanitizesErrorAndLog` (body-marker leak check) | ✅ |
| R18 (invalid intent 400) | `TestGoogleRedirect_InvalidIntentRejected` + `TestGoogleRedirect_InvalidIntent400` | ✅ |
| R19 (missing code/state) | `TestGoogleCallback_MissingParamsOrBadCookie_StateMismatch` (svc) + `TestGoogleCallback_MissingInputsStillHandled` (handler) | ✅ |
| R20 (expired/missing cookie) | same two tests (corrupt/empty cookie cases) | ✅ |
| R21 (JWKS refresh-on-miss) | `TestVerifyIDToken_JWKSRefreshOnMiss` (fetch-counted: exactly one refetch) | ✅ |
| R22 (explicit timeout, NewRequestWithContext) | `TestExchangeCode_TimeoutReturnsCleanError` + code-review check | ✅ |
| R23 (constant-time compare) | nonce half in `VerifyIDToken`, state half in `GoogleCallback` (code-review verifiable; behavior covered via R4/R5 tests) | ✅ |
| R24 (cookie attributes) | `TestGoogleRedirect_LoginNeedsNoAuth` asserts HttpOnly/Secure/Lax/MaxAge=600/Path exactly | ✅ |
| R25 (inline verify, platform/auth untouched) | `TestGoogleRedirect_LinkWithBearerTokenPassesUserID` + `TestGoogleRedirect_TamperedTokenRejected` (foreign-key rejection) | ✅ |
| R26 (forgery generic, only nonce special) | `TestVerifyIDToken_GenericFailuresAreNotNonceMismatch` (8-case table incl. HS256 confusion) | ✅ |

Not yet written (deferred, see §7): integration-tagged tests against real
Postgres for `InsertRefreshToken`/`InsertUserLog` round-trips.

---

## 5. Verification results

| Gate | Command | Result |
|---|---|---|
| Build / Vet / gofmt | `go build ./... && go vet ./... && gofmt -l internal/ cmd/` | ✅ clean |
| Platform race | `go test -race -count=1 ./internal/platform/...` | ✅ all packages pass (googleoauth 12 tests) |
| Domain race | `go test -race -count=1 ./internal/domain/account/` | ✅ pass (~68s; includes 12-goroutine OAuth duplicate race) |
| Transport race | `go test -race -count=1 ./internal/transport/http/` | ✅ pass (14 Google handler tests) |
| Full suite | `go test -race -count=1 ./...` | ✅ all green |
| Migration round-trip | `migrate up` → index/FK inspection → `down 2` → `up` against dev Postgres (:5435) | ✅ clean cycle; partial index predicate verified; R15 precondition (`ux_auth_identities_provider_identifier`) confirmed present |
| FK cascade | user delete removes refresh_tokens/user_logs rows; unknown-user FK rejected (rolled-back tx) | ✅ |
| **Live smoke** | server booted on :8090 (dummy Google creds): `login` → 302 consent URL + correct cookie attrs · unauthenticated `link` → 401 · `intent=evil` → 400 · bare callback → 302 `?error=state_mismatch` | ✅ all four boundary behaviors observed over HTTP |
| `staticcheck`/`gosec`/`gitleaks`/`govulncheck` | not run locally | ⚠️ full `make verify` required before merge |

---

## 6. Process deviations (flagged for audit trail)

All are implementation-level refinements within the techplan's stated
intent; none touch scope, spec, or fenced paths. Grouped for the techplan
owner to ratify or reject:

### 6.1 `NewService` gained an 8th parameter (`frontendURL`)

Techplan §10 enumerated params 6–7 (googleOAuth, authKeys). But §8's flow
has the **service** producing complete redirect URLs (app root, security
page, `?error=` targets), which requires the base URL — while §9 step 8
assigned FRONTEND_URL to the handler. The service-side reading won because
§8 and §10 agree twice (flow pseudocode + CallbackResult doc: "returns …
redirect URL"). Handlers stayed thin translators.

### 6.2 `CallbackResult` gained two fields beyond §10's struct literal

- `Reauth bool` — the transport layer must set the marker store only on
  successful reauth, and it cannot re-decode the cookie (the encoder is
  domain-unexported).
- `UserID uuid.UUID` — the marker is keyed by the session user id (R12);
  populated from the cookie binding where present.

### 6.3 Google handlers take a narrow interface, not `*account.Service`

Existing handlers take the concrete type (`RegisterHandler(svc
*account.Service)`). The two Google handlers take `googleOAuthService`
(two methods) instead: with the concrete type, handler tests would need the
real service wired to a *network-reaching* Google client whose endpoints
are hardcoded in `NewClient` — making the success-path cookie assertions
untestable. `*account.Service` satisfies the interface implicitly; main.go
wiring is unchanged in shape. Style deviation from sibling handlers,
flagged for consistency review.

### 6.4 Session-token transport for link/reauth was unspecified

The techplan says the redirect handler "verifies JWT inline" but never
states where the token comes from. Implemented: the access-token cookie
first (natural carrier for browser-initiated navigations — our own
`setAuthCookies` delivers there), falling back to `Authorization: Bearer`
for non-browser clients/tests.

### 6.5 Task-file drift vs execution (cookie.go)

Task 04's file sketched `setOAuthStateCookie(w, state, nonce, intent,
userID)` encoding in transport; actual execution keeps encoding in the
service (per §8 flow: `GoogleRedirect` returns `cookieValue`) and transport
only wraps it in an `http.Cookie`. Same responsibilities, cleaner boundary;
noted so the task file isn't mistaken for a literal build log.

---

## 7. Risk note

- **Assumptions made:**
  - New Google users get their **email as display name** — our Claims
    carry no separate name claim; users.name is NOT NULL. Cosmetic, fixable
    later if the frontend surfaces a rename flow.
  - Security page path = `/account/security` (mirrors the account-security
    endpoint namespace). Frontend route confirmation pending.
  - A race loss where the winner's identity is not yet visible to re-lookup
    surfaces as `google_email_conflict` — closest recoverable signal;
    mislabels a transient condition as a conflict (recoverable by retry).
  - Link-to-already-linked identity = idempotent success, no second audit
    entry written.
  - `email_verified` id_token claim is **not enforced** — techplan R26
    lists only sig/iss/aud/exp/nonce; adding the check would be a new
    security decision needing human sign-off.
- **Edge cases intentionally NOT handled (and why):**
  - Concurrent JWKS refreshes may duplicate fetches (last-write-wins under
    mutex; harmless).
  - No backoff on JWKS fetch failure — verification fails cleanly and the
    next call retries; backoff adds latency the fail-fast design avoids.
  - Reauth markers lost on restart (in-memory sync.Map per techplan §5
    decision; acceptable for 5-min TTL).
  - Refresh-token rotation/reuse detection deliberately absent (task #3
    builds it on these primitives; INV-account-03/04).
- **Concurrency assumptions:** `Service` remains stateless beyond injected
  goroutine-safe deps; the googleoauth JWKS cache is mutex-guarded and
  proven under `-race` (refresh-on-miss test); the reauth marker store uses
  sync.Map with lazy expiry + periodic sweep. Proven by
  `TestGoogleCallback_ConcurrentDuplicateRegistration_Race` (12 goroutines)
  and the full `-race` suite.
- **What is not tested, and why:**
  - Integration-tagged Postgres round-trips for the two new inserts —
    pattern-identical to `InsertAuthToken` (goqu prepared insert); the
    migration round-trip + FK cascade were verified live against dev
    Postgres. Should be added to the `//go:build integration` suite before
    PR merge.
  - Live end-to-end OAuth with real Google credentials (requires the user's
    GOOGLE_CLIENT_ID/SECRET; boundary behaviors were smoke-tested over HTTP
    with dummies).
  - `email_verified` enforcement (assumption above).
  - staticcheck/gosec/govulncheck/gitleaks — not installed locally; `make
    verify` must gate the PR.

---

## 8. Open items status (techplan §14)

| # | Item | Status | Note |
|---|---|---|---|
| Active 1 | Caddy routing for `/auth/*` | ⚠️ unchanged | Dev workaround (direct :8090) works — smoke-tested; `.env` corrected to :8090. Non-dev fix remains a root-level session. |
| Active 2 | Error-code set v2 + `account_linking` literal — frontend sign-off | ⚠️ unchanged | Backend pins both as specified; six codes surface verbatim (tested). Frontend must confirm before their implementation. |
| Resolved 1–7 | Redis/user_logs ownership/Tier-0 classify/TTL/error codes/FRONTEND_URL/cookie-vs-memory | ✅ honored | All implemented per recorded decisions; nothing reopened. |

---

## 9. Database migrations

Applied against dev Postgres (`localhost:5435/kencleng`) and verified:

```
4/u create_refresh_tokens (58ms)
5/u create_user_logs (82ms)
```

Schema verified via `\d`: all columns, unique `token_hash` index, partial
`ix_refresh_tokens_active` predicate, FK cascades both directions,
`user_logs` minimal shape (no REVOKE/CHECK — deferred to #08 as decided).

Rollback: `make migrate-down` rolls all the way; or `migrate down 2` for
just this ticket's pair.

---

## 10. How to run

```bash
# 0. Fill in backend/.env: GOOGLE_CLIENT_ID / GOOGLE_CLIENT_SECRET
#    (Google Cloud Console; redirect URI http://localhost:8090/auth/google/callback)

set -a; . .env; set +a
make migrate-up          # 000004 + 000005 if not applied

go run ./cmd/server      # listens on :8090

# Manual flow (dev):
#   open http://localhost:8090/auth/google/redirect?intent=login
#   → consent screen → callback → 302 to $FRONTEND_URL with auth cookies

# Tests
go test -count=1 ./...
go test -race -count=1 ./...

# Full gate (before merge)
make verify
```

---

## 11. Next steps (workflow hand-off)

1. **Human review** of §6 deviations (four small items) and §7 assumptions
   (esp. `email_verified` enforcement decision).
2. **Dual-model code review** (mandatory per model-routing Complex tier):
   GLM 5.2 (max) + DeepSeek V4 Pro parallel, diffed manually — review the
   whole working-tree diff; no commits exist yet.
3. Integration tests for the two new repo methods (§7 gap).
4. `make verify` once lint tooling available → PR with risk note referencing
   `docs/spec/1-account/features/02-google-oauth-login-register.md`.
