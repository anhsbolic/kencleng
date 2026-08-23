# Stage 2 — Area 4: Platform dependencies

> Feature: 02-google-oauth-login-register
> Date: 2026-08-22

## Current state

`backend/internal/platform/`:
- `auth/keys.go` — Keys struct with Private/Public ecdsa keys. Load(privatePath, publicPath) reads PEM files. Used for JWT signing/verification (ES256).
- `crypto/` — encryption (AES-GCM) and HMAC. Used by repository for PII encryption.
- `notification/` — Sender interface. FakeSender and DevSender implementations.
- `breachcheck/` — HaveIBeenPwned client.
- `db/` — pgxpool wrapper.
- `secrets/` — password hashing (bcrypt).

No Google OAuth client exists. No JWKS verification library. No cookie utility.

## Requirement

1. Google OAuth client — exchange authorization code for tokens (POST https://oauth2.googleapis.com/token), fetch Google's JWKS, verify id_token signature, issuer, audience, expiry, nonce claim.
2. Cookie utility — set/read short-TTL HttpOnly cookies with structured data.
3. New env vars — GOOGLE_CLIENT_ID, GOOGLE_CLIENT_SECRET, GOOGLE_REDIRECT_URI.

## Gap

1. **No Google OAuth client.** Largest platform gap. Needs: code exchange, JWKS fetch with cache-on-miss, id_token verification (RS256 against JWKS, iss, aud, exp, nonce).
2. **No cookie utility.** Need helper for structured cookies with HttpOnly, Secure, SameSite, Path, Max-Age.
3. **No GOOGLE_REDIRECT_URI in env validation.** cmd/server/main.go's requireEnv doesn't include Google-specific env vars.

## Sniffing

- **Risk:** JWKS verification is most security-critical. Incorrect signature verification (wrong algorithm, missing audience check) would allow token forgery. Spec says "library-backed, not hand-rolled" — but correct configuration is still required.
- **Edge case:** Google's JWKS keys rotate. Cached JWKS might be stale. Need refresh-on-miss behavior, not fetch-once-at-startup.
- **Edge case:** Clock skew between server and Google. Freshly-issued id_token might fail exp check if server clock is ahead. Standard practice: 60-second tolerance.
- **Miscontext:** Feature spec says "standard JWT verification against Google's JWKS" — but Google uses RS256 (RSA), not ES256 (ECDSA which this codebase uses). platform/auth/keys.go loads ECDSA keys for app's own JWTs — completely separate from Google's RSA verification.
- **Misleading signal:** Existing platform/auth/keys.go might suggest "auth key infrastructure exists, just extend it." But it's specifically for ES256 signing — nothing to do with Google's RS256 verification.
- **Inconsistency:** None found. Platform layer is cleanly separated — Google OAuth is genuinely new infrastructure.
