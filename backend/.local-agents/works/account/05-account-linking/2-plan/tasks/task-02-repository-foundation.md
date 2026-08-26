# Task 02: Repository layer foundation (four new operations)

> Back-reference : `.local-agents/works/account/05-account-linking/2-plan/techplan.md` (Status: Approved) — sections 8 (Interface Contract), 9 (step 2), 10 (repository.go / repository_db.go), 5 (D1)
> Depends on    : task-01 (integration truth-tests later assume migration 000010 applied; compile does not)
> Model         : DeepSeek V4 Pro (rule-table-heavy precision work: goqu parameterization, nullable scans, golden-rule SQL discipline)

## Objective

Add the four persistence operations both service flows (tasks 03–04) build on. No business logic here — this is the port + adapter layer only, following the established tx-taking interface and `(nil, nil)`-on-not-found conventions.

## Files

| File | Change |
|---|---|
| `backend/internal/domain/account/repository.go` | +4 interface methods with doc comments |
| `backend/internal/domain/account/repository_db.go` | +4 goqu implementations |

## Interface additions (signatures per techplan §10)

```go
// FindAuthIdentitiesByUser returns ALL identities for userID with
// non-encrypted fields populated; Identifier left empty per the
// read-path convention (lookup-by-hash codebase rule).
FindAuthIdentitiesByUser(ctx context.Context, userID uuid.UUID) ([]AuthIdentity, error)

// DeleteAuthIdentitiesByIDs hard-deletes caller-classified rows.
// Caller (service) owns the guard classification; this method must NOT
// re-check conditions.
DeleteAuthIdentitiesByIDs(ctx context.Context, tx pgx.Tx, ids []uuid.UUID) error

// UpdateCredentialSecret updates credential_secret for the single row
// matching (userID, providerType). Single-row UPDATE. No plaintext ever
// reaches this layer — callers pass bcrypt hashes only.
UpdateCredentialSecret(ctx context.Context, tx pgx.Tx, userID uuid.UUID, providerType, hashedSecret string) error

// RevokeAllRefreshTokensForUser sets revoked_at=now() on every
// non-revoked refresh_tokens row for userID. INV-account-05 primitive —
// FIRST implementation anywhere; Fitur 04 will reuse, not re-derive
// (techplan Resolved #3/#5). Deliberately NO replaced_by_id guard:
// already-rotated rows die too, same reasoning as family revocation.
RevokeAllRefreshTokensForUser(ctx context.Context, tx pgx.Tx, userID uuid.UUID) error
```

Plus the tx-taking locked finder used by unlink's classification (final naming at build per techplan §10): either `FindAuthIdentitiesByUserTx(ctx, tx pgx.Tx, userID)` issuing `SELECT … FOR UPDATE`, or an options variant of `FindAuthIdentitiesByUser`. **The lock must be acquired inside the caller's transaction — never on the pool.**

## Implementation constraints

- goqu only, `Prepared(true)`, parameterized — never string-built (`fmt.Sprintf` of an id list breaks the golden rule; use `goqu.C("id").In(ids)`)
- Nullable columns scanned via established `sql.NullString` / `sql.NullTime` patterns (see existing `FindAuthIdentityByIdentifierHash`)
- `RevokeAllRefreshTokensForUser` predicate exactly: `WHERE user_id = $1 AND revoked_at IS NULL` — index-supported by `ix_refresh_tokens_user_id`
- Compile-time assertion `var _ Repository = (*RepositoryDB)(nil)` stays satisfied

## Rules enabled (not yet proven here)

R7 (UpdateCredentialSecret + RevokeAllRefreshTokensForUser in one tx), R9 (delete + audit), R13 (FOR UPDATE serialization). Their tests live in tasks 03–06.

## Verification

- `go build ./...` and `go vet ./...` clean
- Existing unit suites untouched and green (`go test ./internal/domain/account/...`) — the fake repository in tests must be extended with the new methods so existing fakes still satisfy the interface
