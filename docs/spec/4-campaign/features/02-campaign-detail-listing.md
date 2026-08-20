# Feature Spec — 02: Campaign Detail & Listing

> File: `docs/spec/campaign/features/02-campaign-detail-listing.md`
> Domain: `campaign`
> Task: 02 (see `docs/spec/campaign/tasks.md`)
> Status: draft — authored against `api/openapi/campaign.yaml` 2026-08-20
> Last updated: 2026-08-20

## Summary

Three endpoints: public browse listing, an org-scoped dashboard
listing, and the composite detail endpoint (campaign + organization
summary + progress in one call — deliberate composite, per
`kencleng-backend-tech-stack.md`'s design philosophy, to avoid
client-side N+1 on the platform's primary conversion surface).

## Endpoints

`GET /campaigns`, `GET /organizations/{organizationId}/campaigns`,
`GET /campaigns/{campaignId}` (confirmed, `api/openapi/campaign.yaml`)

## Auth

- `GET /campaigns`: none (`security: []`) — public.
- `GET /organizations/{organizationId}/campaigns`: `bearerAuth` +
  representative of that organization.
- `GET /campaigns/{campaignId}`: none required, but response varies —
  public if `status = 'published'`; otherwise `bearerAuth` +
  representative/Kurator/Admin, `403` otherwise (INV-campaign-14).

## Request

### `GET /campaigns`
| Param | Notes |
|---|---|
| `category` | Optional filter |
| `q` | Optional free-text search against `title` (`# INFERRED`) |
| `cursor`, `limit` | Standard pagination |

### `GET /organizations/{organizationId}/campaigns`
| Param | Notes |
|---|---|
| `status` | Optional filter, any `CampaignStatus` |
| `cursor`, `limit` | Standard pagination |

## Behavior

### Public listing
1. Query `campaigns WHERE status = 'published'`, optionally filtered
   by `category`/`q`.
2. Return `CampaignListResponse` (each item includes `progress`).

### Org-scoped listing
1. Reject (`403`) if caller isn't a representative of the
   organization.
2. Query `campaigns WHERE organization_id = :id`, optionally filtered
   by `status` — **all** statuses, unlike the public listing.
3. Return `CampaignListResponse`.

### Detail
1. Load the campaign — `404` if it doesn't exist.
2. If `status = 'published'`: return `CampaignDetail` to anyone.
3. Otherwise: resolve caller's relationship (representative of the
   owning organization, assigned/historical Kurator, or Admin) — `403`
   if none. Return `CampaignDetail` if a relationship exists.
4. `CampaignDetail` = `Campaign` fields + `organization`
   (`OrganizationSummary` — minimal projection: `id`, `name`, `status`
   only, **not** the full `Organization` resource, so this doesn't
   leak `npwp` or anything gated in `organization` domain) +
   `progress` (`percentage` capped at 100 for display,
   `donor_count`, `days_remaining` — nullable once `closed`).

## Validation & error cases

| Case | Response |
|---|---|
| Org-scoped listing, caller not a representative | `403` |
| Detail, campaign doesn't exist | `404` |
| Detail, non-`published` campaign, caller lacks visibility | `403` (confirmed — see `threat-model.md`'s note on this being existence-confirming, accepted) |
| `limit` out of range | `422` |

## Concurrency & correctness notes

- `CampaignProgress.percentage` and `donor_count` are computed from
  `donations` (a `donation`-domain table, not yet spec'd) — this
  domain's responsibility is just to expose them correctly in the
  composite response; the underlying aggregation's correctness belongs
  to `donation` domain when it's built.
- No caching layer assumed for this composite endpoint in v1 — always
  a fresh read. Flag as a future optimization if the campaign detail
  page becomes a hot path (it's explicitly called out as the
  platform's primary conversion surface).

## Test checklist

- [ ] Public listing returns only `published` campaigns, `category`/
      `q` filters work.
- [ ] Org-scoped listing returns all statuses; non-representative
      caller → `403`.
- [ ] Detail: `published` campaign readable without auth.
- [ ] Detail: non-`published` campaign — representative, Kurator
      (assigned or historical), Admin all succeed; unrelated/
      unauthenticated caller → `403`.
- [ ] `CampaignDetail.organization` never includes `npwp` or any
      owner/Admin-only `organization` field — only the minimal
      `OrganizationSummary` projection.
- [ ] Nonexistent campaign id → `404`.

## References

- `docs/spec/campaign/invariants.md` — INV-campaign-14
- `docs/spec/campaign/threat-model.md` — "Public campaign listing &
  detail" section
- `docs/spec/campaign/tasks.md` — Task 02
- `api/openapi/campaign.yaml` — `CampaignDetail`, `OrganizationSummary`,
  `CampaignProgress` schemas
