# Feature Spec — 01: Submit Donation & Async Settlement

> File: `docs/spec/donation/features/01-submit-donation-settlement.md`
> Domain: `donation`
> Task: 01 (see `docs/spec/donation/tasks.md`)
> Status: draft — authored against `api/openapi/donation.yaml` 2026-08-20
> Last updated: 2026-08-20

## Summary

`POST /campaigns/{campaignId}/donations` — public, works with or
without authentication. Creates a `donations` row (`status =
'pending'`), then an **internal, non-HTTP** settlement process
resolves it after a simulated 2–5s delay (5% failure rate). On
success: atomic `collected_amount` increment, and a check for
campaign closure.

## `[CRITICAL — implementation requirement]`

The settlement process **must never** be reachable via any registered
HTTP route, internal or external. This is the one place in the whole
`donation` domain where a routing mistake creates a severe
vulnerability (anyone could forge a "payment succeeded" callback and
inflate `collected_amount` arbitrarily). Verify explicitly during code
review — don't rely on "we just didn't add a route for it," add a
test that asserts no such route exists if the router setup allows it.

## Endpoint

`POST /campaigns/{campaignId}/donations` (confirmed,
`api/openapi/donation.yaml`) + internal settlement process (no
endpoint)

## Auth

None required (`security: []`). Optional `Authorization` header — if
present and valid, the donation is recorded as the authenticated
user's own; if absent, as a guest donation.

## Request

`IdempotencyKeyHeader` required. Body — `SubmitDonationRequest`:

| Field | Type | Required | Notes |
|---|---|---|---|
| `amount` | decimal string | Yes | `≥ 5000` |
| `payment_method` | enum | Yes | `transfer`/`debit`/`gopay`/`shopeepay`/`ovo`/`qris` — recorded only, no validation logic in v1 |
| `guest_name` | string, nullable | No | Ignored if authenticated |
| `guest_email` | string (email), nullable | No | Ignored if authenticated. If omitted for a guest donation, this donation can never be claimed and the donor gets no outcome notification (deliberate, INV-donation-03) |
| `is_anonymous` | boolean | No | Default `false` |
| `event_id` | UUID, nullable | No | Optional Event context |

## Behavior

### Submission
1. Resolve `current_user_id` if `Authorization` is present and valid;
   otherwise proceed as guest.
2. Validate `amount ≥ 5000` (`422` otherwise).
3. Validate `payment_method` is a known enum value (`422` otherwise).
4. Load the campaign — reject (`409 campaign-not-published`) if
   `status != 'published'`.
5. If authenticated: `donor_user_id = current_user_id`,
   `guest_name`/`guest_email` both `null` on the row regardless of
   what the body contained (INV-donation-07).
6. If guest: `guest_name` stored as-is (nullable). `guest_email`, if
   provided, encrypted (AES-GCM) + hashed (HMAC) before storage
   (INV-donation-04).
7. Generate `status_token` (random, long, unique — retry generation on
   the rare collision against `ux_donations_status_token`).
8. Insert `donations` (`status = 'pending'`).
9. If `guest_email` was provided: best-effort email containing the
   `status_token`/status-check link.
10. Return `201` with the `Donation` object — this is the **only**
    response that ever includes `status_token`.
11. Enqueue the async settlement job (2–5s simulated delay).

### Settlement (internal, no HTTP endpoint)
1. After the delay: 5% chance → `status = 'failed'`. 95% chance →
   proceed to step 2.
2. `UPDATE donations SET status = 'success' WHERE id = :id AND status
   = 'pending'` — guard makes this idempotent (INV-donation-09).
3. If the update affected a row (i.e. wasn't already resolved): within
   the same transaction, `UPDATE campaigns SET collected_amount =
   collected_amount + :amount WHERE id = :campaign_id AND status =
   'published' RETURNING collected_amount` (INV-donation-08).
4. If the returned `collected_amount ≥ max_amount` (when set): trigger
   campaign closure (`closed_reason = 'max_amount_reached'`) — see
   `docs/spec/campaign/features/09-closure.md`.
5. On the 5% failure path: `UPDATE donations SET status = 'failed'
   WHERE id = :id AND status = 'pending'` — no change to
   `collected_amount`.

## Validation & error cases

| Case | Response |
|---|---|
| `amount < 5000` | `422` |
| Invalid `payment_method` | `422` |
| Campaign not `published` | `409 campaign-not-published` |
| Retried request, same `Idempotency-Key` | Original response returned, no new row |

## Concurrency & correctness notes

- Step 3's atomic conditional `UPDATE` is the entire correctness
  mechanism for concurrent donations to the same campaign — Postgres
  row-level locking serializes them. No application-level locking
  needed.
- Settlement idempotency (`WHERE status = 'pending'`) prevents
  double-increment if the job somehow runs twice for the same
  donation.
- `Idempotency-Key` prevents duplicate `donations` rows from a
  double-submit at the HTTP layer — a separate concern from
  settlement idempotency (one guards row creation, the other guards
  the pending→success/failed transition).

## Test checklist

- [ ] `amount < 5000` → `422`.
- [ ] Invalid `payment_method` → `422`.
- [ ] Non-`published` campaign → `409 campaign-not-published`.
- [ ] Authenticated submission with guest fields populated in the body
      → `donor_user_id` set, `guest_name`/`guest_email` both `null` on
      the stored row.
- [ ] Guest submission with no `guest_email` → never appears in any
      claimable list later.
- [ ] Retried submission, same `Idempotency-Key` → exactly one row
      created.
- [ ] `status_token` present only in the `201` response, never again.
- [ ] Settlement invoked twice for the same donation → `collected_amount`
      incremented exactly once, second invocation is a no-op.
- [ ] Concurrent donations to the same campaign: no lost increments
      (load test with N simultaneous submissions, assert final
      `collected_amount` = sum of all successful amounts).
- [ ] A donation crossing `max_amount` triggers closure in the same
      transaction as its own increment.
- [ ] **Route audit**: confirm no HTTP route exists for the settlement
      transition.

## References

- `docs/spec/donation/invariants.md` — INV-donation-01 through 11
- `docs/spec/donation/threat-model.md` — "Submit donation" and
  "Internal settlement process" sections
- `docs/spec/donation/tasks.md` — Task 01
- `docs/spec/campaign/invariants.md` — INV-campaign-13 (referenced)
- `api/openapi/donation.yaml` — `POST /campaigns/{id}/donations`
