# Feature Spec — 05: Guest Donation Claim

> File: `docs/spec/donation/features/05-guest-donation-claim.md`
> Domain: `donation`
> Task: 05 (see `docs/spec/donation/tasks.md`)
> Status: draft — authored against `api/openapi/donation.yaml` 2026-08-20
> Last updated: 2026-08-20

## Summary

Two endpoints: a candidate list matched by the caller's own verified
email, and a one-at-a-time explicit claim action. No auto-claim, no
bulk claim. Permanent (not one-time) — guest donations made even
*after* the user registers can still show up and be claimed later.

## Endpoints

`GET /account/donations/claimable`, `POST
/account/donations/{donationId}/claim` (confirmed,
`api/openapi/donation.yaml`)

## Auth

Both: `bearerAuth` required, caller's email must be verified.

## Request

### Claimable list
`CursorParam`, `LimitParam`.

### Claim
Path: `donationId`. No body.

## Behavior

### Claimable list
1. Reject (`403`) if caller's email isn't verified.
2. Query `donations WHERE guest_email_hash = HMAC(current_user_email)
   AND donor_user_id IS NULL`.
3. Return `ClaimableDonationListResponse` — each item includes the
   **decrypted** `guest_email` (the one place it's returned in full,
   since it's being matched against the requester's own verified
   email — not a third-party reveal), plus `campaign_id`,
   `campaign_title`, `guest_name`, `amount`, `created_at`.

### Claim
1. Reject (`403`) if caller's email isn't verified.
2. Load the donation — reject (`403`) if its `guest_email_hash`
   doesn't match the caller's verified email.
3. `UPDATE donations SET donor_user_id = current_user_id, claimed_at =
   now() WHERE id = :id AND donor_user_id IS NULL` — reject (`409`) if
   0 rows affected (already claimed, by this user or, in the rare
   shared-email edge case, another).
4. Return `200` with the `MyDonation` object.

## Validation & error cases

| Case | Response |
|---|---|
| No/invalid bearer token | `401` |
| Caller's email not verified | `403` (both endpoints) |
| Claim target's `guest_email` doesn't match caller | `403` |
| Already claimed | `409` |

## Concurrency & correctness notes

- The claim's `WHERE donor_user_id IS NULL` guard is the entire
  correctness mechanism for the rare shared-email double-claim race
  (INV-donation-12) — no additional locking needed, Postgres row-level
  semantics on the `UPDATE` handle it.
- Claiming does **not** modify `guest_name`/`guest_email` on the row
  (INV-donation-13) — only `donor_user_id`/`claimed_at` change.

## Test checklist

- [ ] Unverified-email caller → `403` on both endpoints.
- [ ] Claimable list only shows donations matching the caller's own
      email hash, with `donor_user_id IS NULL`.
- [ ] Claimable list never shows a guest donation that had no
      `guest_email` at all.
- [ ] Claim on a mismatched-email donation → `403`.
- [ ] Claim on an already-claimed donation → `409`.
- [ ] Two users with (hypothetically) the same verified email racing
      to claim the same donation: exactly one succeeds.
- [ ] Post-claim: `guest_name`/`guest_email` unchanged on the row.
- [ ] A guest donation made *after* the donor already has an account
      still appears in the claimable list (permanent page, not
      one-time — confirm no time-window restriction exists).

## References

- `docs/spec/donation/invariants.md` — INV-donation-12, 13
- `docs/spec/donation/threat-model.md` — "Account donation history &
  claimable list" and "Claim" sections
- `docs/spec/donation/tasks.md` — Task 05
- `api/openapi/donation.yaml` — `ClaimableDonation`,
  `ClaimableDonationListResponse`, `MyDonation` schemas
