# Feature Spec — 03: Campaign Media

> File: `docs/spec/campaign/features/03-campaign-media.md`
> Domain: `campaign`
> Task: 03 (see `docs/spec/campaign/tasks.md`)
> Status: reconciled for Slice 1 — Public Campaign Understanding
> Last updated: 2026-09-22

## Summary

Slice 1 exposes truthful public media metadata only within
`PublicCampaignDetail`, and delivers bytes through a controlled same-origin
operation. Campaign media remains private in object storage; no public
bucket, direct object URL, redirect, or long-lived signed URL is permitted.

## Endpoints

`GET /campaigns/{campaignId}/media/{mediaId}/content` —
`getPublicCampaignMediaContent`

## Auth

- `security: []`; this is a public-only controlled byte operation.
- Every origin request rechecks the parent Slice-1 public-eligibility
  predicate and that the media member belongs to that parent.
- Optional Authorization never enriches visibility or response shape.

## Request

`campaignId` and `mediaId` are UUID path parameters. Success delivers only
`image/jpeg` or `image/png` bytes.

## Behavior

1. Resolve the parent and media member, then recheck public eligibility.
2. Return JPEG/PNG bytes only when both predicates hold.
3. Return the same `PublicCampaignNotFound` `404` Problem Details response
   for absent/non-public parent or absent/non-member media, without revealing
   which condition failed.
4. Return `PublicCampaignUnavailable` `503` when eligible metadata exists
   but storage/dependency or the object cannot serve bytes; never turn that
   failure into `media.absent`.
5. Apply `Cache-Control: private, no-store` to `200`, `404`, and `503`.
   Retraction prevents new origin fetches through a known content URL;
   already downloaded client-held bytes are outside the guarantee.

## Validation & error cases

| Case | Response |
|---|---|
| Absent/non-public parent or absent/non-member media | identical `404` `PublicCampaignNotFound` |
| Eligible metadata but unavailable bytes/dependency | `503` `PublicCampaignUnavailable` |

## Concurrency & correctness notes

Visibility and membership must be checked at the delivery origin on every
request; a metadata check alone is insufficient for retraction.

## Test checklist

- [ ] Only JPEG/PNG bytes are returned; no object-storage URL/redirect is
      exposed.
- [ ] Missing/non-public/non-member requests have identical `404` behavior,
      including the no-store header.
- [ ] Storage/object failure produces `503`, not a false absence state.
- [ ] Retraction prevents a new origin fetch through a previously known URL.

## References

- `docs/spec/campaign/invariants.md` — INV-campaign-14
- `docs/spec/campaign/threat-model.md` — "Campaign media" section
- `docs/spec/campaign/tasks.md` — Task 03
- `api/openapi/campaign.yaml` — public media metadata/content schemas and operation

## Deferred historical breadth

`GET`/`POST /campaigns/{campaignId}/attachments`, upload validation,
curation, and operational media metadata remain `DEFER`. They must be
reconciled later and do not authorize a public-bucket implementation.
