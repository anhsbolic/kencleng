# Feature Spec — 09: Closure (Auto + Force-Close)

> File: `docs/spec/campaign/features/09-closure.md`
> Domain: `campaign`
> Task: 09 (see `docs/spec/campaign/tasks.md`)
> Status: draft — authored against `api/openapi/campaign.yaml` 2026-08-20
> Last updated: 2026-08-20

## Summary

A `published` campaign closes via exactly one of three independent
triggers, all sharing the same `WHERE status = 'published'` idempotency
guard (INV-campaign-13): `max_amount` reached (donation-domain
transaction, forward reference), `deadline` reached (scheduler), or
Admin force-close (`POST /campaigns/{campaignId}/force-close`).

## Endpoint

`POST /campaigns/{campaignId}/force-close` (confirmed,
`api/openapi/campaign.yaml`) + a deadline scheduler job + a hook into
`donation` domain's donation-success transaction (both no HTTP
endpoint)

## Auth

Force-close: `bearerAuth` + `role = 'admin'` only.

## Request

`ForceCloseRequest`:

| Field | Type | Required | Notes |
|---|---|---|---|
| `decision_note` | string | Yes | Non-empty |

## Behavior

### Force-close
1. Reject (`403`) if caller isn't Admin.
2. Reject (`422`) if `decision_note` is empty.
3. `UPDATE campaigns SET status = 'closed', closed_at = now(),
   closed_reason = 'admin_force_closed', closed_by = current_user_id,
   decision_note = :note WHERE id = :id AND status = 'published'` —
   if 0 rows affected (already closed by another trigger), return
   `409`, not a silent success.
4. Log to `campaign_logs`.
5. Return `200` with the updated `Campaign`.

### Deadline scheduler (background, no endpoint)
1. Periodically: `UPDATE campaigns SET status = 'closed', closed_at =
   now(), closed_reason = 'deadline_reached' WHERE status =
   'published' AND deadline <= now()`.

### `max_amount` trigger (hook, no endpoint — lives in `donation`
domain's write path, forward reference)
1. As part of the successful-donation transaction (not yet spec'd),
   once `collected_amount >= max_amount`: same conditional update,
   `closed_reason = 'max_amount_reached'`, same transaction as the
   donation that crossed the threshold.

## Validation & error cases

| Case | Response |
|---|---|
| No/invalid bearer token | `401` |
| Non-Admin force-close attempt | `403` |
| Empty `decision_note` | `422` |
| Campaign already closed (by any trigger) when force-close is attempted | `409` |

## Concurrency & correctness notes

- All three triggers share the identical `WHERE status = 'published'`
  guard — this is the entire correctness mechanism for
  INV-campaign-13. No additional locking needed; whichever `UPDATE`
  commits first wins, the others affect 0 rows and return their
  respective "not published" error.
- **Explicit 3-way concurrency test required** (per `threat-model.md`):
  simulate all three triggers firing within the same short window on
  one campaign — assert exactly one `closed_reason` is recorded, the
  other two get clean `409`s (force-close) or silently no-op
  (scheduler/donation-trigger, which aren't user-facing requests with
  a response to check — verify via a post-condition assertion instead
  of an error response for those two).

## Test checklist

- [ ] Non-Admin force-close attempt → `403`.
- [ ] Force-close with empty `decision_note` → `422`.
- [ ] Force-close on an already-closed campaign → `409`.
- [ ] Deadline scheduler run twice near-simultaneously: exactly one
      closure, no error on the second run.
- [ ] **3-way race**: simulate max_amount-trigger, deadline-trigger,
      and force-close firing near-simultaneously — exactly one
      `closed_reason` recorded, campaign ends in `status = 'closed'`
      exactly once, no crash.
- [ ] `closed_by` populated only for `admin_force_closed`, `null` for
      the other two reasons.

## References

- `docs/spec/campaign/invariants.md` — INV-campaign-13
- `docs/spec/campaign/threat-model.md` — "Force-close" section
  (3-way race note)
- `docs/spec/campaign/tasks.md` — Task 09
- `api/openapi/campaign.yaml` — `POST /campaigns/{id}/force-close`
- `docs/project/kencleng-phase2-detail.md` — Fitur 3
