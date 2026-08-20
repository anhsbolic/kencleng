# Domain Tasks — donation

> File: `docs/spec/donation/tasks.md`
> Status: draft — authored directly against `api/openapi/donation.yaml` 2026-08-20
> Last updated: 2026-08-20

| # | Task | Endpoint / surface | Depends on | Related invariants |
|---|---|---|---|---|
| 01 | Submit donation & async settlement | `POST /campaigns/{id}/donations` + internal settlement process | campaign domain (Task 01, 02, 07, 09) | INV-donation-01–11 |
| 02 | Donation status check | `GET /donations/{id}/status` | 01 | INV-donation-05 |
| 03 | Public donor list | `GET /campaigns/{id}/donations` | 01 | INV-donation-06 |
| 04 | Account donation history | `GET /account/donations` | 01 | — |
| 05 | Guest donation claim | `GET /account/donations/claimable`, `POST /account/donations/{id}/claim` | 01 | INV-donation-12, 13 |
| 06 | Guest-email reveal (new) | `GET /donations/{id}/guest-email` | 01 | INV-donation-14, 15 |

## Task 01 — Submit donation & async settlement

**What**: `POST /campaigns/{campaignId}/donations` — public, works
guest or authenticated. Creates `donations` (`status = 'pending'`),
then the internal (non-HTTP) settlement process resolves it to
`success`/`failed` after a simulated 2–5s delay (5% failure rate). On
success: atomic `collected_amount` increment + campaign-closure check
(INV-donation-08, references INV-campaign-13).

**KPI / metrics**:
- 0 acceptances below the Rp 5.000 minimum.
- 0 acceptances against a non-`published` campaign.
- Authenticated submissions always use session identity
  (INV-donation-07) — 0 cases where body guest fields leak through.
- Idempotency: retried submission with the same `Idempotency-Key`
  never creates a second row.
- Settlement idempotency: invoking the internal transition twice for
  the same donation never double-increments `collected_amount`.
- **Critical**: the internal settlement process is verified
  unreachable via any HTTP route (see `threat-model.md`'s flagged
  implementation-time requirement).
- Concurrent donations to the same campaign never lose an increment
  (row-level locking test).
- A donation crossing `max_amount` triggers closure in the same
  transaction as the increment (cross-domain integration test with
  `campaign` domain).

## Task 02 — Donation status check

**What**: `GET /donations/{donationId}/status?token=...` — token-based,
no login required, doesn't expire.

**KPI / metrics**:
- `status_token` never appears in this endpoint's response (only in
  the original submission response).
- Wrong token or nonexistent donation both produce the same `401`,
  indistinguishable (no signal leak between the two failure modes —
  explicit test per `threat-model.md`'s note).

## Task 03 — Public donor list

**What**: `GET /campaigns/{campaignId}/donations` — public,
`status = 'success'` only, `display_name` server-computed.

**KPI / metrics**:
- `guest_email` never present in any response from this endpoint,
  under any circumstance.
- `is_anonymous = true` → `display_name: null`, regardless of
  `guest_name`.
- Empty `guest_name` (provided but blank) → generic label substituted
  server-side, not left blank or null (unless also anonymous).

## Task 04 — Account donation history

**What**: `GET /account/donations` — a registered user's full
donation history, including claimed former-guest donations.

**KPI / metrics**:
- Includes both directly-registered donations and claimed donations
  (`claimed_at` populated) in one unified list.
- 0 cross-user leakage.

## Task 05 — Guest donation claim

**What**: `GET /account/donations/claimable` (candidate list, matched
by verified email), `POST /account/donations/{donationId}/claim`
(one-at-a-time confirmation).

**KPI / metrics**:
- Unverified-email caller → `403` on both endpoints.
- Claimable list only shows donations with a matching
  `guest_email_hash` and `donor_user_id IS NULL`.
- Claim on a mismatched-email donation → `403`.
- Already-claimed donation → `409`.
- Two users sharing an email racing to claim the same donation:
  exactly one succeeds.
- Post-claim: `guest_name`/`guest_email` snapshot unchanged, past
  public display unaffected (INV-donation-13).

## Task 06 — Guest-email reveal (new, decided 2026-08-20)

**What**: `GET /donations/{donationId}/guest-email` — Admin, or a
Kurator assigned (current or historical) to the donation's campaign.
Not yet in `api/openapi/donation.yaml` — needs adding at
implementation time (INV-donation-14).

**KPI / metrics**:
- Non-Admin, non-assigned-Kurator caller → `403`.
- An unrelated Kurator (not assigned to this donation's campaign) →
  `403`.
- Every successful reveal produces exactly one `donation_logs` row, in
  the same transaction as the reveal itself (not two independent
  operations where the log could fail silently while the reveal
  succeeds).

## References

- Related domain invariants: `docs/spec/donation/invariants.md`
- Related threat model: `docs/spec/donation/threat-model.md`
- **Actual API (ground truth)**: `api/openapi/donation.yaml` (Task 06's
  endpoint not yet present — added by the 2026-08-20 decision)
- Feature specs (one per task, prefixed with the task `#` above):
  `docs/spec/donation/features/`
