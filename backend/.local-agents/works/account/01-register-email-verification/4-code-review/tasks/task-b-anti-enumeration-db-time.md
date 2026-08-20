# Task B — Anti-Enumeration DB-Time Uniformity (R3/R4)

> Ticket    : 01-register-email-verification
> Sub-task  : B of F (post-review remediation)
> Finding   : S3 (R3/R4 perform no DB writes → timing side-channel)
> Blocking  : yes
> Back-ref  : `../report.md` §1 (S3); `../../2-plan/techplan.md` §4 R7, §5 Decision 8

---

## 1. Scope

The build implemented the CPU-time half of R7 (always-bcrypt on every
branch) but **not** the DB-time half: R3 (verified existing) and R4
(google-only conflict) perform zero DB writes, while R1 (new user)
does 3 inserts + commit and R2 (unverified) does revoke + insert +
commit. Against real Postgres, R3/R4 are measurably faster — an
enumeration side-channel that leaks "verified/google-only" (fast) vs
"new/unverified" (slow).

The techplan explicitly committed to "DB-write-shaped work on all
branches" (`2-plan/techplan.md` §5 Decision 8: "For DB-time uniformity,
all branches perform DB-write-shaped work per Assumption B"). The build
omitted it for R3/R4.

The existing `TestRegister_GenericResponse_Timing` uses the in-memory
`fakeRepo` (DB ops ~microseconds), so it **cannot** catch DB-timing
differences — it only proves bcrypt equivalence. It gives false
confidence.

**In scope:**
- Give R3 and R4 a DB-write-shaped operation equivalent in cost to
  R1/R2, discarded on those no-op branches.
- Add a timing test that actually exercises DB latency (real Postgres
  integration test, or a fake repo that injects per-op latency) and
  asserts all four branches stay within a band.

**Out of scope:**
- Changing the `202` generic response shape (already identical).
- The bcrypt CPU-time half (already correct — keep it).
- `VerifyEmail` (Task A owns it).

## 2. Dependencies

- **Hard deps:** none.
- **Soft deps:** coordinate with Task A (same file, disjoint method).
- **Blocks:** Task F (test assertions).

## 3. Files

| File | Change Type | Why |
|---|---|---|
| `internal/domain/account/service.go` | Edit | R3/R4 branches do a no-op DB write |
| `internal/domain/account/repository.go` | Edit (likely) | add a `NoopWrite(ctx, tx)` or reuse `RevokeTokens` with a synthetic key if a dedicated method reads cleaner |
| `internal/domain/account/repository_db.go` | Edit | implement the no-op write |
| `internal/domain/account/service_test.go` | Edit | fake repo records the no-op call; timing test moves to a latency-injecting fake or integration |
| `internal/domain/account/repository_db_integration_test.go` | Edit (likely) | NEW integration timing test against real Postgres |

## 4. Implementation detail

### Approach: "dummy revoke" on R3/R4

The cleanest equivalent-cost write is a `RevokeTokens`-shaped UPDATE that
affects 0 rows on the no-op branches — same statement shape, same
round-trip count as R2, but touching no real rows. Two options:

**Option 1 (preferred) — revoke against a synthetic user_id:**
On R3/R4, call `RevokeTokens(ctx, tx, uuid.New(), purposeEmailVerify)`.
The `WHERE user_id = $1 AND ...` matches 0 rows (no such user), so it's
a no-op UPDATE — same DB round-trip and lock shape as R2's real revoke,
but affects nothing. Cost: 1 UPDATE + commit, matching R2.

```go
// R3 (verified existing) and R4 (google-only): perform a DB-write-
// shaped no-op so wall-clock time matches R1/R2 (anti-enumeration, R7).
// A revoke against a non-existent user_id affects 0 rows but has the
// same UPDATE+commit cost shape as the real revoke in R2.
if err := s.dummyWrite(ctx); err != nil {
    return fmt.Errorf("account: timing-shaping write: %w", err)
}
```

where `dummyWrite` begins a tx, calls `RevokeTokens(ctx, tx, uuid.New(),
purposeEmailVerify)`, commits. (R1/R2 already do real writes inside
their own tx, so they don't need the dummy.)

**Option 2 — a dedicated `NoopWrite` repository method:**
Adds a `NoopWrite(ctx, tx) error` that runs a trivial UPDATE (e.g.
`UPDATE auth_tokens SET created_at = created_at WHERE id = $1` with a
random uuid — 0 rows, same round-trip). More explicit but more surface
area. Pick Option 1 unless the team prefers the named method.

### Why this satisfies R7

- **CPU time:** bcrypt still runs on all branches (unchanged).
- **DB time:** every branch now does ≥1 write + commit:
  - R1: tx(3 inserts + commit)
  - R2: tx(revoke + insert + commit)
  - R3: tx(dummy revoke 0 rows + commit) + sendNudge
  - R4: tx(dummy revoke 0 rows + commit) + sendNudge
- All four perform a `BeginTx` + ≥1 `UPDATE`/`INSERT` + `Commit` —
  equivalent DB-round-trip shape. Residual difference (3 inserts vs 1)
  is within the noise band a timing test can assert.

### `Register` branch sketch (R3/R4 only)

```go
case identity == nil:
    google, err := s.repo.FindAuthIdentityByIdentifierHash(ctx, providerGoogle, identifierHash)
    if err != nil {
        return fmt.Errorf("account: lookup google identity: %w", err)
    }
    if google != nil {
        // R4: Google-only conflict. DB-write-shaped no-op for timing
        // equivalence with R1/R2 (R7), then the nudge.
        if err := s.dummyWrite(ctx); err != nil {
            return fmt.Errorf("account: timing write: %w", err)
        }
        s.sendNudge(ctx, email, notification.NudgeGoogleOnly)
        return nil
    }
    return s.registerNewUser(ctx, name, email, passwordHash) // R1 unchanged

case identity.VerifiedAt == nil:
    // R2 unchanged (already does revoke + insert in tx)
    ...

default:
    // R3: verified existing. DB-write-shaped no-op for timing
    // equivalence (R7), then the password-reset nudge.
    if err := s.dummyWrite(ctx); err != nil {
        return fmt.Errorf("account: timing write: %w", err)
    }
    s.sendNudge(ctx, email, notification.NudgePasswordReset)
    return nil
```

`dummyWrite` wraps the tx (begin/commit/rollback) — reuse the
`committed := false; defer …` pattern (or the `runTx` helper if Task
F's Q1 lands first).

## 5. Tests to add / update

### Unit (`service_test.go`)

- `fakeRepo.RevokeTokens` already records calls; assert R3/R4 each
  produce one revoke call (with a synthetic user_id) — proves the
  write happens.
- **NEW** `TestRegister_R3R4_PerformTimingWrite` — assert R3 and R4
  each invoke exactly one `RevokeTokens` (the dummy), so the no-op
  branches are no longer write-free.
- Keep `TestRegister_GenericResponse_AllBranches` (1 email per branch
  unchanged).

### Timing — move off the microsecond fake

The current `TestRegister_GenericResponse_Timing` cannot catch DB-timing
gaps because `fakeRepo` ops are ~microseconds. Two acceptable fixes
(pick one):

**Fix 1 (preferred): integration timing test.** Add
`TestRegister_Timing_AllBranches_RealPostgres` under
`//go:build integration` that runs each branch against real Postgres
and asserts `max/min ≤ 2×` (tighter than the 3× bcrypt-only band,
because now DB writes are present on all branches). Run with
`go test -tags=integration -race ./internal/domain/account/...`.

**Fix 2: latency-injecting fake.** Add a `sleepRepo` wrapper over
`fakeRepo` that adds a configurable `time.Sleep` per DB op (e.g. 5ms per
`Exec`), so the timing test on the fake reflects DB-shape cost. Less
realistic but runs in the fast unit suite.

Either way, **add a comment to the timing test** stating it now covers
DB-time equivalence, not just bcrypt — and that Fix 1 is the
authoritative one.

## 6. Verification

```bash
go test -count=1 ./internal/domain/account/...
go test -race -count=1 -timeout 300s ./internal/domain/account/...
go test -tags=integration -count=1 -race ./internal/domain/account/...
```

## 7. Risk note

- **Assumptions made:** `RevokeTokens` against a non-existent
  `user_id` affects 0 rows (no FK violation — `user_id` is not
  constrained to exist in the `WHERE`, only on INSERT via FK). Verify
  with the integration test.
- **Edge cases intentionally NOT handled:** the exact wall-clock
  equivalence is statistical; the test asserts a band (≤2×), not exact
  equality. Residual insert-count difference (R1's 3 vs R3's 1) is
  within that band against real Postgres.
- **Concurrency assumptions:** the dummy write is a per-request tx;
  no shared state. `-race` clean by construction.
- **What is not tested, and why:** network/jitter is not modeled; the
  integration timing test uses real Postgres latency as the realistic
  signal.

## 8. Non-goals

- Do not change the `202` response.
- Do not remove bcrypt (CPU-time half is correct).
- Do not touch `VerifyEmail` (Task A).
