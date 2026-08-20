# Stage 1 — Plan Announcement

> Feature: Register & Email Verification (`01-register-email-verification.md`)
> Domain: account
> Date: 2026-08-19

## Areas to explore (in order)

1. **API contract** (`api/openapi.yaml`, lines ~1726-1830) — the exact
   request/response schemas for `POST /auth/register`,
   `POST /auth/verify-email`, `POST /auth/verify-email/resend`. Must go
   first because the contract defines what the handler signatures and
   DTOs must look like.

2. **Domain data model** (`docs/project/kencleng-erd.md` §1 — `users`,
   `auth_identities`, `auth_tokens` tables) — needed to understand what
   the repository layer must insert/query, and the PII encryption
   pattern (`primary_email` ciphertext + `primary_email_hash` HMAC).

3. **Domain invariants & threat model**
   (`docs/spec/domains/account/invariants.md` INV-account-01 +
   INV-account-08, `docs/spec/domains/account/threat-model.md` §1) —
   the correctness constraints (concurrent uniqueness, single-use
   tokens, anti-enumeration) that the service layer must enforce.

4. **Existing platform scaffolding** (`internal/platform/crypto/`,
   `internal/platform/auth/`, `internal/platform/db/`,
   `internal/platform/ratelimit/`, plus `cmd/server/main.go`) — what's
   already wired and reusable vs. what needs to be built from scratch.

5. **Backend tech dependencies & conventions**
   (`docs/project/kencleng-backend-tech-stack.md` for password hashing,
   HaveIBeenPwned, email sending; `backend/AGENTS.md` for goqu, error
   wrapping, PII pattern) — the implementation constraints that shape
   how the domain layer is written.

## Order rationale

1→2→3 is dependency chain (contract → data model → correctness rules).
4 checks what infra is ready before planning what to build. 5 captures
the library/tooling choices that constrain implementation.

The backend is essentially greenfield (`internal/domain/` has only
`.gitkeep`), so the gap is "everything doesn't exist yet" — the
exploration focuses on precisely *what* needs to exist and *what
constraints* it must satisfy, not on diffing against existing code.
