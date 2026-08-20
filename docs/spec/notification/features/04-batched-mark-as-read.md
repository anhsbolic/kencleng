# Feature Spec — 04: Batched Mark-As-Read

> File: `docs/spec/notification/features/04-batched-mark-as-read.md`
> Domain: `notification`
> Task: 04 (see `docs/spec/notification/tasks.md`)
> Status: draft
> Last updated: 2026-08-19

## Summary

`POST /notifications/mark-read` — marks a batch of the authenticated
user's own notifications as read, either by explicit id list or via
`all` (every currently-unread notification). Batched by design per
Fitur 6 — client-side interactions are debounced into one request
rather than firing per click.

## Endpoint

`POST /notifications/mark-read` (per `api/openapi.yaml`,
`notification` tag)

## Auth

`bearerAuth` required. No role restriction.

## Request

`MarkNotificationsReadRequest`:

| Field | Type | Required | Notes |
|---|---|---|---|
| `notification_ids` | `UUID[]` | Conditionally | Mutually exclusive with `all` |
| `all` | `boolean` | Conditionally | Mutually exclusive with `notification_ids` |

**Open item resolved here**: `notification_ids` had no `maxItems`
bound in `api/openapi.yaml` (flagged in
`docs/spec/notification/threat-model.md`). **Decision — cap at 50**,
matching `LimitParam`'s max page size, since a client only ever has at
most one page's worth of notification ids visible to batch-mark at
once under normal UI flow. `api/openapi.yaml` needs
`MarkNotificationsReadRequest.notification_ids` updated with
`maxItems: 50` before this task is implemented — **not yet applied to
the file**, call this out explicitly when starting implementation.

## Behavior

1. Resolve `current_user_id` from the authenticated session.
2. Validate exactly one of `notification_ids` / `all` is present (not
   both, not neither) — `422` otherwise.
3. If `notification_ids`: `UPDATE notifications SET read_at = now()
   WHERE id = ANY($1) AND recipient_user_id = current_user_id AND
   read_at IS NULL`.
4. If `all`: `UPDATE notifications SET read_at = now() WHERE
   recipient_user_id = current_user_id AND read_at IS NULL AND
   expires_at > now()` — note the explicit `expires_at > now()` guard
   here even though marking an expired notification read is harmless;
   included to avoid unnecessarily touching rows that are about to be
   hard-deleted anyway.
5. Return `204 No Content` regardless of how many rows were actually
   affected (0 affected — e.g. all ids were already read, or all
   foreign — is not an error; see "Validation & error cases").

## Validation & error cases

| Case | Behavior |
|---|---|
| No/invalid bearer token | `401` |
| Both `notification_ids` and `all` present, or neither | `422` |
| `notification_ids` exceeds `maxItems` (50, per decision above) | `422` (schema-level validation) |
| `notification_ids` contains ids not owned by the caller | Silently no-op for those ids — `204`, no error, no signal distinguishing this from "already read" (INV-notification-04, deliberate anti-enumeration behavior) |
| `notification_ids` contains ids that don't exist at all | Same as above — silently no-op, `204` |
| Empty `notification_ids: []` | Treat as a no-op `204` (not an error) — simplest, consistent behavior; avoids a special-case error path for a harmless client-side edge case |

## Concurrency & correctness notes

- The `WHERE read_at IS NULL` guard in the `UPDATE` makes this
  idempotent per-row by construction — submitting the same batch twice
  in a row (e.g. a retried request after a timeout) is a safe no-op
  the second time.
- Two concurrent mark-as-read calls with overlapping ids: standard
  row-level locking in Postgres serializes the two `UPDATE`s; whichever
  commits first sets `read_at`, the second's `WHERE read_at IS NULL`
  guard means it simply updates 0 rows for the already-handled ones —
  no error, no lost update, no double-write.
- `all=true` racing against a new notification being created
  concurrently: the new notification is either included (if its insert
  commits before this `UPDATE`'s snapshot) or not (if after) — either
  outcome is acceptable; there's no correctness requirement that
  `all=true` be atomic with respect to concurrent creation, since "mark
  everything read as of roughly now" is the intended semantics, not a
  hard snapshot guarantee.

## Test checklist

- [ ] Idempotency: same batch submitted twice results in the same end
      state, second call is a no-op, no error.
- [ ] Concurrency: two simultaneous calls with overlapping
      `notification_ids` result in the expected final state, no error
      from either.
- [ ] Ownership: batch mixing caller's own ids with another user's ids
      only mutates the caller's own rows; response is `204` either way
      (no signal).
- [ ] `all=true` marks every currently-unread, non-expired notification
      read and leaves already-read and expired ones untouched.
- [ ] Both `notification_ids` and `all` present → `422`.
- [ ] Neither present → `422`.
- [ ] `notification_ids` over the `maxItems` bound → `422` (once the
      OpenAPI schema is updated per the decision above).
- [ ] Empty `notification_ids: []` → `204`, no error.

## References

- `docs/spec/notification/invariants.md` — INV-notification-02, 04, 05
- `docs/spec/notification/threat-model.md` — `POST
  /notifications/mark-read` section (DoS / `maxItems` gap)
- `docs/spec/notification/tasks.md` — Task 04
- `api/openapi.yaml` — `MarkNotificationsReadRequest` schema (needs
  `maxItems: 50` added to `notification_ids` per the decision above)
