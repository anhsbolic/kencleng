# Feature Spec — 02: Organization Detail View

> File: `docs/spec/organization/features/02-organization-detail-view.md`
> Domain: `organization`
> Task: 02 (see `docs/spec/organization/tasks.md`)
> Status: draft — finalized 2026-08-20
> Last updated: 2026-08-20

## Summary

`GET /organizations/{organizationId}` — single organization's data.
**Broadly readable** (`[RESOLVED — 2026-08-20]`): any authenticated
user may fetch any organization's detail. `npwp` is server-side
gated: present only for owner-level representatives, Admin, or the
assigned/historical Kurator — omitted entirely (not masked) for
`staff` and for any requester with no qualifying relationship.

## Endpoint

`GET /organizations/{organizationId}` (confirmed,
`api/openapi/organization.yaml`)

## Auth

`bearerAuth` required (standard). No relationship check beyond that
for the record as a whole. `npwp` field visibility is gated
separately (see Behavior).

## Request

Path: `organizationId` (UUID).

## Behavior

1. Resolve `current_user_id` and, if any, the caller's relationship to
   the organization (representative row and its `level`, or
   assigned/historical Kurator, or Admin role).
2. Load the organization. `404` if it doesn't exist.
3. Build the response: `id`, `name`, `description`, `contact`,
   `status`, `has_overdue_report`, `my_level` (caller's own
   representative level, `null` if none), `created_at`, `updated_at`
   — always included, regardless of relationship.
4. Include `npwp` (decrypted) **only if** the caller is an
   `owner`-level representative of this organization, holds `role =
   'admin'`, or is the assigned/historical Kurator. Omit the field
   entirely otherwise (not a masked placeholder — genuinely absent
   from the JSON payload).

## Validation & error cases

| Case | Response |
|---|---|
| No/invalid bearer token | `401` |
| Organization id doesn't exist | `404` |

No `403` case — the endpoint itself never rejects based on
relationship; only the `npwp` field's presence varies.

## Concurrency & correctness notes

None specific — plain read.

## Test checklist

- [ ] Any authenticated user (no relationship to the org at all) can
      fetch the record — gets everything except `npwp`.
- [ ] `staff` representative: same as above — no `npwp` in the
      response.
- [ ] `owner` representative: `npwp` present.
- [ ] Admin: `npwp` present, for any organization.
- [ ] Assigned or historical Kurator: `npwp` present.
- [ ] Nonexistent organization id → `404`.
- [ ] `my_level` correctly reflects the caller's own relationship
      (`null` if none, `owner`/`staff` otherwise).

## References

- `docs/spec/organization/invariants.md` — INV-organization-10
  (resolved 2026-08-20)
- `docs/spec/organization/threat-model.md` — "Organization detail
  view" section
- `docs/spec/organization/tasks.md` — Task 02
- `api/openapi/organization.yaml` — `GET
  /organizations/{organizationId}`
