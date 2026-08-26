# Task 01 — Repository foundation: credential update + user-scoped session revoke

> Back-reference (contract): `../techplan.md` — sections 1–8 are the source of truth for all decisions here. If anything in this file seems to contradict the techplan, the techplan wins; stop and flag instead of resolving silently.
> Splitting axis: dependency/sequence chain (see `manifest.md`). This is chain step 1 — no dependencies.

## Scope

**In scope:**
- Two new methods on the `Repository` interface (`internal/domain/account/repository.go`) with doc comments in house style
- Their goqu implementations in `internal/domain/account/repository_db.go`
- Repository-level tests proving guard semantics against real Postgres (integration tag)

**Out of scope (this task):**
- Any service/handler/notification/openapi change (tasks 02–05)
- Migrations — schema already supports everything (see Interface Contract)

## Dependencies

None. First task in the chain. Compiles green on its own.

## Requirements (from techplan §3 — rows relevant to this task)

| Condition | Requirement | Source/Note |
|---|---|---|
| Reset with valid token + passing password | ONE transaction: credential updated, `used_at` set, every refresh token for user revoked | INV-account-05 |
| Concurrent double-submit | Exactly one succeeds (guarded UPDATE), other 404 | INV-account-08 |

The transaction composition itself happens in Task 03 (service layer). This task delivers the two write primitives the transaction will call, plus proof of their guard/idempotency semantics.

## Interface additions (exact signatures — techplan §8)

```go
// repository.go (Repository interface additions; both take caller's tx)
UpdateIdentityCredentialSecret(ctx context.Context, tx pgx.Tx,
    userID uuid.UUID, providerType string, passwordHash string) error
RevokeAllRefreshTokensForUser(ctx context.Context, tx pgx.Tx, userID uuid.UUID) error
```

## Implementation details (redistributed from techplan §10, verbatim intent)

**File**: `internal/domain/account/repository.go`
- Add the two interface methods with doc comments in house style (every exported function gets one — backend/AGENTS.md §2). The `RevokeAllRefreshTokensForUser` doc comment MUST note: idempotency via the `revoked_at IS NULL` guard, and its purpose as the INV-account-05 mass-revoke primitive (called inside the reset transaction).

**File**: `internal/domain/account/repository_db.go`
- Implementations mirroring existing shapes exactly:
  - `UpdateIdentityCredentialSecret`: goqu `Update("auth_identities").Set(goqu.Record{"credential_secret": passwordHash}).Where(goqu.Ex{"user_id": userID, "provider_type": providerType})`, `Prepared(true)`, exec on the caller's `tx`. Shape precedent: `SetUserVerified` (repository_db.go:333–346).
  - `RevokeAllRefreshTokensForUser`: goqu `Update("refresh_tokens").Set(goqu.Record{"revoked_at": time.Now()}).Where(goqu.Ex{"user_id": userID}, goqu.L("revoked_at IS NULL"))`, `Prepared(true)`, exec on the caller's `tx`. Shape precedent: `RevokeRefreshTokenByHash` / `RevokeTokens` (repository_db.go:354–371, 560–579).
- Golden rules that bind this file (backend/AGENTS.md): parameterized SQL via goqu ONLY — never string concatenation; wrap errors with `fmt.Errorf("account: ...: %w", err)` preserving the chain.

## Binding decisions (techplan §5 — rows governing this task)

| Decision | Resolution |
|---|---|
| D1a keying | `(user_id, provider_type)` — NOT identity ID (would widen `RedeemToken`'s RETURNING contract; mirrors `SetUserVerified`; INV-account-01 guarantees ≤1 identity per user+provider) |
| D1b revoke strategy | Single guarded UPDATE by `user_id` inside caller's tx. Rejected: per-family loop (SELECT-first inside critical tx; family created mid-window escapes revocation — genuine INV-05 hole under concurrency). Rejected: outside-tx revoke (violates INV-05 verbatim) |

## Relied-upon schema (NO migration — techplan §8)

```sql
-- migration 000003 (already shipped): auth_tokens purpose CHECK includes 'password_reset'
-- migration 000004 (already shipped):
CREATE INDEX ix_refresh_tokens_user_id ON refresh_tokens (user_id);
CREATE INDEX ix_refresh_tokens_active ON refresh_tokens (user_id)
    WHERE revoked_at IS NULL AND replaced_by_id IS NULL;
```

## Testing checklist (this task's items from techplan §12)

Integration suite (`//go:build integration`, real Postgres via `DATABASE_URL`, skip when unset, `-race`; extend `repository_db_integration_test.go` or add a focused new file):

- [ ] `TestUpdateIdentityCredentialSecret_UpdatesOnlyTargetIdentity` — seeds two identities (email_password + google) for one user; asserts only the targeted provider row's `credential_secret` changes
- [ ] `TestRevokeAllRefreshTokensForUser_RevokesEveryUnrevokedRow` — seeds multiple families incl. already-rotated rows (`replaced_by_id` set) and already-revoked rows; asserts ALL un-revoked rows for the user end revoked, other users untouched, repeat call idempotent (no error, no double-write effect)
- [ ] Guard-semantics assertions ride the existing `-race` integration run command: `go test -tags=integration -race ./internal/domain/account/...`

Unit-level: none required beyond compilation — these methods have no branching logic; their correctness IS the SQL guard, provable only against real Postgres (techplan Area-7 finding: fake repos cannot prove atomicity/guard behavior).

## Common mistakes that apply here (techplan §13)

| Mistake | Fix |
|---|---|
| Revoking only active tokens (`replaced_by_id IS NULL` filter added) | INV-05 says EVERY row for the user — rotated-out rows die too (precedent: `RevokeRefreshTokenFamily` deliberately includes rotated rows) |
| String-building the WHERE clause | goqu only, golden rule |

## Gate

`go build ./...` green; `go test ./...` green; `go test -tags=integration -race ./internal/domain/account/...` green (new tests pass, existing untouched).
