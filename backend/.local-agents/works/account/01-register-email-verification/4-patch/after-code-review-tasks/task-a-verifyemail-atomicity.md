# Task A — `VerifyEmail` Atomicity + Silent-Failure Fix

> Ticket    : 01-register-email-verification
> Sub-task  : A of F (post-review remediation)
> Findings  : S1 (silent success), S2 (non-atomic redeem+verify)
> Blocking  : yes
> Back-ref  : `../report.md` §1 (S1, S2); `../../2-plan/techplan.md` §4 R8-R12

---

## 1. Scope

Make token redemption and identity verification atomic, and stop
swallowing the post-redeem re-fetch error. Today `RedeemToken` commits
`used_at` in one auto-transaction and `SetUserVerified` updates
`verified_at` in a *separate* auto-transaction; between them a
re-fetch (`userIDForToken`) can fail and is silently turned into
`uuid.Nil` → `SetUserVerified` affects 0 rows → `VerifyEmail` returns
`nil` → handler writes `200 "Email verified."` while the identity is
not verified and the token is burned (S1/S2).

**In scope:**
- `Repository.RedeemToken` returns the redeemed token's `userID` (and
  `purpose`) instead of a bare `bool`, eliminating the re-fetch.
- Redeem + set-verified run in a single `pgx.Tx` so a set-verified
  failure rolls back the redeem (no burned token without verification).
- `userIDForToken` removed (no re-fetch path to silently fail).
- Service + unit + integration tests updated to the new signature and
  to assert the error path.

**Out of scope:**
- The 3-clause guard predicate itself (already correct — INV-account-08
  Statement version; do not touch).
- Login / JWT issuance (task #3, Tier 0 fenced).
- The `Register` branches (Task B owns the DB-timing work there).

## 2. Dependencies

- **Hard deps:** none.
- **Soft deps:** none.
- **Blocks:** Task F (its test touches assert on `RedeemToken`'s shape).
- **Coordinate with:** Task B (same file, `service.go`, disjoint
  methods — `VerifyEmail` vs `Register`).

## 3. Files

| File | Change Type | Why |
|---|---|---|
| `internal/domain/account/repository.go` | Edit | `RedeemToken` signature: `(bool, error)` → `(*RedeemResult, error)` or `(userID uuid.UUID, purpose string, ok bool, err error)` |
| `internal/domain/account/repository_db.go` | Edit | `RedeemToken` returns `UserID` via `RETURNING` (`goqu` `Returning("user_id")`) |
| `internal/domain/account/service.go` | Edit | `VerifyEmail`: wrap redeem+set-verified in one tx; delete `userIDForToken` |
| `internal/domain/account/service_test.go` | Edit | update fake `RedeemToken` signature; add post-redeem-failure test |
| `internal/domain/account/repository_db_integration_test.go` | Edit | assert redeem+verify atomicity (set-verified failure rolls back redeem) |

## 4. Implementation detail

### `repository.go` — `RedeemToken` contract

Change the port so the caller gets `user_id` without a second query:

```go
// RedeemToken atomically marks a token used iff it is currently valid:
// used_at IS NULL AND revoked_at IS NULL AND expires_at > now() (full
// 3-clause predicate per INV-account-08 Statement). On success it
// returns the token's user_id and purpose. Returns ok=false if 0 rows
// were affected (not-found / already-used / revoked / expired); the
// caller disambiguates expired vs other via FindAuthTokenByHash.
RedeemToken(ctx context.Context, tokenHash string) (userID uuid.UUID, purpose string, ok bool, err error)
```

(Alternative shape acceptable: a small `*RedeemResult` struct — pick
whichever reads cleaner; the multi-return is fine for 2 values.)

### `repository_db.go` — `RETURNING`

Use goqu's `Returning` to fetch `user_id`/`purpose` from the same
atomic UPDATE so no second round-trip is needed:

```go
var userID uuid.UUID
var purpose string
sqlStr, args, err := pgDialect.Update("auth_tokens").
    Set(goqu.Record{"used_at": time.Now()}).
    Where(
        goqu.Ex{"token_hash": tokenHash},
        goqu.L("used_at IS NULL"),
        goqu.L("revoked_at IS NULL"),
        goqu.L("expires_at > now()"),
    ).
    Returning("user_id", "purpose").
    Prepared(true).
    ToSQL()
if err != nil { return uuid.Nil, "", false, fmt.Errorf("account: build redeem: %w", err) }
err = tx.QueryRow(ctx, sqlStr, args...).Scan(&userID, &purpose)
if errors.Is(err, pgx.ErrNoRows) {
    return uuid.Nil, "", false, nil // 0 rows affected — not redeemed
}
if err != nil {
    return uuid.Nil, "", false, fmt.Errorf("account: redeem: %w", err)
}
return userID, purpose, true, nil
```

Note: `RedeemToken` now needs the caller's `pgx.Tx` so the subsequent
`SetUserVerified` is in the same transaction. Add `tx pgx.Tx` to the
signature (matching the other insert methods).

### `service.go` — `VerifyEmail`

Delete `userIDForToken`. Wrap redeem + set-verified in one tx:

```go
func (s *Service) VerifyEmail(ctx context.Context, token string) error {
    tokenHash := sha256Hex(token)

    tx, err := s.tx.BeginTx(ctx)
    if err != nil {
        return fmt.Errorf("account: begin tx: %w", err)
    }
    committed := false
    defer func() {
        if !committed {
            _ = tx.Rollback(ctx)
        }
    }()

    userID, _, ok, err := s.repo.RedeemToken(ctx, tx, tokenHash)
    if err != nil {
        return fmt.Errorf("account: redeem token: %w", err)
    }
    if !ok {
        // disambiguate expired (R9) vs not-found/used/revoked (R10/R11)
        t, err := s.repo.FindAuthTokenByHash(ctx, tokenHash) // read, ok outside tx
        if err != nil {
            return fmt.Errorf("account: find token: %w", err)
        }
        if t != nil && !t.ExpiresAt.After(time.Now()) {
            return ErrTokenExpired
        }
        return ErrTokenNotFound
    }

    if err := s.repo.SetUserVerified(ctx, tx, userID, providerEmailPassword, time.Now()); err != nil {
        return fmt.Errorf("account: set verified: %w", err) // rolls back the redeem via the deferred Rollback
    }
    if err := tx.Commit(ctx); err != nil {
        return fmt.Errorf("account: commit verify: %w", err)
    }
    committed = true

    log.Printf("account: email verified user_id=%s (token redacted)", userID)
    return nil
}
```

Key points:
- `RedeemToken` now takes `tx` and runs inside the caller's tx.
- If `SetUserVerified` fails, the deferred `Rollback` undoes the
  redeem — the token is *not* burned, the user can retry. (S2 fixed.)
- `FindAuthTokenByHash` for the disambiguation path stays a read; it
  runs *after* the tx is rolled back (ok==false path, nothing to undo).
  Move it outside the tx or run it before beginning the tx — either is
  fine; the disambiguation read is non-mutating.
- No `uuid.Nil` return path remains — every failure returns a real
  error → handler maps to 500 (not a fake 200). (S1 fixed.)

## 5. Tests to add / update

### Unit (`service_test.go`)

- Update `fakeRepo.RedeemToken` to the new signature (`(userID, purpose, ok, err)`).
- `TestVerifyEmail_ValidToken_SetsVerifiedAt` — still asserts 1
  `SetUserVerified` call with the right userID; now sourced from
  `RedeemToken`'s return, not a re-fetch.
- **NEW** `TestVerifyEmail_SetVerifiedFails_RollsBackRedeem` — inject
  `setVerifiedErr`; assert the redeem is rolled back (token *not*
  consumed), `VerifyEmail` returns a wrapped error (→ 500 at handler),
  not `nil`.
- **NEW** `TestVerifyEmail_RedeemReturnsUserID_NoRefetch` — assert no
  `FindAuthTokenByHash` call on the success path (proves the re-fetch
  is gone).
- Keep R9/R10/R11/R12 tests; update to the new signature. R12
  (concurrent double-submit) still asserts exactly one success — now
  the winner also commits the verify in the same tx.

### Integration (`repository_db_integration_test.go`, `//go:build integration`)

- **NEW** `TestRedeemAndVerify_Atomic` — seed a token; redeem + set
  verified in one tx; assert both committed. Then a second case: make
  `SetUserVerified` fail (e.g. non-existent user_id → 0 rows is not an
  error; to force an error, drop the connection mid-tx in a helper or
  inject via a wrapper) — assert the token's `used_at` is **not** set
  (redeem rolled back). If a clean failure-injection is hard, at minimum
  assert the happy path runs both in one tx and document the
  rollback guarantee by code review.
- Keep `TestRedeemToken_Guards` (valid/used/revoked/expired/non-existent)
  — update to the new return shape.

## 6. Verification

```bash
go test -count=1 ./internal/domain/account/...
go test -race -count=1 -timeout 300s ./internal/domain/account/...
go test -tags=integration -count=1 -race ./internal/domain/account/...
go vet ./...
```

`make verify` before merge (includes staticcheck/gosec/govulncheck).

## 7. Risk note (for the remediation PR)

- **Assumptions made:** `RedeemToken` with `RETURNING` is supported by
  goqu's postgres dialect (`Returning(...)` → `RETURNING` clause) and
  pgx `QueryRow.Scan`. Verify with a quick integration test.
- **Edge cases intentionally NOT handled:** the disambiguation
  `FindAuthTokenByHash` (ok==false path) is a read and stays outside
  the redeem tx — it cannot cause a partial commit.
- **Concurrency assumptions:** single-use correctness is still the
  atomic 3-clause `UPDATE … WHERE`; adding `SetUserVerified` to the
  same tx does not weaken it (the guard is on the UPDATE, not the
  app-level read). R12's concurrent test still passes.
- **What is not tested, and why:** forced mid-tx failure injection is
  awkward in pgx; the rollback guarantee is proven by the deferred
  `Rollback` pattern (same as `registerNewUser`) plus the
  `setVerifiedErr` unit test on the fake repo.

## 8. Non-goals

- Do not change the 3-clause predicate.
- Do not change the `202`/`200`/`410`/`404` response mapping (handler
  untouched).
- Do not touch `Register` (Task B).
