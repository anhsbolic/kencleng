# Domain Tasks — notification

> File: `docs/spec/notification/tasks.md`
> Status: draft
> Last updated: 2026-08-19

Task order follows dependency, not the OpenAPI tag's listing order.
Task 01 (internal creation) has no HTTP endpoint but is a hard
prerequisite for everything else — nothing in this domain is
testable end-to-end without a way to create rows first, and it's also
the piece every other domain (`account` already, `organization`,
`campaign`, `donation`, `disbursement` later) directly depends on.

| # | Task | Endpoint / surface | Depends on | Related invariants |
|---|---|---|---|---|
| 01 | Internal notification creation | *(none — internal package function, e.g. `notification.Create(ctx, ...)`)* | — | INV-notification-01, 05, 06, 07 |
| 02 | List notifications | `GET /notifications` | 01 | INV-notification-03, 04 |
| 03 | Unread count | `GET /notifications/unread-count` | 01 | INV-notification-03, 04 |
| 04 | Batched mark-as-read | `POST /notifications/mark-read` | 01 | INV-notification-02, 04, 05 |
| 05 | Hard-delete housekeeping worker | *(none — scheduled job, no HTTP endpoint)* | 01 | INV-notification-03 (state-machine section) |

## Task 01 — Internal notification creation

**What**: the `notification.Create(...)` function (or equivalent) that
every other domain calls, in-process, as a best-effort side effect of
their own transaction (INV-notification-06). Covers: validating
`type` against the known enum (INV-notification-07), enforcing the
recipient `CHECK` at the application layer before insert as a
belt-and-suspenders check ahead of the DB constraint
(INV-notification-01), setting `expires_at = created_at + 30d`
(INV-notification-05), and never returning an error in a way that
would tempt a caller into rolling back its own transaction.

**KPI / metrics**:
- 100% of calls with an unrecognized `type` are rejected before an
  insert is attempted (unit test, no reliance on the DB constraint
  alone).
- Simulated creation failure (e.g. forced DB error in a test) never
  propagates as a failure of the caller's own operation in an
  integration test that exercises a real caller (`account`'s
  `forgot_password_google_only_notice` call site).
- 0 rows created without at least one recipient, verified both by the
  DB `CHECK` and by an application-layer test bypassing the DB layer
  (i.e. testing the Go function's own guard, not just relying on the
  constraint to catch it).

## Task 02 — List notifications

**What**: `GET /notifications`, cursor-paginated (`CursorParam`,
`LimitParam`, max 50), scoped to the authenticated user, excluding
expired rows.

**KPI / metrics**:
- 0 cross-user leakage across a test matrix of at least 2 users with
  interleaved notifications.
- Expired notifications (`expires_at` forced into the past in test
  setup) never appear, regardless of `read_at`.
- Pagination is stable under concurrent inserts (a new notification
  arriving mid-pagination doesn't duplicate or skip a row across
  pages) — cursor is UUIDv7-based (time-ordered), so this should hold
  by construction; test to confirm.

## Task 03 — Unread count

**What**: `GET /notifications/unread-count`, same scoping and
expiry-filtering as list, backed by
`ix_notifications_recipient_unread`.

**KPI / metrics**:
- Count matches the actual number of list-endpoint rows with
  `read_at IS NULL` for the same user, in a combined test (no drift
  between the two endpoints' filtering logic).
- 0 cross-user leakage, same test matrix as Task 02.

## Task 04 — Batched mark-as-read

**What**: `POST /notifications/mark-read`, accepting either
`notification_ids` or `all` (mutually exclusive per the OpenAPI
schema), idempotent per-row update, `204 No Content` response with no
per-id result signal.

**KPI / metrics**:
- Idempotency: submitting the same batch twice results in the same
  end state and no error on the second call.
- Concurrency: two simultaneous mark-as-read calls covering an
  overlapping set of ids result in exactly the expected final state,
  no lost updates, no error.
- Ownership: a batch mixing the caller's own ids with another user's
  ids only mutates the caller's own rows, and the response gives no
  signal distinguishing "foreign id, ignored" from "already read."
- **Resolved 2026-08-19**: `notification_ids` capped at `maxItems: 50`
  in `api/openapi/notification.yaml`, matching `LimitParam`'s max page
  size — closes the DoS gap originally flagged in
  `docs/spec/notification/threat-model.md`.

## Task 05 — Hard-delete housekeeping worker

**What**: scheduled job (weekly, per `kencleng-phase0-detail.md`
Fitur 6) that physically deletes rows `WHERE expires_at < now()`.

**KPI / metrics**:
- 0 non-expired rows ever deleted (test with a mix of expired and
  non-expired rows, assert only expired ones are removed).
- Worker running late/skipped a cycle never affects read-path
  correctness (list/unread-count already filter by `expires_at` —
  this is really a re-confirmation of INV-notification-03's
  independence from this worker, exercised here from the worker's
  side).

## References

- Related domain invariants: `docs/spec/notification/invariants.md`
- Related threat model: `docs/spec/notification/threat-model.md`
- Feature specs (to be written next, one per task, prefixed with the
  task `#` above): `docs/spec/notification/features/`