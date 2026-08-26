# Task C — `ResendVerification` Handler Error Logging

> Ticket    : 01-register-email-verification
> Sub-task  : C of F (post-review remediation)
> Finding   : S4 (swallowed error, no logging)
> Blocking  : yes
> Back-ref  : `../report.md` §1 (S4); root `AGENTS.md` §2 golden rule

---

## 1. Scope

`ResendVerificationHandler` discards the service error entirely:

```go
_ = svc.ResendVerification(r.Context(), req.Email)
// Always 202 generic …
```

If `issueNewVerificationToken` fails (DB error), no token is issued and
no email is sent, but the user gets `202 "you will receive a new
verification link."` Anti-enumeration justifies the 202 *response*; it
does **not** justify the absence of logging. The golden rule allows
"logged with enough context to act on" — this does neither.

**In scope:**
- Log the error server-side before returning the 202.
- A test asserting the log line appears (and contains no PII).

**Out of scope:**
- Changing the 202 response (must stay identical for anti-enumeration).
- The `Register` / `VerifyEmail` handlers (already log via
  `MapServiceError`).

## 2. Dependencies

- **Hard deps:** none.
- **Soft deps:** none.
- **Blocks:** none.

## 3. Files

| File | Change Type | Why |
|---|---|---|
| `internal/transport/http/auth_verify_email.go` | Edit | log the error before 202 |
| `internal/transport/http/auth_verify_email_test.go` (new, or extend an existing handler test file) | New/Edit | assert log line present + no PII |

## 4. Implementation detail

```go
if err := svc.ResendVerification(r.Context(), req.Email); err != nil {
    // Anti-enumeration: the response is still the identical 202 generic
    // (R14). But the internal failure must be visible to ops — log it.
    // Do not log the email (PII); log the error category only.
    log.Printf("transport: resend verification failed (recipient redacted): %v", err)
}

// Always 202 generic — the service returns nil for both match and
// no-match branches (R13/R14).
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusAccepted)
_ = json.NewEncoder(w).Encode(map[string]string{
    "message": "If your email is registered and unverified, you will receive a new verification link.",
})
```

**Note on the `%v` of `err`:** the service error from
`ResendVerification` is wrapped `fmt.Errorf("…: %w", err)` whose leaf
is a `pgconn.PgError` (SQLSTATE + constraint name, no PII values —
queries are parameterized). This is safe to log server-side. If Task E's
L2 hardening lands a sanitized-error helper, prefer that here too; for
now `%v` on this specific error chain is acceptable because the leaf is
a DB driver error, not a third-party HTTP error.

### Why not route through `MapServiceError`?

`Register` routes non-validation errors through `MapServiceError`
(which logs + returns a 500/429). For resend, returning a 500 on the
match branch would *distinguish* it from the no-match branch (202) —
an enumeration leak. So the response must stay 202; only the *log* is
added.

## 5. Tests to add

A handler-level test (the build deferred handler tests — this is a
good first one):

```go
func TestResendVerificationHandler_ServiceError_Still202_ButLogs(t *testing.T) {
    // Inject a service that returns an error from ResendVerification.
    // Assert: response is 202 (not 500), and a log line was emitted
    // containing "resend verification failed" and NOT containing the
    // recipient email.
    var logBuf bytes.Buffer
    origOut := log.Writer()
    log.SetOutput(&logBuf)
    defer log.SetOutput(origOut)

    svc := &failingResendSvc{} // returns errors.New("boom")
    h := ResendVerificationHandler(svc)
    h.ServeHTTP(rr, req)

    if rr.Code != http.StatusAccepted {
        t.Errorf("status = %d, want 202 (anti-enumeration)", rr.Code)
    }
    if !strings.Contains(logBuf.String(), "resend verification failed") {
        t.Errorf("expected failure log, got: %q", logBuf.String())
    }
    if strings.Contains(logBuf.String(), "leak@example.com") {
        t.Errorf("log leaked recipient email: %q", logBuf.String())
    }
}
```

If a full handler test setup is heavy, at minimum add a test that drives
the service with a fake repo returning an error and asserts the handler
emits a log line + 202. (The build's risk note already flagged
handler-level tests as deferred — this task is the right place to
start them, scoped to this one assertion.)

## 6. Verification

```bash
go test -count=1 ./internal/transport/http/...
go vet ./...
```

## 7. Risk note

- **Assumptions made:** the `%v` of the wrapped service error is safe
  to log because the leaf is a pgx/DB error (SQLSTATE, constraint name
  — no PII values, parameterized SQL). If a future resend path adds a
  third-party call, switch to Task E's sanitized helper.
- **Edge cases intentionally NOT handled:** none — the change is a
  one-line log addition.
- **Concurrency assumptions:** none new.
- **What is not tested, and why:** nothing — the test covers the
  behavior fully (202 + log + no PII).

## 8. Non-goals

- Do not change the 202 response or its body.
- Do not add `MapServiceError` routing (would break anti-enumeration).
