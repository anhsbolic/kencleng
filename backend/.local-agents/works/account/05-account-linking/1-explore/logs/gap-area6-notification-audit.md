# Stage 2 — Gap Analysis, Area 6: Notification platform + audit path

> Files: `internal/platform/notification/sender.go`, `dev_sender.go`

## Current state (concrete)

- **`platform/notification.Sender`** interface has exactly two methods:
  `SendVerificationEmail(ctx, to, token)` and `SendNudgeEmail(ctx, to,
  nudgeType)`. Implementations: `FakeSender` (non-dev, redacts) and
  `DevSender` (writes dev outbox file). Nudge-type constants today:
  `password_reset`, `google_only` only.
- **Audit path**: `Repository.InsertUserLog(ctx, tx, entry)` — tx-scoped,
  already exercised by `callbackLink` with
  `actionAccountLinking = "account_linking"`. `user_logs.action_type` is
  unconstrained TEXT at the DB level; vocabulary + REVOKE immutability
  deferred to task #08.

## Requirement vs Gap

1. Branch 1's conflict case requires "**a distinct nudge email** is sent
   instead of a verification email" — no such nudge type exists; a new
   constant (and sender agreement) is needed.
2. Branch 1 success sends a verification email — fully reusable via
   existing `SendVerificationEmail` + token-issuance mechanics.
3. Unlink audit entry: reusable as-is (`InsertUserLog` inside the delete
   tx). Set-password audit: spec says write on *successful verification*
   for Branch 1 — but `/auth/verify-email` has no per-user context
   distinguishing "registration verification" from "set-password
   verification" (identical token purpose). Design wrinkle deferred to
   Stage 3 (D7 there).
4. The spec's post-action user-facing notifications ("Metode login baru
   berhasil ditambahkan ke akunmu", etc.) are an explicit forward
   dependency on the unbuilt notification domain — correctly deferred by
   the spec itself; nothing to build now beyond noting the seam.

## Sniffing

- *Misleading signal*: `SendNudgeEmail(to, nudgeType)` looks generic
  enough that "a distinct nudge" appears trivially supported — but each
  nudge type is a stringly-typed contract with only log/outbox behind
  it, untestable end-to-end until a real sender exists.
- *Edge case*: Branch 1 conflict nudge goes to the submitted email —
  same enumeration-safety reasoning as registration nudges applies;
  reuse, don't reinvent.
