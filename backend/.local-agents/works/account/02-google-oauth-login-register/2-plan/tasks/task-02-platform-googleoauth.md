# Task 02 — Platform Package: `internal/platform/googleoauth/`

> Ticket    : 02-google-oauth-login-register
> Sub-task  : 2 of 4
> Axis      : Dependency/sequence chain (primary) + component/layer alignment
> Status    : Ready (after Task 01 — jwt dependency must be in go.mod)
> Back-ref  : `../techplan.md` — "Tech Plan: Google OAuth Login/Register" (originating contract techplan; cross-check high-level decisions there whenever needed)

---

## 1. Scope

One new shared-infrastructure package: a thin Google OAuth client with no
business rules of its own (platform = shared infra, per backend/AGENTS.md §1
— it must not know what a "campaign" or "account" is). It does three things:
build the consent-screen URL, exchange an authorization code for tokens, and
verify the returned id_token.

**In scope:**
- `internal/platform/googleoauth/client.go` — Client struct, NewClient,
  AuthURL, ExchangeCode, VerifyIDToken, Claims, ErrNonceMismatch
- `internal/platform/googleoauth/client_test.go` — unit tests with mocked
  Google HTTP endpoints

**Out of scope:**
- Any branching on intent (login/link/reauth) — Task 03's service decides;
  this package only reports verification outcomes
- Cookie handling, handler logic, route registration — Task 04
- Changes to `internal/platform/auth/` or `internal/platform/crypto/` — both
  are Tier 0 fenced paths (root AGENTS.md §3); this package is separate from
  and touches neither

## 2. Dependencies

- **Hard deps:** Task 01 (`github.com/golang-jwt/jwt/v5` in go.mod)
- **Soft deps:** none
- **Blocks:** Task 03 (service holds `*googleoauth.Client`; fake client for
  tests mirrors this package's interface)

## 3. Files

| File | Change Type |
|---|---|
| `backend/internal/platform/googleoauth/client.go` | New |
| `backend/internal/platform/googleoauth/client_test.go` | New |

## 4. Why this design (techplan §5 Decision Log rows — do not relitigate)

| Option considered | Why rejected/accepted |
|---|---|
| **A. coreos/go-oidc for OAuth client** | Rejected — full OIDC library is overkill for two HTTP calls + one JWT verification. |
| **B. golang-jwt/jwt/v5 + manual JWKS fetch (chosen)** | Chosen — new dependency but lightweight. ~40 lines for JWKS keyfunc with cache-on-miss. Token exchange is single http.Post. More control than go-oidc, fewer abstractions. |
| **C. Raw net/http + standard library** | Rejected — maximum control but most code to write and maintain. JWT verification without a library is error-prone. |

## 5. Implementation detail

### `backend/internal/platform/googleoauth/client.go` (new)

Per techplan §9 step 4 + §10:

- `Client` struct: `httpClient` (**10s timeout**, constructed once in
  NewClient), `clientID`, `clientSecret`, `redirectURI`, JWKS cache.
- `NewClient(clientID, clientSecret, redirectURI string) *Client`
- `AuthURL(state, nonce string) string` — build the Google consent URL from
  configured client_id / redirect_uri plus scope, state, nonce. Scope must
  include `openid email profile` — the callback flow reads the `email`
  claim from the id_token (techplan §8 business flow). The `nonce` param is
  what makes Google embed it into the id_token; the `state` param is what
  Google echoes back at the callback.
- `ExchangeCode(ctx context.Context, code string) (*TokenResponse, error)` —
  POST to `https://oauth2.googleapis.com/token`. The redirect_uri sent here
  is **always the configured one** (R13): never accept it from a request.
- `VerifyIDToken(ctx context.Context, idToken string, expectedNonce string) (*Claims, error)` —
  parse + verify JWT using golang-jwt/jwt/v5 with a JWKS keyfunc:
  - Signature: **RS256 only**, explicitly (`jwt.WithValidMethods([]string{"RS256"})`)
    — accepting any algorithm enables the HS256/public-key confusion attack
    (techplan §13).
  - Issuer: `accounts.google.com`.
  - Audience: configured GOOGLE_CLIENT_ID.
  - Expiry: with small leeway (~60s) for server–Google clock skew (techplan §7).
  - Nonce: claim must match expectedNonce compared with
    **subtle.ConstantTimeCompare** (R23) — never `==`.
  - Error semantics (R26 vs R5): a replayed/replayed-mismatched nonce returns
    the distinguishable sentinel **`ErrNonceMismatch`** → maps to
    `nonce_mismatch` downstream; **every other** verification failure
    (signature, issuer, audience, expiry, malformed) returns a generic error
    → maps to `google_token_invalid` downstream. Only a nonce mismatch may
    produce nonce_mismatch.
- JWKS cache (R21): keys fetched from
  `https://www.googleapis.com/oauth2/v3/certs`, cached **15 min**, and —
  critically — **refresh-on-miss**: when verification encounters a kid not
  present in the cache, refetch once and retry verification. A cache miss
  never becomes a permanent failure. Never cache forever at startup (Google
  rotates keys).
- `Claims` struct: Email, Sub, Iss, Aud, Exp, Nonce.
- Sanitized error logging (R17): when token exchange or JWKS fetch fails,
  log a sanitized category only ("timeout", "http error", status code) —
  **never** the raw error string, which may embed tokens, code values, or
  response bodies. Never log id_token contents, codes, or client secret.

Requirements verbatim from techplan §3 that this package satisfies:

| Condition | Requirement |
|---|---|
| id_token verification | Signature (RS256 against Google JWKS), issuer (accounts.google.com), audience (GOOGLE_CLIENT_ID), expiry, nonce |
| Fixed redirect_uri | GOOGLE_REDIRECT_URI from env var, never from request |
| Google API timeout/unreachable | Clean error return (caller maps to google_unavailable redirect); not a hang |
| Google OAuth client timeout | HTTP client with explicit timeout for token exchange call |
| State comparison constant-time | subtle.ConstantTimeCompare for nonce here (state compare lives in Task 03) |
| Separate keys per purpose | This package uses only the Google client secret — never the app JWT signing key, never PII encryption/HMAC keys |

Risk rows from techplan §7 that constrain this task:

| Risk | Mitigation |
|---|---|
| Incorrect id_token verification (wrong algorithm, missing aud/iss check) — High severity, account takeover | golang-jwt with explicit iss/aud/exp validation; RS256 only; forged-token tests below |
| Google JWKS keys rotate, cached keys stale | Cache TTL 15 min + refresh-on-miss |
| Clock skew between server and Google | Small leeway (~60s) in exp validation |
| Google API client has no timeout — goroutine hangs indefinitely | http.Client{Timeout: 10s} built once in NewClient, reused; http.NewRequestWithContext everywhere |

## 6. Rules covered

Directly: **R6** (timeout → clean error the caller maps to
google_unavailable), **R13** (fixed redirect_uri), **R17** (sanitized error
logging), **R21** (JWKS refresh-on-miss), **R22** (explicit timeout +
NewRequestWithContext, never http.DefaultClient), **R23** (nonce half of the
constant-time requirement), **R26** (strict sig/iss/aud/exp → generic error;
only nonce mismatch is special).

Enabling for later tasks: R4/R19/R20 ordering guarantees ("no Google API
call before state validation") depend on this package existing as the single
chokepoint behind which all Google I/O hides.

Full rule-to-task mapping: see `manifest.md`.

## 7. Testing checklist (this task's slice)

All HTTP interactions mocked via `httptest.Server` injected as base URL —
no live network calls in tests.

- [ ] ExchangeCode sends client_id, client_secret, code, grant_type, and the
      **configured** redirect_uri to the mock token endpoint; parses the JSON
      response into TokenResponse (R13).
- [ ] ExchangeCode returns an error (not a hang) when the mock endpoint
      delays beyond the timeout / refuses connection — caller can map it to
      google_unavailable (R6, R22). Assert the error message carries no
      response body.
- [ ] VerifyIDToken accepts a correctly RS256-signed token with matching
      iss/aud/exp/nonce and returns populated Claims.
- [ ] Forged-token tests (techplan §7 row 1 explicitly asks for these):
      - bad signature → generic error (→ google_token_invalid path)
      - wrong issuer → generic error
      - wrong audience → generic error
      - expired (with leeway boundary just outside tolerance) → generic error
      - HS256-signed (algorithm confusion attempt) → rejected, not verified (§13)
- [ ] Nonce mismatch returns exactly `ErrNonceMismatch` — distinguishable
      from every other failure (R5/R26 contract).
- [ ] JWKS refresh-on-miss: first verification against a kid absent from the
      initial cached set triggers exactly one refetch and then succeeds
      (R21). Count fetches with a counter in the mock.
- [ ] Constant-time comparison used for nonce (code-review assertion; not
      directly unit-testable).
- [ ] No test log line contains an id_token, code, access_token, or client
      secret (R16/R17 discipline).
- [ ] `go test -race ./internal/platform/googleoauth/...` clean (cache is
      read by verification and written by refresh — concurrency-relevant).

## 8. Common mistakes to avoid (techplan §13 slice)

| Mistake | Fix |
|---|---|
| Using `http.DefaultClient` for token exchange | Goroutine hangs indefinitely if Google is slow — construct http.Client{Timeout: 10s} once in NewClient, use NewRequestWithContext. |
| Logging Google client errors verbatim | May embed tokens/PII — log sanitized category ("timeout", "http error", status) only. |
| Accepting any JWT algorithm | Algorithm confusion attack (HS256 with public key) — explicitly RS256 only via WithValidMethods. |
| Caching JWKS forever at startup | Key rotation causes permanent verification failure — TTL + refresh-on-miss. |
| Not verifying nonce in id_token | Replay attack with stolen id_token — verify claim matches stored value. |
| Comparing nonce with `==` | Timing side-channel — subtle.ConstantTimeCompare. |
| Returning one opaque error for everything | Caller cannot distinguish nonce_mismatch (replay) from google_token_invalid (forgery/bad claims) — keep ErrNonceMismatch distinct. |

## 9. Risk note (to fill in the PR)

- Assumptions made: ...
- Edge cases intentionally NOT handled (and why): ...
- Concurrency assumptions: ...
- What is not tested, and why: ...
