# Feature Spec — Register & Email Verification

> File: `docs/spec/account/features/01-register-email-verification.md`
> Status: draft — all open items resolved, ready for human review
> Risk tier: 1
> Domain: account

## Endpoint

- `POST /auth/register`
- `POST /auth/verify-email`
- `POST /auth/verify-email/resend`

## Acceptance criteria

### `POST /auth/register`

- Given an email not yet registered under `email_password`, and a
  password ≥8 chars not found in the HaveIBeenPwned breach-list (or
  the breach-check API is unreachable — fail-open), When the user
  submits the registration form, Then a new `User` + `AuthIdentity`
  (`provider_type=email_password`, `verified_at=null`) is created, a
  single-use verification token (`auth_tokens`,
  `purpose=email_verification`, expires in 24h) is generated, a
  verification email is sent, and the API responds `202` with a
  generic accepted message (no `user_id` — see Assumption A).
- Given an email already registered under `email_password` and
  **unverified**, When the user submits the registration form again,
  Then no new `User`/`AuthIdentity` is created; instead the same
  internal action as `verify-email/resend` fires (old unused token
  revoked, new one issued, resend-verification email sent), and the
  API responds identically (`202`, same generic message) to the
  new-registration case.
- Given an email already registered under `email_password` and
  **verified**, When the user submits the registration form, Then no
  new record is created, a "you already have an account — forgot your
  password?" nudge email is sent instead, and the API responds
  identically (`202`, same generic message).
- Given a password that fails the length policy or is found in the
  breach-list, When submitted, Then `422` Validation Error — this
  check happens **before** any enumeration-sensitive branching, so it
  doesn't leak whether the email is registered either.
- Given an email that is already registered, but **only** under a
  `google` `AuthIdentity` (no `email_password` identity exists for
  that `User` yet), When the user submits the registration form,
  Then **no new `User` is created** (rejected, matching the
  symmetrical rule already in Fitur 1B/4 — see Assumption C,
  resolved 2026-08-05); instead of a distinguishing status code, this
  follows the **same generic-response pattern** as the other
  already-registered branches — a nudge email is sent ("this email
  is linked via Google — log in with Google, then use 'Atur Password'
  to add a password," mirroring the Google-only notice already
  specced in Fitur 2B's forgot-password flow), and the API responds
  identically (`202`, same generic message) to every other branch.
  Using a distinguishing status code here would re-introduce exactly
  the enumeration leak this endpoint's generic-response redesign was
  meant to close — a mistake caught while drafting this spec (see
  Assumption C).

### `POST /auth/verify-email`

- Given a valid, unexpired, unused token, When submitted, Then
  `AuthIdentity.verified_at` is set to now, response `200`.
- Given an expired token, When submitted, Then `410` Token Expired, no
  state change.
- Given a token that doesn't exist or was already used, When
  submitted, Then `404`, no state change.
- Given the same valid token submitted twice concurrently, Then
  exactly one request succeeds (guarded `UPDATE ... WHERE used_at IS
  NULL AND expires_at > now()`, per INV-account-08); the other gets
  `404`.

### `POST /auth/verify-email/resend`

- Given an email matching an existing **unverified**
  `email_password` identity, When resend is requested, Then the
  previous unused token(s) for that identity are revoked
  (`revoked_at`), a new token is issued, a new verification email is
  sent, and the API responds `202` generic (existing behavior,
  already anti-enumeration by design per `openapi.yaml`).
- Given an email that doesn't match any account, matches an
  already-verified account, or matches only a `google` identity, When
  resend is requested, Then no new token is created and no email is
  sent, but the API response is identical (`202` generic) to the
  success case.
- Rate limit: stricter `/auth/*` limit applies (mitigates
  verification-email flood).

## Error cases

| Condition | Expected response |
|---|---|
| Password fails length policy or found in breach-list | `422` Validation Error |
| Verification token expired | `410` Token Expired |
| Verification token not found / already used | `404` |
| Too many requests (any of the 3 endpoints) | `429` |
| Register / resend, any enumeration-sensitive branch | `202` generic (never a distinguishing status/message — see Assumption A) |

## Applicable invariants

- `docs/spec/account/invariants.md#inv-account-01` — uniqueness is
  per-provider, not global; concurrent duplicate `email_password`
  registration attempts for the same email must not both succeed.
- `docs/spec/account/invariants.md#inv-account-08` — `auth_tokens`
  single-use and time-bound; applies to both the verification token
  and the resend-issued replacement.

## Threat breakdown

Derived from `docs/spec/account/threat-model.md`, component 1:

| Threat | Mitigation at this endpoint's level | Test that proves it |
|---|---|---|
| Concurrent duplicate registration for the same email | DB unique index `(provider_type, identifier_hash)`, INV-account-01 | `TestRegister_ConcurrentDuplicateEmail_Race` |
| Email enumeration via distinguishable register/resend responses | Uniform `202` generic response + equivalent-cost internal branching (Assumption B) | `TestRegister_GenericResponse_AllBranches`, plus a response-timing assertion |
| Verification token replay/double-submit | Guarded `UPDATE ... WHERE used_at IS NULL AND expires_at > now()` | `TestVerifyEmail_TokenSingleUse_Concurrent` |
| Verification-email flood via resend spam | Stricter `/auth/*` rate limit | `TestResend_RateLimited` |
| Weak/breached password accepted | Length policy + HaveIBeenPwned k-anonymity check, fail-open on API outage | `TestRegister_PasswordPolicy`, `TestRegister_BreachCheck_FailOpen` |
| Enumeration via the Google-only-conflict branch specifically | Same generic `202` response as every other branch, notice sent by email instead of by distinguishing status (Assumption C) | `TestRegister_GoogleOnlyConflict_GenericResponse` |

## Risk tier & rationale

**Tier 1** — INV-account-01 requires a genuine concurrency/race test
(project goal #2: concurrency-safe code), and the anti-enumeration
resolution requires a correctness test that response shape *and*
timing don't leak internal state. No Tier 0 sub-area here (password
hashing via bcrypt/argon2 is standard library-backed, not in the same
class as JWT/TOTP core logic called out in `kencleng-agentic-workflow.md`
§13.2).

## Assumptions / open questions

**A. `POST /auth/register` response shape changes from what's
currently in `openapi.yaml`.** Today's spec has `201` +
`RegisterResponse` (`user_id` + message) on success and a distinct
`409` on duplicate email. Per the anti-enumeration decision resolved
in `docs/spec/account/threat-model.md` (2026-08-05), this becomes a
uniform `202` + `GenericAcceptedMessage`, with **no `user_id` in the
response** (there's no new ID to return in the duplicate-email case,
and returning one only sometimes would itself be a signal). Frontend
must get the user's identity later, via the login/verify flow, not
from this response. **Requires an `openapi.yaml` edit** — flagged
here and in `docs/spec/account/threat-model.md`, not yet applied.

**B. Constant-time handling is a Build-stage implementation detail,
not fully specified here.** The three internal branches (new user /
resend nudge / password-reset nudge) must take equivalent wall-clock
time and do DB-write-shaped work, so a timing side-channel doesn't
leak which branch ran. Exact mechanism (e.g. always performing one
write-shaped operation, or a dummy password-hash computation on the
no-op branches) is left to the implementing agent, who must record
the chosen approach in a risk note per `kencleng-agentic-workflow.md`
§9 — this is exactly the kind of thing that must be reported, not
silently assumed correct.

**C. Resolved — 2026-08-05.** `kencleng-phase0-detail.md` Fitur 1
and Fitur 1B never documented what happens when someone registers
`email_password` using an email that already belongs to a
**Google-only** account. Resolved as: **reject, matching the
symmetrical rule already in Fitur 1B/4** — no new `User` is created;
the user is nudged (via email, not a distinguishing API response — see
the acceptance criteria above) to log in with Google and use the
already-specced "Atur Password" flow (Fitur 4,
`POST /account/security/set-password`) to add an `email_password`
identity to their existing account. This achieves "one real-world
email → one account" **without** auto-merging based on an email match
claimed by a different provider — auto-merge was rejected because it
would let anyone who controls a Google-side claim to that email
(e.g. a Google Workspace admin provisioning an address, who never
proved control of the actual inbox to *our* system) silently gain
access to an existing `email_password` account. This mirrors why
Fitur 1B already blocks the reverse direction. No new flow is needed
— Fitur 4's existing set-password sub-flow already does exactly this.

**First draft of this resolution used a distinguishing `409`
response for this branch, which was caught and corrected during
spec review** — it would have reintroduced the exact email-enumeration
leak this endpoint's generic-response redesign (Assumption A) was
meant to close. Left in as a reminder for the Build stage: any new
branch added to an anti-enumeration endpoint later must be checked
against this same failure mode, not just against its own logic in
isolation.

## Audit log entry?

No — per `kencleng-phase0-detail.md` Fitur 9 scope, registration and
email verification are not in the audit-logged action list (they're
self-service, non-destructive actions on the actor's own new account).

## References

- `docs/project/kencleng-phase0-detail.md` Fitur 1
- `docs/project/kencleng-erd.md` §1 (`users`, `auth_identities`,
  `auth_tokens`)
- `api/openapi.yaml` — `POST /auth/register`, `POST /auth/verify-email`,
  `POST /auth/verify-email/resend`