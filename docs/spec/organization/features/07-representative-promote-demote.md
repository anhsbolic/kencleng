# Feature Spec — 07: Representative Promote/Demote

> File: `docs/spec/organization/features/07-representative-promote-demote.md`
> Domain: `organization`
> Task: 07 (see `docs/spec/organization/tasks.md`)
> Status: draft — **contains one remaining `[open]`** (5-org limit only), reconciled 2026-08-20
> Last updated: 2026-08-20

## Summary

`PATCH /organizations/{organizationId}/representatives/{representativeId}`
— an Owner promotes a `staff` representative to `owner`, or demotes an
`owner` (including themselves) to `staff`. Demote enforces the
≥1-owner guard. Promote rejects a target holding `role = 'admin'`
(`409 admin-cannot-be-representative`, `[RESOLVED — 2026-08-20]`).

## `[open]` — still unresolved, needs a decision before this task ships

No documented re-check of the 5-organization-owner cap on promote
(INV-organization-04's own open note — separate from the 5 items
decided 2026-08-20). This spec implements the check as a **proposed**
behavior below (reuse the same `409 organization-limit-reached` type
as registration), but needs explicit confirmation before
implementation, since it wasn't part of today's resolved batch.

## Endpoint

`PATCH /organizations/{organizationId}/representatives/{representativeId}`
(confirmed, `api/openapi/organization.yaml`)

## Auth

`bearerAuth` required, owner-level representative of this organization
only (any owner may promote/demote any representative, including
themselves).

## Request

`application/json`, schema `UpdateRepresentativeLevelRequest`:

| Field | Type | Required | Notes |
|---|---|---|---|
| `level` | string enum (`owner`/`staff`) | Yes | Target level |

## Behavior

1. Reject if caller isn't an owner-level representative (`403`).
2. Load the target row (must belong to this organization).
3. **Promote (`staff → owner`)**:
   - Reject if target holds `role = 'admin'` (`409`, `type:
     admin-cannot-be-representative`).
   - **Proposed, pending confirmation**: reject if promoting would put
     the target's owner-count at 6 (`409`, `type:
     organization-limit-reached`, same type as registration's cap
     error).
   - Update `level = 'owner'`.
4. **Demote (`owner → staff`)**:
   - `SELECT COUNT(*) ... WHERE organization_id = :id AND level =
     'owner' FOR UPDATE`, same transaction as the update — reject
     (`409 last-owner`) if it would drop to 0.
   - Update `level = 'staff'`.
5. Log to `organization_logs`.

## Validation & error cases

| Case | Response |
|---|---|
| No/invalid bearer token | `401` |
| Caller is `staff` or not a representative | `403` |
| Target doesn't belong to this organization | `404` |
| Promote target holds `role = 'admin'` | `409 admin-cannot-be-representative` |
| Promote target already at the 5-org owner limit | `409` (proposed, pending confirmation — see "`[open]`" above) |
| Demote would leave 0 owners | `409 last-owner` |
| `level` equals current value (no-op) | `200`, harmless success |

## Concurrency & correctness notes

- Demote's `COUNT(*) ... FOR UPDATE` guard must be in the same
  transaction as the `UPDATE`.
- Consider sharing this guard's implementation with Task 08's remove
  path rather than duplicating it.

## Test checklist

- [ ] `staff` caller rejected.
- [ ] Promote a `role = 'admin'` user → `409
      admin-cannot-be-representative`.
- [ ] Promote a user already at 5 owned organizations → `409` (pending
      confirmation of the "`[open]`" item above).
- [ ] Demoting the last remaining owner → `409`.
- [ ] Two concurrent demotes against the last two owners of a 2-owner
      org: at most one succeeds.
- [ ] Self-demote allowed as long as ≥1 other owner remains.
- [ ] No-op request (`level` unchanged) succeeds without side effects.

## References

- `docs/spec/organization/invariants.md` — INV-organization-02, 04
  `[open]`, 05 (resolved 2026-08-20)
- `docs/spec/organization/threat-model.md` — "Representative
  promote/demote/remove" section
- `docs/spec/organization/tasks.md` — Task 07
- `api/openapi/organization.yaml` — `PATCH .../representatives/{id}`
