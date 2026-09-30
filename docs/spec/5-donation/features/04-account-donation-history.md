# Feature Spec — 04: Account Donation History

> File: `docs/spec/5-donation/features/04-account-donation-history.md`
> Domain: `donation`
> Task: 04 (see `docs/spec/5-donation/tasks.md`)
> Status: deferred — historical evidence only; excluded from Slice 2
> Last updated: 2026-09-30

> **Slice 2 disposition: DEFER.** Account donation history is not required for the approved guest trust loop, and Account is not a prerequisite. The behavior below is historical evidence, not active acceptance or authorization to create an Account history surface. Reconcile afresh only if a later Product/MVP scope requires it.

## Summary

`GET /account/donations` — a registered user's own full donation
history, including donations originally made as guest and later
claimed. `# INFERRED` in the actual spec (not explicitly named in the
phase doc, but implied by the Donatur persona having a personal
donation record).

## Endpoint

`GET /account/donations` (confirmed, `api/openapi/donation.yaml`)

## Auth

`bearerAuth` required.

## Request

`CursorParam`, `LimitParam`.

## Behavior

1. Resolve `current_user_id`.
2. Query `donations WHERE donor_user_id = current_user_id` — this
   naturally includes both donations made while logged in and
   donations claimed later (claiming sets `donor_user_id`, so claimed
   rows appear here without any special-case query logic).
3. Return `MyDonationListResponse` — each item includes
   `campaign_title`, `amount`, `payment_method`, `is_anonymous`,
   `status`, `claimed_at` (nullable, set only for claimed rows),
   `created_at`.

## Validation & error cases

| Case | Response |
|---|---|
| No/invalid bearer token | `401` |
| `limit` out of range | `422` |

## Concurrency & correctness notes

None specific — plain read.

## Test checklist

- [ ] Returns only the caller's own donations (cross-user leakage
      test with 2+ users).
- [ ] Includes both directly-registered and claimed-former-guest
      donations in one unified, correctly-ordered list.
- [ ] `claimed_at` is `null` for donations made while already
      registered, populated for claimed ones.
- [ ] `pending`/`failed` donations appear too (this is the user's
      *full* history, unlike the public list which is
      `success`-only) — confirm this distinction explicitly.

## References

- `docs/spec/5-donation/invariants.md` — deferred domain context
- `docs/spec/5-donation/tasks.md` — Task 04 (deferred)
- `api/openapi/donation.yaml` — `MyDonation`,
  `MyDonationListResponse` schemas
