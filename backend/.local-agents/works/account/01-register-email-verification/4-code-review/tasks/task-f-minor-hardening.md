# Task F — Minor Hardening (Optional)

> Ticket    : 01-register-email-verification
> Sub-task  : F of F (post-review remediation)
> Findings  : R15 (no rate-limit test), S6 (sweeper no exit), H1 (body
>             not drained), Q1 (runTx helper), Q2 (test && bug), Q3
>             (shared email helper), E1 (HMAC≠enc key runtime check)
> Blocking  : no — optional / follow-up
> Back-ref  : `../report.md` §2 (Quality), §3 (Best Practices)

---

## 1. Scope

The minor/optional findings from the review. None block merge; together
they tidy the code and close the last checklist items. Land as a single
follow-up PR, or fold individual items into Tasks A–E where the file
overlap is natural (e.g. Q2's test fix can ride along with Task A's
test edits).

**In scope:** the seven items below.
**Out of scope:** any behavior change beyond what each item describes.

## 2. Dependencies

- **Hard deps:** Tasks A and B (Q2's `&&`→`||` fix is in
  `service_test.go`, which A and B also edit — land after them to avoid
  conflicts).
- **Soft deps:** none.
- **Blocks:** none.

## 3. Items

### R15 — `TestResend_RateLimited` (rate-limit N+1 rejection)

**Where:** `internal/transport/http/middleware_test.go` (new) or
`middleware_test.go` next to `middleware.go`.
**Best-practices file:** `go/rate-limiting.md` checklist item 4 ("a test
that verifies the limiter actually rejects the N+1th request").

```go
func TestRateLimit_RejectsNPlusOne(t *testing.T) {
    // rps=1, burst=1: first request allowed, second immediately rejected (429).
    h := RateLimit(1, 1)(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
        w.WriteHeader(http.StatusOK)
    }))
    rr1 := newReq("1.2.3.4:5")
    rr2 := newReq("1.2.3.4:5")
    // first OK
    // second → 429 with application/problem+json
}
```

Also assert: different IPs get independent buckets (per-IP granularity,
`go/rate-limiting.md` item 2), and that the eviction map doesn't grow
unbounded (optional).

### S6 — sweeper goroutine exit path

**Where:** `internal/transport/http/middleware.go:29-43`.
**Best-practices file:** `go/goroutine-lifecycle.md` §1.

Accept a `context.Context` in `RateLimit` and have the sweeper
`select` on `ctx.Done()`:

```go
func RateLimit(ctx context.Context, rps float64, burst int) func(http.Handler) http.Handler {
    // ...
    go func() {
        t := time.NewTicker(time.Minute)
        defer t.Stop()
        for {
            select {
            case <-ctx.Done():
                return
            case <-t.C:
                // sweep
            }
        }
    }()
    // ...
}
```

Wire `main.go` to pass the `signal.NotifyContext` ctx (or a derived
process ctx) so the sweeper stops on shutdown. (Note: this changes
`RateLimit`'s signature — update `main.go:119`.)

### H1 — drain breachcheck body on non-OK path

**Where:** `internal/platform/breachcheck/client.go:56-58`.
**Best-practices file:** `go/http-client-and-transport.md` §2 ("Fully
read and close the response body on every code path … an unclosed or
partially-read body prevents the underlying connection from being
returned to the pool for reuse").

On the non-OK status path, drain before close:

```go
defer resp.Body.Close()
if resp.StatusCode != http.StatusOK {
    io.Copy(io.Discard, resp.Body) // drain so the connection returns to the pool
    log.Printf("breachcheck: API returned status %d, proceeding without check", resp.StatusCode)
    return false, nil
}
```

### Q1 — `runTx` helper

**Where:** `internal/domain/account/service.go` (`registerNewUser:178`
and `issueNewVerificationToken:254`).

Extract the duplicated
`committed := false; defer func(){ if !committed { _ = tx.Rollback(ctx) } }()`
pattern:

```go
// runTx begins a transaction, runs fn, and commits on nil error or
// rolls back on non-nil. The deferred rollback is a no-op after a
// successful commit.
func (s *Service) runTx(ctx context.Context, fn func(pgx.Tx) error) error {
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
    if err := fn(tx); err != nil {
        return err
    }
    if err := tx.Commit(ctx); err != nil {
        return fmt.Errorf("account: commit: %w", err)
    }
    committed = true
    return nil
}
```

Then `registerNewUser` and `issueNewVerificationToken` (and Task A/B's
new tx uses) call `s.runTx(ctx, func(tx pgx.Tx) error { … })`. Note:
`registerNewUser` needs `plainToken` returned and the email sent
*after* commit — `runTx` returns only `error`, so the post-commit
side-effect stays in the caller (commit happens inside `runTx`, then
the caller sends). This is fine: `runTx` owns the tx lifecycle, the
caller owns the post-commit side-effect.

### Q2 — `&&`→`||` in `TestResend_NoMatch`

**Where:** `internal/domain/account/service_test.go:810`.

```go
// before:
if len(sender.verificationTo) != 0 && len(sender.nudgeTypes) != 0 {
// after:
if len(sender.verificationTo) != 0 || len(sender.nudgeTypes) != 0 {
```

As written (`&&`) it only fails if *both* are non-zero; should be `||`
to catch either being non-zero.

### Q3 — shared `looksLikeEmail` helper

**Where:** `internal/transport/http/auth_register.go:77-85` and
`auth_verify_email.go:59`.

Move `looksLikeEmail` to a shared file in `transport/http/` (e.g.
`validate.go`) so the register and resend handlers share one
implementation. Pure refactor, no behavior change.

### E1 — runtime check that HMAC ≠ encryption key

**Where:** `internal/platform/crypto/keys.go` (`New`).
**Best-practices file:** `postgresql/encryption-at-rest.md` checklist
item 3 ("The HMAC key and encryption key are different keys").

The keys are structurally separate (separate env vars, separate fields),
but the code does not *enforce* they differ at runtime — config
discipline only. Add a check in `New`:

```go
if bytes.Equal(encryptionKey, hmacKey) {
    return nil, fmt.Errorf("ENCRYPTION_KEY and HMAC_KEY must be different keys (got the same value)")
}
```

This is a one-time startup check; it catches a misconfigured env before
the server boots. (Import `bytes`.)

## 4. Verification

```bash
go test -count=1 ./...
go test -race -count=1 ./internal/domain/account/... ./internal/platform/crypto/...
go vet ./...
make verify
```

## 5. Risk note

- **Assumptions made:** none of these change behavior — R15 adds a
  test, S6 adds an exit path, H1 drains a body, Q1/Q3 are refactors, Q2
  is a test fix, E1 adds a startup guard.
- **Edge cases intentionally NOT handled:** none.
- **Concurrency assumptions:** none new.
- **What is not tested, and why:** nothing — each item is either a test
  addition or a behavior-preserving refactor with existing coverage.

## 6. Non-goals

- Do not fold these into the blocking PR if it would delay it — they
  are explicitly optional and can be a follow-up.
- E1 must not change key *loading* (base64 decode, 32-byte validation)
  — only add the inequality check.
