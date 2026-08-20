# Feature Spec — 05: Curation Assignment

> File: `docs/spec/campaign/features/05-curation-assignment.md`
> Domain: `campaign`
> Task: 05 (see `docs/spec/campaign/tasks.md`)
> Status: draft — authored against `api/openapi/campaign.yaml` 2026-08-20
> Last updated: 2026-08-20

## Summary

`POST /campaigns/{campaignId}/curation/assign` — Admin-only, assigns a
Kurator to review a `pending_curation` campaign. Same pattern as
`organization`'s equivalent (INV-campaign-05, INV-campaign-06,
mirroring INV-organization-06/07).

## Endpoints

`POST /campaigns/{campaignId}/curation/assign`, `GET
/campaigns/curation-queue` (confirmed, `api/openapi/campaign.yaml`)

## Auth

Both: `bearerAuth` + `role = 'admin'`.

## Request

### Assign — `AssignCuratorRequest` (shared schema, from
`organization.yaml`)
| Field | Type | Required | Notes |
|---|---|---|---|
| `kurator_id` | UUID | Yes | Must hold `role = 'kurator'` |

### Queue
| Param | Notes |
|---|---|
| `cursor`, `limit` | Standard pagination |

## Behavior

### Assign
1. Reject (`403`) if caller isn't Admin.
2. Reject (`409`) if `kurator_id` is a representative (any `level`) of
   the campaign's owning organization (INV-campaign-05).
3. Reject (`409`) if a `pending` assignment already exists for this
   campaign (INV-campaign-06, DB unique-partial-index-backed).
4. Insert `campaign_curation_assignments` (`decision = 'pending'`).
5. Best-effort notification to the assigned Kurator.
6. Log to `campaign_logs`.
7. Return `201` with the `CampaignCurationAssignment`.

### Queue
1. Reject (`403`) if caller isn't Admin.
2. Query `campaigns WHERE status = 'pending_curation' AND NOT EXISTS
   (campaign_curation_assignments WHERE decision = 'pending')`.
3. Return `CampaignListResponse`, paginated.

## Validation & error cases

| Case | Response |
|---|---|
| No/invalid bearer token | `401` |
| Non-Admin caller | `403` |
| `kurator_id` is a representative of the campaign's organization | `409` |
| Campaign already has a pending assignment | `409` |

## Concurrency & correctness notes

- Same as `organization`'s equivalent: the DB unique partial index is
  the actual correctness guarantee for the one-active-assignment rule
  — catch the constraint violation, translate to a clean `409`.

## Test checklist

- [ ] Non-Admin caller → `403` on both endpoints.
- [ ] `kurator_id` who is a representative of the campaign's
      organization → `409`.
- [ ] Two concurrent assignment attempts for the same campaign: at
      most one succeeds, clean `409` for the other.
- [ ] Queue excludes campaigns with an active assignment even if still
      `pending_curation`.

## References

- `docs/spec/campaign/invariants.md` — INV-campaign-05, 06
- `docs/spec/campaign/threat-model.md` — references
  `organization/threat-model.md`'s equivalent section
- `docs/spec/campaign/tasks.md` — Task 05
- `docs/spec/organization/features/09-curation-assignment.md` —
  structural precedent
- `api/openapi/campaign.yaml` — `POST .../curation/assign`, `GET
  /campaigns/curation-queue`
