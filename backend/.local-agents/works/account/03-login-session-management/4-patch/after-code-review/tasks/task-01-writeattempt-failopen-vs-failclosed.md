# Task 01: Reconcile `writeAttempt` fail-open doc vs fail-closed code + add test

> Back-reference : `4-code-review/report.md` §2 (Q1) · `2-plan/techplan.md` §13 (common mistakes) · `3-build/report.md` deviation #5
> Priority       : **Blocking** — must land before any commit
> Model          : GLM 5.2 (max) — security-relevant error-path change
> Depends on     : none (surgical edit to existing code)

## Objective

The `writeAttempt` helper's doc comment and the build report's deviation #5
describe **fail-open** behavior (a lost audit-row write is logged and the
credential decision stands), but the code returns the write error —
**fail-closed** (a transient DB hiccup turns a valid login into a 500).
Neither path is tested because `loginFakeRepo.InsertLoginAttempt` always
returns nil. Pick one behavior, make doc/code/test agree.

**Recommended choice: fail-open (option a).** The doc's reasoning is sound
for a non-financial audit table: a lost attempt row can only undercount
toward the lockout threshold, never lock anyone out spuriously. Fail-closed
means an audit-DB outage takes down all logins — a worse failure mode than
a missing audit row. The alternative (option b) is documented below.

## ⚠️ Human decision required

This task defaults to **option (a) fail-open**. If Anhar prefers **option (b)
fail-closed**, adjust the task before execution — the code change and test
shape differ. Do not execute without confirming the choice.

## Files

| File | Change |
|---|---|
| `backend/internal/domain/account/login.go` | Edit `writeAttempt` — implement fail-open (option a) or fix doc (option b) |
| `backend/internal/domain/account/login_test.go` | New test proving the chosen behavior |

## Option (a) — fail-open (RECOMMENDED)

### Code change (`login.go` `writeAttempt`)

Current code returns the error on every write failure. Change it to log the
error and return nil, so the credential decision stands:

```go
// writeAttempt persists one attempt row in its own short transaction. The
// attempt row is bookkeeping, not a state machine: if the write fails after
// the credential decision, the login result stands, the failure is logged,
// and the gap is observable — acceptable for a non-financial audit table
// (a lost attempt row can only undercount toward the lockout threshold,
// never lock anyone out spuriously).
func (s *Service) writeAttempt(ctx context.Context, attempt *LoginAttempt) error {
    tx, err := s.tx.BeginTx(ctx)
    if err != nil {
        log.Printf("account: begin attempt tx failed (fail-open): %v", err)
        return nil // fail-open: credential decision stands
    }
    committed := false
    defer func() {
        if !committed {
            _ = tx.Rollback(ctx)
        }
    }()
    if err := s.repo.InsertLoginAttempt(ctx, tx, attempt); err != nil {
        log.Printf("account: insert login attempt failed (fail-open): stage=%s success=%t: %v",
            attempt.Stage, attempt.Success, err)
        return nil // fail-open: credential decision stands
    }
    if err := tx.Commit(ctx); err != nil {
        log.Printf("account: commit attempt tx failed (fail-open): stage=%s success=%t: %v",
            attempt.Stage, attempt.Success, err)
        return nil // fail-open: credential decision stands
    }
    committed = true
    return nil
}
```

**Log hygiene note (R19):** the error value logged here comes from pgx/goqu
— it may contain SQL detail but NOT user credentials or tokens (the attempt
struct holds only `identifier_hash`, `user_id`, `stage`, `success`). Logging
the error is safe per the "no secrets in logs" golden rule; the attempt
fields logged (`stage`, `success`) are non-sensitive metadata. Do NOT log
`attempt.IdentifierHash` or `attempt.UserID` in the error line — keep it to
`stage` + `success` + the wrapped error.

### Test (`login_test.go`)

Add a fake repo variant whose `InsertLoginAttempt` returns an error, then
assert the login still succeeds:

```go
// failingAttemptRepo wraps loginFakeRepo and forces InsertLoginAttempt
// to fail. Used to prove writeAttempt's fail-open contract: a lost audit
// row must NOT block a valid login.
type failingAttemptRepo struct {
    *loginFakeRepo
}

func (f *failingAttemptRepo) InsertLoginAttempt(_ context.Context, _ pgx.Tx, _ *LoginAttempt) error {
    return errors.New("simulated audit-db outage")
}

// TestWriteAttempt_FailOpen_ValidLoginStillSucceeds proves that a
// login_attempts write failure does not block a valid credential login
// (the audit row is bookkeeping, not a state machine — a lost row can
// only undercount toward lockout, never lock spuriously).
func TestWriteAttempt_FailOpen_ValidLoginStillSucceeds(t *testing.T) {
    h := newLoginHarness(t)
    credHash, _ := secrets.HashPassword("correct-horse-battery")
    h.seedIdentity(t, "failopen@example.com", credHash)

    // Swap in the failing repo — all other methods delegate to loginFakeRepo.
    h.svc.repo = &failingAttemptRepo{loginFakeRepo: h.repo}

    res, err := h.svc.Login(context.Background(), "failopen@example.com", "correct-horse-battery")
    if err != nil {
        t.Fatalf("valid login blocked by audit-write failure (must be fail-open): %v", err)
    }
    if res.Status != "ok" || res.AccessToken == "" {
        t.Errorf("login result wrong despite fail-open: %+v", res)
    }
}
```

Also add a companion asserting a **failed credential** login still returns
`ErrInvalidCredentials` (not a 500) when the audit write also fails — the
fail-open path must not mask the credential rejection:

```go
// TestWriteAttempt_FailOpen_InvalidCredentialsStillRejected proves the
// fail-open path does not mask a credential failure: wrong password +
// audit-write error ⇒ still ErrInvalidCredentials, not nil.
func TestWriteAttempt_FailOpen_InvalidCredentialsStillRejected(t *testing.T) {
    h := newLoginHarness(t)
    credHash, _ := secrets.HashPassword("correct-horse-battery")
    h.seedIdentity(t, "failopen-reject@example.com", credHash)
    h.svc.repo = &failingAttemptRepo{loginFakeRepo: h.repo}

    _, err := h.svc.Login(context.Background(), "failopen-reject@example.com", "wrong-password")
    if !errors.Is(err, ErrInvalidCredentials) {
        t.Fatalf("fail-open masked credential rejection: err = %v, want ErrInvalidCredentials", err)
    }
}
```

### Doc / report updates

- `writeAttempt` doc already matches fail-open — no doc change needed for option (a).
- `3-build/report.md` deviation #5 already describes fail-open — no report change needed for option (a). (If the report has already been committed in a prior state, verify it reads "fail-open" and update if stale.)

## Option (b) — fail-closed (ALTERNATIVE, if Anhar prefers)

### Code change

Keep the current code as-is (it returns the error → 500). Only the doc and
report need updating.

### Doc change (`login.go` `writeAttempt`)

Replace the current doc with:

```go
// writeAttempt persists one attempt row in its own short transaction. If
// the write fails, the error is returned and the login is rejected (fail-
// closed): a transient audit-DB outage will block logins rather than
// silently lose audit rows. This is a deliberate trade-off — the audit
// trail's completeness is prioritized over login availability during a
// DB hiccup. A lost attempt row would undercount toward lockout (never
// lock spuriously), but the chosen stance is that a missing audit row
// for a security-sensitive flow is itself a failure condition.
```

### Report change (`3-build/report.md` deviation #5)

Update deviation #5 to state fail-closed is the actual choice, replacing
the current "undercount toward lockout" framing.

### Test (`login_test.go`)

```go
// TestWriteAttempt_FailClosed_ValidLoginRejected proves fail-closed:
// a valid credential login + audit-write failure ⇒ error (not nil),
// so the audit trail's completeness is prioritized over availability.
func TestWriteAttempt_FailClosed_ValidLoginRejected(t *testing.T) {
    h := newLoginHarness(t)
    credHash, _ := secrets.HashPassword("correct-horse-battery")
    h.seedIdentity(t, "failclosed@example.com", credHash)
    h.svc.repo = &failingAttemptRepo{loginFakeRepo: h.repo}

    _, err := h.svc.Login(context.Background(), "failclosed@example.com", "correct-horse-battery")
    if err == nil {
        t.Fatal("valid login succeeded despite audit-write failure (must be fail-closed)")
    }
    if errors.Is(err, ErrInvalidCredentials) || errors.Is(err, ErrLockedOut) {
        t.Fatalf("audit-write failure surfaced as wrong sentinel: %v", err)
    }
    // err is an internal error (→ 500 at transport) — that's the expected class.
}
```

## Verification

```bash
# Unit tests (option a or b — the relevant test name differs)
go test ./internal/domain/account/ -run TestWriteAttempt -v

# Full unit suite + race
go test -race ./...

# Gate (will still be red on gosec pre-existing baseline — that's expected)
make verify
```

## Out of scope

- The Tier 0 paired rewrite pass (separate human gate)
- task-02 (TTL constant dedup), task-03 (logout doc) — independent
