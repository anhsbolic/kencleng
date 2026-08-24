# Task 04 — Transport + Wiring (cookies, handlers, reauth markers, main.go)

> Ticket    : 02-google-oauth-login-register
> Sub-task  : 4 of 4
> Axis      : Dependency/sequence chain (primary) + component/layer alignment
> Status    : Ready (after Task 03 — handlers call the service methods)
> Back-ref  : `../techplan.md` — "Tech Plan: Google OAuth Login/Register" (originating contract techplan; cross-check high-level decisions there whenever needed)

---

## 1. Scope

The HTTP face of the OAuth flow: cookie utilities, the two GET endpoints,
the inline session verification for link/reauth intents, the in-memory
reauth marker store, and startup wiring.

**In scope:**
- `internal/transport/http/cookie.go` (new) — setOAuthStateCookie /
  readOAuthStateCookie / setAuthCookies
- `internal/transport/http/auth_google.go` (new) — GoogleRedirectHandler,
  GoogleCallbackHandler, reauth marker store, inline ES256 JWT verification
- `internal/transport/http/auth_google_test.go` (new) — handler tests
- `cmd/server/main.go` — requireEnv additions, googleoauth.Client
  construction, NewService call update, route registration

**Out of scope:**
- All branching business rules — they live in Task 03's service; handlers
  translate results into Set-Cookie headers and 302 responses.
- `.env.example` — no change needed (all four vars already present; verified
  techplan §10).
- Caddyfile — root-level boundary; see Open Items below.
- Any modification to `internal/platform/auth/` — **Tier 0 fenced path**
  (root AGENTS.md §3). The public key is *passed to* the handler as a
  dependency; nothing in platform/auth changes. If a shared JWT-verify
  helper extraction ever happens, that is task #3+ and possibly human-paired
  (techplan §9 step 5).

## 2. Dependencies

- **Hard deps:** Task 03 (`GoogleRedirect`/`GoogleCallback` signatures,
  updated NewService), Task 01 (jwt dep for inline verification).
- **Blocks:** none — last task; completes the ticket.

## 3. Files

| File | Change Type |
|---|---|
| `backend/internal/transport/http/cookie.go` | New |
| `backend/internal/transport/http/auth_google.go` | New |
| `backend/internal/transport/http/auth_google_test.go` | New |
| `backend/cmd/server/main.go` | Edit |

## 4. Why this design (techplan §5 Decision Log rows — do not relitigate)

| Option considered | Why rejected/accepted |
|---|---|
| **A. Handler-level cookie utility (chosen)** | Chosen — two functions in transport/http/cookie.go. Keeps cookie logic close to where it is used without over-abstracting. |
| **B. Platform package for cookies** | Rejected — over-engineering for now; only OAuth uses cookies. |
| **A. Handler-level JWT verification (chosen)** | Chosen — handler checks intent first, verifies JWT inline using golang-jwt/jwt/v5 with the public key passed as a dependency. Does NOT touch platform/auth/ (Tier 0 fenced). Explicit, testable. Task #3 can later extract a shared helper if needed. |
| **B. Conditional middleware** | Rejected — middleware reading query params to decide auth is unusual and hard to test. |
| **A. In-memory sync.Map for reauth markers (chosen)** | Chosen — no Redis dependency needed. Markers lost on restart (acceptable for 5-min TTL). Same eviction pattern as rate limiter in middleware.go. |
| **B. Redis-backed marker store** | Rejected (reverted per user decision) — Redis is not in docker-compose.yml, no client in go.mod; adding it requires root-level infra changes outside backend/. |

## 5. API contract (verbatim from techplan §8)

```
GET /auth/google/redirect?intent={login|link|reauth}   (new, public)
  -> 302 to Google consent screen (sets state/nonce cookie)
  -> 401 if intent=link/reauth without session
  -> 400 if intent is invalid

GET /auth/google/callback?code=...&state=...            (new, public)
  -> 302 to frontend (success: tokens as cookies; error: ?error={code})
  Error codes: state_mismatch, nonce_mismatch, google_token_invalid,
               google_unavailable, google_email_conflict, google_link_conflict
  (google_token_invalid = id_token fails sig/iss/aud/exp; nonce_mismatch = replay)
```

Error redirects target FRONTEND_URL-derived pages: login errors → `/login`,
link/reauth errors → security page (per techplan §8 flow). The error-code
set v2 and the `account_linking` audit literal are pending frontend sign-off
(techplan §14 Active item 2) — implement them as pinned here; flag any
frontend divergence in the PR risk note rather than changing literals.

## 6. Implementation detail

### `transport/http/cookie.go` (new)

Per techplan §9 step 7 + §10:

- `setOAuthStateCookie(w http.ResponseWriter, state, nonce, intent string, userID *uuid.UUID)` —
  JSON-encode `{state, nonce, intent, user_id}`, base64-encode, Set-Cookie with:
  - HttpOnly ✓
  - Secure ✓ (**always in non-dev**; conditional on APP_ENV in dev — HTTP dev
    needs Secure off, techplan §7 row "Cookie not Secure in dev")
  - SameSite=**Lax** (NOT Strict — Google's redirect back is cross-origin;
    Strict would drop the cookie; Strict belongs to the *refresh* cookie below)
  - MaxAge=600 (~10 min state TTL, R24)
  - Path=/auth/google (scoped so the browser sends it only on this flow)
- `readOAuthStateCookie(r *http.Request) (state, nonce, intent string, userID *uuid.UUID, err error)` —
  read + decode; error on absent/expired/corrupt cookie → caller maps to
  state_mismatch (R20).
- **Clear the state cookie after reading** at callback (§13 final row): set
  MaxAge < 0 on the response so a replayed callback cannot reuse it.
- `setAuthCookies(w http.ResponseWriter, accessToken, refreshToken string)` —
  two cookies per techplan §10:
  - access token cookie: short TTL, HttpOnly
  - refresh token cookie: HttpOnly, Secure, **SameSite=Strict**, MaxAge=30d

### `transport/http/auth_google.go` (new)

**Handlers** (techplan §9 step 8 + §10):

- `GoogleRedirectHandler(svc, verifyToken func(string) (uuid.UUID, error))`:
  1. Read `intent` query param. Not in {login, link, reauth} → **400 Bad
     Request** (R18).
  2. Intent = link or reauth → verify the app's own ES256 access token
     **inline**: golang-jwt/jwt/v5 with
     `jwt.WithValidMethods([]string{"ES256"})`, public key
     (`*ecdsa.PublicKey`) passed in as a handler dependency (loaded via
     `auth.Load` at startup in main.go). No valid session → **401 BEFORE any
     Google redirect** (R2 — explicit authz check at the handler boundary,
     visible and testable on its own, per AGENTS.md golden rule; never just
     a query filter).
  3. Intent = login → no auth check performed (R1).
  4. Call `svc.GoogleRedirect(...)`; write the state cookie from the returned
     cookieValue; 302 to the returned redirect URL.
- `GoogleCallbackHandler(svc, frontendURL string)`:
  1. Missing code or state query param → 302 error redirect
     `?error=state_mismatch`; **no Google API call** (R19).
  2. Missing/expired/corrupt state cookie → same state_mismatch redirect
     (R20); clear the cookie.
  3. Call `svc.GoogleCallback(code, state, cookieValue)`.
  4. Result has tokens → setAuthCookies; result has Error → 302 to frontend
     with `?error={code}`; either way write the 302 to result.RedirectURL.

**Inline JWT verification constraint (R25):** the verification function is
built here in the transport layer using golang-jwt/jwt/v5 and the injected
public key — `platform/auth/keys.go` remains untouched (Tier 0 fenced, root
AGENTS.md §3).

**Reauth marker store** (techplan §9 step 9):

```go
var reauthMarkers sync.Map   // key: userID, value: expiry time
```

- Background eviction goroutine clearing expired markers — same pattern as
  the rate limiter in middleware.go (RateLimit, middleware.go:20).
- `SetReauthMarker(userID uuid.UUID, expiry time.Time)`,
  `CheckReauthMarker(userID uuid.UUID) bool`.
- TTL = **5 minutes**, matching the state-cookie TTL deliberately (feature
  spec Assumption A; one re-auth-freshness convention, techplan §14 Resolved
  item 4).
- Marker is set by the callback handler after a successful reauth-intent
  result (Task 03 guarantees no AuthIdentity/token side effects for that
  intent). Consumed by task #06 (MFA disable) — expose CheckReauthMarker but
  do not wire consumption now.
- Markers are lost on process restart — accepted trade-off for a 5-min TTL
  (decision recorded in techplan §5/§14; do not add persistence).

### `cmd/server/main.go`

Per techplan §9 step 10 + §10:

- requireEnv += GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET, GOOGLE_REDIRECT_URI
  (already in .env.example) **and FRONTEND_URL**.
- Construct `googleoauth.NewClient(clientID, clientSecret, redirectURI)` once.
- Load keys via `auth.Load` (existing startup path); pass the public key into
  GoogleRedirectHandler and the keys into account.NewService.
- Update the `account.NewService(...)` call for the new 6th/7th parameters
  (Task 03's signature).
- Register on authMux:
  - `mux.HandleFunc("GET /auth/google/redirect", ...)` (Go 1.22 pattern routing)
  - `mux.HandleFunc("GET /auth/google/callback", ...)`
- Dev-only docs routes/outbox behavior is untouched by this task.

Log discipline for everything above (R16/R17, AGENTS.md golden rule): log
operation fact + outcome only. Never log state, nonce, code, id_token,
access_token, refresh_token, cookie values, or raw Google client errors.

## 7. Rules covered

Primary owner for: **R2** (401 before Google redirect — named test below),
**R3** (session user_id encoded into the state cookie), **R18** (invalid
intent → 400), **R24** (full cookie attribute set),
**R25** (inline ES256 verification, platform/auth untouched),
plus the transport half of **R19/R20** (param/cookie presence checks before
any Google call), the 302-writing halves of R1/R4–R12, and the
Set-Cookie halves of R7/R8 (auth cookies on success).

Full rule-to-task mapping: see `manifest.md`.

## 8. Testing checklist (this task's slice)

Handler tests via httptest against the mux; service replaced with a stub
implementing the two methods (Task 03's fake client pattern applies there,
not here).

Named test required by the techplan:

- [ ] `TestGoogleRedirect_LinkReauthRequireAuth` — R2: intent=link and
      intent=reauth without credentials → 401, and the response contains no
      Location header pointing at Google (reject happens BEFORE any redirect).
- [ ] R1: intent=login with no credentials → 302 to Google consent URL,
      state cookie present, no auth attempted.
- [ ] R3: valid access token + intent=link → cookie payload decodes with
      user_id populated; 302 to Google.
- [ ] R18: intent=`evil` → 400.
- [ ] R19: callback without code / without state → 302 Location containing
      `error=state_mismatch`; stub service records zero calls.
- [ ] R20: expired/missing/corrupt state cookie → same; zero calls.
- [ ] Success path: stub returns CallbackResult with tokens → both auth
      cookies written with correct attributes (refresh: HttpOnly, Secure,
      Strict, 30d); 302 to expected page.
- [ ] Error paths: each of the six error codes surfaces verbatim as
      `?error={code}` in the Location header.
- [ ] R24: assert state cookie attributes exactly — HttpOnly, Secure
      (non-dev), Lax, MaxAge=600, Path=/auth/google.
- [ ] State cookie cleared after callback read (MaxAge < 0 in response).
- [ ] Reauth success → SetReauthMarker called; CheckReauthMarker true within
      TTL; false after eviction (use short/injectable clock or TTL for test).
- [ ] Log-output capture across all handler paths: no state, nonce, code,
      id_token, access_token, refresh_token substrings (R16).
- [ ] `go test ./...` green; `go vet ./...` clean; run
      `go test -race ./internal/transport/http/...` (sync.Map store is
      concurrency-relevant).

Manual smoke (dev env only): full round-trip with real Google consent using
GOOGLE_REDIRECT_URI=http://localhost:8090/auth/google/callback (direct to
backend, bypassing Caddy — see Open Items).

## 9. Common mistakes to avoid (techplan §13 slice)

| Mistake | Fix |
|---|---|
| SameSite=Strict on the state cookie | OAuth redirect from Google is cross-origin; Strict blocks the cookie. Lax for state cookie; Strict only for refresh-token cookie. |
| Forgetting the auth check for link/reauth intent | Attacker links Google to any user — conditional check in handler before redirect (R2, named test). |
| Not clearing state cookie after callback | Cookie replay on subsequent callbacks — MaxAge < 0 after reading. |
| Routing JWT verification through platform/auth/ modifications | Tier 0 fenced — inline golang-jwt with injected public key (R25). |
| Taking redirect_uri from the request at wiring time | Only GOOGLE_REDIRECT_URI env feeds the client constructor (R13). |
| Missing timeout on outbound calls added during wiring | Client built once in main with its own 10s timeout (R22). |

## 10. Open items affecting this task (from techplan §14 Active)

1. **Caddy routing for `/auth/*`.** The Caddyfile routes only `/api/*` to the
   backend (:8090); `/auth/google/callback` (no `/api` prefix) would go to
   the frontend (:3000) through Caddy in non-dev. Dev workaround already in
   place: direct `:8090` URI. Non-dev fix requires a root-level session
   (backend/AGENTS.md §5 known infra gap) — do NOT edit the Caddyfile from
   this backend session; surface it in the PR risk note.
2. **Frontend sign-off pending** on error-code set v2 (incl.
   google_token_invalid vs nonce_mismatch split) and the
   `user_logs.action_type='account_linking'` literal. Backend pins both as
   specified; divergence is a flag-for-human item, not a local change.

## 11. Risk note (to fill in the PR)

- Assumptions made: ...
- Edge cases intentionally NOT handled (and why): ...
- Concurrency assumptions: ...
- What is not tested, and why: ...
