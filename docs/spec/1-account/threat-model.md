# Threat Model — account

> Status: draft
> Last updated: 2026-08-05

## Actors & trust boundaries

| Actor | Authenticated? | Trust boundary crossed |
|---|---|---|
| Guest / anonymous visitor | No | Can hit `/auth/register`, `/auth/login`, `/auth/forgot-password`, `/auth/google/redirect` — all public, unauthenticated write/read paths into the system |
| Registered user, unverified email | Partial (has an access token, but `AuthIdentity.verified_at = null`) | Can log in and use account self-service, but is restricted from donation-as-registered / representative actions in **other** domains — enforcement of that restriction lives in those domains, not here |
| Registered user, verified | Yes | Full access to `/account/*` self-service endpoints, scoped to their own `user_id` |
| Admin | Yes + `user_roles.role = 'admin'` | Access to `/admin/users*` — the highest-privilege boundary in this domain (bulk PII read + role elevation) |
| Google (external IdP) | N/A — external trust boundary | `/auth/google/callback` receives an authorization code and, after token exchange, an `id_token` from outside our system. Nothing from this path is trusted until `state`, `nonce`, and JWT signature are all verified |

## STRIDE per component

Endpoints are grouped into 6 components (rather than STRIDE'd
individually per the 21 `account`-tagged OpenAPI paths) — this is the
domain-level view; per-endpoint threat breakdown happens in each
feature's `docs/spec/account/features/<fitur>.md` (§5.3), narrowed
down from this doc.

### 1. Registration & Email Verification

`POST /auth/register`, `POST /auth/verify-email`,
`POST /auth/verify-email/resend`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | Attacker registers using an email they don't own | Verification token (single-use, 24h expiry) required before the identity counts as verified; unverified users are functionally restricted elsewhere | Low — window between registration and verification is inherent to the email-verification pattern itself |
| Tampering | Malicious/extra fields injected into the registration payload | Strict DTO validation against `openapi.yaml` schema (Tier 2 automated gate); no mass-assignment from raw request body | N/A |
| Repudiation | N/A — registration is a non-destructive action, not in the Fitur 9 audit-log scope | — | N/A — no audit log entry required for this action (see `kencleng-phase0-detail.md` Fitur 9) |
| Information disclosure | Email enumeration via `/auth/register` | **Resolved — 2026-08-05.** Response is always generic ("check your email"), regardless of whether the email is new, already registered-unverified, or already registered-verified. Actual email content sent differs by case (registration-verification / resend-verification-nudge / reset-password-nudge), but the API response and response time must not leak which case occurred — same pattern as Fitur 2B, see §1 below | Requires constant-time handling in the handler (send-or-skip must not create a timing side-channel) — flagged for the Build stage, not yet implemented |
| Denial of service | Mass registration / verification-email flood | Stricter rate limit on all `/auth/*` endpoints (per `AGENTS.md` golden rules) | Low |
| Elevation of privilege | N/A — registration only ever produces a base user with no elevated role | — | N/A |

### 2. Login & Session

`POST /auth/login`, `POST /auth/login/mfa`, `POST /auth/refresh`,
`POST /auth/logout`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | Credential stuffing / brute-force login at the password stage, and TOTP/backup-code brute force at `/auth/login/mfa` (MFA stage) **[added 2026-08-26, per feature-spec Assumption C]** | `login_attempts` persistent lockout (5 failed / 15 min window), two stages: password stage keyed by `identifier_hash` (identity not yet known at check time), MFA stage keyed by `user_id` (`stage='mfa'`, checked before any code verification); plus in-memory rate limit on all `/auth/*`. Lockout-rejected attempts write no `login_attempts` row and use the same generic response body as wrong credentials | Lockout is per `identifier_hash` / per `user_id`, not per-IP — see residual risk §2 below (device/IP fingerprinting). A distributed attack can rotate target identifiers to stay under either counter |
| Tampering | Refresh token replay/reuse, or tampering with the cookie. Token-type confusion — an `mfa_pending_token` presented where an access token is expected **[added 2026-08-26, per feature-spec Assumptions A/B]** | Rotate-on-use + reuse detection (INV-account-03, INV-account-04); cookie is HttpOnly + Secure + `SameSite=Strict`, inaccessible to JS. Token confusion mitigated by two independent layers: the `mfa_pending_token` is signed with a separate HS256 secret (`MFA_PENDING_TOKEN_SECRET`, distinct from the access token's ES256 keypair) so it fails access-token signature verification outright — a cryptographic guarantee, not a logic check — plus an explicit `purpose` claim verified as defense-in-depth on top | N/A |
| Repudiation | N/A significant | Every login attempt (success or fail) recorded in `login_attempts` with timestamp | N/A |
| Information disclosure | Login error message distinguishes "wrong email" from "wrong password" | Generic error message regardless of which check failed, consistent with anti-enumeration in Fitur 2B | N/A |
| Denial of service | Refresh-endpoint flooding | Rate limit on `/auth/*` | The limiter keys on `r.RemoteAddr`; behind the reverse proxy every client shares the proxy IP, collapsing all clients into one bucket until X-Forwarded-For support lands (deferred follow-up, accepted 2026-08-26 — see task #3 techplan) |
| Elevation of privilege | Token issued without completing the required MFA step, for a user with MFA enabled | Token issuance is ordered strictly after MFA verification in the flow (Fitur 2, steps 5–7); INV-account-07 ensures MFA state itself can't be silently half-enabled | N/A |

### 3. Forgot & Reset Password

`POST /auth/forgot-password`, `POST /auth/reset-password`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | Attacker triggers a reset for a victim's email | Reset link is only usable by whoever controls the destination inbox — the API response never reveals whether the request succeeded in a way that helps an attacker without inbox access | N/A |
| Tampering | Reset token guessing/tampering | Random, single-use, hashed, 1h expiry, guarded `UPDATE ... WHERE used_at IS NULL AND expires_at > now()` | N/A |
| Repudiation | N/A | — | N/A |
| Information disclosure | User enumeration via forgot-password response | **Resolved** — identical generic response regardless of match, including the Google-only case (same API response, different email content sent to the actual inbox) | N/A |
| Denial of service | Repeated reset requests flooding the mail queue | Rate limit on `/auth/*`; double-submit is safe (idempotent-enough, not financial-critical) | N/A |
| Elevation of privilege | Hijacking another account via password reset | Same mitigation as Spoofing row above (inbox ownership is the actual gate) | N/A |

### 4. Google OAuth & Linking

`GET /auth/google/redirect`, `GET /auth/google/callback`,
`DELETE /account/security/google/unlink`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | Forged OAuth callback / CSRF | `state` param, random, HttpOnly cookie, short-TTL (~10 min), validated on every callback | N/A |
| Tampering | Replay of a stolen/old `id_token` | `nonce` param validated against the `id_token`'s `nonce` claim | N/A |
| Repudiation | Account linking (adding Google to an existing account) | Logged to `user_logs` per Fitur 9 scope ("account linking baru") | N/A |
| Information disclosure | Open-redirect via a manipulated `redirect_uri` | `GOOGLE_REDIRECT_URI` is a fixed env var, exact-match validated against Google Console registration — never taken dynamically from the request | N/A |
| Denial of service | Google API outage blocks login/registration for Google-only users | **Resolved — 2026-08-05.** Accepted (sandbox project, no SLA, Google OAuth uptime is in practice very high) — see §3 below. Mandatory mitigation: the token-exchange call to Google must have an explicit timeout and surface a clean `503` Problem Details response ("Google sign-in is currently unavailable"), not a raw 500/timeout | A frontend-side nudge encouraging Google-only users to set a password as a backup path (Fitur 4) is a separate, later UX decision — not part of this backend resolution, see §3 note below |
| Elevation of privilege | Auto-linking by email match, enabling account takeover via an insufficiently-verified provider | **Explicitly blocked** — no auto-merge; user must go through an explicit, authenticated linking flow (Fitur 1B step 7, Fitur 4) | N/A |

### 5. Account Self-Service

`GET /account/me`, `POST /account/security/set-password`,
`POST /account/security/mfa/enroll(+confirm)`,
`POST /account/security/mfa/disable`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | MFA disabled by an attacker who hijacked a live session | Disable requires re-authentication (password re-entry, or Google re-login prompt for Google-only users) | Google-only re-auth mechanism's exact technical implementation is still deferred to the coding phase (`kencleng-phase0-detail.md` Open Items) |
| Tampering | N/A beyond standard authz | Every action scoped to the authenticated session's own `user_id`, never an ID param from the request | N/A |
| Repudiation | MFA enable/disable, set-password, linking | All logged to `user_logs` per Fitur 9 scope | N/A |
| Information disclosure | `/account/me` leaking another user's data (IDOR) | Endpoint is keyed by the authenticated session, not a request parameter — no ID to tamper with | N/A |
| Denial of service | Repeated MFA enroll requests (secret generation spam) | Low-cost operation; no dedicated mitigation beyond general rate limiting | Low |
| Elevation of privilege | N/A — this component cannot grant `admin`/`kurator` roles | — | N/A |

### 6. Admin Role Management

`GET/POST /admin/users`, `GET/POST/DELETE /admin/users/{userId}/roles`

| Category | Concrete threat | Existing mitigation | Residual risk |
|---|---|---|---|
| Spoofing | N/A — requires an authenticated Admin session | Admin-only authz check | N/A |
| Tampering | Forged `userId` in the role-assignment request to target an unintended user | Explicit authz + target validation at the handler level (not just a query filter — `AGENTS.md` golden rule) | N/A |
| Repudiation | Role assign/revoke | Mandatory `user_logs` entry per Fitur 9 scope; every feature spec for this component must answer "Audit log entry?" (§5.3 of `kencleng-spec-readme.md`) | N/A |
| Information disclosure | `/admin/users` bulk-exposes decrypted PII (email) for many users in one response, each requiring a decrypt operation | **Pagination hard-capped at max 20 for this endpoint specifically** (not the general default max 50 — deviation justified by decrypt cost per row) **[RESOLVED — 2026-08-05]**; Admin-only authz enforced explicitly at the handler | Requires an `openapi.yaml` change: this endpoint needs its own `LimitParam` override (max 20), not the shared one (max 50) — flagged as a follow-up action, not yet applied to the spec |
| Denial of service | A compromised Admin account spamming role assignments | Out of scope for this domain's threat model — this is an insider/compromised-credential scenario, mitigated by the same session/MFA security as component 2, not a separate control | Accepted — Admin account compromise is a general account-security concern, not `/admin/users`-specific |
| Elevation of privilege | Non-Admin calling `/admin/users*` directly | Role-check middleware, explicit at the handler boundary | N/A |
| Elevation of privilege | Assigning `admin` to a user who is also Kurator/Representative (role-conflict bypass) | INV-account-09 (Admin ⊥ Kurator) and INV-account-10 (Admin ⊥ Representative) enforced at assignment time | N/A |

## Knowingly accepted residual risk

1. **Email enumeration at `/auth/register` — resolved, generic
   response.** **[RESOLVED — 2026-08-05]** `POST /auth/register`
   response is generic regardless of the email's actual state in the
   system (new / registered-unverified / registered-verified),
   matching the anti-enumeration pattern already established in Fitur
   2B (forgot-password). The email actually sent differs by case:
   - New email → registration verification email
   - Registered, unverified → a nudge email offering resend-
     verification (functionally equivalent to what
     `/auth/verify-email/resend` already sends)
   - Registered, verified → a nudge email offering password reset
     (functionally equivalent to what `/auth/forgot-password` sends)

   This is an **`openapi.yaml` change** (the current `/auth/register`
   spec likely documents a `409` response for the already-registered
   case, which must be removed in favor of a uniform `202`/`200`) —
   flagged as a follow-up action for whenever `account`'s OpenAPI
   paths are next touched. Implementation must also handle this
   constant-time (no timing side-channel between the three internal
   cases) — flagged for the Build stage.

2. **Device/IP fingerprinting not implemented — resolved, accepted.**
   `login_attempts` lockout is scoped to `identifier_hash` only, not
   per-IP or per-device. A distributed brute-force attempt spreading
   across many different target accounts from a small set of IPs
   would not be caught by this mechanism. Accepted per
   `kencleng-phase0-detail.md` Fitur 2C Open Questions — considered
   over-engineering for this project's current needs, revisit only if
   a concrete need emerges.

3. **Google API outage blocks Google-only users — resolved, accepted
   with graceful degradation.** **[RESOLVED — 2026-08-05]** Unlike
   HaveIBeenPwned (a fail-open/fail-closed choice), there is no
   meaningful fallback for a Google-only user if Google's OAuth/token
   endpoints are unreachable — they simply cannot authenticate via
   that path during the outage. Accepted as low-priority for a sandbox
   project with no SLA obligations and Google's generally very high
   OAuth uptime in practice. Required mitigation (applies to any
   external call, not special-cased for Google): the token-exchange
   call must have an explicit timeout and return a clean `503`
   Problem Details response, never a raw timeout/500.

   **Separate note, not part of this decision**: a frontend nudge
   encouraging Google-only users to set a password (Fitur 4) as a
   backup sign-in path is a UX decision to be made when `account`'s
   frontend work begins — not a backend threat-model resolution.

4. **HaveIBeenPwned fail-open — resolved, accepted.** Referenced here
   for completeness since it affects this domain's registration and
   reset-password flows: if the breach-check API is unreachable,
   registration/reset proceeds without the check (logged via ordinary
   observability logging, not `user_logs`). See
   `kencleng-phase0-detail.md` Fitur 1 for the full resolved rationale.

5. **`mfa_backup_codes` rows accumulate indefinitely across
   disable/re-enable cycles — resolved, accepted.** Per
   `docs/spec/account/invariants.md` INV-account-06's implicit-
   invalidation decision: disabling MFA does not delete old backup
   code rows, it only makes them functionally unusable. No
   housekeeping/cleanup job exists for this table. Accepted as a
   low-severity storage cost, not a security gap (INV-account-06
   guarantees the stale rows can never be redeemed).

6. **Cookie-session CSRF relies on `SameSite=Strict` alone —
   resolved, accepted.** **[RESOLVED — 2026-08-26]** The refresh
   cookie (consumed implicitly by `/auth/login/mfa`,
   `/auth/refresh`, `/auth/logout`) has no second CSRF layer
   (custom-header check or double-submit token). Rationale: Strict
   already blocks cross-site browser requests from carrying the
   cookie; the sandbox topology has no untrusted same-site
   subdomains (the main residual gap of site-scoped SameSite); and
   the worst-case abuse of a forged same-site request is nuisance
   (forced rotation/logout), not data exposure. Revisit trigger: if
   untrusted same-site subdomains ever appear, or when the frontend
   track lands — the centralized React API client is the natural,
   near-zero-cost place to add a custom header then. Decided per
   task #3's techplan Open Item (Anhar, 2026-08-26).

7. **`login_attempts` grows unboundedly — resolved, accepted for
   v1.** **[RESOLVED — 2026-08-26]** The table is append-only audit
   material (successes included, by contract) with no retention job;
   `user_id ON DELETE SET NULL` means user deletion never prunes
   history either. Accepted as a storage cost, not a security gap:
   lockout lookups hit targeted indexes (`identifier_hash`,
   `user_id`+`stage`) whose cost is size-insensitive, time-range
   audit queries ride the BRIN index, and sandbox-scale volume is
   GBs/year at worst. If real usage ever demands housekeeping, the
   right vehicle is a standalone hard-delete-style task (precedent:
   notification domain's `05-hard-delete-worker`), not login-slice
   scope. Decided per task #3's techplan Open Item (Anhar,
   2026-08-26).

## References

- Related ERD: `docs/project/kencleng-erd.md` §1 "Identity & Auth"
- Related business process: `docs/project/kencleng-phase0-detail.md`
  Fitur 1, 1B, 2, 2B, 2C, 3, 4, 5, 9
- Related invariants: `docs/spec/account/invariants.md`
  (INV-account-03, 04, 06, 07, 09, 10 referenced above)
- Related API contract: `api/openapi.yaml`, tag `account`