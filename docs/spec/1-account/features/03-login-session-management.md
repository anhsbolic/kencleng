# Feature Spec — Login & Session Management

> File: `docs/spec/account/features/03-login-session-management.md`
> Status: draft — all open items resolved, ready for human review
> Risk tier: 1 (with a Tier 0 fenced sub-area)
> Domain: account

## Endpoint

- `POST /auth/login`
- `POST /auth/login/mfa`
- `POST /auth/refresh`
- `POST /auth/logout`

Includes Fitur 2C (login attempt lockout) — folded in here rather
than a separate feature file, since it's not a separate endpoint, just
a check embedded in `POST /auth/login`.

## Acceptance criteria

### `POST /auth/login`

- Given valid `email_password` credentials, no MFA enrolled, and the
  identifier is not currently locked out, When submitted, Then a
  `login_attempts` row is recorded (`success=true`), an access token
  (JWT ES256, 15 min TTL) and a new refresh token (30 day TTL, fresh
  `family_id`) are issued, the refresh token is set as an `HttpOnly` +
  `Secure` + `SameSite=Strict` cookie, and the API responds `200`
  `LoginResponse`.
- Given valid credentials and MFA **is** enrolled, When submitted,
  Then a `login_attempts` row is recorded (`success=true` — credential
  check passed, independent of the MFA step per the flow ordering in
  `kencleng-phase0-detail.md` Fitur 2), **no** refresh cookie is set
  yet, and the API responds `200` `LoginMfaRequiredResponse`
  (`mfa_pending_token`).
- Given invalid credentials (wrong email or wrong password — same
  response for both), When submitted, Then a `login_attempts` row is
  recorded (`success=false`), and the API responds `401` with the
  generic "Email atau password salah" message (no distinction between
  the two failure reasons).
- Given the identifier is currently locked out (≥5 failed attempts in
  the trailing 15-minute window for that `identifier_hash`), When
  submitted, Then the lockout check happens **before** credential
  verification (no `login_attempts` row is written for a
  lockout-rejected attempt, since the credential itself was never
  checked), and the API responds `429` with the **same generic detail
  text** as the `401` case — only the status code differs, per the
  anti-enumeration rule already in `openapi.yaml`.
- Given credentials are valid but `AuthIdentity.verified_at IS NULL`
  (unverified `email_password` user), When submitted, Then login
  still succeeds at this endpoint — the verification restriction is
  enforced by other domains at the point of the restricted action
  (donating as registered, becoming a representative), not by
  blocking login itself.

### `POST /auth/login/mfa`

- Given the `mfa_pending_token`'s `user_id` is currently locked out at
  the **MFA stage** (≥5 failed MFA attempts in the trailing 15-minute
  window — same threshold/window as password lockout, see resolved
  Assumption C), When submitted, Then reject **before** checking the
  `totp_code`/`backup_code` at all, `429`, same generic detail
  pattern as the password-lockout case.
- Given a valid, unexpired `mfa_pending_token` and a correct
  `totp_code`, When submitted, Then a `login_attempts` row is recorded
  (`user_id` from the token, `stage='mfa'`, `success=true`), and the
  login completes exactly like the no-MFA branch of `/auth/login`
  (tokens issued, cookie set), `200` `LoginResponse`.
- Given a valid `mfa_pending_token` and a correct, **unused**
  `backup_code`, When submitted, Then the same as above (with the
  `login_attempts` row recorded the same way), **and** that backup
  code's `used_at` is set (INV-account-06) — the
  `mfa_totp_secrets.enabled_at IS NOT NULL` check in INV-account-06 is
  trivially satisfied here since this branch is only reachable when
  MFA is enrolled.
- Given an expired or malformed `mfa_pending_token`, When submitted,
  Then `401`, no tokens issued, no `login_attempts` row (identity
  isn't reliably known — the token itself is what's invalid).
- Given a valid `mfa_pending_token` but an incorrect `totp_code`/
  `backup_code`, When submitted, Then a `login_attempts` row is
  recorded (`user_id` from the token, `stage='mfa'`, `success=false`),
  `401`, no tokens issued.

### `POST /auth/refresh`

- Given a valid refresh token cookie that hasn't been rotated or
  revoked, When called, Then atomically (guarded `UPDATE ... WHERE
  replaced_by_id IS NULL AND revoked_at IS NULL`) a new access token
  is issued, a new refresh token is issued (same `family_id`), the old
  token's `replaced_by_id` is set, and the new token replaces the
  cookie. `200` `RefreshResponse`.
- Given no refresh cookie present, or an expired one, When called,
  Then `401`.
- Given a refresh token that was already rotated (`replaced_by_id IS
  NOT NULL`) being presented again, When called, Then reuse is
  detected: every token in that `family_id` is revoked
  (INV-account-04), and the API responds `401` — client must force a
  full re-login.
- Given two concurrent refresh requests using the **same currently-
  valid** (not-yet-rotated) token, Then exactly one succeeds (`200`);
  the other, having lost the atomic guard race, is treated
  **identically to a reuse-detection case** — the whole family gets
  revoked, forcing a full re-login even though this wasn't an actual
  attack. This is an accepted trade-off of the rotate-on-use design
  (INV-account-03/04), not a bug — see Assumption D.

### `POST /auth/logout`

- Given a valid refresh cookie, When called, Then that token's
  `revoked_at` is set, the cookie is cleared, `204`.
- Given no refresh cookie present (already logged out), When called,
  Then still `204` (idempotent — nothing to revoke, not an error
  condition).

## Error cases

| Condition | Expected response |
|---|---|
| Wrong email or password | `401`, generic message |
| Locked out (≥5 failed / 15 min) | `429`, identical generic message to `401` |
| MFA pending token expired/invalid | `401` |
| Wrong TOTP/backup code | `401` |
| Locked out at MFA stage (≥5 failed / 15 min, per `user_id`) | `429`, same generic detail pattern as password lockout |
| Refresh token missing/expired/revoked/reuse-detected | `401` |

## Applicable invariants

- `docs/spec/account/invariants.md#inv-account-03` — refresh rotation
  single-use.
- `docs/spec/account/invariants.md#inv-account-04` — reuse detection
  revokes the whole family.
- `docs/spec/account/invariants.md#inv-account-06` — backup code
  single-use + implicit invalidation tied to MFA enabled state.
- `docs/spec/account/invariants.md#inv-account-07` — MFA can't be
  half-enabled (relevant here since this endpoint is what actually
  checks `enabled_at` to decide whether to branch into the MFA step).

## Threat breakdown

Derived from `docs/spec/account/threat-model.md`, component 2, plus
one threat found while drafting this spec (not in the original
threat-model — see Assumption B):

| Threat | Mitigation at this endpoint's level | Test that proves it |
|---|---|---|
| Credential stuffing / brute-force login | `login_attempts` persistent lockout (5/15min, `stage=password`, keyed by `identifier_hash`) + in-memory rate limit | `TestLogin_Lockout_5Failed15Min` |
| MFA code brute-force at `/auth/login/mfa` | `login_attempts` persistent lockout (5/15min, `stage=mfa`, keyed by `user_id`) — resolved 2026-08-05, see Assumption C | `TestLoginMfa_Lockout_5Failed15Min` |
| Refresh token replay/reuse | Rotate-on-use + reuse detection (INV-account-03/04) | `TestRefresh_ReuseDetection_FamilyRevoked` |
| Concurrent refresh race | Atomic guarded `UPDATE`, exactly one winner | `TestRefresh_ConcurrentRequests_ExactlyOneWins` |
| Login error message distinguishing failure reason | Generic message for both wrong-email and wrong-password | `TestLogin_GenericErrorMessage` |
| Token issued without completing required MFA | Strict ordering — no code path issues tokens from `/auth/login` when MFA is enrolled | `TestLogin_MfaRequired_NoTokensIssuedYet` |
| `mfa_pending_token` type confusion — a token minted for the MFA-pending step being accepted by protected endpoints as if it were a real access token | **Two independent layers** (resolved 2026-08-05, see Assumption A & B): (1) separate HS256 signing key from the access token's ES256 key — a wrong-purpose token fails signature verification outright; (2) explicit `purpose` claim, checked by auth middleware as defense-in-depth on top of (1) | `TestAuthMiddleware_RejectsWrongSigningKey`, `TestAuthMiddleware_RejectsNonAccessPurposeToken` |

## Risk tier & rationale

**Tier 1**, with a **Tier 0 fenced sub-area**: JWT signing/
verification (the access token's ES256 keypair and the
`mfa_pending_token`'s separate HS256 secret — see Assumption A) and the
refresh-token rotate-on-use/reuse-detection logic. This matches the
Tier 0 examples in `kencleng-agentic-workflow.md` §13.2 verbatim
("JWT/TOTP, refresh-token reuse detection") — these specific files
must be human-authored/paired and marked "no-agent-write" in
`AGENTS.md` once the backend directory structure exists. The rest of
the endpoint (request parsing, `login_attempts` bookkeeping, response
shaping) is ordinary Tier 1 agent-generated work with mandatory
concurrency tests.

## Assumptions / open questions

**A & B — resolved together, 2026-08-05. What `mfa_pending_token`
is and why it needs its own signing key.**

After the password step succeeds but before MFA is verified, the
server needs to hand the client *something* that proves "this
specific user already passed the password check," so the next
request (`/auth/login/mfa`) doesn't have to re-collect the password.
HTTP being stateless, that "something" has to travel back from the
client. `mfa_pending_token` is that carrier.

**Design**: a stateless signed JWT, claims `{sub: user_id,
purpose: mfa_pending, exp}`, **5 minute TTL** (long enough to type a
6-digit code, short enough to bound the exposure window if it ends up
somewhere it shouldn't — browser history, a proxy log, etc.). Not
persisted to any DB table — the token's own signature and expiry are
the only things checked. This is safe *despite* not being
single-use, because knowing the token is never sufficient by itself:
every use still requires a correct `totp_code`/`backup_code` on top
of it.

**Why it needs a genuinely separate signing key from the access
token (not just a `purpose` claim on a shared key)**: a claim check
is application logic, and application logic can have bugs or missing
checks on some code path (that's exactly the "token confusion" risk
originally flagged in Assumption B). A **separate key** turns that
into a cryptographic guarantee instead of a logic guarantee — even if
the `purpose`-claim check is buggy or forgotten somewhere, a token
signed with the MFA-pending key will simply **fail signature
verification** against the access-token verifier, full stop, with no
custom logic involved. **Resolved: use a separate HMAC (HS256)
secret** (`MFA_PENDING_TOKEN_SECRET` — a new env var, distinct from
the access token's ES256 keypair) for this token specifically. HS256
is intentionally chosen over ES256 here (asymmetric isn't needed —
nothing outside this one backend process ever needs to verify this
token, unlike the access token which is designed to be verifiable by
other future services per `kencleng-backend-tech-stack.md`). The
`purpose` claim is **kept as well**, as defense-in-depth on top of
the key separation, not a replacement for it — belt and suspenders,
not either/or.

This lands `mfa_pending_token` handling inside the same Tier 0 fenced
sub-area as the access token (both are JWT signing/verification core
logic), even though it uses a different key.

**C. Resolved — 2026-08-05. MFA code brute-force now uses the same
persistent lockout mechanism as password login, not just the generic
in-memory rate limiter.** `login_attempts` gains a `stage` column
(`password` / `mfa`) — see the `openapi.yaml`/ERD follow-up note
below. Lockout is computed the same way as the password stage (≥5
failed in a trailing 15-minute window), but scoped differently:
password-stage lockout is keyed by `identifier_hash` (identity isn't
known yet at that point), MFA-stage lockout is keyed by `user_id`
(already known and verified via the validated `mfa_pending_token` by
the time this check runs). The check happens **before** verifying the
submitted code, mirroring the password-stage ordering in Fitur 2.

**Requires an ERD change**: `login_attempts` needs a new `stage`
column and a new index for the `user_id`-keyed MFA-stage lookup
(distinct from the existing `identifier_hash`-keyed one, since the
two stages are looked up by different keys). See the updated
`login_attempts` DDL below — this needs to be applied to
`kencleng-erd.md` as a follow-up, alongside the other flagged
`openapi.yaml` changes from earlier feature specs.

```sql
CREATE TABLE login_attempts (
    id               UUID PRIMARY KEY,
    identifier_hash  TEXT NOT NULL,
    user_id          UUID REFERENCES users(id) ON DELETE SET NULL,
    stage            TEXT NOT NULL DEFAULT 'password'
                       CHECK (stage IN ('password', 'mfa')),
    success          BOOLEAN NOT NULL,
    attempted_at     TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- password-stage lockout: identity not yet known, keyed by identifier_hash
CREATE INDEX ix_login_attempts_identifier_time
    ON login_attempts (identifier_hash, attempted_at DESC);

-- MFA-stage lockout: identity already known via the validated
-- mfa_pending_token, keyed by user_id instead
CREATE INDEX ix_login_attempts_user_stage_time
    ON login_attempts (user_id, stage, attempted_at DESC)
    WHERE user_id IS NOT NULL;

CREATE INDEX ix_login_attempts_attempted_at_brin
    ON login_attempts USING BRIN (attempted_at);
```

For MFA-stage rows, `identifier_hash` is still populated (derived
from the user's own `email_password` identifier, looked up via
`user_id`) purely for schema consistency with the `NOT NULL`
constraint — it's not used for the MFA-stage lockout query itself.

**D. Resolved — 2026-08-05.** Accepted: the backend's rotate-on-use +
reuse-detection design (INV-account-03/04) stays strict, unweakened —
no grace window. The multi-tab race is a **frontend-track concern**,
to be solved with cross-tab coordination (`BroadcastChannel` — one
tab acts as the single "refresher," others wait and receive the
rotated access token via the channel instead of independently calling
`/auth/refresh`), not a backend change. This is deferred to when the
`account` domain's frontend track starts
(`kencleng-agentic-workflow.md` §14/§15 — backend-then-frontend
sequential for this first domain) — noted here now so it isn't
rediscovered from scratch later. No action needed in this backend
feature spec beyond this note.

## Audit log entry?

No — per `kencleng-phase0-detail.md` Fitur 9 scope, login/logout/
refresh are not in the audit-logged action list. The `login_attempts`
table already serves as this domain's own record of login activity,
independent of `user_logs`.

## References

- `docs/project/kencleng-phase0-detail.md` Fitur 2, Fitur 2C
- `docs/project/kencleng-backend-tech-stack.md` — Auth Mechanism row
- `docs/spec/account/threat-model.md` component 2
- `api/openapi.yaml` — `POST /auth/login`, `POST /auth/login/mfa`,
  `POST /auth/refresh`, `POST /auth/logout`