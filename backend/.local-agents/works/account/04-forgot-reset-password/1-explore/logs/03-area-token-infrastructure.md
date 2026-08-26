# Area 3 — Token infrastructure

> Stage 2 gap analysis. Files: `internal/domain/account/entity.go`,
> `repository.go`, `repository_db.go`; migrations
> `000003_create_auth_tokens.up.sql`, `000004_create_refresh_tokens.up.sql`.

## Current state

- **Schema already ready** — no migration needed: `auth_tokens.purpose`
  CHECK constraint **already includes `'password_reset'`** (migration
  000003:4); indexes: unique `token_hash`, `(user_id, purpose)`, partial
  `ix_auth_tokens_valid` on un-used/un-revoked. `refresh_tokens` (000004)
  has `family_id` and partial active-token index on `(user_id)`.
- **Entities**: `AuthToken{UserID, Purpose, TokenHash(SHA-256 hex),
  ExpiresAt, UsedAt, RevokedAt}` (entity.go:59–68) — doc comment already
  reserves `"password_reset"`; `RefreshToken` with rotation/reuse fields.
- **Repository port**: `RedeemToken(ctx, tx, tokenHash)` — atomic guarded
  UPDATE with the **full 3-clause predicate** (`used_at IS NULL AND
  revoked_at IS NULL AND expires_at > now()`), `RETURNING user_id,
  purpose`, tx-aware; its doc comment explicitly records that the
  invariant doc's *Verification* field omits `revoked_at` and declares it
  a **known documented spec error** ("techplan §14 Open Item #2") — use
  the Statement's 3-clause version, do not edit the spec. Also:
  `FindAuthTokenByHash` (expired-vs-other disambiguation),
  `RevokeTokens(userID, purpose)` (auth_tokens resend path),
  `InsertRefreshToken`, `RotateRefreshToken` (guarded, tx-atomic),
  `RevokeRefreshTokenByHash` (logout), `RevokeRefreshTokenFamily(familyID)`
  (reuse detection), `FindAuthIdentityByIdentifierHash(providerType,
  identifierHash)` returning the row incl. `credential_secret`.
- **Git ground truth**: commit `16a4bf9` "[BE] account task 03" — the
  login/session slice is live on main.

## Requirement

Reset flow needs: consume token + update `credential_secret` + revoke all
user's refresh tokens in ONE tx (INV-05). Forgot flow needs: distinguish
email_password / google-only / no-account branches via hash lookup.

## Gap

1. **No user-scoped mass session revoke.** Existing revokes are per-token
   (`RevokeRefreshTokenByHash`) or per-family (`RevokeRefreshTokenFamily`).
   INV-05 requires "every `refresh_tokens` row for that `user_id`".
2. **No credential update method.** Nothing writes
   `auth_identities.credential_secret` after insert.
3. Everything else (generate/redeem/disambiguate, identity lookup) is
   reusable as-is.

## Sniffing findings

1. **Risk** — new methods extend an interface shared across account
   slices; but zero schema changes keeps this inside S1 without migration
   collision risk.
2. **Edge cases** — `RedeemToken` returns uniform `ok=false` for
   not-found/used/revoked/expired; 404-vs-410 mapping depends on a
   follow-up `FindAuthTokenByHash`. Inherent race: a concurrent
   double-submit loser sees the token as *used* → maps to 404 — which
   coincidentally matches the spec's expected double-submit outcome.
3. **Miscontext** — Area 1's missing-`revoked_at` finding turns out
   **already adjudicated**: repo documents it as a known spec error and
   mandates the 3-clause guard; the invariant doc itself stays unedited
   per §4 authority rules.
4. **Misleading signals** — `RevokeTokens(userID, "password_reset")`
   looks like forgot-password housekeeping support, but using it would
   violate resolved Assumption A (no proactive revocation). Trap-shaped
   API surface for exactly this slice.
5. **Inconsistency** — tasks.md tracker said task #3 "build not started"
   while git shows it committed (stale tracker, noted not corrected).
6. **Minor observation** — `RedeemToken`/`RevokeTokens` call `time.Now()`
   inside the adapter while `CountRecentFailedAttempts*` take `since`
   from the caller's clock seam; two time conventions coexist in one file.
