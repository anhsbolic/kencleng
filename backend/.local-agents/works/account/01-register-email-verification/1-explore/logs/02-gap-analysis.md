# Stage 2 — Gap Analysis

> Feature: Register & Email Verification (`01-register-email-verification.md`)
> Domain: account
> Date: 2026-08-19

---

## Area 1: API Contract (`api/openapi.yaml`)

### Current state

The `api/openapi.yaml` defines three endpoints and their schemas:

- **`POST /auth/register`** — `RegisterRequest` (name: string 1-255,
  email: email format, password: password format, minLength 8).
  Response: always `202` + `GenericAcceptedMessage` (message: string).
  Also `422` (ValidationProblem) and `429` (Problem). `security: []`
  (public/unauthenticated).

- **`POST /auth/verify-email`** — `VerifyEmailRequest` (token: string).
  Response: `200` (message: string), `410` (token expired), `404` (not
  found/already used/revoked), `429`. `security: []`.

- **`POST /auth/verify-email/resend`** — `ResendVerificationRequest`
  (email: email format). Response: always `202` + `GenericAcceptedMessage`,
  `429`. `security: []`.

- **Error format**: RFC 9457 Problem Details. `ValidationProblem` extends
  `Problem` with `errors[]` (field + message).

### Requirement

The feature spec requires the handlers to match this contract exactly.
Key behavioral requirements embedded in the API description:

- Anti-enumeration: uniform `202` for register/resend regardless of
  email state
- Password validation (`422` before any enumeration-sensitive branching)
- Specific error codes for verify-email (`410` expired, `404` not
  found/used/revoked)

### Gap

No handler code exists. The `internal/transport/http/` directory contains
only `doc.go`. The entire HTTP transport layer — routing, request
parsing, response writing, error mapping — needs to be built from
scratch. No DTOs, no handler functions, no error-to-Problem mapping.

### Sniffing

- **Risk**: The `RegisterRequest` has `minLength: 8` on password, but
  the feature spec also requires a HaveIBeenPwned breach-list check.
  The API contract doesn't mention the breach check in the `422`
  description — this is fine (implementation detail), but the
  `ValidationProblem.errors[]` must include a breach-list failure as a
  field-level error. The handler must decide which field name to use
  (likely `password`).

- **Edge cases**: The `VerifyEmailRequest.token` is just `type: string`
  with no `minLength`/`format` constraint — empty string is technically
  valid per the schema. The handler must reject empty tokens (or the
  service layer will).

- **Miscontext**: None — the API spec and feature spec are aligned
  (the feature spec explicitly documents the `openapi.yaml` changes in
  Assumption A).

- **Misleading signals**: The `GenericAcceptedMessage` schema doesn't
  constrain `properties` to only `message` (no `additionalProperties:
  false`). Harmless but worth noting: the handler must not accidentally
  add extra fields.

- **Inconsistency**: The `POST /auth/verify-email` returns `200` on
  success, while register and resend return `202`. This is intentional
  (verify-email is a synchronous state change, the others trigger async
  side effects). No inconsistency — deliberate design choice.

---

## Area 2: Domain Data Model (`docs/project/kencleng-erd.md` §1)

### Current state

Three tables are relevant, all defined in the ERD:

**`users`**
- `id` (UUID PK), `name` (TEXT NOT NULL)
- `primary_email` (BYTEA, AES-GCM ciphertext)
- `primary_email_hash` (TEXT, HMAC-SHA256, unique index)
- `created_at`, `updated_at`
- No soft-delete, no role column (roles live in `user_roles`)

**`auth_identities`**
- `id` (UUID PK), `user_id` (UUID FK→users CASCADE)
- `provider_type` (TEXT, CHECK: `email_password`/`google`/`phone_otp`)
- `identifier` (BYTEA, encrypted), `identifier_hash` (TEXT, HMAC)
- `credential_secret` (TEXT, nullable — password hash for
  `email_password`, NULL for google)
- `verified_at` (TIMESTAMPTZ, nullable), `created_at`, `updated_at`
- Unique index on `(provider_type, identifier_hash)`

**`auth_tokens`**
- `id` (UUID PK), `user_id` (UUID FK→users CASCADE)
- `purpose` (TEXT, CHECK: `email_verification`/`password_reset`)
- `token_hash` (TEXT, unique), `expires_at`, `used_at`, `revoked_at`,
  `created_at`
- Partial index `ix_auth_tokens_valid` on `(user_id, purpose) WHERE
  used_at IS NULL AND revoked_at IS NULL`

PII pattern: `{field}` (BYTEA, AES-GCM) + `{field}_hash` (TEXT,
HMAC-SHA256). Keys in `internal/platform/crypto/keys.go`.

### Requirement

The register endpoint must insert into all three tables atomically
(user + auth_identity + auth_token in one transaction). Verify-email
must update `auth_tokens.used_at` and `auth_identities.verified_at`.
Resend must revoke old tokens and insert a new one. All queries must
use `goqu` (per AGENTS.md).

### Gap

Everything is missing:

- No entity structs (Go types for User, AuthIdentity, AuthToken)
- No repository interface or implementation (InsertUser,
  InsertAuthIdentity, InsertAuthToken, UpdateVerifiedAt,
  RevokeOldTokens, FindByIdentifierHash, FindValidToken)
- No migration SQL files (`migrations/` dir has only `.gitkeep`)
- No service layer (register orchestration, verify-email token
  redemption, resend with old-token revocation)
- No transaction management (register must be atomic across 3 tables)

The platform layer (`crypto/keys.go`, `db/db.go`) exists and is usable.

### Sniffing

- **Risk**: The `identifier_hash` HMAC is the lookup key for
  duplicate-email detection. If the HMAC key rotates, existing hashes
  become orphaned — no key-rotation strategy documented. Not a blocker
  for this feature but a latent risk.

- **Edge cases**: Concurrent duplicate registration for the same email
  — INV-account-01 requires the DB unique index `(provider_type,
  identifier_hash)` to catch this, and the service must handle the
  unique-violation error cleanly (not crash, not leak internals). The
  race between two simultaneous `INSERT INTO users` + `INSERT INTO
  auth_identities` needs to be handled — either the second
  `auth_identities` INSERT fails on the unique index, or the first
  `users` INSERT needs a rollback. The implementation must use a single
  transaction wrapping both inserts.

- **Miscontext**: None — the ERD and feature spec are aligned.

- **Misleading signals**: The `auth_tokens` table has a `purpose` CHECK
  constraint that includes `password_reset` — someone might think this
  feature needs password-reset logic. It doesn't; this feature only
  uses `purpose = 'email_verification'`.

- **Inconsistency**: The ERD says `users.primary_email` is BYTEA
  (encrypted), but the `RegisterRequest` API schema takes `email` as a
  plain string. The handler must encrypt before insert — this is the
  platform/crypto layer's job. No inconsistency, but a critical wiring
  step.

---

## Area 3: Domain Invariants & Threat Model

### Current state

Two invariants are directly relevant:

**INV-account-01** — Uniqueness is per `(provider_type,
identifier_hash)`. DB-level enforced via
`ux_auth_identities_provider_identifier` unique index. Requires a
concurrent-insert race test.

**INV-account-08** — Token redemption guard: `used_at IS NULL AND
revoked_at IS NULL AND expires_at > now()`. The `UPDATE ... WHERE` must
be atomic. Requires a double-submit concurrency test.

**Threat model §1** (Registration & Email Verification) identifies:
- Email enumeration → mitigated by uniform `202` response + constant-
  time handling (Build-stage detail)
- Concurrent duplicate registration → DB unique index (INV-account-01)
- Verification token replay → guarded UPDATE (INV-account-08)
- Verification-email flood → rate limit on `/auth/*`
- Weak/breached password → length policy + HaveIBeenPwned k-anonymity,
  fail-open

**Feature spec threat breakdown** prescribes test names:
- `TestRegister_GenericResponse_AllBranches`
- `TestRegister_GenericResponse_Timing`
- `TestRegister_ConcurrentDuplicateEmail_Race`
- `TestVerifyEmail_TokenSingleUse_Concurrent`
- `TestVerifyEmail_RevokedToken_Rejected`
- `TestResend_RateLimited`
- `TestRegister_PasswordPolicy`
- `TestRegister_BreachCheck_FailOpen`
- `TestRegister_GoogleOnlyConflict_GenericResponse`

### Requirement

The service/handler layer must enforce:

1. Password validation (≥8 chars + breach check) **before** any
   enumeration-sensitive branching
2. Uniform `202` response for all register branches
   (new/unverified/verified/google-only)
3. Atomic token redemption with the full predicate `used_at IS NULL AND
   revoked_at IS NULL AND expires_at > now()`
4. Old-token revocation on resend (`revoked_at` set on superseded
   tokens)
5. Rate limiting on all three `/auth/*` endpoints

### Gap

No enforcement code exists. The invariants are declared in spec, the DB
constraints are declared in the ERD, but:

- No migration to create the unique index or the partial index
- No service-layer code implementing the token redemption guard
- No rate-limiting middleware wired to these endpoints
- No constant-time handling for the anti-enumeration branches
- No test harness for any of the named tests

### Sniffing

- **Risk**: INV-account-08's statement says `used_at IS NULL AND
  revoked_at IS NULL AND expires_at > now()` but the test description
  only mentions `WHERE used_at IS NULL AND expires_at > now()` —
  missing `revoked_at IS NULL`. The **spec's own statement** is
  authoritative (includes `revoked_at`), but the test description in
  the invariant doc is slightly incomplete. The implementation must use
  the full three-clause predicate.

- **Edge cases**: Concurrent duplicate registration — the race is
  between two goroutines both doing INSERT INTO users + INSERT INTO
  auth_identities. The second INSERT INTO auth_identities will fail on
  the unique index, but the first goroutine's INSERT INTO users row is
  now orphaned if the transaction rolls back. The implementation must
  use a single transaction wrapping both inserts, so the rollback is
  clean.

- **Miscontext**: None — the invariants, threat model, and feature spec
  are consistent.

- **Misleading signals**: The `ix_auth_tokens_valid` partial index in
  the ERD looks like it could be used for the resend "does a valid
  token already exist?" check. But the resend flow doesn't need this
  check — it always revokes old tokens and issues new ones, regardless
  of whether a valid token exists. The index is useful for other
  purposes but not for the resend logic itself.

- **Inconsistency**: INV-account-08's test description omits
  `revoked_at IS NULL` from the WHERE clause, while the invariant
  statement includes it. The implementation must follow the statement,
  not the test description.

---

## Area 4: Existing Platform Scaffolding

### Current state

| Package | What exists | What's missing |
|---|---|---|
| `platform/db` | `db.go` — `Open(ctx, databaseURL) (*pgxpool.Pool, error)` | No query helpers, no transaction wrapper, no goqu integration |
| `platform/crypto` | `keys.go` — `New(encryptionKeyB64, hmacKeyB64) (*Keys, error)` — parses/validates 32-byte base64 keys | **No encrypt/decrypt/HMAC functions** — only the key holder. **Tier 0 fenced** (AGENTS.md §3) |
| `platform/auth` | `keys.go` — `Load(privatePath, publicPath) (*Keys, error)` — ES256 key pair | No JWT sign/verify, no session logic. **Tier 0 fenced** |
| `platform/ratelimit` | `doc.go` only — empty skeleton | No rate-limiting implementation |
| `platform/storage` | `doc.go` only | N/A for this feature |
| `platform/scheduler` | `doc.go` only | N/A for this feature |
| `transport/http` | `doc.go` only — empty skeleton | No handlers, no routing, no middleware |
| `cmd/server` | `main.go` — wired startup (env, db, minio, auth keys), single `GET /healthz` | No domain routes, no service wiring |

`internal/domain/` has only `.gitkeep`.

`go.mod` has: `pgx/v5`, `godotenv`, `minio-go`. Missing for this
feature: `goqu`, `golang.org/x/time/rate`. `golang.org/x/crypto` is
available as indirect dep.

### Requirement

This feature needs:

- Repository layer with goqu-based SQL (parameterized)
- Transaction management (register = atomic insert into 3 tables)
- PII encryption (AES-GCM) and HMAC for lookup columns
- Password hashing (bcrypt/argon2) — in `domain/account/` service or
  `platform/secrets/`, NOT in Tier 0 fenced `platform/crypto/` or
  `platform/auth/`
- Rate limiting middleware for `/auth/*`
- HTTP handlers with request parsing, validation, error-to-Problem
  mapping
- Email sending (fake/logged for v1)

### Gap

Massive — essentially everything is missing. The scaffold provides
connection pooling and key material, but no domain logic, no query
layer, no HTTP layer, no rate limiting, no email sending. The goqu
dependency isn't in `go.mod` yet.

**Critical Tier 0 concern**: `platform/crypto/` is fenced — the
encrypt/decrypt/HMAC functions cannot be added by an agent. The feature
spec says password hashing goes in `domain/account/` or
`platform/secrets/` (non-fenced). But PII encryption (AES-GCM for
`primary_email`, `identifier`) *does* need to live in `platform/crypto/`
— this is a **blocker** that requires human authorship.

### Sniffing

- **Risk**: The Tier 0 fencing on `platform/crypto/` means the PII
  encryption functions cannot be implemented by the agent. This is a
  hard blocker — the register endpoint cannot insert encrypted PII
  without these functions. Either: (a) human provides the
  encrypt/HMAC functions, (b) the feature is re-tiered, or (c) the
  fencing is scoped differently. The feature spec says password hashing
  goes in `platform/secrets/` — but PII encryption is the crypto
  package's stated purpose.

- **Edge cases**: The `goqu` dependency isn't in `go.mod` — adding it
  requires `go get`, straightforward but must happen before any
  repository code.

- **Miscontext**: None — the scaffold doc.go files accurately describe
  their state.

- **Misleading signals**: `platform/crypto/keys.go` has `Keys` struct
  with `EncryptionKey` and `HMACKey` — looks "ready to use" but there
  are no actual encrypt/decrypt/HMAC functions. The key material is
  loaded but unusable without the functions.

- **Inconsistency**: `golang.org/x/crypto` is an indirect dependency
  (pulled by minio), but bcrypt needs explicit import. Having the
  indirect dep doesn't mean bcrypt is wired.

---

## Area 5: Backend Tech Dependencies & Conventions

### Current state

Key tech choices from `docs/project/kencleng-backend-tech-stack.md`:

| Dependency | Status | Notes |
|---|---|---|
| `goqu` | **Not in `go.mod`** | Required by AGENTS.md for all SQL |
| `golang-migrate` | Mentioned in tech stack | CLI-run, no migration files exist |
| `golang.org/x/crypto` | Indirect dep | Available for bcrypt |
| HaveIBeenPwned | Spec requires it | k-anonymity, SHA-1 prefix (5 chars), fail-open. No client code |
| Email sending | **Fake/logged for v1** | `kencleng-phase3-detail.md`: "simply fake/logged" |
| `golang.org/x/time/rate` | Mentioned in tech stack | Not in `go.mod` |
| `net/http` stdlib | Go 1.22+ pattern routing | Hand-written handlers |
| `testify` | Not adopted | stdlib `go test` + `httptest` |

Conventions from `backend/AGENTS.md`:
- One package per domain: `account` needs `{entity.go, repository.go,
  service.go}`
- Repository uses `goqu`, never raw string concatenation
- Errors: wrap with `fmt.Errorf("...: %w", err)`
- Table-driven tests, no comments unless asked

### Requirement

This feature needs:

1. `goqu` for repository queries
2. Password hashing (bcrypt) — placed in `domain/account/service` or
   `platform/secrets/`, NOT in `platform/crypto/`
3. HaveIBeenPwned HTTP client (k-anonymity, SHA-1 prefix)
4. Email sending — fake/logged implementation (interface + stub)
5. Rate limiting middleware — `golang.org/x/time/rate`
6. Token generation (cryptographically random verification token)
7. Migration files for `users`, `auth_identities`, `auth_tokens`
8. HTTP handler wiring in `transport/http/`

### Gap

- `goqu` not in `go.mod` — needs `go get`
- No password hashing code — bcrypt available via indirect dep but no
  wrapper exists
- No HaveIBeenPwned client — needs to be written (HTTP call to
  `api.pwnedpasswords.com`, k-anonymity)
- No email interface — needs a `NotificationSender` or similar
  interface with a fake/logged implementation
- No rate-limiting middleware — `golang.org/x/time/rate` not in
  `go.mod`, no middleware wrapper
- No token generation — needs `crypto/rand` based token + SHA-256 hash
  for storage
- No migration files — `migrations/` is empty
- No handler code — `transport/http/` is empty
- No `main.go` wiring for new routes, services, or middleware

### Sniffing

- **Risk**: The feature spec says password hashing goes in
  `platform/secrets/` (non-fenced) — but this package doesn't exist
  yet. Creating it is fine, but it must not overlap with
  `platform/crypto/`'s purpose. Naming boundary: `crypto/` = PII
  encryption-at-rest (AES-GCM, HMAC), `secrets/` = credential hashing
  (bcrypt, argon2).

- **Edge cases**: HaveIBeenPwned fail-open — implementation must not
  block registration if API is unreachable. Need a timeout on the HTTP
  call and a clean fallback path. The k-anonymity model means only the
  first 5 hex chars of the SHA-1 hash are sent — implementation must
  hash the password, take the prefix, compare suffix locally.

- **Miscontext**: None — the tech stack doc and feature spec are
  aligned.

- **Misleading signals**: `golang.org/x/crypto` in `go.mod` as
  indirect dep — looks "already available" but pulled by minio, not
  intentionally added. bcrypt needs explicit import.

- **Inconsistency**: The tech stack says "no codegen" for OpenAPI —
  handlers are hand-written. But the feature spec defines 9+ test names
  that must exist. The test names are prescriptive — implementation must
  use exactly these names, not just cover the same behavior with
  different names.

---

## Summary of key blockers

1. **Tier 0 blocker**: `platform/crypto/` is fenced — PII encryption
   functions (AES-GCM encrypt/decrypt, HMAC) cannot be added by the
   agent. Hard dependency for inserting encrypted `primary_email` and
   `identifier`.

2. **Missing dependencies**: `goqu`, `golang.org/x/time/rate` not in
   `go.mod`.

3. **Everything is greenfield**: No domain package, no handlers, no
   migrations, no services, no tests exist.

4. **Email is fake**: v1 uses a logged/stub email sender — no real SMTP
   needed.

5. **INV-account-08 doc inconsistency**: Test description omits
   `revoked_at IS NULL` clause that the invariant statement includes.
   Implementation must follow the statement.
