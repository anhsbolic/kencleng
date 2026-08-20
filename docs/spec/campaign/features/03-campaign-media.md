# Feature Spec — 03: Campaign Media

> File: `docs/spec/campaign/features/03-campaign-media.md`
> Domain: `campaign`
> Task: 03 (see `docs/spec/campaign/tasks.md`)
> Status: draft — includes the 2026-08-20 visibility-gating decision
> Last updated: 2026-08-20

## Summary

`GET`/`POST /campaigns/{campaignId}/attachments` — campaign media
(images), public bucket. Upload is owner/staff representative only.
**List visibility is gated to match the campaign detail endpoint**
(`[RESOLVED — 2026-08-20]`) — public only when the parent campaign's
`status = 'published'`, otherwise representative/Kurator/Admin only.
This **changes** `api/openapi/campaign.yaml`'s current `security: []`
on the list endpoint — that needs updating at implementation time to
add the same gating as `GET /campaigns/{campaignId}` (Task 02).

## Endpoints

`GET`/`POST /campaigns/{campaignId}/attachments` (confirmed paths in
`api/openapi/campaign.yaml`; auth on `GET` needs the change described
above)

## Auth

- `GET`: public if the parent campaign's `status = 'published'`;
  otherwise `bearerAuth` + representative/Kurator/Admin, `403`
  otherwise — identical rule to Task 02's detail endpoint.
- `POST`: `bearerAuth` + owner/staff representative of the owning
  organization.

## Request

### `POST` — `UploadCampaignMediaRequest`
`multipart/form-data`:

| Field | Type | Required | Notes |
|---|---|---|---|
| `file` | binary | Yes | JPG/PNG only, 5 MB max |

## Behavior

### List
1. Load the campaign — `404` if it doesn't exist.
2. If `status = 'published'`: return `CampaignAttachment[]` to anyone.
3. Otherwise: resolve caller's relationship (representative/Kurator/
   Admin) — `403` if none.
4. Each item includes a direct public `url` (no signed URL needed —
   public bucket, per `CampaignAttachment`'s schema note) once access
   is granted.

### Upload
1. Reject (`403`) if caller isn't a representative (owner or staff) of
   the owning organization.
2. Validate file type (JPG/PNG) and size (≤ 5 MB) — `422` otherwise.
3. Store in the public bucket, insert `campaign_attachments`.
4. Return `201` with the `CampaignAttachment`.

## Validation & error cases

| Case | Response |
|---|---|
| No/invalid bearer token (on gated list, or on upload) | `401` |
| List, non-`published` campaign, caller lacks visibility | `403` |
| Upload, caller not a representative | `403` |
| Invalid file type/oversized file | `422` |
| Campaign doesn't exist | `404` |

## Concurrency & correctness notes

None specific — plain file upload/list, no cross-row locking needed.

## Test checklist

- [ ] List: `published` campaign → readable without auth.
- [ ] List: non-`published` campaign — representative/Kurator/Admin
      succeed, unrelated/unauthenticated caller → `403` (same test
      matrix as Task 02's detail endpoint).
- [ ] Upload: non-representative → `403`.
- [ ] Upload: invalid file type/oversized → `422`.
- [ ] Uploaded media's `url` is directly publicly fetchable (no signed
      URL, unlike `organization` domain's legal documents) once list
      access is granted.

## References

- `docs/spec/campaign/invariants.md` — INV-campaign-14 (resolved
  2026-08-20)
- `docs/spec/campaign/threat-model.md` — "Campaign media" section
- `docs/spec/campaign/tasks.md` — Task 03
- `api/openapi/campaign.yaml` — attachment endpoints (list `security`
  needs updating per the decision above)
