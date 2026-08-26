# Task 02: Domain entities + Repository port & adapter (rotation foundation)

> Back-reference : `.local-agents/works/account/03-login-session-management/2-plan/techplan.md` (Status: Approved) — sections 8 (business flow), 10 (Implementation Details), 5 (rotation-mechanics decisions)
> Depends on    : task-01 (tables exist)
> Model         : DeepSeek V4 Pro (goqu/invariant precision work)
> Rules touched : R12, R14, R16, R19 (repo half), R20
> ⚠️ TIER 0     : the rotation methods in `repository_db.go` (`RotateRefreshToken`, `RevokeRefreshTokenFamily`) are part of the Tier 0 fenced sub-area — see manifest's paired-pass checklist before commit

## Objective

Give `internal/domain/account` everything the login/session services (task-04) will call: the `LoginAttempt` entity, all new Repository port methods, and their goqu adapter implementations. No HTTP, no service logic here.

## Files

| File | Change |
|---|---|
| `backend/internal/domain/account/entity.go` | Edit |
| `backend/internal/domain/account/repository.go` | Edit |
| `backend/internal/domain/account/repository_db.go` | Edit |
| `backend/internal/domain/account/repository_db_integration_test.go` | Edit (`//go:build integration`) |

## Entity changes (`entity.go`)

```go
// LoginAttempt records one credential-verification outcome for the
// persistent lockout mechanism (Fitur 2C). Stage distinguishes the
// password step from the MFA step; lockout queries count failures per
// stage with different keys. Rows are append-only; user_id is set only
// when identity is reliably known.
type LoginAttempt struct {
    ID            uuid.UUID
    IdentifierHash string   // always populated
    UserID        *uuid.UUID // nil for password-stage rows where identity unknown? NO — see below
    Stage         string     // "password" | "mfa"
    Success       bool
    AttemptedAt   time.Time
}
```

`UserID` population rule per feature spec: password-stage rows carry the `user_id` **when the identity was found** (spec DDL: `user_id … ON DELETE SET NULL`; Assumption C keeps `identifier_hash` populated for MFA-stage rows too). A wrong-email attempt has no known user → `nil`. Doc-comment this precisely.

Also extend the `RefreshToken` doc comment: family/rotation fields are live as of this slice (no longer "first-generation only").

## Port methods (`repository.go`) — signatures indicative

```go
InsertLoginAttempt(ctx context.Context, tx pgx.Tx, a *LoginAttempt) error
CountRecentFailedAttemptsByIdentifier(ctx context.Context, identifierHash string, stage string, window time.Duration) (int, error)
CountRecentFailedAttemptsByUser(ctx context.Context, userID uuid.UUID, stage string, window time.Duration) (int, error)
FindRefreshTokenByHash(ctx context.Context, tokenHash string) (*RefreshToken, bool, error) // false = not found
RotateRefreshToken(ctx context.Context, tx pgx.Tx, oldTokenHash string, child *RefreshToken) (rotated bool, err error)
RevokeRefreshTokenByHash(ctx context.Context, tx pgx.Tx, tokenHash string) error
RevokeRefreshTokenFamily(ctx context.Context, tx pgx.Tx, familyID uuid.UUID) error
GetLoginUserView(ctx context.Context, userID uuid.UUID) (*LoginUserView, error)
```

`LoginUserView` carries exactly what `LoginResponse.user` needs: `ID, Name, Email (plaintext), EmailVerified bool, Roles []string, AuthProviders []string, MFAEnabled bool, CreatedAt`.

## Adapter implementation notes (`repository_db.go`)

- **All SQL via goqu `.Prepared(true)`** — never `fmt.Sprintf` into SQL (R20). Window arithmetic: compute the cutoff timestamp in Go via the injected clock and pass it as a bound arg; do not use `now() - interval` string-building.
- `CountRecentFailed*`: `SELECT count(*) … WHERE <key> AND stage=? AND success=false AND attempted_at > ?`.
- **`RotateRefreshToken` is the invariant core (INV-account-03)** — one transaction containing:
  1. Guarded parent mark:
     `UPDATE refresh_tokens SET replaced_by_id = :childID WHERE token_hash = :oldHash AND replaced_by_id IS NULL AND revoked_at IS NULL AND expires_at > :now RETURNING user_id, family_id`
     → 0 rows ⇒ `rotated=false`, ROLLBACK (caller treats as reuse/race-loser).
  2. Child INSERT (new UUID id, same `family_id` from RETURNING, hash of plain token, `expires_at = now+30d`).
  Single-tx is mandatory: a child-insert failure after a separate parent-mark would brick the family via reuse detection (techplan §7 risk row; Decision Log "Rotation mechanics").
- `RevokeRefreshTokenFamily`: `UPDATE refresh_tokens SET revoked_at = :now WHERE family_id = ? AND revoked_at IS NULL` — deliberately unguarded on `replaced_by_id`: already-rotated descendants must be revoked too (INV-account-04).
- `RevokeRefreshTokenByHash`: guarded `revoked_at IS NULL` so logout stays idempotent.
- `GetLoginUserView`: aggregate query joining `users` + `auth_identities` aggregation + `EXISTS` on `mfa_totp_secrets.enabled_at` + left join aggregate on `user_roles`. **Decrypt-on-read**: `crypto.Decrypt(primary_email_ct, r.keys)` — first decrypt path in this repo; lookups stay on `*_hash` columns, never against ciphertext (encryption-at-rest discipline). Multiple providers/roles → gather with `pgx.CollectRows` or manual loop consistent with existing style.

## Tests (`//go:build integration`, real Postgres)

Happy-path coverage for each new method:
- insert attempt → count by identifier / by user respects stage + window boundary (a failure 15 min ago counts at exactly-window edge per injected-clock semantics; document chosen boundary: strictly `>` cutoff).
- rotate happy path: parent marked, child inserted, same family.
- rotate on already-rotated row ⇒ `rotated=false`, no second child.
- revoke-family revokes parent + rotated descendants.
- user view assembles all fields for a fixture user (incl. decrypt correctness).

Heavy concurrency proof (≥100 goroutines, `-race`) lives in **task-06**, not here.

## Out of scope

Service orchestration, sentinel errors, verifier seam (task-04); anything under `transport/` or `platform/auth` (tasks 03/05).
