# Task 03 — Domain Account: Entity + Repository (Data Layer)

> Ticket    : 01-register-email-verification
> Sub-task  : 3 of 5
> Axis      : Dependency/sequence chain (primary) + layer (vertical slice)
> Status    : Blocked on Tier 0 crypto prerequisite (see manifest)
> Back-ref  : `../2-plan/techplan.md` (originating contract techplan — cross-check high-level decisions there whenever needed)

---

## 1. Scope

The data layer of the new `internal/domain/account/` package: entity
structs, the `Repository` interface (ports), and the goqu+pgx
implementation (adapter) that persists `User`, `AuthIdentity`, and
`AuthToken` rows — including PII encryption of `primary_email` and
`identifier` via the human-authored `platform/crypto/` functions.

No business logic. No service. No HTTP. This is the seam between the
storage schema (Task 01) and the service (Task 04).

**In scope:**
- `entity.go` — `User`, `AuthIdentity`, `AuthToken` structs
- `repository.go` — `Repository` interface (ports)
- `repository_db.go` — goqu + pgx adapter, uses `crypto.Encrypt`/`HMAC`
  for PII columns, full INV-account-08 3-clause predicate in
  `RedeemToken`
- Integration test (`//go:build integration`) that exercises the
  repository against a real Postgres via testcontainers

**Out of scope:**
- `service.go` — Task 04
- Token *generation* (`crypto/rand` → SHA-256) — Task 04
- HTTP handlers — Task 05

## 2. Dependencies

- **Hard deps:**
  - **Tier 0 crypto prerequisite** (HUMAN, fenced — see manifest):
    `platform/crypto/` must export `Encrypt(plaintext []byte, key) ([]byte, error)`,
    `Decrypt(ciphertext []byte, key) ([]byte, error)`,
    `HMAC(data []byte, key) string` before this task can compile.
  - Task 01 (schema must exist for integration tests; `goqu` must be in `go.mod`).
- **Soft deps:** none
- **Blocks:** Task 04 (service depends on `Repository` + entities)

## 3. Files

| File | Change Type |
|---|---|
| `backend/internal/domain/account/entity.go` | New |
| `backend/internal/domain/account/repository.go` | New |
| `backend/internal/domain/account/repository_db.go` | New |
| `backend/internal/domain/account/repository_db_integration_test.go` | New (build tag `integration`) |

## 4. Implementation detail

### `backend/internal/domain/account/entity.go` (new)

Domain entities — plain structs, no methods, no ORM tags (goqu is used
explicitly in the adapter). Field names match the schema columns where
practical, but the ciphertext/hash split is explicit in the type
(`[]byte` for ciphertext, `string` for hex hash).

```go
// Package account implements the account domain: user identity,
// authentication identities (email/password, google, phone OTP), and
// single-use verification tokens.
package account

import (
    "time"

    "github.com/google/uuid"
)

// User is the top-level account entity. PrimaryEmailCiphertext holds
// the AES-GCM-encrypted email; PrimaryEmailHash is the HMAC-SHA256
// hex digest used for uniqueness and lookup.
type User struct {
    ID                    uuid.UUID
    Name                  string
    PrimaryEmailCiphertext []byte
    PrimaryEmailHash      string
    CreatedAt             time.Time
    UpdatedAt             time.Time
}

// AuthIdentity is a provider-scoped credential binding for a User.
// provider_type ∈ {"email_password", "google", "phone_otp"}.
// IdentifierCiphertext / IdentifierHash follow the same PII pattern
// as User.PrimaryEmail*.
type AuthIdentity struct {
    ID                    uuid.UUID
    UserID                uuid.UUID
    ProviderType          string
    IdentifierCiphertext  []byte
    IdentifierHash        string
    CredentialSecret      *string // bcrypt hash for email_password; nil for google
    VerifiedAt            *time.Time
    CreatedAt             time.Time
    UpdatedAt             time.Time
}

// AuthToken is a single-use, time-bound verification or reset token.
// Redemption is guarded by the 3-clause predicate
// (used_at IS NULL AND revoked_at IS NULL AND expires_at > now())
// per INV-account-08 — see repository.RedeemToken.
type AuthToken struct {
    ID         uuid.UUID
    UserID     uuid.UUID
    Purpose    string // "email_verification" | "password_reset"
    TokenHash  string // SHA-256 hex of the plain token
    ExpiresAt  time.Time
    UsedAt     *time.Time
    RevokedAt  *time.Time
    CreatedAt  time.Time
}
```

### `backend/internal/domain/account/repository.go` (new)

The `Repository` port. All methods take `context.Context` first (AGENTS.md
convention). Insert methods take a `pgx.Tx` so the service can wrap
multi-row writes in a single transaction (R16: concurrent duplicate
registration must roll back cleanly — techplan §13 last row).

```go
package account

import (
    "context"
    "time"

    "github.com/google/uuid"
    "github.com/jackc/pgx/v5"
)

// Repository is the persistence port for the account domain.
// Implementations must parameterize all SQL via goqu — never string
// concatenation (AGENTS.md golden rule). PII columns (primary_email,
// identifier) must be encrypted at insert time; *_hash columns hold
// the HMAC for lookup.
type Repository interface {
    // InsertUser inserts a new User. primary_email is encrypted;
    // primary_email_hash is the HMAC. Called within the caller's tx.
    InsertUser(ctx context.Context, tx pgx.Tx, user *User) error

    // InsertAuthIdentity inserts a new AuthIdentity. identifier is
    // encrypted; identifier_hash is the HMAC. Called within tx.
    InsertAuthIdentity(ctx context.Context, tx pgx.Tx, identity *AuthIdentity) error

    // InsertAuthToken inserts a new single-use token. Called within tx.
    InsertAuthToken(ctx context.Context, tx pgx.Tx, token *AuthToken) error

    // FindAuthIdentityByIdentifierHash looks up an identity by
    // (providerType, identifierHash). Returns (nil, nil) if not found.
    FindAuthIdentityByIdentifierHash(ctx context.Context, providerType, identifierHash string) (*AuthIdentity, error)

    // FindAuthTokenByHash looks up a token by its SHA-256 hash.
    // Returns (nil, nil) if not found.
    FindAuthTokenByHash(ctx context.Context, tokenHash string) (*AuthToken, error)

    // RedeemToken atomically marks a token used iff it is currently
    // valid: used_at IS NULL AND revoked_at IS NULL AND expires_at > now()
    // (full 3-clause predicate per INV-account-08 Statement).
    // Returns true if exactly 1 row was affected (success); false if 0
    // rows (not-found / already-used / revoked / expired). On false,
    // the caller disambiguates expired vs other via FindAuthTokenByHash.
    RedeemToken(ctx context.Context, tokenHash string) (bool, error)

    // SetVerifiedAt sets auth_identities.verified_at = verifiedAt for
    // the given identity. Called by VerifyEmail after a successful
    // RedeemToken.
    SetVerifiedAt(ctx context.Context, identityID uuid.UUID, verifiedAt time.Time) error

    // RevokeTokens sets revoked_at = now() for all unused, unrevoked
    // tokens of (userID, purpose). Called by resend before issuing a
    // new token (R13).
    RevokeTokens(ctx context.Context, userID uuid.UUID, purpose string) error
}
```

**Critical correctness note — INV-account-08 guard.** The invariant's
**Statement** says the guard is the 3-clause
`used_at IS NULL AND revoked_at IS NULL AND expires_at > now()`. Its
**Verification** field omits `revoked_at IS NULL` (2 clauses) — that is
a documented spec error (techplan §14 Open Item #2) and **must not** be
followed. Use the 3-clause version. Revoked tokens (superseded by a
later resend) must be rejected. Do not edit the spec — the agent is
forbidden from editing `docs/spec/*` (AGENTS.md §4); the inconsistency
is flagged for a human to fix.

### `backend/internal/domain/account/repository_db.go` (new)

goqu + pgx adapter. Uses `*pgxpool.Pool` for read/standalone queries
and the caller-supplied `pgx.Tx` for inserts (so the service controls
transaction boundaries). PII encryption delegated to `platform/crypto`
(Tier 0 prerequisite).

```go
package account

import (
    "context"
    "fmt"
    "time"

    "github.com/doug-martin/goqu/v9"
    "github.com/google/uuid"
    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgxpool"

    "kencleng/internal/platform/crypto"
)

type RepositoryDB struct {
    db         *pgxpool.Pool
    cryptoKeys *crypto.Keys
}

func NewRepositoryDB(db *pgxpool.Pool, keys *crypto.Keys) *RepositoryDB {
    return &RepositoryDB{db: db, cryptoKeys: keys}
}

func (r *RepositoryDB) InsertUser(ctx context.Context, tx pgx.Tx, user *User) error {
    emailCt, err := crypto.Encrypt([]byte(user.PrimaryEmailHash /* plaintext email is stored separately — see note */), r.cryptoKeys.EncryptionKey)
    // NOTE: the User entity should carry the PLAINTEXT email on the
    // write path (to be encrypted here), not the hash. Coordinate the
    // entity shape with Task 04 — the service computes the HMAC and
    // passes both plaintext (for encryption) and hash (for the column)
    // through to this method. Adjust the signature/entity as needed
    // so that the plaintext is available at insert time and never
    // stored on the struct after insert.
    _ = emailCt // placeholder — see note above
    // ... goqu Insert into users with parameterized values ...
    return nil
}

// InsertAuthIdentity, InsertAuthToken, Find*, RedeemToken,
// SetVerifiedAt, RevokeTokens — all use goqu parameterized SQL.
// RedeemToken specifically:
//   UPDATE auth_tokens SET used_at = now()
//   WHERE token_hash = $1
//     AND used_at IS NULL
//     AND revoked_at IS NULL
//     AND expires_at > now()
// (full 3-clause predicate — INV-account-08 Statement).
// Check pgconn.PgError / rows-affected to decide (true, false, error).
```

**Entity write-path note (resolve with Task 04 author):** the entity
as written stores `PrimaryEmailCiphertext` / `PrimaryEmailHash`. On the
insert path the service has the **plaintext** email (it received it
from the handler) and needs both the ciphertext (for the BYTEA column)
and the HMAC hash (for the lookup column). Either (a) the entity
carries the plaintext on the write path and the adapter encrypts, or
(b) the service encrypts and the adapter receives ciphertext already.
Pick one and keep it consistent with `AuthIdentity` (same pattern for
`identifier`). The AGENTS.md golden rule says "PII fields follow the
established encryption pattern" — the encryption call itself lives in
`platform/crypto/` (fenced); **where** it is *invoked* (service vs
adapter) is a design choice for this task + Task 04 to settle. Prefer
the adapter doing it so the service never handles raw ciphertext and
the encryption pattern is enforced at the storage boundary.

**Error handling.** Use `errors.As` with `*pgconn.PgError` and check
`Code == "23505"` for unique-violation (R16 path) — never
`strings.Contains(err.Error(), "unique")` (techplan §13 row 8). Wrap
with `fmt.Errorf("...: %w", err)`. Never leak SQL errors upward in a
way that reaches the HTTP response — the handler (Task 05) maps
sentinel errors to Problem Details.

### `backend/internal/domain/account/repository_db_integration_test.go` (new)

Build tag `//go:build integration`. Uses `testcontainers-go` to spin up
a real Postgres, applies migrations (Task 01), and exercises:

- InsertUser + InsertAuthIdentity + InsertAuthToken in a single tx;
  commit succeeds; rows are present.
- Concurrent duplicate `InsertAuthIdentity` for the same
  `(provider_type, identifier_hash)` — exactly one succeeds, the other
  returns unique-violation (R16 at the storage layer; the
  service-level ≥100-goroutine test is Task 04).
- `RedeemToken` with the full 3-clause guard:
  - valid token → true, `used_at` set.
  - already-used token → false.
  - revoked token (superseded) → false (proves the 3rd clause is
    enforced; this is the spec-error regression test — techplan §7
    risk row "INV-account-08 guard missing revoked_at IS NULL clause").
  - expired token → false.
  - non-existent token → false.
- `RevokeTokens` revokes only unused, unrevoked tokens of the given
  purpose; already-used tokens keep `used_at`, already-revoked keep
  their earlier `revoked_at`.

## 5. Rules covered (this task's slice)

This task does not satisfy end-user-facing rules alone; it provides the
storage primitives that Task 04's service uses to satisfy:
- **R1** (InsertUser + InsertAuthIdentity + InsertAuthToken)
- **R10/R11/R12** (RedeemToken 3-clause guard — the storage-level
  correctness for R11/R12 lives here; the service-level orchestration
  and concurrent test are Task 04)
- **R13** (RevokeTokens)
- **R16** (unique index `ux_auth_identities_provider_identifier` —
  created in Task 01, enforced and tested here at the storage layer)

## 6. Testing checklist (this task's slice)

- [ ] Integration: insert User+AuthIdentity+AuthToken in one tx → commit,
      rows present with correct ciphertext/hash columns populated.
- [ ] Integration: concurrent duplicate `InsertAuthIdentity` → exactly
      one succeeds, other returns a wrapped unique-violation
      (detectable via `errors.As` on `*pgconn.PgError`, code `23505`).
- [ ] Integration: `RedeemToken` valid token → true, `used_at` set.
- [ ] Integration: `RedeemToken` already-used → false.
- [ ] Integration: `RedeemToken` revoked (superseded) → false
      (**regression for INV-account-08 spec error**).
- [ ] Integration: `RedeemToken` expired → false.
- [ ] Integration: `RedeemToken` non-existent → false.
- [ ] Integration: `RevokeTokens` revokes only unused+unrevoked of the
      given purpose; used/revoke state of other rows unchanged.
- [ ] `go test -race -tags=integration ./internal/domain/account/...` clean.

## 7. Common mistakes to avoid (techplan §13 slice)

| Mistake | Fix |
|---|---|
| Using 2-clause token guard (missing `revoked_at IS NULL`) | Use the full 3-clause predicate from INV-account-08 **Statement**. The Verification field's 2-clause version is a documented spec error. |
| `users` insert + `auth_identities` insert in separate transactions | Inserts share the caller-supplied `pgx.Tx`; the service (Task 04) begins the tx. On unique-violation the whole tx rolls back — no orphaned `users` row. |
| String-matching DB errors | `errors.As(err, &pgErr)` + check `pgErr.Code == "23505"`. |
| `fmt.Sprintf` for SQL | goqu only — parameterized. |
| Not wrapping errors with `%w` | `fmt.Errorf("...: %w", err)`. |

## 8. Risk note

- Assumptions made: `platform/crypto/` exposes `Encrypt`, `Decrypt`,
  `HMAC` with the signatures sketched in the manifest prerequisite;
  the exact entity write-path (who encrypts) is resolved in
  coordination with Task 04 — preferred: adapter encrypts so the
  service never handles ciphertext.
- Edge cases intentionally NOT handled: HMAC key rotation (out of
  scope for v1 per tech stack doc); plaintext-email retention on the
  entity after insert (the entity should not retain plaintext after
  the insert call — Task 04 must not stash it).
- Concurrency assumptions: `RepositoryDB` is safe for concurrent use
  (`pgxpool.Pool` is goroutine-safe; goqu builders are not shared
  across goroutines). `RedeemToken`'s atomic `UPDATE ... WHERE`
  is the single-source-of-truth for single-use — no additional
  application-level locking.
- What is not tested, and why: the service-level 100-goroutine race
  (R16) and the concurrent double-submit (R12) are Task 04 — here we
  prove the storage primitives behave correctly under direct
  concurrent access; the end-to-end orchestration test belongs with
  the service.
