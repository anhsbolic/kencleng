# Area 5 — Notification touchpoint

> Stage 2 gap analysis. Files: `internal/platform/notification/sender.go`,
> `dev_sender.go`, `sender_test.go`.

## Current state

- `notification.Sender` interface has exactly two methods:
  `SendVerificationEmail(ctx, to, token)` and `SendNudgeEmail(ctx, to,
  nudgeType)` with constants `NudgePasswordReset` / `NudgeGoogleOnly`
  (sender.go:14–29). Implementations: `FakeSender` (non-dev), `DevSender`
  (dev outbox file, mode 0600).
- Service calls senders post-commit only; failures logged category-only
  (no PII).

## Requirement

Forgot-password sends (a) a reset email carrying the plain token, (b) a
distinct Google-only notice email — both post-commit.

## Gap

1. **No sender method carries a password-reset token.** New method (e.g.
   `SendPasswordResetEmail`) needed on the interface + `FakeSender` +
   `DevSender`.
2. **Google-only notice satisfiable today** via `SendNudgeEmail(email,
   NudgeGoogleOnly)` — same semantic ("this account uses Google login")
   already used by Register's conflict branch.

## Sniffing findings

1. **Edge case** — DevSender outbox gains a third line type; AGENTS.md §5
   describes outbox contents as "verification token" — trivial doc drift.
2. **Inconsistency risk** — interface change is compile-time ripple: all
   test fakes implementing `Sender` need the new method. Small, but
   touches files outside the two new ones.
