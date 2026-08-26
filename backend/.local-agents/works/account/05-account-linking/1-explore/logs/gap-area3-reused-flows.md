# Stage 2 — Gap Analysis, Area 3: Reused-flow implementations

> Files: `transport/http/auth_register.go`, `auth_verify_email.go`,
> `errors.go`; `docs/spec/1-account/invariants.md`;
> `docs/spec/1-account/tasks.md`

## Current state (concrete)

- **`auth_register.go`**: `RegisterHandler` — JSON decode → boundary
  validation (`looksLikeEmail`, name 1–255, password ≥8 as
  defense-in-depth) → `svc.Register` → **identical generic 202** for all
  four service branches ("If your email is not already registered...").
  The anti-enumeration pattern the spec invokes for Branch 1 is fully
  realized, including timing-equivalence work in the service
  (`dummyWrite`).
- **`auth_verify_email.go`**: `VerifyEmailHandler` +
  `ResendVerificationHandler` exist and map to
  `Service.VerifyEmail`/`ResendVerification` — the exact endpoint
  set-password's Branch 1 step 2 reuses unchanged.
- **`errors.go`**: RFC 9457 Problem Details helpers (`WriteProblem`,
  `WriteValidationError` → 422 with field errors, `write400InvalidJSON`)
  plus `MapServiceError` switch over account sentinels — including
  `ErrInvalidCredentials` → 401 with generic Indonesian message "Email
  atau password salah." from the login slice. New sentinels map cleanly
  into this switch.
- **INV-account-05 implementation status: does not exist.** Grep across
  `internal/` finds no revoke-all-by-user operation, no
  forgot/reset-password handler or service method, and `tasks.md`
  confirms task #4 is **not started**.
- INV statements verified against invariants.md: INV-01 (per-provider
  uniqueness), INV-02 (≥1 identity, unlink the only remover), INV-05
  (reset revokes all sessions, same tx as credential update), INV-08
  (3-clause single-use tokens), INV-12 (remaining identity must be
  verified, distinct error from INV-02's case).

## Requirement vs Gap

1. The feature spec's central claim — Branch 2 "revokes all refresh
   tokens (**same INV-account-05 pattern** as
   `04-forgot-reset-password.md`, **reused here rather than
   re-derived**)" — has **nothing to reuse**. Fitur 04 is unimplemented;
   there is no user-scoped refresh-token revocation anywhere
   (`RevokeRefreshTokenFamily` is family-scoped). Whatever session-
   revocation primitive gets built here will be the *first*
   instantiation of INV-account-05; Fitur 04 will later reuse *it* —
   the dependency direction in the spec is inverted relative to reality.
2. New sentinel errors needed for the two distinct 409 unlink cases
   (spec mandates distinct Indonesian messages — neither exists).
3. No `/account/security/*` handler files exist; authenticated user_id
   extraction unverified at this point (resolved in Area 5).

## Sniffing

- *Miscontext (major)*: the spec author believes Fitur 04's force-logout
  pattern exists to copy; it doesn't. Also tasks.md orders S1 serially
  `#1→#2→#3→#4→#5` — this task (#5) is being worked before #3 and #4
  are finished, contrary to the declared serial order.
- *Inconsistency*: `tasks.md` says task #3 "build not started", yet
  `service.go` (login seams), `login.go`, `mfa_verifier.go`, migrations
  000006–000009, and the login error vocabulary all exist in the tree —
  the tracker appears stale relative to code.
- *Misleading signal*: `ErrInvalidCredentials` → 401 already exists,
  which could suggest "wrong current_password handling is done" — it
  isn't wired to anything here; the comparison target differs (the
  caller's own verified identity's `credential_secret`, not a login-time
  identifier lookup).
- *Risk*: because INV-account-05's atomicity requirement lands here
  first, whatever transaction shape Area 1 showed (`TxRunner` + guarded
  UPDATEs) becomes the precedent both features inherit.
