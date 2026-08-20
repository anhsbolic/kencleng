# Feature Spec — 04: Organization Edit

> File: `docs/spec/organization/features/04-organization-edit.md`
> Domain: `organization`
> Task: 04 (see `docs/spec/organization/tasks.md`)
> Status: draft — finalized 2026-08-20
> Last updated: 2026-08-20

## Summary

`PATCH /organizations/{organizationId}` — edits organization fields.
Access is split by field class (`[RESOLVED — 2026-08-20]`): `staff`
representatives may edit `description`/`contact` only; `name`/`npwp`
(legal/identity, `confirm`-gated, re-curation-triggering) are
owner-only. A legal/identity change requires `confirm: true` and, if
the org is currently `verified`, flips `status` back to
`pending_verification`.

## Endpoint

`PATCH /organizations/{organizationId}` (confirmed,
`api/openapi/organization.yaml`)

## Auth

`bearerAuth` required, plus a field-class-aware check (not a single
blanket role gate):
- Caller must be a representative (any `level`) of this organization
  at minimum — `403` ("not a representative") otherwise.
- If the request body includes `name` or `npwp`: caller must
  additionally be `level = 'owner'` — `403` ("legal/identity fields
  are owner-only") otherwise, even for a `staff` representative.

## Request

`application/json`, schema `OrganizationUpdateRequest` (partial
update):

| Field | Type | Class | Notes |
|---|---|---|---|
| `name` | string | Legal/identity — owner-only | |
| `npwp` | string | Legal/identity — owner-only | Same format pattern as registration |
| `description` | string | Operational — any representative | |
| `contact` | string | Operational — any representative | |
| `confirm` | boolean | — | Required (`true`) if this request changes `name` or `npwp` |

## Behavior

1. Reject (`403`) if caller isn't a representative of this
   organization at all.
2. If the request touches `name` or `npwp`: reject (`403`) if caller's
   `level != 'owner'`.
3. Load the current row.
4. Determine, server-side, whether `name` or `npwp` is present **and
   different** from the stored value — if so, and `confirm != true`,
   reject (`409 confirmation-required`), no partial apply. Never trust
   a client-supplied "this is operational" framing.
5. If `npwp` is changing: validate format, re-check uniqueness
   (`409 npwp-taken` on conflict).
6. Apply the update.
7. If `name`/`npwp` changed (with `confirm: true`) and `status` was
   `'verified'`: set `status = 'pending_verification'`, and in the
   same transaction, auto-unpublish the organization's `published`
   campaigns (INV-organization-09 — stub/test-double until `campaign`
   domain exists).
8. Log to `organization_logs`.
9. Return `200` with the updated `Organization`.

## Validation & error cases

| Case | Response |
|---|---|
| No/invalid bearer token | `401` |
| Caller isn't a representative at all | `403` |
| Caller is `staff` and request includes `name`/`npwp` | `403` |
| `name`/`npwp` changed without `confirm: true` | `409 confirmation-required` |
| `npwp` format invalid | `422` |
| `npwp` duplicate | `409 npwp-taken` |
| No fields present in the request body | `422` |

## Concurrency & correctness notes

- The "did a legal/identity field change" diff must read the
  authoritative stored row within the same transaction as the update.
- Status flip + auto-unpublish (step 7) must be one transaction — no
  committed state where the org is `pending_verification` but a
  campaign under it still reads `published`.

## Test checklist

- [ ] `staff` submitting `description`/`contact`-only → succeeds.
- [ ] `staff` submitting `name` or `npwp` (with or without `confirm`)
      → `403`, no change applied.
- [ ] Non-representative caller → `403`.
- [ ] `owner` submitting `name` without `confirm: true` → `409`, no
      change applied.
- [ ] `owner` submitting `name` with `confirm: true` on a `verified`
      org → status flips, campaigns auto-unpublish (test double).
- [ ] NPWP uniqueness re-checked on edit.
- [ ] Empty request body → `422`.

## References

- `docs/spec/organization/invariants.md` — INV-organization-01, 08
  (resolved 2026-08-20), 09
- `docs/spec/organization/threat-model.md` — "Organization edit"
  section
- `docs/spec/organization/tasks.md` — Task 04
- `api/openapi/organization.yaml` — `PATCH
  /organizations/{organizationId}`
