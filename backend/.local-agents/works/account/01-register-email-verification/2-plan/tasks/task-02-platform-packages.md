# Task 02 — Platform Packages (secrets, breachcheck, notification)

> Ticket    : 01-register-email-verification
> Sub-task  : 2 of 5
> Axis      : Dependency/sequence chain (primary) + component boundary
> Status    : Ready (parallelizable with Task 01)
> Back-ref  : `../2-plan/techplan.md` (originating contract techplan — cross-check high-level decisions there whenever needed)

---

## 1. Scope

Three independent `internal/platform/` packages, each a thin shared-infra
seam with no business rules. They are authored together because none
depends on the others and each is small — but internally they can be
implemented in any order, even in parallel.

**In scope:**
- `platform/secrets/` — bcrypt wrapper (`HashPassword`, `ComparePassword`)
- `platform/breachcheck/` — HaveIBeenPwned k-anonymity HTTP client,
  explicit timeout, fail-open
- `platform/notification/` — `Sender` interface + `FakeSender` (logged,
  no SMTP)

**Out of scope:**
- Real SMTP delivery — v1 uses fake/logged sender per
  `kencleng-phase3-detail.md`
- PII encryption (`platform/crypto/`) — Tier 0 fenced, human-authored
  prerequisite (see manifest)
- Any business logic that *consumes* these packages — that lives in
  Task 04's `domain/account/service.go`

## 2. Dependencies

- **Hard deps:** Task 01 (`golang.org/x/crypto` direct in `go.mod` for
  bcrypt; standard library only for the other two)
- **Soft deps:** none
- **Blocks:** Task 04 (service needs all three)

## 3. Files

| File | Change Type |
|---|---|
| `backend/internal/platform/secrets/secrets.go` | New |
| `backend/internal/platform/secrets/secrets_test.go` | New |
| `backend/internal/platform/breachcheck/client.go` | New |
| `backend/internal/platform/breachcheck/client_test.go` | New |
| `backend/internal/platform/notification/sender.go` | New |
| `backend/internal/platform/notification/sender_test.go` | New |

## 4. Implementation detail

### `backend/internal/platform/secrets/secrets.go` (new)

Thin bcrypt wrapper. Rationale (techplan §5 Decision 2): password
hashing is used by register, reset-password (task #4), set-password
(task #5) — three features — so a wrapper avoids repeating the bcrypt
import pattern. `platform/secrets/` is for **credential hashing**,
distinct from `platform/crypto/`'s PII encryption.

```go
// Package secrets provides credential hashing primitives.
// It is distinct from platform/crypto, which handles PII encryption at rest.
package secrets

import "golang.org/x/crypto/bcrypt"

// HashPassword returns a bcrypt hash of the password using the default cost.
func HashPassword(password string) (string, error) {
    h, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", fmt.Errorf("secrets: hash password: %w", err)
    }
    return string(h), nil
}

// ComparePassword reports whether the password matches the stored hash.
// Returns a non-nil error (bcrypt.ErrMismatchedHashAndPassword) on mismatch.
func ComparePassword(hash, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
```

Notes:
- Default cost (~100ms) is what underpins the "always bcrypt"
  constant-time approach in Task 04 — do not lower it.
- Errors wrapped with `fmt.Errorf("...: %w", err)` per AGENTS.md §2.

### `backend/internal/platform/breachcheck/client.go` (new)

HaveIBeenPwned k-anonymity client. Rationale (techplan §5 Decision 4):
external HTTP call with its own timeout/fail-open logic, reusable by
reset-password (task #4). Client constructed **once** at init with
explicit timeout and **reused** across calls — never
`http.DefaultClient` (no timeout → hung goroutine, techplan §7 risk
row 4).

```go
// Package breachcheck checks passwords against the HaveIBeenPwned
// k-anonymity service. It fails open: if the API is unreachable,
// IsBreached returns (false, nil) and logs the fact.
package breachcheck

import (
    "context"
    "crypto/sha1"
    "encoding/hex"
    "fmt"
    "io"
    "net/http"
    "strings"
    "time"
)

type Client struct {
    httpClient *http.Client
    baseURL    string
}

// NewClient constructs a breach-check client with the given per-call
// timeout. The returned client is safe for concurrent use and should
// be reused — do not construct one per call.
func NewClient(timeout time.Duration) *Client {
    return &Client{
        httpClient: &http.Client{Timeout: timeout},
        baseURL:    "https://api.pwnedpasswords.com/range",
    }
}

// IsBreached reports whether the password has appeared in a known
// breach. On API unreachable it returns (false, nil) — fail-open,
// logged via the provided logger. Never returns the password or its
// hash in any error or log line.
func (c *Client) IsBreached(ctx context.Context, password string) (bool, error) {
    sum := sha1.Sum([]byte(password))
    full := strings.ToUpper(hex.EncodeToString(sum[:]))
    prefix, suffix := full[:5], full[5:]

    req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+"/"+prefix, nil)
    if err != nil {
        return false, fmt.Errorf("breachcheck: build request: %w", err)
    }
    resp, err := c.httpClient.Do(req)
    if err != nil {
        // Fail-open: log the fact + category, not the password.
        log.Printf("breachcheck: API unreachable, proceeding without check: %v", err)
        return false, nil
    }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        log.Printf("breachcheck: API returned status %d, proceeding without check", resp.StatusCode)
        return false, nil
    }
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return false, fmt.Errorf("breachcheck: read response: %w", err)
    }
    for _, line := range strings.Split(string(body), "\n") {
        if strings.HasPrefix(line, suffix+":") {
            return true, nil
        }
    }
    return false, nil
}
```

Notes:
- Only the first 5 hex chars of the SHA-1 prefix leave the process —
  the k-anonymity guarantee.
- Fail-open path: log "API unreachable" + status code only; never the
  password, the full hash, or the raw error string (which may embed
  response bodies). Per `go/secrets-and-sensitive-logging.md` §1.
- `http.NewRequestWithContext` gives ctx-aware deadline on top of the
  client timeout — per `go/http-client-and-transport.md`.

### `backend/internal/platform/notification/sender.go` (new)

Email sender seam. Rationale (techplan §5 Decision 3): v1 email is
fake/logged (no SMTP). The interface creates a clean seam — when real
SMTP is added later (notification domain), only the implementation
swaps, not the service that depends on it.

```go
// Package notification defines the email-sending seam used by the
// account domain. v1 ships FakeSender (logged, no SMTP); a real SMTP
// implementation will be added in the notification domain phase.
package notification

import "context"

// Sender abstracts outbound email delivery.
// Implementations must not block on network I/O inside a caller's
// DB transaction — callers are responsible for sending after commit.
type Sender interface {
    SendVerificationEmail(ctx context.Context, to, token string) error
    SendNudgeEmail(ctx context.Context, to, nudgeType string) error
}

// FakeSender logs the recipient and message type instead of sending.
// Suitable for v1 and for tests. Does not perform any network I/O.
type FakeSender struct{}

// NewFakeSender returns a Sender that logs rather than sends.
func NewFakeSender() *FakeSender { return &FakeSender{} }

// Nudge type constants — keep in sync with service calls.
const (
    NudgeResendVerification = "resend_verification"
    NudgePasswordReset      = "password_reset"
    NudgeGoogleOnly         = "google_only"
)

func (FakeSender) SendVerificationEmail(ctx context.Context, to, token string) error {
    // Log the fact, not the payload. Do not log the token.
    log.Printf("notification: verification email queued (recipient redacted)")
    return nil
}

func (FakeSender) SendNudgeEmail(ctx context.Context, to, nudgeType string) error {
    log.Printf("notification: nudge email queued type=%s (recipient redacted)", nudgeType)
    return nil
}
```

Notes:
- **No PII in logs.** `FakeSender` deliberately does not log the
  recipient address or the verification token — AGENTS.md golden rule,
  reinforced by techplan §13 "Logging email in plaintext" mistake row.
- Interface comment makes the "send after commit" contract explicit so
  the service author (Task 04) doesn't accidentally hold a DB
  transaction open during send.

## 5. Rules covered

This task is infrastructure — it does not directly satisfy any R1-R19
rule on its own. It enables:
- **R5/R18** (password validation — breachcheck client)
- **R6/R19** (fail-open — breachcheck client behavior)
- **R7** (constant-time — bcrypt default cost in `secrets`)
- **R2/R3/R4/R13** (nudge emails — `Sender.SendNudgeEmail`)
- **R1/R2** (verification emails — `Sender.SendVerificationEmail`)

## 6. Testing checklist (this task's slice)

- [ ] `secrets`: `HashPassword` returns a bcrypt hash; `ComparePassword`
      matches the original password and rejects a wrong one.
- [ ] `breachcheck`: `IsBreached` returns `(false, nil)` on a password
      not in the breach list (mock HTTP server returning empty body).
- [ ] `breachcheck`: `IsBreached` returns `(true, nil)` when the suffix
      is present in the response (mock HTTP server).
- [ ] `breachcheck`: `IsBreached` returns `(false, nil)` on API
      unreachable (mock HTTP server closed / returning 5xx) — fail-open,
      and the log line contains no password/hash.
- [ ] `notification`: `FakeSender.SendVerificationEmail` and
      `SendNudgeEmail` return nil; log output contains no recipient
      address or token.
- [ ] `go test -race ./internal/platform/...` clean.

## 7. Common mistakes to avoid (techplan §13 slice)

| Mistake | Fix |
|---|---|
| `http.DefaultClient` (no timeout) for HaveIBeenPwned | Construct `http.Client{Timeout: ...}` once in `NewClient`, reuse. |
| Logging email in plaintext | `FakeSender` logs "recipient redacted", never the address. |
| Logging breach-check error verbatim (may embed response body) | Log sanitized summary + status code, not the raw error string. |
| Not wrapping errors with `%w` | Always `fmt.Errorf("...: %w", err)`. |

## 8. Risk note

- Assumptions made: HaveIBeenPwned API shape (`/range/{5-char-prefix}`,
  `SUFFIX:COUNT` lines) is stable; `bcrypt.DefaultCost` (~10) is the
  intended cost for the constant-time approach in Task 04.
- Edge cases intentionally NOT handled: breach-list API rate-limiting
  (429) is treated as unreachable → fail-open. Adding backoff would
  add latency the fail-open policy is designed to avoid.
- Concurrency assumptions: `breachcheck.Client` and `FakeSender` are
  safe for concurrent use (`http.Client` is goroutine-safe; `FakeSender`
  is stateless).
- What is not tested, and why: no live network call in tests — all
  HTTP interactions are mocked. A live smoke test against
  `api.pwnedpasswords.com` would be flaky and is out of scope for unit
  tests.
