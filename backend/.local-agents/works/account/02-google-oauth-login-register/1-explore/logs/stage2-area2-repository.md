# Stage 2 — Area 2: Domain repository layer

> Feature: 02-google-oauth-login-register
> Date: 2026-08-22

## Current state

`repository.go` defines `Repository` interface with 7 methods:
- `InsertUser(ctx, tx, user)` — encrypts PII, inserts
- `InsertAuthIdentity(ctx, tx, identity)` — encrypts PII, inserts
- `InsertAuthToken(ctx, tx, token)` — inserts
- `FindAuthIdentityByIdentifierHash(ctx, providerType, identifierHash)` — lookup by (provider, hash)
- `FindAuthTokenByHash(ctx, tokenHash)` — lookup by hash
- `RedeemToken(ctx, tx, tokenHash)` — atomic redeem with 3-clause guard
- `SetUserVerified(ctx, tx, userID, providerType, verifiedAt)` — set verified_at
- `RevokeTokens(ctx, tx, userID, purpose)` — revoke unused tokens

`repository_db.go` implements all using goqu query builder. InsertAuthIdentity handles encryption of Identifier and HMAC computation, then clears plaintext after insert.

## Requirement

Google OAuth callback needs:
1. Look up Google identity by email hash — already exists
2. Look up email_password identity by email hash for conflicts — already exists
3. Create new User + AuthIdentity for new Google registrations — already exists
4. Attach Google identity to existing user (link intent) — InsertAuthIdentity already handles this
5. Issue tokens — gap (see Area 1)

## Gap

1. **No FindUserByID.** link intent needs to verify session's user_id exists. Could rely on foreign key constraint, but explicit validation is cleaner per AGENTS.md golden rules.
2. **No new repository methods strictly needed for core OAuth flow.** Existing FindAuthIdentityByIdentifierHash, InsertUser, InsertAuthIdentity cover data access patterns. Main gap is at service/platform level.

## Sniffing

- **Risk:** Low — repository layer is already well-factored for Google use case.
- **Edge case:** link intent's "email already claimed by different user" check requires looking up both Google and email_password identities by email hash. Service needs two separate FindAuthIdentityByIdentifierHash calls.
- **Misleading signal:** InsertAuthIdentity takes pgx.Tx, suggesting it must always run in caller-managed transaction. For link intent (single insert), transaction may not be strictly necessary but is still correct for consistency.
- **Inconsistency:** None found.
