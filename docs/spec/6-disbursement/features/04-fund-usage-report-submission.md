# Feature Spec — 04: Fund-Usage Report Submission & Attachments

> File: `docs/spec/disbursement/features/04-fund-usage-report-submission.md`
> Domain: `disbursement`
> Task: 04 (see `docs/spec/disbursement/tasks.md`)
> Status: draft — authored against `api/openapi/disbursement.yaml` 2026-08-20
> Last updated: 2026-08-20

## Summary

`POST /disbursement-requests/{id}/fund-usage-reports` — Owner-only,
requires `status = 'disbursed'`, strict reconciliation (items must sum
to exactly `requested_amount`), clears `has_overdue_report` in the
same transaction if it was set. `GET` (list + detail with embedded
signed-URL attachments). `POST .../items/{itemId}/attachments` —
Owner-only, private bucket.

## Endpoints

`POST`/`GET
/disbursement-requests/{disbursementRequestId}/fund-usage-reports`,
`GET /fund-usage-reports/{reportId}`, `POST
/fund-usage-reports/{reportId}/items/{itemId}/attachments` (confirmed,
`api/openapi/disbursement.yaml`)

## Auth

`bearerAuth` required. Submission, attachment upload: owner-level
representative only. List, detail: representative/Kurator/Admin —
**assumed default**, confirm at implementation time.

## Request

### Submission — `CreateFundUsageReportRequest`
| Field | Type | Required | Notes |
|---|---|---|---|
| `items` | array, min 1 | Yes | Each: `category` (free-text), `amount`, `description` (optional) |

### Attachment upload — `UploadFundUsageAttachmentRequest`
`multipart/form-data`, `file` (PDF/JPG/PNG, 5 MB max).

## Behavior

### Submission
1. Reject (`403`) if caller isn't an owner-level representative of
   the campaign's organization (via the disbursement request →
   campaign → organization chain).
2. Reject (`409`) if `DisbursementRequest.status != 'disbursed'`.
3. Validate: sum of all `items[].amount` **must exactly equal**
   `requested_amount` — `422` on any mismatch, over or under, however
   small.
4. Within one transaction:
   a. Insert `fund_usage_reports` (`status =
      'pending_verification'`).
   b. Insert one `fund_usage_report_items` row per item.
   c. If the owning organization's `has_overdue_report = true`: clear
      it (`false`), log to `organization_logs` (see "Open item" in
      `threat-model.md` — this domain writes to `organization`'s log
      table, a cross-domain write worth calling out explicitly here
      too).
5. Best-effort notification to Admin (`type =
   admin_new_curation_item`).
6. Return `201` with the `FundUsageReport`.

### Attachment upload
1. Reject (`403`) if caller isn't an owner-level representative.
2. Validate file type (PDF/JPG/PNG) and size (≤ 5 MB) — `422`
   otherwise.
3. Store in the private bucket, insert
   `fund_usage_report_item_attachments`, linked to the specified item.
4. Return `201` with the `FundUsageReportItemAttachment` (no
   `download_url` on this response — that's generated when the report
   detail is fetched, per "List/Detail" below).

### List / Detail
1. Reject (`403`) if caller lacks a qualifying relationship.
2. List: return `FundUsageReportListResponse` — all reports for this
   disbursement, including past `rejected` ones (history).
3. Detail: return `FundUsageReport` with nested `items`, each item's
   `attachments` including a freshly-generated 5-minute signed
   `download_url`.

## Validation & error cases

| Case | Response |
|---|---|
| No/invalid bearer token | `401` |
| Submission/upload, caller is `staff` or not a representative | `403` |
| Submission, disbursement not `disbursed` | `409` |
| Submission, items don't sum to exactly `requested_amount` | `422` |
| Upload, invalid file type/oversized | `422` |
| List/detail, no qualifying relationship | `403` |

## Concurrency & correctness notes

- Steps 4a–4c (report creation + items + `has_overdue_report` clear)
  must be one transaction — no window where the report exists but the
  organization's flag is still `true` (INV-disbursement-06).
- Signed URLs are generated fresh on each detail fetch — no caching,
  same pattern as `organization`'s legal-document downloads.

## Test checklist

- [ ] `staff` submission/upload attempt → `403`.
- [ ] Submission before `disbursed` → `409`.
- [ ] Items summing to anything other than exactly `requested_amount`
      (test both slightly over and slightly under) → `422`.
- [ ] Submission on an org with `has_overdue_report = true` clears it,
      same transaction, no window where both are true.
- [ ] `has_overdue_report` clear logs to `organization_logs` (confirm
      the target table explicitly — this is the open item from
      `threat-model.md`).
- [ ] Invalid attachment file type/oversized → `422`.
- [ ] Detail's embedded `download_url` valid for exactly 5 minutes.
- [ ] List includes past `rejected` reports for this disbursement.

## References

- `docs/spec/disbursement/invariants.md` — INV-disbursement-06, 14
- `docs/spec/disbursement/threat-model.md` — "Fund-usage report
  submission & attachments" section
- `docs/spec/disbursement/tasks.md` — Task 04
- `docs/spec/organization/invariants.md` — INV-organization-13
  (fulfilled here)
- `api/openapi/disbursement.yaml` — fund-usage report endpoints
