# Feature Spec — 05: `has_overdue_report` Scheduler

> File: `docs/spec/disbursement/features/05-overdue-report-scheduler.md`
> Domain: `disbursement`
> Task: 05 (see `docs/spec/disbursement/tasks.md`)
> Status: draft — **contains one open item**, authored 2026-08-20
> Last updated: 2026-08-20

## Summary

A scheduled background job that sets
`organizations.has_overdue_report = true` when a disbursed fund
disbursement has had no fund-usage report submitted for 30+ days. Not
an HTTP endpoint. This is the concrete fulfillment of
`organization/invariants.md`'s INV-organization-13 forward reference.

## `[open item]` — needs resolution before implementation

Does the flag-set/clear log entry land in `organization_logs` (the
field's owning table, consistent with this project's "invariants
owned by the domain that owns the field" convention) or somewhere in
`disbursement`'s own log tables? This spec assumes
**`organization_logs`** as the default — confirm before implementing,
since it's a genuine cross-domain write not yet made explicit
anywhere in the source docs.

## Trigger

Scheduled, periodic (exact interval an implementation detail — daily
is a reasonable default given the 30-day window, but this spec doesn't
fix the exact cadence).

## Behavior

1. `UPDATE organizations SET has_overdue_report = true WHERE id IN
   (SELECT DISTINCT c.organization_id FROM disbursement_requests dr
   JOIN campaigns c ON c.id = dr.campaign_id WHERE dr.status =
   'disbursed' AND dr.disbursed_at + interval '30 days' < now() AND
   NOT EXISTS (SELECT 1 FROM fund_usage_reports WHERE
   disbursement_request_id = dr.id)) AND has_overdue_report = false`
   — the trailing `AND has_overdue_report = false` avoids a
   redundant write (and redundant log entry) if the flag is already
   set.
2. For each organization actually flagged by step 1: insert an
   `organization_logs` row (`action_type =
   'has_overdue_report_set'`, `metadata` referencing the specific
   overdue `disbursement_request_id`).
3. Log a run summary (organizations flagged, duration) —
   operational visibility, not a security-relevant audit entry itself.

## Validation & error cases

| Case | Behavior |
|---|---|
| Job overlaps with a still-running previous run | Guard against concurrent runs (advisory lock or equivalent — same pattern as `notification`'s hard-delete worker) |
| DB error mid-run | Log the error, let the run end; next scheduled run picks up whatever's still overdue — stateless sweep, no progress to reconcile |

## Concurrency & correctness notes

- The `AND has_overdue_report = false` clause makes this naturally
  idempotent against re-running before the flag is cleared — an
  organization already flagged doesn't get re-flagged or re-logged.
- No interaction risk with the clearing logic
  (`fund-usage-report-submission.md`, INV-disbursement-06) beyond
  standard transaction isolation — the clear happens inside the report
  submission's own transaction, checking current state at that moment;
  this scheduler's `NOT EXISTS` check similarly reads current state at
  its own run time. A report submitted in the narrow window between
  this job's read and its write could theoretically still get flagged
  — acceptable, since the next run's `NOT EXISTS` check would find the
  report and never re-flag, and the flag doesn't do anything
  irreversible (it only blocks *new* campaign creation, per
  INV-organization-13).

## Test checklist

- [ ] Disbursement past 30 days with no report → flagged, log entry
      created.
- [ ] Disbursement with a report submitted before the deadline →
      never flagged.
- [ ] Disbursement already flagged → not re-flagged, no duplicate log
      entry, on a subsequent run.
- [ ] Two concurrent runs → no double-flagging, no error.
- [ ] **Once the open item above is resolved**: confirm the log entry
      lands in the correct table (`organization_logs`, per this
      spec's default assumption).

## References

- `docs/spec/disbursement/invariants.md` — INV-disbursement-07, 08, 09
- `docs/spec/disbursement/threat-model.md` — "`has_overdue_report`
  scheduler & clearing logic" section (open item)
- `docs/spec/disbursement/tasks.md` — Task 05
- `docs/spec/organization/invariants.md` — INV-organization-13
- `docs/project/kencleng-phase3-detail.md` — Fitur 4
