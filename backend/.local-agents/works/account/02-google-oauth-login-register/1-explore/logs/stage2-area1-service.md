# Stage 2 — Area 1: Domain service layer

> Feature: 02-google-oauth-login-register
> Date: 2026-08-22

## Current state

`service.go` defines `Service` with dependencies: `repo Repository`, `tx TxRunner`, `breachCheck breachChecker`, `email notification.Sender`, `keys *crypto.Keys`.

Existing methods:
- `Register(ctx, name, email, password)` — handles 4 branches including Google-only conflict detection (line 134: `FindAuthIdentityByIdentifierHash(ctx, providerGoogle, identifierHash)`)
- `VerifyEmail(ctx, token)` — redeem + set-verified in one tx
- `ResendVerification(ctx, email)` — revoke + issue new token
- Helper: `registerNewUser`, `issueNewVerificationToken`, `dummyWrite`, `validatePassword`, `sendVerification`, `sendNudge`

Constants already defined: `providerGoogle = "google"` (line 47).

The `Register` method already does a Google identity lookup (line 134) — the repository layer already supports finding a Google identity by identifier hash.

`service_test.go` uses `fakeRepo` with `fakeTx` — comprehensive test double pattern with hooks for error injection, call recording, and configurable redeem modes.

## Requirement

Two new endpoints:
1. `GET /auth/google/redirect` — generates state+nonce, stores in HttpOnly cookie, 302 to Google consent screen. For link/reauth intents, requires existing session and encodes user_id into cookie.
2. `GET /auth/google/callback` — validates state/nonce, exchanges code with Google, verifies id_token, then branches on intent (login/link/reauth with multiple sub-branches).

## Gap

1. **No Google OAuth client dependency.** No way to exchange authorization code for tokens or verify id_token (signature against Google's JWKS, issuer, audience, expiry, nonce).
2. **No cookie management.** No mechanism to set/read short-TTL HttpOnly cookies for state/nonce/user_id.
3. **No session/auth middleware integration.** link/reauth intents require reading current session's user_id from access token. Existing handlers are all public.
4. **No token issuance in the service.** Register flow doesn't issue JWT access/refresh tokens — only creates user and sends verification email. Google OAuth callback needs to issue tokens on successful login.
5. **No FindUserByID in the repository.** link intent needs to verify session's user_id exists.
6. **No re-auth marker mechanism.** reauth intent needs short-lived server-side marker. No such entity or storage exists.

## Sniffing

- **Risk:** Token issuance dependency is the biggest structural risk. Task #2 needs tokens but token issuance belongs to task #3. Cross-task dependency needs explicit handling.
- **Edge case:** Concurrent Google login attempts for same email — both callbacks could race past "no existing identity" check. INV-account-01's unique index catches this, but error handling needs to be clean.
- **Miscontext:** Feature spec says "issue access + refresh tokens (Fitur 2)" but current codebase has no token issuance. Spec assumes this exists.
- **Misleading signal:** Register method already looks up Google identities (line 134), suggesting "Google support is partially there." But that lookup is only for conflict detection during email-password registration.
- **Inconsistency:** Feature spec's callback table says login + existing Google identity → "issue tokens (Fitur 2)" but Fitur 2 doesn't exist yet. Cross-task dependency needs reconciliation.
