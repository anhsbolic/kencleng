# Stage 2 — Gap Analysis, Area 1: Account domain core

> Files: `internal/domain/account/entity.go`, `repository.go`,
> `repository_db.go`, `service.go`

## Current state (concrete)

- **Entities** (`entity.go`): `AuthIdentity{ID, UserID, ProviderType,
  Identifier, CredentialSecret *string, VerifiedAt *time.Time}` —
  exactly the shape the feature needs (verified_at present, google
  identities have nil secret). PII pattern: plaintext `Identifier` is
  caller-set, adapter encrypts to BYTEA + HMAC at insert, then clears
  plaintext. Lookups are always by `(provider_type, identifier_hash)` —
  never plaintext.
- **Repository port** (`repository.go`) today offers:
  `InsertAuthIdentity`, `FindAuthIdentityByIdentifierHash(providerType,
  identifierHash)`, `SetUserVerified(userID, providerType, verifiedAt)`,
  `RedeemToken` (3-clause guarded, tx-scoped), `RevokeTokens(userID,
  purpose)`, refresh-token ops (`RevokeRefreshTokenByHash`,
  `RevokeRefreshTokenFamily(familyID)` — note: **family-scoped, not
  user-scoped**), `InsertUserLog(tx)`.
- **Service** (`service.go`): `Register` (4-branch anti-enumeration
  dispatch incl. `dummyWrite` timing-equivalence no-op), `VerifyEmail`
  (atomic redeem+set-verified in one tx via `TxRunner` seam),
  `ResendVerification`, `validatePassword` (≥8 chars + breach-check
  fail-open), `issueNewVerificationToken` (revoke-old + insert-new in
  one tx). Sentinel errors `ErrValidation`(→422), `ErrTokenExpired`
  (→410), `ErrTokenNotFound`(→404). Constants `providerEmailPassword`,
  `providerGoogle`, `purposeEmailVerify` already exist.
  **No `SetPassword` or `UnlinkGoogle` method exists anywhere in the
  service.**
- Established transaction idiom: service begins tx via `TxRunner`
  interface, passes `pgx.Tx` into repo methods, deferred rollback,
  commit flag; unique violations detected via `isUniqueViolation`
  (SQLSTATE 23505).

## Requirement vs Gap

| Feature need | Existing support | Gap |
|---|---|---|
| Branch selection: does caller have an `email_password` identity? | Only `FindAuthIdentityByIdentifierHash` (keyed by email hash, not user); also `FindIdentifierHashByUserAndProvider(userID, providerType)` returning `(hash, found, err)` from the MFA slice | A per-user provider lookup exists in a narrow form; no general identity-list-by-user method |
| Unlink guard (INV-account-02/12): count other identities with `verified_at IS NOT NULL` | `GetLoginUserView` reads all identities but as a login read-model | No query returning a user's identities with verified flags suitable for the guard |
| Unlink action: hard-delete the `google` identity row | Nothing | No `DeleteAuthIdentity` (or guarded atomic check-then-delete op) exists |
| Branch 1: create unverified `email_password` identity on existing user | `InsertAuthIdentity` works; `issueNewVerificationToken` covers token issuance | No service flow wires these together for an existing user (registration path always creates a User too) |
| Branch 2: update `credential_secret` in place | Nothing | No `UpdateCredentialSecret(userID, hash)` repo method |
| Both branches / unlink re-auth: compare password | `compare` seam (`secrets.ComparePassword`) exists from login flow | Reusable as-is |
| Branch 2 / reset-password pattern: revoke **all** refresh tokens of a user (INV-account-05) | Only family-scoped or single-hash revocation | **No revoke-all-by-user_id operation** — confirmed in Area 3 that Fitur 04 is unimplemented, so nothing to reuse at all |
| Audit log | `InsertUserLog` exists, tx-scoped; `actionAccountLinking` constant already exists in google_oauth.go | Need to confirm what action_type values are legal today (Area 6: unconstrained TEXT) |

## Sniffing

- *Misleading signal*: `GetLoginUserView` already computes
  `EmailVerified` ("any email_password identity has verified_at set")
  and `AuthProviders` — skimming suggests "we can already see a user's
  identities," but it's a decrypted login read-model, not usable as the
  unlink guard primitive.
- *Misleading signal*: `InsertUserLog`'s doc comment cites "attaching a
  Google AuthIdentity on link intent" as its motivating example — audit
  plumbing for linking-adjacent actions is exercised in
  `google_oauth.go`.
- *Risk*: `SetUserVerified` updates **all** rows matching `(user_id,
  provider_type)` — safe under one-identity-per-provider assumption,
  which INV-account-01 guarantees; but Branch 2's in-place secret update
  will need the same single-row discipline.
- *Inconsistency (minor)*: `registerNewUser`/`VerifyEmail` use raw
  `time.Now()` while lockout paths deliberately use the `s.now` clock
  seam — inconsistent clock discipline within one file.
- *Edge case*: `FindAuthIdentityByIdentifierHash` returns `(nil, nil)`
  on no rows — callers must distinguish; a convention new methods
  should follow.
