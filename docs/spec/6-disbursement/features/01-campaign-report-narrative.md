# Feature Spec — 01: Campaign Report & Narrative

> File: `docs/spec/disbursement/features/01-campaign-report-narrative.md`
> Domain: `disbursement`
> Task: 01 (see `docs/spec/disbursement/tasks.md`)
> Status: draft — authored against `api/openapi/disbursement.yaml` 2026-08-20
> Last updated: 2026-08-20

## Summary

`GET /campaigns/{campaignId}/report` — public, permanent archive for
any `closed` campaign: auto-generated figures plus an optional Owner
narrative. `PATCH .../report-narrative` — Owner-only, ungated,
freely-editable Markdown narrative (≤ 5000 chars).

## `[SECURITY-CRITICAL — frontend coordination note]`

`report_narrative` is raw, unsanitized Markdown displayed on a
**permanent public page**. The backend's only responsibility is
storing it as text and enforcing the length limit — it does **not**
render or sanitize. The frontend **must** render through a sanitizing
pipeline (e.g. `react-markdown` + `rehype-sanitize`), never
`dangerouslySetInnerHTML` on manually-converted HTML. This is called
out here because a backend developer implementing this endpoint might
reasonably assume "just store the string" is the whole job — it is,
on the backend, but the stakes of the frontend getting this wrong are
severe (permanent, public stored XSS).

## Endpoints

`GET /campaigns/{campaignId}/report`, `PATCH
/campaigns/{campaignId}/report-narrative` (confirmed,
`api/openapi/disbursement.yaml`)

## Auth

- Report view: none (`security: []`) — public.
- Narrative edit: `bearerAuth` + owner-level representative of the
  campaign's organization only.

## Request

### Narrative edit — `UpdateReportNarrativeRequest`
| Field | Type | Required | Notes |
|---|---|---|---|
| `report_narrative` | string | Yes | Max 5000 chars, raw Markdown. Empty string clears it. |

## Behavior

### Report view
1. Load the campaign — `404` if it doesn't exist.
2. Reject (`409`) if `status != 'closed'`.
3. Compute: `collected_amount`, `unique_donor_count` (distinct
   successful donors, `donation` domain data), `duration_days`
   (campaign lifetime), `closed_reason`.
4. Return `CampaignReportSummary`, including `report_narrative`
   (nullable — `null` if the Owner never filled it in).

### Narrative edit
1. Reject (`403`) if caller isn't an owner-level representative of
   this campaign's organization.
2. Reject (`409`) if `status != 'closed'`.
3. Reject (`422`) if `report_narrative` exceeds 5000 characters.
4. Update `campaigns.report_narrative` (empty string stored as empty,
   or normalize to `null` — pick one convention and apply
   consistently, since the schema allows either reading).
5. Return `200` with the updated `CampaignReportSummary`.

## Validation & error cases

| Case | Response |
|---|---|
| Report view, campaign doesn't exist | `404` |
| Report view, campaign not `closed` | `409` |
| Narrative edit, caller is `staff` or not a representative | `403` |
| Narrative edit, campaign not `closed` | `409` |
| Narrative edit, over 5000 chars | `422` |

## Concurrency & correctness notes

- No optimistic locking needed — narrative edits are low-frequency,
  single-owner-driven, last-write-wins is fine (same posture as
  `campaign` domain's draft multi-editor note, though narrative is
  even lower-contention since it's owner-only).

## Test checklist

- [ ] Report view on a non-`closed` campaign → `409`.
- [ ] `staff` narrative edit attempt → `403`.
- [ ] Narrative edit before `closed` → `409`.
- [ ] Narrative over 5000 chars → `422`.
- [ ] Repeated narrative edits after the first all succeed.
- [ ] Empty-string narrative clears it (report view then shows `null`
      or empty, whichever convention was chosen — confirm
      consistency).
- [ ] Report view remains accessible indefinitely (no time-based
      expiry test needed, but confirm no such logic was accidentally
      added).

## References

- `docs/spec/disbursement/invariants.md` — INV-disbursement-01, 02
- `docs/spec/disbursement/threat-model.md` — "Campaign report &
  narrative" section (XSS flag)
- `docs/spec/disbursement/tasks.md` — Task 01
- `api/openapi/disbursement.yaml` — `CampaignReportSummary`,
  `UpdateReportNarrativeRequest` schemas
- `docs/project/kencleng-phase3-detail.md` — Fitur 1
