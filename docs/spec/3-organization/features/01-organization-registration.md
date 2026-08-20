# Feature Spec — 01: Organization Registration

> File: `docs/spec/organization/features/01-organization-registration.md`
> Domain: `organization`
> Task: 01 (see `docs/spec/organization/tasks.md`)
> Status: draft — reconciled against `api/openapi/organization.yaml` 2026-08-20
> Last updated: 2026-08-20

## Summary

`POST /organizations` — a logged-in, non-Admin user registers a new
organization. Creates `organizations` (`status =
pending_verification`) and the registrant's own owner
`organization_representatives` row atomically, in one
`multipart/form-data` submission (data + legal documents together, no
separate draft-then-attach step).

## Endpoint

`POST /organizations` (confirmed, `api/openapi/organization.yaml`)

## Auth

`bearerAuth` required. `403` if the caller holds `role = 'admin'`.

## Request

`multipart/form-data`, schema `OrganizationCreateRequest`:

| Field | Type | Required | Notes |
|---|---|---|---|
| `name` | string | Yes | Legal/identity field |
| `description` | string | No | Operational field |
| `contact` | string | No | Operational field |
| `npwp` | string | Yes | Pattern `^\d{2}\.\d{3}\.\d{3}\.\d-\d{3}\.\d{3}$`, format-only validation, no DJP lookup |
| `akta_notaris` | binary | Yes | |
| `sk_kemenkumham` | binary | Yes | |
| `izin_pub` | binary | No | Optional in v1 |

## Behavior

1. Reject if the caller holds `role = 'admin'` (`403`).
2. Validate `npwp` format (regex, plaintext, before encryption).
3. Validate each uploaded file: type/size (`422` if invalid — the
   "invalid file type or exceeds 5 MB max" rule is documented
   explicitly on the attachment-replace endpoint, Task 05; apply the
   same limits here at registration time).
4. Within a single transaction:
   a. `SELECT COUNT(*) ... WHERE user_id = current_user_id AND level =
      'owner' FOR UPDATE` — reject (`409 organization-limit-reached`)
      if already at 5.
   b. Encrypt `npwp` (AES-GCM), compute `npwp_hash`.
   c. Check `npwp_hash` uniqueness — reject (`409 npwp-taken`) if a
      match exists.
   d. Insert `organizations` (`status = 'pending_verification'`).
   e. Insert `organization_representatives` (`level = 'owner'`) — same
      transaction as (d).
   f. Insert `organization_attachments` for each uploaded file
      (private bucket).
5. Commit. Return `201` with the `Organization` object.

## Validation & error cases

| Case | Response |
|---|---|
| No/invalid bearer token | `401` |
| Caller holds `role = 'admin'` | `403` |
| `npwp` fails format regex, or file type/size invalid | `422` (`ValidationError`) |
| `npwp` already registered | `409`, `type: npwp-taken` |
| Caller already owns 5 organizations | `409`, `type: organization-limit-reached` |
| Missing `akta_notaris` or `sk_kemenkumham` | `422` |

## Concurrency & correctness notes

- The 5-organization check and the org/representative inserts must be
  in the same transaction with a locking read — two near-simultaneous
  registrations from a user at 4 owned organizations must not both
  succeed.
- Org insert and owner-representative insert commit together or not
  at all (no orphaned organization).

## Test checklist

- [ ] Admin-role caller rejected.
- [ ] Successful registration creates exactly one `organizations` row
      and one owner `organization_representatives` row, same
      transaction.
- [ ] Fault injection between the two inserts → full rollback.
- [ ] 6th registration attempt for a user at 5 → `409`.
- [ ] Two concurrent registrations from a user at 4 → exactly one
      succeeds.
- [ ] Duplicate NPWP → `409`.
- [ ] Malformed NPWP format → `422`, no insert attempted.
- [ ] Invalid file type/oversized file → `422`.
- [ ] Missing required legal documents → `422`.

## References

- `docs/spec/organization/invariants.md` — INV-organization-01, 02, 03,
  04
- `docs/spec/organization/threat-model.md` — "Organization
  registration" section
- `docs/spec/organization/tasks.md` — Task 01
- `api/openapi/organization.yaml` — `POST /organizations`
