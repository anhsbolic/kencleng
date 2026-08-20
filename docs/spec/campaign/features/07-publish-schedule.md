# Feature Spec — 07: Publish / Schedule / Reschedule / Republish

> File: `docs/spec/campaign/features/07-publish-schedule.md`
> Domain: `campaign`
> Task: 07 (see `docs/spec/campaign/tasks.md`)
> Status: draft — authored against `api/openapi/campaign.yaml` 2026-08-20
> Last updated: 2026-08-20

## Summary

`POST /campaigns/{campaignId}/publish` — one endpoint handling all
four related actions via the same contract: publish now, schedule for
later, reschedule (while `scheduled`), and republish (from
`unpublished`). Owner-only. Plus the background scheduler job that
executes time-triggered `scheduled → published` transitions.

## Endpoint

`POST /campaigns/{campaignId}/publish` (confirmed,
`api/openapi/campaign.yaml`) + a scheduler job (no HTTP endpoint)

## Auth

`bearerAuth` required, owner-level representative only.

## Request

`PublishRequest`:

| Field | Type | Required | Notes |
|---|---|---|---|
| `publish_at` | datetime, nullable | No | Omit to publish immediately; provide a future datetime to schedule/reschedule |

## Behavior

### `POST .../publish`
1. Reject (`403`) if caller isn't owner-level.
2. Reject (`409`) if `status` isn't one of {`approved`, `scheduled`,
   `unpublished`} — first-publish requires `approved`; reschedule
   requires `scheduled`; republish requires `unpublished`.
3. If `publish_at` is provided: validate it's `> now()` and `≤ deadline`
   (`422` otherwise).
4. If `publish_at` omitted: `status = 'published'`, `published_at =
   now()`.
5. If `publish_at` provided: `status = 'scheduled'`, `publish_at`
   stored (this also handles reschedule — same write, whether coming
   from `approved`/`unpublished` or already `scheduled`).
6. Return `200` with the updated `Campaign`.

### Scheduler job (background, no endpoint)
1. Periodically: `UPDATE campaigns SET status = 'published',
   published_at = now() WHERE status = 'scheduled' AND publish_at <=
   now()`.
2. Conditional update makes this naturally idempotent — running the
   job twice in the same window never double-publishes (INV-campaign-09).

## Validation & error cases

| Case | Response |
|---|---|
| No/invalid bearer token | `401` |
| Caller is `staff` or not a representative | `403` |
| `status` not in {`approved`, `scheduled`, `unpublished`} | `409` |
| `publish_at` in the past, or beyond `deadline` | `422` |

## Concurrency & correctness notes

- Scheduler job must use the conditional `WHERE status = 'scheduled'
  AND publish_at <= now()` update — never read-then-write separately,
  to avoid a double-publish race if the job overlaps with itself
  (e.g. a slow run plus the next scheduled tick).

## Test checklist

- [ ] Publish now (no `publish_at`) from `approved` → `status =
      published`.
- [ ] Schedule (`publish_at` in the future) from `approved` → `status
      = scheduled`.
- [ ] Reschedule while `scheduled` → `publish_at` updates, `status`
      unchanged.
- [ ] Republish from `unpublished` → `scheduled`/`published` per the
      same contract.
- [ ] `publish_at` in the past → `422`.
- [ ] `publish_at` beyond `deadline` → `422`.
- [ ] Wrong starting status (e.g. `draft`, `pending_curation`) → `409`.
- [ ] `staff` caller → `403`.
- [ ] Scheduler job run twice near-simultaneously on the same due
      campaign: exactly one transition, second run is a no-op.

## References

- `docs/spec/campaign/invariants.md` — INV-campaign-08, 09
- `docs/spec/campaign/threat-model.md` — "Publish / unpublish /
  republish" section
- `docs/spec/campaign/tasks.md` — Task 07
- `api/openapi/campaign.yaml` — `POST /campaigns/{id}/publish`
