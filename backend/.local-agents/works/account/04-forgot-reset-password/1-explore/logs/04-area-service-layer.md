# Area 4 — Service layer patterns

> Stage 2 gap analysis. File: `internal/domain/account/service.go`.

## Current state

- **Sentinel error vocabulary matches this feature exactly**:
  `ErrValidation`→422, `ErrTokenExpired`→410, `ErrTokenNotFound`→404
  (service.go:30–41), all mapped in `MapServiceError`.
- **Exact anti-enumeration 3-branch pattern exists in `Register`**
  (service.go:165–238): lookup `email_password` identity by HMAC → nil?
  then lookup `google` identity → Google-only nudge vs new-user;
  `dummyWrite` (tx + 0-row UPDATE + commit) shapes no-op branch DB timing
  to match real branches — explicit anti-timing-side-channel device (R7).
- **`issueNewVerificationToken`** (service.go:352): revoke-old +
  insert-new token in one tx, returns plain token for post-commit email.
  Template for issuing reset tokens — *minus* the revoke step (Assumption
  A forbids it here).
- **`VerifyEmail`** (service.go:399): redeem + mutate in one tx; on
  `!ok`, disambiguates expired-vs-other with a read after rollback.
  Direct template for ResetPassword's redeem step.
- **`validatePassword`** (service.go:473): length ≥8 + breach-check
  fail-open; deliberately runs before any enumeration-sensitive work.
  Directly reusable.
- **Conventions**: email sent only after commit; PII-free sanitized
  logging (`notificationErrorCategory`); `generateToken()` = 32
  crypto/rand bytes hex + SHA-256; clock seam `s.now`; `tokenTTL = 24h`
  const documented as the email_verification TTL.

## Requirement

Two new service methods (`ForgotPassword`, `ResetPassword`) reusing the
above building blocks.

## Gap

No forgot/reset methods exist. All building blocks exist.

## Sniffing findings

1. **Risk / inconsistency — cross-purpose token redemption possible
   today**: `VerifyEmail` discards `RedeemToken`'s returned purpose
   (`userID, _, ok`, service.go:413). A `password_reset` token presented
   to `/auth/verify-email` would be consumed and set `verified_at`. The
   reset flow must check `purpose == "password_reset"`; and once reset
   tokens exist in the wild, the verify-email side of the hole grows more
   consequential. Flagged for Stage 3 decision.
2. **Edge cases** — Assumption B ordering maps cleanly onto existing
   structure: `validatePassword` (network I/O to HIBP) runs before
   `BeginTx`; the tx pattern makes "validation failure leaves token
   unconsumed" natural since redemption happens inside the tx.
3. **Misleading signals** — `NudgePasswordReset` constant sounds like the
   reset email but is the registration-flow nudge text ("use
   forgot-password"); `tokenTTL`(24h) must not leak into reset's 1h TTL.
4. **Risk (minor)** — `NewService` takes 14 positional parameters; every
   new seam ripples through all constructor call sites and test fixtures.
