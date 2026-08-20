# Feature Spec — 05: Legal Document Attachments

> File: `docs/spec/organization/features/05-legal-document-attachments.md`
> Domain: `organization`
> Task: 05 (see `docs/spec/organization/tasks.md`)
> Status: draft — finalized 2026-08-20
> Last updated: 2026-08-20

## Summary

Three endpoints, all confirmed in `api/openapi/organization.yaml`:

1. `GET /organizations/{organizationId}/attachments` — metadata list.
2. `PUT /organizations/{organizationId}/attachments/{type}` —
   upload/replace (`type` ∈ `akta_notaris`/`sk_kemenkumham`/
   `izin_pub`, one current file per type — replacing overwrites, no
   version history).
3. `GET /organizations/{organizationId}/attachments/{type}` — returns
   a 5-minute signed download URL.

This replaces the 2026-08-19 draft's attachment-by-`{attachmentId}` +
direct-download design (which predated seeing the actual API).

## `izin_pub` re-curation — `[RESOLVED — 2026-08-20]`

Replacing `akta_notaris`/`sk_kemenkumham` requires `confirm=true` and
triggers re-curation; `izin_pub` does **not** — follows the
implementation as-is. `kencleng-phase1-detail.md` Fitur 1's listing of
Izin PUB as legal/identity-class is now the stale document and should
be updated to match.

## Endpoints

`GET .../attachments`, `PUT .../attachments/{type}`, `GET
.../attachments/{type}`

## Auth

- List, download (signed URL): Owner-level representative, Admin, or
  Kurator (current or historical assignment). `staff` rejected (`403`,
  confirmed explicit on the list endpoint).
- Upload/replace: **owner-only** (confirmed explicit: "Only
  `owner`-level representatives may manage legal documents").

## Request

### `PUT .../attachments/{type}`
`multipart/form-data`, schema `ReplaceAttachmentRequest`:

| Field | Type | Required | Notes |
|---|---|---|---|
| `file` | binary | Yes | |
| `confirm` | boolean | Conditionally | Required (`true`) if `type` is `akta_notaris`/`sk_kemenkumham` |

## Behavior

### List
1. Resolve caller's relationship (owner rep / Admin / Kurator).
2. Reject (`403`) if `staff` or unrelated.
3. Return `OrganizationAttachment[]` — metadata only, no
   `download_url`.

### Upload/replace
1. Reject if caller isn't an owner-level representative.
2. Validate file type/size (`422` if invalid, 5 MB max).
3. If `type` ∈ `{akta_notaris, sk_kemenkumham}` and `confirm != true`:
   reject (`409`), no file stored.
4. Store the file (private bucket), upsert the
   `organization_attachments` row for this `(organization_id, type)`
   — overwrite, not append.
5. If `type` ∈ `{akta_notaris, sk_kemenkumham}` (and thus `confirm ==
   true`): same status-flip + auto-unpublish behavior as Task 04 step
   6, same transaction as the attachment upsert.
6. Return `200` with the `OrganizationAttachment` (no `download_url`
   on this response — that's the separate signed-URL endpoint).

### Signed download URL
1. Resolve caller's relationship (owner rep / Admin / Kurator).
2. Reject (`403`) if `staff` or unrelated; `404` if the organization
   or that attachment type doesn't exist.
3. Generate a signed URL, 5-minute expiry, for the file in private
   storage.
4. Return `200` with `OrganizationAttachment` including `download_url`.

## Validation & error cases

| Case | Response |
|---|---|
| No/invalid bearer token | `401` |
| `staff` or unrelated caller on list/download | `403` |
| Non-owner caller on upload/replace | `403` |
| Invalid file type/size | `422` |
| Legal-class type replaced without `confirm=true` | `409` |
| Attachment type doesn't exist yet for this org (download before any upload) | `404` |

## Concurrency & correctness notes

- Upload/replace + status-flip (when applicable) must be one
  transaction, same reasoning as Task 04.
- Signed URL generation itself needs no locking — it's a read against
  already-stored file metadata.
- **Signed-URL exposure window**: the URL is a bearer credential for
  5 minutes once issued — confirm it isn't logged in plaintext
  anywhere server-side (access logs, error tracking), which would
  extend the effective exposure window (see `threat-model.md`).

## Test checklist

- [ ] `staff` rejected on list and download.
- [ ] Non-owner rejected on upload/replace.
- [ ] Owner, Admin, and assigned/historical Kurator all succeed on
      list and download.
- [ ] `akta_notaris`/`sk_kemenkumham` replace without `confirm=true` →
      `409`, no file stored, no status change.
- [ ] `akta_notaris`/`sk_kemenkumham` replace with `confirm=true` on a
      `verified` org → status flips, campaigns auto-unpublish (test
      double).
- [ ] `izin_pub` replace never requires `confirm` and never changes
      `status`.
- [ ] Invalid file type/oversized file → `422`.
- [ ] Signed URL expires after 5 minutes (test against the storage
      layer's actual TTL enforcement, not just the API response).
- [ ] Uploading a second file for the same `type` overwrites the
      previous one (no history kept).

## References

- `docs/spec/organization/invariants.md` — INV-organization-08
  `[NEEDS DECISION]`, 10, 11
- `docs/spec/organization/threat-model.md` — "Legal document
  attachments" section (signed-URL leakage risk)
- `docs/spec/organization/tasks.md` — Task 05
- `api/openapi/organization.yaml` — attachment endpoints
- `docs/project/kencleng-phase0-detail.md` — Fitur 7 (private bucket)
