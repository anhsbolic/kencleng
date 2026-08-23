# Stage 2 — Area 6: API contract

> Feature: 02-google-oauth-login-register
> Date: 2026-08-22

## Current state

api/openapi/account.yaml defines:
- GET /auth/google/redirect (line 308-330): security: [] (public), intent query param (required, enum: login/link/reauth), response 302.
- GET /auth/google/callback (line 332-363): security: [] (public), code and state query params (required), response 302.

The openapi.yaml description for the callback (line 336-348) matches the feature spec's branching table.

## Requirement

Verify the OpenAPI spec's expectations match the feature spec and codebase capabilities.

## Gap

1. **security: [] on both endpoints.** Spec marks both as public. Correct for intent=login, but intent=link/reauth requires auth. OpenAPI doesn't distinguish — conditional auth is an implementation detail.
2. **No response schema for 302.** Just description text. Error redirect query params (e.g., ?error=no_auto_merge) aren't specified in OpenAPI.
3. **No cookie schema.** state/nonce cookie not documented in OpenAPI (normal — cookies aren't typically in OpenAPI).

## Sniffing

- **Risk:** Lack of error redirect param specification means frontend team doesn't have a contract for error states. Feature spec acknowledges this in Assumption B.
- **Edge case:** 302 response could be success redirect (tokens set) or error redirect (no tokens). Client needs to distinguish. Without specified error param, frontend would need to check for auth cookies or parse URL.
- **Miscontext:** OpenAPI security: [] means "no security scheme defined," not "no auth check at all." Auth check is handler-level.
- **Misleading signal:** OpenAPI callback description lists all branching behavior inline, suggesting endpoint is fully specified. But redirect destinations and error params are deferred to frontend track.
- **Inconsistency:** Feature spec says 302 to /login with error param for no-auto-merge case. OpenAPI description says same. Neither specifies param name or value.
