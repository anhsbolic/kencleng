# Feature Spec — 02: Donation Status Check

> File: `docs/spec/donation/features/02-donation-status-check.md`
> Domain: `donation`
> Task: 02 (see `docs/spec/donation/tasks.md`)
> Status: draft — authored against `api/openapi/donation.yaml` 2026-08-20
> Last updated: 2026-08-20

## Summary

`GET /donations/{donationId}/status?token=...` — token-based
settlement-status lookup, no login required, doesn't expire
(read-only, non-destructive).

## Endpoint

`GET /donations/{donationId}/status` (confirmed,
`api/openapi/donation.yaml`)

## Auth

None — token-guarded, not identity-guarded. `token` (query param) must
match the donation's `status_token`.

## Request

Path: `donationId`. Query: `token` (required).

## Behavior

1. Load the donation by `id`.
2. Compare `token` against the stored `status_token` — if the
   donation doesn't exist, or the token doesn't match, return `401`
   (same response either way — see "Validation" below).
3. Return `200` with the `Donation` object — **without**
   `status_token` (it's the credential to reach this endpoint, not
   data to redisplay).

## Validation & error cases

| Case | Response |
|---|---|
| Donation doesn't exist | `401` |
| Token doesn't match | `401` — **identical** response to the nonexistent-donation case, no distinguishing signal |
| `token` query param missing | `401` |

## Concurrency & correctness notes

None specific — plain, non-destructive read.

## Test checklist

- [ ] Correct token → `200` with `Donation`, no `status_token` field
      in the response.
- [ ] Nonexistent donation id → `401`.
- [ ] Wrong token for an existing donation → `401`.
- [ ] The `401` response body/shape is **identical** for both failure
      cases above — assert this explicitly, not just "both return
      401."
- [ ] Token never expires — a status check performed long after
      creation (test with a manually backdated `created_at`) still
      succeeds.

## References

- `docs/spec/donation/invariants.md` — INV-donation-05
- `docs/spec/donation/threat-model.md` — "Token-based status check"
  section
- `docs/spec/donation/tasks.md` — Task 02
- `api/openapi/donation.yaml` — `GET /donations/{id}/status`
