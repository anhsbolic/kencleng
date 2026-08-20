# Feature Spec — 01: Internal Notification Creation

> File: `docs/spec/notification/features/01-internal-notification-creation.md`
> Domain: `notification`
> Task: 01 (see `docs/spec/notification/tasks.md`)
> Status: draft
> Last updated: 2026-08-19

## Summary

An in-process function, `notification.Create(ctx, ...)` (exact
package/function name TBD at implementation time — this spec fixes
behavior, not Go symbol names), that every other domain calls as a
best-effort side effect of their own operation. This is **not** an
HTTP endpoint — there is no route for it in `api/openapi.yaml` (see
the `notification` tag description). It is the one piece of this
domain every other domain (`account` already, `organization`,
`campaign`, `donation`, `disbursement` later) directly depends on.

## Trigger

Called synchronously, in-process, by another domain's own handler
code — e.g. `account`'s forgot-password flow calls this after
detecting a Google-only account, to create a
`forgot_password_google_only_notice` notification (Fitur 2B).

## Inputs

| Param | Type | Required | Notes |
|---|---|---|---|
| `type` | `NotificationType` (string enum, per `api/openapi.yaml`) | Yes | Must be a recognized value — see "Validation" below (INV-notification-07) |
| `channel` | `NotificationChannel` (string enum) | Yes | `in_app` or `email` |
| `recipient_user_id` | UUID | Conditionally | At least one of `recipient_user_id` / `recipient_email` required (INV-notification-01) |
| `recipient_email` | string | Conditionally | Used for guest recipients (no `users` row); stored encrypted-at-rest per the project's PII pattern (`BYTEA` ciphertext + `TEXT` HMAC lookup hash), consistent with `donations.guest_email` |
| `payload` | JSONB | Yes (may be `{}`) | Minimal, type-specific shape — see "Payload shape convention" below |

## Behavior

1. Validate `type` against the known `NotificationType` set. Reject
   (return an error to the **caller**, not to any end user) if
   unrecognized — this is a caller-side bug being caught, not a
   transient failure, and per INV-notification-06 the caller must
   still not let this fail its own transaction; it should log and move
   on.
2. Validate that at least one of `recipient_user_id` /
   `recipient_email` is present (application-layer check, ahead of the
   DB `CHECK` — belt-and-suspenders, not a replacement for it).
3. Set `expires_at = now() + 30 days` (fixed, per Fitur 6). No caller
   may override this.
4. Insert the row. **Do not** wrap this insert in the caller's
   transaction — it runs as its own, independent unit of work (this is
   the concrete mechanism behind INV-notification-06's "best-effort,
   non-blocking" requirement; see "Concurrency & correctness notes").
5. On insert failure (DB error, connection issue, etc.), log the error
   (structured log, includes `type`, `recipient_user_id` or a
   non-reversible reference to `recipient_email` — never log the raw
   email) and return a non-fatal error value to the caller. The caller
   is expected to ignore it for control-flow purposes (per
   INV-notification-06) but may itself log that a notification was
   attempted.
6. `channel = email` does **not** actually send an email in this
   sandbox — it only records the row with `channel = 'email'`. Actual
   sending is out of scope for this project (per `kencleng-backend-
   tech-stack.md`'s sandbox scope); a future implementer would hook
   real delivery here without changing this function's contract.

## Payload shape convention

`payload` is free-form JSONB with no DB-level schema, but each calling
domain's own feature spec should define a minimal, type-specific shape
when it adds a new `NotificationType` — e.g. `campaign_approved` might
carry `{"campaign_id": "...", "campaign_title": "..."}`, nothing more.
This function does not enforce that shape (see threat model,
"Information disclosure" row for internal creation) — enforcement is a
convention documented per-caller, not a technical control here.

## Validation & error cases

| Case | Behavior |
|---|---|
| Unrecognized `type` | Reject before insert; return error to caller (non-fatal for caller's own transaction) |
| Neither `recipient_user_id` nor `recipient_email` set | Reject before insert; return error to caller |
| DB insert fails (any reason) | Log error, return non-fatal error to caller; caller's own operation is unaffected |

## Concurrency & correctness notes

- This function must **not** accept or require the caller's
  `*sql.Tx` / transaction handle. It opens (or uses a pooled)
  connection independent of the caller's transaction, so a rollback in
  the caller never undoes a notification that was already created, and
  a notification-insert failure never forces a rollback in the caller.
  This is the direct implementation of INV-notification-06's
  best-effort requirement — call this out explicitly in code review,
  since "just pass the caller's tx" is the obvious-but-wrong shortcut
  here.
- No idempotency key / dedup logic at this layer — if a caller
  accidentally calls `Create` twice for the same logical event (e.g. a
  bug causing double-invocation), two notification rows are created.
  Preventing that is the caller's responsibility (e.g. via its own
  transaction ensuring the triggering action itself only happens
  once); not something this function guards against.

## Test checklist

- [ ] Reject unrecognized `type` before attempting insert (unit test,
      no DB required).
- [ ] Reject when both `recipient_user_id` and `recipient_email` are
      absent.
- [ ] Accept `recipient_user_id`-only and `recipient_email`-only cases
      independently.
- [ ] `expires_at` is always `created_at + 30d`, not caller-settable.
- [ ] Simulated insert failure (e.g. mock/fault-injected DB) does not
      panic and returns a non-fatal error.
- [ ] Integration test with a real caller (`account`'s forgot-password
      Google-only path): forcing a notification-insert failure still
      results in the forgot-password request succeeding end-to-end.
- [ ] `recipient_email`, when present, is stored encrypted (ciphertext
      + HMAC lookup hash), never in plaintext.

## References

- `docs/spec/notification/invariants.md` — INV-notification-01, 05,
  06, 07
- `docs/spec/notification/threat-model.md` — "Internal notification
  creation" section
- `docs/spec/notification/tasks.md` — Task 01
- `api/openapi.yaml` — `NotificationType`, `NotificationChannel`
  schemas; `notification` tag description
- `docs/project/kencleng-phase0-detail.md` — Fitur 6
