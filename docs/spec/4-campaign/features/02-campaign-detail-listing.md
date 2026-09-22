# Feature Spec — 02: Campaign Detail & Listing

> File: `docs/spec/campaign/features/02-campaign-detail-listing.md`
> Domain: `campaign`
> Task: 02 (see `docs/spec/campaign/tasks.md`)
> Status: reconciled for Slice 1 — Public Campaign Understanding
> Last updated: 2026-09-22

## Summary

Slice 1 activates only the public composite detail endpoint. It enables a
visitor to understand one persisted eligible Campaign without exposing
internal lifecycle, organization, donor, or operational data.

## Endpoints

`GET /campaigns/{campaignId}` — `getPublicCampaignDetail`

## Auth

- `security: []`; the operation is public-only.
- Only the internal `published` fundraising state is public in Slice 1.
- An absent, malformed/non-resolvable, or non-public Campaign returns the
  same `PublicCampaignNotFound` `404` Problem Details response.
- Optional `Authorization` must not enrich, alter, or reveal a different
  response for any caller.

## Request

### `GET /campaigns/{campaignId}`

`campaignId` is a UUID path parameter. No authenticated/privileged variant
is selected by this endpoint.

## Behavior

### Detail
1. Resolve the Campaign and evaluate public eligibility before mapping.
2. Return identical `404` Problem Details for missing, invalid, and
   ineligible resources; do not distinguish status/body/header by optional
   authentication.
3. For an eligible Campaign, return only the required closed
   `PublicCampaignDetail` projection: `id`, `title`, organizer-owned plain
   text `purpose`/`story`, narrow `steward`, public `lifecycle`, tagged
   `funding`, tagged `media`, and unavailable `donation_action`.
4. `purpose` and `story` are plain strings with `source: organizer`; API
   data neither carries HTML/Markdown nor claims platform verification.
5. Funding amounts and computed percentage are decimal strings. A factual
   zero remains available; percentage is uncapped and relationship preserves
   below/at/above target or not-computable truth. `donor_count` is excluded.
6. `donation_action` is required but unavailable and contains no link or
   activation target.
7. Every success/error response uses `Cache-Control: private, no-store`.

## Validation & error cases

| Case | Response |
|---|---|
| Absent, malformed/non-resolvable, or non-public Campaign | identical `404` `PublicCampaignNotFound` |
| Eligible Campaign whose required detail/funding dependency cannot serve | `503` `PublicCampaignUnavailable` |

## Concurrency & correctness notes

- Backend owns eligibility, exact projection mapping, decimal progress,
  and availability semantics. Frontend may format values but must not
  recalculate business meaning.
- `private, no-store` is a visibility-withdrawal requirement, not a cache
  optimization choice. Timing parity and runtime negative tests belong to
  downstream implementation/testing work.

## Test checklist

- [ ] Eligible published fundraising Campaign returns the exact public
      projection regardless of optional Authorization.
- [ ] Absent, malformed, and non-public IDs have identical `404` status,
      body, and cache header; downstream testing also checks timing parity.
- [ ] Mapping has no inherited internal Campaign/Organization fields,
      `donor_count`, raw status/reasons, or operational metadata.
- [ ] Plain organizer text is rendered safely downstream and carries only
      machine-readable organizer provenance.
- [ ] Decimal funding preserves zero and over-target values without float
      conversion or client-owned calculation.
- [ ] Success and public errors specify `Cache-Control: private, no-store`.

## References

- `docs/spec/campaign/invariants.md` — INV-campaign-14
- `docs/spec/campaign/threat-model.md` — "Public campaign listing &
  detail" section
- `docs/spec/campaign/tasks.md` — Task 02
- `api/openapi/campaign.yaml` — `PublicCampaignDetail` and related public schemas

## Deferred historical breadth

`GET /campaigns`, `GET /organizations/{organizationId}/campaigns`, and any
authenticated/non-public detail semantics are `DEFER`. They remain useful
historical evidence but are not part of Slice 1.
