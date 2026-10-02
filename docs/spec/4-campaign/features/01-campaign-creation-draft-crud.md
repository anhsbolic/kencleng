# Feature Spec — 01: Campaign Creation & Draft CRUD

> File: `docs/spec/4-campaign/features/01-campaign-creation-draft-crud.md`
> Domain: `campaign`
> Task: 01 (see `docs/spec/campaign/tasks.md`)
> Status: draft — authored against `api/openapi/campaign.yaml` 2026-08-20
> Last updated: 2026-10-02
> WU-S2-006 source amendment accepted by Anhar Solehudin on 2026-10-02 after independent Review; acceptance covers this source amendment only.

## Summary

`POST /organizations/{organizationId}/campaigns` (create),
`PATCH`/`DELETE /campaigns/{campaignId}` (edit/delete, only while
`status = draft`). Any representative (owner or staff) of a
`verified`, non-overdue organization may create and edit drafts.

## Endpoints

`POST /organizations/{organizationId}/campaigns`, `PATCH
/campaigns/{campaignId}`, `DELETE /campaigns/{campaignId}` (confirmed,
`api/openapi/campaign.yaml`)

## Auth

`bearerAuth` required, representative (any `level`) of the owning
organization.

## Request

### Create — `CampaignCreateRequest`
| Field | Type | Required | Notes |
|---|---|---|---|
| `title` | string | Yes | |
| `description` | string | No | |
| `category` | enum | Yes | `bencana_alam`/`kesehatan`/`pendidikan`/`sosial`/`lainnya` |
| `location` | string | No | Free-text |
| `beneficiary_description` | string | No | Free-text |
| `target_amount` | decimal string | Yes | `> 0` |
| `max_amount` | decimal string, nullable | No | `≥ target_amount` if set |
| `max_donation_amount` | IDR whole-Rupiah monetary value using the shared major-unit decimal-string and explicit-currency convention | No | Per-Campaign cap; Rp5.000–Rp1.000.000.000 inclusive; omission defaults to Rp1.000.000.000 |
| `deadline` | datetime | Yes | Must be in the future |

### Edit — `CampaignUpdateRequest`
Same fields, all optional (partial update), only while `status =
draft`. If `max_donation_amount` is omitted, retain the Campaign's current
value; if supplied, it must be within the inclusive range and the Campaign
must still be a draft. An unrelated PATCH from an older client must not reset
the stored cap. The cap cannot be changed after publication.

## Behavior

### Create
1. Resolve caller's representative row for `organizationId` — reject
   (`403`) if none.
2. Load the organization — reject (`409 organization-not-verified`) if
   `status != 'verified'`; reject (`409 overdue-report`) if
   `has_overdue_report = true`. Read fresh, not cached.
3. Validate fields (INV-campaign-02): `target_amount > 0`, `max_amount
   ≥ target_amount` if set, `deadline` in the future, `category`
   present and valid, and supplied `max_donation_amount` within its
   per-Campaign whole-IDR range. Apply the default when the field is omitted.
4. Insert `campaigns` (`status = 'draft'`, `created_by =
   current_user_id`).
5. Return `201` with the `Campaign`.

### Edit
1. Resolve caller's representative row — `403` if none.
2. Load the campaign — reject (`409`) if `status != 'draft'`.
3. Apply field validation (same rules as create, for any field
   present). Preserve the stored `max_donation_amount` when omitted.
4. Update. Return `200` with the updated `Campaign`.

### Delete
1. Resolve caller's representative row — `403` if none.
2. Load the campaign — reject (`409`) if `status != 'draft'`.
3. Delete. Return `204`.

## Validation & error cases

| Case | Response |
|---|---|
| No/invalid bearer token | `401` |
| Caller isn't a representative of the organization | `403` |
| Organization not `verified` | `409 organization-not-verified` |
| Organization `has_overdue_report = true` | `409 overdue-report` |
| `target_amount ≤ 0`, `max_amount < target_amount`, past `deadline`, missing `category` | `422` |
| Supplied `max_donation_amount` below Rp5.000 or above Rp1.000.000.000, or not whole IDR | `422` |
| Edit/delete outside `status = draft` | `409` |

## Concurrency & correctness notes

- **Multi-editor drafts**: last-write-wins is sufficient for v1 — no
  optimistic locking (`kencleng-phase1-detail.md` Fitur 3, explicit).
- Organization-state checks (verified, overdue) at creation must read
  fresh state, not a cached/stale organization record.
- Campaign responses include the effective `max_donation_amount`. Existing
  Campaign rows receive the Rp1.000.000.000 default before cap-dependent
  behavior is enabled; exact persisted shape and migration/application are
  delivery-owned and not accepted by this spec amendment.

## Test checklist

- [ ] Creation against unverified org → `409 organization-not-verified`.
- [ ] Creation against a verified-but-overdue-report org → `409
      overdue-report`.
- [ ] Each field-validation rule individually → `422`.
- [ ] Create omission defaults the cap to Rp1.000.000.000; the inclusive
      Rp5.000 and Rp1.000.000.000 boundaries are accepted and out-of-range /
      fractional values are rejected.
- [ ] PATCH omission preserves an existing non-default cap, including when
      an older client updates another field; a supplied cap is draft-only.
- [ ] Existing-row backfill is complete before cap-dependent behavior is
      enabled; published cap is immutable.
- [ ] `staff` can create/edit/delete a draft.
- [ ] Edit/delete outside `draft` → `409`, for each non-draft status.
- [ ] Two concurrent edits to the same draft: last write wins, no
      error either way.
- [ ] An organization that later loses `verified` status doesn't
      retroactively affect an already-created campaign's editability.

## References

- `docs/spec/4-campaign/invariants.md` — INV-campaign-01, 02, 03, 04
- `docs/spec/4-campaign/threat-model.md` — "Campaign draft CRUD" section
- `docs/spec/4-campaign/tasks.md` — Task 01
- `docs/spec/organization/invariants.md` — INV-organization-13
  (referenced, not redefined)
- `api/openapi/campaign.yaml` — creation/edit/delete endpoints (authored API reconciliation remains a later WU-S2-006 stage)
