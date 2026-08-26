# Gap Analysis — Area 2: `internal/domain/account/`

> Files: `entity.go`, `service.go`, `google_oauth.go`, `repository.go`,
> `repository_db.go` (+ test files noted)

## Current state

- **Entities** (`entity.go`): `User`, `AuthIdentity` (`CredentialSecret *string`
  = bcrypt hash, `VerifiedAt`), `AuthToken`, `RefreshToken`, `UserLog`.
  `RefreshToken` already carries `FamilyID`, `RevokedAt`, `ReplacedByID`;
  doc comment says rotation/reuse detection is *"implemented in the
  login/session task; this ticket only issues first-generation tokens."*
- **Service** (`service.go`): `Register`, `VerifyEmail`, `ResendVerification`
  only. House patterns: HMAC identifier lookup (`crypto.HMAC([]byte(email),
  s.keys)`, :143), tx via `TxRunner` seam, sentinel errors
  `ErrValidation`(422)/`ErrTokenExpired`(410)/`ErrTokenNotFound`(404),
  `isUniqueViolation` via SQLSTATE 23505, strict PII-free logging, deliberate
  anti-enumeration timing shaping (`dummyWrite`, always-run-bcrypt R7).
  Password hashing via `secrets.HashPassword` (:137). **No password-verify
  call exists anywhere.**
- **Google OAuth** (`google_oauth.go`): `IssueTokens(ctx, userID)` (:483) is
  the only token issuer — ES256 access JWT claims `{sub, iat, exp}`
  (**no `purpose` claim**, :489-493) + 32-byte-hex refresh token stored as
  SHA-256 hash with fresh `family_id`; insert in its own tx (:523). TTLs:
  15 min / 30 days (:48-52).
- **Repository port** (`repository.go`): `InsertRefreshToken` exists; **no
  read/rotate/revoke methods for refresh_tokens**. No `LoginAttempt`
  anything. `RedeemToken` (repository_db.go:299) is the established atomic
  guarded-redemption pattern: `UPDATE … WHERE <guards> RETURNING`, `ok=false`
  on 0 rows.

## Requirement

Credential verify + generic-failure semantics; `login_attempts` write + lockout
query (2 stages × 2 key types); rotate-on-use guarded UPDATE + family
revocation; logout revoke; `mfa_pending_token` mint/verify;
`LoginResponse.user` assembly (plaintext `email`, `email_verified`, `roles[]`,
`auth_providers[]`, `mfa_enabled`).

## Gap

1. No `Login`/`Refresh`/`Logout`/`LoginMfa` service methods.
2. No repository methods: find-refresh-token-by-hash, guarded rotate,
   revoke-one, revoke-family-by-family_id, insert-login-attempt,
   count-recent-failures (×2 key shapes), identity-by-user_id (needed for
   MFA-stage `identifier_hash` backfill per Assumption C).
3. No `LoginAttempt` entity; no MFA entities (`mfa_totp_secrets`,
   `mfa_backup_codes`) — `/auth/login/mfa` and `mfa_enabled` depend on
   feature-06 territory that doesn't exist yet.
4. Access token has no `purpose` claim — Assumption B's defense-in-depth
   layer 2 has nothing to check today.
5. Repository cannot assemble `LoginResponse.user`: no decrypt-on-read path
   for `primary_email` (doc says decryption "is not on the hot path" — never
   built), no `user_roles` / provider-aggregation / `mfa_enabled` source.

## Sniffing findings

- **Misleading signal (Registrar-pattern):** `RefreshToken.FamilyID/
  ReplacedByID` exist and `IssueTokens` populates `family_id` — skimming
  suggests rotation is "mostly there"; in reality the table is insert-only,
  zero rotation/reuse paths end-to-end.
- **Misleading signal #2:** `s.authKeys` wired into Service makes JWT signing
  look complete; missing `purpose` claim means middleware half of the two-layer
  mitigation is unimplementable without touching minting (Tier 0 concern).
- **Risk:** `IssueTokens` commits its own tx independent of callers — fine for
  OAuth callback, but login wants `login_attempts(success=true)` recorded
  regardless of MFA branch; ordering/atomicity between attempt-write and
  token-issue needs explicit decision → Stage 3 D3/D4.
- **Edge-case precedent:** concurrent-refresh exactly-one-winner will mirror
  `RedeemToken`'s guarded `UPDATE…RETURNING`; no equivalent exists for
  refresh_tokens yet.
- **Inconsistency:** domain sentinel vocabulary lacks 401/429 equivalents —
  transport has no unauthorized/locked-out sentinel to translate (Area 4).
- **Observation (one line):** register's always-run-bcrypt discipline implies
  wrong-email login should burn a dummy bcrypt compare too (Stage 3 D4).
