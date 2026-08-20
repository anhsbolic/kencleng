# Task 05 — Transport Layer + Wiring

> Ticket    : 01-register-email-verification
> Sub-task  : 5 of 5
> Axis      : Dependency/sequence chain (primary) + layer (vertical slice)
> Status    : Blocked on Task 04 (and transitively Task 01 for `golang.org/x/time/rate`)
> Back-ref  : `../2-plan/techplan.md` (originating contract techplan — cross-check high-level decisions there whenever needed)

---

## 1. Scope

The interface/API layer + process wiring: RFC 9457 Problem Details
error mapping, per-endpoint rate-limit middleware with idle-key
eviction, the three HTTP handlers (`POST /auth/register`,
`POST /auth/verify-email`, `POST /auth/verify-email/resend`), and the
`cmd/server/main.go` wiring that assembles `account.Service` with its
dependencies and registers the routes behind the rate limiter.

**In scope:**
- `transport/http/errors.go` — sentinel errors + `WriteProblem` (RFC 9457)
- `transport/http/middleware.go` — `RateLimit(rps, burst)` with
  idle-key eviction, 429 Problem Details
- `transport/http/auth_register.go` — register handler
- `transport/http/auth_verify_email.go` — verify-email + resend handlers
- `cmd/server/main.go` (edit) — wire `account.Service` + deps, register routes

**Out of scope:**
- Any change to `domain/account/*` (Tasks 03/04)
- Any change to `platform/*` (Tasks 02 + Tier 0 crypto prerequisite)
- Rate-limit RPS/burst default values — techplan §14 Open Item #3 is
  unresolved; the middleware is configurable and the wiring takes
  values from env/config (defaults TBD by the human, see §4 below).

## 2. Dependencies

- **Hard deps:**
  - Task 04 (`account.Service`, `account.ErrValidation`,
    `account.ErrTokenExpired`, `account.ErrTokenNotFound`)
  - Task 01 (`golang.org/x/time/rate` in `go.mod`)
- **Soft deps:** none
- **Blocks:** nothing (terminal task in the chain)

## 3. Files

| File | Change Type |
|---|---|
| `backend/internal/transport/http/errors.go` | New |
| `backend/internal/transport/http/middleware.go` | New |
| `backend/internal/transport/http/auth_register.go` | New |
| `backend/internal/transport/http/auth_verify_email.go` | New |
| `backend/cmd/server/main.go` | Edit |

## 4. Implementation detail

### `backend/internal/transport/http/errors.go` (new)

Central RFC 9457 Problem Details writer. No internal leakage
(AGENTS.md golden rule, techplan §7 risk row 8, §13 row 7).

```go
package http

import (
    "encoding/json"
    "errors"
    "log"
    "net/http"

    "kencleng/internal/domain/account"
)

// Problem Details (RFC 9457) — application/problem+json.
// NEVER include stack traces, raw SQL errors, or file paths in
// detail strings. The detail is a stable, human-safe sentence.

type problem struct {
    Type     string       `json:"type"`
    Title    string       `json:"title"`
    Status   int          `json:"status"`
    Detail   string       `json:"detail,omitempty"`
    Instance string       `json:"instance,omitempty"`
    Errors   []fieldError `json:"errors,omitempty"` // for 422 ValidationProblem
}

type fieldError struct {
    Field  string `json:"field"`
    Detail string `json:"detail"`
}

// WriteProblem writes a generic RFC 9457 problem response.
func WriteProblem(w http.ResponseWriter, status int, problemType, title, detail string) {
    w.Header().Set("Content-Type", "application/problem+json")
    w.WriteHeader(status)
    _ = json.NewEncoder(w).Encode(problem{
        Type:   problemType,
        Title:  title,
        Status: status,
        Detail: detail,
    })
}

// WriteValidationError writes a 422 ValidationProblem with field-level errors.
func WriteValidationError(w http.ResponseWriter, errs []fieldError) {
    w.Header().Set("Content-Type", "application/problem+json")
    w.WriteHeader(http.StatusUnprocessableEntity)
    _ = json.NewEncoder(w).Encode(problem{
        Type:   "https://kencleng.dev/problems/validation",
        Title:  "Validation failed",
        Status: http.StatusUnprocessableEntity,
        Errors: errs,
    })
}

// MapServiceError maps an account.Service sentinel error to the
// appropriate HTTP status + Problem Details. Unknown errors map to
// 500 with a generic detail (never the raw error string).
func MapServiceError(w http.ResponseWriter, err error) {
    switch {
    case errors.Is(err, account.ErrValidation):
        // Validation is usually handled by the handler with field
        // detail; this is a fallback.
        WriteProblem(w, http.StatusUnprocessableEntity,
            "https://kencleng.dev/problems/validation",
            "Validation failed", "The request was invalid.")
    case errors.Is(err, account.ErrTokenExpired):
        WriteProblem(w, http.StatusGone,
            "https://kencleng.dev/problems/token-expired",
            "Token expired", "The verification token has expired.")
    case errors.Is(err, account.ErrTokenNotFound):
        WriteProblem(w, http.StatusNotFound,
            "https://kencleng.dev/problems/token-not-found",
            "Token not found", "The verification token was not found.")
    default:
        // Do NOT leak err.Error() — log it server-side, return generic.
        log.Printf("transport: unhandled service error: %v", err)
        WriteProblem(w, http.StatusInternalServerError,
            "https://kencleng.dev/problems/internal",
            "Internal error", "An unexpected error occurred.")
    }
}
```

Notes:
- `Content-Type: application/problem+json` is mandatory per RFC 9457.
- The 500 path logs the real error server-side (for ops) but returns a
  generic detail to the client (no stack trace, no SQL, no path). Per
  AGENTS.md golden rule and techplan §13 row 7.
- Validation ideally carries field-level detail (`WriteValidationError`)
  — the handler uses this for R5/R18 so the client sees which field
  failed. The field name `password` is not sensitive; the password
  value is never echoed.

### `backend/internal/transport/http/middleware.go` (new)

Per-IP rate limiter for public `/auth/*` endpoints. Uses
`golang.org/x/time/rate`. Limiter map with idle-key eviction —
unbounded map is a memory leak (techplan §7 risk row 3, §13 row 3).

```go
package http

import (
    "net"
    "net/http"
    "sync"
    "time"

    "golang.org/x/time/rate"
)

// RateLimit returns middleware that applies a per-IP token bucket of
// rps requests/sec with burst capacity burst. Idle limiters are
// evicted by a background sweeper to bound memory.
func RateLimit(rps float64, burst int) func(http.Handler) http.Handler {
    var (
        mu       sync.Mutex
        limiters = make(map[string]*rate.Limiter)
        lastSeen = make(map[string]time.Time)
    )

    // Background eviction: drop entries not seen in the last TTL.
    const ttl = 10 * time.Minute
    go func() {
        t := time.NewTicker(time.Minute)
        defer t.Stop()
        for range t.C {
            mu.Lock()
            now := time.Now()
            for ip, seen := range lastSeen {
                if now.Sub(seen) > ttl {
                    delete(limiters, ip)
                    delete(lastSeen, ip)
                }
            }
            mu.Unlock()
        }
    }()

    get := func(ip string) *rate.Limiter {
        mu.Lock()
        defer mu.Unlock()
        l, ok := limiters[ip]
        if !ok {
            l = rate.NewLimiter(rate.Limit(rps), burst)
            limiters[ip] = l
        }
        lastSeen[ip] = time.Now()
        return l
    }

    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ip, _, err := net.SplitHostPort(r.RemoteAddr)
            if err != nil {
                ip = r.RemoteAddr // fallback for non-host:port forms
            }
            if !get(ip).Allow() {
                WriteProblem(w, http.StatusTooManyRequests,
                    "https://kencleng.dev/problems/rate-limited",
                    "Rate limited", "Too many requests. Try again later.")
                return
            }
            next.ServeHTTP(w, r)
        })
    }
}
```

Notes:
- Per-IP keying is appropriate for public anonymous endpoints
  (register/verify/resend) — there is no authenticated identity to key
  on.
- `rps` and `burst` are configurable — the wiring (§4 below) takes
  them from env/config. **Defaults are TBD** (techplan §14 Open Item
  #3). The middleware itself is correct for any positive values; the
  human resolves the open item before this task ships.
- Idle-key eviction: background sweeper drops entries not seen in `ttl`.
  Bounds memory under sustained traffic from rotating clients.
- `429` is a Problem Details response, not a bare status.

### `backend/internal/transport/http/auth_register.go` (new)

```go
package http

import (
    "encoding/json"
    "errors"
    "net/http"

    "kencleng/internal/domain/account"
)

type registerRequest struct {
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

// RegisterHandler handles POST /auth/register.
// On success it writes 202 with a generic accepted message —
// identical for all four internal branches (anti-enumeration, R7).
// The service returns nil on every branch; the handler does not
// know which branch ran.
func RegisterHandler(svc *account.Service) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req registerRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            WriteProblem(w, http.StatusBadRequest,
                "https://kencleng.dev/problems/invalid-json",
                "Invalid request", "The request body was not valid JSON.")
            return
        }
        // Boundary validation — reject malformed input before reaching
        // the service. Field names in errors are not sensitive; values
        // are never echoed.
        var fieldErrs []fieldError
        if len(req.Name) < 1 || len(req.Name) > 255 {
            fieldErrs = append(fieldErrs, fieldError{Field: "name", Detail: "must be 1-255 characters"})
        }
        if !looksLikeEmail(req.Email) {
            fieldErrs = append(fieldErrs, fieldError{Field: "email", Detail: "must be a valid email"})
        }
        if len(req.Password) < 8 { // length check also done in service (R5) — defense in depth
            fieldErrs = append(fieldErrs, fieldError{Field: "password", Detail: "must be at least 8 characters"})
        }
        if len(fieldErrs) > 0 {
            WriteValidationError(w, fieldErrs)
            return
        }

        if err := svc.Register(r.Context(), req.Name, req.Email, req.Password); err != nil {
            // Only ErrValidation is expected here; everything else
            // would be a service bug (the four register branches
            // return nil). Map defensively.
            if errors.Is(err, account.ErrValidation) {
                // Service-level validation (e.g. breach-list hit) —
                // surface as field error on password.
                WriteValidationError(w, []fieldError{{Field: "password", Detail: "password is not allowed"}})
                return
            }
            MapServiceError(w, err)
            return
        }
        // Anti-enumeration: identical 202 generic for all branches.
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusAccepted)
        _ = json.NewEncoder(w).Encode(map[string]string{"message": "If your email is not already registered, you will receive a verification link."})
    }
}
```

Notes:
- The 202 message is **generic and identical** regardless of branch —
  no `user_id`, no hint about email state (R1-R4, R7, R17).
- Boundary validation (name length, email shape, password length) is
  defense-in-depth — the service re-checks password policy (R5/R18)
  because the breach check can only run there.
- `looksLikeEmail` is a minimal shape check; full RFC 5322 is out of
  scope (the service's `crypto.HMAC` + DB lookup is the real authority
  on "is this email known").
- No PII in logs — the handler logs at most "register request
  received" / "register completed", never the email.

### `backend/internal/transport/http/auth_verify_email.go` (new)

```go
package http

import (
    "encoding/json"
    "net/http"

    "kencleng/internal/domain/account"
)

type verifyEmailRequest struct {
    Token string `json:"token"`
}

type resendRequest struct {
    Email string `json:"email"`
}

// VerifyEmailHandler handles POST /auth/verify-email.
// R8: valid → 200. R9: expired → 410. R10/R11: not-found/used/revoked → 404.
func VerifyEmailHandler(svc *account.Service) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req verifyEmailRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            WriteProblem(w, http.StatusBadRequest,
                "https://kencleng.dev/problems/invalid-json",
                "Invalid request", "The request body was not valid JSON.")
            return
        }
        // Reject empty token at the boundary — saves a DB round-trip
        // (techplan §13 row 9, §7 risk row "Empty token string").
        if req.Token == "" {
            WriteProblem(w, http.StatusNotFound,
                "https://kencleng.dev/problems/token-not-found",
                "Token not found", "The verification token was not found.")
            return
        }
        if err := svc.VerifyEmail(r.Context(), req.Token); err != nil {
            MapServiceError(w, err) // ErrTokenExpired→410, ErrTokenNotFound→404
            return
        }
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        _ = json.NewEncoder(w).Encode(map[string]string{"message": "Email verified."})
    }
}

// ResendVerificationHandler handles POST /auth/verify-email/resend.
// R13/R14: always 202 generic — identical whether or not a token was issued.
func ResendVerificationHandler(svc *account.Service) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req resendRequest
        if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
            WriteProblem(w, http.StatusBadRequest,
                "https://kencleng.dev/problems/invalid-json",
                "Invalid request", "The request body was not valid JSON.")
            return
        }
        if !looksLikeEmail(req.Email) {
            WriteValidationError(w, []fieldError{{Field: "email", Detail: "must be a valid email"}})
            return
        }
        _ = svc.ResendVerification(r.Context(), req.Email)
        // Always 202 generic — the service returns nil for both
        // match and no-match branches (R13/R14).
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusAccepted)
        _ = json.NewEncoder(w).Encode(map[string]string{"message": "If your email is registered and unverified, you will receive a new verification link."})
    }
}
```

Notes:
- The resend handler **ignores the service error** (it only calls
  `ResendVerification`, which always returns nil per Task 04) — the
  `_ =` is intentional: R13 and R14 produce identical 202 responses,
  so there is no error path to map. If `ResendVerification` ever
  returns a real error, that's a bug — log it server-side and still
  return 202 (better to lose an error than leak email state).
- Empty-token boundary check returns 404 directly (matches R10's
  outcome) — avoids a DB round-trip and avoids any timing distinction
  between "empty" and "not found".
- The 202 messages for register and resend are deliberately similar
  in shape and tone — anti-enumeration at the response-text level too.

### `backend/cmd/server/main.go` (edit)

Wire `account.Service` with its dependencies and register the three
routes behind the rate-limit middleware.

```go
// ... existing setup (db pool, crypto keys, etc.) ...

// Account domain wiring.
accountRepo := account.NewRepositoryDB(dbPool, cryptoKeys)
breachClient := breachcheck.NewClient(5 * time.Second) // explicit timeout (techplan §7 row 4)
emailSender := notification.NewFakeSender()            // v1: logged, no SMTP
accountSvc := account.NewService(accountRepo, dbPool, nil /* hasher pkg */, breachClient, emailSender, cryptoKeys)

// Routes — standard net/http Go 1.22+ pattern routing.
authMux := http.NewServeMux()
authMux.HandleFunc("POST /auth/register", http.RegisterHandler(accountSvc))
authMux.HandleFunc("POST /auth/verify-email", http.VerifyEmailHandler(accountSvc))
authMux.HandleFunc("POST /auth/verify-email/resend", http.ResendVerificationHandler(accountSvc))

// Rate-limited mount — RPS/burst from env (Open Item #3 defaults TBD).
rps := envFloat("AUTH_RATE_RPS", /* TBD */)
burst := envInt("AUTH_RATE_BURST", /* TBD */)
mux.Handle("/auth/", http.RateLimit(rps, burst)(authMux))
```

Notes:
- `breachcheck.NewClient(5 * time.Second)` — explicit timeout, never
  `http.DefaultClient` (techplan §7 risk row 4).
- The rate-limit values come from env so ops can tune without code
  changes; the **defaults are TBD** pending techplan §14 Open Item #3.
  Do not pick arbitrary numbers — flag for the human. The wiring
  compiles with placeholder defaults (`0` would disable the limiter,
  which is wrong — use a sentinel that fails fast if env is unset).
- Standard `net/http` Go 1.22+ pattern routing per `backend/AGENTS.md`
  §2 — no third-party router.

## 5. Rules covered (this task's slice)

| Rule | How this task satisfies it | Test |
|---|---|---|
| R7 (response shape) | Handler writes identical 202 generic for all branches; service returns nil on all four | (Service-level `TestRegister_GenericResponse_AllBranches` in Task 04 proves branch equivalence; handler-level test asserts the 202 body is constant) |
| R8/R9/R10/R11 | `MapServiceError` maps `ErrTokenExpired`→410, `ErrTokenNotFound`→404; handler returns 200 on nil | Handler tests: valid→200, expired→410, not-found/used/revoked→404 |
| R13/R14 | Resend handler always writes 202 generic | Handler test: match and no-match both produce identical 202 |
| R15 | `RateLimit` middleware returns 429 on excess | `TestResend_RateLimited` (handler+middleware integration) |
| Error responses | `WriteProblem` / `MapServiceError` emit RFC 9457, no internal leakage | Handler test: 500 path returns generic detail, not raw error |
| Boundary validation | Empty token rejected at handler (saves DB round-trip); email shape + name length checked | Handler tests for each malformed input |

## 6. Testing checklist (this task's slice)

- [ ] `errors.go`: `WriteProblem` sets `Content-Type: application/problem+json`
      and encodes the problem struct correctly for each status.
- [ ] `errors.go`: `MapServiceError` maps each sentinel to the correct
      status (422/410/404/500); 500 path returns generic detail, not
      `err.Error()`.
- [ ] `middleware.go`: requests under the rate limit pass through;
      the (burst+1)th request returns 429 with Problem Details.
- [ ] `middleware.go`: idle limiters are evicted after the TTL (test
      with a short TTL via internal injection or a fake clock).
- [ ] `auth_register.go`: valid request → 202 with the generic message
      and no `user_id`; malformed JSON → 400; invalid name/email/password
      → 422 with field-level errors; service `ErrValidation` → 422.
- [ ] `auth_verify_email.go`: empty token → 404 (no service call);
      valid → 200; expired → 410; not-found/used/revoked → 404.
- [ ] `auth_verify_email.go` (resend): valid email shape → 202 generic
      regardless of branch; invalid email → 422; rate-limited → 429
      (`TestResend_RateLimited`).
- [ ] `cmd/server/main.go`: server starts; the three routes are
      reachable behind the rate limiter; breach client has a non-zero
      timeout; `golang.org/x/time/rate` is wired.
- [ ] `go test -race ./internal/transport/http/...` clean.
- [ ] `make verify` passes end-to-end (this task is the terminal one,
      so the full gate runs here).

## 7. Common mistakes to avoid (techplan §13 slice)

| Mistake | Fix |
|---|---|
| Raw error string in 4xx/5xx response | `MapServiceError` returns generic detail; log raw server-side. |
| Empty token accepted by verify-email handler | Reject `""` at the boundary, return 404 directly. |
| Rate limiter map without eviction | Background sweeper drops idle entries past TTL. |
| `http.DefaultClient` for breachcheck | `breachcheck.NewClient(timeout)` constructs an `http.Client` with explicit `Timeout` (Task 02); wiring passes `5 * time.Second`. |
| Not wrapping errors with `%w` | `fmt.Errorf("...: %w", err)` everywhere in this package too. |
| Third-party router | Standard `net/http` Go 1.22+ pattern routing per `backend/AGENTS.md` §2. |

## 8. Risk note

- Assumptions made: `account.Service`'s constructor signature matches
  Task 04's `NewService`; the `secrets` package is consumed as
  package-level funcs (adjust the `nil /* hasher pkg */` placeholder
  to match Task 02's actual API — `*secrets.Hasher` vs package funcs).
  Rate-limit RPS/burst defaults are deferred to the human (Open Item
  #3); the wiring reads from env and fails fast if unset rather than
  silently disabling the limiter.
- Edge cases intentionally NOT handled: X-Forwarded-For / trusted-proxy
  IP extraction — `r.RemoteAddr` is used directly. Behind a reverse
  proxy this would key all traffic to the proxy's IP, collapsing the
  per-client limit. This is a known gap for the v1 deployment shape
  (no proxy in front yet); flag it for the human if a proxy is added.
- Concurrency assumptions: the rate limiter map is guarded by a mutex;
  the background evicter goroutine runs for the process lifetime.
  Handler closures capture only the `*account.Service` (safe for
  concurrent use per Task 04).
- What is not tested, and why: the handler tests do not exercise the
  real service — they use a stub/fake service to assert status-code
  and Problem-Details mapping. End-to-end (handler → service → repo →
  Postgres) is covered by the integration build under
  `//go:build integration` in Task 03/04. R15's rate-limit test uses
  a tight RPS (e.g. rps=1, burst=1) and fires burst+2 concurrent
  requests to deterministically trigger 429.
