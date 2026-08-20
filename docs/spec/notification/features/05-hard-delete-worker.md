# Feature Spec — 05: Hard-Delete Housekeeping Worker

> File: `docs/spec/notification/features/05-hard-delete-worker.md`
> Domain: `notification`
> Task: 05 (see `docs/spec/notification/tasks.md`)
> Status: draft
> Last updated: 2026-08-19

## Summary

A scheduled (weekly, per `kencleng-phase0-detail.md` Fitur 6)
background job that physically deletes notification rows whose
`expires_at` has passed. Not an HTTP endpoint — no request/response
contract. Purely a housekeeping pass; the read paths (Tasks 02/03)
already exclude expired rows on their own (INV-notification-03), so
this worker lagging or even failing to run for a cycle never affects
what any user can see — it only affects table size / physical storage.

## Trigger

Scheduled, weekly. Exact scheduling mechanism (cron, a Go
`time.Ticker`-based internal scheduler, external job runner) is an
implementation detail left open for the build phase — this spec fixes
behavior, not the scheduling mechanism.

## Behavior

1. `DELETE FROM notifications WHERE expires_at < now()`.
2. Batch the delete rather than issuing one unbounded statement, e.g.
   `DELETE FROM notifications WHERE id IN (SELECT id FROM notifications
   WHERE expires_at < now() LIMIT 1000)` in a loop until 0 rows are
   affected — avoids holding a long-running lock if the table has
   grown large. (Per threat model: low-severity at this project's
   sandbox scale, but cheap to do correctly from the start rather than
   retrofit later.)
3. Log a summary at the end of each run (rows deleted, duration) — not
   a security-relevant audit entry (this isn't a Fitur-9 sensitive
   action), just operational visibility.

## Validation & error cases

| Case | Behavior |
|---|---|
| Worker run overlaps with a still-running previous run (e.g. a slow run plus the next scheduled tick) | Guard against concurrent runs (e.g. an advisory lock, or a simple in-process mutex/flag if this runs as a singleton within one backend instance) — prevents two runs racing on the same batch-delete loop. Exact mechanism is an implementation detail; the requirement is "no two runs execute concurrently." |
| DB error mid-run (e.g. connection drop) | Log the error, let the run end; the next scheduled run will simply pick up whatever's still expired — no state to reconcile, since this is a stateless sweep over a `WHERE` clause, not a queue with tracked progress. |

## Concurrency & correctness notes

- This worker never deletes a non-expired row — the `WHERE expires_at
  < now()` predicate is the only correctness requirement
  (INV-notification-03's state-machine section). No row-level locking
  concerns with the read paths: `DELETE` and the `SELECT`s in Tasks
  02/03 don't conflict in a way that matters, since a row that's about
  to be deleted was already excluded from read results (its
  `expires_at` already passed) before this worker ever touches it.
- Batching (`LIMIT 1000` loop) means a single worker run may take
  multiple round-trips; this is fine since there's no user-facing
  latency requirement on this worker.

## Test checklist

- [ ] Only rows with `expires_at < now()` are deleted (mixed
      expired/non-expired test data).
- [ ] Non-expired rows are never touched, including rows with
      `read_at IS NULL` (unread doesn't protect from deletion once
      expired — expiry is independent of read state, per
      `kencleng-erd.md`).
- [ ] Batched delete loop correctly terminates (0 rows remain matching
      the predicate after the run) even when the expired set exceeds
      one batch size.
- [ ] Two concurrent invocations of the worker (simulated in test) do
      not error or double-process — guard is effective.

## References

- `docs/spec/notification/invariants.md` — INV-notification-03 (state
  machine section)
- `docs/spec/notification/threat-model.md` — "Weekly hard-delete
  worker" section
- `docs/spec/notification/tasks.md` — Task 05
- `docs/project/kencleng-erd.md` — `ix_notifications_expires_at`
