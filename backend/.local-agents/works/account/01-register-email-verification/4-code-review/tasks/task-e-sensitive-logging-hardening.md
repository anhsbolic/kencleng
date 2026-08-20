# Task E — Sensitive-Error Logging Hardening

> Ticket    : 01-register-email-verification
> Sub-task  : E of F (post-review remediation)
> Findings  : L1 (breachcheck logs http error verbatim), L2 (service logs
>             notification error verbatim)
> Blocking  : yes
> Back-ref  : `../report.md` §3 (`go/secrets-and-sensitive-logging.md`); root `AGENTS.md` §2

---

## 1. Scope

Two log sites log upstream errors verbatim via `%v`, which can embed
sensitive data:

- **L1:** `internal/platform/breachcheck/client.go:52` logs
  `c.httpClient.Do`'s error via `%v`. A `net/http` client error can
  embed the request URL `api.pwnedpasswords.com/range/{5-char-SHA1-prefix}`
  — partial credential-derived data. The k-anonymity prefix is safe to
  *send* to the server, not to *log*.
- **L2:** `internal/domain/account/service.go:385` logs the
  notification-sender error via `%v`. A real SMTP error (when
  `FakeSender` is replaced) could embed the recipient email or token.
  (`FakeSender` returns nil today, so this is latent — but the seam
  will get a real sender, and the pattern should be right first.)

Per `go/secrets-and-sensitive-logging.md` §1: "Before logging an error
from a third-party client/SDK verbatim, check whether that client's
error type can embed request/response payloads — if so, log a
sanitized summary (error code, category) instead of the raw error
string."

**In scope:**
- Replace the verbatim `%v` of upstream errors with a sanitized
  category/code at both sites.
- Tests asserting the log lines contain no URL / SHA-1 prefix (L1) and
  no recipient / token (L2).

**Out of scope:**
- The `log.Printf("transport: unhandled service error: %v", err)` in
  `transport/http/errors.go:76` — that error chain's leaf is a
  `pgconn.PgError` (SQLSTATE + constraint name; parameterized SQL, no
  PII values), safe to log server-side. Leave it.
- Structured logging migration (the codebase uses `log` stdlib
  throughout; do not introduce a logger package in this task).

## 2. Dependencies

- **Hard deps:** none.
- **Soft deps:** coordinate with Task C (the resend handler log uses
  `%v` on the same service-error chain — apply the same sanitization
  there if Task C lands after this one, or note the dependency).
- **Blocks:** none.

## 3. Files

| File | Change Type | Why |
|---|---|---|
| `internal/platform/breachcheck/client.go` | Edit | sanitize the `Do` error log (L1) |
| `internal/platform/breachcheck/client_test.go` | Edit | assert no URL / SHA-1 prefix in the fail-open log |
| `internal/domain/account/service.go` | Edit | sanitize the `sendVerification` error log (L2) |
| `internal/domain/account/service_test.go` | Edit | assert no recipient / token in the send-failure log |

## 4. Implementation detail

### L1 — `breachcheck/client.go`

Replace the verbatim log with a category. A `*url.Error` from
`http.Client.Do` has a clean shape (`Op`, `URL`, `Err`); extract the
operation and a coarse category, drop the URL:

```go
resp, err := c.httpClient.Do(req)
if err != nil {
    // Fail-open: log a sanitized category, never the raw error (it can
    // embed the request URL, which contains the 5-char SHA-1 prefix of
    // the password — partial credential-derived data). Per
    // go/secrets-and-sensitive-logging.md §1.
    log.Printf("breachcheck: API unreachable (%s), proceeding without check",
        breachErrorCategory(err))
    return false, nil
}
```

where:

```go
// breachErrorCategory reduces a breachcheck HTTP error to a safe,
// PII-free category string for logging. It never returns the request
// URL (which carries the 5-char SHA-1 password-hash prefix).
func breachErrorCategory(err error) string {
    var urlErr *url.Error
    if errors.As(err, &urlErr) {
        // urlErr.Op ("Get"/"Post") is safe; urlErr.URL is NOT (contains the prefix).
        if urlErr.Timeout() {
            return "timeout"
        }
        return fmt.Sprintf("%s: %s", urlErr.Op, classifyNetErr(urlErr.Err))
    }
    // Fallback: a coarse category, not the verbatim message.
    return "transport error"
}

// classifyNetErr maps a low-level net error to a short category.
func classifyNetErr(err error) string {
    var netErr net.Error
    if errors.As(err, &netErr) && netErr.Timeout() {
        return "timeout"
    }
    if errors.Is(err, context.Canceled) {
        return "canceled"
    }
    if errors.Is(err, io.EOF) || errors.Is(err, syscall.ECONNRESET) {
        return "connection reset"
    }
    return "network error"
}
```

(Adjust the categories to taste; the invariant is: **no URL, no SHA-1
prefix, no password-derived string in the log**.)

Also apply the same sanitization to the non-OK-status path at line 57
(currently logs `resp.StatusCode` — that's already safe, an int; leave
it, or keep for consistency).

### L2 — `service.go:sendVerification`

```go
func (s *Service) sendVerification(ctx context.Context, email, plainToken string) {
    if err := s.email.SendVerificationEmail(ctx, email, plainToken); err != nil {
        // Post-commit email failure is non-fatal. Log a sanitized
        // category, not the raw error (a real SMTP error can embed the
        // recipient or token). Per go/secrets-and-sensitive-logging.md §1.
        log.Printf("account: send verification email failed (recipient redacted): %s",
            notificationErrorCategory(err))
    }
}
```

`notificationErrorCategory` can be a thin helper (or inline): if the
error type is known (e.g. a future `*smtp.SMTPError`), extract its code;
otherwise return a coarse `"send failed"`. For v1 (`FakeSender` returns
nil), this path is dormant — the helper just needs to never echo the
raw message. A minimal version:

```go
func notificationErrorCategory(err error) string {
    // Never return err.Error() verbatim — it may embed recipient/token.
    var t interface{ Timeout() bool }
    if errors.As(err, &t) && t.Timeout() {
        return "timeout"
    }
    return "send failed"
}
```

Place the helper next to `sendVerification` / `sendNudge`. Apply the
same to `sendNudge`'s log line (`service.go:393`) for consistency —
currently logs `%v` of the nudge error (nudge has no token, but may
embed the recipient via a real sender).

## 5. Tests to add

### L1 — `breachcheck/client_test.go`

```go
func TestIsBreached_APIUnreachable_LogNoURLNoPrefix(t *testing.T) {
    // Point the client at a non-listening port (or a server that closes
    // the connection) to force a *url.Error carrying the request URL.
    // Capture log output; assert:
    //   - log contains "breachcheck: API unreachable"
    //   - log does NOT contain "pwnedpasswords.com"
    //   - log does NOT contain the 5-char SHA-1 prefix of the test password
}
```

### L2 — `service_test.go`

```go
func TestRegister_SendVerificationFails_LogNoPII(t *testing.T) {
    // Wire a captureSender whose SendVerificationEmail returns an error
    // whose Error() contains the recipient email + a fake token (to
    // simulate a leaky SMTP error). Capture log output; assert:
    //   - log contains "send verification email failed"
    //   - log does NOT contain the recipient email
    //   - log does NOT contain the token
}
```

## 6. Verification

```bash
go test -count=1 ./internal/platform/breachcheck/... ./internal/domain/account/...
go vet ./...
```

## 7. Risk note

- **Assumptions made:** `*url.Error` is the concrete type wrapping
  `http.Client.Do` failures (standard since Go 1.0). If a future
  transport change yields a different type, the fallback branch returns
  a coarse category rather than the verbatim message — safe by
  construction.
- **Edge cases intentionally NOT handled:** the `transport/http`
  `MapServiceError` log (`errors.go:76`) is left as `%v` because its
  leaf is a DB driver error with no PII values (parameterized SQL).
  Documented here so a future reviewer doesn't "fix" it the same way
  and accidentally broaden the change.
- **Concurrency assumptions:** none new.
- **What is not tested, and why:** nothing — both sites get positive
  (error present) and negative (no PII) assertions.

## 8. Non-goals

- Do not introduce a structured logger (`zap`/`slog`) — the codebase
  uses `log` stdlib; stay consistent.
- Do not change `transport/http/errors.go:76` (DB-error chain, safe).
- Do not redact `nudgeType` in `sendNudge` (it's a package constant,
  not user input — safe).
