# Task 02: Repository layer foundation (six new MFA operations)

> Back-reference : `.local-agents/works/account/06-mfa-totp/2-plan/techplan.md` (Status: Approved by Anhar) — sections 8 (SQL contract block — authoritative), 9 (step 3), 10 (repository.go / repository_db.go / entity.go), 5 (D3, D4, D5)
> Depends on    : task-01 (nothing in this task imports pquerna/otp, but run order keeps the chain linear; no compile dependency)
> Model         : DeepSeek V4 Pro (rule-table-heavy precision: goqu parameterization golden rule, conflict-arm predicate encoding, nullable scans)

## Objective

Add the six persistence operations all three service methods (task 03) and the verifier (task 04) build on. Port + adapter layer only — zero business logic, zero crypto decisions here beyond encrypt/decrypt at the storage boundary per the established adapter doctrine.

## Files

| File | Change |
|---|---|
| `backend/internal/domain/account/entity.go` | +`MFABackupCode` struct (ID uuid.UUID, UserID uuid.UUID, CodeHash string, CreatedAt time.Time) with doc comment noting ciphertext ownership at the adapter |
| `backend/internal/domain/account/repository.go` | +6 interface methods with full doc comments |
| `backend/internal/domain/account/repository_db.go` | +6 goqu implementations |

## Interface additions (signatures per techplan §10, contracts verbatim)

```go
// D5 guard lives IN the statement — see techplan §8 SQL block.
// inserted=false means "enabled row blocked the upsert" ⇒ service maps 409.
// It is NOT an error signal; distinguish from real exec errors.
UpsertPendingMFASecret(ctx context.Context, userID uuid.UUID, secretCiphertext []byte) (inserted bool, err error)

// Adapter decrypts internally (D3): domain never touches ciphertext.
// Returns found=false when no row exists for userID.
GetMFATOTPSecretForVerify(ctx context.Context, userID uuid.UUID) (secretBase32 string, enabledAt *time.Time, found bool, err error)

// Guarded enable — first statement of the confirm tx (D4-A).
// ok=false ⇔ matched 0 rows ⇔ already enabled or row absent (race loser).
EnableMFATOTPIfPending(ctx context.Context, tx pgx.Tx, userID uuid.UUID) (ok bool, err error)

// Guarded disable — idempotent repeat-safe.
// disabled=false ⇔ 0 rows ⇔ already disabled (idempotent no-op path).
SetMFADisabledIfEnabled(ctx context.Context, tx pgx.Tx, userID uuid.UUID) (disabled bool, err error)

// Batch insert of exactly-backupCodeCount rows inside the confirm tx.
InsertMFABackupCodes(ctx context.Context, tx pgx.Tx, codes []MFABackupCode) error

// INV-account-06 in ONE statement — BOTH clauses mandatory:
//   used_at IS NULL  AND  owner's mfa_totp_secrets.enabled_at IS NOT NULL
// via JOIN, per techplan §8's redemption UPDATE verbatim. The enabled-
// check is a DB clause here, NEVER app-side sequencing (§13 mistake row 2).
RedeemMFABackupCode(ctx context.Context, tx pgx.Tx, userID uuid.UUID, codeHash string) (redeemed bool, err error)
```

## Implementation constraints

- The four guarded shapes in techplan §8 are **contract, not suggestion**:
  - Upsert = `INSERT … ON CONFLICT (user_id) DO UPDATE SET … WHERE mfa_totp_secrets.enabled_at IS NULL` (0 affected rows is the designed 409 signal — §13 mistake row 8)
  - Enable = `UPDATE … SET enabled_at=now() WHERE user_id=$1 AND enabled_at IS NULL`
  - Disable = `UPDATE … SET enabled_at=NULL WHERE user_id=$1 AND enabled_at IS NOT NULL`
  - Redemption = joined UPDATE carrying both INV-account-06 clauses
- Encode through goqu upsert/Returning capabilities; fallback honestly flagged at build if the builder cannot express the conflict-arm predicate: use `tx.Exec` with a goqu-generated **parameterized** literal (`Prepared(true)`) — string concatenation remains forbidden regardless of form
- Nullable scans (`enabled_at`, `secret_encrypted` handled as []byte) follow established `sql.Null*` patterns; decrypt failure returns wrapped error, never plaintext leakage paths
- Compile-time assertion `var _ Repository = (*RepositoryDB)(nil)` stays satisfied
- Extend the fake repository in existing unit tests with the six methods so current suites keep compiling (green before this task counts)

## Rules enabled (not yet proven here)

R1–R4 (upsert path), R5/R8 (guarded enable + batch insert), R9 (redemption), R11 (disable). Their tests live in tasks 03–06.

## Verification

- `go build ./...`, `go vet ./...` clean
- Existing unit suites green (`go test ./internal/domain/account/...`)
- Manual SQL review against §8 block line-by-line (predicate direction, arm placement)
