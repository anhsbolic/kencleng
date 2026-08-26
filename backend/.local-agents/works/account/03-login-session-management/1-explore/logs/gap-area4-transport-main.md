# Gap Analysis — Area 4: `internal/transport/http/` + `cmd/server/main.go`

> Files: `middleware.go`, `cookie.go`, `errors.go`, `auth_register.go`,
> `auth_google.go`, `doc.go`, `cmd/server/main.go`

## Current state

- **Routing** (`main.go:124-144`): Go 1.22 pattern mux; `authMux` holds
  register/verify-email/resend + Google redirect/callback; whole subtree
  wrapped once: `mux.Handle("/auth/", transporthttp.RateLimit(rps, burst)
  (authMux))` — the in-memory rate limit DOES exist:
  `middleware.go:20-72`, per-IP `x/time/rate` bucket, 10-min idle eviction
  sweeper, 429 via `WriteProblem`. New routes registered on `authMux`
  inherit it automatically.
- **Cookie infra** (`cookie.go`): names centralized — `kencleng_access`,
  `kencleng_refresh` (:8-11). `writeAuthCookies` (:72-91) sets BOTH tokens as
  cookies (OAuth 302 contract): refresh = HttpOnly + Secure(conditional) +
  SameSite=Strict, 30 d, Path=/; access = Lax. Only one clear helper exists
  (`clearOAuthStateCookie`) — nothing for logout.
- **Token verification precedent** (`auth_google.go:51-91`):
  `GoogleTokenVerifier(publicKey)` — ES256-only, exp required, 1-min leeway,
  sub→userID; carrier extraction `sessionToken()` = access cookie first,
  then Bearer. Comment (:46-50) explicitly defers shared-helper extraction to
  "the login/session task (#3)... possibly as a human-paired change." No
  `purpose` claim check anywhere.
- **Error mapping** (`errors.go:60-81`): `MapServiceError` knows only
  422/410/404/default-500. No 401 or lockout-429 branch. `WriteProblem`/
  `WriteValidationError` established; unknown errors logged server-side,
  generic 500 out.
- **Handler pattern** (`auth_register.go`): decode DTO → boundary field
  validation → service call → sentinel map → shape response; constructor-
  injected service seams keep transport unit-testable.
- **Wiring**: `cookieSecure = appEnv != "development"` (:138); required-env
  list has no `MFA_PENDING_TOKEN_SECRET`.

## Requirement

Four new POST routes; login/mfa-complete set refresh-only cookie; refresh
rotates it; logout clears idempotently; 401/429 Problem mappings honoring the
identical-generic-detail rule; response shaping for three success schemas;
access-token verifier/middleware with `purpose` defense-in-depth (spec names
`TestAuthMiddleware_*` tests).

## Gap

1. Routes missing for all four endpoints.
2. No refresh-only cookie setter; no refresh-cookie clearer (logout).
3. Error vocabulary missing end-to-end (new sentinels + mapper branches).
4. Verifier is Google-specific, purpose-blind, closure-not-middleware;
   extraction was explicitly deferred as possibly human-paired.
5. `MFA_PENDING_TOKEN_SECRET` plumbing missing.
6. Lockout-vs-rate-limit 429 shapes need disambiguation (see sniffing).

## Sniffing findings

- **Risk:** limiter keys on `r.RemoteAddr` — behind Caddy every client shares
  one IP ⇒ one shared bucket across all users (in-code comment acknowledges;
  X-Forwarded-For deferred). Adding the four hottest endpoints under this
  limiter raises blast radius of the pre-existing flaw. Decision: defer,
  flagged follow-up (Stage 3).
- **Inconsistency #1:** three different 429 bodies in play — middleware's
  `"problems/rate-limited"`/"Too many requests…" (English), contract example
  `"errors/too-many-requests"`/"Terlalu banyak percobaan gagal…" (Indonesian),
  and feature-spec's "identical to 401 detail" rule. No single implementation
  satisfies all three → Stage 3 D5.
- **Misleading signal:** threat model credits rate limiting that lives in
  transport while `platform/ratelimit/doc.go` claims it "arrives in a later
  task" — stale package doc; inverse-direction misleading signal vs Area 3.
- **Edge cases:** browser-dropped-cookie vs server-expired both land 401
  identically (fine); logout with no cookie must still 204 (handler-trivial);
  `sessionToken()` Bearer fallback matters for future protected endpoints.
- **Observation:** reusing `writeAuthCookies` at `/auth/login` would
  over-deliver an access cookie the contract doesn't define — login must set
  refresh-only (Stage 3 D6).
