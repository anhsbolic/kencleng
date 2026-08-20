# Feature Spec — 02: List Notifications

> File: `docs/spec/notification/features/02-list-notifications.md`
> Domain: `notification`
> Task: 02 (see `docs/spec/notification/tasks.md`)
> Status: draft
> Last updated: 2026-08-19

## Summary

`GET /notifications` — cursor-paginated list of the authenticated
user's own notifications, newest first, excluding expired rows.

## Endpoint

`GET /notifications` (per `api/openapi.yaml`, `notification` tag)

## Auth

`bearerAuth` required (global security scheme, same JWT as every other
authenticated endpoint). No role restriction — any authenticated actor
(Donatur, Organisasi Representative, Kurator, Admin) sees their own
notifications only.

## Request

| Param | In | Required | Notes |
|---|---|---|---|
| `cursor` | query | No | `CursorParam` — opaque UUIDv7 from a prior page's `pagination.next_cursor`. Omit for first page. |
| `limit` | query | No | `LimitParam` — default 20, max 50. |
| `unread_only` | query | No | If true, filter to `read_at IS NULL` only (exact param name/shape TBD against `api/openapi.yaml` at implementation time if not already present — confirm before coding). |

## Behavior

1. Resolve `current_user_id` from the authenticated session.
2. Query: `WHERE recipient_user_id = current_user_id AND expires_at >
   now()`, optionally `AND read_at IS NULL` when `unread_only=true`.
3. Order newest-first (by `created_at` or the cursor's UUIDv7 — the
   two are equivalent in ordering since UUIDv7 is time-ordered; use
   whichever the cursor implementation already relies on for
   consistency).
4. Paginate via cursor, `limit` capped at 50.
5. Return `NotificationListResponse` (`items: Notification[]`,
   `pagination`).

## Validation & error cases

| Case | Behavior |
|---|---|
| No/invalid bearer token | `401` (standard auth middleware) |
| Malformed `cursor` (not a valid UUID) | `422` — fail closed, do **not** silently fall back to page 1 (see threat model "Tampering" note for this endpoint) |
| `limit` out of `[1, 50]` range | `422` (schema-level validation, standard OpenAPI request validation) |

## Concurrency & correctness notes

- Cursor is UUIDv7-based (time-ordered), so pagination is stable
  against concurrent inserts: a new notification arriving mid-scroll
  either appears entirely on a page not yet fetched, or not at all in
  the current pagination pass — no duplicate or skipped row across
  pages, by construction.
- Expiry filtering (`expires_at > now()`) is evaluated fresh on every
  call — there's no caching layer here, so a notification that expires
  between two page fetches simply disappears from the next page. This
  is expected and consistent with INV-notification-03.

## Test checklist

- [ ] Returns only the authenticated user's own notifications (2+
      user test matrix, interleaved notifications).
- [ ] Excludes notifications with `expires_at` in the past, regardless
      of `read_at`.
- [ ] `unread_only=true` returns only `read_at IS NULL` rows.
- [ ] Pagination: fetching all pages with `limit=1` returns the same
      full set (in the same order) as fetching with `limit=50` — no
      duplicates, no gaps.
- [ ] Malformed `cursor` returns `422`, not a silent first-page
      fallback.
- [ ] A notification inserted concurrently with an in-progress
      pagination pass does not cause a duplicate or skipped row across
      the already-fetched pages.

## References

- `docs/spec/notification/invariants.md` — INV-notification-03, 04
- `docs/spec/notification/threat-model.md` — `GET /notifications`
  section
- `docs/spec/notification/tasks.md` — Task 02
- `api/openapi.yaml` — `CursorParam`, `LimitParam`,
  `NotificationListResponse`, `Notification` schemas
