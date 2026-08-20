# Feature Spec — 01: Campaign Creation & Draft CRUD

> File: `docs/spec/campaign/features/01-campaign-creation-draft-crud.md`
> Domain: `campaign`
> Task: 01 (see `docs/spec/campaign/tasks.md`)
> Status: draft — authored against `api/openapi/campaign.yaml` 2026-08-20
> Last updated: 2026-08-20

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
| `deadline` | datetime | Yes | Must be in the future |

### Edit — `CampaignUpdateRequest`
Same fields, all optional (partial update), only while `status =
draft`.

## Behavior

### Create
1. Resolve caller's representative row for `organizationId` — reject
   (`403`) if none.
2. Load the organization — reject (`409 organization-not-verified`) if
   `status != 'verified'`; reject (`409 overdue-report`) if
   `has_overdue_report = true`. Read fresh, not cached.
3. Validate fields (INV-campaign-02): `target_amount > 0`, `max_amount
   ≥ target_amount` if set, `deadline` in the future, `category`
   present and valid.
4. Insert `campaigns` (`status = 'draft'`, `created_by =
   current_user_id`).
5. Return `201` with the `Campaign`.

### Edit
1. Resolve caller's representative row — `403` if none.
2. Load the campaign — reject (`409`) if `status != 'draft'`.
3. Apply field validation (same rules as create, for any field
   present).
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
| Edit/delete outside `status = draft` | `409` |

## Concurrency & correctness notes

- **Multi-editor drafts**: last-write-wins is sufficient for v1 — no
  optimistic locking (`kencleng-phase1-detail.md` Fitur 3, explicit).
- Organization-state checks (verified, overdue) at creation must read
  fresh state, not a cached/stale organization record.

## Test checklist

- [ ] Creation against unverified org → `409 organization-not-verified`.
- [ ] Creation against a verified-but-overdue-report org → `409
      overdue-report`.
- [ ] Each field-validation rule individually → `422`.
- [ ] `staff` can create/edit/delete a draft.
- [ ] Edit/delete outside `draft` → `409`, for each non-draft status.
- [ ] Two concurrent edits to the same draft: last write wins, no
      error either way.
- [ ] An organization that later loses `verified` status doesn't
      retroactively affect an already-created campaign's editability.

## References

- `docs/spec/campaign/invariants.md` — INV-campaign-01, 02, 03, 04
- `docs/spec/campaign/threat-model.md` — "Campaign draft CRUD" section
- `docs/spec/campaign/tasks.md` — Task 01
- `docs/spec/organization/invariants.md` — INV-organization-13
  (referenced, not redefined)
- `api/openapi/campaign.yaml` — creation/edit/delete endpoints
