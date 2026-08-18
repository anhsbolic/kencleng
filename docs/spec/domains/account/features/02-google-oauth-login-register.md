# Feature Spec — Google OAuth Login/Register

> File: `docs/spec/account/features/02-google-oauth-login-register.md`
> Status: draft
> Risk tier: 1
> Domain: account

## Endpoint

- `GET /auth/google/redirect?intent={login|link|reauth}`
- `GET /auth/google/callback?code=...&state=...`

**Structural note**: these two endpoints are shared across three
different business flows (login/register — this task; account
linking — `05-account-linking.md`; MFA-disable re-authentication —
`06-mfa-totp.md`), distinguished by the `intent` query param carried
through `state`. Since it's one physical endpoint pair, the full
OAuth mechanics (state/nonce validation, token exchange, `id_token`
verification) are specified once here — `05` and `06` reference this
file for the mechanics and only add their own intent-specific
business rules, rather than re-describing the OAuth flow.

## Acceptance criteria

### `GET /auth/google/redirect`

- Given `intent=login`, When called, Then no auth check is required;
  system generates random `state` + `nonce`, stores both in a
  short-TTL (~10 min) HttpOnly cookie, and responds `302` to Google's
  consent screen.
- Given `intent=link` or `intent=reauth`, When called **without** a
  valid authenticated session, Then `401` — these intents require an
  existing session (per `openapi.yaml`'s description), and this must
  be checked **before** redirecting to Google, not deferred to the
  callback.
- Given `intent=link` or `intent=reauth`, When called **with** a
  valid session, Then the session's `user_id` is also encoded into
  the same short-TTL cookie state (alongside `state`/`nonce`) — needed
  at callback time since intent handling for `link`/`reauth` must act
  on a specific existing user, not create a new one.

### `GET /auth/google/callback`

- Given a valid `state` match, valid `nonce` match in the `id_token`,
  and a successful token exchange + `id_token` signature/issuer/
  audience/expiry verification, the response branches on `intent`:

  | Intent | Existing `google` identity? | Email used by `email_password`? | Result |
  |---|---|---|---|
  | `login` | Yes | — | Treat as login — issue access + refresh tokens (Fitur 2), `302` to app |
  | `login` | No | No | Create `User` + `AuthIdentity` (`provider_type=google`, `verified_at=now`), issue tokens, `302` to app |
  | `login` | No | Yes (different user) | **No auto-merge** — `302` to `/login` with an error query param, no new records created |
  | `link` | — | Yes (different user) | Reject — `302` to the account-security page with an error query param, no new `AuthIdentity` created |
  | `link` | — | No / belongs to the same session's user | Attach new `google` `AuthIdentity` to the session's existing `user_id` (not a new `User`) |
  | `reauth` | — | — | No `AuthIdentity`/token changes — sets a short-lived server-side re-auth marker tied to the session (see Assumption A), `302` back to the security page |

- Given `state` doesn't match the cookie, When called, Then reject
  (`302` to an error route) before any token exchange happens with
  Google — no network call to Google is made on a state mismatch.
- Given the `id_token`'s `nonce` claim doesn't match the stored
  `nonce`, When called, Then reject, no `User`/`AuthIdentity`
  created/modified.
- Given the token-exchange call to Google times out or Google is
  unreachable, When called, Then respond with a clean `503`-equivalent
  error redirect (per `docs/spec/account/threat-model.md` component 4
  resolution) — not a raw 500/timeout.

## Error cases

| Condition | Expected response |
|---|---|
| `intent=link`/`reauth` without an authenticated session | `401` (at `/auth/google/redirect`, before any Google redirect) |
| `state` mismatch at callback | `302` error redirect, no Google API call made |
| `nonce` mismatch at callback | `302` error redirect, no state change |
| `login` intent, email already claimed by `email_password` | `302` error redirect, no auto-merge |
| `link` intent, email already claimed by a different user | `302` error redirect, no new identity created |
| Google API unreachable/timeout during token exchange | Clean `503`-equivalent error redirect, not raw timeout |

## Applicable invariants

- `docs/spec/account/invariants.md#inv-account-01` — uniqueness is
  per-provider; this is what makes it structurally possible for the
  same real-world email to exist under both `google` and
  `email_password` for two different users (the no-auto-merge case is
  a deliberate choice not to collapse that into one, not a constraint
  INV-account-01 itself imposes).
- No dedicated `INV-account-NN` exists for "no auto-merge" — it's a
  threat-model-level mitigation (`docs/spec/account/threat-model.md`
  component 4, Elevation of Privilege row), not a standalone
  invariant. Documented here for traceability rather than inventing a
  new invariant number for it.

## Threat breakdown

Derived from `docs/spec/account/threat-model.md`, component 4:

| Threat | Mitigation at this endpoint's level | Test that proves it |
|---|---|---|
| Forged callback / CSRF | `state` param, HttpOnly cookie, short-TTL, validated before any Google network call | `TestGoogleCallback_StateValidation` |
| `id_token` replay | `nonce` claim validated against the value stored at redirect-time | `TestGoogleCallback_NonceValidation` |
| Open redirect via manipulated `redirect_uri` | Fixed `GOOGLE_REDIRECT_URI` env var, exact-match registered in Google Console — never taken from the request | `TestGoogleRedirect_FixedRedirectURI` |
| Account takeover via auto-merge on email match | Explicitly blocked for both `login` and `link` intents — no code path creates/attaches an identity without an explicit, authenticated action from the account owner | `TestGoogleCallback_NoAutoMerge_Login`, `TestGoogleCallback_NoAutoMerge_Link` |
| `link`/`reauth` intent called without an existing session | `401` at the redirect step, before Google is ever contacted | `TestGoogleRedirect_LinkReauthRequireAuth` |
| Google API outage | Timeout + clean `503`-equivalent redirect, not raw error | `TestGoogleCallback_UpstreamTimeout_CleanError` |

## Risk tier & rationale

**Tier 1** — CSRF/replay protection (`state`/`nonce`) and the
no-auto-merge account-takeover-prevention rule both need human review
even though most of the flow is standard OAuth plumbing; no Tier 0
sub-area (no core crypto beyond standard JWT verification against
Google's published JWKS, which is library-backed, not hand-rolled).

## Assumptions / open questions

**A. `reauth` marker TTL is undefined anywhere in the source docs.**
`openapi.yaml`'s description for `/account/security/mfa/disable`
mentions "a short-lived server-side marker tied to the session
(implementation detail)" but no document specifies how short. Proposed
default: **5 minutes**, matching the OAuth `state`/`nonce` cookie TTL
already established elsewhere in this same flow, so there's one
re-auth-freshness convention rather than two. Not yet confirmed — the
actual acceptance criteria for consuming this marker belong in
`06-mfa-totp.md`, which should reference this default rather than
inventing its own.

**B. Post-redirect frontend destination isn't fully specified.**
Which concrete frontend route each `302` lands on (success dashboard,
specific error-state route per intent, etc.) is deferred to the
frontend track for this domain (`kencleng-agentic-workflow.md` §14) —
this spec only commits to *a* redirect with a distinguishable
success/error query param, not the exact route shape.

## Audit log entry?

**Partial.** Registration via Google (the `login` + no-existing-identity
+ new-user branch) is not in the Fitur 9 audit-log action list (same
reasoning as `01-register-email-verification.md` — self-service,
non-destructive on a new account). The `link` intent's successful
branch **is** in scope — Fitur 9 explicitly lists "account linking
baru" — write a `user_logs` entry there. `reauth` writes no audit
entry (it doesn't change any account state itself; the entry, if any,
belongs to whatever action it gates — e.g. MFA disable in
`06-mfa-totp.md`).

## References

- `docs/project/kencleng-phase0-detail.md` Fitur 1B, Fitur 4
- `docs/spec/account/threat-model.md` component 4
- `api/openapi.yaml` — `GET /auth/google/redirect`,
  `GET /auth/google/callback`