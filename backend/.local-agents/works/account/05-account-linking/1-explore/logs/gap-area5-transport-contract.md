# Stage 2 — Gap Analysis, Area 5: Transport wiring + API contract

> Files: `cmd/server/main.go`, `transport/http/middleware.go`,
> `cookie.go`, `auth_google.go`, `errors.go`; `api/openapi.yaml`

## Current state (concrete)

- **Routing** (`cmd/server/main.go`): all routes are `/auth/*` (mounted
  under `RateLimit`) plus `/healthz`, `/docs`, `/openapi.yaml`.
  **No `/account/*` route group exists** — these two endpoints will be
  the first authenticated, non-`/auth` endpoints.
- **No auth middleware exists.** `auth_google.go` establishes an inline
  pattern instead: `sessionToken(r)` (reads `kencleng_access` HttpOnly
  cookie first, falls back to `Authorization: Bearer`) and
  `GoogleTokenVerifier(publicKey)` — a handler-boundary ES256 verifier
  built over golang-jwt directly, explicitly because `platform/auth/` is
  Tier 0 fenced and must not be modified.
- **Reauth marker store exists** (`reauthMarkers sync.Map`, 5-min TTL,
  in-memory with sweeper): documented as "consumed by task #06"
  (MFA-disable). Per this feature's resolved design, unlink does NOT use
  it — password confirmation instead.
- **openapi.yaml**: both endpoints already fully specified post-redesign
  — `UnlinkGoogleRequest{password}`, `SetPasswordRequest{email?,
  current_password?, password}` with branch semantics in descriptions,
  `200`/`202` split, and two 409 examples with distinct problem-type
  URIs (`errors/only-identity`, `errors/unverified-remaining-identity`)
  whose Indonesian detail strings exactly match the feature spec.
- **Contract defect**: global `security: - bearerAuth []` is declared at
  line 41, but **no `components.securitySchemes` section exists anywhere
  in the file** — `bearerAuth` is referenced but never defined.

## Requirement vs Gap

1. New handlers `SetPasswordHandler` / `UnlinkGoogleHandler` + mounting
   of `/account/security/*` routes don't exist.
2. No shared "authenticated request → user_id" extraction exists; only
   the OAuth-specific `sessionToken` + `GoogleTokenVerifier` pair. Both
   new endpoints need this on every request — the first true
   always-authenticated endpoints in the backend.
3. Spec's References section says both endpoints "need a schema update —
   flagged as follow-up" — **this looks stale**: schemas already carry
   the 2026-08-05 redesign (branching, re-auth password, distinct 409s).
   The real openapi gaps are the missing `securitySchemes` definition.

## Sniffing

- *Misleading signal*: `CheckReauthMarker`/`reauthMarkers` sitting in
  the transport package invites wiring unlink to it; spec says no.
- *Miscontext*: spec says openapi needs updating — partially outdated;
  the contract text is ahead of the feature spec's own reference note.
- *Inconsistency*: contract declares `bearerAuth` globally while the
  implemented session delivery is an HttpOnly cookie (Bearer as test
  fallback); undefined scheme means contract tooling can't describe
  auth.
- *Risk*: no middleware-level session story exists to inherit — each
  authenticated endpoint so far rolled its own check; this task sets
  the pattern for `/account/me` (#7), MFA (#6), and reset-password (#4)
  handlers.
