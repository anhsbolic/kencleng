# Domain Invariant — donation

> File: `docs/spec/donation/invariants.md`
> Status: draft — authored directly against `api/openapi/donation.yaml` 2026-08-20
> Last updated: 2026-08-20

## Domain summary

`donation` owns donation submission (guest or registered), simulated
async payment settlement, the public donor list, token-based
donation-status lookup, a registered user's own donation history, and
the guest-donation claim flow. Covers `donations`, `donation_logs`.
One invariant is **sent** to `campaign` (the `max_amount`-reached
closure trigger, referencing INV-campaign-13 — declared there, this
domain implements the trigger call).

## Decision recorded 2026-08-20 — guest-email reveal endpoint

`donation_logs` (per `kencleng-erd.md`) explicitly anticipates
"reveal of `guest_email` (PII) by Admin/Kurator" as a loggable action,
but no such endpoint existed in `api/openapi/donation.yaml` — a real
gap between schema intent and API surface. **Decided: add one**,
mirroring `organization`'s NPWP-reveal pattern in spirit, but
structurally different since (unlike `npwp`) `guest_email` is **never**
returned in any existing response — so this new endpoint both reveals
*and* logs in one call, rather than logging a reveal of data the
client already has. See INV-donation-14.

## Invariants

### INV-donation-01: Minimum donation amount

- **Statement**: `amount ≥ 5000` (Rp 5.000), enforced both at the
  application layer and via DB `CHECK` (`kencleng-erd.md`:
  `CHECK (amount >= 5000)`).
- **Holds after operations**: `POST /campaigns/{campaignId}/donations`.
- **Verification**: Test — submit `amount < 5000` → `422`.

### INV-donation-02: Donation acceptance requires the campaign to be `published` at submit time

- **Statement**: A donation is only accepted while
  `Campaign.status = 'published'`, checked atomically at submission
  time (not at page-load time) — a campaign closing between page load
  and submit results in rejection, not silent acceptance.
- **Holds after operations**: `POST /campaigns/{campaignId}/donations`.
- **Verification**: Confirmed — `409 campaign-not-published`. Test:
  close a campaign, then submit against its (stale, already-loaded)
  id → `409`.

### INV-donation-03: Guest fields are independently optional; omitting `guest_email` forfeits claim eligibility and outcome notification

- **Statement**: `guest_name` and `guest_email` are each independently
  optional for a guest donation — a guest may provide neither, either,
  or both. If `guest_email` is omitted, this specific donation can
  **never** be claimed later (Guest Donation Claim flow) and the donor
  receives no campaign-outcome notification. This is a **deliberate
  design decision**, not a defect — restated explicitly here since
  it's easy to mistake for a bug during implementation.
- **Holds after operations**: `POST /campaigns/{campaignId}/donations`
  (guest path).
- **Verification**: Test — guest donation with no `guest_email` never
  appears in any `/account/donations/claimable` result for any user,
  regardless of email.

### INV-donation-04: `guest_email` is encrypted at rest

- **Statement**: `guest_email` is stored as `BYTEA` ciphertext
  (AES-GCM) with a separate `guest_email_hash` (HMAC-SHA256) for
  lookup — same PII pattern as `organizations.npwp`. Never stored or
  logged in plaintext outside the request that creates or reveals it.
- **Holds after operations**: donation submission (write), claim-list
  lookup (`WHERE guest_email_hash = ?`, read via hash only, never
  decrypting to compare), the new guest-email reveal endpoint
  (INV-donation-14, the one place plaintext is intentionally
  returned).
- **Verification**: Test — inspect the stored row directly, assert
  `guest_email` is not plaintext-readable without the decryption key.

### INV-donation-05: `status_token` is unique, random, non-expiring, and single-response-only

- **Statement**: `status_token` is cryptographically random, long
  enough to resist brute-force guessing, globally unique
  (`ux_donations_status_token`), and does **not** expire (read-only,
  non-destructive lookup — unlike password-reset tokens). It is
  returned **only** in the immediate submission response (`201`) —
  the status-check endpoint (`GET /donations/{id}/status`) never
  returns it again, since it's the credential used to reach that
  endpoint, not data to display.
- **Holds after operations**: submission (generation), status check
  (never re-exposed).
- **Verification**: DB-level uniqueness constraint. Test — the
  `Donation` object returned by `GET /donations/{id}/status` has no
  `status_token` field, even for the donation's own submitter.

### INV-donation-06: Public-facing donor data never exposes `guest_email`; display name and anonymity are server-computed

- **Statement**: The public donor list (`GET
  /campaigns/{campaignId}/donations`) and any other public-facing
  shape never includes `guest_email` — confirmed, this is explicit in
  the `DonationListItem` schema's description. `display_name` is
  computed server-side: the registered donor's name, or `guest_name`
  (substituted with a generic label if empty), or `null` entirely if
  `is_anonymous = true`. The client never receives raw material to
  reconstruct a masked identity — anonymity is enforced by omission at
  the API boundary, not by a client-side flag the frontend is trusted
  to respect.
- **Holds after operations**: `GET
  /campaigns/{campaignId}/donations`.
- **Verification**: Test — a donation with `is_anonymous = true` has
  `display_name: null` in the public list, regardless of whether
  `guest_name` was provided. `guest_email` never appears in this
  response shape under any circumstance, for any donation.

### INV-donation-07: Authenticated submission always uses session identity, never the body's guest fields

- **Statement**: If `POST /campaigns/{campaignId}/donations` is called
  with a valid `Authorization` header, `guest_name`/`guest_email` in
  the request body are **ignored** — the registered donor's own
  `User.name`/session identity is used instead. There is no path for
  an authenticated user to submit "as guest" by supplying different
  guest fields.
- **Holds after operations**: submission, when authenticated.
- **Verification**: Test — authenticated submission with `guest_name`/
  `guest_email` populated in the body still records `donor_user_id =
  current_user_id`, `guest_name`/`guest_email` both `null` on the
  stored row (registered donations don't populate the guest columns
  at all).

### INV-donation-08: `collected_amount` increment is atomic and conditional; triggers campaign closure (sent — `campaign` domain)

- **Statement**: On donation success,
  `UPDATE campaigns SET collected_amount = collected_amount + :amount
  WHERE id = :id AND status = 'published' RETURNING collected_amount`
  — a single atomic, conditional update. Postgres row-level locking on
  this statement serializes concurrent donations to the same campaign;
  no separate application-level locking is needed. After the
  increment, if the returned `collected_amount ≥ max_amount` (when
  set), trigger campaign closure — see **INV-campaign-13** in
  `docs/spec/campaign/invariants.md` (declared there, this domain
  implements the trigger call within the same transaction as the
  donation's own `status = success` update).
- **Holds after operations**: the donation-success transition (whether
  reached via simulated async settlement or, in a real gateway,
  a callback).
- **Verification**: Test — concurrent donations to the same campaign
  never lose an increment (no lost-update race). A donation that pushes
  `collected_amount` past `max_amount` triggers closure in the same
  transaction as the increment — no window where `collected_amount ≥
  max_amount` but `status` is still `published`.

### INV-donation-09: Async settlement is idempotent

- **Statement**: The simulated payment "callback" (2–5s delayed,
  internal-only — **never** exposed as a public HTTP endpoint,
  distinct from every other endpoint in this domain, all of which are
  public or bearer-authenticated) transitions `pending → success` or
  `pending → failed` guarded by `WHERE status = 'pending'`. If invoked
  twice for the same donation (e.g. a bug), the second invocation is a
  no-op — never double-increments `collected_amount`.
- **Holds after operations**: the internal settlement transition.
- **Verification**: Test — invoke the settlement transition twice for
  the same donation id; assert `collected_amount` is only incremented
  once, second invocation is a no-op with no error.

### INV-donation-10: Donations are accepted in full — no clamping against `max_amount`

- **Statement**: A donation is accepted at its full submitted `amount`
  as long as the campaign is `published` at submit time — there is
  **no** partial-accept or clamping logic against `max_amount`. A
  campaign's final `collected_amount` can slightly exceed
  `max_amount`. This is a deliberate simplicity trade-off, not a
  race-condition gap to "fix."
- **Holds after operations**: submission, settlement.
- **Verification**: Test — a donation that would push
  `collected_amount` past `max_amount` is still accepted at its full
  amount (not truncated), and correctly triggers closure per
  INV-donation-08.

### INV-donation-11: Submission is idempotent via `Idempotency-Key`

- **Statement**: `POST /campaigns/{campaignId}/donations` requires an
  `Idempotency-Key` header (confirmed, `IdempotencyKeyHeader`). A
  retried request with the same key returns the original response
  without creating a second `Donation` row.
- **Holds after operations**: submission.
- **Verification**: Test — submit twice with the same
  `Idempotency-Key`; assert exactly one `donations` row is created,
  second response matches the first.

### INV-donation-12: Guest donation claim is one-at-a-time, verified-email-only, and race-safe

- **Statement**: A donation is claimable by a registered user only if
  `guest_email_hash` matches that user's own **verified**
  `primary_email`, and only while `donor_user_id IS NULL`. Claiming
  requires explicit per-donation confirmation — no bulk or automatic
  claim. The claim update is guarded by `WHERE donor_user_id IS NULL`,
  preventing the same donation from being claimed twice even in the
  rare case where two different users share an email (one claims,
  `donor_user_id` is no longer `NULL`, the second claim attempt
  cleanly fails rather than overwriting).
- **Holds after operations**: `GET /account/donations/claimable`,
  `POST /account/donations/{donationId}/claim`.
- **Verification**: Confirmed — `403` if the caller's email isn't
  verified, `403` if the target donation's `guest_email` doesn't match
  the caller, `409` if already claimed. Test: two users with
  (hypothetically) the same verified email both attempt to claim the
  same donation near-simultaneously — exactly one succeeds, the other
  gets `409`, no double-claim.

### INV-donation-13: A claimed donation's original snapshot is preserved; past public display is not retroactively changed

- **Statement**: Claiming a donation sets `donor_user_id`/`claimed_at`
  but does **not** alter the stored `guest_name`/`guest_email`
  snapshot, and does not retroactively change how the donation
  appeared in any historical public display (e.g. a donor list
  rendered before the claim).
- **Holds after operations**: claim.
- **Verification**: Test — claim a donation, assert `guest_name` on
  the row is unchanged; assert the public donor list's `display_name`
  logic for that row is unaffected by the claim (still computed the
  same way it would've been pre-claim, since `is_anonymous` and
  `guest_name` didn't change).

### INV-donation-14: Guest-email reveal is Admin/scoped-Kurator-only and always logged (new — decided 2026-08-20)

- **Statement**: `guest_email` (decrypted) may be revealed via a new
  endpoint, `GET /donations/{donationId}/guest-email`, to: Admin, or a
  Kurator who is or was assigned to curate the campaign this donation
  belongs to (not any Kurator — same assignment-scoping pattern as
  `organization`'s legal-document access). Every successful reveal
  produces exactly one `donation_logs` row (`action_type =
  'guest_email_revealed'`, `actor_user_id`, `donation_id`) in the same
  request — unlike `organization`'s NPWP pattern (where the value was
  already in an existing response and reveal-logging was a separate,
  data-free call), this endpoint both returns the value *and* logs the
  reveal, since `guest_email` is otherwise never returned anywhere.
- **Holds after operations**: `GET /donations/{donationId}/guest-email`
  (new endpoint, not yet in `api/openapi/donation.yaml` — needs
  adding at implementation time).
- **Verification**: Test — non-Admin, non-assigned-Kurator caller →
  `403`. Admin succeeds for any donation. A Kurator assigned (current
  or historical) to the donation's campaign succeeds; an unrelated
  Kurator does not. Every successful call produces exactly one
  `donation_logs` row.

### INV-donation-15: `donation_logs` is append-only

- **Statement**: No row in `donation_logs`, once inserted, may ever be
  updated or deleted.
- **Holds after operations**: guest-email reveal (INV-donation-14) —
  currently the only anticipated write path per the ERD's comment,
  though the table's shape doesn't preclude future action types.
- **Verification**: DB-level — `REVOKE UPDATE, DELETE ON donation_logs
  FROM kencleng_app`, same pattern as every other `_logs` table in
  this project.

## State machines

### `donations.status`

```
(none) -> pending -> success / failed
```

One-way, terminal at `success`/`failed`. No retry-on-same-row — a
failed payment requires a **new** donation submission
(`kencleng-phase2-detail.md` Fitur 1, explicit: "the donor submits a
new donation, not a retry on the same donation").

### `donations.donor_user_id` / `claimed_at`

```
NULL -> <user_id> (claim, guest donations with guest_email only)
```

One-way, guarded by `WHERE donor_user_id IS NULL` (INV-donation-12).
Registered-donor donations have `donor_user_id` set at creation and
never go through this transition at all.

## Reference for `campaign` domain

`docs/spec/campaign/invariants.md`'s INV-campaign-13 already documents
the `max_amount`-reached closure trigger as received *from*
`donation` — this entry (INV-donation-08) is the implementing side;
no further cross-reference needed beyond what's already written there.

## References

- Related ERD: `docs/project/kencleng-erd.md` §4 (`donations`,
  `donation_logs`)
- Related business process: `docs/project/kencleng-phase2-detail.md`
  Fitur 1, 4
- Related actors/rules: `docs/project/kencleng-actors-entities.md`
  Business Rule 5
- Related invariants: `docs/spec/campaign/invariants.md` —
  INV-campaign-13 (receives INV-donation-08's trigger)
- **Actual API (ground truth)**: `api/openapi/donation.yaml` — note
  INV-donation-14's endpoint is **not yet present**, added by this
  decision
