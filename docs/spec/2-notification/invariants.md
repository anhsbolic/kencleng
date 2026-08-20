# Domain Invariant — notification

> File: `docs/spec/notification/invariants.md`
> Status: draft
> Last updated: 2026-08-19

## Domain summary

`notification` owns notification creation, delivery-channel record
(`in_app` / `email` — email is fake/logged in this sandbox, not
actually sent), read-state, and expiration/housekeeping. It does
**not** expose a public creation endpoint — notifications are created
only via internal calls from other domains (`account`, `organization`,
`campaign`, `donation`, `disbursement`), and that call is **best-effort
and non-blocking** (see INV-notification-06) — the calling domain's own
transaction never rolls back because a notification failed to be
created.

## Invariants

### INV-notification-01: Every notification has at least one recipient

- **Statement**: For every row in `notifications`,
  `recipient_user_id IS NOT NULL OR recipient_email IS NOT NULL`.
- **Holds after operations**: every notification-creation call site,
  across every calling domain (registered-user recipient uses
  `recipient_user_id`; guest recipient — e.g. `donation_success` for a
  guest donor — uses `recipient_email` only).
- **Verification**: DB-level — the existing `CHECK` constraint on
  `notifications` (already in `kencleng-erd.md`). No new test needed
  beyond confirming the constraint is present in the migration.

### INV-notification-02: `read_at` is monotonic and single-use

- **Statement**: A row in `notifications` may transition `read_at`
  from `NULL` to a timestamp **at most once**. No code path may ever
  set `read_at` back to `NULL`, and no code path may overwrite an
  already-set `read_at` with a different timestamp.
- **Holds after operations**: batch mark-as-read (`POST
  /notifications/mark-read`).
- **Verification**: Test — call mark-as-read twice with overlapping
  `notification_ids` (simulating a retried batch); assert the second
  call is a no-op (`WHERE read_at IS NULL` guard per row) and
  `read_at` does not change on rows already marked read. Concurrency
  test: two simultaneous mark-as-read requests covering the same
  notification id; exactly one write takes effect, no error either
  way (idempotent).

### INV-notification-03: Expired notifications never appear in any read path

- **Statement**: Every query that returns notification data to a user
  — list (`GET /notifications`) and unread count (`GET
  /notifications/unread-count`) — must filter `WHERE expires_at >
  now()`. This holds even for rows not yet physically removed by the
  weekly hard-delete worker (soft-hide is logical, independent of
  physical deletion).
- **Holds after operations**: list, unread-count. Must hold
  regardless of `read_at` state — an expired-but-unread notification
  is still excluded.
- **Verification**: Test — create a notification with `expires_at` in
  the past (bypassing the normal 30-day default for the test setup);
  assert it is absent from both the list response and the
  unread-count, even though the hard-delete worker hasn't run yet.

### INV-notification-04: Ownership boundary on read and mutate

- **Statement**: A user may only read or mark-as-read notifications
  where `recipient_user_id` equals their own authenticated user id.
  No endpoint may return or mutate another user's notification rows.
- **Holds after operations**: list, unread-count, mark-as-read — all
  three authenticated endpoints.
- **Verification**: Test — User A attempts to mark-as-read a
  notification id belonging to User B (obtained out-of-band, e.g. a
  guessed/leaked UUID); assert the row is untouched (treated as
  not-found/no-op from A's perspective, not a 403 that would confirm
  the id's existence — consistent with the project's
  anti-enumeration posture established in `account`).

### INV-notification-05: `expires_at` is fixed at creation, immutable afterward

- **Statement**: `expires_at` is set exactly once, at insert time
  (`created_at` + 30 days, per `kencleng-phase0-detail.md` Fitur 6).
  No code path — including mark-as-read — may update `expires_at`
  after creation.
- **Holds after operations**: notification creation (any calling
  domain). Explicitly does **not** hold as a mutable field for any
  later operation.
- **Verification**: Test — mark a notification as read, then assert
  `expires_at` is unchanged from its value at creation.

### INV-notification-06: Notification creation is internal-only and best-effort

- **Statement**: There is no public HTTP endpoint for creating a
  notification. A failure to insert a `notifications` row (e.g. DB
  error) must **never** cause the calling domain's own operation to
  fail or roll back — the calling domain's primary write (e.g.
  organization verification, campaign approval, donation success)
  commits independently of notification-creation success. A
  notification-creation failure is logged as an error, not propagated
  as a failure to the end user of the triggering action.
- **Holds after operations**: every cross-domain call into
  `notification` — from `account` (already built: `Fitur 2B`'s
  `forgot_password_google_only_notice`) and, once built, from
  `organization`, `campaign`, `donation`, `disbursement`.
- **Verification**: Test — simulate a notification-insert failure
  (e.g. inject an error at the internal call site) during a
  triggering operation (e.g. organization verification); assert the
  triggering operation still succeeds and commits, and the failure is
  logged.
- **Decision — best-effort over same-transaction** **[RESOLVED —
  2026-08-19]**: chosen over making notification-insert part of the
  same DB transaction as the triggering event, so that a `notification`
  table/write problem can never block or fail an unrelated domain's
  core operation. Accepted trade-off: a notification can, in a rare
  failure case, silently never be created for an event that otherwise
  succeeded (no retry mechanism in v1) — see "Knowingly accepted
  residual risk" in `docs/spec/notification/threat-model.md` once
  written.

### INV-notification-07: `type` must be a known value at write time

- **Statement**: Although `notifications.type` is a plain `TEXT`
  column with no DB-level `CHECK` (deliberately, to stay extensible —
  see `kencleng-erd.md` "Open Items Carried Into Migration Writing"),
  every application-layer write must validate `type` against the
  known value set (the `NotificationType` enum in `api/openapi.yaml`)
  before insert. No code path may insert an arbitrary/unvalidated
  string.
- **Holds after operations**: every notification-creation call site,
  across every calling domain.
- **Verification**: Test — attempt to create a notification with an
  unrecognized `type` string; assert the internal creation function
  rejects it (returns an error to the caller, which — per
  INV-notification-06 — must still not fail the caller's own
  transaction; this is a caller-side logic bug being caught, not a
  transient failure).

## State machine

### `notifications.read_at`

```
null -> read (timestamp)
```

One-way, single transition (see INV-notification-02). No "unread"
transition exists.

### `notifications.expires_at` vs. row lifecycle

Not a per-row state machine in the traditional sense, but a two-layer
visibility/lifecycle model (per `kencleng-phase0-detail.md` Fitur 6):

```
created (expires_at = created_at + 30d)
  -> [soft-hidden once now() > expires_at, still physically present]
  -> [hard-deleted by weekly worker, WHERE expires_at < now()]
```

The soft-hide filter (INV-notification-03) and the hard-delete worker
both key off the same `expires_at` column and the same index
(`ix_notifications_expires_at`), but soft-hide is enforced on every
read while hard-delete is a periodic housekeeping pass — the two are
independent and hard-delete lagging behind never affects what a user
can see.

## References

- Related ERD: `docs/project/kencleng-erd.md` §6 "Cross-Cutting"
  (`notifications` table)
- Related business process: `docs/project/kencleng-phase0-detail.md`
  Fitur 6 (Notification Infrastructure)
- Related OpenAPI: `api/openapi.yaml` — `notification` tag
  (`GET /notifications`, `GET /notifications/unread-count`,
  `POST /notifications/mark-read`), and schemas `Notification`,
  `NotificationType`, `NotificationChannel`
- Cross-domain callers (forward references — not yet built except
  `account`):
  - `account` (already built) — Fitur 2B, `forgot_password_google_only_notice`
  - `organization` — Fitur "Manage Representative" (`representative
    added` notice), curation-queue notices (`admin_new_curation_item`,
    `kurator_assigned`), org verification (`organisasi_verified`,
    `organisasi_rejected`)
  - `campaign` — `campaign_approved`, `campaign_rejected`,
    `campaign_closed`
  - `donation` — `donation_success`
  - `disbursement` — `fund_usage_report_verified`,
    `fund_usage_report_rejected`, `disbursement_approved`,
    `disbursement_rejected`, `disbursement_completed`
  - Each of the above domains' `invariants.md`, once written, must
    **reference** INV-notification-06 (not redefine it) when
    documenting their own call into `notification` — same
    cross-domain ownership rule used for INV-account-10.
