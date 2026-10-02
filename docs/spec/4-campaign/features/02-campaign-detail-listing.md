# Feature Spec — 02: Campaign Detail & Listing

> File: `docs/spec/4-campaign/features/02-campaign-detail-listing.md`
> Domain: `campaign`
> Task: 02 (see `docs/spec/4-campaign/tasks.md`)
> Status: Slice 1 reconciled; Slice 2 action contract accepted (runtime delivery pending)
> Last updated: 2026-10-02
> WU-S2-006 source amendment accepted by Anhar Solehudin on 2026-10-02 after independent Review; acceptance covers this source amendment only.

## Summary

Slice 1 activates the public composite detail endpoint. Slice 2 reconciles its
donation action affordance with backend Donation eligibility. The detail lets
a visitor understand one persisted eligible Campaign without exposing
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
3. For a publicly eligible Campaign, return only the required closed
   `PublicCampaignDetail` projection: `id`, `title`, required
   `max_donation_amount`, organizer-owned plain text `purpose`/`story`,
   narrow `steward`, public `lifecycle`, tagged `funding`, tagged `media`,
   and `donation_action` as defined below. `max_donation_amount` is the
   effective per-Campaign whole-IDR cap, represented as a closed object
   `{amount, currency_code}`. `amount` is a major-unit decimal string and
   `currency_code` is explicitly `IDR`; both members are required. The donor
   can see it before entering an amount.
4. `purpose` and `story` are plain strings with `source: organizer`; API
   data neither carries HTML/Markdown nor claims platform verification.
5. Funding amounts and computed percentage are decimal strings. A factual
   zero remains available; percentage is uncapped and relationship preserves
   below/at/above target or not-computable truth. `donor_count` is excluded.
6. `donation_action` is required and is a closed union: an eligible Campaign
   has `{availability: available}`; a publicly eligible Campaign that fails
   the backend Donation submission-eligibility predicate at this GET snapshot
   has `{availability: unavailable, reason: campaign_not_eligible}`. The
   reason is generic and does not expose lifecycle or closure details. The
   frontend derives navigation from the Campaign ID; this field contains no
   link or activation target. GET availability is a point-in-time affordance;
   Donation submission independently reevaluates current eligibility and
   validates the amount against the current Campaign cap and remaining
   funding capacity. The disclosed detail value is not authorization or a
   reservation.
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
- In Slice 2, closed Campaigns remain non-public and return the same `404`
  behavior described above. D-04 is a Slice-3 source handoff only: when that
  slice is reconciled, its public closed result retains Campaign identity,
  removes the donation action, and does not present Funding as final while
  accepted Donations remain pending. This note does not activate closed
  detail visibility or settle Slice-3 projection fields in Slice 2.
- `private, no-store` is a visibility-withdrawal requirement, not a cache
  optimization choice. Timing parity and runtime negative tests belong to
  downstream implementation/testing work.

## Test checklist

- [ ] Publicly eligible Campaign returns the exact public projection and
      reports the backend Donation-eligibility predicate at the GET snapshot,
      regardless of optional Authorization; unavailable is emitted only when
      that predicate fails while the detail remains public.
- [ ] The effective per-Campaign `max_donation_amount` is present in public
      detail before amount entry. Submission rechecks the current cap and
      capacity independently; GET does not reserve either.
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

- `docs/spec/4-campaign/invariants.md` — INV-campaign-14
- `docs/spec/4-campaign/threat-model.md` — "Public campaign listing &
  detail" section
- `docs/spec/4-campaign/tasks.md` — Task 02
- `api/openapi/campaign.yaml` — `PublicCampaignDetail` and related public schemas

## Deferred historical breadth

`GET /campaigns`, `GET /organizations/{organizationId}/campaigns`, and any
authenticated/non-public detail semantics are `DEFER`. They remain useful
historical evidence but are not part of Slice 1.
