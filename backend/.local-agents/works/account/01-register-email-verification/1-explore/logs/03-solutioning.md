# Stage 3 — Solutioning

> Feature: Register & Email Verification (`01-register-email-verification.md`)
> Domain: account
> Date: 2026-08-19
> All decisions confirmed by user.

---

## Decision 1: PII Encryption — Tier 0 Blocker

**Problem**: `platform/crypto/` is Tier 0 fenced (AGENTS.md §3). Agent
cannot add encrypt/decrypt/HMAC functions. But register must encrypt
`primary_email`/`identifier` and compute HMAC hashes.

**Chosen: Option A — Human provides crypto functions before this task
starts.**

A human writes `Encrypt(plaintext []byte, key []byte) (ciphertext []byte,
err error)`, `Decrypt(ciphertext, key)`, `HMAC(data, key) (hash string)`
in `platform/crypto/`. The agent consumes them. The crypto package's
stated purpose is exactly this — it already holds the key material.

This is a prerequisite, not part of this feature's scope. Without these
functions, the repository layer cannot insert encrypted PII.

---

## Decision 2: Password Hashing — Package Placement

**Chosen: Option B — New `platform/secrets/` package wrapping bcrypt.**

Agent creates `platform/secrets/` with:
- `HashPassword(password string) (string, error)` — bcrypt
  GenerateFromPassword, default cost
- `ComparePassword(hash, password string) error` — bcrypt
  CompareHashAndPassword

Password hashing is used in register, reset-password, and set-password
— three features across the domain. A thin wrapper avoids repeating the
bcrypt import pattern. The package's doc.go clarifies it's for
credential hashing, distinct from `crypto/`'s PII encryption.

Import: `golang.org/x/crypto/bcrypt` (already an indirect dep via
minio — needs explicit import, not just go.mod presence).

---

## Decision 3: Email Sender Interface

**Chosen: Option A — `platform/notification/sender.go` interface +
`FakeSender` stub.**

```go
// platform/notification/sender.go

type Sender interface {
    SendVerificationEmail(ctx context.Context, to, token string) error
    SendNudgeEmail(ctx context.Context, to, nudgeType string) error
}
```

`FakeSender` logs the email content (recipient, subject, body) via
`log.Printf`. Placed in `platform/` since email sending is shared infra
(notification domain will own real SMTP later).

The interface methods match the feature spec's email types:
- Verification email (new registration)
- Resend-verification nudge (already registered, unverified)
- Password-reset nudge (already registered, verified)
- Google-only nudge (email claimed by google identity)

---

## Decision 4: Domain Package Structure & Breachcheck

**Chosen: `platform/breachcheck/` as a separate package.**

```
internal/domain/account/
├── entity.go          # User, AuthIdentity, AuthToken structs
├── repository.go      # Repository interface (ports)
├── repository_db.go   # goqu-based implementation (adapter)
├── service.go         # Business logic
└── service_test.go    # Table-driven tests

internal/platform/breachcheck/
├── client.go          # HaveIBeenPwned k-anonymity client
└── client_test.go     # Tests with mocked HTTP
```

`repository.go` defines the interface (testable, mockable).
`repository_db.go` implements it with goqu + pgx. The service depends
on the interface, not the concrete implementation.

`breachcheck` is a separate package because it's an external HTTP call
with its own timeout/fail-open logic, reusable by reset-password later.

---

## Decision 5: Token Generation

**Chosen: Option A — Generate in the service using `crypto/rand`.**

32 random bytes, hex-encoded for the user-facing token, SHA-256 for
storage. Self-contained, ~10 lines. If password-reset reuses it,
extract later (YAGNI).

---

## Decision 6: Rate Limiting

**Chosen: Per-endpoint rate limiter middleware in
`transport/http/middleware/`.**

```go
func RateLimit(rps float64, burst int) func(http.Handler) http.Handler
```

Uses `rate.NewLimiter`. Applied to the three `/auth/*` routes. The
`ratelimit` platform package stays empty — middleware is transport-layer
concern.

---

## Decision 7: Migration Strategy

**Chosen: Three migration files (up/down), numbered sequentially.**

```
migrations/
├── 000001_create_users.up.sql
├── 000001_create_users.down.sql
├── 000002_create_auth_identities.up.sql
├── 000002_create_auth_identities.down.sql
├── 000003_create_auth_tokens.up.sql
└── 000003_create_auth_tokens.down.sql
```

Migration 000001 also creates the `set_updated_at()` trigger function
(used by both `users` and `auth_identities`). SQL comes directly from
the ERD. Unique indexes and partial indexes are created in the same
migration as their table.

---

## Decision 8: Anti-Enumeration — Constant-Time Handling

**Chosen: "Always bcrypt" approach — dummy hash on no-op branches.**

The service always performs one `bcrypt.GenerateFromPassword` call
regardless of which branch executes. On branches that don't need a new
password hash (verified user, Google-only), the bcrypt call is a
"dummy" — its result is discarded. This makes all four branches take
~100ms (bcrypt's default cost), eliminating the timing side-channel.

Noted in the risk note per AGENTS.md §5.

---

## Decision 9: Handler Wiring & Error Mapping

One handler file per feature in `transport/http/`:

```
internal/transport/http/
├── auth_register.go        # POST /auth/register
├── auth_verify_email.go    # POST /auth/verify-email + resend
├── middleware.go            # Rate limit, content-type
└── errors.go               # Map domain errors → Problem Details
```

`errors.go` defines sentinel errors (`ErrTokenExpired`,
`ErrTokenNotFound`, `ErrValidation`, etc.) and a central
`WriteProblem(w, err)` function mapping them to the correct HTTP status
+ RFC 9457 body.

In `cmd/server/main.go`:
```go
mux.HandleFunc("POST /auth/register", authHandler.Register)
mux.HandleFunc("POST /auth/verify-email", authHandler.VerifyEmail)
mux.HandleFunc("POST /auth/verify-email/resend", authHandler.ResendVerification)
```

---

## Decision 10: Dependency Order (Implementation Sequence)

```
1.  go get goqu golang.org/x/time/rate
2.  migrations/ (3 up+down files + trigger function)
3.  platform/secrets/ (bcrypt wrapper)
4.  platform/breachcheck/ (HaveIBeenPwned k-anonymity client)
5.  platform/notification/ (Sender interface + FakeSender)
6.  domain/account/entity.go (structs)
7.  domain/account/repository.go (interface)
8.  domain/account/repository_db.go (goqu implementation)
9.  domain/account/service.go (Register, VerifyEmail, ResendVerification)
10. domain/account/service_test.go (table-driven + concurrency tests)
11. transport/http/errors.go (error → Problem mapping)
12. transport/http/middleware.go (rate limit)
13. transport/http/auth_register.go (handler)
14. transport/http/auth_verify_email.go (handler)
15. cmd/server/main.go wiring
```

Steps 2–5 can be parallelized (no interdependencies). Steps 6–9 are
sequential. Steps 10–14 depend on 9. Step 15 last.

**Prerequisite**: human-authored `platform/crypto/` functions (Decision
1). Without them, step 8 cannot encrypt PII.

---

## Assumptions carried forward

- The 9 test names prescribed in the feature spec are the canonical
  names — implementation must use them exactly.
- `golang-migrate` CLI is used for running migrations (per tech stack
  doc) — no auto-run on app start.
- Email sending is fake/logged in v1 — the `FakeSender` implementation
  is sufficient for all tests and local dev.
- `bcrypt` default cost (10) is acceptable for a sandbox project.
- The `set_updated_at()` trigger function is created once in migration
  000001 and reused by subsequent tables.
