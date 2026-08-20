# Feature Spec — 03: Public Donor List

> File: `docs/spec/donation/features/03-public-donor-list.md`
> Domain: `donation`
> Task: 03 (see `docs/spec/donation/tasks.md`)
> Status: draft — authored against `api/openapi/donation.yaml` 2026-08-20
> Last updated: 2026-08-20

## Summary

`GET /campaigns/{campaignId}/donations` — public list of a campaign's
successful donations, with server-computed `display_name` and
`is_anonymous` handling. Never includes `guest_email`.

## Endpoint

`GET /campaigns/{campaignId}/donations` (confirmed,
`api/openapi/donation.yaml`)

## Auth

None (`security: []`).

## Request

`CursorParam`, `LimitParam`.

## Behavior

1. Query `donations WHERE campaign_id = :id AND status = 'success'`.
2. For each row, compute `display_name`:
   - If `is_anonymous = true`: `null`.
   - Else if `donor_user_id` is set: the registered user's `User.name`.
   - Else if `guest_name` is set (non-empty): `guest_name` as-is.
   - Else: a generic placeholder label (exact copy is a frontend
     concern per `kencleng-phase2-detail.md`; backend returns
     `display_name: null` and the frontend supplies the generic label
     — **or** the backend returns a fixed generic string; pick one at
     implementation time and confirm which, since the schema
     description leaves this ambiguous — this is a minor open point,
     not blocking).
3. Return `DonationListResponse` (`id`, `display_name`, `amount`,
   `created_at` per item — **never** `guest_email`, `payment_method`,
   or any other field not in `DonationListItem`).

## Validation & error cases

| Case | Response |
|---|---|
| Campaign doesn't exist | `200` with an empty list (or `404` — pick one at implementation time; low-severity either way since this is a public, non-sensitive listing) |
| `limit` out of range | `422` |

## Concurrency & correctness notes

None specific — plain read.

## Test checklist

- [ ] Only `status = 'success'` donations appear (pending/failed
      excluded).
- [ ] `is_anonymous = true` → `display_name: null`, regardless of
      `guest_name` content.
- [ ] Registered donor → `display_name` is their `User.name`.
- [ ] Guest with `guest_name` provided → shown as-is.
- [ ] Guest with no `guest_name` (and not anonymous) → generic label
      (confirm the exact mechanism per the open point above).
- [ ] `guest_email` never appears in any response item, for any
      donation, under any parameter combination.
- [ ] `payment_method` never appears (confirm the schema's
      `DonationListItem` shape is exactly followed, not accidentally
      widened).

## References

- `docs/spec/donation/invariants.md` — INV-donation-06
- `docs/spec/donation/threat-model.md` — "Public donor list" section
- `docs/spec/donation/tasks.md` — Task 03
- `api/openapi/donation.yaml` — `DonationListItem`,
  `DonationListResponse` schemas
