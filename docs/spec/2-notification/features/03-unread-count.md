# Feature Spec — 03: Unread Count

> File: `docs/spec/notification/features/03-unread-count.md`
> Domain: `notification`
> Task: 03 (see `docs/spec/notification/tasks.md`)
> Status: draft
> Last updated: 2026-08-19

## Summary

`GET /notifications/unread-count` — count of the authenticated user's
unread, non-expired notifications. Powers the notification-center
badge; called on every relevant page load per
`kencleng-phase0-detail.md` Fitur 6, so it needs to stay cheap.

## Endpoint

`GET /notifications/unread-count` (per `api/openapi.yaml`,
`notification` tag)

## Auth

`bearerAuth` required. No role restriction — same as list.

## Request

No parameters.

## Behavior

1. Resolve `current_user_id` from the authenticated session.
2. `SELECT COUNT(*) FROM notifications WHERE recipient_user_id =
   current_user_id AND read_at IS NULL AND expires_at > now()`,
   served by `ix_notifications_recipient_unread` (partial index on
   `(recipient_user_id, read_at) WHERE recipient_user_id IS NOT
   NULL`).
3. Return the count.

## Validation & error cases

| Case | Behavior |
|---|---|
| No/invalid bearer token | `401` |

Nothing else to validate — no input parameters.

## Concurrency & correctness notes

- Must use the exact same filter predicate as the list endpoint's
  `unread_only=true` path (`recipient_user_id = current_user_id AND
  read_at IS NULL AND expires_at > now()`) — any drift between the two
  becomes a visible, confusing bug (badge count disagrees with the
  actual unread list). Recommend sharing the WHERE-clause construction
  in code (e.g. one query-builder function used by both Task 02 and
  Task 03) rather than duplicating the predicate.
- No caching — always a fresh count on each call. Acceptable at this
  project's scale; a production system might debounce/cache this
  client-side, but that's a frontend concern, not this endpoint's.

## Test checklist

- [ ] Count matches the number of rows the list endpoint
      (`unread_only=true`) returns for the same user, in a combined
      test — no drift between the two.
- [ ] Excludes expired notifications, regardless of `read_at`.
- [ ] 0 cross-user leakage (2+ user test matrix).
- [ ] Count updates correctly immediately after a mark-as-read call
      (Task 04) in an integration test spanning both endpoints.

## References

- `docs/spec/notification/invariants.md` — INV-notification-03, 04
- `docs/spec/notification/threat-model.md` — `GET
  /notifications/unread-count` section
- `docs/spec/notification/tasks.md` — Task 03
- `docs/project/kencleng-erd.md` — `ix_notifications_recipient_unread`
