# Feature Spec — 06: Representative Invite

> File: `docs/spec/organization/features/06-representative-invite.md`
> Domain: `organization`
> Task: 06 (see `docs/spec/organization/tasks.md`)
> Status: draft — finalized 2026-08-20
> Last updated: 2026-08-20

## Summary

`POST /organizations/{organizationId}/representatives` — an Owner
directly adds an existing, email-verified user as a `staff`
representative. Direct-add, no accept/consent step. Rejects a target
holding `role = 'admin'` (`409 admin-cannot-be-representative`,
`[RESOLVED — 2026-08-20]`).

## Endpoint

`POST /organizations/{organizationId}/representatives` (confirmed,
`api/openapi/organization.yaml`)

## Auth

`bearerAuth` required, owner-level representative of this organization
only (confirmed explicit `403`).

## Request

`application/json`, schema `InviteRepresentativeRequest`:

| Field | Type | Required | Notes |
|---|---|---|---|
| `email` | string | Yes | Must belong to an existing, email-verified user |

## Behavior

1. Reject if caller isn't an owner-level representative (`403`).
2. Look up a `users` row by `email`.
3. If no such user, or the user's email isn't verified: reject (`404
   user-not-found`, confirmed explicit — single collapsed message for
   both cases, per the 2026-08-19 accepted-risk decision on
   enumeration).
4. If the target already holds `role = 'admin'`: reject (`409`,
   `type: admin-cannot-be-representative`).
5. If the target is already a representative of this organization:
   reject (`409`, confirmed explicit, no specific `type` shown in the
   spec yet — implement with a clear `detail` message either way).
6. Insert `organization_representatives` (`level = 'staff'`).
7. Best-effort notification to the added user (INV-notification-06 —
   must not fail this endpoint's own transaction).
8. Log to `organization_logs`.
9. Return `201` with the `Representative` object.

## Validation & error cases

| Case | Response |
|---|---|
| No/invalid bearer token | `401` |
| Caller is `staff` or not a representative | `403` |
| No user with that email, or email unverified | `404 user-not-found` (confirmed, single message) |
| Target holds `role = 'admin'` | `409`, `type: admin-cannot-be-representative` |
| Target already a representative | `409` (confirmed, exact `type` TBD) |

## Concurrency & correctness notes

- Two concurrent invites for the same email/org: the `UNIQUE (user_id,
  organization_id)` constraint on `organization_representatives`
  makes the second fail cleanly at the DB level.

## Test checklist

- [ ] `staff` caller rejected.
- [ ] Nonexistent email → `404`, documented message.
- [ ] Existing-but-unverified email → **same** `404` message as
      nonexistent (verify this explicitly — easy to accidentally
      implement 3 distinct messages instead of the intended 2).
- [ ] Email belonging to a `role = 'admin'` user → `409
      admin-cannot-be-representative`.
- [ ] Email already representing this organization → `409`.
- [ ] Successful invite creates a `staff` row + best-effort
      notification (verify notification failure doesn't fail this
      endpoint).
- [ ] Two concurrent invites for the same email/org: exactly one
      succeeds.

## References

- `docs/spec/organization/invariants.md` — INV-organization-05
  (resolved 2026-08-20), 12
- `docs/spec/organization/threat-model.md` — "Representative invite"
  section (accepted-risk decision + resolved Admin-exclusion gap)
- `docs/spec/organization/tasks.md` — Task 06
- `api/openapi/organization.yaml` — `POST .../representatives`
